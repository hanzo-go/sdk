# BalanceSheet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AsOf** | Pointer to **string** |  | [optional] 
**Assets** | Pointer to [**[]BalanceLine**](BalanceLine.md) |  | [optional] 
**Balanced** | Pointer to **bool** | TotalAssets &#x3D;&#x3D; TotalLiabilities + TotalEquity | [optional] 
**Equity** | Pointer to [**[]BalanceLine**](BalanceLine.md) |  | [optional] 
**Liabilities** | Pointer to [**[]BalanceLine**](BalanceLine.md) |  | [optional] 
**TotalAssets** | Pointer to **int32** |  | [optional] 
**TotalEquity** | Pointer to **int32** |  | [optional] 
**TotalLiabilities** | Pointer to **int32** |  | [optional] 

## Methods

### NewBalanceSheet

`func NewBalanceSheet() *BalanceSheet`

NewBalanceSheet instantiates a new BalanceSheet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBalanceSheetWithDefaults

`func NewBalanceSheetWithDefaults() *BalanceSheet`

NewBalanceSheetWithDefaults instantiates a new BalanceSheet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsOf

`func (o *BalanceSheet) GetAsOf() string`

GetAsOf returns the AsOf field if non-nil, zero value otherwise.

### GetAsOfOk

`func (o *BalanceSheet) GetAsOfOk() (*string, bool)`

GetAsOfOk returns a tuple with the AsOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsOf

`func (o *BalanceSheet) SetAsOf(v string)`

SetAsOf sets AsOf field to given value.

### HasAsOf

`func (o *BalanceSheet) HasAsOf() bool`

HasAsOf returns a boolean if a field has been set.

### GetAssets

`func (o *BalanceSheet) GetAssets() []BalanceLine`

GetAssets returns the Assets field if non-nil, zero value otherwise.

### GetAssetsOk

`func (o *BalanceSheet) GetAssetsOk() (*[]BalanceLine, bool)`

GetAssetsOk returns a tuple with the Assets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssets

`func (o *BalanceSheet) SetAssets(v []BalanceLine)`

SetAssets sets Assets field to given value.

### HasAssets

`func (o *BalanceSheet) HasAssets() bool`

HasAssets returns a boolean if a field has been set.

### GetBalanced

`func (o *BalanceSheet) GetBalanced() bool`

GetBalanced returns the Balanced field if non-nil, zero value otherwise.

### GetBalancedOk

`func (o *BalanceSheet) GetBalancedOk() (*bool, bool)`

GetBalancedOk returns a tuple with the Balanced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanced

`func (o *BalanceSheet) SetBalanced(v bool)`

SetBalanced sets Balanced field to given value.

### HasBalanced

`func (o *BalanceSheet) HasBalanced() bool`

HasBalanced returns a boolean if a field has been set.

### GetEquity

`func (o *BalanceSheet) GetEquity() []BalanceLine`

GetEquity returns the Equity field if non-nil, zero value otherwise.

### GetEquityOk

`func (o *BalanceSheet) GetEquityOk() (*[]BalanceLine, bool)`

GetEquityOk returns a tuple with the Equity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquity

`func (o *BalanceSheet) SetEquity(v []BalanceLine)`

SetEquity sets Equity field to given value.

### HasEquity

`func (o *BalanceSheet) HasEquity() bool`

HasEquity returns a boolean if a field has been set.

### GetLiabilities

`func (o *BalanceSheet) GetLiabilities() []BalanceLine`

GetLiabilities returns the Liabilities field if non-nil, zero value otherwise.

### GetLiabilitiesOk

`func (o *BalanceSheet) GetLiabilitiesOk() (*[]BalanceLine, bool)`

GetLiabilitiesOk returns a tuple with the Liabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiabilities

`func (o *BalanceSheet) SetLiabilities(v []BalanceLine)`

SetLiabilities sets Liabilities field to given value.

### HasLiabilities

`func (o *BalanceSheet) HasLiabilities() bool`

HasLiabilities returns a boolean if a field has been set.

### GetTotalAssets

`func (o *BalanceSheet) GetTotalAssets() int32`

GetTotalAssets returns the TotalAssets field if non-nil, zero value otherwise.

### GetTotalAssetsOk

`func (o *BalanceSheet) GetTotalAssetsOk() (*int32, bool)`

GetTotalAssetsOk returns a tuple with the TotalAssets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAssets

`func (o *BalanceSheet) SetTotalAssets(v int32)`

SetTotalAssets sets TotalAssets field to given value.

### HasTotalAssets

`func (o *BalanceSheet) HasTotalAssets() bool`

HasTotalAssets returns a boolean if a field has been set.

### GetTotalEquity

`func (o *BalanceSheet) GetTotalEquity() int32`

GetTotalEquity returns the TotalEquity field if non-nil, zero value otherwise.

### GetTotalEquityOk

`func (o *BalanceSheet) GetTotalEquityOk() (*int32, bool)`

GetTotalEquityOk returns a tuple with the TotalEquity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalEquity

`func (o *BalanceSheet) SetTotalEquity(v int32)`

SetTotalEquity sets TotalEquity field to given value.

### HasTotalEquity

`func (o *BalanceSheet) HasTotalEquity() bool`

HasTotalEquity returns a boolean if a field has been set.

### GetTotalLiabilities

`func (o *BalanceSheet) GetTotalLiabilities() int32`

GetTotalLiabilities returns the TotalLiabilities field if non-nil, zero value otherwise.

### GetTotalLiabilitiesOk

`func (o *BalanceSheet) GetTotalLiabilitiesOk() (*int32, bool)`

GetTotalLiabilitiesOk returns a tuple with the TotalLiabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalLiabilities

`func (o *BalanceSheet) SetTotalLiabilities(v int32)`

SetTotalLiabilities sets TotalLiabilities field to given value.

### HasTotalLiabilities

`func (o *BalanceSheet) HasTotalLiabilities() bool`

HasTotalLiabilities returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


