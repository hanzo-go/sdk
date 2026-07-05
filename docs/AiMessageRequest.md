# AiMessageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Model** | **string** |  | 
**Messages** | **[]map[string]interface{}** |  | 
**MaxTokens** | **int32** |  | 
**Stream** | Pointer to **bool** |  | [optional] [default to false]
**System** | Pointer to **interface{}** |  | [optional] 
**Temperature** | Pointer to **float32** |  | [optional] 
**Tools** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewAiMessageRequest

`func NewAiMessageRequest(model string, messages []map[string]interface{}, maxTokens int32, ) *AiMessageRequest`

NewAiMessageRequest instantiates a new AiMessageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiMessageRequestWithDefaults

`func NewAiMessageRequestWithDefaults() *AiMessageRequest`

NewAiMessageRequestWithDefaults instantiates a new AiMessageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModel

`func (o *AiMessageRequest) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *AiMessageRequest) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *AiMessageRequest) SetModel(v string)`

SetModel sets Model field to given value.


### GetMessages

`func (o *AiMessageRequest) GetMessages() []map[string]interface{}`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *AiMessageRequest) GetMessagesOk() (*[]map[string]interface{}, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *AiMessageRequest) SetMessages(v []map[string]interface{})`

SetMessages sets Messages field to given value.


### GetMaxTokens

`func (o *AiMessageRequest) GetMaxTokens() int32`

GetMaxTokens returns the MaxTokens field if non-nil, zero value otherwise.

### GetMaxTokensOk

`func (o *AiMessageRequest) GetMaxTokensOk() (*int32, bool)`

GetMaxTokensOk returns a tuple with the MaxTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTokens

`func (o *AiMessageRequest) SetMaxTokens(v int32)`

SetMaxTokens sets MaxTokens field to given value.


### GetStream

`func (o *AiMessageRequest) GetStream() bool`

GetStream returns the Stream field if non-nil, zero value otherwise.

### GetStreamOk

`func (o *AiMessageRequest) GetStreamOk() (*bool, bool)`

GetStreamOk returns a tuple with the Stream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStream

`func (o *AiMessageRequest) SetStream(v bool)`

SetStream sets Stream field to given value.

### HasStream

`func (o *AiMessageRequest) HasStream() bool`

HasStream returns a boolean if a field has been set.

### GetSystem

`func (o *AiMessageRequest) GetSystem() interface{}`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *AiMessageRequest) GetSystemOk() (*interface{}, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *AiMessageRequest) SetSystem(v interface{})`

SetSystem sets System field to given value.

### HasSystem

`func (o *AiMessageRequest) HasSystem() bool`

HasSystem returns a boolean if a field has been set.

### SetSystemNil

`func (o *AiMessageRequest) SetSystemNil(b bool)`

 SetSystemNil sets the value for System to be an explicit nil

### UnsetSystem
`func (o *AiMessageRequest) UnsetSystem()`

UnsetSystem ensures that no value is present for System, not even an explicit nil
### GetTemperature

`func (o *AiMessageRequest) GetTemperature() float32`

GetTemperature returns the Temperature field if non-nil, zero value otherwise.

### GetTemperatureOk

`func (o *AiMessageRequest) GetTemperatureOk() (*float32, bool)`

GetTemperatureOk returns a tuple with the Temperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperature

`func (o *AiMessageRequest) SetTemperature(v float32)`

SetTemperature sets Temperature field to given value.

### HasTemperature

`func (o *AiMessageRequest) HasTemperature() bool`

HasTemperature returns a boolean if a field has been set.

### GetTools

`func (o *AiMessageRequest) GetTools() []map[string]interface{}`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *AiMessageRequest) GetToolsOk() (*[]map[string]interface{}, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *AiMessageRequest) SetTools(v []map[string]interface{})`

SetTools sets Tools field to given value.

### HasTools

`func (o *AiMessageRequest) HasTools() bool`

HasTools returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


