# EventLLMOverview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | Pointer to **bool** | Available is true whenever the ledger answered — including with no usage in the window, which is honest zeros rather than a missing lens. | [optional] 
**CompletionTokens** | Pointer to **int64** | CompletionTokens is the output half of Tokens. | [optional] 
**ErrorRate** | Pointer to **float64** | ErrorRate is Errors/Requests, 0..1, rounded to three places. Zero when there were no requests. | [optional] 
**Errors** | Pointer to **int64** | Errors is how many of Requests failed. | [optional] 
**Models** | Pointer to **int64** | Models is how many distinct models the org called. | [optional] 
**PromptTokens** | Pointer to **int64** | PromptTokens is the input half of Tokens. | [optional] 
**Providers** | Pointer to **int64** | Providers is how many distinct providers served them. | [optional] 
**Requests** | Pointer to **int64** | Requests is how many LLM calls the org made in the window. | [optional] 
**Source** | Pointer to **string** | Source is the warehouse table the lens read. | [optional] 
**SpendCents** | Pointer to **int64** | SpendCents is what those calls cost, in cents. | [optional] 
**Tokens** | Pointer to **int64** | Tokens is prompt plus completion tokens over those calls. | [optional] 

## Methods

### NewEventLLMOverview

`func NewEventLLMOverview() *EventLLMOverview`

NewEventLLMOverview instantiates a new EventLLMOverview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventLLMOverviewWithDefaults

`func NewEventLLMOverviewWithDefaults() *EventLLMOverview`

NewEventLLMOverviewWithDefaults instantiates a new EventLLMOverview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailable

`func (o *EventLLMOverview) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *EventLLMOverview) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *EventLLMOverview) SetAvailable(v bool)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *EventLLMOverview) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetCompletionTokens

`func (o *EventLLMOverview) GetCompletionTokens() int64`

GetCompletionTokens returns the CompletionTokens field if non-nil, zero value otherwise.

### GetCompletionTokensOk

`func (o *EventLLMOverview) GetCompletionTokensOk() (*int64, bool)`

GetCompletionTokensOk returns a tuple with the CompletionTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionTokens

`func (o *EventLLMOverview) SetCompletionTokens(v int64)`

SetCompletionTokens sets CompletionTokens field to given value.

### HasCompletionTokens

`func (o *EventLLMOverview) HasCompletionTokens() bool`

HasCompletionTokens returns a boolean if a field has been set.

### GetErrorRate

`func (o *EventLLMOverview) GetErrorRate() float64`

GetErrorRate returns the ErrorRate field if non-nil, zero value otherwise.

### GetErrorRateOk

`func (o *EventLLMOverview) GetErrorRateOk() (*float64, bool)`

GetErrorRateOk returns a tuple with the ErrorRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorRate

`func (o *EventLLMOverview) SetErrorRate(v float64)`

SetErrorRate sets ErrorRate field to given value.

### HasErrorRate

`func (o *EventLLMOverview) HasErrorRate() bool`

HasErrorRate returns a boolean if a field has been set.

### GetErrors

`func (o *EventLLMOverview) GetErrors() int64`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *EventLLMOverview) GetErrorsOk() (*int64, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *EventLLMOverview) SetErrors(v int64)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *EventLLMOverview) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetModels

`func (o *EventLLMOverview) GetModels() int64`

GetModels returns the Models field if non-nil, zero value otherwise.

### GetModelsOk

`func (o *EventLLMOverview) GetModelsOk() (*int64, bool)`

GetModelsOk returns a tuple with the Models field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModels

`func (o *EventLLMOverview) SetModels(v int64)`

SetModels sets Models field to given value.

### HasModels

`func (o *EventLLMOverview) HasModels() bool`

HasModels returns a boolean if a field has been set.

### GetPromptTokens

`func (o *EventLLMOverview) GetPromptTokens() int64`

GetPromptTokens returns the PromptTokens field if non-nil, zero value otherwise.

### GetPromptTokensOk

`func (o *EventLLMOverview) GetPromptTokensOk() (*int64, bool)`

GetPromptTokensOk returns a tuple with the PromptTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptTokens

`func (o *EventLLMOverview) SetPromptTokens(v int64)`

SetPromptTokens sets PromptTokens field to given value.

### HasPromptTokens

`func (o *EventLLMOverview) HasPromptTokens() bool`

HasPromptTokens returns a boolean if a field has been set.

### GetProviders

`func (o *EventLLMOverview) GetProviders() int64`

GetProviders returns the Providers field if non-nil, zero value otherwise.

### GetProvidersOk

`func (o *EventLLMOverview) GetProvidersOk() (*int64, bool)`

GetProvidersOk returns a tuple with the Providers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviders

`func (o *EventLLMOverview) SetProviders(v int64)`

SetProviders sets Providers field to given value.

### HasProviders

`func (o *EventLLMOverview) HasProviders() bool`

HasProviders returns a boolean if a field has been set.

### GetRequests

`func (o *EventLLMOverview) GetRequests() int64`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *EventLLMOverview) GetRequestsOk() (*int64, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *EventLLMOverview) SetRequests(v int64)`

SetRequests sets Requests field to given value.

### HasRequests

`func (o *EventLLMOverview) HasRequests() bool`

HasRequests returns a boolean if a field has been set.

### GetSource

`func (o *EventLLMOverview) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *EventLLMOverview) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *EventLLMOverview) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *EventLLMOverview) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSpendCents

`func (o *EventLLMOverview) GetSpendCents() int64`

GetSpendCents returns the SpendCents field if non-nil, zero value otherwise.

### GetSpendCentsOk

`func (o *EventLLMOverview) GetSpendCentsOk() (*int64, bool)`

GetSpendCentsOk returns a tuple with the SpendCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpendCents

`func (o *EventLLMOverview) SetSpendCents(v int64)`

SetSpendCents sets SpendCents field to given value.

### HasSpendCents

`func (o *EventLLMOverview) HasSpendCents() bool`

HasSpendCents returns a boolean if a field has been set.

### GetTokens

`func (o *EventLLMOverview) GetTokens() int64`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *EventLLMOverview) GetTokensOk() (*int64, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *EventLLMOverview) SetTokens(v int64)`

SetTokens sets Tokens field to given value.

### HasTokens

`func (o *EventLLMOverview) HasTokens() bool`

HasTokens returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


