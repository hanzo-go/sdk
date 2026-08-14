# O11yO11yMetricLabel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to [**O11yO11yMetricField**](O11yO11yMetricField.md) | Key is the label&#39;s field. | [optional] 
**Value** | Pointer to **map[string]interface{}** | Value is the label&#39;s value. | [optional] 

## Methods

### NewO11yO11yMetricLabel

`func NewO11yO11yMetricLabel() *O11yO11yMetricLabel`

NewO11yO11yMetricLabel instantiates a new O11yO11yMetricLabel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yMetricLabelWithDefaults

`func NewO11yO11yMetricLabelWithDefaults() *O11yO11yMetricLabel`

NewO11yO11yMetricLabelWithDefaults instantiates a new O11yO11yMetricLabel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *O11yO11yMetricLabel) GetKey() O11yO11yMetricField`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *O11yO11yMetricLabel) GetKeyOk() (*O11yO11yMetricField, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *O11yO11yMetricLabel) SetKey(v O11yO11yMetricField)`

SetKey sets Key field to given value.

### HasKey

`func (o *O11yO11yMetricLabel) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetValue

`func (o *O11yO11yMetricLabel) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *O11yO11yMetricLabel) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *O11yO11yMetricLabel) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *O11yO11yMetricLabel) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


