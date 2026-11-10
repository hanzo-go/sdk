# GitActivityCommit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**At** | Pointer to **string** | At is the commit time, RFC 3339 UTC. | [optional] 
**Message** | Pointer to **string** | Message is the subject line. | [optional] 
**Ours** | Pointer to **bool** | Ours is whether the estate&#39;s people authored it. | [optional] 
**Repo** | Pointer to **string** | Repo is org/name. | [optional] 
**Sha** | Pointer to **string** | SHA is the commit. | [optional] 
**Url** | Pointer to **string** | URL is the commit on the forge. | [optional] 

## Methods

### NewGitActivityCommit

`func NewGitActivityCommit() *GitActivityCommit`

NewGitActivityCommit instantiates a new GitActivityCommit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitActivityCommitWithDefaults

`func NewGitActivityCommitWithDefaults() *GitActivityCommit`

NewGitActivityCommitWithDefaults instantiates a new GitActivityCommit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAt

`func (o *GitActivityCommit) GetAt() string`

GetAt returns the At field if non-nil, zero value otherwise.

### GetAtOk

`func (o *GitActivityCommit) GetAtOk() (*string, bool)`

GetAtOk returns a tuple with the At field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAt

`func (o *GitActivityCommit) SetAt(v string)`

SetAt sets At field to given value.

### HasAt

`func (o *GitActivityCommit) HasAt() bool`

HasAt returns a boolean if a field has been set.

### GetMessage

`func (o *GitActivityCommit) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GitActivityCommit) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GitActivityCommit) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GitActivityCommit) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetOurs

`func (o *GitActivityCommit) GetOurs() bool`

GetOurs returns the Ours field if non-nil, zero value otherwise.

### GetOursOk

`func (o *GitActivityCommit) GetOursOk() (*bool, bool)`

GetOursOk returns a tuple with the Ours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOurs

`func (o *GitActivityCommit) SetOurs(v bool)`

SetOurs sets Ours field to given value.

### HasOurs

`func (o *GitActivityCommit) HasOurs() bool`

HasOurs returns a boolean if a field has been set.

### GetRepo

`func (o *GitActivityCommit) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *GitActivityCommit) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *GitActivityCommit) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *GitActivityCommit) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### GetSha

`func (o *GitActivityCommit) GetSha() string`

GetSha returns the Sha field if non-nil, zero value otherwise.

### GetShaOk

`func (o *GitActivityCommit) GetShaOk() (*string, bool)`

GetShaOk returns a tuple with the Sha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha

`func (o *GitActivityCommit) SetSha(v string)`

SetSha sets Sha field to given value.

### HasSha

`func (o *GitActivityCommit) HasSha() bool`

HasSha returns a boolean if a field has been set.

### GetUrl

`func (o *GitActivityCommit) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *GitActivityCommit) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *GitActivityCommit) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *GitActivityCommit) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


