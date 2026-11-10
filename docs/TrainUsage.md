# TrainUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CostMicroUsd** | Pointer to **int64** | CostMicroUSD is what Seconds were charged, in micro-USD: 0 on the org&#39;s own machines. | [optional] 
**Held** | Pointer to **int64** | Held are the device-seconds held beyond Seconds; 0 once it ended. | [optional] 
**Seconds** | Pointer to **int64** | Seconds are the device-seconds used. | [optional] 

## Methods

### NewTrainUsage

`func NewTrainUsage() *TrainUsage`

NewTrainUsage instantiates a new TrainUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrainUsageWithDefaults

`func NewTrainUsageWithDefaults() *TrainUsage`

NewTrainUsageWithDefaults instantiates a new TrainUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCostMicroUsd

`func (o *TrainUsage) GetCostMicroUsd() int64`

GetCostMicroUsd returns the CostMicroUsd field if non-nil, zero value otherwise.

### GetCostMicroUsdOk

`func (o *TrainUsage) GetCostMicroUsdOk() (*int64, bool)`

GetCostMicroUsdOk returns a tuple with the CostMicroUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostMicroUsd

`func (o *TrainUsage) SetCostMicroUsd(v int64)`

SetCostMicroUsd sets CostMicroUsd field to given value.

### HasCostMicroUsd

`func (o *TrainUsage) HasCostMicroUsd() bool`

HasCostMicroUsd returns a boolean if a field has been set.

### GetHeld

`func (o *TrainUsage) GetHeld() int64`

GetHeld returns the Held field if non-nil, zero value otherwise.

### GetHeldOk

`func (o *TrainUsage) GetHeldOk() (*int64, bool)`

GetHeldOk returns a tuple with the Held field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeld

`func (o *TrainUsage) SetHeld(v int64)`

SetHeld sets Held field to given value.

### HasHeld

`func (o *TrainUsage) HasHeld() bool`

HasHeld returns a boolean if a field has been set.

### GetSeconds

`func (o *TrainUsage) GetSeconds() int64`

GetSeconds returns the Seconds field if non-nil, zero value otherwise.

### GetSecondsOk

`func (o *TrainUsage) GetSecondsOk() (*int64, bool)`

GetSecondsOk returns a tuple with the Seconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeconds

`func (o *TrainUsage) SetSeconds(v int64)`

SetSeconds sets Seconds field to given value.

### HasSeconds

`func (o *TrainUsage) HasSeconds() bool`

HasSeconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


