# IamObjectRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedTime** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Domains** | Pointer to **[]string** |  | [optional] 
**Groups** | Pointer to **[]string** |  | [optional] 
**IsEnabled** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to **string** |  | [optional] 
**Roles** | Pointer to **[]string** |  | [optional] 
**Users** | Pointer to **[]string** |  | [optional] 

## Methods

### NewIamObjectRole

`func NewIamObjectRole() *IamObjectRole`

NewIamObjectRole instantiates a new IamObjectRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamObjectRoleWithDefaults

`func NewIamObjectRoleWithDefaults() *IamObjectRole`

NewIamObjectRoleWithDefaults instantiates a new IamObjectRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedTime

`func (o *IamObjectRole) GetCreatedTime() string`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *IamObjectRole) GetCreatedTimeOk() (*string, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *IamObjectRole) SetCreatedTime(v string)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *IamObjectRole) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetDescription

`func (o *IamObjectRole) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IamObjectRole) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IamObjectRole) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IamObjectRole) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *IamObjectRole) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *IamObjectRole) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *IamObjectRole) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *IamObjectRole) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDomains

`func (o *IamObjectRole) GetDomains() []string`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *IamObjectRole) GetDomainsOk() (*[]string, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *IamObjectRole) SetDomains(v []string)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *IamObjectRole) HasDomains() bool`

HasDomains returns a boolean if a field has been set.

### GetGroups

`func (o *IamObjectRole) GetGroups() []string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *IamObjectRole) GetGroupsOk() (*[]string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *IamObjectRole) SetGroups(v []string)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *IamObjectRole) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetIsEnabled

`func (o *IamObjectRole) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *IamObjectRole) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *IamObjectRole) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *IamObjectRole) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetName

`func (o *IamObjectRole) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IamObjectRole) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IamObjectRole) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IamObjectRole) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwner

`func (o *IamObjectRole) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *IamObjectRole) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *IamObjectRole) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *IamObjectRole) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetRoles

`func (o *IamObjectRole) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *IamObjectRole) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *IamObjectRole) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *IamObjectRole) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetUsers

`func (o *IamObjectRole) GetUsers() []string`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *IamObjectRole) GetUsersOk() (*[]string, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *IamObjectRole) SetUsers(v []string)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *IamObjectRole) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


