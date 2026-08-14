# O11yGCPAccountConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentProjectId** | Pointer to **string** | Project ID where central pub/sub for logs exist | [optional] 
**DeploymentRegion** | Pointer to **string** | Project ID where otel collector will be deployed | [optional] 
**ProjectIds** | Pointer to **[]string** | List of project IDs to monitor | [optional] 

## Methods

### NewO11yGCPAccountConfig

`func NewO11yGCPAccountConfig() *O11yGCPAccountConfig`

NewO11yGCPAccountConfig instantiates a new O11yGCPAccountConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yGCPAccountConfigWithDefaults

`func NewO11yGCPAccountConfigWithDefaults() *O11yGCPAccountConfig`

NewO11yGCPAccountConfigWithDefaults instantiates a new O11yGCPAccountConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentProjectId

`func (o *O11yGCPAccountConfig) GetDeploymentProjectId() string`

GetDeploymentProjectId returns the DeploymentProjectId field if non-nil, zero value otherwise.

### GetDeploymentProjectIdOk

`func (o *O11yGCPAccountConfig) GetDeploymentProjectIdOk() (*string, bool)`

GetDeploymentProjectIdOk returns a tuple with the DeploymentProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentProjectId

`func (o *O11yGCPAccountConfig) SetDeploymentProjectId(v string)`

SetDeploymentProjectId sets DeploymentProjectId field to given value.

### HasDeploymentProjectId

`func (o *O11yGCPAccountConfig) HasDeploymentProjectId() bool`

HasDeploymentProjectId returns a boolean if a field has been set.

### GetDeploymentRegion

`func (o *O11yGCPAccountConfig) GetDeploymentRegion() string`

GetDeploymentRegion returns the DeploymentRegion field if non-nil, zero value otherwise.

### GetDeploymentRegionOk

`func (o *O11yGCPAccountConfig) GetDeploymentRegionOk() (*string, bool)`

GetDeploymentRegionOk returns a tuple with the DeploymentRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentRegion

`func (o *O11yGCPAccountConfig) SetDeploymentRegion(v string)`

SetDeploymentRegion sets DeploymentRegion field to given value.

### HasDeploymentRegion

`func (o *O11yGCPAccountConfig) HasDeploymentRegion() bool`

HasDeploymentRegion returns a boolean if a field has been set.

### GetProjectIds

`func (o *O11yGCPAccountConfig) GetProjectIds() []string`

GetProjectIds returns the ProjectIds field if non-nil, zero value otherwise.

### GetProjectIdsOk

`func (o *O11yGCPAccountConfig) GetProjectIdsOk() (*[]string, bool)`

GetProjectIdsOk returns a tuple with the ProjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectIds

`func (o *O11yGCPAccountConfig) SetProjectIds(v []string)`

SetProjectIds sets ProjectIds field to given value.

### HasProjectIds

`func (o *O11yGCPAccountConfig) HasProjectIds() bool`

HasProjectIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


