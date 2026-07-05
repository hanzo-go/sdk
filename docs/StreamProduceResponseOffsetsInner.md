# StreamProduceResponseOffsetsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Partition** | Pointer to **int32** |  | [optional] 
**Offset** | Pointer to **int32** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewStreamProduceResponseOffsetsInner

`func NewStreamProduceResponseOffsetsInner() *StreamProduceResponseOffsetsInner`

NewStreamProduceResponseOffsetsInner instantiates a new StreamProduceResponseOffsetsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamProduceResponseOffsetsInnerWithDefaults

`func NewStreamProduceResponseOffsetsInnerWithDefaults() *StreamProduceResponseOffsetsInner`

NewStreamProduceResponseOffsetsInnerWithDefaults instantiates a new StreamProduceResponseOffsetsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPartition

`func (o *StreamProduceResponseOffsetsInner) GetPartition() int32`

GetPartition returns the Partition field if non-nil, zero value otherwise.

### GetPartitionOk

`func (o *StreamProduceResponseOffsetsInner) GetPartitionOk() (*int32, bool)`

GetPartitionOk returns a tuple with the Partition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartition

`func (o *StreamProduceResponseOffsetsInner) SetPartition(v int32)`

SetPartition sets Partition field to given value.

### HasPartition

`func (o *StreamProduceResponseOffsetsInner) HasPartition() bool`

HasPartition returns a boolean if a field has been set.

### GetOffset

`func (o *StreamProduceResponseOffsetsInner) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *StreamProduceResponseOffsetsInner) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *StreamProduceResponseOffsetsInner) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *StreamProduceResponseOffsetsInner) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetTimestamp

`func (o *StreamProduceResponseOffsetsInner) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *StreamProduceResponseOffsetsInner) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *StreamProduceResponseOffsetsInner) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *StreamProduceResponseOffsetsInner) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


