# KvPubSubMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewKvPubSubMessage

`func NewKvPubSubMessage() *KvPubSubMessage`

NewKvPubSubMessage instantiates a new KvPubSubMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKvPubSubMessageWithDefaults

`func NewKvPubSubMessageWithDefaults() *KvPubSubMessage`

NewKvPubSubMessageWithDefaults instantiates a new KvPubSubMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *KvPubSubMessage) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *KvPubSubMessage) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *KvPubSubMessage) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *KvPubSubMessage) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetMessage

`func (o *KvPubSubMessage) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *KvPubSubMessage) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *KvPubSubMessage) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *KvPubSubMessage) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetTimestamp

`func (o *KvPubSubMessage) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *KvPubSubMessage) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *KvPubSubMessage) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *KvPubSubMessage) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


