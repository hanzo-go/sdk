package hanzoai

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/hanzoai/go-sdk/v8/audit"
	"github.com/hanzoai/go-sdk/v8/budget"
	"github.com/hanzoai/go-sdk/v8/call"
	"github.com/hanzoai/go-sdk/v8/graph"
	"github.com/hanzoai/go-sdk/v8/kb"
	"github.com/hanzoai/go-sdk/v8/policy"
	"github.com/hanzoai/go-sdk/v8/search"
)

// The six, measured as a caller reaches them: what the SDK SENDS for each
// method, and what it makes of a realistic answer.
//
// Every row states the exact request — method, path, query and body — because
// that is the half of a client contract a decode cannot check. A method that
// posts the right shape to the wrong address passes every test that only reads
// the reply.

// at is the instant the rows are written and read at. RFC 3339 on the wire.
var at = time.Date(2026, 9, 10, 12, 0, 0, 0, time.UTC)

// request is the x-request-id the stand-in stamps, so a row can assert the join
// to the audit trail.
const request = "req-abc123"

// seen is what one call carried.
type seen struct {
	method string
	path   string
	query  string
	body   string
	auth   []string
}

// serve stands up one server answering IAM's mint and the platform API, and
// returns a client pointed at it. The mint answers here too: a client that
// cannot mint cannot call, so every row exercises the credential as well.
func serve(t *testing.T, status int, answer string, got *seen) *Client {
	t.Helper()
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/iam/oauth/token", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"access_token":"tok","expires_in":3600}`))
	})
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		*got = seen{r.Method, r.URL.Path, r.URL.RawQuery, string(body), r.Header.Values("Authorization")}
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-Request-Id", request)
		w.WriteHeader(status)
		w.Write([]byte(answer))
	})
	srv := httptest.NewServer(mux)
	t.Cleanup(srv.Close)
	return New(Options{ID: "cid", Secret: "sec", Base: srv.URL, Issuer: srv.URL, Resource: srv.URL})
}

