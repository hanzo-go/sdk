# CloudLLMOverview

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

### NewCloudLLMOverview

`func NewCloudLLMOverview() *CloudLLMOverview`

NewCloudLLMOverview instantiates a new CloudLLMOverview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudLLMOverviewWithDefaults

`func NewCloudLLMOverviewWithDefaults() *CloudLLMOverview`

NewCloudLLMOverviewWithDefaults instantiates a new CloudLLMOverview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailable

`func (o *CloudLLMOverview) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *CloudLLMOverview) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *CloudLLMOverview) SetAvailable(v bool)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *CloudLLMOverview) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetCompletionTokens

`func (o *CloudLLMOverview) GetCompletionTokens() int32`

GetCompletionTokens returns the CompletionTokens field if non-nil, zero value otherwise.

### GetCompletionTokensOk

`func (o *CloudLLMOverview) GetCompletionTokensOk() (*int32, bool)`

GetCompletionTokensOk returns a tuple with the CompletionTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionTokens

`func (o *CloudLLMOverview) SetCompletionTokens(v int32)`

SetCompletionTokens sets CompletionTokens field to given value.

### HasCompletionTokens

`func (o *CloudLLMOverview) HasCompletionTokens() bool`

HasCompletionTokens returns a boolean if a field has been set.

### GetErrorRate

`func (o *CloudLLMOverview) GetErrorRate() float32`

GetErrorRate returns the ErrorRate field if non-nil, zero value otherwise.

### GetErrorRateOk

`func (o *CloudLLMOverview) GetErrorRateOk() (*float32, bool)`

GetErrorRateOk returns a tuple with the ErrorRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorRate

`func (o *CloudLLMOverview) SetErrorRate(v float32)`

SetErrorRate sets ErrorRate field to given value.

### HasErrorRate

`func (o *CloudLLMOverview) HasErrorRate() bool`

HasErrorRate returns a boolean if a field has been set.

### GetErrors

`func (o *CloudLLMOverview) GetErrors() int32`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *CloudLLMOverview) GetErrorsOk() (*int32, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *CloudLLMOverview) SetErrors(v int32)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *CloudLLMOverview) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetModels

`func (o *CloudLLMOverview) GetModels() int32`

GetModels returns the Models field if non-nil, zero value otherwise.

### GetModelsOk

`func (o *CloudLLMOverview) GetModelsOk() (*int32, bool)`

GetModelsOk returns a tuple with the Models field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModels

`func (o *CloudLLMOverview) SetModels(v int32)`

SetModels sets Models field to given value.

### HasModels

`func (o *CloudLLMOverview) HasModels() bool`

HasModels returns a boolean if a field has been set.

### GetPromptTokens

`func (o *CloudLLMOverview) GetPromptTokens() int32`

GetPromptTokens returns the PromptTokens field if non-nil, zero value otherwise.

### GetPromptTokensOk

`func (o *CloudLLMOverview) GetPromptTokensOk() (*int32, bool)`

GetPromptTokensOk returns a tuple with the PromptTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptTokens

`func (o *CloudLLMOverview) SetPromptTokens(v int32)`

SetPromptTokens sets PromptTokens field to given value.

### HasPromptTokens

`func (o *CloudLLMOverview) HasPromptTokens() bool`

HasPromptTokens returns a boolean if a field has been set.

### GetProviders

`func (o *CloudLLMOverview) GetProviders() int32`

GetProviders returns the Providers field if non-nil, zero value otherwise.

### GetProvidersOk

`func (o *CloudLLMOverview) GetProvidersOk() (*int32, bool)`

GetProvidersOk returns a tuple with the Providers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviders

`func (o *CloudLLMOverview) SetProviders(v int32)`

SetProviders sets Providers field to given value.

### HasProviders

`func (o *CloudLLMOverview) HasProviders() bool`

HasProviders returns a boolean if a field has been set.

### GetRequests

`func (o *CloudLLMOverview) GetRequests() int32`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *CloudLLMOverview) GetRequestsOk() (*int32, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *CloudLLMOverview) SetRequests(v int32)`

SetRequests sets Requests field to given value.

### HasRequests

`func (o *CloudLLMOverview) HasRequests() bool`

HasRequests returns a boolean if a field has been set.

### GetSource

`func (o *CloudLLMOverview) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *CloudLLMOverview) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *CloudLLMOverview) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *CloudLLMOverview) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSpendCents

`func (o *CloudLLMOverview) GetSpendCents() int32`

GetSpendCents returns the SpendCents field if non-nil, zero value otherwise.

### GetSpendCentsOk

`func (o *CloudLLMOverview) GetSpendCentsOk() (*int32, bool)`

GetSpendCentsOk returns a tuple with the SpendCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpendCents

`func (o *CloudLLMOverview) SetSpendCents(v int32)`

SetSpendCents sets SpendCents field to given value.

### HasSpendCents

`func (o *CloudLLMOverview) HasSpendCents() bool`

HasSpendCents returns a boolean if a field has been set.

### GetTokens

`func (o *CloudLLMOverview) GetTokens() int32`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *CloudLLMOverview) GetTokensOk() (*int32, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *CloudLLMOverview) SetTokens(v int32)`

SetTokens sets Tokens field to given value.

### HasTokens

`func (o *CloudLLMOverview) HasTokens() bool`

HasTokens returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


