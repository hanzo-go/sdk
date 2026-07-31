# GitUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Org** | Pointer to **string** |  | [optional] 
**TotalBytes** | Pointer to **int64** |  | [optional] 
**Repos** | Pointer to [**[]GitUsageRepo**](GitUsageRepo.md) |  | [optional] 

## Methods

### NewGitUsage

`func NewGitUsage() *GitUsage`

NewGitUsage instantiates a new GitUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitUsageWithDefaults

`func NewGitUsageWithDefaults() *GitUsage`

NewGitUsageWithDefaults instantiates a new GitUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrg

`func (o *GitUsage) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *GitUsage) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *GitUsage) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *GitUsage) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetTotalBytes

`func (o *GitUsage) GetTotalBytes() int64`

GetTotalBytes returns the TotalBytes field if non-nil, zero value otherwise.

### GetTotalBytesOk

`func (o *GitUsage) GetTotalBytesOk() (*int64, bool)`

GetTotalBytesOk returns a tuple with the TotalBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBytes

`func (o *GitUsage) SetTotalBytes(v int64)`

SetTotalBytes sets TotalBytes field to given value.

### HasTotalBytes

`func (o *GitUsage) HasTotalBytes() bool`

HasTotalBytes returns a boolean if a field has been set.

### GetRepos

`func (o *GitUsage) GetRepos() []GitUsageRepo`

GetRepos returns the Repos field if non-nil, zero value otherwise.

### GetReposOk

`func (o *GitUsage) GetReposOk() (*[]GitUsageRepo, bool)`

GetReposOk returns a tuple with the Repos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepos

`func (o *GitUsage) SetRepos(v []GitUsageRepo)`

SetRepos sets Repos field to given value.

### HasRepos

`func (o *GitUsage) HasRepos() bool`

HasRepos returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


