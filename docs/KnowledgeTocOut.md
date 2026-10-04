# KnowledgeTocOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**File** | Pointer to [**KnowledgeFile**](KnowledgeFile.md) | File is the file it is of. | [optional] 
**Sections** | Pointer to [**[]KnowledgeTocEntry**](KnowledgeTocEntry.md) | Sections are its nodes in document order, the document first; Parent makes them a tree. | [optional] 

## Methods

### NewKnowledgeTocOut

`func NewKnowledgeTocOut() *KnowledgeTocOut`

NewKnowledgeTocOut instantiates a new KnowledgeTocOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKnowledgeTocOutWithDefaults

`func NewKnowledgeTocOutWithDefaults() *KnowledgeTocOut`

NewKnowledgeTocOutWithDefaults instantiates a new KnowledgeTocOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFile

`func (o *KnowledgeTocOut) GetFile() KnowledgeFile`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *KnowledgeTocOut) GetFileOk() (*KnowledgeFile, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *KnowledgeTocOut) SetFile(v KnowledgeFile)`

SetFile sets File field to given value.

### HasFile

`func (o *KnowledgeTocOut) HasFile() bool`

HasFile returns a boolean if a field has been set.

### GetSections

`func (o *KnowledgeTocOut) GetSections() []KnowledgeTocEntry`

GetSections returns the Sections field if non-nil, zero value otherwise.

### GetSectionsOk

`func (o *KnowledgeTocOut) GetSectionsOk() (*[]KnowledgeTocEntry, bool)`

GetSectionsOk returns a tuple with the Sections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSections

`func (o *KnowledgeTocOut) SetSections(v []KnowledgeTocEntry)`

SetSections sets Sections field to given value.

### HasSections

`func (o *KnowledgeTocOut) HasSections() bool`

HasSections returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


