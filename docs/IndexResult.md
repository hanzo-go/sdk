# IndexResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Chunks** | Pointer to **int32** | Chunks is how many AST-boundary chunks the repo holds after this pass. | [optional] 
**Files** | Pointer to **int32** | Files is how many files the repo holds after this pass. | [optional] 
**Indexed** | Pointer to **int32** | Indexed is how many files were parsed and written on this pass. | [optional] 
**Pruned** | Pointer to **int32** | Pruned is how many stored files were deleted because prune was set and they were absent from the request. | [optional] 
**Repo** | Pointer to **string** | Repo is the repository that was indexed. | [optional] 
**Semantic** | Pointer to **bool** | Semantic reports whether the semantic tier was available for this pass. When false the index is lexical + symbolic only and hybrid search still works. | [optional] 
**Skipped** | Pointer to **int32** | Skipped is how many files were unchanged by content hash and left alone. | [optional] 
**Symbols** | Pointer to **int32** | Symbols is how many symbol definitions the repo holds after this pass. | [optional] 
**Vectors** | Pointer to **int32** | Vectors is how many of those chunks carry an embedding. | [optional] 

## Methods

### NewIndexResult

`func NewIndexResult() *IndexResult`

NewIndexResult instantiates a new IndexResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndexResultWithDefaults

`func NewIndexResultWithDefaults() *IndexResult`

NewIndexResultWithDefaults instantiates a new IndexResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChunks

`func (o *IndexResult) GetChunks() int32`

GetChunks returns the Chunks field if non-nil, zero value otherwise.

### GetChunksOk

`func (o *IndexResult) GetChunksOk() (*int32, bool)`

GetChunksOk returns a tuple with the Chunks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChunks

`func (o *IndexResult) SetChunks(v int32)`

SetChunks sets Chunks field to given value.

### HasChunks

`func (o *IndexResult) HasChunks() bool`

HasChunks returns a boolean if a field has been set.

### GetFiles

`func (o *IndexResult) GetFiles() int32`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *IndexResult) GetFilesOk() (*int32, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *IndexResult) SetFiles(v int32)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *IndexResult) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetIndexed

`func (o *IndexResult) GetIndexed() int32`

GetIndexed returns the Indexed field if non-nil, zero value otherwise.

### GetIndexedOk

`func (o *IndexResult) GetIndexedOk() (*int32, bool)`

GetIndexedOk returns a tuple with the Indexed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexed

`func (o *IndexResult) SetIndexed(v int32)`

SetIndexed sets Indexed field to given value.

### HasIndexed

`func (o *IndexResult) HasIndexed() bool`

HasIndexed returns a boolean if a field has been set.

### GetPruned

`func (o *IndexResult) GetPruned() int32`

GetPruned returns the Pruned field if non-nil, zero value otherwise.

### GetPrunedOk

`func (o *IndexResult) GetPrunedOk() (*int32, bool)`

GetPrunedOk returns a tuple with the Pruned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPruned

`func (o *IndexResult) SetPruned(v int32)`

SetPruned sets Pruned field to given value.

### HasPruned

`func (o *IndexResult) HasPruned() bool`

HasPruned returns a boolean if a field has been set.

### GetRepo

`func (o *IndexResult) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *IndexResult) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *IndexResult) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *IndexResult) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### GetSemantic

`func (o *IndexResult) GetSemantic() bool`

GetSemantic returns the Semantic field if non-nil, zero value otherwise.

### GetSemanticOk

`func (o *IndexResult) GetSemanticOk() (*bool, bool)`

GetSemanticOk returns a tuple with the Semantic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSemantic

`func (o *IndexResult) SetSemantic(v bool)`

SetSemantic sets Semantic field to given value.

### HasSemantic

`func (o *IndexResult) HasSemantic() bool`

HasSemantic returns a boolean if a field has been set.

### GetSkipped

`func (o *IndexResult) GetSkipped() int32`

GetSkipped returns the Skipped field if non-nil, zero value otherwise.

### GetSkippedOk

`func (o *IndexResult) GetSkippedOk() (*int32, bool)`

GetSkippedOk returns a tuple with the Skipped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipped

`func (o *IndexResult) SetSkipped(v int32)`

SetSkipped sets Skipped field to given value.

### HasSkipped

`func (o *IndexResult) HasSkipped() bool`

HasSkipped returns a boolean if a field has been set.

### GetSymbols

`func (o *IndexResult) GetSymbols() int32`

GetSymbols returns the Symbols field if non-nil, zero value otherwise.

### GetSymbolsOk

`func (o *IndexResult) GetSymbolsOk() (*int32, bool)`

GetSymbolsOk returns a tuple with the Symbols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbols

`func (o *IndexResult) SetSymbols(v int32)`

SetSymbols sets Symbols field to given value.

### HasSymbols

`func (o *IndexResult) HasSymbols() bool`

HasSymbols returns a boolean if a field has been set.

### GetVectors

`func (o *IndexResult) GetVectors() int32`

GetVectors returns the Vectors field if non-nil, zero value otherwise.

### GetVectorsOk

`func (o *IndexResult) GetVectorsOk() (*int32, bool)`

GetVectorsOk returns a tuple with the Vectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVectors

`func (o *IndexResult) SetVectors(v int32)`

SetVectors sets Vectors field to given value.

### HasVectors

`func (o *IndexResult) HasVectors() bool`

HasVectors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


