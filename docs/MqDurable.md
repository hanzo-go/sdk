# MqDurable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AckPolicy** | Pointer to **string** | Ack is the acknowledgment policy: explicit (default), all, or none. | [optional] 
**AckWait** | Pointer to **string** | AckWait is how long the broker waits for an ack before redelivering, e.g. \&quot;30s\&quot; (default). | [optional] 
**DeliverPolicy** | Pointer to **string** | Deliver is where delivery starts: all (default), last, new, by_start_sequence, by_start_time, or last_per_subject. | [optional] 
**Description** | Pointer to **string** | Description says what this consumer is for. | [optional] 
**DurableName** | Pointer to **string** | Name is the durable consumer name (alphanumeric, hyphens, underscores). | [optional] 
**FilterSubject** | Pointer to **string** | Filter delivers only messages on this org-relative subject (wildcards supported). | [optional] 
**MaxAckPending** | Pointer to **int64** | MaxAckPending caps unacknowledged messages in flight (default 1000). | [optional] 
**MaxDeliver** | Pointer to **int64** | MaxDeliver caps delivery attempts per message; -1 (default) is unlimited. | [optional] 
**OptStartSeq** | Pointer to **int32** | StartSeq is the starting sequence for deliver_policy by_start_sequence. | [optional] 
**OptStartTime** | Pointer to **time.Time** | StartTime is the starting instant for deliver_policy by_start_time. | [optional] 
**ReplayPolicy** | Pointer to **string** | Replay is the replay pacing: instant (default) or original. | [optional] 

## Methods

### NewMqDurable

`func NewMqDurable() *MqDurable`

NewMqDurable instantiates a new MqDurable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMqDurableWithDefaults

`func NewMqDurableWithDefaults() *MqDurable`

NewMqDurableWithDefaults instantiates a new MqDurable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAckPolicy

`func (o *MqDurable) GetAckPolicy() string`

GetAckPolicy returns the AckPolicy field if non-nil, zero value otherwise.

### GetAckPolicyOk

`func (o *MqDurable) GetAckPolicyOk() (*string, bool)`

GetAckPolicyOk returns a tuple with the AckPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAckPolicy

`func (o *MqDurable) SetAckPolicy(v string)`

SetAckPolicy sets AckPolicy field to given value.

### HasAckPolicy

`func (o *MqDurable) HasAckPolicy() bool`

HasAckPolicy returns a boolean if a field has been set.

### GetAckWait

`func (o *MqDurable) GetAckWait() string`

GetAckWait returns the AckWait field if non-nil, zero value otherwise.

### GetAckWaitOk

`func (o *MqDurable) GetAckWaitOk() (*string, bool)`

GetAckWaitOk returns a tuple with the AckWait field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAckWait

`func (o *MqDurable) SetAckWait(v string)`

SetAckWait sets AckWait field to given value.

### HasAckWait

`func (o *MqDurable) HasAckWait() bool`

HasAckWait returns a boolean if a field has been set.

### GetDeliverPolicy

`func (o *MqDurable) GetDeliverPolicy() string`

GetDeliverPolicy returns the DeliverPolicy field if non-nil, zero value otherwise.

### GetDeliverPolicyOk

`func (o *MqDurable) GetDeliverPolicyOk() (*string, bool)`

GetDeliverPolicyOk returns a tuple with the DeliverPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliverPolicy

`func (o *MqDurable) SetDeliverPolicy(v string)`

SetDeliverPolicy sets DeliverPolicy field to given value.

### HasDeliverPolicy

`func (o *MqDurable) HasDeliverPolicy() bool`

HasDeliverPolicy returns a boolean if a field has been set.

### GetDescription

