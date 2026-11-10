# TrainTerm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** | Kind names the term. | [optional] 
**Margin** | Pointer to **float64** | Margin is a hinge&#39;s margin, in logits. | [optional] 
**Weight** | Pointer to **float64** | Weight scales it; 0 is off. | [optional] 

## Methods

### NewTrainTerm

`func NewTrainTerm() *TrainTerm`

NewTrainTerm instantiates a new TrainTerm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrainTermWithDefaults

`func NewTrainTermWithDefaults() *TrainTerm`

NewTrainTermWithDefaults instantiates a new TrainTerm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *TrainTerm) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *TrainTerm) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *TrainTerm) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *TrainTerm) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMargin

`func (o *TrainTerm) GetMargin() float64`

GetMargin returns the Margin field if non-nil, zero value otherwise.

### GetMarginOk

`func (o *TrainTerm) GetMarginOk() (*float64, bool)`

GetMarginOk returns a tuple with the Margin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMargin

`func (o *TrainTerm) SetMargin(v float64)`

SetMargin sets Margin field to given value.

### HasMargin

`func (o *TrainTerm) HasMargin() bool`

HasMargin returns a boolean if a field has been set.

### GetWeight

`func (o *TrainTerm) GetWeight() float64`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *TrainTerm) GetWeightOk() (*float64, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *TrainTerm) SetWeight(v float64)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *TrainTerm) HasWeight() bool`

HasWeight returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


