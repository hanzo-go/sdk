// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hanzoai

import (
	"context"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/hanzoai/go-sdk/internal/apijson"
	"github.com/hanzoai/go-sdk/internal/apiquery"
	"github.com/hanzoai/go-sdk/internal/param"
	"github.com/hanzoai/go-sdk/internal/requestconfig"
	"github.com/hanzoai/go-sdk/option"
)

// OrganizationService contains methods and other services that help with
// interacting with the Hanzo API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrganizationService] method instead.
type OrganizationService struct {
	Options []option.RequestOption
	Info    *OrganizationInfoService
}

// NewOrganizationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewOrganizationService(opts ...option.RequestOption) (r *OrganizationService) {
	r = &OrganizationService{}
	r.Options = opts
	r.Info = NewOrganizationInfoService(opts...)
	return
}

// Allow orgs to own teams
//
// Set org level budgets + model access.
//
// Only admins can create orgs.
//
// # Parameters
//
//   - organization_alias: _str_ - The name of the organization.
//   - models: _List_ - The models the organization has access to.
//   - budget_id: _Optional[str]_ - The id for a budget (tpm/rpm/max budget) for the
//     organization.
//
// ### IF NO BUDGET ID - CREATE ONE WITH THESE PARAMS
//
//   - max_budget: _Optional[float]_ - Max budget for org
//   - tpm_limit: _Optional[int]_ - Max tpm limit for org
//   - rpm_limit: _Optional[int]_ - Max rpm limit for org
//   - model_rpm_limit: _Optional[Dict[str, int]]_ - The RPM (Requests Per Minute)
//     limit per model for this organization.
//   - model_tpm_limit: _Optional[Dict[str, int]]_ - The TPM (Tokens Per Minute)
//     limit per model for this organization.
//   - max_parallel_requests: _Optional[int]_ - [Not Implemented Yet] Max parallel
//     requests for org
//   - soft_budget: _Optional[float]_ - [Not Implemented Yet] Get a slack alert when
//     this soft budget is reached. Don't block requests.
//   - model_max_budget: _Optional[dict]_ - Max budget for a specific model
//   - budget_duration: _Optional[str]_ - Frequency of reseting org budget
//   - metadata: _Optional[dict]_ - Metadata for organization, store information for
//     organization. Example metadata - {"extra_info": "some info"}
//   - blocked: _bool_ - Flag indicating if the org is blocked or not - will stop all
//     calls from keys with this org_id.
//   - tags: _Optional[List[str]]_ - Tags for
//     [tracking spend](https://litellm.vercel.app/docs/proxy/enterprise#tracking-spend-for-custom-tags)
//     and/or doing
//     [tag-based routing](https://litellm.vercel.app/docs/proxy/tag_routing).
//   - organization_id: _Optional[str]_ - The organization id of the team. Default is
//     None. Create via `/organization/new`.
//   - model_aliases: Optional[dict] - Model aliases for the team.
//     [Docs](https://docs.litellm.ai/docs/proxy/team_based_routing#create-team-with-model-alias)
//   - object_permission: Optional[LiteLLM_ObjectPermissionBase] -
//     organization-specific object permission. Example - {"vector_stores":
//     ["vector_store_1", "vector_store_2"]}. IF null or {} then no object
//     permission. Case 1: Create new org **without** a budget_id
//
// ```bash
// curl --location 'http://0.0.0.0:4000/organization/new'
// --header 'Authorization: Bearer sk-1234'
// --header 'Content-Type: application/json'
//
//	--data '{
//	    "organization_alias": "my-secret-org",
//	    "models": ["model1", "model2"],
//	    "max_budget": 100
//	}'
//
// ```
//
// Case 2: Create new org **with** a budget_id
//
// ```bash
// curl --location 'http://0.0.0.0:4000/organization/new'
// --header 'Authorization: Bearer sk-1234'
// --header 'Content-Type: application/json'
//
//	--data '{
//	    "organization_alias": "my-secret-org",
//	    "models": ["model1", "model2"],
//	    "budget_id": "428eeaa8-f3ac-4e85-a8fb-7dc8d7aa8689"
//	}'
//
// ```
func (r *OrganizationService) New(ctx context.Context, body OrganizationNewParams, opts ...option.RequestOption) (res *OrganizationNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "organization/new"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Update an organization
func (r *OrganizationService) Update(ctx context.Context, opts ...option.RequestOption) (res *OrganizationTableWithMembers, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "organization/update"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, nil, &res, opts...)
	return
}

