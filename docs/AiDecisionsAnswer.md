# AiDecisionsAnswer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to [**AiDecisionsAction**](AiDecisionsAction.md) |  | [optional] 
**AnswerConfidence** | Pointer to **float32** |  | [optional] 
**Choice** | Pointer to **string** |  | [optional] 
**Confidence** | Pointer to **float32** |  | [optional] 
**Legend** | Pointer to [**map[string]AiDecisionSidesFalse**](AiDecisionSidesFalse.md) |  | [optional] 
**Noul** | Pointer to **float32** |  | [optional] 
**Probabilities** | Pointer to **map[string]float32** |  | [optional] 
**Score** | Pointer to **float32** |  | [optional] 
**Type** | **string** |  | 

## Methods

### NewAiDecisionsAnswer

`func NewAiDecisionsAnswer(type_ string, ) *AiDecisionsAnswer`

NewAiDecisionsAnswer instantiates a new AiDecisionsAnswer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiDecisionsAnswerWithDefaults

`func NewAiDecisionsAnswerWithDefaults() *AiDecisionsAnswer`

NewAiDecisionsAnswerWithDefaults instantiates a new AiDecisionsAnswer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *AiDecisionsAnswer) GetAction() AiDecisionsAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *AiDecisionsAnswer) GetActionOk() (*AiDecisionsAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *AiDecisionsAnswer) SetAction(v AiDecisionsAction)`

SetAction sets Action field to given value.

### HasAction

`func (o *AiDecisionsAnswer) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetAnswerConfidence

`func (o *AiDecisionsAnswer) GetAnswerConfidence() float32`

GetAnswerConfidence returns the AnswerConfidence field if non-nil, zero value otherwise.

### GetAnswerConfidenceOk

`func (o *AiDecisionsAnswer) GetAnswerConfidenceOk() (*float32, bool)`

GetAnswerConfidenceOk returns a tuple with the AnswerConfidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswerConfidence

`func (o *AiDecisionsAnswer) SetAnswerConfidence(v float32)`

SetAnswerConfidence sets AnswerConfidence field to given value.

### HasAnswerConfidence

`func (o *AiDecisionsAnswer) HasAnswerConfidence() bool`

HasAnswerConfidence returns a boolean if a field has been set.

### GetChoice

`func (o *AiDecisionsAnswer) GetChoice() string`

GetChoice returns the Choice field if non-nil, zero value otherwise.

### GetChoiceOk

`func (o *AiDecisionsAnswer) GetChoiceOk() (*string, bool)`

GetChoiceOk returns a tuple with the Choice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChoice

`func (o *AiDecisionsAnswer) SetChoice(v string)`

SetChoice sets Choice field to given value.

### HasChoice

`func (o *AiDecisionsAnswer) HasChoice() bool`

HasChoice returns a boolean if a field has been set.

### GetConfidence

`func (o *AiDecisionsAnswer) GetConfidence() float32`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *AiDecisionsAnswer) GetConfidenceOk() (*float32, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *AiDecisionsAnswer) SetConfidence(v float32)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *AiDecisionsAnswer) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.

### GetLegend

`func (o *AiDecisionsAnswer) GetLegend() map[string]AiDecisionSidesFalse`

GetLegend returns the Legend field if non-nil, zero value otherwise.

### GetLegendOk

`func (o *AiDecisionsAnswer) GetLegendOk() (*map[string]AiDecisionSidesFalse, bool)`

GetLegendOk returns a tuple with the Legend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegend

`func (o *AiDecisionsAnswer) SetLegend(v map[string]AiDecisionSidesFalse)`

SetLegend sets Legend field to given value.

### HasLegend

`func (o *AiDecisionsAnswer) HasLegend() bool`

HasLegend returns a boolean if a field has been set.

### GetNoul

`func (o *AiDecisionsAnswer) GetNoul() float32`

GetNoul returns the Noul field if non-nil, zero value otherwise.

### GetNoulOk

`func (o *AiDecisionsAnswer) GetNoulOk() (*float32, bool)`

GetNoulOk returns a tuple with the Noul field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoul

`func (o *AiDecisionsAnswer) SetNoul(v float32)`

SetNoul sets Noul field to given value.

### HasNoul

`func (o *AiDecisionsAnswer) HasNoul() bool`

HasNoul returns a boolean if a field has been set.

### GetProbabilities

`func (o *AiDecisionsAnswer) GetProbabilities() map[string]float32`

GetProbabilities returns the Probabilities field if non-nil, zero value otherwise.

### GetProbabilitiesOk

`func (o *AiDecisionsAnswer) GetProbabilitiesOk() (*map[string]float32, bool)`

GetProbabilitiesOk returns a tuple with the Probabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProbabilities

`func (o *AiDecisionsAnswer) SetProbabilities(v map[string]float32)`

SetProbabilities sets Probabilities field to given value.

### HasProbabilities

`func (o *AiDecisionsAnswer) HasProbabilities() bool`

HasProbabilities returns a boolean if a field has been set.

### GetScore

`func (o *AiDecisionsAnswer) GetScore() float32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *AiDecisionsAnswer) GetScoreOk() (*float32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *AiDecisionsAnswer) SetScore(v float32)`

SetScore sets Score field to given value.

### HasScore

`func (o *AiDecisionsAnswer) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetType

`func (o *AiDecisionsAnswer) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AiDecisionsAnswer) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AiDecisionsAnswer) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


