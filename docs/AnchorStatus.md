# AnchorStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChainId** | Pointer to **int32** |  | [optional] 
**Contract** | Pointer to **string** |  | [optional] 
**CurrentRoot** | Pointer to **string** | 0x… root of the journal as it stands now | [optional] 
**EntryCount** | Pointer to **int32** |  | [optional] 
**LastAt** | Pointer to **int32** |  | [optional] 
**LastBlock** | Pointer to **int32** |  | [optional] 
**LastRoot** | Pointer to **string** | The last committed on-chain anchor (nil-fields until the first successful submit). | [optional] 
**LastTxHash** | Pointer to **string** |  | [optional] 
**Note** | Pointer to **string** |  | [optional] 
**RpcConfigured** | Pointer to **bool** |  | [optional] 
**SignerConfigured** | Pointer to **bool** |  | [optional] 
**Status** | Pointer to **string** | pending | anchored | error | [optional] 
**Synced** | Pointer to **bool** | true when the last anchored root &#x3D;&#x3D; the current root | [optional] 

## Methods

### NewAnchorStatus

`func NewAnchorStatus() *AnchorStatus`

NewAnchorStatus instantiates a new AnchorStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnchorStatusWithDefaults

`func NewAnchorStatusWithDefaults() *AnchorStatus`

NewAnchorStatusWithDefaults instantiates a new AnchorStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChainId

`func (o *AnchorStatus) GetChainId() int32`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *AnchorStatus) GetChainIdOk() (*int32, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *AnchorStatus) SetChainId(v int32)`

SetChainId sets ChainId field to given value.

### HasChainId

`func (o *AnchorStatus) HasChainId() bool`

HasChainId returns a boolean if a field has been set.

### GetContract

`func (o *AnchorStatus) GetContract() string`

GetContract returns the Contract field if non-nil, zero value otherwise.

### GetContractOk

`func (o *AnchorStatus) GetContractOk() (*string, bool)`

GetContractOk returns a tuple with the Contract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContract

`func (o *AnchorStatus) SetContract(v string)`

SetContract sets Contract field to given value.

### HasContract

`func (o *AnchorStatus) HasContract() bool`

HasContract returns a boolean if a field has been set.

### GetCurrentRoot

`func (o *AnchorStatus) GetCurrentRoot() string`

GetCurrentRoot returns the CurrentRoot field if non-nil, zero value otherwise.

### GetCurrentRootOk

`func (o *AnchorStatus) GetCurrentRootOk() (*string, bool)`

GetCurrentRootOk returns a tuple with the CurrentRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentRoot

`func (o *AnchorStatus) SetCurrentRoot(v string)`

SetCurrentRoot sets CurrentRoot field to given value.

### HasCurrentRoot

`func (o *AnchorStatus) HasCurrentRoot() bool`

HasCurrentRoot returns a boolean if a field has been set.

### GetEntryCount

`func (o *AnchorStatus) GetEntryCount() int32`

GetEntryCount returns the EntryCount field if non-nil, zero value otherwise.

### GetEntryCountOk

`func (o *AnchorStatus) GetEntryCountOk() (*int32, bool)`

GetEntryCountOk returns a tuple with the EntryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryCount

`func (o *AnchorStatus) SetEntryCount(v int32)`

SetEntryCount sets EntryCount field to given value.

### HasEntryCount

`func (o *AnchorStatus) HasEntryCount() bool`

HasEntryCount returns a boolean if a field has been set.

### GetLastAt

`func (o *AnchorStatus) GetLastAt() int32`

GetLastAt returns the LastAt field if non-nil, zero value otherwise.

### GetLastAtOk

`func (o *AnchorStatus) GetLastAtOk() (*int32, bool)`

GetLastAtOk returns a tuple with the LastAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAt

`func (o *AnchorStatus) SetLastAt(v int32)`

SetLastAt sets LastAt field to given value.

### HasLastAt

`func (o *AnchorStatus) HasLastAt() bool`

HasLastAt returns a boolean if a field has been set.

### GetLastBlock

`func (o *AnchorStatus) GetLastBlock() int32`

GetLastBlock returns the LastBlock field if non-nil, zero value otherwise.

### GetLastBlockOk

`func (o *AnchorStatus) GetLastBlockOk() (*int32, bool)`

GetLastBlockOk returns a tuple with the LastBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastBlock

`func (o *AnchorStatus) SetLastBlock(v int32)`

SetLastBlock sets LastBlock field to given value.

### HasLastBlock

`func (o *AnchorStatus) HasLastBlock() bool`

HasLastBlock returns a boolean if a field has been set.

### GetLastRoot

`func (o *AnchorStatus) GetLastRoot() string`

GetLastRoot returns the LastRoot field if non-nil, zero value otherwise.

### GetLastRootOk

`func (o *AnchorStatus) GetLastRootOk() (*string, bool)`

GetLastRootOk returns a tuple with the LastRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRoot

`func (o *AnchorStatus) SetLastRoot(v string)`

SetLastRoot sets LastRoot field to given value.

### HasLastRoot

`func (o *AnchorStatus) HasLastRoot() bool`

HasLastRoot returns a boolean if a field has been set.

### GetLastTxHash

`func (o *AnchorStatus) GetLastTxHash() string`

GetLastTxHash returns the LastTxHash field if non-nil, zero value otherwise.

### GetLastTxHashOk

`func (o *AnchorStatus) GetLastTxHashOk() (*string, bool)`

GetLastTxHashOk returns a tuple with the LastTxHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTxHash

`func (o *AnchorStatus) SetLastTxHash(v string)`

SetLastTxHash sets LastTxHash field to given value.

### HasLastTxHash

`func (o *AnchorStatus) HasLastTxHash() bool`

HasLastTxHash returns a boolean if a field has been set.

### GetNote

`func (o *AnchorStatus) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *AnchorStatus) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *AnchorStatus) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *AnchorStatus) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetRpcConfigured

