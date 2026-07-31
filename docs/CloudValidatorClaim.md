# CloudValidatorClaim

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nonce** | Pointer to **string** | Nonce is the value GET /v1/validators/challenge issued for this slot. | [optional] 
**Signature** | Pointer to **string** | Signature is the wallet&#39;s personal_sign over the challenge message, hex with a 0x prefix. | [optional] 
**TokenId** | Pointer to **int32** | TokenID is the Validator-tier GenesisNFT token id being claimed. It IS the validator slot. | [optional] 

## Methods

### NewCloudValidatorClaim

`func NewCloudValidatorClaim() *CloudValidatorClaim`

NewCloudValidatorClaim instantiates a new CloudValidatorClaim object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudValidatorClaimWithDefaults

`func NewCloudValidatorClaimWithDefaults() *CloudValidatorClaim`

NewCloudValidatorClaimWithDefaults instantiates a new CloudValidatorClaim object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNonce

`func (o *CloudValidatorClaim) GetNonce() string`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *CloudValidatorClaim) GetNonceOk() (*string, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *CloudValidatorClaim) SetNonce(v string)`

SetNonce sets Nonce field to given value.

### HasNonce

`func (o *CloudValidatorClaim) HasNonce() bool`

HasNonce returns a boolean if a field has been set.

### GetSignature

`func (o *CloudValidatorClaim) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *CloudValidatorClaim) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *CloudValidatorClaim) SetSignature(v string)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *CloudValidatorClaim) HasSignature() bool`

HasSignature returns a boolean if a field has been set.

### GetTokenId

`func (o *CloudValidatorClaim) GetTokenId() int32`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *CloudValidatorClaim) GetTokenIdOk() (*int32, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *CloudValidatorClaim) SetTokenId(v int32)`

SetTokenId sets TokenId field to given value.

### HasTokenId

`func (o *CloudValidatorClaim) HasTokenId() bool`

HasTokenId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


