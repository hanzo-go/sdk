# PlatformCDApp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Applied** | Pointer to **string** | Applied is the universe commit CD last synced, from its deploy history. It trails Revision exactly while a sync is pending. | [optional] 
**Automated** | Pointer to **bool** | Automated is whether CD applies git without being asked. It is cd.automated in the values file, rendered by the ApplicationSet&#39;s templatePatch — false means the Application reports drift and nothing moves. | [optional] 
**Cluster** | Pointer to **string** | Cluster is where the Application deploys: its destination cluster&#39;s name, else its server URL (https://kubernetes.default.svc is CD&#39;s own cluster). | [optional] 
**Health** | Pointer to **string** | Health is the workload&#39;s verdict: Healthy, Progressing, Degraded, Missing. | [optional] 
**Message** | Pointer to **string** | Message is why, when Health is not Healthy. | [optional] 
**Name** | Pointer to **string** | Name is the Application name the generator mints: &lt;namespace&gt;-&lt;app&gt;. It is the join key against a Declaration. | [optional] 
**Namespace** | Pointer to **string** | Namespace is the DESTINATION namespace as the CR declares it — where the workload lands. For a fleet Application that is the org, but this is the OBSERVED field and not our model of it: the two can disagree, and a board whose whole job is drift must be able to show that they do. | [optional] 
**OperationMessage** | Pointer to **string** |  | [optional] 
**Path** | Pointer to **string** | Path is the values file CD renders against, relative to the chart source. | [optional] 
**Phase** | Pointer to **string** | Phase is the last sync operation&#39;s phase (Running, Succeeded, Failed) and OperationMessage is its message. A Failed phase with a Synced verdict is the shape a stuck Application takes. | [optional] 
**Project** | Pointer to **string** | Project is the AppProject fence the sync is admitted under. | [optional] 
**ReconciledAt** | Pointer to **string** | ReconciledAt is when CD last compared this Application. | [optional] 
**Revision** | Pointer to **string** | Revision is the universe commit CD last compared the cluster against — what a sync applies. Empty means it has compared none; never assume it means main. | [optional] 
**SelfHeal** | Pointer to **bool** | SelfHeal is whether CD also corrects drift the cluster introduced. | [optional] 
**Sync** | Pointer to **string** | Sync is CD&#39;s verdict on git-versus-cluster: Synced, OutOfSync, or Unknown. | [optional] 

## Methods

### NewPlatformCDApp

`func NewPlatformCDApp() *PlatformCDApp`

NewPlatformCDApp instantiates a new PlatformCDApp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformCDAppWithDefaults

`func NewPlatformCDAppWithDefaults() *PlatformCDApp`

NewPlatformCDAppWithDefaults instantiates a new PlatformCDApp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplied

`func (o *PlatformCDApp) GetApplied() string`

GetApplied returns the Applied field if non-nil, zero value otherwise.

### GetAppliedOk

`func (o *PlatformCDApp) GetAppliedOk() (*string, bool)`

GetAppliedOk returns a tuple with the Applied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplied

`func (o *PlatformCDApp) SetApplied(v string)`

SetApplied sets Applied field to given value.

### HasApplied

`func (o *PlatformCDApp) HasApplied() bool`

HasApplied returns a boolean if a field has been set.

### GetAutomated

`func (o *PlatformCDApp) GetAutomated() bool`

GetAutomated returns the Automated field if non-nil, zero value otherwise.

### GetAutomatedOk

`func (o *PlatformCDApp) GetAutomatedOk() (*bool, bool)`

GetAutomatedOk returns a tuple with the Automated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomated

`func (o *PlatformCDApp) SetAutomated(v bool)`

SetAutomated sets Automated field to given value.

### HasAutomated

`func (o *PlatformCDApp) HasAutomated() bool`

HasAutomated returns a boolean if a field has been set.

### GetCluster

`func (o *PlatformCDApp) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *PlatformCDApp) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *PlatformCDApp) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *PlatformCDApp) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetHealth

`func (o *PlatformCDApp) GetHealth() string`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *PlatformCDApp) GetHealthOk() (*string, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *PlatformCDApp) SetHealth(v string)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *PlatformCDApp) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetMessage

`func (o *PlatformCDApp) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PlatformCDApp) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PlatformCDApp) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *PlatformCDApp) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetName

`func (o *PlatformCDApp) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PlatformCDApp) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PlatformCDApp) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PlatformCDApp) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *PlatformCDApp) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *PlatformCDApp) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *PlatformCDApp) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *PlatformCDApp) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetOperationMessage

`func (o *PlatformCDApp) GetOperationMessage() string`

GetOperationMessage returns the OperationMessage field if non-nil, zero value otherwise.

### GetOperationMessageOk

`func (o *PlatformCDApp) GetOperationMessageOk() (*string, bool)`

GetOperationMessageOk returns a tuple with the OperationMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationMessage

`func (o *PlatformCDApp) SetOperationMessage(v string)`

SetOperationMessage sets OperationMessage field to given value.

### HasOperationMessage

`func (o *PlatformCDApp) HasOperationMessage() bool`

HasOperationMessage returns a boolean if a field has been set.

### GetPath

`func (o *PlatformCDApp) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *PlatformCDApp) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *PlatformCDApp) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *PlatformCDApp) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetPhase

`func (o *PlatformCDApp) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *PlatformCDApp) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *PlatformCDApp) SetPhase(v string)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *PlatformCDApp) HasPhase() bool`

HasPhase returns a boolean if a field has been set.

### GetProject

`func (o *PlatformCDApp) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *PlatformCDApp) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *PlatformCDApp) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *PlatformCDApp) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetReconciledAt

`func (o *PlatformCDApp) GetReconciledAt() string`

GetReconciledAt returns the ReconciledAt field if non-nil, zero value otherwise.

### GetReconciledAtOk

`func (o *PlatformCDApp) GetReconciledAtOk() (*string, bool)`

GetReconciledAtOk returns a tuple with the ReconciledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReconciledAt

`func (o *PlatformCDApp) SetReconciledAt(v string)`

SetReconciledAt sets ReconciledAt field to given value.

### HasReconciledAt

`func (o *PlatformCDApp) HasReconciledAt() bool`

HasReconciledAt returns a boolean if a field has been set.

### GetRevision

`func (o *PlatformCDApp) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *PlatformCDApp) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *PlatformCDApp) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *PlatformCDApp) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetSelfHeal

`func (o *PlatformCDApp) GetSelfHeal() bool`

GetSelfHeal returns the SelfHeal field if non-nil, zero value otherwise.

### GetSelfHealOk

`func (o *PlatformCDApp) GetSelfHealOk() (*bool, bool)`

GetSelfHealOk returns a tuple with the SelfHeal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfHeal

`func (o *PlatformCDApp) SetSelfHeal(v bool)`

SetSelfHeal sets SelfHeal field to given value.

### HasSelfHeal

`func (o *PlatformCDApp) HasSelfHeal() bool`

HasSelfHeal returns a boolean if a field has been set.

### GetSync

`func (o *PlatformCDApp) GetSync() string`

GetSync returns the Sync field if non-nil, zero value otherwise.

### GetSyncOk

`func (o *PlatformCDApp) GetSyncOk() (*string, bool)`

GetSyncOk returns a tuple with the Sync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSync

`func (o *PlatformCDApp) SetSync(v string)`

SetSync sets Sync field to given value.

### HasSync

`func (o *PlatformCDApp) HasSync() bool`

HasSync returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


