# O11yO11yServiceAccountRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** | CreatedAt is when the role was assigned. | [optional] 
**Id** | Pointer to **string** | ID is the assignment&#39;s own id. | [optional] 
**Role** | Pointer to [**O11yO11yRole**](O11yO11yRole.md) | Role is the role itself. | [optional] 
**RoleId** | Pointer to **string** | RoleID is the role held. | [optional] 
**ServiceAccountId** | Pointer to **string** | ServiceAccountID is the account holding the role. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | UpdatedAt is when the assignment last changed. | [optional] 

## Methods

### NewO11yO11yServiceAccountRole

`func NewO11yO11yServiceAccountRole() *O11yO11yServiceAccountRole`

NewO11yO11yServiceAccountRole instantiates a new O11yO11yServiceAccountRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yServiceAccountRoleWithDefaults

`func NewO11yO11yServiceAccountRoleWithDefaults() *O11yO11yServiceAccountRole`

NewO11yO11yServiceAccountRoleWithDefaults instantiates a new O11yO11yServiceAccountRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *O11yO11yServiceAccountRole) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *O11yO11yServiceAccountRole) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *O11yO11yServiceAccountRole) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *O11yO11yServiceAccountRole) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *O11yO11yServiceAccountRole) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *O11yO11yServiceAccountRole) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *O11yO11yServiceAccountRole) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *O11yO11yServiceAccountRole) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRole

`func (o *O11yO11yServiceAccountRole) GetRole() O11yO11yRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *O11yO11yServiceAccountRole) GetRoleOk() (*O11yO11yRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *O11yO11yServiceAccountRole) SetRole(v O11yO11yRole)`

SetRole sets Role field to given value.

### HasRole

`func (o *O11yO11yServiceAccountRole) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetRoleId

`func (o *O11yO11yServiceAccountRole) GetRoleId() string`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *O11yO11yServiceAccountRole) GetRoleIdOk() (*string, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *O11yO11yServiceAccountRole) SetRoleId(v string)`

SetRoleId sets RoleId field to given value.

### HasRoleId

`func (o *O11yO11yServiceAccountRole) HasRoleId() bool`

HasRoleId returns a boolean if a field has been set.

### GetServiceAccountId

`func (o *O11yO11yServiceAccountRole) GetServiceAccountId() string`

GetServiceAccountId returns the ServiceAccountId field if non-nil, zero value otherwise.

### GetServiceAccountIdOk

`func (o *O11yO11yServiceAccountRole) GetServiceAccountIdOk() (*string, bool)`

GetServiceAccountIdOk returns a tuple with the ServiceAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountId

`func (o *O11yO11yServiceAccountRole) SetServiceAccountId(v string)`

SetServiceAccountId sets ServiceAccountId field to given value.

### HasServiceAccountId

`func (o *O11yO11yServiceAccountRole) HasServiceAccountId() bool`

HasServiceAccountId returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *O11yO11yServiceAccountRole) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *O11yO11yServiceAccountRole) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *O11yO11yServiceAccountRole) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *O11yO11yServiceAccountRole) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


