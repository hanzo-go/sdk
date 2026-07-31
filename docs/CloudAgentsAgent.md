# CloudAgentsAgent

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

## Methods

### NewCloudAgentsAgent

`func NewCloudAgentsAgent() *CloudAgentsAgent`

NewCloudAgentsAgent instantiates a new CloudAgentsAgent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudAgentsAgentWithDefaults

`func NewCloudAgentsAgentWithDefaults() *CloudAgentsAgent`

NewCloudAgentsAgentWithDefaults instantiates a new CloudAgentsAgent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CloudAgentsAgent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudAgentsAgent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudAgentsAgent) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudAgentsAgent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CloudAgentsAgent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudAgentsAgent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudAgentsAgent) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudAgentsAgent) HasName() bool`

HasName returns a boolean if a field has been set.

### GetModel

`func (o *CloudAgentsAgent) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *CloudAgentsAgent) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *CloudAgentsAgent) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *CloudAgentsAgent) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetDescription

`func (o *CloudAgentsAgent) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CloudAgentsAgent) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CloudAgentsAgent) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CloudAgentsAgent) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTools

`func (o *CloudAgentsAgent) GetTools() []string`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *CloudAgentsAgent) GetToolsOk() (*[]string, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *CloudAgentsAgent) SetTools(v []string)`

SetTools sets Tools field to given value.

### HasTools

`func (o *CloudAgentsAgent) HasTools() bool`

HasTools returns a boolean if a field has been set.

### GetStatus

`func (o *CloudAgentsAgent) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudAgentsAgent) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudAgentsAgent) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudAgentsAgent) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetExecutionMode

`func (o *CloudAgentsAgent) GetExecutionMode() string`

GetExecutionMode returns the ExecutionMode field if non-nil, zero value otherwise.

### GetExecutionModeOk

`func (o *CloudAgentsAgent) GetExecutionModeOk() (*string, bool)`

GetExecutionModeOk returns a tuple with the ExecutionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionMode

`func (o *CloudAgentsAgent) SetExecutionMode(v string)`

SetExecutionMode sets ExecutionMode field to given value.

### HasExecutionMode

`func (o *CloudAgentsAgent) HasExecutionMode() bool`

HasExecutionMode returns a boolean if a field has been set.

### GetSchedule

`func (o *CloudAgentsAgent) GetSchedule() string`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *CloudAgentsAgent) GetScheduleOk() (*string, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *CloudAgentsAgent) SetSchedule(v string)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *CloudAgentsAgent) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetComputeRef

`func (o *CloudAgentsAgent) GetComputeRef() string`

GetComputeRef returns the ComputeRef field if non-nil, zero value otherwise.

### GetComputeRefOk

`func (o *CloudAgentsAgent) GetComputeRefOk() (*string, bool)`

GetComputeRefOk returns a tuple with the ComputeRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRef

`func (o *CloudAgentsAgent) SetComputeRef(v string)`

SetComputeRef sets ComputeRef field to given value.

### HasComputeRef

`func (o *CloudAgentsAgent) HasComputeRef() bool`

HasComputeRef returns a boolean if a field has been set.

### GetServiceAccountId

`func (o *CloudAgentsAgent) GetServiceAccountId() string`

GetServiceAccountId returns the ServiceAccountId field if non-nil, zero value otherwise.

### GetServiceAccountIdOk

`func (o *CloudAgentsAgent) GetServiceAccountIdOk() (*string, bool)`

GetServiceAccountIdOk returns a tuple with the ServiceAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountId

`func (o *CloudAgentsAgent) SetServiceAccountId(v string)`

SetServiceAccountId sets ServiceAccountId field to given value.

### HasServiceAccountId

`func (o *CloudAgentsAgent) HasServiceAccountId() bool`

HasServiceAccountId returns a boolean if a field has been set.

### GetRuns

`func (o *CloudAgentsAgent) GetRuns() int32`

GetRuns returns the Runs field if non-nil, zero value otherwise.

### GetRunsOk

`func (o *CloudAgentsAgent) GetRunsOk() (*int32, bool)`

GetRunsOk returns a tuple with the Runs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuns

`func (o *CloudAgentsAgent) SetRuns(v int32)`

SetRuns sets Runs field to given value.

### HasRuns

`func (o *CloudAgentsAgent) HasRuns() bool`

HasRuns returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CloudAgentsAgent) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CloudAgentsAgent) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CloudAgentsAgent) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CloudAgentsAgent) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CloudAgentsAgent) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CloudAgentsAgent) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CloudAgentsAgent) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CloudAgentsAgent) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


