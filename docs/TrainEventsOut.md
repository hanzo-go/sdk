# TrainEventsOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]TrainEvent**](TrainEvent.md) | Data is the events after the cursor, in order. | [optional] 
**Next** | Pointer to **int64** | Next is the cursor to pass as after. | [optional] 
**Status** | Pointer to **string** | Status is the job&#39;s status as of this answer. | [optional] 

## Methods

### NewTrainEventsOut

`func NewTrainEventsOut() *TrainEventsOut`

NewTrainEventsOut instantiates a new TrainEventsOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrainEventsOutWithDefaults

`func NewTrainEventsOutWithDefaults() *TrainEventsOut`

NewTrainEventsOutWithDefaults instantiates a new TrainEventsOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *TrainEventsOut) GetData() []TrainEvent`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TrainEventsOut) GetDataOk() (*[]TrainEvent, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TrainEventsOut) SetData(v []TrainEvent)`

SetData sets Data field to given value.

### HasData

`func (o *TrainEventsOut) HasData() bool`

HasData returns a boolean if a field has been set.

### GetNext

`func (o *TrainEventsOut) GetNext() int64`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *TrainEventsOut) GetNextOk() (*int64, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *TrainEventsOut) SetNext(v int64)`

SetNext sets Next field to given value.

### HasNext

`func (o *TrainEventsOut) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetStatus

`func (o *TrainEventsOut) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TrainEventsOut) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TrainEventsOut) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TrainEventsOut) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


