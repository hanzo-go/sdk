# CloudAgentsSession

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

## Methods

### NewCloudAgentsSession

`func NewCloudAgentsSession() *CloudAgentsSession`

NewCloudAgentsSession instantiates a new CloudAgentsSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudAgentsSessionWithDefaults

`func NewCloudAgentsSessionWithDefaults() *CloudAgentsSession`

NewCloudAgentsSessionWithDefaults instantiates a new CloudAgentsSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CloudAgentsSession) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudAgentsSession) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudAgentsSession) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudAgentsSession) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAgent

`func (o *CloudAgentsSession) GetAgent() string`

GetAgent returns the Agent field if non-nil, zero value otherwise.

### GetAgentOk

`func (o *CloudAgentsSession) GetAgentOk() (*string, bool)`

GetAgentOk returns a tuple with the Agent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgent

`func (o *CloudAgentsSession) SetAgent(v string)`

SetAgent sets Agent field to given value.

### HasAgent

`func (o *CloudAgentsSession) HasAgent() bool`

HasAgent returns a boolean if a field has been set.

### GetActor

`func (o *CloudAgentsSession) GetActor() string`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *CloudAgentsSession) GetActorOk() (*string, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *CloudAgentsSession) SetActor(v string)`

SetActor sets Actor field to given value.

### HasActor

`func (o *CloudAgentsSession) HasActor() bool`

HasActor returns a boolean if a field has been set.

### GetStatus

`func (o *CloudAgentsSession) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudAgentsSession) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudAgentsSession) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudAgentsSession) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetParentSessionId

`func (o *CloudAgentsSession) GetParentSessionId() string`

GetParentSessionId returns the ParentSessionId field if non-nil, zero value otherwise.

### GetParentSessionIdOk

`func (o *CloudAgentsSession) GetParentSessionIdOk() (*string, bool)`

GetParentSessionIdOk returns a tuple with the ParentSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentSessionId

`func (o *CloudAgentsSession) SetParentSessionId(v string)`

SetParentSessionId sets ParentSessionId field to given value.

### HasParentSessionId

`func (o *CloudAgentsSession) HasParentSessionId() bool`

HasParentSessionId returns a boolean if a field has been set.

### GetRootSessionId

`func (o *CloudAgentsSession) GetRootSessionId() string`

GetRootSessionId returns the RootSessionId field if non-nil, zero value otherwise.

### GetRootSessionIdOk

`func (o *CloudAgentsSession) GetRootSessionIdOk() (*string, bool)`

GetRootSessionIdOk returns a tuple with the RootSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootSessionId

`func (o *CloudAgentsSession) SetRootSessionId(v string)`

SetRootSessionId sets RootSessionId field to given value.

### HasRootSessionId

`func (o *CloudAgentsSession) HasRootSessionId() bool`

HasRootSessionId returns a boolean if a field has been set.

### GetTitle

`func (o *CloudAgentsSession) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CloudAgentsSession) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CloudAgentsSession) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CloudAgentsSession) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetTaskWorkflowId

`func (o *CloudAgentsSession) GetTaskWorkflowId() string`

GetTaskWorkflowId returns the TaskWorkflowId field if non-nil, zero value otherwise.

### GetTaskWorkflowIdOk

`func (o *CloudAgentsSession) GetTaskWorkflowIdOk() (*string, bool)`

GetTaskWorkflowIdOk returns a tuple with the TaskWorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskWorkflowId

`func (o *CloudAgentsSession) SetTaskWorkflowId(v string)`

SetTaskWorkflowId sets TaskWorkflowId field to given value.

### HasTaskWorkflowId

`func (o *CloudAgentsSession) HasTaskWorkflowId() bool`

HasTaskWorkflowId returns a boolean if a field has been set.

### GetTaskRunId

`func (o *CloudAgentsSession) GetTaskRunId() string`

GetTaskRunId returns the TaskRunId field if non-nil, zero value otherwise.

### GetTaskRunIdOk

`func (o *CloudAgentsSession) GetTaskRunIdOk() (*string, bool)`

GetTaskRunIdOk returns a tuple with the TaskRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskRunId

`func (o *CloudAgentsSession) SetTaskRunId(v string)`

SetTaskRunId sets TaskRunId field to given value.

### HasTaskRunId

`func (o *CloudAgentsSession) HasTaskRunId() bool`

HasTaskRunId returns a boolean if a field has been set.

### GetEvents

`func (o *CloudAgentsSession) GetEvents() int32`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *CloudAgentsSession) GetEventsOk() (*int32, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *CloudAgentsSession) SetEvents(v int32)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *CloudAgentsSession) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetChildren

`func (o *CloudAgentsSession) GetChildren() int32`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *CloudAgentsSession) GetChildrenOk() (*int32, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *CloudAgentsSession) SetChildren(v int32)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *CloudAgentsSession) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### GetStartedAt

`func (o *CloudAgentsSession) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *CloudAgentsSession) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *CloudAgentsSession) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *CloudAgentsSession) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetEndedAt

`func (o *CloudAgentsSession) GetEndedAt() time.Time`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *CloudAgentsSession) GetEndedAtOk() (*time.Time, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *CloudAgentsSession) SetEndedAt(v time.Time)`

SetEndedAt sets EndedAt field to given value.

### HasEndedAt

`func (o *CloudAgentsSession) HasEndedAt() bool`

HasEndedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CloudAgentsSession) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CloudAgentsSession) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CloudAgentsSession) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CloudAgentsSession) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CloudAgentsSession) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CloudAgentsSession) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CloudAgentsSession) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CloudAgentsSession) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


