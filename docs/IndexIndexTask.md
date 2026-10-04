# IndexIndexTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnqueuedAt** | Pointer to **string** | EnqueuedAt, StartedAt and FinishedAt are the same instant: the write was applied before its task id was minted. | [optional] 
**FinishedAt** | Pointer to **string** | FinishedAt is when the write completed. | [optional] 
**StartedAt** | Pointer to **string** | StartedAt is when the write began. | [optional] 
**Status** | Pointer to **string** | Status is always &#x60;succeeded&#x60;. | [optional] 
**Type** | Pointer to **string** | Type names the kind of write, for a client that inspects it. | [optional] 
**Uid** | Pointer to **int64** | UID echoes the task id that was asked about. | [optional] 

## Methods

### NewIndexIndexTask

`func NewIndexIndexTask() *IndexIndexTask`

NewIndexIndexTask instantiates a new IndexIndexTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndexIndexTaskWithDefaults

`func NewIndexIndexTaskWithDefaults() *IndexIndexTask`

NewIndexIndexTaskWithDefaults instantiates a new IndexIndexTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnqueuedAt

`func (o *IndexIndexTask) GetEnqueuedAt() string`

GetEnqueuedAt returns the EnqueuedAt field if non-nil, zero value otherwise.

### GetEnqueuedAtOk

`func (o *IndexIndexTask) GetEnqueuedAtOk() (*string, bool)`

GetEnqueuedAtOk returns a tuple with the EnqueuedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnqueuedAt

`func (o *IndexIndexTask) SetEnqueuedAt(v string)`

SetEnqueuedAt sets EnqueuedAt field to given value.

### HasEnqueuedAt

`func (o *IndexIndexTask) HasEnqueuedAt() bool`

HasEnqueuedAt returns a boolean if a field has been set.

### GetFinishedAt

`func (o *IndexIndexTask) GetFinishedAt() string`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *IndexIndexTask) GetFinishedAtOk() (*string, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *IndexIndexTask) SetFinishedAt(v string)`

SetFinishedAt sets FinishedAt field to given value.

### HasFinishedAt

`func (o *IndexIndexTask) HasFinishedAt() bool`

HasFinishedAt returns a boolean if a field has been set.

### GetStartedAt

`func (o *IndexIndexTask) GetStartedAt() string`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *IndexIndexTask) GetStartedAtOk() (*string, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *IndexIndexTask) SetStartedAt(v string)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *IndexIndexTask) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetStatus

`func (o *IndexIndexTask) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IndexIndexTask) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IndexIndexTask) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *IndexIndexTask) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *IndexIndexTask) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IndexIndexTask) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IndexIndexTask) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IndexIndexTask) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUid

`func (o *IndexIndexTask) GetUid() int64`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *IndexIndexTask) GetUidOk() (*int64, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *IndexIndexTask) SetUid(v int64)`

SetUid sets Uid field to given value.

### HasUid

`func (o *IndexIndexTask) HasUid() bool`

HasUid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


