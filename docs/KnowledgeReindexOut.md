# KnowledgeReindexOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Failed** | Pointer to **int64** | Failed is how many documents could not be indexed; each is logged with its name and has no passages — the rest of the rebuild went on without it. | [optional] 
**Lexical** | Pointer to **int64** | Lexical is how many rows the org&#39;s lexical index holds now; 0 in a deployment without the index app. | [optional] 
**Removed** | Pointer to **int64** | Removed is how many rows the lexical index held for documents that no longer exist; 0 without the index app. | [optional] 
**Vectors** | Pointer to **int64** | Vectors is how many documents were read back, cut into passages, embedded and written to the org&#39;s passage store, which was emptied first — so it now holds these documents&#39; passages and nothing else. | [optional] 

## Methods

### NewKnowledgeReindexOut

`func NewKnowledgeReindexOut() *KnowledgeReindexOut`

NewKnowledgeReindexOut instantiates a new KnowledgeReindexOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKnowledgeReindexOutWithDefaults

`func NewKnowledgeReindexOutWithDefaults() *KnowledgeReindexOut`

NewKnowledgeReindexOutWithDefaults instantiates a new KnowledgeReindexOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFailed

`func (o *KnowledgeReindexOut) GetFailed() int64`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *KnowledgeReindexOut) GetFailedOk() (*int64, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *KnowledgeReindexOut) SetFailed(v int64)`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *KnowledgeReindexOut) HasFailed() bool`

HasFailed returns a boolean if a field has been set.

### GetLexical

`func (o *KnowledgeReindexOut) GetLexical() int64`

GetLexical returns the Lexical field if non-nil, zero value otherwise.

### GetLexicalOk

`func (o *KnowledgeReindexOut) GetLexicalOk() (*int64, bool)`

GetLexicalOk returns a tuple with the Lexical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLexical

`func (o *KnowledgeReindexOut) SetLexical(v int64)`

SetLexical sets Lexical field to given value.

### HasLexical

`func (o *KnowledgeReindexOut) HasLexical() bool`

HasLexical returns a boolean if a field has been set.

### GetRemoved

`func (o *KnowledgeReindexOut) GetRemoved() int64`

GetRemoved returns the Removed field if non-nil, zero value otherwise.

### GetRemovedOk

`func (o *KnowledgeReindexOut) GetRemovedOk() (*int64, bool)`

GetRemovedOk returns a tuple with the Removed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoved

`func (o *KnowledgeReindexOut) SetRemoved(v int64)`

SetRemoved sets Removed field to given value.

### HasRemoved

`func (o *KnowledgeReindexOut) HasRemoved() bool`

HasRemoved returns a boolean if a field has been set.

### GetVectors

`func (o *KnowledgeReindexOut) GetVectors() int64`

GetVectors returns the Vectors field if non-nil, zero value otherwise.

### GetVectorsOk

`func (o *KnowledgeReindexOut) GetVectorsOk() (*int64, bool)`

GetVectorsOk returns a tuple with the Vectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVectors

`func (o *KnowledgeReindexOut) SetVectors(v int64)`

SetVectors sets Vectors field to given value.

### HasVectors

`func (o *KnowledgeReindexOut) HasVectors() bool`

HasVectors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


