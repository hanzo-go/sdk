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

func TestTeamNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Team.New(context.TODO(), hanzoai.TeamNewParams{
		Admins:                   hanzoai.F([]interface{}{map[string]interface{}{}}),
		AllowedPassthroughRoutes: hanzoai.F([]interface{}{map[string]interface{}{}}),
		AllowedVectorStoreIndexes: hanzoai.F([]hanzoai.TeamNewParamsAllowedVectorStoreIndex{{
			IndexName:        hanzoai.F("index_name"),
			IndexPermissions: hanzoai.F([]hanzoai.TeamNewParamsAllowedVectorStoreIndexesIndexPermission{hanzoai.TeamNewParamsAllowedVectorStoreIndexesIndexPermissionRead}),
		}}),
		Blocked:        hanzoai.F(true),
		BudgetDuration: hanzoai.F("budget_duration"),
		Guardrails:     hanzoai.F([]string{"string"}),
		MaxBudget:      hanzoai.F(0.000000),
		Members:        hanzoai.F([]interface{}{map[string]interface{}{}}),
		MembersWithRoles: hanzoai.F([]hanzoai.MemberParam{{
			Role:      hanzoai.F(hanzoai.MemberRoleAdmin),
			UserEmail: hanzoai.F("user_email"),
			UserID:    hanzoai.F("user_id"),
		}}),
		Metadata: hanzoai.F(map[string]interface{}{
			"foo": "bar",
		}),
		ModelAliases: hanzoai.F(map[string]interface{}{
			"foo": "bar",
		}),
		ModelRpmLimit: hanzoai.F(map[string]int64{
			"foo": int64(0),
		}),
		ModelTpmLimit: hanzoai.F(map[string]int64{
			"foo": int64(0),
		}),
		Models: hanzoai.F([]interface{}{map[string]interface{}{}}),
		ObjectPermission: hanzoai.F(hanzoai.TeamNewParamsObjectPermission{
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
		Prompts:        hanzoai.F([]string{"string"}),
		RouterSettings: hanzoai.F(map[string]interface{}{
			"foo": "bar",
		}),
		RpmLimit:     hanzoai.F(int64(0)),
		RpmLimitType: hanzoai.F(hanzoai.TeamNewParamsRpmLimitTypeGuaranteedThroughput),
		SecretManagerSettings: hanzoai.F(map[string]interface{}{
			"foo": "bar",
		}),
		Tags:                  hanzoai.F([]interface{}{map[string]interface{}{}}),
		TeamAlias:             hanzoai.F("team_alias"),
		TeamID:                hanzoai.F("team_id"),
		TeamMemberBudget:      hanzoai.F(0.000000),
		TeamMemberKeyDuration: hanzoai.F("team_member_key_duration"),
		TeamMemberPermissions: hanzoai.F([]string{"string"}),
		TeamMemberRpmLimit:    hanzoai.F(int64(0)),
		TeamMemberTpmLimit:    hanzoai.F(int64(0)),
		TpmLimit:              hanzoai.F(int64(0)),
		TpmLimitType:          hanzoai.F(hanzoai.TeamNewParamsTpmLimitTypeGuaranteedThroughput),
		LitellmChangedBy:      hanzoai.F("litellm-changed-by"),
	})
	if err != nil {
		var apierr *hanzoai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTeamUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Team.Update(context.TODO(), hanzoai.TeamUpdateParams{
		TeamID:                   hanzoai.F("team_id"),
		AllowedPassthroughRoutes: hanzoai.F([]interface{}{map[string]interface{}{}}),
		AllowedVectorStoreIndexes: hanzoai.F([]hanzoai.TeamUpdateParamsAllowedVectorStoreIndex{{
			IndexName:        hanzoai.F("index_name"),
			IndexPermissions: hanzoai.F([]hanzoai.TeamUpdateParamsAllowedVectorStoreIndexesIndexPermission{hanzoai.TeamUpdateParamsAllowedVectorStoreIndexesIndexPermissionRead}),
		}}),
		Blocked:        hanzoai.F(true),
		BudgetDuration: hanzoai.F("budget_duration"),
		Guardrails:     hanzoai.F([]string{"string"}),
		MaxBudget:      hanzoai.F(0.000000),
		Metadata: hanzoai.F(map[string]interface{}{
			"foo": "bar",
		}),
		ModelAliases: hanzoai.F(map[string]interface{}{
			"foo": "bar",
		}),
		ModelRpmLimit: hanzoai.F(map[string]int64{
			"foo": int64(0),
		}),
		ModelTpmLimit: hanzoai.F(map[string]int64{
			"foo": int64(0),
		}),
		Models: hanzoai.F([]interface{}{map[string]interface{}{}}),
		ObjectPermission: hanzoai.F(hanzoai.TeamUpdateParamsObjectPermission{
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
		Prompts:        hanzoai.F([]string{"string"}),
		RouterSettings: hanzoai.F(map[string]interface{}{
			"foo": "bar",
		}),
		RpmLimit: hanzoai.F(int64(0)),
		SecretManagerSettings: hanzoai.F(map[string]interface{}{
			"foo": "bar",
		}),
		Tags:                     hanzoai.F([]interface{}{map[string]interface{}{}}),
		TeamAlias:                hanzoai.F("team_alias"),
		TeamMemberBudget:         hanzoai.F(0.000000),
		TeamMemberBudgetDuration: hanzoai.F("team_member_budget_duration"),
		TeamMemberKeyDuration:    hanzoai.F("team_member_key_duration"),
		TeamMemberRpmLimit:       hanzoai.F(int64(0)),
		TeamMemberTpmLimit:       hanzoai.F(int64(0)),
		TpmLimit:                 hanzoai.F(int64(0)),
		LitellmChangedBy:         hanzoai.F("litellm-changed-by"),
	})
	if err != nil {
		var apierr *hanzoai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTeamListWithOptionalParams(t *testing.T) {
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
	_, err := client.Team.List(context.TODO(), hanzoai.TeamListParams{
		OrganizationID: hanzoai.F("organization_id"),
		UserID:         hanzoai.F("user_id"),
	})
	if err != nil {
		var apierr *hanzoai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTeamDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.Team.Delete(context.TODO(), hanzoai.TeamDeleteParams{
		TeamIDs:          hanzoai.F([]string{"string"}),
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

func TestTeamAddMemberWithOptionalParams(t *testing.T) {
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
	_, err := client.Team.AddMember(context.TODO(), hanzoai.TeamAddMemberParams{
		Member: hanzoai.F[hanzoai.TeamAddMemberParamsMemberUnion](hanzoai.TeamAddMemberParamsMemberArray([]hanzoai.MemberParam{{
			Role:      hanzoai.F(hanzoai.MemberRoleAdmin),
			UserEmail: hanzoai.F("user_email"),
			UserID:    hanzoai.F("user_id"),
		}})),
		TeamID:          hanzoai.F("team_id"),
		MaxBudgetInTeam: hanzoai.F(0.000000),
	})
	if err != nil {
		var apierr *hanzoai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTeamBlock(t *testing.T) {
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
	_, err := client.Team.Block(context.TODO(), hanzoai.TeamBlockParams{
		BlockTeamRequest: hanzoai.BlockTeamRequestParam{
			TeamID: hanzoai.F("team_id"),
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

func TestTeamDisableLogging(t *testing.T) {
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
	_, err := client.Team.DisableLogging(context.TODO(), "team_id")
	if err != nil {
		var apierr *hanzoai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTeamListAvailableWithOptionalParams(t *testing.T) {
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
	_, err := client.Team.ListAvailable(context.TODO(), hanzoai.TeamListAvailableParams{
		ResponseModel: hanzoai.F[any](map[string]interface{}{}),
	})
	if err != nil {
		var apierr *hanzoai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTeamRemoveMemberWithOptionalParams(t *testing.T) {
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
	_, err := client.Team.RemoveMember(context.TODO(), hanzoai.TeamRemoveMemberParams{
		TeamID:    hanzoai.F("team_id"),
		UserEmail: hanzoai.F("user_email"),
		UserID:    hanzoai.F("user_id"),
	})
	if err != nil {
		var apierr *hanzoai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTeamGetInfoWithOptionalParams(t *testing.T) {
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
	_, err := client.Team.GetInfo(context.TODO(), hanzoai.TeamGetInfoParams{
		TeamID: hanzoai.F("team_id"),
	})
	if err != nil {
		var apierr *hanzoai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTeamUnblock(t *testing.T) {
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
	_, err := client.Team.Unblock(context.TODO(), hanzoai.TeamUnblockParams{
		BlockTeamRequest: hanzoai.BlockTeamRequestParam{
			TeamID: hanzoai.F("team_id"),
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

func TestTeamUpdateMemberWithOptionalParams(t *testing.T) {
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
	_, err := client.Team.UpdateMember(context.TODO(), hanzoai.TeamUpdateMemberParams{
		TeamID:          hanzoai.F("team_id"),
		MaxBudgetInTeam: hanzoai.F(0.000000),
		Role:            hanzoai.F(hanzoai.TeamUpdateMemberParamsRoleAdmin),
		RpmLimit:        hanzoai.F(int64(0)),
		TpmLimit:        hanzoai.F(int64(0)),
		UserEmail:       hanzoai.F("user_email"),
		UserID:          hanzoai.F("user_id"),
	})
	if err != nil {
		var apierr *hanzoai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
