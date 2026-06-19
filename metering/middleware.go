package metering

import (
	"context"
	"net/http"
	"strings"
)

// Identity headers minted by the Hanzo gateway (the trust boundary). A product
// behind the gateway reads the caller identity from these — it never trusts a
// client-supplied value (the gateway strips them on ingress). See
// commerce/CLAUDE.md "Gateway Trust Headers".
const (
	HeaderUserID = "X-User-Id"
	HeaderOrgID  = "X-Org-Id"
)

// PriceFunc computes the cost (in cents) to record for a completed request.
// It is called AFTER the wrapped handler runs, with the captured status code
// and the per-request context, so it can price by outcome (e.g. charge only on
// success) and by work done (bytes, rows, units recorded on the context by the
// handler). Return 0 to record nothing.
type PriceFunc func(r *http.Request, status int, in AuthInput) int64

// MiddlewareConfig configures Middleware.
type MiddlewareConfig struct {
	// Price computes the per-request cost in cents. Required — without it the
	// middleware would gate but never charge, which is not metering.
	Price PriceFunc

	// Provider labels the recorded usage (e.g. "search", "functions"). It is
	// stored on the commerce transaction so spend can be attributed per product.
	Provider string

	// Identify extracts the billing identity (IAM user + org) from a request.
	// Defaults to IdentityFromGatewayHeaders, which reads the gateway-minted
	// X-User-Id / X-Org-Id headers.
	Identify func(*http.Request) AuthInput

	// Skip lets a request bypass metering entirely (health checks, public
	// paths). Returning true means: no gate, no record. Optional.
	Skip func(*http.Request) bool

	// OnDenied renders the response when Authorize denies. Defaults to a JSON
	// 402 (insufficient balance) / 503 (balance unknown, fail-closed) — the
	// same status mapping the gateway uses.
	OnDenied func(w http.ResponseWriter, r *http.Request, err error)

	// OnRecordError is invoked (best-effort, async) if recording usage fails.
	// Optional — typically wired to the product's logger/metrics. The request
	// has already succeeded; recording failure must not affect the response.
	OnRecordError func(r *http.Request, u Usage, err error)
}

// Middleware returns net/http middleware that gates every request on the
// caller's commerce balance (fail-closed by default) and records usage after a
// successful response. It is the ONE way a non-LLM product opts into
// pay-for-everything: wrap the handler once and every request is metered.
//
// It is plain net/http middleware (func(http.Handler) http.Handler), so it
// composes with the standard library, gorilla/mux (.Use), chi, and anything
// that speaks http.Handler — no per-framework variants.
func (c *Client) Middleware(cfg MiddlewareConfig) func(http.Handler) http.Handler {
	identify := cfg.Identify
	if identify == nil {
		identify = IdentityFromGatewayHeaders
	}
	onDenied := cfg.OnDenied
	if onDenied == nil {
		onDenied = defaultOnDenied
	}
	price := cfg.Price
	if price == nil {
		// A nil Price would gate-without-charging. Refuse silently to charge
		// nothing rather than panic, but this is a misconfiguration: callers
		// MUST supply Price. We treat it as "record nothing".
		price = func(*http.Request, int, AuthInput) int64 { return 0 }
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if cfg.Skip != nil && cfg.Skip(r) {
				next.ServeHTTP(w, r)
				return
			}

			in := identify(r)

			// Pre-request gate. When the client is not configured this always
			// allows (Authorize handles that), so a product can ship the wrap
			// before its billing is wired.
			if err := c.Authorize(r.Context(), in); err != nil {
				onDenied(w, r, err)
				return
			}

			// Run the handler, capturing the status code for outcome pricing.
			sw := &statusWriter{ResponseWriter: w, status: http.StatusOK}
			next.ServeHTTP(sw, r)

			// Post-request record. Decoupled from the response: priced by
			// outcome, recorded best-effort so billing never corrupts the
			// reply the user already received.
			cents := price(r, sw.status, in)
			if cents <= 0 || !c.Enabled() {
				return
			}
			u := Usage{
				User:        in.User,
				Org:         in.Org,
				Currency:    in.Currency,
				AmountCents: cents,
				Provider:    cfg.Provider,
				RequestID:   r.Header.Get("X-Request-Id"),
				Status:      statusLabel(sw.status),
				ClientIP:    clientIP(r),
			}
			c.recordAsync(r, u, cfg.OnRecordError)
		})
	}
}

// recordAsync records usage without blocking the response. The detach uses a
// background context so a client disconnect can't cancel the debit.
func (c *Client) recordAsync(r *http.Request, u Usage, onErr func(*http.Request, Usage, error)) {
	go func() {
		if _, err := c.Record(context.Background(), u); err != nil && onErr != nil {
			onErr(r, u, err)
		}
	}()
}

// IdentityFromGatewayHeaders builds an AuthInput from the gateway-minted
// identity headers. The user is "{org}/{sub}" (commerce's iam-user form) when
// both are present, else the bare sub.
func IdentityFromGatewayHeaders(r *http.Request) AuthInput {
	org := strings.TrimSpace(r.Header.Get(HeaderOrgID))
	sub := strings.TrimSpace(r.Header.Get(HeaderUserID))
	user := sub
	if org != "" && sub != "" {
		user = org + "/" + sub
	}
	return AuthInput{User: user, Org: org}
}

func defaultOnDenied(w http.ResponseWriter, _ *http.Request, err error) {
	w.Header().Set("Content-Type", "application/json")
	if err == ErrInsufficientBalance {
		w.WriteHeader(http.StatusPaymentRequired) // 402
		_, _ = w.Write([]byte(`{"error":{"message":"Insufficient balance. Please add credits at console.hanzo.ai","type":"billing_error","code":"insufficient_balance"}}`))
		return
	}
	// Balance unknown -> fail-closed -> 503 (commerce unreachable / not verifiable).
	w.WriteHeader(http.StatusServiceUnavailable) // 503
	_, _ = w.Write([]byte(`{"error":{"message":"Billing temporarily unavailable","type":"billing_error","code":"balance_unavailable"}}`))
}

func statusLabel(code int) string {
	if code >= 200 && code < 400 {
		return "success"
	}
	return "error"
}

func clientIP(r *http.Request) string {
	if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
		if i := strings.IndexByte(xff, ','); i > 0 {
			return strings.TrimSpace(xff[:i])
		}
		return strings.TrimSpace(xff)
	}
	return r.RemoteAddr
}

// statusWriter captures the response status code for outcome-based pricing.
type statusWriter struct {
	http.ResponseWriter
	status      int
	wroteHeader bool
}

func (s *statusWriter) WriteHeader(code int) {
	if !s.wroteHeader {
		s.status = code
		s.wroteHeader = true
	}
	s.ResponseWriter.WriteHeader(code)
}

func (s *statusWriter) Write(b []byte) (int, error) {
	if !s.wroteHeader {
		s.wroteHeader = true // implicit 200
	}
	return s.ResponseWriter.Write(b)
}

// Flush propagates flushing if the underlying writer supports it (streaming).
func (s *statusWriter) Flush() {
	if f, ok := s.ResponseWriter.(http.Flusher); ok {
		f.Flush()
	}
}
