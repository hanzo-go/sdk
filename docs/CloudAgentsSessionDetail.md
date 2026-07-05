# CloudAgentsSessionDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Session id (sess_...). | [optional] 
**Agent** | Pointer to **string** | Agent name/type label (need not be a cloud Agent row). | [optional] 
**Actor** | Pointer to **string** | The principal that started it. | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**ParentSessionId** | Pointer to **string** | Empty for a root (the outer agent). | [optional] 
**RootSessionId** | Pointer to **string** | The tree key; equals id for a root. | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**TaskWorkflowId** | Pointer to **string** | The hanzoai/tasks workflow that executes this session, when task-backed. | [optional] 
**TaskRunId** | Pointer to **string** |  | [optional] 
**Events** | Pointer to **int32** |  | [optional] 
**Children** | Pointer to **int32** | Direct-child (subagent) fan-out count. | [optional] 
**StartedAt** | Pointer to **time.Time** |  | [optional] 
**EndedAt** | Pointer to **time.Time** | Set once a terminal status (done/error) is reached. | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**ChildSessions** | Pointer to [**[]CloudAgentsSession**](CloudAgentsSession.md) |  | [optional] 
**RecentEvents** | Pointer to [**[]CloudAgentsEvent**](CloudAgentsEvent.md) |  | [optional] 

## Methods

### NewCloudAgentsSessionDetail

`func NewCloudAgentsSessionDetail() *CloudAgentsSessionDetail`

NewCloudAgentsSessionDetail instantiates a new CloudAgentsSessionDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudAgentsSessionDetailWithDefaults

`func NewCloudAgentsSessionDetailWithDefaults() *CloudAgentsSessionDetail`

NewCloudAgentsSessionDetailWithDefaults instantiates a new CloudAgentsSessionDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CloudAgentsSessionDetail) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudAgentsSessionDetail) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudAgentsSessionDetail) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudAgentsSessionDetail) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAgent

`func (o *CloudAgentsSessionDetail) GetAgent() string`

GetAgent returns the Agent field if non-nil, zero value otherwise.

### GetAgentOk

`func (o *CloudAgentsSessionDetail) GetAgentOk() (*string, bool)`

GetAgentOk returns a tuple with the Agent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgent

`func (o *CloudAgentsSessionDetail) SetAgent(v string)`

SetAgent sets Agent field to given value.

### HasAgent

`func (o *CloudAgentsSessionDetail) HasAgent() bool`

HasAgent returns a boolean if a field has been set.

### GetActor

`func (o *CloudAgentsSessionDetail) GetActor() string`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *CloudAgentsSessionDetail) GetActorOk() (*string, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *CloudAgentsSessionDetail) SetActor(v string)`

SetActor sets Actor field to given value.

### HasActor

`func (o *CloudAgentsSessionDetail) HasActor() bool`

HasActor returns a boolean if a field has been set.

### GetStatus

`func (o *CloudAgentsSessionDetail) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudAgentsSessionDetail) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudAgentsSessionDetail) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudAgentsSessionDetail) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetParentSessionId

`func (o *CloudAgentsSessionDetail) GetParentSessionId() string`

GetParentSessionId returns the ParentSessionId field if non-nil, zero value otherwise.

### GetParentSessionIdOk

`func (o *CloudAgentsSessionDetail) GetParentSessionIdOk() (*string, bool)`

GetParentSessionIdOk returns a tuple with the ParentSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentSessionId

`func (o *CloudAgentsSessionDetail) SetParentSessionId(v string)`

SetParentSessionId sets ParentSessionId field to given value.

### HasParentSessionId

`func (o *CloudAgentsSessionDetail) HasParentSessionId() bool`

HasParentSessionId returns a boolean if a field has been set.

### GetRootSessionId

`func (o *CloudAgentsSessionDetail) GetRootSessionId() string`

GetRootSessionId returns the RootSessionId field if non-nil, zero value otherwise.

### GetRootSessionIdOk

`func (o *CloudAgentsSessionDetail) GetRootSessionIdOk() (*string, bool)`

GetRootSessionIdOk returns a tuple with the RootSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootSessionId

`func (o *CloudAgentsSessionDetail) SetRootSessionId(v string)`

SetRootSessionId sets RootSessionId field to given value.

### HasRootSessionId

`func (o *CloudAgentsSessionDetail) HasRootSessionId() bool`

HasRootSessionId returns a boolean if a field has been set.

### GetTitle

