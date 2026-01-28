// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hanzoai_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/hanzoai/go-sdk"
	"github.com/hanzoai/go-sdk/internal/testutil"
	"github.com/hanzoai/go-sdk/option"
	"github.com/hanzoai/go-sdk/shared"
)

func TestEmbeddingNewWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := hanzoai.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Embeddings.New(context.TODO(), hanzoai.EmbeddingNewParams{
		Model:             hanzoai.F("model"),
		APIBase:           hanzoai.F("api_base"),
		APIKey:            hanzoai.F("api_key"),
		APIType:           hanzoai.F("api_type"),
		APIVersion:        hanzoai.F("api_version"),
		Caching:           hanzoai.F(true),
		CustomLlmProvider: hanzoai.F[hanzoai.EmbeddingNewParamsCustomLlmProviderUnion](shared.UnionString("string")),
		Input:             hanzoai.F([]string{"string"}),
		LitellmCallID:     hanzoai.F("litellm_call_id"),
		LitellmLoggingObj: hanzoai.F(map[string]interface{}{
			"foo": "bar",
		}),
		LoggerFn: hanzoai.F("logger_fn"),
		Timeout:  hanzoai.F(int64(0)),
		User:     hanzoai.F("user"),
	})
	if err != nil {
		var apierr *hanzoai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
