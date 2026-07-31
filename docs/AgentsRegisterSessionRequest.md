# AgentsRegisterSessionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Agent** | **string** | Agent label (max 128 chars). | 
**Actor** | Pointer to **string** | Defaults to the validated principal (org/sub). | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** | Defaults to running. | [optional] 
**ParentSessionId** | Pointer to **string** | Parent session id in the same org (for subagent linkage). | [optional] 
**TaskWorkflowId** | Pointer to **string** |  | [optional] 
**TaskRunId** | Pointer to **string** |  | [optional] 

## Methods

### NewAgentsRegisterSessionRequest

`func NewAgentsRegisterSessionRequest(agent string, ) *AgentsRegisterSessionRequest`

NewAgentsRegisterSessionRequest instantiates a new AgentsRegisterSessionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentsRegisterSessionRequestWithDefaults

`func NewAgentsRegisterSessionRequestWithDefaults() *AgentsRegisterSessionRequest`

NewAgentsRegisterSessionRequestWithDefaults instantiates a new AgentsRegisterSessionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgent

`func (o *AgentsRegisterSessionRequest) GetAgent() string`

GetAgent returns the Agent field if non-nil, zero value otherwise.

### GetAgentOk

`func (o *AgentsRegisterSessionRequest) GetAgentOk() (*string, bool)`

GetAgentOk returns a tuple with the Agent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgent

`func (o *AgentsRegisterSessionRequest) SetAgent(v string)`

SetAgent sets Agent field to given value.


### GetActor

`func (o *AgentsRegisterSessionRequest) GetActor() string`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *AgentsRegisterSessionRequest) GetActorOk() (*string, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *AgentsRegisterSessionRequest) SetActor(v string)`

SetActor sets Actor field to given value.

### HasActor

`func (o *AgentsRegisterSessionRequest) HasActor() bool`

HasActor returns a boolean if a field has been set.

### GetTitle

`func (o *AgentsRegisterSessionRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AgentsRegisterSessionRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AgentsRegisterSessionRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AgentsRegisterSessionRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetStatus

`func (o *AgentsRegisterSessionRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AgentsRegisterSessionRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AgentsRegisterSessionRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AgentsRegisterSessionRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetParentSessionId

`func (o *AgentsRegisterSessionRequest) GetParentSessionId() string`

GetParentSessionId returns the ParentSessionId field if non-nil, zero value otherwise.

### GetParentSessionIdOk

`func (o *AgentsRegisterSessionRequest) GetParentSessionIdOk() (*string, bool)`

GetParentSessionIdOk returns a tuple with the ParentSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentSessionId

`func (o *AgentsRegisterSessionRequest) SetParentSessionId(v string)`

SetParentSessionId sets ParentSessionId field to given value.

### HasParentSessionId

`func (o *AgentsRegisterSessionRequest) HasParentSessionId() bool`

HasParentSessionId returns a boolean if a field has been set.

### GetTaskWorkflowId

`func (o *AgentsRegisterSessionRequest) GetTaskWorkflowId() string`

GetTaskWorkflowId returns the TaskWorkflowId field if non-nil, zero value otherwise.

### GetTaskWorkflowIdOk

`func (o *AgentsRegisterSessionRequest) GetTaskWorkflowIdOk() (*string, bool)`

GetTaskWorkflowIdOk returns a tuple with the TaskWorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskWorkflowId

`func (o *AgentsRegisterSessionRequest) SetTaskWorkflowId(v string)`

SetTaskWorkflowId sets TaskWorkflowId field to given value.

### HasTaskWorkflowId

`func (o *AgentsRegisterSessionRequest) HasTaskWorkflowId() bool`

HasTaskWorkflowId returns a boolean if a field has been set.

### GetTaskRunId

`func (o *AgentsRegisterSessionRequest) GetTaskRunId() string`

GetTaskRunId returns the TaskRunId field if non-nil, zero value otherwise.

### GetTaskRunIdOk

`func (o *AgentsRegisterSessionRequest) GetTaskRunIdOk() (*string, bool)`

GetTaskRunIdOk returns a tuple with the TaskRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskRunId

`func (o *AgentsRegisterSessionRequest) SetTaskRunId(v string)`

SetTaskRunId sets TaskRunId field to given value.

### HasTaskRunId

`func (o *AgentsRegisterSessionRequest) HasTaskRunId() bool`

HasTaskRunId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


