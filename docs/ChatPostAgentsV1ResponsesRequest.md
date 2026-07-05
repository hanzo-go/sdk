# ChatPostAgentsV1ResponsesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Model** | **string** | Agent ID | 
**Input** | [**ChatPostAgentsV1ResponsesRequestInput**](ChatPostAgentsV1ResponsesRequestInput.md) |  | 
**Stream** | Pointer to **bool** |  | [optional] [default to false]
**PreviousResponseId** | Pointer to **string** |  | [optional] 
**Instructions** | Pointer to **string** |  | [optional] 
**Tools** | Pointer to **[]map[string]interface{}** |  | [optional] 
**ToolChoice** | Pointer to **string** |  | [optional] 
**MaxOutputTokens** | Pointer to **int32** |  | [optional] 
**Temperature** | Pointer to **float32** |  | [optional] 

## Methods

### NewChatPostAgentsV1ResponsesRequest

`func NewChatPostAgentsV1ResponsesRequest(model string, input ChatPostAgentsV1ResponsesRequestInput, ) *ChatPostAgentsV1ResponsesRequest`

NewChatPostAgentsV1ResponsesRequest instantiates a new ChatPostAgentsV1ResponsesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatPostAgentsV1ResponsesRequestWithDefaults

`func NewChatPostAgentsV1ResponsesRequestWithDefaults() *ChatPostAgentsV1ResponsesRequest`

NewChatPostAgentsV1ResponsesRequestWithDefaults instantiates a new ChatPostAgentsV1ResponsesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModel

`func (o *ChatPostAgentsV1ResponsesRequest) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ChatPostAgentsV1ResponsesRequest) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ChatPostAgentsV1ResponsesRequest) SetModel(v string)`

SetModel sets Model field to given value.


### GetInput

`func (o *ChatPostAgentsV1ResponsesRequest) GetInput() ChatPostAgentsV1ResponsesRequestInput`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *ChatPostAgentsV1ResponsesRequest) GetInputOk() (*ChatPostAgentsV1ResponsesRequestInput, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *ChatPostAgentsV1ResponsesRequest) SetInput(v ChatPostAgentsV1ResponsesRequestInput)`

SetInput sets Input field to given value.


### GetStream

`func (o *ChatPostAgentsV1ResponsesRequest) GetStream() bool`

GetStream returns the Stream field if non-nil, zero value otherwise.

### GetStreamOk

`func (o *ChatPostAgentsV1ResponsesRequest) GetStreamOk() (*bool, bool)`

GetStreamOk returns a tuple with the Stream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStream

`func (o *ChatPostAgentsV1ResponsesRequest) SetStream(v bool)`

SetStream sets Stream field to given value.

### HasStream

`func (o *ChatPostAgentsV1ResponsesRequest) HasStream() bool`

HasStream returns a boolean if a field has been set.

### GetPreviousResponseId

`func (o *ChatPostAgentsV1ResponsesRequest) GetPreviousResponseId() string`

GetPreviousResponseId returns the PreviousResponseId field if non-nil, zero value otherwise.

### GetPreviousResponseIdOk

`func (o *ChatPostAgentsV1ResponsesRequest) GetPreviousResponseIdOk() (*string, bool)`

GetPreviousResponseIdOk returns a tuple with the PreviousResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousResponseId

`func (o *ChatPostAgentsV1ResponsesRequest) SetPreviousResponseId(v string)`

SetPreviousResponseId sets PreviousResponseId field to given value.

### HasPreviousResponseId

`func (o *ChatPostAgentsV1ResponsesRequest) HasPreviousResponseId() bool`

HasPreviousResponseId returns a boolean if a field has been set.

### GetInstructions

`func (o *ChatPostAgentsV1ResponsesRequest) GetInstructions() string`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *ChatPostAgentsV1ResponsesRequest) GetInstructionsOk() (*string, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *ChatPostAgentsV1ResponsesRequest) SetInstructions(v string)`

SetInstructions sets Instructions field to given value.

### HasInstructions

`func (o *ChatPostAgentsV1ResponsesRequest) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.

### GetTools

`func (o *ChatPostAgentsV1ResponsesRequest) GetTools() []map[string]interface{}`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *ChatPostAgentsV1ResponsesRequest) GetToolsOk() (*[]map[string]interface{}, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *ChatPostAgentsV1ResponsesRequest) SetTools(v []map[string]interface{})`

SetTools sets Tools field to given value.

### HasTools

`func (o *ChatPostAgentsV1ResponsesRequest) HasTools() bool`

HasTools returns a boolean if a field has been set.

### GetToolChoice

`func (o *ChatPostAgentsV1ResponsesRequest) GetToolChoice() string`

GetToolChoice returns the ToolChoice field if non-nil, zero value otherwise.

### GetToolChoiceOk

`func (o *ChatPostAgentsV1ResponsesRequest) GetToolChoiceOk() (*string, bool)`

GetToolChoiceOk returns a tuple with the ToolChoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolChoice

`func (o *ChatPostAgentsV1ResponsesRequest) SetToolChoice(v string)`

SetToolChoice sets ToolChoice field to given value.

### HasToolChoice

`func (o *ChatPostAgentsV1ResponsesRequest) HasToolChoice() bool`

HasToolChoice returns a boolean if a field has been set.

### GetMaxOutputTokens

`func (o *ChatPostAgentsV1ResponsesRequest) GetMaxOutputTokens() int32`

GetMaxOutputTokens returns the MaxOutputTokens field if non-nil, zero value otherwise.

### GetMaxOutputTokensOk

`func (o *ChatPostAgentsV1ResponsesRequest) GetMaxOutputTokensOk() (*int32, bool)`

GetMaxOutputTokensOk returns a tuple with the MaxOutputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxOutputTokens

`func (o *ChatPostAgentsV1ResponsesRequest) SetMaxOutputTokens(v int32)`

SetMaxOutputTokens sets MaxOutputTokens field to given value.

### HasMaxOutputTokens

`func (o *ChatPostAgentsV1ResponsesRequest) HasMaxOutputTokens() bool`

HasMaxOutputTokens returns a boolean if a field has been set.

### GetTemperature

`func (o *ChatPostAgentsV1ResponsesRequest) GetTemperature() float32`

GetTemperature returns the Temperature field if non-nil, zero value otherwise.

### GetTemperatureOk

`func (o *ChatPostAgentsV1ResponsesRequest) GetTemperatureOk() (*float32, bool)`

GetTemperatureOk returns a tuple with the Temperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperature

`func (o *ChatPostAgentsV1ResponsesRequest) SetTemperature(v float32)`

SetTemperature sets Temperature field to given value.

### HasTemperature

`func (o *ChatPostAgentsV1ResponsesRequest) HasTemperature() bool`

HasTemperature returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


