# MqStreamState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Messages** | Pointer to **int32** | Total number of messages in the stream. | [optional] 
**Bytes** | Pointer to **int32** | Total bytes used by the stream. | [optional] 
**FirstSeq** | Pointer to **int32** | Sequence number of the first message. | [optional] 
**FirstTs** | Pointer to **time.Time** | Timestamp of the first message. | [optional] 
**LastSeq** | Pointer to **int32** | Sequence number of the last message. | [optional] 
**LastTs** | Pointer to **time.Time** | Timestamp of the last message. | [optional] 
**ConsumerCount** | Pointer to **int32** | Number of consumers attached to this stream. | [optional] 
**NumSubjects** | Pointer to **int32** | Number of unique subjects in the stream. | [optional] 
**NumDeleted** | Pointer to **int32** | Number of deleted messages (gaps in sequence). | [optional] 

## Methods

### NewMqStreamState

`func NewMqStreamState() *MqStreamState`

NewMqStreamState instantiates a new MqStreamState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMqStreamStateWithDefaults

`func NewMqStreamStateWithDefaults() *MqStreamState`

NewMqStreamStateWithDefaults instantiates a new MqStreamState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessages

`func (o *MqStreamState) GetMessages() int32`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *MqStreamState) GetMessagesOk() (*int32, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *MqStreamState) SetMessages(v int32)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *MqStreamState) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### GetBytes

`func (o *MqStreamState) GetBytes() int32`

GetBytes returns the Bytes field if non-nil, zero value otherwise.

### GetBytesOk

`func (o *MqStreamState) GetBytesOk() (*int32, bool)`

GetBytesOk returns a tuple with the Bytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytes

`func (o *MqStreamState) SetBytes(v int32)`

SetBytes sets Bytes field to given value.

### HasBytes

`func (o *MqStreamState) HasBytes() bool`

HasBytes returns a boolean if a field has been set.

### GetFirstSeq

`func (o *MqStreamState) GetFirstSeq() int32`

GetFirstSeq returns the FirstSeq field if non-nil, zero value otherwise.

### GetFirstSeqOk

`func (o *MqStreamState) GetFirstSeqOk() (*int32, bool)`

GetFirstSeqOk returns a tuple with the FirstSeq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstSeq

`func (o *MqStreamState) SetFirstSeq(v int32)`

SetFirstSeq sets FirstSeq field to given value.

### HasFirstSeq

`func (o *MqStreamState) HasFirstSeq() bool`

HasFirstSeq returns a boolean if a field has been set.

### GetFirstTs

`func (o *MqStreamState) GetFirstTs() time.Time`

GetFirstTs returns the FirstTs field if non-nil, zero value otherwise.

### GetFirstTsOk

`func (o *MqStreamState) GetFirstTsOk() (*time.Time, bool)`

GetFirstTsOk returns a tuple with the FirstTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstTs

`func (o *MqStreamState) SetFirstTs(v time.Time)`

SetFirstTs sets FirstTs field to given value.

### HasFirstTs

`func (o *MqStreamState) HasFirstTs() bool`

HasFirstTs returns a boolean if a field has been set.

### GetLastSeq

`func (o *MqStreamState) GetLastSeq() int32`

GetLastSeq returns the LastSeq field if non-nil, zero value otherwise.

### GetLastSeqOk

`func (o *MqStreamState) GetLastSeqOk() (*int32, bool)`

GetLastSeqOk returns a tuple with the LastSeq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeq

`func (o *MqStreamState) SetLastSeq(v int32)`

SetLastSeq sets LastSeq field to given value.

### HasLastSeq

`func (o *MqStreamState) HasLastSeq() bool`

HasLastSeq returns a boolean if a field has been set.

### GetLastTs

`func (o *MqStreamState) GetLastTs() time.Time`

GetLastTs returns the LastTs field if non-nil, zero value otherwise.

### GetLastTsOk

`func (o *MqStreamState) GetLastTsOk() (*time.Time, bool)`

GetLastTsOk returns a tuple with the LastTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTs

`func (o *MqStreamState) SetLastTs(v time.Time)`

SetLastTs sets LastTs field to given value.

### HasLastTs

`func (o *MqStreamState) HasLastTs() bool`

HasLastTs returns a boolean if a field has been set.

### GetConsumerCount

`func (o *MqStreamState) GetConsumerCount() int32`

GetConsumerCount returns the ConsumerCount field if non-nil, zero value otherwise.

### GetConsumerCountOk

`func (o *MqStreamState) GetConsumerCountOk() (*int32, bool)`

GetConsumerCountOk returns a tuple with the ConsumerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerCount

`func (o *MqStreamState) SetConsumerCount(v int32)`

SetConsumerCount sets ConsumerCount field to given value.

### HasConsumerCount

`func (o *MqStreamState) HasConsumerCount() bool`

HasConsumerCount returns a boolean if a field has been set.

### GetNumSubjects

`func (o *MqStreamState) GetNumSubjects() int32`

GetNumSubjects returns the NumSubjects field if non-nil, zero value otherwise.

### GetNumSubjectsOk

`func (o *MqStreamState) GetNumSubjectsOk() (*int32, bool)`

GetNumSubjectsOk returns a tuple with the NumSubjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumSubjects

`func (o *MqStreamState) SetNumSubjects(v int32)`

SetNumSubjects sets NumSubjects field to given value.

### HasNumSubjects

`func (o *MqStreamState) HasNumSubjects() bool`

HasNumSubjects returns a boolean if a field has been set.

### GetNumDeleted

`func (o *MqStreamState) GetNumDeleted() int32`

GetNumDeleted returns the NumDeleted field if non-nil, zero value otherwise.

### GetNumDeletedOk

`func (o *MqStreamState) GetNumDeletedOk() (*int32, bool)`

GetNumDeletedOk returns a tuple with the NumDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumDeleted

`func (o *MqStreamState) SetNumDeleted(v int32)`

SetNumDeleted sets NumDeleted field to given value.

### HasNumDeleted

`func (o *MqStreamState) HasNumDeleted() bool`

HasNumDeleted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


