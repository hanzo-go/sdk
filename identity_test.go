package hanzoai

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
)

// estate stands up one server answering both of IAM's mints and the platform
// API, so a test can watch a whole round trip: client credentials in, an access
// token out, and — where a client acts as a subject — an act grant on top.
type estate struct {
	srv *httptest.Server

	oauth atomic.Int32 // client_credentials exchanges
	acts  atomic.Int32 // act grants
	calls atomic.Int32 // platform calls

	form    atomic.Value // the last exchange's form, parsed
	basic   atomic.Value // the clientId the last exchange presented
	actURL  atomic.Value // path and query of the last act grant
	actAuth atomic.Value // the credential the last act grant presented

	mu      sync.Mutex
	bearers [][]string // Authorization seen on each platform call, in order
	bodies  []string   // what each platform call carried, in order
}

// mints is what IAM answers, per exchange, so a test can make the second one
// differ from the first.
type mints struct {
	oauth func(n int32, w http.ResponseWriter)
	act   func(n int32, w http.ResponseWriter)
	api   func(n int32, w http.ResponseWriter, r *http.Request)
}

func newEstate(t *testing.T, m mints) *estate {
	t.Helper()
	e := &estate{}

	mux := http.NewServeMux()
	mux.HandleFunc("/v1/iam/oauth/token", func(w http.ResponseWriter, r *http.Request) {
		n := e.oauth.Add(1)
		if r.Method != http.MethodPost {
			t.Errorf("exchange method = %s, want POST", r.Method)
		}
		body, _ := io.ReadAll(r.Body)
		form, _ := url.ParseQuery(string(body))
		e.form.Store(form)
		id, _, _ := r.BasicAuth()
		e.basic.Store(id)
		w.Header().Set("Content-Type", "application/json")
		if m.oauth != nil {
			m.oauth(n, w)
			return
		}
		w.Write([]byte(`{"access_token":"tok-operator","expires_in":3600}`))
	})
	mux.HandleFunc("/v1/iam/tokens/issue", func(w http.ResponseWriter, r *http.Request) {
		n := e.acts.Add(1)
		if r.ContentLength > 0 {
			t.Errorf("act grant carried a body of %d bytes; the subject rides as the id query", r.ContentLength)
		}
		e.actURL.Store(r.URL.String())
		e.actAuth.Store(r.Header.Get("Authorization"))
		w.Header().Set("Content-Type", "application/json")
		if m.act != nil {
			m.act(n, w)
			return
		}
		w.Write([]byte(`{"accessToken":"tok-subject","expiresIn":3600}`))
	})
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n := e.calls.Add(1)
		body, _ := io.ReadAll(r.Body)
		e.mu.Lock()
		e.bearers = append(e.bearers, r.Header.Values("Authorization"))
		e.bodies = append(e.bodies, string(body))
		e.mu.Unlock()
		w.Header().Set("Content-Type", "application/json")
		if m.api != nil {
			m.api(n, w, r)
			return
		}
		w.Write([]byte(`{}`))
	})

	e.srv = httptest.NewServer(mux)
	t.Cleanup(e.srv.Close)
	return e
}

// client is an operator client pointed at the estate.
func (e *estate) client() *Client {
	return New(Options{ID: "cid", Secret: "shh", Base: e.srv.URL, Issuer: e.srv.URL, Resource: e.srv.URL})
}

// bearer is the single Authorization value the nth platform call carried, and
// the assertion that there was exactly one of them.
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

// keys is one ordinary read, the same one throughout.
func keys(c *Client) error {
	_, _, err := c.AccountAPI.GetAccountKeys(context.Background()).Execute()
	return err
}

