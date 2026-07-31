# CloudSyncQueued

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID is the sync the reconcile was queued for. | [optional] 
**Queued** | Pointer to **bool** | Queued is true when the reconcile was accepted; it has not run yet. | [optional] 

## Methods

### NewCloudSyncQueued

`func NewCloudSyncQueued() *CloudSyncQueued`

NewCloudSyncQueued instantiates a new CloudSyncQueued object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudSyncQueuedWithDefaults

`func NewCloudSyncQueuedWithDefaults() *CloudSyncQueued`

NewCloudSyncQueuedWithDefaults instantiates a new CloudSyncQueued object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CloudSyncQueued) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudSyncQueued) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudSyncQueued) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudSyncQueued) HasId() bool`

HasId returns a boolean if a field has been set.

### GetQueued

`func (o *CloudSyncQueued) GetQueued() bool`

GetQueued returns the Queued field if non-nil, zero value otherwise.

### GetQueuedOk

`func (o *CloudSyncQueued) GetQueuedOk() (*bool, bool)`

GetQueuedOk returns a tuple with the Queued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueued

`func (o *CloudSyncQueued) SetQueued(v bool)`

SetQueued sets Queued field to given value.

### HasQueued

`func (o *CloudSyncQueued) HasQueued() bool`

HasQueued returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


