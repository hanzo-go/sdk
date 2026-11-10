# BillingCryptoOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assets** | Pointer to [**[]BillingCryptoPair**](BillingCryptoPair.md) | Assets is every (chain, token) pair the deposit-address rail accepts. Chains and Tokens are its two axes: a chain and a token each on offer are not a pair on offer, so a picker that crosses them can name an asset the rail will refuse. | [optional] 
**Chains** | Pointer to **[]string** |  | [optional] 
**Gateway** | Pointer to **[]string** |  | [optional] 
**Tokens** | Pointer to **[]string** |  | [optional] 

## Methods

### NewBillingCryptoOptions

`func NewBillingCryptoOptions() *BillingCryptoOptions`

NewBillingCryptoOptions instantiates a new BillingCryptoOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingCryptoOptionsWithDefaults

`func NewBillingCryptoOptionsWithDefaults() *BillingCryptoOptions`

NewBillingCryptoOptionsWithDefaults instantiates a new BillingCryptoOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssets

`func (o *BillingCryptoOptions) GetAssets() []BillingCryptoPair`

GetAssets returns the Assets field if non-nil, zero value otherwise.

### GetAssetsOk

`func (o *BillingCryptoOptions) GetAssetsOk() (*[]BillingCryptoPair, bool)`

GetAssetsOk returns a tuple with the Assets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssets

`func (o *BillingCryptoOptions) SetAssets(v []BillingCryptoPair)`

SetAssets sets Assets field to given value.

### HasAssets

`func (o *BillingCryptoOptions) HasAssets() bool`

HasAssets returns a boolean if a field has been set.

### GetChains

`func (o *BillingCryptoOptions) GetChains() []string`

GetChains returns the Chains field if non-nil, zero value otherwise.

### GetChainsOk

`func (o *BillingCryptoOptions) GetChainsOk() (*[]string, bool)`

GetChainsOk returns a tuple with the Chains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChains

`func (o *BillingCryptoOptions) SetChains(v []string)`

SetChains sets Chains field to given value.

### HasChains

`func (o *BillingCryptoOptions) HasChains() bool`

HasChains returns a boolean if a field has been set.

### GetGateway

`func (o *BillingCryptoOptions) GetGateway() []string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *BillingCryptoOptions) GetGatewayOk() (*[]string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *BillingCryptoOptions) SetGateway(v []string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *BillingCryptoOptions) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetTokens

`func (o *BillingCryptoOptions) GetTokens() []string`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *BillingCryptoOptions) GetTokensOk() (*[]string, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *BillingCryptoOptions) SetTokens(v []string)`

SetTokens sets Tokens field to given value.

### HasTokens

`func (o *BillingCryptoOptions) HasTokens() bool`

HasTokens returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


