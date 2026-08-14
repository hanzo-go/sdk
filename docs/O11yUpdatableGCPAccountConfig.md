# O11yUpdatableGCPAccountConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentProjectId** | Pointer to **string** | Project ID where central pub/sub for logs exist | [optional] 
**DeploymentRegion** | Pointer to **string** | Compute service region where otel collector will be deployed | [optional] 
**ProjectIds** | Pointer to **[]string** | List of project IDs to monitor | [optional] 

## Methods

### NewO11yUpdatableGCPAccountConfig

`func NewO11yUpdatableGCPAccountConfig() *O11yUpdatableGCPAccountConfig`

NewO11yUpdatableGCPAccountConfig instantiates a new O11yUpdatableGCPAccountConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yUpdatableGCPAccountConfigWithDefaults

`func NewO11yUpdatableGCPAccountConfigWithDefaults() *O11yUpdatableGCPAccountConfig`

NewO11yUpdatableGCPAccountConfigWithDefaults instantiates a new O11yUpdatableGCPAccountConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentProjectId

`func (o *O11yUpdatableGCPAccountConfig) GetDeploymentProjectId() string`

GetDeploymentProjectId returns the DeploymentProjectId field if non-nil, zero value otherwise.

### GetDeploymentProjectIdOk

`func (o *O11yUpdatableGCPAccountConfig) GetDeploymentProjectIdOk() (*string, bool)`

GetDeploymentProjectIdOk returns a tuple with the DeploymentProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentProjectId

`func (o *O11yUpdatableGCPAccountConfig) SetDeploymentProjectId(v string)`

SetDeploymentProjectId sets DeploymentProjectId field to given value.

### HasDeploymentProjectId

`func (o *O11yUpdatableGCPAccountConfig) HasDeploymentProjectId() bool`

HasDeploymentProjectId returns a boolean if a field has been set.

### GetDeploymentRegion

`func (o *O11yUpdatableGCPAccountConfig) GetDeploymentRegion() string`

GetDeploymentRegion returns the DeploymentRegion field if non-nil, zero value otherwise.

### GetDeploymentRegionOk

`func (o *O11yUpdatableGCPAccountConfig) GetDeploymentRegionOk() (*string, bool)`

GetDeploymentRegionOk returns a tuple with the DeploymentRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentRegion

`func (o *O11yUpdatableGCPAccountConfig) SetDeploymentRegion(v string)`

SetDeploymentRegion sets DeploymentRegion field to given value.

### HasDeploymentRegion

`func (o *O11yUpdatableGCPAccountConfig) HasDeploymentRegion() bool`

HasDeploymentRegion returns a boolean if a field has been set.

### GetProjectIds

`func (o *O11yUpdatableGCPAccountConfig) GetProjectIds() []string`

GetProjectIds returns the ProjectIds field if non-nil, zero value otherwise.

### GetProjectIdsOk

`func (o *O11yUpdatableGCPAccountConfig) GetProjectIdsOk() (*[]string, bool)`

GetProjectIdsOk returns a tuple with the ProjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectIds

`func (o *O11yUpdatableGCPAccountConfig) SetProjectIds(v []string)`

SetProjectIds sets ProjectIds field to given value.

### HasProjectIds

`func (o *O11yUpdatableGCPAccountConfig) HasProjectIds() bool`

HasProjectIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


