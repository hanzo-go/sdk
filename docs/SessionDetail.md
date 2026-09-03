# SessionDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** |  | [optional] 
**Actor** | Pointer to **string** |  | [optional] 
**Agent** | Pointer to **string** |  | [optional] 
**ChildSessions** | Pointer to [**[]SessionView**](SessionView.md) | Children is the session&#39;s DIRECT children, one level down, each with its own counts. The promoted &#x60;children&#x60; integer beside it is how many there are; this is who they are. For the whole subtree, read the tree. | [optional] 
**Children** | Pointer to **int64** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**Cwd** | Pointer to **string** |  | [optional] 
**EndedAt** | Pointer to **string** |  | [optional] 
**Events** | Pointer to **int64** |  | [optional] 
**Host** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**LastEvent** | Pointer to [**LastEventView**](LastEventView.md) |  | [optional] 
**Org** | Pointer to **string** |  | [optional] 
**ParentSessionId** | Pointer to **string** |  | [optional] 
**Progress** | Pointer to [**SessionProgress**](SessionProgress.md) |  | [optional] 
**Project** | Pointer to **string** |  | [optional] 
**Provider** | Pointer to **string** |  | [optional] 
**Published** | Pointer to **bool** |  | [optional] 
**RecentEvents** | Pointer to [**[]EventView**](EventView.md) | RecentEvents is the 50 most recent turns, OLDEST of those first — a transcript to read down, not a feed. The promoted &#x60;events&#x60; integer says how many the log holds in total; page the rest from a seq. | [optional] 
**Repo** | Pointer to **string** |  | [optional] 
**Room** | Pointer to **string** |  | [optional] 
**RootSessionId** | Pointer to **string** |  | [optional] 
**StartedAt** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Target** | Pointer to **string** |  | [optional] 
**TaskRunId** | Pointer to **string** |  | [optional] 
**TaskWorkflowId** | Pointer to **string** |  | [optional] 
**Terminal** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewSessionDetail

`func NewSessionDetail() *SessionDetail`

NewSessionDetail instantiates a new SessionDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionDetailWithDefaults

`func NewSessionDetailWithDefaults() *SessionDetail`

NewSessionDetailWithDefaults instantiates a new SessionDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *SessionDetail) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *SessionDetail) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *SessionDetail) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *SessionDetail) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetActor

`func (o *SessionDetail) GetActor() string`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *SessionDetail) GetActorOk() (*string, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *SessionDetail) SetActor(v string)`

SetActor sets Actor field to given value.

### HasActor

`func (o *SessionDetail) HasActor() bool`

HasActor returns a boolean if a field has been set.

### GetAgent

`func (o *SessionDetail) GetAgent() string`

GetAgent returns the Agent field if non-nil, zero value otherwise.

### GetAgentOk

`func (o *SessionDetail) GetAgentOk() (*string, bool)`

GetAgentOk returns a tuple with the Agent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgent

`func (o *SessionDetail) SetAgent(v string)`

SetAgent sets Agent field to given value.

### HasAgent

`func (o *SessionDetail) HasAgent() bool`

HasAgent returns a boolean if a field has been set.

### GetChildSessions

`func (o *SessionDetail) GetChildSessions() []SessionView`

GetChildSessions returns the ChildSessions field if non-nil, zero value otherwise.

### GetChildSessionsOk

`func (o *SessionDetail) GetChildSessionsOk() (*[]SessionView, bool)`

GetChildSessionsOk returns a tuple with the ChildSessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildSessions

`func (o *SessionDetail) SetChildSessions(v []SessionView)`

SetChildSessions sets ChildSessions field to given value.

### HasChildSessions

`func (o *SessionDetail) HasChildSessions() bool`

HasChildSessions returns a boolean if a field has been set.

### GetChildren

`func (o *SessionDetail) GetChildren() int64`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *SessionDetail) GetChildrenOk() (*int64, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *SessionDetail) SetChildren(v int64)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *SessionDetail) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SessionDetail) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SessionDetail) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SessionDetail) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SessionDetail) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCwd

`func (o *SessionDetail) GetCwd() string`

GetCwd returns the Cwd field if non-nil, zero value otherwise.

### GetCwdOk

`func (o *SessionDetail) GetCwdOk() (*string, bool)`

GetCwdOk returns a tuple with the Cwd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCwd

`func (o *SessionDetail) SetCwd(v string)`

SetCwd sets Cwd field to given value.

### HasCwd

`func (o *SessionDetail) HasCwd() bool`

HasCwd returns a boolean if a field has been set.

### GetEndedAt

`func (o *SessionDetail) GetEndedAt() string`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *SessionDetail) GetEndedAtOk() (*string, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *SessionDetail) SetEndedAt(v string)`

