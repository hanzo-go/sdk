# LinkRoutedUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** | Account is the provider-side account identifier. | [optional] 
**Billing** | Pointer to **string** | Billing is how the routed inference bills: plan or commerce. | [optional] 
**CompletionTokens** | Pointer to **int64** | CompletionTokens is the routed completion-token count. | [optional] 
**CostCents** | Pointer to **int64** | CostCents is the routed cost in cents. | [optional] 
**Kind** | Pointer to **string** | Kind is how the account authenticates: subscription or apikey. | [optional] 
**PromptTokens** | Pointer to **int64** | PromptTokens is the routed prompt-token count. | [optional] 
**Provider** | Pointer to **string** | Provider is the AI provider the row&#39;s account belongs to. | [optional] 
**Requests** | Pointer to **int64** | Requests is how many requests the gateway routed through this account. | [optional] 
**TotalTokens** | Pointer to **int64** | TotalTokens is the routed total token count. | [optional] 

## Methods

### NewLinkRoutedUsage

`func NewLinkRoutedUsage() *LinkRoutedUsage`

NewLinkRoutedUsage instantiates a new LinkRoutedUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkRoutedUsageWithDefaults

`func NewLinkRoutedUsageWithDefaults() *LinkRoutedUsage`

NewLinkRoutedUsageWithDefaults instantiates a new LinkRoutedUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *LinkRoutedUsage) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *LinkRoutedUsage) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *LinkRoutedUsage) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *LinkRoutedUsage) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetBilling

`func (o *LinkRoutedUsage) GetBilling() string`

GetBilling returns the Billing field if non-nil, zero value otherwise.

### GetBillingOk

`func (o *LinkRoutedUsage) GetBillingOk() (*string, bool)`

GetBillingOk returns a tuple with the Billing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBilling

`func (o *LinkRoutedUsage) SetBilling(v string)`

SetBilling sets Billing field to given value.

### HasBilling

`func (o *LinkRoutedUsage) HasBilling() bool`

HasBilling returns a boolean if a field has been set.

### GetCompletionTokens

`func (o *LinkRoutedUsage) GetCompletionTokens() int64`

GetCompletionTokens returns the CompletionTokens field if non-nil, zero value otherwise.

### GetCompletionTokensOk

`func (o *LinkRoutedUsage) GetCompletionTokensOk() (*int64, bool)`

GetCompletionTokensOk returns a tuple with the CompletionTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionTokens

`func (o *LinkRoutedUsage) SetCompletionTokens(v int64)`

SetCompletionTokens sets CompletionTokens field to given value.

### HasCompletionTokens

`func (o *LinkRoutedUsage) HasCompletionTokens() bool`

HasCompletionTokens returns a boolean if a field has been set.

### GetCostCents

`func (o *LinkRoutedUsage) GetCostCents() int64`

GetCostCents returns the CostCents field if non-nil, zero value otherwise.

### GetCostCentsOk

`func (o *LinkRoutedUsage) GetCostCentsOk() (*int64, bool)`

GetCostCentsOk returns a tuple with the CostCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostCents

`func (o *LinkRoutedUsage) SetCostCents(v int64)`

SetCostCents sets CostCents field to given value.

### HasCostCents

`func (o *LinkRoutedUsage) HasCostCents() bool`

HasCostCents returns a boolean if a field has been set.

### GetKind

`func (o *LinkRoutedUsage) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *LinkRoutedUsage) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *LinkRoutedUsage) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *LinkRoutedUsage) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetPromptTokens

`func (o *LinkRoutedUsage) GetPromptTokens() int64`

GetPromptTokens returns the PromptTokens field if non-nil, zero value otherwise.

### GetPromptTokensOk

`func (o *LinkRoutedUsage) GetPromptTokensOk() (*int64, bool)`

GetPromptTokensOk returns a tuple with the PromptTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptTokens

`func (o *LinkRoutedUsage) SetPromptTokens(v int64)`

SetPromptTokens sets PromptTokens field to given value.

### HasPromptTokens

`func (o *LinkRoutedUsage) HasPromptTokens() bool`

HasPromptTokens returns a boolean if a field has been set.

### GetProvider

`func (o *LinkRoutedUsage) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *LinkRoutedUsage) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *LinkRoutedUsage) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *LinkRoutedUsage) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetRequests

`func (o *LinkRoutedUsage) GetRequests() int64`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *LinkRoutedUsage) GetRequestsOk() (*int64, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *LinkRoutedUsage) SetRequests(v int64)`

SetRequests sets Requests field to given value.

### HasRequests

`func (o *LinkRoutedUsage) HasRequests() bool`

HasRequests returns a boolean if a field has been set.

### GetTotalTokens

`func (o *LinkRoutedUsage) GetTotalTokens() int64`

GetTotalTokens returns the TotalTokens field if non-nil, zero value otherwise.

### GetTotalTokensOk

`func (o *LinkRoutedUsage) GetTotalTokensOk() (*int64, bool)`

GetTotalTokensOk returns a tuple with the TotalTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTokens

`func (o *LinkRoutedUsage) SetTotalTokens(v int64)`

SetTotalTokens sets TotalTokens field to given value.

### HasTotalTokens

`func (o *LinkRoutedUsage) HasTotalTokens() bool`

HasTotalTokens returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


