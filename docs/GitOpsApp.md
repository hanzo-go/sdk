# GitOpsApp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Automated** | Pointer to **bool** |  | [optional] 
**Health** | Pointer to **string** | Healthy|Degraded|Progressing|… | [optional] 
**History** | Pointer to [**[]GitOpsDeploy**](GitOpsDeploy.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**Operation** | Pointer to [**GitOpsOperation**](GitOpsOperation.md) |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**Project** | Pointer to **string** |  | [optional] 
**ReconciledAt** | Pointer to **string** |  | [optional] 
**RepoURL** | Pointer to **string** |  | [optional] 
**Resources** | Pointer to **int32** |  | [optional] 
**Revision** | Pointer to **string** | the commit last applied | [optional] 
**SelfHeal** | Pointer to **bool** |  | [optional] 
**Sync** | Pointer to **string** | Synced|OutOfSync|Unknown | [optional] 
**TargetRevision** | Pointer to **string** |  | [optional] 

## Methods

### NewGitOpsApp

`func NewGitOpsApp() *GitOpsApp`

NewGitOpsApp instantiates a new GitOpsApp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitOpsAppWithDefaults

`func NewGitOpsAppWithDefaults() *GitOpsApp`

NewGitOpsAppWithDefaults instantiates a new GitOpsApp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutomated

`func (o *GitOpsApp) GetAutomated() bool`

GetAutomated returns the Automated field if non-nil, zero value otherwise.

### GetAutomatedOk

`func (o *GitOpsApp) GetAutomatedOk() (*bool, bool)`

GetAutomatedOk returns a tuple with the Automated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomated

`func (o *GitOpsApp) SetAutomated(v bool)`

SetAutomated sets Automated field to given value.

### HasAutomated

`func (o *GitOpsApp) HasAutomated() bool`

HasAutomated returns a boolean if a field has been set.

### GetHealth

`func (o *GitOpsApp) GetHealth() string`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *GitOpsApp) GetHealthOk() (*string, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *GitOpsApp) SetHealth(v string)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *GitOpsApp) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetHistory

`func (o *GitOpsApp) GetHistory() []GitOpsDeploy`

GetHistory returns the History field if non-nil, zero value otherwise.

### GetHistoryOk

`func (o *GitOpsApp) GetHistoryOk() (*[]GitOpsDeploy, bool)`

GetHistoryOk returns a tuple with the History field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistory

`func (o *GitOpsApp) SetHistory(v []GitOpsDeploy)`

SetHistory sets History field to given value.

### HasHistory

`func (o *GitOpsApp) HasHistory() bool`

HasHistory returns a boolean if a field has been set.

### GetName

`func (o *GitOpsApp) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GitOpsApp) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GitOpsApp) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GitOpsApp) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *GitOpsApp) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *GitOpsApp) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *GitOpsApp) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *GitOpsApp) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetOperation

`func (o *GitOpsApp) GetOperation() GitOpsOperation`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *GitOpsApp) GetOperationOk() (*GitOpsOperation, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *GitOpsApp) SetOperation(v GitOpsOperation)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *GitOpsApp) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetPath

`func (o *GitOpsApp) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *GitOpsApp) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *GitOpsApp) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *GitOpsApp) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetProject

`func (o *GitOpsApp) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *GitOpsApp) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *GitOpsApp) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *GitOpsApp) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetReconciledAt

`func (o *GitOpsApp) GetReconciledAt() string`

GetReconciledAt returns the ReconciledAt field if non-nil, zero value otherwise.

### GetReconciledAtOk

`func (o *GitOpsApp) GetReconciledAtOk() (*string, bool)`

GetReconciledAtOk returns a tuple with the ReconciledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReconciledAt

`func (o *GitOpsApp) SetReconciledAt(v string)`

SetReconciledAt sets ReconciledAt field to given value.

### HasReconciledAt

`func (o *GitOpsApp) HasReconciledAt() bool`

HasReconciledAt returns a boolean if a field has been set.

### GetRepoURL

`func (o *GitOpsApp) GetRepoURL() string`

GetRepoURL returns the RepoURL field if non-nil, zero value otherwise.

### GetRepoURLOk

`func (o *GitOpsApp) GetRepoURLOk() (*string, bool)`

GetRepoURLOk returns a tuple with the RepoURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoURL

`func (o *GitOpsApp) SetRepoURL(v string)`

SetRepoURL sets RepoURL field to given value.

### HasRepoURL

`func (o *GitOpsApp) HasRepoURL() bool`

HasRepoURL returns a boolean if a field has been set.

### GetResources

`func (o *GitOpsApp) GetResources() int32`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *GitOpsApp) GetResourcesOk() (*int32, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *GitOpsApp) SetResources(v int32)`

SetResources sets Resources field to given value.

### HasResources

`func (o *GitOpsApp) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetRevision

`func (o *GitOpsApp) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *GitOpsApp) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *GitOpsApp) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *GitOpsApp) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetSelfHeal

`func (o *GitOpsApp) GetSelfHeal() bool`

GetSelfHeal returns the SelfHeal field if non-nil, zero value otherwise.

### GetSelfHealOk

`func (o *GitOpsApp) GetSelfHealOk() (*bool, bool)`

GetSelfHealOk returns a tuple with the SelfHeal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfHeal

`func (o *GitOpsApp) SetSelfHeal(v bool)`

SetSelfHeal sets SelfHeal field to given value.

### HasSelfHeal

`func (o *GitOpsApp) HasSelfHeal() bool`

HasSelfHeal returns a boolean if a field has been set.

### GetSync

`func (o *GitOpsApp) GetSync() string`

GetSync returns the Sync field if non-nil, zero value otherwise.

### GetSyncOk

`func (o *GitOpsApp) GetSyncOk() (*string, bool)`

GetSyncOk returns a tuple with the Sync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSync

`func (o *GitOpsApp) SetSync(v string)`

SetSync sets Sync field to given value.

### HasSync

`func (o *GitOpsApp) HasSync() bool`

HasSync returns a boolean if a field has been set.

### GetTargetRevision

`func (o *GitOpsApp) GetTargetRevision() string`

GetTargetRevision returns the TargetRevision field if non-nil, zero value otherwise.

### GetTargetRevisionOk

`func (o *GitOpsApp) GetTargetRevisionOk() (*string, bool)`

GetTargetRevisionOk returns a tuple with the TargetRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetRevision

`func (o *GitOpsApp) SetTargetRevision(v string)`

SetTargetRevision sets TargetRevision field to given value.

### HasTargetRevision

`func (o *GitOpsApp) HasTargetRevision() bool`

HasTargetRevision returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


