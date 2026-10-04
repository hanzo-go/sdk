# KnowledgeSectionLink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**File** | Pointer to [**KnowledgeFileBrief**](KnowledgeFileBrief.md) | File is the file the linked section is in. | [optional] 
**Kind** | Pointer to **string** | Kind is href, xref, cites, or entity (the two sections name the same entities). | [optional] 
**Label** | Pointer to **string** | Label is the link&#39;s text, or the entities the two share. | [optional] 
**Section** | Pointer to [**KnowledgeSectionBrief**](KnowledgeSectionBrief.md) | Section is the linked section. | [optional] 

## Methods

### NewKnowledgeSectionLink

`func NewKnowledgeSectionLink() *KnowledgeSectionLink`

NewKnowledgeSectionLink instantiates a new KnowledgeSectionLink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKnowledgeSectionLinkWithDefaults

`func NewKnowledgeSectionLinkWithDefaults() *KnowledgeSectionLink`

NewKnowledgeSectionLinkWithDefaults instantiates a new KnowledgeSectionLink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFile

`func (o *KnowledgeSectionLink) GetFile() KnowledgeFileBrief`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *KnowledgeSectionLink) GetFileOk() (*KnowledgeFileBrief, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *KnowledgeSectionLink) SetFile(v KnowledgeFileBrief)`

SetFile sets File field to given value.

### HasFile

`func (o *KnowledgeSectionLink) HasFile() bool`

HasFile returns a boolean if a field has been set.

### GetKind

`func (o *KnowledgeSectionLink) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *KnowledgeSectionLink) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *KnowledgeSectionLink) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *KnowledgeSectionLink) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLabel

`func (o *KnowledgeSectionLink) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *KnowledgeSectionLink) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *KnowledgeSectionLink) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *KnowledgeSectionLink) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetSection

`func (o *KnowledgeSectionLink) GetSection() KnowledgeSectionBrief`

GetSection returns the Section field if non-nil, zero value otherwise.

### GetSectionOk

`func (o *KnowledgeSectionLink) GetSectionOk() (*KnowledgeSectionBrief, bool)`

GetSectionOk returns a tuple with the Section field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSection

`func (o *KnowledgeSectionLink) SetSection(v KnowledgeSectionBrief)`

SetSection sets Section field to given value.

### HasSection

`func (o *KnowledgeSectionLink) HasSection() bool`

HasSection returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


