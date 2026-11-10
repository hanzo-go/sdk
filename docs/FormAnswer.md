# FormAnswer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Boolean** | Pointer to **bool** |  | [optional] 
**Choices** | Pointer to **[]string** |  | [optional] 
**Date** | Pointer to **string** |  | [optional] 
**Field** | Pointer to **string** | Field is the ref of the field answered. | [optional] 
**Number** | Pointer to **float64** |  | [optional] 
**Text** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** | Type is the field&#39;s type. Set on a read; ignored on a write. | [optional] 

## Methods

### NewFormAnswer

`func NewFormAnswer() *FormAnswer`

NewFormAnswer instantiates a new FormAnswer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormAnswerWithDefaults

`func NewFormAnswerWithDefaults() *FormAnswer`

NewFormAnswerWithDefaults instantiates a new FormAnswer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBoolean

`func (o *FormAnswer) GetBoolean() bool`

GetBoolean returns the Boolean field if non-nil, zero value otherwise.

### GetBooleanOk

`func (o *FormAnswer) GetBooleanOk() (*bool, bool)`

GetBooleanOk returns a tuple with the Boolean field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoolean

`func (o *FormAnswer) SetBoolean(v bool)`

SetBoolean sets Boolean field to given value.

### HasBoolean

`func (o *FormAnswer) HasBoolean() bool`

HasBoolean returns a boolean if a field has been set.

### GetChoices

`func (o *FormAnswer) GetChoices() []string`

GetChoices returns the Choices field if non-nil, zero value otherwise.

### GetChoicesOk

`func (o *FormAnswer) GetChoicesOk() (*[]string, bool)`

GetChoicesOk returns a tuple with the Choices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChoices

`func (o *FormAnswer) SetChoices(v []string)`

SetChoices sets Choices field to given value.

### HasChoices

`func (o *FormAnswer) HasChoices() bool`

HasChoices returns a boolean if a field has been set.

### GetDate

`func (o *FormAnswer) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *FormAnswer) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *FormAnswer) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *FormAnswer) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetField

`func (o *FormAnswer) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *FormAnswer) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *FormAnswer) SetField(v string)`

SetField sets Field field to given value.

### HasField

`func (o *FormAnswer) HasField() bool`

HasField returns a boolean if a field has been set.

### GetNumber

`func (o *FormAnswer) GetNumber() float64`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *FormAnswer) GetNumberOk() (*float64, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *FormAnswer) SetNumber(v float64)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *FormAnswer) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetText

`func (o *FormAnswer) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *FormAnswer) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *FormAnswer) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *FormAnswer) HasText() bool`

HasText returns a boolean if a field has been set.

### GetType

`func (o *FormAnswer) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FormAnswer) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FormAnswer) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FormAnswer) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


