# MqState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bytes** | Pointer to **int32** | Bytes is the total stored size. | [optional] 
**ConsumerCount** | Pointer to **int64** | Consumers is the number of consumers attached to this stream. | [optional] 
**FirstSeq** | Pointer to **int32** | FirstSeq is the sequence of the first stored message. | [optional] 
**FirstTs** | Pointer to **time.Time** | FirstTS is the timestamp of the first stored message. | [optional] 
**LastSeq** | Pointer to **int32** | LastSeq is the sequence of the last stored message. | [optional] 
**LastTs** | Pointer to **time.Time** | LastTS is the timestamp of the last stored message. | [optional] 
**Messages** | Pointer to **int32** | Messages is the number of messages currently stored. | [optional] 
**NumDeleted** | Pointer to **int64** | Deleted is the number of deleted messages (sequence gaps). | [optional] 
**NumSubjects** | Pointer to **int32** | Subjects is the number of distinct subjects stored. | [optional] 

## Methods

### NewMqState

`func NewMqState() *MqState`

NewMqState instantiates a new MqState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMqStateWithDefaults

`func NewMqStateWithDefaults() *MqState`

NewMqStateWithDefaults instantiates a new MqState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBytes

`func (o *MqState) GetBytes() int32`

GetBytes returns the Bytes field if non-nil, zero value otherwise.

### GetBytesOk

`func (o *MqState) GetBytesOk() (*int32, bool)`

GetBytesOk returns a tuple with the Bytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytes

`func (o *MqState) SetBytes(v int32)`

SetBytes sets Bytes field to given value.

### HasBytes

`func (o *MqState) HasBytes() bool`

HasBytes returns a boolean if a field has been set.

### GetConsumerCount

`func (o *MqState) GetConsumerCount() int64`

GetConsumerCount returns the ConsumerCount field if non-nil, zero value otherwise.

### GetConsumerCountOk

`func (o *MqState) GetConsumerCountOk() (*int64, bool)`

GetConsumerCountOk returns a tuple with the ConsumerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerCount

`func (o *MqState) SetConsumerCount(v int64)`

SetConsumerCount sets ConsumerCount field to given value.

### HasConsumerCount

`func (o *MqState) HasConsumerCount() bool`

HasConsumerCount returns a boolean if a field has been set.

### GetFirstSeq

`func (o *MqState) GetFirstSeq() int32`

GetFirstSeq returns the FirstSeq field if non-nil, zero value otherwise.

### GetFirstSeqOk

`func (o *MqState) GetFirstSeqOk() (*int32, bool)`

GetFirstSeqOk returns a tuple with the FirstSeq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstSeq

`func (o *MqState) SetFirstSeq(v int32)`

SetFirstSeq sets FirstSeq field to given value.

### HasFirstSeq

`func (o *MqState) HasFirstSeq() bool`

HasFirstSeq returns a boolean if a field has been set.

### GetFirstTs

`func (o *MqState) GetFirstTs() time.Time`

GetFirstTs returns the FirstTs field if non-nil, zero value otherwise.

### GetFirstTsOk

`func (o *MqState) GetFirstTsOk() (*time.Time, bool)`

GetFirstTsOk returns a tuple with the FirstTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstTs

`func (o *MqState) SetFirstTs(v time.Time)`

SetFirstTs sets FirstTs field to given value.

### HasFirstTs

`func (o *MqState) HasFirstTs() bool`

HasFirstTs returns a boolean if a field has been set.

### GetLastSeq

`func (o *MqState) GetLastSeq() int32`

GetLastSeq returns the LastSeq field if non-nil, zero value otherwise.

### GetLastSeqOk

`func (o *MqState) GetLastSeqOk() (*int32, bool)`

GetLastSeqOk returns a tuple with the LastSeq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeq

`func (o *MqState) SetLastSeq(v int32)`

SetLastSeq sets LastSeq field to given value.

### HasLastSeq

`func (o *MqState) HasLastSeq() bool`

HasLastSeq returns a boolean if a field has been set.

### GetLastTs

`func (o *MqState) GetLastTs() time.Time`

GetLastTs returns the LastTs field if non-nil, zero value otherwise.

### GetLastTsOk

`func (o *MqState) GetLastTsOk() (*time.Time, bool)`

GetLastTsOk returns a tuple with the LastTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTs

`func (o *MqState) SetLastTs(v time.Time)`

SetLastTs sets LastTs field to given value.

### HasLastTs

`func (o *MqState) HasLastTs() bool`

HasLastTs returns a boolean if a field has been set.

### GetMessages

`func (o *MqState) GetMessages() int32`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *MqState) GetMessagesOk() (*int32, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *MqState) SetMessages(v int32)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *MqState) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### GetNumDeleted

`func (o *MqState) GetNumDeleted() int64`

GetNumDeleted returns the NumDeleted field if non-nil, zero value otherwise.

### GetNumDeletedOk

`func (o *MqState) GetNumDeletedOk() (*int64, bool)`

GetNumDeletedOk returns a tuple with the NumDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumDeleted

`func (o *MqState) SetNumDeleted(v int64)`

SetNumDeleted sets NumDeleted field to given value.

### HasNumDeleted

`func (o *MqState) HasNumDeleted() bool`

HasNumDeleted returns a boolean if a field has been set.

### GetNumSubjects

`func (o *MqState) GetNumSubjects() int32`

GetNumSubjects returns the NumSubjects field if non-nil, zero value otherwise.

### GetNumSubjectsOk

`func (o *MqState) GetNumSubjectsOk() (*int32, bool)`

GetNumSubjectsOk returns a tuple with the NumSubjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumSubjects

`func (o *MqState) SetNumSubjects(v int32)`

SetNumSubjects sets NumSubjects field to given value.

### HasNumSubjects

`func (o *MqState) HasNumSubjects() bool`

HasNumSubjects returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