SetEndedAt sets EndedAt field to given value.

### HasEndedAt

`func (o *SessionDetail) HasEndedAt() bool`

HasEndedAt returns a boolean if a field has been set.

### GetEvents

`func (o *SessionDetail) GetEvents() int64`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *SessionDetail) GetEventsOk() (*int64, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *SessionDetail) SetEvents(v int64)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *SessionDetail) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetHost

`func (o *SessionDetail) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *SessionDetail) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *SessionDetail) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *SessionDetail) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetId

`func (o *SessionDetail) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SessionDetail) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SessionDetail) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SessionDetail) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastEvent

`func (o *SessionDetail) GetLastEvent() LastEventView`

GetLastEvent returns the LastEvent field if non-nil, zero value otherwise.

### GetLastEventOk

`func (o *SessionDetail) GetLastEventOk() (*LastEventView, bool)`

GetLastEventOk returns a tuple with the LastEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEvent

`func (o *SessionDetail) SetLastEvent(v LastEventView)`

SetLastEvent sets LastEvent field to given value.

### HasLastEvent

`func (o *SessionDetail) HasLastEvent() bool`

HasLastEvent returns a boolean if a field has been set.

### GetOrg

`func (o *SessionDetail) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *SessionDetail) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *SessionDetail) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *SessionDetail) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetParentSessionId

`func (o *SessionDetail) GetParentSessionId() string`

GetParentSessionId returns the ParentSessionId field if non-nil, zero value otherwise.

### GetParentSessionIdOk

`func (o *SessionDetail) GetParentSessionIdOk() (*string, bool)`

GetParentSessionIdOk returns a tuple with the ParentSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentSessionId

`func (o *SessionDetail) SetParentSessionId(v string)`

SetParentSessionId sets ParentSessionId field to given value.

### HasParentSessionId

`func (o *SessionDetail) HasParentSessionId() bool`

HasParentSessionId returns a boolean if a field has been set.

### GetProgress

`func (o *SessionDetail) GetProgress() SessionProgress`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *SessionDetail) GetProgressOk() (*SessionProgress, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *SessionDetail) SetProgress(v SessionProgress)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *SessionDetail) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetProject

`func (o *SessionDetail) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *SessionDetail) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *SessionDetail) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *SessionDetail) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetProvider

`func (o *SessionDetail) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *SessionDetail) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *SessionDetail) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *SessionDetail) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetPublished

`func (o *SessionDetail) GetPublished() bool`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *SessionDetail) GetPublishedOk() (*bool, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *SessionDetail) SetPublished(v bool)`

SetPublished sets Published field to given value.

### HasPublished

`func (o *SessionDetail) HasPublished() bool`

HasPublished returns a boolean if a field has been set.

### GetRecentEvents

`func (o *SessionDetail) GetRecentEvents() []EventView`

GetRecentEvents returns the RecentEvents field if non-nil, zero value otherwise.

### GetRecentEventsOk

`func (o *SessionDetail) GetRecentEventsOk() (*[]EventView, bool)`

GetRecentEventsOk returns a tuple with the RecentEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecentEvents

`func (o *SessionDetail) SetRecentEvents(v []EventView)`

