# O11yO11yPreference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedScopes** | Pointer to **[]string** | AllowedScopes are the scopes the preference may be set at — org, user. | [optional] 
**AllowedValues** | Pointer to **[]string** | AllowedValues restricts a string preference to these values. | [optional] 
**DefaultValue** | Pointer to **map[string]interface{}** | DefaultValue is the value before anyone set one. | [optional] 
**Description** | Pointer to **string** | Description says what the preference does. | [optional] 
**Name** | Pointer to **string** | Name is the preference name. | [optional] 
**Value** | Pointer to **map[string]interface{}** | Value is the current value. | [optional] 
**ValueType** | Pointer to **string** | ValueType is the JSON type a value must have — string, integer, float or boolean. | [optional] 

## Methods

### NewO11yO11yPreference

`func NewO11yO11yPreference() *O11yO11yPreference`

NewO11yO11yPreference instantiates a new O11yO11yPreference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yPreferenceWithDefaults

`func NewO11yO11yPreferenceWithDefaults() *O11yO11yPreference`

NewO11yO11yPreferenceWithDefaults instantiates a new O11yO11yPreference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedScopes

`func (o *O11yO11yPreference) GetAllowedScopes() []string`

GetAllowedScopes returns the AllowedScopes field if non-nil, zero value otherwise.

### GetAllowedScopesOk

`func (o *O11yO11yPreference) GetAllowedScopesOk() (*[]string, bool)`

GetAllowedScopesOk returns a tuple with the AllowedScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedScopes

`func (o *O11yO11yPreference) SetAllowedScopes(v []string)`

SetAllowedScopes sets AllowedScopes field to given value.

### HasAllowedScopes

`func (o *O11yO11yPreference) HasAllowedScopes() bool`

HasAllowedScopes returns a boolean if a field has been set.

### GetAllowedValues

`func (o *O11yO11yPreference) GetAllowedValues() []string`

GetAllowedValues returns the AllowedValues field if non-nil, zero value otherwise.

### GetAllowedValuesOk

`func (o *O11yO11yPreference) GetAllowedValuesOk() (*[]string, bool)`

GetAllowedValuesOk returns a tuple with the AllowedValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedValues

`func (o *O11yO11yPreference) SetAllowedValues(v []string)`

SetAllowedValues sets AllowedValues field to given value.

### HasAllowedValues

`func (o *O11yO11yPreference) HasAllowedValues() bool`

HasAllowedValues returns a boolean if a field has been set.

### GetDefaultValue

`func (o *O11yO11yPreference) GetDefaultValue() map[string]interface{}`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *O11yO11yPreference) GetDefaultValueOk() (*map[string]interface{}, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *O11yO11yPreference) SetDefaultValue(v map[string]interface{})`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *O11yO11yPreference) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### GetDescription

`func (o *O11yO11yPreference) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *O11yO11yPreference) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *O11yO11yPreference) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *O11yO11yPreference) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *O11yO11yPreference) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *O11yO11yPreference) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *O11yO11yPreference) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *O11yO11yPreference) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValue

`func (o *O11yO11yPreference) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *O11yO11yPreference) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *O11yO11yPreference) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *O11yO11yPreference) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetValueType

`func (o *O11yO11yPreference) GetValueType() string`

GetValueType returns the ValueType field if non-nil, zero value otherwise.

### GetValueTypeOk

`func (o *O11yO11yPreference) GetValueTypeOk() (*string, bool)`

GetValueTypeOk returns a tuple with the ValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueType

`func (o *O11yO11yPreference) SetValueType(v string)`

SetValueType sets ValueType field to given value.

### HasValueType

`func (o *O11yO11yPreference) HasValueType() bool`

HasValueType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


