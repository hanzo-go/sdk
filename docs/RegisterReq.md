# RegisterReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** | Account is which subscription or API account under that provider served the run, up to 256 characters. It is what lets a revoke of that login stop exactly the sessions it was paying for. | [optional] 
**Actor** | Pointer to **string** | Actor is the \&quot;org/sub\&quot; identity to record the session under, up to 256 characters. Omit it and the calling principal is used, which is almost always what you want: it is what a login revoke matches on to stop this session. | [optional] 
**Agent** | Pointer to **string** | Agent is the label the surface opening this session calls itself by (\&quot;hanzo-dev\&quot;). REQUIRED, up to 128 characters, and free text — nothing resolves it against a defined agent. | [optional] 
**Cwd** | Pointer to **string** | Cwd is the directory the session starts in, up to 1024 characters. It can be moved later, because a linked shell walks around. | [optional] 
**Host** | Pointer to **string** | Execution context — where this session runs (all optional). | [optional] 
**ParentSessionId** | Pointer to **string** | ParentSessionID makes this a subagent of that session: it inherits the parent&#39;s root, so one flow stays one tree. The parent must exist IN THE SAME ORG — a foreign or unknown id is a 400, never a tree across tenants. Empty opens a root session. | [optional] 
**Project** | Pointer to **string** | The readable build (provenance.go): which product this session builds, and whether its story may be read by the world. | [optional] 
**Provider** | Pointer to **string** | Account tag — the linked AI account this session ran under (login manager). | [optional] 
**Published** | Pointer to **bool** | Published opens this session&#39;s story to the public build route. It is refused without a Project, because that route is keyed on (org, project) — a build with no product is not a story anyone can open. False keeps it org-only. | [optional] 
**Repo** | Pointer to **string** | Repo is the code being worked on, up to 512 characters. A label the surface states; nothing resolves it against the forge. | [optional] 
**Room** | Pointer to **string** | Room is the collaborative room this run was started in (HIP-0523), so a space view can list the sessions of one room. It is PROVENANCE and is set only here: there is deliberately no way to move a session to another room, so it is absent from the patch input and from UpdateSession&#39;s SET list. | [optional] 
**Status** | Pointer to **string** | Status opens the session in one of running, paused, done or error. Empty means running. A TERMINAL status here (done, error) records a session that has already finished — its end time is stamped now — and nothing can move it afterwards. | [optional] 
**Target** | Pointer to **string** | Target names a run-target the org has registered. Unlike Host and Repo it IS resolved: a target that does not exist in this org is a 400, so a session can never claim to run on another tenant&#39;s machine. Empty names no machine. | [optional] 
**TaskRunId** | Pointer to **string** | TaskRunID is that workflow&#39;s particular run, same bound. Recorded, not resolved: this surface does not check the workflow exists. | [optional] 
**TaskWorkflowId** | Pointer to **string** | TaskWorkflowID links this session to the hanzoai/tasks workflow that executes it, up to 256 characters. Set it and control commands are forwarded to that engine; leave it and the running surface polls for them instead. | [optional] 
**Terminal** | Pointer to **string** | Terminal is the URL this session&#39;s live terminal is published at, so the console can watch it. Optional — a session that publishes nothing is still a session. | [optional] 
**Title** | Pointer to **string** | Title is the human line a card shows, up to 512 characters. Optional, and changeable later. | [optional] 

## Methods

### NewRegisterReq

`func NewRegisterReq() *RegisterReq`

NewRegisterReq instantiates a new RegisterReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterReqWithDefaults

`func NewRegisterReqWithDefaults() *RegisterReq`

NewRegisterReqWithDefaults instantiates a new RegisterReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *RegisterReq) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *RegisterReq) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *RegisterReq) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *RegisterReq) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetActor

`func (o *RegisterReq) GetActor() string`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *RegisterReq) GetActorOk() (*string, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *RegisterReq) SetActor(v string)`

SetActor sets Actor field to given value.

### HasActor

`func (o *RegisterReq) HasActor() bool`

HasActor returns a boolean if a field has been set.

### GetAgent

`func (o *RegisterReq) GetAgent() string`

GetAgent returns the Agent field if non-nil, zero value otherwise.

### GetAgentOk

`func (o *RegisterReq) GetAgentOk() (*string, bool)`

GetAgentOk returns a tuple with the Agent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgent

`func (o *RegisterReq) SetAgent(v string)`

SetAgent sets Agent field to given value.

### HasAgent

`func (o *RegisterReq) HasAgent() bool`

HasAgent returns a boolean if a field has been set.

### GetCwd

`func (o *RegisterReq) GetCwd() string`

GetCwd returns the Cwd field if non-nil, zero value otherwise.

### GetCwdOk

`func (o *RegisterReq) GetCwdOk() (*string, bool)`

GetCwdOk returns a tuple with the Cwd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCwd

`func (o *RegisterReq) SetCwd(v string)`

SetCwd sets Cwd field to given value.

### HasCwd

`func (o *RegisterReq) HasCwd() bool`

HasCwd returns a boolean if a field has been set.

### GetHost

`func (o *RegisterReq) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *RegisterReq) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *RegisterReq) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *RegisterReq) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetParentSessionId

