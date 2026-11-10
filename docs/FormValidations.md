# FormValidations

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxLength** | Pointer to **int64** | MaxLength is the most characters a text answer holds, at most the platform&#39;s bound for the type. | [optional] 
**MaxValue** | Pointer to **float64** | MaxValue is the most a number answer may be. | [optional] 
**MinValue** | Pointer to **float64** | MinValue is the least a number answer may be. | [optional] 

## Methods

### NewFormValidations

`func NewFormValidations() *FormValidations`

NewFormValidations instantiates a new FormValidations object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormValidationsWithDefaults

`func NewFormValidationsWithDefaults() *FormValidations`

NewFormValidationsWithDefaults instantiates a new FormValidations object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxLength

`func (o *FormValidations) GetMaxLength() int64`

GetMaxLength returns the MaxLength field if non-nil, zero value otherwise.

### GetMaxLengthOk

`func (o *FormValidations) GetMaxLengthOk() (*int64, bool)`

GetMaxLengthOk returns a tuple with the MaxLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLength

`func (o *FormValidations) SetMaxLength(v int64)`

SetMaxLength sets MaxLength field to given value.

### HasMaxLength

`func (o *FormValidations) HasMaxLength() bool`

HasMaxLength returns a boolean if a field has been set.

### GetMaxValue

`func (o *FormValidations) GetMaxValue() float64`

GetMaxValue returns the MaxValue field if non-nil, zero value otherwise.

### GetMaxValueOk

`func (o *FormValidations) GetMaxValueOk() (*float64, bool)`

GetMaxValueOk returns a tuple with the MaxValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxValue

`func (o *FormValidations) SetMaxValue(v float64)`

SetMaxValue sets MaxValue field to given value.

### HasMaxValue

`func (o *FormValidations) HasMaxValue() bool`

HasMaxValue returns a boolean if a field has been set.

### GetMinValue

`func (o *FormValidations) GetMinValue() float64`

GetMinValue returns the MinValue field if non-nil, zero value otherwise.

### GetMinValueOk

`func (o *FormValidations) GetMinValueOk() (*float64, bool)`

GetMinValueOk returns a tuple with the MinValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinValue

`func (o *FormValidations) SetMinValue(v float64)`

SetMinValue sets MinValue field to given value.

### HasMinValue

`func (o *FormValidations) HasMinValue() bool`

HasMinValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


