# KnowledgeSectionBrief

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | ID is the section&#39;s number in its file&#39;s table of contents; 0 is the document itself. | [optional] 
**Path** | Pointer to **string** | Path is the section&#39;s place in its document: the titles from the document&#39;s own down to it, joined by \&quot; › \&quot;. | [optional] 
**Title** | Pointer to **string** | Title is the section&#39;s heading. | [optional] 

## Methods

### NewKnowledgeSectionBrief

`func NewKnowledgeSectionBrief() *KnowledgeSectionBrief`

NewKnowledgeSectionBrief instantiates a new KnowledgeSectionBrief object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKnowledgeSectionBriefWithDefaults

`func NewKnowledgeSectionBriefWithDefaults() *KnowledgeSectionBrief`

NewKnowledgeSectionBriefWithDefaults instantiates a new KnowledgeSectionBrief object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KnowledgeSectionBrief) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KnowledgeSectionBrief) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KnowledgeSectionBrief) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *KnowledgeSectionBrief) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPath

`func (o *KnowledgeSectionBrief) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *KnowledgeSectionBrief) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *KnowledgeSectionBrief) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *KnowledgeSectionBrief) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetTitle

`func (o *KnowledgeSectionBrief) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *KnowledgeSectionBrief) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *KnowledgeSectionBrief) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *KnowledgeSectionBrief) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


