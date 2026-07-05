# IamControllersLdapSyncResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Exist** | Pointer to [**[]IamObjectLdapUser**](IamObjectLdapUser.md) |  | [optional] 
**Failed** | Pointer to [**[]IamObjectLdapUser**](IamObjectLdapUser.md) |  | [optional] 

## Methods

### NewIamControllersLdapSyncResp

`func NewIamControllersLdapSyncResp() *IamControllersLdapSyncResp`

NewIamControllersLdapSyncResp instantiates a new IamControllersLdapSyncResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamControllersLdapSyncRespWithDefaults

`func NewIamControllersLdapSyncRespWithDefaults() *IamControllersLdapSyncResp`

NewIamControllersLdapSyncRespWithDefaults instantiates a new IamControllersLdapSyncResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExist

`func (o *IamControllersLdapSyncResp) GetExist() []IamObjectLdapUser`

GetExist returns the Exist field if non-nil, zero value otherwise.

### GetExistOk

`func (o *IamControllersLdapSyncResp) GetExistOk() (*[]IamObjectLdapUser, bool)`

GetExistOk returns a tuple with the Exist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExist

`func (o *IamControllersLdapSyncResp) SetExist(v []IamObjectLdapUser)`

SetExist sets Exist field to given value.

### HasExist

`func (o *IamControllersLdapSyncResp) HasExist() bool`

HasExist returns a boolean if a field has been set.

### GetFailed

`func (o *IamControllersLdapSyncResp) GetFailed() []IamObjectLdapUser`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *IamControllersLdapSyncResp) GetFailedOk() (*[]IamObjectLdapUser, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *IamControllersLdapSyncResp) SetFailed(v []IamObjectLdapUser)`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *IamControllersLdapSyncResp) HasFailed() bool`

HasFailed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


