# O11yServiceConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aws** | Pointer to [**O11yAWSServiceConfig**](O11yAWSServiceConfig.md) |  | [optional] 
**Azure** | Pointer to [**O11yAzureServiceConfig**](O11yAzureServiceConfig.md) |  | [optional] 
**Gcp** | Pointer to [**O11yGCPServiceConfig**](O11yGCPServiceConfig.md) |  | [optional] 

## Methods

### NewO11yServiceConfig

`func NewO11yServiceConfig() *O11yServiceConfig`

NewO11yServiceConfig instantiates a new O11yServiceConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yServiceConfigWithDefaults

`func NewO11yServiceConfigWithDefaults() *O11yServiceConfig`

NewO11yServiceConfigWithDefaults instantiates a new O11yServiceConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAws

`func (o *O11yServiceConfig) GetAws() O11yAWSServiceConfig`

GetAws returns the Aws field if non-nil, zero value otherwise.

### GetAwsOk

`func (o *O11yServiceConfig) GetAwsOk() (*O11yAWSServiceConfig, bool)`

GetAwsOk returns a tuple with the Aws field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAws

`func (o *O11yServiceConfig) SetAws(v O11yAWSServiceConfig)`

SetAws sets Aws field to given value.

### HasAws

`func (o *O11yServiceConfig) HasAws() bool`

HasAws returns a boolean if a field has been set.

### GetAzure

`func (o *O11yServiceConfig) GetAzure() O11yAzureServiceConfig`

GetAzure returns the Azure field if non-nil, zero value otherwise.

### GetAzureOk

`func (o *O11yServiceConfig) GetAzureOk() (*O11yAzureServiceConfig, bool)`

GetAzureOk returns a tuple with the Azure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzure

`func (o *O11yServiceConfig) SetAzure(v O11yAzureServiceConfig)`

SetAzure sets Azure field to given value.

### HasAzure

`func (o *O11yServiceConfig) HasAzure() bool`

HasAzure returns a boolean if a field has been set.

### GetGcp

`func (o *O11yServiceConfig) GetGcp() O11yGCPServiceConfig`

GetGcp returns the Gcp field if non-nil, zero value otherwise.

### GetGcpOk

`func (o *O11yServiceConfig) GetGcpOk() (*O11yGCPServiceConfig, bool)`

GetGcpOk returns a tuple with the Gcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcp

`func (o *O11yServiceConfig) SetGcp(v O11yGCPServiceConfig)`

SetGcp sets Gcp field to given value.

### HasGcp

`func (o *O11yServiceConfig) HasGcp() bool`

HasGcp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


