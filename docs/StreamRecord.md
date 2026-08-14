# StreamRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bytes** | Pointer to **int32** | Bytes is how many bytes the stream holds right now. | [optional] 
**Consumers** | Pointer to **int32** | Consumers is how many consumers the stream carries. | [optional] 
**Created** | Pointer to **string** | Created is when the stream was created, RFC3339. | [optional] 
**Discard** | Pointer to **string** | Discard says which end gives way at the limits: old or new. | [optional] 
**FirstSeq** | Pointer to **int32** | FirstSeq is the sequence of the oldest retained message. | [optional] 
**LastSeq** | Pointer to **int32** | LastSeq is the sequence of the newest message. | [optional] 
**MaxAge** | Pointer to **int32** | MaxAge is the age cap in seconds; 0 means unlimited. | [optional] 
**MaxBytes** | Pointer to **int32** | MaxBytes is the retained-byte cap; -1 means unlimited. | [optional] 
**MaxMsgs** | Pointer to **int32** | MaxMsgs is the retained-message cap; -1 means unlimited. | [optional] 
**Messages** | Pointer to **int32** | Messages is how many messages the stream holds right now. | [optional] 
**Name** | Pointer to **string** | Name is the stream&#39;s name within the org. | [optional] 
**Retention** | Pointer to **string** | Retention is the discipline: limits, interest or workqueue. | [optional] 
**Storage** | Pointer to **string** | Storage is the backend: file or memory. | [optional] 
**Subjects** | Pointer to **[]string** | Subjects are the subjects it captures, in the org&#39;s namespace. | [optional] 

## Methods

### NewStreamRecord

`func NewStreamRecord() *StreamRecord`

NewStreamRecord instantiates a new StreamRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamRecordWithDefaults

`func NewStreamRecordWithDefaults() *StreamRecord`

NewStreamRecordWithDefaults instantiates a new StreamRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBytes

`func (o *StreamRecord) GetBytes() int32`

GetBytes returns the Bytes field if non-nil, zero value otherwise.

### GetBytesOk

`func (o *StreamRecord) GetBytesOk() (*int32, bool)`

GetBytesOk returns a tuple with the Bytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytes

`func (o *StreamRecord) SetBytes(v int32)`

SetBytes sets Bytes field to given value.

### HasBytes

`func (o *StreamRecord) HasBytes() bool`

HasBytes returns a boolean if a field has been set.

### GetConsumers

`func (o *StreamRecord) GetConsumers() int32`

GetConsumers returns the Consumers field if non-nil, zero value otherwise.

### GetConsumersOk

`func (o *StreamRecord) GetConsumersOk() (*int32, bool)`

GetConsumersOk returns a tuple with the Consumers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumers

`func (o *StreamRecord) SetConsumers(v int32)`

SetConsumers sets Consumers field to given value.

### HasConsumers

`func (o *StreamRecord) HasConsumers() bool`

HasConsumers returns a boolean if a field has been set.

### GetCreated

