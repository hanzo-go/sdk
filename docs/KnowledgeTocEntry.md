# KnowledgeTocEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | ID is the section&#39;s number; 0 is the document itself. | [optional] 
**Level** | Pointer to **int64** | Level is its depth: 0 for the document, 1 for a top-level heading. | [optional] 
**Parent** | Pointer to **int64** | Parent is the id of the section it sits in; -1 for the document. | [optional] 
**Size** | Pointer to **int64** | Size is how many bytes of the document it spans, subsections included. | [optional] 
**Summary** | Pointer to **string** | Summary is a one-line account of what it covers. | [optional] 
**Synthetic** | Pointer to **bool** | Synthetic is true for a section the index cut out of a long run of text the document gave no structure to. | [optional] 
**Title** | Pointer to **string** | Title is its heading. | [optional] 

## Methods

### NewKnowledgeTocEntry

`func NewKnowledgeTocEntry() *KnowledgeTocEntry`

NewKnowledgeTocEntry instantiates a new KnowledgeTocEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKnowledgeTocEntryWithDefaults

`func NewKnowledgeTocEntryWithDefaults() *KnowledgeTocEntry`

NewKnowledgeTocEntryWithDefaults instantiates a new KnowledgeTocEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KnowledgeTocEntry) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KnowledgeTocEntry) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KnowledgeTocEntry) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *KnowledgeTocEntry) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLevel

`func (o *KnowledgeTocEntry) GetLevel() int64`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *KnowledgeTocEntry) GetLevelOk() (*int64, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *KnowledgeTocEntry) SetLevel(v int64)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *KnowledgeTocEntry) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetParent

`func (o *KnowledgeTocEntry) GetParent() int64`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *KnowledgeTocEntry) GetParentOk() (*int64, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *KnowledgeTocEntry) SetParent(v int64)`

SetParent sets Parent field to given value.

### HasParent

`func (o *KnowledgeTocEntry) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetSize

`func (o *KnowledgeTocEntry) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *KnowledgeTocEntry) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *KnowledgeTocEntry) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *KnowledgeTocEntry) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSummary

`func (o *KnowledgeTocEntry) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *KnowledgeTocEntry) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *KnowledgeTocEntry) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *KnowledgeTocEntry) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetSynthetic

`func (o *KnowledgeTocEntry) GetSynthetic() bool`

GetSynthetic returns the Synthetic field if non-nil, zero value otherwise.

### GetSyntheticOk

`func (o *KnowledgeTocEntry) GetSyntheticOk() (*bool, bool)`

GetSyntheticOk returns a tuple with the Synthetic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynthetic

`func (o *KnowledgeTocEntry) SetSynthetic(v bool)`

SetSynthetic sets Synthetic field to given value.

### HasSynthetic

`func (o *KnowledgeTocEntry) HasSynthetic() bool`

HasSynthetic returns a boolean if a field has been set.

### GetTitle

`func (o *KnowledgeTocEntry) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *KnowledgeTocEntry) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *KnowledgeTocEntry) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *KnowledgeTocEntry) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


