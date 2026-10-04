# FrameworkModule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Doctypes** | Pointer to **[]string** |  | [optional] 
**Enabled** | Pointer to **bool** | Enabled is whether this org has turned the module on. A module that is off answers 404 on every DocType it owns (elective.go). | [optional] 
**Module** | Pointer to **string** |  | [optional] 

## Methods

### NewFrameworkModule

`func NewFrameworkModule() *FrameworkModule`

NewFrameworkModule instantiates a new FrameworkModule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFrameworkModuleWithDefaults

`func NewFrameworkModuleWithDefaults() *FrameworkModule`

NewFrameworkModuleWithDefaults instantiates a new FrameworkModule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDoctypes

`func (o *FrameworkModule) GetDoctypes() []string`

GetDoctypes returns the Doctypes field if non-nil, zero value otherwise.

### GetDoctypesOk

`func (o *FrameworkModule) GetDoctypesOk() (*[]string, bool)`

GetDoctypesOk returns a tuple with the Doctypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoctypes

`func (o *FrameworkModule) SetDoctypes(v []string)`

SetDoctypes sets Doctypes field to given value.

### HasDoctypes

`func (o *FrameworkModule) HasDoctypes() bool`

HasDoctypes returns a boolean if a field has been set.

### GetEnabled

`func (o *FrameworkModule) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *FrameworkModule) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *FrameworkModule) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *FrameworkModule) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetModule

`func (o *FrameworkModule) GetModule() string`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *FrameworkModule) GetModuleOk() (*string, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *FrameworkModule) SetModule(v string)`

SetModule sets Module field to given value.

### HasModule

`func (o *FrameworkModule) HasModule() bool`

HasModule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


