// Package kb is the corpus you write and the connectors that fill it.
//
// Searching it belongs to search: kb.Find would be a second way to ask one
// question. This package is the write side, the connectors, and the corpus's
// own link structure.
//
// Cloud's shape here is awkward and the SDK's is not. KB writes are not under
// /v1/knowledge at all — they are generic doctype CRUD at /v1/framework/kb.page,
// kb.memory and kb.source, whose field names differ per kind: a page holds its
// text in `body` and is named by its slug, a memory holds it in `content`, a
// source carries a `url` back to where it came from. [Doc] is the one shape
// over the three, and [Doc.MarshalJSON] is the single place that knows the
// split.
package kb

import (
	"context"
	"encoding/json"
	"io"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/hanzoai/go-sdk/v8/call"
)

// Client reads and writes the caller's own org's corpus.
type Client struct{ e *call.Endpoint }

// New returns the kb capability over an endpoint.
func New(e *call.Endpoint) *Client { return &Client{e: e} }

// The three kinds of knowledge a caller writes. Anything else in the framework
// is not this capability's.
const (
	Page   = "page"
	Memory = "memory"
	Source = "source"
)

// Doc is one document, whichever of the three kinds it is.
type Doc struct {
	// Kind is page, memory or source.
	Kind string
	// Name is the document's id within its kind. A page is named by its slug,
	// so a page needs one; a memory and a source are named by the store.
	Name    string
	Title   string
	Body    string
	Project string
	// URL is where the document came from. Only a source carries one.
	URL string
}

// Filter narrows a listing.
type Filter struct {
	Project string
	Limit   int
}

// Connector is one provider and this org's connection to it.
type Connector struct {
	Provider string
	// Kind is "native" for a first-party connector, "piece" for a long-tail one.
	Kind string
	// Status is connected, disconnected, syncing or error.
	Status string
	// Configured is whether this deployment holds OAuth credentials for the
	// provider at all.
	Configured bool
	Account    string
	// Docs is the live count of this provider's documents in the org's store.
	Docs   int
	Synced time.Time
	Error  string
}

// Link is where to send a person to authorize a connector.
type Link struct{ URL string }

// Sync is what one connector pull landed.
type Sync struct {
	Provider string `json:"provider"`
	Ingested int    `json:"ingested"`
}

// Export is an archive to import: an Obsidian or Notion vault zip, a Roam JSON,
// an Evernote .enex.
type Export struct {
	// Format is obsidian, notion, roam or evernote. It picks the normalizer.
	Format string
	// Project narrows every imported page to one scope.
	Project string
	Body    io.Reader
}

// Import is what an export actually filed — what landed, never what was sent.
type Import struct {
	Format   string   `json:"format"`
	Imported int      `json:"imported"`
	Pages    []string `json:"pages"`
}

// Reindex is what a rebuild of the org's indexes moved.
type Reindex struct {
	Vectors int `json:"vectors"`
	Lexical int `json:"lexical"`
	Removed int `json:"removed"`
	Failed  int `json:"failed"`
}

// Links is the corpus's own parent, wikilink and provenance edges.
//
// This is NOT graph. It describes documents, not assertions, and the two never
// share a type.
type Links struct {
	Nodes []Node `json:"nodes"`
	Edges []Edge `json:"edges"`
	// Partial is true where the store was unreachable and this answer is
	// honestly empty rather than genuinely so.
	Partial bool `json:"degraded"`
}

// Node is one document, connector or unresolved link target.
type Node struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Title string `json:"title"`
	// Kind is what the node is: kb.page, kb.memory, kb.source, kb.connector or
	// unresolved. The route spells it `type`; `kind` is the word this client
	// uses for what a thing is, everywhere.
	Kind    string `json:"type"`
	Project string `json:"project"`
}

// Edge is one parent, link or provenance relation between two nodes.
type Edge struct {
	From string `json:"from"`
	To   string `json:"to"`
	Kind string `json:"kind"`
}

// Put writes a document.
//
// A doc carrying a Name is written AT that name and replaces what stands there;
// a doc without one is created and the store names it. A page is named by its
// slug, so a page needs a Name.
func (c *Client) Put(ctx context.Context, doc Doc) call.Answer[Doc] {
	if doc.Name == "" {
		return call.Ask[Doc](ctx, c.e, "POST", doctype(doc.Kind), nil, doc)
	}
	return call.Ask[Doc](ctx, c.e, "PUT", doctype(doc.Kind)+"/"+url.PathEscape(doc.Name), nil, doc)
}

// Get reads one document by kind and name.
func (c *Client) Get(ctx context.Context, kind, name string) (Doc, error) {
	var doc Doc
	_, err := call.Do(ctx, c.e, "GET", doctype(kind)+"/"+url.PathEscape(name), nil, nil, &doc)
	return doc, err
}

// List reads a page of documents of one kind, newest-updated first.
func (c *Client) List(ctx context.Context, kind string, f Filter) (call.Page[Doc], error) {
	query := url.Values{}
	if f.Project != "" {
		// The route takes its narrowing as a JSON object of field to value.
		narrow, err := json.Marshal(map[string]string{"project": f.Project})
		if err != nil {
			return call.Page[Doc]{}, err
		}
		query.Set("filters", string(narrow))
	}
	if f.Limit > 0 {
		query.Set("limit", strconv.Itoa(f.Limit))
	}

	var wire struct {
		Data []Doc `json:"data"`
	}
	if _, err := call.Do(ctx, c.e, "GET", doctype(kind), query, nil, &wire); err != nil {
		return call.Page[Doc]{}, err
	}
	// The route publishes no count, so the total is what came back. It stops
	// being a floor the day the listing carries one.
	return call.Page[Doc]{Items: wire.Data, Total: int64(len(wire.Data))}, nil
}

