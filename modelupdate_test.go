// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hanzoai_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/hanzoai/go-sdk"
	"github.com/hanzoai/go-sdk/internal/testutil"
	"github.com/hanzoai/go-sdk/option"
	"github.com/hanzoai/go-sdk/shared"
)

func TestModelUpdateFullWithOptionalParams(t *testing.T) {
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
	_, err := client.Model.Update.Full(context.TODO(), hanzoai.ModelUpdateFullParams{
		UpdateDeployment: hanzoai.UpdateDeploymentParam{
			LitellmParams: hanzoai.F(hanzoai.UpdateDeploymentLitellmParamsParam{
				APIBase:                                    hanzoai.F("api_base"),
				APIKey:                                     hanzoai.F("api_key"),
				APIVersion:                                 hanzoai.F("api_version"),
				AutoRouterConfig:                           hanzoai.F("auto_router_config"),
				AutoRouterConfigPath:                       hanzoai.F("auto_router_config_path"),
				AutoRouterDefaultModel:                     hanzoai.F("auto_router_default_model"),
				AutoRouterEmbeddingModel:                   hanzoai.F("auto_router_embedding_model"),
				AwsAccessKeyID:                             hanzoai.F("aws_access_key_id"),
				AwsBedrockRuntimeEndpoint:                  hanzoai.F("aws_bedrock_runtime_endpoint"),
				AwsRegionName:                              hanzoai.F("aws_region_name"),
				AwsSecretAccessKey:                         hanzoai.F("aws_secret_access_key"),
				BudgetDuration:                             hanzoai.F("budget_duration"),
				CacheCreationInputAudioTokenCost:           hanzoai.F(0.000000),
				CacheCreationInputTokenCost:                hanzoai.F(0.000000),
				CacheCreationInputTokenCostAbove1hr:        hanzoai.F(0.000000),
				CacheCreationInputTokenCostAbove200kTokens: hanzoai.F(0.000000),
				CacheReadInputAudioTokenCost:               hanzoai.F(0.000000),
				CacheReadInputTokenCost:                    hanzoai.F(0.000000),
				CacheReadInputTokenCostAbove200kTokens:     hanzoai.F(0.000000),
				CacheReadInputTokenCostFlex:                hanzoai.F(0.000000),
				CacheReadInputTokenCostPriority:            hanzoai.F(0.000000),
				CitationCostPerToken:                       hanzoai.F(0.000000),
				ConfigurableClientsideAuthParams:           hanzoai.F([]hanzoai.UpdateDeploymentLitellmParamsConfigurableClientsideAuthParamsUnionParam{shared.UnionString("string")}),
				CustomLlmProvider:                          hanzoai.F("custom_llm_provider"),
				GcsBucketName:                              hanzoai.F("gcs_bucket_name"),
				InputCostPerAudioPerSecond:                 hanzoai.F(0.000000),
				InputCostPerAudioPerSecondAbove128kTokens:  hanzoai.F(0.000000),
				InputCostPerAudioToken:                     hanzoai.F(0.000000),
				InputCostPerCharacter:                      hanzoai.F(0.000000),
				InputCostPerCharacterAbove128kTokens:       hanzoai.F(0.000000),
				InputCostPerImage:                          hanzoai.F(0.000000),
				InputCostPerImageAbove128kTokens:           hanzoai.F(0.000000),
				InputCostPerPixel:                          hanzoai.F(0.000000),
				InputCostPerQuery:                          hanzoai.F(0.000000),
				InputCostPerSecond:                         hanzoai.F(0.000000),
				InputCostPerToken:                          hanzoai.F(0.000000),
				InputCostPerTokenAbove128kTokens:           hanzoai.F(0.000000),
				InputCostPerTokenAbove200kTokens:           hanzoai.F(0.000000),
				InputCostPerTokenBatches:                   hanzoai.F(0.000000),
				InputCostPerTokenCacheHit:                  hanzoai.F(0.000000),
				InputCostPerTokenFlex:                      hanzoai.F(0.000000),
				InputCostPerTokenPriority:                  hanzoai.F(0.000000),
				InputCostPerVideoPerSecond:                 hanzoai.F(0.000000),
				InputCostPerVideoPerSecondAbove128kTokens:  hanzoai.F(0.000000),
				InputCostPerVideoPerSecondAbove15sInterval: hanzoai.F(0.000000),
				InputCostPerVideoPerSecondAbove8sInterval:  hanzoai.F(0.000000),
				LitellmCredentialName:                      hanzoai.F("litellm_credential_name"),
				LitellmTraceID:                             hanzoai.F("litellm_trace_id"),
				MaxBudget:                                  hanzoai.F(0.000000),
				MaxFileSizeMB:                              hanzoai.F(0.000000),
				MaxRetries:                                 hanzoai.F(int64(0)),
				MergeReasoningContentInChoices:             hanzoai.F(true),
				MilvusTextField:                            hanzoai.F("milvus_text_field"),
				MockResponse:                               hanzoai.F[hanzoai.UpdateDeploymentLitellmParamsMockResponseUnionParam](shared.UnionString("string")),
				Model:                                      hanzoai.F("model"),
				ModelInfo: hanzoai.F(map[string]interface{}{
					"foo": "bar",
				}),
				Organization:                          hanzoai.F("organization"),
				OutputCostPerAudioPerSecond:           hanzoai.F(0.000000),
				OutputCostPerAudioToken:               hanzoai.F(0.000000),
				OutputCostPerCharacter:                hanzoai.F(0.000000),
				OutputCostPerCharacterAbove128kTokens: hanzoai.F(0.000000),
				OutputCostPerImage:                    hanzoai.F(0.000000),
				OutputCostPerImageToken:               hanzoai.F(0.000000),
				OutputCostPerPixel:                    hanzoai.F(0.000000),
				OutputCostPerReasoningToken:           hanzoai.F(0.000000),
				OutputCostPerSecond:                   hanzoai.F(0.000000),
				OutputCostPerToken:                    hanzoai.F(0.000000),
				OutputCostPerTokenAbove128kTokens:     hanzoai.F(0.000000),
				OutputCostPerTokenAbove200kTokens:     hanzoai.F(0.000000),
				OutputCostPerTokenBatches:             hanzoai.F(0.000000),
				OutputCostPerTokenFlex:                hanzoai.F(0.000000),
				OutputCostPerTokenPriority:            hanzoai.F(0.000000),
				OutputCostPerVideoPerSecond:           hanzoai.F(0.000000),
				RegionName:                            hanzoai.F("region_name"),
				Rpm:                                   hanzoai.F(int64(0)),
				S3BucketName:                          hanzoai.F("s3_bucket_name"),
				S3EncryptionKeyID:                     hanzoai.F("s3_encryption_key_id"),
				SearchContextCostPerQuery: hanzoai.F(map[string]interface{}{
					"foo": "bar",
				}),
				StreamTimeout: hanzoai.F[hanzoai.UpdateDeploymentLitellmParamsStreamTimeoutUnionParam](shared.UnionFloat(0.000000)),
				TieredPricing: hanzoai.F([]map[string]interface{}{{
					"foo": "bar",
				}}),
				Timeout:           hanzoai.F[hanzoai.UpdateDeploymentLitellmParamsTimeoutUnionParam](shared.UnionFloat(0.000000)),
				Tpm:               hanzoai.F(int64(0)),
				UseInPassThrough:  hanzoai.F(true),
				UseLitellmProxy:   hanzoai.F(true),
				VectorStoreID:     hanzoai.F("vector_store_id"),
				VertexCredentials: hanzoai.F[hanzoai.UpdateDeploymentLitellmParamsVertexCredentialsUnionParam](shared.UnionString("string")),
				VertexLocation:    hanzoai.F("vertex_location"),
				VertexProject:     hanzoai.F("vertex_project"),
				WatsonxRegionName: hanzoai.F("watsonx_region_name"),
			}),
			ModelInfo: hanzoai.F(hanzoai.ModelInfoParam{
				ID:                  hanzoai.F("id"),
				BaseModel:           hanzoai.F("base_model"),
				CreatedAt:           hanzoai.F(time.Now()),
				CreatedBy:           hanzoai.F("created_by"),
				DBModel:             hanzoai.F(true),
				TeamID:              hanzoai.F("team_id"),
				TeamPublicModelName: hanzoai.F("team_public_model_name"),
				Tier:                hanzoai.F(hanzoai.ModelInfoTierFree),
				UpdatedAt:           hanzoai.F(time.Now()),
				UpdatedBy:           hanzoai.F("updated_by"),
			}),
			ModelName: hanzoai.F("model_name"),
		},
	})
	if err != nil {
		var apierr *hanzoai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestModelUpdatePartialWithOptionalParams(t *testing.T) {
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
	_, err := client.Model.Update.Partial(
		context.TODO(),
		"model_id",
		hanzoai.ModelUpdatePartialParams{
			UpdateDeployment: hanzoai.UpdateDeploymentParam{
				LitellmParams: hanzoai.F(hanzoai.UpdateDeploymentLitellmParamsParam{
					APIBase:                                    hanzoai.F("api_base"),
					APIKey:                                     hanzoai.F("api_key"),
					APIVersion:                                 hanzoai.F("api_version"),
					AutoRouterConfig:                           hanzoai.F("auto_router_config"),
					AutoRouterConfigPath:                       hanzoai.F("auto_router_config_path"),
					AutoRouterDefaultModel:                     hanzoai.F("auto_router_default_model"),
					AutoRouterEmbeddingModel:                   hanzoai.F("auto_router_embedding_model"),
					AwsAccessKeyID:                             hanzoai.F("aws_access_key_id"),
					AwsBedrockRuntimeEndpoint:                  hanzoai.F("aws_bedrock_runtime_endpoint"),
					AwsRegionName:                              hanzoai.F("aws_region_name"),
					AwsSecretAccessKey:                         hanzoai.F("aws_secret_access_key"),
					BudgetDuration:                             hanzoai.F("budget_duration"),
					CacheCreationInputAudioTokenCost:           hanzoai.F(0.000000),
					CacheCreationInputTokenCost:                hanzoai.F(0.000000),
					CacheCreationInputTokenCostAbove1hr:        hanzoai.F(0.000000),
					CacheCreationInputTokenCostAbove200kTokens: hanzoai.F(0.000000),
					CacheReadInputAudioTokenCost:               hanzoai.F(0.000000),
					CacheReadInputTokenCost:                    hanzoai.F(0.000000),
					CacheReadInputTokenCostAbove200kTokens:     hanzoai.F(0.000000),
					CacheReadInputTokenCostFlex:                hanzoai.F(0.000000),
					CacheReadInputTokenCostPriority:            hanzoai.F(0.000000),
					CitationCostPerToken:                       hanzoai.F(0.000000),
					ConfigurableClientsideAuthParams:           hanzoai.F([]hanzoai.UpdateDeploymentLitellmParamsConfigurableClientsideAuthParamsUnionParam{shared.UnionString("string")}),
					CustomLlmProvider:                          hanzoai.F("custom_llm_provider"),
					GcsBucketName:                              hanzoai.F("gcs_bucket_name"),
					InputCostPerAudioPerSecond:                 hanzoai.F(0.000000),
					InputCostPerAudioPerSecondAbove128kTokens:  hanzoai.F(0.000000),
					InputCostPerAudioToken:                     hanzoai.F(0.000000),
					InputCostPerCharacter:                      hanzoai.F(0.000000),
					InputCostPerCharacterAbove128kTokens:       hanzoai.F(0.000000),
					InputCostPerImage:                          hanzoai.F(0.000000),
					InputCostPerImageAbove128kTokens:           hanzoai.F(0.000000),
					InputCostPerPixel:                          hanzoai.F(0.000000),
					InputCostPerQuery:                          hanzoai.F(0.000000),
					InputCostPerSecond:                         hanzoai.F(0.000000),
					InputCostPerToken:                          hanzoai.F(0.000000),
					InputCostPerTokenAbove128kTokens:           hanzoai.F(0.000000),
					InputCostPerTokenAbove200kTokens:           hanzoai.F(0.000000),
					InputCostPerTokenBatches:                   hanzoai.F(0.000000),
					InputCostPerTokenCacheHit:                  hanzoai.F(0.000000),
					InputCostPerTokenFlex:                      hanzoai.F(0.000000),
					InputCostPerTokenPriority:                  hanzoai.F(0.000000),
					InputCostPerVideoPerSecond:                 hanzoai.F(0.000000),
					InputCostPerVideoPerSecondAbove128kTokens:  hanzoai.F(0.000000),
					InputCostPerVideoPerSecondAbove15sInterval: hanzoai.F(0.000000),
					InputCostPerVideoPerSecondAbove8sInterval:  hanzoai.F(0.000000),
					LitellmCredentialName:                      hanzoai.F("litellm_credential_name"),
					LitellmTraceID:                             hanzoai.F("litellm_trace_id"),
					MaxBudget:                                  hanzoai.F(0.000000),
					MaxFileSizeMB:                              hanzoai.F(0.000000),
					MaxRetries:                                 hanzoai.F(int64(0)),
					MergeReasoningContentInChoices:             hanzoai.F(true),
					MilvusTextField:                            hanzoai.F("milvus_text_field"),
					MockResponse:                               hanzoai.F[hanzoai.UpdateDeploymentLitellmParamsMockResponseUnionParam](shared.UnionString("string")),
					Model:                                      hanzoai.F("model"),
					ModelInfo: hanzoai.F(map[string]interface{}{
						"foo": "bar",
					}),
					Organization:                          hanzoai.F("organization"),
					OutputCostPerAudioPerSecond:           hanzoai.F(0.000000),
					OutputCostPerAudioToken:               hanzoai.F(0.000000),
					OutputCostPerCharacter:                hanzoai.F(0.000000),
					OutputCostPerCharacterAbove128kTokens: hanzoai.F(0.000000),
					OutputCostPerImage:                    hanzoai.F(0.000000),
					OutputCostPerImageToken:               hanzoai.F(0.000000),
					OutputCostPerPixel:                    hanzoai.F(0.000000),
					OutputCostPerReasoningToken:           hanzoai.F(0.000000),
					OutputCostPerSecond:                   hanzoai.F(0.000000),
					OutputCostPerToken:                    hanzoai.F(0.000000),
					OutputCostPerTokenAbove128kTokens:     hanzoai.F(0.000000),
					OutputCostPerTokenAbove200kTokens:     hanzoai.F(0.000000),
					OutputCostPerTokenBatches:             hanzoai.F(0.000000),
					OutputCostPerTokenFlex:                hanzoai.F(0.000000),
					OutputCostPerTokenPriority:            hanzoai.F(0.000000),
					OutputCostPerVideoPerSecond:           hanzoai.F(0.000000),
					RegionName:                            hanzoai.F("region_name"),
					Rpm:                                   hanzoai.F(int64(0)),
					S3BucketName:                          hanzoai.F("s3_bucket_name"),
					S3EncryptionKeyID:                     hanzoai.F("s3_encryption_key_id"),
					SearchContextCostPerQuery: hanzoai.F(map[string]interface{}{
						"foo": "bar",
					}),
					StreamTimeout: hanzoai.F[hanzoai.UpdateDeploymentLitellmParamsStreamTimeoutUnionParam](shared.UnionFloat(0.000000)),
					TieredPricing: hanzoai.F([]map[string]interface{}{{
						"foo": "bar",
					}}),
					Timeout:           hanzoai.F[hanzoai.UpdateDeploymentLitellmParamsTimeoutUnionParam](shared.UnionFloat(0.000000)),
					Tpm:               hanzoai.F(int64(0)),
					UseInPassThrough:  hanzoai.F(true),
					UseLitellmProxy:   hanzoai.F(true),
					VectorStoreID:     hanzoai.F("vector_store_id"),
					VertexCredentials: hanzoai.F[hanzoai.UpdateDeploymentLitellmParamsVertexCredentialsUnionParam](shared.UnionString("string")),
					VertexLocation:    hanzoai.F("vertex_location"),
					VertexProject:     hanzoai.F("vertex_project"),
					WatsonxRegionName: hanzoai.F("watsonx_region_name"),
				}),
				ModelInfo: hanzoai.F(hanzoai.ModelInfoParam{
					ID:                  hanzoai.F("id"),
					BaseModel:           hanzoai.F("base_model"),
					CreatedAt:           hanzoai.F(time.Now()),
					CreatedBy:           hanzoai.F("created_by"),
					DBModel:             hanzoai.F(true),
					TeamID:              hanzoai.F("team_id"),
					TeamPublicModelName: hanzoai.F("team_public_model_name"),
					Tier:                hanzoai.F(hanzoai.ModelInfoTierFree),
					UpdatedAt:           hanzoai.F(time.Now()),
					UpdatedBy:           hanzoai.F("updated_by"),
				}),
				ModelName: hanzoai.F("model_name"),
			},
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
