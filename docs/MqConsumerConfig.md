# MqConsumerConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DurableName** | Pointer to **string** | Durable consumer name. If set, the consumer survives client disconnections.  | [optional] 
**FilterSubject** | Pointer to **string** | Subject filter. Only messages matching this subject are delivered. Supports wildcards.  | [optional] 
**AckPolicy** | Pointer to **string** | Acknowledgment policy. &#x60;explicit&#x60; requires per-message ack. &#x60;all&#x60; acks all messages up to the acked sequence. &#x60;none&#x60; disables acks.  | [optional] [default to "explicit"]
**DeliverPolicy** | Pointer to **string** | Where to start delivery. &#x60;all&#x60; delivers from the beginning. &#x60;last&#x60; delivers the last message. &#x60;new&#x60; delivers only new messages. &#x60;by_start_sequence&#x60; and &#x60;by_start_time&#x60; start from a specific point.  | [optional] [default to "all"]
**OptStartSeq** | Pointer to **int32** | Starting sequence number (used with deliver_policy &#x60;by_start_sequence&#x60;).  | [optional] 
**OptStartTime** | Pointer to **time.Time** | Starting timestamp (used with deliver_policy &#x60;by_start_time&#x60;).  | [optional] 
**MaxDeliver** | Pointer to **int32** | Maximum delivery attempts before the message is dropped or sent to a dead letter subject (-1 for unlimited).  | [optional] [default to -1]
**AckWait** | Pointer to **string** | Time to wait for acknowledgment before redelivery (e.g., \&quot;30s\&quot;, \&quot;5m\&quot;). Defaults to \&quot;30s\&quot;.  | [optional] [default to "30s"]
**ReplayPolicy** | Pointer to **string** | Replay policy for historical messages. &#x60;instant&#x60; delivers as fast as possible. &#x60;original&#x60; preserves original timing gaps.  | [optional] [default to "instant"]
**MaxAckPending** | Pointer to **int32** | Maximum number of unacknowledged messages before delivery pauses.  | [optional] [default to 1000]
**Description** | Pointer to **string** | Optional human-readable description. | [optional] 

## Methods

### NewMqConsumerConfig

`func NewMqConsumerConfig() *MqConsumerConfig`

NewMqConsumerConfig instantiates a new MqConsumerConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMqConsumerConfigWithDefaults

`func NewMqConsumerConfigWithDefaults() *MqConsumerConfig`

NewMqConsumerConfigWithDefaults instantiates a new MqConsumerConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDurableName

`func (o *MqConsumerConfig) GetDurableName() string`

GetDurableName returns the DurableName field if non-nil, zero value otherwise.

### GetDurableNameOk

`func (o *MqConsumerConfig) GetDurableNameOk() (*string, bool)`

GetDurableNameOk returns a tuple with the DurableName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurableName

`func (o *MqConsumerConfig) SetDurableName(v string)`

SetDurableName sets DurableName field to given value.

### HasDurableName

`func (o *MqConsumerConfig) HasDurableName() bool`

HasDurableName returns a boolean if a field has been set.

### GetFilterSubject

`func (o *MqConsumerConfig) GetFilterSubject() string`

GetFilterSubject returns the FilterSubject field if non-nil, zero value otherwise.

### GetFilterSubjectOk

`func (o *MqConsumerConfig) GetFilterSubjectOk() (*string, bool)`

GetFilterSubjectOk returns a tuple with the FilterSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterSubject

`func (o *MqConsumerConfig) SetFilterSubject(v string)`

SetFilterSubject sets FilterSubject field to given value.

### HasFilterSubject

`func (o *MqConsumerConfig) HasFilterSubject() bool`

HasFilterSubject returns a boolean if a field has been set.

### GetAckPolicy

`func (o *MqConsumerConfig) GetAckPolicy() string`

GetAckPolicy returns the AckPolicy field if non-nil, zero value otherwise.

### GetAckPolicyOk

`func (o *MqConsumerConfig) GetAckPolicyOk() (*string, bool)`

GetAckPolicyOk returns a tuple with the AckPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAckPolicy

`func (o *MqConsumerConfig) SetAckPolicy(v string)`

SetAckPolicy sets AckPolicy field to given value.

### HasAckPolicy

`func (o *MqConsumerConfig) HasAckPolicy() bool`

HasAckPolicy returns a boolean if a field has been set.

### GetDeliverPolicy

`func (o *MqConsumerConfig) GetDeliverPolicy() string`

GetDeliverPolicy returns the DeliverPolicy field if non-nil, zero value otherwise.

### GetDeliverPolicyOk

`func (o *MqConsumerConfig) GetDeliverPolicyOk() (*string, bool)`

GetDeliverPolicyOk returns a tuple with the DeliverPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliverPolicy

`func (o *MqConsumerConfig) SetDeliverPolicy(v string)`

SetDeliverPolicy sets DeliverPolicy field to given value.

### HasDeliverPolicy

`func (o *MqConsumerConfig) HasDeliverPolicy() bool`

HasDeliverPolicy returns a boolean if a field has been set.

### GetOptStartSeq

`func (o *MqConsumerConfig) GetOptStartSeq() int32`

GetOptStartSeq returns the OptStartSeq field if non-nil, zero value otherwise.

### GetOptStartSeqOk

`func (o *MqConsumerConfig) GetOptStartSeqOk() (*int32, bool)`

GetOptStartSeqOk returns a tuple with the OptStartSeq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptStartSeq

`func (o *MqConsumerConfig) SetOptStartSeq(v int32)`

SetOptStartSeq sets OptStartSeq field to given value.

### HasOptStartSeq

`func (o *MqConsumerConfig) HasOptStartSeq() bool`

HasOptStartSeq returns a boolean if a field has been set.

### GetOptStartTime

`func (o *MqConsumerConfig) GetOptStartTime() time.Time`

GetOptStartTime returns the OptStartTime field if non-nil, zero value otherwise.

### GetOptStartTimeOk

`func (o *MqConsumerConfig) GetOptStartTimeOk() (*time.Time, bool)`

GetOptStartTimeOk returns a tuple with the OptStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptStartTime

`func (o *MqConsumerConfig) SetOptStartTime(v time.Time)`

SetOptStartTime sets OptStartTime field to given value.

### HasOptStartTime

`func (o *MqConsumerConfig) HasOptStartTime() bool`

HasOptStartTime returns a boolean if a field has been set.

### GetMaxDeliver

`func (o *MqConsumerConfig) GetMaxDeliver() int32`

GetMaxDeliver returns the MaxDeliver field if non-nil, zero value otherwise.

### GetMaxDeliverOk

`func (o *MqConsumerConfig) GetMaxDeliverOk() (*int32, bool)`

GetMaxDeliverOk returns a tuple with the MaxDeliver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDeliver

`func (o *MqConsumerConfig) SetMaxDeliver(v int32)`

SetMaxDeliver sets MaxDeliver field to given value.

### HasMaxDeliver

`func (o *MqConsumerConfig) HasMaxDeliver() bool`

HasMaxDeliver returns a boolean if a field has been set.

### GetAckWait

`func (o *MqConsumerConfig) GetAckWait() string`

GetAckWait returns the AckWait field if non-nil, zero value otherwise.

### GetAckWaitOk

`func (o *MqConsumerConfig) GetAckWaitOk() (*string, bool)`

GetAckWaitOk returns a tuple with the AckWait field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAckWait

`func (o *MqConsumerConfig) SetAckWait(v string)`

SetAckWait sets AckWait field to given value.

### HasAckWait

`func (o *MqConsumerConfig) HasAckWait() bool`

HasAckWait returns a boolean if a field has been set.

### GetReplayPolicy

`func (o *MqConsumerConfig) GetReplayPolicy() string`

GetReplayPolicy returns the ReplayPolicy field if non-nil, zero value otherwise.

### GetReplayPolicyOk

`func (o *MqConsumerConfig) GetReplayPolicyOk() (*string, bool)`

GetReplayPolicyOk returns a tuple with the ReplayPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplayPolicy

`func (o *MqConsumerConfig) SetReplayPolicy(v string)`

SetReplayPolicy sets ReplayPolicy field to given value.

### HasReplayPolicy

`func (o *MqConsumerConfig) HasReplayPolicy() bool`

HasReplayPolicy returns a boolean if a field has been set.

### GetMaxAckPending

`func (o *MqConsumerConfig) GetMaxAckPending() int32`

GetMaxAckPending returns the MaxAckPending field if non-nil, zero value otherwise.

### GetMaxAckPendingOk

`func (o *MqConsumerConfig) GetMaxAckPendingOk() (*int32, bool)`

GetMaxAckPendingOk returns a tuple with the MaxAckPending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAckPending

`func (o *MqConsumerConfig) SetMaxAckPending(v int32)`

SetMaxAckPending sets MaxAckPending field to given value.

### HasMaxAckPending

`func (o *MqConsumerConfig) HasMaxAckPending() bool`

HasMaxAckPending returns a boolean if a field has been set.

### GetDescription

`func (o *MqConsumerConfig) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MqConsumerConfig) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MqConsumerConfig) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MqConsumerConfig) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


