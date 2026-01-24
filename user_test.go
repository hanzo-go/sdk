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
)

func TestUserNewWithOptionalParams(t *testing.T) {
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
	_, err := client.User.New(context.TODO(), hanzoai.UserNewParams{
		Aliases: hanzoai.F(map[string]interface{}{
			"foo": "bar",
		}),
		AllowedCacheControls: hanzoai.F([]interface{}{map[string]interface{}{}}),
		AutoCreateKey:        hanzoai.F(true),
		Blocked:              hanzoai.F(true),
		BudgetDuration:       hanzoai.F("budget_duration"),
		Config: hanzoai.F(map[string]interface{}{
			"foo": "bar",
		}),
		Duration:            hanzoai.F("duration"),
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
		ObjectPermission: hanzoai.F(hanzoai.UserNewParamsObjectPermission{
			AgentAccessGroups: hanzoai.F([]string{"string"}),
			Agents:            hanzoai.F([]string{"string"}),
			McpAccessGroups:   hanzoai.F([]string{"string"}),
			McpServers:        hanzoai.F([]string{"string"}),
			McpToolPermissions: hanzoai.F(map[string][]string{
				"foo": {"string"},
			}),
			VectorStores: hanzoai.F([]string{"string"}),
		}),
		Organizations: hanzoai.F([]string{"string"}),
		Permissions: hanzoai.F(map[string]interface{}{
			"foo": "bar",
		}),
		Prompts:         hanzoai.F([]string{"string"}),
		RpmLimit:        hanzoai.F(int64(0)),
		SendInviteEmail: hanzoai.F(true),
		Spend:           hanzoai.F(0.000000),
		SSOUserID:       hanzoai.F("sso_user_id"),
		TeamID:          hanzoai.F("team_id"),
		Teams:           hanzoai.F[hanzoai.UserNewParamsTeamsUnion](hanzoai.UserNewParamsTeamsArray([]string{"string"})),
		TpmLimit:        hanzoai.F(int64(0)),
		UserAlias:       hanzoai.F("user_alias"),
		UserEmail:       hanzoai.F("user_email"),
		UserID:          hanzoai.F("user_id"),
		UserRole:        hanzoai.F(hanzoai.UserNewParamsUserRoleProxyAdmin),
	})
	if err != nil {
		var apierr *hanzoai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.User.Update(context.TODO(), hanzoai.UserUpdateParams{
		Aliases: hanzoai.F(map[string]interface{}{
			"foo": "bar",
		}),
		AllowedCacheControls: hanzoai.F([]interface{}{map[string]interface{}{}}),
		Blocked:              hanzoai.F(true),
		BudgetDuration:       hanzoai.F("budget_duration"),
		Config: hanzoai.F(map[string]interface{}{
			"foo": "bar",
		}),
		Duration:            hanzoai.F("duration"),
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
		ObjectPermission: hanzoai.F(hanzoai.UserUpdateParamsObjectPermission{
			AgentAccessGroups: hanzoai.F([]string{"string"}),
			Agents:            hanzoai.F([]string{"string"}),
			McpAccessGroups:   hanzoai.F([]string{"string"}),
			McpServers:        hanzoai.F([]string{"string"}),
			McpToolPermissions: hanzoai.F(map[string][]string{
				"foo": {"string"},
			}),
			VectorStores: hanzoai.F([]string{"string"}),
		}),
		Password: hanzoai.F("password"),
		Permissions: hanzoai.F(map[string]interface{}{
			"foo": "bar",
		}),
		Prompts:   hanzoai.F([]string{"string"}),
		RpmLimit:  hanzoai.F(int64(0)),
		Spend:     hanzoai.F(0.000000),
		TeamID:    hanzoai.F("team_id"),
		TpmLimit:  hanzoai.F(int64(0)),
		UserAlias: hanzoai.F("user_alias"),
		UserEmail: hanzoai.F("user_email"),
		UserID:    hanzoai.F("user_id"),
		UserRole:  hanzoai.F(hanzoai.UserUpdateParamsUserRoleProxyAdmin),
	})
	if err != nil {
		var apierr *hanzoai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.User.Delete(context.TODO(), hanzoai.UserDeleteParams{
		UserIDs:          hanzoai.F([]string{"string"}),
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

func TestUserGetInfoWithOptionalParams(t *testing.T) {
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
	_, err := client.User.GetInfo(context.TODO(), hanzoai.UserGetInfoParams{
		UserID: hanzoai.F("user_id"),
	})
	if err != nil {
		var apierr *hanzoai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
