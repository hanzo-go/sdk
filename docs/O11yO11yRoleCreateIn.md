# O11yO11yRoleCreateIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description says what the role is for. | [optional] 
**Name** | Pointer to **string** | Name is the role&#39;s name: lowercase letters and hyphens, at most 50 characters, not starting with the reserved managed-role prefix. Required. | [optional] 
**TransactionGroups** | Pointer to [**[]O11yO11yTransactionGroup**](O11yO11yTransactionGroup.md) | TransactionGroups are the grants the role carries. | [optional] 

## Methods

### NewO11yO11yRoleCreateIn

`func NewO11yO11yRoleCreateIn() *O11yO11yRoleCreateIn`

NewO11yO11yRoleCreateIn instantiates a new O11yO11yRoleCreateIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yRoleCreateInWithDefaults

`func NewO11yO11yRoleCreateInWithDefaults() *O11yO11yRoleCreateIn`

NewO11yO11yRoleCreateInWithDefaults instantiates a new O11yO11yRoleCreateIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *O11yO11yRoleCreateIn) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *O11yO11yRoleCreateIn) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *O11yO11yRoleCreateIn) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *O11yO11yRoleCreateIn) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *O11yO11yRoleCreateIn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *O11yO11yRoleCreateIn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *O11yO11yRoleCreateIn) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *O11yO11yRoleCreateIn) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTransactionGroups

`func (o *O11yO11yRoleCreateIn) GetTransactionGroups() []O11yO11yTransactionGroup`

GetTransactionGroups returns the TransactionGroups field if non-nil, zero value otherwise.

### GetTransactionGroupsOk

`func (o *O11yO11yRoleCreateIn) GetTransactionGroupsOk() (*[]O11yO11yTransactionGroup, bool)`

GetTransactionGroupsOk returns a tuple with the TransactionGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionGroups

`func (o *O11yO11yRoleCreateIn) SetTransactionGroups(v []O11yO11yTransactionGroup)`

SetTransactionGroups sets TransactionGroups field to given value.

### HasTransactionGroups

`func (o *O11yO11yRoleCreateIn) HasTransactionGroups() bool`

HasTransactionGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


