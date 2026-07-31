# CloudModuleState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Doctypes** | Pointer to **[]string** |  | [optional] 
**Installed** | Pointer to **[]string** |  | [optional] 
**Module** | Pointer to **string** |  | [optional] 

## Methods

### NewCloudModuleState

`func NewCloudModuleState() *CloudModuleState`

NewCloudModuleState instantiates a new CloudModuleState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudModuleStateWithDefaults

`func NewCloudModuleStateWithDefaults() *CloudModuleState`

NewCloudModuleStateWithDefaults instantiates a new CloudModuleState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDoctypes

`func (o *CloudModuleState) GetDoctypes() []string`

GetDoctypes returns the Doctypes field if non-nil, zero value otherwise.

### GetDoctypesOk

`func (o *CloudModuleState) GetDoctypesOk() (*[]string, bool)`

GetDoctypesOk returns a tuple with the Doctypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoctypes

`func (o *CloudModuleState) SetDoctypes(v []string)`

SetDoctypes sets Doctypes field to given value.

### HasDoctypes

`func (o *CloudModuleState) HasDoctypes() bool`

HasDoctypes returns a boolean if a field has been set.

### GetInstalled

`func (o *CloudModuleState) GetInstalled() []string`

GetInstalled returns the Installed field if non-nil, zero value otherwise.

### GetInstalledOk

`func (o *CloudModuleState) GetInstalledOk() (*[]string, bool)`

GetInstalledOk returns a tuple with the Installed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstalled

`func (o *CloudModuleState) SetInstalled(v []string)`

SetInstalled sets Installed field to given value.

### HasInstalled

`func (o *CloudModuleState) HasInstalled() bool`

HasInstalled returns a boolean if a field has been set.

### GetModule

`func (o *CloudModuleState) GetModule() string`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *CloudModuleState) GetModuleOk() (*string, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *CloudModuleState) SetModule(v string)`

SetModule sets Module field to given value.

### HasModule

`func (o *CloudModuleState) HasModule() bool`

HasModule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


