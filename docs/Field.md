# Field

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** |  | [optional] 
**Label** | Pointer to **string** |  | [optional] 

## Methods

### NewField

`func NewField() *Field`

NewField instantiates a new Field object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFieldWithDefaults

`func NewFieldWithDefaults() *Field`

NewFieldWithDefaults instantiates a new Field object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *Field) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *Field) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *Field) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *Field) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetLabel

`func (o *Field) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *Field) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *Field) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *Field) HasLabel() bool`

HasLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


