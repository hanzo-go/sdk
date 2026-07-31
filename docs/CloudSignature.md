# CloudSignature

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | Address is the wallet&#39;s on-chain address, the one this signature recovers to. | [optional] 
**Digest** | Pointer to **string** | Digest is the 32-byte digest that was signed, hex with an 0x prefix. | [optional] 
**Signature** | Pointer to **string** | Signature is the 65-byte secp256k1 signature, hex with an 0x prefix. | [optional] 
**WalletId** | Pointer to **string** | WalletID is the wallet that signed. | [optional] 

## Methods

### NewCloudSignature

`func NewCloudSignature() *CloudSignature`

NewCloudSignature instantiates a new CloudSignature object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudSignatureWithDefaults

`func NewCloudSignatureWithDefaults() *CloudSignature`

NewCloudSignatureWithDefaults instantiates a new CloudSignature object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *CloudSignature) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *CloudSignature) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *CloudSignature) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *CloudSignature) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetDigest

`func (o *CloudSignature) GetDigest() string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *CloudSignature) GetDigestOk() (*string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *CloudSignature) SetDigest(v string)`

SetDigest sets Digest field to given value.

### HasDigest

`func (o *CloudSignature) HasDigest() bool`

HasDigest returns a boolean if a field has been set.

### GetSignature

`func (o *CloudSignature) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *CloudSignature) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *CloudSignature) SetSignature(v string)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *CloudSignature) HasSignature() bool`

HasSignature returns a boolean if a field has been set.

### GetWalletId

`func (o *CloudSignature) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *CloudSignature) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *CloudSignature) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.

### HasWalletId

`func (o *CloudSignature) HasWalletId() bool`

HasWalletId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


