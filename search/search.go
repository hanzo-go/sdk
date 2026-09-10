// Package search is one ranked result set over everything the org has stored.
//
// Degradation is stated, never silent. A leg that is down produces the
// surviving legs' results plus a degraded [Backend] entry carrying the failure.
// That is not a refusal and not an error: it is a first-class field, because a
// caller that ignores [Hits.Partial] reads a truncated corpus as a complete one.
package search

import (
	"context"
	"encoding/json"
	"time"

	"github.com/hanzoai/go-sdk/v8/call"
	"github.com/hanzoai/go-sdk/v8/kb"
)

// Client searches the caller's own org.
type Client struct{ e *call.Endpoint }

// New returns the search capability over an endpoint.
func New(e *call.Endpoint) *Client { return &Client{e: e} }

// Opts narrows a search. The zero value searches everything the org has, in
// auto mode.
type Opts struct {
	// Mode is auto, text, semantic or hybrid; empty is auto. The modes name
	// retrieval KINDS, not backends — a caller chooses how to search, never
	// which subsystem answers.
	Mode string
	// Project narrows to one project scope within the org.
	Project string
	// Kinds restricts retrieval to a subset of the indexed knowledge kinds:
	// page, memory, source — the same word kb.Doc.Kind takes. The route filters
	// on the doctype address and silently ignores anything that is not one, so
	// the mapping happens here rather than in a caller's head.
	Kinds []string
	// Index names the lexical index to query.
	Index string
	// Limit bounds the fused result set; Offset pages it.
	Limit  int
	Offset int
}

// Hits is one fused, ranked result set with the honesty of the query beside it.
type Hits struct {
	// Status is "ok" where every consulted leg answered, "partial" where one
	// did not, "unavailable" where none did.
	Status string
	// Partial is Status != "ok" — the one field a caller must read.
	Partial bool
	// Mode is the mode actually used, after auto resolved.
	Mode     string
	Items    []Hit
	Backends []Backend
	Took     time.Duration
}

// Hit is one document the query matched.
type Hit struct {
	// ID is the document's identity inside its corpus.
	ID     string
	Corpus string
	// Kind is the knowledge kind that matched: page, memory or source, the same
	// word kb.Doc.Kind takes. A hit out of another corpus carries that corpus's
	// own type and keeps it.
	Kind    string
	Title   string
	URL     string
	Project string
	// Score is a reciprocal-rank fusion sum, comparable ONLY within one
	// response. It is not a relevance and not a similarity: never compare it
	// across queries, and never against a backend's own score, which stays in
	// Matched.
	Score   float64
	Matched []Match
}

// Match is one leg's contribution to a hit, on that leg's own scale.
type Match struct {
	Backend string  `json:"backend"`
	Rank    int     `json:"rank"`
	Score   float64 `json:"score"`
}

// Backend is one leg's report. Every consulted leg appears, working or not.
type Backend struct {
	Name string
	// Status is ok, degraded, disabled or skipped — four distinct operational
	// facts, never folded together.
	Status string
	// Hits is what this leg returned before fusion.
	Hits  int
	Took  time.Duration
	Error string
}

// Find searches the org's whole corpus.
//
// The rerank leg runs through the metered AI gateway, so a search can be
// refused on money: the answer's denied arm carries insufficient_balance or
// spend_cap_exceeded and the cures that clear it.
//
// The semantic leg has a second address of its own, POST /v1/knowledge/search,
// with a second request shape and a second hit shape. No SDK binds it; it is
// reached as Find(ctx, q, Opts{Mode: "semantic"}).
func (c *Client) Find(ctx context.Context, query string, o Opts) call.Answer[Hits] {
	in := struct {
		Query    string   `json:"query"`
		Mode     string   `json:"mode,omitempty"`
		Project  string   `json:"project,omitempty"`
		Doctypes []string `json:"doctypes,omitempty"`
		Index    string   `json:"index,omitempty"`
		Limit    int      `json:"limit,omitempty"`
		Offset   int      `json:"offset,omitempty"`
	}{query, o.Mode, o.Project, doctypes(o.Kinds), o.Index, o.Limit, o.Offset}

	return call.Ask[Hits](ctx, c.e, "POST", "/v1/search", nil, in)
}

// doctypes addresses the kinds a caller named. The route matches the doctype
// address exactly and drops anything else, so a query narrowed to "page" would
// come back narrowed to nothing at all.
func doctypes(kinds []string) []string {
	if len(kinds) == 0 {
		return nil
	}
	out := make([]string, len(kinds))
	for i, kind := range kinds {
		out[i] = kb.Doctype(kind)
	}
	return out
}

// UnmarshalJSON reads the fusion answer. The wire spells the result set `hits`
// and its durations in milliseconds; this type spells them as the reader wants
// them, and derives Partial from Status so the honesty signal is a field rather
// than a string comparison every caller repeats.
func (h *Hits) UnmarshalJSON(raw []byte) error {
	var wire struct {
		Status   string `json:"status"`
		Mode     string `json:"mode"`
		TookMS   int64  `json:"took_ms"`
		Backends []struct {
			Name   string `json:"name"`
			Status string `json:"status"`
			Hits   int    `json:"hits"`
			TookMS int64  `json:"took_ms"`
			Error  string `json:"error"`
		} `json:"backends"`
		Hits []struct {
			ID      string  `json:"id"`
			Corpus  string  `json:"corpus"`
			Doctype string  `json:"doctype"`
			Title   string  `json:"title"`
			URL     string  `json:"url"`
			Project string  `json:"project"`
			Score   float64 `json:"score"`
			Matched []Match `json:"matched"`
		} `json:"hits"`
	}
	if err := json.Unmarshal(raw, &wire); err != nil {
		return err
	}

	*h = Hits{
		Status:  wire.Status,
		Partial: wire.Status != "ok",
		Mode:    wire.Mode,
		Took:    time.Duration(wire.TookMS) * time.Millisecond,
	}
	for _, hit := range wire.Hits {
		h.Items = append(h.Items, Hit{
			ID: hit.ID, Corpus: hit.Corpus, Kind: kb.Kind(hit.Doctype), Title: hit.Title,
			URL: hit.URL, Project: hit.Project, Score: hit.Score, Matched: hit.Matched,
		})
	}
	for _, backend := range wire.Backends {
		h.Backends = append(h.Backends, Backend{
			Name: backend.Name, Status: backend.Status, Hits: backend.Hits,
			Took: time.Duration(backend.TookMS) * time.Millisecond, Error: backend.Error,
		})
	}
	return nil
}
