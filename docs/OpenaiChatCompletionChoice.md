# OpenaiChatCompletionChoice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentFilterResults** | Pointer to [**OpenaiContentFilterResults**](OpenaiContentFilterResults.md) |  | [optional] 
**FinishReason** | Pointer to **string** |  | [optional] 
**Index** | Pointer to **int32** |  | [optional] 
**Logprobs** | Pointer to [**OpenaiLogProbs**](OpenaiLogProbs.md) |  | [optional] 
**Message** | Pointer to [**OpenaiChatCompletionMessage**](OpenaiChatCompletionMessage.md) |  | [optional] 

## Methods

### NewOpenaiChatCompletionChoice

`func NewOpenaiChatCompletionChoice() *OpenaiChatCompletionChoice`

NewOpenaiChatCompletionChoice instantiates a new OpenaiChatCompletionChoice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenaiChatCompletionChoiceWithDefaults

`func NewOpenaiChatCompletionChoiceWithDefaults() *OpenaiChatCompletionChoice`

NewOpenaiChatCompletionChoiceWithDefaults instantiates a new OpenaiChatCompletionChoice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentFilterResults

`func (o *OpenaiChatCompletionChoice) GetContentFilterResults() OpenaiContentFilterResults`

GetContentFilterResults returns the ContentFilterResults field if non-nil, zero value otherwise.

### GetContentFilterResultsOk

`func (o *OpenaiChatCompletionChoice) GetContentFilterResultsOk() (*OpenaiContentFilterResults, bool)`

GetContentFilterResultsOk returns a tuple with the ContentFilterResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentFilterResults

`func (o *OpenaiChatCompletionChoice) SetContentFilterResults(v OpenaiContentFilterResults)`

SetContentFilterResults sets ContentFilterResults field to given value.

### HasContentFilterResults

`func (o *OpenaiChatCompletionChoice) HasContentFilterResults() bool`

HasContentFilterResults returns a boolean if a field has been set.

### GetFinishReason

`func (o *OpenaiChatCompletionChoice) GetFinishReason() string`

GetFinishReason returns the FinishReason field if non-nil, zero value otherwise.

### GetFinishReasonOk

`func (o *OpenaiChatCompletionChoice) GetFinishReasonOk() (*string, bool)`

GetFinishReasonOk returns a tuple with the FinishReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishReason

`func (o *OpenaiChatCompletionChoice) SetFinishReason(v string)`

SetFinishReason sets FinishReason field to given value.

### HasFinishReason

`func (o *OpenaiChatCompletionChoice) HasFinishReason() bool`

HasFinishReason returns a boolean if a field has been set.

### GetIndex

`func (o *OpenaiChatCompletionChoice) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *OpenaiChatCompletionChoice) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *OpenaiChatCompletionChoice) SetIndex(v int32)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *OpenaiChatCompletionChoice) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetLogprobs

`func (o *OpenaiChatCompletionChoice) GetLogprobs() OpenaiLogProbs`

GetLogprobs returns the Logprobs field if non-nil, zero value otherwise.

### GetLogprobsOk

`func (o *OpenaiChatCompletionChoice) GetLogprobsOk() (*OpenaiLogProbs, bool)`

GetLogprobsOk returns a tuple with the Logprobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogprobs

`func (o *OpenaiChatCompletionChoice) SetLogprobs(v OpenaiLogProbs)`

SetLogprobs sets Logprobs field to given value.

### HasLogprobs

`func (o *OpenaiChatCompletionChoice) HasLogprobs() bool`

HasLogprobs returns a boolean if a field has been set.

### GetMessage

`func (o *OpenaiChatCompletionChoice) GetMessage() OpenaiChatCompletionMessage`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *OpenaiChatCompletionChoice) GetMessageOk() (*OpenaiChatCompletionMessage, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *OpenaiChatCompletionChoice) SetMessage(v OpenaiChatCompletionMessage)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *OpenaiChatCompletionChoice) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


