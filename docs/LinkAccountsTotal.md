# LinkAccountsTotal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accounts** | Pointer to **int64** | Accounts is how many linked accounts the total folds. | [optional] 
**CompletionTokens** | Pointer to **int64** | CompletionTokens is the total completion-token count. | [optional] 
**CostCents** | Pointer to **int64** | CostCents is the total cost in cents. | [optional] 
**PromptTokens** | Pointer to **int64** | PromptTokens is the total prompt-token count. | [optional] 
**Requests** | Pointer to **int64** | Requests is the total request count the gateway routed. | [optional] 
**TotalTokens** | Pointer to **int64** | TotalTokens is the total token count. | [optional] 

## Methods

### NewLinkAccountsTotal

`func NewLinkAccountsTotal() *LinkAccountsTotal`

NewLinkAccountsTotal instantiates a new LinkAccountsTotal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkAccountsTotalWithDefaults

`func NewLinkAccountsTotalWithDefaults() *LinkAccountsTotal`

NewLinkAccountsTotalWithDefaults instantiates a new LinkAccountsTotal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccounts

`func (o *LinkAccountsTotal) GetAccounts() int64`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *LinkAccountsTotal) GetAccountsOk() (*int64, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *LinkAccountsTotal) SetAccounts(v int64)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *LinkAccountsTotal) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetCompletionTokens

`func (o *LinkAccountsTotal) GetCompletionTokens() int64`

GetCompletionTokens returns the CompletionTokens field if non-nil, zero value otherwise.

### GetCompletionTokensOk

`func (o *LinkAccountsTotal) GetCompletionTokensOk() (*int64, bool)`

GetCompletionTokensOk returns a tuple with the CompletionTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionTokens

`func (o *LinkAccountsTotal) SetCompletionTokens(v int64)`

SetCompletionTokens sets CompletionTokens field to given value.

### HasCompletionTokens

`func (o *LinkAccountsTotal) HasCompletionTokens() bool`

HasCompletionTokens returns a boolean if a field has been set.

### GetCostCents

`func (o *LinkAccountsTotal) GetCostCents() int64`

GetCostCents returns the CostCents field if non-nil, zero value otherwise.

### GetCostCentsOk

`func (o *LinkAccountsTotal) GetCostCentsOk() (*int64, bool)`

GetCostCentsOk returns a tuple with the CostCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostCents

`func (o *LinkAccountsTotal) SetCostCents(v int64)`

SetCostCents sets CostCents field to given value.

### HasCostCents

`func (o *LinkAccountsTotal) HasCostCents() bool`

HasCostCents returns a boolean if a field has been set.

### GetPromptTokens

`func (o *LinkAccountsTotal) GetPromptTokens() int64`

GetPromptTokens returns the PromptTokens field if non-nil, zero value otherwise.

### GetPromptTokensOk

`func (o *LinkAccountsTotal) GetPromptTokensOk() (*int64, bool)`

GetPromptTokensOk returns a tuple with the PromptTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptTokens

`func (o *LinkAccountsTotal) SetPromptTokens(v int64)`

SetPromptTokens sets PromptTokens field to given value.

### HasPromptTokens

`func (o *LinkAccountsTotal) HasPromptTokens() bool`

HasPromptTokens returns a boolean if a field has been set.

### GetRequests

`func (o *LinkAccountsTotal) GetRequests() int64`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *LinkAccountsTotal) GetRequestsOk() (*int64, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *LinkAccountsTotal) SetRequests(v int64)`

SetRequests sets Requests field to given value.

### HasRequests

`func (o *LinkAccountsTotal) HasRequests() bool`

HasRequests returns a boolean if a field has been set.

### GetTotalTokens

`func (o *LinkAccountsTotal) GetTotalTokens() int64`

GetTotalTokens returns the TotalTokens field if non-nil, zero value otherwise.

### GetTotalTokensOk

`func (o *LinkAccountsTotal) GetTotalTokensOk() (*int64, bool)`

GetTotalTokensOk returns a tuple with the TotalTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTokens

`func (o *LinkAccountsTotal) SetTotalTokens(v int64)`

SetTotalTokens sets TotalTokens field to given value.

### HasTotalTokens

`func (o *LinkAccountsTotal) HasTotalTokens() bool`

HasTotalTokens returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


