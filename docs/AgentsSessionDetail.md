# AgentsSessionDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Agent** | Pointer to **string** |  | [optional] 
**Actor** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**ParentSessionId** | Pointer to **string** |  | [optional] 
**RootSessionId** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**TaskWorkflowId** | Pointer to **string** |  | [optional] 
**TaskRunId** | Pointer to **string** |  | [optional] 
**Events** | Pointer to **int32** |  | [optional] 
**Children** | Pointer to **int32** | Direct fan-out count. | [optional] 
**StartedAt** | Pointer to **time.Time** |  | [optional] 
**EndedAt** | Pointer to **time.Time** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**ChildSessions** | Pointer to [**[]AgentsSessionView**](AgentsSessionView.md) |  | [optional] 
**RecentEvents** | Pointer to [**[]AgentsEventView**](AgentsEventView.md) |  | [optional] 

## Methods

### NewAgentsSessionDetail

`func NewAgentsSessionDetail() *AgentsSessionDetail`

NewAgentsSessionDetail instantiates a new AgentsSessionDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentsSessionDetailWithDefaults

`func NewAgentsSessionDetailWithDefaults() *AgentsSessionDetail`

NewAgentsSessionDetailWithDefaults instantiates a new AgentsSessionDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AgentsSessionDetail) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AgentsSessionDetail) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AgentsSessionDetail) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AgentsSessionDetail) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAgent

`func (o *AgentsSessionDetail) GetAgent() string`

GetAgent returns the Agent field if non-nil, zero value otherwise.

### GetAgentOk

`func (o *AgentsSessionDetail) GetAgentOk() (*string, bool)`

GetAgentOk returns a tuple with the Agent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgent

`func (o *AgentsSessionDetail) SetAgent(v string)`

SetAgent sets Agent field to given value.

### HasAgent

`func (o *AgentsSessionDetail) HasAgent() bool`

HasAgent returns a boolean if a field has been set.

### GetActor

`func (o *AgentsSessionDetail) GetActor() string`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *AgentsSessionDetail) GetActorOk() (*string, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *AgentsSessionDetail) SetActor(v string)`

SetActor sets Actor field to given value.

### HasActor

`func (o *AgentsSessionDetail) HasActor() bool`

HasActor returns a boolean if a field has been set.

### GetStatus

`func (o *AgentsSessionDetail) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AgentsSessionDetail) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AgentsSessionDetail) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AgentsSessionDetail) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetParentSessionId

`func (o *AgentsSessionDetail) GetParentSessionId() string`

GetParentSessionId returns the ParentSessionId field if non-nil, zero value otherwise.

### GetParentSessionIdOk

`func (o *AgentsSessionDetail) GetParentSessionIdOk() (*string, bool)`

GetParentSessionIdOk returns a tuple with the ParentSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentSessionId

`func (o *AgentsSessionDetail) SetParentSessionId(v string)`

SetParentSessionId sets ParentSessionId field to given value.

### HasParentSessionId

`func (o *AgentsSessionDetail) HasParentSessionId() bool`

HasParentSessionId returns a boolean if a field has been set.

### GetRootSessionId

`func (o *AgentsSessionDetail) GetRootSessionId() string`

GetRootSessionId returns the RootSessionId field if non-nil, zero value otherwise.

### GetRootSessionIdOk

`func (o *AgentsSessionDetail) GetRootSessionIdOk() (*string, bool)`

GetRootSessionIdOk returns a tuple with the RootSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootSessionId

`func (o *AgentsSessionDetail) SetRootSessionId(v string)`

SetRootSessionId sets RootSessionId field to given value.

### HasRootSessionId

`func (o *AgentsSessionDetail) HasRootSessionId() bool`

HasRootSessionId returns a boolean if a field has been set.

### GetTitle

