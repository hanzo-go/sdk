# Consumer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AckFloor** | Pointer to [**Sequences**](Sequences.md) | AckFloor is the highest contiguously acknowledged sequence pair. | [optional] 
**Config** | Pointer to [**Durable**](Durable.md) | Config is the consumer&#39;s configuration. | [optional] 
**Created** | Pointer to **time.Time** | Created is when the consumer was created. | [optional] 
**Delivered** | Pointer to [**Sequences**](Sequences.md) | Delivered is the highest delivered sequence pair. | [optional] 
**Name** | Pointer to **string** | Name is the consumer name. | [optional] 
**NumAckPending** | Pointer to **int64** | AckPending is the number of delivered, not yet acknowledged messages. | [optional] 
**NumPending** | Pointer to **int32** | Pending is the number of messages yet to be delivered. | [optional] 
**NumRedelivered** | Pointer to **int64** | Redelivered is the number of messages currently being redelivered. | [optional] 
**NumWaiting** | Pointer to **int64** | Waiting is the number of pull requests waiting for messages. | [optional] 
**StreamName** | Pointer to **string** | Stream is the stream this consumer reads. | [optional] 

## Methods

### NewConsumer

`func NewConsumer() *Consumer`

NewConsumer instantiates a new Consumer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsumerWithDefaults

`func NewConsumerWithDefaults() *Consumer`

NewConsumerWithDefaults instantiates a new Consumer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAckFloor

`func (o *Consumer) GetAckFloor() Sequences`

GetAckFloor returns the AckFloor field if non-nil, zero value otherwise.

### GetAckFloorOk

`func (o *Consumer) GetAckFloorOk() (*Sequences, bool)`

GetAckFloorOk returns a tuple with the AckFloor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAckFloor

`func (o *Consumer) SetAckFloor(v Sequences)`

SetAckFloor sets AckFloor field to given value.

### HasAckFloor

`func (o *Consumer) HasAckFloor() bool`

HasAckFloor returns a boolean if a field has been set.

### GetConfig

`func (o *Consumer) GetConfig() Durable`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *Consumer) GetConfigOk() (*Durable, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *Consumer) SetConfig(v Durable)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *Consumer) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetCreated

`func (o *Consumer) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Consumer) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Consumer) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Consumer) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetDelivered

`func (o *Consumer) GetDelivered() Sequences`

GetDelivered returns the Delivered field if non-nil, zero value otherwise.

### GetDeliveredOk

`func (o *Consumer) GetDeliveredOk() (*Sequences, bool)`

GetDeliveredOk returns a tuple with the Delivered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelivered

`func (o *Consumer) SetDelivered(v Sequences)`

SetDelivered sets Delivered field to given value.

### HasDelivered

`func (o *Consumer) HasDelivered() bool`

HasDelivered returns a boolean if a field has been set.

### GetName

`func (o *Consumer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Consumer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Consumer) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Consumer) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNumAckPending

`func (o *Consumer) GetNumAckPending() int64`

GetNumAckPending returns the NumAckPending field if non-nil, zero value otherwise.

### GetNumAckPendingOk

`func (o *Consumer) GetNumAckPendingOk() (*int64, bool)`

GetNumAckPendingOk returns a tuple with the NumAckPending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumAckPending

`func (o *Consumer) SetNumAckPending(v int64)`

SetNumAckPending sets NumAckPending field to given value.

### HasNumAckPending

`func (o *Consumer) HasNumAckPending() bool`

HasNumAckPending returns a boolean if a field has been set.

### GetNumPending

`func (o *Consumer) GetNumPending() int32`

GetNumPending returns the NumPending field if non-nil, zero value otherwise.

### GetNumPendingOk

`func (o *Consumer) GetNumPendingOk() (*int32, bool)`

GetNumPendingOk returns a tuple with the NumPending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumPending

`func (o *Consumer) SetNumPending(v int32)`

SetNumPending sets NumPending field to given value.

### HasNumPending

`func (o *Consumer) HasNumPending() bool`

HasNumPending returns a boolean if a field has been set.

### GetNumRedelivered

`func (o *Consumer) GetNumRedelivered() int64`

GetNumRedelivered returns the NumRedelivered field if non-nil, zero value otherwise.

### GetNumRedeliveredOk

`func (o *Consumer) GetNumRedeliveredOk() (*int64, bool)`

GetNumRedeliveredOk returns a tuple with the NumRedelivered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumRedelivered

`func (o *Consumer) SetNumRedelivered(v int64)`

SetNumRedelivered sets NumRedelivered field to given value.

### HasNumRedelivered

`func (o *Consumer) HasNumRedelivered() bool`

HasNumRedelivered returns a boolean if a field has been set.

### GetNumWaiting

`func (o *Consumer) GetNumWaiting() int64`

GetNumWaiting returns the NumWaiting field if non-nil, zero value otherwise.

### GetNumWaitingOk

`func (o *Consumer) GetNumWaitingOk() (*int64, bool)`

GetNumWaitingOk returns a tuple with the NumWaiting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumWaiting

`func (o *Consumer) SetNumWaiting(v int64)`

SetNumWaiting sets NumWaiting field to given value.

### HasNumWaiting

`func (o *Consumer) HasNumWaiting() bool`

HasNumWaiting returns a boolean if a field has been set.

### GetStreamName

`func (o *Consumer) GetStreamName() string`

GetStreamName returns the StreamName field if non-nil, zero value otherwise.

### GetStreamNameOk

`func (o *Consumer) GetStreamNameOk() (*string, bool)`

GetStreamNameOk returns a tuple with the StreamName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamName

`func (o *Consumer) SetStreamName(v string)`

SetStreamName sets StreamName field to given value.

### HasStreamName

`func (o *Consumer) HasStreamName() bool`

HasStreamName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


