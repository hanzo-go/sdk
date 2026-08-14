# ConsumerRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ack** | Pointer to **string** | Ack is the acknowledgement discipline: explicit, none or all. | [optional] 
**AckWait** | Pointer to **int32** | AckWait is the redelivery timeout in seconds. | [optional] 
**Acked** | Pointer to **int32** | Acked is the stream sequence acknowledged furthest. | [optional] 
**Deliver** | Pointer to **string** | Deliver is the starting point: all, last, new or lastPerSubject. | [optional] 
**Delivered** | Pointer to **int32** | Delivered is the stream sequence delivered furthest. | [optional] 
**Filter** | Pointer to **string** | Filter is the subject filter, in the org&#39;s namespace; empty means all. | [optional] 
**MaxDeliver** | Pointer to **int32** | MaxDeliver is the delivery-attempt cap; -1 means unlimited. | [optional] 
**Name** | Pointer to **string** | Name is the durable consumer name. | [optional] 
**Pending** | Pointer to **int32** | Pending is how many messages await delivery. | [optional] 
**Redelivered** | Pointer to **int32** | Redelivered is how many messages are being redelivered. | [optional] 
**Stream** | Pointer to **string** | Stream is the stream it consumes, in the org&#39;s view. | [optional] 

## Methods

### NewConsumerRecord

`func NewConsumerRecord() *ConsumerRecord`

NewConsumerRecord instantiates a new ConsumerRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsumerRecordWithDefaults

`func NewConsumerRecordWithDefaults() *ConsumerRecord`

NewConsumerRecordWithDefaults instantiates a new ConsumerRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAck

`func (o *ConsumerRecord) GetAck() string`

GetAck returns the Ack field if non-nil, zero value otherwise.

### GetAckOk

`func (o *ConsumerRecord) GetAckOk() (*string, bool)`

GetAckOk returns a tuple with the Ack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAck

`func (o *ConsumerRecord) SetAck(v string)`

SetAck sets Ack field to given value.

### HasAck

`func (o *ConsumerRecord) HasAck() bool`

HasAck returns a boolean if a field has been set.

### GetAckWait

`func (o *ConsumerRecord) GetAckWait() int32`

GetAckWait returns the AckWait field if non-nil, zero value otherwise.

### GetAckWaitOk

`func (o *ConsumerRecord) GetAckWaitOk() (*int32, bool)`

GetAckWaitOk returns a tuple with the AckWait field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAckWait

`func (o *ConsumerRecord) SetAckWait(v int32)`

SetAckWait sets AckWait field to given value.

### HasAckWait

`func (o *ConsumerRecord) HasAckWait() bool`

HasAckWait returns a boolean if a field has been set.

### GetAcked

`func (o *ConsumerRecord) GetAcked() int32`

GetAcked returns the Acked field if non-nil, zero value otherwise.

### GetAckedOk

`func (o *ConsumerRecord) GetAckedOk() (*int32, bool)`

GetAckedOk returns a tuple with the Acked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcked

`func (o *ConsumerRecord) SetAcked(v int32)`

SetAcked sets Acked field to given value.

### HasAcked

`func (o *ConsumerRecord) HasAcked() bool`

HasAcked returns a boolean if a field has been set.

### GetDeliver

`func (o *ConsumerRecord) GetDeliver() string`

GetDeliver returns the Deliver field if non-nil, zero value otherwise.

### GetDeliverOk

`func (o *ConsumerRecord) GetDeliverOk() (*string, bool)`

GetDeliverOk returns a tuple with the Deliver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliver

`func (o *ConsumerRecord) SetDeliver(v string)`

SetDeliver sets Deliver field to given value.

### HasDeliver

`func (o *ConsumerRecord) HasDeliver() bool`

HasDeliver returns a boolean if a field has been set.

### GetDelivered

`func (o *ConsumerRecord) GetDelivered() int32`

GetDelivered returns the Delivered field if non-nil, zero value otherwise.

### GetDeliveredOk

`func (o *ConsumerRecord) GetDeliveredOk() (*int32, bool)`

GetDeliveredOk returns a tuple with the Delivered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelivered

`func (o *ConsumerRecord) SetDelivered(v int32)`

SetDelivered sets Delivered field to given value.

### HasDelivered

`func (o *ConsumerRecord) HasDelivered() bool`

HasDelivered returns a boolean if a field has been set.

### GetFilter

`func (o *ConsumerRecord) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *ConsumerRecord) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *ConsumerRecord) SetFilter(v string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *ConsumerRecord) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetMaxDeliver

`func (o *ConsumerRecord) GetMaxDeliver() int32`

GetMaxDeliver returns the MaxDeliver field if non-nil, zero value otherwise.

### GetMaxDeliverOk

`func (o *ConsumerRecord) GetMaxDeliverOk() (*int32, bool)`

GetMaxDeliverOk returns a tuple with the MaxDeliver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDeliver

`func (o *ConsumerRecord) SetMaxDeliver(v int32)`

SetMaxDeliver sets MaxDeliver field to given value.

### HasMaxDeliver

`func (o *ConsumerRecord) HasMaxDeliver() bool`

HasMaxDeliver returns a boolean if a field has been set.

### GetName

`func (o *ConsumerRecord) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConsumerRecord) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConsumerRecord) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConsumerRecord) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPending

`func (o *ConsumerRecord) GetPending() int32`

GetPending returns the Pending field if non-nil, zero value otherwise.

### GetPendingOk

`func (o *ConsumerRecord) GetPendingOk() (*int32, bool)`

GetPendingOk returns a tuple with the Pending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPending

`func (o *ConsumerRecord) SetPending(v int32)`

SetPending sets Pending field to given value.

### HasPending

`func (o *ConsumerRecord) HasPending() bool`

HasPending returns a boolean if a field has been set.

### GetRedelivered

`func (o *ConsumerRecord) GetRedelivered() int32`

GetRedelivered returns the Redelivered field if non-nil, zero value otherwise.

### GetRedeliveredOk

`func (o *ConsumerRecord) GetRedeliveredOk() (*int32, bool)`

GetRedeliveredOk returns a tuple with the Redelivered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedelivered

`func (o *ConsumerRecord) SetRedelivered(v int32)`

SetRedelivered sets Redelivered field to given value.

### HasRedelivered

`func (o *ConsumerRecord) HasRedelivered() bool`

HasRedelivered returns a boolean if a field has been set.

### GetStream

`func (o *ConsumerRecord) GetStream() string`

GetStream returns the Stream field if non-nil, zero value otherwise.

### GetStreamOk

`func (o *ConsumerRecord) GetStreamOk() (*string, bool)`

GetStreamOk returns a tuple with the Stream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStream

`func (o *ConsumerRecord) SetStream(v string)`

SetStream sets Stream field to given value.

### HasStream

`func (o *ConsumerRecord) HasStream() bool`

HasStream returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


