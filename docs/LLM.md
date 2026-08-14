# LLM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | Pointer to **bool** | Available is false when the warehouse was not connected or a query blipped. The totals below are then honest zeros, NOT measured ones. | [optional] 
**CompletionTokens** | Pointer to **int32** | CompletionTokens is the output half. | [optional] 
**CostCents** | Pointer to **int32** | CostCents is what they cost the org, in US cents. This IS a Hanzo charge. | [optional] 
**Models** | Pointer to **int32** | Models is how many distinct models were used. | [optional] 
**PromptTokens** | Pointer to **int32** | PromptTokens is the input half of that total. | [optional] 
**Requests** | Pointer to **int32** | Requests is how many completions the org made in the window. | [optional] 
**Source** | Pointer to **string** | Source names the warehouse table the totals came from. | [optional] 
**Tokens** | Pointer to **int32** | Tokens is the total tokens those completions consumed. | [optional] 

## Methods

### NewLLM

`func NewLLM() *LLM`

NewLLM instantiates a new LLM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLLMWithDefaults

`func NewLLMWithDefaults() *LLM`

NewLLMWithDefaults instantiates a new LLM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailable

`func (o *LLM) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *LLM) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *LLM) SetAvailable(v bool)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *LLM) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetCompletionTokens

`func (o *LLM) GetCompletionTokens() int32`

GetCompletionTokens returns the CompletionTokens field if non-nil, zero value otherwise.

### GetCompletionTokensOk

`func (o *LLM) GetCompletionTokensOk() (*int32, bool)`

GetCompletionTokensOk returns a tuple with the CompletionTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionTokens

`func (o *LLM) SetCompletionTokens(v int32)`

SetCompletionTokens sets CompletionTokens field to given value.

### HasCompletionTokens

`func (o *LLM) HasCompletionTokens() bool`

HasCompletionTokens returns a boolean if a field has been set.

### GetCostCents

`func (o *LLM) GetCostCents() int32`

GetCostCents returns the CostCents field if non-nil, zero value otherwise.

### GetCostCentsOk

`func (o *LLM) GetCostCentsOk() (*int32, bool)`

GetCostCentsOk returns a tuple with the CostCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostCents

`func (o *LLM) SetCostCents(v int32)`

SetCostCents sets CostCents field to given value.

### HasCostCents

`func (o *LLM) HasCostCents() bool`

HasCostCents returns a boolean if a field has been set.

### GetModels

`func (o *LLM) GetModels() int32`

GetModels returns the Models field if non-nil, zero value otherwise.

### GetModelsOk

`func (o *LLM) GetModelsOk() (*int32, bool)`

GetModelsOk returns a tuple with the Models field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModels

`func (o *LLM) SetModels(v int32)`

SetModels sets Models field to given value.

### HasModels

`func (o *LLM) HasModels() bool`

HasModels returns a boolean if a field has been set.

### GetPromptTokens

`func (o *LLM) GetPromptTokens() int32`

GetPromptTokens returns the PromptTokens field if non-nil, zero value otherwise.

### GetPromptTokensOk

`func (o *LLM) GetPromptTokensOk() (*int32, bool)`

GetPromptTokensOk returns a tuple with the PromptTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptTokens

`func (o *LLM) SetPromptTokens(v int32)`

SetPromptTokens sets PromptTokens field to given value.

### HasPromptTokens

`func (o *LLM) HasPromptTokens() bool`

HasPromptTokens returns a boolean if a field has been set.

### GetRequests

`func (o *LLM) GetRequests() int32`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *LLM) GetRequestsOk() (*int32, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *LLM) SetRequests(v int32)`

SetRequests sets Requests field to given value.

### HasRequests

`func (o *LLM) HasRequests() bool`

HasRequests returns a boolean if a field has been set.

### GetSource

`func (o *LLM) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *LLM) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *LLM) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *LLM) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetTokens

`func (o *LLM) GetTokens() int32`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *LLM) GetTokensOk() (*int32, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *LLM) SetTokens(v int32)`

SetTokens sets Tokens field to given value.

### HasTokens

`func (o *LLM) HasTokens() bool`

HasTokens returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


