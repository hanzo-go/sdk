# O11yO11yDeprecatedUserUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | DisplayName is the new display name; empty leaves it unchanged. | [optional] 
**Role** | Pointer to **string** | Role is the legacy role to move to — ADMIN, EDITOR or VIEWER; empty leaves it unchanged. | [optional] 

## Methods

### NewO11yO11yDeprecatedUserUpdate

`func NewO11yO11yDeprecatedUserUpdate() *O11yO11yDeprecatedUserUpdate`

NewO11yO11yDeprecatedUserUpdate instantiates a new O11yO11yDeprecatedUserUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yDeprecatedUserUpdateWithDefaults

`func NewO11yO11yDeprecatedUserUpdateWithDefaults() *O11yO11yDeprecatedUserUpdate`

NewO11yO11yDeprecatedUserUpdateWithDefaults instantiates a new O11yO11yDeprecatedUserUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *O11yO11yDeprecatedUserUpdate) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *O11yO11yDeprecatedUserUpdate) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *O11yO11yDeprecatedUserUpdate) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *O11yO11yDeprecatedUserUpdate) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetRole

`func (o *O11yO11yDeprecatedUserUpdate) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *O11yO11yDeprecatedUserUpdate) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *O11yO11yDeprecatedUserUpdate) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *O11yO11yDeprecatedUserUpdate) HasRole() bool`

HasRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


