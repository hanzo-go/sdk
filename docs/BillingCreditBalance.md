# BillingCreditBalance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Balances** | Pointer to [**[]BillingCreditEntry**](BillingCreditEntry.md) |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 

## Methods

### NewBillingCreditBalance

`func NewBillingCreditBalance() *BillingCreditBalance`

NewBillingCreditBalance instantiates a new BillingCreditBalance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingCreditBalanceWithDefaults

`func NewBillingCreditBalanceWithDefaults() *BillingCreditBalance`

NewBillingCreditBalanceWithDefaults instantiates a new BillingCreditBalance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBalances

`func (o *BillingCreditBalance) GetBalances() []BillingCreditEntry`

GetBalances returns the Balances field if non-nil, zero value otherwise.

### GetBalancesOk

`func (o *BillingCreditBalance) GetBalancesOk() (*[]BillingCreditEntry, bool)`

GetBalancesOk returns a tuple with the Balances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalances

`func (o *BillingCreditBalance) SetBalances(v []BillingCreditEntry)`

SetBalances sets Balances field to given value.

### HasBalances

`func (o *BillingCreditBalance) HasBalances() bool`

HasBalances returns a boolean if a field has been set.

### GetUserId

`func (o *BillingCreditBalance) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *BillingCreditBalance) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *BillingCreditBalance) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *BillingCreditBalance) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


