# MqStreamMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subject** | Pointer to **string** |  | [optional] 
**Data** | Pointer to **string** |  | [optional] 
**Headers** | Pointer to **map[string][]string** |  | [optional] 
**Sequence** | Pointer to **int32** | Stream sequence number. | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**NumDelivered** | Pointer to **int32** | Number of times this message has been delivered. | [optional] 
**NumPending** | Pointer to **int32** | Messages remaining after this one (consumer context). | [optional] 

## Methods

### NewMqStreamMessage

`func NewMqStreamMessage() *MqStreamMessage`

NewMqStreamMessage instantiates a new MqStreamMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMqStreamMessageWithDefaults

`func NewMqStreamMessageWithDefaults() *MqStreamMessage`

NewMqStreamMessageWithDefaults instantiates a new MqStreamMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubject

`func (o *MqStreamMessage) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *MqStreamMessage) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *MqStreamMessage) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *MqStreamMessage) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetData

`func (o *MqStreamMessage) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MqStreamMessage) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MqStreamMessage) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *MqStreamMessage) HasData() bool`

HasData returns a boolean if a field has been set.

### GetHeaders

`func (o *MqStreamMessage) GetHeaders() map[string][]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *MqStreamMessage) GetHeadersOk() (*map[string][]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *MqStreamMessage) SetHeaders(v map[string][]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *MqStreamMessage) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetSequence

`func (o *MqStreamMessage) GetSequence() int32`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *MqStreamMessage) GetSequenceOk() (*int32, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *MqStreamMessage) SetSequence(v int32)`

SetSequence sets Sequence field to given value.

### HasSequence

`func (o *MqStreamMessage) HasSequence() bool`

HasSequence returns a boolean if a field has been set.

### GetTimestamp

`func (o *MqStreamMessage) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *MqStreamMessage) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *MqStreamMessage) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *MqStreamMessage) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetNumDelivered

`func (o *MqStreamMessage) GetNumDelivered() int32`

GetNumDelivered returns the NumDelivered field if non-nil, zero value otherwise.

### GetNumDeliveredOk

`func (o *MqStreamMessage) GetNumDeliveredOk() (*int32, bool)`

GetNumDeliveredOk returns a tuple with the NumDelivered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumDelivered

`func (o *MqStreamMessage) SetNumDelivered(v int32)`

SetNumDelivered sets NumDelivered field to given value.

### HasNumDelivered

`func (o *MqStreamMessage) HasNumDelivered() bool`

HasNumDelivered returns a boolean if a field has been set.

### GetNumPending

`func (o *MqStreamMessage) GetNumPending() int32`

GetNumPending returns the NumPending field if non-nil, zero value otherwise.

### GetNumPendingOk

`func (o *MqStreamMessage) GetNumPendingOk() (*int32, bool)`

GetNumPendingOk returns a tuple with the NumPending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumPending

`func (o *MqStreamMessage) SetNumPending(v int32)`

SetNumPending sets NumPending field to given value.

### HasNumPending

`func (o *MqStreamMessage) HasNumPending() bool`

HasNumPending returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


