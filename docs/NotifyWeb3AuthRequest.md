# NotifyWeb3AuthRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | Ethereum wallet address | 
**Nonce** | **string** | Nonce obtained from /v1/iam/auth/web3-nonce | 
**CreateAt** | Pointer to **int32** | Timestamp when signature was created | [optional] 
**TypedData** | Pointer to **string** | The message that was signed | [optional] 
**Signature** | **string** | Wallet signature | 
**WalletType** | Pointer to **string** | Type of wallet used | [optional] 

## Methods

### NewNotifyWeb3AuthRequest

`func NewNotifyWeb3AuthRequest(address string, nonce string, signature string, ) *NotifyWeb3AuthRequest`

NewNotifyWeb3AuthRequest instantiates a new NotifyWeb3AuthRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotifyWeb3AuthRequestWithDefaults

`func NewNotifyWeb3AuthRequestWithDefaults() *NotifyWeb3AuthRequest`

NewNotifyWeb3AuthRequestWithDefaults instantiates a new NotifyWeb3AuthRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *NotifyWeb3AuthRequest) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *NotifyWeb3AuthRequest) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *NotifyWeb3AuthRequest) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetNonce

`func (o *NotifyWeb3AuthRequest) GetNonce() string`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *NotifyWeb3AuthRequest) GetNonceOk() (*string, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *NotifyWeb3AuthRequest) SetNonce(v string)`

SetNonce sets Nonce field to given value.


### GetCreateAt

`func (o *NotifyWeb3AuthRequest) GetCreateAt() int32`

GetCreateAt returns the CreateAt field if non-nil, zero value otherwise.

### GetCreateAtOk

`func (o *NotifyWeb3AuthRequest) GetCreateAtOk() (*int32, bool)`

GetCreateAtOk returns a tuple with the CreateAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateAt

`func (o *NotifyWeb3AuthRequest) SetCreateAt(v int32)`

SetCreateAt sets CreateAt field to given value.

### HasCreateAt

`func (o *NotifyWeb3AuthRequest) HasCreateAt() bool`

HasCreateAt returns a boolean if a field has been set.

### GetTypedData

`func (o *NotifyWeb3AuthRequest) GetTypedData() string`

GetTypedData returns the TypedData field if non-nil, zero value otherwise.

### GetTypedDataOk

`func (o *NotifyWeb3AuthRequest) GetTypedDataOk() (*string, bool)`

GetTypedDataOk returns a tuple with the TypedData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypedData

`func (o *NotifyWeb3AuthRequest) SetTypedData(v string)`

SetTypedData sets TypedData field to given value.

### HasTypedData

`func (o *NotifyWeb3AuthRequest) HasTypedData() bool`

HasTypedData returns a boolean if a field has been set.

### GetSignature

`func (o *NotifyWeb3AuthRequest) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *NotifyWeb3AuthRequest) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *NotifyWeb3AuthRequest) SetSignature(v string)`

SetSignature sets Signature field to given value.


### GetWalletType

`func (o *NotifyWeb3AuthRequest) GetWalletType() string`

GetWalletType returns the WalletType field if non-nil, zero value otherwise.

### GetWalletTypeOk

`func (o *NotifyWeb3AuthRequest) GetWalletTypeOk() (*string, bool)`

GetWalletTypeOk returns a tuple with the WalletType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletType

`func (o *NotifyWeb3AuthRequest) SetWalletType(v string)`

SetWalletType sets WalletType field to given value.

### HasWalletType

`func (o *NotifyWeb3AuthRequest) HasWalletType() bool`

HasWalletType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


