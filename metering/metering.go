// Package metering is the ONE way every Hanzo product meters usage to
// commerce — the single billing source of truth — so that every product
// (not only the LLM/cloud path) can be paid for.
//
// It provides two operations, matching the proven cloud/gateway path:
//
//   - Authorize: a pre-request balance gate. Fail-closed by default — if the
//     balance cannot be determined the request is denied, exactly like the
//     gateway's prepaid-balance gate (gateway/auth_middleware.go). With
//     TierAware enabled it consults the tier-aware effective balance, which
//     folds in the tenant's included plan allotment (e.g. the free-tier
//     daily credit) so included usage is honored before prepaid funds.
//
//   - Record: a post-request usage write. Records a usage event (cost in
//     cents) against commerce, which debits the user's balance ledger.
//
// The HTTP contract is commerce's canonical billing API, mounted under /v1
// (commerce/api/billing/handlers.go):
//
//	GET  {BaseURL}/v1/billing/balance?user={user}&currency={cur}
//	GET  {BaseURL}/v1/billing/tier?user={user}            (tier-aware)
//	POST {BaseURL}/v1/billing/usage
//
// Auth is the commerce service token (admin-scoped S2S), sent as
//
//	Authorization: Bearer {Token}
//
// plus the tenant org as the X-IAM-Org-Id header. The token is a secret and
// MUST be sourced from KMS (never plaintext); this package never reads it from
// disk — the caller supplies it (typically from an env var the operator wires
// from a KMS-backed secret, e.g. COMMERCE_SERVICE_TOKEN).
//
// This package has no dependency on the commerce module, so any product —
// Go service, CLI, or job — can import it without pulling in a server.
package metering

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Canonical commerce billing paths (mounted under /v1). Keep in lockstep with
// commerce/api/billing/handlers.go — a wrong prefix 404s and, fail-closed,
// denies every request.
const (
	pathBalance = "/v1/billing/balance"
	pathTier    = "/v1/billing/tier"
	pathUsage   = "/v1/billing/usage"
)

// Header carrying the tenant org slug for commerce namespace resolution.
// Commerce's S2S auth reads the org from X-IAM-Org-Id (commerce/middleware
// /accesstoken.go); without it commerce falls back to COMMERCE_SERVICE_ORG.
const headerOrg = "X-IAM-Org-Id"

// ErrInsufficientBalance is returned by Authorize when commerce confirms the
// user's available balance is non-positive. It is distinct from a connectivity
// failure so callers can map it to HTTP 402 (vs 503 for "unknown").
var ErrInsufficientBalance = errors.New("metering: insufficient balance")

// HTTPDoer is the minimal HTTP surface the client needs. *http.Client
// satisfies it; tests and instrumented transports can substitute their own.
type HTTPDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Config configures a Client. Only BaseURL is conceptually required; an empty
// BaseURL puts the client in "not configured" mode where Authorize allows and
// Record is a no-op — matching the gateway's behavior when no billing URL is
// set, so a product can adopt metering before its tenant billing is wired.
type Config struct {
	// BaseURL is the commerce service base, e.g.
	// "http://commerce.hanzo.svc.cluster.local:8001". No trailing /v1 — the
	// client appends the canonical billing paths itself.
	BaseURL string

	// Token is the commerce service token (admin-scoped). MUST come from KMS;
	// never hard-code or read from a file. Sent as "Authorization: Bearer".
	Token string

	// Org is the tenant org slug (e.g. "hanzo") sent as X-IAM-Org-Id so
	// commerce resolves the right tenant namespace. Per-request Org on the
	// Usage/AuthInput overrides this default.
	Org string

	// TierAware, when true, makes Authorize consult GET /v1/billing/tier and
	// gate on the effective balance (prepaid + included plan allotment such as
	// the free-tier daily credit) instead of the bare prepaid balance. This is
	// the same effectiveAvailable commerce computes in GetTier.
	TierAware bool

	// FailOpen inverts the default fail-closed posture: when commerce cannot be
	// reached, Authorize allows the request instead of denying it. Leave false
	// for paid products; set true only where availability outranks billing
	// (and accept the revenue leak). Mirrors the gateway, which is fail-closed.
	FailOpen bool

	// Timeout bounds each commerce HTTP call. Default 5s (the gateway's value).
	Timeout time.Duration

	// HTTPClient overrides the underlying HTTP client. When nil a client with
	// Timeout is created.
	HTTPClient HTTPDoer
}

