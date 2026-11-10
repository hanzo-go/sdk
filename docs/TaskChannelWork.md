# TaskChannelWork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | Pointer to **string** | Channel is the channel these counts are for, echoed back as it was resolved. | [optional] 
**Open** | Pointer to **int64** | Open is how many items are still work: everything whose status does not end it. It is the number a channel header shows. | [optional] 
**Status** | Pointer to **map[string]int64** | Status is the count per board column, carrying EVERY column this surface knows — an empty column reads 0 rather than being absent, so a caller can render the board without inventing the vocabulary. The keys are the same closed set every other operation here validates against. | [optional] 
**Total** | Pointer to **int64** | Total is every item bound to this channel, settled ones included, so Total minus Open is what the channel has finished. | [optional] 
**Updated** | Pointer to **int64** | Updated is when anything in this channel&#39;s work last moved, in unix seconds. ABSENT when the channel has no work at all: zero would read as the epoch, and a channel nobody has filed anything in has no last activity rather than an infinitely old one. Total is 0 in exactly that case. | [optional] 

## Methods

### NewTaskChannelWork

`func NewTaskChannelWork() *TaskChannelWork`

NewTaskChannelWork instantiates a new TaskChannelWork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskChannelWorkWithDefaults

`func NewTaskChannelWorkWithDefaults() *TaskChannelWork`

NewTaskChannelWorkWithDefaults instantiates a new TaskChannelWork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *TaskChannelWork) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *TaskChannelWork) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *TaskChannelWork) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *TaskChannelWork) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetOpen

`func (o *TaskChannelWork) GetOpen() int64`

GetOpen returns the Open field if non-nil, zero value otherwise.

### GetOpenOk

`func (o *TaskChannelWork) GetOpenOk() (*int64, bool)`

GetOpenOk returns a tuple with the Open field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpen

`func (o *TaskChannelWork) SetOpen(v int64)`

SetOpen sets Open field to given value.

### HasOpen

`func (o *TaskChannelWork) HasOpen() bool`

HasOpen returns a boolean if a field has been set.

### GetStatus

`func (o *TaskChannelWork) GetStatus() map[string]int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TaskChannelWork) GetStatusOk() (*map[string]int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TaskChannelWork) SetStatus(v map[string]int64)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TaskChannelWork) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTotal

`func (o *TaskChannelWork) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *TaskChannelWork) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *TaskChannelWork) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *TaskChannelWork) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetUpdated

`func (o *TaskChannelWork) GetUpdated() int64`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *TaskChannelWork) GetUpdatedOk() (*int64, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *TaskChannelWork) SetUpdated(v int64)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *TaskChannelWork) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


