# StreamCreateTopicRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Topic name | 
**Partitions** | Pointer to **int32** | Number of partitions | [optional] [default to 1]
**ReplicationFactor** | Pointer to **int32** | Replication factor for underlying PubSub streams | [optional] [default to 1]
**Config** | Pointer to [**StreamCreateTopicRequestConfig**](StreamCreateTopicRequestConfig.md) |  | [optional] 

## Methods

### NewStreamCreateTopicRequest

`func NewStreamCreateTopicRequest(name string, ) *StreamCreateTopicRequest`

NewStreamCreateTopicRequest instantiates a new StreamCreateTopicRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamCreateTopicRequestWithDefaults

`func NewStreamCreateTopicRequestWithDefaults() *StreamCreateTopicRequest`

NewStreamCreateTopicRequestWithDefaults instantiates a new StreamCreateTopicRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *StreamCreateTopicRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StreamCreateTopicRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StreamCreateTopicRequest) SetName(v string)`

SetName sets Name field to given value.


### GetPartitions

`func (o *StreamCreateTopicRequest) GetPartitions() int32`

GetPartitions returns the Partitions field if non-nil, zero value otherwise.

### GetPartitionsOk

`func (o *StreamCreateTopicRequest) GetPartitionsOk() (*int32, bool)`

GetPartitionsOk returns a tuple with the Partitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitions

`func (o *StreamCreateTopicRequest) SetPartitions(v int32)`

SetPartitions sets Partitions field to given value.

### HasPartitions

`func (o *StreamCreateTopicRequest) HasPartitions() bool`

HasPartitions returns a boolean if a field has been set.

### GetReplicationFactor

`func (o *StreamCreateTopicRequest) GetReplicationFactor() int32`

GetReplicationFactor returns the ReplicationFactor field if non-nil, zero value otherwise.

### GetReplicationFactorOk

`func (o *StreamCreateTopicRequest) GetReplicationFactorOk() (*int32, bool)`

GetReplicationFactorOk returns a tuple with the ReplicationFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationFactor

`func (o *StreamCreateTopicRequest) SetReplicationFactor(v int32)`

SetReplicationFactor sets ReplicationFactor field to given value.

### HasReplicationFactor

`func (o *StreamCreateTopicRequest) HasReplicationFactor() bool`

HasReplicationFactor returns a boolean if a field has been set.

### GetConfig

`func (o *StreamCreateTopicRequest) GetConfig() StreamCreateTopicRequestConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *StreamCreateTopicRequest) GetConfigOk() (*StreamCreateTopicRequestConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *StreamCreateTopicRequest) SetConfig(v StreamCreateTopicRequestConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *StreamCreateTopicRequest) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


