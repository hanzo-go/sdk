# O11yMetricsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Product** | Pointer to **string** | Product is the service these numbers are about, echoed back. | [optional] 
**Range** | Pointer to [**O11yMetricsResponseRange**](O11yMetricsResponseRange.md) |  | [optional] 
**Series** | Pointer to [**O11yMetricsResponseSeries**](O11yMetricsResponseSeries.md) |  | [optional] 
**Summary** | Pointer to [**O11yMetricsResponseSummary**](O11yMetricsResponseSummary.md) |  | [optional] 
**Usage** | Pointer to [**O11yMetricsResponseUsage**](O11yMetricsResponseUsage.md) |  | [optional] 

## Methods

### NewO11yMetricsResponse

`func NewO11yMetricsResponse() *O11yMetricsResponse`

NewO11yMetricsResponse instantiates a new O11yMetricsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yMetricsResponseWithDefaults

`func NewO11yMetricsResponseWithDefaults() *O11yMetricsResponse`

NewO11yMetricsResponseWithDefaults instantiates a new O11yMetricsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProduct

`func (o *O11yMetricsResponse) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *O11yMetricsResponse) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *O11yMetricsResponse) SetProduct(v string)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *O11yMetricsResponse) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetRange

`func (o *O11yMetricsResponse) GetRange() O11yMetricsResponseRange`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *O11yMetricsResponse) GetRangeOk() (*O11yMetricsResponseRange, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *O11yMetricsResponse) SetRange(v O11yMetricsResponseRange)`

SetRange sets Range field to given value.

### HasRange

`func (o *O11yMetricsResponse) HasRange() bool`

HasRange returns a boolean if a field has been set.

### GetSeries

`func (o *O11yMetricsResponse) GetSeries() O11yMetricsResponseSeries`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *O11yMetricsResponse) GetSeriesOk() (*O11yMetricsResponseSeries, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *O11yMetricsResponse) SetSeries(v O11yMetricsResponseSeries)`

SetSeries sets Series field to given value.

### HasSeries

`func (o *O11yMetricsResponse) HasSeries() bool`

HasSeries returns a boolean if a field has been set.

### GetSummary

`func (o *O11yMetricsResponse) GetSummary() O11yMetricsResponseSummary`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *O11yMetricsResponse) GetSummaryOk() (*O11yMetricsResponseSummary, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *O11yMetricsResponse) SetSummary(v O11yMetricsResponseSummary)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *O11yMetricsResponse) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetUsage

`func (o *O11yMetricsResponse) GetUsage() O11yMetricsResponseUsage`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *O11yMetricsResponse) GetUsageOk() (*O11yMetricsResponseUsage, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *O11yMetricsResponse) SetUsage(v O11yMetricsResponseUsage)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *O11yMetricsResponse) HasUsage() bool`

HasUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


