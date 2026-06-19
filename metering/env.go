package metering

import (
	"os"
	"strings"
)

// Environment variables a product reads to wire the metering Client. These are
// the canonical names the operator injects (the token from a KMS-backed
// secret); products MUST NOT invent their own. See gateway/auth_middleware.go
// (DefaultAuthConfig) for the same names on the gateway side.
const (
	// EnvBaseURL is the commerce service base URL.
	// Default: http://commerce.hanzo.svc.cluster.local:8001
	EnvBaseURL = "COMMERCE_URL"

	// EnvToken is the commerce service token (admin-scoped S2S). The operator
	// wires this from a KMS-backed secret; it is NEVER stored in plaintext in
	// the repo or image.
	EnvToken = "COMMERCE_SERVICE_TOKEN"

	// EnvOrg is the default tenant org slug (X-IAM-Org-Id) for S2S calls when a
	// request carries no org. Default: hanzo.
	EnvOrg = "COMMERCE_SERVICE_ORG"

	// EnvTierAware ("true") gates on the tier-aware effective balance
	// (prepaid + included plan allotment) instead of bare prepaid balance.
	EnvTierAware = "METERING_TIER_AWARE"

	// EnvFailOpen ("true") flips the gate to allow-on-error. Default is
	// fail-closed (deny), matching the gateway. Set only where availability
	// outranks billing.
	EnvFailOpen = "METERING_FAIL_OPEN"

	// EnvDisabled ("true") forces "not configured" mode regardless of
	// COMMERCE_URL — Authorize allows, Record is a no-op. For local dev.
	EnvDisabled = "METERING_DISABLED"
)

// DefaultBaseURL is the in-cluster commerce address. Matches the gateway's
// AUTH_BILLING_URL default so both gate on the same balance source.
const DefaultBaseURL = "http://commerce.hanzo.svc.cluster.local:8001"

// ConfigFromEnv builds a Config from the canonical environment variables. It
// applies the in-cluster commerce default and the fail-closed/usd defaults.
func ConfigFromEnv() Config {
	if envTrue(EnvDisabled) {
		return Config{} // not configured -> allow + no-op.
	}

	base := strings.TrimSpace(os.Getenv(EnvBaseURL))
	if base == "" {
		base = DefaultBaseURL
	}
	org := strings.TrimSpace(os.Getenv(EnvOrg))
	if org == "" {
		org = "hanzo"
	}

	return Config{
		BaseURL:   base,
		Token:     strings.TrimSpace(os.Getenv(EnvToken)),
		Org:       org,
		TierAware: envTrue(EnvTierAware),
		FailOpen:  envTrue(EnvFailOpen),
	}
}

// FromEnv builds a Client from the canonical environment variables. This is the
// one-liner products use at startup:
//
//	meter, _ := metering.FromEnv()
//	mux.Use(meter.Middleware(metering.MiddlewareConfig{Provider: "search", Price: priceSearch}))
func FromEnv() (*Client, error) {
	return New(ConfigFromEnv())
}

func envTrue(key string) bool {
	v := strings.ToLower(strings.TrimSpace(os.Getenv(key)))
	return v == "true" || v == "1" || v == "yes"
}
