# MetricBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Labels** | Pointer to **map[string]interface{}** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **float32** |  | [optional] 

## Methods

### NewMetricBody

`func NewMetricBody() *MetricBody`

NewMetricBody instantiates a new MetricBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricBodyWithDefaults

`func NewMetricBodyWithDefaults() *MetricBody`

NewMetricBodyWithDefaults instantiates a new MetricBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabels

`func (o *MetricBody) GetLabels() map[string]interface{}`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *MetricBody) GetLabelsOk() (*map[string]interface{}, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *MetricBody) SetLabels(v map[string]interface{})`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *MetricBody) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetName

`func (o *MetricBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MetricBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MetricBody) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MetricBody) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValue

`func (o *MetricBody) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *MetricBody) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *MetricBody) SetValue(v float32)`

SetValue sets Value field to given value.

### HasValue

`func (o *MetricBody) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


