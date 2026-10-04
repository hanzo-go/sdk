# KnowledgeRetrieveOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Degraded** | Pointer to **bool** | Degraded is true when a search leg failed along the way. | [optional] 
**Passages** | Pointer to [**[]KnowledgeCitedPassage**](KnowledgeCitedPassage.md) | Passages are the passages drilled out of those sections, then those the graph reached, each citing file › section › paragraph. | [optional] 
**Picked** | Pointer to **string** | Picked says how the sections were chosen: whole when the documents were short enough to read every section, model when a model read the tables of contents, search when the sections of the best passages stood in — the documents had no headings to choose by, or no model answered. | [optional] 
**Sections** | Pointer to [**[]KnowledgePickedSection**](KnowledgePickedSection.md) | Sections are the sections chosen from the candidate documents&#39; tables of contents. | [optional] 

## Methods

### NewKnowledgeRetrieveOut

`func NewKnowledgeRetrieveOut() *KnowledgeRetrieveOut`

NewKnowledgeRetrieveOut instantiates a new KnowledgeRetrieveOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKnowledgeRetrieveOutWithDefaults

`func NewKnowledgeRetrieveOutWithDefaults() *KnowledgeRetrieveOut`

NewKnowledgeRetrieveOutWithDefaults instantiates a new KnowledgeRetrieveOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDegraded

`func (o *KnowledgeRetrieveOut) GetDegraded() bool`

GetDegraded returns the Degraded field if non-nil, zero value otherwise.

### GetDegradedOk

`func (o *KnowledgeRetrieveOut) GetDegradedOk() (*bool, bool)`

GetDegradedOk returns a tuple with the Degraded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDegraded

`func (o *KnowledgeRetrieveOut) SetDegraded(v bool)`

SetDegraded sets Degraded field to given value.

### HasDegraded

`func (o *KnowledgeRetrieveOut) HasDegraded() bool`

HasDegraded returns a boolean if a field has been set.

### GetPassages

`func (o *KnowledgeRetrieveOut) GetPassages() []KnowledgeCitedPassage`

GetPassages returns the Passages field if non-nil, zero value otherwise.

### GetPassagesOk

`func (o *KnowledgeRetrieveOut) GetPassagesOk() (*[]KnowledgeCitedPassage, bool)`

GetPassagesOk returns a tuple with the Passages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassages

`func (o *KnowledgeRetrieveOut) SetPassages(v []KnowledgeCitedPassage)`

SetPassages sets Passages field to given value.

### HasPassages

`func (o *KnowledgeRetrieveOut) HasPassages() bool`

HasPassages returns a boolean if a field has been set.

### GetPicked

`func (o *KnowledgeRetrieveOut) GetPicked() string`

GetPicked returns the Picked field if non-nil, zero value otherwise.

### GetPickedOk

`func (o *KnowledgeRetrieveOut) GetPickedOk() (*string, bool)`

GetPickedOk returns a tuple with the Picked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPicked

`func (o *KnowledgeRetrieveOut) SetPicked(v string)`

SetPicked sets Picked field to given value.

### HasPicked

`func (o *KnowledgeRetrieveOut) HasPicked() bool`

HasPicked returns a boolean if a field has been set.

### GetSections

`func (o *KnowledgeRetrieveOut) GetSections() []KnowledgePickedSection`

GetSections returns the Sections field if non-nil, zero value otherwise.

### GetSectionsOk

`func (o *KnowledgeRetrieveOut) GetSectionsOk() (*[]KnowledgePickedSection, bool)`

GetSectionsOk returns a tuple with the Sections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSections

`func (o *KnowledgeRetrieveOut) SetSections(v []KnowledgePickedSection)`

SetSections sets Sections field to given value.

### HasSections

`func (o *KnowledgeRetrieveOut) HasSections() bool`

HasSections returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


