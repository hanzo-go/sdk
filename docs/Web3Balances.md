# Web3Balances

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | Address is the account they belong to. | [optional] 
**Chain** | Pointer to **string** | Chain is the chain the balances were read from. | [optional] 
**Native** | Pointer to **string** | Native is the chain&#39;s own currency, as a 0x-quantity — the RPC&#39;s own encoding, not a float, because a wei value does not survive float64. | [optional] 

## Methods

### NewWeb3Balances

`func NewWeb3Balances() *Web3Balances`

NewWeb3Balances instantiates a new Web3Balances object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWeb3BalancesWithDefaults

`func NewWeb3BalancesWithDefaults() *Web3Balances`

NewWeb3BalancesWithDefaults instantiates a new Web3Balances object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *Web3Balances) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *Web3Balances) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *Web3Balances) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *Web3Balances) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetChain

`func (o *Web3Balances) GetChain() string`

GetChain returns the Chain field if non-nil, zero value otherwise.

### GetChainOk

`func (o *Web3Balances) GetChainOk() (*string, bool)`

GetChainOk returns a tuple with the Chain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChain

`func (o *Web3Balances) SetChain(v string)`

SetChain sets Chain field to given value.

### HasChain

`func (o *Web3Balances) HasChain() bool`

HasChain returns a boolean if a field has been set.

### GetNative

`func (o *Web3Balances) GetNative() string`

GetNative returns the Native field if non-nil, zero value otherwise.

### GetNativeOk

`func (o *Web3Balances) GetNativeOk() (*string, bool)`

GetNativeOk returns a tuple with the Native field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNative

`func (o *Web3Balances) SetNative(v string)`

SetNative sets Native field to given value.

### HasNative

`func (o *Web3Balances) HasNative() bool`

HasNative returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


