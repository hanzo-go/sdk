package metering_test

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"sync"
	"testing"

	"github.com/hanzoai/go-sdk/metering"
)

// fakeCommerce records the last request and replies with a canned status+body.
type fakeCommerce struct {
	mu     sync.Mutex
	method string
	path   string
	query  url.Values
	auth   string
	org    string
	ctype  string
	body   []byte
	status int
	reply  string
}

func (f *fakeCommerce) handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		f.mu.Lock()
		defer f.mu.Unlock()
		f.method = r.Method
		f.path = r.URL.Path
		f.query = r.URL.Query()
		f.auth = r.Header.Get("Authorization")
		f.org = r.Header.Get("X-IAM-Org-Id")
		f.ctype = r.Header.Get("Content-Type")
		f.body, _ = io.ReadAll(r.Body)
		if f.status == 0 {
			f.status = 200
		}
		w.WriteHeader(f.status)
		_, _ = io.WriteString(w, f.reply)
	}
}

func newClient(t *testing.T, srv *httptest.Server, cfg metering.Config) *metering.Client {
	t.Helper()
	cfg.BaseURL = srv.URL
	if cfg.Token == "" {
		cfg.Token = "svc-token"
	}
	if cfg.Org == "" {
		cfg.Org = "hanzo"
	}
	c, err := metering.New(cfg)
	if err != nil {
		t.Fatalf("New: %v", err)
	}
	return c
}

func TestAuthorize_Allows_WhenAvailablePositive(t *testing.T) {
	fc := &fakeCommerce{reply: `{"user":"hanzo/alice","currency":"usd","balance":5000,"holds":0,"available":5000}`}
	srv := httptest.NewServer(fc.handler())
	defer srv.Close()

	c := newClient(t, srv, metering.Config{})
	if err := c.Authorize(context.Background(), metering.AuthInput{User: "hanzo/alice"}); err != nil {
		t.Fatalf("Authorize allowed should be nil, got %v", err)
	}

	// Verify the exact commerce contract.
	if fc.method != http.MethodGet {
		t.Errorf("method = %s, want GET", fc.method)
	}
	if fc.path != "/v1/billing/balance" {
		t.Errorf("path = %s, want /v1/billing/balance", fc.path)
	}
	if got := fc.query.Get("user"); got != "hanzo/alice" {
		t.Errorf("user query = %q, want hanzo/alice", got)
	}
	if got := fc.query.Get("currency"); got != "usd" {
		t.Errorf("currency query = %q, want usd", got)
	}
	if fc.auth != "Bearer svc-token" {
		t.Errorf("auth = %q, want Bearer svc-token", fc.auth)
	}
	if fc.org != "hanzo" {
		t.Errorf("X-IAM-Org-Id = %q, want hanzo", fc.org)
	}
}

func TestAuthorize_Denies_WhenAvailableZero(t *testing.T) {
	fc := &fakeCommerce{reply: `{"available":0}`}
	srv := httptest.NewServer(fc.handler())
	defer srv.Close()

	c := newClient(t, srv, metering.Config{})
	err := c.Authorize(context.Background(), metering.AuthInput{User: "hanzo/alice"})
	if err != metering.ErrInsufficientBalance {
		t.Fatalf("want ErrInsufficientBalance, got %v", err)
	}
}

func TestAuthorize_FailClosed_OnCommerceError(t *testing.T) {
	fc := &fakeCommerce{status: 500, reply: `boom`}
	srv := httptest.NewServer(fc.handler())
	defer srv.Close()

	c := newClient(t, srv, metering.Config{})
	err := c.Authorize(context.Background(), metering.AuthInput{User: "hanzo/alice"})
	if err == nil {
		t.Fatal("fail-closed: commerce 500 must deny, got nil")
	}
	if err == metering.ErrInsufficientBalance {
		t.Fatal("a 500 is 'unknown', not 'insufficient' — must be a connectivity error")
	}
}

func TestAuthorize_FailOpen_OnCommerceError(t *testing.T) {
	fc := &fakeCommerce{status: 503, reply: `down`}
	srv := httptest.NewServer(fc.handler())
	defer srv.Close()

	c := newClient(t, srv, metering.Config{FailOpen: true})
	if err := c.Authorize(context.Background(), metering.AuthInput{User: "hanzo/alice"}); err != nil {
		t.Fatalf("fail-open: commerce down must allow, got %v", err)
	}
}

