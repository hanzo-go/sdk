# GitWorkflowRun

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actor** | Pointer to **string** | Actor is who caused it. | [optional] 
**Commit** | Pointer to **string** | Commit is the commit being run, in full. | [optional] 
**CreatedAt** | Pointer to **int64** | CreatedAt is when the run was opened, in unix seconds. | [optional] 
**Event** | Pointer to **string** | Event is what caused the run, e.g. \&quot;push\&quot;. | [optional] 
**Id** | Pointer to **string** | ID addresses this run. | [optional] 
**Number** | Pointer to **int64** | Number is the run&#39;s position in its repository, counting from 1 — the handle a person uses (\&quot;#4\&quot;). | [optional] 
**Ref** | Pointer to **string** | Ref is the full ref, e.g. \&quot;refs/heads/main\&quot;. | [optional] 
**Repo** | Pointer to **string** | Repo is the repository the run is for. | [optional] 
**Status** | Pointer to **string** | Status is waiting, running, success, failure, cancelled or skipped. | [optional] 
**UpdatedAt** | Pointer to **int64** | UpdatedAt is when its status last changed, in unix seconds. | [optional] 
**Workflow** | Pointer to **string** | Workflow is the path of the document being run, e.g. \&quot;.hanzo/workflows/ci.yml\&quot;. | [optional] 

## Methods

### NewGitWorkflowRun

`func NewGitWorkflowRun() *GitWorkflowRun`

NewGitWorkflowRun instantiates a new GitWorkflowRun object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitWorkflowRunWithDefaults

`func NewGitWorkflowRunWithDefaults() *GitWorkflowRun`

NewGitWorkflowRunWithDefaults instantiates a new GitWorkflowRun object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActor

`func (o *GitWorkflowRun) GetActor() string`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *GitWorkflowRun) GetActorOk() (*string, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *GitWorkflowRun) SetActor(v string)`

SetActor sets Actor field to given value.

### HasActor

`func (o *GitWorkflowRun) HasActor() bool`

HasActor returns a boolean if a field has been set.

### GetCommit

`func (o *GitWorkflowRun) GetCommit() string`

GetCommit returns the Commit field if non-nil, zero value otherwise.

### GetCommitOk

`func (o *GitWorkflowRun) GetCommitOk() (*string, bool)`

GetCommitOk returns a tuple with the Commit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommit

`func (o *GitWorkflowRun) SetCommit(v string)`

SetCommit sets Commit field to given value.

### HasCommit

`func (o *GitWorkflowRun) HasCommit() bool`

HasCommit returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GitWorkflowRun) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GitWorkflowRun) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GitWorkflowRun) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GitWorkflowRun) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEvent

`func (o *GitWorkflowRun) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *GitWorkflowRun) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *GitWorkflowRun) SetEvent(v string)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *GitWorkflowRun) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetId

`func (o *GitWorkflowRun) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GitWorkflowRun) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GitWorkflowRun) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GitWorkflowRun) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNumber

`func (o *GitWorkflowRun) GetNumber() int64`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *GitWorkflowRun) GetNumberOk() (*int64, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *GitWorkflowRun) SetNumber(v int64)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *GitWorkflowRun) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetRef

`func (o *GitWorkflowRun) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GitWorkflowRun) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GitWorkflowRun) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GitWorkflowRun) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetRepo

`func (o *GitWorkflowRun) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *GitWorkflowRun) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *GitWorkflowRun) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *GitWorkflowRun) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### GetStatus

`func (o *GitWorkflowRun) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GitWorkflowRun) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GitWorkflowRun) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GitWorkflowRun) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GitWorkflowRun) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GitWorkflowRun) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GitWorkflowRun) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GitWorkflowRun) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetWorkflow

`func (o *GitWorkflowRun) GetWorkflow() string`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *GitWorkflowRun) GetWorkflowOk() (*string, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *GitWorkflowRun) SetWorkflow(v string)`

SetWorkflow sets Workflow field to given value.

### HasWorkflow

`func (o *GitWorkflowRun) HasWorkflow() bool`

HasWorkflow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