// Client meters usage to commerce. It is safe for concurrent use.
type Client struct {
	baseURL   string
	token     string
	org       string
	tierAware bool
	failOpen  bool
	http      HTTPDoer
}

// New builds a metering Client from cfg. It returns an error only for an
// unparseable BaseURL; an empty BaseURL is valid ("not configured" mode).
func New(cfg Config) (*Client, error) {
	base := strings.TrimRight(strings.TrimSpace(cfg.BaseURL), "/")
	if base != "" {
		if _, err := url.Parse(base); err != nil {
			return nil, fmt.Errorf("metering: invalid BaseURL %q: %w", cfg.BaseURL, err)
		}
	}

	doer := cfg.HTTPClient
	if doer == nil {
		timeout := cfg.Timeout
		if timeout <= 0 {
			timeout = 5 * time.Second
		}
		doer = &http.Client{Timeout: timeout}
	}

	return &Client{
		baseURL:   base,
		token:     strings.TrimSpace(cfg.Token),
		org:       strings.TrimSpace(cfg.Org),
		tierAware: cfg.TierAware,
		failOpen:  cfg.FailOpen,
		http:      doer,
	}, nil
}

// Enabled reports whether a commerce BaseURL is configured. When false,
// Authorize always allows and Record is a no-op.
func (c *Client) Enabled() bool { return c != nil && c.baseURL != "" }

// AuthInput identifies who to authorize. User is the IAM "org/sub" identity
// (e.g. "hanzo/alice"). Org overrides the client default org for this call;
// Currency defaults to "usd".
type AuthInput struct {
	User     string
	Org      string
	Currency string
}

// Authorize is the pre-request balance gate.
//
// Contract — three outcomes, matching the gateway:
//
//	(nil)                       -> allow.
//	(ErrInsufficientBalance)    -> deny: out of funds          (map to HTTP 402).
//	(other error)               -> balance unknown; with the default fail-closed
//	                               posture this denies          (map to HTTP 503).
//	                               With FailOpen it returns nil (allow).
//
// When the client is not configured (no BaseURL) it always allows.
func (c *Client) Authorize(ctx context.Context, in AuthInput) error {
	if !c.Enabled() {
		return nil
	}
	user := strings.TrimSpace(in.User)
	if user == "" {
		// No identity -> cannot bill. Fail-closed denies (anonymous traffic
		// must be handled by a public-path bypass before reaching here).
		if c.failOpen {
			return nil
		}
		return fmt.Errorf("metering: empty user")
	}

	available, err := c.fetchAvailable(ctx, user, c.orgFor(in.Org), currencyOr(in.Currency))
	if err != nil {
		if c.failOpen {
			return nil
		}
		return err // unknown -> deny (fail-closed).
	}
	if available > 0 {
		return nil
	}
	return ErrInsufficientBalance
}

// fetchAvailable returns the spendable balance in cents. With TierAware it uses
// the tier endpoint's effectiveAvailable (prepaid + included allotment);
// otherwise the bare prepaid available from the balance endpoint.
func (c *Client) fetchAvailable(ctx context.Context, user, org, cur string) (int64, error) {
	if c.tierAware {
		q := url.Values{"user": {user}}
		body, err := c.get(ctx, pathTier, q, org)
		if err != nil {
			return 0, err
		}
		var tr tierResponse
		if err := json.Unmarshal(body, &tr); err != nil {
			return 0, fmt.Errorf("metering: decode tier: %w", err)
		}
		return tr.Balance.EffectiveAvailable, nil
	}

	q := url.Values{"user": {user}, "currency": {cur}}
	body, err := c.get(ctx, pathBalance, q, org)
	if err != nil {
		return 0, err
	}
	var br balanceResponse
	if err := json.Unmarshal(body, &br); err != nil {
		return 0, fmt.Errorf("metering: decode balance: %w", err)
	}
	return br.Available, nil
}

// Usage is one usage event to record. User (IAM "org/sub") and AmountCents
// (the cost to debit) are the essentials; the rest is descriptive metadata
// commerce stores on the transaction. Fields mirror commerce's usageRequest
// (commerce/api/billing/usage.go) one-for-one.
type Usage struct {
	User             string `json:"user"`
	Org              string `json:"-"` // routed via X-IAM-Org-Id, not the body.
	Currency         string `json:"currency,omitempty"`
	AmountCents      int64  `json:"amount"`
	Model            string `json:"model,omitempty"`
	Provider         string `json:"provider,omitempty"`
	PromptTokens     int    `json:"promptTokens,omitempty"`
	CompletionTokens int    `json:"completionTokens,omitempty"`
	TotalTokens      int    `json:"totalTokens,omitempty"`
	RequestID        string `json:"requestId,omitempty"`
	Premium          bool   `json:"premium,omitempty"`
	Stream           bool   `json:"stream,omitempty"`
	Status           string `json:"status,omitempty"`
	ClientIP         string `json:"clientIp,omitempty"`
}

