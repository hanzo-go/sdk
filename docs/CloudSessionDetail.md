# CloudSessionDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** |  | [optional] 
**Actor** | Pointer to **string** |  | [optional] 
**Agent** | Pointer to **string** |  | [optional] 
**ChildSessions** | Pointer to [**[]CloudSessionView**](CloudSessionView.md) |  | [optional] 
**Children** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**Cwd** | Pointer to **string** |  | [optional] 
**EndedAt** | Pointer to **string** |  | [optional] 
**Events** | Pointer to **int32** |  | [optional] 
**Host** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**LastEvent** | Pointer to [**CloudLastEventView**](CloudLastEventView.md) |  | [optional] 
**Org** | Pointer to **string** |  | [optional] 
**ParentSessionId** | Pointer to **string** |  | [optional] 
**Project** | Pointer to **string** |  | [optional] 
**Provider** | Pointer to **string** |  | [optional] 
**Published** | Pointer to **bool** |  | [optional] 
**RecentEvents** | Pointer to [**[]CloudEventView**](CloudEventView.md) |  | [optional] 
**Repo** | Pointer to **string** |  | [optional] 
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

### NewCloudSessionDetail

`func NewCloudSessionDetail() *CloudSessionDetail`

NewCloudSessionDetail instantiates a new CloudSessionDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudSessionDetailWithDefaults

`func NewCloudSessionDetailWithDefaults() *CloudSessionDetail`

NewCloudSessionDetailWithDefaults instantiates a new CloudSessionDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *CloudSessionDetail) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *CloudSessionDetail) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *CloudSessionDetail) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *CloudSessionDetail) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetActor

`func (o *CloudSessionDetail) GetActor() string`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *CloudSessionDetail) GetActorOk() (*string, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *CloudSessionDetail) SetActor(v string)`

SetActor sets Actor field to given value.

### HasActor

`func (o *CloudSessionDetail) HasActor() bool`

HasActor returns a boolean if a field has been set.

### GetAgent

`func (o *CloudSessionDetail) GetAgent() string`

GetAgent returns the Agent field if non-nil, zero value otherwise.

### GetAgentOk

`func (o *CloudSessionDetail) GetAgentOk() (*string, bool)`

GetAgentOk returns a tuple with the Agent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgent

`func (o *CloudSessionDetail) SetAgent(v string)`

SetAgent sets Agent field to given value.

### HasAgent

`func (o *CloudSessionDetail) HasAgent() bool`

HasAgent returns a boolean if a field has been set.

### GetChildSessions

`func (o *CloudSessionDetail) GetChildSessions() []CloudSessionView`

GetChildSessions returns the ChildSessions field if non-nil, zero value otherwise.

### GetChildSessionsOk

`func (o *CloudSessionDetail) GetChildSessionsOk() (*[]CloudSessionView, bool)`

GetChildSessionsOk returns a tuple with the ChildSessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildSessions

`func (o *CloudSessionDetail) SetChildSessions(v []CloudSessionView)`

SetChildSessions sets ChildSessions field to given value.

### HasChildSessions

`func (o *CloudSessionDetail) HasChildSessions() bool`

HasChildSessions returns a boolean if a field has been set.

### GetChildren

`func (o *CloudSessionDetail) GetChildren() int32`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *CloudSessionDetail) GetChildrenOk() (*int32, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *CloudSessionDetail) SetChildren(v int32)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *CloudSessionDetail) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CloudSessionDetail) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CloudSessionDetail) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CloudSessionDetail) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CloudSessionDetail) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCwd

`func (o *CloudSessionDetail) GetCwd() string`

GetCwd returns the Cwd field if non-nil, zero value otherwise.

### GetCwdOk

`func (o *CloudSessionDetail) GetCwdOk() (*string, bool)`

GetCwdOk returns a tuple with the Cwd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCwd

`func (o *CloudSessionDetail) SetCwd(v string)`

SetCwd sets Cwd field to given value.

### HasCwd

`func (o *CloudSessionDetail) HasCwd() bool`

HasCwd returns a boolean if a field has been set.

### GetEndedAt

`func (o *CloudSessionDetail) GetEndedAt() string`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *CloudSessionDetail) GetEndedAtOk() (*string, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *CloudSessionDetail) SetEndedAt(v string)`

SetEndedAt sets EndedAt field to given value.

### HasEndedAt

`func (o *CloudSessionDetail) HasEndedAt() bool`

HasEndedAt returns a boolean if a field has been set.

### GetEvents

`func (o *CloudSessionDetail) GetEvents() int32`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *CloudSessionDetail) GetEventsOk() (*int32, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *CloudSessionDetail) SetEvents(v int32)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *CloudSessionDetail) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetHost

`func (o *CloudSessionDetail) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *CloudSessionDetail) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *CloudSessionDetail) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *CloudSessionDetail) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetId

`func (o *CloudSessionDetail) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudSessionDetail) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudSessionDetail) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudSessionDetail) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastEvent

`func (o *CloudSessionDetail) GetLastEvent() CloudLastEventView`

GetLastEvent returns the LastEvent field if non-nil, zero value otherwise.

### GetLastEventOk

`func (o *CloudSessionDetail) GetLastEventOk() (*CloudLastEventView, bool)`

GetLastEventOk returns a tuple with the LastEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEvent

`func (o *CloudSessionDetail) SetLastEvent(v CloudLastEventView)`

SetLastEvent sets LastEvent field to given value.

### HasLastEvent

`func (o *CloudSessionDetail) HasLastEvent() bool`

HasLastEvent returns a boolean if a field has been set.

### GetOrg

