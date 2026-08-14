# ModuleState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Doctypes** | Pointer to **[]string** |  | [optional] 
**Installed** | Pointer to **[]string** |  | [optional] 
**Module** | Pointer to **string** |  | [optional] 

## Methods

### NewModuleState

`func NewModuleState() *ModuleState`

NewModuleState instantiates a new ModuleState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModuleStateWithDefaults

`func NewModuleStateWithDefaults() *ModuleState`

NewModuleStateWithDefaults instantiates a new ModuleState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDoctypes

`func (o *ModuleState) GetDoctypes() []string`

GetDoctypes returns the Doctypes field if non-nil, zero value otherwise.

### GetDoctypesOk

`func (o *ModuleState) GetDoctypesOk() (*[]string, bool)`

GetDoctypesOk returns a tuple with the Doctypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoctypes

`func (o *ModuleState) SetDoctypes(v []string)`

SetDoctypes sets Doctypes field to given value.

### HasDoctypes

`func (o *ModuleState) HasDoctypes() bool`

HasDoctypes returns a boolean if a field has been set.

### GetInstalled

`func (o *ModuleState) GetInstalled() []string`

GetInstalled returns the Installed field if non-nil, zero value otherwise.

### GetInstalledOk

`func (o *ModuleState) GetInstalledOk() (*[]string, bool)`

GetInstalledOk returns a tuple with the Installed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstalled

`func (o *ModuleState) SetInstalled(v []string)`

SetInstalled sets Installed field to given value.

### HasInstalled

`func (o *ModuleState) HasInstalled() bool`

HasInstalled returns a boolean if a field has been set.

### GetModule

`func (o *ModuleState) GetModule() string`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *ModuleState) GetModuleOk() (*string, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *ModuleState) SetModule(v string)`

SetModule sets Module field to given value.

### HasModule

`func (o *ModuleState) HasModule() bool`

HasModule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


