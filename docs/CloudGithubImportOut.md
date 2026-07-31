# CloudGithubImportOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Queued** | Pointer to **int32** | Queued is how many repositories were handed to the background importer. | [optional] 
**Repos** | Pointer to **[]string** | Repos names those repositories, in the installation&#39;s listing order. | [optional] 

## Methods

### NewCloudGithubImportOut

`func NewCloudGithubImportOut() *CloudGithubImportOut`

NewCloudGithubImportOut instantiates a new CloudGithubImportOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudGithubImportOutWithDefaults

`func NewCloudGithubImportOutWithDefaults() *CloudGithubImportOut`

NewCloudGithubImportOutWithDefaults instantiates a new CloudGithubImportOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueued

`func (o *CloudGithubImportOut) GetQueued() int32`

GetQueued returns the Queued field if non-nil, zero value otherwise.

### GetQueuedOk

`func (o *CloudGithubImportOut) GetQueuedOk() (*int32, bool)`

GetQueuedOk returns a tuple with the Queued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueued

`func (o *CloudGithubImportOut) SetQueued(v int32)`

SetQueued sets Queued field to given value.

### HasQueued

`func (o *CloudGithubImportOut) HasQueued() bool`

HasQueued returns a boolean if a field has been set.

### GetRepos

`func (o *CloudGithubImportOut) GetRepos() []string`

GetRepos returns the Repos field if non-nil, zero value otherwise.

### GetReposOk

`func (o *CloudGithubImportOut) GetReposOk() (*[]string, bool)`

GetReposOk returns a tuple with the Repos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepos

`func (o *CloudGithubImportOut) SetRepos(v []string)`

SetRepos sets Repos field to given value.

### HasRepos

`func (o *CloudGithubImportOut) HasRepos() bool`

HasRepos returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


