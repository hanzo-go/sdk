# CloudMetricsResponseUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Calls** | Pointer to **int32** |  | [optional] 
**CostCents** | Pointer to **int32** |  | [optional] 
**Series** | Pointer to [**[]CloudUsageBucket**](CloudUsageBucket.md) |  | [optional] 
**Tokens** | Pointer to **int32** |  | [optional] 

## Methods

### NewCloudMetricsResponseUsage

`func NewCloudMetricsResponseUsage() *CloudMetricsResponseUsage`

NewCloudMetricsResponseUsage instantiates a new CloudMetricsResponseUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudMetricsResponseUsageWithDefaults

`func NewCloudMetricsResponseUsageWithDefaults() *CloudMetricsResponseUsage`

NewCloudMetricsResponseUsageWithDefaults instantiates a new CloudMetricsResponseUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCalls

`func (o *CloudMetricsResponseUsage) GetCalls() int32`

GetCalls returns the Calls field if non-nil, zero value otherwise.

### GetCallsOk

`func (o *CloudMetricsResponseUsage) GetCallsOk() (*int32, bool)`

GetCallsOk returns a tuple with the Calls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalls

`func (o *CloudMetricsResponseUsage) SetCalls(v int32)`

SetCalls sets Calls field to given value.

### HasCalls

`func (o *CloudMetricsResponseUsage) HasCalls() bool`

HasCalls returns a boolean if a field has been set.

### GetCostCents

`func (o *CloudMetricsResponseUsage) GetCostCents() int32`

GetCostCents returns the CostCents field if non-nil, zero value otherwise.

### GetCostCentsOk

`func (o *CloudMetricsResponseUsage) GetCostCentsOk() (*int32, bool)`

GetCostCentsOk returns a tuple with the CostCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostCents

`func (o *CloudMetricsResponseUsage) SetCostCents(v int32)`

SetCostCents sets CostCents field to given value.

### HasCostCents

`func (o *CloudMetricsResponseUsage) HasCostCents() bool`

HasCostCents returns a boolean if a field has been set.

### GetSeries

`func (o *CloudMetricsResponseUsage) GetSeries() []CloudUsageBucket`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *CloudMetricsResponseUsage) GetSeriesOk() (*[]CloudUsageBucket, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *CloudMetricsResponseUsage) SetSeries(v []CloudUsageBucket)`

SetSeries sets Series field to given value.

### HasSeries

`func (o *CloudMetricsResponseUsage) HasSeries() bool`

HasSeries returns a boolean if a field has been set.

### GetTokens

`func (o *CloudMetricsResponseUsage) GetTokens() int32`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *CloudMetricsResponseUsage) GetTokensOk() (*int32, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *CloudMetricsResponseUsage) SetTokens(v int32)`

SetTokens sets Tokens field to given value.

### HasTokens

`func (o *CloudMetricsResponseUsage) HasTokens() bool`

HasTokens returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


