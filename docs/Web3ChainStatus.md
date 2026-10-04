# Web3ChainStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChainId** | Pointer to **int64** |  | [optional] 
**Height** | Pointer to **int64** | Height is the latest block, omitted when the chain did not answer rather than reported as zero — a zero height is a real value on a fresh chain. | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Live** | Pointer to **bool** | Live is whether the upstream answered eth_blockNumber. | [optional] 
**Name** | Pointer to **string** |  | [optional] 

## Methods

### NewWeb3ChainStatus

`func NewWeb3ChainStatus() *Web3ChainStatus`

NewWeb3ChainStatus instantiates a new Web3ChainStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWeb3ChainStatusWithDefaults

`func NewWeb3ChainStatusWithDefaults() *Web3ChainStatus`

NewWeb3ChainStatusWithDefaults instantiates a new Web3ChainStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChainId

`func (o *Web3ChainStatus) GetChainId() int64`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *Web3ChainStatus) GetChainIdOk() (*int64, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *Web3ChainStatus) SetChainId(v int64)`

SetChainId sets ChainId field to given value.

### HasChainId

`func (o *Web3ChainStatus) HasChainId() bool`

HasChainId returns a boolean if a field has been set.

### GetHeight

`func (o *Web3ChainStatus) GetHeight() int64`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *Web3ChainStatus) GetHeightOk() (*int64, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *Web3ChainStatus) SetHeight(v int64)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *Web3ChainStatus) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetId

`func (o *Web3ChainStatus) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Web3ChainStatus) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Web3ChainStatus) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Web3ChainStatus) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLive

`func (o *Web3ChainStatus) GetLive() bool`

GetLive returns the Live field if non-nil, zero value otherwise.

### GetLiveOk

`func (o *Web3ChainStatus) GetLiveOk() (*bool, bool)`

GetLiveOk returns a tuple with the Live field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLive

`func (o *Web3ChainStatus) SetLive(v bool)`

SetLive sets Live field to given value.

### HasLive

`func (o *Web3ChainStatus) HasLive() bool`

HasLive returns a boolean if a field has been set.

### GetName

`func (o *Web3ChainStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Web3ChainStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Web3ChainStatus) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Web3ChainStatus) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


