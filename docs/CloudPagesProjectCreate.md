# CloudPagesProjectCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuildConfig** | Pointer to [**CloudPagesBuildConfig**](CloudPagesBuildConfig.md) |  | [optional] 
**DeploymentConfigs** | Pointer to [**CloudPagesDeploymentConfigs**](CloudPagesDeploymentConfigs.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ProductionBranch** | Pointer to **string** |  | [optional] 

## Methods

### NewCloudPagesProjectCreate

`func NewCloudPagesProjectCreate() *CloudPagesProjectCreate`

NewCloudPagesProjectCreate instantiates a new CloudPagesProjectCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudPagesProjectCreateWithDefaults

`func NewCloudPagesProjectCreateWithDefaults() *CloudPagesProjectCreate`

NewCloudPagesProjectCreateWithDefaults instantiates a new CloudPagesProjectCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuildConfig

`func (o *CloudPagesProjectCreate) GetBuildConfig() CloudPagesBuildConfig`

GetBuildConfig returns the BuildConfig field if non-nil, zero value otherwise.

### GetBuildConfigOk

`func (o *CloudPagesProjectCreate) GetBuildConfigOk() (*CloudPagesBuildConfig, bool)`

GetBuildConfigOk returns a tuple with the BuildConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildConfig

`func (o *CloudPagesProjectCreate) SetBuildConfig(v CloudPagesBuildConfig)`

SetBuildConfig sets BuildConfig field to given value.

### HasBuildConfig

`func (o *CloudPagesProjectCreate) HasBuildConfig() bool`

HasBuildConfig returns a boolean if a field has been set.

### GetDeploymentConfigs

`func (o *CloudPagesProjectCreate) GetDeploymentConfigs() CloudPagesDeploymentConfigs`

GetDeploymentConfigs returns the DeploymentConfigs field if non-nil, zero value otherwise.

### GetDeploymentConfigsOk

`func (o *CloudPagesProjectCreate) GetDeploymentConfigsOk() (*CloudPagesDeploymentConfigs, bool)`

GetDeploymentConfigsOk returns a tuple with the DeploymentConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentConfigs

`func (o *CloudPagesProjectCreate) SetDeploymentConfigs(v CloudPagesDeploymentConfigs)`

SetDeploymentConfigs sets DeploymentConfigs field to given value.

### HasDeploymentConfigs

`func (o *CloudPagesProjectCreate) HasDeploymentConfigs() bool`

HasDeploymentConfigs returns a boolean if a field has been set.

### GetName

`func (o *CloudPagesProjectCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudPagesProjectCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudPagesProjectCreate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudPagesProjectCreate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProductionBranch

`func (o *CloudPagesProjectCreate) GetProductionBranch() string`

GetProductionBranch returns the ProductionBranch field if non-nil, zero value otherwise.

### GetProductionBranchOk

`func (o *CloudPagesProjectCreate) GetProductionBranchOk() (*string, bool)`

GetProductionBranchOk returns a tuple with the ProductionBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductionBranch

`func (o *CloudPagesProjectCreate) SetProductionBranch(v string)`

SetProductionBranch sets ProductionBranch field to given value.

### HasProductionBranch

`func (o *CloudPagesProjectCreate) HasProductionBranch() bool`

HasProductionBranch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


