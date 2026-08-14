# O11yO11yMetricTreemap

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Samples** | Pointer to [**[]O11yO11yTreemapEntry**](O11yO11yTreemapEntry.md) | Samples are the entries when measuring by sample count. | [optional] 
**Timeseries** | Pointer to [**[]O11yO11yTreemapEntry**](O11yO11yTreemapEntry.md) | TimeSeries are the entries when measuring by time-series count. | [optional] 

## Methods

### NewO11yO11yMetricTreemap

`func NewO11yO11yMetricTreemap() *O11yO11yMetricTreemap`

NewO11yO11yMetricTreemap instantiates a new O11yO11yMetricTreemap object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yMetricTreemapWithDefaults

`func NewO11yO11yMetricTreemapWithDefaults() *O11yO11yMetricTreemap`

NewO11yO11yMetricTreemapWithDefaults instantiates a new O11yO11yMetricTreemap object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSamples

`func (o *O11yO11yMetricTreemap) GetSamples() []O11yO11yTreemapEntry`

GetSamples returns the Samples field if non-nil, zero value otherwise.

### GetSamplesOk

`func (o *O11yO11yMetricTreemap) GetSamplesOk() (*[]O11yO11yTreemapEntry, bool)`

GetSamplesOk returns a tuple with the Samples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamples

`func (o *O11yO11yMetricTreemap) SetSamples(v []O11yO11yTreemapEntry)`

SetSamples sets Samples field to given value.

### HasSamples

`func (o *O11yO11yMetricTreemap) HasSamples() bool`

HasSamples returns a boolean if a field has been set.

### GetTimeseries

`func (o *O11yO11yMetricTreemap) GetTimeseries() []O11yO11yTreemapEntry`

GetTimeseries returns the Timeseries field if non-nil, zero value otherwise.

### GetTimeseriesOk

`func (o *O11yO11yMetricTreemap) GetTimeseriesOk() (*[]O11yO11yTreemapEntry, bool)`

GetTimeseriesOk returns a tuple with the Timeseries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeseries

`func (o *O11yO11yMetricTreemap) SetTimeseries(v []O11yO11yTreemapEntry)`

SetTimeseries sets Timeseries field to given value.

### HasTimeseries

`func (o *O11yO11yMetricTreemap) HasTimeseries() bool`

HasTimeseries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


