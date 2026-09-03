# ContextBundle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BudgetTokens** | Pointer to **int64** | BudgetTokens is the ceiling the caller asked for. Packing stops under it, so this is a bound and not a target. | [optional] 
**Query** | Pointer to **string** | Query is the ask this bundle was packed for, echoed back so a cached or forwarded bundle still says what it answers. | [optional] 
**Repo** | Pointer to **string** | Repo narrows the retrieval to one repository. Absent means every indexed repo was searched. | [optional] 
**Spans** | Pointer to [**[]Span**](Span.md) | Spans are the packed chunks, most relevant first, each expanded with the definitions it calls and its notable callers. The top match is always present even if it had to be truncated to fit, so a matched query never comes back with nothing. | [optional] 
**UsedTokens** | Pointer to **int64** | UsedTokens is what the returned spans actually cost, by the same estimate the packer used (roughly one token per four characters — an estimate, not a tokenizer&#39;s count, so size a real window with headroom). | [optional] 

## Methods

### NewContextBundle

`func NewContextBundle() *ContextBundle`

NewContextBundle instantiates a new ContextBundle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContextBundleWithDefaults

`func NewContextBundleWithDefaults() *ContextBundle`

NewContextBundleWithDefaults instantiates a new ContextBundle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBudgetTokens

`func (o *ContextBundle) GetBudgetTokens() int64`

GetBudgetTokens returns the BudgetTokens field if non-nil, zero value otherwise.

### GetBudgetTokensOk

`func (o *ContextBundle) GetBudgetTokensOk() (*int64, bool)`

GetBudgetTokensOk returns a tuple with the BudgetTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetTokens

`func (o *ContextBundle) SetBudgetTokens(v int64)`

SetBudgetTokens sets BudgetTokens field to given value.

### HasBudgetTokens

`func (o *ContextBundle) HasBudgetTokens() bool`

HasBudgetTokens returns a boolean if a field has been set.

### GetQuery

`func (o *ContextBundle) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *ContextBundle) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *ContextBundle) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *ContextBundle) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetRepo

`func (o *ContextBundle) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *ContextBundle) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *ContextBundle) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *ContextBundle) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### GetSpans

`func (o *ContextBundle) GetSpans() []Span`

GetSpans returns the Spans field if non-nil, zero value otherwise.

### GetSpansOk

`func (o *ContextBundle) GetSpansOk() (*[]Span, bool)`

GetSpansOk returns a tuple with the Spans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpans

`func (o *ContextBundle) SetSpans(v []Span)`

SetSpans sets Spans field to given value.

### HasSpans

`func (o *ContextBundle) HasSpans() bool`

HasSpans returns a boolean if a field has been set.

### GetUsedTokens

`func (o *ContextBundle) GetUsedTokens() int64`

GetUsedTokens returns the UsedTokens field if non-nil, zero value otherwise.

### GetUsedTokensOk

`func (o *ContextBundle) GetUsedTokensOk() (*int64, bool)`

GetUsedTokensOk returns a tuple with the UsedTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedTokens

`func (o *ContextBundle) SetUsedTokens(v int64)`

SetUsedTokens sets UsedTokens field to given value.

### HasUsedTokens

`func (o *ContextBundle) HasUsedTokens() bool`

HasUsedTokens returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


