# CloudStreamRecord

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

### NewCloudStreamRecord

`func NewCloudStreamRecord() *CloudStreamRecord`

NewCloudStreamRecord instantiates a new CloudStreamRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudStreamRecordWithDefaults

`func NewCloudStreamRecordWithDefaults() *CloudStreamRecord`

NewCloudStreamRecordWithDefaults instantiates a new CloudStreamRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBytes

`func (o *CloudStreamRecord) GetBytes() int32`

GetBytes returns the Bytes field if non-nil, zero value otherwise.

### GetBytesOk

`func (o *CloudStreamRecord) GetBytesOk() (*int32, bool)`

GetBytesOk returns a tuple with the Bytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytes

`func (o *CloudStreamRecord) SetBytes(v int32)`

SetBytes sets Bytes field to given value.

### HasBytes

`func (o *CloudStreamRecord) HasBytes() bool`

HasBytes returns a boolean if a field has been set.

### GetConsumers

`func (o *CloudStreamRecord) GetConsumers() int32`

GetConsumers returns the Consumers field if non-nil, zero value otherwise.

### GetConsumersOk

`func (o *CloudStreamRecord) GetConsumersOk() (*int32, bool)`

GetConsumersOk returns a tuple with the Consumers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumers

`func (o *CloudStreamRecord) SetConsumers(v int32)`

SetConsumers sets Consumers field to given value.

### HasConsumers

`func (o *CloudStreamRecord) HasConsumers() bool`

HasConsumers returns a boolean if a field has been set.

### GetCreated

`func (o *CloudStreamRecord) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CloudStreamRecord) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *CloudStreamRecord) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *CloudStreamRecord) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetDiscard

`func (o *CloudStreamRecord) GetDiscard() string`

GetDiscard returns the Discard field if non-nil, zero value otherwise.

### GetDiscardOk

`func (o *CloudStreamRecord) GetDiscardOk() (*string, bool)`

GetDiscardOk returns a tuple with the Discard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscard

`func (o *CloudStreamRecord) SetDiscard(v string)`

SetDiscard sets Discard field to given value.

### HasDiscard

`func (o *CloudStreamRecord) HasDiscard() bool`

HasDiscard returns a boolean if a field has been set.

### GetFirstSeq

`func (o *CloudStreamRecord) GetFirstSeq() int32`

GetFirstSeq returns the FirstSeq field if non-nil, zero value otherwise.

### GetFirstSeqOk

`func (o *CloudStreamRecord) GetFirstSeqOk() (*int32, bool)`

GetFirstSeqOk returns a tuple with the FirstSeq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstSeq

`func (o *CloudStreamRecord) SetFirstSeq(v int32)`

SetFirstSeq sets FirstSeq field to given value.

### HasFirstSeq

`func (o *CloudStreamRecord) HasFirstSeq() bool`

HasFirstSeq returns a boolean if a field has been set.

### GetLastSeq

`func (o *CloudStreamRecord) GetLastSeq() int32`

GetLastSeq returns the LastSeq field if non-nil, zero value otherwise.

### GetLastSeqOk

`func (o *CloudStreamRecord) GetLastSeqOk() (*int32, bool)`

GetLastSeqOk returns a tuple with the LastSeq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeq

`func (o *CloudStreamRecord) SetLastSeq(v int32)`

SetLastSeq sets LastSeq field to given value.

### HasLastSeq

`func (o *CloudStreamRecord) HasLastSeq() bool`

HasLastSeq returns a boolean if a field has been set.

### GetMaxAge

`func (o *CloudStreamRecord) GetMaxAge() int32`

GetMaxAge returns the MaxAge field if non-nil, zero value otherwise.

### GetMaxAgeOk

`func (o *CloudStreamRecord) GetMaxAgeOk() (*int32, bool)`

GetMaxAgeOk returns a tuple with the MaxAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAge

`func (o *CloudStreamRecord) SetMaxAge(v int32)`

SetMaxAge sets MaxAge field to given value.

### HasMaxAge

`func (o *CloudStreamRecord) HasMaxAge() bool`

HasMaxAge returns a boolean if a field has been set.

### GetMaxBytes

`func (o *CloudStreamRecord) GetMaxBytes() int32`

GetMaxBytes returns the MaxBytes field if non-nil, zero value otherwise.

### GetMaxBytesOk

`func (o *CloudStreamRecord) GetMaxBytesOk() (*int32, bool)`

GetMaxBytesOk returns a tuple with the MaxBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxBytes

`func (o *CloudStreamRecord) SetMaxBytes(v int32)`

SetMaxBytes sets MaxBytes field to given value.

### HasMaxBytes

`func (o *CloudStreamRecord) HasMaxBytes() bool`

HasMaxBytes returns a boolean if a field has been set.

### GetMaxMsgs

`func (o *CloudStreamRecord) GetMaxMsgs() int32`

GetMaxMsgs returns the MaxMsgs field if non-nil, zero value otherwise.

### GetMaxMsgsOk

`func (o *CloudStreamRecord) GetMaxMsgsOk() (*int32, bool)`

GetMaxMsgsOk returns a tuple with the MaxMsgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMsgs

`func (o *CloudStreamRecord) SetMaxMsgs(v int32)`

SetMaxMsgs sets MaxMsgs field to given value.

### HasMaxMsgs

`func (o *CloudStreamRecord) HasMaxMsgs() bool`

HasMaxMsgs returns a boolean if a field has been set.

### GetMessages

`func (o *CloudStreamRecord) GetMessages() int32`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *CloudStreamRecord) GetMessagesOk() (*int32, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *CloudStreamRecord) SetMessages(v int32)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *CloudStreamRecord) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### GetName

`func (o *CloudStreamRecord) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudStreamRecord) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudStreamRecord) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudStreamRecord) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRetention

`func (o *CloudStreamRecord) GetRetention() string`

GetRetention returns the Retention field if non-nil, zero value otherwise.

### GetRetentionOk

`func (o *CloudStreamRecord) GetRetentionOk() (*string, bool)`

GetRetentionOk returns a tuple with the Retention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetention

`func (o *CloudStreamRecord) SetRetention(v string)`

SetRetention sets Retention field to given value.

### HasRetention

`func (o *CloudStreamRecord) HasRetention() bool`

HasRetention returns a boolean if a field has been set.

### GetStorage

`func (o *CloudStreamRecord) GetStorage() string`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *CloudStreamRecord) GetStorageOk() (*string, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *CloudStreamRecord) SetStorage(v string)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *CloudStreamRecord) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### GetSubjects

`func (o *CloudStreamRecord) GetSubjects() []string`

GetSubjects returns the Subjects field if non-nil, zero value otherwise.

### GetSubjectsOk

`func (o *CloudStreamRecord) GetSubjectsOk() (*[]string, bool)`

GetSubjectsOk returns a tuple with the Subjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjects

`func (o *CloudStreamRecord) SetSubjects(v []string)`

SetSubjects sets Subjects field to given value.

### HasSubjects

`func (o *CloudStreamRecord) HasSubjects() bool`

HasSubjects returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


