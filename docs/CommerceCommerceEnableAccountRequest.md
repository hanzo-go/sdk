# CommerceCommerceEnableAccountRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | Pointer to **string** | New account password; applied only when non-empty. | [optional] 
**PasswordConfirm** | Pointer to **string** | Password confirmation. | [optional] 

## Methods

### NewCommerceCommerceEnableAccountRequest

`func NewCommerceCommerceEnableAccountRequest() *CommerceCommerceEnableAccountRequest`

NewCommerceCommerceEnableAccountRequest instantiates a new CommerceCommerceEnableAccountRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommerceCommerceEnableAccountRequestWithDefaults

`func NewCommerceCommerceEnableAccountRequestWithDefaults() *CommerceCommerceEnableAccountRequest`

NewCommerceCommerceEnableAccountRequestWithDefaults instantiates a new CommerceCommerceEnableAccountRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *CommerceCommerceEnableAccountRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CommerceCommerceEnableAccountRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CommerceCommerceEnableAccountRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *CommerceCommerceEnableAccountRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPasswordConfirm

`func (o *CommerceCommerceEnableAccountRequest) GetPasswordConfirm() string`

GetPasswordConfirm returns the PasswordConfirm field if non-nil, zero value otherwise.

### GetPasswordConfirmOk

`func (o *CommerceCommerceEnableAccountRequest) GetPasswordConfirmOk() (*string, bool)`

GetPasswordConfirmOk returns a tuple with the PasswordConfirm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordConfirm

`func (o *CommerceCommerceEnableAccountRequest) SetPasswordConfirm(v string)`

SetPasswordConfirm sets PasswordConfirm field to given value.

### HasPasswordConfirm

`func (o *CommerceCommerceEnableAccountRequest) HasPasswordConfirm() bool`

HasPasswordConfirm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


