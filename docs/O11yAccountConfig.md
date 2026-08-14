# O11yAccountConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aws** | Pointer to [**O11yAWSAccountConfig**](O11yAWSAccountConfig.md) |  | [optional] 
**Azure** | Pointer to [**O11yAzureAccountConfig**](O11yAzureAccountConfig.md) |  | [optional] 
**Gcp** | Pointer to [**O11yGCPAccountConfig**](O11yGCPAccountConfig.md) |  | [optional] 

## Methods

### NewO11yAccountConfig

`func NewO11yAccountConfig() *O11yAccountConfig`

NewO11yAccountConfig instantiates a new O11yAccountConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yAccountConfigWithDefaults

`func NewO11yAccountConfigWithDefaults() *O11yAccountConfig`

NewO11yAccountConfigWithDefaults instantiates a new O11yAccountConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAws

`func (o *O11yAccountConfig) GetAws() O11yAWSAccountConfig`

GetAws returns the Aws field if non-nil, zero value otherwise.

### GetAwsOk

`func (o *O11yAccountConfig) GetAwsOk() (*O11yAWSAccountConfig, bool)`

GetAwsOk returns a tuple with the Aws field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAws

`func (o *O11yAccountConfig) SetAws(v O11yAWSAccountConfig)`

SetAws sets Aws field to given value.

### HasAws

`func (o *O11yAccountConfig) HasAws() bool`

HasAws returns a boolean if a field has been set.

### GetAzure

`func (o *O11yAccountConfig) GetAzure() O11yAzureAccountConfig`

GetAzure returns the Azure field if non-nil, zero value otherwise.

### GetAzureOk

`func (o *O11yAccountConfig) GetAzureOk() (*O11yAzureAccountConfig, bool)`

GetAzureOk returns a tuple with the Azure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzure

`func (o *O11yAccountConfig) SetAzure(v O11yAzureAccountConfig)`

SetAzure sets Azure field to given value.

### HasAzure

`func (o *O11yAccountConfig) HasAzure() bool`

HasAzure returns a boolean if a field has been set.

### GetGcp

`func (o *O11yAccountConfig) GetGcp() O11yGCPAccountConfig`

GetGcp returns the Gcp field if non-nil, zero value otherwise.

### GetGcpOk

`func (o *O11yAccountConfig) GetGcpOk() (*O11yGCPAccountConfig, bool)`

GetGcpOk returns a tuple with the Gcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcp

`func (o *O11yAccountConfig) SetGcp(v O11yGCPAccountConfig)`

SetGcp sets Gcp field to given value.

### HasGcp

`func (o *O11yAccountConfig) HasGcp() bool`

HasGcp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


