# CloudConsumer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AckFloor** | Pointer to [**CloudSequences**](CloudSequences.md) | AckFloor is the highest contiguously acknowledged sequence pair. | [optional] 
**Config** | Pointer to [**CloudDurable**](CloudDurable.md) | Config is the consumer&#39;s configuration. | [optional] 
**Created** | Pointer to **time.Time** | Created is when the consumer was created. | [optional] 
**Delivered** | Pointer to [**CloudSequences**](CloudSequences.md) | Delivered is the highest delivered sequence pair. | [optional] 
**Name** | Pointer to **string** | Name is the consumer name. | [optional] 
**NumAckPending** | Pointer to **int32** | AckPending is the number of delivered, not yet acknowledged messages. | [optional] 
**NumPending** | Pointer to **int32** | Pending is the number of messages yet to be delivered. | [optional] 
**NumRedelivered** | Pointer to **int32** | Redelivered is the number of messages currently being redelivered. | [optional] 
**NumWaiting** | Pointer to **int32** | Waiting is the number of pull requests waiting for messages. | [optional] 
**StreamName** | Pointer to **string** | Stream is the stream this consumer reads. | [optional] 

## Methods

### NewCloudConsumer

`func NewCloudConsumer() *CloudConsumer`

NewCloudConsumer instantiates a new CloudConsumer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudConsumerWithDefaults

`func NewCloudConsumerWithDefaults() *CloudConsumer`

NewCloudConsumerWithDefaults instantiates a new CloudConsumer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAckFloor

`func (o *CloudConsumer) GetAckFloor() CloudSequences`

GetAckFloor returns the AckFloor field if non-nil, zero value otherwise.

### GetAckFloorOk

`func (o *CloudConsumer) GetAckFloorOk() (*CloudSequences, bool)`

GetAckFloorOk returns a tuple with the AckFloor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAckFloor

`func (o *CloudConsumer) SetAckFloor(v CloudSequences)`

SetAckFloor sets AckFloor field to given value.

### HasAckFloor

`func (o *CloudConsumer) HasAckFloor() bool`

HasAckFloor returns a boolean if a field has been set.

### GetConfig

`func (o *CloudConsumer) GetConfig() CloudDurable`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *CloudConsumer) GetConfigOk() (*CloudDurable, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *CloudConsumer) SetConfig(v CloudDurable)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *CloudConsumer) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetCreated

`func (o *CloudConsumer) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CloudConsumer) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *CloudConsumer) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *CloudConsumer) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetDelivered

`func (o *CloudConsumer) GetDelivered() CloudSequences`

GetDelivered returns the Delivered field if non-nil, zero value otherwise.

### GetDeliveredOk

`func (o *CloudConsumer) GetDeliveredOk() (*CloudSequences, bool)`

GetDeliveredOk returns a tuple with the Delivered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelivered

`func (o *CloudConsumer) SetDelivered(v CloudSequences)`

SetDelivered sets Delivered field to given value.

### HasDelivered

`func (o *CloudConsumer) HasDelivered() bool`

HasDelivered returns a boolean if a field has been set.

### GetName

`func (o *CloudConsumer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudConsumer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudConsumer) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudConsumer) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNumAckPending

`func (o *CloudConsumer) GetNumAckPending() int32`

GetNumAckPending returns the NumAckPending field if non-nil, zero value otherwise.

### GetNumAckPendingOk

`func (o *CloudConsumer) GetNumAckPendingOk() (*int32, bool)`

GetNumAckPendingOk returns a tuple with the NumAckPending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumAckPending

`func (o *CloudConsumer) SetNumAckPending(v int32)`

SetNumAckPending sets NumAckPending field to given value.

### HasNumAckPending

`func (o *CloudConsumer) HasNumAckPending() bool`

HasNumAckPending returns a boolean if a field has been set.

### GetNumPending

`func (o *CloudConsumer) GetNumPending() int32`

GetNumPending returns the NumPending field if non-nil, zero value otherwise.

### GetNumPendingOk

`func (o *CloudConsumer) GetNumPendingOk() (*int32, bool)`

GetNumPendingOk returns a tuple with the NumPending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumPending

`func (o *CloudConsumer) SetNumPending(v int32)`

SetNumPending sets NumPending field to given value.

### HasNumPending

`func (o *CloudConsumer) HasNumPending() bool`

HasNumPending returns a boolean if a field has been set.

### GetNumRedelivered

`func (o *CloudConsumer) GetNumRedelivered() int32`

GetNumRedelivered returns the NumRedelivered field if non-nil, zero value otherwise.

### GetNumRedeliveredOk

`func (o *CloudConsumer) GetNumRedeliveredOk() (*int32, bool)`

GetNumRedeliveredOk returns a tuple with the NumRedelivered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumRedelivered

`func (o *CloudConsumer) SetNumRedelivered(v int32)`

SetNumRedelivered sets NumRedelivered field to given value.

### HasNumRedelivered

`func (o *CloudConsumer) HasNumRedelivered() bool`

HasNumRedelivered returns a boolean if a field has been set.

### GetNumWaiting

`func (o *CloudConsumer) GetNumWaiting() int32`

GetNumWaiting returns the NumWaiting field if non-nil, zero value otherwise.

### GetNumWaitingOk

`func (o *CloudConsumer) GetNumWaitingOk() (*int32, bool)`

GetNumWaitingOk returns a tuple with the NumWaiting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumWaiting

`func (o *CloudConsumer) SetNumWaiting(v int32)`

SetNumWaiting sets NumWaiting field to given value.

### HasNumWaiting

`func (o *CloudConsumer) HasNumWaiting() bool`

HasNumWaiting returns a boolean if a field has been set.

### GetStreamName

`func (o *CloudConsumer) GetStreamName() string`

GetStreamName returns the StreamName field if non-nil, zero value otherwise.

### GetStreamNameOk

`func (o *CloudConsumer) GetStreamNameOk() (*string, bool)`

GetStreamNameOk returns a tuple with the StreamName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamName

`func (o *CloudConsumer) SetStreamName(v string)`

SetStreamName sets StreamName field to given value.

### HasStreamName

`func (o *CloudConsumer) HasStreamName() bool`

HasStreamName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


