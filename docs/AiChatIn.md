# AiChatIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxTokens** | Pointer to **int64** | MaxTokens bounds the answer in tokens; 0 leaves it to the model. | [optional] 
**Model** | Pointer to **string** | Model is the model to answer, by id (zen5, enso, kai, anthropic/claude-sonnet-4.5 — aiModels lists them). Empty asks the deployment&#39;s default Hanzo model. | [optional] 
**Prompt** | Pointer to **string** | Prompt is what to ask. | [optional] 
**System** | Pointer to **string** | System is the system message ahead of the prompt, when set. | [optional] 

## Methods

### NewAiChatIn

`func NewAiChatIn() *AiChatIn`

NewAiChatIn instantiates a new AiChatIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiChatInWithDefaults

`func NewAiChatInWithDefaults() *AiChatIn`

NewAiChatInWithDefaults instantiates a new AiChatIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxTokens

`func (o *AiChatIn) GetMaxTokens() int64`

GetMaxTokens returns the MaxTokens field if non-nil, zero value otherwise.

### GetMaxTokensOk

`func (o *AiChatIn) GetMaxTokensOk() (*int64, bool)`

GetMaxTokensOk returns a tuple with the MaxTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTokens

`func (o *AiChatIn) SetMaxTokens(v int64)`

SetMaxTokens sets MaxTokens field to given value.

### HasMaxTokens

`func (o *AiChatIn) HasMaxTokens() bool`

HasMaxTokens returns a boolean if a field has been set.

### GetModel

`func (o *AiChatIn) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *AiChatIn) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *AiChatIn) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *AiChatIn) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetPrompt

`func (o *AiChatIn) GetPrompt() string`

GetPrompt returns the Prompt field if non-nil, zero value otherwise.

### GetPromptOk

`func (o *AiChatIn) GetPromptOk() (*string, bool)`

GetPromptOk returns a tuple with the Prompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompt

`func (o *AiChatIn) SetPrompt(v string)`

SetPrompt sets Prompt field to given value.

### HasPrompt

`func (o *AiChatIn) HasPrompt() bool`

HasPrompt returns a boolean if a field has been set.

### GetSystem

`func (o *AiChatIn) GetSystem() string`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *AiChatIn) GetSystemOk() (*string, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *AiChatIn) SetSystem(v string)`

SetSystem sets System field to given value.

### HasSystem

`func (o *AiChatIn) HasSystem() bool`

HasSystem returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


