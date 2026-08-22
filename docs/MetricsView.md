# MetricsView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Range** | Pointer to **string** | echoes the requested window (24H|7D|30D) | [optional] 
**Resource** | Pointer to [**ResourceUsage**](ResourceUsage.md) | Resource is the Resource Usage panel&#39;s rollup, and every field of it is currently null — see resourceUsage. It is present rather than omitted so a panel renders \&quot;—\&quot; instead of guessing. | [optional] 
**Series** | Pointer to [**[]SeriesLine**](SeriesLine.md) | per-agent invocation histogram (real) | [optional] 

## Methods

### NewMetricsView

`func NewMetricsView() *MetricsView`

NewMetricsView instantiates a new MetricsView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricsViewWithDefaults

`func NewMetricsViewWithDefaults() *MetricsView`

NewMetricsViewWithDefaults instantiates a new MetricsView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRange

`func (o *MetricsView) GetRange() string`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *MetricsView) GetRangeOk() (*string, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *MetricsView) SetRange(v string)`

SetRange sets Range field to given value.

### HasRange

`func (o *MetricsView) HasRange() bool`

HasRange returns a boolean if a field has been set.

### GetResource

`func (o *MetricsView) GetResource() ResourceUsage`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *MetricsView) GetResourceOk() (*ResourceUsage, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *MetricsView) SetResource(v ResourceUsage)`

SetResource sets Resource field to given value.

### HasResource

`func (o *MetricsView) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetSeries

`func (o *MetricsView) GetSeries() []SeriesLine`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *MetricsView) GetSeriesOk() (*[]SeriesLine, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *MetricsView) SetSeries(v []SeriesLine)`

SetSeries sets Series field to given value.

### HasSeries

`func (o *MetricsView) HasSeries() bool`

HasSeries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


