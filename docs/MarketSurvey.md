# MarketSurvey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Carries** | Pointer to [**[]MarketPrecompile**](MarketPrecompile.md) | Carries is all four where the node answered and &#x60;null&#x60; where it did not. It is never a short list: a partial read reporting three would let a reader count them and conclude the chain is missing one. | [optional] 
**Chain** | Pointer to **string** |  | [optional] 
**Reach** | Pointer to [**MarketReach**](MarketReach.md) |  | [optional] 
**Rpc** | Pointer to **string** | RPC is the endpoint that was asked, so an answer names where it came from. | [optional] 

## Methods

### NewMarketSurvey

`func NewMarketSurvey() *MarketSurvey`

NewMarketSurvey instantiates a new MarketSurvey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketSurveyWithDefaults

`func NewMarketSurveyWithDefaults() *MarketSurvey`

NewMarketSurveyWithDefaults instantiates a new MarketSurvey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCarries

`func (o *MarketSurvey) GetCarries() []MarketPrecompile`

GetCarries returns the Carries field if non-nil, zero value otherwise.

### GetCarriesOk

`func (o *MarketSurvey) GetCarriesOk() (*[]MarketPrecompile, bool)`

GetCarriesOk returns a tuple with the Carries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarries

`func (o *MarketSurvey) SetCarries(v []MarketPrecompile)`

SetCarries sets Carries field to given value.

### HasCarries

`func (o *MarketSurvey) HasCarries() bool`

HasCarries returns a boolean if a field has been set.

### GetChain

`func (o *MarketSurvey) GetChain() string`

GetChain returns the Chain field if non-nil, zero value otherwise.

### GetChainOk

`func (o *MarketSurvey) GetChainOk() (*string, bool)`

GetChainOk returns a tuple with the Chain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChain

`func (o *MarketSurvey) SetChain(v string)`

SetChain sets Chain field to given value.

### HasChain

`func (o *MarketSurvey) HasChain() bool`

HasChain returns a boolean if a field has been set.

### GetReach

`func (o *MarketSurvey) GetReach() MarketReach`

GetReach returns the Reach field if non-nil, zero value otherwise.

### GetReachOk

`func (o *MarketSurvey) GetReachOk() (*MarketReach, bool)`

GetReachOk returns a tuple with the Reach field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReach

`func (o *MarketSurvey) SetReach(v MarketReach)`

SetReach sets Reach field to given value.

### HasReach

`func (o *MarketSurvey) HasReach() bool`

HasReach returns a boolean if a field has been set.

### GetRpc

`func (o *MarketSurvey) GetRpc() string`

GetRpc returns the Rpc field if non-nil, zero value otherwise.

### GetRpcOk

`func (o *MarketSurvey) GetRpcOk() (*string, bool)`

GetRpcOk returns a tuple with the Rpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpc

`func (o *MarketSurvey) SetRpc(v string)`

SetRpc sets Rpc field to given value.

### HasRpc

`func (o *MarketSurvey) HasRpc() bool`

HasRpc returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


