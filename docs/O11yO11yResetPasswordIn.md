# O11yO11yResetPasswordIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | Pointer to **string** | Password is the new password. | [optional] 
**Token** | Pointer to **string** | Token is the reset-password token authorizing the change. | [optional] 

## Methods

### NewO11yO11yResetPasswordIn

`func NewO11yO11yResetPasswordIn() *O11yO11yResetPasswordIn`

NewO11yO11yResetPasswordIn instantiates a new O11yO11yResetPasswordIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yResetPasswordInWithDefaults

`func NewO11yO11yResetPasswordInWithDefaults() *O11yO11yResetPasswordIn`

NewO11yO11yResetPasswordInWithDefaults instantiates a new O11yO11yResetPasswordIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *O11yO11yResetPasswordIn) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *O11yO11yResetPasswordIn) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *O11yO11yResetPasswordIn) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *O11yO11yResetPasswordIn) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetToken

`func (o *O11yO11yResetPasswordIn) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *O11yO11yResetPasswordIn) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *O11yO11yResetPasswordIn) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *O11yO11yResetPasswordIn) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


