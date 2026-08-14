# GithubImportOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Queued** | Pointer to **int32** | Queued is how many repositories were handed to the background importer. | [optional] 
**Repos** | Pointer to **[]string** | Repos names those repositories, in the installation&#39;s listing order. | [optional] 

## Methods

### NewGithubImportOut

`func NewGithubImportOut() *GithubImportOut`

NewGithubImportOut instantiates a new GithubImportOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGithubImportOutWithDefaults

`func NewGithubImportOutWithDefaults() *GithubImportOut`

NewGithubImportOutWithDefaults instantiates a new GithubImportOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueued

`func (o *GithubImportOut) GetQueued() int32`

GetQueued returns the Queued field if non-nil, zero value otherwise.

### GetQueuedOk

`func (o *GithubImportOut) GetQueuedOk() (*int32, bool)`

GetQueuedOk returns a tuple with the Queued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueued

`func (o *GithubImportOut) SetQueued(v int32)`

SetQueued sets Queued field to given value.

### HasQueued

`func (o *GithubImportOut) HasQueued() bool`

HasQueued returns a boolean if a field has been set.

### GetRepos

`func (o *GithubImportOut) GetRepos() []string`

GetRepos returns the Repos field if non-nil, zero value otherwise.

### GetReposOk

`func (o *GithubImportOut) GetReposOk() (*[]string, bool)`

GetReposOk returns a tuple with the Repos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepos

`func (o *GithubImportOut) SetRepos(v []string)`

SetRepos sets Repos field to given value.

### HasRepos

`func (o *GithubImportOut) HasRepos() bool`

HasRepos returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


