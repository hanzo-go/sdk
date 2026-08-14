# O11yConnectionArtifact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aws** | Pointer to [**O11yAWSConnectionArtifact**](O11yAWSConnectionArtifact.md) | required till new providers are added | [optional] 
**Azure** | Pointer to [**O11yAzureConnectionArtifact**](O11yAzureConnectionArtifact.md) |  | [optional] 
**Gcp** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewO11yConnectionArtifact

`func NewO11yConnectionArtifact() *O11yConnectionArtifact`

NewO11yConnectionArtifact instantiates a new O11yConnectionArtifact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yConnectionArtifactWithDefaults

`func NewO11yConnectionArtifactWithDefaults() *O11yConnectionArtifact`

NewO11yConnectionArtifactWithDefaults instantiates a new O11yConnectionArtifact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAws

`func (o *O11yConnectionArtifact) GetAws() O11yAWSConnectionArtifact`

GetAws returns the Aws field if non-nil, zero value otherwise.

### GetAwsOk

`func (o *O11yConnectionArtifact) GetAwsOk() (*O11yAWSConnectionArtifact, bool)`

GetAwsOk returns a tuple with the Aws field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAws

`func (o *O11yConnectionArtifact) SetAws(v O11yAWSConnectionArtifact)`

SetAws sets Aws field to given value.

### HasAws

`func (o *O11yConnectionArtifact) HasAws() bool`

HasAws returns a boolean if a field has been set.

### GetAzure

`func (o *O11yConnectionArtifact) GetAzure() O11yAzureConnectionArtifact`

GetAzure returns the Azure field if non-nil, zero value otherwise.

### GetAzureOk

`func (o *O11yConnectionArtifact) GetAzureOk() (*O11yAzureConnectionArtifact, bool)`

GetAzureOk returns a tuple with the Azure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzure

`func (o *O11yConnectionArtifact) SetAzure(v O11yAzureConnectionArtifact)`

SetAzure sets Azure field to given value.

### HasAzure

`func (o *O11yConnectionArtifact) HasAzure() bool`

HasAzure returns a boolean if a field has been set.

### GetGcp

`func (o *O11yConnectionArtifact) GetGcp() map[string]interface{}`

GetGcp returns the Gcp field if non-nil, zero value otherwise.

### GetGcpOk

`func (o *O11yConnectionArtifact) GetGcpOk() (*map[string]interface{}, bool)`

GetGcpOk returns a tuple with the Gcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcp

`func (o *O11yConnectionArtifact) SetGcp(v map[string]interface{})`

SetGcp sets Gcp field to given value.

### HasGcp

`func (o *O11yConnectionArtifact) HasGcp() bool`

HasGcp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


