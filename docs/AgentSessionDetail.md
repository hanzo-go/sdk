# AgentSessionDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** |  | [optional] 
**Actor** | Pointer to **string** |  | [optional] 
**Agent** | Pointer to **string** |  | [optional] 
**Base** | Pointer to **string** |  | [optional] 
**Branch** | Pointer to **string** |  | [optional] 
**ChildSessions** | Pointer to [**[]AgentSessionView**](AgentSessionView.md) | Children is the session&#39;s DIRECT children, one level down, each with its own counts. The promoted &#x60;children&#x60; integer beside it is how many there are; this is who they are. For the whole subtree, read the tree. | [optional] 
**Children** | Pointer to **int64** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**Cwd** | Pointer to **string** |  | [optional] 
**EndedAt** | Pointer to **string** |  | [optional] 
**Environment** | Pointer to **string** |  | [optional] 
**Events** | Pointer to **int64** |  | [optional] 
**Host** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**LastEvent** | Pointer to [**AgentLastEventView**](AgentLastEventView.md) |  | [optional] 
**Mode** | Pointer to **string** |  | [optional] 
**Org** | Pointer to **string** |  | [optional] 
**ParentSessionId** | Pointer to **string** |  | [optional] 
**Pr** | Pointer to **string** |  | [optional] 
**Progress** | Pointer to [**AgentSessionProgress**](AgentSessionProgress.md) |  | [optional] 
**Project** | Pointer to **string** |  | [optional] 
**Provider** | Pointer to **string** |  | [optional] 
**Published** | Pointer to **bool** |  | [optional] 
**RecentEvents** | Pointer to [**[]AgentEventView**](AgentEventView.md) | RecentEvents is the 50 most recent turns, OLDEST of those first — a transcript to read down, not a feed. The promoted &#x60;events&#x60; integer says how many the log holds in total; page the rest from a seq. | [optional] 
**Repo** | Pointer to **string** |  | [optional] 
**Room** | Pointer to **string** |  | [optional] 
**RootSessionId** | Pointer to **string** |  | [optional] 
**Sandbox** | Pointer to **string** |  | [optional] 
**StartedAt** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Target** | Pointer to **string** |  | [optional] 
**TaskRunId** | Pointer to **string** |  | [optional] 
**TaskWorkflowId** | Pointer to **string** |  | [optional] 
**Terminal** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewAgentSessionDetail

`func NewAgentSessionDetail() *AgentSessionDetail`

NewAgentSessionDetail instantiates a new AgentSessionDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentSessionDetailWithDefaults

`func NewAgentSessionDetailWithDefaults() *AgentSessionDetail`

NewAgentSessionDetailWithDefaults instantiates a new AgentSessionDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *AgentSessionDetail) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *AgentSessionDetail) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *AgentSessionDetail) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *AgentSessionDetail) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetActor

`func (o *AgentSessionDetail) GetActor() string`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *AgentSessionDetail) GetActorOk() (*string, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *AgentSessionDetail) SetActor(v string)`

SetActor sets Actor field to given value.

### HasActor

`func (o *AgentSessionDetail) HasActor() bool`

HasActor returns a boolean if a field has been set.

### GetAgent

`func (o *AgentSessionDetail) GetAgent() string`

GetAgent returns the Agent field if non-nil, zero value otherwise.

### GetAgentOk

`func (o *AgentSessionDetail) GetAgentOk() (*string, bool)`

GetAgentOk returns a tuple with the Agent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgent

`func (o *AgentSessionDetail) SetAgent(v string)`

SetAgent sets Agent field to given value.

### HasAgent

`func (o *AgentSessionDetail) HasAgent() bool`

HasAgent returns a boolean if a field has been set.

### GetBase

`func (o *AgentSessionDetail) GetBase() string`

GetBase returns the Base field if non-nil, zero value otherwise.

### GetBaseOk

`func (o *AgentSessionDetail) GetBaseOk() (*string, bool)`

GetBaseOk returns a tuple with the Base field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBase

`func (o *AgentSessionDetail) SetBase(v string)`

SetBase sets Base field to given value.

### HasBase

`func (o *AgentSessionDetail) HasBase() bool`

HasBase returns a boolean if a field has been set.

### GetBranch

`func (o *AgentSessionDetail) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *AgentSessionDetail) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *AgentSessionDetail) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *AgentSessionDetail) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetChildSessions

