# O11yO11yResetToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpiresAt** | Pointer to **time.Time** | ExpiresAt is when it stops working. | [optional] 
**Id** | Pointer to **string** | ID is the grant&#39;s id. | [optional] 
**PasswordId** | Pointer to **string** | PasswordID is the password record it resets. | [optional] 
**Token** | Pointer to **string** | Token is the secret that redeems it. | [optional] 

## Methods

### NewO11yO11yResetToken

`func NewO11yO11yResetToken() *O11yO11yResetToken`

NewO11yO11yResetToken instantiates a new O11yO11yResetToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yResetTokenWithDefaults

`func NewO11yO11yResetTokenWithDefaults() *O11yO11yResetToken`

NewO11yO11yResetTokenWithDefaults instantiates a new O11yO11yResetToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiresAt

`func (o *O11yO11yResetToken) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *O11yO11yResetToken) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *O11yO11yResetToken) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *O11yO11yResetToken) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetId

`func (o *O11yO11yResetToken) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *O11yO11yResetToken) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *O11yO11yResetToken) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *O11yO11yResetToken) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPasswordId

`func (o *O11yO11yResetToken) GetPasswordId() string`

GetPasswordId returns the PasswordId field if non-nil, zero value otherwise.

### GetPasswordIdOk

`func (o *O11yO11yResetToken) GetPasswordIdOk() (*string, bool)`

GetPasswordIdOk returns a tuple with the PasswordId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordId

`func (o *O11yO11yResetToken) SetPasswordId(v string)`

SetPasswordId sets PasswordId field to given value.

### HasPasswordId

`func (o *O11yO11yResetToken) HasPasswordId() bool`

HasPasswordId returns a boolean if a field has been set.

### GetToken

`func (o *O11yO11yResetToken) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *O11yO11yResetToken) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *O11yO11yResetToken) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *O11yO11yResetToken) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


