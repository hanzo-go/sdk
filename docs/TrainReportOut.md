# TrainReportOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cancel** | Pointer to **bool** | Cancel says the job was cancelled, or stopped at what its payer or its own bounds could hold: kill the trainer, report nothing more. | [optional] 
**Held** | Pointer to **int64** | Held is the device-seconds held for the task now. | [optional] 
**Seq** | Pointer to **int64** | Seq is the job&#39;s last event seq. | [optional] 
**Status** | Pointer to **string** | Status is the job&#39;s status after this report. | [optional] 

## Methods

### NewTrainReportOut

`func NewTrainReportOut() *TrainReportOut`

NewTrainReportOut instantiates a new TrainReportOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrainReportOutWithDefaults

`func NewTrainReportOutWithDefaults() *TrainReportOut`

NewTrainReportOutWithDefaults instantiates a new TrainReportOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCancel

`func (o *TrainReportOut) GetCancel() bool`

GetCancel returns the Cancel field if non-nil, zero value otherwise.

### GetCancelOk

`func (o *TrainReportOut) GetCancelOk() (*bool, bool)`

GetCancelOk returns a tuple with the Cancel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancel

`func (o *TrainReportOut) SetCancel(v bool)`

SetCancel sets Cancel field to given value.

### HasCancel

`func (o *TrainReportOut) HasCancel() bool`

HasCancel returns a boolean if a field has been set.

### GetHeld

`func (o *TrainReportOut) GetHeld() int64`

GetHeld returns the Held field if non-nil, zero value otherwise.

### GetHeldOk

`func (o *TrainReportOut) GetHeldOk() (*int64, bool)`

GetHeldOk returns a tuple with the Held field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeld

`func (o *TrainReportOut) SetHeld(v int64)`

SetHeld sets Held field to given value.

### HasHeld

`func (o *TrainReportOut) HasHeld() bool`

HasHeld returns a boolean if a field has been set.

### GetSeq

`func (o *TrainReportOut) GetSeq() int64`

GetSeq returns the Seq field if non-nil, zero value otherwise.

### GetSeqOk

`func (o *TrainReportOut) GetSeqOk() (*int64, bool)`

GetSeqOk returns a tuple with the Seq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeq

`func (o *TrainReportOut) SetSeq(v int64)`

SetSeq sets Seq field to given value.

### HasSeq

`func (o *TrainReportOut) HasSeq() bool`

HasSeq returns a boolean if a field has been set.

### GetStatus

`func (o *TrainReportOut) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TrainReportOut) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TrainReportOut) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TrainReportOut) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


