# BuildSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Agent** | Pointer to **string** | Agent is the label the surface that did the work calls itself by. | [optional] 
**EndedAt** | Pointer to **string** | EndedAt is when it finished, same format. Empty means it is still going. | [optional] 
**Org** | Pointer to **string** | Org and Project are the build&#39;s public ADDRESS — the pair the full story is read at, and the pair a visitor sees in the URL bar. Not a tenant key: this index is anonymous and lists only what authors published. | [optional] 
**Project** | Pointer to **string** | Project is the product&#39;s slug, the second half of that address. | [optional] 
**Repo** | Pointer to **string** | Repo is the repository the work was done in, as the session reported it. | [optional] 
**Session** | Pointer to **string** | Session is the agent session behind the build, and the value its commits name in their &#x60;Hanzo-Session:&#x60; trailer. | [optional] 
**StartedAt** | Pointer to **string** | StartedAt is when the session opened, RFC 3339 in UTC. | [optional] 
**Status** | Pointer to **string** | Status is the session&#39;s own: running, paused, done or error — so a card can show a build still being written. | [optional] 
**Title** | Pointer to **string** | Title is the human line for the card. Sent even when empty, like every field here, because that is what this route has always sent. | [optional] 
**Turns** | Pointer to **int32** | Turns is HOW MANY turns the transcript holds — a COUNT, unlike the &#x60;turn&#x60; on each turn of the full story, which is that turn&#39;s position. The full read returns at most 1000 of them; this number is not capped. | [optional] 

## Methods

### NewBuildSummary

`func NewBuildSummary() *BuildSummary`

NewBuildSummary instantiates a new BuildSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuildSummaryWithDefaults

`func NewBuildSummaryWithDefaults() *BuildSummary`

NewBuildSummaryWithDefaults instantiates a new BuildSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgent

`func (o *BuildSummary) GetAgent() string`

GetAgent returns the Agent field if non-nil, zero value otherwise.

### GetAgentOk

`func (o *BuildSummary) GetAgentOk() (*string, bool)`

GetAgentOk returns a tuple with the Agent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgent

`func (o *BuildSummary) SetAgent(v string)`

SetAgent sets Agent field to given value.

### HasAgent

`func (o *BuildSummary) HasAgent() bool`

HasAgent returns a boolean if a field has been set.

### GetEndedAt

`func (o *BuildSummary) GetEndedAt() string`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *BuildSummary) GetEndedAtOk() (*string, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *BuildSummary) SetEndedAt(v string)`

SetEndedAt sets EndedAt field to given value.

### HasEndedAt

`func (o *BuildSummary) HasEndedAt() bool`

HasEndedAt returns a boolean if a field has been set.

### GetOrg

`func (o *BuildSummary) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *BuildSummary) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *BuildSummary) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *BuildSummary) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetProject

`func (o *BuildSummary) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *BuildSummary) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *BuildSummary) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *BuildSummary) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetRepo

`func (o *BuildSummary) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *BuildSummary) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *BuildSummary) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *BuildSummary) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### GetSession

`func (o *BuildSummary) GetSession() string`

GetSession returns the Session field if non-nil, zero value otherwise.

### GetSessionOk

`func (o *BuildSummary) GetSessionOk() (*string, bool)`

GetSessionOk returns a tuple with the Session field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSession

`func (o *BuildSummary) SetSession(v string)`

SetSession sets Session field to given value.

### HasSession

`func (o *BuildSummary) HasSession() bool`

HasSession returns a boolean if a field has been set.

### GetStartedAt

`func (o *BuildSummary) GetStartedAt() string`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *BuildSummary) GetStartedAtOk() (*string, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *BuildSummary) SetStartedAt(v string)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *BuildSummary) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetStatus

`func (o *BuildSummary) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BuildSummary) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BuildSummary) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BuildSummary) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTitle

`func (o *BuildSummary) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *BuildSummary) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *BuildSummary) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *BuildSummary) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetTurns

`func (o *BuildSummary) GetTurns() int32`

GetTurns returns the Turns field if non-nil, zero value otherwise.

### GetTurnsOk

`func (o *BuildSummary) GetTurnsOk() (*int32, bool)`

GetTurnsOk returns a tuple with the Turns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTurns

`func (o *BuildSummary) SetTurns(v int32)`

SetTurns sets Turns field to given value.

### HasTurns

`func (o *BuildSummary) HasTurns() bool`

HasTurns returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


