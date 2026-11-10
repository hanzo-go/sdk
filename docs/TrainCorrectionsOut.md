# TrainCorrectionsOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]TrainCorrection**](TrainCorrection.md) |  | [optional] 
**Questions** | Pointer to [**[]TrainCorrectionGroup**](TrainCorrectionGroup.md) |  | [optional] 

## Methods

### NewTrainCorrectionsOut

`func NewTrainCorrectionsOut() *TrainCorrectionsOut`

NewTrainCorrectionsOut instantiates a new TrainCorrectionsOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrainCorrectionsOutWithDefaults

`func NewTrainCorrectionsOutWithDefaults() *TrainCorrectionsOut`

NewTrainCorrectionsOutWithDefaults instantiates a new TrainCorrectionsOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *TrainCorrectionsOut) GetData() []TrainCorrection`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TrainCorrectionsOut) GetDataOk() (*[]TrainCorrection, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TrainCorrectionsOut) SetData(v []TrainCorrection)`

SetData sets Data field to given value.

### HasData

`func (o *TrainCorrectionsOut) HasData() bool`

HasData returns a boolean if a field has been set.

### GetQuestions

`func (o *TrainCorrectionsOut) GetQuestions() []TrainCorrectionGroup`

GetQuestions returns the Questions field if non-nil, zero value otherwise.

### GetQuestionsOk

`func (o *TrainCorrectionsOut) GetQuestionsOk() (*[]TrainCorrectionGroup, bool)`

GetQuestionsOk returns a tuple with the Questions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestions

`func (o *TrainCorrectionsOut) SetQuestions(v []TrainCorrectionGroup)`

SetQuestions sets Questions field to given value.

### HasQuestions

`func (o *TrainCorrectionsOut) HasQuestions() bool`

HasQuestions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


