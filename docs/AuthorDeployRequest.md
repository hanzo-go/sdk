# AuthorDeployRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Project** | Pointer to **string** | Project is the deployed project&#39;s id. Required. | [optional] 
**RepoUrl** | Pointer to **string** | RepoURL is the source repository the project was built from. Empty means a hand-built project with nothing to attribute — an honest no-op, not an error. | [optional] 

## Methods

### NewAuthorDeployRequest

`func NewAuthorDeployRequest() *AuthorDeployRequest`

NewAuthorDeployRequest instantiates a new AuthorDeployRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorDeployRequestWithDefaults

`func NewAuthorDeployRequestWithDefaults() *AuthorDeployRequest`

NewAuthorDeployRequestWithDefaults instantiates a new AuthorDeployRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProject

`func (o *AuthorDeployRequest) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *AuthorDeployRequest) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *AuthorDeployRequest) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *AuthorDeployRequest) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetRepoUrl

`func (o *AuthorDeployRequest) GetRepoUrl() string`

GetRepoUrl returns the RepoUrl field if non-nil, zero value otherwise.

### GetRepoUrlOk

`func (o *AuthorDeployRequest) GetRepoUrlOk() (*string, bool)`

GetRepoUrlOk returns a tuple with the RepoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoUrl

`func (o *AuthorDeployRequest) SetRepoUrl(v string)`

SetRepoUrl sets RepoUrl field to given value.

### HasRepoUrl

`func (o *AuthorDeployRequest) HasRepoUrl() bool`

HasRepoUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


