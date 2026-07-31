# CloudCreateAgentIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComputeRef** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ExecutionMode** | Pointer to **string** |  | [optional] 
**Instructions** | Pointer to **string** |  | [optional] 
**Model** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Schedule** | Pointer to **string** |  | [optional] 
**ServiceAccountId** | Pointer to **string** |  | [optional] 
**Tools** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCloudCreateAgentIn

`func NewCloudCreateAgentIn() *CloudCreateAgentIn`

NewCloudCreateAgentIn instantiates a new CloudCreateAgentIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudCreateAgentInWithDefaults

`func NewCloudCreateAgentInWithDefaults() *CloudCreateAgentIn`

NewCloudCreateAgentInWithDefaults instantiates a new CloudCreateAgentIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComputeRef

`func (o *CloudCreateAgentIn) GetComputeRef() string`

GetComputeRef returns the ComputeRef field if non-nil, zero value otherwise.

### GetComputeRefOk

`func (o *CloudCreateAgentIn) GetComputeRefOk() (*string, bool)`

GetComputeRefOk returns a tuple with the ComputeRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRef

`func (o *CloudCreateAgentIn) SetComputeRef(v string)`

SetComputeRef sets ComputeRef field to given value.

### HasComputeRef

`func (o *CloudCreateAgentIn) HasComputeRef() bool`

HasComputeRef returns a boolean if a field has been set.

### GetDescription

`func (o *CloudCreateAgentIn) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CloudCreateAgentIn) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CloudCreateAgentIn) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CloudCreateAgentIn) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExecutionMode

`func (o *CloudCreateAgentIn) GetExecutionMode() string`

GetExecutionMode returns the ExecutionMode field if non-nil, zero value otherwise.

### GetExecutionModeOk

`func (o *CloudCreateAgentIn) GetExecutionModeOk() (*string, bool)`

GetExecutionModeOk returns a tuple with the ExecutionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionMode

`func (o *CloudCreateAgentIn) SetExecutionMode(v string)`

SetExecutionMode sets ExecutionMode field to given value.

### HasExecutionMode

`func (o *CloudCreateAgentIn) HasExecutionMode() bool`

HasExecutionMode returns a boolean if a field has been set.

### GetInstructions

`func (o *CloudCreateAgentIn) GetInstructions() string`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *CloudCreateAgentIn) GetInstructionsOk() (*string, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *CloudCreateAgentIn) SetInstructions(v string)`

SetInstructions sets Instructions field to given value.

### HasInstructions

`func (o *CloudCreateAgentIn) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.

### GetModel

`func (o *CloudCreateAgentIn) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *CloudCreateAgentIn) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *CloudCreateAgentIn) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *CloudCreateAgentIn) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetName

`func (o *CloudCreateAgentIn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudCreateAgentIn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudCreateAgentIn) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudCreateAgentIn) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSchedule

`func (o *CloudCreateAgentIn) GetSchedule() string`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *CloudCreateAgentIn) GetScheduleOk() (*string, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *CloudCreateAgentIn) SetSchedule(v string)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *CloudCreateAgentIn) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetServiceAccountId

`func (o *CloudCreateAgentIn) GetServiceAccountId() string`

GetServiceAccountId returns the ServiceAccountId field if non-nil, zero value otherwise.

### GetServiceAccountIdOk

`func (o *CloudCreateAgentIn) GetServiceAccountIdOk() (*string, bool)`

GetServiceAccountIdOk returns a tuple with the ServiceAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountId

`func (o *CloudCreateAgentIn) SetServiceAccountId(v string)`

SetServiceAccountId sets ServiceAccountId field to given value.

### HasServiceAccountId

`func (o *CloudCreateAgentIn) HasServiceAccountId() bool`

HasServiceAccountId returns a boolean if a field has been set.

### GetTools

`func (o *CloudCreateAgentIn) GetTools() []string`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *CloudCreateAgentIn) GetToolsOk() (*[]string, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *CloudCreateAgentIn) SetTools(v []string)`

SetTools sets Tools field to given value.

### HasTools

`func (o *CloudCreateAgentIn) HasTools() bool`

HasTools returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


