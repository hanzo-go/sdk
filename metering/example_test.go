package metering_test

import (
	"net/http"

	"github.com/hanzoai/go-sdk/metering"
)

// Example shows the ONE way a non-LLM product opts into pay-for-everything:
// build the client from the operator-wired env vars and wrap the handler. Every
// request is then balance-gated (fail-closed) and metered to commerce.
func Example() {
	// Token comes from a KMS-backed secret via COMMERCE_SERVICE_TOKEN — never
	// plaintext. COMMERCE_URL defaults to the in-cluster commerce address.
	meter, err := metering.FromEnv()
	if err != nil {
		panic(err)
	}

	// Price by outcome: charge a flat 5 cents per successful search request.
	priceSearch := func(r *http.Request, status int, in metering.AuthInput) int64 {
		if status >= 200 && status < 300 {
			return 5
		}
		return 0
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/v1/search", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"results":[]}`))
	})

	// One wrap meters the whole surface.
	handler := meter.Middleware(metering.MiddlewareConfig{
		Provider: "search",
		Price:    priceSearch,
		Skip:     func(r *http.Request) bool { return r.URL.Path == "/healthz" },
	})(mux)

	_ = handler // http.ListenAndServe(":8080", handler)
}

// Example_imperative shows direct use without middleware — for products that
// price per unit of work discovered during the request (e.g. functions billed
// by execution time), gating up front and recording the measured cost after.
func Example_imperative() {
	meter, _ := metering.FromEnv()

	handle := func(w http.ResponseWriter, r *http.Request) {
		in := metering.IdentityFromGatewayHeaders(r)

		// Pre-request gate (fail-closed).
		if err := meter.Authorize(r.Context(), in); err != nil {
			if err == metering.ErrInsufficientBalance {
				http.Error(w, "insufficient balance", http.StatusPaymentRequired)
			} else {
				http.Error(w, "billing unavailable", http.StatusServiceUnavailable)
			}
			return
		}

		// ... do the work, measuring cost ...
		costCents := int64(12)

		w.WriteHeader(http.StatusOK)

		// Post-request record (best-effort).
		_, _ = meter.Record(r.Context(), metering.Usage{
			User:        in.User,
			Org:         in.Org,
			AmountCents: costCents,
			Provider:    "functions",
			RequestID:   r.Header.Get("X-Request-Id"),
			Status:      "success",
		})
	}
	_ = handle
}
