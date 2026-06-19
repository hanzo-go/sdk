package metering_test

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"

	"github.com/hanzoai/go-sdk/metering"
)

// commerceStub serves balance (GET) and records usage (POST), capturing the
// recorded amount so the test can assert post-request metering happened.
type commerceStub struct {
	available int64

	mu          sync.Mutex
	recordedAmt int64
	recordedCnt int
	recordedUsr string
}

func (s *commerceStub) server() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			s.mu.Lock()
			defer s.mu.Unlock()
			body, _ := io.ReadAll(r.Body)
			// crude parse: amount + user
			s.recordedCnt++
			s.recordedAmt = extractInt(body, `"amount":`)
			s.recordedUsr = extractStr(body, `"user":"`)
			w.WriteHeader(201)
			_, _ = io.WriteString(w, `{"transactionId":"tx","type":"withdraw"}`)
			return
		}
		// balance
		_, _ = io.WriteString(w, `{"available":`+itoa(s.available)+`}`)
	}))
}

func (s *commerceStub) records() (int, int64, string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.recordedCnt, s.recordedAmt, s.recordedUsr
}

func gatewayReq() *http.Request {
	r := httptest.NewRequest(http.MethodGet, "/search?q=foo", nil)
	r.Header.Set(metering.HeaderOrgID, "hanzo")
	r.Header.Set(metering.HeaderUserID, "alice")
	return r
}

func TestMiddleware_GatesAndRecords(t *testing.T) {
	stub := &commerceStub{available: 5000}
	srv := stub.server()
	defer srv.Close()

	c, _ := metering.New(metering.Config{BaseURL: srv.URL, Token: "t", Org: "hanzo"})

	handlerHit := false
	h := c.Middleware(metering.MiddlewareConfig{
		Provider: "search",
		Price: func(_ *http.Request, status int, _ metering.AuthInput) int64 {
			if status == http.StatusOK {
				return 7 // 7 cents per successful search
			}
			return 0
		},
	})(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlerHit = true
		w.WriteHeader(200)
		_, _ = w.Write([]byte("results"))
	}))

	rr := httptest.NewRecorder()
	h.ServeHTTP(rr, gatewayReq())

	if !handlerHit {
		t.Fatal("handler should have run (balance positive)")
	}
	if rr.Code != 200 {
		t.Fatalf("status = %d, want 200", rr.Code)
	}

	// Record happens async; poll briefly.
	waitFor(t, func() bool { n, _, _ := stub.records(); return n == 1 })
	n, amt, usr := stub.records()
	if n != 1 || amt != 7 {
		t.Fatalf("recorded (count=%d amount=%d), want (1, 7)", n, amt)
	}
	if usr != "hanzo/alice" {
		t.Errorf("recorded user = %q, want hanzo/alice (org/sub)", usr)
	}
}

func TestMiddleware_Denies402_WhenNoBalance(t *testing.T) {
	stub := &commerceStub{available: 0}
	srv := stub.server()
	defer srv.Close()

	c, _ := metering.New(metering.Config{BaseURL: srv.URL, Token: "t", Org: "hanzo"})
	handlerHit := false
	h := c.Middleware(metering.MiddlewareConfig{
		Provider: "search",
		Price:    func(*http.Request, int, metering.AuthInput) int64 { return 7 },
	})(http.HandlerFunc(func(http.ResponseWriter, *http.Request) { handlerHit = true }))

	rr := httptest.NewRecorder()
	h.ServeHTTP(rr, gatewayReq())

	if handlerHit {
		t.Fatal("handler must NOT run when balance is zero")
	}
	if rr.Code != http.StatusPaymentRequired {
		t.Fatalf("status = %d, want 402", rr.Code)
	}
	if n, _, _ := stub.records(); n != 0 {
		t.Fatalf("denied request must not record usage, got %d records", n)
	}
}

func TestMiddleware_FailClosed503_WhenCommerceDown(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
	}))
	defer srv.Close()

	c, _ := metering.New(metering.Config{BaseURL: srv.URL, Token: "t", Org: "hanzo"})
	h := c.Middleware(metering.MiddlewareConfig{
		Provider: "search",
		Price:    func(*http.Request, int, metering.AuthInput) int64 { return 7 },
	})(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Fatal("handler must not run when balance is unknown (fail-closed)")
	}))

	rr := httptest.NewRecorder()
	h.ServeHTTP(rr, gatewayReq())
	if rr.Code != http.StatusServiceUnavailable {
		t.Fatalf("status = %d, want 503 (fail-closed)", rr.Code)
	}
}

