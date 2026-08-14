# InstallState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Installed** | Pointer to **bool** | Installed is its activation after the write. | [optional] 
**Tool** | Pointer to **string** | Tool is the capability the write applied to. | [optional] 

## Methods

### NewInstallState

`func NewInstallState() *InstallState`

NewInstallState instantiates a new InstallState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstallStateWithDefaults

`func NewInstallStateWithDefaults() *InstallState`

NewInstallStateWithDefaults instantiates a new InstallState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstalled

`func (o *InstallState) GetInstalled() bool`

GetInstalled returns the Installed field if non-nil, zero value otherwise.

### GetInstalledOk

`func (o *InstallState) GetInstalledOk() (*bool, bool)`

GetInstalledOk returns a tuple with the Installed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstalled

`func (o *InstallState) SetInstalled(v bool)`

SetInstalled sets Installed field to given value.

### HasInstalled

`func (o *InstallState) HasInstalled() bool`

HasInstalled returns a boolean if a field has been set.

### GetTool

`func (o *InstallState) GetTool() string`

GetTool returns the Tool field if non-nil, zero value otherwise.

### GetToolOk

`func (o *InstallState) GetToolOk() (*string, bool)`

GetToolOk returns a tuple with the Tool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTool

`func (o *InstallState) SetTool(v string)`

SetTool sets Tool field to given value.

### HasTool

`func (o *InstallState) HasTool() bool`

HasTool returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


