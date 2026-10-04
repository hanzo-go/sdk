# EvalBoardTotals

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompletionTokens** | Pointer to **int64** | tokens the models answered with | [optional] 
**CostCents** | Pointer to **int64** | what the window cost, in cents | [optional] 
**Errors** | Pointer to **int64** | calls that did not succeed | [optional] 
**Generations** | Pointer to **int64** | how many model calls the window holds | [optional] 
**Models** | Pointer to **int64** | how many distinct models were called | [optional] 
**PromptTokens** | Pointer to **int64** | tokens sent to the models | [optional] 
**SuccessRate** | Pointer to **float64** | share of calls that succeeded, 0..1 | [optional] 
**TotalTokens** | Pointer to **int64** | prompt plus completion | [optional] 
**Users** | Pointer to **int64** | how many distinct users called them | [optional] 

## Methods

### NewEvalBoardTotals

`func NewEvalBoardTotals() *EvalBoardTotals`

NewEvalBoardTotals instantiates a new EvalBoardTotals object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEvalBoardTotalsWithDefaults

`func NewEvalBoardTotalsWithDefaults() *EvalBoardTotals`

NewEvalBoardTotalsWithDefaults instantiates a new EvalBoardTotals object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompletionTokens

`func (o *EvalBoardTotals) GetCompletionTokens() int64`

GetCompletionTokens returns the CompletionTokens field if non-nil, zero value otherwise.

### GetCompletionTokensOk

`func (o *EvalBoardTotals) GetCompletionTokensOk() (*int64, bool)`

GetCompletionTokensOk returns a tuple with the CompletionTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionTokens

`func (o *EvalBoardTotals) SetCompletionTokens(v int64)`

SetCompletionTokens sets CompletionTokens field to given value.

### HasCompletionTokens

`func (o *EvalBoardTotals) HasCompletionTokens() bool`

HasCompletionTokens returns a boolean if a field has been set.

### GetCostCents

`func (o *EvalBoardTotals) GetCostCents() int64`

GetCostCents returns the CostCents field if non-nil, zero value otherwise.

### GetCostCentsOk

`func (o *EvalBoardTotals) GetCostCentsOk() (*int64, bool)`

GetCostCentsOk returns a tuple with the CostCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostCents

`func (o *EvalBoardTotals) SetCostCents(v int64)`

SetCostCents sets CostCents field to given value.

### HasCostCents

`func (o *EvalBoardTotals) HasCostCents() bool`

HasCostCents returns a boolean if a field has been set.

### GetErrors

`func (o *EvalBoardTotals) GetErrors() int64`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *EvalBoardTotals) GetErrorsOk() (*int64, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *EvalBoardTotals) SetErrors(v int64)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *EvalBoardTotals) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetGenerations

`func (o *EvalBoardTotals) GetGenerations() int64`

GetGenerations returns the Generations field if non-nil, zero value otherwise.

### GetGenerationsOk

`func (o *EvalBoardTotals) GetGenerationsOk() (*int64, bool)`

GetGenerationsOk returns a tuple with the Generations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerations

`func (o *EvalBoardTotals) SetGenerations(v int64)`

SetGenerations sets Generations field to given value.

### HasGenerations

`func (o *EvalBoardTotals) HasGenerations() bool`

HasGenerations returns a boolean if a field has been set.

### GetModels

`func (o *EvalBoardTotals) GetModels() int64`

GetModels returns the Models field if non-nil, zero value otherwise.

### GetModelsOk

`func (o *EvalBoardTotals) GetModelsOk() (*int64, bool)`

GetModelsOk returns a tuple with the Models field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModels

`func (o *EvalBoardTotals) SetModels(v int64)`

SetModels sets Models field to given value.

### HasModels

`func (o *EvalBoardTotals) HasModels() bool`

HasModels returns a boolean if a field has been set.

### GetPromptTokens

`func (o *EvalBoardTotals) GetPromptTokens() int64`

GetPromptTokens returns the PromptTokens field if non-nil, zero value otherwise.

### GetPromptTokensOk

`func (o *EvalBoardTotals) GetPromptTokensOk() (*int64, bool)`

GetPromptTokensOk returns a tuple with the PromptTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptTokens

`func (o *EvalBoardTotals) SetPromptTokens(v int64)`

SetPromptTokens sets PromptTokens field to given value.

### HasPromptTokens

`func (o *EvalBoardTotals) HasPromptTokens() bool`

HasPromptTokens returns a boolean if a field has been set.

### GetSuccessRate

`func (o *EvalBoardTotals) GetSuccessRate() float64`

GetSuccessRate returns the SuccessRate field if non-nil, zero value otherwise.

### GetSuccessRateOk

`func (o *EvalBoardTotals) GetSuccessRateOk() (*float64, bool)`

GetSuccessRateOk returns a tuple with the SuccessRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessRate

`func (o *EvalBoardTotals) SetSuccessRate(v float64)`

SetSuccessRate sets SuccessRate field to given value.

### HasSuccessRate

`func (o *EvalBoardTotals) HasSuccessRate() bool`

HasSuccessRate returns a boolean if a field has been set.

### GetTotalTokens

`func (o *EvalBoardTotals) GetTotalTokens() int64`

GetTotalTokens returns the TotalTokens field if non-nil, zero value otherwise.

### GetTotalTokensOk

`func (o *EvalBoardTotals) GetTotalTokensOk() (*int64, bool)`

GetTotalTokensOk returns a tuple with the TotalTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTokens

`func (o *EvalBoardTotals) SetTotalTokens(v int64)`

SetTotalTokens sets TotalTokens field to given value.

### HasTotalTokens

`func (o *EvalBoardTotals) HasTotalTokens() bool`

HasTotalTokens returns a boolean if a field has been set.

### GetUsers

`func (o *EvalBoardTotals) GetUsers() int64`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *EvalBoardTotals) GetUsersOk() (*int64, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *EvalBoardTotals) SetUsers(v int64)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *EvalBoardTotals) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


