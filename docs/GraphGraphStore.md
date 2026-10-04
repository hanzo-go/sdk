# GraphGraphStore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Detail** | Pointer to **string** | Detail says what an unsupported store is missing before an erasure could reach it. | [optional] 
**Removed** | Pointer to **int64** | Removed is how many of its rows were removed. | [optional] 
**Status** | Pointer to **string** | Status is erased (rows naming the entity were removed), absent (it held none) or unsupported (this op cannot reach it, so it may still hold the subject). | [optional] 
**Store** | Pointer to **string** | Store names it: &#x60;graph&#x60; is the caller&#39;s own (org, project) assertion file, its full-text index included; &#x60;knowledge&#x60; is the org&#39;s wiki pages, memories and sources with their vectors; &#x60;ai.memory&#x60; is /v1/ai/memory. | [optional] 

## Methods

### NewGraphGraphStore

`func NewGraphGraphStore() *GraphGraphStore`

NewGraphGraphStore instantiates a new GraphGraphStore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphGraphStoreWithDefaults

`func NewGraphGraphStoreWithDefaults() *GraphGraphStore`

NewGraphGraphStoreWithDefaults instantiates a new GraphGraphStore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetail

`func (o *GraphGraphStore) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *GraphGraphStore) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *GraphGraphStore) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *GraphGraphStore) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetRemoved

`func (o *GraphGraphStore) GetRemoved() int64`

GetRemoved returns the Removed field if non-nil, zero value otherwise.

### GetRemovedOk

`func (o *GraphGraphStore) GetRemovedOk() (*int64, bool)`

GetRemovedOk returns a tuple with the Removed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoved

`func (o *GraphGraphStore) SetRemoved(v int64)`

SetRemoved sets Removed field to given value.

### HasRemoved

`func (o *GraphGraphStore) HasRemoved() bool`

HasRemoved returns a boolean if a field has been set.

### GetStatus

`func (o *GraphGraphStore) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GraphGraphStore) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GraphGraphStore) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GraphGraphStore) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStore

`func (o *GraphGraphStore) GetStore() string`

GetStore returns the Store field if non-nil, zero value otherwise.

### GetStoreOk

`func (o *GraphGraphStore) GetStoreOk() (*string, bool)`

GetStoreOk returns a tuple with the Store field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStore

`func (o *GraphGraphStore) SetStore(v string)`

SetStore sets Store field to given value.

### HasStore

`func (o *GraphGraphStore) HasStore() bool`

HasStore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


