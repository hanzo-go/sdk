# CloudCaptureBatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Batch** | Pointer to [**[]CloudCaptureEvent**](CloudCaptureEvent.md) |  | [optional] 
**Events** | Pointer to [**[]CloudCaptureEvent**](CloudCaptureEvent.md) |  | [optional] 

## Methods

### NewCloudCaptureBatch

`func NewCloudCaptureBatch() *CloudCaptureBatch`

NewCloudCaptureBatch instantiates a new CloudCaptureBatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudCaptureBatchWithDefaults

`func NewCloudCaptureBatchWithDefaults() *CloudCaptureBatch`

NewCloudCaptureBatchWithDefaults instantiates a new CloudCaptureBatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBatch

`func (o *CloudCaptureBatch) GetBatch() []CloudCaptureEvent`

GetBatch returns the Batch field if non-nil, zero value otherwise.

### GetBatchOk

`func (o *CloudCaptureBatch) GetBatchOk() (*[]CloudCaptureEvent, bool)`

GetBatchOk returns a tuple with the Batch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatch

`func (o *CloudCaptureBatch) SetBatch(v []CloudCaptureEvent)`

SetBatch sets Batch field to given value.

### HasBatch

`func (o *CloudCaptureBatch) HasBatch() bool`

HasBatch returns a boolean if a field has been set.

### GetEvents

`func (o *CloudCaptureBatch) GetEvents() []CloudCaptureEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *CloudCaptureBatch) GetEventsOk() (*[]CloudCaptureEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *CloudCaptureBatch) SetEvents(v []CloudCaptureEvent)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *CloudCaptureBatch) HasEvents() bool`

HasEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


