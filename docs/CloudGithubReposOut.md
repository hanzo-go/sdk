# CloudGithubReposOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Repos** | Pointer to [**[]CloudGithubRepoView**](CloudGithubRepoView.md) | Repos is every repo the installation grants. Never null; [] when none. | [optional] 

## Methods

### NewCloudGithubReposOut

`func NewCloudGithubReposOut() *CloudGithubReposOut`

NewCloudGithubReposOut instantiates a new CloudGithubReposOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudGithubReposOutWithDefaults

`func NewCloudGithubReposOutWithDefaults() *CloudGithubReposOut`

NewCloudGithubReposOutWithDefaults instantiates a new CloudGithubReposOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepos

`func (o *CloudGithubReposOut) GetRepos() []CloudGithubRepoView`

GetRepos returns the Repos field if non-nil, zero value otherwise.

### GetReposOk

`func (o *CloudGithubReposOut) GetReposOk() (*[]CloudGithubRepoView, bool)`

GetReposOk returns a tuple with the Repos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepos

`func (o *CloudGithubReposOut) SetRepos(v []CloudGithubRepoView)`

SetRepos sets Repos field to given value.

### HasRepos

`func (o *CloudGithubReposOut) HasRepos() bool`

HasRepos returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