func TestClientMintsFromItsOwnCredentials(t *testing.T) {
	e := newEstate(t, mints{})

	if err := keys(e.client()); err != nil {
		t.Fatalf("keys: %v", err)
	}

	form := e.form.Load().(url.Values)
	if got, want := form.Get("grant_type"), "client_credentials"; got != want {
		t.Errorf("grant_type = %q, want %q", got, want)
	}
	// RFC 8707: the token names what it may spend at and is useless elsewhere.
	if got, want := form.Get("resource"), e.srv.URL; got != want {
		t.Errorf("resource = %q, want %q", got, want)
	}
	// client_secret_basic, not a secret in the form.
	if got := e.basic.Load().(string); got != "cid" {
		t.Errorf("basic id = %q, want cid", got)
	}
	if form.Get("client_secret") != "" {
		t.Error("the secret rode in the form; it belongs in the Basic credential")
	}
	if got, want := e.bearer(t, 0), "Bearer tok-operator"; got != want {
		t.Errorf("Authorization = %q, want %q", got, want)
	}
}

// A client holding no credentials presents nothing rather than failing. The
// operations that take none still answer.
func TestClientWithoutCredentialsPresentsNothing(t *testing.T) {
	t.Setenv("HANZO_CLIENT_ID", "")
	t.Setenv("HANZO_CLIENT_SECRET", "")
	e := newEstate(t, mints{})

	client := New(Options{Base: e.srv.URL, Issuer: e.srv.URL})
	if err := keys(client); err != nil {
		t.Fatalf("keys: %v", err)
	}
	if got := e.oauth.Load(); got != 0 {
		t.Errorf("exchanges = %d, want 0 — there was nothing to exchange", got)
	}
	e.mu.Lock()
	defer e.mu.Unlock()
	if len(e.bearers[0]) != 0 {
		t.Errorf("Authorization = %q, want none", e.bearers[0])
	}
}

func TestTokenIsCachedToExpiry(t *testing.T) {
	e := newEstate(t, mints{})

	client := e.client()
	for i := range 3 {
		if err := keys(client); err != nil {
			t.Fatalf("call %d: %v", i, err)
		}
	}
	if got := e.oauth.Load(); got != 1 {
		t.Errorf("exchanges = %d, want 1 — the token is held to expiry", got)
	}
	if got := e.calls.Load(); got != 3 {
		t.Errorf("calls = %d, want 3", got)
	}
}

// A token IAM says lives one second is already inside the re-mint window, so
// the next call exchanges again rather than riding a token that expires in
// flight.
func TestShortLivedTokenIsReplacedBeforeItExpires(t *testing.T) {
	e := newEstate(t, mints{oauth: func(n int32, w http.ResponseWriter) {
		w.Write([]byte(`{"access_token":"tok","expires_in":1}`))
	}})

	client := e.client()
	for i := range 2 {
		if err := keys(client); err != nil {
			t.Fatalf("call %d: %v", i, err)
		}
	}
	if got := e.oauth.Load(); got != 2 {
		t.Errorf("exchanges = %d, want 2", got)
	}
}

func TestUnauthorizedMintsOnceAndReplays(t *testing.T) {
	e := newEstate(t, mints{
		oauth: func(n int32, w http.ResponseWriter) {
			if n == 1 {
				w.Write([]byte(`{"access_token":"stale","expires_in":3600}`))
				return
			}
			w.Write([]byte(`{"access_token":"fresh","expires_in":3600}`))
		},
		api: func(n int32, w http.ResponseWriter, r *http.Request) {
			if n == 1 {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte(`{"error":"token expired"}`))
				return
			}
			w.Write([]byte(`{"id":"agent_2"}`))
		},
	})

	name := "scribe"
	agent, _, err := e.client().AgentAPI.PostAgent(context.Background()).
		CreateAgentIn(CreateAgentIn{Name: &name}).Execute()
	if err != nil {
		t.Fatalf("PostAgent: %v", err)
	}
	if got := e.oauth.Load(); got != 2 {
		t.Errorf("exchanges = %d, want 2 — a 401 drops the token and mints once more", got)
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

// A second 401 is the server saying no, not a stale token.
func TestSecondUnauthorizedStops(t *testing.T) {
	e := newEstate(t, mints{api: func(n int32, w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"error":"nope"}`))
	}})

	_, res, err := e.client().AccountAPI.GetAccountKeys(context.Background()).Execute()
	var apiErr *GenericOpenAPIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("err = %v (%T), want *GenericOpenAPIError", err, err)
	}
	if res.StatusCode != http.StatusUnauthorized {
		t.Errorf("StatusCode = %d, want 401", res.StatusCode)
	}
	if got := e.calls.Load(); got != 2 {
		t.Errorf("calls = %d, want 2 — one original and exactly one replay", got)
	}
}

// A refused exchange names the identity that was refused, and no call goes out.
func TestRefusedExchangeSurfaces(t *testing.T) {
	e := newEstate(t, mints{oauth: func(n int32, w http.ResponseWriter) {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"error":"invalid_client","error_description":"unknown client"}`))
	}})

	err := keys(e.client())
	if err == nil {
		t.Fatal("want an error when IAM refuses the exchange")
	}
	if !strings.Contains(err.Error(), `"cid"`) || !strings.Contains(err.Error(), "invalid_client") {
		t.Errorf("error = %q, want the client and the refusal in it", err)
	}
	if got := e.calls.Load(); got != 0 {
		t.Errorf("calls = %d, want 0 — no token, no call", got)
	}
}

