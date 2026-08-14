# Usage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CostCents** | Pointer to **int32** | null — no per-invocation cost source | [optional] 
**Series** | Pointer to [**[]CostLine**](CostLine.md) | one line per function that ran in the window | [optional] 
**Status** | Pointer to [**StatusBreakdown**](StatusBreakdown.md) | how those invocations ended | [optional] 

## Methods

### NewUsage

`func NewUsage() *Usage`

NewUsage instantiates a new Usage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageWithDefaults

`func NewUsageWithDefaults() *Usage`

NewUsageWithDefaults instantiates a new Usage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCostCents

`func (o *Usage) GetCostCents() int32`

GetCostCents returns the CostCents field if non-nil, zero value otherwise.

### GetCostCentsOk

`func (o *Usage) GetCostCentsOk() (*int32, bool)`

GetCostCentsOk returns a tuple with the CostCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostCents

`func (o *Usage) SetCostCents(v int32)`

SetCostCents sets CostCents field to given value.

### HasCostCents

`func (o *Usage) HasCostCents() bool`

HasCostCents returns a boolean if a field has been set.

### GetSeries

`func (o *Usage) GetSeries() []CostLine`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *Usage) GetSeriesOk() (*[]CostLine, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *Usage) SetSeries(v []CostLine)`

SetSeries sets Series field to given value.

### HasSeries

`func (o *Usage) HasSeries() bool`

HasSeries returns a boolean if a field has been set.

### GetStatus

`func (o *Usage) GetStatus() StatusBreakdown`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Usage) GetStatusOk() (*StatusBreakdown, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Usage) SetStatus(v StatusBreakdown)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Usage) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


