# CloudIndexIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Files** | Pointer to [**[]CloudFileInput**](CloudFileInput.md) | Files is the full set of files to index. Required and non-empty; max 20000 files, 1 MiB per file and 1 GiB in total. Unchanged files are skipped by content hash, so re-sending the whole tree is cheap. | [optional] 
**Prune** | Pointer to **bool** | Prune deletes indexed files that are NOT in this request — which makes the call a full sync of the repo rather than an upsert. Only pass it when Files is the complete tree. | [optional] 
**Repo** | Pointer to **string** | Repo is the repository label to index under. Required, max 200 bytes. It is a stored column value, not a filesystem path. | [optional] 

## Methods

### NewCloudIndexIn

`func NewCloudIndexIn() *CloudIndexIn`

NewCloudIndexIn instantiates a new CloudIndexIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudIndexInWithDefaults

`func NewCloudIndexInWithDefaults() *CloudIndexIn`

NewCloudIndexInWithDefaults instantiates a new CloudIndexIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFiles

`func (o *CloudIndexIn) GetFiles() []CloudFileInput`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *CloudIndexIn) GetFilesOk() (*[]CloudFileInput, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *CloudIndexIn) SetFiles(v []CloudFileInput)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *CloudIndexIn) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetPrune

`func (o *CloudIndexIn) GetPrune() bool`

GetPrune returns the Prune field if non-nil, zero value otherwise.

### GetPruneOk

`func (o *CloudIndexIn) GetPruneOk() (*bool, bool)`

GetPruneOk returns a tuple with the Prune field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrune

`func (o *CloudIndexIn) SetPrune(v bool)`

SetPrune sets Prune field to given value.

### HasPrune

`func (o *CloudIndexIn) HasPrune() bool`

HasPrune returns a boolean if a field has been set.

### GetRepo

`func (o *CloudIndexIn) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *CloudIndexIn) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *CloudIndexIn) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *CloudIndexIn) HasRepo() bool`

HasRepo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


