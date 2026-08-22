# CreditBalance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Balances** | Pointer to [**[]CreditEntry**](CreditEntry.md) |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 

## Methods

### NewCreditBalance

`func NewCreditBalance() *CreditBalance`

NewCreditBalance instantiates a new CreditBalance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditBalanceWithDefaults

`func NewCreditBalanceWithDefaults() *CreditBalance`

NewCreditBalanceWithDefaults instantiates a new CreditBalance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBalances

`func (o *CreditBalance) GetBalances() []CreditEntry`

GetBalances returns the Balances field if non-nil, zero value otherwise.

### GetBalancesOk

`func (o *CreditBalance) GetBalancesOk() (*[]CreditEntry, bool)`

GetBalancesOk returns a tuple with the Balances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalances

`func (o *CreditBalance) SetBalances(v []CreditEntry)`

SetBalances sets Balances field to given value.

### HasBalances

`func (o *CreditBalance) HasBalances() bool`

HasBalances returns a boolean if a field has been set.

### GetUserId

`func (o *CreditBalance) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *CreditBalance) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *CreditBalance) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *CreditBalance) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


