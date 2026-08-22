# PagesProjectCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuildConfig** | Pointer to [**PagesBuildConfig**](PagesBuildConfig.md) | BuildConfig says how to build the site. Omitted means no build step. | [optional] 
**DeploymentConfigs** | Pointer to [**PagesDeploymentConfigs**](PagesDeploymentConfigs.md) | DeploymentConfigs carries the preview and production runtime configs — the bindings and variables the built site&#39;s functions run with. | [optional] 
**Name** | Pointer to **string** | Name is the project name, and it is also the address: the site answers at &lt;name&gt;.pages.dev. Cloudflare will not rename a project afterwards. | [optional] 
**ProductionBranch** | Pointer to **string** | ProductionBranch is which git branch builds to production; every other branch builds a preview. Omitted leaves Cloudflare&#39;s own default. | [optional] 

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


