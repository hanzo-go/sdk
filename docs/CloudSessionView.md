# CloudSessionView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** |  | [optional] 
**Actor** | Pointer to **string** |  | [optional] 
**Agent** | Pointer to **string** |  | [optional] 
**Children** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**Cwd** | Pointer to **string** |  | [optional] 
**EndedAt** | Pointer to **string** |  | [optional] 
**Events** | Pointer to **int32** |  | [optional] 
**Host** | Pointer to **string** | Execution context (mission-control): the machine/repo/cwd a card shows and the run-target a session is dispatched to. Omitted when a surface didn&#39;t report it. | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**LastEvent** | Pointer to [**CloudLastEventView**](CloudLastEventView.md) | LastEvent is the compact latest-activity line for the list projection (nil in register/patch/tree responses; set by list + detail). It lets a swipe card show a live one-line preview without fetching full detail. | [optional] 
**Org** | Pointer to **string** | Org is the caller&#39;s OWN tenant, echoed so a client can build the public build URL (/builds/:org/:project) without a second call or a guess. It is never another tenant&#39;s — every read is org-scoped before it gets here. | [optional] 
**ParentSessionId** | Pointer to **string** |  | [optional] 
**Project** | Pointer to **string** | The readable build: the product this session built and whether its story is public (provenance.go). | [optional] 
**Provider** | Pointer to **string** |  | [optional] 
**Published** | Pointer to **bool** |  | [optional] 
**Repo** | Pointer to **string** |  | [optional] 
**RootSessionId** | Pointer to **string** |  | [optional] 
**StartedAt** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Target** | Pointer to **string** |  | [optional] 
**TaskRunId** | Pointer to **string** |  | [optional] 
**TaskWorkflowId** | Pointer to **string** |  | [optional] 
**Terminal** | Pointer to **string** | Terminal is where this session can be WATCHED — the URL the machine published for its live terminal. Omitted when it publishes none. | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewCloudSessionView

`func NewCloudSessionView() *CloudSessionView`

NewCloudSessionView instantiates a new CloudSessionView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudSessionViewWithDefaults

`func NewCloudSessionViewWithDefaults() *CloudSessionView`

NewCloudSessionViewWithDefaults instantiates a new CloudSessionView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *CloudSessionView) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *CloudSessionView) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *CloudSessionView) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *CloudSessionView) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetActor

`func (o *CloudSessionView) GetActor() string`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *CloudSessionView) GetActorOk() (*string, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *CloudSessionView) SetActor(v string)`

SetActor sets Actor field to given value.

### HasActor

`func (o *CloudSessionView) HasActor() bool`

HasActor returns a boolean if a field has been set.

### GetAgent

`func (o *CloudSessionView) GetAgent() string`

GetAgent returns the Agent field if non-nil, zero value otherwise.

### GetAgentOk

`func (o *CloudSessionView) GetAgentOk() (*string, bool)`

GetAgentOk returns a tuple with the Agent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgent

`func (o *CloudSessionView) SetAgent(v string)`

SetAgent sets Agent field to given value.

### HasAgent

`func (o *CloudSessionView) HasAgent() bool`

HasAgent returns a boolean if a field has been set.

### GetChildren

`func (o *CloudSessionView) GetChildren() int32`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *CloudSessionView) GetChildrenOk() (*int32, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *CloudSessionView) SetChildren(v int32)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *CloudSessionView) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CloudSessionView) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CloudSessionView) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CloudSessionView) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CloudSessionView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCwd

`func (o *CloudSessionView) GetCwd() string`

GetCwd returns the Cwd field if non-nil, zero value otherwise.

### GetCwdOk

`func (o *CloudSessionView) GetCwdOk() (*string, bool)`

GetCwdOk returns a tuple with the Cwd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCwd

`func (o *CloudSessionView) SetCwd(v string)`

SetCwd sets Cwd field to given value.

### HasCwd

`func (o *CloudSessionView) HasCwd() bool`

HasCwd returns a boolean if a field has been set.

### GetEndedAt

`func (o *CloudSessionView) GetEndedAt() string`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *CloudSessionView) GetEndedAtOk() (*string, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *CloudSessionView) SetEndedAt(v string)`

SetEndedAt sets EndedAt field to given value.

### HasEndedAt

`func (o *CloudSessionView) HasEndedAt() bool`

HasEndedAt returns a boolean if a field has been set.

### GetEvents

`func (o *CloudSessionView) GetEvents() int32`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *CloudSessionView) GetEventsOk() (*int32, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *CloudSessionView) SetEvents(v int32)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *CloudSessionView) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetHost

