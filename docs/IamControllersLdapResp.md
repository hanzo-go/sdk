# IamControllersLdapResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExistUuids** | Pointer to **[]string** |  | [optional] 
**Users** | Pointer to [**[]IamObjectLdapUser**](IamObjectLdapUser.md) |  | [optional] 

## Methods

### NewIamControllersLdapResp

`func NewIamControllersLdapResp() *IamControllersLdapResp`

NewIamControllersLdapResp instantiates a new IamControllersLdapResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamControllersLdapRespWithDefaults

`func NewIamControllersLdapRespWithDefaults() *IamControllersLdapResp`

NewIamControllersLdapRespWithDefaults instantiates a new IamControllersLdapResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExistUuids

`func (o *IamControllersLdapResp) GetExistUuids() []string`

GetExistUuids returns the ExistUuids field if non-nil, zero value otherwise.

### GetExistUuidsOk

`func (o *IamControllersLdapResp) GetExistUuidsOk() (*[]string, bool)`

GetExistUuidsOk returns a tuple with the ExistUuids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistUuids

`func (o *IamControllersLdapResp) SetExistUuids(v []string)`

SetExistUuids sets ExistUuids field to given value.

### HasExistUuids

`func (o *IamControllersLdapResp) HasExistUuids() bool`

HasExistUuids returns a boolean if a field has been set.

### GetUsers

`func (o *IamControllersLdapResp) GetUsers() []IamObjectLdapUser`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *IamControllersLdapResp) GetUsersOk() (*[]IamObjectLdapUser, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *IamControllersLdapResp) SetUsers(v []IamObjectLdapUser)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *IamControllersLdapResp) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


