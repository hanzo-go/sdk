# GitCommitsJSON

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Commits** | Pointer to [**[]GitCommitJSON**](GitCommitJSON.md) | Commits are newest first. | [optional] 

## Methods

### NewGitCommitsJSON

`func NewGitCommitsJSON() *GitCommitsJSON`

NewGitCommitsJSON instantiates a new GitCommitsJSON object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitCommitsJSONWithDefaults

`func NewGitCommitsJSONWithDefaults() *GitCommitsJSON`

NewGitCommitsJSONWithDefaults instantiates a new GitCommitsJSON object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommits

`func (o *GitCommitsJSON) GetCommits() []GitCommitJSON`

GetCommits returns the Commits field if non-nil, zero value otherwise.

### GetCommitsOk

`func (o *GitCommitsJSON) GetCommitsOk() (*[]GitCommitJSON, bool)`

GetCommitsOk returns a tuple with the Commits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommits

`func (o *GitCommitsJSON) SetCommits(v []GitCommitJSON)`

SetCommits sets Commits field to given value.

### HasCommits

`func (o *GitCommitsJSON) HasCommits() bool`

HasCommits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


