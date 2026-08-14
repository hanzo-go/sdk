# O11yAzureIntegrationConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentRegion** | Pointer to **string** |  | [optional] 
**ResourceGroups** | Pointer to **[]string** |  | [optional] 
**TelemetryCollectionStrategy** | Pointer to [**[]O11yAzureTelemetryCollectionStrategy**](O11yAzureTelemetryCollectionStrategy.md) |  | [optional] 

## Methods

### NewO11yAzureIntegrationConfig

`func NewO11yAzureIntegrationConfig() *O11yAzureIntegrationConfig`

NewO11yAzureIntegrationConfig instantiates a new O11yAzureIntegrationConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yAzureIntegrationConfigWithDefaults

`func NewO11yAzureIntegrationConfigWithDefaults() *O11yAzureIntegrationConfig`

NewO11yAzureIntegrationConfigWithDefaults instantiates a new O11yAzureIntegrationConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentRegion

`func (o *O11yAzureIntegrationConfig) GetDeploymentRegion() string`

GetDeploymentRegion returns the DeploymentRegion field if non-nil, zero value otherwise.

### GetDeploymentRegionOk

`func (o *O11yAzureIntegrationConfig) GetDeploymentRegionOk() (*string, bool)`

GetDeploymentRegionOk returns a tuple with the DeploymentRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentRegion

`func (o *O11yAzureIntegrationConfig) SetDeploymentRegion(v string)`

SetDeploymentRegion sets DeploymentRegion field to given value.

### HasDeploymentRegion

`func (o *O11yAzureIntegrationConfig) HasDeploymentRegion() bool`

HasDeploymentRegion returns a boolean if a field has been set.

### GetResourceGroups

`func (o *O11yAzureIntegrationConfig) GetResourceGroups() []string`

GetResourceGroups returns the ResourceGroups field if non-nil, zero value otherwise.

### GetResourceGroupsOk

`func (o *O11yAzureIntegrationConfig) GetResourceGroupsOk() (*[]string, bool)`

GetResourceGroupsOk returns a tuple with the ResourceGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroups

`func (o *O11yAzureIntegrationConfig) SetResourceGroups(v []string)`

SetResourceGroups sets ResourceGroups field to given value.

### HasResourceGroups

`func (o *O11yAzureIntegrationConfig) HasResourceGroups() bool`

HasResourceGroups returns a boolean if a field has been set.

### GetTelemetryCollectionStrategy

`func (o *O11yAzureIntegrationConfig) GetTelemetryCollectionStrategy() []O11yAzureTelemetryCollectionStrategy`

GetTelemetryCollectionStrategy returns the TelemetryCollectionStrategy field if non-nil, zero value otherwise.

### GetTelemetryCollectionStrategyOk

`func (o *O11yAzureIntegrationConfig) GetTelemetryCollectionStrategyOk() (*[]O11yAzureTelemetryCollectionStrategy, bool)`

GetTelemetryCollectionStrategyOk returns a tuple with the TelemetryCollectionStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelemetryCollectionStrategy

`func (o *O11yAzureIntegrationConfig) SetTelemetryCollectionStrategy(v []O11yAzureTelemetryCollectionStrategy)`

SetTelemetryCollectionStrategy sets TelemetryCollectionStrategy field to given value.

### HasTelemetryCollectionStrategy

`func (o *O11yAzureIntegrationConfig) HasTelemetryCollectionStrategy() bool`

HasTelemetryCollectionStrategy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


