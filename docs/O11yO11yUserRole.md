# O11yO11yUserRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** | CreatedAt is when it was assigned. | [optional] 
**Id** | Pointer to **string** | ID is the assignment&#39;s own id. | [optional] 
**Role** | Pointer to [**O11yO11yRole**](O11yO11yRole.md) | Role is the role itself. | [optional] 
**RoleId** | Pointer to **string** | RoleID is the role held. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | UpdatedAt is when the assignment last changed. | [optional] 
**UserId** | Pointer to **string** | UserID is the user holding the role. | [optional] 

## Methods

### NewO11yO11yUserRole

`func NewO11yO11yUserRole() *O11yO11yUserRole`

NewO11yO11yUserRole instantiates a new O11yO11yUserRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yUserRoleWithDefaults

`func NewO11yO11yUserRoleWithDefaults() *O11yO11yUserRole`

NewO11yO11yUserRoleWithDefaults instantiates a new O11yO11yUserRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *O11yO11yUserRole) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *O11yO11yUserRole) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *O11yO11yUserRole) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *O11yO11yUserRole) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *O11yO11yUserRole) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *O11yO11yUserRole) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *O11yO11yUserRole) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *O11yO11yUserRole) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRole

`func (o *O11yO11yUserRole) GetRole() O11yO11yRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *O11yO11yUserRole) GetRoleOk() (*O11yO11yRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *O11yO11yUserRole) SetRole(v O11yO11yRole)`

SetRole sets Role field to given value.

### HasRole

`func (o *O11yO11yUserRole) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetRoleId

`func (o *O11yO11yUserRole) GetRoleId() string`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *O11yO11yUserRole) GetRoleIdOk() (*string, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *O11yO11yUserRole) SetRoleId(v string)`

SetRoleId sets RoleId field to given value.

### HasRoleId

`func (o *O11yO11yUserRole) HasRoleId() bool`

HasRoleId returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *O11yO11yUserRole) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *O11yO11yUserRole) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *O11yO11yUserRole) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *O11yO11yUserRole) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUserId

`func (o *O11yO11yUserRole) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *O11yO11yUserRole) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *O11yO11yUserRole) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *O11yO11yUserRole) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


