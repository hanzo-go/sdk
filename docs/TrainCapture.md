# TrainCapture

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**On** | Pointer to **bool** |  | [optional] 
**RetentionDays** | Pointer to **int64** | RetentionDays is how long a decision&#39;s states are kept unless a correction names it. | [optional] 

## Methods

### NewTrainCapture

`func NewTrainCapture() *TrainCapture`

NewTrainCapture instantiates a new TrainCapture object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrainCaptureWithDefaults

`func NewTrainCaptureWithDefaults() *TrainCapture`

NewTrainCaptureWithDefaults instantiates a new TrainCapture object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOn

`func (o *TrainCapture) GetOn() bool`

GetOn returns the On field if non-nil, zero value otherwise.

### GetOnOk

`func (o *TrainCapture) GetOnOk() (*bool, bool)`

GetOnOk returns a tuple with the On field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOn

`func (o *TrainCapture) SetOn(v bool)`

SetOn sets On field to given value.

### HasOn

`func (o *TrainCapture) HasOn() bool`

HasOn returns a boolean if a field has been set.

### GetRetentionDays

`func (o *TrainCapture) GetRetentionDays() int64`

GetRetentionDays returns the RetentionDays field if non-nil, zero value otherwise.

### GetRetentionDaysOk

`func (o *TrainCapture) GetRetentionDaysOk() (*int64, bool)`

GetRetentionDaysOk returns a tuple with the RetentionDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionDays

`func (o *TrainCapture) SetRetentionDays(v int64)`

SetRetentionDays sets RetentionDays field to given value.

### HasRetentionDays

`func (o *TrainCapture) HasRetentionDays() bool`

HasRetentionDays returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


