# AgentCodingChanges

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Base** | Pointer to **string** | Base is the branch the run started from: the one it was given, or the repository&#39;s default. | [optional] 
**Commits** | Pointer to [**[]AgentCodingCommit**](AgentCodingCommit.md) | Commits are the commits on head that base does not have, newest first. | [optional] 
**Files** | Pointer to [**[]AgentCodingFile**](AgentCodingFile.md) | Files is the net change head makes against its merge base with base, one entry per file. | [optional] 
**Head** | Pointer to **string** | Head is the run&#39;s own branch. Empty for a plan or a setup, which write none. | [optional] 
**MoreCommits** | Pointer to **bool** | MoreCommits says the branch carries more commits than the newest 250 that Commits lists. | [optional] 
**MoreFiles** | Pointer to **bool** | MoreFiles says the change is larger than one read answers — past 8 MiB of diff or 3000 files — so Files lists the files before that, and the last of them is marked truncated with the counts of the part read. | [optional] 
**Pull** | Pointer to [**AgentCodingPull**](AgentCodingPull.md) | Pull is the run&#39;s pull request on the forge, or null while it has none — and for a run proposed somewhere else, a repository the forge only mirrors from GitHub. | [optional] 
**Repo** | Pointer to **string** | Repo is the run&#39;s repository on the forge, owner/name. | [optional] 

## Methods

### NewAgentCodingChanges

`func NewAgentCodingChanges() *AgentCodingChanges`

NewAgentCodingChanges instantiates a new AgentCodingChanges object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentCodingChangesWithDefaults

`func NewAgentCodingChangesWithDefaults() *AgentCodingChanges`

NewAgentCodingChangesWithDefaults instantiates a new AgentCodingChanges object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBase

`func (o *AgentCodingChanges) GetBase() string`

GetBase returns the Base field if non-nil, zero value otherwise.

### GetBaseOk

`func (o *AgentCodingChanges) GetBaseOk() (*string, bool)`

GetBaseOk returns a tuple with the Base field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBase

`func (o *AgentCodingChanges) SetBase(v string)`

SetBase sets Base field to given value.

### HasBase

`func (o *AgentCodingChanges) HasBase() bool`

HasBase returns a boolean if a field has been set.

### GetCommits

`func (o *AgentCodingChanges) GetCommits() []AgentCodingCommit`

GetCommits returns the Commits field if non-nil, zero value otherwise.

### GetCommitsOk

`func (o *AgentCodingChanges) GetCommitsOk() (*[]AgentCodingCommit, bool)`

GetCommitsOk returns a tuple with the Commits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommits

`func (o *AgentCodingChanges) SetCommits(v []AgentCodingCommit)`

SetCommits sets Commits field to given value.

### HasCommits

`func (o *AgentCodingChanges) HasCommits() bool`

HasCommits returns a boolean if a field has been set.

### GetFiles

`func (o *AgentCodingChanges) GetFiles() []AgentCodingFile`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *AgentCodingChanges) GetFilesOk() (*[]AgentCodingFile, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *AgentCodingChanges) SetFiles(v []AgentCodingFile)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *AgentCodingChanges) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetHead

`func (o *AgentCodingChanges) GetHead() string`

GetHead returns the Head field if non-nil, zero value otherwise.

### GetHeadOk

`func (o *AgentCodingChanges) GetHeadOk() (*string, bool)`

GetHeadOk returns a tuple with the Head field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHead

`func (o *AgentCodingChanges) SetHead(v string)`

SetHead sets Head field to given value.

### HasHead

`func (o *AgentCodingChanges) HasHead() bool`

HasHead returns a boolean if a field has been set.

### GetMoreCommits

`func (o *AgentCodingChanges) GetMoreCommits() bool`

GetMoreCommits returns the MoreCommits field if non-nil, zero value otherwise.

### GetMoreCommitsOk

`func (o *AgentCodingChanges) GetMoreCommitsOk() (*bool, bool)`

GetMoreCommitsOk returns a tuple with the MoreCommits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoreCommits

`func (o *AgentCodingChanges) SetMoreCommits(v bool)`

SetMoreCommits sets MoreCommits field to given value.

### HasMoreCommits

`func (o *AgentCodingChanges) HasMoreCommits() bool`

HasMoreCommits returns a boolean if a field has been set.

### GetMoreFiles

`func (o *AgentCodingChanges) GetMoreFiles() bool`

GetMoreFiles returns the MoreFiles field if non-nil, zero value otherwise.

### GetMoreFilesOk

`func (o *AgentCodingChanges) GetMoreFilesOk() (*bool, bool)`

GetMoreFilesOk returns a tuple with the MoreFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoreFiles

`func (o *AgentCodingChanges) SetMoreFiles(v bool)`

SetMoreFiles sets MoreFiles field to given value.

### HasMoreFiles

`func (o *AgentCodingChanges) HasMoreFiles() bool`

HasMoreFiles returns a boolean if a field has been set.

### GetPull

`func (o *AgentCodingChanges) GetPull() AgentCodingPull`

GetPull returns the Pull field if non-nil, zero value otherwise.

### GetPullOk

`func (o *AgentCodingChanges) GetPullOk() (*AgentCodingPull, bool)`

GetPullOk returns a tuple with the Pull field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPull

`func (o *AgentCodingChanges) SetPull(v AgentCodingPull)`

SetPull sets Pull field to given value.

### HasPull

`func (o *AgentCodingChanges) HasPull() bool`

HasPull returns a boolean if a field has been set.

### GetRepo

`func (o *AgentCodingChanges) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *AgentCodingChanges) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *AgentCodingChanges) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *AgentCodingChanges) HasRepo() bool`

HasRepo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


