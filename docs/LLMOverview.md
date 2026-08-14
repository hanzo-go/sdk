# LLMOverview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | Pointer to **bool** | Available is true whenever the ledger answered — including with no usage in the window, which is honest zeros rather than a missing lens. | [optional] 
**CompletionTokens** | Pointer to **int32** | CompletionTokens is the output half of Tokens. | [optional] 
**ErrorRate** | Pointer to **float32** | ErrorRate is Errors/Requests, 0..1, rounded to three places. Zero when there were no requests. | [optional] 
**Errors** | Pointer to **int32** | Errors is how many of Requests failed. | [optional] 
**Models** | Pointer to **int32** | Models is how many distinct models the org called. | [optional] 
**PromptTokens** | Pointer to **int32** | PromptTokens is the input half of Tokens. | [optional] 
**Providers** | Pointer to **int32** | Providers is how many distinct providers served them. | [optional] 
**Requests** | Pointer to **int32** | Requests is how many LLM calls the org made in the window. | [optional] 
**Source** | Pointer to **string** | Source is the warehouse table the lens read. | [optional] 
**SpendCents** | Pointer to **int32** | SpendCents is what those calls cost, in cents. | [optional] 
**Tokens** | Pointer to **int32** | Tokens is prompt plus completion tokens over those calls. | [optional] 

## Methods

### NewLLMOverview

`func NewLLMOverview() *LLMOverview`

NewLLMOverview instantiates a new LLMOverview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLLMOverviewWithDefaults

`func NewLLMOverviewWithDefaults() *LLMOverview`

NewLLMOverviewWithDefaults instantiates a new LLMOverview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailable

`func (o *LLMOverview) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *LLMOverview) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *LLMOverview) SetAvailable(v bool)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *LLMOverview) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetCompletionTokens

`func (o *LLMOverview) GetCompletionTokens() int32`

GetCompletionTokens returns the CompletionTokens field if non-nil, zero value otherwise.

### GetCompletionTokensOk

`func (o *LLMOverview) GetCompletionTokensOk() (*int32, bool)`

GetCompletionTokensOk returns a tuple with the CompletionTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionTokens

`func (o *LLMOverview) SetCompletionTokens(v int32)`

SetCompletionTokens sets CompletionTokens field to given value.

### HasCompletionTokens

`func (o *LLMOverview) HasCompletionTokens() bool`

HasCompletionTokens returns a boolean if a field has been set.

### GetErrorRate

`func (o *LLMOverview) GetErrorRate() float32`

GetErrorRate returns the ErrorRate field if non-nil, zero value otherwise.

### GetErrorRateOk

`func (o *LLMOverview) GetErrorRateOk() (*float32, bool)`

GetErrorRateOk returns a tuple with the ErrorRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorRate

`func (o *LLMOverview) SetErrorRate(v float32)`

SetErrorRate sets ErrorRate field to given value.

### HasErrorRate

`func (o *LLMOverview) HasErrorRate() bool`

HasErrorRate returns a boolean if a field has been set.

### GetErrors

`func (o *LLMOverview) GetErrors() int32`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *LLMOverview) GetErrorsOk() (*int32, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *LLMOverview) SetErrors(v int32)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *LLMOverview) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetModels

`func (o *LLMOverview) GetModels() int32`

GetModels returns the Models field if non-nil, zero value otherwise.

### GetModelsOk

`func (o *LLMOverview) GetModelsOk() (*int32, bool)`

GetModelsOk returns a tuple with the Models field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModels

`func (o *LLMOverview) SetModels(v int32)`

SetModels sets Models field to given value.

### HasModels

`func (o *LLMOverview) HasModels() bool`

HasModels returns a boolean if a field has been set.

### GetPromptTokens

`func (o *LLMOverview) GetPromptTokens() int32`

GetPromptTokens returns the PromptTokens field if non-nil, zero value otherwise.

### GetPromptTokensOk

`func (o *LLMOverview) GetPromptTokensOk() (*int32, bool)`

GetPromptTokensOk returns a tuple with the PromptTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptTokens

`func (o *LLMOverview) SetPromptTokens(v int32)`

SetPromptTokens sets PromptTokens field to given value.

### HasPromptTokens

`func (o *LLMOverview) HasPromptTokens() bool`

HasPromptTokens returns a boolean if a field has been set.

### GetProviders

`func (o *LLMOverview) GetProviders() int32`

GetProviders returns the Providers field if non-nil, zero value otherwise.

### GetProvidersOk

`func (o *LLMOverview) GetProvidersOk() (*int32, bool)`

GetProvidersOk returns a tuple with the Providers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviders

`func (o *LLMOverview) SetProviders(v int32)`

SetProviders sets Providers field to given value.

### HasProviders

`func (o *LLMOverview) HasProviders() bool`

HasProviders returns a boolean if a field has been set.

### GetRequests

`func (o *LLMOverview) GetRequests() int32`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *LLMOverview) GetRequestsOk() (*int32, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *LLMOverview) SetRequests(v int32)`

SetRequests sets Requests field to given value.

### HasRequests

`func (o *LLMOverview) HasRequests() bool`

HasRequests returns a boolean if a field has been set.

### GetSource

`func (o *LLMOverview) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *LLMOverview) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *LLMOverview) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *LLMOverview) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSpendCents

`func (o *LLMOverview) GetSpendCents() int32`

GetSpendCents returns the SpendCents field if non-nil, zero value otherwise.

### GetSpendCentsOk

`func (o *LLMOverview) GetSpendCentsOk() (*int32, bool)`

GetSpendCentsOk returns a tuple with the SpendCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpendCents

`func (o *LLMOverview) SetSpendCents(v int32)`

SetSpendCents sets SpendCents field to given value.

### HasSpendCents

`func (o *LLMOverview) HasSpendCents() bool`

HasSpendCents returns a boolean if a field has been set.

### GetTokens

`func (o *LLMOverview) GetTokens() int32`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *LLMOverview) GetTokensOk() (*int32, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *LLMOverview) SetTokens(v int32)`

SetTokens sets Tokens field to given value.

### HasTokens

`func (o *LLMOverview) HasTokens() bool`

HasTokens returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


