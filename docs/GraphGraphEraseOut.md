# GraphGraphEraseOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**At** | Pointer to **string** | At is the instant of the erasure, RFC 3339, and the time on its audit record. | [optional] 
**By** | Pointer to **string** | By is the identity that erased — &#x60;owner&#x60; or &#x60;owner/user&#x60; — from the validated principal. | [optional] 
**Complete** | Pointer to **bool** | Complete is true only when no store may still hold the subject: every store erased or absent and nothing Remaining. While a store is unsupported it is false, and that store is handled out of band before the erasure is done. | [optional] 
**Digest** | Pointer to **string** | Digest is SHA-256, hex, over IDs in their ascending order, each written &#x60;&lt;length&gt;:&lt;id&gt;&#x60; — the same content address an assertion&#39;s own ID is. The audit trail records it in place of the ids, so the receipt verifies against the trail and the trail names nothing. | [optional] 
**Entity** | Pointer to **string** | Entity is the key that was erased, folded as every key here is (NFC, no control characters) and echoed so a stored receipt says what it is for. This plane keeps no copy of it. | [optional] 
**Ids** | Pointer to **[]string** | IDs are the content addresses of the assertions removed, ascending. | [optional] 
**Remaining** | Pointer to **int64** | Remaining is how many assertions still name the entity because one call removes at most ten thousand. Non-zero means call again; each call is its own receipt. | [optional] 
**Sources** | Pointer to **[]string** | Sources are the documents the removed assertions were filed from. This plane does not own them: ingesting one again files the subject again. | [optional] 
**Stores** | Pointer to [**[]GraphGraphStore**](GraphGraphStore.md) | Stores is every store in this deployment that may hold what the org says about the subject, and what happened in each: this graph first, then the stores this op cannot reach. | [optional] 

## Methods

### NewGraphGraphEraseOut

`func NewGraphGraphEraseOut() *GraphGraphEraseOut`

NewGraphGraphEraseOut instantiates a new GraphGraphEraseOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphGraphEraseOutWithDefaults

`func NewGraphGraphEraseOutWithDefaults() *GraphGraphEraseOut`

NewGraphGraphEraseOutWithDefaults instantiates a new GraphGraphEraseOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAt

`func (o *GraphGraphEraseOut) GetAt() string`

GetAt returns the At field if non-nil, zero value otherwise.

### GetAtOk

`func (o *GraphGraphEraseOut) GetAtOk() (*string, bool)`

GetAtOk returns a tuple with the At field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAt

`func (o *GraphGraphEraseOut) SetAt(v string)`

SetAt sets At field to given value.

### HasAt

`func (o *GraphGraphEraseOut) HasAt() bool`

HasAt returns a boolean if a field has been set.

### GetBy

`func (o *GraphGraphEraseOut) GetBy() string`

GetBy returns the By field if non-nil, zero value otherwise.

### GetByOk

`func (o *GraphGraphEraseOut) GetByOk() (*string, bool)`

GetByOk returns a tuple with the By field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBy

`func (o *GraphGraphEraseOut) SetBy(v string)`

SetBy sets By field to given value.

### HasBy

`func (o *GraphGraphEraseOut) HasBy() bool`

HasBy returns a boolean if a field has been set.

### GetComplete

`func (o *GraphGraphEraseOut) GetComplete() bool`

GetComplete returns the Complete field if non-nil, zero value otherwise.

### GetCompleteOk

`func (o *GraphGraphEraseOut) GetCompleteOk() (*bool, bool)`

GetCompleteOk returns a tuple with the Complete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplete

`func (o *GraphGraphEraseOut) SetComplete(v bool)`

SetComplete sets Complete field to given value.

### HasComplete

`func (o *GraphGraphEraseOut) HasComplete() bool`

HasComplete returns a boolean if a field has been set.

### GetDigest

`func (o *GraphGraphEraseOut) GetDigest() string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *GraphGraphEraseOut) GetDigestOk() (*string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *GraphGraphEraseOut) SetDigest(v string)`

SetDigest sets Digest field to given value.

### HasDigest

`func (o *GraphGraphEraseOut) HasDigest() bool`

HasDigest returns a boolean if a field has been set.

### GetEntity

`func (o *GraphGraphEraseOut) GetEntity() string`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *GraphGraphEraseOut) GetEntityOk() (*string, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *GraphGraphEraseOut) SetEntity(v string)`

SetEntity sets Entity field to given value.

### HasEntity

`func (o *GraphGraphEraseOut) HasEntity() bool`

HasEntity returns a boolean if a field has been set.

### GetIds

`func (o *GraphGraphEraseOut) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *GraphGraphEraseOut) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *GraphGraphEraseOut) SetIds(v []string)`

SetIds sets Ids field to given value.

### HasIds

`func (o *GraphGraphEraseOut) HasIds() bool`

HasIds returns a boolean if a field has been set.

### GetRemaining

`func (o *GraphGraphEraseOut) GetRemaining() int64`

GetRemaining returns the Remaining field if non-nil, zero value otherwise.

### GetRemainingOk

`func (o *GraphGraphEraseOut) GetRemainingOk() (*int64, bool)`

GetRemainingOk returns a tuple with the Remaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemaining

`func (o *GraphGraphEraseOut) SetRemaining(v int64)`

SetRemaining sets Remaining field to given value.

### HasRemaining

`func (o *GraphGraphEraseOut) HasRemaining() bool`

HasRemaining returns a boolean if a field has been set.

### GetSources

`func (o *GraphGraphEraseOut) GetSources() []string`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *GraphGraphEraseOut) GetSourcesOk() (*[]string, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *GraphGraphEraseOut) SetSources(v []string)`

SetSources sets Sources field to given value.

### HasSources

`func (o *GraphGraphEraseOut) HasSources() bool`

HasSources returns a boolean if a field has been set.

### GetStores

`func (o *GraphGraphEraseOut) GetStores() []GraphGraphStore`

GetStores returns the Stores field if non-nil, zero value otherwise.

### GetStoresOk

`func (o *GraphGraphEraseOut) GetStoresOk() (*[]GraphGraphStore, bool)`

GetStoresOk returns a tuple with the Stores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStores

`func (o *GraphGraphEraseOut) SetStores(v []GraphGraphStore)`

SetStores sets Stores field to given value.

### HasStores

`func (o *GraphGraphEraseOut) HasStores() bool`

HasStores returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


