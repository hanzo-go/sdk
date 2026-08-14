# PushResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Branch** | Pointer to **string** | Branch is the branch that was advanced, resolved (never empty). | [optional] 
**CloneUrl** | Pointer to **string** | CloneURL is the repo&#39;s HTTPS remote. | [optional] 
**Commit** | Pointer to **string** | Commit is the new commit&#39;s full hash. | [optional] 
**SshUrl** | Pointer to **string** | SSHURL is the repo&#39;s scp-style SSH remote. | [optional] 

## Methods

### NewPushResp

`func NewPushResp() *PushResp`

NewPushResp instantiates a new PushResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPushRespWithDefaults

`func NewPushRespWithDefaults() *PushResp`

NewPushRespWithDefaults instantiates a new PushResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBranch

`func (o *PushResp) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *PushResp) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *PushResp) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *PushResp) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetCloneUrl

`func (o *PushResp) GetCloneUrl() string`

GetCloneUrl returns the CloneUrl field if non-nil, zero value otherwise.

### GetCloneUrlOk

`func (o *PushResp) GetCloneUrlOk() (*string, bool)`

GetCloneUrlOk returns a tuple with the CloneUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloneUrl

`func (o *PushResp) SetCloneUrl(v string)`

SetCloneUrl sets CloneUrl field to given value.

### HasCloneUrl

`func (o *PushResp) HasCloneUrl() bool`

HasCloneUrl returns a boolean if a field has been set.

### GetCommit

`func (o *PushResp) GetCommit() string`

GetCommit returns the Commit field if non-nil, zero value otherwise.

### GetCommitOk

`func (o *PushResp) GetCommitOk() (*string, bool)`

GetCommitOk returns a tuple with the Commit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommit

`func (o *PushResp) SetCommit(v string)`

SetCommit sets Commit field to given value.

### HasCommit

`func (o *PushResp) HasCommit() bool`

HasCommit returns a boolean if a field has been set.

### GetSshUrl

`func (o *PushResp) GetSshUrl() string`

GetSshUrl returns the SshUrl field if non-nil, zero value otherwise.

### GetSshUrlOk

`func (o *PushResp) GetSshUrlOk() (*string, bool)`

GetSshUrlOk returns a tuple with the SshUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshUrl

`func (o *PushResp) SetSshUrl(v string)`

SetSshUrl sets SshUrl field to given value.

### HasSshUrl

`func (o *PushResp) HasSshUrl() bool`

HasSshUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