`func (o *AgentsSessionDetail) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AgentsSessionDetail) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AgentsSessionDetail) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AgentsSessionDetail) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetTaskWorkflowId

`func (o *AgentsSessionDetail) GetTaskWorkflowId() string`

GetTaskWorkflowId returns the TaskWorkflowId field if non-nil, zero value otherwise.

### GetTaskWorkflowIdOk

`func (o *AgentsSessionDetail) GetTaskWorkflowIdOk() (*string, bool)`

GetTaskWorkflowIdOk returns a tuple with the TaskWorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskWorkflowId

`func (o *AgentsSessionDetail) SetTaskWorkflowId(v string)`

SetTaskWorkflowId sets TaskWorkflowId field to given value.

### HasTaskWorkflowId

`func (o *AgentsSessionDetail) HasTaskWorkflowId() bool`

HasTaskWorkflowId returns a boolean if a field has been set.

### GetTaskRunId

`func (o *AgentsSessionDetail) GetTaskRunId() string`

GetTaskRunId returns the TaskRunId field if non-nil, zero value otherwise.

### GetTaskRunIdOk

`func (o *AgentsSessionDetail) GetTaskRunIdOk() (*string, bool)`

GetTaskRunIdOk returns a tuple with the TaskRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskRunId

`func (o *AgentsSessionDetail) SetTaskRunId(v string)`

SetTaskRunId sets TaskRunId field to given value.

### HasTaskRunId

`func (o *AgentsSessionDetail) HasTaskRunId() bool`

HasTaskRunId returns a boolean if a field has been set.

### GetEvents

`func (o *AgentsSessionDetail) GetEvents() int32`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *AgentsSessionDetail) GetEventsOk() (*int32, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *AgentsSessionDetail) SetEvents(v int32)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *AgentsSessionDetail) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetChildren

`func (o *AgentsSessionDetail) GetChildren() int32`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *AgentsSessionDetail) GetChildrenOk() (*int32, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *AgentsSessionDetail) SetChildren(v int32)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *AgentsSessionDetail) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### GetStartedAt

`func (o *AgentsSessionDetail) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *AgentsSessionDetail) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *AgentsSessionDetail) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *AgentsSessionDetail) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetEndedAt

`func (o *AgentsSessionDetail) GetEndedAt() time.Time`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *AgentsSessionDetail) GetEndedAtOk() (*time.Time, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *AgentsSessionDetail) SetEndedAt(v time.Time)`

SetEndedAt sets EndedAt field to given value.

### HasEndedAt

`func (o *AgentsSessionDetail) HasEndedAt() bool`

HasEndedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AgentsSessionDetail) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AgentsSessionDetail) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AgentsSessionDetail) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AgentsSessionDetail) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AgentsSessionDetail) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AgentsSessionDetail) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AgentsSessionDetail) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AgentsSessionDetail) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetChildSessions

`func (o *AgentsSessionDetail) GetChildSessions() []AgentsSessionView`

GetChildSessions returns the ChildSessions field if non-nil, zero value otherwise.

### GetChildSessionsOk

`func (o *AgentsSessionDetail) GetChildSessionsOk() (*[]AgentsSessionView, bool)`

GetChildSessionsOk returns a tuple with the ChildSessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildSessions

`func (o *AgentsSessionDetail) SetChildSessions(v []AgentsSessionView)`

SetChildSessions sets ChildSessions field to given value.

### HasChildSessions

`func (o *AgentsSessionDetail) HasChildSessions() bool`

HasChildSessions returns a boolean if a field has been set.

### GetRecentEvents

`func (o *AgentsSessionDetail) GetRecentEvents() []AgentsEventView`

GetRecentEvents returns the RecentEvents field if non-nil, zero value otherwise.

### GetRecentEventsOk

`func (o *AgentsSessionDetail) GetRecentEventsOk() (*[]AgentsEventView, bool)`

GetRecentEventsOk returns a tuple with the RecentEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecentEvents

`func (o *AgentsSessionDetail) SetRecentEvents(v []AgentsEventView)`

SetRecentEvents sets RecentEvents field to given value.

### HasRecentEvents

`func (o *AgentsSessionDetail) HasRecentEvents() bool`

HasRecentEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


