# CompanyGenesis

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**At** | Pointer to **int64** | At is the unix second the genesis root was computed. | [optional] 
**Block** | Pointer to **int32** | Block is the L1 block the anchoring transaction landed in. Set only once the receipt has been read; absent otherwise. | [optional] 
**ChainId** | Pointer to **int64** | ChainID is the EVM chain the root is committed to — the Hanzo L1 by default. | [optional] 
**Note** | Pointer to **string** | Note explains an unanchored genesis honestly — anchor wiring absent, or the submit error — rather than reporting a commit that did not happen. | [optional] 
**Root** | Pointer to **string** | Root is the 0x-prefixed keccak256 root of the founding allocation. It is ALWAYS computed, whether or not the on-chain anchor is wired, because the root is the tamper-evident witness. | [optional] 
**Status** | Pointer to **string** | Status is pending (root computed, not yet on-chain) or anchored (committed). | [optional] 
**TxHash** | Pointer to **string** | TxHash is the L1 transaction hash of the anchoring commit. Empty until anchored. | [optional] 

## Methods

### NewCompanyGenesis

`func NewCompanyGenesis() *CompanyGenesis`

NewCompanyGenesis instantiates a new CompanyGenesis object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyGenesisWithDefaults

`func NewCompanyGenesisWithDefaults() *CompanyGenesis`

NewCompanyGenesisWithDefaults instantiates a new CompanyGenesis object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAt

`func (o *CompanyGenesis) GetAt() int64`

GetAt returns the At field if non-nil, zero value otherwise.

### GetAtOk

`func (o *CompanyGenesis) GetAtOk() (*int64, bool)`

GetAtOk returns a tuple with the At field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAt

`func (o *CompanyGenesis) SetAt(v int64)`

SetAt sets At field to given value.

### HasAt

`func (o *CompanyGenesis) HasAt() bool`

HasAt returns a boolean if a field has been set.

### GetBlock

`func (o *CompanyGenesis) GetBlock() int32`

GetBlock returns the Block field if non-nil, zero value otherwise.

### GetBlockOk

`func (o *CompanyGenesis) GetBlockOk() (*int32, bool)`

GetBlockOk returns a tuple with the Block field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlock

`func (o *CompanyGenesis) SetBlock(v int32)`

SetBlock sets Block field to given value.

### HasBlock

`func (o *CompanyGenesis) HasBlock() bool`

HasBlock returns a boolean if a field has been set.

### GetChainId

`func (o *CompanyGenesis) GetChainId() int64`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *CompanyGenesis) GetChainIdOk() (*int64, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *CompanyGenesis) SetChainId(v int64)`

SetChainId sets ChainId field to given value.

### HasChainId

`func (o *CompanyGenesis) HasChainId() bool`

HasChainId returns a boolean if a field has been set.

### GetNote

`func (o *CompanyGenesis) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *CompanyGenesis) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *CompanyGenesis) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *CompanyGenesis) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetRoot

`func (o *CompanyGenesis) GetRoot() string`

GetRoot returns the Root field if non-nil, zero value otherwise.

### GetRootOk

`func (o *CompanyGenesis) GetRootOk() (*string, bool)`

GetRootOk returns a tuple with the Root field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoot

`func (o *CompanyGenesis) SetRoot(v string)`

SetRoot sets Root field to given value.

### HasRoot

`func (o *CompanyGenesis) HasRoot() bool`

HasRoot returns a boolean if a field has been set.

### GetStatus

`func (o *CompanyGenesis) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CompanyGenesis) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CompanyGenesis) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CompanyGenesis) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTxHash

`func (o *CompanyGenesis) GetTxHash() string`

GetTxHash returns the TxHash field if non-nil, zero value otherwise.

### GetTxHashOk

`func (o *CompanyGenesis) GetTxHashOk() (*string, bool)`

GetTxHashOk returns a tuple with the TxHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxHash

`func (o *CompanyGenesis) SetTxHash(v string)`

SetTxHash sets TxHash field to given value.

### HasTxHash

`func (o *CompanyGenesis) HasTxHash() bool`

HasTxHash returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


