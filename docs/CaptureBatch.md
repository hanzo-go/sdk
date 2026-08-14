# CaptureBatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Batch** | Pointer to [**[]CaptureEvent**](CaptureEvent.md) |  | [optional] 
**Events** | Pointer to [**[]CaptureEvent**](CaptureEvent.md) |  | [optional] 

## Methods

### NewCaptureBatch

`func NewCaptureBatch() *CaptureBatch`

NewCaptureBatch instantiates a new CaptureBatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCaptureBatchWithDefaults

`func NewCaptureBatchWithDefaults() *CaptureBatch`

NewCaptureBatchWithDefaults instantiates a new CaptureBatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBatch

`func (o *CaptureBatch) GetBatch() []CaptureEvent`

GetBatch returns the Batch field if non-nil, zero value otherwise.

### GetBatchOk

`func (o *CaptureBatch) GetBatchOk() (*[]CaptureEvent, bool)`

GetBatchOk returns a tuple with the Batch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatch

`func (o *CaptureBatch) SetBatch(v []CaptureEvent)`

SetBatch sets Batch field to given value.

### HasBatch

`func (o *CaptureBatch) HasBatch() bool`

HasBatch returns a boolean if a field has been set.

### GetEvents

`func (o *CaptureBatch) GetEvents() []CaptureEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *CaptureBatch) GetEventsOk() (*[]CaptureEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *CaptureBatch) SetEvents(v []CaptureEvent)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *CaptureBatch) HasEvents() bool`

HasEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


