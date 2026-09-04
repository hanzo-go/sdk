package hanzoai

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
)

// estate stands up one server answering both IAM's mint and the platform API,
// so a test can watch the whole act-grant round trip.
type estate struct {
	srv *httptest.Server

	mints   atomic.Int32
	calls   atomic.Int32
	mintURL atomic.Value // path and query of the last mint
	mintKey atomic.Value // credential seen on the last mint

	mu      sync.Mutex
	bearers [][]string // Authorization seen on each API call, in order
	bodies  []string   // what each API call carried, in order
}

func newEstate(t *testing.T, issue func(n int32, w http.ResponseWriter), api func(n int32, w http.ResponseWriter, r *http.Request)) *estate {
	t.Helper()
	e := &estate{}
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/iam/tokens/issue", func(w http.ResponseWriter, r *http.Request) {
		n := e.mints.Add(1)
		if r.Method != http.MethodPost {
			t.Errorf("mint method = %s, want POST", r.Method)
		}
		if r.ContentLength > 0 {
			t.Errorf("mint carried a body of %d bytes; the target rides as the id query", r.ContentLength)
		}
		e.mintURL.Store(r.URL.String())
		e.mintKey.Store(r.Header.Get("Authorization"))
		w.Header().Set("Content-Type", "application/json")
		issue(n, w)
	})
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n := e.calls.Add(1)
		body, _ := io.ReadAll(r.Body)
		e.mu.Lock()
		e.bearers = append(e.bearers, r.Header.Values("Authorization"))
		e.bodies = append(e.bodies, string(body))
		e.mu.Unlock()
		w.Header().Set("Content-Type", "application/json")
		api(n, w, r)
	})
	e.srv = httptest.NewServer(mux)
	t.Cleanup(e.srv.Close)
	t.Setenv("HANZO_BASE_URL", e.srv.URL)
	t.Setenv("HANZO_ISSUER_URL", e.srv.URL)
	return e
}

// create is one gated operation, the same one throughout: the document requires
// a body on it, so the tests state one rather than each restating why.
func create(c *APIClient) AgentsAPIPostAgentsRequest {
	name := "scribe"
	return c.AgentsAPI.PostAgents(context.Background()).CreateAgentIn(CreateAgentIn{Name: &name})
}

// scoped builds an operator client pointed at the estate and acts as subject.
func (e *estate) scoped(subject string) *APIClient {
	return NewClient("hk-operator").As(subject)
}

// bearer is the single Authorization value the nth API call carried, and the
// assertion that there was exactly one of them.
func (e *estate) bearer(t *testing.T, n int) string {
	t.Helper()
	e.mu.Lock()
	defer e.mu.Unlock()
	if n >= len(e.bearers) {
		t.Fatalf("call %d never arrived; %d did", n, len(e.bearers))
	}
	got := e.bearers[n]
	if len(got) != 1 {
		t.Fatalf("call %d carried Authorization %q, want exactly one — two credentials on one request is two answers to who is calling", n, got)
	}
	return got[0]
}

func TestAsMintsFromIamAndScopesTheCall(t *testing.T) {
	e := newEstate(t,
		func(n int32, w http.ResponseWriter) {
			w.Write([]byte(`{"accessToken":"tok-abc","expiresIn":3600}`))
		},
		func(n int32, w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(`{"keys":[{"prefix":"sk-live"}]}`))
		})

	keys, _, err := e.scoped("usr_7").AccountAPI.GetAccountKeys(context.Background()).Execute()
	if err != nil {
		t.Fatalf("GetAccountKeys: %v", err)
	}

	// The mint is IAM's own path, and the subject rides as the id query.
	if got, want := e.mintURL.Load().(string), "/v1/iam/tokens/issue?id=usr_7"; got != want {
		t.Errorf("mint URL = %q, want %q", got, want)
	}
	// IAM reads a bearer, and the operator credential is already spelled that
	// way here — so it travels verbatim and nothing else travels with it.
	if got, want := e.mintKey.Load().(string), "Bearer hk-operator"; got != want {
		t.Errorf("mint credential = %q, want %q", got, want)
	}
	// camelCase accessToken became the bearer on the scoped call, and the
	// operator credential left with it.
	if got, want := e.bearer(t, 0), "Bearer tok-abc"; got != want {
		t.Errorf("scoped Authorization = %q, want %q", got, want)
	}
	if len(keys.Keys) != 1 || keys.Keys[0].GetPrefix() != "sk-live" {
		t.Errorf("keys = %+v, want the one the estate answered with", keys.Keys)
	}
}

