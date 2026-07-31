# CloudPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CachePaths** | Pointer to **map[string]int32** | CachePaths overrides CacheTTLSec per path PREFIX (key \&quot;/v1/models\&quot; → seconds). The longest matching prefix wins. | [optional] 
**CacheTtlSec** | Pointer to **int32** | CacheTTLSec is the org&#39;s default edge-cache TTL for its responses, in seconds; 0 means no caching. Unset inherits the platform default. | [optional] 
**CorsOrigins** | Pointer to **[]string** | CORSOrigins is the PLATFORM-scope CORS allowlist EdgeCORS admits: an exact origin, a bare host, or a \&quot;*.host\&quot; wildcard. Writable only by a SuperAdmin — CORS is evaluated before identity, so it has no tenant to scope to. | [optional] 
**Methods** | Pointer to **[]string** | Methods is the allowlist of HTTP methods the edge accepts for this org. Empty means all are accepted. | [optional] 
**OrgRpm** | Pointer to **int32** | OrgRPM is the org&#39;s OWN authenticated rate ceiling, requests per minute, as ScopeRateLimit enforces it. Unset inherits the platform default, then the static boot default. | [optional] 
**PerIpRpm** | Pointer to **int32** | PerIPRPM is the PLATFORM-scope pre-auth flood cap: requests EdgeRateLimit admits per WindowSec from one client IP. SuperAdmin-only, same reason. | [optional] 
**UpdatedAt** | Pointer to **int32** | UpdatedAt is the unix second this policy row was last written. Server-stamped; a client-supplied value is ignored. | [optional] 
**UpdatedBy** | Pointer to **string** | UpdatedBy is the validated user id that wrote this policy row. Server-stamped; a client-supplied value is ignored. | [optional] 
**WindowSec** | Pointer to **int32** | WindowSec is the window PerIPRPM is counted over, in seconds. SuperAdmin-only. | [optional] 

## Methods

### NewCloudPolicy

`func NewCloudPolicy() *CloudPolicy`

NewCloudPolicy instantiates a new CloudPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudPolicyWithDefaults

`func NewCloudPolicyWithDefaults() *CloudPolicy`

NewCloudPolicyWithDefaults instantiates a new CloudPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCachePaths

`func (o *CloudPolicy) GetCachePaths() map[string]int32`

GetCachePaths returns the CachePaths field if non-nil, zero value otherwise.

### GetCachePathsOk

`func (o *CloudPolicy) GetCachePathsOk() (*map[string]int32, bool)`

GetCachePathsOk returns a tuple with the CachePaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCachePaths

`func (o *CloudPolicy) SetCachePaths(v map[string]int32)`

SetCachePaths sets CachePaths field to given value.

### HasCachePaths

`func (o *CloudPolicy) HasCachePaths() bool`

HasCachePaths returns a boolean if a field has been set.

### GetCacheTtlSec

`func (o *CloudPolicy) GetCacheTtlSec() int32`

GetCacheTtlSec returns the CacheTtlSec field if non-nil, zero value otherwise.

### GetCacheTtlSecOk

`func (o *CloudPolicy) GetCacheTtlSecOk() (*int32, bool)`

GetCacheTtlSecOk returns a tuple with the CacheTtlSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheTtlSec

`func (o *CloudPolicy) SetCacheTtlSec(v int32)`

SetCacheTtlSec sets CacheTtlSec field to given value.

### HasCacheTtlSec

`func (o *CloudPolicy) HasCacheTtlSec() bool`

HasCacheTtlSec returns a boolean if a field has been set.

### GetCorsOrigins

`func (o *CloudPolicy) GetCorsOrigins() []string`

GetCorsOrigins returns the CorsOrigins field if non-nil, zero value otherwise.

### GetCorsOriginsOk

`func (o *CloudPolicy) GetCorsOriginsOk() (*[]string, bool)`

GetCorsOriginsOk returns a tuple with the CorsOrigins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorsOrigins

`func (o *CloudPolicy) SetCorsOrigins(v []string)`

SetCorsOrigins sets CorsOrigins field to given value.

### HasCorsOrigins

`func (o *CloudPolicy) HasCorsOrigins() bool`

HasCorsOrigins returns a boolean if a field has been set.

### GetMethods

`func (o *CloudPolicy) GetMethods() []string`

GetMethods returns the Methods field if non-nil, zero value otherwise.

### GetMethodsOk

`func (o *CloudPolicy) GetMethodsOk() (*[]string, bool)`

GetMethodsOk returns a tuple with the Methods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethods

`func (o *CloudPolicy) SetMethods(v []string)`

SetMethods sets Methods field to given value.

### HasMethods

`func (o *CloudPolicy) HasMethods() bool`

HasMethods returns a boolean if a field has been set.

### GetOrgRpm

`func (o *CloudPolicy) GetOrgRpm() int32`

GetOrgRpm returns the OrgRpm field if non-nil, zero value otherwise.

### GetOrgRpmOk

`func (o *CloudPolicy) GetOrgRpmOk() (*int32, bool)`

GetOrgRpmOk returns a tuple with the OrgRpm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgRpm

`func (o *CloudPolicy) SetOrgRpm(v int32)`

SetOrgRpm sets OrgRpm field to given value.

### HasOrgRpm

`func (o *CloudPolicy) HasOrgRpm() bool`

HasOrgRpm returns a boolean if a field has been set.

### GetPerIpRpm

`func (o *CloudPolicy) GetPerIpRpm() int32`

GetPerIpRpm returns the PerIpRpm field if non-nil, zero value otherwise.

### GetPerIpRpmOk

`func (o *CloudPolicy) GetPerIpRpmOk() (*int32, bool)`

GetPerIpRpmOk returns a tuple with the PerIpRpm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerIpRpm

`func (o *CloudPolicy) SetPerIpRpm(v int32)`

SetPerIpRpm sets PerIpRpm field to given value.

### HasPerIpRpm

`func (o *CloudPolicy) HasPerIpRpm() bool`

HasPerIpRpm returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CloudPolicy) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CloudPolicy) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CloudPolicy) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CloudPolicy) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *CloudPolicy) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *CloudPolicy) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *CloudPolicy) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *CloudPolicy) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetWindowSec

`func (o *CloudPolicy) GetWindowSec() int32`

GetWindowSec returns the WindowSec field if non-nil, zero value otherwise.

### GetWindowSecOk

`func (o *CloudPolicy) GetWindowSecOk() (*int32, bool)`

GetWindowSecOk returns a tuple with the WindowSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowSec

`func (o *CloudPolicy) SetWindowSec(v int32)`

SetWindowSec sets WindowSec field to given value.

### HasWindowSec

`func (o *CloudPolicy) HasWindowSec() bool`

HasWindowSec returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


