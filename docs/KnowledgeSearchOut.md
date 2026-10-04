# KnowledgeSearchOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Degraded** | Pointer to **bool** | Degraded is true when a leg FAILED — the query could not be embedded, or a store could not be read — and Hits are what the other leg found, possibly nothing. A RAG caller continues with what it got instead of failing the turn. Absent on a normal answer, including one from a deployment that runs only one of the legs. | [optional] 
**Hits** | Pointer to [**[]KnowledgeHit**](KnowledgeHit.md) | Hits are the matching documents, most relevant first: the order is the two legs fused, so a document both the semantic and the keyword leg found comes before one only a single leg found. | [optional] 

## Methods

### NewKnowledgeSearchOut

`func NewKnowledgeSearchOut() *KnowledgeSearchOut`

NewKnowledgeSearchOut instantiates a new KnowledgeSearchOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKnowledgeSearchOutWithDefaults

`func NewKnowledgeSearchOutWithDefaults() *KnowledgeSearchOut`

NewKnowledgeSearchOutWithDefaults instantiates a new KnowledgeSearchOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDegraded

`func (o *KnowledgeSearchOut) GetDegraded() bool`

GetDegraded returns the Degraded field if non-nil, zero value otherwise.

### GetDegradedOk

`func (o *KnowledgeSearchOut) GetDegradedOk() (*bool, bool)`

GetDegradedOk returns a tuple with the Degraded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDegraded

`func (o *KnowledgeSearchOut) SetDegraded(v bool)`

SetDegraded sets Degraded field to given value.

### HasDegraded

`func (o *KnowledgeSearchOut) HasDegraded() bool`

HasDegraded returns a boolean if a field has been set.

### GetHits

`func (o *KnowledgeSearchOut) GetHits() []KnowledgeHit`

GetHits returns the Hits field if non-nil, zero value otherwise.

### GetHitsOk

`func (o *KnowledgeSearchOut) GetHitsOk() (*[]KnowledgeHit, bool)`

GetHitsOk returns a tuple with the Hits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHits

`func (o *KnowledgeSearchOut) SetHits(v []KnowledgeHit)`

SetHits sets Hits field to given value.

### HasHits

`func (o *KnowledgeSearchOut) HasHits() bool`

HasHits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


