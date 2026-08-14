# O11yO11yDeprecatedUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** | CreatedAt is when they joined. | [optional] 
**DisplayName** | Pointer to **string** | DisplayName is what the console shows for them. | [optional] 
**Email** | Pointer to **string** | Email is their address. | [optional] 
**Id** | Pointer to **string** | ID is the user id. | [optional] 
**IsRoot** | Pointer to **bool** | IsRoot marks the org&#39;s root user. | [optional] 
**OrgId** | Pointer to **string** | OrgID is the org they belong to. | [optional] 
**Role** | Pointer to **string** | Role is their legacy role — ADMIN, EDITOR or VIEWER. | [optional] 
**Status** | Pointer to **string** | Status is their lifecycle state — active, pending_invite or deleted. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | UpdatedAt is when their record last changed. | [optional] 

## Methods

### NewO11yO11yDeprecatedUser

`func NewO11yO11yDeprecatedUser() *O11yO11yDeprecatedUser`

NewO11yO11yDeprecatedUser instantiates a new O11yO11yDeprecatedUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yDeprecatedUserWithDefaults

`func NewO11yO11yDeprecatedUserWithDefaults() *O11yO11yDeprecatedUser`

NewO11yO11yDeprecatedUserWithDefaults instantiates a new O11yO11yDeprecatedUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *O11yO11yDeprecatedUser) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *O11yO11yDeprecatedUser) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *O11yO11yDeprecatedUser) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *O11yO11yDeprecatedUser) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDisplayName

`func (o *O11yO11yDeprecatedUser) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *O11yO11yDeprecatedUser) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *O11yO11yDeprecatedUser) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *O11yO11yDeprecatedUser) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEmail

`func (o *O11yO11yDeprecatedUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *O11yO11yDeprecatedUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *O11yO11yDeprecatedUser) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *O11yO11yDeprecatedUser) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetId

`func (o *O11yO11yDeprecatedUser) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *O11yO11yDeprecatedUser) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *O11yO11yDeprecatedUser) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *O11yO11yDeprecatedUser) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsRoot

`func (o *O11yO11yDeprecatedUser) GetIsRoot() bool`

GetIsRoot returns the IsRoot field if non-nil, zero value otherwise.

### GetIsRootOk

`func (o *O11yO11yDeprecatedUser) GetIsRootOk() (*bool, bool)`

GetIsRootOk returns a tuple with the IsRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRoot

`func (o *O11yO11yDeprecatedUser) SetIsRoot(v bool)`

SetIsRoot sets IsRoot field to given value.

### HasIsRoot

`func (o *O11yO11yDeprecatedUser) HasIsRoot() bool`

HasIsRoot returns a boolean if a field has been set.

### GetOrgId

`func (o *O11yO11yDeprecatedUser) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *O11yO11yDeprecatedUser) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *O11yO11yDeprecatedUser) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *O11yO11yDeprecatedUser) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetRole

`func (o *O11yO11yDeprecatedUser) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *O11yO11yDeprecatedUser) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *O11yO11yDeprecatedUser) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *O11yO11yDeprecatedUser) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetStatus

`func (o *O11yO11yDeprecatedUser) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *O11yO11yDeprecatedUser) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *O11yO11yDeprecatedUser) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *O11yO11yDeprecatedUser) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *O11yO11yDeprecatedUser) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *O11yO11yDeprecatedUser) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *O11yO11yDeprecatedUser) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *O11yO11yDeprecatedUser) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


