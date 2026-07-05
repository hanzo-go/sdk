# MqKVBucket

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Config** | Pointer to [**MqKVBucketConfig**](MqKVBucketConfig.md) |  | [optional] 
**Values** | Pointer to **int32** | Number of keys in the bucket. | [optional] 
**Bytes** | Pointer to **int32** | Total bytes used. | [optional] 
**History** | Pointer to **int32** | Maximum revisions per key. | [optional] 
**Ttl** | Pointer to **string** | Default TTL for keys. | [optional] 
**BackingStream** | Pointer to **string** | Name of the underlying JetStream stream. | [optional] 
**Created** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewMqKVBucket

`func NewMqKVBucket() *MqKVBucket`

NewMqKVBucket instantiates a new MqKVBucket object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMqKVBucketWithDefaults

`func NewMqKVBucketWithDefaults() *MqKVBucket`

NewMqKVBucketWithDefaults instantiates a new MqKVBucket object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MqKVBucket) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MqKVBucket) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MqKVBucket) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MqKVBucket) HasName() bool`

HasName returns a boolean if a field has been set.

### GetConfig

`func (o *MqKVBucket) GetConfig() MqKVBucketConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *MqKVBucket) GetConfigOk() (*MqKVBucketConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *MqKVBucket) SetConfig(v MqKVBucketConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *MqKVBucket) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetValues

`func (o *MqKVBucket) GetValues() int32`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *MqKVBucket) GetValuesOk() (*int32, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *MqKVBucket) SetValues(v int32)`

SetValues sets Values field to given value.

### HasValues

`func (o *MqKVBucket) HasValues() bool`

HasValues returns a boolean if a field has been set.

### GetBytes

`func (o *MqKVBucket) GetBytes() int32`

GetBytes returns the Bytes field if non-nil, zero value otherwise.

### GetBytesOk

`func (o *MqKVBucket) GetBytesOk() (*int32, bool)`

GetBytesOk returns a tuple with the Bytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytes

`func (o *MqKVBucket) SetBytes(v int32)`

SetBytes sets Bytes field to given value.

### HasBytes

`func (o *MqKVBucket) HasBytes() bool`

HasBytes returns a boolean if a field has been set.

### GetHistory

`func (o *MqKVBucket) GetHistory() int32`

GetHistory returns the History field if non-nil, zero value otherwise.

### GetHistoryOk

`func (o *MqKVBucket) GetHistoryOk() (*int32, bool)`

GetHistoryOk returns a tuple with the History field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistory

`func (o *MqKVBucket) SetHistory(v int32)`

SetHistory sets History field to given value.

### HasHistory

`func (o *MqKVBucket) HasHistory() bool`

HasHistory returns a boolean if a field has been set.

### GetTtl

`func (o *MqKVBucket) GetTtl() string`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *MqKVBucket) GetTtlOk() (*string, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *MqKVBucket) SetTtl(v string)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *MqKVBucket) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetBackingStream

`func (o *MqKVBucket) GetBackingStream() string`

GetBackingStream returns the BackingStream field if non-nil, zero value otherwise.

### GetBackingStreamOk

`func (o *MqKVBucket) GetBackingStreamOk() (*string, bool)`

GetBackingStreamOk returns a tuple with the BackingStream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackingStream

`func (o *MqKVBucket) SetBackingStream(v string)`

SetBackingStream sets BackingStream field to given value.

### HasBackingStream

`func (o *MqKVBucket) HasBackingStream() bool`

HasBackingStream returns a boolean if a field has been set.

### GetCreated

`func (o *MqKVBucket) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *MqKVBucket) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *MqKVBucket) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *MqKVBucket) HasCreated() bool`

HasCreated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


