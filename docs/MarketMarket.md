# MarketMarket

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amm** | Pointer to **bool** | Amm reports whether an automated market maker is deployed on this chain, which is the registry&#39;s factory addresses being present and not the indexer having rows. The two disagree in exactly the interesting case: a chain with a factory and nothing traded yet is a live venue with no history, and a chain with neither has no venue at all. | [optional] 
**Coin** | Pointer to **string** |  | [optional] 
**Day** | Pointer to [**MarketDay**](MarketDay.md) |  | [optional] 
**Factory** | Pointer to **map[string]string** | Factory is the AMM&#39;s factory contracts, by generation, omitted where none is deployed. Addresses come from the registry because that is what the indexer itself ingested from; anything else is a second copy free to drift. | [optional] 
**Figures** | Pointer to [**MarketFigures**](MarketFigures.md) | Figures is the chain&#39;s whole market maker, and Day its most recent active one. Both are absent unless Reach says Read, so a caller cannot mistake a zero this process never received for one the indexer computed. Day is also absent on a chain that has never traded — which reach reports as Read, so the two absences are told apart by the state beside them and never by the gap itself. | [optional] 
**Graph** | Pointer to **string** | Graph is where this chain&#39;s indexer answers, empty where it has none. | [optional] 
**Id** | Pointer to **int64** | ID is the EVM chain id, which is what a wallet must agree with. | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Reach** | Pointer to [**MarketReach**](MarketReach.md) | Reach is how far the read of this chain&#39;s FIGURES got — its own, so one indexer being down describes one row and leaves the others to answer. | [optional] 
**Rpc** | Pointer to **string** | RPC is the chain&#39;s PUBLIC JSON-RPC, empty where the registry names only a route this process happens to have. The registry&#39;s own &#x60;rpc&#x60; field is the INDEXER&#39;s route to the node and is sometimes inside its cluster — plain HTTP on a &#x60;.svc.cluster.local&#x60; name — which is reachable from the indexer, from nothing else, and from no browser. Publishing that as the chain&#39;s endpoint hands every caller an address that cannot answer them. See [endpoint]. | [optional] 
**Slug** | Pointer to **string** | Slug is the chain&#39;s word in every indexer path — &#x60;cchain&#x60;, &#x60;zoo&#x60;. It is the value a caller passes back as &#x60;chain&#x60;, and it is NOT the chain id: &#x60;96369&#x60;, &#x60;C&#x60; and &#x60;c-chain&#x60; all answer 404 in that position. | [optional] 

## Methods

### NewMarketMarket

`func NewMarketMarket() *MarketMarket`

NewMarketMarket instantiates a new MarketMarket object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketMarketWithDefaults

`func NewMarketMarketWithDefaults() *MarketMarket`

NewMarketMarketWithDefaults instantiates a new MarketMarket object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmm

`func (o *MarketMarket) GetAmm() bool`

GetAmm returns the Amm field if non-nil, zero value otherwise.

### GetAmmOk

`func (o *MarketMarket) GetAmmOk() (*bool, bool)`

GetAmmOk returns a tuple with the Amm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmm

`func (o *MarketMarket) SetAmm(v bool)`

SetAmm sets Amm field to given value.

### HasAmm

`func (o *MarketMarket) HasAmm() bool`

HasAmm returns a boolean if a field has been set.

### GetCoin

`func (o *MarketMarket) GetCoin() string`

GetCoin returns the Coin field if non-nil, zero value otherwise.

### GetCoinOk

`func (o *MarketMarket) GetCoinOk() (*string, bool)`

GetCoinOk returns a tuple with the Coin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoin

`func (o *MarketMarket) SetCoin(v string)`

SetCoin sets Coin field to given value.

### HasCoin

`func (o *MarketMarket) HasCoin() bool`

HasCoin returns a boolean if a field has been set.

### GetDay

`func (o *MarketMarket) GetDay() MarketDay`

GetDay returns the Day field if non-nil, zero value otherwise.

### GetDayOk

`func (o *MarketMarket) GetDayOk() (*MarketDay, bool)`

GetDayOk returns a tuple with the Day field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDay

`func (o *MarketMarket) SetDay(v MarketDay)`

SetDay sets Day field to given value.

### HasDay

`func (o *MarketMarket) HasDay() bool`

HasDay returns a boolean if a field has been set.

### GetFactory

`func (o *MarketMarket) GetFactory() map[string]string`

GetFactory returns the Factory field if non-nil, zero value otherwise.

### GetFactoryOk

`func (o *MarketMarket) GetFactoryOk() (*map[string]string, bool)`

GetFactoryOk returns a tuple with the Factory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactory

`func (o *MarketMarket) SetFactory(v map[string]string)`

SetFactory sets Factory field to given value.

### HasFactory

`func (o *MarketMarket) HasFactory() bool`

HasFactory returns a boolean if a field has been set.

### GetFigures

`func (o *MarketMarket) GetFigures() MarketFigures`

GetFigures returns the Figures field if non-nil, zero value otherwise.

### GetFiguresOk

`func (o *MarketMarket) GetFiguresOk() (*MarketFigures, bool)`

GetFiguresOk returns a tuple with the Figures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFigures

`func (o *MarketMarket) SetFigures(v MarketFigures)`

SetFigures sets Figures field to given value.

### HasFigures

`func (o *MarketMarket) HasFigures() bool`

HasFigures returns a boolean if a field has been set.

### GetGraph

`func (o *MarketMarket) GetGraph() string`

GetGraph returns the Graph field if non-nil, zero value otherwise.

### GetGraphOk

`func (o *MarketMarket) GetGraphOk() (*string, bool)`

GetGraphOk returns a tuple with the Graph field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraph

`func (o *MarketMarket) SetGraph(v string)`

SetGraph sets Graph field to given value.

### HasGraph

`func (o *MarketMarket) HasGraph() bool`

HasGraph returns a boolean if a field has been set.

### GetId

`func (o *MarketMarket) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MarketMarket) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MarketMarket) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *MarketMarket) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *MarketMarket) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MarketMarket) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MarketMarket) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MarketMarket) HasName() bool`

HasName returns a boolean if a field has been set.

### GetReach

`func (o *MarketMarket) GetReach() MarketReach`

GetReach returns the Reach field if non-nil, zero value otherwise.

### GetReachOk

`func (o *MarketMarket) GetReachOk() (*MarketReach, bool)`

GetReachOk returns a tuple with the Reach field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReach

`func (o *MarketMarket) SetReach(v MarketReach)`

SetReach sets Reach field to given value.

### HasReach

`func (o *MarketMarket) HasReach() bool`

HasReach returns a boolean if a field has been set.

### GetRpc

`func (o *MarketMarket) GetRpc() string`

GetRpc returns the Rpc field if non-nil, zero value otherwise.

### GetRpcOk

`func (o *MarketMarket) GetRpcOk() (*string, bool)`

GetRpcOk returns a tuple with the Rpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpc

`func (o *MarketMarket) SetRpc(v string)`

SetRpc sets Rpc field to given value.

### HasRpc

`func (o *MarketMarket) HasRpc() bool`

HasRpc returns a boolean if a field has been set.

### GetSlug

`func (o *MarketMarket) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *MarketMarket) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *MarketMarket) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *MarketMarket) HasSlug() bool`

HasSlug returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


