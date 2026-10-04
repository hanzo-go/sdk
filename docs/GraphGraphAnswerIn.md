# GraphGraphAnswerIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AsKnown** | Pointer to **string** | AsKnown is how much this plane had heard, RFC 3339: the answer as it would have been given then, which is what makes a past answer reproducible. Absent is now. | [optional] 
**AsOf** | Pointer to **string** | AsOf answers from the graph as it stood at an instant of the world, RFC 3339: what held then. Absent is now. | [optional] 
**Question** | **string** | Question is what to answer, in words, 1000 bytes at most. Its words also decide which communities are read first, so naming the things it is about is what aims it. | 

## Methods

### NewGraphGraphAnswerIn

`func NewGraphGraphAnswerIn(question string, ) *GraphGraphAnswerIn`

NewGraphGraphAnswerIn instantiates a new GraphGraphAnswerIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphGraphAnswerInWithDefaults

`func NewGraphGraphAnswerInWithDefaults() *GraphGraphAnswerIn`

NewGraphGraphAnswerInWithDefaults instantiates a new GraphGraphAnswerIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsKnown

`func (o *GraphGraphAnswerIn) GetAsKnown() string`

GetAsKnown returns the AsKnown field if non-nil, zero value otherwise.

### GetAsKnownOk

`func (o *GraphGraphAnswerIn) GetAsKnownOk() (*string, bool)`

GetAsKnownOk returns a tuple with the AsKnown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsKnown

`func (o *GraphGraphAnswerIn) SetAsKnown(v string)`

SetAsKnown sets AsKnown field to given value.

### HasAsKnown

`func (o *GraphGraphAnswerIn) HasAsKnown() bool`

HasAsKnown returns a boolean if a field has been set.

### GetAsOf

`func (o *GraphGraphAnswerIn) GetAsOf() string`

GetAsOf returns the AsOf field if non-nil, zero value otherwise.

### GetAsOfOk

`func (o *GraphGraphAnswerIn) GetAsOfOk() (*string, bool)`

GetAsOfOk returns a tuple with the AsOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsOf

`func (o *GraphGraphAnswerIn) SetAsOf(v string)`

SetAsOf sets AsOf field to given value.

### HasAsOf

`func (o *GraphGraphAnswerIn) HasAsOf() bool`

HasAsOf returns a boolean if a field has been set.

### GetQuestion

`func (o *GraphGraphAnswerIn) GetQuestion() string`

GetQuestion returns the Question field if non-nil, zero value otherwise.

### GetQuestionOk

`func (o *GraphGraphAnswerIn) GetQuestionOk() (*string, bool)`

GetQuestionOk returns a tuple with the Question field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestion

`func (o *GraphGraphAnswerIn) SetQuestion(v string)`

SetQuestion sets Question field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


