# SessionView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** | Account is which subscription or API account under that provider served it. Together with Provider it is what a login revoke matches on to stop the sessions a withdrawn account was paying for. | [optional] 
**Actor** | Pointer to **string** | Actor is WHO this session belongs to, as \&quot;org/sub\&quot; — the same identity a run is billed under. A register that names none takes the calling principal. It is what scopes a login revoke, so a session with the wrong actor is a session the right person cannot stop. | [optional] 
**Agent** | Pointer to **string** | Agent is the label the surface running this session calls itself by (\&quot;hanzo-dev\&quot;), up to 128 characters. Required at register. It is free text, not a reference: it need not name a defined agent, and nothing resolves it. | [optional] 
**Children** | Pointer to **int32** | Children is the DIRECT fan-out — how many sessions name this one as parent — and not the size of the subtree. Read the tree for that. | [optional] 
**CreatedAt** | Pointer to **string** | CreatedAt is when the row was written, same format. Every path that opens a session stamps it and StartedAt from one clock reading, so the two are equal on every session this surface has ever produced. | [optional] 
**Cwd** | Pointer to **string** | Cwd is the directory the session is working in NOW, not the one it started in: a linked shell moves around, and a card showing where &#x60;hanzo link&#x60; was run answers \&quot;which work is this\&quot; with something that was true once. | [optional] 
**EndedAt** | Pointer to **string** | EndedAt is when it reached done or error, same format. Empty while it is still running or paused, which is how absence reads here: not over yet. | [optional] 
**Events** | Pointer to **int32** | Events is how many turns the session&#39;s log holds, counted at read time. It is the whole log, however few of them RecentEvents carries. | [optional] 
**Host** | Pointer to **string** | Execution context (mission-control): the machine/repo/cwd a card shows and the run-target a session is dispatched to. Omitted when a surface didn&#39;t report it. | [optional] 
**Id** | Pointer to **string** | ID is the session&#39;s handle, minted here as \&quot;sess_\&quot; + 32 hex characters. Every later read, patch, event append and control command is addressed with it, and a caller cannot choose it. | [optional] 
**LastEvent** | Pointer to [**LastEventView**](LastEventView.md) | LastEvent is the compact latest-activity line for the list projection (nil in register/patch/tree responses; set by list + detail). It lets a swipe card show a live one-line preview without fetching full detail. | [optional] 
**Org** | Pointer to **string** | Org is the caller&#39;s OWN tenant, echoed so a client can build the public build URL (/builds/:org/:project) without a second call or a guess. It is never another tenant&#39;s — every read is org-scoped before it gets here. | [optional] 
**ParentSessionId** | Pointer to **string** | ParentSessionID is the session that spawned this one, making this a subagent of it. Empty means this session is a root — a flow of its own. A parent always belongs to the same org, so a tree never crosses a tenant. | [optional] 
**Progress** | Pointer to [**SessionProgress**](SessionProgress.md) | Progress is how far along this run is — a share of its goal, a phase, and a line saying what it is doing. Always present, so a board never branches on whether it is there; &#x60;phase&#x60; says \&quot;unknown\&quot; when nothing has estimated it. It is a MODEL ESTIMATE wherever &#x60;estimated&#x60; is true, and the row&#39;s own word where it is false. See progress.go. | [optional] 
**Project** | Pointer to **string** | The readable build: the product this session built and whether its story is public (provenance.go). | [optional] 
**Provider** | Pointer to **string** | Provider is the linked AI account&#39;s provider (claude | codex | hanzo | …) that served this run. Empty when the surface did not say. | [optional] 
**Published** | Pointer to **bool** | Published is the author&#39;s decision to let anyone read this session&#39;s story at the public build route. It only ever widens READ access to a session that already exists and grants nothing else; false, an unpublished session is invisible there no matter who asks. It cannot be true without a Project, because that route is keyed on (org, project). | [optional] 
**Repo** | Pointer to **string** | Repo is the code the session is working on, as the surface reported it. It is truth the SURFACE states, so it is a label rather than something resolved here. | [optional] 
**Room** | Pointer to **string** | Room is the collaborative room this run was started in (HIP-0523), empty when it came from anywhere else — a CLI, a schedule, an API call. It is what lets a space view show the runs of one room beside its messages. | [optional] 
**RootSessionId** | Pointer to **string** | RootSessionID is the top of this session&#39;s tree, inherited from the parent and shared by every node in one flow. A root session&#39;s own id, when it has no parent. It is the key one indexed read pulls a whole flow by, and what ?root&#x3D; narrows a list or a stream to. | [optional] 
**StartedAt** | Pointer to **string** | StartedAt is when the session opened, RFC 3339 in UTC to the second. | [optional] 
**Status** | Pointer to **string** | Status is one of exactly four: running, paused, done, error. running and paused are LIVE; done and error are TERMINAL and monotonic — once a session reaches one it can never go back, because reopening a finished run would fabricate liveness. A control command never moves it: the surface running the agent reports the new status, and until it does the command is only recorded. | [optional] 
**Target** | Pointer to **string** | Target is the registered run-target this session is dispatched to — a machine the org claimed, resolved same-org when it was set, so it can never point at another tenant&#39;s computer. Empty means the session names no machine. | [optional] 
**TaskRunId** | Pointer to **string** | TaskRunID is that workflow&#39;s particular run. A workflow is the definition and a run is one execution of it, which is why both are carried. | [optional] 
**TaskWorkflowId** | Pointer to **string** | TaskWorkflowID is the hanzoai/tasks durable workflow that actually EXECUTES this session — this registry is the view, control and stream layer over it. Set, a control command is FORWARDED to that engine; empty, the running surface polls for commands instead, which is every session today. | [optional] 
**Terminal** | Pointer to **string** | Terminal is where this session can be WATCHED — the URL the machine published for its live terminal. Omitted when it publishes none. | [optional] 
**Title** | Pointer to **string** | Title is the human line a card shows (\&quot;ship the landing page\&quot;), up to 512 characters. Free text, and the one field a surface may rewrite as the work turns out to be something else. | [optional] 
**UpdatedAt** | Pointer to **string** | UpdatedAt is the session&#39;s last-activity clock, same format. It moves on a write to the row — a status, a title, a re-dispatch — AND on every appended turn, because the append bumps it in the same transaction. The list is ordered on CreatedAt, so this is the field that says whether a session is still saying anything. | [optional] 

