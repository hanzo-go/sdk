# ReindexOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Failed** | Pointer to **int64** | Failed is how many documents could not be embedded; each is logged with its name, and the rest of the rebuild went on without it. | [optional] 
**Lexical** | Pointer to **int64** | Lexical is how many rows the org&#39;s lexical index holds now; 0 in a deployment without the index app. | [optional] 
**Removed** | Pointer to **int64** | Removed is how many rows the lexical index held for documents that no longer exist; 0 without the index app. | [optional] 
**Vectors** | Pointer to **int64** | Vectors is how many documents were embedded and written to the org&#39;s collection, which was dropped and created again at the configured size. | [optional] 

## Methods

### NewReindexOut

`func NewReindexOut() *ReindexOut`

NewReindexOut instantiates a new ReindexOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReindexOutWithDefaults

`func NewReindexOutWithDefaults() *ReindexOut`

NewReindexOutWithDefaults instantiates a new ReindexOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFailed

`func (o *ReindexOut) GetFailed() int64`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *ReindexOut) GetFailedOk() (*int64, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *ReindexOut) SetFailed(v int64)`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *ReindexOut) HasFailed() bool`

HasFailed returns a boolean if a field has been set.

### GetLexical

`func (o *ReindexOut) GetLexical() int64`

GetLexical returns the Lexical field if non-nil, zero value otherwise.

### GetLexicalOk

`func (o *ReindexOut) GetLexicalOk() (*int64, bool)`

GetLexicalOk returns a tuple with the Lexical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLexical

`func (o *ReindexOut) SetLexical(v int64)`

SetLexical sets Lexical field to given value.

### HasLexical

`func (o *ReindexOut) HasLexical() bool`

HasLexical returns a boolean if a field has been set.

### GetRemoved

`func (o *ReindexOut) GetRemoved() int64`

GetRemoved returns the Removed field if non-nil, zero value otherwise.

### GetRemovedOk

`func (o *ReindexOut) GetRemovedOk() (*int64, bool)`

GetRemovedOk returns a tuple with the Removed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoved

`func (o *ReindexOut) SetRemoved(v int64)`

SetRemoved sets Removed field to given value.

### HasRemoved

`func (o *ReindexOut) HasRemoved() bool`

HasRemoved returns a boolean if a field has been set.

### GetVectors

`func (o *ReindexOut) GetVectors() int64`

GetVectors returns the Vectors field if non-nil, zero value otherwise.

### GetVectorsOk

`func (o *ReindexOut) GetVectorsOk() (*int64, bool)`

GetVectorsOk returns a tuple with the Vectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVectors

`func (o *ReindexOut) SetVectors(v int64)`

SetVectors sets Vectors field to given value.

### HasVectors

`func (o *ReindexOut) HasVectors() bool`

HasVectors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