`func (o *CloudSessionDetail) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *CloudSessionDetail) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *CloudSessionDetail) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *CloudSessionDetail) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetParentSessionId

`func (o *CloudSessionDetail) GetParentSessionId() string`

GetParentSessionId returns the ParentSessionId field if non-nil, zero value otherwise.

### GetParentSessionIdOk

`func (o *CloudSessionDetail) GetParentSessionIdOk() (*string, bool)`

GetParentSessionIdOk returns a tuple with the ParentSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentSessionId

`func (o *CloudSessionDetail) SetParentSessionId(v string)`

SetParentSessionId sets ParentSessionId field to given value.

### HasParentSessionId

`func (o *CloudSessionDetail) HasParentSessionId() bool`

HasParentSessionId returns a boolean if a field has been set.

### GetProject

`func (o *CloudSessionDetail) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *CloudSessionDetail) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *CloudSessionDetail) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *CloudSessionDetail) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetProvider

`func (o *CloudSessionDetail) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *CloudSessionDetail) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *CloudSessionDetail) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *CloudSessionDetail) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetPublished

`func (o *CloudSessionDetail) GetPublished() bool`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *CloudSessionDetail) GetPublishedOk() (*bool, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *CloudSessionDetail) SetPublished(v bool)`

SetPublished sets Published field to given value.

### HasPublished

`func (o *CloudSessionDetail) HasPublished() bool`

HasPublished returns a boolean if a field has been set.

### GetRecentEvents

`func (o *CloudSessionDetail) GetRecentEvents() []CloudEventView`

GetRecentEvents returns the RecentEvents field if non-nil, zero value otherwise.

### GetRecentEventsOk

`func (o *CloudSessionDetail) GetRecentEventsOk() (*[]CloudEventView, bool)`

GetRecentEventsOk returns a tuple with the RecentEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecentEvents

`func (o *CloudSessionDetail) SetRecentEvents(v []CloudEventView)`

SetRecentEvents sets RecentEvents field to given value.

### HasRecentEvents

`func (o *CloudSessionDetail) HasRecentEvents() bool`

HasRecentEvents returns a boolean if a field has been set.

### GetRepo

`func (o *CloudSessionDetail) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *CloudSessionDetail) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *CloudSessionDetail) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *CloudSessionDetail) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### GetRootSessionId

`func (o *CloudSessionDetail) GetRootSessionId() string`

GetRootSessionId returns the RootSessionId field if non-nil, zero value otherwise.

### GetRootSessionIdOk

`func (o *CloudSessionDetail) GetRootSessionIdOk() (*string, bool)`

GetRootSessionIdOk returns a tuple with the RootSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootSessionId

`func (o *CloudSessionDetail) SetRootSessionId(v string)`

SetRootSessionId sets RootSessionId field to given value.

### HasRootSessionId

`func (o *CloudSessionDetail) HasRootSessionId() bool`

HasRootSessionId returns a boolean if a field has been set.

### GetStartedAt

`func (o *CloudSessionDetail) GetStartedAt() string`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *CloudSessionDetail) GetStartedAtOk() (*string, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *CloudSessionDetail) SetStartedAt(v string)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *CloudSessionDetail) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetStatus

`func (o *CloudSessionDetail) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudSessionDetail) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudSessionDetail) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudSessionDetail) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTarget

`func (o *CloudSessionDetail) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *CloudSessionDetail) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *CloudSessionDetail) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *CloudSessionDetail) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetTaskRunId

`func (o *CloudSessionDetail) GetTaskRunId() string`

GetTaskRunId returns the TaskRunId field if non-nil, zero value otherwise.

### GetTaskRunIdOk

`func (o *CloudSessionDetail) GetTaskRunIdOk() (*string, bool)`

GetTaskRunIdOk returns a tuple with the TaskRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskRunId

`func (o *CloudSessionDetail) SetTaskRunId(v string)`

SetTaskRunId sets TaskRunId field to given value.

### HasTaskRunId

`func (o *CloudSessionDetail) HasTaskRunId() bool`

HasTaskRunId returns a boolean if a field has been set.

### GetTaskWorkflowId

`func (o *CloudSessionDetail) GetTaskWorkflowId() string`

GetTaskWorkflowId returns the TaskWorkflowId field if non-nil, zero value otherwise.

### GetTaskWorkflowIdOk

`func (o *CloudSessionDetail) GetTaskWorkflowIdOk() (*string, bool)`

GetTaskWorkflowIdOk returns a tuple with the TaskWorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskWorkflowId

`func (o *CloudSessionDetail) SetTaskWorkflowId(v string)`

SetTaskWorkflowId sets TaskWorkflowId field to given value.

### HasTaskWorkflowId

`func (o *CloudSessionDetail) HasTaskWorkflowId() bool`

HasTaskWorkflowId returns a boolean if a field has been set.

### GetTerminal

`func (o *CloudSessionDetail) GetTerminal() string`

GetTerminal returns the Terminal field if non-nil, zero value otherwise.

### GetTerminalOk

`func (o *CloudSessionDetail) GetTerminalOk() (*string, bool)`

GetTerminalOk returns a tuple with the Terminal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminal

`func (o *CloudSessionDetail) SetTerminal(v string)`

SetTerminal sets Terminal field to given value.

### HasTerminal

`func (o *CloudSessionDetail) HasTerminal() bool`

HasTerminal returns a boolean if a field has been set.

### GetTitle

`func (o *CloudSessionDetail) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CloudSessionDetail) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CloudSessionDetail) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CloudSessionDetail) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CloudSessionDetail) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CloudSessionDetail) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CloudSessionDetail) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CloudSessionDetail) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


