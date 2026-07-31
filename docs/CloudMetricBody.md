# CloudMetricBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Labels** | Pointer to **map[string]interface{}** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **float32** |  | [optional] 

## Methods

### NewCloudMetricBody

`func NewCloudMetricBody() *CloudMetricBody`

NewCloudMetricBody instantiates a new CloudMetricBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudMetricBodyWithDefaults

`func NewCloudMetricBodyWithDefaults() *CloudMetricBody`

NewCloudMetricBodyWithDefaults instantiates a new CloudMetricBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabels

`func (o *CloudMetricBody) GetLabels() map[string]interface{}`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *CloudMetricBody) GetLabelsOk() (*map[string]interface{}, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *CloudMetricBody) SetLabels(v map[string]interface{})`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *CloudMetricBody) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetName

`func (o *CloudMetricBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudMetricBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudMetricBody) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudMetricBody) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValue

`func (o *CloudMetricBody) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CloudMetricBody) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CloudMetricBody) SetValue(v float32)`

SetValue sets Value field to given value.

### HasValue

`func (o *CloudMetricBody) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