// Drop deletes one document.
func (c *Client) Drop(ctx context.Context, kind, name string) call.Answer[struct{}] {
	return call.Ask[struct{}](ctx, c.e, "DELETE", doctype(kind)+"/"+url.PathEscape(name), nil, nil)
}

// Import files an export as a tree of pages with its link structure intact.
func (c *Client) Import(ctx context.Context, export Export) call.Answer[Import] {
	query := url.Values{"format": {export.Format}}
	if export.Project != "" {
		query.Set("project", export.Project)
	}
	return call.Ask[Import](ctx, c.e, "POST", "/v1/knowledge/import", query, export.Body)
}

// Reindex rebuilds the org's retrieval indexes over what it already holds.
func (c *Client) Reindex(ctx context.Context) call.Answer[Reindex] {
	return call.Ask[Reindex](ctx, c.e, "POST", "/v1/knowledge/reindex", nil, nil)
}

// Connectors lists every supported provider with this org's connection state.
func (c *Client) Connectors(ctx context.Context) ([]Connector, error) {
	var wire struct {
		Connectors []struct {
			Provider   string    `json:"provider"`
			Kind       string    `json:"kind"`
			Status     string    `json:"status"`
			Configured bool      `json:"configured"`
			Account    string    `json:"account"`
			DocCount   int       `json:"docCount"`
			LastSync   time.Time `json:"lastSync"`
			Error      string    `json:"error"`
		} `json:"connectors"`
	}
	if _, err := call.Do(ctx, c.e, "GET", "/v1/knowledge/connectors", nil, nil, &wire); err != nil {
		return nil, err
	}
	out := make([]Connector, 0, len(wire.Connectors))
	for _, row := range wire.Connectors {
		out = append(out, Connector{
			Provider: row.Provider, Kind: row.Kind, Status: row.Status,
			Configured: row.Configured, Account: row.Account, Docs: row.DocCount,
			Synced: row.LastSync, Error: row.Error,
		})
	}
	return out, nil
}

// Connect answers where to send a person to authorize a provider.
//
// The SDK never follows it: an OAuth consent screen is not a client's to
// complete.
func (c *Client) Connect(ctx context.Context, provider string) (Link, error) {
	var wire struct {
		AuthorizeURL string `json:"authorizeUrl"`
	}
	_, err := call.Do(ctx, c.e, "GET", "/v1/knowledge/connectors/"+url.PathEscape(provider)+"/connect", nil, nil, &wire)
	return Link{URL: wire.AuthorizeURL}, err
}

// Sync pulls a connected provider now.
func (c *Client) Sync(ctx context.Context, provider string) call.Answer[Sync] {
	return call.Ask[Sync](ctx, c.e, "POST", "/v1/knowledge/connectors/"+url.PathEscape(provider)+"/sync", nil, nil)
}

// Revoke disconnects a provider and forgets its credential.
func (c *Client) Revoke(ctx context.Context, provider string) call.Answer[struct{}] {
	return call.Ask[struct{}](ctx, c.e, "DELETE", "/v1/knowledge/connectors/"+url.PathEscape(provider), nil, nil)
}

// Links reads the corpus's own graph of documents.
func (c *Client) Links(ctx context.Context) (Links, error) {
	var links Links
	_, err := call.Do(ctx, c.e, "GET", "/v1/knowledge/graph", nil, nil, &links)
	return links, err
}

// Install creates the kb doctypes in the caller's org, once, before the first
// write.
//
// It is in the surface only because cloud requires it. It leaves the day a
// write to a module's doctype installs the module itself.
func (c *Client) Install(ctx context.Context) call.Answer[struct{}] {
	return call.Ask[struct{}](ctx, c.e, "POST", "/v1/framework/modules/kb/install", nil, nil)
}

// doctype is the address of one kind. The kb module owns three doctypes and
// this is the only place their spelling appears.
func doctype(kind string) string { return "/v1/framework/kb." + kind }

// MarshalJSON writes a document as the kind's own doctype declares it: a page
// keeps its text in `body` and its name in `slug`, a memory keeps it in
// `content`, a source carries `url` beside `body`.
func (d Doc) MarshalJSON() ([]byte, error) {
	out := map[string]any{}
	if d.Name != "" {
		out["name"] = d.Name
	}
	if d.Title != "" {
		out["title"] = d.Title
	}
	if d.Project != "" {
		out["project"] = d.Project
	}
	switch d.Kind {
	case Memory:
		out["content"] = d.Body
	case Source:
		out["body"] = d.Body
		if d.URL != "" {
			out["url"] = d.URL
		}
	default:
		out["body"] = d.Body
		out["slug"] = d.Name
	}
	return json.Marshal(out)
}

// UnmarshalJSON reads a stored document back into the one shape, taking the
// kind from the doctype the store stamped on it.
func (d *Doc) UnmarshalJSON(raw []byte) error {
	var wire struct {
		Doctype string `json:"doctype"`
		Name    string `json:"name"`
		Title   string `json:"title"`
		Body    string `json:"body"`
		Content string `json:"content"`
		Project string `json:"project"`
		URL     string `json:"url"`
	}
	if err := json.Unmarshal(raw, &wire); err != nil {
		return err
	}
	body := wire.Body
	if body == "" {
		body = wire.Content
	}
	*d = Doc{
		Kind:    kindOf(wire.Doctype),
		Name:    wire.Name,
		Title:   wire.Title,
		Body:    body,
		Project: wire.Project,
		URL:     wire.URL,
	}
	return nil
}

// kindOf takes the kind out of a doctype address: kb.page is page.
func kindOf(doctype string) string { return strings.TrimPrefix(doctype, "kb.") }
