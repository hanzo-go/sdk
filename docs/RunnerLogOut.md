# RunnerLogOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ack** | Pointer to **int64** | Ack is how far the log is durable — index plus lines accepted — and where the runner resends from. | [optional] 

## Methods

### NewRunnerLogOut

`func NewRunnerLogOut() *RunnerLogOut`

NewRunnerLogOut instantiates a new RunnerLogOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunnerLogOutWithDefaults

`func NewRunnerLogOutWithDefaults() *RunnerLogOut`

NewRunnerLogOutWithDefaults instantiates a new RunnerLogOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAck

`func (o *RunnerLogOut) GetAck() int64`

GetAck returns the Ack field if non-nil, zero value otherwise.

### GetAckOk

`func (o *RunnerLogOut) GetAckOk() (*int64, bool)`

GetAckOk returns a tuple with the Ack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAck

`func (o *RunnerLogOut) SetAck(v int64)`

SetAck sets Ack field to given value.

### HasAck

`func (o *RunnerLogOut) HasAck() bool`

HasAck returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