func TestCapabilities(t *testing.T) {
	ctx := context.Background()

	for _, tc := range []struct {
		name   string
		status int
		answer string
		call   func(*Client) (any, error)

		method string
		path   string
		query  string
		body   string
		want   any
	}{{
		name:   "budget.left",
		answer: `{"plan":"pro","limit":20,"used":3,"spent":false,"window":"day","resets":1757548800}`,
		call:   func(c *Client) (any, error) { return c.Budget.Left(ctx) },
		method: "GET", path: "/v1/allowance",
		want: budget.Allowance{
			Plan: "pro", Limit: 20, Used: 3, Left: left(17), Spent: false, Window: "day",
			Resets: when(time.Unix(1757548800, 0).UTC()),
		},
	}, {
		// An unbounded plan has no remainder and no period to end, so neither
		// is reported as a zero.
		name:   "budget.left unbounded",
		answer: `{"plan":"enterprise","limit":0,"used":9000,"window":""}`,
		call:   func(c *Client) (any, error) { return c.Budget.Left(ctx) },
		method: "GET", path: "/v1/allowance",
		want: budget.Allowance{Plan: "enterprise", Limit: 0, Used: 9000},
	}, {
		name:   "budget.balance",
		answer: `{"balance":12345,"holds":0,"available":12345,"account":"acme"}`,
		call:   func(c *Client) (any, error) { return c.Budget.Balance(ctx) },
		method: "GET", path: "/v1/billing/balance",
		want: budget.Balance{
			Available: call.Money{Cents: 12345, Currency: "usd"},
			Held:      call.Money{Cents: 0, Currency: "usd"},
			Account:   "acme",
		},
	}, {
		name:   "budget.plan",
		answer: `{"tier":"pro","apps":{"chat":true,"dataroom":false}}`,
		call:   func(c *Client) (any, error) { return c.Budget.Plan(ctx) },
		method: "GET", path: "/v1/entitlement",
		want: budget.Plan{Tier: "pro", Apps: map[string]bool{"chat": true, "dataroom": false}},
	}, {
		name:   "budget.spent",
		answer: `{"user":"acme","count":1,"usage":[{"transactionId":"tx_1","amount":42,"metadata":{"model":"zen-1"},"createdAt":"2026-09-10T12:00:00Z"}]}`,
		call: func(c *Client) (any, error) {
			return c.Budget.Spent(ctx, budget.Filter{Product: "inference", Since: at})
		},
		method: "GET", path: "/v1/billing/usage",
		query: "product=inference&start=2026-09-10T12%3A00%3A00Z",
		want: call.Page[budget.Charge]{
			Total: 1,
			Items: []budget.Charge{{ID: "tx_1", At: at, Model: "zen-1", Amount: call.Money{Cents: 42, Currency: "usd"}}},
		},
	}, {
		// The arguments read as the sentence does; the body goes out in the
		// members the route binds.
		name:   "policy.check",
		answer: `{"allow":true,"subject":"usr_7","verb":"write","path":"acme/graph"}`,
		call:   func(c *Client) (any, error) { return c.Policy.Check(ctx, "usr_7", "write", "acme/graph") },
		method: "POST", path: "/v1/authz/check",
		body: `{"subject":"usr_7","verb":"write","path":"acme/graph"}`,
		want: policy.Decision{Allow: true, Sub: "usr_7", Act: "write", Obj: "acme/graph"},
	}, {
		name:   "audit.list",
		answer: `{"status":"ok","total":1,"data":[{"seq":7,"org":"acme","sub":"usr_7","action":"graph.assert","resource":"graph","resourceId":"e1","result":"success","status":200,"method":"POST","path":"/v1/graph","requestId":"req-abc123","sourceIp":"203.0.113.4","userAgent":"go-sdk","time":"2026-09-10T12:00:00Z"}]}`,
		call: func(c *Client) (any, error) {
			return c.Audit.List(ctx, audit.Filter{Actor: "usr_7", Action: "graph.assert", Since: at, Size: 50, Page: 2})
		},
		method: "GET", path: "/v1/audit",
		query: "action=graph.assert&p=2&pageSize=50&since=2026-09-10T12%3A00%3A00Z&sub=usr_7",
		want: call.Page[audit.Event]{
			Total: 1,
			Items: []audit.Event{{
				Seq: 7, Org: "acme", Actor: "usr_7", Action: "graph.assert", Resource: "graph",
				ID: "e1", Method: "POST", Path: "/v1/graph", Result: "success", Status: 200,
				Request: request, IP: "203.0.113.4", Agent: "go-sdk", At: at,
			}},
		},
	}, {
		name:   "search.find",
		answer: `{"status":"partial","mode":"hybrid","took_ms":42,"hits":[{"id":"kb.page/runbook","corpus":"kb","doctype":"kb.page","title":"Runbook","url":"/kb/runbook","project":"ops","score":0.031,"matched":[{"backend":"index","rank":1,"score":8.2}]}],"backends":[{"name":"vector","status":"degraded","hits":0,"took_ms":5,"error":"collection cold"}]}`,
		call: func(c *Client) (any, error) {
			return value(c.Search.Find(ctx, "runbook", search.Opts{Mode: "hybrid", Kinds: []string{"kb.page"}, Limit: 5}))
		},
		method: "POST", path: "/v1/search",
		body: `{"query":"runbook","mode":"hybrid","doctypes":["kb.page"],"limit":5}`,
		want: search.Hits{
			Status: "partial", Partial: true, Mode: "hybrid", Took: 42 * time.Millisecond,
			Items: []search.Hit{{
				ID: "kb.page/runbook", Corpus: "kb", Kind: "kb.page", Title: "Runbook",
				URL: "/kb/runbook", Project: "ops", Score: 0.031,
				Matched: []search.Match{{Backend: "index", Rank: 1, Score: 8.2}},
			}},
			Backends: []search.Backend{{Name: "vector", Status: "degraded", Took: 5 * time.Millisecond, Error: "collection cold"}},
		},
	}, {
		// One method, three kinds: the doctype is the address and the kind's
		// own field names are the body.
		name:   "kb.put memory",
		answer: `{"doctype":"kb.memory","name":"m_1","title":"Note","content":"the text","project":"ops"}`,
		call: func(c *Client) (any, error) {
			return value(c.KB.Put(ctx, kb.Doc{Kind: kb.Memory, Title: "Note", Body: "the text", Project: "ops"}))
		},
		method: "POST", path: "/v1/framework/kb.memory",
		body: `{"content":"the text","project":"ops","title":"Note"}`,
		want: kb.Doc{Kind: "memory", Name: "m_1", Title: "Note", Body: "the text", Project: "ops"},
	}, {
		name:   "kb.put page at a name",
		answer: `{"doctype":"kb.page","name":"runbook","title":"Runbook","body":"the text"}`,
		call: func(c *Client) (any, error) {
			return value(c.KB.Put(ctx, kb.Doc{Kind: kb.Page, Name: "runbook", Title: "Runbook", Body: "the text"}))
		},
		method: "PUT", path: "/v1/framework/kb.page/runbook",
		body: `{"body":"the text","name":"runbook","slug":"runbook","title":"Runbook"}`,
		want: kb.Doc{Kind: "page", Name: "runbook", Title: "Runbook", Body: "the text"},
	}, {
		name:   "kb.list",
		answer: `{"data":[{"doctype":"kb.page","name":"runbook","title":"Runbook","body":"b","project":"ops"}]}`,
		call:   func(c *Client) (any, error) { return c.KB.List(ctx, kb.Page, kb.Filter{Project: "ops", Limit: 10}) },
		method: "GET", path: "/v1/framework/kb.page",
		query: "filters=%7B%22project%22%3A%22ops%22%7D&limit=10",
		want: call.Page[kb.Doc]{
			Total: 1,
			Items: []kb.Doc{{Kind: "page", Name: "runbook", Title: "Runbook", Body: "b", Project: "ops"}},
		},
	}, {
		name:   "kb.connectors",
		answer: `{"connectors":[{"provider":"github","kind":"native","status":"connected","configured":true,"account":"acme","docCount":12,"lastSync":"2026-09-10T12:00:00Z"}]}`,
		call:   func(c *Client) (any, error) { return c.KB.Connectors(ctx) },
		method: "GET", path: "/v1/knowledge/connectors",
		want: []kb.Connector{{
			Provider: "github", Kind: "native", Status: "connected",
			Configured: true, Account: "acme", Docs: 12, Synced: at,
		}},
	}, {
		name:   "kb.links",
		answer: `{"nodes":[{"id":"kb.page:runbook","name":"runbook","title":"Runbook","type":"kb.page"}],"edges":[{"from":"kb.page:runbook","to":"kb.page:index","kind":"parent"}],"degraded":true}`,
		call:   func(c *Client) (any, error) { return c.KB.Links(ctx) },
		method: "GET", path: "/v1/knowledge/graph",
		want: kb.Links{
			Nodes:   []kb.Node{{ID: "kb.page:runbook", Name: "runbook", Title: "Runbook", Type: "kb.page"}},
			Edges:   []kb.Edge{{From: "kb.page:runbook", To: "kb.page:index", Kind: "parent"}},
			Partial: true,
		},
	}, {
		name:   "graph.assert",
		answer: `{"recorded":1,"duplicate":0,"refused":0,"reasons":[]}`,
		call: func(c *Client) (any, error) {
			return value(c.Graph.Assert(ctx, []graph.Fact{{
				Entity: "svc/api", Relation: "owner", Value: "team/core",
				Names: true, At: at, Source: "runbook",
			}}))
		},
		method: "POST", path: "/v1/graph",
		body: `{"assertions":[{"entity":"svc/api","relation":"owner","value":"team/core","names":true,"at":"2026-09-10T12:00:00Z","source":"runbook"}]}`,
		want: graph.Wrote{Recorded: 1, Reasons: []string{}},
	}, {
		name:   "graph.read",
		answer: `{"assertions":[{"id":"a1","entity":"svc/api","relation":"owner","value":"team/core","names":true,"at":"2026-09-10T12:00:00Z","seen":"2026-09-10T12:00:00Z","knowable":"2026-09-10T12:00:00Z","by":"acme","source":"runbook","confidence":0.9}]}`,
		call: func(c *Client) (any, error) {
			return c.Graph.Read(ctx, graph.Filter{Entity: "svc/api", Relation: "owner", At: at, Limit: 10})
		},
		method: "GET", path: "/v1/graph",
		query: "as_of=2026-09-10T12%3A00%3A00Z&entity=svc%2Fapi&limit=10&relation=owner",
		want: []graph.Fact{{
			ID: "a1", Entity: "svc/api", Relation: "owner", Value: "team/core", Names: true,
			At: at, Seen: at, Knowable: at, By: "acme", Source: "runbook", Confidence: 0.9,
		}},
	}, {
		name:   "graph.find",
		answer: `{"assertions":[]}`,
		call: func(c *Client) (any, error) {
			return c.Graph.Find(ctx, "core", graph.Filter{Relation: "owner", Limit: 5})
		},
		method: "GET", path: "/v1/graph/search",
		query: "limit=5&q=core&relation=owner",
		want:  []graph.Fact{},
	}, {
		name:   "graph.resolve",
		answer: `{"entity":"svc/api","relation":"owner","as_of":"2026-09-10T12:00:00Z","known":true,"contested":false,"conflicts":[],"winner":{"id":"a1","entity":"svc/api","relation":"owner","value":"team/core","at":"2026-09-10T12:00:00Z"}}`,
		call:   func(c *Client) (any, error) { return c.Graph.Resolve(ctx, "svc/api", "owner", time.Time{}) },
		method: "POST", path: "/v1/graph/resolve",
		body: `{"entity":"svc/api","relation":"owner"}`,
		want: graph.Resolution{
			Entity: "svc/api", Relation: "owner", At: at, Known: true, Conflicts: []graph.Fact{},
			Winner: &graph.Fact{ID: "a1", Entity: "svc/api", Relation: "owner", Value: "team/core", At: at},
		},
	}, {
		name:   "graph.walk",
		answer: `{"entities":["svc/api","team/core"],"depth":1,"bound":500,"truncated":false}`,
		call: func(c *Client) (any, error) {
			return c.Graph.Walk(ctx, []string{"svc/api"}, graph.Opts{Relation: "owner", Direction: "out", Depth: 2})
		},
		method: "POST", path: "/v1/graph/neighbors",
		body: `{"seeds":["svc/api"],"relation":"owner","direction":"out","depth":2}`,
		want: graph.Walk{Entities: []string{"svc/api", "team/core"}, Depth: 1, Bound: 500},
	}, {
		name:   "graph.vocabulary",
		answer: `{"relations":["owner","depends"],"rule":["seen","confidence","source"],"bound":500}`,
		call:   func(c *Client) (any, error) { return c.Graph.Vocabulary(ctx) },
		method: "GET", path: "/v1/graph/vocabulary",
		want: graph.Vocabulary{
			Relations: []string{"owner", "depends"},
			Rule:      []string{"seen", "confidence", "source"},
			Bound:     500,
		},
	}, {
		name:   "graph.extract",
		answer: `{"triples":[{"subject":"svc/api","predicate":"owner","object":"team/core","names":true,"section":0}]}`,
		call: func(c *Client) (any, error) {
			return c.Graph.Extract(ctx, graph.Source{Source: "runbook.md", Subject: "svc/api", Text: "owner :: [[team/core]]", At: at})
		},
		method: "POST", path: "/v1/graph/extract",
		body: `{"source":"runbook.md","subject":"svc/api","text":"owner :: [[team/core]]","at":"2026-09-10T12:00:00Z"}`,
		want: []graph.Triple{{Subject: "svc/api", Predicate: "owner", Object: "team/core", Names: true}},
	}} {
		t.Run(tc.name, func(t *testing.T) {
			var got seen
			status := tc.status
			if status == 0 {
				status = http.StatusOK
			}
			client := serve(t, status, tc.answer, &got)

			value, err := tc.call(client)
			if err != nil {
				t.Fatalf("call: %v", err)
			}
			if got.method != tc.method || got.path != tc.path {
				t.Errorf("sent %s %s, want %s %s", got.method, got.path, tc.method, tc.path)
			}
			if got.query != tc.query {
				t.Errorf("query = %q, want %q", got.query, tc.query)
			}
			if got.body != tc.body {
				t.Errorf("body = %q, want %q", got.body, tc.body)
			}
			// One header, never two. Two credentials on one request are two
			// answers to who is calling.
			if want := []string{"Bearer tok"}; !reflect.DeepEqual(got.auth, want) {
				t.Errorf("Authorization = %q, want %q", got.auth, want)
			}
			if !reflect.DeepEqual(value, tc.want) {
				t.Errorf("value =\n  %#v\nwant\n  %#v", value, tc.want)
			}
		})
	}
}

