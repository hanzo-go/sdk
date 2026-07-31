# ObserveMetricsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Product** | Pointer to **string** |  | [optional] 
**Range** | Pointer to [**O11yMetricsResponseRange**](O11yMetricsResponseRange.md) |  | [optional] 
**Series** | Pointer to [**ObserveMetricsResponseSeries**](ObserveMetricsResponseSeries.md) |  | [optional] 
**Usage** | Pointer to [**ObserveMetricsResponseUsage**](ObserveMetricsResponseUsage.md) |  | [optional] 
**Summary** | Pointer to [**O11yMetricsResponseSummary**](O11yMetricsResponseSummary.md) |  | [optional] 

## Methods

### NewObserveMetricsResponse

`func NewObserveMetricsResponse() *ObserveMetricsResponse`

NewObserveMetricsResponse instantiates a new ObserveMetricsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObserveMetricsResponseWithDefaults

`func NewObserveMetricsResponseWithDefaults() *ObserveMetricsResponse`

NewObserveMetricsResponseWithDefaults instantiates a new ObserveMetricsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProduct

`func (o *ObserveMetricsResponse) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *ObserveMetricsResponse) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *ObserveMetricsResponse) SetProduct(v string)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *ObserveMetricsResponse) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetRange

`func (o *ObserveMetricsResponse) GetRange() O11yMetricsResponseRange`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *ObserveMetricsResponse) GetRangeOk() (*O11yMetricsResponseRange, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *ObserveMetricsResponse) SetRange(v O11yMetricsResponseRange)`

SetRange sets Range field to given value.

### HasRange

`func (o *ObserveMetricsResponse) HasRange() bool`

HasRange returns a boolean if a field has been set.

### GetSeries

`func (o *ObserveMetricsResponse) GetSeries() ObserveMetricsResponseSeries`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *ObserveMetricsResponse) GetSeriesOk() (*ObserveMetricsResponseSeries, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *ObserveMetricsResponse) SetSeries(v ObserveMetricsResponseSeries)`

SetSeries sets Series field to given value.

### HasSeries

`func (o *ObserveMetricsResponse) HasSeries() bool`

HasSeries returns a boolean if a field has been set.

### GetUsage

`func (o *ObserveMetricsResponse) GetUsage() ObserveMetricsResponseUsage`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *ObserveMetricsResponse) GetUsageOk() (*ObserveMetricsResponseUsage, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *ObserveMetricsResponse) SetUsage(v ObserveMetricsResponseUsage)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *ObserveMetricsResponse) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### GetSummary

`func (o *ObserveMetricsResponse) GetSummary() O11yMetricsResponseSummary`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *ObserveMetricsResponse) GetSummaryOk() (*O11yMetricsResponseSummary, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *ObserveMetricsResponse) SetSummary(v O11yMetricsResponseSummary)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *ObserveMetricsResponse) HasSummary() bool`

HasSummary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