`func (o *AgentSessionDetail) GetChildSessions() []AgentSessionView`

GetChildSessions returns the ChildSessions field if non-nil, zero value otherwise.

### GetChildSessionsOk

`func (o *AgentSessionDetail) GetChildSessionsOk() (*[]AgentSessionView, bool)`

GetChildSessionsOk returns a tuple with the ChildSessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildSessions

`func (o *AgentSessionDetail) SetChildSessions(v []AgentSessionView)`

SetChildSessions sets ChildSessions field to given value.

### HasChildSessions

`func (o *AgentSessionDetail) HasChildSessions() bool`

HasChildSessions returns a boolean if a field has been set.

### GetChildren

`func (o *AgentSessionDetail) GetChildren() int64`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *AgentSessionDetail) GetChildrenOk() (*int64, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *AgentSessionDetail) SetChildren(v int64)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *AgentSessionDetail) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AgentSessionDetail) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AgentSessionDetail) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AgentSessionDetail) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AgentSessionDetail) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCwd

`func (o *AgentSessionDetail) GetCwd() string`

GetCwd returns the Cwd field if non-nil, zero value otherwise.

### GetCwdOk

`func (o *AgentSessionDetail) GetCwdOk() (*string, bool)`

GetCwdOk returns a tuple with the Cwd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCwd

`func (o *AgentSessionDetail) SetCwd(v string)`

SetCwd sets Cwd field to given value.

### HasCwd

`func (o *AgentSessionDetail) HasCwd() bool`

HasCwd returns a boolean if a field has been set.

### GetEndedAt

`func (o *AgentSessionDetail) GetEndedAt() string`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *AgentSessionDetail) GetEndedAtOk() (*string, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *AgentSessionDetail) SetEndedAt(v string)`

SetEndedAt sets EndedAt field to given value.

### HasEndedAt

`func (o *AgentSessionDetail) HasEndedAt() bool`

HasEndedAt returns a boolean if a field has been set.

### GetEnvironment

`func (o *AgentSessionDetail) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *AgentSessionDetail) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *AgentSessionDetail) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *AgentSessionDetail) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetEvents

`func (o *AgentSessionDetail) GetEvents() int64`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *AgentSessionDetail) GetEventsOk() (*int64, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *AgentSessionDetail) SetEvents(v int64)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *AgentSessionDetail) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetHost

`func (o *AgentSessionDetail) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *AgentSessionDetail) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *AgentSessionDetail) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *AgentSessionDetail) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetId

`func (o *AgentSessionDetail) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AgentSessionDetail) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AgentSessionDetail) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AgentSessionDetail) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *AgentSessionDetail) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *AgentSessionDetail) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *AgentSessionDetail) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *AgentSessionDetail) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLastEvent

`func (o *AgentSessionDetail) GetLastEvent() AgentLastEventView`

GetLastEvent returns the LastEvent field if non-nil, zero value otherwise.

### GetLastEventOk

`func (o *AgentSessionDetail) GetLastEventOk() (*AgentLastEventView, bool)`

GetLastEventOk returns a tuple with the LastEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEvent

`func (o *AgentSessionDetail) SetLastEvent(v AgentLastEventView)`

SetLastEvent sets LastEvent field to given value.

### HasLastEvent

`func (o *AgentSessionDetail) HasLastEvent() bool`

HasLastEvent returns a boolean if a field has been set.