// RecordResult is the commerce response to a usage write.
type RecordResult struct {
	TransactionID string `json:"transactionId"`
	User          string `json:"user"`
	Amount        int64  `json:"amount"`
	Currency      string `json:"currency"`
	Type          string `json:"type"`
}

// Record writes a usage event to commerce, debiting the user's balance.
//
// It is a no-op (nil, nil) when the client is not configured or when
// AmountCents <= 0 (commerce treats zero-cost usage as "skipped"). Usage
// recording is deliberately decoupled from gating: the work already happened
// and must be recorded, so balance is NOT re-checked here — exactly as
// commerce's RecordUsage documents.
//
// Provider is the service name doing the metering when no model/provider is
// natural (e.g. "search", "functions"); set it on Usage.Provider.
func (c *Client) Record(ctx context.Context, u Usage) (*RecordResult, error) {
	if !c.Enabled() || u.AmountCents <= 0 {
		return nil, nil
	}
	if strings.TrimSpace(u.User) == "" {
		return nil, fmt.Errorf("metering: Record requires a user")
	}
	if u.Currency == "" {
		u.Currency = "usd"
	}

	payload, err := json.Marshal(u)
	if err != nil {
		return nil, fmt.Errorf("metering: encode usage: %w", err)
	}

	body, err := c.post(ctx, pathUsage, payload, c.orgFor(u.Org))
	if err != nil {
		return nil, err
	}

	var res RecordResult
	if err := json.Unmarshal(body, &res); err != nil {
		// Commerce returned 2xx but an unexpected shape; the debit still
		// happened, so don't treat decode failure as a hard error.
		return nil, nil
	}
	return &res, nil
}

// ---- HTTP plumbing -------------------------------------------------------

func (c *Client) get(ctx context.Context, path string, q url.Values, org string) ([]byte, error) {
	u := c.baseURL + path
	if enc := q.Encode(); enc != "" {
		u += "?" + enc
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}
	return c.do(req, org)
}

func (c *Client) post(ctx context.Context, path string, body []byte, org string) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseURL+path, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	return c.do(req, org)
}

func (c *Client) do(req *http.Request, org string) ([]byte, error) {
	if c.token != "" {
		req.Header.Set("Authorization", "Bearer "+c.token)
	}
	if org != "" {
		req.Header.Set(headerOrg, org)
	}

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("metering: commerce unreachable: %w", err)
	}
	defer resp.Body.Close()

	// Read a bounded body — these are tiny JSON objects.
	body, err := io.ReadAll(io.LimitReader(resp.Body, 1<<16))
	if err != nil {
		return nil, fmt.Errorf("metering: read commerce response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("metering: commerce status %d: %s", resp.StatusCode, strings.TrimSpace(string(body)))
	}
	return body, nil
}

func (c *Client) orgFor(perCall string) string {
	if perCall = strings.TrimSpace(perCall); perCall != "" {
		return perCall
	}
	return c.org
}

func currencyOr(cur string) string {
	if cur = strings.TrimSpace(cur); cur != "" {
		return cur
	}
	return "usd"
}

// balanceResponse mirrors commerce GET /v1/billing/balance. Amounts are cents;
// available = balance - holds.
type balanceResponse struct {
	User      string `json:"user"`
	Currency  string `json:"currency"`
	Balance   int64  `json:"balance"`
	Holds     int64  `json:"holds"`
	Available int64  `json:"available"`
}

// tierResponse mirrors commerce GET /v1/billing/tier. effectiveAvailable folds
// the tenant's included plan allotment (e.g. free-tier daily credit) into the
// prepaid available balance.
type tierResponse struct {
	User    string `json:"user"`
	Balance struct {
		Currency           string `json:"currency"`
		PrepaidAvailable   int64  `json:"prepaidAvailable"`
		DailyRemaining     int64  `json:"dailyRemaining"`
		EffectiveAvailable int64  `json:"effectiveAvailable"`
	} `json:"balance"`
}
