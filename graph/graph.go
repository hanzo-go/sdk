// Package graph is assertions with provenance and time.
//
// Nothing overwrites anything. A retraction is an assertion, a superseded claim
// stays readable beside the one that superseded it, and which of two claims is
// in force is a question [Client.Resolve] answers rather than a fact the store
// destroys.
//
// The substrate is HIP-0526's. The SDK talks to /v1/graph; it embeds no store.
package graph

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/hanzoai/go-sdk/v8/call"
)

// Client reads and writes the caller's own org's assertions.
type Client struct{ e *call.Endpoint }

// New returns the graph capability over an endpoint.
func New(e *call.Endpoint) *Client { return &Client{e: e} }

// Fact is one assertion: an entity, what is claimed of it, and who says so.
//
// Times are RFC 3339 on the wire and time.Time here. At is when the thing was
// SO; Seen is when it became knowable and defaults to At. Cloud refuses an At
// more than five minutes ahead of its clock.
type Fact struct {
	// Entity is the thing being described, in the org's own namespace.
	Entity string `json:"entity"`
	// Relation is what is asserted of it: depends, owner, same, title.
	Relation string `json:"relation"`
	// Value is what the relation points at.
	Value string `json:"value"`
	// Names declares that Value is another entity's key, which makes the
	// assertion an edge. A walk reads only edges, so this is a decision, not a
	// hint.
	Names bool      `json:"names,omitempty"`
	At    time.Time `json:"at,omitzero"`
	Seen  time.Time `json:"seen,omitzero"`
	// Source names who asserted. An assertion nobody is named for cannot be
	// weighed against another.
	Source string `json:"source,omitempty"`
	// Evidence points at the record the claim came from.
	Evidence string `json:"evidence,omitempty"`
	// Confidence in [0,1] breaks a tie within the order, never substitutes for
	// it.
	Confidence float64 `json:"confidence,omitempty"`

	// ID is the assertion's content address, minted by the server. Read only.
	ID string `json:"id,omitempty"`
	// By is the identity that filed it, stamped from the validated principal.
	// Read only.
	By string `json:"by,omitempty"`
	// Knowable is the first instant this plane could have answered with the
	// assertion, which is what a read at an instant is bounded by. Read only.
	Knowable time.Time `json:"knowable,omitzero"`
}

// Wrote is what a batch of assertions did. Each member is judged on its own:
// one malformed fact does not discard the batch.
type Wrote struct {
	// Recorded is how many members became new rows.
	Recorded int `json:"recorded"`
	// Duplicate is how many the plane already held.
	Duplicate int `json:"duplicate"`
	// Refused is how many were turned away on arrival.
	Refused int `json:"refused"`
	// Reasons names why each refused member was refused, in the order sent.
	Reasons []string `json:"reasons"`
}

// Resolution is which claim about one (entity, relation) is in force.
type Resolution struct {
	Entity   string    `json:"entity"`
	Relation string    `json:"relation"`
	At       time.Time `json:"as_of"`
	// Known is false where the plane held nothing knowable at At. That is an
	// answer, not an error.
	Known bool `json:"known"`
	// Winner is the strongest assertion knowable at At.
	Winner *Fact `json:"winner,omitempty"`
	// Conflicts is every other assertion knowable at At, strongest first.
	Conflicts []Fact `json:"conflicts"`
	// Contested is true only where a conflict claims a DIFFERENT value than the
	// winner. Two sources agreeing is not a conflict.
	Contested bool `json:"contested"`
	// Truncated says the pair holds more assertions than one read returns.
	Truncated bool `json:"truncated"`
}

// Walk is what following edges from a set of seeds reached.
type Walk struct {
	// Entities is everything reached, the seeds included, ordered by the fewest
	// hops to it.
	Entities []string `json:"entities"`
	// Depth is the deepest hop count actually reached.
	Depth int `json:"depth"`
	// Bound is the ceiling this walk was held to, the same for every caller.
	Bound int `json:"bound"`
	// Truncated says the bound stopped the walk.
	Truncated bool `json:"truncated"`
}

// Triple is one relation a document states, read without recording it.
type Triple struct {
	Subject   string `json:"subject"`
	Predicate string `json:"predicate"`
	Object    string `json:"object"`
	Names     bool   `json:"names"`
	// Section is which section of the source stated it, counting from zero.
	Section int `json:"section"`
}

// Vocabulary is the relations an org has actually asserted and the ordering
// that decides a conflict between two of them.
type Vocabulary struct {
	Relations []string `json:"relations"`
	// Rule names the terms of the precedence order, in the order they apply.
	Rule []string `json:"rule"`
	// Bound is the ceiling on one walk.
	Bound int `json:"bound"`
}

// Source is a document to read relations out of, and where it came from.
type Source struct {
	// Source names the origin — a URL, a document id, a page title.
	Source string `json:"source"`
	// Subject is the entity the text is about before any heading names one.
	Subject string    `json:"subject"`
	Text    string    `json:"text"`
	At      time.Time `json:"at,omitzero"`
}

// Filter narrows a read by key.
//
// [Client.Find] shares it: a text search is a read whose key is words, so
// Relation, At and Limit apply there too. Entity and Value name a key and a
// text search has none, so Find refuses a filter carrying either rather than
// dropping it silently.
type Filter struct {
	Entity   string
	Relation string
	Value    string
	// At bounds the read to what was knowable at an instant. Zero reads now.
	At    time.Time
	Limit int
}

