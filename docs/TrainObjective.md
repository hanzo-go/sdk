# TrainObjective

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Loss** | Pointer to **string** | Loss is &#x60;cross_entropy&#x60;. | [optional] 
**Terms** | Pointer to [**[]TrainTerm**](TrainTerm.md) | Terms are added to each row&#39;s loss. | [optional] 

## Methods

### NewTrainObjective

`func NewTrainObjective() *TrainObjective`

NewTrainObjective instantiates a new TrainObjective object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrainObjectiveWithDefaults

`func NewTrainObjectiveWithDefaults() *TrainObjective`

NewTrainObjectiveWithDefaults instantiates a new TrainObjective object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLoss

`func (o *TrainObjective) GetLoss() string`

GetLoss returns the Loss field if non-nil, zero value otherwise.

### GetLossOk

`func (o *TrainObjective) GetLossOk() (*string, bool)`

GetLossOk returns a tuple with the Loss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoss

`func (o *TrainObjective) SetLoss(v string)`

SetLoss sets Loss field to given value.

### HasLoss

`func (o *TrainObjective) HasLoss() bool`

HasLoss returns a boolean if a field has been set.

### GetTerms

`func (o *TrainObjective) GetTerms() []TrainTerm`

GetTerms returns the Terms field if non-nil, zero value otherwise.

### GetTermsOk

`func (o *TrainObjective) GetTermsOk() (*[]TrainTerm, bool)`

GetTermsOk returns a tuple with the Terms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerms

`func (o *TrainObjective) SetTerms(v []TrainTerm)`

SetTerms sets Terms field to given value.

### HasTerms

`func (o *TrainObjective) HasTerms() bool`

HasTerms returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


