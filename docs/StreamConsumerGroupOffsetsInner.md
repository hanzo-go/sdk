# StreamConsumerGroupOffsetsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Topic** | Pointer to **string** |  | [optional] 
**Partition** | Pointer to **int32** |  | [optional] 
**CommittedOffset** | Pointer to **int32** |  | [optional] 
**Lag** | Pointer to **int32** |  | [optional] 

## Methods

### NewStreamConsumerGroupOffsetsInner

`func NewStreamConsumerGroupOffsetsInner() *StreamConsumerGroupOffsetsInner`

NewStreamConsumerGroupOffsetsInner instantiates a new StreamConsumerGroupOffsetsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamConsumerGroupOffsetsInnerWithDefaults

`func NewStreamConsumerGroupOffsetsInnerWithDefaults() *StreamConsumerGroupOffsetsInner`

NewStreamConsumerGroupOffsetsInnerWithDefaults instantiates a new StreamConsumerGroupOffsetsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTopic

`func (o *StreamConsumerGroupOffsetsInner) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *StreamConsumerGroupOffsetsInner) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *StreamConsumerGroupOffsetsInner) SetTopic(v string)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *StreamConsumerGroupOffsetsInner) HasTopic() bool`

HasTopic returns a boolean if a field has been set.

### GetPartition

`func (o *StreamConsumerGroupOffsetsInner) GetPartition() int32`

GetPartition returns the Partition field if non-nil, zero value otherwise.

### GetPartitionOk

`func (o *StreamConsumerGroupOffsetsInner) GetPartitionOk() (*int32, bool)`

GetPartitionOk returns a tuple with the Partition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartition

`func (o *StreamConsumerGroupOffsetsInner) SetPartition(v int32)`

SetPartition sets Partition field to given value.

### HasPartition

`func (o *StreamConsumerGroupOffsetsInner) HasPartition() bool`

HasPartition returns a boolean if a field has been set.

### GetCommittedOffset

`func (o *StreamConsumerGroupOffsetsInner) GetCommittedOffset() int32`

GetCommittedOffset returns the CommittedOffset field if non-nil, zero value otherwise.

### GetCommittedOffsetOk

`func (o *StreamConsumerGroupOffsetsInner) GetCommittedOffsetOk() (*int32, bool)`

GetCommittedOffsetOk returns a tuple with the CommittedOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommittedOffset

`func (o *StreamConsumerGroupOffsetsInner) SetCommittedOffset(v int32)`

SetCommittedOffset sets CommittedOffset field to given value.

### HasCommittedOffset

`func (o *StreamConsumerGroupOffsetsInner) HasCommittedOffset() bool`

HasCommittedOffset returns a boolean if a field has been set.

### GetLag

`func (o *StreamConsumerGroupOffsetsInner) GetLag() int32`

GetLag returns the Lag field if non-nil, zero value otherwise.

### GetLagOk

`func (o *StreamConsumerGroupOffsetsInner) GetLagOk() (*int32, bool)`

GetLagOk returns a tuple with the Lag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLag

`func (o *StreamConsumerGroupOffsetsInner) SetLag(v int32)`

SetLag sets Lag field to given value.

### HasLag

`func (o *StreamConsumerGroupOffsetsInner) HasLag() bool`

HasLag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


