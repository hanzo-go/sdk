# CloudAgentDetail

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
**RecentRuns** | Pointer to [**[]CloudAgentRunView**](CloudAgentRunView.md) |  | [optional] 
**Runs** | Pointer to **int32** |  | [optional] 
**Schedule** | Pointer to **string** |  | [optional] 
**ServiceAccountId** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Tools** | Pointer to **[]string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewCloudAgentDetail

`func NewCloudAgentDetail() *CloudAgentDetail`

NewCloudAgentDetail instantiates a new CloudAgentDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudAgentDetailWithDefaults

`func NewCloudAgentDetailWithDefaults() *CloudAgentDetail`

NewCloudAgentDetailWithDefaults instantiates a new CloudAgentDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComputeRef

`func (o *CloudAgentDetail) GetComputeRef() string`

GetComputeRef returns the ComputeRef field if non-nil, zero value otherwise.

### GetComputeRefOk

`func (o *CloudAgentDetail) GetComputeRefOk() (*string, bool)`

GetComputeRefOk returns a tuple with the ComputeRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRef

`func (o *CloudAgentDetail) SetComputeRef(v string)`

SetComputeRef sets ComputeRef field to given value.

### HasComputeRef

`func (o *CloudAgentDetail) HasComputeRef() bool`

HasComputeRef returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CloudAgentDetail) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CloudAgentDetail) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CloudAgentDetail) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CloudAgentDetail) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *CloudAgentDetail) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CloudAgentDetail) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CloudAgentDetail) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CloudAgentDetail) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExecutionMode

`func (o *CloudAgentDetail) GetExecutionMode() string`

GetExecutionMode returns the ExecutionMode field if non-nil, zero value otherwise.

### GetExecutionModeOk

`func (o *CloudAgentDetail) GetExecutionModeOk() (*string, bool)`

GetExecutionModeOk returns a tuple with the ExecutionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionMode

`func (o *CloudAgentDetail) SetExecutionMode(v string)`

SetExecutionMode sets ExecutionMode field to given value.

### HasExecutionMode

`func (o *CloudAgentDetail) HasExecutionMode() bool`

HasExecutionMode returns a boolean if a field has been set.

### GetId

`func (o *CloudAgentDetail) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudAgentDetail) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudAgentDetail) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudAgentDetail) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInstructions

`func (o *CloudAgentDetail) GetInstructions() string`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *CloudAgentDetail) GetInstructionsOk() (*string, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *CloudAgentDetail) SetInstructions(v string)`

SetInstructions sets Instructions field to given value.

### HasInstructions

`func (o *CloudAgentDetail) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.

### GetModel

`func (o *CloudAgentDetail) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *CloudAgentDetail) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *CloudAgentDetail) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *CloudAgentDetail) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetName

`func (o *CloudAgentDetail) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudAgentDetail) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudAgentDetail) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudAgentDetail) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRecentRuns

`func (o *CloudAgentDetail) GetRecentRuns() []CloudAgentRunView`

GetRecentRuns returns the RecentRuns field if non-nil, zero value otherwise.

### GetRecentRunsOk

`func (o *CloudAgentDetail) GetRecentRunsOk() (*[]CloudAgentRunView, bool)`

GetRecentRunsOk returns a tuple with the RecentRuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecentRuns

`func (o *CloudAgentDetail) SetRecentRuns(v []CloudAgentRunView)`

SetRecentRuns sets RecentRuns field to given value.

### HasRecentRuns

`func (o *CloudAgentDetail) HasRecentRuns() bool`

HasRecentRuns returns a boolean if a field has been set.

### GetRuns

`func (o *CloudAgentDetail) GetRuns() int32`

GetRuns returns the Runs field if non-nil, zero value otherwise.

### GetRunsOk

`func (o *CloudAgentDetail) GetRunsOk() (*int32, bool)`

GetRunsOk returns a tuple with the Runs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuns

`func (o *CloudAgentDetail) SetRuns(v int32)`

SetRuns sets Runs field to given value.

### HasRuns

`func (o *CloudAgentDetail) HasRuns() bool`

HasRuns returns a boolean if a field has been set.

### GetSchedule

`func (o *CloudAgentDetail) GetSchedule() string`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *CloudAgentDetail) GetScheduleOk() (*string, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *CloudAgentDetail) SetSchedule(v string)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *CloudAgentDetail) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetServiceAccountId

`func (o *CloudAgentDetail) GetServiceAccountId() string`

GetServiceAccountId returns the ServiceAccountId field if non-nil, zero value otherwise.

### GetServiceAccountIdOk

`func (o *CloudAgentDetail) GetServiceAccountIdOk() (*string, bool)`

GetServiceAccountIdOk returns a tuple with the ServiceAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountId

`func (o *CloudAgentDetail) SetServiceAccountId(v string)`

SetServiceAccountId sets ServiceAccountId field to given value.

### HasServiceAccountId

`func (o *CloudAgentDetail) HasServiceAccountId() bool`

HasServiceAccountId returns a boolean if a field has been set.

### GetStatus

`func (o *CloudAgentDetail) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudAgentDetail) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudAgentDetail) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudAgentDetail) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTools

`func (o *CloudAgentDetail) GetTools() []string`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *CloudAgentDetail) GetToolsOk() (*[]string, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *CloudAgentDetail) SetTools(v []string)`

SetTools sets Tools field to given value.

### HasTools

`func (o *CloudAgentDetail) HasTools() bool`

HasTools returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CloudAgentDetail) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CloudAgentDetail) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CloudAgentDetail) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CloudAgentDetail) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


