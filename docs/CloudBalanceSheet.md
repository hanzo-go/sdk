# CloudBalanceSheet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AsOf** | Pointer to **string** |  | [optional] 
**Assets** | Pointer to [**[]CloudBalanceLine**](CloudBalanceLine.md) |  | [optional] 
**Balanced** | Pointer to **bool** | TotalAssets &#x3D;&#x3D; TotalLiabilities + TotalEquity | [optional] 
**Equity** | Pointer to [**[]CloudBalanceLine**](CloudBalanceLine.md) |  | [optional] 
**Liabilities** | Pointer to [**[]CloudBalanceLine**](CloudBalanceLine.md) |  | [optional] 
**TotalAssets** | Pointer to **int32** |  | [optional] 
**TotalEquity** | Pointer to **int32** |  | [optional] 
**TotalLiabilities** | Pointer to **int32** |  | [optional] 

## Methods

### NewCloudBalanceSheet

`func NewCloudBalanceSheet() *CloudBalanceSheet`

NewCloudBalanceSheet instantiates a new CloudBalanceSheet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudBalanceSheetWithDefaults

`func NewCloudBalanceSheetWithDefaults() *CloudBalanceSheet`

NewCloudBalanceSheetWithDefaults instantiates a new CloudBalanceSheet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsOf

`func (o *CloudBalanceSheet) GetAsOf() string`

GetAsOf returns the AsOf field if non-nil, zero value otherwise.

### GetAsOfOk

`func (o *CloudBalanceSheet) GetAsOfOk() (*string, bool)`

GetAsOfOk returns a tuple with the AsOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsOf

`func (o *CloudBalanceSheet) SetAsOf(v string)`

SetAsOf sets AsOf field to given value.

### HasAsOf

`func (o *CloudBalanceSheet) HasAsOf() bool`

HasAsOf returns a boolean if a field has been set.

### GetAssets

`func (o *CloudBalanceSheet) GetAssets() []CloudBalanceLine`

GetAssets returns the Assets field if non-nil, zero value otherwise.

### GetAssetsOk

`func (o *CloudBalanceSheet) GetAssetsOk() (*[]CloudBalanceLine, bool)`

GetAssetsOk returns a tuple with the Assets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssets

`func (o *CloudBalanceSheet) SetAssets(v []CloudBalanceLine)`

SetAssets sets Assets field to given value.

### HasAssets

`func (o *CloudBalanceSheet) HasAssets() bool`

HasAssets returns a boolean if a field has been set.

### GetBalanced

`func (o *CloudBalanceSheet) GetBalanced() bool`

GetBalanced returns the Balanced field if non-nil, zero value otherwise.

### GetBalancedOk

`func (o *CloudBalanceSheet) GetBalancedOk() (*bool, bool)`

GetBalancedOk returns a tuple with the Balanced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanced

`func (o *CloudBalanceSheet) SetBalanced(v bool)`

SetBalanced sets Balanced field to given value.

### HasBalanced

`func (o *CloudBalanceSheet) HasBalanced() bool`

HasBalanced returns a boolean if a field has been set.

### GetEquity

`func (o *CloudBalanceSheet) GetEquity() []CloudBalanceLine`

GetEquity returns the Equity field if non-nil, zero value otherwise.

### GetEquityOk

`func (o *CloudBalanceSheet) GetEquityOk() (*[]CloudBalanceLine, bool)`

GetEquityOk returns a tuple with the Equity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquity

`func (o *CloudBalanceSheet) SetEquity(v []CloudBalanceLine)`

SetEquity sets Equity field to given value.

### HasEquity

`func (o *CloudBalanceSheet) HasEquity() bool`

HasEquity returns a boolean if a field has been set.

### GetLiabilities

`func (o *CloudBalanceSheet) GetLiabilities() []CloudBalanceLine`

GetLiabilities returns the Liabilities field if non-nil, zero value otherwise.

### GetLiabilitiesOk

`func (o *CloudBalanceSheet) GetLiabilitiesOk() (*[]CloudBalanceLine, bool)`

GetLiabilitiesOk returns a tuple with the Liabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiabilities

`func (o *CloudBalanceSheet) SetLiabilities(v []CloudBalanceLine)`

SetLiabilities sets Liabilities field to given value.

### HasLiabilities

`func (o *CloudBalanceSheet) HasLiabilities() bool`

HasLiabilities returns a boolean if a field has been set.

### GetTotalAssets

`func (o *CloudBalanceSheet) GetTotalAssets() int32`

GetTotalAssets returns the TotalAssets field if non-nil, zero value otherwise.

### GetTotalAssetsOk

`func (o *CloudBalanceSheet) GetTotalAssetsOk() (*int32, bool)`

GetTotalAssetsOk returns a tuple with the TotalAssets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAssets

`func (o *CloudBalanceSheet) SetTotalAssets(v int32)`

SetTotalAssets sets TotalAssets field to given value.

### HasTotalAssets

`func (o *CloudBalanceSheet) HasTotalAssets() bool`

HasTotalAssets returns a boolean if a field has been set.

### GetTotalEquity

`func (o *CloudBalanceSheet) GetTotalEquity() int32`

GetTotalEquity returns the TotalEquity field if non-nil, zero value otherwise.

### GetTotalEquityOk

`func (o *CloudBalanceSheet) GetTotalEquityOk() (*int32, bool)`

GetTotalEquityOk returns a tuple with the TotalEquity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalEquity

`func (o *CloudBalanceSheet) SetTotalEquity(v int32)`

SetTotalEquity sets TotalEquity field to given value.

### HasTotalEquity

`func (o *CloudBalanceSheet) HasTotalEquity() bool`

HasTotalEquity returns a boolean if a field has been set.

### GetTotalLiabilities

`func (o *CloudBalanceSheet) GetTotalLiabilities() int32`

GetTotalLiabilities returns the TotalLiabilities field if non-nil, zero value otherwise.

### GetTotalLiabilitiesOk

`func (o *CloudBalanceSheet) GetTotalLiabilitiesOk() (*int32, bool)`

GetTotalLiabilitiesOk returns a tuple with the TotalLiabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalLiabilities

`func (o *CloudBalanceSheet) SetTotalLiabilities(v int32)`

SetTotalLiabilities sets TotalLiabilities field to given value.

### HasTotalLiabilities

`func (o *CloudBalanceSheet) HasTotalLiabilities() bool`

HasTotalLiabilities returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


