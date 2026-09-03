# BoardTotals

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

### NewBoardTotals

`func NewBoardTotals() *BoardTotals`

NewBoardTotals instantiates a new BoardTotals object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBoardTotalsWithDefaults

`func NewBoardTotalsWithDefaults() *BoardTotals`

NewBoardTotalsWithDefaults instantiates a new BoardTotals object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompletionTokens

`func (o *BoardTotals) GetCompletionTokens() int64`

GetCompletionTokens returns the CompletionTokens field if non-nil, zero value otherwise.

### GetCompletionTokensOk

`func (o *BoardTotals) GetCompletionTokensOk() (*int64, bool)`

GetCompletionTokensOk returns a tuple with the CompletionTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionTokens

`func (o *BoardTotals) SetCompletionTokens(v int64)`

SetCompletionTokens sets CompletionTokens field to given value.

### HasCompletionTokens

`func (o *BoardTotals) HasCompletionTokens() bool`

HasCompletionTokens returns a boolean if a field has been set.

### GetCostCents

`func (o *BoardTotals) GetCostCents() int64`

GetCostCents returns the CostCents field if non-nil, zero value otherwise.

### GetCostCentsOk

`func (o *BoardTotals) GetCostCentsOk() (*int64, bool)`

GetCostCentsOk returns a tuple with the CostCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostCents

`func (o *BoardTotals) SetCostCents(v int64)`

SetCostCents sets CostCents field to given value.

### HasCostCents

`func (o *BoardTotals) HasCostCents() bool`

HasCostCents returns a boolean if a field has been set.

### GetErrors

`func (o *BoardTotals) GetErrors() int64`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *BoardTotals) GetErrorsOk() (*int64, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *BoardTotals) SetErrors(v int64)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *BoardTotals) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetGenerations

`func (o *BoardTotals) GetGenerations() int64`

GetGenerations returns the Generations field if non-nil, zero value otherwise.

### GetGenerationsOk

`func (o *BoardTotals) GetGenerationsOk() (*int64, bool)`

GetGenerationsOk returns a tuple with the Generations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerations

`func (o *BoardTotals) SetGenerations(v int64)`

SetGenerations sets Generations field to given value.

### HasGenerations

`func (o *BoardTotals) HasGenerations() bool`

HasGenerations returns a boolean if a field has been set.

### GetModels

`func (o *BoardTotals) GetModels() int64`

GetModels returns the Models field if non-nil, zero value otherwise.

### GetModelsOk

`func (o *BoardTotals) GetModelsOk() (*int64, bool)`

GetModelsOk returns a tuple with the Models field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModels

`func (o *BoardTotals) SetModels(v int64)`

SetModels sets Models field to given value.

### HasModels

`func (o *BoardTotals) HasModels() bool`

HasModels returns a boolean if a field has been set.

### GetPromptTokens

`func (o *BoardTotals) GetPromptTokens() int64`

GetPromptTokens returns the PromptTokens field if non-nil, zero value otherwise.

### GetPromptTokensOk

`func (o *BoardTotals) GetPromptTokensOk() (*int64, bool)`

GetPromptTokensOk returns a tuple with the PromptTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptTokens

`func (o *BoardTotals) SetPromptTokens(v int64)`

SetPromptTokens sets PromptTokens field to given value.

### HasPromptTokens

`func (o *BoardTotals) HasPromptTokens() bool`

HasPromptTokens returns a boolean if a field has been set.

### GetSuccessRate

`func (o *BoardTotals) GetSuccessRate() float64`

GetSuccessRate returns the SuccessRate field if non-nil, zero value otherwise.

### GetSuccessRateOk

`func (o *BoardTotals) GetSuccessRateOk() (*float64, bool)`

GetSuccessRateOk returns a tuple with the SuccessRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessRate

`func (o *BoardTotals) SetSuccessRate(v float64)`

SetSuccessRate sets SuccessRate field to given value.

### HasSuccessRate

`func (o *BoardTotals) HasSuccessRate() bool`

HasSuccessRate returns a boolean if a field has been set.

### GetTotalTokens

`func (o *BoardTotals) GetTotalTokens() int64`

GetTotalTokens returns the TotalTokens field if non-nil, zero value otherwise.

### GetTotalTokensOk

`func (o *BoardTotals) GetTotalTokensOk() (*int64, bool)`

GetTotalTokensOk returns a tuple with the TotalTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTokens

`func (o *BoardTotals) SetTotalTokens(v int64)`

SetTotalTokens sets TotalTokens field to given value.

### HasTotalTokens

`func (o *BoardTotals) HasTotalTokens() bool`

HasTotalTokens returns a boolean if a field has been set.

### GetUsers

`func (o *BoardTotals) GetUsers() int64`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *BoardTotals) GetUsersOk() (*int64, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *BoardTotals) SetUsers(v int64)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *BoardTotals) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