func TestMiddleware_Skip_Bypasses(t *testing.T) {
	stub := &commerceStub{available: 0} // would deny if gated
	srv := stub.server()
	defer srv.Close()

	c, _ := metering.New(metering.Config{BaseURL: srv.URL, Token: "t", Org: "hanzo"})
	handlerHit := false
	h := c.Middleware(metering.MiddlewareConfig{
		Provider: "search",
		Price:    func(*http.Request, int, metering.AuthInput) int64 { return 7 },
		Skip:     func(r *http.Request) bool { return r.URL.Path == "/healthz" },
	})(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlerHit = true
		w.WriteHeader(200)
	}))

	rr := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	h.ServeHTTP(rr, r)

	if !handlerHit {
		t.Fatal("skipped path must run handler without gating")
	}
	if n, _, _ := stub.records(); n != 0 {
		t.Fatal("skipped path must not record usage")
	}
}

func TestMiddleware_OnlyChargesSuccess(t *testing.T) {
	stub := &commerceStub{available: 5000}
	srv := stub.server()
	defer srv.Close()

	c, _ := metering.New(metering.Config{BaseURL: srv.URL, Token: "t", Org: "hanzo"})
	h := c.Middleware(metering.MiddlewareConfig{
		Provider: "search",
		Price: func(_ *http.Request, status int, _ metering.AuthInput) int64 {
			if status == http.StatusOK {
				return 7
			}
			return 0 // don't charge for failures
		},
	})(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError) // handler failed
	}))

	rr := httptest.NewRecorder()
	h.ServeHTTP(rr, gatewayReq())
	if rr.Code != http.StatusInternalServerError {
		t.Fatalf("status = %d, want 500", rr.Code)
	}
	// Give any (erroneous) async record a chance, then assert none happened.
	time.Sleep(50 * time.Millisecond)
	if n, _, _ := stub.records(); n != 0 {
		t.Fatalf("failed request must not be charged, got %d records", n)
	}
}

func TestIdentityFromGatewayHeaders(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	r.Header.Set(metering.HeaderOrgID, "zoo")
	r.Header.Set(metering.HeaderUserID, "bob")
	in := metering.IdentityFromGatewayHeaders(r)
	if in.User != "zoo/bob" {
		t.Errorf("User = %q, want zoo/bob", in.User)
	}
	if in.Org != "zoo" {
		t.Errorf("Org = %q, want zoo", in.Org)
	}
}

// ---- tiny helpers (avoid pulling strconv/fmt into hot asserts) ----

func waitFor(t *testing.T, cond func() bool) {
	t.Helper()
	deadline := time.Now().Add(2 * time.Second)
	for time.Now().Before(deadline) {
		if cond() {
			return
		}
		time.Sleep(5 * time.Millisecond)
	}
	t.Fatal("condition not met within deadline")
}

func itoa(n int64) string {
	if n == 0 {
		return "0"
	}
	neg := n < 0
	if neg {
		n = -n
	}
	var b [20]byte
	i := len(b)
	for n > 0 {
		i--
		b[i] = byte('0' + n%10)
		n /= 10
	}
	if neg {
		i--
		b[i] = '-'
	}
	return string(b[i:])
}

func extractInt(body []byte, key string) int64 {
	s := string(body)
	idx := indexOf(s, key)
	if idx < 0 {
		return 0
	}
	idx += len(key)
	var n int64
	for idx < len(s) && (s[idx] == ' ') {
		idx++
	}
	neg := false
	if idx < len(s) && s[idx] == '-' {
		neg = true
		idx++
	}
	for idx < len(s) && s[idx] >= '0' && s[idx] <= '9' {
		n = n*10 + int64(s[idx]-'0')
		idx++
	}
	if neg {
		return -n
	}
	return n
}

func extractStr(body []byte, key string) string {
	s := string(body)
	idx := indexOf(s, key)
	if idx < 0 {
		return ""
	}
	idx += len(key)
	end := idx
	for end < len(s) && s[end] != '"' {
		end++
	}
	return s[idx:end]
}

func indexOf(s, sub string) int {
	for i := 0; i+len(sub) <= len(s); i++ {
		if s[i:i+len(sub)] == sub {
			return i
		}
	}
	return -1
}

var _ = context.Background
