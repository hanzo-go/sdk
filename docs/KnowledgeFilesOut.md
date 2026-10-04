# KnowledgeFilesOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Files** | Pointer to [**[]KnowledgeFile**](KnowledgeFile.md) | Files are the org&#39;s files, most recently changed first. | [optional] 

## Methods

### NewKnowledgeFilesOut

`func NewKnowledgeFilesOut() *KnowledgeFilesOut`

NewKnowledgeFilesOut instantiates a new KnowledgeFilesOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKnowledgeFilesOutWithDefaults

`func NewKnowledgeFilesOutWithDefaults() *KnowledgeFilesOut`

NewKnowledgeFilesOutWithDefaults instantiates a new KnowledgeFilesOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFiles

`func (o *KnowledgeFilesOut) GetFiles() []KnowledgeFile`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *KnowledgeFilesOut) GetFilesOk() (*[]KnowledgeFile, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *KnowledgeFilesOut) SetFiles(v []KnowledgeFile)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *KnowledgeFilesOut) HasFiles() bool`

HasFiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


