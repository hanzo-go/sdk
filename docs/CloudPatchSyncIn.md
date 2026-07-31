# CloudPatchSyncIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actor** | Pointer to **string** | Actor is the loop-guard identity the sync writes as. Omitted, the stored actor stands. | [optional] 
**Direction** | Pointer to **string** | Direction is both, pull, push or off. Omitted, the stored direction stands. | [optional] 
**Id** | Pointer to **string** | ID is the sync to update, from the path. | [optional] 
**Trigger** | Pointer to **string** | Trigger is webhook, poll or manual. Omitted, the stored trigger stands. | [optional] 

## Methods

### NewCloudPatchSyncIn

`func NewCloudPatchSyncIn() *CloudPatchSyncIn`

NewCloudPatchSyncIn instantiates a new CloudPatchSyncIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudPatchSyncInWithDefaults

`func NewCloudPatchSyncInWithDefaults() *CloudPatchSyncIn`

NewCloudPatchSyncInWithDefaults instantiates a new CloudPatchSyncIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActor

`func (o *CloudPatchSyncIn) GetActor() string`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *CloudPatchSyncIn) GetActorOk() (*string, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *CloudPatchSyncIn) SetActor(v string)`

SetActor sets Actor field to given value.

### HasActor

`func (o *CloudPatchSyncIn) HasActor() bool`

HasActor returns a boolean if a field has been set.

### GetDirection

`func (o *CloudPatchSyncIn) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *CloudPatchSyncIn) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *CloudPatchSyncIn) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *CloudPatchSyncIn) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetId

`func (o *CloudPatchSyncIn) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudPatchSyncIn) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudPatchSyncIn) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudPatchSyncIn) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTrigger

`func (o *CloudPatchSyncIn) GetTrigger() string`

GetTrigger returns the Trigger field if non-nil, zero value otherwise.

### GetTriggerOk

`func (o *CloudPatchSyncIn) GetTriggerOk() (*string, bool)`

GetTriggerOk returns a tuple with the Trigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigger

`func (o *CloudPatchSyncIn) SetTrigger(v string)`

SetTrigger sets Trigger field to given value.

### HasTrigger

`func (o *CloudPatchSyncIn) HasTrigger() bool`

HasTrigger returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


