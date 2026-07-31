# CloudGitOpsApp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Automated** | Pointer to **bool** |  | [optional] 
**Health** | Pointer to **string** | Healthy|Degraded|Progressing|… | [optional] 
**History** | Pointer to [**[]CloudGitOpsDeploy**](CloudGitOpsDeploy.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**Operation** | Pointer to [**CloudGitOpsOperation**](CloudGitOpsOperation.md) |  | [optional] 
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

### NewCloudGitOpsApp

`func NewCloudGitOpsApp() *CloudGitOpsApp`

NewCloudGitOpsApp instantiates a new CloudGitOpsApp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudGitOpsAppWithDefaults

`func NewCloudGitOpsAppWithDefaults() *CloudGitOpsApp`

NewCloudGitOpsAppWithDefaults instantiates a new CloudGitOpsApp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutomated

`func (o *CloudGitOpsApp) GetAutomated() bool`

GetAutomated returns the Automated field if non-nil, zero value otherwise.

### GetAutomatedOk

`func (o *CloudGitOpsApp) GetAutomatedOk() (*bool, bool)`

GetAutomatedOk returns a tuple with the Automated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomated

`func (o *CloudGitOpsApp) SetAutomated(v bool)`

SetAutomated sets Automated field to given value.

### HasAutomated

`func (o *CloudGitOpsApp) HasAutomated() bool`

HasAutomated returns a boolean if a field has been set.

### GetHealth

`func (o *CloudGitOpsApp) GetHealth() string`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *CloudGitOpsApp) GetHealthOk() (*string, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *CloudGitOpsApp) SetHealth(v string)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *CloudGitOpsApp) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetHistory

`func (o *CloudGitOpsApp) GetHistory() []CloudGitOpsDeploy`

GetHistory returns the History field if non-nil, zero value otherwise.

### GetHistoryOk

`func (o *CloudGitOpsApp) GetHistoryOk() (*[]CloudGitOpsDeploy, bool)`

GetHistoryOk returns a tuple with the History field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistory

`func (o *CloudGitOpsApp) SetHistory(v []CloudGitOpsDeploy)`

SetHistory sets History field to given value.

### HasHistory

`func (o *CloudGitOpsApp) HasHistory() bool`

HasHistory returns a boolean if a field has been set.

### GetName

`func (o *CloudGitOpsApp) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudGitOpsApp) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudGitOpsApp) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudGitOpsApp) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *CloudGitOpsApp) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *CloudGitOpsApp) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *CloudGitOpsApp) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *CloudGitOpsApp) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetOperation

`func (o *CloudGitOpsApp) GetOperation() CloudGitOpsOperation`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *CloudGitOpsApp) GetOperationOk() (*CloudGitOpsOperation, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *CloudGitOpsApp) SetOperation(v CloudGitOpsOperation)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *CloudGitOpsApp) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetPath

`func (o *CloudGitOpsApp) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *CloudGitOpsApp) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *CloudGitOpsApp) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *CloudGitOpsApp) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetProject

`func (o *CloudGitOpsApp) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *CloudGitOpsApp) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *CloudGitOpsApp) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *CloudGitOpsApp) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetReconciledAt

`func (o *CloudGitOpsApp) GetReconciledAt() string`

GetReconciledAt returns the ReconciledAt field if non-nil, zero value otherwise.

### GetReconciledAtOk

`func (o *CloudGitOpsApp) GetReconciledAtOk() (*string, bool)`

GetReconciledAtOk returns a tuple with the ReconciledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReconciledAt

`func (o *CloudGitOpsApp) SetReconciledAt(v string)`

SetReconciledAt sets ReconciledAt field to given value.

### HasReconciledAt

`func (o *CloudGitOpsApp) HasReconciledAt() bool`

HasReconciledAt returns a boolean if a field has been set.

### GetRepoURL

`func (o *CloudGitOpsApp) GetRepoURL() string`

GetRepoURL returns the RepoURL field if non-nil, zero value otherwise.

### GetRepoURLOk

`func (o *CloudGitOpsApp) GetRepoURLOk() (*string, bool)`

GetRepoURLOk returns a tuple with the RepoURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoURL

`func (o *CloudGitOpsApp) SetRepoURL(v string)`

SetRepoURL sets RepoURL field to given value.

### HasRepoURL

`func (o *CloudGitOpsApp) HasRepoURL() bool`

HasRepoURL returns a boolean if a field has been set.

### GetResources

`func (o *CloudGitOpsApp) GetResources() int32`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *CloudGitOpsApp) GetResourcesOk() (*int32, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *CloudGitOpsApp) SetResources(v int32)`

SetResources sets Resources field to given value.

### HasResources

`func (o *CloudGitOpsApp) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetRevision

`func (o *CloudGitOpsApp) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *CloudGitOpsApp) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *CloudGitOpsApp) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *CloudGitOpsApp) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetSelfHeal

`func (o *CloudGitOpsApp) GetSelfHeal() bool`

GetSelfHeal returns the SelfHeal field if non-nil, zero value otherwise.

### GetSelfHealOk

`func (o *CloudGitOpsApp) GetSelfHealOk() (*bool, bool)`

GetSelfHealOk returns a tuple with the SelfHeal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfHeal

`func (o *CloudGitOpsApp) SetSelfHeal(v bool)`

SetSelfHeal sets SelfHeal field to given value.

### HasSelfHeal

`func (o *CloudGitOpsApp) HasSelfHeal() bool`

HasSelfHeal returns a boolean if a field has been set.

### GetSync

`func (o *CloudGitOpsApp) GetSync() string`

GetSync returns the Sync field if non-nil, zero value otherwise.

### GetSyncOk

`func (o *CloudGitOpsApp) GetSyncOk() (*string, bool)`

GetSyncOk returns a tuple with the Sync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSync

`func (o *CloudGitOpsApp) SetSync(v string)`

SetSync sets Sync field to given value.

### HasSync

`func (o *CloudGitOpsApp) HasSync() bool`

HasSync returns a boolean if a field has been set.

### GetTargetRevision

`func (o *CloudGitOpsApp) GetTargetRevision() string`

GetTargetRevision returns the TargetRevision field if non-nil, zero value otherwise.

### GetTargetRevisionOk

`func (o *CloudGitOpsApp) GetTargetRevisionOk() (*string, bool)`

GetTargetRevisionOk returns a tuple with the TargetRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetRevision

`func (o *CloudGitOpsApp) SetTargetRevision(v string)`

SetTargetRevision sets TargetRevision field to given value.

### HasTargetRevision

`func (o *CloudGitOpsApp) HasTargetRevision() bool`

HasTargetRevision returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