### GetMode

`func (o *AgentSessionDetail) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *AgentSessionDetail) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *AgentSessionDetail) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *AgentSessionDetail) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetOrg

`func (o *AgentSessionDetail) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *AgentSessionDetail) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *AgentSessionDetail) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *AgentSessionDetail) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetParentSessionId

`func (o *AgentSessionDetail) GetParentSessionId() string`

GetParentSessionId returns the ParentSessionId field if non-nil, zero value otherwise.

### GetParentSessionIdOk

`func (o *AgentSessionDetail) GetParentSessionIdOk() (*string, bool)`

GetParentSessionIdOk returns a tuple with the ParentSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentSessionId

`func (o *AgentSessionDetail) SetParentSessionId(v string)`

SetParentSessionId sets ParentSessionId field to given value.

### HasParentSessionId

`func (o *AgentSessionDetail) HasParentSessionId() bool`

HasParentSessionId returns a boolean if a field has been set.

### GetPr

`func (o *AgentSessionDetail) GetPr() string`

GetPr returns the Pr field if non-nil, zero value otherwise.

### GetPrOk

`func (o *AgentSessionDetail) GetPrOk() (*string, bool)`

GetPrOk returns a tuple with the Pr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPr

`func (o *AgentSessionDetail) SetPr(v string)`

SetPr sets Pr field to given value.

### HasPr

`func (o *AgentSessionDetail) HasPr() bool`

HasPr returns a boolean if a field has been set.

### GetProgress

`func (o *AgentSessionDetail) GetProgress() AgentSessionProgress`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *AgentSessionDetail) GetProgressOk() (*AgentSessionProgress, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *AgentSessionDetail) SetProgress(v AgentSessionProgress)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *AgentSessionDetail) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetProject

`func (o *AgentSessionDetail) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *AgentSessionDetail) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *AgentSessionDetail) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *AgentSessionDetail) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetProvider

`func (o *AgentSessionDetail) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *AgentSessionDetail) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *AgentSessionDetail) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *AgentSessionDetail) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetPublished

`func (o *AgentSessionDetail) GetPublished() bool`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *AgentSessionDetail) GetPublishedOk() (*bool, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *AgentSessionDetail) SetPublished(v bool)`

SetPublished sets Published field to given value.

### HasPublished

`func (o *AgentSessionDetail) HasPublished() bool`

HasPublished returns a boolean if a field has been set.

### GetRecentEvents

`func (o *AgentSessionDetail) GetRecentEvents() []AgentEventView`

GetRecentEvents returns the RecentEvents field if non-nil, zero value otherwise.

### GetRecentEventsOk

`func (o *AgentSessionDetail) GetRecentEventsOk() (*[]AgentEventView, bool)`

GetRecentEventsOk returns a tuple with the RecentEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecentEvents

`func (o *AgentSessionDetail) SetRecentEvents(v []AgentEventView)`

SetRecentEvents sets RecentEvents field to given value.

### HasRecentEvents

`func (o *AgentSessionDetail) HasRecentEvents() bool`

HasRecentEvents returns a boolean if a field has been set.

### GetRepo

`func (o *AgentSessionDetail) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *AgentSessionDetail) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *AgentSessionDetail) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *AgentSessionDetail) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### GetRoom

`func (o *AgentSessionDetail) GetRoom() string`

GetRoom returns the Room field if non-nil, zero value otherwise.

### GetRoomOk

`func (o *AgentSessionDetail) GetRoomOk() (*string, bool)`

GetRoomOk returns a tuple with the Room field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoom

`func (o *AgentSessionDetail) SetRoom(v string)`

SetRoom sets Room field to given value.

### HasRoom

`func (o *AgentSessionDetail) HasRoom() bool`

HasRoom returns a boolean if a field has been set.

### GetRootSessionId

`func (o *AgentSessionDetail) GetRootSessionId() string`

