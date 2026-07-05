# StreamConsumeRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Topic** | Pointer to **string** |  | [optional] 
**Partition** | Pointer to **int32** |  | [optional] 
**Offset** | Pointer to **int32** |  | [optional] 
**Key** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Headers** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewStreamConsumeRecord

`func NewStreamConsumeRecord() *StreamConsumeRecord`

NewStreamConsumeRecord instantiates a new StreamConsumeRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamConsumeRecordWithDefaults

`func NewStreamConsumeRecordWithDefaults() *StreamConsumeRecord`

NewStreamConsumeRecordWithDefaults instantiates a new StreamConsumeRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTopic

`func (o *StreamConsumeRecord) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *StreamConsumeRecord) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *StreamConsumeRecord) SetTopic(v string)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *StreamConsumeRecord) HasTopic() bool`

HasTopic returns a boolean if a field has been set.

### GetPartition

`func (o *StreamConsumeRecord) GetPartition() int32`

GetPartition returns the Partition field if non-nil, zero value otherwise.

### GetPartitionOk

`func (o *StreamConsumeRecord) GetPartitionOk() (*int32, bool)`

GetPartitionOk returns a tuple with the Partition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartition

`func (o *StreamConsumeRecord) SetPartition(v int32)`

SetPartition sets Partition field to given value.

### HasPartition

`func (o *StreamConsumeRecord) HasPartition() bool`

HasPartition returns a boolean if a field has been set.

### GetOffset

`func (o *StreamConsumeRecord) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *StreamConsumeRecord) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *StreamConsumeRecord) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *StreamConsumeRecord) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetKey

`func (o *StreamConsumeRecord) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *StreamConsumeRecord) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *StreamConsumeRecord) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *StreamConsumeRecord) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetValue

`func (o *StreamConsumeRecord) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *StreamConsumeRecord) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *StreamConsumeRecord) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *StreamConsumeRecord) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetTimestamp

`func (o *StreamConsumeRecord) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *StreamConsumeRecord) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *StreamConsumeRecord) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *StreamConsumeRecord) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetHeaders

`func (o *StreamConsumeRecord) GetHeaders() map[string]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *StreamConsumeRecord) GetHeadersOk() (*map[string]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *StreamConsumeRecord) SetHeaders(v map[string]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *StreamConsumeRecord) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


