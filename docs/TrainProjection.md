# TrainProjection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Strength** | Pointer to **float64** | Strength is λ in [0, 1]; absent is 1, the whole of P g removed. | [optional] 

## Methods

### NewTrainProjection

`func NewTrainProjection() *TrainProjection`

NewTrainProjection instantiates a new TrainProjection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrainProjectionWithDefaults

`func NewTrainProjectionWithDefaults() *TrainProjection`

NewTrainProjectionWithDefaults instantiates a new TrainProjection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStrength

`func (o *TrainProjection) GetStrength() float64`

GetStrength returns the Strength field if non-nil, zero value otherwise.

### GetStrengthOk

`func (o *TrainProjection) GetStrengthOk() (*float64, bool)`

GetStrengthOk returns a tuple with the Strength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrength

`func (o *TrainProjection) SetStrength(v float64)`

SetStrength sets Strength field to given value.

### HasStrength

`func (o *TrainProjection) HasStrength() bool`

HasStrength returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


