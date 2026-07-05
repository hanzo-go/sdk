# GraphIndexer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Chain** | Pointer to **string** |  | [optional] 
**Network** | Pointer to **string** | mainnet | testnet | devnet | [optional] 
**Height** | Pointer to **string** | Latest indexed block height | [optional] 
**Lag** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewGraphIndexer

`func NewGraphIndexer() *GraphIndexer`

NewGraphIndexer instantiates a new GraphIndexer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphIndexerWithDefaults

`func NewGraphIndexerWithDefaults() *GraphIndexer`

NewGraphIndexerWithDefaults instantiates a new GraphIndexer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GraphIndexer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GraphIndexer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GraphIndexer) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GraphIndexer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetChain

`func (o *GraphIndexer) GetChain() string`

GetChain returns the Chain field if non-nil, zero value otherwise.

### GetChainOk

`func (o *GraphIndexer) GetChainOk() (*string, bool)`

GetChainOk returns a tuple with the Chain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChain

`func (o *GraphIndexer) SetChain(v string)`

SetChain sets Chain field to given value.

### HasChain

`func (o *GraphIndexer) HasChain() bool`

HasChain returns a boolean if a field has been set.

### GetNetwork

`func (o *GraphIndexer) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *GraphIndexer) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *GraphIndexer) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *GraphIndexer) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetHeight

`func (o *GraphIndexer) GetHeight() string`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *GraphIndexer) GetHeightOk() (*string, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *GraphIndexer) SetHeight(v string)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *GraphIndexer) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetLag

`func (o *GraphIndexer) GetLag() string`

GetLag returns the Lag field if non-nil, zero value otherwise.

### GetLagOk

`func (o *GraphIndexer) GetLagOk() (*string, bool)`

GetLagOk returns a tuple with the Lag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLag

`func (o *GraphIndexer) SetLag(v string)`

SetLag sets Lag field to given value.

### HasLag

`func (o *GraphIndexer) HasLag() bool`

HasLag returns a boolean if a field has been set.

### GetStatus

`func (o *GraphIndexer) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GraphIndexer) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GraphIndexer) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GraphIndexer) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GraphIndexer) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GraphIndexer) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GraphIndexer) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GraphIndexer) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


