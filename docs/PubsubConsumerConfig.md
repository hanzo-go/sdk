# PubsubConsumerConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DurableName** | **string** | Durable consumer name | 
**DeliverPolicy** | Pointer to **string** |  | [optional] [default to "all"]
**AckPolicy** | Pointer to **string** |  | [optional] [default to "explicit"]
**FilterSubject** | Pointer to **string** | Subject filter for this consumer | [optional] 
**MaxDeliver** | Pointer to **int32** | Maximum delivery attempts per message | [optional] [default to -1]
**AckWait** | Pointer to **int64** | Ack wait timeout in nanoseconds | [optional] [default to 30000000000]

## Methods

### NewPubsubConsumerConfig

`func NewPubsubConsumerConfig(durableName string, ) *PubsubConsumerConfig`

NewPubsubConsumerConfig instantiates a new PubsubConsumerConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPubsubConsumerConfigWithDefaults

`func NewPubsubConsumerConfigWithDefaults() *PubsubConsumerConfig`

NewPubsubConsumerConfigWithDefaults instantiates a new PubsubConsumerConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDurableName

`func (o *PubsubConsumerConfig) GetDurableName() string`

GetDurableName returns the DurableName field if non-nil, zero value otherwise.

### GetDurableNameOk

`func (o *PubsubConsumerConfig) GetDurableNameOk() (*string, bool)`

GetDurableNameOk returns a tuple with the DurableName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurableName

`func (o *PubsubConsumerConfig) SetDurableName(v string)`

SetDurableName sets DurableName field to given value.


### GetDeliverPolicy

`func (o *PubsubConsumerConfig) GetDeliverPolicy() string`

GetDeliverPolicy returns the DeliverPolicy field if non-nil, zero value otherwise.

### GetDeliverPolicyOk

`func (o *PubsubConsumerConfig) GetDeliverPolicyOk() (*string, bool)`

GetDeliverPolicyOk returns a tuple with the DeliverPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliverPolicy

`func (o *PubsubConsumerConfig) SetDeliverPolicy(v string)`

SetDeliverPolicy sets DeliverPolicy field to given value.

### HasDeliverPolicy

`func (o *PubsubConsumerConfig) HasDeliverPolicy() bool`

HasDeliverPolicy returns a boolean if a field has been set.

### GetAckPolicy

`func (o *PubsubConsumerConfig) GetAckPolicy() string`

GetAckPolicy returns the AckPolicy field if non-nil, zero value otherwise.

### GetAckPolicyOk

`func (o *PubsubConsumerConfig) GetAckPolicyOk() (*string, bool)`

GetAckPolicyOk returns a tuple with the AckPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAckPolicy

`func (o *PubsubConsumerConfig) SetAckPolicy(v string)`

SetAckPolicy sets AckPolicy field to given value.

### HasAckPolicy

`func (o *PubsubConsumerConfig) HasAckPolicy() bool`

HasAckPolicy returns a boolean if a field has been set.

### GetFilterSubject

`func (o *PubsubConsumerConfig) GetFilterSubject() string`

GetFilterSubject returns the FilterSubject field if non-nil, zero value otherwise.

### GetFilterSubjectOk

`func (o *PubsubConsumerConfig) GetFilterSubjectOk() (*string, bool)`

GetFilterSubjectOk returns a tuple with the FilterSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterSubject

`func (o *PubsubConsumerConfig) SetFilterSubject(v string)`

SetFilterSubject sets FilterSubject field to given value.

### HasFilterSubject

`func (o *PubsubConsumerConfig) HasFilterSubject() bool`

HasFilterSubject returns a boolean if a field has been set.

### GetMaxDeliver

`func (o *PubsubConsumerConfig) GetMaxDeliver() int32`

GetMaxDeliver returns the MaxDeliver field if non-nil, zero value otherwise.

### GetMaxDeliverOk

`func (o *PubsubConsumerConfig) GetMaxDeliverOk() (*int32, bool)`

GetMaxDeliverOk returns a tuple with the MaxDeliver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDeliver

`func (o *PubsubConsumerConfig) SetMaxDeliver(v int32)`

SetMaxDeliver sets MaxDeliver field to given value.

### HasMaxDeliver

`func (o *PubsubConsumerConfig) HasMaxDeliver() bool`

HasMaxDeliver returns a boolean if a field has been set.

### GetAckWait

`func (o *PubsubConsumerConfig) GetAckWait() int64`

GetAckWait returns the AckWait field if non-nil, zero value otherwise.

### GetAckWaitOk

`func (o *PubsubConsumerConfig) GetAckWaitOk() (*int64, bool)`

GetAckWaitOk returns a tuple with the AckWait field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAckWait

`func (o *PubsubConsumerConfig) SetAckWait(v int64)`

SetAckWait sets AckWait field to given value.

### HasAckWait

`func (o *PubsubConsumerConfig) HasAckWait() bool`

HasAckWait returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


