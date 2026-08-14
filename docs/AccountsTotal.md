# AccountsTotal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accounts** | Pointer to **int32** | Accounts is how many linked accounts the total folds. | [optional] 
**CompletionTokens** | Pointer to **int32** | CompletionTokens is the total completion-token count. | [optional] 
**CostCents** | Pointer to **int32** | CostCents is the total cost in cents. | [optional] 
**PromptTokens** | Pointer to **int32** | PromptTokens is the total prompt-token count. | [optional] 
**Requests** | Pointer to **int32** | Requests is the total request count the gateway routed. | [optional] 
**TotalTokens** | Pointer to **int32** | TotalTokens is the total token count. | [optional] 

## Methods

### NewAccountsTotal

`func NewAccountsTotal() *AccountsTotal`

NewAccountsTotal instantiates a new AccountsTotal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountsTotalWithDefaults

`func NewAccountsTotalWithDefaults() *AccountsTotal`

NewAccountsTotalWithDefaults instantiates a new AccountsTotal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccounts

`func (o *AccountsTotal) GetAccounts() int32`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *AccountsTotal) GetAccountsOk() (*int32, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *AccountsTotal) SetAccounts(v int32)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *AccountsTotal) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetCompletionTokens

`func (o *AccountsTotal) GetCompletionTokens() int32`

GetCompletionTokens returns the CompletionTokens field if non-nil, zero value otherwise.

### GetCompletionTokensOk

`func (o *AccountsTotal) GetCompletionTokensOk() (*int32, bool)`

GetCompletionTokensOk returns a tuple with the CompletionTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionTokens

`func (o *AccountsTotal) SetCompletionTokens(v int32)`

SetCompletionTokens sets CompletionTokens field to given value.

### HasCompletionTokens

`func (o *AccountsTotal) HasCompletionTokens() bool`

HasCompletionTokens returns a boolean if a field has been set.

### GetCostCents

`func (o *AccountsTotal) GetCostCents() int32`

GetCostCents returns the CostCents field if non-nil, zero value otherwise.

### GetCostCentsOk

`func (o *AccountsTotal) GetCostCentsOk() (*int32, bool)`

GetCostCentsOk returns a tuple with the CostCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostCents

`func (o *AccountsTotal) SetCostCents(v int32)`

SetCostCents sets CostCents field to given value.

### HasCostCents

`func (o *AccountsTotal) HasCostCents() bool`

HasCostCents returns a boolean if a field has been set.

### GetPromptTokens

`func (o *AccountsTotal) GetPromptTokens() int32`

GetPromptTokens returns the PromptTokens field if non-nil, zero value otherwise.

### GetPromptTokensOk

`func (o *AccountsTotal) GetPromptTokensOk() (*int32, bool)`

GetPromptTokensOk returns a tuple with the PromptTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptTokens

`func (o *AccountsTotal) SetPromptTokens(v int32)`

SetPromptTokens sets PromptTokens field to given value.

### HasPromptTokens

`func (o *AccountsTotal) HasPromptTokens() bool`

HasPromptTokens returns a boolean if a field has been set.

### GetRequests

`func (o *AccountsTotal) GetRequests() int32`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *AccountsTotal) GetRequestsOk() (*int32, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *AccountsTotal) SetRequests(v int32)`

SetRequests sets Requests field to given value.

### HasRequests

`func (o *AccountsTotal) HasRequests() bool`

HasRequests returns a boolean if a field has been set.

### GetTotalTokens

`func (o *AccountsTotal) GetTotalTokens() int32`

GetTotalTokens returns the TotalTokens field if non-nil, zero value otherwise.

### GetTotalTokensOk

`func (o *AccountsTotal) GetTotalTokensOk() (*int32, bool)`

GetTotalTokensOk returns a tuple with the TotalTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTokens

`func (o *AccountsTotal) SetTotalTokens(v int32)`

SetTotalTokens sets TotalTokens field to given value.

### HasTotalTokens

`func (o *AccountsTotal) HasTotalTokens() bool`

HasTotalTokens returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


