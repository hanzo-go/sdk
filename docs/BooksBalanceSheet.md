# BooksBalanceSheet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AsOf** | Pointer to **string** | AsOf is the posting time the statement is taken at, inclusive. A balance sheet is a snapshot, not a window, so there is no From. Absent means as of now. | [optional] 
**Assets** | Pointer to [**[]BooksBalanceLine**](BooksBalanceLine.md) | Assets are what the org OWNS at that instant, one line per account that has a balance. Cash, receivables, funds captured but not yet settled. | [optional] 
**Balanced** | Pointer to **bool** | Balanced is whether assets equal liabilities plus equity — the accounting equation, computed from the totals above rather than assumed. False means the ledger is broken, not that the statement is. | [optional] 
**Equity** | Pointer to [**[]BooksBalanceLine**](BooksBalanceLine.md) | Equity is what is left over for the owners. It carries a DERIVED retained earnings line holding cumulative income minus expense, because this ledger has no period close that sweeps the P&amp;L into equity — without that line the equation would not close. | [optional] 
**Liabilities** | Pointer to [**[]BooksBalanceLine**](BooksBalanceLine.md) | Liabilities are what the org OWES — including customers&#39; unspent prepaid credit, which is their money until it is consumed and so is carried here rather than counted as revenue. | [optional] 
**TotalAssets** | Pointer to **int64** | TotalAssets is the sum of the asset lines, in cents. | [optional] 
**TotalEquity** | Pointer to **int64** | TotalEquity is the sum of the equity lines including retained earnings, in cents. | [optional] 
**TotalLiabilities** | Pointer to **int64** | TotalLiabilities is the sum of the liability lines, in cents. | [optional] 

## Methods

### NewBooksBalanceSheet

`func NewBooksBalanceSheet() *BooksBalanceSheet`

NewBooksBalanceSheet instantiates a new BooksBalanceSheet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBooksBalanceSheetWithDefaults

`func NewBooksBalanceSheetWithDefaults() *BooksBalanceSheet`

NewBooksBalanceSheetWithDefaults instantiates a new BooksBalanceSheet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsOf

`func (o *BooksBalanceSheet) GetAsOf() string`

GetAsOf returns the AsOf field if non-nil, zero value otherwise.

### GetAsOfOk

`func (o *BooksBalanceSheet) GetAsOfOk() (*string, bool)`

GetAsOfOk returns a tuple with the AsOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsOf

`func (o *BooksBalanceSheet) SetAsOf(v string)`

SetAsOf sets AsOf field to given value.

### HasAsOf

`func (o *BooksBalanceSheet) HasAsOf() bool`

HasAsOf returns a boolean if a field has been set.

### GetAssets

`func (o *BooksBalanceSheet) GetAssets() []BooksBalanceLine`

GetAssets returns the Assets field if non-nil, zero value otherwise.

### GetAssetsOk

`func (o *BooksBalanceSheet) GetAssetsOk() (*[]BooksBalanceLine, bool)`

GetAssetsOk returns a tuple with the Assets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssets

`func (o *BooksBalanceSheet) SetAssets(v []BooksBalanceLine)`

SetAssets sets Assets field to given value.

### HasAssets

`func (o *BooksBalanceSheet) HasAssets() bool`

HasAssets returns a boolean if a field has been set.

### GetBalanced

`func (o *BooksBalanceSheet) GetBalanced() bool`

GetBalanced returns the Balanced field if non-nil, zero value otherwise.

### GetBalancedOk

`func (o *BooksBalanceSheet) GetBalancedOk() (*bool, bool)`

GetBalancedOk returns a tuple with the Balanced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanced

`func (o *BooksBalanceSheet) SetBalanced(v bool)`

SetBalanced sets Balanced field to given value.

### HasBalanced

`func (o *BooksBalanceSheet) HasBalanced() bool`

HasBalanced returns a boolean if a field has been set.

### GetEquity

`func (o *BooksBalanceSheet) GetEquity() []BooksBalanceLine`

GetEquity returns the Equity field if non-nil, zero value otherwise.

### GetEquityOk

`func (o *BooksBalanceSheet) GetEquityOk() (*[]BooksBalanceLine, bool)`

GetEquityOk returns a tuple with the Equity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquity

`func (o *BooksBalanceSheet) SetEquity(v []BooksBalanceLine)`

SetEquity sets Equity field to given value.

### HasEquity

`func (o *BooksBalanceSheet) HasEquity() bool`

HasEquity returns a boolean if a field has been set.

### GetLiabilities

`func (o *BooksBalanceSheet) GetLiabilities() []BooksBalanceLine`

GetLiabilities returns the Liabilities field if non-nil, zero value otherwise.

### GetLiabilitiesOk

`func (o *BooksBalanceSheet) GetLiabilitiesOk() (*[]BooksBalanceLine, bool)`

GetLiabilitiesOk returns a tuple with the Liabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiabilities

`func (o *BooksBalanceSheet) SetLiabilities(v []BooksBalanceLine)`

SetLiabilities sets Liabilities field to given value.

### HasLiabilities

`func (o *BooksBalanceSheet) HasLiabilities() bool`

HasLiabilities returns a boolean if a field has been set.

### GetTotalAssets

`func (o *BooksBalanceSheet) GetTotalAssets() int64`

GetTotalAssets returns the TotalAssets field if non-nil, zero value otherwise.

### GetTotalAssetsOk

`func (o *BooksBalanceSheet) GetTotalAssetsOk() (*int64, bool)`

GetTotalAssetsOk returns a tuple with the TotalAssets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAssets

`func (o *BooksBalanceSheet) SetTotalAssets(v int64)`

SetTotalAssets sets TotalAssets field to given value.

### HasTotalAssets

`func (o *BooksBalanceSheet) HasTotalAssets() bool`

HasTotalAssets returns a boolean if a field has been set.

### GetTotalEquity

`func (o *BooksBalanceSheet) GetTotalEquity() int64`

GetTotalEquity returns the TotalEquity field if non-nil, zero value otherwise.

### GetTotalEquityOk

`func (o *BooksBalanceSheet) GetTotalEquityOk() (*int64, bool)`

GetTotalEquityOk returns a tuple with the TotalEquity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalEquity

`func (o *BooksBalanceSheet) SetTotalEquity(v int64)`

SetTotalEquity sets TotalEquity field to given value.

### HasTotalEquity

`func (o *BooksBalanceSheet) HasTotalEquity() bool`

HasTotalEquity returns a boolean if a field has been set.

### GetTotalLiabilities

`func (o *BooksBalanceSheet) GetTotalLiabilities() int64`

GetTotalLiabilities returns the TotalLiabilities field if non-nil, zero value otherwise.

### GetTotalLiabilitiesOk

`func (o *BooksBalanceSheet) GetTotalLiabilitiesOk() (*int64, bool)`

GetTotalLiabilitiesOk returns a tuple with the TotalLiabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalLiabilities

`func (o *BooksBalanceSheet) SetTotalLiabilities(v int64)`

SetTotalLiabilities sets TotalLiabilities field to given value.

### HasTotalLiabilities

`func (o *BooksBalanceSheet) HasTotalLiabilities() bool`

HasTotalLiabilities returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


