# CodeIndexIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Files** | Pointer to [**[]CodeFileInput**](CodeFileInput.md) | Files is the full set of files to index. Required and non-empty; max 20000 files, 1 MiB per file and 1 GiB in total. Unchanged files are skipped by content hash, so re-sending the whole tree is cheap. | [optional] 
**Prune** | Pointer to **bool** | Prune deletes indexed files that are NOT in this request — which makes the call a full sync of the repo rather than an upsert. Only pass it when Files is the complete tree. | [optional] 
**Repo** | Pointer to **string** | Repo is the repository label to index under. Required, max 200 bytes. It is a stored column value, not a filesystem path. | [optional] 

## Methods

### NewCodeIndexIn

`func NewCodeIndexIn() *CodeIndexIn`

NewCodeIndexIn instantiates a new CodeIndexIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCodeIndexInWithDefaults

`func NewCodeIndexInWithDefaults() *CodeIndexIn`

NewCodeIndexInWithDefaults instantiates a new CodeIndexIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFiles

`func (o *CodeIndexIn) GetFiles() []CodeFileInput`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *CodeIndexIn) GetFilesOk() (*[]CodeFileInput, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *CodeIndexIn) SetFiles(v []CodeFileInput)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *CodeIndexIn) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetPrune

`func (o *CodeIndexIn) GetPrune() bool`

GetPrune returns the Prune field if non-nil, zero value otherwise.

### GetPruneOk

`func (o *CodeIndexIn) GetPruneOk() (*bool, bool)`

GetPruneOk returns a tuple with the Prune field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrune

`func (o *CodeIndexIn) SetPrune(v bool)`

SetPrune sets Prune field to given value.

### HasPrune

`func (o *CodeIndexIn) HasPrune() bool`

HasPrune returns a boolean if a field has been set.

### GetRepo

`func (o *CodeIndexIn) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *CodeIndexIn) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *CodeIndexIn) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *CodeIndexIn) HasRepo() bool`

HasRepo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


