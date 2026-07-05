# CommerceWalletPayRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**To** | Pointer to **string** |  | [optional] 
**Amount** | Pointer to **int32** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 

## Methods

### NewCommerceWalletPayRequest

`func NewCommerceWalletPayRequest() *CommerceWalletPayRequest`

NewCommerceWalletPayRequest instantiates a new CommerceWalletPayRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommerceWalletPayRequestWithDefaults

`func NewCommerceWalletPayRequestWithDefaults() *CommerceWalletPayRequest`

NewCommerceWalletPayRequestWithDefaults instantiates a new CommerceWalletPayRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTo

`func (o *CommerceWalletPayRequest) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *CommerceWalletPayRequest) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *CommerceWalletPayRequest) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *CommerceWalletPayRequest) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetAmount

`func (o *CommerceWalletPayRequest) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CommerceWalletPayRequest) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CommerceWalletPayRequest) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *CommerceWalletPayRequest) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCurrency

`func (o *CommerceWalletPayRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CommerceWalletPayRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CommerceWalletPayRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CommerceWalletPayRequest) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


