# CloudMetricsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Product** | Pointer to **string** |  | [optional] 
**Range** | Pointer to [**CloudMetricsResponseRange**](CloudMetricsResponseRange.md) |  | [optional] 
**Series** | Pointer to [**CloudMetricsResponseSeries**](CloudMetricsResponseSeries.md) |  | [optional] 
**Summary** | Pointer to [**CloudMetricsResponseSummary**](CloudMetricsResponseSummary.md) |  | [optional] 
**Usage** | Pointer to [**CloudMetricsResponseUsage**](CloudMetricsResponseUsage.md) |  | [optional] 

## Methods

### NewCloudMetricsResponse

`func NewCloudMetricsResponse() *CloudMetricsResponse`

NewCloudMetricsResponse instantiates a new CloudMetricsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudMetricsResponseWithDefaults

`func NewCloudMetricsResponseWithDefaults() *CloudMetricsResponse`

NewCloudMetricsResponseWithDefaults instantiates a new CloudMetricsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProduct

`func (o *CloudMetricsResponse) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *CloudMetricsResponse) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *CloudMetricsResponse) SetProduct(v string)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *CloudMetricsResponse) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetRange

`func (o *CloudMetricsResponse) GetRange() CloudMetricsResponseRange`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *CloudMetricsResponse) GetRangeOk() (*CloudMetricsResponseRange, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *CloudMetricsResponse) SetRange(v CloudMetricsResponseRange)`

SetRange sets Range field to given value.

### HasRange

`func (o *CloudMetricsResponse) HasRange() bool`

HasRange returns a boolean if a field has been set.

### GetSeries

`func (o *CloudMetricsResponse) GetSeries() CloudMetricsResponseSeries`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *CloudMetricsResponse) GetSeriesOk() (*CloudMetricsResponseSeries, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *CloudMetricsResponse) SetSeries(v CloudMetricsResponseSeries)`

SetSeries sets Series field to given value.

### HasSeries

`func (o *CloudMetricsResponse) HasSeries() bool`

HasSeries returns a boolean if a field has been set.

### GetSummary

`func (o *CloudMetricsResponse) GetSummary() CloudMetricsResponseSummary`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *CloudMetricsResponse) GetSummaryOk() (*CloudMetricsResponseSummary, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *CloudMetricsResponse) SetSummary(v CloudMetricsResponseSummary)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *CloudMetricsResponse) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetUsage

`func (o *CloudMetricsResponse) GetUsage() CloudMetricsResponseUsage`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *CloudMetricsResponse) GetUsageOk() (*CloudMetricsResponseUsage, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *CloudMetricsResponse) SetUsage(v CloudMetricsResponseUsage)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *CloudMetricsResponse) HasUsage() bool`

HasUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


