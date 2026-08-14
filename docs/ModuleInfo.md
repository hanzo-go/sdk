# ModuleInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Doctypes** | Pointer to **[]string** |  | [optional] 
**Module** | Pointer to **string** |  | [optional] 

## Methods

### NewModuleInfo

`func NewModuleInfo() *ModuleInfo`

NewModuleInfo instantiates a new ModuleInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModuleInfoWithDefaults

`func NewModuleInfoWithDefaults() *ModuleInfo`

NewModuleInfoWithDefaults instantiates a new ModuleInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDoctypes

`func (o *ModuleInfo) GetDoctypes() []string`

GetDoctypes returns the Doctypes field if non-nil, zero value otherwise.

### GetDoctypesOk

`func (o *ModuleInfo) GetDoctypesOk() (*[]string, bool)`

GetDoctypesOk returns a tuple with the Doctypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoctypes

`func (o *ModuleInfo) SetDoctypes(v []string)`

SetDoctypes sets Doctypes field to given value.

### HasDoctypes

`func (o *ModuleInfo) HasDoctypes() bool`

HasDoctypes returns a boolean if a field has been set.

### GetModule

`func (o *ModuleInfo) GetModule() string`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *ModuleInfo) GetModuleOk() (*string, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *ModuleInfo) SetModule(v string)`

SetModule sets Module field to given value.

### HasModule

`func (o *ModuleInfo) HasModule() bool`

HasModule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


