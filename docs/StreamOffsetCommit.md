# StreamOffsetCommit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Topic** | Pointer to **string** |  | [optional] 
**Partition** | Pointer to **int32** |  | [optional] 
**Offset** | Pointer to **int32** |  | [optional] 
**Metadata** | Pointer to **string** |  | [optional] 

## Methods

### NewStreamOffsetCommit

`func NewStreamOffsetCommit() *StreamOffsetCommit`

NewStreamOffsetCommit instantiates a new StreamOffsetCommit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamOffsetCommitWithDefaults

`func NewStreamOffsetCommitWithDefaults() *StreamOffsetCommit`

NewStreamOffsetCommitWithDefaults instantiates a new StreamOffsetCommit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTopic

`func (o *StreamOffsetCommit) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *StreamOffsetCommit) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *StreamOffsetCommit) SetTopic(v string)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *StreamOffsetCommit) HasTopic() bool`

HasTopic returns a boolean if a field has been set.

### GetPartition

`func (o *StreamOffsetCommit) GetPartition() int32`

GetPartition returns the Partition field if non-nil, zero value otherwise.

### GetPartitionOk

`func (o *StreamOffsetCommit) GetPartitionOk() (*int32, bool)`

GetPartitionOk returns a tuple with the Partition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartition

`func (o *StreamOffsetCommit) SetPartition(v int32)`

SetPartition sets Partition field to given value.

### HasPartition

`func (o *StreamOffsetCommit) HasPartition() bool`

HasPartition returns a boolean if a field has been set.

### GetOffset

`func (o *StreamOffsetCommit) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *StreamOffsetCommit) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *StreamOffsetCommit) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *StreamOffsetCommit) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetMetadata

`func (o *StreamOffsetCommit) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *StreamOffsetCommit) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *StreamOffsetCommit) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *StreamOffsetCommit) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


