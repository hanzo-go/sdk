# CloudSyncView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actor** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**Direction** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Source** | Pointer to [**CloudEndpointView**](CloudEndpointView.md) |  | [optional] 
**Target** | Pointer to [**CloudEndpointView**](CloudEndpointView.md) |  | [optional] 
**Trigger** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** | bumped on every reconcile — the last-synced time | [optional] 

## Methods

### NewCloudSyncView

`func NewCloudSyncView() *CloudSyncView`

NewCloudSyncView instantiates a new CloudSyncView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudSyncViewWithDefaults

`func NewCloudSyncViewWithDefaults() *CloudSyncView`

NewCloudSyncViewWithDefaults instantiates a new CloudSyncView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActor

`func (o *CloudSyncView) GetActor() string`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *CloudSyncView) GetActorOk() (*string, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *CloudSyncView) SetActor(v string)`

SetActor sets Actor field to given value.

### HasActor

`func (o *CloudSyncView) HasActor() bool`

HasActor returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CloudSyncView) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CloudSyncView) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CloudSyncView) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CloudSyncView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDirection

`func (o *CloudSyncView) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *CloudSyncView) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *CloudSyncView) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *CloudSyncView) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetId

`func (o *CloudSyncView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudSyncView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudSyncView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudSyncView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *CloudSyncView) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CloudSyncView) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CloudSyncView) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *CloudSyncView) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetSource

`func (o *CloudSyncView) GetSource() CloudEndpointView`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *CloudSyncView) GetSourceOk() (*CloudEndpointView, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *CloudSyncView) SetSource(v CloudEndpointView)`

SetSource sets Source field to given value.

### HasSource

`func (o *CloudSyncView) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetTarget

`func (o *CloudSyncView) GetTarget() CloudEndpointView`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *CloudSyncView) GetTargetOk() (*CloudEndpointView, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *CloudSyncView) SetTarget(v CloudEndpointView)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *CloudSyncView) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetTrigger

`func (o *CloudSyncView) GetTrigger() string`

GetTrigger returns the Trigger field if non-nil, zero value otherwise.

### GetTriggerOk

`func (o *CloudSyncView) GetTriggerOk() (*string, bool)`

GetTriggerOk returns a tuple with the Trigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigger

`func (o *CloudSyncView) SetTrigger(v string)`

SetTrigger sets Trigger field to given value.

### HasTrigger

`func (o *CloudSyncView) HasTrigger() bool`

HasTrigger returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CloudSyncView) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CloudSyncView) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CloudSyncView) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CloudSyncView) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


