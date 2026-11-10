# FormForward

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** | Email is the address. Required. | [optional] 
**Name** | Pointer to **string** | Name is who it is, for the greeting. Optional. | [optional] 

## Methods

### NewFormForward

`func NewFormForward() *FormForward`

NewFormForward instantiates a new FormForward object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormForwardWithDefaults

`func NewFormForwardWithDefaults() *FormForward`

NewFormForwardWithDefaults instantiates a new FormForward object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *FormForward) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *FormForward) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *FormForward) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *FormForward) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetName

`func (o *FormForward) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FormForward) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FormForward) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FormForward) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