GetRootSessionId returns the RootSessionId field if non-nil, zero value otherwise.

### GetRootSessionIdOk

`func (o *AgentSessionDetail) GetRootSessionIdOk() (*string, bool)`

GetRootSessionIdOk returns a tuple with the RootSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootSessionId

`func (o *AgentSessionDetail) SetRootSessionId(v string)`

SetRootSessionId sets RootSessionId field to given value.

### HasRootSessionId

`func (o *AgentSessionDetail) HasRootSessionId() bool`

HasRootSessionId returns a boolean if a field has been set.

### GetSandbox

`func (o *AgentSessionDetail) GetSandbox() string`

GetSandbox returns the Sandbox field if non-nil, zero value otherwise.

### GetSandboxOk

`func (o *AgentSessionDetail) GetSandboxOk() (*string, bool)`

GetSandboxOk returns a tuple with the Sandbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSandbox

`func (o *AgentSessionDetail) SetSandbox(v string)`

SetSandbox sets Sandbox field to given value.

### HasSandbox

`func (o *AgentSessionDetail) HasSandbox() bool`

HasSandbox returns a boolean if a field has been set.

### GetStartedAt

`func (o *AgentSessionDetail) GetStartedAt() string`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *AgentSessionDetail) GetStartedAtOk() (*string, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *AgentSessionDetail) SetStartedAt(v string)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *AgentSessionDetail) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetStatus

`func (o *AgentSessionDetail) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AgentSessionDetail) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AgentSessionDetail) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AgentSessionDetail) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTarget

`func (o *AgentSessionDetail) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *AgentSessionDetail) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *AgentSessionDetail) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *AgentSessionDetail) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetTaskRunId

`func (o *AgentSessionDetail) GetTaskRunId() string`

GetTaskRunId returns the TaskRunId field if non-nil, zero value otherwise.

### GetTaskRunIdOk

`func (o *AgentSessionDetail) GetTaskRunIdOk() (*string, bool)`

GetTaskRunIdOk returns a tuple with the TaskRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskRunId

`func (o *AgentSessionDetail) SetTaskRunId(v string)`

SetTaskRunId sets TaskRunId field to given value.

### HasTaskRunId

`func (o *AgentSessionDetail) HasTaskRunId() bool`

HasTaskRunId returns a boolean if a field has been set.

### GetTaskWorkflowId

`func (o *AgentSessionDetail) GetTaskWorkflowId() string`

GetTaskWorkflowId returns the TaskWorkflowId field if non-nil, zero value otherwise.

### GetTaskWorkflowIdOk

`func (o *AgentSessionDetail) GetTaskWorkflowIdOk() (*string, bool)`

GetTaskWorkflowIdOk returns a tuple with the TaskWorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskWorkflowId

`func (o *AgentSessionDetail) SetTaskWorkflowId(v string)`

SetTaskWorkflowId sets TaskWorkflowId field to given value.

### HasTaskWorkflowId

`func (o *AgentSessionDetail) HasTaskWorkflowId() bool`

HasTaskWorkflowId returns a boolean if a field has been set.

### GetTerminal

`func (o *AgentSessionDetail) GetTerminal() string`

GetTerminal returns the Terminal field if non-nil, zero value otherwise.

### GetTerminalOk

`func (o *AgentSessionDetail) GetTerminalOk() (*string, bool)`

GetTerminalOk returns a tuple with the Terminal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminal

`func (o *AgentSessionDetail) SetTerminal(v string)`

SetTerminal sets Terminal field to given value.

### HasTerminal

`func (o *AgentSessionDetail) HasTerminal() bool`

HasTerminal returns a boolean if a field has been set.

### GetTitle

`func (o *AgentSessionDetail) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AgentSessionDetail) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AgentSessionDetail) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AgentSessionDetail) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AgentSessionDetail) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AgentSessionDetail) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AgentSessionDetail) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AgentSessionDetail) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


