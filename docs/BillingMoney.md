# BillingMoney

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | **string** |  | 
**Decimal** | **string** |  | 

## Methods

### NewBillingMoney

`func NewBillingMoney(currency string, decimal string, ) *BillingMoney`

NewBillingMoney instantiates a new BillingMoney object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingMoneyWithDefaults

`func NewBillingMoneyWithDefaults() *BillingMoney`

NewBillingMoneyWithDefaults instantiates a new BillingMoney object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *BillingMoney) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *BillingMoney) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *BillingMoney) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetDecimal

`func (o *BillingMoney) GetDecimal() string`

GetDecimal returns the Decimal field if non-nil, zero value otherwise.

### GetDecimalOk

`func (o *BillingMoney) GetDecimalOk() (*string, bool)`

GetDecimalOk returns a tuple with the Decimal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecimal

`func (o *BillingMoney) SetDecimal(v string)`

SetDecimal sets Decimal field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