`func (o *StreamRecord) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *StreamRecord) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *StreamRecord) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *StreamRecord) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetDiscard

`func (o *StreamRecord) GetDiscard() string`

GetDiscard returns the Discard field if non-nil, zero value otherwise.

### GetDiscardOk

`func (o *StreamRecord) GetDiscardOk() (*string, bool)`

GetDiscardOk returns a tuple with the Discard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscard

`func (o *StreamRecord) SetDiscard(v string)`

SetDiscard sets Discard field to given value.

### HasDiscard

`func (o *StreamRecord) HasDiscard() bool`

HasDiscard returns a boolean if a field has been set.

### GetFirstSeq

`func (o *StreamRecord) GetFirstSeq() int32`

GetFirstSeq returns the FirstSeq field if non-nil, zero value otherwise.

### GetFirstSeqOk

`func (o *StreamRecord) GetFirstSeqOk() (*int32, bool)`

GetFirstSeqOk returns a tuple with the FirstSeq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstSeq

`func (o *StreamRecord) SetFirstSeq(v int32)`

SetFirstSeq sets FirstSeq field to given value.

### HasFirstSeq

`func (o *StreamRecord) HasFirstSeq() bool`

HasFirstSeq returns a boolean if a field has been set.

### GetLastSeq

`func (o *StreamRecord) GetLastSeq() int32`

GetLastSeq returns the LastSeq field if non-nil, zero value otherwise.

### GetLastSeqOk

`func (o *StreamRecord) GetLastSeqOk() (*int32, bool)`

GetLastSeqOk returns a tuple with the LastSeq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeq

`func (o *StreamRecord) SetLastSeq(v int32)`

SetLastSeq sets LastSeq field to given value.

### HasLastSeq

`func (o *StreamRecord) HasLastSeq() bool`

HasLastSeq returns a boolean if a field has been set.

### GetMaxAge

`func (o *StreamRecord) GetMaxAge() int32`

GetMaxAge returns the MaxAge field if non-nil, zero value otherwise.

### GetMaxAgeOk

`func (o *StreamRecord) GetMaxAgeOk() (*int32, bool)`

GetMaxAgeOk returns a tuple with the MaxAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAge

`func (o *StreamRecord) SetMaxAge(v int32)`

SetMaxAge sets MaxAge field to given value.

### HasMaxAge

`func (o *StreamRecord) HasMaxAge() bool`

HasMaxAge returns a boolean if a field has been set.

### GetMaxBytes

`func (o *StreamRecord) GetMaxBytes() int32`

GetMaxBytes returns the MaxBytes field if non-nil, zero value otherwise.

### GetMaxBytesOk

`func (o *StreamRecord) GetMaxBytesOk() (*int32, bool)`

GetMaxBytesOk returns a tuple with the MaxBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxBytes

`func (o *StreamRecord) SetMaxBytes(v int32)`

SetMaxBytes sets MaxBytes field to given value.

### HasMaxBytes

`func (o *StreamRecord) HasMaxBytes() bool`

HasMaxBytes returns a boolean if a field has been set.

### GetMaxMsgs

`func (o *StreamRecord) GetMaxMsgs() int32`

GetMaxMsgs returns the MaxMsgs field if non-nil, zero value otherwise.

### GetMaxMsgsOk

`func (o *StreamRecord) GetMaxMsgsOk() (*int32, bool)`

GetMaxMsgsOk returns a tuple with the MaxMsgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMsgs

`func (o *StreamRecord) SetMaxMsgs(v int32)`

SetMaxMsgs sets MaxMsgs field to given value.

### HasMaxMsgs

`func (o *StreamRecord) HasMaxMsgs() bool`

HasMaxMsgs returns a boolean if a field has been set.

### GetMessages

`func (o *StreamRecord) GetMessages() int32`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *StreamRecord) GetMessagesOk() (*int32, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *StreamRecord) SetMessages(v int32)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *StreamRecord) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### GetName

`func (o *StreamRecord) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StreamRecord) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StreamRecord) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StreamRecord) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRetention

`func (o *StreamRecord) GetRetention() string`

GetRetention returns the Retention field if non-nil, zero value otherwise.

### GetRetentionOk

`func (o *StreamRecord) GetRetentionOk() (*string, bool)`

GetRetentionOk returns a tuple with the Retention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetention

`func (o *StreamRecord) SetRetention(v string)`

SetRetention sets Retention field to given value.

### HasRetention

`func (o *StreamRecord) HasRetention() bool`

HasRetention returns a boolean if a field has been set.

### GetStorage

`func (o *StreamRecord) GetStorage() string`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *StreamRecord) GetStorageOk() (*string, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *StreamRecord) SetStorage(v string)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *StreamRecord) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### GetSubjects

`func (o *StreamRecord) GetSubjects() []string`

GetSubjects returns the Subjects field if non-nil, zero value otherwise.

### GetSubjectsOk

`func (o *StreamRecord) GetSubjectsOk() (*[]string, bool)`

GetSubjectsOk returns a tuple with the Subjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjects

`func (o *StreamRecord) SetSubjects(v []string)`

SetSubjects sets Subjects field to given value.

### HasSubjects

`func (o *StreamRecord) HasSubjects() bool`

HasSubjects returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


