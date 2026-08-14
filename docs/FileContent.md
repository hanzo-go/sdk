# FileContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to **string** | Content is the file&#39;s text as the index stored it. It is NOT guaranteed byte-verbatim — the git object plane is the source of record for exact bytes, history and blame. | [optional] 
**Lang** | Pointer to **string** | Lang is the detected language. | [optional] 
**Path** | Pointer to **string** | Path echoes the file that was read. | [optional] 
**Repo** | Pointer to **string** | Repo echoes the repository it came from. | [optional] 

## Methods

### NewFileContent

`func NewFileContent() *FileContent`

NewFileContent instantiates a new FileContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileContentWithDefaults

`func NewFileContentWithDefaults() *FileContent`

NewFileContentWithDefaults instantiates a new FileContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *FileContent) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *FileContent) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *FileContent) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *FileContent) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetLang

`func (o *FileContent) GetLang() string`

GetLang returns the Lang field if non-nil, zero value otherwise.

### GetLangOk

`func (o *FileContent) GetLangOk() (*string, bool)`

GetLangOk returns a tuple with the Lang field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLang

`func (o *FileContent) SetLang(v string)`

SetLang sets Lang field to given value.

### HasLang

`func (o *FileContent) HasLang() bool`

HasLang returns a boolean if a field has been set.

### GetPath

`func (o *FileContent) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *FileContent) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *FileContent) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *FileContent) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetRepo

`func (o *FileContent) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *FileContent) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *FileContent) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *FileContent) HasRepo() bool`

HasRepo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


