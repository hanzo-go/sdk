# UsageLLM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | Pointer to **bool** | Available is false when the warehouse was not connected or a query blipped. The totals below are then honest zeros, NOT measured ones. | [optional] 
**CacheReadTokens** | Pointer to **int64** | CacheReadTokens is how many of the prompt tokens the provider&#39;s prompt cache served, at the cache-read price. | [optional] 
**CacheSavedCents** | Pointer to **int64** | CacheSavedCents is what the prompt cache took off the bill, in US cents: each call&#39;s cache-read tokens at the model&#39;s input price less the same tokens at its cache-read price, from the price that call was billed at. A call recorded before this was measured counts 0, so it is a floor. | [optional] 
**CacheWriteTokens** | Pointer to **int64** | CacheWriteTokens is how many prompt tokens were written to the prompt cache. | [optional] 
**CachedShare** | Pointer to **float64** | CachedShare is CacheReadTokens over PromptTokens, 0 to 1 at four decimals; 0 when nothing was prompted. | [optional] 
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

### GetCacheReadTokens

`func (o *UsageLLM) GetCacheReadTokens() int64`

GetCacheReadTokens returns the CacheReadTokens field if non-nil, zero value otherwise.

### GetCacheReadTokensOk

`func (o *UsageLLM) GetCacheReadTokensOk() (*int64, bool)`

GetCacheReadTokensOk returns a tuple with the CacheReadTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheReadTokens

`func (o *UsageLLM) SetCacheReadTokens(v int64)`

SetCacheReadTokens sets CacheReadTokens field to given value.

### HasCacheReadTokens

`func (o *UsageLLM) HasCacheReadTokens() bool`

HasCacheReadTokens returns a boolean if a field has been set.

### GetCacheSavedCents

`func (o *UsageLLM) GetCacheSavedCents() int64`

GetCacheSavedCents returns the CacheSavedCents field if non-nil, zero value otherwise.

### GetCacheSavedCentsOk

`func (o *UsageLLM) GetCacheSavedCentsOk() (*int64, bool)`

GetCacheSavedCentsOk returns a tuple with the CacheSavedCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheSavedCents

`func (o *UsageLLM) SetCacheSavedCents(v int64)`

SetCacheSavedCents sets CacheSavedCents field to given value.

### HasCacheSavedCents

`func (o *UsageLLM) HasCacheSavedCents() bool`

HasCacheSavedCents returns a boolean if a field has been set.

### GetCacheWriteTokens

`func (o *UsageLLM) GetCacheWriteTokens() int64`

GetCacheWriteTokens returns the CacheWriteTokens field if non-nil, zero value otherwise.

### GetCacheWriteTokensOk

`func (o *UsageLLM) GetCacheWriteTokensOk() (*int64, bool)`

GetCacheWriteTokensOk returns a tuple with the CacheWriteTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheWriteTokens

`func (o *UsageLLM) SetCacheWriteTokens(v int64)`

SetCacheWriteTokens sets CacheWriteTokens field to given value.

### HasCacheWriteTokens

`func (o *UsageLLM) HasCacheWriteTokens() bool`

HasCacheWriteTokens returns a boolean if a field has been set.

### GetCachedShare

`func (o *UsageLLM) GetCachedShare() float64`

GetCachedShare returns the CachedShare field if non-nil, zero value otherwise.

### GetCachedShareOk

`func (o *UsageLLM) GetCachedShareOk() (*float64, bool)`

GetCachedShareOk returns a tuple with the CachedShare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCachedShare

`func (o *UsageLLM) SetCachedShare(v float64)`

SetCachedShare sets CachedShare field to given value.

### HasCachedShare

`func (o *UsageLLM) HasCachedShare() bool`

HasCachedShare returns a boolean if a field has been set.

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