## Methods

### NewSessionView

`func NewSessionView() *SessionView`

NewSessionView instantiates a new SessionView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionViewWithDefaults

`func NewSessionViewWithDefaults() *SessionView`

NewSessionViewWithDefaults instantiates a new SessionView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *SessionView) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *SessionView) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *SessionView) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *SessionView) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetActor

`func (o *SessionView) GetActor() string`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *SessionView) GetActorOk() (*string, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *SessionView) SetActor(v string)`

SetActor sets Actor field to given value.

### HasActor

`func (o *SessionView) HasActor() bool`

HasActor returns a boolean if a field has been set.

### GetAgent

`func (o *SessionView) GetAgent() string`

GetAgent returns the Agent field if non-nil, zero value otherwise.

### GetAgentOk

`func (o *SessionView) GetAgentOk() (*string, bool)`

GetAgentOk returns a tuple with the Agent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgent

`func (o *SessionView) SetAgent(v string)`

SetAgent sets Agent field to given value.

### HasAgent

`func (o *SessionView) HasAgent() bool`

HasAgent returns a boolean if a field has been set.

### GetChildren

`func (o *SessionView) GetChildren() int32`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *SessionView) GetChildrenOk() (*int32, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *SessionView) SetChildren(v int32)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *SessionView) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SessionView) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SessionView) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SessionView) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SessionView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCwd

`func (o *SessionView) GetCwd() string`

GetCwd returns the Cwd field if non-nil, zero value otherwise.

### GetCwdOk

`func (o *SessionView) GetCwdOk() (*string, bool)`

GetCwdOk returns a tuple with the Cwd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCwd

`func (o *SessionView) SetCwd(v string)`

SetCwd sets Cwd field to given value.

### HasCwd

`func (o *SessionView) HasCwd() bool`

HasCwd returns a boolean if a field has been set.

### GetEndedAt

`func (o *SessionView) GetEndedAt() string`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *SessionView) GetEndedAtOk() (*string, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *SessionView) SetEndedAt(v string)`

SetEndedAt sets EndedAt field to given value.

### HasEndedAt

`func (o *SessionView) HasEndedAt() bool`

HasEndedAt returns a boolean if a field has been set.

### GetEvents

`func (o *SessionView) GetEvents() int32`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *SessionView) GetEventsOk() (*int32, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *SessionView) SetEvents(v int32)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *SessionView) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetHost

