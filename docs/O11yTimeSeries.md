# O11yTimeSeries

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Labels** | Pointer to [**[]O11yLabel**](O11yLabel.md) |  | [optional] 
**Values** | Pointer to **[]interface{}** |  | [optional] 

## Methods

### NewO11yTimeSeries

`func NewO11yTimeSeries() *O11yTimeSeries`

NewO11yTimeSeries instantiates a new O11yTimeSeries object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yTimeSeriesWithDefaults

`func NewO11yTimeSeriesWithDefaults() *O11yTimeSeries`

NewO11yTimeSeriesWithDefaults instantiates a new O11yTimeSeries object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabels

`func (o *O11yTimeSeries) GetLabels() []O11yLabel`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *O11yTimeSeries) GetLabelsOk() (*[]O11yLabel, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *O11yTimeSeries) SetLabels(v []O11yLabel)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *O11yTimeSeries) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetValues

`func (o *O11yTimeSeries) GetValues() []interface{}`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *O11yTimeSeries) GetValuesOk() (*[]interface{}, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *O11yTimeSeries) SetValues(v []interface{})`

SetValues sets Values field to given value.

### HasValues

`func (o *O11yTimeSeries) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