`func (o *CloudSessionView) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *CloudSessionView) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *CloudSessionView) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *CloudSessionView) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetId

`func (o *CloudSessionView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudSessionView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudSessionView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudSessionView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastEvent

`func (o *CloudSessionView) GetLastEvent() CloudLastEventView`

GetLastEvent returns the LastEvent field if non-nil, zero value otherwise.

### GetLastEventOk

`func (o *CloudSessionView) GetLastEventOk() (*CloudLastEventView, bool)`

GetLastEventOk returns a tuple with the LastEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEvent

`func (o *CloudSessionView) SetLastEvent(v CloudLastEventView)`

SetLastEvent sets LastEvent field to given value.

### HasLastEvent

`func (o *CloudSessionView) HasLastEvent() bool`

HasLastEvent returns a boolean if a field has been set.

### GetOrg

`func (o *CloudSessionView) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *CloudSessionView) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *CloudSessionView) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *CloudSessionView) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetParentSessionId

`func (o *CloudSessionView) GetParentSessionId() string`

GetParentSessionId returns the ParentSessionId field if non-nil, zero value otherwise.

### GetParentSessionIdOk

`func (o *CloudSessionView) GetParentSessionIdOk() (*string, bool)`

GetParentSessionIdOk returns a tuple with the ParentSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentSessionId

`func (o *CloudSessionView) SetParentSessionId(v string)`

SetParentSessionId sets ParentSessionId field to given value.

### HasParentSessionId

`func (o *CloudSessionView) HasParentSessionId() bool`

HasParentSessionId returns a boolean if a field has been set.

### GetProject

`func (o *CloudSessionView) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *CloudSessionView) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *CloudSessionView) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *CloudSessionView) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetProvider

`func (o *CloudSessionView) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *CloudSessionView) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *CloudSessionView) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *CloudSessionView) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetPublished

`func (o *CloudSessionView) GetPublished() bool`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *CloudSessionView) GetPublishedOk() (*bool, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *CloudSessionView) SetPublished(v bool)`

SetPublished sets Published field to given value.

### HasPublished

`func (o *CloudSessionView) HasPublished() bool`

HasPublished returns a boolean if a field has been set.

### GetRepo

`func (o *CloudSessionView) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *CloudSessionView) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *CloudSessionView) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *CloudSessionView) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### GetRootSessionId

`func (o *CloudSessionView) GetRootSessionId() string`

GetRootSessionId returns the RootSessionId field if non-nil, zero value otherwise.

### GetRootSessionIdOk

`func (o *CloudSessionView) GetRootSessionIdOk() (*string, bool)`

GetRootSessionIdOk returns a tuple with the RootSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootSessionId

`func (o *CloudSessionView) SetRootSessionId(v string)`

SetRootSessionId sets RootSessionId field to given value.

### HasRootSessionId

`func (o *CloudSessionView) HasRootSessionId() bool`

HasRootSessionId returns a boolean if a field has been set.

### GetStartedAt

`func (o *CloudSessionView) GetStartedAt() string`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *CloudSessionView) GetStartedAtOk() (*string, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *CloudSessionView) SetStartedAt(v string)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *CloudSessionView) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetStatus

`func (o *CloudSessionView) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudSessionView) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudSessionView) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudSessionView) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTarget

`func (o *CloudSessionView) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *CloudSessionView) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *CloudSessionView) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *CloudSessionView) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetTaskRunId

`func (o *CloudSessionView) GetTaskRunId() string`

GetTaskRunId returns the TaskRunId field if non-nil, zero value otherwise.

### GetTaskRunIdOk

`func (o *CloudSessionView) GetTaskRunIdOk() (*string, bool)`

GetTaskRunIdOk returns a tuple with the TaskRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskRunId

`func (o *CloudSessionView) SetTaskRunId(v string)`

SetTaskRunId sets TaskRunId field to given value.

### HasTaskRunId

`func (o *CloudSessionView) HasTaskRunId() bool`

HasTaskRunId returns a boolean if a field has been set.

### GetTaskWorkflowId

`func (o *CloudSessionView) GetTaskWorkflowId() string`

GetTaskWorkflowId returns the TaskWorkflowId field if non-nil, zero value otherwise.

### GetTaskWorkflowIdOk

`func (o *CloudSessionView) GetTaskWorkflowIdOk() (*string, bool)`

GetTaskWorkflowIdOk returns a tuple with the TaskWorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskWorkflowId

`func (o *CloudSessionView) SetTaskWorkflowId(v string)`

SetTaskWorkflowId sets TaskWorkflowId field to given value.

### HasTaskWorkflowId

`func (o *CloudSessionView) HasTaskWorkflowId() bool`

HasTaskWorkflowId returns a boolean if a field has been set.

### GetTerminal

`func (o *CloudSessionView) GetTerminal() string`

GetTerminal returns the Terminal field if non-nil, zero value otherwise.

### GetTerminalOk

`func (o *CloudSessionView) GetTerminalOk() (*string, bool)`

GetTerminalOk returns a tuple with the Terminal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminal

`func (o *CloudSessionView) SetTerminal(v string)`

SetTerminal sets Terminal field to given value.

### HasTerminal

`func (o *CloudSessionView) HasTerminal() bool`

HasTerminal returns a boolean if a field has been set.

### GetTitle

`func (o *CloudSessionView) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CloudSessionView) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CloudSessionView) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CloudSessionView) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CloudSessionView) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CloudSessionView) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CloudSessionView) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CloudSessionView) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


