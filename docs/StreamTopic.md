# StreamTopic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Partitions** | Pointer to **int32** |  | [optional] 
**ReplicationFactor** | Pointer to **int32** |  | [optional] 
**Config** | Pointer to [**StreamTopicConfig**](StreamTopicConfig.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewStreamTopic

`func NewStreamTopic() *StreamTopic`

NewStreamTopic instantiates a new StreamTopic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamTopicWithDefaults

`func NewStreamTopicWithDefaults() *StreamTopic`

NewStreamTopicWithDefaults instantiates a new StreamTopic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *StreamTopic) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StreamTopic) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StreamTopic) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StreamTopic) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPartitions

`func (o *StreamTopic) GetPartitions() int32`

GetPartitions returns the Partitions field if non-nil, zero value otherwise.

### GetPartitionsOk

`func (o *StreamTopic) GetPartitionsOk() (*int32, bool)`

GetPartitionsOk returns a tuple with the Partitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitions

`func (o *StreamTopic) SetPartitions(v int32)`

SetPartitions sets Partitions field to given value.

### HasPartitions

`func (o *StreamTopic) HasPartitions() bool`

HasPartitions returns a boolean if a field has been set.

### GetReplicationFactor

`func (o *StreamTopic) GetReplicationFactor() int32`

GetReplicationFactor returns the ReplicationFactor field if non-nil, zero value otherwise.

### GetReplicationFactorOk

`func (o *StreamTopic) GetReplicationFactorOk() (*int32, bool)`

GetReplicationFactorOk returns a tuple with the ReplicationFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationFactor

`func (o *StreamTopic) SetReplicationFactor(v int32)`

SetReplicationFactor sets ReplicationFactor field to given value.

### HasReplicationFactor

`func (o *StreamTopic) HasReplicationFactor() bool`

HasReplicationFactor returns a boolean if a field has been set.

### GetConfig

`func (o *StreamTopic) GetConfig() StreamTopicConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *StreamTopic) GetConfigOk() (*StreamTopicConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *StreamTopic) SetConfig(v StreamTopicConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *StreamTopic) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetCreatedAt

`func (o *StreamTopic) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *StreamTopic) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *StreamTopic) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *StreamTopic) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


