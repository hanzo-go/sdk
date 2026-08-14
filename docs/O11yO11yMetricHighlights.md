# O11yO11yMetricHighlights

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveTimeSeries** | Pointer to **int32** | ActiveTimeSeries is how many of them are active. | [optional] 
**DataPoints** | Pointer to **int32** | DataPoints is how many data points the metric has. | [optional] 
**LastReceived** | Pointer to **int32** | LastReceived is when the metric last arrived, as a Unix timestamp in milliseconds. | [optional] 
**TotalTimeSeries** | Pointer to **int32** | TotalTimeSeries is how many time series the metric has ever had. | [optional] 

## Methods

### NewO11yO11yMetricHighlights

`func NewO11yO11yMetricHighlights() *O11yO11yMetricHighlights`

NewO11yO11yMetricHighlights instantiates a new O11yO11yMetricHighlights object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yMetricHighlightsWithDefaults

`func NewO11yO11yMetricHighlightsWithDefaults() *O11yO11yMetricHighlights`

NewO11yO11yMetricHighlightsWithDefaults instantiates a new O11yO11yMetricHighlights object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveTimeSeries

`func (o *O11yO11yMetricHighlights) GetActiveTimeSeries() int32`

GetActiveTimeSeries returns the ActiveTimeSeries field if non-nil, zero value otherwise.

### GetActiveTimeSeriesOk

`func (o *O11yO11yMetricHighlights) GetActiveTimeSeriesOk() (*int32, bool)`

GetActiveTimeSeriesOk returns a tuple with the ActiveTimeSeries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveTimeSeries

`func (o *O11yO11yMetricHighlights) SetActiveTimeSeries(v int32)`

SetActiveTimeSeries sets ActiveTimeSeries field to given value.

### HasActiveTimeSeries

`func (o *O11yO11yMetricHighlights) HasActiveTimeSeries() bool`

HasActiveTimeSeries returns a boolean if a field has been set.

### GetDataPoints

`func (o *O11yO11yMetricHighlights) GetDataPoints() int32`

GetDataPoints returns the DataPoints field if non-nil, zero value otherwise.

### GetDataPointsOk

`func (o *O11yO11yMetricHighlights) GetDataPointsOk() (*int32, bool)`

GetDataPointsOk returns a tuple with the DataPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataPoints

`func (o *O11yO11yMetricHighlights) SetDataPoints(v int32)`

SetDataPoints sets DataPoints field to given value.

### HasDataPoints

`func (o *O11yO11yMetricHighlights) HasDataPoints() bool`

HasDataPoints returns a boolean if a field has been set.

### GetLastReceived

`func (o *O11yO11yMetricHighlights) GetLastReceived() int32`

GetLastReceived returns the LastReceived field if non-nil, zero value otherwise.

### GetLastReceivedOk

`func (o *O11yO11yMetricHighlights) GetLastReceivedOk() (*int32, bool)`

GetLastReceivedOk returns a tuple with the LastReceived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReceived

`func (o *O11yO11yMetricHighlights) SetLastReceived(v int32)`

SetLastReceived sets LastReceived field to given value.

### HasLastReceived

`func (o *O11yO11yMetricHighlights) HasLastReceived() bool`

HasLastReceived returns a boolean if a field has been set.

### GetTotalTimeSeries

`func (o *O11yO11yMetricHighlights) GetTotalTimeSeries() int32`

GetTotalTimeSeries returns the TotalTimeSeries field if non-nil, zero value otherwise.

### GetTotalTimeSeriesOk

`func (o *O11yO11yMetricHighlights) GetTotalTimeSeriesOk() (*int32, bool)`

GetTotalTimeSeriesOk returns a tuple with the TotalTimeSeries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTimeSeries

`func (o *O11yO11yMetricHighlights) SetTotalTimeSeries(v int32)`

SetTotalTimeSeries sets TotalTimeSeries field to given value.

### HasTotalTimeSeries

`func (o *O11yO11yMetricHighlights) HasTotalTimeSeries() bool`

HasTotalTimeSeries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


