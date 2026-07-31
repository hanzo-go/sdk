# CloudContextIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BudgetTokens** | Pointer to **int32** | BudgetTokens caps the bundle&#39;s size. Clamped to [256, 32000]; 0 or absent uses 4000. | [optional] 
**Query** | Pointer to **string** | Query is what to retrieve context for. Required, max 4000 bytes. | [optional] 
**Repo** | Pointer to **string** | Repo narrows retrieval to one repository. Empty searches every repo the org has indexed. | [optional] 

## Methods

### NewCloudContextIn

`func NewCloudContextIn() *CloudContextIn`

NewCloudContextIn instantiates a new CloudContextIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudContextInWithDefaults

`func NewCloudContextInWithDefaults() *CloudContextIn`

NewCloudContextInWithDefaults instantiates a new CloudContextIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBudgetTokens

`func (o *CloudContextIn) GetBudgetTokens() int32`

GetBudgetTokens returns the BudgetTokens field if non-nil, zero value otherwise.

### GetBudgetTokensOk

`func (o *CloudContextIn) GetBudgetTokensOk() (*int32, bool)`

GetBudgetTokensOk returns a tuple with the BudgetTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetTokens

`func (o *CloudContextIn) SetBudgetTokens(v int32)`

SetBudgetTokens sets BudgetTokens field to given value.

### HasBudgetTokens

`func (o *CloudContextIn) HasBudgetTokens() bool`

HasBudgetTokens returns a boolean if a field has been set.

### GetQuery

`func (o *CloudContextIn) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *CloudContextIn) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *CloudContextIn) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *CloudContextIn) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetRepo

`func (o *CloudContextIn) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *CloudContextIn) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *CloudContextIn) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *CloudContextIn) HasRepo() bool`

HasRepo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


