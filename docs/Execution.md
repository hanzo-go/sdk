# Execution

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actor** | Pointer to **string** |  | [optional] 
**Branch** | Pointer to **string** |  | [optional] 
**Conclusion** | Pointer to **string** |  | [optional] 
**EndedAt** | Pointer to **time.Time** |  | [optional] 
**Event** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Number** | Pointer to **int32** |  | [optional] 
**Org** | Pointer to **string** |  | [optional] 
**Repo** | Pointer to **string** |  | [optional] 
**Sha** | Pointer to **string** |  | [optional] 
**StartedAt** | Pointer to **time.Time** |  | [optional] 
**Status** | Pointer to **string** | Status and Conclusion are BOTH required to know how a run went, and reading only one is wrong in a way that looks fine. Status answers \&quot;is it over\&quot; (queued | in_progress | completed); Conclusion answers \&quot;how did it end\&quot; and is empty until it is over. A view that buckets on Status alone sees &#x60;completed&#x60; and cannot tell a pass from a failure, so it draws every finished run — successes and cancellations included — the same way. | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**Workflow** | Pointer to **string** |  | [optional] 

## Methods

### NewExecution

`func NewExecution() *Execution`

NewExecution instantiates a new Execution object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecutionWithDefaults

`func NewExecutionWithDefaults() *Execution`

NewExecutionWithDefaults instantiates a new Execution object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActor

`func (o *Execution) GetActor() string`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *Execution) GetActorOk() (*string, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *Execution) SetActor(v string)`

SetActor sets Actor field to given value.

### HasActor

`func (o *Execution) HasActor() bool`

HasActor returns a boolean if a field has been set.

### GetBranch

`func (o *Execution) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *Execution) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *Execution) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *Execution) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetConclusion

`func (o *Execution) GetConclusion() string`

GetConclusion returns the Conclusion field if non-nil, zero value otherwise.

### GetConclusionOk

`func (o *Execution) GetConclusionOk() (*string, bool)`

GetConclusionOk returns a tuple with the Conclusion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConclusion

`func (o *Execution) SetConclusion(v string)`

SetConclusion sets Conclusion field to given value.

### HasConclusion

`func (o *Execution) HasConclusion() bool`

HasConclusion returns a boolean if a field has been set.

### GetEndedAt

`func (o *Execution) GetEndedAt() time.Time`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *Execution) GetEndedAtOk() (*time.Time, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *Execution) SetEndedAt(v time.Time)`

SetEndedAt sets EndedAt field to given value.

### HasEndedAt

`func (o *Execution) HasEndedAt() bool`

HasEndedAt returns a boolean if a field has been set.

### GetEvent

`func (o *Execution) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *Execution) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *Execution) SetEvent(v string)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *Execution) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetId

`func (o *Execution) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Execution) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Execution) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Execution) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNumber

`func (o *Execution) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *Execution) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *Execution) SetNumber(v int32)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *Execution) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetOrg

`func (o *Execution) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *Execution) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *Execution) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *Execution) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetRepo

`func (o *Execution) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *Execution) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *Execution) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *Execution) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### GetSha

`func (o *Execution) GetSha() string`

GetSha returns the Sha field if non-nil, zero value otherwise.

### GetShaOk

`func (o *Execution) GetShaOk() (*string, bool)`

GetShaOk returns a tuple with the Sha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha

`func (o *Execution) SetSha(v string)`

SetSha sets Sha field to given value.

### HasSha

`func (o *Execution) HasSha() bool`

HasSha returns a boolean if a field has been set.

### GetStartedAt

`func (o *Execution) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *Execution) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *Execution) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *Execution) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetStatus

`func (o *Execution) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Execution) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Execution) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Execution) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTitle

`func (o *Execution) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Execution) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Execution) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Execution) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUrl

`func (o *Execution) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Execution) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Execution) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Execution) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetWorkflow

`func (o *Execution) GetWorkflow() string`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *Execution) GetWorkflowOk() (*string, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *Execution) SetWorkflow(v string)`

SetWorkflow sets Workflow field to given value.

### HasWorkflow

`func (o *Execution) HasWorkflow() bool`

HasWorkflow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


