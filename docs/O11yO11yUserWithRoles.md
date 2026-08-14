# O11yO11yUserWithRoles

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** | CreatedAt is when they joined. | [optional] 
**DisplayName** | Pointer to **string** | DisplayName is what the console shows for them. | [optional] 
**Email** | Pointer to **string** | Email is their address. | [optional] 
**Id** | Pointer to **string** | ID is the user id. | [optional] 
**IsRoot** | Pointer to **bool** | IsRoot marks the org&#39;s root user. | [optional] 
**OrgId** | Pointer to **string** | OrgID is the org they belong to. | [optional] 
**Status** | Pointer to **string** | Status is their lifecycle state — active, pending_invite or deleted. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | UpdatedAt is when their record last changed. | [optional] 
**UserRoles** | Pointer to [**[]O11yO11yUserRole**](O11yO11yUserRole.md) | UserRoles are their role assignments. | [optional] 

## Methods

### NewO11yO11yUserWithRoles

`func NewO11yO11yUserWithRoles() *O11yO11yUserWithRoles`

NewO11yO11yUserWithRoles instantiates a new O11yO11yUserWithRoles object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yUserWithRolesWithDefaults

`func NewO11yO11yUserWithRolesWithDefaults() *O11yO11yUserWithRoles`

NewO11yO11yUserWithRolesWithDefaults instantiates a new O11yO11yUserWithRoles object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *O11yO11yUserWithRoles) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *O11yO11yUserWithRoles) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *O11yO11yUserWithRoles) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *O11yO11yUserWithRoles) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDisplayName

`func (o *O11yO11yUserWithRoles) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *O11yO11yUserWithRoles) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *O11yO11yUserWithRoles) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *O11yO11yUserWithRoles) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEmail

`func (o *O11yO11yUserWithRoles) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *O11yO11yUserWithRoles) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *O11yO11yUserWithRoles) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *O11yO11yUserWithRoles) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetId

`func (o *O11yO11yUserWithRoles) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *O11yO11yUserWithRoles) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *O11yO11yUserWithRoles) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *O11yO11yUserWithRoles) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsRoot

`func (o *O11yO11yUserWithRoles) GetIsRoot() bool`

GetIsRoot returns the IsRoot field if non-nil, zero value otherwise.

### GetIsRootOk

`func (o *O11yO11yUserWithRoles) GetIsRootOk() (*bool, bool)`

GetIsRootOk returns a tuple with the IsRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRoot

`func (o *O11yO11yUserWithRoles) SetIsRoot(v bool)`

SetIsRoot sets IsRoot field to given value.

### HasIsRoot

`func (o *O11yO11yUserWithRoles) HasIsRoot() bool`

HasIsRoot returns a boolean if a field has been set.

### GetOrgId

`func (o *O11yO11yUserWithRoles) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *O11yO11yUserWithRoles) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *O11yO11yUserWithRoles) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *O11yO11yUserWithRoles) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetStatus

`func (o *O11yO11yUserWithRoles) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *O11yO11yUserWithRoles) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *O11yO11yUserWithRoles) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *O11yO11yUserWithRoles) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *O11yO11yUserWithRoles) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *O11yO11yUserWithRoles) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *O11yO11yUserWithRoles) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *O11yO11yUserWithRoles) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUserRoles

`func (o *O11yO11yUserWithRoles) GetUserRoles() []O11yO11yUserRole`

GetUserRoles returns the UserRoles field if non-nil, zero value otherwise.

### GetUserRolesOk

`func (o *O11yO11yUserWithRoles) GetUserRolesOk() (*[]O11yO11yUserRole, bool)`

GetUserRolesOk returns a tuple with the UserRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserRoles

`func (o *O11yO11yUserWithRoles) SetUserRoles(v []O11yO11yUserRole)`

SetUserRoles sets UserRoles field to given value.

### HasUserRoles

`func (o *O11yO11yUserWithRoles) HasUserRoles() bool`

HasUserRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


