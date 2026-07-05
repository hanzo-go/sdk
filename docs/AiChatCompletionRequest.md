# AiChatCompletionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Model** | **string** |  | 
**Messages** | [**[]AiChatMessage**](AiChatMessage.md) |  | 
**Stream** | Pointer to **bool** |  | [optional] [default to false]
**Temperature** | Pointer to **float32** |  | [optional] 
**TopP** | Pointer to **float32** |  | [optional] 
**MaxTokens** | Pointer to **int32** |  | [optional] 
**Tools** | Pointer to **[]map[string]interface{}** |  | [optional] 
**ToolChoice** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewAiChatCompletionRequest

`func NewAiChatCompletionRequest(model string, messages []AiChatMessage, ) *AiChatCompletionRequest`

NewAiChatCompletionRequest instantiates a new AiChatCompletionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiChatCompletionRequestWithDefaults

`func NewAiChatCompletionRequestWithDefaults() *AiChatCompletionRequest`

NewAiChatCompletionRequestWithDefaults instantiates a new AiChatCompletionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModel

`func (o *AiChatCompletionRequest) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *AiChatCompletionRequest) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *AiChatCompletionRequest) SetModel(v string)`

SetModel sets Model field to given value.


### GetMessages

`func (o *AiChatCompletionRequest) GetMessages() []AiChatMessage`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *AiChatCompletionRequest) GetMessagesOk() (*[]AiChatMessage, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *AiChatCompletionRequest) SetMessages(v []AiChatMessage)`

SetMessages sets Messages field to given value.


### GetStream

`func (o *AiChatCompletionRequest) GetStream() bool`

GetStream returns the Stream field if non-nil, zero value otherwise.

### GetStreamOk

`func (o *AiChatCompletionRequest) GetStreamOk() (*bool, bool)`

GetStreamOk returns a tuple with the Stream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStream

`func (o *AiChatCompletionRequest) SetStream(v bool)`

SetStream sets Stream field to given value.

### HasStream

`func (o *AiChatCompletionRequest) HasStream() bool`

HasStream returns a boolean if a field has been set.

### GetTemperature

`func (o *AiChatCompletionRequest) GetTemperature() float32`

GetTemperature returns the Temperature field if non-nil, zero value otherwise.

### GetTemperatureOk

`func (o *AiChatCompletionRequest) GetTemperatureOk() (*float32, bool)`

GetTemperatureOk returns a tuple with the Temperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperature

`func (o *AiChatCompletionRequest) SetTemperature(v float32)`

SetTemperature sets Temperature field to given value.

### HasTemperature

`func (o *AiChatCompletionRequest) HasTemperature() bool`

HasTemperature returns a boolean if a field has been set.

### GetTopP

`func (o *AiChatCompletionRequest) GetTopP() float32`

GetTopP returns the TopP field if non-nil, zero value otherwise.

### GetTopPOk

`func (o *AiChatCompletionRequest) GetTopPOk() (*float32, bool)`

GetTopPOk returns a tuple with the TopP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopP

`func (o *AiChatCompletionRequest) SetTopP(v float32)`

SetTopP sets TopP field to given value.

### HasTopP

`func (o *AiChatCompletionRequest) HasTopP() bool`

HasTopP returns a boolean if a field has been set.

### GetMaxTokens

`func (o *AiChatCompletionRequest) GetMaxTokens() int32`

GetMaxTokens returns the MaxTokens field if non-nil, zero value otherwise.

### GetMaxTokensOk

`func (o *AiChatCompletionRequest) GetMaxTokensOk() (*int32, bool)`

GetMaxTokensOk returns a tuple with the MaxTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTokens

`func (o *AiChatCompletionRequest) SetMaxTokens(v int32)`

SetMaxTokens sets MaxTokens field to given value.

### HasMaxTokens

`func (o *AiChatCompletionRequest) HasMaxTokens() bool`

HasMaxTokens returns a boolean if a field has been set.

### GetTools

`func (o *AiChatCompletionRequest) GetTools() []map[string]interface{}`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *AiChatCompletionRequest) GetToolsOk() (*[]map[string]interface{}, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *AiChatCompletionRequest) SetTools(v []map[string]interface{})`

SetTools sets Tools field to given value.

### HasTools

`func (o *AiChatCompletionRequest) HasTools() bool`

HasTools returns a boolean if a field has been set.

### GetToolChoice

`func (o *AiChatCompletionRequest) GetToolChoice() interface{}`

GetToolChoice returns the ToolChoice field if non-nil, zero value otherwise.

### GetToolChoiceOk

`func (o *AiChatCompletionRequest) GetToolChoiceOk() (*interface{}, bool)`

GetToolChoiceOk returns a tuple with the ToolChoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolChoice

`func (o *AiChatCompletionRequest) SetToolChoice(v interface{})`

SetToolChoice sets ToolChoice field to given value.

### HasToolChoice

`func (o *AiChatCompletionRequest) HasToolChoice() bool`

HasToolChoice returns a boolean if a field has been set.

### SetToolChoiceNil

`func (o *AiChatCompletionRequest) SetToolChoiceNil(b bool)`

 SetToolChoiceNil sets the value for ToolChoice to be an explicit nil

### UnsetToolChoice
`func (o *AiChatCompletionRequest) UnsetToolChoice()`

UnsetToolChoice ensures that no value is present for ToolChoice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


