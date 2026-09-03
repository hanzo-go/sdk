# LimitsBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiRateLimit** | Pointer to **int64** | APIRateLimit is requests per minute allowed against the REST /v1/world surface. -1 means unlimited. | [optional] 
**MaxAlerts** | Pointer to **int64** | MaxAlerts is how many saved OSINT alert rules the plan allows. -1 means unlimited. | [optional] 
**McpRateLimit** | Pointer to **int64** | MCPRateLimit is requests per minute allowed against the MCP surface. -1 means unlimited. | [optional] 
**ModelApi** | Pointer to **bool** | ModelAPI is whether the plan reaches the World model endpoint and the SSE stream. The free floor is false, and that is what a catalog outage resolves to. | [optional] 

## Methods

### NewLimitsBlock

`func NewLimitsBlock() *LimitsBlock`

NewLimitsBlock instantiates a new LimitsBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLimitsBlockWithDefaults

`func NewLimitsBlockWithDefaults() *LimitsBlock`

NewLimitsBlockWithDefaults instantiates a new LimitsBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiRateLimit

`func (o *LimitsBlock) GetApiRateLimit() int64`

GetApiRateLimit returns the ApiRateLimit field if non-nil, zero value otherwise.

### GetApiRateLimitOk

`func (o *LimitsBlock) GetApiRateLimitOk() (*int64, bool)`

GetApiRateLimitOk returns a tuple with the ApiRateLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiRateLimit

`func (o *LimitsBlock) SetApiRateLimit(v int64)`

SetApiRateLimit sets ApiRateLimit field to given value.

### HasApiRateLimit

`func (o *LimitsBlock) HasApiRateLimit() bool`

HasApiRateLimit returns a boolean if a field has been set.

### GetMaxAlerts

`func (o *LimitsBlock) GetMaxAlerts() int64`

GetMaxAlerts returns the MaxAlerts field if non-nil, zero value otherwise.

### GetMaxAlertsOk

`func (o *LimitsBlock) GetMaxAlertsOk() (*int64, bool)`

GetMaxAlertsOk returns a tuple with the MaxAlerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAlerts

`func (o *LimitsBlock) SetMaxAlerts(v int64)`

SetMaxAlerts sets MaxAlerts field to given value.

### HasMaxAlerts

`func (o *LimitsBlock) HasMaxAlerts() bool`

HasMaxAlerts returns a boolean if a field has been set.

### GetMcpRateLimit

`func (o *LimitsBlock) GetMcpRateLimit() int64`

GetMcpRateLimit returns the McpRateLimit field if non-nil, zero value otherwise.

### GetMcpRateLimitOk

`func (o *LimitsBlock) GetMcpRateLimitOk() (*int64, bool)`

GetMcpRateLimitOk returns a tuple with the McpRateLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcpRateLimit

`func (o *LimitsBlock) SetMcpRateLimit(v int64)`

SetMcpRateLimit sets McpRateLimit field to given value.

### HasMcpRateLimit

`func (o *LimitsBlock) HasMcpRateLimit() bool`

HasMcpRateLimit returns a boolean if a field has been set.

### GetModelApi

`func (o *LimitsBlock) GetModelApi() bool`

GetModelApi returns the ModelApi field if non-nil, zero value otherwise.

### GetModelApiOk

`func (o *LimitsBlock) GetModelApiOk() (*bool, bool)`

GetModelApiOk returns a tuple with the ModelApi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelApi

`func (o *LimitsBlock) SetModelApi(v bool)`

SetModelApi sets ModelApi field to given value.

### HasModelApi

`func (o *LimitsBlock) HasModelApi() bool`

HasModelApi returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


