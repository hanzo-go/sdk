# CloudAgentsAgentDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Public handle (agent_...). | [optional] 
**Name** | Pointer to **string** | Org-unique name matching ^[A-Za-z0-9][A-Za-z0-9._-]{0,63}$. | [optional] 
**Model** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Tools** | Pointer to **[]string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**ExecutionMode** | Pointer to **string** |  | [optional] 
**Schedule** | Pointer to **string** | 5-field cron; required and evaluated only when long-running. | [optional] 
**ComputeRef** | Pointer to **string** | Optional visor machine id the bot is bound to. | [optional] 
**ServiceAccountId** | Pointer to **string** | Optional IAM agent service-account (&lt;org&gt;-&lt;agent&gt;) recorded as the actor on scheduled-run billing. | [optional] 
**Runs** | Pointer to **int32** | Recorded run count. | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**Instructions** | Pointer to **string** |  | [optional] 
**RecentRuns** | Pointer to [**[]CloudAgentsRun**](CloudAgentsRun.md) |  | [optional] 

## Methods

### NewCloudAgentsAgentDetail

`func NewCloudAgentsAgentDetail() *CloudAgentsAgentDetail`

NewCloudAgentsAgentDetail instantiates a new CloudAgentsAgentDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudAgentsAgentDetailWithDefaults

`func NewCloudAgentsAgentDetailWithDefaults() *CloudAgentsAgentDetail`

NewCloudAgentsAgentDetailWithDefaults instantiates a new CloudAgentsAgentDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CloudAgentsAgentDetail) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudAgentsAgentDetail) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudAgentsAgentDetail) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudAgentsAgentDetail) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CloudAgentsAgentDetail) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudAgentsAgentDetail) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudAgentsAgentDetail) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudAgentsAgentDetail) HasName() bool`

HasName returns a boolean if a field has been set.

### GetModel

`func (o *CloudAgentsAgentDetail) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *CloudAgentsAgentDetail) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *CloudAgentsAgentDetail) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *CloudAgentsAgentDetail) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetDescription

`func (o *CloudAgentsAgentDetail) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CloudAgentsAgentDetail) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CloudAgentsAgentDetail) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CloudAgentsAgentDetail) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTools

`func (o *CloudAgentsAgentDetail) GetTools() []string`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *CloudAgentsAgentDetail) GetToolsOk() (*[]string, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *CloudAgentsAgentDetail) SetTools(v []string)`

SetTools sets Tools field to given value.

### HasTools

`func (o *CloudAgentsAgentDetail) HasTools() bool`

HasTools returns a boolean if a field has been set.

### GetStatus

`func (o *CloudAgentsAgentDetail) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudAgentsAgentDetail) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudAgentsAgentDetail) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudAgentsAgentDetail) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetExecutionMode

`func (o *CloudAgentsAgentDetail) GetExecutionMode() string`

GetExecutionMode returns the ExecutionMode field if non-nil, zero value otherwise.

### GetExecutionModeOk

`func (o *CloudAgentsAgentDetail) GetExecutionModeOk() (*string, bool)`

GetExecutionModeOk returns a tuple with the ExecutionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionMode

`func (o *CloudAgentsAgentDetail) SetExecutionMode(v string)`

SetExecutionMode sets ExecutionMode field to given value.

### HasExecutionMode

`func (o *CloudAgentsAgentDetail) HasExecutionMode() bool`

HasExecutionMode returns a boolean if a field has been set.

### GetSchedule

`func (o *CloudAgentsAgentDetail) GetSchedule() string`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *CloudAgentsAgentDetail) GetScheduleOk() (*string, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *CloudAgentsAgentDetail) SetSchedule(v string)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *CloudAgentsAgentDetail) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetComputeRef

`func (o *CloudAgentsAgentDetail) GetComputeRef() string`

GetComputeRef returns the ComputeRef field if non-nil, zero value otherwise.

### GetComputeRefOk

`func (o *CloudAgentsAgentDetail) GetComputeRefOk() (*string, bool)`

GetComputeRefOk returns a tuple with the ComputeRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRef

`func (o *CloudAgentsAgentDetail) SetComputeRef(v string)`

SetComputeRef sets ComputeRef field to given value.

### HasComputeRef

`func (o *CloudAgentsAgentDetail) HasComputeRef() bool`

HasComputeRef returns a boolean if a field has been set.

### GetServiceAccountId

`func (o *CloudAgentsAgentDetail) GetServiceAccountId() string`

GetServiceAccountId returns the ServiceAccountId field if non-nil, zero value otherwise.

### GetServiceAccountIdOk

`func (o *CloudAgentsAgentDetail) GetServiceAccountIdOk() (*string, bool)`

GetServiceAccountIdOk returns a tuple with the ServiceAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountId

`func (o *CloudAgentsAgentDetail) SetServiceAccountId(v string)`

SetServiceAccountId sets ServiceAccountId field to given value.

### HasServiceAccountId

`func (o *CloudAgentsAgentDetail) HasServiceAccountId() bool`

HasServiceAccountId returns a boolean if a field has been set.

### GetRuns

`func (o *CloudAgentsAgentDetail) GetRuns() int32`

GetRuns returns the Runs field if non-nil, zero value otherwise.

### GetRunsOk

`func (o *CloudAgentsAgentDetail) GetRunsOk() (*int32, bool)`

GetRunsOk returns a tuple with the Runs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuns

`func (o *CloudAgentsAgentDetail) SetRuns(v int32)`

SetRuns sets Runs field to given value.

### HasRuns

`func (o *CloudAgentsAgentDetail) HasRuns() bool`

HasRuns returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CloudAgentsAgentDetail) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CloudAgentsAgentDetail) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CloudAgentsAgentDetail) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CloudAgentsAgentDetail) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CloudAgentsAgentDetail) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CloudAgentsAgentDetail) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CloudAgentsAgentDetail) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CloudAgentsAgentDetail) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetInstructions

`func (o *CloudAgentsAgentDetail) GetInstructions() string`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *CloudAgentsAgentDetail) GetInstructionsOk() (*string, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *CloudAgentsAgentDetail) SetInstructions(v string)`

SetInstructions sets Instructions field to given value.

### HasInstructions

`func (o *CloudAgentsAgentDetail) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.

### GetRecentRuns

`func (o *CloudAgentsAgentDetail) GetRecentRuns() []CloudAgentsRun`

GetRecentRuns returns the RecentRuns field if non-nil, zero value otherwise.

### GetRecentRunsOk

`func (o *CloudAgentsAgentDetail) GetRecentRunsOk() (*[]CloudAgentsRun, bool)`

GetRecentRunsOk returns a tuple with the RecentRuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecentRuns

`func (o *CloudAgentsAgentDetail) SetRecentRuns(v []CloudAgentsRun)`

SetRecentRuns sets RecentRuns field to given value.

### HasRecentRuns

`func (o *CloudAgentsAgentDetail) HasRecentRuns() bool`

HasRecentRuns returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


