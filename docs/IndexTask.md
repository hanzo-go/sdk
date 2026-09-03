# IndexTask

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

### NewIndexTask

`func NewIndexTask() *IndexTask`

NewIndexTask instantiates a new IndexTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndexTaskWithDefaults

`func NewIndexTaskWithDefaults() *IndexTask`

NewIndexTaskWithDefaults instantiates a new IndexTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnqueuedAt

`func (o *IndexTask) GetEnqueuedAt() string`

GetEnqueuedAt returns the EnqueuedAt field if non-nil, zero value otherwise.

### GetEnqueuedAtOk

`func (o *IndexTask) GetEnqueuedAtOk() (*string, bool)`

GetEnqueuedAtOk returns a tuple with the EnqueuedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnqueuedAt

`func (o *IndexTask) SetEnqueuedAt(v string)`

SetEnqueuedAt sets EnqueuedAt field to given value.

### HasEnqueuedAt

`func (o *IndexTask) HasEnqueuedAt() bool`

HasEnqueuedAt returns a boolean if a field has been set.

### GetFinishedAt

`func (o *IndexTask) GetFinishedAt() string`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *IndexTask) GetFinishedAtOk() (*string, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *IndexTask) SetFinishedAt(v string)`

SetFinishedAt sets FinishedAt field to given value.

### HasFinishedAt

`func (o *IndexTask) HasFinishedAt() bool`

HasFinishedAt returns a boolean if a field has been set.

### GetStartedAt

`func (o *IndexTask) GetStartedAt() string`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *IndexTask) GetStartedAtOk() (*string, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *IndexTask) SetStartedAt(v string)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *IndexTask) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetStatus

`func (o *IndexTask) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IndexTask) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IndexTask) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *IndexTask) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *IndexTask) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IndexTask) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IndexTask) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IndexTask) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUid

`func (o *IndexTask) GetUid() int64`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *IndexTask) GetUidOk() (*int64, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *IndexTask) SetUid(v int64)`

SetUid sets Uid field to given value.

### HasUid

`func (o *IndexTask) HasUid() bool`

HasUid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


