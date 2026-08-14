# O11yO11yRoleUpdateIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description says what the role is for. Required — send an empty string to clear it. | [optional] 
**TransactionGroups** | Pointer to [**[]O11yO11yTransactionGroup**](O11yO11yTransactionGroup.md) | TransactionGroups are the grants the role carries. Required — send an empty array to clear them. | [optional] 

## Methods

### NewO11yO11yRoleUpdateIn

`func NewO11yO11yRoleUpdateIn() *O11yO11yRoleUpdateIn`

NewO11yO11yRoleUpdateIn instantiates a new O11yO11yRoleUpdateIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yRoleUpdateInWithDefaults

`func NewO11yO11yRoleUpdateInWithDefaults() *O11yO11yRoleUpdateIn`

NewO11yO11yRoleUpdateInWithDefaults instantiates a new O11yO11yRoleUpdateIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *O11yO11yRoleUpdateIn) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *O11yO11yRoleUpdateIn) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *O11yO11yRoleUpdateIn) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *O11yO11yRoleUpdateIn) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTransactionGroups

`func (o *O11yO11yRoleUpdateIn) GetTransactionGroups() []O11yO11yTransactionGroup`

GetTransactionGroups returns the TransactionGroups field if non-nil, zero value otherwise.

### GetTransactionGroupsOk

`func (o *O11yO11yRoleUpdateIn) GetTransactionGroupsOk() (*[]O11yO11yTransactionGroup, bool)`

GetTransactionGroupsOk returns a tuple with the TransactionGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionGroups

`func (o *O11yO11yRoleUpdateIn) SetTransactionGroups(v []O11yO11yTransactionGroup)`

SetTransactionGroups sets TransactionGroups field to given value.

### HasTransactionGroups

`func (o *O11yO11yRoleUpdateIn) HasTransactionGroups() bool`

HasTransactionGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


