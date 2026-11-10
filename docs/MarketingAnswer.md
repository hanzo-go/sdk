# MarketingAnswer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Answer** | Pointer to **string** | Answer is what the person said; options they picked are joined by commas. | [optional] 
**Question** | Pointer to **string** | Question is the question as the form asked it, on one line. | [optional] 

## Methods

### NewMarketingAnswer

`func NewMarketingAnswer() *MarketingAnswer`

NewMarketingAnswer instantiates a new MarketingAnswer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketingAnswerWithDefaults

`func NewMarketingAnswerWithDefaults() *MarketingAnswer`

NewMarketingAnswerWithDefaults instantiates a new MarketingAnswer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnswer

`func (o *MarketingAnswer) GetAnswer() string`

GetAnswer returns the Answer field if non-nil, zero value otherwise.

### GetAnswerOk

`func (o *MarketingAnswer) GetAnswerOk() (*string, bool)`

GetAnswerOk returns a tuple with the Answer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswer

`func (o *MarketingAnswer) SetAnswer(v string)`

SetAnswer sets Answer field to given value.

### HasAnswer

`func (o *MarketingAnswer) HasAnswer() bool`

HasAnswer returns a boolean if a field has been set.

### GetQuestion

`func (o *MarketingAnswer) GetQuestion() string`

GetQuestion returns the Question field if non-nil, zero value otherwise.

### GetQuestionOk

`func (o *MarketingAnswer) GetQuestionOk() (*string, bool)`

GetQuestionOk returns a tuple with the Question field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestion

`func (o *MarketingAnswer) SetQuestion(v string)`

SetQuestion sets Question field to given value.

### HasQuestion

`func (o *MarketingAnswer) HasQuestion() bool`

HasQuestion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


