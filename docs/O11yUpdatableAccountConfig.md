# O11yUpdatableAccountConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aws** | Pointer to [**O11yAWSAccountConfig**](O11yAWSAccountConfig.md) |  | [optional] 
**Azure** | Pointer to [**O11yUpdatableAzureAccountConfig**](O11yUpdatableAzureAccountConfig.md) |  | [optional] 
**Gcp** | Pointer to [**O11yUpdatableGCPAccountConfig**](O11yUpdatableGCPAccountConfig.md) |  | [optional] 

## Methods

### NewO11yUpdatableAccountConfig

`func NewO11yUpdatableAccountConfig() *O11yUpdatableAccountConfig`

NewO11yUpdatableAccountConfig instantiates a new O11yUpdatableAccountConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yUpdatableAccountConfigWithDefaults

`func NewO11yUpdatableAccountConfigWithDefaults() *O11yUpdatableAccountConfig`

NewO11yUpdatableAccountConfigWithDefaults instantiates a new O11yUpdatableAccountConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAws

`func (o *O11yUpdatableAccountConfig) GetAws() O11yAWSAccountConfig`

GetAws returns the Aws field if non-nil, zero value otherwise.

### GetAwsOk

`func (o *O11yUpdatableAccountConfig) GetAwsOk() (*O11yAWSAccountConfig, bool)`

GetAwsOk returns a tuple with the Aws field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAws

`func (o *O11yUpdatableAccountConfig) SetAws(v O11yAWSAccountConfig)`

SetAws sets Aws field to given value.

### HasAws

`func (o *O11yUpdatableAccountConfig) HasAws() bool`

HasAws returns a boolean if a field has been set.

### GetAzure

`func (o *O11yUpdatableAccountConfig) GetAzure() O11yUpdatableAzureAccountConfig`

GetAzure returns the Azure field if non-nil, zero value otherwise.

### GetAzureOk

`func (o *O11yUpdatableAccountConfig) GetAzureOk() (*O11yUpdatableAzureAccountConfig, bool)`

GetAzureOk returns a tuple with the Azure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzure

`func (o *O11yUpdatableAccountConfig) SetAzure(v O11yUpdatableAzureAccountConfig)`

SetAzure sets Azure field to given value.

### HasAzure

`func (o *O11yUpdatableAccountConfig) HasAzure() bool`

HasAzure returns a boolean if a field has been set.

### GetGcp

`func (o *O11yUpdatableAccountConfig) GetGcp() O11yUpdatableGCPAccountConfig`

GetGcp returns the Gcp field if non-nil, zero value otherwise.

### GetGcpOk

`func (o *O11yUpdatableAccountConfig) GetGcpOk() (*O11yUpdatableGCPAccountConfig, bool)`

GetGcpOk returns a tuple with the Gcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcp

`func (o *O11yUpdatableAccountConfig) SetGcp(v O11yUpdatableGCPAccountConfig)`

SetGcp sets Gcp field to given value.

### HasGcp

`func (o *O11yUpdatableAccountConfig) HasGcp() bool`

HasGcp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


