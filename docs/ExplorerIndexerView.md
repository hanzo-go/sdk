# ExplorerIndexerView

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

### NewExplorerIndexerView

`func NewExplorerIndexerView() *ExplorerIndexerView`

NewExplorerIndexerView instantiates a new ExplorerIndexerView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExplorerIndexerViewWithDefaults

`func NewExplorerIndexerViewWithDefaults() *ExplorerIndexerView`

NewExplorerIndexerViewWithDefaults instantiates a new ExplorerIndexerView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChain

`func (o *ExplorerIndexerView) GetChain() string`

GetChain returns the Chain field if non-nil, zero value otherwise.

### GetChainOk

`func (o *ExplorerIndexerView) GetChainOk() (*string, bool)`

GetChainOk returns a tuple with the Chain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChain

`func (o *ExplorerIndexerView) SetChain(v string)`

SetChain sets Chain field to given value.

### HasChain

`func (o *ExplorerIndexerView) HasChain() bool`

HasChain returns a boolean if a field has been set.

### GetHeight

`func (o *ExplorerIndexerView) GetHeight() string`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *ExplorerIndexerView) GetHeightOk() (*string, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *ExplorerIndexerView) SetHeight(v string)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *ExplorerIndexerView) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetId

`func (o *ExplorerIndexerView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExplorerIndexerView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExplorerIndexerView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ExplorerIndexerView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLag

`func (o *ExplorerIndexerView) GetLag() string`

GetLag returns the Lag field if non-nil, zero value otherwise.

### GetLagOk

`func (o *ExplorerIndexerView) GetLagOk() (*string, bool)`

GetLagOk returns a tuple with the Lag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLag

`func (o *ExplorerIndexerView) SetLag(v string)`

SetLag sets Lag field to given value.

### HasLag

`func (o *ExplorerIndexerView) HasLag() bool`

HasLag returns a boolean if a field has been set.

### GetNetwork

`func (o *ExplorerIndexerView) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *ExplorerIndexerView) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *ExplorerIndexerView) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *ExplorerIndexerView) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetStatus

`func (o *ExplorerIndexerView) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ExplorerIndexerView) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ExplorerIndexerView) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ExplorerIndexerView) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ExplorerIndexerView) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ExplorerIndexerView) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ExplorerIndexerView) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ExplorerIndexerView) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