`func (o *RegisterReq) GetParentSessionId() string`

GetParentSessionId returns the ParentSessionId field if non-nil, zero value otherwise.

### GetParentSessionIdOk

`func (o *RegisterReq) GetParentSessionIdOk() (*string, bool)`

GetParentSessionIdOk returns a tuple with the ParentSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentSessionId

`func (o *RegisterReq) SetParentSessionId(v string)`

SetParentSessionId sets ParentSessionId field to given value.

### HasParentSessionId

`func (o *RegisterReq) HasParentSessionId() bool`

HasParentSessionId returns a boolean if a field has been set.

### GetProject

`func (o *RegisterReq) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *RegisterReq) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *RegisterReq) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *RegisterReq) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetProvider

`func (o *RegisterReq) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *RegisterReq) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *RegisterReq) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *RegisterReq) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetPublished

`func (o *RegisterReq) GetPublished() bool`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *RegisterReq) GetPublishedOk() (*bool, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *RegisterReq) SetPublished(v bool)`

SetPublished sets Published field to given value.

### HasPublished

`func (o *RegisterReq) HasPublished() bool`

HasPublished returns a boolean if a field has been set.

### GetRepo

`func (o *RegisterReq) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *RegisterReq) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *RegisterReq) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *RegisterReq) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### GetRoom

`func (o *RegisterReq) GetRoom() string`

GetRoom returns the Room field if non-nil, zero value otherwise.

### GetRoomOk

`func (o *RegisterReq) GetRoomOk() (*string, bool)`

GetRoomOk returns a tuple with the Room field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoom

`func (o *RegisterReq) SetRoom(v string)`

SetRoom sets Room field to given value.

### HasRoom

`func (o *RegisterReq) HasRoom() bool`

HasRoom returns a boolean if a field has been set.

### GetStatus

`func (o *RegisterReq) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RegisterReq) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RegisterReq) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RegisterReq) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTarget

`func (o *RegisterReq) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *RegisterReq) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *RegisterReq) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *RegisterReq) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetTaskRunId

`func (o *RegisterReq) GetTaskRunId() string`

GetTaskRunId returns the TaskRunId field if non-nil, zero value otherwise.

### GetTaskRunIdOk

`func (o *RegisterReq) GetTaskRunIdOk() (*string, bool)`

GetTaskRunIdOk returns a tuple with the TaskRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskRunId

`func (o *RegisterReq) SetTaskRunId(v string)`

SetTaskRunId sets TaskRunId field to given value.

### HasTaskRunId

`func (o *RegisterReq) HasTaskRunId() bool`

HasTaskRunId returns a boolean if a field has been set.

### GetTaskWorkflowId

`func (o *RegisterReq) GetTaskWorkflowId() string`

GetTaskWorkflowId returns the TaskWorkflowId field if non-nil, zero value otherwise.

### GetTaskWorkflowIdOk

`func (o *RegisterReq) GetTaskWorkflowIdOk() (*string, bool)`

GetTaskWorkflowIdOk returns a tuple with the TaskWorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskWorkflowId

`func (o *RegisterReq) SetTaskWorkflowId(v string)`

SetTaskWorkflowId sets TaskWorkflowId field to given value.

### HasTaskWorkflowId

`func (o *RegisterReq) HasTaskWorkflowId() bool`

HasTaskWorkflowId returns a boolean if a field has been set.

### GetTerminal

`func (o *RegisterReq) GetTerminal() string`

GetTerminal returns the Terminal field if non-nil, zero value otherwise.

### GetTerminalOk

`func (o *RegisterReq) GetTerminalOk() (*string, bool)`

GetTerminalOk returns a tuple with the Terminal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminal

`func (o *RegisterReq) SetTerminal(v string)`

SetTerminal sets Terminal field to given value.

### HasTerminal

`func (o *RegisterReq) HasTerminal() bool`

HasTerminal returns a boolean if a field has been set.

### GetTitle

`func (o *RegisterReq) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *RegisterReq) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *RegisterReq) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *RegisterReq) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


