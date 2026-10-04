# WalletSafeProposal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**R** | Pointer to **string** | R is the r component of the MPC threshold signature over the Safe-tx hash. | [optional] 
**S** | Pointer to **string** | S is the s component of that signature. | [optional] 
**SafeAddress** | Pointer to **string** | SafeAddress is the Safe contract this transaction is for. | [optional] 
**SafeTxHash** | Pointer to **string** | SafeTxHash is the EIP-712 Safe transaction hash, bound to the Safe contract and the chain id — the value the owner approval signs. | [optional] 
**WalletId** | Pointer to **string** | WalletID is the wallet whose Safe this is. | [optional] 

## Methods

### NewWalletSafeProposal

`func NewWalletSafeProposal() *WalletSafeProposal`

NewWalletSafeProposal instantiates a new WalletSafeProposal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWalletSafeProposalWithDefaults

`func NewWalletSafeProposalWithDefaults() *WalletSafeProposal`

NewWalletSafeProposalWithDefaults instantiates a new WalletSafeProposal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetR

`func (o *WalletSafeProposal) GetR() string`

GetR returns the R field if non-nil, zero value otherwise.

### GetROk

`func (o *WalletSafeProposal) GetROk() (*string, bool)`

GetROk returns a tuple with the R field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetR

`func (o *WalletSafeProposal) SetR(v string)`

SetR sets R field to given value.

### HasR

`func (o *WalletSafeProposal) HasR() bool`

HasR returns a boolean if a field has been set.

### GetS

`func (o *WalletSafeProposal) GetS() string`

GetS returns the S field if non-nil, zero value otherwise.

### GetSOk

`func (o *WalletSafeProposal) GetSOk() (*string, bool)`

GetSOk returns a tuple with the S field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS

`func (o *WalletSafeProposal) SetS(v string)`

SetS sets S field to given value.

### HasS

`func (o *WalletSafeProposal) HasS() bool`

HasS returns a boolean if a field has been set.

### GetSafeAddress

`func (o *WalletSafeProposal) GetSafeAddress() string`

GetSafeAddress returns the SafeAddress field if non-nil, zero value otherwise.

### GetSafeAddressOk

`func (o *WalletSafeProposal) GetSafeAddressOk() (*string, bool)`

GetSafeAddressOk returns a tuple with the SafeAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafeAddress

`func (o *WalletSafeProposal) SetSafeAddress(v string)`

SetSafeAddress sets SafeAddress field to given value.

### HasSafeAddress

`func (o *WalletSafeProposal) HasSafeAddress() bool`

HasSafeAddress returns a boolean if a field has been set.

### GetSafeTxHash

`func (o *WalletSafeProposal) GetSafeTxHash() string`

GetSafeTxHash returns the SafeTxHash field if non-nil, zero value otherwise.

### GetSafeTxHashOk

`func (o *WalletSafeProposal) GetSafeTxHashOk() (*string, bool)`

GetSafeTxHashOk returns a tuple with the SafeTxHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafeTxHash

`func (o *WalletSafeProposal) SetSafeTxHash(v string)`

SetSafeTxHash sets SafeTxHash field to given value.

### HasSafeTxHash

`func (o *WalletSafeProposal) HasSafeTxHash() bool`

HasSafeTxHash returns a boolean if a field has been set.

### GetWalletId

`func (o *WalletSafeProposal) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *WalletSafeProposal) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *WalletSafeProposal) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.

### HasWalletId

`func (o *WalletSafeProposal) HasWalletId() bool`

HasWalletId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


