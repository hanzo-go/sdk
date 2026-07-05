# CloudAgentsRegisterSessionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Agent** | **string** |  | 
**Actor** | Pointer to **string** | Defaults to the validated principal when omitted. | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** | Defaults to running. | [optional] 
**ParentSessionId** | Pointer to **string** | Links this session as a subagent child of an existing same-org session (inherits its rootSessionId). | [optional] 
**TaskWorkflowId** | Pointer to **string** |  | [optional] 
**TaskRunId** | Pointer to **string** |  | [optional] 

## Methods

### NewCloudAgentsRegisterSessionRequest

`func NewCloudAgentsRegisterSessionRequest(agent string, ) *CloudAgentsRegisterSessionRequest`

NewCloudAgentsRegisterSessionRequest instantiates a new CloudAgentsRegisterSessionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudAgentsRegisterSessionRequestWithDefaults

`func NewCloudAgentsRegisterSessionRequestWithDefaults() *CloudAgentsRegisterSessionRequest`

NewCloudAgentsRegisterSessionRequestWithDefaults instantiates a new CloudAgentsRegisterSessionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgent

`func (o *CloudAgentsRegisterSessionRequest) GetAgent() string`

GetAgent returns the Agent field if non-nil, zero value otherwise.

### GetAgentOk

`func (o *CloudAgentsRegisterSessionRequest) GetAgentOk() (*string, bool)`

GetAgentOk returns a tuple with the Agent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgent

`func (o *CloudAgentsRegisterSessionRequest) SetAgent(v string)`

SetAgent sets Agent field to given value.


### GetActor

`func (o *CloudAgentsRegisterSessionRequest) GetActor() string`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *CloudAgentsRegisterSessionRequest) GetActorOk() (*string, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *CloudAgentsRegisterSessionRequest) SetActor(v string)`

SetActor sets Actor field to given value.

### HasActor

`func (o *CloudAgentsRegisterSessionRequest) HasActor() bool`

HasActor returns a boolean if a field has been set.

### GetTitle

`func (o *CloudAgentsRegisterSessionRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CloudAgentsRegisterSessionRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CloudAgentsRegisterSessionRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CloudAgentsRegisterSessionRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetStatus

`func (o *CloudAgentsRegisterSessionRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudAgentsRegisterSessionRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudAgentsRegisterSessionRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudAgentsRegisterSessionRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetParentSessionId

`func (o *CloudAgentsRegisterSessionRequest) GetParentSessionId() string`

GetParentSessionId returns the ParentSessionId field if non-nil, zero value otherwise.

### GetParentSessionIdOk

`func (o *CloudAgentsRegisterSessionRequest) GetParentSessionIdOk() (*string, bool)`

GetParentSessionIdOk returns a tuple with the ParentSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentSessionId

`func (o *CloudAgentsRegisterSessionRequest) SetParentSessionId(v string)`

SetParentSessionId sets ParentSessionId field to given value.

### HasParentSessionId

`func (o *CloudAgentsRegisterSessionRequest) HasParentSessionId() bool`

HasParentSessionId returns a boolean if a field has been set.

### GetTaskWorkflowId

`func (o *CloudAgentsRegisterSessionRequest) GetTaskWorkflowId() string`

GetTaskWorkflowId returns the TaskWorkflowId field if non-nil, zero value otherwise.

### GetTaskWorkflowIdOk

`func (o *CloudAgentsRegisterSessionRequest) GetTaskWorkflowIdOk() (*string, bool)`

GetTaskWorkflowIdOk returns a tuple with the TaskWorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskWorkflowId

`func (o *CloudAgentsRegisterSessionRequest) SetTaskWorkflowId(v string)`

SetTaskWorkflowId sets TaskWorkflowId field to given value.

### HasTaskWorkflowId

`func (o *CloudAgentsRegisterSessionRequest) HasTaskWorkflowId() bool`

HasTaskWorkflowId returns a boolean if a field has been set.

### GetTaskRunId

`func (o *CloudAgentsRegisterSessionRequest) GetTaskRunId() string`

GetTaskRunId returns the TaskRunId field if non-nil, zero value otherwise.

### GetTaskRunIdOk

`func (o *CloudAgentsRegisterSessionRequest) GetTaskRunIdOk() (*string, bool)`

GetTaskRunIdOk returns a tuple with the TaskRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskRunId

`func (o *CloudAgentsRegisterSessionRequest) SetTaskRunId(v string)`

SetTaskRunId sets TaskRunId field to given value.

### HasTaskRunId

`func (o *CloudAgentsRegisterSessionRequest) HasTaskRunId() bool`

HasTaskRunId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