// An externalId is what an operator usually files a member under, and it can
// hold characters that must survive the query.
func TestAsPassesExternalIDAsIDQuery(t *testing.T) {
	e := newEstate(t,
		func(n int32, w http.ResponseWriter) { w.Write([]byte(`{"accessToken":"t","expiresIn":60}`)) },
		func(n int32, w http.ResponseWriter, r *http.Request) { w.Write([]byte(`{}`)) })

	if _, _, err := e.scoped("acme/user@example.com").AccountAPI.GetAccountKeys(context.Background()).Execute(); err != nil {
		t.Fatalf("GetAccountKeys: %v", err)
	}
	if got, want := e.mintURL.Load().(string), "/v1/iam/tokens/issue?id=acme%2Fuser%40example.com"; got != want {
		t.Errorf("mint URL = %q, want %q", got, want)
	}
}

func TestGrantCachesTokenToExpiry(t *testing.T) {
	e := newEstate(t,
		func(n int32, w http.ResponseWriter) { w.Write([]byte(`{"accessToken":"tok-1","expiresIn":3600}`)) },
		func(n int32, w http.ResponseWriter, r *http.Request) { w.Write([]byte(`{}`)) })

	c := e.scoped("usr_7")
	for i := 0; i < 3; i++ {
		if _, _, err := c.AccountAPI.GetAccountKeys(context.Background()).Execute(); err != nil {
			t.Fatalf("call %d: %v", i, err)
		}
	}
	if got := e.mints.Load(); got != 1 {
		t.Errorf("mints = %d, want 1 — the token should be cached to expiry", got)
	}
	if got := e.calls.Load(); got != 3 {
		t.Errorf("api calls = %d, want 3", got)
	}
}

// expiresIn is honoured: a token that IAM says lives one second is already
// inside the re-mint skew, so the next call mints again.
func TestGrantReMintsWhenExpiresInIsShort(t *testing.T) {
	e := newEstate(t,
		func(n int32, w http.ResponseWriter) { w.Write([]byte(`{"accessToken":"tok","expiresIn":1}`)) },
		func(n int32, w http.ResponseWriter, r *http.Request) { w.Write([]byte(`{}`)) })

	c := e.scoped("usr_7")
	for i := 0; i < 2; i++ {
		if _, _, err := c.AccountAPI.GetAccountKeys(context.Background()).Execute(); err != nil {
			t.Fatalf("call %d: %v", i, err)
		}
	}
	if got := e.mints.Load(); got != 2 {
		t.Errorf("mints = %d, want 2", got)
	}
}

func TestGrantReMintsOnUnauthorized(t *testing.T) {
	e := newEstate(t,
		func(n int32, w http.ResponseWriter) {
			if n == 1 {
				w.Write([]byte(`{"accessToken":"stale","expiresIn":3600}`))
				return
			}
			w.Write([]byte(`{"accessToken":"fresh","expiresIn":3600}`))
		},
		func(n int32, w http.ResponseWriter, r *http.Request) {
			if n == 1 {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte(`{"error":"token expired"}`))
				return
			}
			w.Write([]byte(`{"id":"agent_2"}`))
		})

	agent, _, err := create(e.scoped("usr_7")).Execute()
	if err != nil {
		t.Fatalf("PostAgents: %v", err)
	}
	if got := e.mints.Load(); got != 2 {
		t.Errorf("mints = %d, want 2 — a 401 should drop the token and mint once more", got)
	}
	if got, want := e.bearer(t, 0), "Bearer stale"; got != want {
		t.Errorf("first bearer = %q, want %q", got, want)
	}
	if got, want := e.bearer(t, 1), "Bearer fresh"; got != want {
		t.Errorf("replayed bearer = %q, want %q", got, want)
	}
	e.mu.Lock()
	replayed := e.bodies[1]
	e.mu.Unlock()
	if !strings.Contains(replayed, `"name":"scribe"`) {
		t.Errorf("replayed body = %q, want the body the first attempt carried", replayed)
	}
	if agent.GetId() != "agent_2" {
		t.Errorf("agent = %+v, want agent_2 from the replay", agent)
	}
}