`func (o *CloudAgentsSessionDetail) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CloudAgentsSessionDetail) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CloudAgentsSessionDetail) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CloudAgentsSessionDetail) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetTaskWorkflowId

`func (o *CloudAgentsSessionDetail) GetTaskWorkflowId() string`

GetTaskWorkflowId returns the TaskWorkflowId field if non-nil, zero value otherwise.

### GetTaskWorkflowIdOk

`func (o *CloudAgentsSessionDetail) GetTaskWorkflowIdOk() (*string, bool)`

GetTaskWorkflowIdOk returns a tuple with the TaskWorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskWorkflowId

`func (o *CloudAgentsSessionDetail) SetTaskWorkflowId(v string)`

SetTaskWorkflowId sets TaskWorkflowId field to given value.

### HasTaskWorkflowId

`func (o *CloudAgentsSessionDetail) HasTaskWorkflowId() bool`

HasTaskWorkflowId returns a boolean if a field has been set.

### GetTaskRunId

`func (o *CloudAgentsSessionDetail) GetTaskRunId() string`

GetTaskRunId returns the TaskRunId field if non-nil, zero value otherwise.

### GetTaskRunIdOk

`func (o *CloudAgentsSessionDetail) GetTaskRunIdOk() (*string, bool)`

GetTaskRunIdOk returns a tuple with the TaskRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskRunId

`func (o *CloudAgentsSessionDetail) SetTaskRunId(v string)`

SetTaskRunId sets TaskRunId field to given value.

### HasTaskRunId

`func (o *CloudAgentsSessionDetail) HasTaskRunId() bool`

HasTaskRunId returns a boolean if a field has been set.

### GetEvents

`func (o *CloudAgentsSessionDetail) GetEvents() int32`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *CloudAgentsSessionDetail) GetEventsOk() (*int32, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *CloudAgentsSessionDetail) SetEvents(v int32)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *CloudAgentsSessionDetail) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetChildren

`func (o *CloudAgentsSessionDetail) GetChildren() int32`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *CloudAgentsSessionDetail) GetChildrenOk() (*int32, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *CloudAgentsSessionDetail) SetChildren(v int32)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *CloudAgentsSessionDetail) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### GetStartedAt

`func (o *CloudAgentsSessionDetail) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *CloudAgentsSessionDetail) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *CloudAgentsSessionDetail) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *CloudAgentsSessionDetail) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetEndedAt

`func (o *CloudAgentsSessionDetail) GetEndedAt() time.Time`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *CloudAgentsSessionDetail) GetEndedAtOk() (*time.Time, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *CloudAgentsSessionDetail) SetEndedAt(v time.Time)`

SetEndedAt sets EndedAt field to given value.

### HasEndedAt

`func (o *CloudAgentsSessionDetail) HasEndedAt() bool`

HasEndedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CloudAgentsSessionDetail) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CloudAgentsSessionDetail) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CloudAgentsSessionDetail) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CloudAgentsSessionDetail) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CloudAgentsSessionDetail) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CloudAgentsSessionDetail) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CloudAgentsSessionDetail) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CloudAgentsSessionDetail) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetChildSessions

`func (o *CloudAgentsSessionDetail) GetChildSessions() []CloudAgentsSession`

GetChildSessions returns the ChildSessions field if non-nil, zero value otherwise.

### GetChildSessionsOk

`func (o *CloudAgentsSessionDetail) GetChildSessionsOk() (*[]CloudAgentsSession, bool)`

GetChildSessionsOk returns a tuple with the ChildSessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildSessions

`func (o *CloudAgentsSessionDetail) SetChildSessions(v []CloudAgentsSession)`

SetChildSessions sets ChildSessions field to given value.

### HasChildSessions

`func (o *CloudAgentsSessionDetail) HasChildSessions() bool`

HasChildSessions returns a boolean if a field has been set.

### GetRecentEvents

`func (o *CloudAgentsSessionDetail) GetRecentEvents() []CloudAgentsEvent`

GetRecentEvents returns the RecentEvents field if non-nil, zero value otherwise.

### GetRecentEventsOk

`func (o *CloudAgentsSessionDetail) GetRecentEventsOk() (*[]CloudAgentsEvent, bool)`

GetRecentEventsOk returns a tuple with the RecentEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecentEvents

`func (o *CloudAgentsSessionDetail) SetRecentEvents(v []CloudAgentsEvent)`

SetRecentEvents sets RecentEvents field to given value.

### HasRecentEvents

`func (o *CloudAgentsSessionDetail) HasRecentEvents() bool`

HasRecentEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


