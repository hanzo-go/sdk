# KnowledgeFileSearchOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Degraded** | Pointer to **bool** | Degraded is true when a leg failed — the query could not be embedded, or the full-text index could not be read — and Passages are what the other leg found. | [optional] 
**Passages** | Pointer to [**[]KnowledgeCitedPassage**](KnowledgeCitedPassage.md) | Passages are the matching passages, best first, at most Limit. | [optional] 

## Methods

### NewKnowledgeFileSearchOut

`func NewKnowledgeFileSearchOut() *KnowledgeFileSearchOut`

NewKnowledgeFileSearchOut instantiates a new KnowledgeFileSearchOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKnowledgeFileSearchOutWithDefaults

`func NewKnowledgeFileSearchOutWithDefaults() *KnowledgeFileSearchOut`

NewKnowledgeFileSearchOutWithDefaults instantiates a new KnowledgeFileSearchOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDegraded

`func (o *KnowledgeFileSearchOut) GetDegraded() bool`

GetDegraded returns the Degraded field if non-nil, zero value otherwise.

### GetDegradedOk

`func (o *KnowledgeFileSearchOut) GetDegradedOk() (*bool, bool)`

GetDegradedOk returns a tuple with the Degraded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDegraded

`func (o *KnowledgeFileSearchOut) SetDegraded(v bool)`

SetDegraded sets Degraded field to given value.

### HasDegraded

`func (o *KnowledgeFileSearchOut) HasDegraded() bool`

HasDegraded returns a boolean if a field has been set.

### GetPassages

`func (o *KnowledgeFileSearchOut) GetPassages() []KnowledgeCitedPassage`

GetPassages returns the Passages field if non-nil, zero value otherwise.

### GetPassagesOk

`func (o *KnowledgeFileSearchOut) GetPassagesOk() (*[]KnowledgeCitedPassage, bool)`

GetPassagesOk returns a tuple with the Passages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassages

`func (o *KnowledgeFileSearchOut) SetPassages(v []KnowledgeCitedPassage)`

SetPassages sets Passages field to given value.

### HasPassages

`func (o *KnowledgeFileSearchOut) HasPassages() bool`

HasPassages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


