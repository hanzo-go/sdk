# FunctionsMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Series** | Pointer to [**[]FunctionsMetricsSeriesInner**](FunctionsMetricsSeriesInner.md) |  | [optional] 
**Status** | Pointer to [**FunctionsMetricsStatus**](FunctionsMetricsStatus.md) |  | [optional] 
**CostCents** | Pointer to **int64** |  | [optional] 

## Methods

### NewFunctionsMetrics

`func NewFunctionsMetrics() *FunctionsMetrics`

NewFunctionsMetrics instantiates a new FunctionsMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionsMetricsWithDefaults

`func NewFunctionsMetricsWithDefaults() *FunctionsMetrics`

NewFunctionsMetricsWithDefaults instantiates a new FunctionsMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSeries

`func (o *FunctionsMetrics) GetSeries() []FunctionsMetricsSeriesInner`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *FunctionsMetrics) GetSeriesOk() (*[]FunctionsMetricsSeriesInner, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *FunctionsMetrics) SetSeries(v []FunctionsMetricsSeriesInner)`

SetSeries sets Series field to given value.

### HasSeries

`func (o *FunctionsMetrics) HasSeries() bool`

HasSeries returns a boolean if a field has been set.

### GetStatus

`func (o *FunctionsMetrics) GetStatus() FunctionsMetricsStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FunctionsMetrics) GetStatusOk() (*FunctionsMetricsStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FunctionsMetrics) SetStatus(v FunctionsMetricsStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FunctionsMetrics) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCostCents

`func (o *FunctionsMetrics) GetCostCents() int64`

GetCostCents returns the CostCents field if non-nil, zero value otherwise.

### GetCostCentsOk

`func (o *FunctionsMetrics) GetCostCentsOk() (*int64, bool)`

GetCostCentsOk returns a tuple with the CostCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostCents

`func (o *FunctionsMetrics) SetCostCents(v int64)`

SetCostCents sets CostCents field to given value.

### HasCostCents

`func (o *FunctionsMetrics) HasCostCents() bool`

HasCostCents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


