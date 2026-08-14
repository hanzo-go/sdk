# FilesJSON

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Files** | Pointer to [**[]FileJSON**](FileJSON.md) | Files are the selected files, sorted by path. Directories are never returned. | [optional] 
**Rev** | Pointer to **string** | Rev is the full revision the ref resolved to — pin follow-up reads to it. | [optional] 

## Methods

### NewFilesJSON

`func NewFilesJSON() *FilesJSON`

NewFilesJSON instantiates a new FilesJSON object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilesJSONWithDefaults

`func NewFilesJSONWithDefaults() *FilesJSON`

NewFilesJSONWithDefaults instantiates a new FilesJSON object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFiles

`func (o *FilesJSON) GetFiles() []FileJSON`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *FilesJSON) GetFilesOk() (*[]FileJSON, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *FilesJSON) SetFiles(v []FileJSON)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *FilesJSON) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetRev

`func (o *FilesJSON) GetRev() string`

GetRev returns the Rev field if non-nil, zero value otherwise.

### GetRevOk

`func (o *FilesJSON) GetRevOk() (*string, bool)`

GetRevOk returns a tuple with the Rev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRev

`func (o *FilesJSON) SetRev(v string)`

SetRev sets Rev field to given value.

### HasRev

`func (o *FilesJSON) HasRev() bool`

HasRev returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


