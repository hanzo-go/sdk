# RepoTree

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Files** | Pointer to [**[]TreeEntry**](TreeEntry.md) | Files are the repo&#39;s indexed files in path order, each with its language and how many symbols it defines. Never null. | [optional] 
**Repo** | Pointer to **string** | Repo echoes the repository that was walked. | [optional] 

## Methods

### NewRepoTree

`func NewRepoTree() *RepoTree`

NewRepoTree instantiates a new RepoTree object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepoTreeWithDefaults

`func NewRepoTreeWithDefaults() *RepoTree`

NewRepoTreeWithDefaults instantiates a new RepoTree object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFiles

`func (o *RepoTree) GetFiles() []TreeEntry`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *RepoTree) GetFilesOk() (*[]TreeEntry, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *RepoTree) SetFiles(v []TreeEntry)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *RepoTree) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetRepo

`func (o *RepoTree) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *RepoTree) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *RepoTree) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *RepoTree) HasRepo() bool`

HasRepo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


