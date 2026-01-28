// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hanzoai

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"slices"
	"time"

	"github.com/hanzoai/go-sdk/internal/apijson"
	"github.com/hanzoai/go-sdk/internal/apiquery"
	"github.com/hanzoai/go-sdk/internal/param"
	"github.com/hanzoai/go-sdk/internal/requestconfig"
	"github.com/hanzoai/go-sdk/option"
	"github.com/hanzoai/go-sdk/shared"
	"github.com/tidwall/gjson"
)

// TeamService contains methods and other services that help with interacting with
// the Hanzo API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTeamService] method instead.
type TeamService struct {
	Options  []option.RequestOption
	Model    *TeamModelService
	Callback *TeamCallbackService
}

// NewTeamService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewTeamService(opts ...option.RequestOption) (r *TeamService) {
	r = &TeamService{}
	r.Options = opts
	r.Model = NewTeamModelService(opts...)
	r.Callback = NewTeamCallbackService(opts...)
	return
}

// Allow users to create a new team. Apply user permissions to their team.
//
// 👉
// [Detailed Doc on setting team budgets](https://docs.litellm.ai/docs/proxy/team_budgets)
//
// Parameters:
//
//   - team_alias: Optional[str] - User defined team alias
//   - team_id: Optional[str] - The team id of the user. If none passed, we'll
//     generate it.
//   - members_with_roles: List[{"role": "admin" or "user", "user_id":
//     "<user-id>"}] - A list of users and their roles in the team. Get user_id when
//     making a new user via `/user/new`.
//   - team_member_permissions: Optional[List[str]] - A list of routes that non-admin
//     team members can access. example: ["/key/generate", "/key/update",
//     "/key/delete"]
//   - metadata: Optional[dict] - Metadata for team, store information for team.
//     Example metadata = {"extra_info": "some info"}
//   - model_rpm_limit: Optional[Dict[str, int]] - The RPM (Requests Per Minute)
//     limit for this team - applied across all keys for this team.
//   - model_tpm_limit: Optional[Dict[str, int]] - The TPM (Tokens Per Minute) limit
//     for this team - applied across all keys for this team.
//   - tpm_limit: Optional[int] - The TPM (Tokens Per Minute) limit for this team -
//     all keys with this team_id will have at max this TPM limit
//   - rpm_limit: Optional[int] - The RPM (Requests Per Minute) limit for this team -
//     all keys associated with this team_id will have at max this RPM limit
//   - rpm_limit_type: Optional[Literal["guaranteed_throughput",
//     "best_effort_throughput"]] - The type of RPM limit enforcement. Use
//     "guaranteed_throughput" to raise an error if overallocating RPM, or
//     "best_effort_throughput" for best effort enforcement.
//   - tpm_limit_type: Optional[Literal["guaranteed_throughput",
//     "best_effort_throughput"]] - The type of TPM limit enforcement. Use
//     "guaranteed_throughput" to raise an error if overallocating TPM, or
//     "best_effort_throughput" for best effort enforcement.
//   - max_budget: Optional[float] - The maximum budget allocated to the team - all
//     keys for this team_id will have at max this max_budget
//   - budget_duration: Optional[str] - The duration of the budget for the team. Doc
//     [here](https://docs.litellm.ai/docs/proxy/team_budgets)
//   - models: Optional[list] - A list of models associated with the team - all keys
//     for this team_id will have at most, these models. If empty, assumes all models
//     are allowed.
//   - blocked: bool - Flag indicating if the team is blocked or not - will stop all
//     calls from keys with this team_id.
//   - members: Optional[List] - Control team members via `/team/member/add` and
//     `/team/member/delete`.
//   - tags: Optional[List[str]] - Tags for
//     [tracking spend](https://litellm.vercel.app/docs/proxy/enterprise#tracking-spend-for-custom-tags)
//     and/or doing
//     [tag-based routing](https://litellm.vercel.app/docs/proxy/tag_routing).
//   - prompts: Optional[List[str]] - List of prompts that the team is allowed to
//     use.
//   - organization_id: Optional[str] - The organization id of the team. Default is
//     None. Create via `/organization/new`.
//   - model_aliases: Optional[dict] - Model aliases for the team.
//     [Docs](https://docs.litellm.ai/docs/proxy/team_based_routing#create-team-with-model-alias)
//   - guardrails: Optional[List[str]] - Guardrails for the team.
//     [Docs](https://docs.litellm.ai/docs/proxy/guardrails)
//   - disable_global_guardrails: Optional[bool] - Whether to disable global
//     guardrails for the key.
//   - object_permission: Optional[LiteLLM_ObjectPermissionBase] - team-specific
//     object permission. Example - {"vector_stores": ["vector_store_1",
//     "vector_store_2"], "agents": ["agent_1", "agent_2"], "agent_access_groups":
//     ["dev_group"]}. IF null or {} then no object permission.
//   - team_member_budget: Optional[float] - The maximum budget allocated to an
//     individual team member.
//   - team_member_rpm_limit: Optional[int] - The RPM (Requests Per Minute) limit for
//     individual team members.
//   - team_member_tpm_limit: Optional[int] - The TPM (Tokens Per Minute) limit for
//     individual team members.
//   - team_member_key_duration: Optional[str] - The duration for a team member's
//     key. e.g. "1d", "1w", "1mo"
//   - allowed_passthrough_routes: Optional[List[str]] - List of allowed pass through
//     routes for the team.
//   - allowed_vector_store_indexes: Optional[List[dict]] - List of allowed vector
//     store indexes for the key. Example - [{"index_name": "my-index",
//     "index_permissions": ["write", "read"]}]. If specified, the key will only be
//     able to use these specific vector store indexes. Create index, using
//     `/v1/indexes` endpoint.
//   - secret_manager_settings: Optional[dict] - Secret manager settings for the
//     team. [Docs](https://docs.litellm.ai/docs/secret_managers/overview)
//   - router_settings: Optional[UpdateRouterConfig] - team-specific router settings.
//     Example - {"model_group_retry_policy": {"max_retries": 5}}. IF null or {} then
//     no router settings.
//
// Returns:
//
//   - team_id: (str) Unique team id - used for tracking spend across multiple keys
//     for same team id.
//
// \_deprecated_params:
//
// - admins: list - A list of user_id's for the admin role
// - users: list - A list of user_id's for the user role
//
// Example Request:
//
// ```
//
//	curl --location 'http://0.0.0.0:4000/team/new'     --header 'Authorization: Bearer sk-1234'     --header 'Content-Type: application/json'     --data '{
//	  "team_alias": "my-new-team_2",
//	  "members_with_roles": [{"role": "admin", "user_id": "user-1234"},
//	    {"role": "user", "user_id": "user-2434"}]
//	}'
//
// ```
//
// ```
//
//	curl --location 'http://0.0.0.0:4000/team/new'     --header 'Authorization: Bearer sk-1234'     --header 'Content-Type: application/json'     --data '{
//	           "team_alias": "QA Prod Bot",
//	           "max_budget": 0.000000001,
//	           "budget_duration": "1d"
//	       }'
//
// ```
func (r *TeamService) New(ctx context.Context, params TeamNewParams, opts ...option.RequestOption) (res *TeamNewResponse, err error) {
	if params.LitellmChangedBy.Present {
		opts = append(opts, option.WithHeader("litellm-changed-by", fmt.Sprintf("%s", params.LitellmChangedBy)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "team/new"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Use `/team/member_add` AND `/team/member/delete` to add/remove new team members
//
// You can now update team budget / rate limits via /team/update
//
// Parameters:
//
//   - team_id: str - The team id of the user. Required param.
//   - team_alias: Optional[str] - User defined team alias
//   - team_member_permissions: Optional[List[str]] - A list of routes that non-admin
//     team members can access. example: ["/key/generate", "/key/update",
//     "/key/delete"]
//   - metadata: Optional[dict] - Metadata for team, store information for team.
//     Example metadata = {"team": "core-infra", "app": "app2", "email":
//     "ishaan@berri.ai" }
//   - tpm_limit: Optional[int] - The TPM (Tokens Per Minute) limit for this team -
//     all keys with this team_id will have at max this TPM limit
//   - rpm_limit: Optional[int] - The RPM (Requests Per Minute) limit for this team -
//     all keys associated with this team_id will have at max this RPM limit
//   - max_budget: Optional[float] - The maximum budget allocated to the team - all
//     keys for this team_id will have at max this max_budget
//   - budget_duration: Optional[str] - The duration of the budget for the team. Doc
//     [here](https://docs.litellm.ai/docs/proxy/team_budgets)
//   - models: Optional[list] - A list of models associated with the team - all keys
//     for this team_id will have at most, these models. If empty, assumes all models
//     are allowed.
//   - prompts: Optional[List[str]] - List of prompts that the team is allowed to
//     use.
//   - blocked: bool - Flag indicating if the team is blocked or not - will stop all
//     calls from keys with this team_id.
//   - tags: Optional[List[str]] - Tags for
//     [tracking spend](https://litellm.vercel.app/docs/proxy/enterprise#tracking-spend-for-custom-tags)
//     and/or doing
//     [tag-based routing](https://litellm.vercel.app/docs/proxy/tag_routing).
//   - organization_id: Optional[str] - The organization id of the team. Default is
//     None. Create via `/organization/new`.
//   - model_aliases: Optional[dict] - Model aliases for the team.
//     [Docs](https://docs.litellm.ai/docs/proxy/team_based_routing#create-team-with-model-alias)
//   - guardrails: Optional[List[str]] - Guardrails for the team.
//     [Docs](https://docs.litellm.ai/docs/proxy/guardrails)
//   - disable_global_guardrails: Optional[bool] - Whether to disable global
//     guardrails for the key.
//   - object_permission: Optional[LiteLLM_ObjectPermissionBase] - team-specific
//     object permission. Example - {"vector_stores": ["vector_store_1",
//     "vector_store_2"], "agents": ["agent_1", "agent_2"], "agent_access_groups":
//     ["dev_group"]}. IF null or {} then no object permission.
//   - team_member_budget: Optional[float] - The maximum budget allocated to an
//     individual team member.
//   - team_member_budget_duration: Optional[str] - The duration of the budget for
//     the team member. Doc [here](https://docs.litellm.ai/docs/proxy/team_budgets)
//   - team_member_rpm_limit: Optional[int] - The RPM (Requests Per Minute) limit for
//     individual team members.
//   - team_member_tpm_limit: Optional[int] - The TPM (Tokens Per Minute) limit for
//     individual team members.
//   - team_member_key_duration: Optional[str] - The duration for a team member's
//     key. e.g. "1d", "1w", "1mo"
//   - allowed_passthrough_routes: Optional[List[str]] - List of allowed pass through
//     routes for the team.
//   - model_rpm_limit: Optional[Dict[str, int]] - The RPM (Requests Per Minute)
//     limit per model for this team. Example: {"gpt-4": 100, "gpt-3.5-turbo": 200}
//   - model_tpm_limit: Optional[Dict[str, int]] - The TPM (Tokens Per Minute) limit
//     per model for this team. Example: {"gpt-4": 10000, "gpt-3.5-turbo": 20000}
//     Example - update team TPM Limit
//   - allowed_vector_store_indexes: Optional[List[dict]] - List of allowed vector
//     store indexes for the key. Example - [{"index_name": "my-index",
//     "index_permissions": ["write", "read"]}]. If specified, the key will only be
//     able to use these specific vector store indexes. Create index, using
//     `/v1/indexes` endpoint.
//   - secret_manager_settings: Optional[dict] - Secret manager settings for the
//     team. [Docs](https://docs.litellm.ai/docs/secret_managers/overview)
//   - router_settings: Optional[UpdateRouterConfig] - team-specific router settings.
//     Example - {"model_group_retry_policy": {"max_retries": 5}}. IF null or {} then
//     no router settings.
//
// ```
//
//	curl --location 'http://0.0.0.0:4000/team/update'     --header 'Authorization: Bearer sk-1234'     --header 'Content-Type: application/json'     --data-raw '{
//	    "team_id": "8d916b1c-510d-4894-a334-1c16a93344f5",
//	    "tpm_limit": 100
//	}'
//
// ```
//
// Example - Update Team `max_budget` budget
//
// ```
//
//	curl --location 'http://0.0.0.0:4000/team/update'     --header 'Authorization: Bearer sk-1234'     --header 'Content-Type: application/json'     --data-raw '{
//	    "team_id": "8d916b1c-510d-4894-a334-1c16a93344f5",
//	    "max_budget": 10
//	}'
//
// ```
func (r *TeamService) Update(ctx context.Context, params TeamUpdateParams, opts ...option.RequestOption) (res *TeamUpdateResponse, err error) {
	if params.LitellmChangedBy.Present {
		opts = append(opts, option.WithHeader("litellm-changed-by", fmt.Sprintf("%s", params.LitellmChangedBy)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "team/update"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// ```
// curl --location --request GET 'http://0.0.0.0:4000/team/list'         --header 'Authorization: Bearer sk-1234'
// ```
//
// Parameters:
//
//   - user_id: str - Optional. If passed will only return teams that the user_id is
//     a member of.
//   - organization_id: str - Optional. If passed will only return teams that belong
//     to the organization_id. Pass 'default_organization' to get all teams without
//     organization_id.
func (r *TeamService) List(ctx context.Context, query TeamListParams, opts ...option.RequestOption) (res *TeamListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "team/list"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// delete team and associated team keys
//
// Parameters:
//
//   - team_ids: List[str] - Required. List of team IDs to delete. Example:
//     ["team-1234", "team-5678"]
//
// ```
//
//	curl --location 'http://0.0.0.0:4000/team/delete'     --header 'Authorization: Bearer sk-1234'     --header 'Content-Type: application/json'     --data-raw '{
//	    "team_ids": ["8d916b1c-510d-4894-a334-1c16a93344f5"]
//	}'
//
// ```
func (r *TeamService) Delete(ctx context.Context, params TeamDeleteParams, opts ...option.RequestOption) (res *TeamDeleteResponse, err error) {
	if params.LitellmChangedBy.Present {
		opts = append(opts, option.WithHeader("litellm-changed-by", fmt.Sprintf("%s", params.LitellmChangedBy)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "team/delete"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Add new members (either via user_email or user_id) to a team
//
// If user doesn't exist, new user row will also be added to User Table
//
// Only proxy_admin or admin of team, allowed to access this endpoint.
//
// ```
//
// curl -X POST 'http://0.0.0.0:4000/team/member_add'     -H 'Authorization: Bearer sk-1234'     -H 'Content-Type: application/json'     -d '{"team_id": "45e3e396-ee08-4a61-a88e-16b3ce7e0849", "member": {"role": "user", "user_id": "krrish247652@berri.ai"}}'
//
// ```
func (r *TeamService) AddMember(ctx context.Context, body TeamAddMemberParams, opts ...option.RequestOption) (res *TeamAddMemberResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "team/member_add"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Blocks all calls from keys with this team id.
//
// Parameters:
//
// - team_id: str - Required. The unique identifier of the team to block.
//
// Example:
//
// ```
//
//	curl --location 'http://0.0.0.0:4000/team/block'     --header 'Authorization: Bearer sk-1234'     --header 'Content-Type: application/json'     --data '{
//	    "team_id": "team-1234"
//	}'
//
// ```
//
// Returns:
//
// - The updated team record with blocked=True
func (r *TeamService) Block(ctx context.Context, body TeamBlockParams, opts ...option.RequestOption) (res *TeamBlockResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "team/block"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Disable all logging callbacks for a team
//
// Parameters:
//
// - team_id (str, required): The unique identifier for the team
//
// Example curl:
//
// ```
// curl -X POST 'http://localhost:4000/team/dbe2f686-a686-4896-864a-4c3924458709/disable_logging'         -H 'Authorization: Bearer sk-1234'
// ```
func (r *TeamService) DisableLogging(ctx context.Context, teamID string, opts ...option.RequestOption) (res *TeamDisableLoggingResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if teamID == "" {
		err = errors.New("missing required team_id parameter")
		return
	}
	path := fmt.Sprintf("team/%s/disable_logging", teamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// List Available Teams
func (r *TeamService) ListAvailable(ctx context.Context, query TeamListAvailableParams, opts ...option.RequestOption) (res *TeamListAvailableResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "team/available"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// [BETA]
//
// delete members (either via user_email or user_id) from a team
//
// If user doesn't exist, an exception will be raised
//
// ```
// curl -X POST 'http://0.0.0.0:8000/team/member_delete'
// -H 'Authorization: Bearer sk-1234'
// -H 'Content-Type: application/json'
//
//	-d '{
//	    "team_id": "45e3e396-ee08-4a61-a88e-16b3ce7e0849",
//	    "user_id": "krrish247652@berri.ai"
//	}'
//
// ```
func (r *TeamService) RemoveMember(ctx context.Context, body TeamRemoveMemberParams, opts ...option.RequestOption) (res *TeamRemoveMemberResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "team/member_delete"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// get info on team + related keys
//
// Parameters:
//
// - team_id: str - Required. The unique identifier of the team to get info on.
//
// ```
// curl --location 'http://localhost:4000/team/info?team_id=your_team_id_here'     --header 'Authorization: Bearer your_api_key_here'
// ```
func (r *TeamService) GetInfo(ctx context.Context, query TeamGetInfoParams, opts ...option.RequestOption) (res *TeamGetInfoResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "team/info"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Blocks all calls from keys with this team id.
//
// Parameters:
//
// - team_id: str - Required. The unique identifier of the team to unblock.
//
// Example:
//
// ```
//
//	curl --location 'http://0.0.0.0:4000/team/unblock'     --header 'Authorization: Bearer sk-1234'     --header 'Content-Type: application/json'     --data '{
//	    "team_id": "team-1234"
//	}'
//
// ```
func (r *TeamService) Unblock(ctx context.Context, body TeamUnblockParams, opts ...option.RequestOption) (res *TeamUnblockResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "team/unblock"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// [BETA]
//
// Update team member budgets and team member role
func (r *TeamService) UpdateMember(ctx context.Context, body TeamUpdateMemberParams, opts ...option.RequestOption) (res *TeamUpdateMemberResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "team/member_update"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type BlockTeamRequestParam struct {
	TeamID param.Field[string] `json:"team_id,required"`
}

func (r BlockTeamRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type Member struct {
	// The role of the user within the team. 'admin' users can manage team settings and
	// members, 'user' is a regular team member
	Role MemberRole `json:"role,required"`
	// The email address of the user to add. Either user_id or user_email must be
	// provided
	UserEmail string `json:"user_email,nullable"`
	// The unique ID of the user to add. Either user_id or user_email must be provided
	UserID string     `json:"user_id,nullable"`
	JSON   memberJSON `json:"-"`
}

// memberJSON contains the JSON metadata for the struct [Member]
type memberJSON struct {
	Role        apijson.Field
	UserEmail   apijson.Field
	UserID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Member) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memberJSON) RawJSON() string {
	return r.raw
}

// The role of the user within the team. 'admin' users can manage team settings and
// members, 'user' is a regular team member
type MemberRole string

const (
	MemberRoleAdmin MemberRole = "admin"
	MemberRoleUser  MemberRole = "user"
)

func (r MemberRole) IsKnown() bool {
	switch r {
	case MemberRoleAdmin, MemberRoleUser:
		return true
	}
	return false
}

type MemberParam struct {
	// The role of the user within the team. 'admin' users can manage team settings and
	// members, 'user' is a regular team member
	Role param.Field[MemberRole] `json:"role,required"`
	// The email address of the user to add. Either user_id or user_email must be
	// provided
	UserEmail param.Field[string] `json:"user_email"`
	// The unique ID of the user to add. Either user_id or user_email must be provided
	UserID param.Field[string] `json:"user_id"`
}

func (r MemberParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MemberParam) implementsTeamAddMemberParamsMemberUnion() {}

type TeamNewResponse struct {
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
	ObjectPermission      TeamNewResponseObjectPermission `json:"object_permission,nullable"`
	ObjectPermissionID    string                          `json:"object_permission_id,nullable"`
	OrganizationID        string                          `json:"organization_id,nullable"`
	RouterSettings        map[string]interface{}          `json:"router_settings,nullable"`
	RpmLimit              int64                           `json:"rpm_limit,nullable"`
	Spend                 float64                         `json:"spend,nullable"`
	TeamAlias             string                          `json:"team_alias,nullable"`
	TeamMemberPermissions []string                        `json:"team_member_permissions,nullable"`
	TpmLimit              int64                           `json:"tpm_limit,nullable"`
	UpdatedAt             time.Time                       `json:"updated_at,nullable" format:"date-time"`
	JSON                  teamNewResponseJSON             `json:"-"`
}

// teamNewResponseJSON contains the JSON metadata for the struct [TeamNewResponse]
type teamNewResponseJSON struct {
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

func (r *TeamNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r teamNewResponseJSON) RawJSON() string {
	return r.raw
}

// Represents a LiteLLM_ObjectPermissionTable record
type TeamNewResponseObjectPermission struct {
	ObjectPermissionID string                              `json:"object_permission_id,required"`
	AgentAccessGroups  []string                            `json:"agent_access_groups,nullable"`
	Agents             []string                            `json:"agents,nullable"`
	McpAccessGroups    []string                            `json:"mcp_access_groups,nullable"`
	McpServers         []string                            `json:"mcp_servers,nullable"`
	McpToolPermissions map[string][]string                 `json:"mcp_tool_permissions,nullable"`
	VectorStores       []string                            `json:"vector_stores,nullable"`
	JSON               teamNewResponseObjectPermissionJSON `json:"-"`
}

// teamNewResponseObjectPermissionJSON contains the JSON metadata for the struct
// [TeamNewResponseObjectPermission]
type teamNewResponseObjectPermissionJSON struct {
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

func (r *TeamNewResponseObjectPermission) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r teamNewResponseObjectPermissionJSON) RawJSON() string {
	return r.raw
}

type TeamUpdateResponse = interface{}

type TeamListResponse = interface{}

type TeamDeleteResponse = interface{}

type TeamAddMemberResponse struct {
	TeamID                 string                                       `json:"team_id,required"`
	UpdatedTeamMemberships []TeamAddMemberResponseUpdatedTeamMembership `json:"updated_team_memberships,required"`
	UpdatedUsers           []TeamAddMemberResponseUpdatedUser           `json:"updated_users,required"`
	Admins                 []interface{}                                `json:"admins"`
	Blocked                bool                                         `json:"blocked"`
	BudgetDuration         string                                       `json:"budget_duration,nullable"`
	BudgetResetAt          time.Time                                    `json:"budget_reset_at,nullable" format:"date-time"`
	CreatedAt              time.Time                                    `json:"created_at,nullable" format:"date-time"`
	LitellmModelTable      TeamAddMemberResponseLitellmModelTable       `json:"litellm_model_table,nullable"`
	MaxBudget              float64                                      `json:"max_budget,nullable"`
	MaxParallelRequests    int64                                        `json:"max_parallel_requests,nullable"`
	Members                []interface{}                                `json:"members"`
	MembersWithRoles       []Member                                     `json:"members_with_roles"`
	Metadata               map[string]interface{}                       `json:"metadata,nullable"`
	ModelID                int64                                        `json:"model_id,nullable"`
	Models                 []interface{}                                `json:"models"`
	// Represents a LiteLLM_ObjectPermissionTable record
	ObjectPermission      TeamAddMemberResponseObjectPermission `json:"object_permission,nullable"`
	ObjectPermissionID    string                                `json:"object_permission_id,nullable"`
	OrganizationID        string                                `json:"organization_id,nullable"`
	RouterSettings        map[string]interface{}                `json:"router_settings,nullable"`
	RpmLimit              int64                                 `json:"rpm_limit,nullable"`
	Spend                 float64                               `json:"spend,nullable"`
	TeamAlias             string                                `json:"team_alias,nullable"`
	TeamMemberPermissions []string                              `json:"team_member_permissions,nullable"`
	TpmLimit              int64                                 `json:"tpm_limit,nullable"`
	UpdatedAt             time.Time                             `json:"updated_at,nullable" format:"date-time"`
	JSON                  teamAddMemberResponseJSON             `json:"-"`
}

// teamAddMemberResponseJSON contains the JSON metadata for the struct
// [TeamAddMemberResponse]
type teamAddMemberResponseJSON struct {
	TeamID                 apijson.Field
	UpdatedTeamMemberships apijson.Field
	UpdatedUsers           apijson.Field
	Admins                 apijson.Field
	Blocked                apijson.Field
	BudgetDuration         apijson.Field
	BudgetResetAt          apijson.Field
	CreatedAt              apijson.Field
	LitellmModelTable      apijson.Field
	MaxBudget              apijson.Field
	MaxParallelRequests    apijson.Field
	Members                apijson.Field
	MembersWithRoles       apijson.Field
	Metadata               apijson.Field
	ModelID                apijson.Field
	Models                 apijson.Field
	ObjectPermission       apijson.Field
	ObjectPermissionID     apijson.Field
	OrganizationID         apijson.Field
	RouterSettings         apijson.Field
	RpmLimit               apijson.Field
	Spend                  apijson.Field
	TeamAlias              apijson.Field
	TeamMemberPermissions  apijson.Field
	TpmLimit               apijson.Field
	UpdatedAt              apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *TeamAddMemberResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r teamAddMemberResponseJSON) RawJSON() string {
	return r.raw
}

type TeamAddMemberResponseUpdatedTeamMembership struct {
	// Represents user-controllable params for a LiteLLM_BudgetTable record
	LitellmBudgetTable BudgetTable                                    `json:"litellm_budget_table,required,nullable"`
	TeamID             string                                         `json:"team_id,required"`
	UserID             string                                         `json:"user_id,required"`
	BudgetID           string                                         `json:"budget_id,nullable"`
	Spend              float64                                        `json:"spend,nullable"`
	JSON               teamAddMemberResponseUpdatedTeamMembershipJSON `json:"-"`
}

// teamAddMemberResponseUpdatedTeamMembershipJSON contains the JSON metadata for
// the struct [TeamAddMemberResponseUpdatedTeamMembership]
type teamAddMemberResponseUpdatedTeamMembershipJSON struct {
	LitellmBudgetTable apijson.Field
	TeamID             apijson.Field
	UserID             apijson.Field
	BudgetID           apijson.Field
	Spend              apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *TeamAddMemberResponseUpdatedTeamMembership) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r teamAddMemberResponseUpdatedTeamMembershipJSON) RawJSON() string {
	return r.raw
}

type TeamAddMemberResponseUpdatedUser struct {
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
	ObjectPermission        TeamAddMemberResponseUpdatedUsersObjectPermission `json:"object_permission,nullable"`
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
	JSON                    teamAddMemberResponseUpdatedUserJSON              `json:"-"`
}

// teamAddMemberResponseUpdatedUserJSON contains the JSON metadata for the struct
// [TeamAddMemberResponseUpdatedUser]
type teamAddMemberResponseUpdatedUserJSON struct {
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

func (r *TeamAddMemberResponseUpdatedUser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r teamAddMemberResponseUpdatedUserJSON) RawJSON() string {
	return r.raw
}

// Represents a LiteLLM_ObjectPermissionTable record
type TeamAddMemberResponseUpdatedUsersObjectPermission struct {
	ObjectPermissionID string                                                `json:"object_permission_id,required"`
	AgentAccessGroups  []string                                              `json:"agent_access_groups,nullable"`
	Agents             []string                                              `json:"agents,nullable"`
	McpAccessGroups    []string                                              `json:"mcp_access_groups,nullable"`
	McpServers         []string                                              `json:"mcp_servers,nullable"`
	McpToolPermissions map[string][]string                                   `json:"mcp_tool_permissions,nullable"`
	VectorStores       []string                                              `json:"vector_stores,nullable"`
	JSON               teamAddMemberResponseUpdatedUsersObjectPermissionJSON `json:"-"`
}

// teamAddMemberResponseUpdatedUsersObjectPermissionJSON contains the JSON metadata
// for the struct [TeamAddMemberResponseUpdatedUsersObjectPermission]
type teamAddMemberResponseUpdatedUsersObjectPermissionJSON struct {
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

func (r *TeamAddMemberResponseUpdatedUsersObjectPermission) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r teamAddMemberResponseUpdatedUsersObjectPermissionJSON) RawJSON() string {
	return r.raw
}

type TeamAddMemberResponseLitellmModelTable struct {
	CreatedBy    string                                                  `json:"created_by,required"`
	UpdatedBy    string                                                  `json:"updated_by,required"`
	ID           int64                                                   `json:"id,nullable"`
	ModelAliases TeamAddMemberResponseLitellmModelTableModelAliasesUnion `json:"model_aliases,nullable"`
	Team         interface{}                                             `json:"team"`
	JSON         teamAddMemberResponseLitellmModelTableJSON              `json:"-"`
}

// teamAddMemberResponseLitellmModelTableJSON contains the JSON metadata for the
// struct [TeamAddMemberResponseLitellmModelTable]
type teamAddMemberResponseLitellmModelTableJSON struct {
	CreatedBy    apijson.Field
	UpdatedBy    apijson.Field
	ID           apijson.Field
	ModelAliases apijson.Field
	Team         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *TeamAddMemberResponseLitellmModelTable) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r teamAddMemberResponseLitellmModelTableJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [TeamAddMemberResponseLitellmModelTableModelAliasesMap] or
// [shared.UnionString].
type TeamAddMemberResponseLitellmModelTableModelAliasesUnion interface {
	ImplementsTeamAddMemberResponseLitellmModelTableModelAliasesUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TeamAddMemberResponseLitellmModelTableModelAliasesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TeamAddMemberResponseLitellmModelTableModelAliasesMap{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type TeamAddMemberResponseLitellmModelTableModelAliasesMap map[string]interface{}

func (r TeamAddMemberResponseLitellmModelTableModelAliasesMap) ImplementsTeamAddMemberResponseLitellmModelTableModelAliasesUnion() {
}

// Represents a LiteLLM_ObjectPermissionTable record
type TeamAddMemberResponseObjectPermission struct {
	ObjectPermissionID string                                    `json:"object_permission_id,required"`
	AgentAccessGroups  []string                                  `json:"agent_access_groups,nullable"`
	Agents             []string                                  `json:"agents,nullable"`
	McpAccessGroups    []string                                  `json:"mcp_access_groups,nullable"`
	McpServers         []string                                  `json:"mcp_servers,nullable"`
	McpToolPermissions map[string][]string                       `json:"mcp_tool_permissions,nullable"`
	VectorStores       []string                                  `json:"vector_stores,nullable"`
	JSON               teamAddMemberResponseObjectPermissionJSON `json:"-"`
}

// teamAddMemberResponseObjectPermissionJSON contains the JSON metadata for the
// struct [TeamAddMemberResponseObjectPermission]
type teamAddMemberResponseObjectPermissionJSON struct {
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

func (r *TeamAddMemberResponseObjectPermission) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r teamAddMemberResponseObjectPermissionJSON) RawJSON() string {
	return r.raw
}

type TeamBlockResponse = interface{}

type TeamDisableLoggingResponse = interface{}

type TeamListAvailableResponse = interface{}

type TeamRemoveMemberResponse = interface{}

type TeamGetInfoResponse = interface{}

type TeamUnblockResponse = interface{}

type TeamUpdateMemberResponse struct {
	TeamID          string                       `json:"team_id,required"`
	UserID          string                       `json:"user_id,required"`
	MaxBudgetInTeam float64                      `json:"max_budget_in_team,nullable"`
	RpmLimit        int64                        `json:"rpm_limit,nullable"`
	TpmLimit        int64                        `json:"tpm_limit,nullable"`
	UserEmail       string                       `json:"user_email,nullable"`
	JSON            teamUpdateMemberResponseJSON `json:"-"`
}

// teamUpdateMemberResponseJSON contains the JSON metadata for the struct
// [TeamUpdateMemberResponse]
type teamUpdateMemberResponseJSON struct {
	TeamID          apijson.Field
	UserID          apijson.Field
	MaxBudgetInTeam apijson.Field
	RpmLimit        apijson.Field
	TpmLimit        apijson.Field
	UserEmail       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *TeamUpdateMemberResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r teamUpdateMemberResponseJSON) RawJSON() string {
	return r.raw
}

type TeamNewParams struct {
	Admins                    param.Field[[]interface{}]                          `json:"admins"`
	AllowedPassthroughRoutes  param.Field[[]interface{}]                          `json:"allowed_passthrough_routes"`
	AllowedVectorStoreIndexes param.Field[[]TeamNewParamsAllowedVectorStoreIndex] `json:"allowed_vector_store_indexes"`
	Blocked                   param.Field[bool]                                   `json:"blocked"`
	BudgetDuration            param.Field[string]                                 `json:"budget_duration"`
	Guardrails                param.Field[[]string]                               `json:"guardrails"`
	MaxBudget                 param.Field[float64]                                `json:"max_budget"`
	Members                   param.Field[[]interface{}]                          `json:"members"`
	MembersWithRoles          param.Field[[]MemberParam]                          `json:"members_with_roles"`
	Metadata                  param.Field[map[string]interface{}]                 `json:"metadata"`
	ModelAliases              param.Field[map[string]interface{}]                 `json:"model_aliases"`
	ModelRpmLimit             param.Field[map[string]int64]                       `json:"model_rpm_limit"`
	ModelTpmLimit             param.Field[map[string]int64]                       `json:"model_tpm_limit"`
	Models                    param.Field[[]interface{}]                          `json:"models"`
	ObjectPermission          param.Field[TeamNewParamsObjectPermission]          `json:"object_permission"`
	OrganizationID            param.Field[string]                                 `json:"organization_id"`
	Prompts                   param.Field[[]string]                               `json:"prompts"`
	RouterSettings            param.Field[map[string]interface{}]                 `json:"router_settings"`
	RpmLimit                  param.Field[int64]                                  `json:"rpm_limit"`
	RpmLimitType              param.Field[TeamNewParamsRpmLimitType]              `json:"rpm_limit_type"`
	SecretManagerSettings     param.Field[map[string]interface{}]                 `json:"secret_manager_settings"`
	Tags                      param.Field[[]interface{}]                          `json:"tags"`
	TeamAlias                 param.Field[string]                                 `json:"team_alias"`
	TeamID                    param.Field[string]                                 `json:"team_id"`
	TeamMemberBudget          param.Field[float64]                                `json:"team_member_budget"`
	TeamMemberKeyDuration     param.Field[string]                                 `json:"team_member_key_duration"`
	TeamMemberPermissions     param.Field[[]string]                               `json:"team_member_permissions"`
	TeamMemberRpmLimit        param.Field[int64]                                  `json:"team_member_rpm_limit"`
	TeamMemberTpmLimit        param.Field[int64]                                  `json:"team_member_tpm_limit"`
	TpmLimit                  param.Field[int64]                                  `json:"tpm_limit"`
	TpmLimitType              param.Field[TeamNewParamsTpmLimitType]              `json:"tpm_limit_type"`
	// The litellm-changed-by header enables tracking of actions performed by
	// authorized users on behalf of other users, providing an audit trail for
	// accountability
	LitellmChangedBy param.Field[string] `header:"litellm-changed-by"`
}

func (r TeamNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TeamNewParamsAllowedVectorStoreIndex struct {
	IndexName        param.Field[string]                                                  `json:"index_name,required"`
	IndexPermissions param.Field[[]TeamNewParamsAllowedVectorStoreIndexesIndexPermission] `json:"index_permissions,required"`
}

func (r TeamNewParamsAllowedVectorStoreIndex) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TeamNewParamsAllowedVectorStoreIndexesIndexPermission string

const (
	TeamNewParamsAllowedVectorStoreIndexesIndexPermissionRead  TeamNewParamsAllowedVectorStoreIndexesIndexPermission = "read"
	TeamNewParamsAllowedVectorStoreIndexesIndexPermissionWrite TeamNewParamsAllowedVectorStoreIndexesIndexPermission = "write"
)

func (r TeamNewParamsAllowedVectorStoreIndexesIndexPermission) IsKnown() bool {
	switch r {
	case TeamNewParamsAllowedVectorStoreIndexesIndexPermissionRead, TeamNewParamsAllowedVectorStoreIndexesIndexPermissionWrite:
		return true
	}
	return false
}

type TeamNewParamsObjectPermission struct {
	AgentAccessGroups  param.Field[[]string]            `json:"agent_access_groups"`
	Agents             param.Field[[]string]            `json:"agents"`
	McpAccessGroups    param.Field[[]string]            `json:"mcp_access_groups"`
	McpServers         param.Field[[]string]            `json:"mcp_servers"`
	McpToolPermissions param.Field[map[string][]string] `json:"mcp_tool_permissions"`
	VectorStores       param.Field[[]string]            `json:"vector_stores"`
}

func (r TeamNewParamsObjectPermission) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TeamNewParamsRpmLimitType string

const (
	TeamNewParamsRpmLimitTypeGuaranteedThroughput TeamNewParamsRpmLimitType = "guaranteed_throughput"
	TeamNewParamsRpmLimitTypeBestEffortThroughput TeamNewParamsRpmLimitType = "best_effort_throughput"
)

func (r TeamNewParamsRpmLimitType) IsKnown() bool {
	switch r {
	case TeamNewParamsRpmLimitTypeGuaranteedThroughput, TeamNewParamsRpmLimitTypeBestEffortThroughput:
		return true
	}
	return false
}

type TeamNewParamsTpmLimitType string

const (
	TeamNewParamsTpmLimitTypeGuaranteedThroughput TeamNewParamsTpmLimitType = "guaranteed_throughput"
	TeamNewParamsTpmLimitTypeBestEffortThroughput TeamNewParamsTpmLimitType = "best_effort_throughput"
)

func (r TeamNewParamsTpmLimitType) IsKnown() bool {
	switch r {
	case TeamNewParamsTpmLimitTypeGuaranteedThroughput, TeamNewParamsTpmLimitTypeBestEffortThroughput:
		return true
	}
	return false
}

type TeamUpdateParams struct {
	TeamID                    param.Field[string]                                    `json:"team_id,required"`
	AllowedPassthroughRoutes  param.Field[[]interface{}]                             `json:"allowed_passthrough_routes"`
	AllowedVectorStoreIndexes param.Field[[]TeamUpdateParamsAllowedVectorStoreIndex] `json:"allowed_vector_store_indexes"`
	Blocked                   param.Field[bool]                                      `json:"blocked"`
	BudgetDuration            param.Field[string]                                    `json:"budget_duration"`
	Guardrails                param.Field[[]string]                                  `json:"guardrails"`
	MaxBudget                 param.Field[float64]                                   `json:"max_budget"`
	Metadata                  param.Field[map[string]interface{}]                    `json:"metadata"`
	ModelAliases              param.Field[map[string]interface{}]                    `json:"model_aliases"`
	ModelRpmLimit             param.Field[map[string]int64]                          `json:"model_rpm_limit"`
	ModelTpmLimit             param.Field[map[string]int64]                          `json:"model_tpm_limit"`
	Models                    param.Field[[]interface{}]                             `json:"models"`
	ObjectPermission          param.Field[TeamUpdateParamsObjectPermission]          `json:"object_permission"`
	OrganizationID            param.Field[string]                                    `json:"organization_id"`
	Prompts                   param.Field[[]string]                                  `json:"prompts"`
	RouterSettings            param.Field[map[string]interface{}]                    `json:"router_settings"`
	RpmLimit                  param.Field[int64]                                     `json:"rpm_limit"`
	SecretManagerSettings     param.Field[map[string]interface{}]                    `json:"secret_manager_settings"`
	Tags                      param.Field[[]interface{}]                             `json:"tags"`
	TeamAlias                 param.Field[string]                                    `json:"team_alias"`
	TeamMemberBudget          param.Field[float64]                                   `json:"team_member_budget"`
	TeamMemberBudgetDuration  param.Field[string]                                    `json:"team_member_budget_duration"`
	TeamMemberKeyDuration     param.Field[string]                                    `json:"team_member_key_duration"`
	TeamMemberRpmLimit        param.Field[int64]                                     `json:"team_member_rpm_limit"`
	TeamMemberTpmLimit        param.Field[int64]                                     `json:"team_member_tpm_limit"`
	TpmLimit                  param.Field[int64]                                     `json:"tpm_limit"`
	// The litellm-changed-by header enables tracking of actions performed by
	// authorized users on behalf of other users, providing an audit trail for
	// accountability
	LitellmChangedBy param.Field[string] `header:"litellm-changed-by"`
}

func (r TeamUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TeamUpdateParamsAllowedVectorStoreIndex struct {
	IndexName        param.Field[string]                                                     `json:"index_name,required"`
	IndexPermissions param.Field[[]TeamUpdateParamsAllowedVectorStoreIndexesIndexPermission] `json:"index_permissions,required"`
}

func (r TeamUpdateParamsAllowedVectorStoreIndex) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TeamUpdateParamsAllowedVectorStoreIndexesIndexPermission string

const (
	TeamUpdateParamsAllowedVectorStoreIndexesIndexPermissionRead  TeamUpdateParamsAllowedVectorStoreIndexesIndexPermission = "read"
	TeamUpdateParamsAllowedVectorStoreIndexesIndexPermissionWrite TeamUpdateParamsAllowedVectorStoreIndexesIndexPermission = "write"
)

func (r TeamUpdateParamsAllowedVectorStoreIndexesIndexPermission) IsKnown() bool {
	switch r {
	case TeamUpdateParamsAllowedVectorStoreIndexesIndexPermissionRead, TeamUpdateParamsAllowedVectorStoreIndexesIndexPermissionWrite:
		return true
	}
	return false
}

type TeamUpdateParamsObjectPermission struct {
	AgentAccessGroups  param.Field[[]string]            `json:"agent_access_groups"`
	Agents             param.Field[[]string]            `json:"agents"`
	McpAccessGroups    param.Field[[]string]            `json:"mcp_access_groups"`
	McpServers         param.Field[[]string]            `json:"mcp_servers"`
	McpToolPermissions param.Field[map[string][]string] `json:"mcp_tool_permissions"`
	VectorStores       param.Field[[]string]            `json:"vector_stores"`
}

func (r TeamUpdateParamsObjectPermission) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TeamListParams struct {
	OrganizationID param.Field[string] `query:"organization_id"`
	// Only return teams which this 'user_id' belongs to
	UserID param.Field[string] `query:"user_id"`
}

// URLQuery serializes [TeamListParams]'s query parameters as `url.Values`.
func (r TeamListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type TeamDeleteParams struct {
	TeamIDs param.Field[[]string] `json:"team_ids,required"`
	// The litellm-changed-by header enables tracking of actions performed by
	// authorized users on behalf of other users, providing an audit trail for
	// accountability
	LitellmChangedBy param.Field[string] `header:"litellm-changed-by"`
}

func (r TeamDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TeamAddMemberParams struct {
	// Member object or list of member objects to add. Each member must include either
	// user_id or user_email, and a role
	Member param.Field[TeamAddMemberParamsMemberUnion] `json:"member,required"`
	// The ID of the team to add the member to
	TeamID param.Field[string] `json:"team_id,required"`
	// Maximum budget allocated to this user within the team. If not set, user has
	// unlimited budget within team limits
	MaxBudgetInTeam param.Field[float64] `json:"max_budget_in_team"`
}

func (r TeamAddMemberParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Member object or list of member objects to add. Each member must include either
// user_id or user_email, and a role
//
// Satisfied by [TeamAddMemberParamsMemberArray], [MemberParam].
type TeamAddMemberParamsMemberUnion interface {
	implementsTeamAddMemberParamsMemberUnion()
}

type TeamAddMemberParamsMemberArray []MemberParam

func (r TeamAddMemberParamsMemberArray) implementsTeamAddMemberParamsMemberUnion() {}

type TeamBlockParams struct {
	BlockTeamRequest BlockTeamRequestParam `json:"block_team_request,required"`
}

func (r TeamBlockParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.BlockTeamRequest)
}

type TeamListAvailableParams struct {
	ResponseModel param.Field[interface{}] `query:"response_model"`
}

// URLQuery serializes [TeamListAvailableParams]'s query parameters as
// `url.Values`.
func (r TeamListAvailableParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type TeamRemoveMemberParams struct {
	TeamID    param.Field[string] `json:"team_id,required"`
	UserEmail param.Field[string] `json:"user_email"`
	UserID    param.Field[string] `json:"user_id"`
}

func (r TeamRemoveMemberParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TeamGetInfoParams struct {
	// Team ID in the request parameters
	TeamID param.Field[string] `query:"team_id"`
}

// URLQuery serializes [TeamGetInfoParams]'s query parameters as `url.Values`.
func (r TeamGetInfoParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type TeamUnblockParams struct {
	BlockTeamRequest BlockTeamRequestParam `json:"block_team_request,required"`
}

func (r TeamUnblockParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.BlockTeamRequest)
}

type TeamUpdateMemberParams struct {
	TeamID          param.Field[string]                     `json:"team_id,required"`
	MaxBudgetInTeam param.Field[float64]                    `json:"max_budget_in_team"`
	Role            param.Field[TeamUpdateMemberParamsRole] `json:"role"`
	// Requests per minute limit for this team member
	RpmLimit param.Field[int64] `json:"rpm_limit"`
	// Tokens per minute limit for this team member
	TpmLimit  param.Field[int64]  `json:"tpm_limit"`
	UserEmail param.Field[string] `json:"user_email"`
	UserID    param.Field[string] `json:"user_id"`
}

func (r TeamUpdateMemberParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TeamUpdateMemberParamsRole string

const (
	TeamUpdateMemberParamsRoleAdmin TeamUpdateMemberParamsRole = "admin"
	TeamUpdateMemberParamsRoleUser  TeamUpdateMemberParamsRole = "user"
)

func (r TeamUpdateMemberParamsRole) IsKnown() bool {
	switch r {
	case TeamUpdateMemberParamsRoleAdmin, TeamUpdateMemberParamsRoleUser:
		return true
	}
	return false
}
