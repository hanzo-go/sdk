# O11yO11yRoleDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** | CreatedAt is when the role was created. | [optional] 
**Description** | Pointer to **string** | Description says what the role is for. | [optional] 
**Id** | Pointer to **string** | ID is the role id. | [optional] 
**Name** | Pointer to **string** | Name is the role&#39;s name. | [optional] 
**OrgId** | Pointer to **string** | OrgID is the org the role belongs to. | [optional] 
**TransactionGroups** | Pointer to [**[]O11yO11yTransactionGroup**](O11yO11yTransactionGroup.md) | TransactionGroups are the grants the role carries. | [optional] 
**Type** | Pointer to **string** | Type is custom or managed. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | UpdatedAt is when it last changed. | [optional] 

## Methods

### NewO11yO11yRoleDetail

`func NewO11yO11yRoleDetail() *O11yO11yRoleDetail`

NewO11yO11yRoleDetail instantiates a new O11yO11yRoleDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yRoleDetailWithDefaults

`func NewO11yO11yRoleDetailWithDefaults() *O11yO11yRoleDetail`

NewO11yO11yRoleDetailWithDefaults instantiates a new O11yO11yRoleDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *O11yO11yRoleDetail) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *O11yO11yRoleDetail) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *O11yO11yRoleDetail) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *O11yO11yRoleDetail) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *O11yO11yRoleDetail) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *O11yO11yRoleDetail) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *O11yO11yRoleDetail) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *O11yO11yRoleDetail) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *O11yO11yRoleDetail) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *O11yO11yRoleDetail) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *O11yO11yRoleDetail) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *O11yO11yRoleDetail) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *O11yO11yRoleDetail) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *O11yO11yRoleDetail) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *O11yO11yRoleDetail) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *O11yO11yRoleDetail) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrgId

`func (o *O11yO11yRoleDetail) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *O11yO11yRoleDetail) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *O11yO11yRoleDetail) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *O11yO11yRoleDetail) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetTransactionGroups

`func (o *O11yO11yRoleDetail) GetTransactionGroups() []O11yO11yTransactionGroup`

GetTransactionGroups returns the TransactionGroups field if non-nil, zero value otherwise.

### GetTransactionGroupsOk

`func (o *O11yO11yRoleDetail) GetTransactionGroupsOk() (*[]O11yO11yTransactionGroup, bool)`

GetTransactionGroupsOk returns a tuple with the TransactionGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionGroups

`func (o *O11yO11yRoleDetail) SetTransactionGroups(v []O11yO11yTransactionGroup)`

SetTransactionGroups sets TransactionGroups field to given value.

### HasTransactionGroups

`func (o *O11yO11yRoleDetail) HasTransactionGroups() bool`

HasTransactionGroups returns a boolean if a field has been set.

### GetType

`func (o *O11yO11yRoleDetail) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *O11yO11yRoleDetail) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *O11yO11yRoleDetail) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *O11yO11yRoleDetail) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *O11yO11yRoleDetail) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *O11yO11yRoleDetail) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *O11yO11yRoleDetail) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *O11yO11yRoleDetail) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


