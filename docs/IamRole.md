# IamRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**CreatedTime** | Pointer to **string** |  | [optional] 
**Deleted** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Domains** | Pointer to **[]string** |  | [optional] 
**Groups** | Pointer to **[]string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**IsEnabled** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to **string** |  | [optional] 
**Roles** | Pointer to **[]string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**Users** | Pointer to **[]string** |  | [optional] 

## Methods

### NewIamRole

`func NewIamRole() *IamRole`

NewIamRole instantiates a new IamRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamRoleWithDefaults

`func NewIamRoleWithDefaults() *IamRole`

NewIamRoleWithDefaults instantiates a new IamRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *IamRole) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IamRole) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IamRole) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *IamRole) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedTime

`func (o *IamRole) GetCreatedTime() string`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *IamRole) GetCreatedTimeOk() (*string, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *IamRole) SetCreatedTime(v string)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *IamRole) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetDeleted

`func (o *IamRole) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *IamRole) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *IamRole) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *IamRole) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetDescription

`func (o *IamRole) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IamRole) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IamRole) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IamRole) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *IamRole) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *IamRole) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *IamRole) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *IamRole) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDomains

`func (o *IamRole) GetDomains() []string`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *IamRole) GetDomainsOk() (*[]string, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *IamRole) SetDomains(v []string)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *IamRole) HasDomains() bool`

HasDomains returns a boolean if a field has been set.

### GetGroups

`func (o *IamRole) GetGroups() []string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *IamRole) GetGroupsOk() (*[]string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *IamRole) SetGroups(v []string)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *IamRole) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetId

`func (o *IamRole) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IamRole) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IamRole) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IamRole) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsEnabled

`func (o *IamRole) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *IamRole) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *IamRole) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *IamRole) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetName

`func (o *IamRole) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IamRole) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IamRole) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IamRole) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwner

`func (o *IamRole) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *IamRole) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *IamRole) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *IamRole) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetRoles

`func (o *IamRole) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *IamRole) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *IamRole) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *IamRole) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *IamRole) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *IamRole) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *IamRole) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *IamRole) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUsers

`func (o *IamRole) GetUsers() []string`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *IamRole) GetUsersOk() (*[]string, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *IamRole) SetUsers(v []string)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *IamRole) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