// Get a list of organizations with optional filtering.
//
// Parameters: org_id: Optional[str] Filter organizations by exact organization_id
// match org_alias: Optional[str] Filter organizations by partial
// organization_alias match (case-insensitive)
//
// Example:
//
// ```
// curl --location --request GET 'http://0.0.0.0:4000/organization/list?org_alias=my-org'         --header 'Authorization: Bearer sk-1234'
// ```
//
// Example with org_id:
//
// ```
// curl --location --request GET 'http://0.0.0.0:4000/organization/list?org_id=123e4567-e89b-12d3-a456-426614174000'         --header 'Authorization: Bearer sk-1234'
// ```
func (r *OrganizationService) List(ctx context.Context, query OrganizationListParams, opts ...option.RequestOption) (res *[]OrganizationTableWithMembers, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "organization/list"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete an organization
//
// # Parameters:
//
// - organization_ids: List[str] - The organization ids to delete.
func (r *OrganizationService) Delete(ctx context.Context, body OrganizationDeleteParams, opts ...option.RequestOption) (res *[]OrganizationTableWithMembers, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "organization/delete"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// [BETA]
//
// Add new members (either via user_email or user_id) to an organization
//
// If user doesn't exist, new user row will also be added to User Table
//
// Only proxy_admin or org_admin of organization, allowed to access this endpoint.
//
// # Parameters:
//
// - organization_id: str (required)
// - member: Union[List[Member], Member] (required)
//   - role: Literal[LitellmUserRoles] (required)
//   - user_id: Optional[str]
//   - user_email: Optional[str]
//
// Note: Either user_id or user_email must be provided for each member.
//
// Example:
//
// ```
//
//	curl -X POST 'http://0.0.0.0:4000/organization/member_add'     -H 'Authorization: Bearer sk-1234'     -H 'Content-Type: application/json'     -d '{
//	    "organization_id": "45e3e396-ee08-4a61-a88e-16b3ce7e0849",
//	    "member": {
//	        "role": "internal_user",
//	        "user_id": "krrish247652@berri.ai"
//	    },
//	    "max_budget_in_organization": 100.0
//	}'
//
// ```
//
// The following is executed in this function:
//
//  1. Check if organization exists
//  2. Creates a new Internal User if the user_id or user_email is not found in
//     LiteLLM_UserTable
//  3. Add Internal User to the `LiteLLM_OrganizationMembership` table
func (r *OrganizationService) AddMember(ctx context.Context, body OrganizationAddMemberParams, opts ...option.RequestOption) (res *OrganizationAddMemberResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "organization/member_add"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Delete a member from an organization
func (r *OrganizationService) DeleteMember(ctx context.Context, body OrganizationDeleteMemberParams, opts ...option.RequestOption) (res *OrganizationDeleteMemberResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "organization/member_delete"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// Update a member's role in an organization
func (r *OrganizationService) UpdateMember(ctx context.Context, body OrganizationUpdateMemberParams, opts ...option.RequestOption) (res *OrganizationMembershipTable, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "organization/member_update"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Represents user-controllable params for a LiteLLM_BudgetTable record
type BudgetTable struct {
	BudgetDuration      string                 `json:"budget_duration,nullable"`
	BudgetID            string                 `json:"budget_id,nullable"`
	MaxBudget           float64                `json:"max_budget,nullable"`
	MaxParallelRequests int64                  `json:"max_parallel_requests,nullable"`
	ModelMaxBudget      map[string]interface{} `json:"model_max_budget,nullable"`
	RpmLimit            int64                  `json:"rpm_limit,nullable"`
	SoftBudget          float64                `json:"soft_budget,nullable"`
	TpmLimit            int64                  `json:"tpm_limit,nullable"`
	JSON                budgetTableJSON        `json:"-"`
}

// budgetTableJSON contains the JSON metadata for the struct [BudgetTable]
type budgetTableJSON struct {
	BudgetDuration      apijson.Field
	BudgetID            apijson.Field
	MaxBudget           apijson.Field
	MaxParallelRequests apijson.Field
	ModelMaxBudget      apijson.Field
	RpmLimit            apijson.Field
	SoftBudget          apijson.Field
	TpmLimit            apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *BudgetTable) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r budgetTableJSON) RawJSON() string {
	return r.raw
}

type OrgMemberParam struct {
	Role param.Field[OrgMemberRole] `json:"role,required"`
	// The email address of the user to add. Either user_id or user_email must be
	// provided
	UserEmail param.Field[string] `json:"user_email"`
	// The unique ID of the user to add. Either user_id or user_email must be provided
	UserID param.Field[string] `json:"user_id"`
}

func (r OrgMemberParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r OrgMemberParam) implementsOrganizationAddMemberParamsMemberUnion() {}

type OrgMemberRole string

const (
	OrgMemberRoleOrgAdmin           OrgMemberRole = "org_admin"
	OrgMemberRoleInternalUser       OrgMemberRole = "internal_user"
	OrgMemberRoleInternalUserViewer OrgMemberRole = "internal_user_viewer"
)

func (r OrgMemberRole) IsKnown() bool {
	switch r {
	case OrgMemberRoleOrgAdmin, OrgMemberRoleInternalUser, OrgMemberRoleInternalUserViewer:
		return true
	}
	return false
}

// This is the table that track what organizations a user belongs to and users
// spend within the organization
type OrganizationMembershipTable struct {
	CreatedAt      time.Time `json:"created_at,required" format:"date-time"`
	OrganizationID string    `json:"organization_id,required"`
	UpdatedAt      time.Time `json:"updated_at,required" format:"date-time"`
	UserID         string    `json:"user_id,required"`
	BudgetID       string    `json:"budget_id,nullable"`
	// Represents user-controllable params for a LiteLLM_BudgetTable record
	LitellmBudgetTable BudgetTable                     `json:"litellm_budget_table,nullable"`
	Spend              float64                         `json:"spend"`
	User               interface{}                     `json:"user"`
	UserRole           string                          `json:"user_role,nullable"`
	JSON               organizationMembershipTableJSON `json:"-"`
}

// organizationMembershipTableJSON contains the JSON metadata for the struct
// [OrganizationMembershipTable]
type organizationMembershipTableJSON struct {
	CreatedAt          apijson.Field
	OrganizationID     apijson.Field
	UpdatedAt          apijson.Field
	UserID             apijson.Field
	BudgetID           apijson.Field
	LitellmBudgetTable apijson.Field
	Spend              apijson.Field
	User               apijson.Field
	UserRole           apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *OrganizationMembershipTable) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationMembershipTableJSON) RawJSON() string {
	return r.raw
}

// Returned by the /organization/info endpoint and /organization/list endpoint
type OrganizationTableWithMembers struct {
	BudgetID  string    `json:"budget_id,required"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	CreatedBy string    `json:"created_by,required"`
	Models    []string  `json:"models,required"`
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	UpdatedBy string    `json:"updated_by,required"`
	// Represents user-controllable params for a LiteLLM_BudgetTable record
	LitellmBudgetTable BudgetTable                   `json:"litellm_budget_table,nullable"`
	Members            []OrganizationMembershipTable `json:"members"`
	Metadata           map[string]interface{}        `json:"metadata,nullable"`
	// Represents a LiteLLM_ObjectPermissionTable record
	ObjectPermission   OrganizationTableWithMembersObjectPermission `json:"object_permission,nullable"`
	ObjectPermissionID string                                       `json:"object_permission_id,nullable"`
	OrganizationAlias  string                                       `json:"organization_alias,nullable"`
	OrganizationID     string                                       `json:"organization_id,nullable"`
	Spend              float64                                      `json:"spend"`
	Teams              []OrganizationTableWithMembersTeam           `json:"teams"`
	Users              []OrganizationTableWithMembersUser           `json:"users,nullable"`
	JSON               organizationTableWithMembersJSON             `json:"-"`
}

// organizationTableWithMembersJSON contains the JSON metadata for the struct
// [OrganizationTableWithMembers]
type organizationTableWithMembersJSON struct {
	BudgetID           apijson.Field
	CreatedAt          apijson.Field
	CreatedBy          apijson.Field
	Models             apijson.Field
	UpdatedAt          apijson.Field
	UpdatedBy          apijson.Field
	LitellmBudgetTable apijson.Field
	Members            apijson.Field
	Metadata           apijson.Field
	ObjectPermission   apijson.Field
	ObjectPermissionID apijson.Field
	OrganizationAlias  apijson.Field
	OrganizationID     apijson.Field
	Spend              apijson.Field
	Teams              apijson.Field
	Users              apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *OrganizationTableWithMembers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationTableWithMembersJSON) RawJSON() string {
	return r.raw
}

// Represents a LiteLLM_ObjectPermissionTable record
type OrganizationTableWithMembersObjectPermission struct {
	ObjectPermissionID string                                           `json:"object_permission_id,required"`
	AgentAccessGroups  []string                                         `json:"agent_access_groups,nullable"`
	Agents             []string                                         `json:"agents,nullable"`
	McpAccessGroups    []string                                         `json:"mcp_access_groups,nullable"`
	McpServers         []string                                         `json:"mcp_servers,nullable"`
	McpToolPermissions map[string][]string                              `json:"mcp_tool_permissions,nullable"`
	VectorStores       []string                                         `json:"vector_stores,nullable"`
	JSON               organizationTableWithMembersObjectPermissionJSON `json:"-"`
}

// organizationTableWithMembersObjectPermissionJSON contains the JSON metadata for
// the struct [OrganizationTableWithMembersObjectPermission]
type organizationTableWithMembersObjectPermissionJSON struct {
	ObjectPermissionID apijson.Field
	AgentAccessGroups  apijson.Field
	Agents             apijson.Field
	McpAccessGroups    apijson.Field
	McpServers         apijson.Field
	McpToolPermissions apijson.Field
	VectorStores       apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *OrganizationTableWithMembersObjectPermission) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationTableWithMembersObjectPermissionJSON) RawJSON() string {
	return r.raw
}

type OrganizationTableWithMembersTeam struct {
	TeamID              string                 `json:"team_id,required"`
	Admins              []interface{}          `json:"admins"`
	Blocked             bool                   `json:"blocked"`
	BudgetDuration      string                 `json:"budget_duration,nullable"`
	BudgetResetAt       time.Time              `json:"budget_reset_at,nullable" format:"date-time"`
	CreatedAt           time.Time              `json:"created_at,nullable" format:"date-time"`
	LitellmModelTable   interface{}            `json:"litellm_model_table"`
	MaxBudget           float64                `json:"max_budget,nullable"`
	MaxParallelRequests int64                  `json:"max_parallel_requests,nullable"`
	Members             []interface{}          `json:"members"`
	MembersWithRoles    []Member               `json:"members_with_roles"`
	Metadata            map[string]interface{} `json:"metadata,nullable"`
	ModelID             int64                  `json:"model_id,nullable"`
	Models              []interface{}          `json:"models"`
	// Represents a LiteLLM_ObjectPermissionTable record
	ObjectPermission      OrganizationTableWithMembersTeamsObjectPermission `json:"object_permission,nullable"`
	ObjectPermissionID    string                                            `json:"object_permission_id,nullable"`
	OrganizationID        string                                            `json:"organization_id,nullable"`
	RouterSettings        map[string]interface{}                            `json:"router_settings,nullable"`
	RpmLimit              int64                                             `json:"rpm_limit,nullable"`
	Spend                 float64                                           `json:"spend,nullable"`
	TeamAlias             string                                            `json:"team_alias,nullable"`
	TeamMemberPermissions []string                                          `json:"team_member_permissions,nullable"`
	TpmLimit              int64                                             `json:"tpm_limit,nullable"`
	UpdatedAt             time.Time                                         `json:"updated_at,nullable" format:"date-time"`
	JSON                  organizationTableWithMembersTeamJSON              `json:"-"`
}

// organizationTableWithMembersTeamJSON contains the JSON metadata for the struct
// [OrganizationTableWithMembersTeam]
type organizationTableWithMembersTeamJSON struct {
	TeamID                apijson.Field
	Admins                apijson.Field
	Blocked               apijson.Field
	BudgetDuration        apijson.Field
	BudgetResetAt         apijson.Field
	CreatedAt             apijson.Field
	LitellmModelTable     apijson.Field
	MaxBudget             apijson.Field
	MaxParallelRequests   apijson.Field
	Members               apijson.Field
	MembersWithRoles      apijson.Field
	Metadata              apijson.Field
	ModelID               apijson.Field
	Models                apijson.Field
	ObjectPermission      apijson.Field
	ObjectPermissionID    apijson.Field
	OrganizationID        apijson.Field
	RouterSettings        apijson.Field
	RpmLimit              apijson.Field
	Spend                 apijson.Field
	TeamAlias             apijson.Field
	TeamMemberPermissions apijson.Field
	TpmLimit              apijson.Field
	UpdatedAt             apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *OrganizationTableWithMembersTeam) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationTableWithMembersTeamJSON) RawJSON() string {
	return r.raw
}

// Represents a LiteLLM_ObjectPermissionTable record
type OrganizationTableWithMembersTeamsObjectPermission struct {
	ObjectPermissionID string                                                `json:"object_permission_id,required"`
	AgentAccessGroups  []string                                              `json:"agent_access_groups,nullable"`
	Agents             []string                                              `json:"agents,nullable"`
	McpAccessGroups    []string                                              `json:"mcp_access_groups,nullable"`
	McpServers         []string                                              `json:"mcp_servers,nullable"`
	McpToolPermissions map[string][]string                                   `json:"mcp_tool_permissions,nullable"`
	VectorStores       []string                                              `json:"vector_stores,nullable"`
	JSON               organizationTableWithMembersTeamsObjectPermissionJSON `json:"-"`
}

// organizationTableWithMembersTeamsObjectPermissionJSON contains the JSON metadata
// for the struct [OrganizationTableWithMembersTeamsObjectPermission]
type organizationTableWithMembersTeamsObjectPermissionJSON struct {
	ObjectPermissionID apijson.Field
	AgentAccessGroups  apijson.Field
	Agents             apijson.Field
	McpAccessGroups    apijson.Field
	McpServers         apijson.Field
	McpToolPermissions apijson.Field
	VectorStores       apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *OrganizationTableWithMembersTeamsObjectPermission) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationTableWithMembersTeamsObjectPermissionJSON) RawJSON() string {
	return r.raw
}

type OrganizationTableWithMembersUser struct {
	UserID         string                 `json:"user_id,required"`
	BudgetDuration string                 `json:"budget_duration,nullable"`
	BudgetResetAt  time.Time              `json:"budget_reset_at,nullable" format:"date-time"`
	CreatedAt      time.Time              `json:"created_at,nullable" format:"date-time"`
	MaxBudget      float64                `json:"max_budget,nullable"`
	Metadata       map[string]interface{} `json:"metadata,nullable"`
	ModelMaxBudget map[string]interface{} `json:"model_max_budget,nullable"`
	ModelSpend     map[string]interface{} `json:"model_spend,nullable"`
	Models         []interface{}          `json:"models"`
	// Represents a LiteLLM_ObjectPermissionTable record
	ObjectPermission        OrganizationTableWithMembersUsersObjectPermission `json:"object_permission,nullable"`
	OrganizationMemberships []OrganizationMembershipTable                     `json:"organization_memberships,nullable"`
	RpmLimit                int64                                             `json:"rpm_limit,nullable"`
	Spend                   float64                                           `json:"spend"`
	SSOUserID               string                                            `json:"sso_user_id,nullable"`
	Teams                   []string                                          `json:"teams"`
	TpmLimit                int64                                             `json:"tpm_limit,nullable"`
	UpdatedAt               time.Time                                         `json:"updated_at,nullable" format:"date-time"`
	UserAlias               string                                            `json:"user_alias,nullable"`
	UserEmail               string                                            `json:"user_email,nullable"`
	UserRole                string                                            `json:"user_role,nullable"`
	JSON                    organizationTableWithMembersUserJSON              `json:"-"`
}

// organizationTableWithMembersUserJSON contains the JSON metadata for the struct
// [OrganizationTableWithMembersUser]
type organizationTableWithMembersUserJSON struct {
	UserID                  apijson.Field
	BudgetDuration          apijson.Field
	BudgetResetAt           apijson.Field
	CreatedAt               apijson.Field
	MaxBudget               apijson.Field
	Metadata                apijson.Field
	ModelMaxBudget          apijson.Field
	ModelSpend              apijson.Field
	Models                  apijson.Field
	ObjectPermission        apijson.Field
	OrganizationMemberships apijson.Field
	RpmLimit                apijson.Field
	Spend                   apijson.Field
	SSOUserID               apijson.Field
	Teams                   apijson.Field
	TpmLimit                apijson.Field
	UpdatedAt               apijson.Field
	UserAlias               apijson.Field
	UserEmail               apijson.Field
	UserRole                apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *OrganizationTableWithMembersUser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationTableWithMembersUserJSON) RawJSON() string {
	return r.raw
}

// Represents a LiteLLM_ObjectPermissionTable record
type OrganizationTableWithMembersUsersObjectPermission struct {
	ObjectPermissionID string                                                `json:"object_permission_id,required"`
	AgentAccessGroups  []string                                              `json:"agent_access_groups,nullable"`
	Agents             []string                                              `json:"agents,nullable"`
	McpAccessGroups    []string                                              `json:"mcp_access_groups,nullable"`
	McpServers         []string                                              `json:"mcp_servers,nullable"`
	McpToolPermissions map[string][]string                                   `json:"mcp_tool_permissions,nullable"`
	VectorStores       []string                                              `json:"vector_stores,nullable"`
	JSON               organizationTableWithMembersUsersObjectPermissionJSON `json:"-"`
}

// organizationTableWithMembersUsersObjectPermissionJSON contains the JSON metadata
// for the struct [OrganizationTableWithMembersUsersObjectPermission]
type organizationTableWithMembersUsersObjectPermissionJSON struct {
	ObjectPermissionID apijson.Field
	AgentAccessGroups  apijson.Field
	Agents             apijson.Field
	McpAccessGroups    apijson.Field
	McpServers         apijson.Field
	McpToolPermissions apijson.Field
	VectorStores       apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *OrganizationTableWithMembersUsersObjectPermission) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationTableWithMembersUsersObjectPermissionJSON) RawJSON() string {
	return r.raw
}

// Admin Roles: PROXY_ADMIN: admin over the platform PROXY_ADMIN_VIEW_ONLY: can
// login, view all own keys, view all spend ORG_ADMIN: admin over a specific
// organization, can create teams, users only within their organization
//
// Internal User Roles: INTERNAL_USER: can login, view/create/delete their own
// keys, view their spend INTERNAL_USER_VIEW_ONLY: can login, view their own keys,
// view their own spend
//
// Team Roles: TEAM: used for JWT auth
//
// Customer Roles: CUSTOMER: External users -> these are customers
type UserRoles string

const (
	UserRolesProxyAdmin         UserRoles = "proxy_admin"
	UserRolesProxyAdminViewer   UserRoles = "proxy_admin_viewer"
	UserRolesOrgAdmin           UserRoles = "org_admin"
	UserRolesInternalUser       UserRoles = "internal_user"
	UserRolesInternalUserViewer UserRoles = "internal_user_viewer"
	UserRolesTeam               UserRoles = "team"
	UserRolesCustomer           UserRoles = "customer"
)

func (r UserRoles) IsKnown() bool {
	switch r {
	case UserRolesProxyAdmin, UserRolesProxyAdminViewer, UserRolesOrgAdmin, UserRolesInternalUser, UserRolesInternalUserViewer, UserRolesTeam, UserRolesCustomer:
		return true
	}
	return false
}

type OrganizationNewResponse struct {
	BudgetID       string    `json:"budget_id,required"`
	CreatedAt      time.Time `json:"created_at,required" format:"date-time"`
	CreatedBy      string    `json:"created_by,required"`
	Models         []string  `json:"models,required"`
	OrganizationID string    `json:"organization_id,required"`
	UpdatedAt      time.Time `json:"updated_at,required" format:"date-time"`
	UpdatedBy      string    `json:"updated_by,required"`
	// Represents user-controllable params for a LiteLLM_BudgetTable record
	LitellmBudgetTable BudgetTable            `json:"litellm_budget_table,nullable"`
	Metadata           map[string]interface{} `json:"metadata,nullable"`
	// Represents a LiteLLM_ObjectPermissionTable record
	ObjectPermission   OrganizationNewResponseObjectPermission `json:"object_permission,nullable"`
	ObjectPermissionID string                                  `json:"object_permission_id,nullable"`
	OrganizationAlias  string                                  `json:"organization_alias,nullable"`
	Spend              float64                                 `json:"spend"`
	Users              []OrganizationNewResponseUser           `json:"users,nullable"`
	JSON               organizationNewResponseJSON             `json:"-"`
}

// organizationNewResponseJSON contains the JSON metadata for the struct
// [OrganizationNewResponse]
type organizationNewResponseJSON struct {
	BudgetID           apijson.Field
	CreatedAt          apijson.Field
	CreatedBy          apijson.Field
	Models             apijson.Field
	OrganizationID     apijson.Field
	UpdatedAt          apijson.Field
	UpdatedBy          apijson.Field
	LitellmBudgetTable apijson.Field
	Metadata           apijson.Field
	ObjectPermission   apijson.Field
	ObjectPermissionID apijson.Field
	OrganizationAlias  apijson.Field
	Spend              apijson.Field
	Users              apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *OrganizationNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationNewResponseJSON) RawJSON() string {
	return r.raw
}

// Represents a LiteLLM_ObjectPermissionTable record
type OrganizationNewResponseObjectPermission struct {
	ObjectPermissionID string                                      `json:"object_permission_id,required"`
	AgentAccessGroups  []string                                    `json:"agent_access_groups,nullable"`
	Agents             []string                                    `json:"agents,nullable"`
	McpAccessGroups    []string                                    `json:"mcp_access_groups,nullable"`
	McpServers         []string                                    `json:"mcp_servers,nullable"`
	McpToolPermissions map[string][]string                         `json:"mcp_tool_permissions,nullable"`
	VectorStores       []string                                    `json:"vector_stores,nullable"`
	JSON               organizationNewResponseObjectPermissionJSON `json:"-"`
}

// organizationNewResponseObjectPermissionJSON contains the JSON metadata for the
// struct [OrganizationNewResponseObjectPermission]
type organizationNewResponseObjectPermissionJSON struct {
	ObjectPermissionID apijson.Field
	AgentAccessGroups  apijson.Field
	Agents             apijson.Field
	McpAccessGroups    apijson.Field
	McpServers         apijson.Field
	McpToolPermissions apijson.Field
	VectorStores       apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *OrganizationNewResponseObjectPermission) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationNewResponseObjectPermissionJSON) RawJSON() string {
	return r.raw
}

type OrganizationNewResponseUser struct {
	UserID         string                 `json:"user_id,required"`
	BudgetDuration string                 `json:"budget_duration,nullable"`
	BudgetResetAt  time.Time              `json:"budget_reset_at,nullable" format:"date-time"`
	CreatedAt      time.Time              `json:"created_at,nullable" format:"date-time"`
	MaxBudget      float64                `json:"max_budget,nullable"`
	Metadata       map[string]interface{} `json:"metadata,nullable"`
	ModelMaxBudget map[string]interface{} `json:"model_max_budget,nullable"`
	ModelSpend     map[string]interface{} `json:"model_spend,nullable"`
	Models         []interface{}          `json:"models"`
	// Represents a LiteLLM_ObjectPermissionTable record
	ObjectPermission        OrganizationNewResponseUsersObjectPermission `json:"object_permission,nullable"`
	OrganizationMemberships []OrganizationMembershipTable                `json:"organization_memberships,nullable"`
	RpmLimit                int64                                        `json:"rpm_limit,nullable"`
	Spend                   float64                                      `json:"spend"`
	SSOUserID               string                                       `json:"sso_user_id,nullable"`
	Teams                   []string                                     `json:"teams"`
	TpmLimit                int64                                        `json:"tpm_limit,nullable"`
	UpdatedAt               time.Time                                    `json:"updated_at,nullable" format:"date-time"`
	UserAlias               string                                       `json:"user_alias,nullable"`
	UserEmail               string                                       `json:"user_email,nullable"`
	UserRole                string                                       `json:"user_role,nullable"`
	JSON                    organizationNewResponseUserJSON              `json:"-"`
}

// organizationNewResponseUserJSON contains the JSON metadata for the struct
// [OrganizationNewResponseUser]
type organizationNewResponseUserJSON struct {
	UserID                  apijson.Field
	BudgetDuration          apijson.Field
	BudgetResetAt           apijson.Field
	CreatedAt               apijson.Field
	MaxBudget               apijson.Field
	Metadata                apijson.Field
	ModelMaxBudget          apijson.Field
	ModelSpend              apijson.Field
	Models                  apijson.Field
	ObjectPermission        apijson.Field
	OrganizationMemberships apijson.Field
	RpmLimit                apijson.Field
	Spend                   apijson.Field
	SSOUserID               apijson.Field
	Teams                   apijson.Field
	TpmLimit                apijson.Field
	UpdatedAt               apijson.Field
	UserAlias               apijson.Field
	UserEmail               apijson.Field
	UserRole                apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *OrganizationNewResponseUser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationNewResponseUserJSON) RawJSON() string {
	return r.raw
}

// Represents a LiteLLM_ObjectPermissionTable record
type OrganizationNewResponseUsersObjectPermission struct {
	ObjectPermissionID string                                           `json:"object_permission_id,required"`
	AgentAccessGroups  []string                                         `json:"agent_access_groups,nullable"`
	Agents             []string                                         `json:"agents,nullable"`
	McpAccessGroups    []string                                         `json:"mcp_access_groups,nullable"`
	McpServers         []string                                         `json:"mcp_servers,nullable"`
	McpToolPermissions map[string][]string                              `json:"mcp_tool_permissions,nullable"`
	VectorStores       []string                                         `json:"vector_stores,nullable"`
	JSON               organizationNewResponseUsersObjectPermissionJSON `json:"-"`
}

// organizationNewResponseUsersObjectPermissionJSON contains the JSON metadata for
// the struct [OrganizationNewResponseUsersObjectPermission]
type organizationNewResponseUsersObjectPermissionJSON struct {
	ObjectPermissionID apijson.Field
	AgentAccessGroups  apijson.Field
	Agents             apijson.Field
	McpAccessGroups    apijson.Field
	McpServers         apijson.Field
	McpToolPermissions apijson.Field
	VectorStores       apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *OrganizationNewResponseUsersObjectPermission) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationNewResponseUsersObjectPermissionJSON) RawJSON() string {
	return r.raw
}

type OrganizationAddMemberResponse struct {
	OrganizationID                 string                                     `json:"organization_id,required"`
	UpdatedOrganizationMemberships []OrganizationMembershipTable              `json:"updated_organization_memberships,required"`
	UpdatedUsers                   []OrganizationAddMemberResponseUpdatedUser `json:"updated_users,required"`
	JSON                           organizationAddMemberResponseJSON          `json:"-"`
}

// organizationAddMemberResponseJSON contains the JSON metadata for the struct
// [OrganizationAddMemberResponse]
type organizationAddMemberResponseJSON struct {
	OrganizationID                 apijson.Field
	UpdatedOrganizationMemberships apijson.Field
	UpdatedUsers                   apijson.Field
	raw                            string
	ExtraFields                    map[string]apijson.Field
}

func (r *OrganizationAddMemberResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationAddMemberResponseJSON) RawJSON() string {
	return r.raw
}

type OrganizationAddMemberResponseUpdatedUser struct {
	UserID         string                 `json:"user_id,required"`
	BudgetDuration string                 `json:"budget_duration,nullable"`
	BudgetResetAt  time.Time              `json:"budget_reset_at,nullable" format:"date-time"`
	CreatedAt      time.Time              `json:"created_at,nullable" format:"date-time"`
	MaxBudget      float64                `json:"max_budget,nullable"`
	Metadata       map[string]interface{} `json:"metadata,nullable"`
	ModelMaxBudget map[string]interface{} `json:"model_max_budget,nullable"`
	ModelSpend     map[string]interface{} `json:"model_spend,nullable"`
	Models         []interface{}          `json:"models"`
	// Represents a LiteLLM_ObjectPermissionTable record
	ObjectPermission        OrganizationAddMemberResponseUpdatedUsersObjectPermission `json:"object_permission,nullable"`
	OrganizationMemberships []OrganizationMembershipTable                             `json:"organization_memberships,nullable"`
	RpmLimit                int64                                                     `json:"rpm_limit,nullable"`
	Spend                   float64                                                   `json:"spend"`
	SSOUserID               string                                                    `json:"sso_user_id,nullable"`
	Teams                   []string                                                  `json:"teams"`
	TpmLimit                int64                                                     `json:"tpm_limit,nullable"`
	UpdatedAt               time.Time                                                 `json:"updated_at,nullable" format:"date-time"`
	UserAlias               string                                                    `json:"user_alias,nullable"`
	UserEmail               string                                                    `json:"user_email,nullable"`
	UserRole                string                                                    `json:"user_role,nullable"`
	JSON                    organizationAddMemberResponseUpdatedUserJSON              `json:"-"`
}

// organizationAddMemberResponseUpdatedUserJSON contains the JSON metadata for the
// struct [OrganizationAddMemberResponseUpdatedUser]
type organizationAddMemberResponseUpdatedUserJSON struct {
	UserID                  apijson.Field
	BudgetDuration          apijson.Field
	BudgetResetAt           apijson.Field
	CreatedAt               apijson.Field
	MaxBudget               apijson.Field
	Metadata                apijson.Field
	ModelMaxBudget          apijson.Field
	ModelSpend              apijson.Field
	Models                  apijson.Field
	ObjectPermission        apijson.Field
	OrganizationMemberships apijson.Field
	RpmLimit                apijson.Field
	Spend                   apijson.Field
	SSOUserID               apijson.Field
	Teams                   apijson.Field
	TpmLimit                apijson.Field
	UpdatedAt               apijson.Field
	UserAlias               apijson.Field
	UserEmail               apijson.Field
	UserRole                apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *OrganizationAddMemberResponseUpdatedUser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationAddMemberResponseUpdatedUserJSON) RawJSON() string {
	return r.raw
}

// Represents a LiteLLM_ObjectPermissionTable record
type OrganizationAddMemberResponseUpdatedUsersObjectPermission struct {
	ObjectPermissionID string                                                        `json:"object_permission_id,required"`
	AgentAccessGroups  []string                                                      `json:"agent_access_groups,nullable"`
	Agents             []string                                                      `json:"agents,nullable"`
	McpAccessGroups    []string                                                      `json:"mcp_access_groups,nullable"`
	McpServers         []string                                                      `json:"mcp_servers,nullable"`
	McpToolPermissions map[string][]string                                           `json:"mcp_tool_permissions,nullable"`
	VectorStores       []string                                                      `json:"vector_stores,nullable"`
	JSON               organizationAddMemberResponseUpdatedUsersObjectPermissionJSON `json:"-"`
}

// organizationAddMemberResponseUpdatedUsersObjectPermissionJSON contains the JSON
// metadata for the struct
// [OrganizationAddMemberResponseUpdatedUsersObjectPermission]
type organizationAddMemberResponseUpdatedUsersObjectPermissionJSON struct {
	ObjectPermissionID apijson.Field
	AgentAccessGroups  apijson.Field
	Agents             apijson.Field
	McpAccessGroups    apijson.Field
	McpServers         apijson.Field
	McpToolPermissions apijson.Field
	VectorStores       apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *OrganizationAddMemberResponseUpdatedUsersObjectPermission) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationAddMemberResponseUpdatedUsersObjectPermissionJSON) RawJSON() string {
	return r.raw
}

type OrganizationDeleteMemberResponse = interface{}

type OrganizationNewParams struct {
	OrganizationAlias   param.Field[string]                                `json:"organization_alias,required"`
	BudgetDuration      param.Field[string]                                `json:"budget_duration"`
	BudgetID            param.Field[string]                                `json:"budget_id"`
	MaxBudget           param.Field[float64]                               `json:"max_budget"`
	MaxParallelRequests param.Field[int64]                                 `json:"max_parallel_requests"`
	Metadata            param.Field[map[string]interface{}]                `json:"metadata"`
	ModelMaxBudget      param.Field[map[string]interface{}]                `json:"model_max_budget"`
	ModelRpmLimit       param.Field[map[string]int64]                      `json:"model_rpm_limit"`
	ModelTpmLimit       param.Field[map[string]int64]                      `json:"model_tpm_limit"`
	Models              param.Field[[]interface{}]                         `json:"models"`
	ObjectPermission    param.Field[OrganizationNewParamsObjectPermission] `json:"object_permission"`
	OrganizationID      param.Field[string]                                `json:"organization_id"`
	RpmLimit            param.Field[int64]                                 `json:"rpm_limit"`
	SoftBudget          param.Field[float64]                               `json:"soft_budget"`
	TpmLimit            param.Field[int64]                                 `json:"tpm_limit"`
}

func (r OrganizationNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OrganizationNewParamsObjectPermission struct {
	AgentAccessGroups  param.Field[[]string]            `json:"agent_access_groups"`
	Agents             param.Field[[]string]            `json:"agents"`
	McpAccessGroups    param.Field[[]string]            `json:"mcp_access_groups"`
	McpServers         param.Field[[]string]            `json:"mcp_servers"`
	McpToolPermissions param.Field[map[string][]string] `json:"mcp_tool_permissions"`
	VectorStores       param.Field[[]string]            `json:"vector_stores"`
}

func (r OrganizationNewParamsObjectPermission) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OrganizationListParams struct {
	// Filter organizations by partial organization_alias match. Supports
	// case-insensitive search.
	OrgAlias param.Field[string] `query:"org_alias"`
	// Filter organizations by exact organization_id match
	OrgID param.Field[string] `query:"org_id"`
}

// URLQuery serializes [OrganizationListParams]'s query parameters as `url.Values`.
func (r OrganizationListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OrganizationDeleteParams struct {
	OrganizationIDs param.Field[[]string] `json:"organization_ids,required"`
}

func (r OrganizationDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OrganizationAddMemberParams struct {
	Member                  param.Field[OrganizationAddMemberParamsMemberUnion] `json:"member,required"`
	OrganizationID          param.Field[string]                                 `json:"organization_id,required"`
	MaxBudgetInOrganization param.Field[float64]                                `json:"max_budget_in_organization"`
}

func (r OrganizationAddMemberParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [OrganizationAddMemberParamsMemberArray], [OrgMemberParam].
type OrganizationAddMemberParamsMemberUnion interface {
	implementsOrganizationAddMemberParamsMemberUnion()
}

type OrganizationAddMemberParamsMemberArray []OrgMemberParam

func (r OrganizationAddMemberParamsMemberArray) implementsOrganizationAddMemberParamsMemberUnion() {}

type OrganizationDeleteMemberParams struct {
	OrganizationID param.Field[string] `json:"organization_id,required"`
	UserEmail      param.Field[string] `json:"user_email"`
	UserID         param.Field[string] `json:"user_id"`
}

func (r OrganizationDeleteMemberParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OrganizationUpdateMemberParams struct {
	OrganizationID          param.Field[string]  `json:"organization_id,required"`
	MaxBudgetInOrganization param.Field[float64] `json:"max_budget_in_organization"`
	// Admin Roles: PROXY_ADMIN: admin over the platform PROXY_ADMIN_VIEW_ONLY: can
	// login, view all own keys, view all spend ORG_ADMIN: admin over a specific
	// organization, can create teams, users only within their organization
	//
	// Internal User Roles: INTERNAL_USER: can login, view/create/delete their own
	// keys, view their spend INTERNAL_USER_VIEW_ONLY: can login, view their own keys,
	// view their own spend
	//
	// Team Roles: TEAM: used for JWT auth
	//
	// Customer Roles: CUSTOMER: External users -> these are customers
	Role      param.Field[UserRoles] `json:"role"`
	UserEmail param.Field[string]    `json:"user_email"`
	UserID    param.Field[string]    `json:"user_id"`
}

func (r OrganizationUpdateMemberParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
