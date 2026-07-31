package hanzoai

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewConfigAuthenticates(t *testing.T) {
	t.Setenv("HANZO_API_KEY", "from-env")
	t.Setenv("HANZO_BASE_URL", "")

	if got := NewConfig("explicit").DefaultHeader["Authorization"]; got != "Bearer explicit" {
		t.Errorf("explicit key: Authorization = %q, want %q", got, "Bearer explicit")
	}
	if got := NewConfig("").DefaultHeader["Authorization"]; got != "Bearer from-env" {
		t.Errorf("empty key: Authorization = %q, want %q", got, "Bearer from-env")
	}
	if got := NewConfig("").Servers[0].URL; got != DefaultBaseURL {
		t.Errorf("base URL = %q, want %q", got, DefaultBaseURL)
	}
}

func TestNewConfigHonoursBaseURLOverride(t *testing.T) {
	t.Setenv("HANZO_BASE_URL", "http://localhost:9999")
	if got := NewConfig("k").Servers[0].URL; got != "http://localhost:9999" {
		t.Errorf("base URL = %q, want override", got)
	}
}

// TestFlows pins the six canonical example flows to the routes hanzoai/openapi
// `flows.yaml` names — the manifest that makes "the same six examples in every
// SDK" a fact rather than six repos independently remembering to agree. If a
// regeneration moves an operation, this fails instead of the examples silently
// calling the wrong endpoint.
func TestFlows(t *testing.T) {
	ctx := context.Background()

	for _, tc := range []struct {
		flow, method, path string
		call               func(c *APIClient)
	}{
		{"hello", "GET", "/v1/bot/auth/me", func(c *APIClient) {
			c.AuthAPI.BotAuthMe(ctx).Execute()
		}},
		{"chat", "POST", "/v1/chat/completions", func(c *APIClient) {
			c.OpenAICompatibleAPI.AiCreateChatCompletion(ctx).AiChatCompletionRequest(
				AiChatCompletionRequest{Model: "zen-1", Messages: []AiChatMessage{{Role: "user"}}},
			).Execute()
		}},
		{"money", "GET", "/v1/billing/balance", func(c *APIClient) {
			c.BillingAPI.CloudGetV1BillingBalance(ctx).Execute()
		}},
		{"store", "POST", "/v1/kv", func(c *APIClient) {
			name := "n"
			c.KvAPI.CloudPostV1Kv(ctx).CloudProvisionRequest(
				CloudProvisionRequest{Name: &name},
			).Execute()
		}},
		{"agent", "POST", "/v1/agents", func(c *APIClient) {
			name, model := "n", "zen-1"
			c.AgentsAPI.CloudPostV1Agents(ctx).CloudCreateAgentIn(
				CloudCreateAgentIn{Name: &name, Model: &model},
			).Execute()
		}},
		{"tools", "GET", "/v1/tools", func(c *APIClient) {
			c.ToolsAPI.CloudGetV1Tools(ctx).Execute()
		}},
	} {
		t.Run(tc.flow, func(t *testing.T) {
			var got *http.Request
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				got = r.Clone(context.Background())
				w.Header().Set("Content-Type", "application/json")
				w.Write([]byte(`{}`))
			}))
			defer srv.Close()
			t.Setenv("HANZO_BASE_URL", srv.URL)

			tc.call(NewClient("test-key"))

			if got == nil {
				t.Fatal("client made no request")
			}
			if got.Method != tc.method || got.URL.Path != tc.path {
				t.Errorf("%s %s, want %s %s", got.Method, got.URL.Path, tc.method, tc.path)
			}
			if auth := got.Header.Get("Authorization"); auth != "Bearer test-key" {
				t.Errorf("Authorization = %q, want %q", auth, "Bearer test-key")
			}
		})
	}
}
