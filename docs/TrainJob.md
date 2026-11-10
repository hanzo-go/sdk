# TrainJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Adaptation** | Pointer to [**TrainAdaptation**](TrainAdaptation.md) | Adaptation is how the update is parameterized. | [optional] 
**Artifacts** | Pointer to [**[]TrainArtifact**](TrainArtifact.md) | Artifacts are its outputs. | [optional] 
**BaseModel** | Pointer to **string** | BaseModel is the base the capability is taught to: &#x60;kai&#x60;, or a model the engine loads for a client. | [optional] 
**Created** | Pointer to **int64** | Created, Started and Ended are unix seconds; 0 is not yet. | [optional] 
**CreatedBy** | Pointer to **string** | CreatedBy is the principal that created it. | [optional] 
**Dataset** | Pointer to [**TrainDataset**](TrainDataset.md) | Dataset is what the job trains on. | [optional] 
**Ended** | Pointer to **int64** |  | [optional] 
**Error** | Pointer to **string** | Error says why a job failed or stopped. A job stopped at what its payer or its own bounds could hold says budget, max_seconds, insufficient_balance, spend_cap, balance_unavailable or no_payer. | [optional] 
**Evaluation** | Pointer to [**TrainEvaluation**](TrainEvaluation.md) | Evaluation names the target suites the result reports. | [optional] 
**From** | Pointer to **string** | From is an artifact the job starts from, by sha256. | [optional] 
**Id** | Pointer to **string** | ID names the job. | [optional] 
**Objective** | Pointer to [**TrainObjective**](TrainObjective.md) | Objective is the loss and the terms beside it. | [optional] 
**Output** | Pointer to [**TrainOutput**](TrainOutput.md) | Output is what the job produces. | [optional] 
**Project** | Pointer to **string** | Project is the creator&#39;s project scope. | [optional] 
**Protect** | Pointer to [**TrainProtect**](TrainProtect.md) | Protect names what the base already does that must hold, and the budget. | [optional] 
**Published** | Pointer to **interface{}** |  | [optional] 
**Resources** | Pointer to [**TrainResources**](TrainResources.md) | Resources places and bounds the run. | [optional] 
**Result** | Pointer to **interface{}** |  | [optional] 
**Revision** | Pointer to **string** | Revision pins the base. For kai: a finished run&#39;s name on the executor, or &#x60;hub:&lt;repo&gt;@&lt;rev&gt;&#x60;. | [optional] 
**Started** | Pointer to **int64** |  | [optional] 
**Status** | Pointer to **string** | Status is queued, running, evaluating, succeeded, rejected, failed, cancelled or stopped. | [optional] 
**Tasks** | Pointer to [**[]TrainTaskView**](TrainTaskView.md) | Tasks are its units of execution. | [optional] 
**Usage** | Pointer to [**TrainUsage**](TrainUsage.md) | Usage is what its tasks used and what that cost. | [optional] 

## Methods

### NewTrainJob

`func NewTrainJob() *TrainJob`

NewTrainJob instantiates a new TrainJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrainJobWithDefaults

`func NewTrainJobWithDefaults() *TrainJob`

NewTrainJobWithDefaults instantiates a new TrainJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdaptation

`func (o *TrainJob) GetAdaptation() TrainAdaptation`

GetAdaptation returns the Adaptation field if non-nil, zero value otherwise.

### GetAdaptationOk

`func (o *TrainJob) GetAdaptationOk() (*TrainAdaptation, bool)`

GetAdaptationOk returns a tuple with the Adaptation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdaptation

`func (o *TrainJob) SetAdaptation(v TrainAdaptation)`

SetAdaptation sets Adaptation field to given value.

### HasAdaptation

`func (o *TrainJob) HasAdaptation() bool`

HasAdaptation returns a boolean if a field has been set.

### GetArtifacts

`func (o *TrainJob) GetArtifacts() []TrainArtifact`

GetArtifacts returns the Artifacts field if non-nil, zero value otherwise.

### GetArtifactsOk

`func (o *TrainJob) GetArtifactsOk() (*[]TrainArtifact, bool)`

GetArtifactsOk returns a tuple with the Artifacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifacts

`func (o *TrainJob) SetArtifacts(v []TrainArtifact)`

SetArtifacts sets Artifacts field to given value.

### HasArtifacts

`func (o *TrainJob) HasArtifacts() bool`

HasArtifacts returns a boolean if a field has been set.

### GetBaseModel

`func (o *TrainJob) GetBaseModel() string`

GetBaseModel returns the BaseModel field if non-nil, zero value otherwise.

### GetBaseModelOk

`func (o *TrainJob) GetBaseModelOk() (*string, bool)`

GetBaseModelOk returns a tuple with the BaseModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseModel

`func (o *TrainJob) SetBaseModel(v string)`

SetBaseModel sets BaseModel field to given value.

### HasBaseModel

`func (o *TrainJob) HasBaseModel() bool`

HasBaseModel returns a boolean if a field has been set.

### GetCreated

`func (o *TrainJob) GetCreated() int64`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *TrainJob) GetCreatedOk() (*int64, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *TrainJob) SetCreated(v int64)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *TrainJob) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCreatedBy

