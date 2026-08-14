# O11yBasicAuth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | Pointer to **interface{}** |  | [optional] 
**PasswordFile** | Pointer to **string** |  | [optional] 
**PasswordRef** | Pointer to **string** | PasswordRef is the name of the secret within the secret manager to use as the password. | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**UsernameFile** | Pointer to **string** |  | [optional] 
**UsernameRef** | Pointer to **string** | UsernameRef is the name of the secret within the secret manager to use as the username. | [optional] 

## Methods

### NewO11yBasicAuth

`func NewO11yBasicAuth() *O11yBasicAuth`

NewO11yBasicAuth instantiates a new O11yBasicAuth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yBasicAuthWithDefaults

`func NewO11yBasicAuthWithDefaults() *O11yBasicAuth`

NewO11yBasicAuthWithDefaults instantiates a new O11yBasicAuth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *O11yBasicAuth) GetPassword() interface{}`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *O11yBasicAuth) GetPasswordOk() (*interface{}, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *O11yBasicAuth) SetPassword(v interface{})`

SetPassword sets Password field to given value.

### HasPassword

`func (o *O11yBasicAuth) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *O11yBasicAuth) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *O11yBasicAuth) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetPasswordFile

`func (o *O11yBasicAuth) GetPasswordFile() string`

GetPasswordFile returns the PasswordFile field if non-nil, zero value otherwise.

### GetPasswordFileOk

`func (o *O11yBasicAuth) GetPasswordFileOk() (*string, bool)`

GetPasswordFileOk returns a tuple with the PasswordFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordFile

`func (o *O11yBasicAuth) SetPasswordFile(v string)`

SetPasswordFile sets PasswordFile field to given value.

### HasPasswordFile

`func (o *O11yBasicAuth) HasPasswordFile() bool`

HasPasswordFile returns a boolean if a field has been set.

### GetPasswordRef

`func (o *O11yBasicAuth) GetPasswordRef() string`

GetPasswordRef returns the PasswordRef field if non-nil, zero value otherwise.

### GetPasswordRefOk

`func (o *O11yBasicAuth) GetPasswordRefOk() (*string, bool)`

GetPasswordRefOk returns a tuple with the PasswordRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordRef

`func (o *O11yBasicAuth) SetPasswordRef(v string)`

SetPasswordRef sets PasswordRef field to given value.

### HasPasswordRef

`func (o *O11yBasicAuth) HasPasswordRef() bool`

HasPasswordRef returns a boolean if a field has been set.

### GetUsername

`func (o *O11yBasicAuth) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *O11yBasicAuth) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *O11yBasicAuth) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *O11yBasicAuth) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetUsernameFile

`func (o *O11yBasicAuth) GetUsernameFile() string`

GetUsernameFile returns the UsernameFile field if non-nil, zero value otherwise.

### GetUsernameFileOk

`func (o *O11yBasicAuth) GetUsernameFileOk() (*string, bool)`

GetUsernameFileOk returns a tuple with the UsernameFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsernameFile

`func (o *O11yBasicAuth) SetUsernameFile(v string)`

SetUsernameFile sets UsernameFile field to given value.

### HasUsernameFile

`func (o *O11yBasicAuth) HasUsernameFile() bool`

HasUsernameFile returns a boolean if a field has been set.

### GetUsernameRef

`func (o *O11yBasicAuth) GetUsernameRef() string`

GetUsernameRef returns the UsernameRef field if non-nil, zero value otherwise.

### GetUsernameRefOk

`func (o *O11yBasicAuth) GetUsernameRefOk() (*string, bool)`

GetUsernameRefOk returns a tuple with the UsernameRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsernameRef

`func (o *O11yBasicAuth) SetUsernameRef(v string)`

SetUsernameRef sets UsernameRef field to given value.

### HasUsernameRef

`func (o *O11yBasicAuth) HasUsernameRef() bool`

HasUsernameRef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


