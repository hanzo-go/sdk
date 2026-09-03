# NextIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Batch** | Pointer to **int64** | Batch is how many messages to pull (1–1000, default 1). | [optional] 
**Expires** | Pointer to **string** | Expires is how long to wait for messages, e.g. \&quot;5s\&quot; (default \&quot;30s\&quot;, max \&quot;60s\&quot;). | [optional] 
**Name** | Pointer to **string** | Name is the consumer name, from the path. | [optional] 
**NoWait** | Pointer to **bool** | NoWait answers immediately with whatever is available instead of waiting. | [optional] 
**Stream** | Pointer to **string** | Stream is the stream name, from the path. | [optional] 

## Methods

### NewNextIn

`func NewNextIn() *NextIn`

NewNextIn instantiates a new NextIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNextInWithDefaults

`func NewNextInWithDefaults() *NextIn`

NewNextInWithDefaults instantiates a new NextIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBatch

`func (o *NextIn) GetBatch() int64`

GetBatch returns the Batch field if non-nil, zero value otherwise.

### GetBatchOk

`func (o *NextIn) GetBatchOk() (*int64, bool)`

GetBatchOk returns a tuple with the Batch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatch

`func (o *NextIn) SetBatch(v int64)`

SetBatch sets Batch field to given value.

### HasBatch

`func (o *NextIn) HasBatch() bool`

HasBatch returns a boolean if a field has been set.

### GetExpires

`func (o *NextIn) GetExpires() string`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *NextIn) GetExpiresOk() (*string, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *NextIn) SetExpires(v string)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *NextIn) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### GetName

`func (o *NextIn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NextIn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NextIn) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NextIn) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNoWait

`func (o *NextIn) GetNoWait() bool`

GetNoWait returns the NoWait field if non-nil, zero value otherwise.

### GetNoWaitOk

`func (o *NextIn) GetNoWaitOk() (*bool, bool)`

GetNoWaitOk returns a tuple with the NoWait field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoWait

`func (o *NextIn) SetNoWait(v bool)`

SetNoWait sets NoWait field to given value.

### HasNoWait

`func (o *NextIn) HasNoWait() bool`

HasNoWait returns a boolean if a field has been set.

### GetStream

`func (o *NextIn) GetStream() string`

GetStream returns the Stream field if non-nil, zero value otherwise.

### GetStreamOk

`func (o *NextIn) GetStreamOk() (*string, bool)`

GetStreamOk returns a tuple with the Stream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStream

`func (o *NextIn) SetStream(v string)`

SetStream sets Stream field to given value.

### HasStream

`func (o *NextIn) HasStream() bool`

HasStream returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


