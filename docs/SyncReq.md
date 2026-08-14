# SyncReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actor** | Pointer to **string** | Actor is the identity the sync writes as, used as the loop guard so its own writes do not re-trigger it. Defaults to the deployment&#39;s GIT_SYNC_ACTOR. | [optional] 
**Direction** | Pointer to **string** | Direction is both (the default), pull, push or off. | [optional] 
**Kind** | Pointer to **string** | Kind is what is being synced. Only \&quot;git\&quot; today, which is also the default. | [optional] 
**Run** | Pointer to **bool** | Run reconciles once immediately after the upsert, in the background. | [optional] 
**Source** | Pointer to [**EndpointReq**](EndpointReq.md) | Source is the upstream end. Required. | [optional] 
**Target** | Pointer to [**EndpointReq**](EndpointReq.md) | Target is the downstream end. Optional for git: a native repository named after the source is derived when it is omitted. | [optional] 
**Trigger** | Pointer to **string** | Trigger is what starts a reconcile: webhook (the default), poll or manual. | [optional] 

## Methods

### NewSyncReq

`func NewSyncReq() *SyncReq`

NewSyncReq instantiates a new SyncReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyncReqWithDefaults

`func NewSyncReqWithDefaults() *SyncReq`

NewSyncReqWithDefaults instantiates a new SyncReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActor

`func (o *SyncReq) GetActor() string`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *SyncReq) GetActorOk() (*string, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *SyncReq) SetActor(v string)`

SetActor sets Actor field to given value.

### HasActor

`func (o *SyncReq) HasActor() bool`

HasActor returns a boolean if a field has been set.

### GetDirection

`func (o *SyncReq) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *SyncReq) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *SyncReq) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *SyncReq) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetKind

`func (o *SyncReq) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *SyncReq) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *SyncReq) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *SyncReq) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetRun

`func (o *SyncReq) GetRun() bool`

GetRun returns the Run field if non-nil, zero value otherwise.

### GetRunOk

`func (o *SyncReq) GetRunOk() (*bool, bool)`

GetRunOk returns a tuple with the Run field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRun

`func (o *SyncReq) SetRun(v bool)`

SetRun sets Run field to given value.

### HasRun

`func (o *SyncReq) HasRun() bool`

HasRun returns a boolean if a field has been set.

### GetSource

`func (o *SyncReq) GetSource() EndpointReq`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *SyncReq) GetSourceOk() (*EndpointReq, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *SyncReq) SetSource(v EndpointReq)`

SetSource sets Source field to given value.

### HasSource

`func (o *SyncReq) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetTarget

`func (o *SyncReq) GetTarget() EndpointReq`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *SyncReq) GetTargetOk() (*EndpointReq, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *SyncReq) SetTarget(v EndpointReq)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *SyncReq) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetTrigger

`func (o *SyncReq) GetTrigger() string`

GetTrigger returns the Trigger field if non-nil, zero value otherwise.

### GetTriggerOk

`func (o *SyncReq) GetTriggerOk() (*string, bool)`

GetTriggerOk returns a tuple with the Trigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigger

`func (o *SyncReq) SetTrigger(v string)`

SetTrigger sets Trigger field to given value.

### HasTrigger

`func (o *SyncReq) HasTrigger() bool`

HasTrigger returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


