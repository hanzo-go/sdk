package hanzoai

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestOptionsFallBackToTheEnvironment(t *testing.T) {
	t.Setenv("HANZO_CLIENT_ID", "env-id")
	t.Setenv("HANZO_CLIENT_SECRET", "env-secret")
	t.Setenv("HANZO_BASE_URL", "")
	t.Setenv("HANZO_ISSUER_URL", "")
	t.Setenv("HANZO_RESOURCE", "")

	got := filled(Options{})
	want := Options{
		ID: "env-id", Secret: "env-secret",
		Base: DefaultBaseURL, Issuer: DefaultIssuerURL, Resource: DefaultBaseURL,
	}
	if got != want {
		t.Errorf("filled = %+v, want %+v", got, want)
	}

	// An option given wins over the environment, and the resource follows the
	// base it was not given beside.
	got = filled(Options{ID: "explicit", Base: "http://localhost:9999"})
	if got.ID != "explicit" {
		t.Errorf("ID = %q, want the explicit one", got.ID)
	}
	if got.Resource != "http://localhost:9999" {
		t.Errorf("Resource = %q, want the base it defaults to", got.Resource)
	}
}

// The six canonical flows every Hanzo SDK ships, pinned to the addresses the
// document serves them at. If a regeneration moves an operation this fails,
// rather than the examples silently calling the wrong endpoint.
func TestFlows(t *testing.T) {
	ctx := context.Background()

	for _, tc := range []struct {
		flow, method, path string
		call               func(*Client)
	}{
		{"hello", "GET", "/v1/account/keys", func(c *Client) {
			c.AccountAPI.GetAccountKeys(ctx).Execute()
		}},
		{"chat", "POST", "/v1/chat/completions", func(c *Client) {
			c.AiAPI.PostChatCompletions(ctx).Execute()
		}},
		{"money", "GET", "/v1/billing/balance", func(c *Client) {
			c.Budget.Balance(ctx)
		}},
		{"store", "POST", "/v1/provisioning/kv", func(c *Client) {
			name := "n"
			c.ProvisioningAPI.PostProvisioningKv(ctx).ProvisionRequest(
				ProvisionRequest{Name: &name},
			).Execute()
		}},
		{"agent", "POST", "/v1/agent", func(c *Client) {
			name, model := "n", "zen-1"
			c.AgentAPI.PostAgent(ctx).CreateAgentIn(
				CreateAgentIn{Name: &name, Model: &model},
			).Execute()
		}},
		{"tools", "GET", "/v1/tool", func(c *Client) {
			c.ToolAPI.GetTool(ctx).Execute()
		}},
		{"six", "POST", "/v1/graph", func(c *Client) {
			c.Graph.Assert(ctx, nil)
		}},
	} {
		t.Run(tc.flow, func(t *testing.T) {
			var got *http.Request
			mux := http.NewServeMux()
			mux.HandleFunc("/v1/iam/oauth/token", func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.Write([]byte(`{"access_token":"tok","expires_in":3600}`))
			})
			mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
				got = r.Clone(context.Background())
				w.Header().Set("Content-Type", "application/json")
				w.Write([]byte(`{}`))
			})
			srv := httptest.NewServer(mux)
			defer srv.Close()

			tc.call(New(Options{ID: "cid", Secret: "sec", Base: srv.URL, Issuer: srv.URL, Resource: srv.URL}))

			if got == nil {
				t.Fatal("client made no request")
			}
			if got.Method != tc.method || got.URL.Path != tc.path {
				t.Errorf("%s %s, want %s %s", got.Method, got.URL.Path, tc.method, tc.path)
			}
			// One header, not two. The generated client adds the
			// Configuration's default headers AFTER reading
			// ContextAccessToken; the credential rides on the transport, which
			// is the one place it is presented.
			if auth := got.Header.Values("Authorization"); len(auth) != 1 || auth[0] != "Bearer tok" {
				t.Errorf("Authorization = %q, want exactly [%q]", auth, "Bearer tok")
			}
		})
	}
}
