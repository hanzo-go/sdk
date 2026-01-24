// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hanzoai

import (
	"github.com/hanzoai/go-sdk/internal/apijson"
	"github.com/hanzoai/go-sdk/internal/param"
	"github.com/hanzoai/go-sdk/option"
)

// KeyRegenerateService contains methods and other services that help with
// interacting with the Hanzo API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewKeyRegenerateService] method instead.
type KeyRegenerateService struct {
	Options []option.RequestOption
}

// NewKeyRegenerateService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewKeyRegenerateService(opts ...option.RequestOption) (r *KeyRegenerateService) {
	r = &KeyRegenerateService{}
	r.Options = opts
	return
}

type RegenerateKeyRequestParam struct {
	Aliases                   param.Field[map[string]interface{}]                             `json:"aliases"`
	AllowedCacheControls      param.Field[[]interface{}]                                      `json:"allowed_cache_controls"`
	AllowedPassthroughRoutes  param.Field[[]interface{}]                                      `json:"allowed_passthrough_routes"`
	AllowedRoutes             param.Field[[]interface{}]                                      `json:"allowed_routes"`
	AllowedVectorStoreIndexes param.Field[[]RegenerateKeyRequestAllowedVectorStoreIndexParam] `json:"allowed_vector_store_indexes"`
	// Whether this key should be automatically rotated
	AutoRotate     param.Field[bool]                   `json:"auto_rotate"`
	Blocked        param.Field[bool]                   `json:"blocked"`
	BudgetDuration param.Field[string]                 `json:"budget_duration"`
	BudgetID       param.Field[string]                 `json:"budget_id"`
	Config         param.Field[map[string]interface{}] `json:"config"`
	Duration       param.Field[string]                 `json:"duration"`
	EnforcedParams param.Field[[]string]               `json:"enforced_params"`
	Guardrails     param.Field[[]string]               `json:"guardrails"`
	Key            param.Field[string]                 `json:"key"`
	KeyAlias       param.Field[string]                 `json:"key_alias"`
	// Enum for key types that determine what routes a key can access
	KeyType             param.Field[RegenerateKeyRequestKeyType]               `json:"key_type"`
	MaxBudget           param.Field[float64]                                   `json:"max_budget"`
	MaxParallelRequests param.Field[int64]                                     `json:"max_parallel_requests"`
	Metadata            param.Field[map[string]interface{}]                    `json:"metadata"`
	ModelMaxBudget      param.Field[map[string]interface{}]                    `json:"model_max_budget"`
	ModelRpmLimit       param.Field[map[string]interface{}]                    `json:"model_rpm_limit"`
	ModelTpmLimit       param.Field[map[string]interface{}]                    `json:"model_tpm_limit"`
	Models              param.Field[[]interface{}]                             `json:"models"`
	NewKey              param.Field[string]                                    `json:"new_key"`
	NewMasterKey        param.Field[string]                                    `json:"new_master_key"`
	ObjectPermission    param.Field[RegenerateKeyRequestObjectPermissionParam] `json:"object_permission"`
	OrganizationID      param.Field[string]                                    `json:"organization_id"`
	Permissions         param.Field[map[string]interface{}]                    `json:"permissions"`
	Prompts             param.Field[[]string]                                  `json:"prompts"`
	// How often to rotate this key (e.g., '30d', '90d'). Required if auto_rotate=True
	RotationInterval param.Field[string] `json:"rotation_interval"`
	// Set of params that you can modify via `router.update_settings()`.
	RouterSettings  param.Field[RegenerateKeyRequestRouterSettingsParam] `json:"router_settings"`
	RpmLimit        param.Field[int64]                                   `json:"rpm_limit"`
	RpmLimitType    param.Field[RegenerateKeyRequestRpmLimitType]        `json:"rpm_limit_type"`
	SendInviteEmail param.Field[bool]                                    `json:"send_invite_email"`
	SoftBudget      param.Field[float64]                                 `json:"soft_budget"`
	Spend           param.Field[float64]                                 `json:"spend"`
	Tags            param.Field[[]string]                                `json:"tags"`
	TeamID          param.Field[string]                                  `json:"team_id"`
	TpmLimit        param.Field[int64]                                   `json:"tpm_limit"`
	TpmLimitType    param.Field[RegenerateKeyRequestTpmLimitType]        `json:"tpm_limit_type"`
	UserID          param.Field[string]                                  `json:"user_id"`
}

