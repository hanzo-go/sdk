# O11yPostableAccountConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentVersion** | Pointer to **string** | as agent version is common for all providers, we can keep it at top level of this struct | [optional] 
**Aws** | Pointer to [**O11yAWSPostableAccountConfig**](O11yAWSPostableAccountConfig.md) |  | [optional] 
**Azure** | Pointer to [**O11yAzureAccountConfig**](O11yAzureAccountConfig.md) |  | [optional] 
**Gcp** | Pointer to [**O11yGCPAccountConfig**](O11yGCPAccountConfig.md) |  | [optional] 

## Methods

### NewO11yPostableAccountConfig

`func NewO11yPostableAccountConfig() *O11yPostableAccountConfig`

NewO11yPostableAccountConfig instantiates a new O11yPostableAccountConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yPostableAccountConfigWithDefaults

`func NewO11yPostableAccountConfigWithDefaults() *O11yPostableAccountConfig`

NewO11yPostableAccountConfigWithDefaults instantiates a new O11yPostableAccountConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentVersion

`func (o *O11yPostableAccountConfig) GetAgentVersion() string`

GetAgentVersion returns the AgentVersion field if non-nil, zero value otherwise.

### GetAgentVersionOk

`func (o *O11yPostableAccountConfig) GetAgentVersionOk() (*string, bool)`

GetAgentVersionOk returns a tuple with the AgentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentVersion

`func (o *O11yPostableAccountConfig) SetAgentVersion(v string)`

SetAgentVersion sets AgentVersion field to given value.

### HasAgentVersion

`func (o *O11yPostableAccountConfig) HasAgentVersion() bool`

HasAgentVersion returns a boolean if a field has been set.

### GetAws

`func (o *O11yPostableAccountConfig) GetAws() O11yAWSPostableAccountConfig`

GetAws returns the Aws field if non-nil, zero value otherwise.

### GetAwsOk

`func (o *O11yPostableAccountConfig) GetAwsOk() (*O11yAWSPostableAccountConfig, bool)`

GetAwsOk returns a tuple with the Aws field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAws

`func (o *O11yPostableAccountConfig) SetAws(v O11yAWSPostableAccountConfig)`

SetAws sets Aws field to given value.

### HasAws

`func (o *O11yPostableAccountConfig) HasAws() bool`

HasAws returns a boolean if a field has been set.

### GetAzure

`func (o *O11yPostableAccountConfig) GetAzure() O11yAzureAccountConfig`

GetAzure returns the Azure field if non-nil, zero value otherwise.

### GetAzureOk

`func (o *O11yPostableAccountConfig) GetAzureOk() (*O11yAzureAccountConfig, bool)`

GetAzureOk returns a tuple with the Azure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzure

`func (o *O11yPostableAccountConfig) SetAzure(v O11yAzureAccountConfig)`

SetAzure sets Azure field to given value.

### HasAzure

`func (o *O11yPostableAccountConfig) HasAzure() bool`

HasAzure returns a boolean if a field has been set.

### GetGcp

`func (o *O11yPostableAccountConfig) GetGcp() O11yGCPAccountConfig`

GetGcp returns the Gcp field if non-nil, zero value otherwise.

### GetGcpOk

`func (o *O11yPostableAccountConfig) GetGcpOk() (*O11yGCPAccountConfig, bool)`

GetGcpOk returns a tuple with the Gcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcp

`func (o *O11yPostableAccountConfig) SetGcp(v O11yGCPAccountConfig)`

SetGcp sets Gcp field to given value.

### HasGcp

`func (o *O11yPostableAccountConfig) HasGcp() bool`

HasGcp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


