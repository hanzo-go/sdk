# MarketplacePayoutChallenge

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | Address is its address, which the signature must recover to. | [optional] 
**Digest** | Pointer to **string** | Digest is the EIP-191 hash of Message, 0x hex: what POST /v1/wallet/{id}/sign signs, for a wallet the platform holds. | [optional] 
**Expires** | Pointer to **int64** | Expires is when the challenge may no longer be signed, unix seconds. | [optional] 
**Message** | Pointer to **string** | Message is the text to sign with the wallet (EIP-191 personal_sign). | [optional] 
**Wallet** | Pointer to **string** | Wallet is the wallet to prove. | [optional] 

## Methods

### NewMarketplacePayoutChallenge

`func NewMarketplacePayoutChallenge() *MarketplacePayoutChallenge`

NewMarketplacePayoutChallenge instantiates a new MarketplacePayoutChallenge object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketplacePayoutChallengeWithDefaults

`func NewMarketplacePayoutChallengeWithDefaults() *MarketplacePayoutChallenge`

NewMarketplacePayoutChallengeWithDefaults instantiates a new MarketplacePayoutChallenge object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *MarketplacePayoutChallenge) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *MarketplacePayoutChallenge) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *MarketplacePayoutChallenge) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *MarketplacePayoutChallenge) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetDigest

`func (o *MarketplacePayoutChallenge) GetDigest() string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *MarketplacePayoutChallenge) GetDigestOk() (*string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *MarketplacePayoutChallenge) SetDigest(v string)`

SetDigest sets Digest field to given value.

### HasDigest

`func (o *MarketplacePayoutChallenge) HasDigest() bool`

HasDigest returns a boolean if a field has been set.

### GetExpires

`func (o *MarketplacePayoutChallenge) GetExpires() int64`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *MarketplacePayoutChallenge) GetExpiresOk() (*int64, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *MarketplacePayoutChallenge) SetExpires(v int64)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *MarketplacePayoutChallenge) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### GetMessage

`func (o *MarketplacePayoutChallenge) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *MarketplacePayoutChallenge) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *MarketplacePayoutChallenge) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *MarketplacePayoutChallenge) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetWallet

`func (o *MarketplacePayoutChallenge) GetWallet() string`

GetWallet returns the Wallet field if non-nil, zero value otherwise.

### GetWalletOk

`func (o *MarketplacePayoutChallenge) GetWalletOk() (*string, bool)`

GetWalletOk returns a tuple with the Wallet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWallet

`func (o *MarketplacePayoutChallenge) SetWallet(v string)`

SetWallet sets Wallet field to given value.

### HasWallet

`func (o *MarketplacePayoutChallenge) HasWallet() bool`

HasWallet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


