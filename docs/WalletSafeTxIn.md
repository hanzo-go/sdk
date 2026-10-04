# WalletSafeTxIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChainId** | Pointer to **int64** | ChainID is the EVM chain the Safe transaction is bound to. 0 uses the wallet&#39;s own chain, or the Hanzo L1 (36963) when it is chain-agnostic. | [optional] 
**Data** | Pointer to **string** | Data is the call data, hex-encoded. | [optional] 
**Nonce** | Pointer to **int64** | Nonce is the Safe&#39;s transaction nonce. | [optional] 
**To** | Pointer to **string** | To is the transaction&#39;s target address. | [optional] 
**Value** | Pointer to **string** | Value is the native-token amount to send, as a decimal string in wei. | [optional] 

## Methods

### NewWalletSafeTxIn

`func NewWalletSafeTxIn() *WalletSafeTxIn`

NewWalletSafeTxIn instantiates a new WalletSafeTxIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWalletSafeTxInWithDefaults

`func NewWalletSafeTxInWithDefaults() *WalletSafeTxIn`

NewWalletSafeTxInWithDefaults instantiates a new WalletSafeTxIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChainId

`func (o *WalletSafeTxIn) GetChainId() int64`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *WalletSafeTxIn) GetChainIdOk() (*int64, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *WalletSafeTxIn) SetChainId(v int64)`

SetChainId sets ChainId field to given value.

### HasChainId

`func (o *WalletSafeTxIn) HasChainId() bool`

HasChainId returns a boolean if a field has been set.

### GetData

`func (o *WalletSafeTxIn) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *WalletSafeTxIn) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *WalletSafeTxIn) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *WalletSafeTxIn) HasData() bool`

HasData returns a boolean if a field has been set.

### GetNonce

`func (o *WalletSafeTxIn) GetNonce() int64`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *WalletSafeTxIn) GetNonceOk() (*int64, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *WalletSafeTxIn) SetNonce(v int64)`

SetNonce sets Nonce field to given value.

### HasNonce

`func (o *WalletSafeTxIn) HasNonce() bool`

HasNonce returns a boolean if a field has been set.

### GetTo

`func (o *WalletSafeTxIn) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *WalletSafeTxIn) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *WalletSafeTxIn) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *WalletSafeTxIn) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetValue

`func (o *WalletSafeTxIn) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *WalletSafeTxIn) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *WalletSafeTxIn) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *WalletSafeTxIn) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