// A second 401 is the server saying no. The client stops and surfaces it.
func TestGrantStopsAfterOneReMint(t *testing.T) {
	e := newEstate(t,
		func(n int32, w http.ResponseWriter) { w.Write([]byte(`{"accessToken":"t","expiresIn":3600}`)) },
		func(n int32, w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(`{"error":"nope"}`))
		})

	_, res, err := e.scoped("usr_7").AccountAPI.GetAccountKeys(context.Background()).Execute()
	var apiErr *GenericOpenAPIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("err = %v (%T), want *GenericOpenAPIError", err, err)
	}
	if res.StatusCode != http.StatusUnauthorized {
		t.Errorf("StatusCode = %d, want 401", res.StatusCode)
	}
	if got := e.calls.Load(); got != 2 {
		t.Errorf("api calls = %d, want 2 — one original and exactly one replay", got)
	}
}

func TestMintFailureSurfacesAsTypedError(t *testing.T) {
	e := newEstate(t,
		func(n int32, w http.ResponseWriter) {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte(`{"error":"no act grant on this key"}`))
		},
		func(n int32, w http.ResponseWriter, r *http.Request) { w.Write([]byte(`{}`)) })

	_, _, err := e.scoped("usr_7").AccountAPI.GetAccountKeys(context.Background()).Execute()
	var apiErr *GenericOpenAPIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("err = %v (%T), want *GenericOpenAPIError", err, err)
	}
	if !strings.Contains(apiErr.Error(), "403") {
		t.Errorf("error = %q, want the refusal's status", apiErr.Error())
	}
	if got, want := string(apiErr.Body()), `{"error":"no act grant on this key"}`; got != want {
		t.Errorf("body = %q, want %q", got, want)
	}
	if got := e.calls.Load(); got != 0 {
		t.Errorf("api calls = %d, want 0 — no token, no call", got)
	}
	// A refused mint is an answer, not a network blip. It must not be retried.
	if got := e.mints.Load(); got != 1 {
		t.Errorf("mints = %d, want 1", got)
	}
}

// Concurrent first calls mint once between them, not once each.
func TestGrantMintsOnceUnderConcurrency(t *testing.T) {
	e := newEstate(t,
		func(n int32, w http.ResponseWriter) { w.Write([]byte(`{"accessToken":"t","expiresIn":3600}`)) },
		func(n int32, w http.ResponseWriter, r *http.Request) { w.Write([]byte(`{}`)) })

	c := e.scoped("usr_7")
	var wg sync.WaitGroup
	for i := 0; i < 8; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if _, _, err := c.AccountAPI.GetAccountKeys(context.Background()).Execute(); err != nil {
				t.Error(err)
			}
		}()
	}
	wg.Wait()
	if got := e.mints.Load(); got != 1 {
		t.Errorf("mints = %d, want 1", got)
	}
}

func TestMintWithoutTokenIsAnError(t *testing.T) {
	e := newEstate(t,
		func(n int32, w http.ResponseWriter) { w.Write([]byte(`{"expiresIn":3600}`)) },
		func(n int32, w http.ResponseWriter, r *http.Request) { w.Write([]byte(`{}`)) })

	_, _, err := e.scoped("usr_7").AccountAPI.GetAccountKeys(context.Background()).Execute()
	if err == nil {
		t.Fatal("want an error when IAM issues no token")
	}
	if got := e.calls.Load(); got != 0 {
		t.Errorf("api calls = %d, want 0", got)
	}
}