`func (o *AnchorStatus) GetRpcConfigured() bool`

GetRpcConfigured returns the RpcConfigured field if non-nil, zero value otherwise.

### GetRpcConfiguredOk

`func (o *AnchorStatus) GetRpcConfiguredOk() (*bool, bool)`

GetRpcConfiguredOk returns a tuple with the RpcConfigured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpcConfigured

`func (o *AnchorStatus) SetRpcConfigured(v bool)`

SetRpcConfigured sets RpcConfigured field to given value.

### HasRpcConfigured

`func (o *AnchorStatus) HasRpcConfigured() bool`

HasRpcConfigured returns a boolean if a field has been set.

### GetSignerConfigured

`func (o *AnchorStatus) GetSignerConfigured() bool`

GetSignerConfigured returns the SignerConfigured field if non-nil, zero value otherwise.

### GetSignerConfiguredOk

`func (o *AnchorStatus) GetSignerConfiguredOk() (*bool, bool)`

GetSignerConfiguredOk returns a tuple with the SignerConfigured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignerConfigured

`func (o *AnchorStatus) SetSignerConfigured(v bool)`

SetSignerConfigured sets SignerConfigured field to given value.

### HasSignerConfigured

`func (o *AnchorStatus) HasSignerConfigured() bool`

HasSignerConfigured returns a boolean if a field has been set.

### GetStatus

`func (o *AnchorStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AnchorStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AnchorStatus) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AnchorStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSynced

`func (o *AnchorStatus) GetSynced() bool`

GetSynced returns the Synced field if non-nil, zero value otherwise.

### GetSyncedOk

`func (o *AnchorStatus) GetSyncedOk() (*bool, bool)`

GetSyncedOk returns a tuple with the Synced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynced

`func (o *AnchorStatus) SetSynced(v bool)`

SetSynced sets Synced field to given value.

### HasSynced

`func (o *AnchorStatus) HasSynced() bool`

HasSynced returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


