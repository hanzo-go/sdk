# CloudSafeProposal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**R** | Pointer to **string** | R is the r component of the MPC threshold signature over the Safe-tx hash. | [optional] 
**S** | Pointer to **string** | S is the s component of that signature. | [optional] 
**SafeAddress** | Pointer to **string** | SafeAddress is the Safe contract this transaction is for. | [optional] 
**SafeTxHash** | Pointer to **string** | SafeTxHash is the EIP-712 Safe transaction hash, bound to the Safe contract and the chain id — the value the owner approval signs. | [optional] 
**WalletId** | Pointer to **string** | WalletID is the wallet whose Safe this is. | [optional] 

## Methods

### NewCloudSafeProposal

`func NewCloudSafeProposal() *CloudSafeProposal`

NewCloudSafeProposal instantiates a new CloudSafeProposal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudSafeProposalWithDefaults

`func NewCloudSafeProposalWithDefaults() *CloudSafeProposal`

NewCloudSafeProposalWithDefaults instantiates a new CloudSafeProposal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetR

`func (o *CloudSafeProposal) GetR() string`

GetR returns the R field if non-nil, zero value otherwise.

### GetROk

`func (o *CloudSafeProposal) GetROk() (*string, bool)`

GetROk returns a tuple with the R field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetR

`func (o *CloudSafeProposal) SetR(v string)`

SetR sets R field to given value.

### HasR

`func (o *CloudSafeProposal) HasR() bool`

HasR returns a boolean if a field has been set.

### GetS

`func (o *CloudSafeProposal) GetS() string`

GetS returns the S field if non-nil, zero value otherwise.

### GetSOk

`func (o *CloudSafeProposal) GetSOk() (*string, bool)`

GetSOk returns a tuple with the S field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS

`func (o *CloudSafeProposal) SetS(v string)`

SetS sets S field to given value.

### HasS

`func (o *CloudSafeProposal) HasS() bool`

HasS returns a boolean if a field has been set.

### GetSafeAddress

`func (o *CloudSafeProposal) GetSafeAddress() string`

GetSafeAddress returns the SafeAddress field if non-nil, zero value otherwise.

### GetSafeAddressOk

`func (o *CloudSafeProposal) GetSafeAddressOk() (*string, bool)`

GetSafeAddressOk returns a tuple with the SafeAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafeAddress

`func (o *CloudSafeProposal) SetSafeAddress(v string)`

SetSafeAddress sets SafeAddress field to given value.

### HasSafeAddress

`func (o *CloudSafeProposal) HasSafeAddress() bool`

HasSafeAddress returns a boolean if a field has been set.

### GetSafeTxHash

`func (o *CloudSafeProposal) GetSafeTxHash() string`

GetSafeTxHash returns the SafeTxHash field if non-nil, zero value otherwise.

### GetSafeTxHashOk

`func (o *CloudSafeProposal) GetSafeTxHashOk() (*string, bool)`

GetSafeTxHashOk returns a tuple with the SafeTxHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafeTxHash

`func (o *CloudSafeProposal) SetSafeTxHash(v string)`

SetSafeTxHash sets SafeTxHash field to given value.

### HasSafeTxHash

`func (o *CloudSafeProposal) HasSafeTxHash() bool`

HasSafeTxHash returns a boolean if a field has been set.

### GetWalletId

`func (o *CloudSafeProposal) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *CloudSafeProposal) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *CloudSafeProposal) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.

### HasWalletId

`func (o *CloudSafeProposal) HasWalletId() bool`

HasWalletId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


