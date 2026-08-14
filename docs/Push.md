# Push

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**After** | Pointer to **string** |  | [optional] 
**Before** | Pointer to **string** |  | [optional] 
**Pusher** | Pointer to [**PushPusher**](PushPusher.md) |  | [optional] 
**Ref** | Pointer to **string** |  | [optional] 
**Repository** | Pointer to [**PushRepository**](PushRepository.md) |  | [optional] 

## Methods

### NewPush

`func NewPush() *Push`

NewPush instantiates a new Push object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPushWithDefaults

`func NewPushWithDefaults() *Push`

NewPushWithDefaults instantiates a new Push object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAfter

`func (o *Push) GetAfter() string`

GetAfter returns the After field if non-nil, zero value otherwise.

### GetAfterOk

`func (o *Push) GetAfterOk() (*string, bool)`

GetAfterOk returns a tuple with the After field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfter

`func (o *Push) SetAfter(v string)`

SetAfter sets After field to given value.

### HasAfter

`func (o *Push) HasAfter() bool`

HasAfter returns a boolean if a field has been set.

### GetBefore

`func (o *Push) GetBefore() string`

GetBefore returns the Before field if non-nil, zero value otherwise.

### GetBeforeOk

`func (o *Push) GetBeforeOk() (*string, bool)`

GetBeforeOk returns a tuple with the Before field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBefore

`func (o *Push) SetBefore(v string)`

SetBefore sets Before field to given value.

### HasBefore

`func (o *Push) HasBefore() bool`

HasBefore returns a boolean if a field has been set.

### GetPusher

`func (o *Push) GetPusher() PushPusher`

GetPusher returns the Pusher field if non-nil, zero value otherwise.

### GetPusherOk

`func (o *Push) GetPusherOk() (*PushPusher, bool)`

GetPusherOk returns a tuple with the Pusher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPusher

`func (o *Push) SetPusher(v PushPusher)`

SetPusher sets Pusher field to given value.

### HasPusher

`func (o *Push) HasPusher() bool`

HasPusher returns a boolean if a field has been set.

### GetRef

`func (o *Push) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *Push) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *Push) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *Push) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetRepository

`func (o *Push) GetRepository() PushRepository`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *Push) GetRepositoryOk() (*PushRepository, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *Push) SetRepository(v PushRepository)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *Push) HasRepository() bool`

HasRepository returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


