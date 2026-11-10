# AiChat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Answer** | Pointer to **string** | Answer is the model&#39;s reply. | [optional] 
**Model** | Pointer to **string** | Model is the model asked: the caller&#39;s, or the default when it named none. | [optional] 
**Receipt** | Pointer to [**AiReceipt**](AiReceipt.md) | Receipt says which model served the reply and who paid for it. | [optional] 

## Methods

### NewAiChat

`func NewAiChat() *AiChat`

NewAiChat instantiates a new AiChat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiChatWithDefaults

`func NewAiChatWithDefaults() *AiChat`

NewAiChatWithDefaults instantiates a new AiChat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnswer

`func (o *AiChat) GetAnswer() string`

GetAnswer returns the Answer field if non-nil, zero value otherwise.

### GetAnswerOk

`func (o *AiChat) GetAnswerOk() (*string, bool)`

GetAnswerOk returns a tuple with the Answer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswer

`func (o *AiChat) SetAnswer(v string)`

SetAnswer sets Answer field to given value.

### HasAnswer

`func (o *AiChat) HasAnswer() bool`

HasAnswer returns a boolean if a field has been set.

### GetModel

`func (o *AiChat) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *AiChat) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *AiChat) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *AiChat) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetReceipt

`func (o *AiChat) GetReceipt() AiReceipt`

GetReceipt returns the Receipt field if non-nil, zero value otherwise.

### GetReceiptOk

`func (o *AiChat) GetReceiptOk() (*AiReceipt, bool)`

GetReceiptOk returns a tuple with the Receipt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceipt

`func (o *AiChat) SetReceipt(v AiReceipt)`

SetReceipt sets Receipt field to given value.

### HasReceipt

`func (o *AiChat) HasReceipt() bool`

HasReceipt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


