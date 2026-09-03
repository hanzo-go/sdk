# IamUsersListOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int64** |  | [optional] 
**Users** | Pointer to [**[]IamUser**](IamUser.md) |  | [optional] 

## Methods

### NewIamUsersListOutput

`func NewIamUsersListOutput() *IamUsersListOutput`

NewIamUsersListOutput instantiates a new IamUsersListOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamUsersListOutputWithDefaults

`func NewIamUsersListOutputWithDefaults() *IamUsersListOutput`

NewIamUsersListOutputWithDefaults instantiates a new IamUsersListOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *IamUsersListOutput) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *IamUsersListOutput) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *IamUsersListOutput) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *IamUsersListOutput) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetUsers

`func (o *IamUsersListOutput) GetUsers() []IamUser`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *IamUsersListOutput) GetUsersOk() (*[]IamUser, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *IamUsersListOutput) SetUsers(v []IamUser)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *IamUsersListOutput) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


