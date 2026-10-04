# GitPushResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Branch** | Pointer to **string** | Branch is the branch that was advanced, resolved (never empty). | [optional] 
**CloneUrl** | Pointer to **string** | CloneURL is the repo&#39;s HTTPS remote. | [optional] 
**Commit** | Pointer to **string** | Commit is the new commit&#39;s full hash. | [optional] 
**SshUrl** | Pointer to **string** | SSHURL is the repo&#39;s scp-style SSH remote. | [optional] 

## Methods

### NewGitPushResp

`func NewGitPushResp() *GitPushResp`

NewGitPushResp instantiates a new GitPushResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitPushRespWithDefaults

`func NewGitPushRespWithDefaults() *GitPushResp`

NewGitPushRespWithDefaults instantiates a new GitPushResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBranch

`func (o *GitPushResp) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *GitPushResp) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *GitPushResp) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *GitPushResp) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetCloneUrl

`func (o *GitPushResp) GetCloneUrl() string`

GetCloneUrl returns the CloneUrl field if non-nil, zero value otherwise.

### GetCloneUrlOk

`func (o *GitPushResp) GetCloneUrlOk() (*string, bool)`

GetCloneUrlOk returns a tuple with the CloneUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloneUrl

`func (o *GitPushResp) SetCloneUrl(v string)`

SetCloneUrl sets CloneUrl field to given value.

### HasCloneUrl

`func (o *GitPushResp) HasCloneUrl() bool`

HasCloneUrl returns a boolean if a field has been set.

### GetCommit

`func (o *GitPushResp) GetCommit() string`

GetCommit returns the Commit field if non-nil, zero value otherwise.

### GetCommitOk

`func (o *GitPushResp) GetCommitOk() (*string, bool)`

GetCommitOk returns a tuple with the Commit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommit

`func (o *GitPushResp) SetCommit(v string)`

SetCommit sets Commit field to given value.

### HasCommit

`func (o *GitPushResp) HasCommit() bool`

HasCommit returns a boolean if a field has been set.

### GetSshUrl

`func (o *GitPushResp) GetSshUrl() string`

GetSshUrl returns the SshUrl field if non-nil, zero value otherwise.

### GetSshUrlOk

`func (o *GitPushResp) GetSshUrlOk() (*string, bool)`

GetSshUrlOk returns a tuple with the SshUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshUrl

`func (o *GitPushResp) SetSshUrl(v string)`

SetSshUrl sets SshUrl field to given value.

### HasSshUrl

`func (o *GitPushResp) HasSshUrl() bool`

HasSshUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


