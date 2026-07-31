# CloudFilesJSON

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Files** | Pointer to [**[]CloudFileJSON**](CloudFileJSON.md) | Files are the selected files, sorted by path. Directories are never returned. | [optional] 
**Rev** | Pointer to **string** | Rev is the full revision the ref resolved to — pin follow-up reads to it. | [optional] 

## Methods

### NewCloudFilesJSON

`func NewCloudFilesJSON() *CloudFilesJSON`

NewCloudFilesJSON instantiates a new CloudFilesJSON object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudFilesJSONWithDefaults

`func NewCloudFilesJSONWithDefaults() *CloudFilesJSON`

NewCloudFilesJSONWithDefaults instantiates a new CloudFilesJSON object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFiles

`func (o *CloudFilesJSON) GetFiles() []CloudFileJSON`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *CloudFilesJSON) GetFilesOk() (*[]CloudFileJSON, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *CloudFilesJSON) SetFiles(v []CloudFileJSON)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *CloudFilesJSON) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetRev

`func (o *CloudFilesJSON) GetRev() string`

GetRev returns the Rev field if non-nil, zero value otherwise.

### GetRevOk

`func (o *CloudFilesJSON) GetRevOk() (*string, bool)`

GetRevOk returns a tuple with the Rev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRev

`func (o *CloudFilesJSON) SetRev(v string)`

SetRev sets Rev field to given value.

### HasRev

`func (o *CloudFilesJSON) HasRev() bool`

HasRev returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


