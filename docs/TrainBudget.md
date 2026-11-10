# TrainBudget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accuracy** | Pointer to **float64** | Accuracy is the largest accuracy drop allowed. | [optional] 
**Ece** | Pointer to **float64** | ECE is the largest calibration-error rise allowed. | [optional] 

## Methods

### NewTrainBudget

`func NewTrainBudget() *TrainBudget`

NewTrainBudget instantiates a new TrainBudget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrainBudgetWithDefaults

`func NewTrainBudgetWithDefaults() *TrainBudget`

NewTrainBudgetWithDefaults instantiates a new TrainBudget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccuracy

`func (o *TrainBudget) GetAccuracy() float64`

GetAccuracy returns the Accuracy field if non-nil, zero value otherwise.

### GetAccuracyOk

`func (o *TrainBudget) GetAccuracyOk() (*float64, bool)`

GetAccuracyOk returns a tuple with the Accuracy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccuracy

`func (o *TrainBudget) SetAccuracy(v float64)`

SetAccuracy sets Accuracy field to given value.

### HasAccuracy

`func (o *TrainBudget) HasAccuracy() bool`

HasAccuracy returns a boolean if a field has been set.

### GetEce

`func (o *TrainBudget) GetEce() float64`

GetEce returns the Ece field if non-nil, zero value otherwise.

### GetEceOk

`func (o *TrainBudget) GetEceOk() (*float64, bool)`

GetEceOk returns a tuple with the Ece field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEce

`func (o *TrainBudget) SetEce(v float64)`

SetEce sets Ece field to given value.

### HasEce

`func (o *TrainBudget) HasEce() bool`

HasEce returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


