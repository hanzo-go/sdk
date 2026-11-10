# DeployRunning

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Health** | Pointer to **string** | Health is CD&#39;s verdict on the objects it manages: Healthy, Progressing, Degraded, Suspended, Missing or Unknown. | [optional] 
**Images** | Pointer to [**[]DeployImage**](DeployImage.md) | Images are the images the service&#39;s pods are running now, as CD collected them from the live pods. Two of one repository is a rollout in progress. | [optional] 
**Name** | Pointer to **string** | Name is the CD Application&#39;s name, which the fleet mints as &lt;namespace&gt;-&lt;app&gt;. | [optional] 
**Namespace** | Pointer to **string** | Namespace is where the service&#39;s workloads run (spec.destination.namespace). | [optional] 
**RepoURL** | Pointer to **string** | RepoURL is the git repository CD applies this service from, usually universe. | [optional] 
**Revision** | Pointer to **string** | Revision is the commit of that repository CD last applied. Empty means it has applied none. | [optional] 
**Sync** | Pointer to **string** | Sync is CD&#39;s verdict on git against the cluster: Synced, OutOfSync or Unknown. | [optional] 
**SyncedAt** | Pointer to **string** | SyncedAt is when CD last finished applying a revision, RFC 3339. Empty when it never has. | [optional] 

## Methods

### NewDeployRunning

`func NewDeployRunning() *DeployRunning`

NewDeployRunning instantiates a new DeployRunning object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeployRunningWithDefaults

`func NewDeployRunningWithDefaults() *DeployRunning`

NewDeployRunningWithDefaults instantiates a new DeployRunning object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHealth

`func (o *DeployRunning) GetHealth() string`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *DeployRunning) GetHealthOk() (*string, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *DeployRunning) SetHealth(v string)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *DeployRunning) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetImages

`func (o *DeployRunning) GetImages() []DeployImage`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *DeployRunning) GetImagesOk() (*[]DeployImage, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *DeployRunning) SetImages(v []DeployImage)`

SetImages sets Images field to given value.

### HasImages

`func (o *DeployRunning) HasImages() bool`

HasImages returns a boolean if a field has been set.

### GetName

`func (o *DeployRunning) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeployRunning) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeployRunning) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DeployRunning) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *DeployRunning) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *DeployRunning) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *DeployRunning) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *DeployRunning) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetRepoURL

`func (o *DeployRunning) GetRepoURL() string`

GetRepoURL returns the RepoURL field if non-nil, zero value otherwise.

### GetRepoURLOk

`func (o *DeployRunning) GetRepoURLOk() (*string, bool)`

GetRepoURLOk returns a tuple with the RepoURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoURL

`func (o *DeployRunning) SetRepoURL(v string)`

SetRepoURL sets RepoURL field to given value.

### HasRepoURL

`func (o *DeployRunning) HasRepoURL() bool`

HasRepoURL returns a boolean if a field has been set.

### GetRevision

`func (o *DeployRunning) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *DeployRunning) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *DeployRunning) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *DeployRunning) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetSync

`func (o *DeployRunning) GetSync() string`

GetSync returns the Sync field if non-nil, zero value otherwise.

### GetSyncOk

`func (o *DeployRunning) GetSyncOk() (*string, bool)`

GetSyncOk returns a tuple with the Sync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSync

`func (o *DeployRunning) SetSync(v string)`

SetSync sets Sync field to given value.

### HasSync

`func (o *DeployRunning) HasSync() bool`

HasSync returns a boolean if a field has been set.

### GetSyncedAt

`func (o *DeployRunning) GetSyncedAt() string`

GetSyncedAt returns the SyncedAt field if non-nil, zero value otherwise.

### GetSyncedAtOk

`func (o *DeployRunning) GetSyncedAtOk() (*string, bool)`

GetSyncedAtOk returns a tuple with the SyncedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncedAt

`func (o *DeployRunning) SetSyncedAt(v string)`

SetSyncedAt sets SyncedAt field to given value.

### HasSyncedAt

`func (o *DeployRunning) HasSyncedAt() bool`

HasSyncedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


