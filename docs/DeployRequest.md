# DeployRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Project** | Pointer to **string** | Project is the deployed project&#39;s id. Required. | [optional] 
**RepoUrl** | Pointer to **string** | RepoURL is the source repository the project was built from. Empty means a hand-built project with nothing to attribute — an honest no-op, not an error. | [optional] 

## Methods

### NewDeployRequest

`func NewDeployRequest() *DeployRequest`

NewDeployRequest instantiates a new DeployRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeployRequestWithDefaults

`func NewDeployRequestWithDefaults() *DeployRequest`

NewDeployRequestWithDefaults instantiates a new DeployRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProject

`func (o *DeployRequest) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *DeployRequest) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *DeployRequest) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *DeployRequest) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetRepoUrl

`func (o *DeployRequest) GetRepoUrl() string`

GetRepoUrl returns the RepoUrl field if non-nil, zero value otherwise.

### GetRepoUrlOk

`func (o *DeployRequest) GetRepoUrlOk() (*string, bool)`

GetRepoUrlOk returns a tuple with the RepoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoUrl

`func (o *DeployRequest) SetRepoUrl(v string)`

SetRepoUrl sets RepoUrl field to given value.

### HasRepoUrl

`func (o *DeployRequest) HasRepoUrl() bool`

HasRepoUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


