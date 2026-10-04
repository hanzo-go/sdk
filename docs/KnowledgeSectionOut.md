# KnowledgeSectionOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Children** | Pointer to [**[]KnowledgeTocEntry**](KnowledgeTocEntry.md) | Children are its subsections, to open next. | [optional] 
**Entities** | Pointer to [**[]KnowledgeFileEntity**](KnowledgeFileEntity.md) | Entities are the entities it names. | [optional] 
**File** | Pointer to [**KnowledgeFileBrief**](KnowledgeFileBrief.md) | File is the file it is in. | [optional] 
**Links** | Pointer to [**[]KnowledgeSectionLink**](KnowledgeSectionLink.md) | Links are the sections the graph reaches from it, strongest first. | [optional] 
**Section** | Pointer to [**KnowledgeSectionBrief**](KnowledgeSectionBrief.md) | Section is where it is. | [optional] 
**Summary** | Pointer to **string** | Summary is its one-line account. | [optional] 
**Text** | Pointer to **string** | Text is its own text — up to its first subsection — at most 64 KB. | [optional] 

## Methods

### NewKnowledgeSectionOut

`func NewKnowledgeSectionOut() *KnowledgeSectionOut`

NewKnowledgeSectionOut instantiates a new KnowledgeSectionOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKnowledgeSectionOutWithDefaults

`func NewKnowledgeSectionOutWithDefaults() *KnowledgeSectionOut`

NewKnowledgeSectionOutWithDefaults instantiates a new KnowledgeSectionOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChildren

`func (o *KnowledgeSectionOut) GetChildren() []KnowledgeTocEntry`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *KnowledgeSectionOut) GetChildrenOk() (*[]KnowledgeTocEntry, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *KnowledgeSectionOut) SetChildren(v []KnowledgeTocEntry)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *KnowledgeSectionOut) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### GetEntities

`func (o *KnowledgeSectionOut) GetEntities() []KnowledgeFileEntity`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *KnowledgeSectionOut) GetEntitiesOk() (*[]KnowledgeFileEntity, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *KnowledgeSectionOut) SetEntities(v []KnowledgeFileEntity)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *KnowledgeSectionOut) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### GetFile

`func (o *KnowledgeSectionOut) GetFile() KnowledgeFileBrief`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *KnowledgeSectionOut) GetFileOk() (*KnowledgeFileBrief, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *KnowledgeSectionOut) SetFile(v KnowledgeFileBrief)`

SetFile sets File field to given value.

### HasFile

`func (o *KnowledgeSectionOut) HasFile() bool`

HasFile returns a boolean if a field has been set.

### GetLinks

`func (o *KnowledgeSectionOut) GetLinks() []KnowledgeSectionLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *KnowledgeSectionOut) GetLinksOk() (*[]KnowledgeSectionLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *KnowledgeSectionOut) SetLinks(v []KnowledgeSectionLink)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *KnowledgeSectionOut) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetSection

`func (o *KnowledgeSectionOut) GetSection() KnowledgeSectionBrief`

GetSection returns the Section field if non-nil, zero value otherwise.

### GetSectionOk

`func (o *KnowledgeSectionOut) GetSectionOk() (*KnowledgeSectionBrief, bool)`

GetSectionOk returns a tuple with the Section field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSection

`func (o *KnowledgeSectionOut) SetSection(v KnowledgeSectionBrief)`

SetSection sets Section field to given value.

### HasSection

`func (o *KnowledgeSectionOut) HasSection() bool`

HasSection returns a boolean if a field has been set.

### GetSummary

`func (o *KnowledgeSectionOut) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *KnowledgeSectionOut) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *KnowledgeSectionOut) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *KnowledgeSectionOut) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetText

`func (o *KnowledgeSectionOut) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *KnowledgeSectionOut) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *KnowledgeSectionOut) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *KnowledgeSectionOut) HasText() bool`

HasText returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


