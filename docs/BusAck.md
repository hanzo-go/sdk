# BusAck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duplicate** | Pointer to **bool** | Duplicate is true when JetStream deduplicated the message by its Nats-Msg-Id instead of storing it again. | [optional] 
**Ok** | Pointer to **bool** | OK is true when the bus accepted the message. | [optional] 
**Seq** | Pointer to **int32** | Seq is the message&#39;s sequence in that stream. | [optional] 
**Stream** | Pointer to **string** | Stream is the stream that stored the message — absent when no stream captures the subject and the message went out core (fire-and-forget). | [optional] 

## Methods

### NewBusAck

`func NewBusAck() *BusAck`

NewBusAck instantiates a new BusAck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBusAckWithDefaults

`func NewBusAckWithDefaults() *BusAck`

NewBusAckWithDefaults instantiates a new BusAck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuplicate

`func (o *BusAck) GetDuplicate() bool`

GetDuplicate returns the Duplicate field if non-nil, zero value otherwise.

### GetDuplicateOk

`func (o *BusAck) GetDuplicateOk() (*bool, bool)`

GetDuplicateOk returns a tuple with the Duplicate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplicate

`func (o *BusAck) SetDuplicate(v bool)`

SetDuplicate sets Duplicate field to given value.

### HasDuplicate

`func (o *BusAck) HasDuplicate() bool`

HasDuplicate returns a boolean if a field has been set.

### GetOk

`func (o *BusAck) GetOk() bool`

GetOk returns the Ok field if non-nil, zero value otherwise.

### GetOkOk

`func (o *BusAck) GetOkOk() (*bool, bool)`

GetOkOk returns a tuple with the Ok field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOk

`func (o *BusAck) SetOk(v bool)`

SetOk sets Ok field to given value.

### HasOk

`func (o *BusAck) HasOk() bool`

HasOk returns a boolean if a field has been set.

### GetSeq

`func (o *BusAck) GetSeq() int32`

GetSeq returns the Seq field if non-nil, zero value otherwise.

### GetSeqOk

`func (o *BusAck) GetSeqOk() (*int32, bool)`

GetSeqOk returns a tuple with the Seq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeq

`func (o *BusAck) SetSeq(v int32)`

SetSeq sets Seq field to given value.

### HasSeq

`func (o *BusAck) HasSeq() bool`

HasSeq returns a boolean if a field has been set.

### GetStream

`func (o *BusAck) GetStream() string`

GetStream returns the Stream field if non-nil, zero value otherwise.

### GetStreamOk

`func (o *BusAck) GetStreamOk() (*string, bool)`

GetStreamOk returns a tuple with the Stream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStream

`func (o *BusAck) SetStream(v string)`

SetStream sets Stream field to given value.

### HasStream

`func (o *BusAck) HasStream() bool`

HasStream returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


