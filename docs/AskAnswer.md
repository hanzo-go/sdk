# AskAnswer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Answer** | Pointer to **string** |  | [optional] 
**Citations** | Pointer to [**[]Citation**](Citation.md) |  | [optional] 
**Degraded** | Pointer to **bool** |  | [optional] 
**Question** | Pointer to **string** |  | [optional] 

## Methods

### NewAskAnswer

`func NewAskAnswer() *AskAnswer`

NewAskAnswer instantiates a new AskAnswer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAskAnswerWithDefaults

`func NewAskAnswerWithDefaults() *AskAnswer`

NewAskAnswerWithDefaults instantiates a new AskAnswer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnswer

`func (o *AskAnswer) GetAnswer() string`

GetAnswer returns the Answer field if non-nil, zero value otherwise.

### GetAnswerOk

`func (o *AskAnswer) GetAnswerOk() (*string, bool)`

GetAnswerOk returns a tuple with the Answer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswer

`func (o *AskAnswer) SetAnswer(v string)`

SetAnswer sets Answer field to given value.

### HasAnswer

`func (o *AskAnswer) HasAnswer() bool`

HasAnswer returns a boolean if a field has been set.

### GetCitations

`func (o *AskAnswer) GetCitations() []Citation`

GetCitations returns the Citations field if non-nil, zero value otherwise.

### GetCitationsOk

`func (o *AskAnswer) GetCitationsOk() (*[]Citation, bool)`

GetCitationsOk returns a tuple with the Citations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitations

`func (o *AskAnswer) SetCitations(v []Citation)`

SetCitations sets Citations field to given value.

### HasCitations

`func (o *AskAnswer) HasCitations() bool`

HasCitations returns a boolean if a field has been set.

### GetDegraded

`func (o *AskAnswer) GetDegraded() bool`

GetDegraded returns the Degraded field if non-nil, zero value otherwise.

### GetDegradedOk

`func (o *AskAnswer) GetDegradedOk() (*bool, bool)`

GetDegradedOk returns a tuple with the Degraded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDegraded

`func (o *AskAnswer) SetDegraded(v bool)`

SetDegraded sets Degraded field to given value.

### HasDegraded

`func (o *AskAnswer) HasDegraded() bool`

HasDegraded returns a boolean if a field has been set.

### GetQuestion

`func (o *AskAnswer) GetQuestion() string`

GetQuestion returns the Question field if non-nil, zero value otherwise.

### GetQuestionOk

`func (o *AskAnswer) GetQuestionOk() (*string, bool)`

GetQuestionOk returns a tuple with the Question field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestion

`func (o *AskAnswer) SetQuestion(v string)`

SetQuestion sets Question field to given value.

### HasQuestion

`func (o *AskAnswer) HasQuestion() bool`

HasQuestion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


