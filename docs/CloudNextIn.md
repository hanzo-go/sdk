# CloudNextIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Batch** | Pointer to **int32** | Batch is how many messages to pull (1–1000, default 1). | [optional] 
**Expires** | Pointer to **string** | Expires is how long to wait for messages, e.g. \&quot;5s\&quot; (default \&quot;30s\&quot;, max \&quot;60s\&quot;). | [optional] 
**Name** | Pointer to **string** | Name is the consumer name, from the path. | [optional] 
**NoWait** | Pointer to **bool** | NoWait answers immediately with whatever is available instead of waiting. | [optional] 
**Stream** | Pointer to **string** | Stream is the stream name, from the path. | [optional] 

## Methods

### NewCloudNextIn

`func NewCloudNextIn() *CloudNextIn`

NewCloudNextIn instantiates a new CloudNextIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudNextInWithDefaults

`func NewCloudNextInWithDefaults() *CloudNextIn`

NewCloudNextInWithDefaults instantiates a new CloudNextIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBatch

`func (o *CloudNextIn) GetBatch() int32`

GetBatch returns the Batch field if non-nil, zero value otherwise.

### GetBatchOk

`func (o *CloudNextIn) GetBatchOk() (*int32, bool)`

GetBatchOk returns a tuple with the Batch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatch

`func (o *CloudNextIn) SetBatch(v int32)`

SetBatch sets Batch field to given value.

### HasBatch

`func (o *CloudNextIn) HasBatch() bool`

HasBatch returns a boolean if a field has been set.

### GetExpires

`func (o *CloudNextIn) GetExpires() string`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *CloudNextIn) GetExpiresOk() (*string, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *CloudNextIn) SetExpires(v string)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *CloudNextIn) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### GetName

`func (o *CloudNextIn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudNextIn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudNextIn) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudNextIn) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNoWait

`func (o *CloudNextIn) GetNoWait() bool`

GetNoWait returns the NoWait field if non-nil, zero value otherwise.

### GetNoWaitOk

`func (o *CloudNextIn) GetNoWaitOk() (*bool, bool)`

GetNoWaitOk returns a tuple with the NoWait field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoWait

`func (o *CloudNextIn) SetNoWait(v bool)`

SetNoWait sets NoWait field to given value.

### HasNoWait

`func (o *CloudNextIn) HasNoWait() bool`

HasNoWait returns a boolean if a field has been set.

### GetStream

`func (o *CloudNextIn) GetStream() string`

GetStream returns the Stream field if non-nil, zero value otherwise.

### GetStreamOk

`func (o *CloudNextIn) GetStreamOk() (*string, bool)`

GetStreamOk returns a tuple with the Stream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStream

`func (o *CloudNextIn) SetStream(v string)`

SetStream sets Stream field to given value.

### HasStream

`func (o *CloudNextIn) HasStream() bool`

HasStream returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


