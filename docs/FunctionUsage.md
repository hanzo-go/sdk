# FunctionUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CostCents** | Pointer to **int64** | null — no per-invocation cost source | [optional] 
**Series** | Pointer to [**[]FunctionCostLine**](FunctionCostLine.md) | one line per function that ran in the window | [optional] 
**Status** | Pointer to [**FunctionStatusBreakdown**](FunctionStatusBreakdown.md) | how those invocations ended | [optional] 

## Methods

### NewFunctionUsage

`func NewFunctionUsage() *FunctionUsage`

NewFunctionUsage instantiates a new FunctionUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionUsageWithDefaults

`func NewFunctionUsageWithDefaults() *FunctionUsage`

NewFunctionUsageWithDefaults instantiates a new FunctionUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCostCents

`func (o *FunctionUsage) GetCostCents() int64`

GetCostCents returns the CostCents field if non-nil, zero value otherwise.

### GetCostCentsOk

`func (o *FunctionUsage) GetCostCentsOk() (*int64, bool)`

GetCostCentsOk returns a tuple with the CostCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostCents

`func (o *FunctionUsage) SetCostCents(v int64)`

SetCostCents sets CostCents field to given value.

### HasCostCents

`func (o *FunctionUsage) HasCostCents() bool`

HasCostCents returns a boolean if a field has been set.

### GetSeries

`func (o *FunctionUsage) GetSeries() []FunctionCostLine`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *FunctionUsage) GetSeriesOk() (*[]FunctionCostLine, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *FunctionUsage) SetSeries(v []FunctionCostLine)`

SetSeries sets Series field to given value.

### HasSeries

`func (o *FunctionUsage) HasSeries() bool`

HasSeries returns a boolean if a field has been set.

### GetStatus

`func (o *FunctionUsage) GetStatus() FunctionStatusBreakdown`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FunctionUsage) GetStatusOk() (*FunctionStatusBreakdown, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FunctionUsage) SetStatus(v FunctionStatusBreakdown)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FunctionUsage) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


