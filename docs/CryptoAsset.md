# CryptoAsset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountCents** | Pointer to **int32** | AmountCents is what the payer intends to send, for the record. Optional — the credit is what actually arrives, never what was announced. | [optional] 
**Chain** | Pointer to **string** | Chain is the network to receive on. Empty takes the rail&#39;s default. | [optional] 
**Token** | Pointer to **string** | Token is the asset on that chain. Empty takes the chain&#39;s native one. | [optional] 

## Methods

### NewCryptoAsset

`func NewCryptoAsset() *CryptoAsset`

NewCryptoAsset instantiates a new CryptoAsset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCryptoAssetWithDefaults

`func NewCryptoAssetWithDefaults() *CryptoAsset`

NewCryptoAssetWithDefaults instantiates a new CryptoAsset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountCents

`func (o *CryptoAsset) GetAmountCents() int32`

GetAmountCents returns the AmountCents field if non-nil, zero value otherwise.

### GetAmountCentsOk

`func (o *CryptoAsset) GetAmountCentsOk() (*int32, bool)`

GetAmountCentsOk returns a tuple with the AmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCents

`func (o *CryptoAsset) SetAmountCents(v int32)`

SetAmountCents sets AmountCents field to given value.

### HasAmountCents

`func (o *CryptoAsset) HasAmountCents() bool`

HasAmountCents returns a boolean if a field has been set.

### GetChain

`func (o *CryptoAsset) GetChain() string`

GetChain returns the Chain field if non-nil, zero value otherwise.

### GetChainOk

`func (o *CryptoAsset) GetChainOk() (*string, bool)`

GetChainOk returns a tuple with the Chain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChain

`func (o *CryptoAsset) SetChain(v string)`

SetChain sets Chain field to given value.

### HasChain

`func (o *CryptoAsset) HasChain() bool`

HasChain returns a boolean if a field has been set.

### GetToken

`func (o *CryptoAsset) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CryptoAsset) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CryptoAsset) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *CryptoAsset) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


