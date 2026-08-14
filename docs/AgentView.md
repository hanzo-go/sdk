# AgentView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComputeRef** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ExecutionMode** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Model** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Runs** | Pointer to **int32** |  | [optional] 
**Schedule** | Pointer to **string** |  | [optional] 
**ServiceAccountId** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Tools** | Pointer to **[]string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewAgentView

`func NewAgentView() *AgentView`

NewAgentView instantiates a new AgentView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentViewWithDefaults

`func NewAgentViewWithDefaults() *AgentView`

NewAgentViewWithDefaults instantiates a new AgentView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComputeRef

`func (o *AgentView) GetComputeRef() string`

GetComputeRef returns the ComputeRef field if non-nil, zero value otherwise.

### GetComputeRefOk

`func (o *AgentView) GetComputeRefOk() (*string, bool)`

GetComputeRefOk returns a tuple with the ComputeRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRef

`func (o *AgentView) SetComputeRef(v string)`

SetComputeRef sets ComputeRef field to given value.

### HasComputeRef

`func (o *AgentView) HasComputeRef() bool`

HasComputeRef returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AgentView) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AgentView) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AgentView) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AgentView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *AgentView) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AgentView) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AgentView) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AgentView) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExecutionMode

`func (o *AgentView) GetExecutionMode() string`

GetExecutionMode returns the ExecutionMode field if non-nil, zero value otherwise.

### GetExecutionModeOk

`func (o *AgentView) GetExecutionModeOk() (*string, bool)`

GetExecutionModeOk returns a tuple with the ExecutionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionMode

`func (o *AgentView) SetExecutionMode(v string)`

SetExecutionMode sets ExecutionMode field to given value.

### HasExecutionMode

`func (o *AgentView) HasExecutionMode() bool`

HasExecutionMode returns a boolean if a field has been set.

### GetId

`func (o *AgentView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AgentView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AgentView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AgentView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModel

`func (o *AgentView) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *AgentView) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *AgentView) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *AgentView) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetName

`func (o *AgentView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AgentView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AgentView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AgentView) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRuns

`func (o *AgentView) GetRuns() int32`

GetRuns returns the Runs field if non-nil, zero value otherwise.

### GetRunsOk

`func (o *AgentView) GetRunsOk() (*int32, bool)`

GetRunsOk returns a tuple with the Runs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuns

`func (o *AgentView) SetRuns(v int32)`

SetRuns sets Runs field to given value.

### HasRuns

`func (o *AgentView) HasRuns() bool`

HasRuns returns a boolean if a field has been set.

### GetSchedule

`func (o *AgentView) GetSchedule() string`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *AgentView) GetScheduleOk() (*string, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *AgentView) SetSchedule(v string)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *AgentView) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetServiceAccountId

`func (o *AgentView) GetServiceAccountId() string`

GetServiceAccountId returns the ServiceAccountId field if non-nil, zero value otherwise.

### GetServiceAccountIdOk

`func (o *AgentView) GetServiceAccountIdOk() (*string, bool)`

GetServiceAccountIdOk returns a tuple with the ServiceAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountId

`func (o *AgentView) SetServiceAccountId(v string)`

SetServiceAccountId sets ServiceAccountId field to given value.

### HasServiceAccountId

`func (o *AgentView) HasServiceAccountId() bool`

HasServiceAccountId returns a boolean if a field has been set.

### GetStatus

`func (o *AgentView) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AgentView) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AgentView) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AgentView) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTools

`func (o *AgentView) GetTools() []string`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *AgentView) GetToolsOk() (*[]string, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *AgentView) SetTools(v []string)`

SetTools sets Tools field to given value.

### HasTools

`func (o *AgentView) HasTools() bool`

HasTools returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AgentView) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AgentView) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AgentView) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AgentView) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


