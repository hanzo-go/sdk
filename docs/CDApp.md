# CDApp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Automated** | Pointer to **bool** | Automated is whether CD applies git without being asked. It is cd.automated in the values file, rendered by the ApplicationSet&#39;s templatePatch — false means the Application reports drift and nothing moves. | [optional] 
**Health** | Pointer to **string** | Health is the workload&#39;s verdict: Healthy, Progressing, Degraded, Missing. | [optional] 
**Message** | Pointer to **string** | Message is why, when Health is not Healthy. | [optional] 
**Name** | Pointer to **string** | Name is the Application name the generator mints: &lt;namespace&gt;-&lt;app&gt;. It is the join key against a Declaration. | [optional] 
**Namespace** | Pointer to **string** | Namespace is the DESTINATION namespace as the CR declares it — where the workload lands. For a fleet Application that is the org, but this is the OBSERVED field and not our model of it: the two can disagree, and a board whose whole job is drift must be able to show that they do. | [optional] 
**OperationMessage** | Pointer to **string** |  | [optional] 
**Path** | Pointer to **string** | Path is the values file CD renders against, relative to the chart source. | [optional] 
**Phase** | Pointer to **string** | Phase is the last sync operation&#39;s phase (Running, Succeeded, Failed) and OperationMessage is its message. A Failed phase with a Synced verdict is the shape a stuck Application takes. | [optional] 
**Project** | Pointer to **string** | Project is the AppProject fence the sync is admitted under. | [optional] 
**ReconciledAt** | Pointer to **string** | ReconciledAt is when CD last compared this Application. | [optional] 
**Revision** | Pointer to **string** | Revision is the universe commit CD last applied. Empty means it has not applied one — never assume it means main. | [optional] 
**SelfHeal** | Pointer to **bool** | SelfHeal is whether CD also corrects drift the cluster introduced. | [optional] 
**Sync** | Pointer to **string** | Sync is CD&#39;s verdict on git-versus-cluster: Synced, OutOfSync, or Unknown. | [optional] 

## Methods

### NewCDApp

`func NewCDApp() *CDApp`

NewCDApp instantiates a new CDApp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCDAppWithDefaults

`func NewCDAppWithDefaults() *CDApp`

NewCDAppWithDefaults instantiates a new CDApp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutomated

`func (o *CDApp) GetAutomated() bool`

GetAutomated returns the Automated field if non-nil, zero value otherwise.

### GetAutomatedOk

`func (o *CDApp) GetAutomatedOk() (*bool, bool)`

GetAutomatedOk returns a tuple with the Automated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomated

`func (o *CDApp) SetAutomated(v bool)`

SetAutomated sets Automated field to given value.

### HasAutomated

`func (o *CDApp) HasAutomated() bool`

HasAutomated returns a boolean if a field has been set.

### GetHealth

`func (o *CDApp) GetHealth() string`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *CDApp) GetHealthOk() (*string, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *CDApp) SetHealth(v string)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *CDApp) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetMessage

`func (o *CDApp) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CDApp) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CDApp) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *CDApp) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetName

`func (o *CDApp) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CDApp) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CDApp) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CDApp) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *CDApp) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *CDApp) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *CDApp) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *CDApp) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetOperationMessage

`func (o *CDApp) GetOperationMessage() string`

GetOperationMessage returns the OperationMessage field if non-nil, zero value otherwise.

### GetOperationMessageOk

`func (o *CDApp) GetOperationMessageOk() (*string, bool)`

GetOperationMessageOk returns a tuple with the OperationMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationMessage

`func (o *CDApp) SetOperationMessage(v string)`

SetOperationMessage sets OperationMessage field to given value.

### HasOperationMessage

`func (o *CDApp) HasOperationMessage() bool`

HasOperationMessage returns a boolean if a field has been set.

### GetPath

`func (o *CDApp) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *CDApp) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *CDApp) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *CDApp) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetPhase

`func (o *CDApp) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *CDApp) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *CDApp) SetPhase(v string)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *CDApp) HasPhase() bool`

HasPhase returns a boolean if a field has been set.

### GetProject

`func (o *CDApp) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *CDApp) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *CDApp) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *CDApp) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetReconciledAt

`func (o *CDApp) GetReconciledAt() string`

GetReconciledAt returns the ReconciledAt field if non-nil, zero value otherwise.

### GetReconciledAtOk

`func (o *CDApp) GetReconciledAtOk() (*string, bool)`

GetReconciledAtOk returns a tuple with the ReconciledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReconciledAt

`func (o *CDApp) SetReconciledAt(v string)`

SetReconciledAt sets ReconciledAt field to given value.

### HasReconciledAt

`func (o *CDApp) HasReconciledAt() bool`

HasReconciledAt returns a boolean if a field has been set.

### GetRevision

`func (o *CDApp) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *CDApp) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *CDApp) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *CDApp) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetSelfHeal

`func (o *CDApp) GetSelfHeal() bool`

GetSelfHeal returns the SelfHeal field if non-nil, zero value otherwise.

### GetSelfHealOk

`func (o *CDApp) GetSelfHealOk() (*bool, bool)`

GetSelfHealOk returns a tuple with the SelfHeal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfHeal

`func (o *CDApp) SetSelfHeal(v bool)`

SetSelfHeal sets SelfHeal field to given value.

### HasSelfHeal

`func (o *CDApp) HasSelfHeal() bool`

HasSelfHeal returns a boolean if a field has been set.

### GetSync

`func (o *CDApp) GetSync() string`

GetSync returns the Sync field if non-nil, zero value otherwise.

### GetSyncOk

`func (o *CDApp) GetSyncOk() (*string, bool)`

GetSyncOk returns a tuple with the Sync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSync

`func (o *CDApp) SetSync(v string)`

SetSync sets Sync field to given value.

### HasSync

`func (o *CDApp) HasSync() bool`

HasSync returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