`func (o *MqDurable) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MqDurable) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MqDurable) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MqDurable) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDurableName

`func (o *MqDurable) GetDurableName() string`

GetDurableName returns the DurableName field if non-nil, zero value otherwise.

### GetDurableNameOk

`func (o *MqDurable) GetDurableNameOk() (*string, bool)`

GetDurableNameOk returns a tuple with the DurableName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurableName

`func (o *MqDurable) SetDurableName(v string)`

SetDurableName sets DurableName field to given value.

### HasDurableName

`func (o *MqDurable) HasDurableName() bool`

HasDurableName returns a boolean if a field has been set.

### GetFilterSubject

`func (o *MqDurable) GetFilterSubject() string`

GetFilterSubject returns the FilterSubject field if non-nil, zero value otherwise.

### GetFilterSubjectOk

`func (o *MqDurable) GetFilterSubjectOk() (*string, bool)`

GetFilterSubjectOk returns a tuple with the FilterSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterSubject

`func (o *MqDurable) SetFilterSubject(v string)`

SetFilterSubject sets FilterSubject field to given value.

### HasFilterSubject

`func (o *MqDurable) HasFilterSubject() bool`

HasFilterSubject returns a boolean if a field has been set.

### GetMaxAckPending

`func (o *MqDurable) GetMaxAckPending() int64`

GetMaxAckPending returns the MaxAckPending field if non-nil, zero value otherwise.

### GetMaxAckPendingOk

`func (o *MqDurable) GetMaxAckPendingOk() (*int64, bool)`

GetMaxAckPendingOk returns a tuple with the MaxAckPending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAckPending

`func (o *MqDurable) SetMaxAckPending(v int64)`

SetMaxAckPending sets MaxAckPending field to given value.

### HasMaxAckPending

`func (o *MqDurable) HasMaxAckPending() bool`

HasMaxAckPending returns a boolean if a field has been set.

### GetMaxDeliver

`func (o *MqDurable) GetMaxDeliver() int64`

GetMaxDeliver returns the MaxDeliver field if non-nil, zero value otherwise.

### GetMaxDeliverOk

`func (o *MqDurable) GetMaxDeliverOk() (*int64, bool)`

GetMaxDeliverOk returns a tuple with the MaxDeliver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDeliver

`func (o *MqDurable) SetMaxDeliver(v int64)`

SetMaxDeliver sets MaxDeliver field to given value.

### HasMaxDeliver

`func (o *MqDurable) HasMaxDeliver() bool`

HasMaxDeliver returns a boolean if a field has been set.

### GetOptStartSeq

`func (o *MqDurable) GetOptStartSeq() int32`

GetOptStartSeq returns the OptStartSeq field if non-nil, zero value otherwise.

### GetOptStartSeqOk

`func (o *MqDurable) GetOptStartSeqOk() (*int32, bool)`

GetOptStartSeqOk returns a tuple with the OptStartSeq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptStartSeq

`func (o *MqDurable) SetOptStartSeq(v int32)`

SetOptStartSeq sets OptStartSeq field to given value.

### HasOptStartSeq

`func (o *MqDurable) HasOptStartSeq() bool`

HasOptStartSeq returns a boolean if a field has been set.

### GetOptStartTime

`func (o *MqDurable) GetOptStartTime() time.Time`

GetOptStartTime returns the OptStartTime field if non-nil, zero value otherwise.

### GetOptStartTimeOk

`func (o *MqDurable) GetOptStartTimeOk() (*time.Time, bool)`

GetOptStartTimeOk returns a tuple with the OptStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptStartTime

`func (o *MqDurable) SetOptStartTime(v time.Time)`

SetOptStartTime sets OptStartTime field to given value.

### HasOptStartTime

`func (o *MqDurable) HasOptStartTime() bool`

HasOptStartTime returns a boolean if a field has been set.

### GetReplayPolicy

`func (o *MqDurable) GetReplayPolicy() string`

GetReplayPolicy returns the ReplayPolicy field if non-nil, zero value otherwise.

### GetReplayPolicyOk

`func (o *MqDurable) GetReplayPolicyOk() (*string, bool)`

GetReplayPolicyOk returns a tuple with the ReplayPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplayPolicy

`func (o *MqDurable) SetReplayPolicy(v string)`

SetReplayPolicy sets ReplayPolicy field to given value.

### HasReplayPolicy

`func (o *MqDurable) HasReplayPolicy() bool`

HasReplayPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


