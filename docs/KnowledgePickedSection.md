# KnowledgePickedSection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**File** | Pointer to [**KnowledgeFileBrief**](KnowledgeFileBrief.md) | File is the file it is in. | [optional] 
**Section** | Pointer to [**KnowledgeSectionBrief**](KnowledgeSectionBrief.md) | Section is where it is. | [optional] 
**Summary** | Pointer to **string** | Summary is its one-line account. | [optional] 

## Methods

### NewKnowledgePickedSection

`func NewKnowledgePickedSection() *KnowledgePickedSection`

NewKnowledgePickedSection instantiates a new KnowledgePickedSection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKnowledgePickedSectionWithDefaults

`func NewKnowledgePickedSectionWithDefaults() *KnowledgePickedSection`

NewKnowledgePickedSectionWithDefaults instantiates a new KnowledgePickedSection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFile

`func (o *KnowledgePickedSection) GetFile() KnowledgeFileBrief`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *KnowledgePickedSection) GetFileOk() (*KnowledgeFileBrief, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *KnowledgePickedSection) SetFile(v KnowledgeFileBrief)`

SetFile sets File field to given value.

### HasFile

`func (o *KnowledgePickedSection) HasFile() bool`

HasFile returns a boolean if a field has been set.

### GetSection

`func (o *KnowledgePickedSection) GetSection() KnowledgeSectionBrief`

GetSection returns the Section field if non-nil, zero value otherwise.

### GetSectionOk

`func (o *KnowledgePickedSection) GetSectionOk() (*KnowledgeSectionBrief, bool)`

GetSectionOk returns a tuple with the Section field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSection

`func (o *KnowledgePickedSection) SetSection(v KnowledgeSectionBrief)`

SetSection sets Section field to given value.

### HasSection

`func (o *KnowledgePickedSection) HasSection() bool`

HasSection returns a boolean if a field has been set.

### GetSummary

`func (o *KnowledgePickedSection) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *KnowledgePickedSection) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *KnowledgePickedSection) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *KnowledgePickedSection) HasSummary() bool`

HasSummary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


