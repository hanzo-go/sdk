# CorrectIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Answer** | Pointer to **string** | Answer is the right option label; a score&#39;s level may be named by its index. | [optional] 
**Cohort** | Pointer to **string** | Cohort groups corrections the org files together. | [optional] 
**Decision** | Pointer to **string** | Decision is the decision&#39;s id. | [optional] 
**Note** | Pointer to **string** | Note says why, for the person reading it later. | [optional] 
**Question** | Pointer to **string** | Question is the question&#39;s key. | [optional] 

## Methods

### NewCorrectIn

`func NewCorrectIn() *CorrectIn`

NewCorrectIn instantiates a new CorrectIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCorrectInWithDefaults

`func NewCorrectInWithDefaults() *CorrectIn`

NewCorrectInWithDefaults instantiates a new CorrectIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnswer

`func (o *CorrectIn) GetAnswer() string`

GetAnswer returns the Answer field if non-nil, zero value otherwise.

### GetAnswerOk

`func (o *CorrectIn) GetAnswerOk() (*string, bool)`

GetAnswerOk returns a tuple with the Answer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswer

`func (o *CorrectIn) SetAnswer(v string)`

SetAnswer sets Answer field to given value.

### HasAnswer

`func (o *CorrectIn) HasAnswer() bool`

HasAnswer returns a boolean if a field has been set.

### GetCohort

`func (o *CorrectIn) GetCohort() string`

GetCohort returns the Cohort field if non-nil, zero value otherwise.

### GetCohortOk

`func (o *CorrectIn) GetCohortOk() (*string, bool)`

GetCohortOk returns a tuple with the Cohort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCohort

`func (o *CorrectIn) SetCohort(v string)`

SetCohort sets Cohort field to given value.

### HasCohort

`func (o *CorrectIn) HasCohort() bool`

HasCohort returns a boolean if a field has been set.

### GetDecision

`func (o *CorrectIn) GetDecision() string`

GetDecision returns the Decision field if non-nil, zero value otherwise.

### GetDecisionOk

`func (o *CorrectIn) GetDecisionOk() (*string, bool)`

GetDecisionOk returns a tuple with the Decision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecision

`func (o *CorrectIn) SetDecision(v string)`

SetDecision sets Decision field to given value.

### HasDecision

`func (o *CorrectIn) HasDecision() bool`

HasDecision returns a boolean if a field has been set.

### GetNote

`func (o *CorrectIn) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *CorrectIn) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *CorrectIn) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *CorrectIn) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetQuestion

`func (o *CorrectIn) GetQuestion() string`

GetQuestion returns the Question field if non-nil, zero value otherwise.

### GetQuestionOk

`func (o *CorrectIn) GetQuestionOk() (*string, bool)`

GetQuestionOk returns a tuple with the Question field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestion

`func (o *CorrectIn) SetQuestion(v string)`

SetQuestion sets Question field to given value.

### HasQuestion

`func (o *CorrectIn) HasQuestion() bool`

HasQuestion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


