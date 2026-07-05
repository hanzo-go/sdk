# GatewayCreateCompletionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Model** | **string** |  | 
**Prompt** | [**GatewayCreateCompletionRequestPrompt**](GatewayCreateCompletionRequestPrompt.md) |  | 
**MaxTokens** | Pointer to **int32** |  | [optional] 
**Temperature** | Pointer to **float32** |  | [optional] 
**Stream** | Pointer to **bool** |  | [optional] 

## Methods

### NewGatewayCreateCompletionRequest

`func NewGatewayCreateCompletionRequest(model string, prompt GatewayCreateCompletionRequestPrompt, ) *GatewayCreateCompletionRequest`

NewGatewayCreateCompletionRequest instantiates a new GatewayCreateCompletionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayCreateCompletionRequestWithDefaults

`func NewGatewayCreateCompletionRequestWithDefaults() *GatewayCreateCompletionRequest`

NewGatewayCreateCompletionRequestWithDefaults instantiates a new GatewayCreateCompletionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModel

`func (o *GatewayCreateCompletionRequest) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *GatewayCreateCompletionRequest) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *GatewayCreateCompletionRequest) SetModel(v string)`

SetModel sets Model field to given value.


### GetPrompt

`func (o *GatewayCreateCompletionRequest) GetPrompt() GatewayCreateCompletionRequestPrompt`

GetPrompt returns the Prompt field if non-nil, zero value otherwise.

### GetPromptOk

`func (o *GatewayCreateCompletionRequest) GetPromptOk() (*GatewayCreateCompletionRequestPrompt, bool)`

GetPromptOk returns a tuple with the Prompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompt

`func (o *GatewayCreateCompletionRequest) SetPrompt(v GatewayCreateCompletionRequestPrompt)`

SetPrompt sets Prompt field to given value.


### GetMaxTokens

`func (o *GatewayCreateCompletionRequest) GetMaxTokens() int32`

GetMaxTokens returns the MaxTokens field if non-nil, zero value otherwise.

### GetMaxTokensOk

`func (o *GatewayCreateCompletionRequest) GetMaxTokensOk() (*int32, bool)`

GetMaxTokensOk returns a tuple with the MaxTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTokens

`func (o *GatewayCreateCompletionRequest) SetMaxTokens(v int32)`

SetMaxTokens sets MaxTokens field to given value.

### HasMaxTokens

`func (o *GatewayCreateCompletionRequest) HasMaxTokens() bool`

HasMaxTokens returns a boolean if a field has been set.

### GetTemperature

`func (o *GatewayCreateCompletionRequest) GetTemperature() float32`

GetTemperature returns the Temperature field if non-nil, zero value otherwise.

### GetTemperatureOk

`func (o *GatewayCreateCompletionRequest) GetTemperatureOk() (*float32, bool)`

GetTemperatureOk returns a tuple with the Temperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperature

`func (o *GatewayCreateCompletionRequest) SetTemperature(v float32)`

SetTemperature sets Temperature field to given value.

### HasTemperature

`func (o *GatewayCreateCompletionRequest) HasTemperature() bool`

HasTemperature returns a boolean if a field has been set.

### GetStream

`func (o *GatewayCreateCompletionRequest) GetStream() bool`

GetStream returns the Stream field if non-nil, zero value otherwise.

### GetStreamOk

`func (o *GatewayCreateCompletionRequest) GetStreamOk() (*bool, bool)`

GetStreamOk returns a tuple with the Stream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStream

`func (o *GatewayCreateCompletionRequest) SetStream(v bool)`

SetStream sets Stream field to given value.

### HasStream

`func (o *GatewayCreateCompletionRequest) HasStream() bool`

HasStream returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


