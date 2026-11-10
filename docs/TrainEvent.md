# TrainEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**At** | Pointer to **int64** | At is when it was recorded, unix seconds. | [optional] 
**Event** | Pointer to **interface{}** |  | [optional] 
**Seq** | Pointer to **int64** | Seq orders the job&#39;s events from 1. | [optional] 
**Task** | Pointer to **string** | Task is the task that reported it. | [optional] 

## Methods

### NewTrainEvent

`func NewTrainEvent() *TrainEvent`

NewTrainEvent instantiates a new TrainEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrainEventWithDefaults

`func NewTrainEventWithDefaults() *TrainEvent`

NewTrainEventWithDefaults instantiates a new TrainEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAt

`func (o *TrainEvent) GetAt() int64`

GetAt returns the At field if non-nil, zero value otherwise.

### GetAtOk

`func (o *TrainEvent) GetAtOk() (*int64, bool)`

GetAtOk returns a tuple with the At field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAt

`func (o *TrainEvent) SetAt(v int64)`

SetAt sets At field to given value.

### HasAt

`func (o *TrainEvent) HasAt() bool`

HasAt returns a boolean if a field has been set.

### GetEvent

`func (o *TrainEvent) GetEvent() interface{}`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *TrainEvent) GetEventOk() (*interface{}, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *TrainEvent) SetEvent(v interface{})`

SetEvent sets Event field to given value.

### HasEvent

`func (o *TrainEvent) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### SetEventNil

`func (o *TrainEvent) SetEventNil(b bool)`

 SetEventNil sets the value for Event to be an explicit nil

### UnsetEvent
`func (o *TrainEvent) UnsetEvent()`

UnsetEvent ensures that no value is present for Event, not even an explicit nil
### GetSeq

`func (o *TrainEvent) GetSeq() int64`

GetSeq returns the Seq field if non-nil, zero value otherwise.

### GetSeqOk

`func (o *TrainEvent) GetSeqOk() (*int64, bool)`

GetSeqOk returns a tuple with the Seq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeq

`func (o *TrainEvent) SetSeq(v int64)`

SetSeq sets Seq field to given value.

### HasSeq

`func (o *TrainEvent) HasSeq() bool`

HasSeq returns a boolean if a field has been set.

### GetTask

`func (o *TrainEvent) GetTask() string`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *TrainEvent) GetTaskOk() (*string, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *TrainEvent) SetTask(v string)`

SetTask sets Task field to given value.

### HasTask

`func (o *TrainEvent) HasTask() bool`

HasTask returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


