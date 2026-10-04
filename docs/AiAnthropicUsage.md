# AiAnthropicUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CacheCreation** | Pointer to [**AiCacheWrites**](AiCacheWrites.md) |  | [optional] 
**CacheCreationInputTokens** | Pointer to **int32** |  | [optional] 
**CacheReadInputTokens** | Pointer to **int32** |  | [optional] 
**InputTokens** | Pointer to **int32** |  | [optional] 
**Iterations** | Pointer to [**[]AiAnthropicUsage**](AiAnthropicUsage.md) |  | [optional] 
**OutputTokens** | Pointer to **int32** |  | [optional] 

## Methods

### NewAiAnthropicUsage

`func NewAiAnthropicUsage() *AiAnthropicUsage`

NewAiAnthropicUsage instantiates a new AiAnthropicUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiAnthropicUsageWithDefaults

`func NewAiAnthropicUsageWithDefaults() *AiAnthropicUsage`

NewAiAnthropicUsageWithDefaults instantiates a new AiAnthropicUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCacheCreation

`func (o *AiAnthropicUsage) GetCacheCreation() AiCacheWrites`

GetCacheCreation returns the CacheCreation field if non-nil, zero value otherwise.

### GetCacheCreationOk

`func (o *AiAnthropicUsage) GetCacheCreationOk() (*AiCacheWrites, bool)`

GetCacheCreationOk returns a tuple with the CacheCreation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheCreation

`func (o *AiAnthropicUsage) SetCacheCreation(v AiCacheWrites)`

SetCacheCreation sets CacheCreation field to given value.

### HasCacheCreation

`func (o *AiAnthropicUsage) HasCacheCreation() bool`

HasCacheCreation returns a boolean if a field has been set.

### GetCacheCreationInputTokens

`func (o *AiAnthropicUsage) GetCacheCreationInputTokens() int32`

GetCacheCreationInputTokens returns the CacheCreationInputTokens field if non-nil, zero value otherwise.

### GetCacheCreationInputTokensOk

`func (o *AiAnthropicUsage) GetCacheCreationInputTokensOk() (*int32, bool)`

GetCacheCreationInputTokensOk returns a tuple with the CacheCreationInputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheCreationInputTokens

`func (o *AiAnthropicUsage) SetCacheCreationInputTokens(v int32)`

SetCacheCreationInputTokens sets CacheCreationInputTokens field to given value.

### HasCacheCreationInputTokens

`func (o *AiAnthropicUsage) HasCacheCreationInputTokens() bool`

HasCacheCreationInputTokens returns a boolean if a field has been set.

### GetCacheReadInputTokens

`func (o *AiAnthropicUsage) GetCacheReadInputTokens() int32`

GetCacheReadInputTokens returns the CacheReadInputTokens field if non-nil, zero value otherwise.

### GetCacheReadInputTokensOk

`func (o *AiAnthropicUsage) GetCacheReadInputTokensOk() (*int32, bool)`

GetCacheReadInputTokensOk returns a tuple with the CacheReadInputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheReadInputTokens

`func (o *AiAnthropicUsage) SetCacheReadInputTokens(v int32)`

SetCacheReadInputTokens sets CacheReadInputTokens field to given value.

### HasCacheReadInputTokens

`func (o *AiAnthropicUsage) HasCacheReadInputTokens() bool`

HasCacheReadInputTokens returns a boolean if a field has been set.

### GetInputTokens

`func (o *AiAnthropicUsage) GetInputTokens() int32`

GetInputTokens returns the InputTokens field if non-nil, zero value otherwise.

### GetInputTokensOk

`func (o *AiAnthropicUsage) GetInputTokensOk() (*int32, bool)`

GetInputTokensOk returns a tuple with the InputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputTokens

`func (o *AiAnthropicUsage) SetInputTokens(v int32)`

SetInputTokens sets InputTokens field to given value.

### HasInputTokens

`func (o *AiAnthropicUsage) HasInputTokens() bool`

HasInputTokens returns a boolean if a field has been set.

### GetIterations

`func (o *AiAnthropicUsage) GetIterations() []AiAnthropicUsage`

GetIterations returns the Iterations field if non-nil, zero value otherwise.

### GetIterationsOk

`func (o *AiAnthropicUsage) GetIterationsOk() (*[]AiAnthropicUsage, bool)`

GetIterationsOk returns a tuple with the Iterations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIterations

`func (o *AiAnthropicUsage) SetIterations(v []AiAnthropicUsage)`

SetIterations sets Iterations field to given value.

### HasIterations

`func (o *AiAnthropicUsage) HasIterations() bool`

HasIterations returns a boolean if a field has been set.

### GetOutputTokens

`func (o *AiAnthropicUsage) GetOutputTokens() int32`

GetOutputTokens returns the OutputTokens field if non-nil, zero value otherwise.

### GetOutputTokensOk

`func (o *AiAnthropicUsage) GetOutputTokensOk() (*int32, bool)`

GetOutputTokensOk returns a tuple with the OutputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputTokens

`func (o *AiAnthropicUsage) SetOutputTokens(v int32)`

SetOutputTokens sets OutputTokens field to given value.

### HasOutputTokens

`func (o *AiAnthropicUsage) HasOutputTokens() bool`

HasOutputTokens returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


