# O11yProviderIntegrationConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aws** | Pointer to [**O11yAWSIntegrationConfig**](O11yAWSIntegrationConfig.md) |  | [optional] 
**Azure** | Pointer to [**O11yAzureIntegrationConfig**](O11yAzureIntegrationConfig.md) |  | [optional] 
**Gcp** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewO11yProviderIntegrationConfig

`func NewO11yProviderIntegrationConfig() *O11yProviderIntegrationConfig`

NewO11yProviderIntegrationConfig instantiates a new O11yProviderIntegrationConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yProviderIntegrationConfigWithDefaults

`func NewO11yProviderIntegrationConfigWithDefaults() *O11yProviderIntegrationConfig`

NewO11yProviderIntegrationConfigWithDefaults instantiates a new O11yProviderIntegrationConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAws

`func (o *O11yProviderIntegrationConfig) GetAws() O11yAWSIntegrationConfig`

GetAws returns the Aws field if non-nil, zero value otherwise.

### GetAwsOk

`func (o *O11yProviderIntegrationConfig) GetAwsOk() (*O11yAWSIntegrationConfig, bool)`

GetAwsOk returns a tuple with the Aws field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAws

`func (o *O11yProviderIntegrationConfig) SetAws(v O11yAWSIntegrationConfig)`

SetAws sets Aws field to given value.

### HasAws

`func (o *O11yProviderIntegrationConfig) HasAws() bool`

HasAws returns a boolean if a field has been set.

### GetAzure

`func (o *O11yProviderIntegrationConfig) GetAzure() O11yAzureIntegrationConfig`

GetAzure returns the Azure field if non-nil, zero value otherwise.

### GetAzureOk

`func (o *O11yProviderIntegrationConfig) GetAzureOk() (*O11yAzureIntegrationConfig, bool)`

GetAzureOk returns a tuple with the Azure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzure

`func (o *O11yProviderIntegrationConfig) SetAzure(v O11yAzureIntegrationConfig)`

SetAzure sets Azure field to given value.

### HasAzure

`func (o *O11yProviderIntegrationConfig) HasAzure() bool`

HasAzure returns a boolean if a field has been set.

### GetGcp

`func (o *O11yProviderIntegrationConfig) GetGcp() map[string]interface{}`

GetGcp returns the Gcp field if non-nil, zero value otherwise.

### GetGcpOk

`func (o *O11yProviderIntegrationConfig) GetGcpOk() (*map[string]interface{}, bool)`

GetGcpOk returns a tuple with the Gcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcp

`func (o *O11yProviderIntegrationConfig) SetGcp(v map[string]interface{})`

SetGcp sets Gcp field to given value.

### HasGcp

`func (o *O11yProviderIntegrationConfig) HasGcp() bool`

HasGcp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


