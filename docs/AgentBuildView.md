# AgentBuildView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Agent** | Pointer to **string** | Agent is the label the surface that did the work calls itself by. | [optional] 
**EndedAt** | Pointer to **string** | EndedAt is when it finished, same format. Empty means it has not — the build is still going. | [optional] 
**Model** | Pointer to **string** | Model is the model that did the work, taken from the FIRST turn whose body names one — a transcript states it, this route does not resolve it. Empty when no turn said. | [optional] 
**Org** | Pointer to **string** | Org is the org that published this build, echoed from the URL. It is part of the build&#39;s public ADDRESS and not a tenant key — this route is anonymous, and the only rows it can reach are ones an author explicitly published. | [optional] 
**Project** | Pointer to **string** | Project is the product&#39;s slug, the other half of that address. | [optional] 
**Repo** | Pointer to **string** | Repo is the repository the work was done in, as the session reported it. | [optional] 
**Session** | Pointer to **string** | Session is the id of the agent session this story IS — the same value a produced commit carries in its &#x60;Hanzo-Session:&#x60; trailer, which is what ties the repository&#39;s history to this page. | [optional] 
**StartedAt** | Pointer to **string** | StartedAt is when the session opened, RFC 3339 in UTC. | [optional] 
**Status** | Pointer to **string** | Status is the session&#39;s own: running, paused, done or error. A build can be read while it is still being written, so this is not always terminal — and an &#x60;error&#x60; build is still a readable story, not a missing page. | [optional] 
**Title** | Pointer to **string** | Title is the human line the session was opened or renamed with. Empty when nobody gave it one. | [optional] 
**Turns** | Pointer to [**[]AgentBuildTurn**](AgentBuildTurn.md) | Turns is the whole transcript, oldest first, capped at 1000: a published build is a story to read down, not an archive to page. | [optional] 
**Verify** | Pointer to **string** | Verify is the exact command that re-derives every commit binding below straight from git, so nothing here has to be taken on trust. | [optional] 

## Methods

### NewAgentBuildView

`func NewAgentBuildView() *AgentBuildView`

NewAgentBuildView instantiates a new AgentBuildView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentBuildViewWithDefaults

`func NewAgentBuildViewWithDefaults() *AgentBuildView`

NewAgentBuildViewWithDefaults instantiates a new AgentBuildView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgent

`func (o *AgentBuildView) GetAgent() string`

GetAgent returns the Agent field if non-nil, zero value otherwise.

### GetAgentOk

`func (o *AgentBuildView) GetAgentOk() (*string, bool)`

GetAgentOk returns a tuple with the Agent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgent

`func (o *AgentBuildView) SetAgent(v string)`

SetAgent sets Agent field to given value.

### HasAgent

`func (o *AgentBuildView) HasAgent() bool`

HasAgent returns a boolean if a field has been set.

### GetEndedAt

`func (o *AgentBuildView) GetEndedAt() string`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *AgentBuildView) GetEndedAtOk() (*string, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *AgentBuildView) SetEndedAt(v string)`

SetEndedAt sets EndedAt field to given value.

### HasEndedAt

`func (o *AgentBuildView) HasEndedAt() bool`

HasEndedAt returns a boolean if a field has been set.

### GetModel

`func (o *AgentBuildView) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *AgentBuildView) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *AgentBuildView) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *AgentBuildView) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetOrg

`func (o *AgentBuildView) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *AgentBuildView) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *AgentBuildView) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *AgentBuildView) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetProject

`func (o *AgentBuildView) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *AgentBuildView) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *AgentBuildView) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *AgentBuildView) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetRepo

`func (o *AgentBuildView) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *AgentBuildView) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *AgentBuildView) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *AgentBuildView) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### GetSession

`func (o *AgentBuildView) GetSession() string`

GetSession returns the Session field if non-nil, zero value otherwise.

### GetSessionOk

`func (o *AgentBuildView) GetSessionOk() (*string, bool)`

GetSessionOk returns a tuple with the Session field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSession

`func (o *AgentBuildView) SetSession(v string)`

SetSession sets Session field to given value.

### HasSession

`func (o *AgentBuildView) HasSession() bool`

HasSession returns a boolean if a field has been set.

### GetStartedAt

`func (o *AgentBuildView) GetStartedAt() string`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *AgentBuildView) GetStartedAtOk() (*string, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *AgentBuildView) SetStartedAt(v string)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *AgentBuildView) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetStatus

`func (o *AgentBuildView) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AgentBuildView) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AgentBuildView) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AgentBuildView) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTitle

`func (o *AgentBuildView) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AgentBuildView) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AgentBuildView) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AgentBuildView) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetTurns

`func (o *AgentBuildView) GetTurns() []AgentBuildTurn`

GetTurns returns the Turns field if non-nil, zero value otherwise.

### GetTurnsOk

`func (o *AgentBuildView) GetTurnsOk() (*[]AgentBuildTurn, bool)`

GetTurnsOk returns a tuple with the Turns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTurns

`func (o *AgentBuildView) SetTurns(v []AgentBuildTurn)`

SetTurns sets Turns field to given value.

### HasTurns

`func (o *AgentBuildView) HasTurns() bool`

HasTurns returns a boolean if a field has been set.

### GetVerify

`func (o *AgentBuildView) GetVerify() string`

GetVerify returns the Verify field if non-nil, zero value otherwise.

### GetVerifyOk

`func (o *AgentBuildView) GetVerifyOk() (*string, bool)`

GetVerifyOk returns a tuple with the Verify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerify

`func (o *AgentBuildView) SetVerify(v string)`

SetVerify sets Verify field to given value.

### HasVerify

`func (o *AgentBuildView) HasVerify() bool`

HasVerify returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


