# CloudFileContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to **string** | Content is the file&#39;s text as the index stored it. It is NOT guaranteed byte-verbatim — the git object plane is the source of record for exact bytes, history and blame. | [optional] 
**Lang** | Pointer to **string** | Lang is the detected language. | [optional] 
**Path** | Pointer to **string** | Path echoes the file that was read. | [optional] 
**Repo** | Pointer to **string** | Repo echoes the repository it came from. | [optional] 

## Methods

### NewCloudFileContent

`func NewCloudFileContent() *CloudFileContent`

NewCloudFileContent instantiates a new CloudFileContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudFileContentWithDefaults

`func NewCloudFileContentWithDefaults() *CloudFileContent`

NewCloudFileContentWithDefaults instantiates a new CloudFileContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *CloudFileContent) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *CloudFileContent) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *CloudFileContent) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *CloudFileContent) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetLang

`func (o *CloudFileContent) GetLang() string`

GetLang returns the Lang field if non-nil, zero value otherwise.

### GetLangOk

`func (o *CloudFileContent) GetLangOk() (*string, bool)`

GetLangOk returns a tuple with the Lang field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLang

`func (o *CloudFileContent) SetLang(v string)`

SetLang sets Lang field to given value.

### HasLang

`func (o *CloudFileContent) HasLang() bool`

HasLang returns a boolean if a field has been set.

### GetPath

`func (o *CloudFileContent) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *CloudFileContent) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *CloudFileContent) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *CloudFileContent) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetRepo

`func (o *CloudFileContent) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *CloudFileContent) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *CloudFileContent) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *CloudFileContent) HasRepo() bool`

HasRepo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


