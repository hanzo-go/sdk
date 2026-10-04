# EnvironmentEnvSave

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Install** | Pointer to **string** | Install is run in every fresh checkout before the agent starts, from the repository root, and must be non-interactive. Empty installs nothing. | [optional] 
**Repo** | Pointer to **string** | Repo is the repository&#39;s name in the caller&#39;s org. | [optional] 
**Start** | Pointer to **string** | Start is left running in the background after the install. Empty starts nothing. | [optional] 

## Methods

### NewEnvironmentEnvSave

`func NewEnvironmentEnvSave() *EnvironmentEnvSave`

NewEnvironmentEnvSave instantiates a new EnvironmentEnvSave object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentEnvSaveWithDefaults

`func NewEnvironmentEnvSaveWithDefaults() *EnvironmentEnvSave`

NewEnvironmentEnvSaveWithDefaults instantiates a new EnvironmentEnvSave object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstall

`func (o *EnvironmentEnvSave) GetInstall() string`

GetInstall returns the Install field if non-nil, zero value otherwise.

### GetInstallOk

`func (o *EnvironmentEnvSave) GetInstallOk() (*string, bool)`

GetInstallOk returns a tuple with the Install field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstall

`func (o *EnvironmentEnvSave) SetInstall(v string)`

SetInstall sets Install field to given value.

### HasInstall

`func (o *EnvironmentEnvSave) HasInstall() bool`

HasInstall returns a boolean if a field has been set.

### GetRepo

`func (o *EnvironmentEnvSave) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *EnvironmentEnvSave) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *EnvironmentEnvSave) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *EnvironmentEnvSave) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### GetStart

`func (o *EnvironmentEnvSave) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *EnvironmentEnvSave) GetStartOk() (*string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *EnvironmentEnvSave) SetStart(v string)`

SetStart sets Start field to given value.

### HasStart

`func (o *EnvironmentEnvSave) HasStart() bool`

HasStart returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


