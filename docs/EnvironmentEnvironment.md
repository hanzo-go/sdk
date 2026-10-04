# EnvironmentEnvironment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Install** | Pointer to **string** | Install is the shell a run executes in its fresh checkout before the agent starts, from the repository root. Empty installs nothing. | [optional] 
**Proposal** | Pointer to [**EnvironmentProposal**](EnvironmentProposal.md) | Proposal is what the last setup run found, until it is saved or replaced. | [optional] 
**Repo** | Pointer to **string** | Repo is the codebase: a repository&#39;s name in the caller&#39;s org. | [optional] 
**Secrets** | Pointer to **[]string** | Secrets are the names set on this codebase, each exported to the run&#39;s commands under that name. Values are never answered. | [optional] 
**Session** | Pointer to **string** | Session is the setup run that produced the proposal, when there is one. | [optional] 
**Start** | Pointer to **string** | Start is the one command a run leaves running in the background once the install has finished — a dev server, a database. Empty starts nothing. | [optional] 
**State** | Pointer to **string** | State is &#x60;none&#x60; when nothing is saved or proposed, &#x60;proposed&#x60; when a setup run&#39;s answer is waiting to be reviewed, and &#x60;ready&#x60; otherwise. | [optional] 
**UpdatedAt** | Pointer to **string** | UpdatedAt is when this environment last changed, RFC 3339. Absent for none. | [optional] 

## Methods

### NewEnvironmentEnvironment

`func NewEnvironmentEnvironment() *EnvironmentEnvironment`

NewEnvironmentEnvironment instantiates a new EnvironmentEnvironment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentEnvironmentWithDefaults

`func NewEnvironmentEnvironmentWithDefaults() *EnvironmentEnvironment`

NewEnvironmentEnvironmentWithDefaults instantiates a new EnvironmentEnvironment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstall

`func (o *EnvironmentEnvironment) GetInstall() string`

GetInstall returns the Install field if non-nil, zero value otherwise.

### GetInstallOk

`func (o *EnvironmentEnvironment) GetInstallOk() (*string, bool)`

GetInstallOk returns a tuple with the Install field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstall

`func (o *EnvironmentEnvironment) SetInstall(v string)`

SetInstall sets Install field to given value.

### HasInstall

`func (o *EnvironmentEnvironment) HasInstall() bool`

HasInstall returns a boolean if a field has been set.

### GetProposal

`func (o *EnvironmentEnvironment) GetProposal() EnvironmentProposal`

GetProposal returns the Proposal field if non-nil, zero value otherwise.

### GetProposalOk

`func (o *EnvironmentEnvironment) GetProposalOk() (*EnvironmentProposal, bool)`

GetProposalOk returns a tuple with the Proposal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProposal

`func (o *EnvironmentEnvironment) SetProposal(v EnvironmentProposal)`

SetProposal sets Proposal field to given value.

### HasProposal

`func (o *EnvironmentEnvironment) HasProposal() bool`

HasProposal returns a boolean if a field has been set.

### GetRepo

`func (o *EnvironmentEnvironment) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *EnvironmentEnvironment) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *EnvironmentEnvironment) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *EnvironmentEnvironment) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### GetSecrets

`func (o *EnvironmentEnvironment) GetSecrets() []string`

GetSecrets returns the Secrets field if non-nil, zero value otherwise.

### GetSecretsOk

`func (o *EnvironmentEnvironment) GetSecretsOk() (*[]string, bool)`

GetSecretsOk returns a tuple with the Secrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecrets

`func (o *EnvironmentEnvironment) SetSecrets(v []string)`

SetSecrets sets Secrets field to given value.

### HasSecrets

`func (o *EnvironmentEnvironment) HasSecrets() bool`

HasSecrets returns a boolean if a field has been set.

### GetSession

`func (o *EnvironmentEnvironment) GetSession() string`

GetSession returns the Session field if non-nil, zero value otherwise.

### GetSessionOk

`func (o *EnvironmentEnvironment) GetSessionOk() (*string, bool)`

GetSessionOk returns a tuple with the Session field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSession

`func (o *EnvironmentEnvironment) SetSession(v string)`

SetSession sets Session field to given value.

### HasSession

`func (o *EnvironmentEnvironment) HasSession() bool`

HasSession returns a boolean if a field has been set.

### GetStart

`func (o *EnvironmentEnvironment) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *EnvironmentEnvironment) GetStartOk() (*string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *EnvironmentEnvironment) SetStart(v string)`

SetStart sets Start field to given value.

### HasStart

`func (o *EnvironmentEnvironment) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetState

`func (o *EnvironmentEnvironment) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *EnvironmentEnvironment) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *EnvironmentEnvironment) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *EnvironmentEnvironment) HasState() bool`

HasState returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *EnvironmentEnvironment) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EnvironmentEnvironment) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EnvironmentEnvironment) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *EnvironmentEnvironment) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


