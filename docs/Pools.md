# Pools

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Chain** | Pointer to **string** |  | [optional] 
**Pools** | Pointer to [**[]Pool**](Pool.md) | Pools is &#x60;[]&#x60; where the chain has none and &#x60;null&#x60; where the read failed.  The two are different sentences and the wire says which: an empty ARRAY is the indexer answering that nothing is deployed there, and &#x60;null&#x60; is nobody having answered. &#x60;omitempty&#x60; would collapse both to an absent key — which is the exact flattening the reach beside it exists to prevent, reintroduced one struct tag lower down. | [optional] 
**Reach** | Pointer to [**Reach**](Reach.md) |  | [optional] 

## Methods

### NewPools

`func NewPools() *Pools`

NewPools instantiates a new Pools object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoolsWithDefaults

`func NewPoolsWithDefaults() *Pools`

NewPoolsWithDefaults instantiates a new Pools object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChain

`func (o *Pools) GetChain() string`

GetChain returns the Chain field if non-nil, zero value otherwise.

### GetChainOk

`func (o *Pools) GetChainOk() (*string, bool)`

GetChainOk returns a tuple with the Chain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChain

`func (o *Pools) SetChain(v string)`

SetChain sets Chain field to given value.

### HasChain

`func (o *Pools) HasChain() bool`

HasChain returns a boolean if a field has been set.

### GetPools

`func (o *Pools) GetPools() []Pool`

GetPools returns the Pools field if non-nil, zero value otherwise.

### GetPoolsOk

`func (o *Pools) GetPoolsOk() (*[]Pool, bool)`

GetPoolsOk returns a tuple with the Pools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPools

`func (o *Pools) SetPools(v []Pool)`

SetPools sets Pools field to given value.

### HasPools

`func (o *Pools) HasPools() bool`

HasPools returns a boolean if a field has been set.

### GetReach

`func (o *Pools) GetReach() Reach`

GetReach returns the Reach field if non-nil, zero value otherwise.

### GetReachOk

`func (o *Pools) GetReachOk() (*Reach, bool)`

GetReachOk returns a tuple with the Reach field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReach

`func (o *Pools) SetReach(v Reach)`

SetReach sets Reach field to given value.

### HasReach

`func (o *Pools) HasReach() bool`

HasReach returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


