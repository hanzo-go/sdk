# IamCreateInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | Pointer to **string** |  | [optional] 
**User** | Pointer to [**IamUser**](IamUser.md) |  | [optional] 

## Methods

### NewIamCreateInput

`func NewIamCreateInput() *IamCreateInput`

NewIamCreateInput instantiates a new IamCreateInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamCreateInputWithDefaults

`func NewIamCreateInputWithDefaults() *IamCreateInput`

NewIamCreateInputWithDefaults instantiates a new IamCreateInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *IamCreateInput) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *IamCreateInput) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *IamCreateInput) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *IamCreateInput) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetUser

`func (o *IamCreateInput) GetUser() IamUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *IamCreateInput) GetUserOk() (*IamUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *IamCreateInput) SetUser(v IamUser)`

SetUser sets User field to given value.

### HasUser

`func (o *IamCreateInput) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


