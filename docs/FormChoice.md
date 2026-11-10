# FormChoice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID is the choice&#39;s id, assigned when empty. | [optional] 
**Label** | Pointer to **string** | Label is the choice as shown. At most 200 characters. | [optional] 
**Ref** | Pointer to **string** | Ref is how an answer names the choice; unique in the field, defaults to the id. | [optional] 

## Methods

### NewFormChoice

`func NewFormChoice() *FormChoice`

NewFormChoice instantiates a new FormChoice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormChoiceWithDefaults

`func NewFormChoiceWithDefaults() *FormChoice`

NewFormChoiceWithDefaults instantiates a new FormChoice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FormChoice) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FormChoice) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FormChoice) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FormChoice) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *FormChoice) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *FormChoice) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *FormChoice) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *FormChoice) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetRef

`func (o *FormChoice) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *FormChoice) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *FormChoice) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *FormChoice) HasRef() bool`

HasRef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


