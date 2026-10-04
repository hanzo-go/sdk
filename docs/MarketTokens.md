# MarketTokens

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Chain** | Pointer to **string** |  | [optional] 
**Reach** | Pointer to [**MarketReach**](MarketReach.md) |  | [optional] 
**Tokens** | Pointer to [**[]MarketToken**](MarketToken.md) | Tokens is &#x60;[]&#x60; where the indexer holds none and &#x60;null&#x60; where the read failed. | [optional] 

## Methods

### NewMarketTokens

`func NewMarketTokens() *MarketTokens`

NewMarketTokens instantiates a new MarketTokens object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketTokensWithDefaults

`func NewMarketTokensWithDefaults() *MarketTokens`

NewMarketTokensWithDefaults instantiates a new MarketTokens object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChain

`func (o *MarketTokens) GetChain() string`

GetChain returns the Chain field if non-nil, zero value otherwise.

### GetChainOk

`func (o *MarketTokens) GetChainOk() (*string, bool)`

GetChainOk returns a tuple with the Chain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChain

`func (o *MarketTokens) SetChain(v string)`

SetChain sets Chain field to given value.

### HasChain

`func (o *MarketTokens) HasChain() bool`

HasChain returns a boolean if a field has been set.

### GetReach

`func (o *MarketTokens) GetReach() MarketReach`

GetReach returns the Reach field if non-nil, zero value otherwise.

### GetReachOk

`func (o *MarketTokens) GetReachOk() (*MarketReach, bool)`

GetReachOk returns a tuple with the Reach field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReach

`func (o *MarketTokens) SetReach(v MarketReach)`

SetReach sets Reach field to given value.

### HasReach

`func (o *MarketTokens) HasReach() bool`

HasReach returns a boolean if a field has been set.

### GetTokens

`func (o *MarketTokens) GetTokens() []MarketToken`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *MarketTokens) GetTokensOk() (*[]MarketToken, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *MarketTokens) SetTokens(v []MarketToken)`

SetTokens sets Tokens field to given value.

### HasTokens

`func (o *MarketTokens) HasTokens() bool`

HasTokens returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


