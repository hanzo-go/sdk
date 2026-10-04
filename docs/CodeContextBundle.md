# CodeContextBundle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BudgetTokens** | Pointer to **int64** | BudgetTokens is the ceiling the caller asked for. Packing stops under it, so this is a bound and not a target. | [optional] 
**Capped** | Pointer to **bool** | Capped is true when the semantic tier stopped at its scan cap — it compares a query with at most 50000 stored vectors — with more left unread, so code past the cap could not be picked by meaning. The lexical and symbol tiers read the whole index either way. Absent otherwise. | [optional] 
**Query** | Pointer to **string** | Query is the ask this bundle was packed for, echoed back so a cached or forwarded bundle still says what it answers. | [optional] 
**Repo** | Pointer to **string** | Repo narrows the retrieval to one repository. Absent means every indexed repo was searched. | [optional] 
**Spans** | Pointer to [**[]CodeSpan**](CodeSpan.md) | Spans are the packed chunks, most relevant first, each expanded with the definitions it calls and its notable callers. The top match is always present even if it had to be truncated to fit, so a matched query never comes back with nothing. | [optional] 
**UsedTokens** | Pointer to **int64** | UsedTokens is what the returned spans actually cost, by the same estimate the packer used (roughly one token per four characters — an estimate, not a tokenizer&#39;s count, so size a real window with headroom). | [optional] 

## Methods

### NewCodeContextBundle

`func NewCodeContextBundle() *CodeContextBundle`

NewCodeContextBundle instantiates a new CodeContextBundle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCodeContextBundleWithDefaults

`func NewCodeContextBundleWithDefaults() *CodeContextBundle`

NewCodeContextBundleWithDefaults instantiates a new CodeContextBundle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBudgetTokens

`func (o *CodeContextBundle) GetBudgetTokens() int64`

GetBudgetTokens returns the BudgetTokens field if non-nil, zero value otherwise.

### GetBudgetTokensOk

`func (o *CodeContextBundle) GetBudgetTokensOk() (*int64, bool)`

GetBudgetTokensOk returns a tuple with the BudgetTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetTokens

`func (o *CodeContextBundle) SetBudgetTokens(v int64)`

SetBudgetTokens sets BudgetTokens field to given value.

### HasBudgetTokens

`func (o *CodeContextBundle) HasBudgetTokens() bool`

HasBudgetTokens returns a boolean if a field has been set.

### GetCapped

`func (o *CodeContextBundle) GetCapped() bool`

GetCapped returns the Capped field if non-nil, zero value otherwise.

### GetCappedOk

`func (o *CodeContextBundle) GetCappedOk() (*bool, bool)`

GetCappedOk returns a tuple with the Capped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapped

`func (o *CodeContextBundle) SetCapped(v bool)`

SetCapped sets Capped field to given value.

### HasCapped

`func (o *CodeContextBundle) HasCapped() bool`

HasCapped returns a boolean if a field has been set.

### GetQuery

`func (o *CodeContextBundle) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *CodeContextBundle) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *CodeContextBundle) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *CodeContextBundle) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetRepo

`func (o *CodeContextBundle) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *CodeContextBundle) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *CodeContextBundle) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *CodeContextBundle) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### GetSpans

`func (o *CodeContextBundle) GetSpans() []CodeSpan`

GetSpans returns the Spans field if non-nil, zero value otherwise.

### GetSpansOk

`func (o *CodeContextBundle) GetSpansOk() (*[]CodeSpan, bool)`

GetSpansOk returns a tuple with the Spans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpans

`func (o *CodeContextBundle) SetSpans(v []CodeSpan)`

SetSpans sets Spans field to given value.

### HasSpans

`func (o *CodeContextBundle) HasSpans() bool`

HasSpans returns a boolean if a field has been set.

### GetUsedTokens

`func (o *CodeContextBundle) GetUsedTokens() int64`

GetUsedTokens returns the UsedTokens field if non-nil, zero value otherwise.

### GetUsedTokensOk

`func (o *CodeContextBundle) GetUsedTokensOk() (*int64, bool)`

GetUsedTokensOk returns a tuple with the UsedTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedTokens

`func (o *CodeContextBundle) SetUsedTokens(v int64)`

SetUsedTokens sets UsedTokens field to given value.

### HasUsedTokens

`func (o *CodeContextBundle) HasUsedTokens() bool`

HasUsedTokens returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


