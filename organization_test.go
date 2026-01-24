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

func TestOrganizationNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Organization.New(context.TODO(), hanzoai.OrganizationNewParams{
		OrganizationAlias:   hanzoai.F("organization_alias"),
		BudgetDuration:      hanzoai.F("budget_duration"),
		BudgetID:            hanzoai.F("budget_id"),
		MaxBudget:           hanzoai.F(0.000000),
		MaxParallelRequests: hanzoai.F(int64(0)),
		Metadata: hanzoai.F(map[string]interface{}{
			"foo": "bar",
		}),
		ModelMaxBudget: hanzoai.F(map[string]interface{}{
			"foo": "bar",
		}),
		ModelRpmLimit: hanzoai.F(map[string]int64{
			"foo": int64(0),
		}),
		ModelTpmLimit: hanzoai.F(map[string]int64{
			"foo": int64(0),
		}),
		Models: hanzoai.F([]interface{}{map[string]interface{}{}}),
		ObjectPermission: hanzoai.F(hanzoai.OrganizationNewParamsObjectPermission{
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
		RpmLimit:       hanzoai.F(int64(0)),
		SoftBudget:     hanzoai.F(0.000000),
		TpmLimit:       hanzoai.F(int64(0)),
	})
	if err != nil {
		var apierr *hanzoai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOrganizationUpdate(t *testing.T) {
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
	_, err := client.Organization.Update(context.TODO())
	if err != nil {
		var apierr *hanzoai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOrganizationListWithOptionalParams(t *testing.T) {
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
	_, err := client.Organization.List(context.TODO(), hanzoai.OrganizationListParams{
		OrgAlias: hanzoai.F("org_alias"),
		OrgID:    hanzoai.F("org_id"),
	})
	if err != nil {
		var apierr *hanzoai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOrganizationDelete(t *testing.T) {
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
	_, err := client.Organization.Delete(context.TODO(), hanzoai.OrganizationDeleteParams{
		OrganizationIDs: hanzoai.F([]string{"string"}),
	})
	if err != nil {
		var apierr *hanzoai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOrganizationAddMemberWithOptionalParams(t *testing.T) {
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
	_, err := client.Organization.AddMember(context.TODO(), hanzoai.OrganizationAddMemberParams{
		Member: hanzoai.F[hanzoai.OrganizationAddMemberParamsMemberUnion](hanzoai.OrganizationAddMemberParamsMemberArray([]hanzoai.OrgMemberParam{{
			Role:      hanzoai.F(hanzoai.OrgMemberRoleOrgAdmin),
			UserEmail: hanzoai.F("user_email"),
			UserID:    hanzoai.F("user_id"),
		}})),
		OrganizationID:          hanzoai.F("organization_id"),
		MaxBudgetInOrganization: hanzoai.F(0.000000),
	})
	if err != nil {
		var apierr *hanzoai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOrganizationDeleteMemberWithOptionalParams(t *testing.T) {
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
	_, err := client.Organization.DeleteMember(context.TODO(), hanzoai.OrganizationDeleteMemberParams{
		OrganizationID: hanzoai.F("organization_id"),
		UserEmail:      hanzoai.F("user_email"),
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

func TestOrganizationUpdateMemberWithOptionalParams(t *testing.T) {
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
	_, err := client.Organization.UpdateMember(context.TODO(), hanzoai.OrganizationUpdateMemberParams{
		OrganizationID:          hanzoai.F("organization_id"),
		MaxBudgetInOrganization: hanzoai.F(0.000000),
		Role:                    hanzoai.F(hanzoai.UserRolesProxyAdmin),
		UserEmail:               hanzoai.F("user_email"),
		UserID:                  hanzoai.F("user_id"),
	})
	if err != nil {
		var apierr *hanzoai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
