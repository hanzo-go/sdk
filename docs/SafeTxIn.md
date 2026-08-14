# SafeTxIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChainId** | Pointer to **int32** | ChainID is the EVM chain the Safe transaction is bound to. 0 uses the wallet&#39;s own chain, or the Hanzo L1 (36963) when it is chain-agnostic. | [optional] 
**Data** | Pointer to **string** | Data is the call data, hex-encoded. | [optional] 
**Nonce** | Pointer to **int32** | Nonce is the Safe&#39;s transaction nonce. | [optional] 
**To** | Pointer to **string** | To is the transaction&#39;s target address. | [optional] 
**Value** | Pointer to **string** | Value is the native-token amount to send, as a decimal string in wei. | [optional] 

## Methods

### NewSafeTxIn

`func NewSafeTxIn() *SafeTxIn`

NewSafeTxIn instantiates a new SafeTxIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSafeTxInWithDefaults

`func NewSafeTxInWithDefaults() *SafeTxIn`

NewSafeTxInWithDefaults instantiates a new SafeTxIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChainId

`func (o *SafeTxIn) GetChainId() int32`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *SafeTxIn) GetChainIdOk() (*int32, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *SafeTxIn) SetChainId(v int32)`

SetChainId sets ChainId field to given value.

### HasChainId

`func (o *SafeTxIn) HasChainId() bool`

HasChainId returns a boolean if a field has been set.

### GetData

`func (o *SafeTxIn) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SafeTxIn) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SafeTxIn) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *SafeTxIn) HasData() bool`

HasData returns a boolean if a field has been set.

### GetNonce

`func (o *SafeTxIn) GetNonce() int32`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *SafeTxIn) GetNonceOk() (*int32, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *SafeTxIn) SetNonce(v int32)`

SetNonce sets Nonce field to given value.

### HasNonce

`func (o *SafeTxIn) HasNonce() bool`

HasNonce returns a boolean if a field has been set.

### GetTo

`func (o *SafeTxIn) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *SafeTxIn) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *SafeTxIn) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *SafeTxIn) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetValue

`func (o *SafeTxIn) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SafeTxIn) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SafeTxIn) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *SafeTxIn) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


