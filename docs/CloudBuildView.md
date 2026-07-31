# CloudBuildView

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
**Turns** | Pointer to [**[]CloudBuildTurn**](CloudBuildTurn.md) |  | [optional] 
**Verify** | Pointer to **string** | Verify is the exact command that re-derives every commit binding below straight from git, so nothing here has to be taken on trust. | [optional] 

## Methods

### NewCloudBuildView

`func NewCloudBuildView() *CloudBuildView`

NewCloudBuildView instantiates a new CloudBuildView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudBuildViewWithDefaults

`func NewCloudBuildViewWithDefaults() *CloudBuildView`

NewCloudBuildViewWithDefaults instantiates a new CloudBuildView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgent

`func (o *CloudBuildView) GetAgent() string`

GetAgent returns the Agent field if non-nil, zero value otherwise.

### GetAgentOk

`func (o *CloudBuildView) GetAgentOk() (*string, bool)`

GetAgentOk returns a tuple with the Agent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgent

`func (o *CloudBuildView) SetAgent(v string)`

SetAgent sets Agent field to given value.

### HasAgent

`func (o *CloudBuildView) HasAgent() bool`

HasAgent returns a boolean if a field has been set.

### GetEndedAt

`func (o *CloudBuildView) GetEndedAt() string`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *CloudBuildView) GetEndedAtOk() (*string, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *CloudBuildView) SetEndedAt(v string)`

SetEndedAt sets EndedAt field to given value.

### HasEndedAt

`func (o *CloudBuildView) HasEndedAt() bool`

HasEndedAt returns a boolean if a field has been set.

### GetModel

`func (o *CloudBuildView) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *CloudBuildView) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *CloudBuildView) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *CloudBuildView) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetOrg

`func (o *CloudBuildView) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *CloudBuildView) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *CloudBuildView) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *CloudBuildView) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetProject

`func (o *CloudBuildView) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *CloudBuildView) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *CloudBuildView) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *CloudBuildView) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetRepo

`func (o *CloudBuildView) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *CloudBuildView) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *CloudBuildView) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *CloudBuildView) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### GetSession

`func (o *CloudBuildView) GetSession() string`

GetSession returns the Session field if non-nil, zero value otherwise.

### GetSessionOk

`func (o *CloudBuildView) GetSessionOk() (*string, bool)`

GetSessionOk returns a tuple with the Session field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSession

`func (o *CloudBuildView) SetSession(v string)`

SetSession sets Session field to given value.

### HasSession

`func (o *CloudBuildView) HasSession() bool`

HasSession returns a boolean if a field has been set.

### GetStartedAt

`func (o *CloudBuildView) GetStartedAt() string`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *CloudBuildView) GetStartedAtOk() (*string, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *CloudBuildView) SetStartedAt(v string)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *CloudBuildView) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetStatus

`func (o *CloudBuildView) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudBuildView) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudBuildView) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudBuildView) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTitle

`func (o *CloudBuildView) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CloudBuildView) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CloudBuildView) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CloudBuildView) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetTurns

`func (o *CloudBuildView) GetTurns() []CloudBuildTurn`

GetTurns returns the Turns field if non-nil, zero value otherwise.

### GetTurnsOk

`func (o *CloudBuildView) GetTurnsOk() (*[]CloudBuildTurn, bool)`

GetTurnsOk returns a tuple with the Turns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTurns

`func (o *CloudBuildView) SetTurns(v []CloudBuildTurn)`

SetTurns sets Turns field to given value.

### HasTurns

`func (o *CloudBuildView) HasTurns() bool`

HasTurns returns a boolean if a field has been set.

### GetVerify

`func (o *CloudBuildView) GetVerify() string`

GetVerify returns the Verify field if non-nil, zero value otherwise.

### GetVerifyOk

`func (o *CloudBuildView) GetVerifyOk() (*string, bool)`

GetVerifyOk returns a tuple with the Verify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerify

`func (o *CloudBuildView) SetVerify(v string)`

SetVerify sets Verify field to given value.

### HasVerify

`func (o *CloudBuildView) HasVerify() bool`

HasVerify returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


