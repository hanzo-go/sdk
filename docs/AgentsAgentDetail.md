# AgentsAgentDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Model** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Tools** | Pointer to **[]string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**ExecutionMode** | Pointer to **string** |  | [optional] 
**Schedule** | Pointer to **string** |  | [optional] 
**ComputeRef** | Pointer to **string** |  | [optional] 
**ServiceAccountId** | Pointer to **string** |  | [optional] 
**Runs** | Pointer to **int32** | Recorded run count. | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**Instructions** | Pointer to **string** |  | [optional] 
**RecentRuns** | Pointer to [**[]AgentsRunView**](AgentsRunView.md) |  | [optional] 

## Methods

### NewAgentsAgentDetail

`func NewAgentsAgentDetail() *AgentsAgentDetail`

NewAgentsAgentDetail instantiates a new AgentsAgentDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentsAgentDetailWithDefaults

`func NewAgentsAgentDetailWithDefaults() *AgentsAgentDetail`

NewAgentsAgentDetailWithDefaults instantiates a new AgentsAgentDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AgentsAgentDetail) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AgentsAgentDetail) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AgentsAgentDetail) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AgentsAgentDetail) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AgentsAgentDetail) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AgentsAgentDetail) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AgentsAgentDetail) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AgentsAgentDetail) HasName() bool`

HasName returns a boolean if a field has been set.

### GetModel

`func (o *AgentsAgentDetail) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *AgentsAgentDetail) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *AgentsAgentDetail) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *AgentsAgentDetail) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetDescription

`func (o *AgentsAgentDetail) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AgentsAgentDetail) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AgentsAgentDetail) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AgentsAgentDetail) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTools

`func (o *AgentsAgentDetail) GetTools() []string`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *AgentsAgentDetail) GetToolsOk() (*[]string, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *AgentsAgentDetail) SetTools(v []string)`

SetTools sets Tools field to given value.

### HasTools

`func (o *AgentsAgentDetail) HasTools() bool`

HasTools returns a boolean if a field has been set.

### GetStatus

`func (o *AgentsAgentDetail) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AgentsAgentDetail) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AgentsAgentDetail) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AgentsAgentDetail) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetExecutionMode

`func (o *AgentsAgentDetail) GetExecutionMode() string`

GetExecutionMode returns the ExecutionMode field if non-nil, zero value otherwise.

### GetExecutionModeOk

`func (o *AgentsAgentDetail) GetExecutionModeOk() (*string, bool)`

GetExecutionModeOk returns a tuple with the ExecutionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionMode

`func (o *AgentsAgentDetail) SetExecutionMode(v string)`

SetExecutionMode sets ExecutionMode field to given value.

### HasExecutionMode

`func (o *AgentsAgentDetail) HasExecutionMode() bool`

HasExecutionMode returns a boolean if a field has been set.

### GetSchedule

`func (o *AgentsAgentDetail) GetSchedule() string`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *AgentsAgentDetail) GetScheduleOk() (*string, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *AgentsAgentDetail) SetSchedule(v string)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *AgentsAgentDetail) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetComputeRef

`func (o *AgentsAgentDetail) GetComputeRef() string`

GetComputeRef returns the ComputeRef field if non-nil, zero value otherwise.

### GetComputeRefOk

`func (o *AgentsAgentDetail) GetComputeRefOk() (*string, bool)`

GetComputeRefOk returns a tuple with the ComputeRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRef

`func (o *AgentsAgentDetail) SetComputeRef(v string)`

SetComputeRef sets ComputeRef field to given value.

### HasComputeRef

`func (o *AgentsAgentDetail) HasComputeRef() bool`

HasComputeRef returns a boolean if a field has been set.

### GetServiceAccountId

`func (o *AgentsAgentDetail) GetServiceAccountId() string`

GetServiceAccountId returns the ServiceAccountId field if non-nil, zero value otherwise.

### GetServiceAccountIdOk

`func (o *AgentsAgentDetail) GetServiceAccountIdOk() (*string, bool)`

GetServiceAccountIdOk returns a tuple with the ServiceAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountId

`func (o *AgentsAgentDetail) SetServiceAccountId(v string)`

SetServiceAccountId sets ServiceAccountId field to given value.

### HasServiceAccountId

`func (o *AgentsAgentDetail) HasServiceAccountId() bool`

HasServiceAccountId returns a boolean if a field has been set.

### GetRuns

`func (o *AgentsAgentDetail) GetRuns() int32`

GetRuns returns the Runs field if non-nil, zero value otherwise.

### GetRunsOk

`func (o *AgentsAgentDetail) GetRunsOk() (*int32, bool)`

GetRunsOk returns a tuple with the Runs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuns

`func (o *AgentsAgentDetail) SetRuns(v int32)`

SetRuns sets Runs field to given value.

### HasRuns

`func (o *AgentsAgentDetail) HasRuns() bool`

HasRuns returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AgentsAgentDetail) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AgentsAgentDetail) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AgentsAgentDetail) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AgentsAgentDetail) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AgentsAgentDetail) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AgentsAgentDetail) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AgentsAgentDetail) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AgentsAgentDetail) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetInstructions

`func (o *AgentsAgentDetail) GetInstructions() string`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *AgentsAgentDetail) GetInstructionsOk() (*string, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *AgentsAgentDetail) SetInstructions(v string)`

SetInstructions sets Instructions field to given value.

### HasInstructions

`func (o *AgentsAgentDetail) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.

### GetRecentRuns

`func (o *AgentsAgentDetail) GetRecentRuns() []AgentsRunView`

GetRecentRuns returns the RecentRuns field if non-nil, zero value otherwise.

### GetRecentRunsOk

`func (o *AgentsAgentDetail) GetRecentRunsOk() (*[]AgentsRunView, bool)`

GetRecentRunsOk returns a tuple with the RecentRuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecentRuns

`func (o *AgentsAgentDetail) SetRecentRuns(v []AgentsRunView)`

SetRecentRuns sets RecentRuns field to given value.

### HasRecentRuns

`func (o *AgentsAgentDetail) HasRecentRuns() bool`

HasRecentRuns returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


