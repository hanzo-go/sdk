# CiRun

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attempt** | Pointer to **int64** | Attempt counts re-runs of this run, starting at 1. | [optional] 
**Conclusion** | Pointer to **string** | Conclusion is how it ended — success, failure, cancelled or skipped — and is empty until Status is completed. Read both: completed alone does not say whether it passed. | [optional] 
**EndedAt** | Pointer to **time.Time** | EndedAt is when it finished; absent until it has. | [optional] 
**Event** | Pointer to **string** | Event is what started the run: push, pull_request, workflow_dispatch, schedule. | [optional] 
**Id** | Pointer to **int64** | ID is the run&#39;s id on the forge. | [optional] 
**Jobs** | Pointer to [**[]CiJob**](CiJob.md) | Jobs are the run&#39;s jobs in the order the forge created them. | [optional] 
**Number** | Pointer to **int64** | Number is the run&#39;s number within the repository, as the forge&#39;s page shows it. | [optional] 
**Ref** | Pointer to **string** | Ref is the branch or tag that started the run. A tag and the branch it was cut from share a commit, so one answer can hold a run for each. | [optional] 
**Sha** | Pointer to **string** | SHA is the commit the run checked out. | [optional] 
**StartedAt** | Pointer to **time.Time** | StartedAt is when the run started; absent until it has. | [optional] 
**Status** | Pointer to **string** | Status is whether the run is over: queued, waiting, in_progress or completed. | [optional] 
**Title** | Pointer to **string** | Title is the run&#39;s display title, usually the commit&#39;s subject. | [optional] 
**Url** | Pointer to **string** | URL is the run&#39;s page on the forge. | [optional] 
**Workflow** | Pointer to **string** | Workflow is the workflow file that ran, such as cicd.yml. | [optional] 

## Methods

### NewCiRun

`func NewCiRun() *CiRun`

NewCiRun instantiates a new CiRun object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCiRunWithDefaults

`func NewCiRunWithDefaults() *CiRun`

NewCiRunWithDefaults instantiates a new CiRun object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttempt

`func (o *CiRun) GetAttempt() int64`

GetAttempt returns the Attempt field if non-nil, zero value otherwise.

### GetAttemptOk

`func (o *CiRun) GetAttemptOk() (*int64, bool)`

GetAttemptOk returns a tuple with the Attempt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttempt

`func (o *CiRun) SetAttempt(v int64)`

SetAttempt sets Attempt field to given value.

### HasAttempt

`func (o *CiRun) HasAttempt() bool`

HasAttempt returns a boolean if a field has been set.

### GetConclusion

`func (o *CiRun) GetConclusion() string`

GetConclusion returns the Conclusion field if non-nil, zero value otherwise.

### GetConclusionOk

`func (o *CiRun) GetConclusionOk() (*string, bool)`

GetConclusionOk returns a tuple with the Conclusion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConclusion

`func (o *CiRun) SetConclusion(v string)`

SetConclusion sets Conclusion field to given value.

### HasConclusion

`func (o *CiRun) HasConclusion() bool`

HasConclusion returns a boolean if a field has been set.

### GetEndedAt

`func (o *CiRun) GetEndedAt() time.Time`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *CiRun) GetEndedAtOk() (*time.Time, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *CiRun) SetEndedAt(v time.Time)`

SetEndedAt sets EndedAt field to given value.

### HasEndedAt

`func (o *CiRun) HasEndedAt() bool`

HasEndedAt returns a boolean if a field has been set.

### GetEvent

`func (o *CiRun) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *CiRun) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *CiRun) SetEvent(v string)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *CiRun) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetId

`func (o *CiRun) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CiRun) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CiRun) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *CiRun) HasId() bool`

HasId returns a boolean if a field has been set.

### GetJobs

`func (o *CiRun) GetJobs() []CiJob`

GetJobs returns the Jobs field if non-nil, zero value otherwise.

### GetJobsOk

`func (o *CiRun) GetJobsOk() (*[]CiJob, bool)`

GetJobsOk returns a tuple with the Jobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobs

`func (o *CiRun) SetJobs(v []CiJob)`

SetJobs sets Jobs field to given value.

### HasJobs

`func (o *CiRun) HasJobs() bool`

HasJobs returns a boolean if a field has been set.

### GetNumber

`func (o *CiRun) GetNumber() int64`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *CiRun) GetNumberOk() (*int64, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *CiRun) SetNumber(v int64)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *CiRun) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetRef

`func (o *CiRun) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *CiRun) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *CiRun) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *CiRun) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetSha

`func (o *CiRun) GetSha() string`

GetSha returns the Sha field if non-nil, zero value otherwise.

### GetShaOk

`func (o *CiRun) GetShaOk() (*string, bool)`

GetShaOk returns a tuple with the Sha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha

`func (o *CiRun) SetSha(v string)`

SetSha sets Sha field to given value.

### HasSha

`func (o *CiRun) HasSha() bool`

HasSha returns a boolean if a field has been set.

### GetStartedAt

`func (o *CiRun) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *CiRun) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *CiRun) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *CiRun) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetStatus

`func (o *CiRun) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CiRun) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CiRun) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CiRun) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTitle

`func (o *CiRun) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CiRun) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CiRun) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CiRun) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUrl

`func (o *CiRun) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CiRun) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CiRun) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *CiRun) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetWorkflow

`func (o *CiRun) GetWorkflow() string`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *CiRun) GetWorkflowOk() (*string, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *CiRun) SetWorkflow(v string)`

SetWorkflow sets Workflow field to given value.

### HasWorkflow

`func (o *CiRun) HasWorkflow() bool`

HasWorkflow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


