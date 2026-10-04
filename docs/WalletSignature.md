# WalletSignature

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | Address is the wallet&#39;s on-chain address, the one this signature recovers to. | [optional] 
**Digest** | Pointer to **string** | Digest is the 32-byte digest that was signed, hex with an 0x prefix. | [optional] 
**Signature** | Pointer to **string** | Signature is the 65-byte secp256k1 signature, hex with an 0x prefix. | [optional] 
**WalletId** | Pointer to **string** | WalletID is the wallet that signed. | [optional] 

## Methods

### NewWalletSignature

`func NewWalletSignature() *WalletSignature`

NewWalletSignature instantiates a new WalletSignature object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWalletSignatureWithDefaults

`func NewWalletSignatureWithDefaults() *WalletSignature`

NewWalletSignatureWithDefaults instantiates a new WalletSignature object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *WalletSignature) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *WalletSignature) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *WalletSignature) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *WalletSignature) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetDigest

`func (o *WalletSignature) GetDigest() string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *WalletSignature) GetDigestOk() (*string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *WalletSignature) SetDigest(v string)`

SetDigest sets Digest field to given value.

### HasDigest

`func (o *WalletSignature) HasDigest() bool`

HasDigest returns a boolean if a field has been set.

### GetSignature

`func (o *WalletSignature) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *WalletSignature) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *WalletSignature) SetSignature(v string)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *WalletSignature) HasSignature() bool`

HasSignature returns a boolean if a field has been set.

### GetWalletId

`func (o *WalletSignature) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *WalletSignature) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *WalletSignature) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.

### HasWalletId

`func (o *WalletSignature) HasWalletId() bool`

HasWalletId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