func TestAuthorize_NotConfigured_Allows(t *testing.T) {
	c, err := metering.New(metering.Config{}) // no BaseURL
	if err != nil {
		t.Fatalf("New: %v", err)
	}
	if c.Enabled() {
		t.Fatal("client with no BaseURL should report Enabled()=false")
	}
	if err := c.Authorize(context.Background(), metering.AuthInput{User: "hanzo/alice"}); err != nil {
		t.Fatalf("not-configured Authorize must allow, got %v", err)
	}
}

func TestAuthorize_TierAware_UsesEffectiveAvailable(t *testing.T) {
	// Bare prepaid is 0 but the free-tier daily credit gives effective 100.
	fc := &fakeCommerce{reply: `{"user":"hanzo/alice","balance":{"currency":"usd","prepaidAvailable":0,"dailyRemaining":100,"effectiveAvailable":100}}`}
	srv := httptest.NewServer(fc.handler())
	defer srv.Close()

	c := newClient(t, srv, metering.Config{TierAware: true})
	if err := c.Authorize(context.Background(), metering.AuthInput{User: "hanzo/alice"}); err != nil {
		t.Fatalf("tier-aware allow (included allotment) should be nil, got %v", err)
	}
	if fc.path != "/v1/billing/tier" {
		t.Errorf("tier-aware must hit /v1/billing/tier, got %s", fc.path)
	}
	if fc.query.Get("user") != "hanzo/alice" {
		t.Errorf("tier user query = %q", fc.query.Get("user"))
	}
}

func TestAuthorize_TierAware_DeniesWhenEffectiveZero(t *testing.T) {
	fc := &fakeCommerce{reply: `{"balance":{"effectiveAvailable":0}}`}
	srv := httptest.NewServer(fc.handler())
	defer srv.Close()

	c := newClient(t, srv, metering.Config{TierAware: true})
	if err := c.Authorize(context.Background(), metering.AuthInput{User: "hanzo/alice"}); err != metering.ErrInsufficientBalance {
		t.Fatalf("tier-aware exhausted must deny with ErrInsufficientBalance, got %v", err)
	}
}

func TestAuthorize_PerCallOrgOverride(t *testing.T) {
	fc := &fakeCommerce{reply: `{"available":1}`}
	srv := httptest.NewServer(fc.handler())
	defer srv.Close()

	c := newClient(t, srv, metering.Config{}) // default org hanzo
	_ = c.Authorize(context.Background(), metering.AuthInput{User: "zoo/bob", Org: "zoo"})
	if fc.org != "zoo" {
		t.Errorf("per-call org override: X-IAM-Org-Id = %q, want zoo", fc.org)
	}
}

func TestRecord_PostsCanonicalPayload(t *testing.T) {
	fc := &fakeCommerce{status: 201, reply: `{"transactionId":"tx_123","user":"hanzo/alice","amount":250,"currency":"usd","type":"withdraw"}`}
	srv := httptest.NewServer(fc.handler())
	defer srv.Close()

	c := newClient(t, srv, metering.Config{})
	res, err := c.Record(context.Background(), metering.Usage{
		User:        "hanzo/alice",
		AmountCents: 250,
		Provider:    "search",
		RequestID:   "req-9",
		Status:      "success",
	})
	if err != nil {
		t.Fatalf("Record: %v", err)
	}
	if res == nil || res.TransactionID != "tx_123" || res.Amount != 250 {
		t.Fatalf("unexpected RecordResult: %+v", res)
	}

	if fc.method != http.MethodPost {
		t.Errorf("method = %s, want POST", fc.method)
	}
	if fc.path != "/v1/billing/usage" {
		t.Errorf("path = %s, want /v1/billing/usage", fc.path)
	}
	if fc.ctype != "application/json" {
		t.Errorf("content-type = %q", fc.ctype)
	}
	if fc.auth != "Bearer svc-token" {
		t.Errorf("auth = %q", fc.auth)
	}
	if fc.org != "hanzo" {
		t.Errorf("X-IAM-Org-Id = %q, want hanzo", fc.org)
	}

	// Verify the JSON body matches commerce's usageRequest field names.
	var body map[string]any
	if err := json.Unmarshal(fc.body, &body); err != nil {
		t.Fatalf("decode body: %v", err)
	}
	if body["user"] != "hanzo/alice" {
		t.Errorf("body.user = %v", body["user"])
	}
	if body["amount"].(float64) != 250 {
		t.Errorf("body.amount = %v, want 250", body["amount"])
	}
	if body["currency"] != "usd" {
		t.Errorf("body.currency = %v, want usd (defaulted)", body["currency"])
	}
	if body["provider"] != "search" {
		t.Errorf("body.provider = %v, want search", body["provider"])
	}
	if body["requestId"] != "req-9" {
		t.Errorf("body.requestId = %v", body["requestId"])
	}
	// Org must NOT be in the body (it travels via the header).
	if _, ok := body["Org"]; ok {
		t.Error("Org leaked into the JSON body; it must be a header only")
	}
}

