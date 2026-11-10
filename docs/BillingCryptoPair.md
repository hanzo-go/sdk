# BillingCryptoPair

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Chain** | Pointer to **string** | Chain is the network, as a deposit names it (\&quot;base\&quot;). | [optional] 
**Token** | Pointer to **string** | Token is the asset on that chain, as a deposit names it (\&quot;usdc\&quot;). | [optional] 

## Methods

### NewBillingCryptoPair

`func NewBillingCryptoPair() *BillingCryptoPair`

NewBillingCryptoPair instantiates a new BillingCryptoPair object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingCryptoPairWithDefaults

`func NewBillingCryptoPairWithDefaults() *BillingCryptoPair`

NewBillingCryptoPairWithDefaults instantiates a new BillingCryptoPair object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChain

`func (o *BillingCryptoPair) GetChain() string`

GetChain returns the Chain field if non-nil, zero value otherwise.

### GetChainOk

`func (o *BillingCryptoPair) GetChainOk() (*string, bool)`

GetChainOk returns a tuple with the Chain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChain

`func (o *BillingCryptoPair) SetChain(v string)`

SetChain sets Chain field to given value.

### HasChain

`func (o *BillingCryptoPair) HasChain() bool`

HasChain returns a boolean if a field has been set.

### GetToken

`func (o *BillingCryptoPair) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *BillingCryptoPair) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *BillingCryptoPair) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *BillingCryptoPair) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