`func (o *TrainJob) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *TrainJob) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *TrainJob) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *TrainJob) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetDataset

`func (o *TrainJob) GetDataset() TrainDataset`

GetDataset returns the Dataset field if non-nil, zero value otherwise.

### GetDatasetOk

`func (o *TrainJob) GetDatasetOk() (*TrainDataset, bool)`

GetDatasetOk returns a tuple with the Dataset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataset

`func (o *TrainJob) SetDataset(v TrainDataset)`

SetDataset sets Dataset field to given value.

### HasDataset

`func (o *TrainJob) HasDataset() bool`

HasDataset returns a boolean if a field has been set.

### GetEnded

`func (o *TrainJob) GetEnded() int64`

GetEnded returns the Ended field if non-nil, zero value otherwise.

### GetEndedOk

`func (o *TrainJob) GetEndedOk() (*int64, bool)`

GetEndedOk returns a tuple with the Ended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnded

`func (o *TrainJob) SetEnded(v int64)`

SetEnded sets Ended field to given value.

### HasEnded

`func (o *TrainJob) HasEnded() bool`

HasEnded returns a boolean if a field has been set.

### GetError

`func (o *TrainJob) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *TrainJob) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *TrainJob) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *TrainJob) HasError() bool`

HasError returns a boolean if a field has been set.

### GetEvaluation

`func (o *TrainJob) GetEvaluation() TrainEvaluation`

GetEvaluation returns the Evaluation field if non-nil, zero value otherwise.

### GetEvaluationOk

`func (o *TrainJob) GetEvaluationOk() (*TrainEvaluation, bool)`

GetEvaluationOk returns a tuple with the Evaluation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluation

`func (o *TrainJob) SetEvaluation(v TrainEvaluation)`

SetEvaluation sets Evaluation field to given value.

### HasEvaluation

`func (o *TrainJob) HasEvaluation() bool`

HasEvaluation returns a boolean if a field has been set.

### GetFrom

`func (o *TrainJob) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *TrainJob) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *TrainJob) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *TrainJob) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetId

`func (o *TrainJob) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TrainJob) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TrainJob) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TrainJob) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObjective

`func (o *TrainJob) GetObjective() TrainObjective`

GetObjective returns the Objective field if non-nil, zero value otherwise.

### GetObjectiveOk

`func (o *TrainJob) GetObjectiveOk() (*TrainObjective, bool)`

GetObjectiveOk returns a tuple with the Objective field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjective

`func (o *TrainJob) SetObjective(v TrainObjective)`

SetObjective sets Objective field to given value.

### HasObjective

`func (o *TrainJob) HasObjective() bool`

HasObjective returns a boolean if a field has been set.

### GetOutput

`func (o *TrainJob) GetOutput() TrainOutput`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *TrainJob) GetOutputOk() (*TrainOutput, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *TrainJob) SetOutput(v TrainOutput)`

SetOutput sets Output field to given value.

### HasOutput

`func (o *TrainJob) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### GetProject

`func (o *TrainJob) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *TrainJob) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *TrainJob) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *TrainJob) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetProtect

`func (o *TrainJob) GetProtect() TrainProtect`

GetProtect returns the Protect field if non-nil, zero value otherwise.

### GetProtectOk

`func (o *TrainJob) GetProtectOk() (*TrainProtect, bool)`

GetProtectOk returns a tuple with the Protect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtect

`func (o *TrainJob) SetProtect(v TrainProtect)`

SetProtect sets Protect field to given value.

### HasProtect

`func (o *TrainJob) HasProtect() bool`

HasProtect returns a boolean if a field has been set.

### GetPublished

`func (o *TrainJob) GetPublished() interface{}`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *TrainJob) GetPublishedOk() (*interface{}, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *TrainJob) SetPublished(v interface{})`

SetPublished sets Published field to given value.

### HasPublished

`func (o *TrainJob) HasPublished() bool`

HasPublished returns a boolean if a field has been set.

### SetPublishedNil

`func (o *TrainJob) SetPublishedNil(b bool)`

 SetPublishedNil sets the value for Published to be an explicit nil

### UnsetPublished
`func (o *TrainJob) UnsetPublished()`

UnsetPublished ensures that no value is present for Published, not even an explicit nil
### GetResources

`func (o *TrainJob) GetResources() TrainResources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *TrainJob) GetResourcesOk() (*TrainResources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *TrainJob) SetResources(v TrainResources)`

SetResources sets Resources field to given value.

### HasResources

`func (o *TrainJob) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetResult

`func (o *TrainJob) GetResult() interface{}`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *TrainJob) GetResultOk() (*interface{}, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *TrainJob) SetResult(v interface{})`

SetResult sets Result field to given value.

### HasResult

`func (o *TrainJob) HasResult() bool`

HasResult returns a boolean if a field has been set.

### SetResultNil

`func (o *TrainJob) SetResultNil(b bool)`

 SetResultNil sets the value for Result to be an explicit nil

### UnsetResult
`func (o *TrainJob) UnsetResult()`

UnsetResult ensures that no value is present for Result, not even an explicit nil
### GetRevision

`func (o *TrainJob) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *TrainJob) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *TrainJob) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *TrainJob) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetStarted

`func (o *TrainJob) GetStarted() int64`

GetStarted returns the Started field if non-nil, zero value otherwise.

### GetStartedOk

`func (o *TrainJob) GetStartedOk() (*int64, bool)`

GetStartedOk returns a tuple with the Started field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarted

`func (o *TrainJob) SetStarted(v int64)`

SetStarted sets Started field to given value.

### HasStarted

`func (o *TrainJob) HasStarted() bool`

HasStarted returns a boolean if a field has been set.

### GetStatus

`func (o *TrainJob) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TrainJob) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TrainJob) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TrainJob) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTasks

`func (o *TrainJob) GetTasks() []TrainTaskView`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *TrainJob) GetTasksOk() (*[]TrainTaskView, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *TrainJob) SetTasks(v []TrainTaskView)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *TrainJob) HasTasks() bool`

HasTasks returns a boolean if a field has been set.

### GetUsage

`func (o *TrainJob) GetUsage() TrainUsage`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *TrainJob) GetUsageOk() (*TrainUsage, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *TrainJob) SetUsage(v TrainUsage)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *TrainJob) HasUsage() bool`

HasUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


