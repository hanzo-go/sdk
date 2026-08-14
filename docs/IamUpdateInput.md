# IamUpdateInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | Pointer to **string** |  | [optional] 
**User** | Pointer to [**IamUser**](IamUser.md) |  | [optional] 

## Methods

### NewIamUpdateInput

`func NewIamUpdateInput() *IamUpdateInput`

NewIamUpdateInput instantiates a new IamUpdateInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamUpdateInputWithDefaults

`func NewIamUpdateInputWithDefaults() *IamUpdateInput`

NewIamUpdateInputWithDefaults instantiates a new IamUpdateInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *IamUpdateInput) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *IamUpdateInput) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *IamUpdateInput) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *IamUpdateInput) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetUser

`func (o *IamUpdateInput) GetUser() IamUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *IamUpdateInput) GetUserOk() (*IamUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *IamUpdateInput) SetUser(v IamUser)`

SetUser sets User field to given value.

### HasUser

`func (o *IamUpdateInput) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


