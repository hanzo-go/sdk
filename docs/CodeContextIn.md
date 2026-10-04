# CodeContextIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BudgetTokens** | Pointer to **int64** | BudgetTokens caps the bundle&#39;s size. Clamped to [256, 32000]; 0 or absent uses 4000. | [optional] 
**Query** | Pointer to **string** | Query is what to retrieve context for. Required, max 4000 bytes. | [optional] 
**Repo** | Pointer to **string** | Repo narrows retrieval to one repository. Empty searches every repo the org has indexed. | [optional] 

## Methods

### NewCodeContextIn

`func NewCodeContextIn() *CodeContextIn`

NewCodeContextIn instantiates a new CodeContextIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCodeContextInWithDefaults

`func NewCodeContextInWithDefaults() *CodeContextIn`

NewCodeContextInWithDefaults instantiates a new CodeContextIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBudgetTokens

`func (o *CodeContextIn) GetBudgetTokens() int64`

GetBudgetTokens returns the BudgetTokens field if non-nil, zero value otherwise.

### GetBudgetTokensOk

`func (o *CodeContextIn) GetBudgetTokensOk() (*int64, bool)`

GetBudgetTokensOk returns a tuple with the BudgetTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetTokens

`func (o *CodeContextIn) SetBudgetTokens(v int64)`

SetBudgetTokens sets BudgetTokens field to given value.

### HasBudgetTokens

`func (o *CodeContextIn) HasBudgetTokens() bool`

HasBudgetTokens returns a boolean if a field has been set.

### GetQuery

`func (o *CodeContextIn) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *CodeContextIn) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *CodeContextIn) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *CodeContextIn) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetRepo

`func (o *CodeContextIn) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *CodeContextIn) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *CodeContextIn) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *CodeContextIn) HasRepo() bool`

HasRepo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


