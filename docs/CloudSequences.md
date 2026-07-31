# CloudSequences

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConsumerSeq** | Pointer to **int32** | Consumer is the consumer&#39;s own sequence. | [optional] 
**StreamSeq** | Pointer to **int32** | Stream is the corresponding stream sequence. | [optional] 

## Methods

### NewCloudSequences

`func NewCloudSequences() *CloudSequences`

NewCloudSequences instantiates a new CloudSequences object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudSequencesWithDefaults

`func NewCloudSequencesWithDefaults() *CloudSequences`

NewCloudSequencesWithDefaults instantiates a new CloudSequences object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConsumerSeq

`func (o *CloudSequences) GetConsumerSeq() int32`

GetConsumerSeq returns the ConsumerSeq field if non-nil, zero value otherwise.

### GetConsumerSeqOk

`func (o *CloudSequences) GetConsumerSeqOk() (*int32, bool)`

GetConsumerSeqOk returns a tuple with the ConsumerSeq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerSeq

`func (o *CloudSequences) SetConsumerSeq(v int32)`

SetConsumerSeq sets ConsumerSeq field to given value.

### HasConsumerSeq

`func (o *CloudSequences) HasConsumerSeq() bool`

HasConsumerSeq returns a boolean if a field has been set.

### GetStreamSeq

`func (o *CloudSequences) GetStreamSeq() int32`

GetStreamSeq returns the StreamSeq field if non-nil, zero value otherwise.

### GetStreamSeqOk

`func (o *CloudSequences) GetStreamSeqOk() (*int32, bool)`

GetStreamSeqOk returns a tuple with the StreamSeq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamSeq

`func (o *CloudSequences) SetStreamSeq(v int32)`

SetStreamSeq sets StreamSeq field to given value.

### HasStreamSeq

`func (o *CloudSequences) HasStreamSeq() bool`

HasStreamSeq returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


