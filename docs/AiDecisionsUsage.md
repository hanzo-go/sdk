# AiDecisionsUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cost** | Pointer to **float32** |  | [optional] 
**InputTokens** | **int32** |  | 
**OutputTokens** | **int32** |  | 

## Methods

### NewAiDecisionsUsage

`func NewAiDecisionsUsage(inputTokens int32, outputTokens int32, ) *AiDecisionsUsage`

NewAiDecisionsUsage instantiates a new AiDecisionsUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiDecisionsUsageWithDefaults

`func NewAiDecisionsUsageWithDefaults() *AiDecisionsUsage`

NewAiDecisionsUsageWithDefaults instantiates a new AiDecisionsUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCost

`func (o *AiDecisionsUsage) GetCost() float32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *AiDecisionsUsage) GetCostOk() (*float32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *AiDecisionsUsage) SetCost(v float32)`

SetCost sets Cost field to given value.

### HasCost

`func (o *AiDecisionsUsage) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetInputTokens

`func (o *AiDecisionsUsage) GetInputTokens() int32`

GetInputTokens returns the InputTokens field if non-nil, zero value otherwise.

### GetInputTokensOk

`func (o *AiDecisionsUsage) GetInputTokensOk() (*int32, bool)`

GetInputTokensOk returns a tuple with the InputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputTokens

`func (o *AiDecisionsUsage) SetInputTokens(v int32)`

SetInputTokens sets InputTokens field to given value.


### GetOutputTokens

`func (o *AiDecisionsUsage) GetOutputTokens() int32`

GetOutputTokens returns the OutputTokens field if non-nil, zero value otherwise.

### GetOutputTokensOk

`func (o *AiDecisionsUsage) GetOutputTokensOk() (*int32, bool)`

GetOutputTokensOk returns a tuple with the OutputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputTokens

`func (o *AiDecisionsUsage) SetOutputTokens(v int32)`

SetOutputTokens sets OutputTokens field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