func (r RegenerateKeyRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RegenerateKeyRequestAllowedVectorStoreIndexParam struct {
	IndexName        param.Field[string]                                                         `json:"index_name,required"`
	IndexPermissions param.Field[[]RegenerateKeyRequestAllowedVectorStoreIndexesIndexPermission] `json:"index_permissions,required"`
}

func (r RegenerateKeyRequestAllowedVectorStoreIndexParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RegenerateKeyRequestAllowedVectorStoreIndexesIndexPermission string

const (
	RegenerateKeyRequestAllowedVectorStoreIndexesIndexPermissionRead  RegenerateKeyRequestAllowedVectorStoreIndexesIndexPermission = "read"
	RegenerateKeyRequestAllowedVectorStoreIndexesIndexPermissionWrite RegenerateKeyRequestAllowedVectorStoreIndexesIndexPermission = "write"
)

func (r RegenerateKeyRequestAllowedVectorStoreIndexesIndexPermission) IsKnown() bool {
	switch r {
	case RegenerateKeyRequestAllowedVectorStoreIndexesIndexPermissionRead, RegenerateKeyRequestAllowedVectorStoreIndexesIndexPermissionWrite:
		return true
	}
	return false
}

// Enum for key types that determine what routes a key can access
type RegenerateKeyRequestKeyType string

const (
	RegenerateKeyRequestKeyTypeLlmAPI     RegenerateKeyRequestKeyType = "llm_api"
	RegenerateKeyRequestKeyTypeManagement RegenerateKeyRequestKeyType = "management"
	RegenerateKeyRequestKeyTypeReadOnly   RegenerateKeyRequestKeyType = "read_only"
	RegenerateKeyRequestKeyTypeDefault    RegenerateKeyRequestKeyType = "default"
)

func (r RegenerateKeyRequestKeyType) IsKnown() bool {
	switch r {
	case RegenerateKeyRequestKeyTypeLlmAPI, RegenerateKeyRequestKeyTypeManagement, RegenerateKeyRequestKeyTypeReadOnly, RegenerateKeyRequestKeyTypeDefault:
		return true
	}
	return false
}

type RegenerateKeyRequestObjectPermissionParam struct {
	AgentAccessGroups  param.Field[[]string]            `json:"agent_access_groups"`
	Agents             param.Field[[]string]            `json:"agents"`
	McpAccessGroups    param.Field[[]string]            `json:"mcp_access_groups"`
	McpServers         param.Field[[]string]            `json:"mcp_servers"`
	McpToolPermissions param.Field[map[string][]string] `json:"mcp_tool_permissions"`
	VectorStores       param.Field[[]string]            `json:"vector_stores"`
}

func (r RegenerateKeyRequestObjectPermissionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Set of params that you can modify via `router.update_settings()`.
type RegenerateKeyRequestRouterSettingsParam struct {
	AllowedFails           param.Field[int64]                                                                  `json:"allowed_fails"`
	ContextWindowFallbacks param.Field[[]map[string]interface{}]                                               `json:"context_window_fallbacks"`
	CooldownTime           param.Field[float64]                                                                `json:"cooldown_time"`
	Fallbacks              param.Field[[]map[string]interface{}]                                               `json:"fallbacks"`
	MaxRetries             param.Field[int64]                                                                  `json:"max_retries"`
	ModelGroupAlias        param.Field[map[string]RegenerateKeyRequestRouterSettingsModelGroupAliasUnionParam] `json:"model_group_alias"`
	ModelGroupRetryPolicy  param.Field[map[string]interface{}]                                                 `json:"model_group_retry_policy"`
	NumRetries             param.Field[int64]                                                                  `json:"num_retries"`
	RetryAfter             param.Field[float64]                                                                `json:"retry_after"`
	RoutingStrategy        param.Field[string]                                                                 `json:"routing_strategy"`
	RoutingStrategyArgs    param.Field[map[string]interface{}]                                                 `json:"routing_strategy_args"`
	Timeout                param.Field[float64]                                                                `json:"timeout"`
}

func (r RegenerateKeyRequestRouterSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [shared.UnionString],
// [RegenerateKeyRequestRouterSettingsModelGroupAliasMapParam].
type RegenerateKeyRequestRouterSettingsModelGroupAliasUnionParam interface {
	ImplementsRegenerateKeyRequestRouterSettingsModelGroupAliasUnionParam()
}

type RegenerateKeyRequestRouterSettingsModelGroupAliasMapParam map[string]interface{}

func (r RegenerateKeyRequestRouterSettingsModelGroupAliasMapParam) ImplementsRegenerateKeyRequestRouterSettingsModelGroupAliasUnionParam() {
}

type RegenerateKeyRequestRpmLimitType string

const (
	RegenerateKeyRequestRpmLimitTypeGuaranteedThroughput RegenerateKeyRequestRpmLimitType = "guaranteed_throughput"
	RegenerateKeyRequestRpmLimitTypeBestEffortThroughput RegenerateKeyRequestRpmLimitType = "best_effort_throughput"
	RegenerateKeyRequestRpmLimitTypeDynamic              RegenerateKeyRequestRpmLimitType = "dynamic"
)

func (r RegenerateKeyRequestRpmLimitType) IsKnown() bool {
	switch r {
	case RegenerateKeyRequestRpmLimitTypeGuaranteedThroughput, RegenerateKeyRequestRpmLimitTypeBestEffortThroughput, RegenerateKeyRequestRpmLimitTypeDynamic:
		return true
	}
	return false
}

type RegenerateKeyRequestTpmLimitType string

const (
	RegenerateKeyRequestTpmLimitTypeGuaranteedThroughput RegenerateKeyRequestTpmLimitType = "guaranteed_throughput"
	RegenerateKeyRequestTpmLimitTypeBestEffortThroughput RegenerateKeyRequestTpmLimitType = "best_effort_throughput"
	RegenerateKeyRequestTpmLimitTypeDynamic              RegenerateKeyRequestTpmLimitType = "dynamic"
)

func (r RegenerateKeyRequestTpmLimitType) IsKnown() bool {
	switch r {
	case RegenerateKeyRequestTpmLimitTypeGuaranteedThroughput, RegenerateKeyRequestTpmLimitTypeBestEffortThroughput, RegenerateKeyRequestTpmLimitTypeDynamic:
		return true
	}
	return false
}
