# ReconcileReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Declared** | Pointer to **int64** | Declared is how many objects the rendered source declares — the denominator the three outcome counts below are read against. Zero means the render produced nothing, which trips the prune fuse rather than sweeping the fleet. | [optional] 
**Failed** | Pointer to **int64** | Failed is how many objects the apply could not reconcile. Non-zero is a PARTIAL run reported at 200: the engine applied what it could and each failure names itself in Results, so a caller reads this number rather than the status code to learn whether the fleet matches the source. | [optional] 
**Instance** | Pointer to **string** | Instance is the tracking id this run stamps on everything it manages, so a later run can tell the objects it owns from objects another instance declares. DEPLOY_ENGINE_INSTANCE names it; the default is &#x60;universe&#x60;. | [optional] 
**Prune** | Pointer to **bool** | Prune reports whether DELETION was enabled for this run. False means an object the source no longer declares was left alone rather than removed, so a zero Pruned below means \&quot;nothing to delete\&quot; only when this is true. | [optional] 
**Pruned** | Pointer to **int64** | Pruned is how many live objects this run DELETED because the source no longer declares them. Always 0 when Prune is false. | [optional] 
**Results** | Pointer to [**[]AppliedResource**](AppliedResource.md) | Results is one entry per object the run acted on, in the order the engine applied them. Empty (never null) when the run reconciled nothing. | [optional] 
**Revision** | Pointer to **string** | Revision is the source commit this run applied, as the source resolved it — a git commit SHA, not an image tag. It is what an operator cites when asking what the cluster was last made to match. | [optional] 
**Source** | Pointer to [**ReconcileSource**](ReconcileSource.md) | Source is the git coordinate the run rendered. It is this deployment&#39;s own configuration echoed back, never a request parameter, and it is reported so a reader of the answer knows WHICH tree the revision names. | [optional] 
**Synced** | Pointer to **int64** | Synced is how many objects the run applied successfully. | [optional] 

## Methods

### NewReconcileReport

`func NewReconcileReport() *ReconcileReport`

NewReconcileReport instantiates a new ReconcileReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReconcileReportWithDefaults

`func NewReconcileReportWithDefaults() *ReconcileReport`

NewReconcileReportWithDefaults instantiates a new ReconcileReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeclared

`func (o *ReconcileReport) GetDeclared() int64`

GetDeclared returns the Declared field if non-nil, zero value otherwise.

### GetDeclaredOk

`func (o *ReconcileReport) GetDeclaredOk() (*int64, bool)`

GetDeclaredOk returns a tuple with the Declared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclared

`func (o *ReconcileReport) SetDeclared(v int64)`

SetDeclared sets Declared field to given value.

### HasDeclared

`func (o *ReconcileReport) HasDeclared() bool`

HasDeclared returns a boolean if a field has been set.

### GetFailed

`func (o *ReconcileReport) GetFailed() int64`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *ReconcileReport) GetFailedOk() (*int64, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *ReconcileReport) SetFailed(v int64)`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *ReconcileReport) HasFailed() bool`

HasFailed returns a boolean if a field has been set.

### GetInstance

`func (o *ReconcileReport) GetInstance() string`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *ReconcileReport) GetInstanceOk() (*string, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *ReconcileReport) SetInstance(v string)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *ReconcileReport) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetPrune

`func (o *ReconcileReport) GetPrune() bool`

GetPrune returns the Prune field if non-nil, zero value otherwise.

### GetPruneOk

`func (o *ReconcileReport) GetPruneOk() (*bool, bool)`

GetPruneOk returns a tuple with the Prune field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrune

`func (o *ReconcileReport) SetPrune(v bool)`

SetPrune sets Prune field to given value.

### HasPrune

`func (o *ReconcileReport) HasPrune() bool`

HasPrune returns a boolean if a field has been set.

### GetPruned

`func (o *ReconcileReport) GetPruned() int64`

GetPruned returns the Pruned field if non-nil, zero value otherwise.

### GetPrunedOk

`func (o *ReconcileReport) GetPrunedOk() (*int64, bool)`

GetPrunedOk returns a tuple with the Pruned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPruned

`func (o *ReconcileReport) SetPruned(v int64)`

SetPruned sets Pruned field to given value.

### HasPruned

`func (o *ReconcileReport) HasPruned() bool`

HasPruned returns a boolean if a field has been set.

### GetResults

`func (o *ReconcileReport) GetResults() []AppliedResource`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ReconcileReport) GetResultsOk() (*[]AppliedResource, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ReconcileReport) SetResults(v []AppliedResource)`

SetResults sets Results field to given value.

### HasResults

`func (o *ReconcileReport) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetRevision

`func (o *ReconcileReport) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *ReconcileReport) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *ReconcileReport) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *ReconcileReport) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetSource

`func (o *ReconcileReport) GetSource() ReconcileSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ReconcileReport) GetSourceOk() (*ReconcileSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ReconcileReport) SetSource(v ReconcileSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *ReconcileReport) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSynced

`func (o *ReconcileReport) GetSynced() int64`

GetSynced returns the Synced field if non-nil, zero value otherwise.

### GetSyncedOk

`func (o *ReconcileReport) GetSyncedOk() (*int64, bool)`

GetSyncedOk returns a tuple with the Synced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynced

`func (o *ReconcileReport) SetSynced(v int64)`

SetSynced sets Synced field to given value.

### HasSynced

`func (o *ReconcileReport) HasSynced() bool`

HasSynced returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


