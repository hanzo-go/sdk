# MarketingBetaAnswer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Answer** | Pointer to **[]string** | Answer is what the person answered: one entry for a typed answer or a single choice, one entry per option for several. | [optional] 
**Id** | Pointer to **string** | ID is the question&#39;s id on the form, such as \&quot;builds\&quot;. | [optional] 
**Question** | Pointer to **string** | Question is the question as the form asked it. | [optional] 

## Methods

### NewMarketingBetaAnswer

`func NewMarketingBetaAnswer() *MarketingBetaAnswer`

NewMarketingBetaAnswer instantiates a new MarketingBetaAnswer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketingBetaAnswerWithDefaults

`func NewMarketingBetaAnswerWithDefaults() *MarketingBetaAnswer`

NewMarketingBetaAnswerWithDefaults instantiates a new MarketingBetaAnswer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnswer

`func (o *MarketingBetaAnswer) GetAnswer() []string`

GetAnswer returns the Answer field if non-nil, zero value otherwise.

### GetAnswerOk

`func (o *MarketingBetaAnswer) GetAnswerOk() (*[]string, bool)`

GetAnswerOk returns a tuple with the Answer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswer

`func (o *MarketingBetaAnswer) SetAnswer(v []string)`

SetAnswer sets Answer field to given value.

### HasAnswer

`func (o *MarketingBetaAnswer) HasAnswer() bool`

HasAnswer returns a boolean if a field has been set.

### GetId

`func (o *MarketingBetaAnswer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MarketingBetaAnswer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MarketingBetaAnswer) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MarketingBetaAnswer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetQuestion

`func (o *MarketingBetaAnswer) GetQuestion() string`

GetQuestion returns the Question field if non-nil, zero value otherwise.

### GetQuestionOk

`func (o *MarketingBetaAnswer) GetQuestionOk() (*string, bool)`

GetQuestionOk returns a tuple with the Question field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestion

`func (o *MarketingBetaAnswer) SetQuestion(v string)`

SetQuestion sets Question field to given value.

### HasQuestion

`func (o *MarketingBetaAnswer) HasQuestion() bool`

HasQuestion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


