# SyncSyncQueued

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID is the sync the reconcile was queued for. | [optional] 
**Queued** | Pointer to **bool** | Queued is true when the reconcile was accepted; it has not run yet. | [optional] 

## Methods

### NewSyncSyncQueued

`func NewSyncSyncQueued() *SyncSyncQueued`

NewSyncSyncQueued instantiates a new SyncSyncQueued object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyncSyncQueuedWithDefaults

`func NewSyncSyncQueuedWithDefaults() *SyncSyncQueued`

NewSyncSyncQueuedWithDefaults instantiates a new SyncSyncQueued object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SyncSyncQueued) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SyncSyncQueued) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SyncSyncQueued) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SyncSyncQueued) HasId() bool`

HasId returns a boolean if a field has been set.

### GetQueued

`func (o *SyncSyncQueued) GetQueued() bool`

GetQueued returns the Queued field if non-nil, zero value otherwise.

### GetQueuedOk

`func (o *SyncSyncQueued) GetQueuedOk() (*bool, bool)`

GetQueuedOk returns a tuple with the Queued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueued

`func (o *SyncSyncQueued) SetQueued(v bool)`

SetQueued sets Queued field to given value.

### HasQueued

`func (o *SyncSyncQueued) HasQueued() bool`

HasQueued returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