// A refusal is an answer. These are the arms a caller reads instead of catching,
// through the capabilities that produce them.
func TestRefusalsAreAnswers(t *testing.T) {
	ctx := context.Background()

	t.Run("denied on money, RFC 9457", func(t *testing.T) {
		var got seen
		client := serve(t, http.StatusPaymentRequired,
			`{"type":"about:blank","title":"Payment Required","status":402,"detail":"Add credits at console.hanzo.ai","code":"insufficient_balance"}`, &got)

		answer := client.Search.Find(ctx, "q3 incident postmortems", search.Opts{})
		_, err := answer.Value()

		var denied *Denied
		if !errors.As(err, &denied) {
			t.Fatalf("err = %v (%T), want *Denied", err, err)
		}
		if denied.Code != "insufficient_balance" {
			t.Errorf("code = %q, want insufficient_balance", denied.Code)
		}
		if denied.Reason != "Add credits at console.hanzo.ai" {
			t.Errorf("reason = %q", denied.Reason)
		}
		if answer.Request != request {
			t.Errorf("request = %q, want %q — every arm names the row", answer.Request, request)
		}
	})

	t.Run("denied on money, the second 402 body", func(t *testing.T) {
		var got seen
		client := serve(t, http.StatusPaymentRequired,
			`{"error":"payment_required","product":"graph","reason":"unpaid","message":"no active subscription for graph and no prepaid credit","cure":[{"kind":"subscribe","url":"/v1/plan"},{"kind":"credit","url":"/v1/billing"}]}`, &got)

		_, err := client.Graph.Assert(ctx, []graph.Fact{{Entity: "e", Relation: "r", Value: "v", At: at}}).Value()

		var denied *Denied
		if !errors.As(err, &denied) {
			t.Fatalf("err = %v (%T), want *Denied", err, err)
		}
		if denied.Code != "payment_required" || denied.Product != "graph" {
			t.Errorf("denied = %+v, want payment_required for graph", denied)
		}
		want := []Cure{{Kind: "subscribe", URL: "/v1/plan"}, {Kind: "credit", URL: "/v1/billing"}}
		if !reflect.DeepEqual(denied.Cures, want) {
			t.Errorf("cures = %+v, want %+v", denied.Cures, want)
		}
	})

	t.Run("denied by policy", func(t *testing.T) {
		var got seen
		client := serve(t, http.StatusForbidden,
			`{"type":"about:blank","title":"Forbidden","status":403,"detail":"clause graph.write refuses usr_7","code":"policy_denied"}`, &got)

		_, err := client.KB.Put(ctx, kb.Doc{Kind: kb.Page, Name: "runbook", Title: "Runbook"}).Value()

		var denied *Denied
		if !errors.As(err, &denied) {
			t.Fatalf("err = %v (%T), want *Denied", err, err)
		}
		if denied.Code != "policy_denied" {
			t.Errorf("code = %q, want policy_denied", denied.Code)
		}
	})

	t.Run("held for a person", func(t *testing.T) {
		var got seen
		client := serve(t, http.StatusAccepted,
			`{"status":"held","id":"apr_9","clause":"graph.write","reason":"a reviewer was asked"}`, &got)

		answer := client.Graph.Assert(ctx, []graph.Fact{{Entity: "e", Relation: "r", Value: "v", At: at}})
		_, err := answer.Value()

		var held *Held
		if !errors.As(err, &held) {
			t.Fatalf("err = %v (%T), want *Held", err, err)
		}
		if held.ID != "apr_9" || held.Clause != "graph.write" {
			t.Errorf("held = %+v", held)
		}
		if answer.Request != request {
			t.Errorf("request = %q, want %q", answer.Request, request)
		}
	})

	// An absent principal is not a decision, so it is not a refusal to read.
	// Cloud spells it 403 forbidden today, which is why the code matters.
	t.Run("no principal is a fault, not a denial", func(t *testing.T) {
		var got seen
		client := serve(t, http.StatusForbidden,
			`{"type":"about:blank","title":"Forbidden","status":403,"detail":"a validated principal is required","code":"forbidden"}`, &got)

		_, err := client.Search.Find(ctx, "anything", search.Opts{}).Value()

		var denied *Denied
		if errors.As(err, &denied) {
			t.Fatalf("a bare forbidden read as a denial: %+v", denied)
		}
		var fault *Fault
		if !errors.As(err, &fault) {
			t.Fatalf("err = %v (%T), want *Fault", err, err)
		}
		if fault.Status != http.StatusForbidden || fault.Request != request {
			t.Errorf("fault = %+v", fault)
		}
	})

	// The batch is accounted for, member by member. A count that does not add
	// up is a transport fault, not an answer.
	t.Run("a batch that does not add up", func(t *testing.T) {
		var got seen
		client := serve(t, http.StatusOK, `{"recorded":1,"duplicate":0,"refused":0}`, &got)

		_, err := client.Graph.Assert(ctx, []graph.Fact{
			{Entity: "a", Relation: "r", Value: "v", At: at},
			{Entity: "b", Relation: "r", Value: "v", At: at},
		}).Value()
		if err == nil {
			t.Fatal("want an error when the counts do not account for the batch")
		}
		var denied *Denied
		if errors.As(err, &denied) {
			t.Fatal("an inconsistent count is not a refusal")
		}
	})
}

