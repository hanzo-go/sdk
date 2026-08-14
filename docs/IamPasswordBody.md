# IamPasswordBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | Code is the one-time code delivered to the account&#39;s own address. | [optional] 
**OldPassword** | Pointer to **string** | OldPassword is the credential being replaced, the proof a signed-in caller gives instead of a code. | [optional] 
**Organization** | Pointer to **string** | The account being recovered, for a caller who cannot be signed in. Read on the CODE arm only — a signed-in caller is resolved from its own session or token, never from these. | [optional] 
**Password** | Pointer to **string** | Password is the new credential. It must satisfy the platform floor and the organization&#39;s own complexity options. | [optional] 
**Username** | Pointer to **string** | email, username OR phone | [optional] 

## Methods

### NewIamPasswordBody

`func NewIamPasswordBody() *IamPasswordBody`

NewIamPasswordBody instantiates a new IamPasswordBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamPasswordBodyWithDefaults

`func NewIamPasswordBodyWithDefaults() *IamPasswordBody`

NewIamPasswordBodyWithDefaults instantiates a new IamPasswordBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *IamPasswordBody) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *IamPasswordBody) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *IamPasswordBody) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *IamPasswordBody) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetOldPassword

`func (o *IamPasswordBody) GetOldPassword() string`

GetOldPassword returns the OldPassword field if non-nil, zero value otherwise.

### GetOldPasswordOk

`func (o *IamPasswordBody) GetOldPasswordOk() (*string, bool)`

GetOldPasswordOk returns a tuple with the OldPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldPassword

`func (o *IamPasswordBody) SetOldPassword(v string)`

SetOldPassword sets OldPassword field to given value.

### HasOldPassword

`func (o *IamPasswordBody) HasOldPassword() bool`

HasOldPassword returns a boolean if a field has been set.

### GetOrganization

`func (o *IamPasswordBody) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *IamPasswordBody) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *IamPasswordBody) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *IamPasswordBody) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetPassword

`func (o *IamPasswordBody) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *IamPasswordBody) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *IamPasswordBody) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *IamPasswordBody) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetUsername

`func (o *IamPasswordBody) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *IamPasswordBody) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *IamPasswordBody) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *IamPasswordBody) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


