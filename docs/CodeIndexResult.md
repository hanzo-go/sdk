# CodeIndexResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Chunks** | Pointer to **int64** | Chunks is how many AST-boundary chunks the repo holds after this pass. | [optional] 
**Files** | Pointer to **int64** | Files is how many files the repo holds after this pass. | [optional] 
**Indexed** | Pointer to **int64** | Indexed is how many files were parsed and written on this pass. | [optional] 
**Pruned** | Pointer to **int64** | Pruned is how many stored files were deleted because prune was set and they were absent from the request. | [optional] 
**Repo** | Pointer to **string** | Repo is the repository that was indexed. | [optional] 
**Semantic** | Pointer to **bool** | Semantic reports whether the semantic tier was available for this pass. When false the index is lexical + symbolic only and hybrid search still works. | [optional] 
**Skipped** | Pointer to **int64** | Skipped is how many files were unchanged by content hash and left alone. | [optional] 
**Symbols** | Pointer to **int64** | Symbols is how many symbol definitions the repo holds after this pass. | [optional] 
**Vectors** | Pointer to **int64** | Vectors is how many of those chunks carry an embedding. | [optional] 

## Methods

### NewCodeIndexResult

`func NewCodeIndexResult() *CodeIndexResult`

NewCodeIndexResult instantiates a new CodeIndexResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCodeIndexResultWithDefaults

`func NewCodeIndexResultWithDefaults() *CodeIndexResult`

NewCodeIndexResultWithDefaults instantiates a new CodeIndexResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChunks

`func (o *CodeIndexResult) GetChunks() int64`

GetChunks returns the Chunks field if non-nil, zero value otherwise.

### GetChunksOk

`func (o *CodeIndexResult) GetChunksOk() (*int64, bool)`

GetChunksOk returns a tuple with the Chunks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChunks

`func (o *CodeIndexResult) SetChunks(v int64)`

SetChunks sets Chunks field to given value.

### HasChunks

`func (o *CodeIndexResult) HasChunks() bool`

HasChunks returns a boolean if a field has been set.

### GetFiles

`func (o *CodeIndexResult) GetFiles() int64`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *CodeIndexResult) GetFilesOk() (*int64, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *CodeIndexResult) SetFiles(v int64)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *CodeIndexResult) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetIndexed

`func (o *CodeIndexResult) GetIndexed() int64`

GetIndexed returns the Indexed field if non-nil, zero value otherwise.

### GetIndexedOk

`func (o *CodeIndexResult) GetIndexedOk() (*int64, bool)`

GetIndexedOk returns a tuple with the Indexed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexed

`func (o *CodeIndexResult) SetIndexed(v int64)`

SetIndexed sets Indexed field to given value.

### HasIndexed

`func (o *CodeIndexResult) HasIndexed() bool`

HasIndexed returns a boolean if a field has been set.

### GetPruned

`func (o *CodeIndexResult) GetPruned() int64`

GetPruned returns the Pruned field if non-nil, zero value otherwise.

### GetPrunedOk

`func (o *CodeIndexResult) GetPrunedOk() (*int64, bool)`

GetPrunedOk returns a tuple with the Pruned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPruned

`func (o *CodeIndexResult) SetPruned(v int64)`

SetPruned sets Pruned field to given value.

### HasPruned

`func (o *CodeIndexResult) HasPruned() bool`

HasPruned returns a boolean if a field has been set.

### GetRepo

`func (o *CodeIndexResult) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *CodeIndexResult) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *CodeIndexResult) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *CodeIndexResult) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### GetSemantic

`func (o *CodeIndexResult) GetSemantic() bool`

GetSemantic returns the Semantic field if non-nil, zero value otherwise.

### GetSemanticOk

`func (o *CodeIndexResult) GetSemanticOk() (*bool, bool)`

GetSemanticOk returns a tuple with the Semantic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSemantic

`func (o *CodeIndexResult) SetSemantic(v bool)`

SetSemantic sets Semantic field to given value.

### HasSemantic

`func (o *CodeIndexResult) HasSemantic() bool`

HasSemantic returns a boolean if a field has been set.

### GetSkipped

`func (o *CodeIndexResult) GetSkipped() int64`

GetSkipped returns the Skipped field if non-nil, zero value otherwise.

### GetSkippedOk

`func (o *CodeIndexResult) GetSkippedOk() (*int64, bool)`

GetSkippedOk returns a tuple with the Skipped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipped

`func (o *CodeIndexResult) SetSkipped(v int64)`

SetSkipped sets Skipped field to given value.

### HasSkipped

`func (o *CodeIndexResult) HasSkipped() bool`

HasSkipped returns a boolean if a field has been set.

### GetSymbols

`func (o *CodeIndexResult) GetSymbols() int64`

GetSymbols returns the Symbols field if non-nil, zero value otherwise.

### GetSymbolsOk

`func (o *CodeIndexResult) GetSymbolsOk() (*int64, bool)`

GetSymbolsOk returns a tuple with the Symbols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbols

`func (o *CodeIndexResult) SetSymbols(v int64)`

SetSymbols sets Symbols field to given value.

### HasSymbols

`func (o *CodeIndexResult) HasSymbols() bool`

HasSymbols returns a boolean if a field has been set.

### GetVectors

`func (o *CodeIndexResult) GetVectors() int64`

GetVectors returns the Vectors field if non-nil, zero value otherwise.

### GetVectorsOk

`func (o *CodeIndexResult) GetVectorsOk() (*int64, bool)`

GetVectorsOk returns a tuple with the Vectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVectors

`func (o *CodeIndexResult) SetVectors(v int64)`

SetVectors sets Vectors field to given value.

### HasVectors

`func (o *CodeIndexResult) HasVectors() bool`

HasVectors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


