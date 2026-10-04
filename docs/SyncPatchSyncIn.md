# SyncPatchSyncIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actor** | Pointer to **string** | Actor is the loop-guard identity the sync writes as. Omitted, the stored actor stands. | [optional] 
**Direction** | Pointer to **string** | Direction is both, pull, push or off. Omitted, the stored direction stands. | [optional] 
**Id** | Pointer to **string** | ID is the sync to update, from the path. | [optional] 
**Kind** | Pointer to **string** | Kind names a different kind of sync, and is refused, for the same reason. | [optional] 
**Source** | Pointer to [**SyncEndpointReq**](SyncEndpointReq.md) | Source, Target and Kind are DECLARED HERE IN ORDER TO BE REFUSED.  They are immutable by design — re-pointing a sync is a delete and a create, so a link can never silently start syncing somewhere else — but an UNDECLARED field is dropped by the binder before the handler sees it, so a request asking to repoint answered 200, changed nothing, and said nothing. The operator then believes a moved repository has been repointed and it has not.  Live: a sync still naming github.com/hanzoai/cloud after the repository moved to hanzo-inc/cloud failed every reconcile with \&quot;Repository not found\&quot;, and the PATCH that appeared to fix it did nothing at all. Declaring the fields is what lets the documented immutability actually answer. Source names a new upstream, and is refused. Delete this sync and create the one you want. | [optional] 
**Target** | Pointer to [**SyncEndpointReq**](SyncEndpointReq.md) | Target names a new native repository, and is refused, for the same reason. | [optional] 
**Trigger** | Pointer to **string** | Trigger is webhook, poll or manual. Omitted, the stored trigger stands. | [optional] 

## Methods

### NewSyncPatchSyncIn

`func NewSyncPatchSyncIn() *SyncPatchSyncIn`

NewSyncPatchSyncIn instantiates a new SyncPatchSyncIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyncPatchSyncInWithDefaults

`func NewSyncPatchSyncInWithDefaults() *SyncPatchSyncIn`

NewSyncPatchSyncInWithDefaults instantiates a new SyncPatchSyncIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActor

`func (o *SyncPatchSyncIn) GetActor() string`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *SyncPatchSyncIn) GetActorOk() (*string, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *SyncPatchSyncIn) SetActor(v string)`

SetActor sets Actor field to given value.

### HasActor

`func (o *SyncPatchSyncIn) HasActor() bool`

HasActor returns a boolean if a field has been set.

### GetDirection

`func (o *SyncPatchSyncIn) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *SyncPatchSyncIn) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *SyncPatchSyncIn) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *SyncPatchSyncIn) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetId

`func (o *SyncPatchSyncIn) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SyncPatchSyncIn) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SyncPatchSyncIn) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SyncPatchSyncIn) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *SyncPatchSyncIn) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *SyncPatchSyncIn) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *SyncPatchSyncIn) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *SyncPatchSyncIn) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetSource

`func (o *SyncPatchSyncIn) GetSource() SyncEndpointReq`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *SyncPatchSyncIn) GetSourceOk() (*SyncEndpointReq, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *SyncPatchSyncIn) SetSource(v SyncEndpointReq)`

SetSource sets Source field to given value.

### HasSource

`func (o *SyncPatchSyncIn) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetTarget

`func (o *SyncPatchSyncIn) GetTarget() SyncEndpointReq`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *SyncPatchSyncIn) GetTargetOk() (*SyncEndpointReq, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *SyncPatchSyncIn) SetTarget(v SyncEndpointReq)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *SyncPatchSyncIn) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetTrigger

`func (o *SyncPatchSyncIn) GetTrigger() string`

GetTrigger returns the Trigger field if non-nil, zero value otherwise.

### GetTriggerOk

`func (o *SyncPatchSyncIn) GetTriggerOk() (*string, bool)`

GetTriggerOk returns a tuple with the Trigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigger

`func (o *SyncPatchSyncIn) SetTrigger(v string)`

SetTrigger sets Trigger field to given value.

### HasTrigger

`func (o *SyncPatchSyncIn) HasTrigger() bool`

HasTrigger returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


