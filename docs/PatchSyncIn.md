# PatchSyncIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actor** | Pointer to **string** | Actor is the loop-guard identity the sync writes as. Omitted, the stored actor stands. | [optional] 
**Direction** | Pointer to **string** | Direction is both, pull, push or off. Omitted, the stored direction stands. | [optional] 
**Id** | Pointer to **string** | ID is the sync to update, from the path. | [optional] 
**Kind** | Pointer to **string** | Kind names a different kind of sync, and is refused, for the same reason. | [optional] 
**Source** | Pointer to [**EndpointReq**](EndpointReq.md) | Source, Target and Kind are DECLARED HERE IN ORDER TO BE REFUSED.  They are immutable by design — re-pointing a sync is a delete and a create, so a link can never silently start syncing somewhere else — but an UNDECLARED field is dropped by the binder before the handler sees it, so a request asking to repoint answered 200, changed nothing, and said nothing. The operator then believes a moved repository has been repointed and it has not.  Live: a sync still naming github.com/hanzoai/cloud after the repository moved to hanzo-inc/cloud failed every reconcile with \&quot;Repository not found\&quot;, and the PATCH that appeared to fix it did nothing at all. Declaring the fields is what lets the documented immutability actually answer. Source names a new upstream, and is refused. Delete this sync and create the one you want. | [optional] 
**Target** | Pointer to [**EndpointReq**](EndpointReq.md) | Target names a new native repository, and is refused, for the same reason. | [optional] 
**Trigger** | Pointer to **string** | Trigger is webhook, poll or manual. Omitted, the stored trigger stands. | [optional] 

## Methods

### NewPatchSyncIn

`func NewPatchSyncIn() *PatchSyncIn`

NewPatchSyncIn instantiates a new PatchSyncIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchSyncInWithDefaults

`func NewPatchSyncInWithDefaults() *PatchSyncIn`

NewPatchSyncInWithDefaults instantiates a new PatchSyncIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActor

`func (o *PatchSyncIn) GetActor() string`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *PatchSyncIn) GetActorOk() (*string, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *PatchSyncIn) SetActor(v string)`

SetActor sets Actor field to given value.

### HasActor

`func (o *PatchSyncIn) HasActor() bool`

HasActor returns a boolean if a field has been set.

### GetDirection

`func (o *PatchSyncIn) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *PatchSyncIn) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *PatchSyncIn) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *PatchSyncIn) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetId

`func (o *PatchSyncIn) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchSyncIn) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchSyncIn) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchSyncIn) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *PatchSyncIn) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *PatchSyncIn) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *PatchSyncIn) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *PatchSyncIn) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetSource

`func (o *PatchSyncIn) GetSource() EndpointReq`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *PatchSyncIn) GetSourceOk() (*EndpointReq, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *PatchSyncIn) SetSource(v EndpointReq)`

SetSource sets Source field to given value.

### HasSource

`func (o *PatchSyncIn) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetTarget

`func (o *PatchSyncIn) GetTarget() EndpointReq`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *PatchSyncIn) GetTargetOk() (*EndpointReq, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *PatchSyncIn) SetTarget(v EndpointReq)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *PatchSyncIn) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetTrigger

`func (o *PatchSyncIn) GetTrigger() string`

GetTrigger returns the Trigger field if non-nil, zero value otherwise.

### GetTriggerOk

`func (o *PatchSyncIn) GetTriggerOk() (*string, bool)`

GetTriggerOk returns a tuple with the Trigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigger

`func (o *PatchSyncIn) SetTrigger(v string)`

SetTrigger sets Trigger field to given value.

### HasTrigger

`func (o *PatchSyncIn) HasTrigger() bool`

HasTrigger returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


