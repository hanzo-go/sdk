# O11yO11yChangePasswordIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewPassword** | Pointer to **string** | NewPassword is the password to set. | [optional] 
**OldPassword** | Pointer to **string** | OldPassword is the current password; the change is refused when it does not match. | [optional] 

## Methods

### NewO11yO11yChangePasswordIn

`func NewO11yO11yChangePasswordIn() *O11yO11yChangePasswordIn`

NewO11yO11yChangePasswordIn instantiates a new O11yO11yChangePasswordIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yChangePasswordInWithDefaults

`func NewO11yO11yChangePasswordInWithDefaults() *O11yO11yChangePasswordIn`

NewO11yO11yChangePasswordInWithDefaults instantiates a new O11yO11yChangePasswordIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewPassword

`func (o *O11yO11yChangePasswordIn) GetNewPassword() string`

GetNewPassword returns the NewPassword field if non-nil, zero value otherwise.

### GetNewPasswordOk

`func (o *O11yO11yChangePasswordIn) GetNewPasswordOk() (*string, bool)`

GetNewPasswordOk returns a tuple with the NewPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPassword

`func (o *O11yO11yChangePasswordIn) SetNewPassword(v string)`

SetNewPassword sets NewPassword field to given value.

### HasNewPassword

`func (o *O11yO11yChangePasswordIn) HasNewPassword() bool`

HasNewPassword returns a boolean if a field has been set.

### GetOldPassword

`func (o *O11yO11yChangePasswordIn) GetOldPassword() string`

GetOldPassword returns the OldPassword field if non-nil, zero value otherwise.

### GetOldPasswordOk

`func (o *O11yO11yChangePasswordIn) GetOldPasswordOk() (*string, bool)`

GetOldPasswordOk returns a tuple with the OldPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldPassword

`func (o *O11yO11yChangePasswordIn) SetOldPassword(v string)`

SetOldPassword sets OldPassword field to given value.

### HasOldPassword

`func (o *O11yO11yChangePasswordIn) HasOldPassword() bool`

HasOldPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


