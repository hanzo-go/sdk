# PubsubStreamInfoState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Messages** | Pointer to **int32** |  | [optional] 
**Bytes** | Pointer to **int32** |  | [optional] 
**FirstSeq** | Pointer to **int32** |  | [optional] 
**LastSeq** | Pointer to **int32** |  | [optional] 
**ConsumerCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewPubsubStreamInfoState

`func NewPubsubStreamInfoState() *PubsubStreamInfoState`

NewPubsubStreamInfoState instantiates a new PubsubStreamInfoState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPubsubStreamInfoStateWithDefaults

`func NewPubsubStreamInfoStateWithDefaults() *PubsubStreamInfoState`

NewPubsubStreamInfoStateWithDefaults instantiates a new PubsubStreamInfoState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessages

`func (o *PubsubStreamInfoState) GetMessages() int32`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *PubsubStreamInfoState) GetMessagesOk() (*int32, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *PubsubStreamInfoState) SetMessages(v int32)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *PubsubStreamInfoState) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### GetBytes

`func (o *PubsubStreamInfoState) GetBytes() int32`

GetBytes returns the Bytes field if non-nil, zero value otherwise.

### GetBytesOk

`func (o *PubsubStreamInfoState) GetBytesOk() (*int32, bool)`

GetBytesOk returns a tuple with the Bytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytes

`func (o *PubsubStreamInfoState) SetBytes(v int32)`

SetBytes sets Bytes field to given value.

### HasBytes

`func (o *PubsubStreamInfoState) HasBytes() bool`

HasBytes returns a boolean if a field has been set.

### GetFirstSeq

`func (o *PubsubStreamInfoState) GetFirstSeq() int32`

GetFirstSeq returns the FirstSeq field if non-nil, zero value otherwise.

### GetFirstSeqOk

`func (o *PubsubStreamInfoState) GetFirstSeqOk() (*int32, bool)`

GetFirstSeqOk returns a tuple with the FirstSeq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstSeq

`func (o *PubsubStreamInfoState) SetFirstSeq(v int32)`

SetFirstSeq sets FirstSeq field to given value.

### HasFirstSeq

`func (o *PubsubStreamInfoState) HasFirstSeq() bool`

HasFirstSeq returns a boolean if a field has been set.

### GetLastSeq

`func (o *PubsubStreamInfoState) GetLastSeq() int32`

GetLastSeq returns the LastSeq field if non-nil, zero value otherwise.

### GetLastSeqOk

`func (o *PubsubStreamInfoState) GetLastSeqOk() (*int32, bool)`

GetLastSeqOk returns a tuple with the LastSeq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeq

`func (o *PubsubStreamInfoState) SetLastSeq(v int32)`

SetLastSeq sets LastSeq field to given value.

### HasLastSeq

`func (o *PubsubStreamInfoState) HasLastSeq() bool`

HasLastSeq returns a boolean if a field has been set.

### GetConsumerCount

`func (o *PubsubStreamInfoState) GetConsumerCount() int32`

GetConsumerCount returns the ConsumerCount field if non-nil, zero value otherwise.

### GetConsumerCountOk

`func (o *PubsubStreamInfoState) GetConsumerCountOk() (*int32, bool)`

GetConsumerCountOk returns a tuple with the ConsumerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerCount

`func (o *PubsubStreamInfoState) SetConsumerCount(v int32)`

SetConsumerCount sets ConsumerCount field to given value.

### HasConsumerCount

`func (o *PubsubStreamInfoState) HasConsumerCount() bool`

HasConsumerCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


