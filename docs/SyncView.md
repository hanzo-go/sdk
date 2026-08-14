# SyncView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actor** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**Direction** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Source** | Pointer to [**EndpointView**](EndpointView.md) |  | [optional] 
**Target** | Pointer to [**EndpointView**](EndpointView.md) |  | [optional] 
**Trigger** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** | bumped on every reconcile — the last-synced time | [optional] 

## Methods

### NewSyncView

`func NewSyncView() *SyncView`

NewSyncView instantiates a new SyncView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyncViewWithDefaults

`func NewSyncViewWithDefaults() *SyncView`

NewSyncViewWithDefaults instantiates a new SyncView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActor

`func (o *SyncView) GetActor() string`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *SyncView) GetActorOk() (*string, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *SyncView) SetActor(v string)`

SetActor sets Actor field to given value.

### HasActor

`func (o *SyncView) HasActor() bool`

HasActor returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SyncView) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SyncView) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SyncView) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SyncView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDirection

`func (o *SyncView) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *SyncView) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *SyncView) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *SyncView) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetId

`func (o *SyncView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SyncView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SyncView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SyncView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *SyncView) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *SyncView) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *SyncView) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *SyncView) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetSource

`func (o *SyncView) GetSource() EndpointView`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *SyncView) GetSourceOk() (*EndpointView, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *SyncView) SetSource(v EndpointView)`

SetSource sets Source field to given value.

### HasSource

`func (o *SyncView) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetTarget

`func (o *SyncView) GetTarget() EndpointView`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *SyncView) GetTargetOk() (*EndpointView, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *SyncView) SetTarget(v EndpointView)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *SyncView) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetTrigger

`func (o *SyncView) GetTrigger() string`

GetTrigger returns the Trigger field if non-nil, zero value otherwise.

### GetTriggerOk

`func (o *SyncView) GetTriggerOk() (*string, bool)`

GetTriggerOk returns a tuple with the Trigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigger

`func (o *SyncView) SetTrigger(v string)`

SetTrigger sets Trigger field to given value.

### HasTrigger

`func (o *SyncView) HasTrigger() bool`

HasTrigger returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *SyncView) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SyncView) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SyncView) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *SyncView) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


