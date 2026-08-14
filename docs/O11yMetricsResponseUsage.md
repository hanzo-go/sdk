# O11yMetricsResponseUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Calls** | Pointer to **int32** |  | [optional] 
**CostCents** | Pointer to **int32** |  | [optional] 
**Series** | Pointer to [**[]O11yUsageBucket**](O11yUsageBucket.md) |  | [optional] 
**Tokens** | Pointer to **int32** |  | [optional] 

## Methods

### NewO11yMetricsResponseUsage

`func NewO11yMetricsResponseUsage() *O11yMetricsResponseUsage`

NewO11yMetricsResponseUsage instantiates a new O11yMetricsResponseUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yMetricsResponseUsageWithDefaults

`func NewO11yMetricsResponseUsageWithDefaults() *O11yMetricsResponseUsage`

NewO11yMetricsResponseUsageWithDefaults instantiates a new O11yMetricsResponseUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCalls

`func (o *O11yMetricsResponseUsage) GetCalls() int32`

GetCalls returns the Calls field if non-nil, zero value otherwise.

### GetCallsOk

`func (o *O11yMetricsResponseUsage) GetCallsOk() (*int32, bool)`

GetCallsOk returns a tuple with the Calls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalls

`func (o *O11yMetricsResponseUsage) SetCalls(v int32)`

SetCalls sets Calls field to given value.

### HasCalls

`func (o *O11yMetricsResponseUsage) HasCalls() bool`

HasCalls returns a boolean if a field has been set.

### GetCostCents

`func (o *O11yMetricsResponseUsage) GetCostCents() int32`

GetCostCents returns the CostCents field if non-nil, zero value otherwise.

### GetCostCentsOk

`func (o *O11yMetricsResponseUsage) GetCostCentsOk() (*int32, bool)`

GetCostCentsOk returns a tuple with the CostCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostCents

`func (o *O11yMetricsResponseUsage) SetCostCents(v int32)`

SetCostCents sets CostCents field to given value.

### HasCostCents

`func (o *O11yMetricsResponseUsage) HasCostCents() bool`

HasCostCents returns a boolean if a field has been set.

### GetSeries

`func (o *O11yMetricsResponseUsage) GetSeries() []O11yUsageBucket`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *O11yMetricsResponseUsage) GetSeriesOk() (*[]O11yUsageBucket, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *O11yMetricsResponseUsage) SetSeries(v []O11yUsageBucket)`

SetSeries sets Series field to given value.

### HasSeries

`func (o *O11yMetricsResponseUsage) HasSeries() bool`

HasSeries returns a boolean if a field has been set.

### GetTokens

`func (o *O11yMetricsResponseUsage) GetTokens() int32`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *O11yMetricsResponseUsage) GetTokensOk() (*int32, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *O11yMetricsResponseUsage) SetTokens(v int32)`

SetTokens sets Tokens field to given value.

### HasTokens

`func (o *O11yMetricsResponseUsage) HasTokens() bool`

HasTokens returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


