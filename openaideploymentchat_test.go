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

func TestOpenAIDeploymentChatCompleteWithOptionalParams(t *testing.T) {
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
	_, err := client.OpenAI.Deployments.Chat.Complete(
		context.TODO(),
		"model",
		hanzoai.OpenAIDeploymentChatCompleteParams{
			Messages: hanzoai.F([]hanzoai.OpenAIDeploymentChatCompleteParamsMessageUnion{hanzoai.OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessage{
				Content: hanzoai.F[hanzoai.OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentUnion](shared.UnionString("Hello, how are you?")),
				Role:    hanzoai.F(hanzoai.OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageRoleUser),
				CacheControl: hanzoai.F(hanzoai.OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageCacheControl{
					Type: hanzoai.F(hanzoai.OpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageCacheControlTypeEphemeral),
				}),
			}}),
			Model:   hanzoai.F("model"),
			Caching: hanzoai.F(true),
			ContextWindowFallbackDict: hanzoai.F(map[string]string{
				"foo": "string",
			}),
			Fallbacks:        hanzoai.F([]string{"string"}),
			FrequencyPenalty: hanzoai.F(0.000000),
			FunctionCall:     hanzoai.F[hanzoai.OpenAIDeploymentChatCompleteParamsFunctionCallUnion](shared.UnionString("string")),
			Functions: hanzoai.F([]map[string]interface{}{{
				"foo": "bar",
			}}),
			Guardrails: hanzoai.F([]string{"string"}),
			LogitBias: hanzoai.F(map[string]float64{
				"foo": 0.000000,
			}),
			Logprobs:  hanzoai.F(true),
			MaxTokens: hanzoai.F(int64(0)),
			Metadata: hanzoai.F(map[string]interface{}{
				"foo": "bar",
			}),
			N:                 hanzoai.F(int64(0)),
			NumRetries:        hanzoai.F(int64(0)),
			ParallelToolCalls: hanzoai.F(true),
			PresencePenalty:   hanzoai.F(0.000000),
			ResponseFormat: hanzoai.F(map[string]interface{}{
				"foo": "bar",
			}),
			Seed:        hanzoai.F(int64(0)),
			ServiceTier: hanzoai.F("service_tier"),
			Stop:        hanzoai.F[hanzoai.OpenAIDeploymentChatCompleteParamsStopUnion](shared.UnionString("string")),
			Stream:      hanzoai.F(true),
			StreamOptions: hanzoai.F(map[string]interface{}{
				"foo": "bar",
			}),
			Temperature: hanzoai.F(0.000000),
			ToolChoice:  hanzoai.F[hanzoai.OpenAIDeploymentChatCompleteParamsToolChoiceUnion](shared.UnionString("string")),
			Tools: hanzoai.F([]map[string]interface{}{{
				"foo": "bar",
			}}),
			TopLogprobs: hanzoai.F(int64(0)),
			TopP:        hanzoai.F(0.000000),
			User:        hanzoai.F("user"),
		},
	)
	if err != nil {
		var apierr *hanzoai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
