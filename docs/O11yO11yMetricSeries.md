# O11yO11yMetricSeries

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Labels** | Pointer to [**[]O11yO11yMetricLabel**](O11yO11yMetricLabel.md) | Labels identify the series. | [optional] 
**Values** | Pointer to [**[]O11yO11yMetricPoint**](O11yO11yMetricPoint.md) | Values are the series&#39; points, in time order. | [optional] 

## Methods

### NewO11yO11yMetricSeries

`func NewO11yO11yMetricSeries() *O11yO11yMetricSeries`

NewO11yO11yMetricSeries instantiates a new O11yO11yMetricSeries object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yMetricSeriesWithDefaults

`func NewO11yO11yMetricSeriesWithDefaults() *O11yO11yMetricSeries`

NewO11yO11yMetricSeriesWithDefaults instantiates a new O11yO11yMetricSeries object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabels

`func (o *O11yO11yMetricSeries) GetLabels() []O11yO11yMetricLabel`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *O11yO11yMetricSeries) GetLabelsOk() (*[]O11yO11yMetricLabel, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *O11yO11yMetricSeries) SetLabels(v []O11yO11yMetricLabel)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *O11yO11yMetricSeries) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetValues

`func (o *O11yO11yMetricSeries) GetValues() []O11yO11yMetricPoint`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *O11yO11yMetricSeries) GetValuesOk() (*[]O11yO11yMetricPoint, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *O11yO11yMetricSeries) SetValues(v []O11yO11yMetricPoint)`

SetValues sets Values field to given value.

### HasValues

`func (o *O11yO11yMetricSeries) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


