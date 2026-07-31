# WorldWorldLimits200ResponseLimits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiRateLimit** | Pointer to **int32** | world.api_rate_limit, req/min; -1 &#x3D; unlimited | [optional] 
**McpRateLimit** | Pointer to **int32** | world.mcp_rate_limit, req/min; -1 &#x3D; unlimited | [optional] 
**MaxAlerts** | Pointer to **int32** | world.max_alerts; -1 &#x3D; unlimited | [optional] 
**ModelApi** | Pointer to **bool** | world.model_api — /v1/world/model + SSE stream access | [optional] 

## Methods

### NewWorldWorldLimits200ResponseLimits

`func NewWorldWorldLimits200ResponseLimits() *WorldWorldLimits200ResponseLimits`

NewWorldWorldLimits200ResponseLimits instantiates a new WorldWorldLimits200ResponseLimits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorldWorldLimits200ResponseLimitsWithDefaults

`func NewWorldWorldLimits200ResponseLimitsWithDefaults() *WorldWorldLimits200ResponseLimits`

NewWorldWorldLimits200ResponseLimitsWithDefaults instantiates a new WorldWorldLimits200ResponseLimits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiRateLimit

`func (o *WorldWorldLimits200ResponseLimits) GetApiRateLimit() int32`

GetApiRateLimit returns the ApiRateLimit field if non-nil, zero value otherwise.

### GetApiRateLimitOk

`func (o *WorldWorldLimits200ResponseLimits) GetApiRateLimitOk() (*int32, bool)`

GetApiRateLimitOk returns a tuple with the ApiRateLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiRateLimit

`func (o *WorldWorldLimits200ResponseLimits) SetApiRateLimit(v int32)`

SetApiRateLimit sets ApiRateLimit field to given value.

### HasApiRateLimit

`func (o *WorldWorldLimits200ResponseLimits) HasApiRateLimit() bool`

HasApiRateLimit returns a boolean if a field has been set.

### GetMcpRateLimit

`func (o *WorldWorldLimits200ResponseLimits) GetMcpRateLimit() int32`

GetMcpRateLimit returns the McpRateLimit field if non-nil, zero value otherwise.

### GetMcpRateLimitOk

`func (o *WorldWorldLimits200ResponseLimits) GetMcpRateLimitOk() (*int32, bool)`

GetMcpRateLimitOk returns a tuple with the McpRateLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcpRateLimit

`func (o *WorldWorldLimits200ResponseLimits) SetMcpRateLimit(v int32)`

SetMcpRateLimit sets McpRateLimit field to given value.

### HasMcpRateLimit

`func (o *WorldWorldLimits200ResponseLimits) HasMcpRateLimit() bool`

HasMcpRateLimit returns a boolean if a field has been set.

### GetMaxAlerts

`func (o *WorldWorldLimits200ResponseLimits) GetMaxAlerts() int32`

GetMaxAlerts returns the MaxAlerts field if non-nil, zero value otherwise.

### GetMaxAlertsOk

`func (o *WorldWorldLimits200ResponseLimits) GetMaxAlertsOk() (*int32, bool)`

GetMaxAlertsOk returns a tuple with the MaxAlerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAlerts

`func (o *WorldWorldLimits200ResponseLimits) SetMaxAlerts(v int32)`

SetMaxAlerts sets MaxAlerts field to given value.

### HasMaxAlerts

`func (o *WorldWorldLimits200ResponseLimits) HasMaxAlerts() bool`

HasMaxAlerts returns a boolean if a field has been set.

### GetModelApi

`func (o *WorldWorldLimits200ResponseLimits) GetModelApi() bool`

GetModelApi returns the ModelApi field if non-nil, zero value otherwise.

### GetModelApiOk

`func (o *WorldWorldLimits200ResponseLimits) GetModelApiOk() (*bool, bool)`

GetModelApiOk returns a tuple with the ModelApi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelApi

`func (o *WorldWorldLimits200ResponseLimits) SetModelApi(v bool)`

SetModelApi sets ModelApi field to given value.

### HasModelApi

`func (o *WorldWorldLimits200ResponseLimits) HasModelApi() bool`

HasModelApi returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


