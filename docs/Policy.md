# Policy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CachePaths** | Pointer to **map[string]int64** | CachePaths overrides CacheTTLSec per path PREFIX (key \&quot;/v1/models\&quot; → seconds). The longest matching prefix wins. | [optional] 
**CacheTtlSec** | Pointer to **int64** | CacheTTLSec is the org&#39;s default edge-cache TTL for its responses, in seconds; 0 means no caching. Unset inherits the platform default. | [optional] 
**CorsOrigins** | Pointer to **[]string** | CORSOrigins is the PLATFORM-scope CORS allowlist EdgeCORS admits: an exact origin, a bare host, or a \&quot;*.host\&quot; wildcard. Writable only by a SuperAdmin — CORS is evaluated before identity, so it has no tenant to scope to. | [optional] 
**Methods** | Pointer to **[]string** | Methods is the allowlist of HTTP methods the edge accepts for this org. Empty means all are accepted. | [optional] 
**Mode** | Pointer to **string** | Mode is the abuse gate&#39;s posture for THIS scope: \&quot;shadow\&quot; scores traffic and records the verdict without acting on it, \&quot;live\&quot; enforces it. Unset means shadow.  It is the one per-org field that does NOT inherit. Every other field here layers a platform default under the org&#39;s own value, which is right for a default: a tenant that sets no rate ceiling should get the platform&#39;s. Mode is not a default, it is an ARMING DECISION — it is what makes a statistical judgement start refusing real traffic — and inheriting it means arming one scope arms every tenant that never asked for it, without a write to their row and without anything in their config changing. So a tenant is live only if that tenant&#39;s OWN row says live, and the platform row&#39;s mode governs exactly one scope: the anonymous lane, which has no tenant of its own.  It is also not self-service. Writing it requires SuperAdmin (see the /v1/gateway config op): the subject of an abuse control does not get to switch the control off. | [optional] 
**OrgRpm** | Pointer to **int64** | OrgRPM is the org&#39;s OWN authenticated rate ceiling, requests per minute, as ScopeRateLimit enforces it. Unset inherits the platform default, then the static boot default. | [optional] 
**PerIpRpm** | Pointer to **int64** | PerIPRPM is the PLATFORM-scope pre-auth flood cap: requests EdgeRateLimit admits per WindowSec from one client IP. SuperAdmin-only, same reason. | [optional] 
**UpdatedAt** | Pointer to **int64** | UpdatedAt is the unix second this policy row was last written. Server-stamped; a client-supplied value is ignored. | [optional] 
**UpdatedBy** | Pointer to **string** | UpdatedBy is the validated user id that wrote this policy row. Server-stamped; a client-supplied value is ignored. | [optional] 
**WindowSec** | Pointer to **int64** | WindowSec is the window PerIPRPM is counted over, in seconds. SuperAdmin-only. | [optional] 

## Methods

### NewPolicy

`func NewPolicy() *Policy`

NewPolicy instantiates a new Policy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyWithDefaults

`func NewPolicyWithDefaults() *Policy`

NewPolicyWithDefaults instantiates a new Policy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCachePaths

`func (o *Policy) GetCachePaths() map[string]int64`

GetCachePaths returns the CachePaths field if non-nil, zero value otherwise.

### GetCachePathsOk

`func (o *Policy) GetCachePathsOk() (*map[string]int64, bool)`

GetCachePathsOk returns a tuple with the CachePaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCachePaths

`func (o *Policy) SetCachePaths(v map[string]int64)`

SetCachePaths sets CachePaths field to given value.

### HasCachePaths

`func (o *Policy) HasCachePaths() bool`

HasCachePaths returns a boolean if a field has been set.

### GetCacheTtlSec

`func (o *Policy) GetCacheTtlSec() int64`

GetCacheTtlSec returns the CacheTtlSec field if non-nil, zero value otherwise.

### GetCacheTtlSecOk

`func (o *Policy) GetCacheTtlSecOk() (*int64, bool)`

GetCacheTtlSecOk returns a tuple with the CacheTtlSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheTtlSec

`func (o *Policy) SetCacheTtlSec(v int64)`

SetCacheTtlSec sets CacheTtlSec field to given value.

### HasCacheTtlSec

`func (o *Policy) HasCacheTtlSec() bool`

HasCacheTtlSec returns a boolean if a field has been set.

### GetCorsOrigins

`func (o *Policy) GetCorsOrigins() []string`

GetCorsOrigins returns the CorsOrigins field if non-nil, zero value otherwise.

### GetCorsOriginsOk

`func (o *Policy) GetCorsOriginsOk() (*[]string, bool)`

GetCorsOriginsOk returns a tuple with the CorsOrigins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorsOrigins

`func (o *Policy) SetCorsOrigins(v []string)`

SetCorsOrigins sets CorsOrigins field to given value.

### HasCorsOrigins

`func (o *Policy) HasCorsOrigins() bool`

HasCorsOrigins returns a boolean if a field has been set.

### GetMethods

`func (o *Policy) GetMethods() []string`

GetMethods returns the Methods field if non-nil, zero value otherwise.

### GetMethodsOk

`func (o *Policy) GetMethodsOk() (*[]string, bool)`

GetMethodsOk returns a tuple with the Methods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethods

`func (o *Policy) SetMethods(v []string)`

SetMethods sets Methods field to given value.

### HasMethods

`func (o *Policy) HasMethods() bool`

HasMethods returns a boolean if a field has been set.

### GetMode

`func (o *Policy) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *Policy) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *Policy) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *Policy) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetOrgRpm

`func (o *Policy) GetOrgRpm() int64`

GetOrgRpm returns the OrgRpm field if non-nil, zero value otherwise.

### GetOrgRpmOk

`func (o *Policy) GetOrgRpmOk() (*int64, bool)`

GetOrgRpmOk returns a tuple with the OrgRpm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgRpm

`func (o *Policy) SetOrgRpm(v int64)`

SetOrgRpm sets OrgRpm field to given value.

### HasOrgRpm

`func (o *Policy) HasOrgRpm() bool`

HasOrgRpm returns a boolean if a field has been set.

### GetPerIpRpm

`func (o *Policy) GetPerIpRpm() int64`

GetPerIpRpm returns the PerIpRpm field if non-nil, zero value otherwise.

### GetPerIpRpmOk

`func (o *Policy) GetPerIpRpmOk() (*int64, bool)`

GetPerIpRpmOk returns a tuple with the PerIpRpm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerIpRpm

`func (o *Policy) SetPerIpRpm(v int64)`

SetPerIpRpm sets PerIpRpm field to given value.

### HasPerIpRpm

`func (o *Policy) HasPerIpRpm() bool`

HasPerIpRpm returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Policy) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Policy) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Policy) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Policy) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *Policy) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *Policy) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *Policy) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *Policy) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetWindowSec

`func (o *Policy) GetWindowSec() int64`

GetWindowSec returns the WindowSec field if non-nil, zero value otherwise.

### GetWindowSecOk

`func (o *Policy) GetWindowSecOk() (*int64, bool)`

GetWindowSecOk returns a tuple with the WindowSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowSec

`func (o *Policy) SetWindowSec(v int64)`

SetWindowSec sets WindowSec field to given value.

### HasWindowSec

`func (o *Policy) HasWindowSec() bool`

HasWindowSec returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