SetRecentEvents sets RecentEvents field to given value.

### HasRecentEvents

`func (o *SessionDetail) HasRecentEvents() bool`

HasRecentEvents returns a boolean if a field has been set.

### GetRepo

`func (o *SessionDetail) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *SessionDetail) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *SessionDetail) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *SessionDetail) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### GetRoom

`func (o *SessionDetail) GetRoom() string`

GetRoom returns the Room field if non-nil, zero value otherwise.

### GetRoomOk

`func (o *SessionDetail) GetRoomOk() (*string, bool)`

GetRoomOk returns a tuple with the Room field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoom

`func (o *SessionDetail) SetRoom(v string)`

SetRoom sets Room field to given value.

### HasRoom

`func (o *SessionDetail) HasRoom() bool`

HasRoom returns a boolean if a field has been set.

### GetRootSessionId

`func (o *SessionDetail) GetRootSessionId() string`

GetRootSessionId returns the RootSessionId field if non-nil, zero value otherwise.

### GetRootSessionIdOk

`func (o *SessionDetail) GetRootSessionIdOk() (*string, bool)`

GetRootSessionIdOk returns a tuple with the RootSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootSessionId

`func (o *SessionDetail) SetRootSessionId(v string)`

SetRootSessionId sets RootSessionId field to given value.

### HasRootSessionId

`func (o *SessionDetail) HasRootSessionId() bool`

HasRootSessionId returns a boolean if a field has been set.

### GetStartedAt

`func (o *SessionDetail) GetStartedAt() string`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *SessionDetail) GetStartedAtOk() (*string, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *SessionDetail) SetStartedAt(v string)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *SessionDetail) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetStatus

`func (o *SessionDetail) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SessionDetail) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SessionDetail) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SessionDetail) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTarget

`func (o *SessionDetail) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *SessionDetail) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *SessionDetail) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *SessionDetail) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetTaskRunId

`func (o *SessionDetail) GetTaskRunId() string`

GetTaskRunId returns the TaskRunId field if non-nil, zero value otherwise.

### GetTaskRunIdOk

`func (o *SessionDetail) GetTaskRunIdOk() (*string, bool)`

GetTaskRunIdOk returns a tuple with the TaskRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskRunId

`func (o *SessionDetail) SetTaskRunId(v string)`

SetTaskRunId sets TaskRunId field to given value.

### HasTaskRunId

`func (o *SessionDetail) HasTaskRunId() bool`

HasTaskRunId returns a boolean if a field has been set.

### GetTaskWorkflowId

`func (o *SessionDetail) GetTaskWorkflowId() string`

GetTaskWorkflowId returns the TaskWorkflowId field if non-nil, zero value otherwise.

### GetTaskWorkflowIdOk

`func (o *SessionDetail) GetTaskWorkflowIdOk() (*string, bool)`

GetTaskWorkflowIdOk returns a tuple with the TaskWorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskWorkflowId

`func (o *SessionDetail) SetTaskWorkflowId(v string)`

SetTaskWorkflowId sets TaskWorkflowId field to given value.

### HasTaskWorkflowId

`func (o *SessionDetail) HasTaskWorkflowId() bool`

HasTaskWorkflowId returns a boolean if a field has been set.

### GetTerminal

`func (o *SessionDetail) GetTerminal() string`

GetTerminal returns the Terminal field if non-nil, zero value otherwise.

### GetTerminalOk

`func (o *SessionDetail) GetTerminalOk() (*string, bool)`

GetTerminalOk returns a tuple with the Terminal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminal

`func (o *SessionDetail) SetTerminal(v string)`

SetTerminal sets Terminal field to given value.

### HasTerminal

`func (o *SessionDetail) HasTerminal() bool`

HasTerminal returns a boolean if a field has been set.

### GetTitle

`func (o *SessionDetail) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SessionDetail) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SessionDetail) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *SessionDetail) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *SessionDetail) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SessionDetail) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SessionDetail) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *SessionDetail) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


