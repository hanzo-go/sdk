package hanzoai_test

import (
	"testing"

	"github.com/hanzoai/go-sdk"
	"github.com/hanzoai/go-sdk/option"
)

func TestNewClient(t *testing.T) {
	tests := []struct {
		name    string
		opts    []option.RequestOption
		wantNil bool
	}{
		{
			name:    "default client",
			opts:    nil,
			wantNil: false,
		},
		{
			name:    "client with API key",
			opts:    []option.RequestOption{option.WithAPIKey("test-key")},
			wantNil: false,
		},
		{
			name:    "client with base URL",
			opts:    []option.RequestOption{option.WithBaseURL("https://api.test.com")},
			wantNil: false,
		},
		{
			name: "client with multiple options",
			opts: []option.RequestOption{
				option.WithAPIKey("test-key"),
				option.WithBaseURL("https://api.test.com"),
			},
			wantNil: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := hanzoai.NewClient(tt.opts...)
			if (client == nil) != tt.wantNil {
				t.Errorf("NewClient() returned nil = %v, want %v", client == nil, tt.wantNil)
			}

			// Verify all services are initialized
			if client != nil {
				verifyServices(t, client)
			}
		})
	}
}

func verifyServices(t *testing.T, client *hanzoai.Client) {
	if client.Models == nil {
		t.Error("Models service is nil")
	}
	if client.OpenAI == nil {
		t.Error("OpenAI service is nil")
	}
	if client.Chat == nil {
		t.Error("Chat service is nil")
	}
	if client.Completions == nil {
		t.Error("Completions service is nil")
	}
	if client.Embeddings == nil {
		t.Error("Embeddings service is nil")
	}
	if client.Images == nil {
		t.Error("Images service is nil")
	}
	if client.Audio == nil {
		t.Error("Audio service is nil")
	}
	if client.Assistants == nil {
		t.Error("Assistants service is nil")
	}
	if client.Threads == nil {
		t.Error("Threads service is nil")
	}
	if client.Moderations == nil {
		t.Error("Moderations service is nil")
	}
	if client.Utils == nil {
		t.Error("Utils service is nil")
	}
	if client.Model == nil {
		t.Error("Model service is nil")
	}
	if client.Model.Info == nil {
		t.Error("Model.Info service is nil")
	}
	if client.Model.Update == nil {
		t.Error("Model.Update service is nil")
	}
	if client.ModelGroup == nil {
		t.Error("ModelGroup service is nil")
	}
	if client.Routes == nil {
		t.Error("Routes service is nil")
	}
	if client.Responses == nil {
		t.Error("Responses service is nil")
	}
	if client.Batches == nil {
		t.Error("Batches service is nil")
	}
	if client.Rerank == nil {
		t.Error("Rerank service is nil")
	}
	if client.FineTuning == nil {
		t.Error("FineTuning service is nil")
	}
	if client.Credentials == nil {
		t.Error("Credentials service is nil")
	}
	if client.VertexAI == nil {
		t.Error("VertexAI service is nil")
	}
	if client.Gemini == nil {
		t.Error("Gemini service is nil")
	}
	if client.Cohere == nil {
		t.Error("Cohere service is nil")
	}
	if client.Anthropic == nil {
		t.Error("Anthropic service is nil")
	}
	if client.Bedrock == nil {
		t.Error("Bedrock service is nil")
	}
	if client.Azure == nil {
		t.Error("Azure service is nil")
	}
	if client.Langfuse == nil {
		t.Error("Langfuse service is nil")
	}
	if client.Config == nil {
		t.Error("Config service is nil")
	}
	if client.Test == nil {
		t.Error("Test service is nil")
	}
	if client.Health == nil {
		t.Error("Health service is nil")
	}
	if client.Active == nil {
		t.Error("Active service is nil")
	}
	if client.Settings == nil {
		t.Error("Settings service is nil")
	}
	if client.Key == nil {
		t.Error("Key service is nil")
	}
	if client.User == nil {
		t.Error("User service is nil")
	}
	if client.Team == nil {
		t.Error("Team service is nil")
	}
	if client.Organization == nil {
		t.Error("Organization service is nil")
	}
	if client.Customer == nil {
		t.Error("Customer service is nil")
	}
	if client.Spend == nil {
		t.Error("Spend service is nil")
	}
	if client.Global == nil {
		t.Error("Global service is nil")
	}
	if client.Provider == nil {
		t.Error("Provider service is nil")
	}
	if client.Cache == nil {
		t.Error("Cache service is nil")
	}
	if client.Guardrails == nil {
		t.Error("Guardrails service is nil")
	}
	if client.Add == nil {
		t.Error("Add service is nil")
	}
	if client.Deletes == nil {
		t.Error("Deletes service is nil")
	}
	if client.Files == nil {
		t.Error("Files service is nil")
	}
	if client.Budget == nil {
		t.Error("Budget service is nil")
	}
}

func TestDefaultClientOptions(t *testing.T) {
	// Test that default options are returned
	opts := hanzoai.DefaultClientOptions()
	if opts == nil {
		t.Error("DefaultClientOptions() returned nil")
	}
	if len(opts) == 0 {
		t.Error("DefaultClientOptions() returned empty slice")
	}
}