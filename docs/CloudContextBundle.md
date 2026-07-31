# CloudContextBundle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BudgetTokens** | Pointer to **int32** |  | [optional] 
**Query** | Pointer to **string** |  | [optional] 
**Repo** | Pointer to **string** |  | [optional] 
**Spans** | Pointer to [**[]CloudSpan**](CloudSpan.md) |  | [optional] 
**UsedTokens** | Pointer to **int32** |  | [optional] 

## Methods

### NewCloudContextBundle

`func NewCloudContextBundle() *CloudContextBundle`

NewCloudContextBundle instantiates a new CloudContextBundle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudContextBundleWithDefaults

`func NewCloudContextBundleWithDefaults() *CloudContextBundle`

NewCloudContextBundleWithDefaults instantiates a new CloudContextBundle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBudgetTokens

`func (o *CloudContextBundle) GetBudgetTokens() int32`

GetBudgetTokens returns the BudgetTokens field if non-nil, zero value otherwise.

### GetBudgetTokensOk

`func (o *CloudContextBundle) GetBudgetTokensOk() (*int32, bool)`

GetBudgetTokensOk returns a tuple with the BudgetTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetTokens

`func (o *CloudContextBundle) SetBudgetTokens(v int32)`

SetBudgetTokens sets BudgetTokens field to given value.

### HasBudgetTokens

`func (o *CloudContextBundle) HasBudgetTokens() bool`

HasBudgetTokens returns a boolean if a field has been set.

### GetQuery

`func (o *CloudContextBundle) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *CloudContextBundle) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *CloudContextBundle) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *CloudContextBundle) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetRepo

`func (o *CloudContextBundle) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *CloudContextBundle) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *CloudContextBundle) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *CloudContextBundle) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### GetSpans

`func (o *CloudContextBundle) GetSpans() []CloudSpan`

GetSpans returns the Spans field if non-nil, zero value otherwise.

### GetSpansOk

`func (o *CloudContextBundle) GetSpansOk() (*[]CloudSpan, bool)`

GetSpansOk returns a tuple with the Spans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpans

`func (o *CloudContextBundle) SetSpans(v []CloudSpan)`

SetSpans sets Spans field to given value.

### HasSpans

`func (o *CloudContextBundle) HasSpans() bool`

HasSpans returns a boolean if a field has been set.

### GetUsedTokens

`func (o *CloudContextBundle) GetUsedTokens() int32`

GetUsedTokens returns the UsedTokens field if non-nil, zero value otherwise.

### GetUsedTokensOk

`func (o *CloudContextBundle) GetUsedTokensOk() (*int32, bool)`

GetUsedTokensOk returns a tuple with the UsedTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedTokens

`func (o *CloudContextBundle) SetUsedTokens(v int32)`

SetUsedTokens sets UsedTokens field to given value.

### HasUsedTokens

`func (o *CloudContextBundle) HasUsedTokens() bool`

HasUsedTokens returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


