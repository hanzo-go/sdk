# O11yO11yRoleMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultRole** | Pointer to **string** | DefaultRole is the role when no group mapping applies. | [optional] 
**GroupMappings** | Pointer to **map[string]string** | GroupMappings maps a provider group name to a role name. | [optional] 
**UseRoleAttribute** | Pointer to **bool** | UseRoleAttribute reads the role straight from the provider&#39;s role claim instead of the group mappings. | [optional] 

## Methods

### NewO11yO11yRoleMapping

`func NewO11yO11yRoleMapping() *O11yO11yRoleMapping`

NewO11yO11yRoleMapping instantiates a new O11yO11yRoleMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yRoleMappingWithDefaults

`func NewO11yO11yRoleMappingWithDefaults() *O11yO11yRoleMapping`

NewO11yO11yRoleMappingWithDefaults instantiates a new O11yO11yRoleMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultRole

`func (o *O11yO11yRoleMapping) GetDefaultRole() string`

GetDefaultRole returns the DefaultRole field if non-nil, zero value otherwise.

### GetDefaultRoleOk

`func (o *O11yO11yRoleMapping) GetDefaultRoleOk() (*string, bool)`

GetDefaultRoleOk returns a tuple with the DefaultRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRole

`func (o *O11yO11yRoleMapping) SetDefaultRole(v string)`

SetDefaultRole sets DefaultRole field to given value.

### HasDefaultRole

`func (o *O11yO11yRoleMapping) HasDefaultRole() bool`

HasDefaultRole returns a boolean if a field has been set.

### GetGroupMappings

`func (o *O11yO11yRoleMapping) GetGroupMappings() map[string]string`

GetGroupMappings returns the GroupMappings field if non-nil, zero value otherwise.

### GetGroupMappingsOk

`func (o *O11yO11yRoleMapping) GetGroupMappingsOk() (*map[string]string, bool)`

GetGroupMappingsOk returns a tuple with the GroupMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupMappings

`func (o *O11yO11yRoleMapping) SetGroupMappings(v map[string]string)`

SetGroupMappings sets GroupMappings field to given value.

### HasGroupMappings

`func (o *O11yO11yRoleMapping) HasGroupMappings() bool`

HasGroupMappings returns a boolean if a field has been set.

### GetUseRoleAttribute

`func (o *O11yO11yRoleMapping) GetUseRoleAttribute() bool`

GetUseRoleAttribute returns the UseRoleAttribute field if non-nil, zero value otherwise.

### GetUseRoleAttributeOk

`func (o *O11yO11yRoleMapping) GetUseRoleAttributeOk() (*bool, bool)`

GetUseRoleAttributeOk returns a tuple with the UseRoleAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseRoleAttribute

`func (o *O11yO11yRoleMapping) SetUseRoleAttribute(v bool)`

SetUseRoleAttribute sets UseRoleAttribute field to given value.

### HasUseRoleAttribute

`func (o *O11yO11yRoleMapping) HasUseRoleAttribute() bool`

HasUseRoleAttribute returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


