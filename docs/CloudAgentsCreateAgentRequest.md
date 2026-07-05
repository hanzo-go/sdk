# CloudAgentsCreateAgentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Org-unique name matching ^[A-Za-z0-9][A-Za-z0-9._-]{0,63}$. | 
**Model** | **string** |  | 
**Instructions** | Pointer to **string** | System prompt (max 32 KiB). | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Tools** | Pointer to **[]string** |  | [optional] 
**ExecutionMode** | Pointer to **string** | Defaults to one-shot. | [optional] 
**Schedule** | Pointer to **string** | 5-field cron; required when executionMode is long-running. | [optional] 
**ComputeRef** | Pointer to **string** |  | [optional] 
**ServiceAccountId** | Pointer to **string** |  | [optional] 

## Methods

### NewCloudAgentsCreateAgentRequest

`func NewCloudAgentsCreateAgentRequest(name string, model string, ) *CloudAgentsCreateAgentRequest`

NewCloudAgentsCreateAgentRequest instantiates a new CloudAgentsCreateAgentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudAgentsCreateAgentRequestWithDefaults

`func NewCloudAgentsCreateAgentRequestWithDefaults() *CloudAgentsCreateAgentRequest`

NewCloudAgentsCreateAgentRequestWithDefaults instantiates a new CloudAgentsCreateAgentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CloudAgentsCreateAgentRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudAgentsCreateAgentRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudAgentsCreateAgentRequest) SetName(v string)`

SetName sets Name field to given value.


### GetModel

`func (o *CloudAgentsCreateAgentRequest) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *CloudAgentsCreateAgentRequest) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *CloudAgentsCreateAgentRequest) SetModel(v string)`

SetModel sets Model field to given value.


### GetInstructions

`func (o *CloudAgentsCreateAgentRequest) GetInstructions() string`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *CloudAgentsCreateAgentRequest) GetInstructionsOk() (*string, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *CloudAgentsCreateAgentRequest) SetInstructions(v string)`

SetInstructions sets Instructions field to given value.

### HasInstructions

`func (o *CloudAgentsCreateAgentRequest) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.

### GetDescription

`func (o *CloudAgentsCreateAgentRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CloudAgentsCreateAgentRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CloudAgentsCreateAgentRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CloudAgentsCreateAgentRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTools

`func (o *CloudAgentsCreateAgentRequest) GetTools() []string`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *CloudAgentsCreateAgentRequest) GetToolsOk() (*[]string, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *CloudAgentsCreateAgentRequest) SetTools(v []string)`

SetTools sets Tools field to given value.

### HasTools

`func (o *CloudAgentsCreateAgentRequest) HasTools() bool`

HasTools returns a boolean if a field has been set.

### GetExecutionMode

`func (o *CloudAgentsCreateAgentRequest) GetExecutionMode() string`

GetExecutionMode returns the ExecutionMode field if non-nil, zero value otherwise.

### GetExecutionModeOk

`func (o *CloudAgentsCreateAgentRequest) GetExecutionModeOk() (*string, bool)`

GetExecutionModeOk returns a tuple with the ExecutionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionMode

`func (o *CloudAgentsCreateAgentRequest) SetExecutionMode(v string)`

SetExecutionMode sets ExecutionMode field to given value.

### HasExecutionMode

`func (o *CloudAgentsCreateAgentRequest) HasExecutionMode() bool`

HasExecutionMode returns a boolean if a field has been set.

### GetSchedule

`func (o *CloudAgentsCreateAgentRequest) GetSchedule() string`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *CloudAgentsCreateAgentRequest) GetScheduleOk() (*string, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *CloudAgentsCreateAgentRequest) SetSchedule(v string)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *CloudAgentsCreateAgentRequest) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetComputeRef

`func (o *CloudAgentsCreateAgentRequest) GetComputeRef() string`

GetComputeRef returns the ComputeRef field if non-nil, zero value otherwise.

### GetComputeRefOk

`func (o *CloudAgentsCreateAgentRequest) GetComputeRefOk() (*string, bool)`

GetComputeRefOk returns a tuple with the ComputeRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRef

`func (o *CloudAgentsCreateAgentRequest) SetComputeRef(v string)`

SetComputeRef sets ComputeRef field to given value.

### HasComputeRef

`func (o *CloudAgentsCreateAgentRequest) HasComputeRef() bool`

HasComputeRef returns a boolean if a field has been set.

### GetServiceAccountId

`func (o *CloudAgentsCreateAgentRequest) GetServiceAccountId() string`

GetServiceAccountId returns the ServiceAccountId field if non-nil, zero value otherwise.

### GetServiceAccountIdOk

`func (o *CloudAgentsCreateAgentRequest) GetServiceAccountIdOk() (*string, bool)`

GetServiceAccountIdOk returns a tuple with the ServiceAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountId

`func (o *CloudAgentsCreateAgentRequest) SetServiceAccountId(v string)`

SetServiceAccountId sets ServiceAccountId field to given value.

### HasServiceAccountId

`func (o *CloudAgentsCreateAgentRequest) HasServiceAccountId() bool`

HasServiceAccountId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


