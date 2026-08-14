# UpdateAgentIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComputeRef** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ExecutionMode** | Pointer to **string** |  | [optional] 
**Instructions** | Pointer to **string** |  | [optional] 
**Model** | Pointer to **string** |  | [optional] 
**Ref** | Pointer to **string** | Ref is the agent to update — its public id or org-unique name, from the path. | [optional] 
**Schedule** | Pointer to **string** |  | [optional] 
**ServiceAccountId** | Pointer to **string** |  | [optional] 
**Tools** | Pointer to **[]string** |  | [optional] 

## Methods

### NewUpdateAgentIn

`func NewUpdateAgentIn() *UpdateAgentIn`

NewUpdateAgentIn instantiates a new UpdateAgentIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAgentInWithDefaults

`func NewUpdateAgentInWithDefaults() *UpdateAgentIn`

NewUpdateAgentInWithDefaults instantiates a new UpdateAgentIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComputeRef

`func (o *UpdateAgentIn) GetComputeRef() string`

GetComputeRef returns the ComputeRef field if non-nil, zero value otherwise.

### GetComputeRefOk

`func (o *UpdateAgentIn) GetComputeRefOk() (*string, bool)`

GetComputeRefOk returns a tuple with the ComputeRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRef

`func (o *UpdateAgentIn) SetComputeRef(v string)`

SetComputeRef sets ComputeRef field to given value.

### HasComputeRef

`func (o *UpdateAgentIn) HasComputeRef() bool`

HasComputeRef returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateAgentIn) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateAgentIn) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateAgentIn) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateAgentIn) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExecutionMode

`func (o *UpdateAgentIn) GetExecutionMode() string`

GetExecutionMode returns the ExecutionMode field if non-nil, zero value otherwise.

### GetExecutionModeOk

`func (o *UpdateAgentIn) GetExecutionModeOk() (*string, bool)`

GetExecutionModeOk returns a tuple with the ExecutionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionMode

`func (o *UpdateAgentIn) SetExecutionMode(v string)`

SetExecutionMode sets ExecutionMode field to given value.

### HasExecutionMode

`func (o *UpdateAgentIn) HasExecutionMode() bool`

HasExecutionMode returns a boolean if a field has been set.

### GetInstructions

`func (o *UpdateAgentIn) GetInstructions() string`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *UpdateAgentIn) GetInstructionsOk() (*string, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *UpdateAgentIn) SetInstructions(v string)`

SetInstructions sets Instructions field to given value.

### HasInstructions

`func (o *UpdateAgentIn) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.

### GetModel

`func (o *UpdateAgentIn) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *UpdateAgentIn) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *UpdateAgentIn) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *UpdateAgentIn) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetRef

`func (o *UpdateAgentIn) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *UpdateAgentIn) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *UpdateAgentIn) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *UpdateAgentIn) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetSchedule

`func (o *UpdateAgentIn) GetSchedule() string`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *UpdateAgentIn) GetScheduleOk() (*string, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *UpdateAgentIn) SetSchedule(v string)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *UpdateAgentIn) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetServiceAccountId

`func (o *UpdateAgentIn) GetServiceAccountId() string`

GetServiceAccountId returns the ServiceAccountId field if non-nil, zero value otherwise.

### GetServiceAccountIdOk

`func (o *UpdateAgentIn) GetServiceAccountIdOk() (*string, bool)`

GetServiceAccountIdOk returns a tuple with the ServiceAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountId

`func (o *UpdateAgentIn) SetServiceAccountId(v string)`

SetServiceAccountId sets ServiceAccountId field to given value.

### HasServiceAccountId

`func (o *UpdateAgentIn) HasServiceAccountId() bool`

HasServiceAccountId returns a boolean if a field has been set.

### GetTools

`func (o *UpdateAgentIn) GetTools() []string`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *UpdateAgentIn) GetToolsOk() (*[]string, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *UpdateAgentIn) SetTools(v []string)`

SetTools sets Tools field to given value.

### HasTools

`func (o *UpdateAgentIn) HasTools() bool`

HasTools returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