// Concurrent first calls exchange once between them, not once each.
func TestConcurrentCallsMintOnce(t *testing.T) {
	e := newEstate(t, mints{})

	client := e.client()
	var wg sync.WaitGroup
	for range 8 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := keys(client); err != nil {
				t.Error(err)
			}
		}()
	}
	wg.Wait()
	if got := e.oauth.Load(); got != 1 {
		t.Errorf("exchanges = %d, want 1", got)
	}
}

func TestAsScopesToASubject(t *testing.T) {
	e := newEstate(t, mints{})

	if err := keys(e.client().As("usr_7")); err != nil {
		t.Fatalf("keys: %v", err)
	}

	// The act grant is IAM's own path and the subject rides as the id query.
	if got, want := e.actURL.Load().(string), "/v1/iam/tokens/issue?id=usr_7"; got != want {
		t.Errorf("act grant URL = %q, want %q", got, want)
	}
	// IAM reads the grant off the operator's own minted token.
	if got, want := e.actAuth.Load().(string), "Bearer tok-operator"; got != want {
		t.Errorf("act grant credential = %q, want %q", got, want)
	}
	// The subject-bound token is the whole identity of the scoped call, and the
	// operator credential left with the scope.
	if got, want := e.bearer(t, 0), "Bearer tok-subject"; got != want {
		t.Errorf("scoped Authorization = %q, want %q", got, want)
	}
}

// An externalId is what an operator usually files a member under, and it can
// hold characters that must survive the query.
func TestAsEscapesAnExternalID(t *testing.T) {
	e := newEstate(t, mints{})

	if err := keys(e.client().As("acme/user@example.com")); err != nil {
		t.Fatalf("keys: %v", err)
	}
	if got, want := e.actURL.Load().(string), "/v1/iam/tokens/issue?id=acme%2Fuser%40example.com"; got != want {
		t.Errorf("act grant URL = %q, want %q", got, want)
	}
}

// A scoped client shares the operator's held token, so scoping costs one act
// grant and not a second exchange.
func TestScopingReusesTheOperatorToken(t *testing.T) {
	e := newEstate(t, mints{})

	client := e.client()
	if err := keys(client); err != nil {
		t.Fatalf("operator: %v", err)
	}
	if err := keys(client.As("usr_7")); err != nil {
		t.Fatalf("scoped: %v", err)
	}
	if got := e.oauth.Load(); got != 1 {
		t.Errorf("exchanges = %d, want 1", got)
	}
	if got := e.acts.Load(); got != 1 {
		t.Errorf("act grants = %d, want 1", got)
	}
}

// The six reach the same endpoint under the same credential as the generated
// surface, so a scoped client scopes all of it.
func TestScopingReachesTheSix(t *testing.T) {
	e := newEstate(t, mints{api: func(n int32, w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"plan":"pro","limit":20,"used":1}`))
	}})

	if _, err := e.client().As("usr_7").Budget.Left(context.Background()); err != nil {
		t.Fatalf("left: %v", err)
	}
	if got, want := e.bearer(t, 0), "Bearer tok-subject"; got != want {
		t.Errorf("Authorization = %q, want %q", got, want)
	}
}
