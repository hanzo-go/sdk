# KnowledgeCitedPassage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cite** | Pointer to **string** | Cite is how an answer cites it: file › section path › ¶part. | [optional] 
**File** | Pointer to [**KnowledgeFileBrief**](KnowledgeFileBrief.md) | File is the file the passage is from. | [optional] 
**Part** | Pointer to **int64** | Part is the passage&#39;s place in its section, from 1. | [optional] 
**Score** | Pointer to **float64** | Score is the fused rank score, comparable within one response only. | [optional] 
**Section** | Pointer to [**KnowledgeSectionBrief**](KnowledgeSectionBrief.md) | Section is the section it is in. | [optional] 
**Text** | Pointer to **string** | Text is the passage: a span of its section of about 2000 bytes, never crossing into another section. | [optional] 
**Via** | Pointer to **string** | Via is how retrieval reached it: search (it matched), toc (it is in a section the table of contents pointed at) or graph (a linked section). | [optional] 
**Why** | Pointer to **string** | Why says, for a passage reached through the graph, which edge led there. | [optional] 

## Methods

### NewKnowledgeCitedPassage

`func NewKnowledgeCitedPassage() *KnowledgeCitedPassage`

NewKnowledgeCitedPassage instantiates a new KnowledgeCitedPassage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKnowledgeCitedPassageWithDefaults

`func NewKnowledgeCitedPassageWithDefaults() *KnowledgeCitedPassage`

NewKnowledgeCitedPassageWithDefaults instantiates a new KnowledgeCitedPassage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCite

`func (o *KnowledgeCitedPassage) GetCite() string`

GetCite returns the Cite field if non-nil, zero value otherwise.

### GetCiteOk

`func (o *KnowledgeCitedPassage) GetCiteOk() (*string, bool)`

GetCiteOk returns a tuple with the Cite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCite

`func (o *KnowledgeCitedPassage) SetCite(v string)`

SetCite sets Cite field to given value.

### HasCite

`func (o *KnowledgeCitedPassage) HasCite() bool`

HasCite returns a boolean if a field has been set.

### GetFile

`func (o *KnowledgeCitedPassage) GetFile() KnowledgeFileBrief`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *KnowledgeCitedPassage) GetFileOk() (*KnowledgeFileBrief, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *KnowledgeCitedPassage) SetFile(v KnowledgeFileBrief)`

SetFile sets File field to given value.

### HasFile

`func (o *KnowledgeCitedPassage) HasFile() bool`

HasFile returns a boolean if a field has been set.

### GetPart

`func (o *KnowledgeCitedPassage) GetPart() int64`

GetPart returns the Part field if non-nil, zero value otherwise.

### GetPartOk

`func (o *KnowledgeCitedPassage) GetPartOk() (*int64, bool)`

GetPartOk returns a tuple with the Part field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPart

`func (o *KnowledgeCitedPassage) SetPart(v int64)`

SetPart sets Part field to given value.

### HasPart

`func (o *KnowledgeCitedPassage) HasPart() bool`

HasPart returns a boolean if a field has been set.

### GetScore

`func (o *KnowledgeCitedPassage) GetScore() float64`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *KnowledgeCitedPassage) GetScoreOk() (*float64, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *KnowledgeCitedPassage) SetScore(v float64)`

SetScore sets Score field to given value.

### HasScore

`func (o *KnowledgeCitedPassage) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetSection

`func (o *KnowledgeCitedPassage) GetSection() KnowledgeSectionBrief`

GetSection returns the Section field if non-nil, zero value otherwise.

### GetSectionOk

`func (o *KnowledgeCitedPassage) GetSectionOk() (*KnowledgeSectionBrief, bool)`

GetSectionOk returns a tuple with the Section field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSection

`func (o *KnowledgeCitedPassage) SetSection(v KnowledgeSectionBrief)`

SetSection sets Section field to given value.

### HasSection

`func (o *KnowledgeCitedPassage) HasSection() bool`

HasSection returns a boolean if a field has been set.

### GetText

`func (o *KnowledgeCitedPassage) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *KnowledgeCitedPassage) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *KnowledgeCitedPassage) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *KnowledgeCitedPassage) HasText() bool`

HasText returns a boolean if a field has been set.

### GetVia

`func (o *KnowledgeCitedPassage) GetVia() string`

GetVia returns the Via field if non-nil, zero value otherwise.

### GetViaOk

`func (o *KnowledgeCitedPassage) GetViaOk() (*string, bool)`

GetViaOk returns a tuple with the Via field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVia

`func (o *KnowledgeCitedPassage) SetVia(v string)`

SetVia sets Via field to given value.

### HasVia

`func (o *KnowledgeCitedPassage) HasVia() bool`

HasVia returns a boolean if a field has been set.

### GetWhy

`func (o *KnowledgeCitedPassage) GetWhy() string`

GetWhy returns the Why field if non-nil, zero value otherwise.

### GetWhyOk

`func (o *KnowledgeCitedPassage) GetWhyOk() (*string, bool)`

GetWhyOk returns a tuple with the Why field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhy

`func (o *KnowledgeCitedPassage) SetWhy(v string)`

SetWhy sets Why field to given value.

### HasWhy

`func (o *KnowledgeCitedPassage) HasWhy() bool`

HasWhy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