// Find searches by text. Entity and Value name a key, and a text search has
// none, so a filter carrying either is refused rather than dropped.
func TestGraphFindRefusesAKey(t *testing.T) {
	var got seen
	client := serve(t, http.StatusOK, `{"assertions":[]}`, &got)

	if _, err := client.Graph.Find(context.Background(), "core", graph.Filter{Entity: "svc/api"}); err == nil {
		t.Fatal("want an error for a key on a text search")
	}
	if got.path != "" {
		t.Errorf("the call went out as %s %s; it should not have left", got.method, got.path)
	}
}

// All walks every page, stopping when the running count reaches the total.
func TestAuditAllWalksPages(t *testing.T) {
	pages := []string{
		`{"total":3,"data":[{"seq":1,"requestId":"req-1"},{"seq":2,"requestId":"req-2"}]}`,
		`{"total":3,"data":[{"seq":3,"requestId":"req-1"}]}`,
	}
	var asked []string
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/iam/oauth/token", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"access_token":"tok","expires_in":3600}`))
	})
	mux.HandleFunc("/v1/audit", func(w http.ResponseWriter, r *http.Request) {
		asked = append(asked, r.URL.RawQuery)
		w.Header().Set("Content-Type", "application/json")
		if len(asked) <= len(pages) {
			w.Write([]byte(pages[len(asked)-1]))
			return
		}
		w.Write([]byte(`{"total":3,"data":[]}`))
	})
	srv := httptest.NewServer(mux)
	t.Cleanup(srv.Close)
	client := New(Options{ID: "cid", Secret: "sec", Base: srv.URL, Issuer: srv.URL, Resource: srv.URL})

	var seqs []int64
	for event, err := range client.Audit.All(context.Background(), audit.Filter{Size: 2}) {
		if err != nil {
			t.Fatalf("all: %v", err)
		}
		seqs = append(seqs, event.Seq)
	}
	if want := []int64{1, 2, 3}; !reflect.DeepEqual(seqs, want) {
		t.Errorf("seqs = %v, want %v", seqs, want)
	}
	if want := []string{"p=1&pageSize=2", "p=2&pageSize=2"}; !reflect.DeepEqual(asked, want) {
		t.Errorf("asked %v, want %v", asked, want)
	}

	// A Request filter narrows client-side, so the server's total cannot end
	// the walk: it runs until a page comes back empty.
	asked = nil
	seqs = nil
	for event, err := range client.Audit.All(context.Background(), audit.Filter{Size: 2, Request: "req-1"}) {
		if err != nil {
			t.Fatalf("all: %v", err)
		}
		seqs = append(seqs, event.Seq)
	}
	if want := []int64{1, 3}; !reflect.DeepEqual(seqs, want) {
		t.Errorf("seqs = %v, want %v — only the rows one request produced", seqs, want)
	}
	if len(asked) != 3 {
		t.Errorf("asked %d pages, want 3 — the walk ends on an empty page", len(asked))
	}

	// A listing that publishes no total cannot end the walk on one, so it runs
	// to the empty page rather than truncating at the first.
	asked = nil
	seqs = nil
	pages = []string{`{"data":[{"seq":1}]}`, `{"data":[{"seq":2}]}`}
	for event, err := range client.Audit.All(context.Background(), audit.Filter{Size: 1}) {
		if err != nil {
			t.Fatalf("all: %v", err)
		}
		seqs = append(seqs, event.Seq)
	}
	if want := []int64{1, 2}; !reflect.DeepEqual(seqs, want) {
		t.Errorf("seqs = %v, want %v", seqs, want)
	}
}

// value is Answer.Value as a table row uses it: the value alone, boxed.
func value[T any](a call.Answer[T]) (any, error) {
	v, err := a.Value()
	return v, err
}

func left(n int64) *int64         { return &n }
func when(t time.Time) *time.Time { return &t }
