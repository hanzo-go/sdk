# CryptoOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Chains** | Pointer to **[]string** |  | [optional] 
**Tokens** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCryptoOptions

`func NewCryptoOptions() *CryptoOptions`

NewCryptoOptions instantiates a new CryptoOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCryptoOptionsWithDefaults

`func NewCryptoOptionsWithDefaults() *CryptoOptions`

NewCryptoOptionsWithDefaults instantiates a new CryptoOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChains

`func (o *CryptoOptions) GetChains() []string`

GetChains returns the Chains field if non-nil, zero value otherwise.

### GetChainsOk

`func (o *CryptoOptions) GetChainsOk() (*[]string, bool)`

GetChainsOk returns a tuple with the Chains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChains

`func (o *CryptoOptions) SetChains(v []string)`

SetChains sets Chains field to given value.

### HasChains

`func (o *CryptoOptions) HasChains() bool`

HasChains returns a boolean if a field has been set.

### GetTokens

`func (o *CryptoOptions) GetTokens() []string`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *CryptoOptions) GetTokensOk() (*[]string, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *CryptoOptions) SetTokens(v []string)`

SetTokens sets Tokens field to given value.

### HasTokens

`func (o *CryptoOptions) HasTokens() bool`

HasTokens returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


