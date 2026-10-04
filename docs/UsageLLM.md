# UsageLLM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | Pointer to **bool** | Available is false when the warehouse was not connected or a query blipped. The totals below are then honest zeros, NOT measured ones. | [optional] 
**CompletionTokens** | Pointer to **int64** | CompletionTokens is the output half. | [optional] 
**CostCents** | Pointer to **int64** | CostCents is what they cost the org, in US cents. This IS a Hanzo charge. | [optional] 
**Models** | Pointer to **int64** | Models is how many distinct models were used. | [optional] 
**PromptTokens** | Pointer to **int64** | PromptTokens is the input half of that total. | [optional] 
**Requests** | Pointer to **int64** | Requests is how many completions the org made in the window. | [optional] 
**Source** | Pointer to **string** | Source names the warehouse table the totals came from. | [optional] 
**Tokens** | Pointer to **int64** | Tokens is the total tokens those completions consumed. | [optional] 

## Methods

### NewUsageLLM

`func NewUsageLLM() *UsageLLM`

NewUsageLLM instantiates a new UsageLLM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageLLMWithDefaults

`func NewUsageLLMWithDefaults() *UsageLLM`

NewUsageLLMWithDefaults instantiates a new UsageLLM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailable

`func (o *UsageLLM) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *UsageLLM) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *UsageLLM) SetAvailable(v bool)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *UsageLLM) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetCompletionTokens

`func (o *UsageLLM) GetCompletionTokens() int64`

GetCompletionTokens returns the CompletionTokens field if non-nil, zero value otherwise.

### GetCompletionTokensOk

`func (o *UsageLLM) GetCompletionTokensOk() (*int64, bool)`

GetCompletionTokensOk returns a tuple with the CompletionTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionTokens

`func (o *UsageLLM) SetCompletionTokens(v int64)`

SetCompletionTokens sets CompletionTokens field to given value.

### HasCompletionTokens

`func (o *UsageLLM) HasCompletionTokens() bool`

HasCompletionTokens returns a boolean if a field has been set.

### GetCostCents

`func (o *UsageLLM) GetCostCents() int64`

GetCostCents returns the CostCents field if non-nil, zero value otherwise.

### GetCostCentsOk

`func (o *UsageLLM) GetCostCentsOk() (*int64, bool)`

GetCostCentsOk returns a tuple with the CostCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostCents

`func (o *UsageLLM) SetCostCents(v int64)`

SetCostCents sets CostCents field to given value.

### HasCostCents

`func (o *UsageLLM) HasCostCents() bool`

HasCostCents returns a boolean if a field has been set.

### GetModels

`func (o *UsageLLM) GetModels() int64`

GetModels returns the Models field if non-nil, zero value otherwise.

### GetModelsOk

`func (o *UsageLLM) GetModelsOk() (*int64, bool)`

GetModelsOk returns a tuple with the Models field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModels

`func (o *UsageLLM) SetModels(v int64)`

SetModels sets Models field to given value.

### HasModels

`func (o *UsageLLM) HasModels() bool`

HasModels returns a boolean if a field has been set.

### GetPromptTokens

`func (o *UsageLLM) GetPromptTokens() int64`

GetPromptTokens returns the PromptTokens field if non-nil, zero value otherwise.

### GetPromptTokensOk

`func (o *UsageLLM) GetPromptTokensOk() (*int64, bool)`

GetPromptTokensOk returns a tuple with the PromptTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptTokens

`func (o *UsageLLM) SetPromptTokens(v int64)`

SetPromptTokens sets PromptTokens field to given value.

### HasPromptTokens

`func (o *UsageLLM) HasPromptTokens() bool`

HasPromptTokens returns a boolean if a field has been set.

### GetRequests

`func (o *UsageLLM) GetRequests() int64`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *UsageLLM) GetRequestsOk() (*int64, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *UsageLLM) SetRequests(v int64)`

SetRequests sets Requests field to given value.

### HasRequests

`func (o *UsageLLM) HasRequests() bool`

HasRequests returns a boolean if a field has been set.

### GetSource

`func (o *UsageLLM) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *UsageLLM) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *UsageLLM) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *UsageLLM) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetTokens

`func (o *UsageLLM) GetTokens() int64`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *UsageLLM) GetTokensOk() (*int64, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *UsageLLM) SetTokens(v int64)`

SetTokens sets Tokens field to given value.

### HasTokens

`func (o *UsageLLM) HasTokens() bool`

HasTokens returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


