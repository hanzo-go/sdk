# MqSequencePair

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConsumerSeq** | Pointer to **int32** | Consumer sequence number. | [optional] 
**StreamSeq** | Pointer to **int32** | Corresponding stream sequence number. | [optional] 

## Methods

### NewMqSequencePair

`func NewMqSequencePair() *MqSequencePair`

NewMqSequencePair instantiates a new MqSequencePair object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMqSequencePairWithDefaults

`func NewMqSequencePairWithDefaults() *MqSequencePair`

NewMqSequencePairWithDefaults instantiates a new MqSequencePair object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConsumerSeq

`func (o *MqSequencePair) GetConsumerSeq() int32`

GetConsumerSeq returns the ConsumerSeq field if non-nil, zero value otherwise.

### GetConsumerSeqOk

`func (o *MqSequencePair) GetConsumerSeqOk() (*int32, bool)`

GetConsumerSeqOk returns a tuple with the ConsumerSeq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerSeq

`func (o *MqSequencePair) SetConsumerSeq(v int32)`

SetConsumerSeq sets ConsumerSeq field to given value.

### HasConsumerSeq

`func (o *MqSequencePair) HasConsumerSeq() bool`

HasConsumerSeq returns a boolean if a field has been set.

### GetStreamSeq

`func (o *MqSequencePair) GetStreamSeq() int32`

GetStreamSeq returns the StreamSeq field if non-nil, zero value otherwise.

### GetStreamSeqOk

`func (o *MqSequencePair) GetStreamSeqOk() (*int32, bool)`

GetStreamSeqOk returns a tuple with the StreamSeq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamSeq

`func (o *MqSequencePair) SetStreamSeq(v int32)`

SetStreamSeq sets StreamSeq field to given value.

### HasStreamSeq

`func (o *MqSequencePair) HasStreamSeq() bool`

HasStreamSeq returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


