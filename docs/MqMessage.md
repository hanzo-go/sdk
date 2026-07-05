# MqMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subject** | Pointer to **string** | Subject the message was published to. | [optional] 
**Data** | Pointer to **string** | Message payload (base64-encoded for binary data). | [optional] 
**Headers** | Pointer to **map[string][]string** | Message headers (key to list of values). | [optional] 
**Reply** | Pointer to **string** | Reply-to subject for request/reply patterns. | [optional] 
**Timestamp** | Pointer to **time.Time** | Server timestamp when the message was received. | [optional] 
**Sequence** | Pointer to **int32** | Stream sequence number (present only for JetStream messages).  | [optional] 

## Methods

### NewMqMessage

`func NewMqMessage() *MqMessage`

NewMqMessage instantiates a new MqMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMqMessageWithDefaults

`func NewMqMessageWithDefaults() *MqMessage`

NewMqMessageWithDefaults instantiates a new MqMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubject

`func (o *MqMessage) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *MqMessage) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *MqMessage) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *MqMessage) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetData

`func (o *MqMessage) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MqMessage) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MqMessage) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *MqMessage) HasData() bool`

HasData returns a boolean if a field has been set.

### GetHeaders

`func (o *MqMessage) GetHeaders() map[string][]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *MqMessage) GetHeadersOk() (*map[string][]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *MqMessage) SetHeaders(v map[string][]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *MqMessage) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetReply

`func (o *MqMessage) GetReply() string`

GetReply returns the Reply field if non-nil, zero value otherwise.

### GetReplyOk

`func (o *MqMessage) GetReplyOk() (*string, bool)`

GetReplyOk returns a tuple with the Reply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReply

`func (o *MqMessage) SetReply(v string)`

SetReply sets Reply field to given value.

### HasReply

`func (o *MqMessage) HasReply() bool`

HasReply returns a boolean if a field has been set.

### GetTimestamp

`func (o *MqMessage) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *MqMessage) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *MqMessage) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *MqMessage) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetSequence

`func (o *MqMessage) GetSequence() int32`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *MqMessage) GetSequenceOk() (*int32, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *MqMessage) SetSequence(v int32)`

SetSequence sets Sequence field to given value.

### HasSequence

`func (o *MqMessage) HasSequence() bool`

HasSequence returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