`func (o *SessionView) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *SessionView) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *SessionView) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *SessionView) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetId

`func (o *SessionView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SessionView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SessionView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SessionView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastEvent

`func (o *SessionView) GetLastEvent() LastEventView`

GetLastEvent returns the LastEvent field if non-nil, zero value otherwise.

### GetLastEventOk

`func (o *SessionView) GetLastEventOk() (*LastEventView, bool)`

GetLastEventOk returns a tuple with the LastEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEvent

`func (o *SessionView) SetLastEvent(v LastEventView)`

SetLastEvent sets LastEvent field to given value.

### HasLastEvent

`func (o *SessionView) HasLastEvent() bool`

HasLastEvent returns a boolean if a field has been set.

### GetOrg

`func (o *SessionView) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *SessionView) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *SessionView) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *SessionView) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetParentSessionId

`func (o *SessionView) GetParentSessionId() string`

GetParentSessionId returns the ParentSessionId field if non-nil, zero value otherwise.

### GetParentSessionIdOk

`func (o *SessionView) GetParentSessionIdOk() (*string, bool)`

GetParentSessionIdOk returns a tuple with the ParentSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentSessionId

`func (o *SessionView) SetParentSessionId(v string)`

SetParentSessionId sets ParentSessionId field to given value.

### HasParentSessionId

`func (o *SessionView) HasParentSessionId() bool`

HasParentSessionId returns a boolean if a field has been set.

### GetProgress

`func (o *SessionView) GetProgress() SessionProgress`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *SessionView) GetProgressOk() (*SessionProgress, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *SessionView) SetProgress(v SessionProgress)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *SessionView) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetProject

`func (o *SessionView) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *SessionView) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *SessionView) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *SessionView) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetProvider

`func (o *SessionView) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *SessionView) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *SessionView) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *SessionView) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetPublished

`func (o *SessionView) GetPublished() bool`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *SessionView) GetPublishedOk() (*bool, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *SessionView) SetPublished(v bool)`

SetPublished sets Published field to given value.

### HasPublished

`func (o *SessionView) HasPublished() bool`

HasPublished returns a boolean if a field has been set.

### GetRepo

`func (o *SessionView) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *SessionView) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *SessionView) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *SessionView) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### GetRoom

`func (o *SessionView) GetRoom() string`

GetRoom returns the Room field if non-nil, zero value otherwise.

### GetRoomOk

`func (o *SessionView) GetRoomOk() (*string, bool)`

GetRoomOk returns a tuple with the Room field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoom

`func (o *SessionView) SetRoom(v string)`

SetRoom sets Room field to given value.

### HasRoom

`func (o *SessionView) HasRoom() bool`

HasRoom returns a boolean if a field has been set.

### GetRootSessionId

`func (o *SessionView) GetRootSessionId() string`

GetRootSessionId returns the RootSessionId field if non-nil, zero value otherwise.

### GetRootSessionIdOk

`func (o *SessionView) GetRootSessionIdOk() (*string, bool)`

GetRootSessionIdOk returns a tuple with the RootSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootSessionId

`func (o *SessionView) SetRootSessionId(v string)`

SetRootSessionId sets RootSessionId field to given value.

### HasRootSessionId

`func (o *SessionView) HasRootSessionId() bool`

HasRootSessionId returns a boolean if a field has been set.

### GetStartedAt

`func (o *SessionView) GetStartedAt() string`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *SessionView) GetStartedAtOk() (*string, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *SessionView) SetStartedAt(v string)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *SessionView) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetStatus

`func (o *SessionView) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SessionView) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SessionView) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SessionView) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTarget

`func (o *SessionView) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *SessionView) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *SessionView) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *SessionView) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetTaskRunId

`func (o *SessionView) GetTaskRunId() string`

GetTaskRunId returns the TaskRunId field if non-nil, zero value otherwise.

### GetTaskRunIdOk

`func (o *SessionView) GetTaskRunIdOk() (*string, bool)`

GetTaskRunIdOk returns a tuple with the TaskRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskRunId

`func (o *SessionView) SetTaskRunId(v string)`

SetTaskRunId sets TaskRunId field to given value.

### HasTaskRunId

`func (o *SessionView) HasTaskRunId() bool`

HasTaskRunId returns a boolean if a field has been set.

### GetTaskWorkflowId

`func (o *SessionView) GetTaskWorkflowId() string`

GetTaskWorkflowId returns the TaskWorkflowId field if non-nil, zero value otherwise.

### GetTaskWorkflowIdOk

`func (o *SessionView) GetTaskWorkflowIdOk() (*string, bool)`

GetTaskWorkflowIdOk returns a tuple with the TaskWorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskWorkflowId

`func (o *SessionView) SetTaskWorkflowId(v string)`

SetTaskWorkflowId sets TaskWorkflowId field to given value.

### HasTaskWorkflowId

`func (o *SessionView) HasTaskWorkflowId() bool`

HasTaskWorkflowId returns a boolean if a field has been set.

### GetTerminal

`func (o *SessionView) GetTerminal() string`

GetTerminal returns the Terminal field if non-nil, zero value otherwise.

### GetTerminalOk

`func (o *SessionView) GetTerminalOk() (*string, bool)`

GetTerminalOk returns a tuple with the Terminal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminal

`func (o *SessionView) SetTerminal(v string)`

SetTerminal sets Terminal field to given value.

### HasTerminal

`func (o *SessionView) HasTerminal() bool`

HasTerminal returns a boolean if a field has been set.

### GetTitle

`func (o *SessionView) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SessionView) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SessionView) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *SessionView) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *SessionView) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SessionView) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SessionView) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *SessionView) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


