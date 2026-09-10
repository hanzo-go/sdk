// Package audit reads the org's own trail, newest first.
//
// It is read-only by construction. The server writes the trail; an SDK that
// offered a write would let a caller forge their own evidence.
package audit

import (
	"context"
	"iter"
	"net/url"
	"strconv"
	"time"

	"github.com/hanzoai/go-sdk/v8/call"
)

// Client reads the caller's own org's trail.
type Client struct{ e *call.Endpoint }

// New returns the audit capability over an endpoint.
func New(e *call.Endpoint) *Client { return &Client{e: e} }

// Event is one row of the trail.
//
// Six wire fields are renamed and nothing else moves: sub is token jargon where
// the reader wants who did it; time becomes at, one word for an instant, shared
// with graph.Fact.At; resourceId sits beside Resource and the compound says
// nothing more; requestId becomes Request, the same word as call.Answer.Request,
// which is what makes the join work; and sourceIp and userAgent lose a qualifier
// they are the only candidate for.
type Event struct {
	// Seq is the row's position in the chain, 0-based and gapless.
	Seq int64
	Org string
	// Home is present only on a cross-org action, and marks it as an
	// impersonation: the org the actor came FROM, where Org is the one they
	// acted IN.
	Home string
	// Actor is the acting IAM subject. Empty for a machine principal.
	Actor  string
	Email  string
	Action string
	// Resource is the KIND of thing acted upon; ID is the instance, absent
	// where the kind alone identifies it.
	Resource string
	ID       string
	Method   string
	Path     string
	// Result is "success", "deny" or "error".
	Result string
	// Status is the HTTP status the caller received.
	Status int
	// Reason is a short explanation for a deny or an error.
	Reason string
	// Request is the x-request-id, the same value call.Answer.Request carries.
	Request string
	IP      string
	Agent   string
	At      time.Time
}

// Filter narrows a read. Every field is optional and every field narrows within
// the caller's own org.
type Filter struct {
	Actor    string
	Action   string
	Resource string
	ID       string
	Result   string
	Since    time.Time
	Until    time.Time
	// Size is rows per page, 100 where unset.
	Size int
	// Page is 1-based.
	Page int
	// Request narrows to the rows one call produced — the value
	// call.Answer.Request carries.
	//
	// The route accepts no such filter, though every row carries the id, so
	// this one narrows CLIENT-SIDE: the request goes out under the rest of the
	// filter and the rows are matched here. It is correct and it is slow, so
	// pair it with Action or Since. It stops being a scan the day GET /v1/audit
	// accepts requestId.
	Request string
}

// pageSize is the trail's own default, restated so paging can count.
const pageSize = 100

// List answers one page of the trail.
//
// Total is the server's own count for the filter the server applied. Under a
// Request filter, which the route does not accept, Items are this page's
// matches and Total still counts what the server was asked: a count the SDK
// substituted would be a number nobody made.
func (c *Client) List(ctx context.Context, f Filter) (call.Page[Event], error) {
	var wire struct {
		Data []struct {
			Seq       int64     `json:"seq"`
			Org       string    `json:"org"`
			Home      string    `json:"home"`
			Sub       string    `json:"sub"`
			Email     string    `json:"email"`
			Action    string    `json:"action"`
			Resource  string    `json:"resource"`
			ResourceI string    `json:"resourceId"`
			Method    string    `json:"method"`
			Path      string    `json:"path"`
			Result    string    `json:"result"`
			Status    int       `json:"status"`
			Reason    string    `json:"reason"`
			RequestID string    `json:"requestId"`
			SourceIP  string    `json:"sourceIp"`
			UserAgent string    `json:"userAgent"`
			Time      time.Time `json:"time"`
		} `json:"data"`
		Total int64 `json:"total"`
	}
	if _, err := call.Do(ctx, c.e, "GET", "/v1/audit", f.query(), nil, &wire); err != nil {
		return call.Page[Event]{}, err
	}

	page := call.Page[Event]{Total: wire.Total, Items: make([]Event, 0, len(wire.Data))}
	for _, row := range wire.Data {
		if f.Request != "" && row.RequestID != f.Request {
			continue
		}
		page.Items = append(page.Items, Event{
			Seq: row.Seq, Org: row.Org, Home: row.Home, Actor: row.Sub, Email: row.Email,
			Action: row.Action, Resource: row.Resource, ID: row.ResourceI,
			Method: row.Method, Path: row.Path, Result: row.Result, Status: row.Status,
			Reason: row.Reason, Request: row.RequestID, IP: row.SourceIP,
			Agent: row.UserAgent, At: row.Time,
		})
	}
	return page, nil
}

// All walks every page of the trail under one filter.
//
// It asks for page 1 at the caller's size and stops when a page comes back
// empty or the running count reaches the total. A read that fails yields its
// error once and ends.
//
// The walk asks for the pages UNFILTERED by Request and matches here, so what
// it counts against the total is what the server sent. Counting the survivors
// of a narrowing the server never made would end the walk early or not at all.
func (c *Client) All(ctx context.Context, f Filter) iter.Seq2[Event, error] {
	return func(yield func(Event, error) bool) {
		want := f.Request
		f.Request = ""
		if f.Size <= 0 {
			f.Size = pageSize
		}
		if f.Page <= 0 {
			f.Page = 1
		}
		var seen int64
		for {
			page, err := c.List(ctx, f)
			if err != nil {
				yield(Event{}, err)
				return
			}
			if len(page.Items) == 0 {
				return
			}
			for _, event := range page.Items {
				if want != "" && event.Request != want {
					continue
				}
				if !yield(event, nil) {
					return
				}
			}
			seen += int64(len(page.Items))
			// A total of zero beside rows is a listing that published none, and
			// falls back to the empty page, which always ends the walk.
			if page.Total > 0 && seen >= page.Total {
				return
			}
			f.Page++
		}
	}
}

// query renders the filter as the route's parameters. Paging is stringly typed
// on the wire; the SDK's is not.
func (f Filter) query() url.Values {
	query := url.Values{}
	for name, value := range map[string]string{
		"sub":        f.Actor,
		"action":     f.Action,
		"resource":   f.Resource,
		"resourceId": f.ID,
		"result":     f.Result,
	} {
		if value != "" {
			query.Set(name, value)
		}
	}
	if !f.Since.IsZero() {
		query.Set("since", f.Since.UTC().Format(time.RFC3339))
	}
	if !f.Until.IsZero() {
		query.Set("until", f.Until.UTC().Format(time.RFC3339))
	}
	if f.Size > 0 {
		query.Set("pageSize", strconv.Itoa(f.Size))
	}
	if f.Page > 0 {
		query.Set("p", strconv.Itoa(f.Page))
	}
	return query
}
