# PagesProjectCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuildConfig** | Pointer to [**PagesBuildConfig**](PagesBuildConfig.md) |  | [optional] 
**DeploymentConfigs** | Pointer to [**PagesDeploymentConfigs**](PagesDeploymentConfigs.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ProductionBranch** | Pointer to **string** |  | [optional] 

## Methods

### NewPagesProjectCreate

`func NewPagesProjectCreate() *PagesProjectCreate`

NewPagesProjectCreate instantiates a new PagesProjectCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPagesProjectCreateWithDefaults

`func NewPagesProjectCreateWithDefaults() *PagesProjectCreate`

NewPagesProjectCreateWithDefaults instantiates a new PagesProjectCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuildConfig

`func (o *PagesProjectCreate) GetBuildConfig() PagesBuildConfig`

GetBuildConfig returns the BuildConfig field if non-nil, zero value otherwise.

### GetBuildConfigOk

`func (o *PagesProjectCreate) GetBuildConfigOk() (*PagesBuildConfig, bool)`

GetBuildConfigOk returns a tuple with the BuildConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildConfig

`func (o *PagesProjectCreate) SetBuildConfig(v PagesBuildConfig)`

SetBuildConfig sets BuildConfig field to given value.

### HasBuildConfig

`func (o *PagesProjectCreate) HasBuildConfig() bool`

HasBuildConfig returns a boolean if a field has been set.

### GetDeploymentConfigs

`func (o *PagesProjectCreate) GetDeploymentConfigs() PagesDeploymentConfigs`

GetDeploymentConfigs returns the DeploymentConfigs field if non-nil, zero value otherwise.

### GetDeploymentConfigsOk

`func (o *PagesProjectCreate) GetDeploymentConfigsOk() (*PagesDeploymentConfigs, bool)`

GetDeploymentConfigsOk returns a tuple with the DeploymentConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentConfigs

`func (o *PagesProjectCreate) SetDeploymentConfigs(v PagesDeploymentConfigs)`

SetDeploymentConfigs sets DeploymentConfigs field to given value.

### HasDeploymentConfigs

`func (o *PagesProjectCreate) HasDeploymentConfigs() bool`

HasDeploymentConfigs returns a boolean if a field has been set.

### GetName

`func (o *PagesProjectCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PagesProjectCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PagesProjectCreate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PagesProjectCreate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProductionBranch

`func (o *PagesProjectCreate) GetProductionBranch() string`

GetProductionBranch returns the ProductionBranch field if non-nil, zero value otherwise.

### GetProductionBranchOk

`func (o *PagesProjectCreate) GetProductionBranchOk() (*string, bool)`

GetProductionBranchOk returns a tuple with the ProductionBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductionBranch

`func (o *PagesProjectCreate) SetProductionBranch(v string)`

SetProductionBranch sets ProductionBranch field to given value.

### HasProductionBranch

`func (o *PagesProjectCreate) HasProductionBranch() bool`

HasProductionBranch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


