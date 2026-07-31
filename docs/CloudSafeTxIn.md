# CloudSafeTxIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChainId** | Pointer to **int32** | ChainID is the EVM chain the Safe transaction is bound to. 0 uses the wallet&#39;s own chain, or the Hanzo L1 (36963) when it is chain-agnostic. | [optional] 
**Data** | Pointer to **string** | Data is the call data, hex-encoded. | [optional] 
**Nonce** | Pointer to **int32** | Nonce is the Safe&#39;s transaction nonce. | [optional] 
**To** | Pointer to **string** | To is the transaction&#39;s target address. | [optional] 
**Value** | Pointer to **string** | Value is the native-token amount to send, as a decimal string in wei. | [optional] 

## Methods

### NewCloudSafeTxIn

`func NewCloudSafeTxIn() *CloudSafeTxIn`

NewCloudSafeTxIn instantiates a new CloudSafeTxIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudSafeTxInWithDefaults

`func NewCloudSafeTxInWithDefaults() *CloudSafeTxIn`

NewCloudSafeTxInWithDefaults instantiates a new CloudSafeTxIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChainId

`func (o *CloudSafeTxIn) GetChainId() int32`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *CloudSafeTxIn) GetChainIdOk() (*int32, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *CloudSafeTxIn) SetChainId(v int32)`

SetChainId sets ChainId field to given value.

### HasChainId

`func (o *CloudSafeTxIn) HasChainId() bool`

HasChainId returns a boolean if a field has been set.

### GetData

`func (o *CloudSafeTxIn) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CloudSafeTxIn) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CloudSafeTxIn) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *CloudSafeTxIn) HasData() bool`

HasData returns a boolean if a field has been set.

### GetNonce

`func (o *CloudSafeTxIn) GetNonce() int32`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *CloudSafeTxIn) GetNonceOk() (*int32, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *CloudSafeTxIn) SetNonce(v int32)`

SetNonce sets Nonce field to given value.

### HasNonce

`func (o *CloudSafeTxIn) HasNonce() bool`

HasNonce returns a boolean if a field has been set.

### GetTo

`func (o *CloudSafeTxIn) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *CloudSafeTxIn) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *CloudSafeTxIn) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *CloudSafeTxIn) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetValue

`func (o *CloudSafeTxIn) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CloudSafeTxIn) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CloudSafeTxIn) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *CloudSafeTxIn) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


