# MarketPools

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Chain** | Pointer to **string** |  | [optional] 
**Pools** | Pointer to [**[]MarketPool**](MarketPool.md) | Pools is &#x60;[]&#x60; where the chain has none and &#x60;null&#x60; where the read failed.  The two are different sentences and the wire says which: an empty ARRAY is the indexer answering that nothing is deployed there, and &#x60;null&#x60; is nobody having answered. &#x60;omitempty&#x60; would collapse both to an absent key — which is the exact flattening the reach beside it exists to prevent, reintroduced one struct tag lower down. | [optional] 
**Reach** | Pointer to [**MarketReach**](MarketReach.md) |  | [optional] 

## Methods

### NewMarketPools

`func NewMarketPools() *MarketPools`

NewMarketPools instantiates a new MarketPools object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketPoolsWithDefaults

`func NewMarketPoolsWithDefaults() *MarketPools`

NewMarketPoolsWithDefaults instantiates a new MarketPools object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChain

`func (o *MarketPools) GetChain() string`

GetChain returns the Chain field if non-nil, zero value otherwise.

### GetChainOk

`func (o *MarketPools) GetChainOk() (*string, bool)`

GetChainOk returns a tuple with the Chain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChain

`func (o *MarketPools) SetChain(v string)`

SetChain sets Chain field to given value.

### HasChain

`func (o *MarketPools) HasChain() bool`

HasChain returns a boolean if a field has been set.

### GetPools

`func (o *MarketPools) GetPools() []MarketPool`

GetPools returns the Pools field if non-nil, zero value otherwise.

### GetPoolsOk

`func (o *MarketPools) GetPoolsOk() (*[]MarketPool, bool)`

GetPoolsOk returns a tuple with the Pools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPools

`func (o *MarketPools) SetPools(v []MarketPool)`

SetPools sets Pools field to given value.

### HasPools

`func (o *MarketPools) HasPools() bool`

HasPools returns a boolean if a field has been set.

### GetReach

`func (o *MarketPools) GetReach() MarketReach`

GetReach returns the Reach field if non-nil, zero value otherwise.

### GetReachOk

`func (o *MarketPools) GetReachOk() (*MarketReach, bool)`

GetReachOk returns a tuple with the Reach field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReach

`func (o *MarketPools) SetReach(v MarketReach)`

SetReach sets Reach field to given value.

### HasReach

`func (o *MarketPools) HasReach() bool`

HasReach returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


