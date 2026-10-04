# MqConsumer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AckFloor** | Pointer to [**MqSequences**](MqSequences.md) | AckFloor is the highest contiguously acknowledged sequence pair. | [optional] 
**Config** | Pointer to [**MqDurable**](MqDurable.md) | Config is the consumer&#39;s configuration. | [optional] 
**Created** | Pointer to **time.Time** | Created is when the consumer was created. | [optional] 
**Delivered** | Pointer to [**MqSequences**](MqSequences.md) | Delivered is the highest delivered sequence pair. | [optional] 
**Name** | Pointer to **string** | Name is the consumer name. | [optional] 
**NumAckPending** | Pointer to **int64** | AckPending is the number of delivered, not yet acknowledged messages. | [optional] 
**NumPending** | Pointer to **int32** | Pending is the number of messages yet to be delivered. | [optional] 
**NumRedelivered** | Pointer to **int64** | Redelivered is the number of messages currently being redelivered. | [optional] 
**NumWaiting** | Pointer to **int64** | Waiting is the number of pull requests waiting for messages. | [optional] 
**StreamName** | Pointer to **string** | Stream is the stream this consumer reads. | [optional] 

## Methods

### NewMqConsumer

`func NewMqConsumer() *MqConsumer`

NewMqConsumer instantiates a new MqConsumer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMqConsumerWithDefaults

`func NewMqConsumerWithDefaults() *MqConsumer`

NewMqConsumerWithDefaults instantiates a new MqConsumer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAckFloor

`func (o *MqConsumer) GetAckFloor() MqSequences`

GetAckFloor returns the AckFloor field if non-nil, zero value otherwise.

### GetAckFloorOk

`func (o *MqConsumer) GetAckFloorOk() (*MqSequences, bool)`

GetAckFloorOk returns a tuple with the AckFloor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAckFloor

`func (o *MqConsumer) SetAckFloor(v MqSequences)`

SetAckFloor sets AckFloor field to given value.

### HasAckFloor

`func (o *MqConsumer) HasAckFloor() bool`

HasAckFloor returns a boolean if a field has been set.

### GetConfig

`func (o *MqConsumer) GetConfig() MqDurable`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *MqConsumer) GetConfigOk() (*MqDurable, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *MqConsumer) SetConfig(v MqDurable)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *MqConsumer) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetCreated

`func (o *MqConsumer) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *MqConsumer) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *MqConsumer) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *MqConsumer) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetDelivered

`func (o *MqConsumer) GetDelivered() MqSequences`

GetDelivered returns the Delivered field if non-nil, zero value otherwise.

### GetDeliveredOk

`func (o *MqConsumer) GetDeliveredOk() (*MqSequences, bool)`

GetDeliveredOk returns a tuple with the Delivered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelivered

`func (o *MqConsumer) SetDelivered(v MqSequences)`

SetDelivered sets Delivered field to given value.

### HasDelivered

`func (o *MqConsumer) HasDelivered() bool`

HasDelivered returns a boolean if a field has been set.

### GetName

`func (o *MqConsumer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MqConsumer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MqConsumer) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MqConsumer) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNumAckPending

`func (o *MqConsumer) GetNumAckPending() int64`

GetNumAckPending returns the NumAckPending field if non-nil, zero value otherwise.

### GetNumAckPendingOk

`func (o *MqConsumer) GetNumAckPendingOk() (*int64, bool)`

GetNumAckPendingOk returns a tuple with the NumAckPending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumAckPending

`func (o *MqConsumer) SetNumAckPending(v int64)`

SetNumAckPending sets NumAckPending field to given value.

### HasNumAckPending

`func (o *MqConsumer) HasNumAckPending() bool`

HasNumAckPending returns a boolean if a field has been set.

### GetNumPending

`func (o *MqConsumer) GetNumPending() int32`

GetNumPending returns the NumPending field if non-nil, zero value otherwise.

### GetNumPendingOk

`func (o *MqConsumer) GetNumPendingOk() (*int32, bool)`

GetNumPendingOk returns a tuple with the NumPending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumPending

`func (o *MqConsumer) SetNumPending(v int32)`

SetNumPending sets NumPending field to given value.

### HasNumPending

`func (o *MqConsumer) HasNumPending() bool`

HasNumPending returns a boolean if a field has been set.

### GetNumRedelivered

`func (o *MqConsumer) GetNumRedelivered() int64`

GetNumRedelivered returns the NumRedelivered field if non-nil, zero value otherwise.

### GetNumRedeliveredOk

`func (o *MqConsumer) GetNumRedeliveredOk() (*int64, bool)`

GetNumRedeliveredOk returns a tuple with the NumRedelivered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumRedelivered

`func (o *MqConsumer) SetNumRedelivered(v int64)`

SetNumRedelivered sets NumRedelivered field to given value.

### HasNumRedelivered

`func (o *MqConsumer) HasNumRedelivered() bool`

HasNumRedelivered returns a boolean if a field has been set.

### GetNumWaiting

`func (o *MqConsumer) GetNumWaiting() int64`

GetNumWaiting returns the NumWaiting field if non-nil, zero value otherwise.

### GetNumWaitingOk

`func (o *MqConsumer) GetNumWaitingOk() (*int64, bool)`

GetNumWaitingOk returns a tuple with the NumWaiting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumWaiting

`func (o *MqConsumer) SetNumWaiting(v int64)`

SetNumWaiting sets NumWaiting field to given value.

### HasNumWaiting

`func (o *MqConsumer) HasNumWaiting() bool`

HasNumWaiting returns a boolean if a field has been set.

### GetStreamName

`func (o *MqConsumer) GetStreamName() string`

GetStreamName returns the StreamName field if non-nil, zero value otherwise.

### GetStreamNameOk

`func (o *MqConsumer) GetStreamNameOk() (*string, bool)`

GetStreamNameOk returns a tuple with the StreamName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamName

`func (o *MqConsumer) SetStreamName(v string)`

SetStreamName sets StreamName field to given value.

### HasStreamName

`func (o *MqConsumer) HasStreamName() bool`

HasStreamName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


