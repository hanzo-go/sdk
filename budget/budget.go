// Package budget is what you may spend, what is left, and what it cost.
//
// Allowance and money never stand in for each other. [Client.Left] answers a
// COUNT — the free calls a plan grants in a period — and [Client.Balance]
// answers a SUM, the prepaid wallet the metered calls draw on. A product that
// shows one where the other belongs tells a funded customer they are out of
// calls, or an out-of-calls customer they are broke.
//
// Every read here is a mirror: the subject is the validated principal and can
// never be named in a request, so no method takes an org.
package budget

import (
	"context"
	"net/url"
	"time"

	"github.com/hanzoai/go-sdk/v8/call"
)

// Client reads the caller's own budget.
type Client struct{ e *call.Endpoint }

// New returns the budget capability over an endpoint. The client at the module
// root builds one; there is no reason to call this directly.
func New(e *call.Endpoint) *Client { return &Client{e: e} }

// Allowance is the plan's free-call ceiling for the period that will stop the
// caller next.
type Allowance struct {
	// Plan is the tier the limit came from.
	Plan string
	// Limit is the calls the plan allows per period. Zero means unbounded.
	Limit int64
	// Used is how many zero-priced calls have been SERVED this period. It stops
	// at Limit rather than climbing past it.
	Used int64
	// Left is Limit-Used, and is absent when the plan is unbounded: there is no
	// remainder of an unbounded count, and reporting zero would read as spent.
	Left *int64
	// Spent reports that the subject is at the limit.
	Spent bool
	// Window is which ceiling these numbers describe, "hour" or "day" — the
	// caller is held to both and this is the one that refused, or the one with
	// least left. Empty where no window bounds the subject at all.
	Window string
	// Resets is when the window starts again, absent when the plan is
	// unbounded, for the same reason Left is.
	Resets *time.Time
}

// Balance is the prepaid wallet the metered calls draw on.
type Balance struct {
	// Available is what can still be spent.
	Available call.Money
	// Reserved is what a reservation has claimed and not yet posted. The wire
	// spells it `holds`; the SDK renames it because `held` is an arm of an
	// answer, and one word for two facts is what the naming rule prevents.
	Reserved call.Money
	// Account is the wallet key the ledger resolved for this caller — the org's
	// shared pool, or a personal account. It is echoed rather than guessed: a
	// guess that disagrees with the server is how money lands in an account the
	// gate never reads.
	Account string
}

// Plan is which apps the caller's org may open, and the tier that decides it.
type Plan struct {
	Tier string
	// Apps is false both where the plan does not grant the app and where the
	// licence authority could not be reached. A read that decides what to show
	// fails to locked, not to an error.
	Apps map[string]bool
}

// Charge is one billed call.
type Charge struct {
	ID     string
	At     time.Time
	Model  string
	Amount call.Money
}

// Filter narrows a spend read. Every field is optional and every field narrows
// within the caller's own org.
type Filter struct {
	Product string
	Since   time.Time
	Until   time.Time
}

// Left answers the free-call allowance for the period.
//
// It READS: asking does not spend, so a page that polls it costs nothing.
func (c *Client) Left(ctx context.Context) (Allowance, error) {
	var wire struct {
		Limit  int64  `json:"limit"`
		Plan   string `json:"plan"`
		Resets int64  `json:"resets"`
		Spent  bool   `json:"spent"`
		Used   int64  `json:"used"`
		Window string `json:"window"`
	}
	if _, err := call.Do(ctx, c.e, "GET", "/v1/allowance", nil, nil, &wire); err != nil {
		return Allowance{}, err
	}
	out := Allowance{Plan: wire.Plan, Limit: wire.Limit, Used: wire.Used, Spent: wire.Spent, Window: wire.Window}
	if wire.Limit > 0 {
		left := max(wire.Limit-wire.Used, 0)
		out.Left = &left
		if wire.Resets > 0 {
			resets := time.Unix(wire.Resets, 0).UTC()
			out.Resets = &resets
		}
	}
	return out, nil
}

// usd is the currency cloud's ledger answers in, ISO 4217. It is stated on
// every Money this package builds rather than assumed by whoever reads one.
const usd = "USD"

// Balance answers the wallet the metered calls draw on.
func (c *Client) Balance(ctx context.Context) (Balance, error) {
	// Cloud declares no response schema for this route, so the shape is
	// modelled here from what the handler writes: {balance, holds, available,
	// account}, whole USD cents.
	var wire struct {
		Holds     int64  `json:"holds"`
		Available int64  `json:"available"`
		Account   string `json:"account"`
	}
	if _, err := call.Do(ctx, c.e, "GET", "/v1/billing/balance", nil, nil, &wire); err != nil {
		return Balance{}, err
	}
	return Balance{
		Available: call.Money{Minor: wire.Available, Currency: usd},
		Reserved:  call.Money{Minor: wire.Holds, Currency: usd},
		Account:   wire.Account,
	}, nil
}

// Plan answers which apps the caller's org may open.
func (c *Client) Plan(ctx context.Context) (Plan, error) {
	var wire struct {
		Tier string          `json:"tier"`
		Apps map[string]bool `json:"apps"`
	}
	if _, err := call.Do(ctx, c.e, "GET", "/v1/entitlement", nil, nil, &wire); err != nil {
		return Plan{}, err
	}
	return Plan{Tier: wire.Tier, Apps: wire.Apps}, nil
}

// Spent answers the billed calls, newest first — one row per charge, not a
// rollup.
func (c *Client) Spent(ctx context.Context, f Filter) (call.Page[Charge], error) {
	query := url.Values{}
	if f.Product != "" {
		query.Set("product", f.Product)
	}
	if !f.Since.IsZero() {
		query.Set("start", f.Since.UTC().Format(time.RFC3339))
	}
	if !f.Until.IsZero() {
		query.Set("end", f.Until.UTC().Format(time.RFC3339))
	}

	// Modelled here for the same reason Balance is: the route declares an
	// address and not a shape.
	var wire struct {
		Count int64 `json:"count"`
		Usage []struct {
			ID        string         `json:"transactionId"`
			Amount    int64          `json:"amount"`
			Metadata  map[string]any `json:"metadata"`
			CreatedAt time.Time      `json:"createdAt"`
		} `json:"usage"`
	}
	if _, err := call.Do(ctx, c.e, "GET", "/v1/billing/usage", query, nil, &wire); err != nil {
		return call.Page[Charge]{}, err
	}

	page := call.Page[Charge]{Total: wire.Count, Items: make([]Charge, 0, len(wire.Usage))}
	for _, row := range wire.Usage {
		model, _ := row.Metadata["model"].(string)
		page.Items = append(page.Items, Charge{
			ID:     row.ID,
			At:     row.CreatedAt,
			Model:  model,
			Amount: call.Money{Minor: row.Amount, Currency: usd},
		})
	}
	return page, nil
}
