# IndexEnqueued

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnqueuedAt** | Pointer to **string** | EnqueuedAt is when the task was recorded, RFC 3339 — which is also when it completed. | [optional] 
**IndexUid** | Pointer to **string** | IndexUID names the index the write landed in. | [optional] 
**Status** | Pointer to **string** | Status is always &#x60;enqueued&#x60;, for dialect compatibility. The work is already done. | [optional] 
**TaskUid** | Pointer to **int32** | TaskUID identifies the task for a client that polls it. Polling resolves immediately. | [optional] 
**Type** | Pointer to **string** | Type is the dialect&#39;s name for the kind of write: indexCreation, indexDeletion, settingsUpdate, documentAdditionOrUpdate, documentDeletion. | [optional] 

## Methods

### NewIndexEnqueued

`func NewIndexEnqueued() *IndexEnqueued`

NewIndexEnqueued instantiates a new IndexEnqueued object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndexEnqueuedWithDefaults

`func NewIndexEnqueuedWithDefaults() *IndexEnqueued`

NewIndexEnqueuedWithDefaults instantiates a new IndexEnqueued object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnqueuedAt

`func (o *IndexEnqueued) GetEnqueuedAt() string`

GetEnqueuedAt returns the EnqueuedAt field if non-nil, zero value otherwise.

### GetEnqueuedAtOk

`func (o *IndexEnqueued) GetEnqueuedAtOk() (*string, bool)`

GetEnqueuedAtOk returns a tuple with the EnqueuedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnqueuedAt

`func (o *IndexEnqueued) SetEnqueuedAt(v string)`

SetEnqueuedAt sets EnqueuedAt field to given value.

### HasEnqueuedAt

`func (o *IndexEnqueued) HasEnqueuedAt() bool`

HasEnqueuedAt returns a boolean if a field has been set.

### GetIndexUid

`func (o *IndexEnqueued) GetIndexUid() string`

GetIndexUid returns the IndexUid field if non-nil, zero value otherwise.

### GetIndexUidOk

`func (o *IndexEnqueued) GetIndexUidOk() (*string, bool)`

GetIndexUidOk returns a tuple with the IndexUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexUid

`func (o *IndexEnqueued) SetIndexUid(v string)`

SetIndexUid sets IndexUid field to given value.

### HasIndexUid

`func (o *IndexEnqueued) HasIndexUid() bool`

HasIndexUid returns a boolean if a field has been set.

### GetStatus

`func (o *IndexEnqueued) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IndexEnqueued) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IndexEnqueued) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *IndexEnqueued) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTaskUid

`func (o *IndexEnqueued) GetTaskUid() int32`

GetTaskUid returns the TaskUid field if non-nil, zero value otherwise.

### GetTaskUidOk

`func (o *IndexEnqueued) GetTaskUidOk() (*int32, bool)`

GetTaskUidOk returns a tuple with the TaskUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskUid

`func (o *IndexEnqueued) SetTaskUid(v int32)`

SetTaskUid sets TaskUid field to given value.

### HasTaskUid

`func (o *IndexEnqueued) HasTaskUid() bool`

HasTaskUid returns a boolean if a field has been set.

### GetType

`func (o *IndexEnqueued) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IndexEnqueued) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IndexEnqueued) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IndexEnqueued) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


