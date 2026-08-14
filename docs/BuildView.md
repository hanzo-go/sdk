# BuildView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Agent** | Pointer to **string** |  | [optional] 
**EndedAt** | Pointer to **string** |  | [optional] 
**Model** | Pointer to **string** |  | [optional] 
**Org** | Pointer to **string** |  | [optional] 
**Project** | Pointer to **string** |  | [optional] 
**Repo** | Pointer to **string** |  | [optional] 
**Session** | Pointer to **string** |  | [optional] 
**StartedAt** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Turns** | Pointer to [**[]BuildTurn**](BuildTurn.md) |  | [optional] 
**Verify** | Pointer to **string** | Verify is the exact command that re-derives every commit binding below straight from git, so nothing here has to be taken on trust. | [optional] 

## Methods

### NewBuildView

`func NewBuildView() *BuildView`

NewBuildView instantiates a new BuildView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuildViewWithDefaults

`func NewBuildViewWithDefaults() *BuildView`

NewBuildViewWithDefaults instantiates a new BuildView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgent

`func (o *BuildView) GetAgent() string`

GetAgent returns the Agent field if non-nil, zero value otherwise.

### GetAgentOk

`func (o *BuildView) GetAgentOk() (*string, bool)`

GetAgentOk returns a tuple with the Agent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgent

`func (o *BuildView) SetAgent(v string)`

SetAgent sets Agent field to given value.

### HasAgent

`func (o *BuildView) HasAgent() bool`

HasAgent returns a boolean if a field has been set.

### GetEndedAt

`func (o *BuildView) GetEndedAt() string`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *BuildView) GetEndedAtOk() (*string, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *BuildView) SetEndedAt(v string)`

SetEndedAt sets EndedAt field to given value.

### HasEndedAt

`func (o *BuildView) HasEndedAt() bool`

HasEndedAt returns a boolean if a field has been set.

### GetModel

`func (o *BuildView) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *BuildView) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *BuildView) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *BuildView) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetOrg

`func (o *BuildView) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *BuildView) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *BuildView) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *BuildView) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetProject

`func (o *BuildView) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *BuildView) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *BuildView) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *BuildView) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetRepo

`func (o *BuildView) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *BuildView) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *BuildView) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *BuildView) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### GetSession

`func (o *BuildView) GetSession() string`

GetSession returns the Session field if non-nil, zero value otherwise.

### GetSessionOk

`func (o *BuildView) GetSessionOk() (*string, bool)`

GetSessionOk returns a tuple with the Session field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSession

`func (o *BuildView) SetSession(v string)`

SetSession sets Session field to given value.

### HasSession

`func (o *BuildView) HasSession() bool`

HasSession returns a boolean if a field has been set.

### GetStartedAt

`func (o *BuildView) GetStartedAt() string`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *BuildView) GetStartedAtOk() (*string, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *BuildView) SetStartedAt(v string)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *BuildView) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetStatus

`func (o *BuildView) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BuildView) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BuildView) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BuildView) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTitle

`func (o *BuildView) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *BuildView) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *BuildView) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *BuildView) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetTurns

`func (o *BuildView) GetTurns() []BuildTurn`

GetTurns returns the Turns field if non-nil, zero value otherwise.

### GetTurnsOk

`func (o *BuildView) GetTurnsOk() (*[]BuildTurn, bool)`

GetTurnsOk returns a tuple with the Turns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTurns

`func (o *BuildView) SetTurns(v []BuildTurn)`

SetTurns sets Turns field to given value.

### HasTurns

`func (o *BuildView) HasTurns() bool`

HasTurns returns a boolean if a field has been set.

### GetVerify

`func (o *BuildView) GetVerify() string`

GetVerify returns the Verify field if non-nil, zero value otherwise.

### GetVerifyOk

`func (o *BuildView) GetVerifyOk() (*string, bool)`

GetVerifyOk returns a tuple with the Verify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerify

`func (o *BuildView) SetVerify(v string)`

SetVerify sets Verify field to given value.

### HasVerify

`func (o *BuildView) HasVerify() bool`

HasVerify returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


