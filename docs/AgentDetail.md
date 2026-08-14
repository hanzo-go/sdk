# AgentDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComputeRef** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ExecutionMode** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Instructions** | Pointer to **string** |  | [optional] 
**Model** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**RecentRuns** | Pointer to [**[]AgentRunView**](AgentRunView.md) |  | [optional] 
**Runs** | Pointer to **int32** |  | [optional] 
**Schedule** | Pointer to **string** |  | [optional] 
**ServiceAccountId** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Tools** | Pointer to **[]string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewAgentDetail

`func NewAgentDetail() *AgentDetail`

NewAgentDetail instantiates a new AgentDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentDetailWithDefaults

`func NewAgentDetailWithDefaults() *AgentDetail`

NewAgentDetailWithDefaults instantiates a new AgentDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComputeRef

`func (o *AgentDetail) GetComputeRef() string`

GetComputeRef returns the ComputeRef field if non-nil, zero value otherwise.

### GetComputeRefOk

`func (o *AgentDetail) GetComputeRefOk() (*string, bool)`

GetComputeRefOk returns a tuple with the ComputeRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRef

`func (o *AgentDetail) SetComputeRef(v string)`

SetComputeRef sets ComputeRef field to given value.

### HasComputeRef

`func (o *AgentDetail) HasComputeRef() bool`

HasComputeRef returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AgentDetail) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AgentDetail) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AgentDetail) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AgentDetail) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *AgentDetail) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AgentDetail) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AgentDetail) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AgentDetail) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExecutionMode

`func (o *AgentDetail) GetExecutionMode() string`

GetExecutionMode returns the ExecutionMode field if non-nil, zero value otherwise.

### GetExecutionModeOk

`func (o *AgentDetail) GetExecutionModeOk() (*string, bool)`

GetExecutionModeOk returns a tuple with the ExecutionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionMode

`func (o *AgentDetail) SetExecutionMode(v string)`

SetExecutionMode sets ExecutionMode field to given value.

### HasExecutionMode

`func (o *AgentDetail) HasExecutionMode() bool`

HasExecutionMode returns a boolean if a field has been set.

### GetId

`func (o *AgentDetail) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AgentDetail) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AgentDetail) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AgentDetail) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInstructions

`func (o *AgentDetail) GetInstructions() string`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *AgentDetail) GetInstructionsOk() (*string, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *AgentDetail) SetInstructions(v string)`

SetInstructions sets Instructions field to given value.

### HasInstructions

`func (o *AgentDetail) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.

### GetModel

`func (o *AgentDetail) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *AgentDetail) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *AgentDetail) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *AgentDetail) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetName

`func (o *AgentDetail) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AgentDetail) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AgentDetail) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AgentDetail) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRecentRuns

`func (o *AgentDetail) GetRecentRuns() []AgentRunView`

GetRecentRuns returns the RecentRuns field if non-nil, zero value otherwise.

### GetRecentRunsOk

`func (o *AgentDetail) GetRecentRunsOk() (*[]AgentRunView, bool)`

GetRecentRunsOk returns a tuple with the RecentRuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecentRuns

`func (o *AgentDetail) SetRecentRuns(v []AgentRunView)`

SetRecentRuns sets RecentRuns field to given value.

### HasRecentRuns

`func (o *AgentDetail) HasRecentRuns() bool`

HasRecentRuns returns a boolean if a field has been set.

### GetRuns

`func (o *AgentDetail) GetRuns() int32`

GetRuns returns the Runs field if non-nil, zero value otherwise.

### GetRunsOk

`func (o *AgentDetail) GetRunsOk() (*int32, bool)`

GetRunsOk returns a tuple with the Runs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuns

`func (o *AgentDetail) SetRuns(v int32)`

SetRuns sets Runs field to given value.

### HasRuns

`func (o *AgentDetail) HasRuns() bool`

HasRuns returns a boolean if a field has been set.

### GetSchedule

`func (o *AgentDetail) GetSchedule() string`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *AgentDetail) GetScheduleOk() (*string, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *AgentDetail) SetSchedule(v string)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *AgentDetail) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetServiceAccountId

`func (o *AgentDetail) GetServiceAccountId() string`

GetServiceAccountId returns the ServiceAccountId field if non-nil, zero value otherwise.

### GetServiceAccountIdOk

`func (o *AgentDetail) GetServiceAccountIdOk() (*string, bool)`

GetServiceAccountIdOk returns a tuple with the ServiceAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountId

`func (o *AgentDetail) SetServiceAccountId(v string)`

SetServiceAccountId sets ServiceAccountId field to given value.

### HasServiceAccountId

`func (o *AgentDetail) HasServiceAccountId() bool`

HasServiceAccountId returns a boolean if a field has been set.

### GetStatus

`func (o *AgentDetail) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AgentDetail) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AgentDetail) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AgentDetail) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTools

`func (o *AgentDetail) GetTools() []string`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *AgentDetail) GetToolsOk() (*[]string, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *AgentDetail) SetTools(v []string)`

SetTools sets Tools field to given value.

### HasTools

`func (o *AgentDetail) HasTools() bool`

HasTools returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AgentDetail) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AgentDetail) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AgentDetail) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AgentDetail) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