// Opts narrows a walk.
type Opts struct {
	// Relation follows one edge relation. Empty follows all.
	Relation string
	// Direction is out, in or both. Empty is out.
	Direction string
	// Depth is how many hops. Zero is one.
	Depth int
	// At walks the graph as it stood at an instant. Zero walks it as it stands.
	At time.Time
}

// Assert records a batch of assertions.
//
// The answer accounts for every member: Recorded plus Duplicate plus Refused
// equals the batch. A batch that does not add up is a transport fault and
// surfaces as an error rather than as an answer, because a count that does not
// account for what was sent cannot be acted on.
func (c *Client) Assert(ctx context.Context, facts []Fact) call.Answer[Wrote] {
	in := struct {
		Assertions []Fact `json:"assertions"`
	}{asserted(facts)}

	answer := call.Ask[Wrote](ctx, c.e, "POST", "/v1/graph", nil, in)
	wrote, err := answer.Value()
	if err != nil {
		return answer
	}
	if got := wrote.Recorded + wrote.Duplicate + wrote.Refused; got != len(facts) {
		return call.Failed[Wrote](answer.Request,
			fmt.Errorf("hanzoai: graph accounted for %d of %d assertions", got, len(facts)))
	}
	return answer
}

// Read answers assertions by key.
//
// It resolves nothing and withholds nothing: a superseded claim and the one
// that superseded it both appear.
func (c *Client) Read(ctx context.Context, f Filter) ([]Fact, error) {
	query := url.Values{}
	if f.Entity != "" {
		query.Set("entity", f.Entity)
	}
	if f.Value != "" {
		query.Set("value", f.Value)
	}
	f.narrow(query)
	return c.assertions(ctx, "/v1/graph", query)
}

// Find answers assertions by text. The same word as search.Find, because it is
// the same act on a different corpus.
func (c *Client) Find(ctx context.Context, query string, f Filter) ([]Fact, error) {
	if f.Entity != "" || f.Value != "" {
		return nil, errors.New("hanzoai: graph.Find searches by text; Entity and Value name a key, which graph.Read takes")
	}
	params := url.Values{"q": {query}}
	f.narrow(params)
	return c.assertions(ctx, "/v1/graph/search", params)
}

// Resolve answers which claim about one entity and relation is in force. A zero
// at asks about now.
func (c *Client) Resolve(ctx context.Context, entity, relation string, at time.Time) (Resolution, error) {
	in := struct {
		Entity   string    `json:"entity"`
		Relation string    `json:"relation"`
		AsOf     time.Time `json:"as_of,omitzero"`
	}{entity, relation, at}

	var out Resolution
	_, err := call.Do(ctx, c.e, "POST", "/v1/graph/resolve", nil, in, &out)
	return out, err
}

// Walk follows edges out of a set of seeds. Only edges are followed: an
// assertion whose value is not another entity's key is not a hop.
func (c *Client) Walk(ctx context.Context, seeds []string, o Opts) (Walk, error) {
	in := struct {
		Seeds     []string  `json:"seeds"`
		Relation  string    `json:"relation,omitempty"`
		Direction string    `json:"direction,omitempty"`
		Depth     int       `json:"depth,omitempty"`
		AsOf      time.Time `json:"as_of,omitzero"`
	}{seeds, o.Relation, o.Direction, o.Depth, o.At}

	var out Walk
	_, err := call.Do(ctx, c.e, "POST", "/v1/graph/neighbors", nil, in, &out)
	return out, err
}

// Extract reads what a document states without recording any of it.
func (c *Client) Extract(ctx context.Context, source Source) ([]Triple, error) {
	var out struct {
		Triples []Triple `json:"triples"`
	}
	_, err := call.Do(ctx, c.e, "POST", "/v1/graph/extract", nil, source, &out)
	return out.Triples, err
}

// Ingest extracts and asserts in one call, answering the same [Wrote].
func (c *Client) Ingest(ctx context.Context, source Source) call.Answer[Wrote] {
	return call.Ask[Wrote](ctx, c.e, "POST", "/v1/graph/ingest", nil, source)
}

// Vocabulary answers the relations in use and the ordering that decides a
// conflict.
func (c *Client) Vocabulary(ctx context.Context) (Vocabulary, error) {
	var out Vocabulary
	_, err := call.Do(ctx, c.e, "GET", "/v1/graph/vocabulary", nil, nil, &out)
	return out, err
}

// asserted is the batch as an asserter states it. ID, By and Knowable are the
// server's — it mints the content address, stamps the filer and derives when
// the row became knowable — so a fact read back and asserted again sends the
// nine members the route declares and none of the three it does not.
func asserted(facts []Fact) []Fact {
	out := make([]Fact, len(facts))
	for i, fact := range facts {
		fact.ID, fact.By, fact.Knowable = "", "", time.Time{}
		out[i] = fact
	}
	return out
}

// assertions is the one decode both reads share: the two routes answer the same
// body.
func (c *Client) assertions(ctx context.Context, path string, query url.Values) ([]Fact, error) {
	var out struct {
		Assertions []Fact `json:"assertions"`
	}
	_, err := call.Do(ctx, c.e, "GET", path, query, nil, &out)
	return out.Assertions, err
}

// narrow adds the parameters a read and a text search share.
func (f Filter) narrow(query url.Values) {
	if f.Relation != "" {
		query.Set("relation", f.Relation)
	}
	if !f.At.IsZero() {
		query.Set("as_of", f.At.UTC().Format(time.RFC3339))
	}
	if f.Limit > 0 {
		query.Set("limit", strconv.Itoa(f.Limit))
	}
}
