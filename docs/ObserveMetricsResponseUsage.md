# ObserveMetricsResponseUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Calls** | Pointer to **int64** |  | [optional] 
**Tokens** | Pointer to **int64** |  | [optional] 
**CostCents** | Pointer to **int64** |  | [optional] 
**Series** | Pointer to [**[]ObserveUsagePoint**](ObserveUsagePoint.md) |  | [optional] 

## Methods

### NewObserveMetricsResponseUsage

`func NewObserveMetricsResponseUsage() *ObserveMetricsResponseUsage`

NewObserveMetricsResponseUsage instantiates a new ObserveMetricsResponseUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObserveMetricsResponseUsageWithDefaults

`func NewObserveMetricsResponseUsageWithDefaults() *ObserveMetricsResponseUsage`

NewObserveMetricsResponseUsageWithDefaults instantiates a new ObserveMetricsResponseUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCalls

`func (o *ObserveMetricsResponseUsage) GetCalls() int64`

GetCalls returns the Calls field if non-nil, zero value otherwise.

### GetCallsOk

`func (o *ObserveMetricsResponseUsage) GetCallsOk() (*int64, bool)`

GetCallsOk returns a tuple with the Calls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalls

`func (o *ObserveMetricsResponseUsage) SetCalls(v int64)`

SetCalls sets Calls field to given value.

### HasCalls

`func (o *ObserveMetricsResponseUsage) HasCalls() bool`

HasCalls returns a boolean if a field has been set.

### GetTokens

`func (o *ObserveMetricsResponseUsage) GetTokens() int64`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *ObserveMetricsResponseUsage) GetTokensOk() (*int64, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *ObserveMetricsResponseUsage) SetTokens(v int64)`

SetTokens sets Tokens field to given value.

### HasTokens

`func (o *ObserveMetricsResponseUsage) HasTokens() bool`

HasTokens returns a boolean if a field has been set.

### GetCostCents

`func (o *ObserveMetricsResponseUsage) GetCostCents() int64`

GetCostCents returns the CostCents field if non-nil, zero value otherwise.

### GetCostCentsOk

`func (o *ObserveMetricsResponseUsage) GetCostCentsOk() (*int64, bool)`

GetCostCentsOk returns a tuple with the CostCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostCents

`func (o *ObserveMetricsResponseUsage) SetCostCents(v int64)`

SetCostCents sets CostCents field to given value.

### HasCostCents

`func (o *ObserveMetricsResponseUsage) HasCostCents() bool`

HasCostCents returns a boolean if a field has been set.

### GetSeries

`func (o *ObserveMetricsResponseUsage) GetSeries() []ObserveUsagePoint`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *ObserveMetricsResponseUsage) GetSeriesOk() (*[]ObserveUsagePoint, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *ObserveMetricsResponseUsage) SetSeries(v []ObserveUsagePoint)`

SetSeries sets Series field to given value.

### HasSeries

`func (o *ObserveMetricsResponseUsage) HasSeries() bool`

HasSeries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


