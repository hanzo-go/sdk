# AiAnthropicRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxTokens** | Pointer to **int32** |  | [optional] 
**Messages** | Pointer to [**[]AiAnthropicMessage**](AiAnthropicMessage.md) |  | [optional] 
**Model** | Pointer to **string** |  | [optional] 
**Stream** | Pointer to **bool** |  | [optional] 
**System** | Pointer to **interface{}** |  | [optional] 
**Temperature** | Pointer to **float32** |  | [optional] 
**Thinking** | Pointer to **interface{}** |  | [optional] 
**ToolChoice** | Pointer to **interface{}** |  | [optional] 
**Tools** | Pointer to [**[]AiAnthropicTool**](AiAnthropicTool.md) |  | [optional] 

## Methods

### NewAiAnthropicRequest

`func NewAiAnthropicRequest() *AiAnthropicRequest`

NewAiAnthropicRequest instantiates a new AiAnthropicRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiAnthropicRequestWithDefaults

`func NewAiAnthropicRequestWithDefaults() *AiAnthropicRequest`

NewAiAnthropicRequestWithDefaults instantiates a new AiAnthropicRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxTokens

`func (o *AiAnthropicRequest) GetMaxTokens() int32`

GetMaxTokens returns the MaxTokens field if non-nil, zero value otherwise.

### GetMaxTokensOk

`func (o *AiAnthropicRequest) GetMaxTokensOk() (*int32, bool)`

GetMaxTokensOk returns a tuple with the MaxTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTokens

`func (o *AiAnthropicRequest) SetMaxTokens(v int32)`

SetMaxTokens sets MaxTokens field to given value.

### HasMaxTokens

`func (o *AiAnthropicRequest) HasMaxTokens() bool`

HasMaxTokens returns a boolean if a field has been set.

### GetMessages

`func (o *AiAnthropicRequest) GetMessages() []AiAnthropicMessage`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *AiAnthropicRequest) GetMessagesOk() (*[]AiAnthropicMessage, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *AiAnthropicRequest) SetMessages(v []AiAnthropicMessage)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *AiAnthropicRequest) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### GetModel

`func (o *AiAnthropicRequest) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *AiAnthropicRequest) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *AiAnthropicRequest) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *AiAnthropicRequest) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetStream

`func (o *AiAnthropicRequest) GetStream() bool`

GetStream returns the Stream field if non-nil, zero value otherwise.

### GetStreamOk

`func (o *AiAnthropicRequest) GetStreamOk() (*bool, bool)`

GetStreamOk returns a tuple with the Stream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStream

`func (o *AiAnthropicRequest) SetStream(v bool)`

SetStream sets Stream field to given value.

### HasStream

`func (o *AiAnthropicRequest) HasStream() bool`

HasStream returns a boolean if a field has been set.

### GetSystem

`func (o *AiAnthropicRequest) GetSystem() interface{}`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *AiAnthropicRequest) GetSystemOk() (*interface{}, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *AiAnthropicRequest) SetSystem(v interface{})`

SetSystem sets System field to given value.

### HasSystem

`func (o *AiAnthropicRequest) HasSystem() bool`

HasSystem returns a boolean if a field has been set.

### SetSystemNil

`func (o *AiAnthropicRequest) SetSystemNil(b bool)`

 SetSystemNil sets the value for System to be an explicit nil

### UnsetSystem
`func (o *AiAnthropicRequest) UnsetSystem()`

UnsetSystem ensures that no value is present for System, not even an explicit nil
### GetTemperature

`func (o *AiAnthropicRequest) GetTemperature() float32`

GetTemperature returns the Temperature field if non-nil, zero value otherwise.

### GetTemperatureOk

`func (o *AiAnthropicRequest) GetTemperatureOk() (*float32, bool)`

GetTemperatureOk returns a tuple with the Temperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperature

`func (o *AiAnthropicRequest) SetTemperature(v float32)`

SetTemperature sets Temperature field to given value.

### HasTemperature

`func (o *AiAnthropicRequest) HasTemperature() bool`

HasTemperature returns a boolean if a field has been set.

### GetThinking

`func (o *AiAnthropicRequest) GetThinking() interface{}`

GetThinking returns the Thinking field if non-nil, zero value otherwise.

### GetThinkingOk

`func (o *AiAnthropicRequest) GetThinkingOk() (*interface{}, bool)`

GetThinkingOk returns a tuple with the Thinking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThinking

`func (o *AiAnthropicRequest) SetThinking(v interface{})`

SetThinking sets Thinking field to given value.

### HasThinking

`func (o *AiAnthropicRequest) HasThinking() bool`

HasThinking returns a boolean if a field has been set.

### SetThinkingNil

`func (o *AiAnthropicRequest) SetThinkingNil(b bool)`

 SetThinkingNil sets the value for Thinking to be an explicit nil

### UnsetThinking
`func (o *AiAnthropicRequest) UnsetThinking()`

UnsetThinking ensures that no value is present for Thinking, not even an explicit nil
### GetToolChoice

`func (o *AiAnthropicRequest) GetToolChoice() interface{}`

GetToolChoice returns the ToolChoice field if non-nil, zero value otherwise.

### GetToolChoiceOk

`func (o *AiAnthropicRequest) GetToolChoiceOk() (*interface{}, bool)`

GetToolChoiceOk returns a tuple with the ToolChoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolChoice

`func (o *AiAnthropicRequest) SetToolChoice(v interface{})`

SetToolChoice sets ToolChoice field to given value.

### HasToolChoice

`func (o *AiAnthropicRequest) HasToolChoice() bool`

HasToolChoice returns a boolean if a field has been set.

### SetToolChoiceNil

`func (o *AiAnthropicRequest) SetToolChoiceNil(b bool)`

 SetToolChoiceNil sets the value for ToolChoice to be an explicit nil

### UnsetToolChoice
`func (o *AiAnthropicRequest) UnsetToolChoice()`

UnsetToolChoice ensures that no value is present for ToolChoice, not even an explicit nil
### GetTools

`func (o *AiAnthropicRequest) GetTools() []AiAnthropicTool`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *AiAnthropicRequest) GetToolsOk() (*[]AiAnthropicTool, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *AiAnthropicRequest) SetTools(v []AiAnthropicTool)`

SetTools sets Tools field to given value.

### HasTools

`func (o *AiAnthropicRequest) HasTools() bool`

HasTools returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