func TestRecord_ZeroAmount_IsNoOp(t *testing.T) {
	called := false
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		called = true
	}))
	defer srv.Close()

	c := newClient(t, srv, metering.Config{})
	res, err := c.Record(context.Background(), metering.Usage{User: "hanzo/alice", AmountCents: 0})
	if err != nil || res != nil {
		t.Fatalf("zero-amount Record should be (nil,nil), got (%v,%v)", res, err)
	}
	if called {
		t.Fatal("zero-amount Record must not call commerce")
	}
}

func TestRecord_NotConfigured_IsNoOp(t *testing.T) {
	c, _ := metering.New(metering.Config{})
	res, err := c.Record(context.Background(), metering.Usage{User: "hanzo/alice", AmountCents: 100})
	if err != nil || res != nil {
		t.Fatalf("not-configured Record should be (nil,nil), got (%v,%v)", res, err)
	}
}

func TestNew_RejectsBadURL(t *testing.T) {
	if _, err := metering.New(metering.Config{BaseURL: "://bad"}); err == nil {
		t.Fatal("expected error for unparseable BaseURL")
	}
}

func TestConfigFromEnv_Defaults(t *testing.T) {
	t.Setenv(metering.EnvBaseURL, "")
	t.Setenv(metering.EnvOrg, "")
	t.Setenv(metering.EnvDisabled, "")
	cfg := metering.ConfigFromEnv()
	if cfg.BaseURL != metering.DefaultBaseURL {
		t.Errorf("default BaseURL = %q, want %q", cfg.BaseURL, metering.DefaultBaseURL)
	}
	if cfg.Org != "hanzo" {
		t.Errorf("default Org = %q, want hanzo", cfg.Org)
	}
	if cfg.FailOpen {
		t.Error("default must be fail-closed")
	}
}

func TestConfigFromEnv_Disabled(t *testing.T) {
	t.Setenv(metering.EnvDisabled, "true")
	t.Setenv(metering.EnvBaseURL, "http://commerce:8001")
	cfg := metering.ConfigFromEnv()
	if cfg.BaseURL != "" {
		t.Errorf("METERING_DISABLED must yield empty BaseURL, got %q", cfg.BaseURL)
	}
}

func TestConfigFromEnv_ReadsToken(t *testing.T) {
	t.Setenv(metering.EnvToken, "kms-sourced-token")
	t.Setenv(metering.EnvTierAware, "true")
	cfg := metering.ConfigFromEnv()
	if cfg.Token != "kms-sourced-token" {
		t.Errorf("token = %q", cfg.Token)
	}
	if !cfg.TierAware {
		t.Error("METERING_TIER_AWARE=true should set TierAware")
	}
}

// TestContractMatchesGateway pins the wire format the gateway already uses, so
// this client and the gateway gate on the identical balance source.
func TestContractMatchesGateway(t *testing.T) {
	var gotURL string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotURL = r.URL.String()
		_, _ = io.WriteString(w, `{"available":1}`)
	}))
	defer srv.Close()

	c := newClient(t, srv, metering.Config{})
	_ = c.Authorize(context.Background(), metering.AuthInput{User: "hanzo/alice"})

	// Gateway: GET {base}/v1/billing/balance?user=hanzo%2Falice&currency=usd
	if !strings.HasPrefix(gotURL, "/v1/billing/balance?") {
		t.Fatalf("URL %q must start with /v1/billing/balance?", gotURL)
	}
	if !strings.Contains(gotURL, "user=hanzo%2Falice") {
		t.Errorf("URL %q must url-encode the user as hanzo%%2Falice", gotURL)
	}
	if !strings.Contains(gotURL, "currency=usd") {
		t.Errorf("URL %q must carry currency=usd", gotURL)
	}
}
