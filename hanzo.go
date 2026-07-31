// Package hanzoai is the Hanzo API client for Go.
//
// Everything else in this package is generated from hanzoai/openapi
// `hanzo.yaml` by scripts/generate.sh. This file is not: it is the hand-written
// seam that turns the generated transport into an authenticated client.
package hanzoai

import "os"

// DefaultBaseURL is the Hanzo gateway. Every route is <base>/v1/<service>/*.
const DefaultBaseURL = "https://api.hanzo.ai"

// NewConfig returns a Configuration pointed at the Hanzo API and authenticated
// with apiKey. An empty apiKey reads HANZO_API_KEY from the environment; an
// empty HANZO_BASE_URL leaves the base URL at DefaultBaseURL.
//
// Auth is a default header rather than generated per-operation code on purpose.
// hanzo.yaml declares `security: [{bearerAuth: []}]` once at the document root,
// which OpenAPI applies to every operation, but openapi-generator 7.14.0's `go`
// generator only emits auth code for operations that carry their own explicit
// `security` block — 65 of 1519 here. Relying on that would ship a client that
// cannot authenticate. One default header covers every operation instead.
//
// Use NewConfig when a request needs more than credentials — an org scope, say:
//
//	cfg := hanzoai.NewConfig("")
//	cfg.AddDefaultHeader("X-Org-Id", org)
//	client := hanzoai.NewAPIClient(cfg)
func NewConfig(apiKey string) *Configuration {
	if apiKey == "" {
		apiKey = os.Getenv("HANZO_API_KEY")
	}
	cfg := NewConfiguration()
	if base := os.Getenv("HANZO_BASE_URL"); base != "" {
		cfg.Servers = ServerConfigurations{{URL: base}}
	}
	if apiKey != "" {
		cfg.AddDefaultHeader("Authorization", "Bearer "+apiKey)
	}
	return cfg
}

// NewClient returns a client for the Hanzo API, authenticated with apiKey. An
// empty apiKey reads HANZO_API_KEY from the environment.
//
//	client := hanzoai.NewClient("")
//	models, _, err := client.OpenAICompatibleAPI.AiListModels(ctx).Execute()
func NewClient(apiKey string) *APIClient {
	return NewAPIClient(NewConfig(apiKey))
}
