# CloudflarePagesDeploymentConfigs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Preview** | Pointer to [**CloudflarePagesDeploymentConfig**](CloudflarePagesDeploymentConfig.md) | Preview is the config every branch build other than the production branch runs under. It is a SEPARATE set of bindings and variables, which is what lets a preview point at test data. | [optional] 
**Production** | Pointer to [**CloudflarePagesDeploymentConfig**](CloudflarePagesDeploymentConfig.md) | Production is the config the production branch builds under. | [optional] 

## Methods

### NewCloudflarePagesDeploymentConfigs

`func NewCloudflarePagesDeploymentConfigs() *CloudflarePagesDeploymentConfigs`

NewCloudflarePagesDeploymentConfigs instantiates a new CloudflarePagesDeploymentConfigs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudflarePagesDeploymentConfigsWithDefaults

`func NewCloudflarePagesDeploymentConfigsWithDefaults() *CloudflarePagesDeploymentConfigs`

NewCloudflarePagesDeploymentConfigsWithDefaults instantiates a new CloudflarePagesDeploymentConfigs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPreview

`func (o *CloudflarePagesDeploymentConfigs) GetPreview() CloudflarePagesDeploymentConfig`

GetPreview returns the Preview field if non-nil, zero value otherwise.

### GetPreviewOk

`func (o *CloudflarePagesDeploymentConfigs) GetPreviewOk() (*CloudflarePagesDeploymentConfig, bool)`

GetPreviewOk returns a tuple with the Preview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreview

`func (o *CloudflarePagesDeploymentConfigs) SetPreview(v CloudflarePagesDeploymentConfig)`

SetPreview sets Preview field to given value.

### HasPreview

`func (o *CloudflarePagesDeploymentConfigs) HasPreview() bool`

HasPreview returns a boolean if a field has been set.

### GetProduction

`func (o *CloudflarePagesDeploymentConfigs) GetProduction() CloudflarePagesDeploymentConfig`

GetProduction returns the Production field if non-nil, zero value otherwise.

### GetProductionOk

`func (o *CloudflarePagesDeploymentConfigs) GetProductionOk() (*CloudflarePagesDeploymentConfig, bool)`

GetProductionOk returns a tuple with the Production field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduction

`func (o *CloudflarePagesDeploymentConfigs) SetProduction(v CloudflarePagesDeploymentConfig)`

SetProduction sets Production field to given value.

### HasProduction

`func (o *CloudflarePagesDeploymentConfigs) HasProduction() bool`

HasProduction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


