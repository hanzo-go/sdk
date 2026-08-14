# ValidatorClaim

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nonce** | Pointer to **string** | Nonce is the value GET /v1/validators/challenge issued for this slot. | [optional] 
**Signature** | Pointer to **string** | Signature is the wallet&#39;s personal_sign over the challenge message, hex with a 0x prefix. | [optional] 
**TokenId** | Pointer to **int32** | TokenID is the Validator-tier GenesisNFT token id being claimed. It IS the validator slot. | [optional] 

## Methods

### NewValidatorClaim

`func NewValidatorClaim() *ValidatorClaim`

NewValidatorClaim instantiates a new ValidatorClaim object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidatorClaimWithDefaults

`func NewValidatorClaimWithDefaults() *ValidatorClaim`

NewValidatorClaimWithDefaults instantiates a new ValidatorClaim object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNonce

`func (o *ValidatorClaim) GetNonce() string`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *ValidatorClaim) GetNonceOk() (*string, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *ValidatorClaim) SetNonce(v string)`

SetNonce sets Nonce field to given value.

### HasNonce

`func (o *ValidatorClaim) HasNonce() bool`

HasNonce returns a boolean if a field has been set.

### GetSignature

`func (o *ValidatorClaim) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *ValidatorClaim) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *ValidatorClaim) SetSignature(v string)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *ValidatorClaim) HasSignature() bool`

HasSignature returns a boolean if a field has been set.

### GetTokenId

`func (o *ValidatorClaim) GetTokenId() int32`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *ValidatorClaim) GetTokenIdOk() (*int32, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *ValidatorClaim) SetTokenId(v int32)`

SetTokenId sets TokenId field to given value.

### HasTokenId

`func (o *ValidatorClaim) HasTokenId() bool`

HasTokenId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


