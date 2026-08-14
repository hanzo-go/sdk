# IndexerView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Chain** | Pointer to **string** | Chain is the chain this indexer indexes, as the indexer names it. | [optional] 
**Height** | Pointer to **string** | Height is the latest INDEXED block height, as a decimal string. Absent when nothing has been indexed yet. | [optional] 
**Id** | Pointer to **string** | ID identifies the indexer: its chain name, else its chain id, else the brand. | [optional] 
**Lag** | Pointer to **string** | Lag is how far behind the chain HEAD this indexer is. The indexer REST does not expose the head, so it is always absent rather than a fabricated zero. | [optional] 
**Network** | Pointer to **string** | Network is the deployment&#39;s network tier: mainnet, testnet or devnet. | [optional] 
**Status** | Pointer to **string** | Status is \&quot;degraded\&quot; when /health explicitly reports unhealthy, else \&quot;active\&quot;. | [optional] 
**UpdatedAt** | Pointer to **string** | UpdatedAt is the latest indexed block&#39;s timestamp, RFC 3339 UTC. | [optional] 

## Methods

### NewIndexerView

`func NewIndexerView() *IndexerView`

NewIndexerView instantiates a new IndexerView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndexerViewWithDefaults

`func NewIndexerViewWithDefaults() *IndexerView`

NewIndexerViewWithDefaults instantiates a new IndexerView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChain

`func (o *IndexerView) GetChain() string`

GetChain returns the Chain field if non-nil, zero value otherwise.

### GetChainOk

`func (o *IndexerView) GetChainOk() (*string, bool)`

GetChainOk returns a tuple with the Chain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChain

`func (o *IndexerView) SetChain(v string)`

SetChain sets Chain field to given value.

### HasChain

`func (o *IndexerView) HasChain() bool`

HasChain returns a boolean if a field has been set.

### GetHeight

`func (o *IndexerView) GetHeight() string`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *IndexerView) GetHeightOk() (*string, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *IndexerView) SetHeight(v string)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *IndexerView) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetId

`func (o *IndexerView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IndexerView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IndexerView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IndexerView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLag

`func (o *IndexerView) GetLag() string`

GetLag returns the Lag field if non-nil, zero value otherwise.

### GetLagOk

`func (o *IndexerView) GetLagOk() (*string, bool)`

GetLagOk returns a tuple with the Lag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLag

`func (o *IndexerView) SetLag(v string)`

SetLag sets Lag field to given value.

### HasLag

`func (o *IndexerView) HasLag() bool`

HasLag returns a boolean if a field has been set.

### GetNetwork

`func (o *IndexerView) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *IndexerView) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *IndexerView) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *IndexerView) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetStatus

`func (o *IndexerView) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IndexerView) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IndexerView) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *IndexerView) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *IndexerView) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *IndexerView) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *IndexerView) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *IndexerView) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


