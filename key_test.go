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

func TestKeyUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Key.Update(context.TODO(), hanzoai.KeyUpdateParams{
		Key: hanzoai.F("key"),
		Aliases: hanzoai.F(map[string]interface{}{
			"foo": "bar",
		}),
		AllowedCacheControls:     hanzoai.F([]interface{}{map[string]interface{}{}}),
		AllowedPassthroughRoutes: hanzoai.F([]interface{}{map[string]interface{}{}}),
		AllowedRoutes:            hanzoai.F([]interface{}{map[string]interface{}{}}),
		AllowedVectorStoreIndexes: hanzoai.F([]hanzoai.KeyUpdateParamsAllowedVectorStoreIndex{{
			IndexName:        hanzoai.F("index_name"),
			IndexPermissions: hanzoai.F([]hanzoai.KeyUpdateParamsAllowedVectorStoreIndexesIndexPermission{hanzoai.KeyUpdateParamsAllowedVectorStoreIndexesIndexPermissionRead}),
		}}),
		AutoRotate:     hanzoai.F(true),
		Blocked:        hanzoai.F(true),
		BudgetDuration: hanzoai.F("budget_duration"),
		BudgetID:       hanzoai.F("budget_id"),
		Config: hanzoai.F(map[string]interface{}{
			"foo": "bar",
		}),
		Duration:            hanzoai.F("duration"),
		EnforcedParams:      hanzoai.F([]string{"string"}),
		Guardrails:          hanzoai.F([]string{"string"}),
		KeyAlias:            hanzoai.F("key_alias"),
		MaxBudget:           hanzoai.F(0.000000),
		MaxParallelRequests: hanzoai.F(int64(0)),
		Metadata: hanzoai.F(map[string]interface{}{
			"foo": "bar",
		}),
		ModelMaxBudget: hanzoai.F(map[string]interface{}{
			"foo": "bar",
		}),
		ModelRpmLimit: hanzoai.F(map[string]interface{}{
			"foo": "bar",
		}),
		ModelTpmLimit: hanzoai.F(map[string]interface{}{
			"foo": "bar",
		}),
		Models: hanzoai.F([]interface{}{map[string]interface{}{}}),
		ObjectPermission: hanzoai.F(hanzoai.KeyUpdateParamsObjectPermission{
			AgentAccessGroups: hanzoai.F([]string{"string"}),
			Agents:            hanzoai.F([]string{"string"}),
			McpAccessGroups:   hanzoai.F([]string{"string"}),
			McpServers:        hanzoai.F([]string{"string"}),
			McpToolPermissions: hanzoai.F(map[string][]string{
				"foo": {"string"},
			}),
			VectorStores: hanzoai.F([]string{"string"}),
		}),
		Permissions: hanzoai.F(map[string]interface{}{
			"foo": "bar",
		}),
		Prompts:          hanzoai.F([]string{"string"}),
		RotationInterval: hanzoai.F("rotation_interval"),
		RouterSettings: hanzoai.F(hanzoai.KeyUpdateParamsRouterSettings{
			AllowedFails: hanzoai.F(int64(0)),
			ContextWindowFallbacks: hanzoai.F([]map[string]interface{}{{
				"foo": "bar",
			}}),
			CooldownTime: hanzoai.F(0.000000),
			Fallbacks: hanzoai.F([]map[string]interface{}{{
				"foo": "bar",
			}}),
			MaxRetries: hanzoai.F(int64(0)),
			ModelGroupAlias: hanzoai.F(map[string]hanzoai.KeyUpdateParamsRouterSettingsModelGroupAliasUnion{
				"foo": shared.UnionString("string"),
			}),
			ModelGroupRetryPolicy: hanzoai.F(map[string]interface{}{
				"foo": "bar",
			}),
			NumRetries:      hanzoai.F(int64(0)),
			RetryAfter:      hanzoai.F(0.000000),
			RoutingStrategy: hanzoai.F("routing_strategy"),
			RoutingStrategyArgs: hanzoai.F(map[string]interface{}{
				"foo": "bar",
			}),
			Timeout: hanzoai.F(0.000000),
		}),
		RpmLimit:           hanzoai.F(int64(0)),
		RpmLimitType:       hanzoai.F(hanzoai.KeyUpdateParamsRpmLimitTypeGuaranteedThroughput),
		Spend:              hanzoai.F(0.000000),
		Tags:               hanzoai.F([]string{"string"}),
		TeamID:             hanzoai.F("team_id"),
		TempBudgetExpiry:   hanzoai.F(time.Now()),
		TempBudgetIncrease: hanzoai.F(0.000000),
		TpmLimit:           hanzoai.F(int64(0)),
		TpmLimitType:       hanzoai.F(hanzoai.KeyUpdateParamsTpmLimitTypeGuaranteedThroughput),
		UserID:             hanzoai.F("user_id"),
		LitellmChangedBy:   hanzoai.F("litellm-changed-by"),
	})
	if err != nil {
		var apierr *hanzoai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestKeyListWithOptionalParams(t *testing.T) {
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
	_, err := client.Key.List(context.TODO(), hanzoai.KeyListParams{
		Expand:               hanzoai.F([]string{"string"}),
		IncludeCreatedByKeys: hanzoai.F(true),
		IncludeTeamKeys:      hanzoai.F(true),
		KeyAlias:             hanzoai.F("key_alias"),
		KeyHash:              hanzoai.F("key_hash"),
		OrganizationID:       hanzoai.F("organization_id"),
		Page:                 hanzoai.F(int64(1)),
		ReturnFullObject:     hanzoai.F(true),
		Size:                 hanzoai.F(int64(1)),
		SortBy:               hanzoai.F("sort_by"),
		SortOrder:            hanzoai.F("sort_order"),
		Status:               hanzoai.F("status"),
		TeamID:               hanzoai.F("team_id"),
		UserID:               hanzoai.F("user_id"),
	})
	if err != nil {
		var apierr *hanzoai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestKeyDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.Key.Delete(context.TODO(), hanzoai.KeyDeleteParams{
		KeyAliases:       hanzoai.F([]string{"string"}),
		Keys:             hanzoai.F([]string{"string"}),
		LitellmChangedBy: hanzoai.F("litellm-changed-by"),
	})
	if err != nil {
		var apierr *hanzoai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestKeyBlockWithOptionalParams(t *testing.T) {
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
	_, err := client.Key.Block(context.TODO(), hanzoai.KeyBlockParams{
		BlockKeyRequest: hanzoai.BlockKeyRequestParam{
			Key: hanzoai.F("key"),
		},
		LitellmChangedBy: hanzoai.F("litellm-changed-by"),
	})
	if err != nil {
		var apierr *hanzoai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestKeyCheckHealth(t *testing.T) {
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
	_, err := client.Key.CheckHealth(context.TODO())
	if err != nil {
		var apierr *hanzoai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestKeyGenerateWithOptionalParams(t *testing.T) {
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
	_, err := client.Key.Generate(context.TODO(), hanzoai.KeyGenerateParams{
		Aliases: hanzoai.F(map[string]interface{}{
			"foo": "bar",
		}),
		AllowedCacheControls:     hanzoai.F([]interface{}{map[string]interface{}{}}),
		AllowedPassthroughRoutes: hanzoai.F([]interface{}{map[string]interface{}{}}),
		AllowedRoutes:            hanzoai.F([]interface{}{map[string]interface{}{}}),
		AllowedVectorStoreIndexes: hanzoai.F([]hanzoai.KeyGenerateParamsAllowedVectorStoreIndex{{
			IndexName:        hanzoai.F("index_name"),
			IndexPermissions: hanzoai.F([]hanzoai.KeyGenerateParamsAllowedVectorStoreIndexesIndexPermission{hanzoai.KeyGenerateParamsAllowedVectorStoreIndexesIndexPermissionRead}),
		}}),
		AutoRotate:     hanzoai.F(true),
		Blocked:        hanzoai.F(true),
		BudgetDuration: hanzoai.F("budget_duration"),
		BudgetID:       hanzoai.F("budget_id"),
		Config: hanzoai.F(map[string]interface{}{
			"foo": "bar",
		}),
		Duration:            hanzoai.F("duration"),
		EnforcedParams:      hanzoai.F([]string{"string"}),
		Guardrails:          hanzoai.F([]string{"string"}),
		Key:                 hanzoai.F("key"),
		KeyAlias:            hanzoai.F("key_alias"),
		KeyType:             hanzoai.F(hanzoai.KeyGenerateParamsKeyTypeLlmAPI),
		MaxBudget:           hanzoai.F(0.000000),
		MaxParallelRequests: hanzoai.F(int64(0)),
		Metadata: hanzoai.F(map[string]interface{}{
			"foo": "bar",
		}),
		ModelMaxBudget: hanzoai.F(map[string]interface{}{
			"foo": "bar",
		}),
		ModelRpmLimit: hanzoai.F(map[string]interface{}{
			"foo": "bar",
		}),
		ModelTpmLimit: hanzoai.F(map[string]interface{}{
			"foo": "bar",
		}),
		Models: hanzoai.F([]interface{}{map[string]interface{}{}}),
		ObjectPermission: hanzoai.F(hanzoai.KeyGenerateParamsObjectPermission{
			AgentAccessGroups: hanzoai.F([]string{"string"}),
			Agents:            hanzoai.F([]string{"string"}),
			McpAccessGroups:   hanzoai.F([]string{"string"}),
			McpServers:        hanzoai.F([]string{"string"}),
			McpToolPermissions: hanzoai.F(map[string][]string{
				"foo": {"string"},
			}),
			VectorStores: hanzoai.F([]string{"string"}),
		}),
		OrganizationID: hanzoai.F("organization_id"),
		Permissions: hanzoai.F(map[string]interface{}{
			"foo": "bar",
		}),
		Prompts:          hanzoai.F([]string{"string"}),
		RotationInterval: hanzoai.F("rotation_interval"),
		RouterSettings: hanzoai.F(hanzoai.KeyGenerateParamsRouterSettings{
			AllowedFails: hanzoai.F(int64(0)),
			ContextWindowFallbacks: hanzoai.F([]map[string]interface{}{{
				"foo": "bar",
			}}),
			CooldownTime: hanzoai.F(0.000000),
			Fallbacks: hanzoai.F([]map[string]interface{}{{
				"foo": "bar",
			}}),
			MaxRetries: hanzoai.F(int64(0)),
			ModelGroupAlias: hanzoai.F(map[string]hanzoai.KeyGenerateParamsRouterSettingsModelGroupAliasUnion{
				"foo": shared.UnionString("string"),
			}),
			ModelGroupRetryPolicy: hanzoai.F(map[string]interface{}{
				"foo": "bar",
			}),
			NumRetries:      hanzoai.F(int64(0)),
			RetryAfter:      hanzoai.F(0.000000),
			RoutingStrategy: hanzoai.F("routing_strategy"),
			RoutingStrategyArgs: hanzoai.F(map[string]interface{}{
				"foo": "bar",
			}),
			Timeout: hanzoai.F(0.000000),
		}),
		RpmLimit:         hanzoai.F(int64(0)),
		RpmLimitType:     hanzoai.F(hanzoai.KeyGenerateParamsRpmLimitTypeGuaranteedThroughput),
		SendInviteEmail:  hanzoai.F(true),
		SoftBudget:       hanzoai.F(0.000000),
		Spend:            hanzoai.F(0.000000),
		Tags:             hanzoai.F([]string{"string"}),
		TeamID:           hanzoai.F("team_id"),
		TpmLimit:         hanzoai.F(int64(0)),
		TpmLimitType:     hanzoai.F(hanzoai.KeyGenerateParamsTpmLimitTypeGuaranteedThroughput),
		UserID:           hanzoai.F("user_id"),
		LitellmChangedBy: hanzoai.F("litellm-changed-by"),
	})
	if err != nil {
		var apierr *hanzoai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestKeyRegenerateByKeyWithOptionalParams(t *testing.T) {
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
	_, err := client.Key.RegenerateByKey(
		context.TODO(),
		"key",
		hanzoai.KeyRegenerateByKeyParams{
			RegenerateKeyRequest: hanzoai.RegenerateKeyRequestParam{
				Aliases: hanzoai.F(map[string]interface{}{
					"foo": "bar",
				}),
				AllowedCacheControls:     hanzoai.F([]interface{}{map[string]interface{}{}}),
				AllowedPassthroughRoutes: hanzoai.F([]interface{}{map[string]interface{}{}}),
				AllowedRoutes:            hanzoai.F([]interface{}{map[string]interface{}{}}),
				AllowedVectorStoreIndexes: hanzoai.F([]hanzoai.RegenerateKeyRequestAllowedVectorStoreIndexParam{{
					IndexName:        hanzoai.F("index_name"),
					IndexPermissions: hanzoai.F([]hanzoai.RegenerateKeyRequestAllowedVectorStoreIndexesIndexPermission{hanzoai.RegenerateKeyRequestAllowedVectorStoreIndexesIndexPermissionRead}),
				}}),
				AutoRotate:     hanzoai.F(true),
				Blocked:        hanzoai.F(true),
				BudgetDuration: hanzoai.F("budget_duration"),
				BudgetID:       hanzoai.F("budget_id"),
				Config: hanzoai.F(map[string]interface{}{
					"foo": "bar",
				}),
				Duration:            hanzoai.F("duration"),
				EnforcedParams:      hanzoai.F([]string{"string"}),
				Guardrails:          hanzoai.F([]string{"string"}),
				Key:                 hanzoai.F("key"),
				KeyAlias:            hanzoai.F("key_alias"),
				KeyType:             hanzoai.F(hanzoai.RegenerateKeyRequestKeyTypeLlmAPI),
				MaxBudget:           hanzoai.F(0.000000),
				MaxParallelRequests: hanzoai.F(int64(0)),
				Metadata: hanzoai.F(map[string]interface{}{
					"foo": "bar",
				}),
				ModelMaxBudget: hanzoai.F(map[string]interface{}{
					"foo": "bar",
				}),
				ModelRpmLimit: hanzoai.F(map[string]interface{}{
					"foo": "bar",
				}),
				ModelTpmLimit: hanzoai.F(map[string]interface{}{
					"foo": "bar",
				}),
				Models:       hanzoai.F([]interface{}{map[string]interface{}{}}),
				NewKey:       hanzoai.F("new_key"),
				NewMasterKey: hanzoai.F("new_master_key"),
				ObjectPermission: hanzoai.F(hanzoai.RegenerateKeyRequestObjectPermissionParam{
					AgentAccessGroups: hanzoai.F([]string{"string"}),
					Agents:            hanzoai.F([]string{"string"}),
					McpAccessGroups:   hanzoai.F([]string{"string"}),
					McpServers:        hanzoai.F([]string{"string"}),
					McpToolPermissions: hanzoai.F(map[string][]string{
						"foo": {"string"},
					}),
					VectorStores: hanzoai.F([]string{"string"}),
				}),
				OrganizationID: hanzoai.F("organization_id"),
				Permissions: hanzoai.F(map[string]interface{}{
					"foo": "bar",
				}),
				Prompts:          hanzoai.F([]string{"string"}),
				RotationInterval: hanzoai.F("rotation_interval"),
				RouterSettings: hanzoai.F(hanzoai.RegenerateKeyRequestRouterSettingsParam{
					AllowedFails: hanzoai.F(int64(0)),
					ContextWindowFallbacks: hanzoai.F([]map[string]interface{}{{
						"foo": "bar",
					}}),
					CooldownTime: hanzoai.F(0.000000),
					Fallbacks: hanzoai.F([]map[string]interface{}{{
						"foo": "bar",
					}}),
					MaxRetries: hanzoai.F(int64(0)),
					ModelGroupAlias: hanzoai.F(map[string]hanzoai.RegenerateKeyRequestRouterSettingsModelGroupAliasUnionParam{
						"foo": shared.UnionString("string"),
					}),
					ModelGroupRetryPolicy: hanzoai.F(map[string]interface{}{
						"foo": "bar",
					}),
					NumRetries:      hanzoai.F(int64(0)),
					RetryAfter:      hanzoai.F(0.000000),
					RoutingStrategy: hanzoai.F("routing_strategy"),
					RoutingStrategyArgs: hanzoai.F(map[string]interface{}{
						"foo": "bar",
					}),
					Timeout: hanzoai.F(0.000000),
				}),
				RpmLimit:        hanzoai.F(int64(0)),
				RpmLimitType:    hanzoai.F(hanzoai.RegenerateKeyRequestRpmLimitTypeGuaranteedThroughput),
				SendInviteEmail: hanzoai.F(true),
				SoftBudget:      hanzoai.F(0.000000),
				Spend:           hanzoai.F(0.000000),
				Tags:            hanzoai.F([]string{"string"}),
				TeamID:          hanzoai.F("team_id"),
				TpmLimit:        hanzoai.F(int64(0)),
				TpmLimitType:    hanzoai.F(hanzoai.RegenerateKeyRequestTpmLimitTypeGuaranteedThroughput),
				UserID:          hanzoai.F("user_id"),
			},
			LitellmChangedBy: hanzoai.F("litellm-changed-by"),
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

func TestKeyGetInfoWithOptionalParams(t *testing.T) {
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
	_, err := client.Key.GetInfo(context.TODO(), hanzoai.KeyGetInfoParams{
		Key: hanzoai.F("key"),
	})
	if err != nil {
		var apierr *hanzoai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestKeyUnblockWithOptionalParams(t *testing.T) {
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
	_, err := client.Key.Unblock(context.TODO(), hanzoai.KeyUnblockParams{
		BlockKeyRequest: hanzoai.BlockKeyRequestParam{
			Key: hanzoai.F("key"),
		},
		LitellmChangedBy: hanzoai.F("litellm-changed-by"),
	})
	if err != nil {
		var apierr *hanzoai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
