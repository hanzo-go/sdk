# AgentsCreateAgentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Org-unique name; must match ^[A-Za-z0-9][A-Za-z0-9._-]{0,63}$. | 
**Model** | Pointer to **string** | Served model id. Validated against the gateway catalog; omit to use the deployment default. | [optional] 
**Instructions** | Pointer to **string** | System prompt (capped at 32 KiB). | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Tools** | Pointer to **[]string** |  | [optional] 
**ExecutionMode** | Pointer to **string** | Defaults to one-shot when empty. | [optional] 
**Schedule** | Pointer to **string** | Required 5-field cron for a long-running agent; cleared for one-shot. | [optional] 
**ComputeRef** | Pointer to **string** |  | [optional] 
**ServiceAccountId** | Pointer to **string** |  | [optional] 

## Methods

### NewAgentsCreateAgentRequest

`func NewAgentsCreateAgentRequest(name string, ) *AgentsCreateAgentRequest`

NewAgentsCreateAgentRequest instantiates a new AgentsCreateAgentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentsCreateAgentRequestWithDefaults

`func NewAgentsCreateAgentRequestWithDefaults() *AgentsCreateAgentRequest`

NewAgentsCreateAgentRequestWithDefaults instantiates a new AgentsCreateAgentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AgentsCreateAgentRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AgentsCreateAgentRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AgentsCreateAgentRequest) SetName(v string)`

SetName sets Name field to given value.


### GetModel

`func (o *AgentsCreateAgentRequest) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *AgentsCreateAgentRequest) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *AgentsCreateAgentRequest) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *AgentsCreateAgentRequest) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetInstructions

`func (o *AgentsCreateAgentRequest) GetInstructions() string`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *AgentsCreateAgentRequest) GetInstructionsOk() (*string, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *AgentsCreateAgentRequest) SetInstructions(v string)`

SetInstructions sets Instructions field to given value.

### HasInstructions

`func (o *AgentsCreateAgentRequest) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.

### GetDescription

`func (o *AgentsCreateAgentRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AgentsCreateAgentRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AgentsCreateAgentRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AgentsCreateAgentRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTools

`func (o *AgentsCreateAgentRequest) GetTools() []string`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *AgentsCreateAgentRequest) GetToolsOk() (*[]string, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *AgentsCreateAgentRequest) SetTools(v []string)`

SetTools sets Tools field to given value.

### HasTools

`func (o *AgentsCreateAgentRequest) HasTools() bool`

HasTools returns a boolean if a field has been set.

### GetExecutionMode

`func (o *AgentsCreateAgentRequest) GetExecutionMode() string`

GetExecutionMode returns the ExecutionMode field if non-nil, zero value otherwise.

### GetExecutionModeOk

`func (o *AgentsCreateAgentRequest) GetExecutionModeOk() (*string, bool)`

GetExecutionModeOk returns a tuple with the ExecutionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionMode

`func (o *AgentsCreateAgentRequest) SetExecutionMode(v string)`

SetExecutionMode sets ExecutionMode field to given value.

### HasExecutionMode

`func (o *AgentsCreateAgentRequest) HasExecutionMode() bool`

HasExecutionMode returns a boolean if a field has been set.

### GetSchedule

`func (o *AgentsCreateAgentRequest) GetSchedule() string`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *AgentsCreateAgentRequest) GetScheduleOk() (*string, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *AgentsCreateAgentRequest) SetSchedule(v string)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *AgentsCreateAgentRequest) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetComputeRef

`func (o *AgentsCreateAgentRequest) GetComputeRef() string`

GetComputeRef returns the ComputeRef field if non-nil, zero value otherwise.

### GetComputeRefOk

`func (o *AgentsCreateAgentRequest) GetComputeRefOk() (*string, bool)`

GetComputeRefOk returns a tuple with the ComputeRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRef

`func (o *AgentsCreateAgentRequest) SetComputeRef(v string)`

SetComputeRef sets ComputeRef field to given value.

### HasComputeRef

`func (o *AgentsCreateAgentRequest) HasComputeRef() bool`

HasComputeRef returns a boolean if a field has been set.

### GetServiceAccountId

`func (o *AgentsCreateAgentRequest) GetServiceAccountId() string`

GetServiceAccountId returns the ServiceAccountId field if non-nil, zero value otherwise.

### GetServiceAccountIdOk

`func (o *AgentsCreateAgentRequest) GetServiceAccountIdOk() (*string, bool)`

GetServiceAccountIdOk returns a tuple with the ServiceAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountId

`func (o *AgentsCreateAgentRequest) SetServiceAccountId(v string)`

SetServiceAccountId sets ServiceAccountId field to given value.

### HasServiceAccountId

`func (o *AgentsCreateAgentRequest) HasServiceAccountId() bool`

HasServiceAccountId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


