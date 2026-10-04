# GitRefsJSON

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Branches** | Pointer to [**[]GitRefJSON**](GitRefJSON.md) | Branches are the repo&#39;s heads; empty on a repo with no commits. | [optional] 
**Default** | Pointer to **string** | Default is the branch name a caller gets when it asks for no ref. | [optional] 
**Tags** | Pointer to [**[]GitRefJSON**](GitRefJSON.md) | Tags are the repo&#39;s tags; empty when there are none. | [optional] 

## Methods

### NewGitRefsJSON

`func NewGitRefsJSON() *GitRefsJSON`

NewGitRefsJSON instantiates a new GitRefsJSON object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitRefsJSONWithDefaults

`func NewGitRefsJSONWithDefaults() *GitRefsJSON`

NewGitRefsJSONWithDefaults instantiates a new GitRefsJSON object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBranches

`func (o *GitRefsJSON) GetBranches() []GitRefJSON`

GetBranches returns the Branches field if non-nil, zero value otherwise.

### GetBranchesOk

`func (o *GitRefsJSON) GetBranchesOk() (*[]GitRefJSON, bool)`

GetBranchesOk returns a tuple with the Branches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranches

`func (o *GitRefsJSON) SetBranches(v []GitRefJSON)`

SetBranches sets Branches field to given value.

### HasBranches

`func (o *GitRefsJSON) HasBranches() bool`

HasBranches returns a boolean if a field has been set.

### GetDefault

`func (o *GitRefsJSON) GetDefault() string`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *GitRefsJSON) GetDefaultOk() (*string, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *GitRefsJSON) SetDefault(v string)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *GitRefsJSON) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetTags

`func (o *GitRefsJSON) GetTags() []GitRefJSON`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GitRefsJSON) GetTagsOk() (*[]GitRefJSON, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GitRefsJSON) SetTags(v []GitRefJSON)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GitRefsJSON) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


