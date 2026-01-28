// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hanzoai

import (
	"context"
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

// UserService contains methods and other services that help with interacting with
// the Hanzo API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserService] method instead.
type UserService struct {
	Options []option.RequestOption
}

// NewUserService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewUserService(opts ...option.RequestOption) (r *UserService) {
	r = &UserService{}
	r.Options = opts
	return
}

// Use this to create a new INTERNAL user with a budget. Internal Users can access
// LiteLLM Admin UI to make keys, request access to models. This creates a new user
// and generates a new api key for the new user. The new api key is returned.
//
// Returns user id, budget + new key.
//
// Parameters:
//
//   - user_id: Optional[str] - Specify a user id. If not set, a unique id will be
//     generated.
//   - user_alias: Optional[str] - A descriptive name for you to know who this user
//     id refers to.
//   - teams: Optional[list] - specify a list of team id's a user belongs to.
//   - user_email: Optional[str] - Specify a user email.
//   - send_invite_email: Optional[bool] - Specify if an invite email should be sent.
//   - user_role: Optional[str] - Specify a user role - "proxy_admin",
//     "proxy_admin_viewer", "internal_user", "internal_user_viewer", "team",
//     "customer". Info about each role here:
//     `https://github.com/BerriAI/litellm/litellm/proxy/_types.py#L20`
//   - max_budget: Optional[float] - Specify max budget for a given user.
//   - budget_duration: Optional[str] - Budget is reset at the end of specified
//     duration. If not set, budget is never reset. You can set duration as seconds
//     ("30s"), minutes ("30m"), hours ("30h"), days ("30d"), months ("1mo").
//   - models: Optional[list] - Model_name's a user is allowed to call. (if empty,
//     key is allowed to call all models). Set to ['no-default-models'] to block all
//     model access. Restricting user to only team-based model access.
//   - tpm_limit: Optional[int] - Specify tpm limit for a given user (Tokens per
//     minute)
//   - rpm_limit: Optional[int] - Specify rpm limit for a given user (Requests per
//     minute)
//   - auto_create_key: bool - Default=True. Flag used for returning a key as part of
//     the /user/new response
//   - aliases: Optional[dict] - Model aliases for the user -
//     [Docs](https://litellm.vercel.app/docs/proxy/virtual_keys#model-aliases)
//   - config: Optional[dict] - [DEPRECATED PARAM] User-specific config.
//   - allowed_cache_controls: Optional[list] - List of allowed cache control values.
//     Example - ["no-cache", "no-store"]. See all values -
//     https://docs.litellm.ai/docs/proxy/caching#turn-on--off-caching-per-request-
//   - blocked: Optional[bool] - [Not Implemented Yet] Whether the user is blocked.
//   - guardrails: Optional[List[str]] - [Not Implemented Yet] List of active
//     guardrails for the user
//   - permissions: Optional[dict] - [Not Implemented Yet] User-specific permissions,
//     eg. turning off pii masking.
//   - metadata: Optional[dict] - Metadata for user, store information for user.
//     Example metadata = {"team": "core-infra", "app": "app2", "email":
//     "ishaan@berri.ai" }
//   - max_parallel_requests: Optional[int] - Rate limit a user based on the number
//     of parallel requests. Raises 429 error, if user's parallel requests > x.
//   - soft_budget: Optional[float] - Get alerts when user crosses given budget,
//     doesn't block requests.
//   - model_max_budget: Optional[dict] - Model-specific max budget for user.
//     [Docs](https://docs.litellm.ai/docs/proxy/users#add-model-specific-budgets-to-keys)
//   - model_rpm_limit: Optional[float] - Model-specific rpm limit for user.
//     [Docs](https://docs.litellm.ai/docs/proxy/users#add-model-specific-limits-to-keys)
//   - model_tpm_limit: Optional[float] - Model-specific tpm limit for user.
//     [Docs](https://docs.litellm.ai/docs/proxy/users#add-model-specific-limits-to-keys)
//   - spend: Optional[float] - Amount spent by user. Default is 0. Will be updated
//     by proxy whenever user is used. You can set duration as seconds ("30s"),
//     minutes ("30m"), hours ("30h"), days ("30d"), months ("1mo").
//   - team_id: Optional[str] - [DEPRECATED PARAM] The team id of the user. Default
//     is None.
//   - duration: Optional[str] - Duration for the key auto-created on `/user/new`.
//     Default is None.
//   - key_alias: Optional[str] - Alias for the key auto-created on `/user/new`.
//     Default is None.
//   - sso_user_id: Optional[str] - The id of the user in the SSO provider.
//   - object_permission: Optional[LiteLLM_ObjectPermissionBase] - internal
//     user-specific object permission. Example - {"vector_stores":
//     ["vector_store_1", "vector_store_2"]}. IF null or {} then no object
//     permission.
//   - prompts: Optional[List[str]] - List of allowed prompts for the user. If
//     specified, the user will only be able to use these specific prompts.
//   - organizations: List[str] - List of organization id's the user is a member of
//     Returns:
//   - key: (str) The generated api key for the user
//   - expires: (datetime) Datetime object for when key expires.
//   - user_id: (str) Unique user id - used for tracking spend across multiple keys
//     for same user id.
//   - max_budget: (float|None) Max budget for given user.
//
// # Usage Example
//
// ```shell
//
//	curl -X POST "http://localhost:4000/user/new"      -H "Content-Type: application/json"      -H "Authorization: Bearer sk-1234"      -d '{
//	    "username": "new_user",
//	    "email": "new_user@example.com"
//	}'
//
// ```
func (r *UserService) New(ctx context.Context, body UserNewParams, opts ...option.RequestOption) (res *UserNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "user/new"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Example curl
//
// ```
//
//	curl --location 'http://0.0.0.0:4000/user/update'     --header 'Authorization: Bearer sk-1234'     --header 'Content-Type: application/json'     --data '{
//	    "user_id": "test-litellm-user-4",
//	    "user_role": "proxy_admin_viewer"
//	}'
//
// ```
//
// Parameters: - user_id: Optional[str] - Specify a user id. If not set, a unique
// id will be generated. - user_email: Optional[str] - Specify a user email. -
// password: Optional[str] - Specify a user password. - user_alias: Optional[str] -
// A descriptive name for you to know who this user id refers to. - teams:
// Optional[list] - specify a list of team id's a user belongs to. -
// send_invite_email: Optional[bool] - Specify if an invite email should be sent. -
// user_role: Optional[str] - Specify a user role - "proxy_admin",
// "proxy_admin_viewer", "internal_user", "internal_user_viewer", "team",
// "customer". Info about each role here:
// `https://github.com/BerriAI/litellm/litellm/proxy/_types.py#L20` - max_budget:
// Optional[float] - Specify max budget for a given user. - budget_duration:
// Optional[str] - Budget is reset at the end of specified duration. If not set,
// budget is never reset. You can set duration as seconds ("30s"), minutes ("30m"),
// hours ("30h"), days ("30d"), months ("1mo"). - models: Optional[list] -
// Model_name's a user is allowed to call. (if empty, key is allowed to call all
// models) - tpm_limit: Optional[int] - Specify tpm limit for a given user (Tokens
// per minute) - rpm_limit: Optional[int] - Specify rpm limit for a given user
// (Requests per minute) - auto_create_key: bool - Default=True. Flag used for
// returning a key as part of the /user/new response - aliases: Optional[dict] -
// Model aliases for the user -
// [Docs](https://litellm.vercel.app/docs/proxy/virtual_keys#model-aliases) -
// config: Optional[dict] - [DEPRECATED PARAM] User-specific config. -
// allowed_cache_controls: Optional[list] - List of allowed cache control values.
// Example - ["no-cache", "no-store"]. See all values -
// https://docs.litellm.ai/docs/proxy/caching#turn-on--off-caching-per-request- -
// blocked: Optional[bool] - [Not Implemented Yet] Whether the user is blocked. -
// guardrails: Optional[List[str]] - [Not Implemented Yet] List of active
// guardrails for the user - permissions: Optional[dict] - [Not Implemented Yet]
// User-specific permissions, eg. turning off pii masking. - metadata:
// Optional[dict] - Metadata for user, store information for user. Example metadata
// = {"team": "core-infra", "app": "app2", "email": "ishaan@berri.ai" } -
// max_parallel_requests: Optional[int] - Rate limit a user based on the number of
// parallel requests. Raises 429 error, if user's parallel requests > x. -
// soft_budget: Optional[float] - Get alerts when user crosses given budget,
// doesn't block requests. - model_max_budget: Optional[dict] - Model-specific max
// budget for user.
// [Docs](https://docs.litellm.ai/docs/proxy/users#add-model-specific-budgets-to-keys) -
// model_rpm_limit: Optional[float] - Model-specific rpm limit for user.
// [Docs](https://docs.litellm.ai/docs/proxy/users#add-model-specific-limits-to-keys) -
// model_tpm_limit: Optional[float] - Model-specific tpm limit for user.
// [Docs](https://docs.litellm.ai/docs/proxy/users#add-model-specific-limits-to-keys) -
// spend: Optional[float] - Amount spent by user. Default is 0. Will be updated by
// proxy whenever user is used. You can set duration as seconds ("30s"), minutes
// ("30m"), hours ("30h"), days ("30d"), months ("1mo"). - team_id: Optional[str] -
// [DEPRECATED PARAM] The team id of the user. Default is None. - duration:
// Optional[str] - [NOT IMPLEMENTED]. - key_alias: Optional[str] - [NOT
// IMPLEMENTED]. - object_permission: Optional[LiteLLM_ObjectPermissionBase] -
// internal user-specific object permission. Example - {"vector_stores":
// ["vector_store_1", "vector_store_2"]}. IF null or {} then no object
// permission. - prompts: Optional[List[str]] - List of allowed prompts for the
// user. If specified, the user will only be able to use these specific prompts.
func (r *UserService) Update(ctx context.Context, body UserUpdateParams, opts ...option.RequestOption) (res *UserUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "user/update"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// delete user and associated user keys
//
// ```
// curl --location 'http://0.0.0.0:4000/user/delete'
// --header 'Authorization: Bearer sk-1234'
// --header 'Content-Type: application/json'
//
//	--data-raw '{
//	    "user_ids": ["45e3e396-ee08-4a61-a88e-16b3ce7e0849"]
//	}'
//
// ```
//
// Parameters:
//
// - user_ids: List[str] - The list of user id's to be deleted.
func (r *UserService) Delete(ctx context.Context, params UserDeleteParams, opts ...option.RequestOption) (res *UserDeleteResponse, err error) {
	if params.LitellmChangedBy.Present {
		opts = append(opts, option.WithHeader("litellm-changed-by", fmt.Sprintf("%s", params.LitellmChangedBy)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "user/delete"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// [10/07/2024] Note: To get all users (+pagination), use `/user/list` endpoint.
//
// Use this to get user information. (user row + all user key info)
//
// # Example request
//
// ```
// curl -X GET 'http://localhost:4000/user/info?user_id=krrish7%40berri.ai'     --header 'Authorization: Bearer sk-1234'
// ```
func (r *UserService) GetInfo(ctx context.Context, query UserGetInfoParams, opts ...option.RequestOption) (res *UserGetInfoResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "user/info"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type UserNewResponse struct {
	Key                       string                                   `json:"key,required"`
	Token                     string                                   `json:"token,nullable"`
	Aliases                   map[string]interface{}                   `json:"aliases,nullable"`
	AllowedCacheControls      []interface{}                            `json:"allowed_cache_controls,nullable"`
	AllowedPassthroughRoutes  []interface{}                            `json:"allowed_passthrough_routes,nullable"`
	AllowedRoutes             []interface{}                            `json:"allowed_routes,nullable"`
	AllowedVectorStoreIndexes []UserNewResponseAllowedVectorStoreIndex `json:"allowed_vector_store_indexes,nullable"`
	Blocked                   bool                                     `json:"blocked,nullable"`
	BudgetDuration            string                                   `json:"budget_duration,nullable"`
	BudgetID                  string                                   `json:"budget_id,nullable"`
	Config                    map[string]interface{}                   `json:"config,nullable"`
	CreatedAt                 time.Time                                `json:"created_at,nullable" format:"date-time"`
	CreatedBy                 string                                   `json:"created_by,nullable"`
	Duration                  string                                   `json:"duration,nullable"`
	EnforcedParams            []string                                 `json:"enforced_params,nullable"`
	Expires                   time.Time                                `json:"expires,nullable" format:"date-time"`
	Guardrails                []string                                 `json:"guardrails,nullable"`
	KeyAlias                  string                                   `json:"key_alias,nullable"`
	KeyName                   string                                   `json:"key_name,nullable"`
	LitellmBudgetTable        interface{}                              `json:"litellm_budget_table"`
	MaxBudget                 float64                                  `json:"max_budget,nullable"`
	MaxParallelRequests       int64                                    `json:"max_parallel_requests,nullable"`
	Metadata                  map[string]interface{}                   `json:"metadata,nullable"`
	ModelMaxBudget            map[string]interface{}                   `json:"model_max_budget,nullable"`
	ModelRpmLimit             map[string]interface{}                   `json:"model_rpm_limit,nullable"`
	ModelTpmLimit             map[string]interface{}                   `json:"model_tpm_limit,nullable"`
	Models                    []interface{}                            `json:"models,nullable"`
	ObjectPermission          UserNewResponseObjectPermission          `json:"object_permission,nullable"`
	OrganizationID            string                                   `json:"organization_id,nullable"`
	Permissions               map[string]interface{}                   `json:"permissions,nullable"`
	Prompts                   []string                                 `json:"prompts,nullable"`
	// Set of params that you can modify via `router.update_settings()`.
	RouterSettings UserNewResponseRouterSettings `json:"router_settings,nullable"`
	RpmLimit       int64                         `json:"rpm_limit,nullable"`
	RpmLimitType   UserNewResponseRpmLimitType   `json:"rpm_limit_type,nullable"`
	Spend          float64                       `json:"spend,nullable"`
	Tags           []string                      `json:"tags,nullable"`
	TeamID         string                        `json:"team_id,nullable"`
	Teams          []interface{}                 `json:"teams,nullable"`
	TokenID        string                        `json:"token_id,nullable"`
	TpmLimit       int64                         `json:"tpm_limit,nullable"`
	TpmLimitType   UserNewResponseTpmLimitType   `json:"tpm_limit_type,nullable"`
	UpdatedAt      time.Time                     `json:"updated_at,nullable" format:"date-time"`
	UpdatedBy      string                        `json:"updated_by,nullable"`
	UserAlias      string                        `json:"user_alias,nullable"`
	UserEmail      string                        `json:"user_email,nullable"`
	UserID         string                        `json:"user_id,nullable"`
	UserRole       UserNewResponseUserRole       `json:"user_role,nullable"`
	JSON           userNewResponseJSON           `json:"-"`
}

// userNewResponseJSON contains the JSON metadata for the struct [UserNewResponse]
type userNewResponseJSON struct {
	Key                       apijson.Field
	Token                     apijson.Field
	Aliases                   apijson.Field
	AllowedCacheControls      apijson.Field
	AllowedPassthroughRoutes  apijson.Field
	AllowedRoutes             apijson.Field
	AllowedVectorStoreIndexes apijson.Field
	Blocked                   apijson.Field
	BudgetDuration            apijson.Field
	BudgetID                  apijson.Field
	Config                    apijson.Field
	CreatedAt                 apijson.Field
	CreatedBy                 apijson.Field
	Duration                  apijson.Field
	EnforcedParams            apijson.Field
	Expires                   apijson.Field
	Guardrails                apijson.Field
	KeyAlias                  apijson.Field
	KeyName                   apijson.Field
	LitellmBudgetTable        apijson.Field
	MaxBudget                 apijson.Field
	MaxParallelRequests       apijson.Field
	Metadata                  apijson.Field
	ModelMaxBudget            apijson.Field
	ModelRpmLimit             apijson.Field
	ModelTpmLimit             apijson.Field
	Models                    apijson.Field
	ObjectPermission          apijson.Field
	OrganizationID            apijson.Field
	Permissions               apijson.Field
	Prompts                   apijson.Field
	RouterSettings            apijson.Field
	RpmLimit                  apijson.Field
	RpmLimitType              apijson.Field
	Spend                     apijson.Field
	Tags                      apijson.Field
	TeamID                    apijson.Field
	Teams                     apijson.Field
	TokenID                   apijson.Field
	TpmLimit                  apijson.Field
	TpmLimitType              apijson.Field
	UpdatedAt                 apijson.Field
	UpdatedBy                 apijson.Field
	UserAlias                 apijson.Field
	UserEmail                 apijson.Field
	UserID                    apijson.Field
	UserRole                  apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *UserNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userNewResponseJSON) RawJSON() string {
	return r.raw
}

type UserNewResponseAllowedVectorStoreIndex struct {
	IndexName        string                                                    `json:"index_name,required"`
	IndexPermissions []UserNewResponseAllowedVectorStoreIndexesIndexPermission `json:"index_permissions,required"`
	JSON             userNewResponseAllowedVectorStoreIndexJSON                `json:"-"`
}

// userNewResponseAllowedVectorStoreIndexJSON contains the JSON metadata for the
// struct [UserNewResponseAllowedVectorStoreIndex]
type userNewResponseAllowedVectorStoreIndexJSON struct {
	IndexName        apijson.Field
	IndexPermissions apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *UserNewResponseAllowedVectorStoreIndex) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userNewResponseAllowedVectorStoreIndexJSON) RawJSON() string {
	return r.raw
}

type UserNewResponseAllowedVectorStoreIndexesIndexPermission string

const (
	UserNewResponseAllowedVectorStoreIndexesIndexPermissionRead  UserNewResponseAllowedVectorStoreIndexesIndexPermission = "read"
	UserNewResponseAllowedVectorStoreIndexesIndexPermissionWrite UserNewResponseAllowedVectorStoreIndexesIndexPermission = "write"
)

func (r UserNewResponseAllowedVectorStoreIndexesIndexPermission) IsKnown() bool {
	switch r {
	case UserNewResponseAllowedVectorStoreIndexesIndexPermissionRead, UserNewResponseAllowedVectorStoreIndexesIndexPermissionWrite:
		return true
	}
	return false
}

type UserNewResponseObjectPermission struct {
	AgentAccessGroups  []string                            `json:"agent_access_groups,nullable"`
	Agents             []string                            `json:"agents,nullable"`
	McpAccessGroups    []string                            `json:"mcp_access_groups,nullable"`
	McpServers         []string                            `json:"mcp_servers,nullable"`
	McpToolPermissions map[string][]string                 `json:"mcp_tool_permissions,nullable"`
	VectorStores       []string                            `json:"vector_stores,nullable"`
	JSON               userNewResponseObjectPermissionJSON `json:"-"`
}

// userNewResponseObjectPermissionJSON contains the JSON metadata for the struct
// [UserNewResponseObjectPermission]
type userNewResponseObjectPermissionJSON struct {
	AgentAccessGroups  apijson.Field
	Agents             apijson.Field
	McpAccessGroups    apijson.Field
	McpServers         apijson.Field
	McpToolPermissions apijson.Field
	VectorStores       apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *UserNewResponseObjectPermission) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userNewResponseObjectPermissionJSON) RawJSON() string {
	return r.raw
}

// Set of params that you can modify via `router.update_settings()`.
type UserNewResponseRouterSettings struct {
	AllowedFails           int64                                                        `json:"allowed_fails,nullable"`
	ContextWindowFallbacks []map[string]interface{}                                     `json:"context_window_fallbacks,nullable"`
	CooldownTime           float64                                                      `json:"cooldown_time,nullable"`
	Fallbacks              []map[string]interface{}                                     `json:"fallbacks,nullable"`
	MaxRetries             int64                                                        `json:"max_retries,nullable"`
	ModelGroupAlias        map[string]UserNewResponseRouterSettingsModelGroupAliasUnion `json:"model_group_alias,nullable"`
	ModelGroupRetryPolicy  map[string]interface{}                                       `json:"model_group_retry_policy,nullable"`
	NumRetries             int64                                                        `json:"num_retries,nullable"`
	RetryAfter             float64                                                      `json:"retry_after,nullable"`
	RoutingStrategy        string                                                       `json:"routing_strategy,nullable"`
	RoutingStrategyArgs    map[string]interface{}                                       `json:"routing_strategy_args,nullable"`
	Timeout                float64                                                      `json:"timeout,nullable"`
	JSON                   userNewResponseRouterSettingsJSON                            `json:"-"`
}

// userNewResponseRouterSettingsJSON contains the JSON metadata for the struct
// [UserNewResponseRouterSettings]
type userNewResponseRouterSettingsJSON struct {
	AllowedFails           apijson.Field
	ContextWindowFallbacks apijson.Field
	CooldownTime           apijson.Field
	Fallbacks              apijson.Field
	MaxRetries             apijson.Field
	ModelGroupAlias        apijson.Field
	ModelGroupRetryPolicy  apijson.Field
	NumRetries             apijson.Field
	RetryAfter             apijson.Field
	RoutingStrategy        apijson.Field
	RoutingStrategyArgs    apijson.Field
	Timeout                apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *UserNewResponseRouterSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userNewResponseRouterSettingsJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [shared.UnionString] or
// [UserNewResponseRouterSettingsModelGroupAliasMap].
type UserNewResponseRouterSettingsModelGroupAliasUnion interface {
	ImplementsUserNewResponseRouterSettingsModelGroupAliasUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserNewResponseRouterSettingsModelGroupAliasUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserNewResponseRouterSettingsModelGroupAliasMap{}),
		},
	)
}

type UserNewResponseRouterSettingsModelGroupAliasMap map[string]interface{}

func (r UserNewResponseRouterSettingsModelGroupAliasMap) ImplementsUserNewResponseRouterSettingsModelGroupAliasUnion() {
}

type UserNewResponseRpmLimitType string

const (
	UserNewResponseRpmLimitTypeGuaranteedThroughput UserNewResponseRpmLimitType = "guaranteed_throughput"
	UserNewResponseRpmLimitTypeBestEffortThroughput UserNewResponseRpmLimitType = "best_effort_throughput"
	UserNewResponseRpmLimitTypeDynamic              UserNewResponseRpmLimitType = "dynamic"
)

func (r UserNewResponseRpmLimitType) IsKnown() bool {
	switch r {
	case UserNewResponseRpmLimitTypeGuaranteedThroughput, UserNewResponseRpmLimitTypeBestEffortThroughput, UserNewResponseRpmLimitTypeDynamic:
		return true
	}
	return false
}

type UserNewResponseTpmLimitType string

const (
	UserNewResponseTpmLimitTypeGuaranteedThroughput UserNewResponseTpmLimitType = "guaranteed_throughput"
	UserNewResponseTpmLimitTypeBestEffortThroughput UserNewResponseTpmLimitType = "best_effort_throughput"
	UserNewResponseTpmLimitTypeDynamic              UserNewResponseTpmLimitType = "dynamic"
)

func (r UserNewResponseTpmLimitType) IsKnown() bool {
	switch r {
	case UserNewResponseTpmLimitTypeGuaranteedThroughput, UserNewResponseTpmLimitTypeBestEffortThroughput, UserNewResponseTpmLimitTypeDynamic:
		return true
	}
	return false
}

type UserNewResponseUserRole string

const (
	UserNewResponseUserRoleProxyAdmin         UserNewResponseUserRole = "proxy_admin"
	UserNewResponseUserRoleProxyAdminViewer   UserNewResponseUserRole = "proxy_admin_viewer"
	UserNewResponseUserRoleInternalUser       UserNewResponseUserRole = "internal_user"
	UserNewResponseUserRoleInternalUserViewer UserNewResponseUserRole = "internal_user_viewer"
)

func (r UserNewResponseUserRole) IsKnown() bool {
	switch r {
	case UserNewResponseUserRoleProxyAdmin, UserNewResponseUserRoleProxyAdminViewer, UserNewResponseUserRoleInternalUser, UserNewResponseUserRoleInternalUserViewer:
		return true
	}
	return false
}

type UserUpdateResponse = interface{}

type UserDeleteResponse = interface{}

type UserGetInfoResponse = interface{}

type UserNewParams struct {
	Aliases              param.Field[map[string]interface{}]        `json:"aliases"`
	AllowedCacheControls param.Field[[]interface{}]                 `json:"allowed_cache_controls"`
	AutoCreateKey        param.Field[bool]                          `json:"auto_create_key"`
	Blocked              param.Field[bool]                          `json:"blocked"`
	BudgetDuration       param.Field[string]                        `json:"budget_duration"`
	Config               param.Field[map[string]interface{}]        `json:"config"`
	Duration             param.Field[string]                        `json:"duration"`
	Guardrails           param.Field[[]string]                      `json:"guardrails"`
	KeyAlias             param.Field[string]                        `json:"key_alias"`
	MaxBudget            param.Field[float64]                       `json:"max_budget"`
	MaxParallelRequests  param.Field[int64]                         `json:"max_parallel_requests"`
	Metadata             param.Field[map[string]interface{}]        `json:"metadata"`
	ModelMaxBudget       param.Field[map[string]interface{}]        `json:"model_max_budget"`
	ModelRpmLimit        param.Field[map[string]interface{}]        `json:"model_rpm_limit"`
	ModelTpmLimit        param.Field[map[string]interface{}]        `json:"model_tpm_limit"`
	Models               param.Field[[]interface{}]                 `json:"models"`
	ObjectPermission     param.Field[UserNewParamsObjectPermission] `json:"object_permission"`
	Organizations        param.Field[[]string]                      `json:"organizations"`
	Permissions          param.Field[map[string]interface{}]        `json:"permissions"`
	Prompts              param.Field[[]string]                      `json:"prompts"`
	RpmLimit             param.Field[int64]                         `json:"rpm_limit"`
	SendInviteEmail      param.Field[bool]                          `json:"send_invite_email"`
	Spend                param.Field[float64]                       `json:"spend"`
	SSOUserID            param.Field[string]                        `json:"sso_user_id"`
	TeamID               param.Field[string]                        `json:"team_id"`
	Teams                param.Field[UserNewParamsTeamsUnion]       `json:"teams"`
	TpmLimit             param.Field[int64]                         `json:"tpm_limit"`
	UserAlias            param.Field[string]                        `json:"user_alias"`
	UserEmail            param.Field[string]                        `json:"user_email"`
	UserID               param.Field[string]                        `json:"user_id"`
	UserRole             param.Field[UserNewParamsUserRole]         `json:"user_role"`
}

func (r UserNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserNewParamsObjectPermission struct {
	AgentAccessGroups  param.Field[[]string]            `json:"agent_access_groups"`
	Agents             param.Field[[]string]            `json:"agents"`
	McpAccessGroups    param.Field[[]string]            `json:"mcp_access_groups"`
	McpServers         param.Field[[]string]            `json:"mcp_servers"`
	McpToolPermissions param.Field[map[string][]string] `json:"mcp_tool_permissions"`
	VectorStores       param.Field[[]string]            `json:"vector_stores"`
}

func (r UserNewParamsObjectPermission) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [UserNewParamsTeamsArray], [UserNewParamsTeamsArray].
type UserNewParamsTeamsUnion interface {
	implementsUserNewParamsTeamsUnion()
}

type UserNewParamsTeamsArray []string

func (r UserNewParamsTeamsArray) implementsUserNewParamsTeamsUnion() {}

type UserNewParamsUserRole string

const (
	UserNewParamsUserRoleProxyAdmin         UserNewParamsUserRole = "proxy_admin"
	UserNewParamsUserRoleProxyAdminViewer   UserNewParamsUserRole = "proxy_admin_viewer"
	UserNewParamsUserRoleInternalUser       UserNewParamsUserRole = "internal_user"
	UserNewParamsUserRoleInternalUserViewer UserNewParamsUserRole = "internal_user_viewer"
)

func (r UserNewParamsUserRole) IsKnown() bool {
	switch r {
	case UserNewParamsUserRoleProxyAdmin, UserNewParamsUserRoleProxyAdminViewer, UserNewParamsUserRoleInternalUser, UserNewParamsUserRoleInternalUserViewer:
		return true
	}
	return false
}

type UserUpdateParams struct {
	Aliases              param.Field[map[string]interface{}]           `json:"aliases"`
	AllowedCacheControls param.Field[[]interface{}]                    `json:"allowed_cache_controls"`
	Blocked              param.Field[bool]                             `json:"blocked"`
	BudgetDuration       param.Field[string]                           `json:"budget_duration"`
	Config               param.Field[map[string]interface{}]           `json:"config"`
	Duration             param.Field[string]                           `json:"duration"`
	Guardrails           param.Field[[]string]                         `json:"guardrails"`
	KeyAlias             param.Field[string]                           `json:"key_alias"`
	MaxBudget            param.Field[float64]                          `json:"max_budget"`
	MaxParallelRequests  param.Field[int64]                            `json:"max_parallel_requests"`
	Metadata             param.Field[map[string]interface{}]           `json:"metadata"`
	ModelMaxBudget       param.Field[map[string]interface{}]           `json:"model_max_budget"`
	ModelRpmLimit        param.Field[map[string]interface{}]           `json:"model_rpm_limit"`
	ModelTpmLimit        param.Field[map[string]interface{}]           `json:"model_tpm_limit"`
	Models               param.Field[[]interface{}]                    `json:"models"`
	ObjectPermission     param.Field[UserUpdateParamsObjectPermission] `json:"object_permission"`
	Password             param.Field[string]                           `json:"password"`
	Permissions          param.Field[map[string]interface{}]           `json:"permissions"`
	Prompts              param.Field[[]string]                         `json:"prompts"`
	RpmLimit             param.Field[int64]                            `json:"rpm_limit"`
	Spend                param.Field[float64]                          `json:"spend"`
	TeamID               param.Field[string]                           `json:"team_id"`
	TpmLimit             param.Field[int64]                            `json:"tpm_limit"`
	UserAlias            param.Field[string]                           `json:"user_alias"`
	UserEmail            param.Field[string]                           `json:"user_email"`
	UserID               param.Field[string]                           `json:"user_id"`
	UserRole             param.Field[UserUpdateParamsUserRole]         `json:"user_role"`
}

func (r UserUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserUpdateParamsObjectPermission struct {
	AgentAccessGroups  param.Field[[]string]            `json:"agent_access_groups"`
	Agents             param.Field[[]string]            `json:"agents"`
	McpAccessGroups    param.Field[[]string]            `json:"mcp_access_groups"`
	McpServers         param.Field[[]string]            `json:"mcp_servers"`
	McpToolPermissions param.Field[map[string][]string] `json:"mcp_tool_permissions"`
	VectorStores       param.Field[[]string]            `json:"vector_stores"`
}

func (r UserUpdateParamsObjectPermission) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserUpdateParamsUserRole string

const (
	UserUpdateParamsUserRoleProxyAdmin         UserUpdateParamsUserRole = "proxy_admin"
	UserUpdateParamsUserRoleProxyAdminViewer   UserUpdateParamsUserRole = "proxy_admin_viewer"
	UserUpdateParamsUserRoleInternalUser       UserUpdateParamsUserRole = "internal_user"
	UserUpdateParamsUserRoleInternalUserViewer UserUpdateParamsUserRole = "internal_user_viewer"
)

func (r UserUpdateParamsUserRole) IsKnown() bool {
	switch r {
	case UserUpdateParamsUserRoleProxyAdmin, UserUpdateParamsUserRoleProxyAdminViewer, UserUpdateParamsUserRoleInternalUser, UserUpdateParamsUserRoleInternalUserViewer:
		return true
	}
	return false
}

type UserDeleteParams struct {
	UserIDs param.Field[[]string] `json:"user_ids,required"`
	// The litellm-changed-by header enables tracking of actions performed by
	// authorized users on behalf of other users, providing an audit trail for
	// accountability
	LitellmChangedBy param.Field[string] `header:"litellm-changed-by"`
}

func (r UserDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserGetInfoParams struct {
	// User ID in the request parameters
	UserID param.Field[string] `query:"user_id"`
}

// URLQuery serializes [UserGetInfoParams]'s query parameters as `url.Values`.
func (r UserGetInfoParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
