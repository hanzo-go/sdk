# MqSubjectInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subject** | Pointer to **string** | Subject name. | [optional] 
**Subscribers** | Pointer to **int32** | Number of active subscribers. | [optional] 
**MessagesPerSec** | Pointer to **float64** | Current message rate (messages per second). | [optional] 
**BytesPerSec** | Pointer to **int64** | Current data rate (bytes per second). | [optional] 

## Methods

### NewMqSubjectInfo

`func NewMqSubjectInfo() *MqSubjectInfo`

NewMqSubjectInfo instantiates a new MqSubjectInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMqSubjectInfoWithDefaults

`func NewMqSubjectInfoWithDefaults() *MqSubjectInfo`

NewMqSubjectInfoWithDefaults instantiates a new MqSubjectInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubject

`func (o *MqSubjectInfo) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *MqSubjectInfo) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *MqSubjectInfo) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *MqSubjectInfo) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetSubscribers

`func (o *MqSubjectInfo) GetSubscribers() int32`

GetSubscribers returns the Subscribers field if non-nil, zero value otherwise.

### GetSubscribersOk

`func (o *MqSubjectInfo) GetSubscribersOk() (*int32, bool)`

GetSubscribersOk returns a tuple with the Subscribers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribers

`func (o *MqSubjectInfo) SetSubscribers(v int32)`

SetSubscribers sets Subscribers field to given value.

### HasSubscribers

`func (o *MqSubjectInfo) HasSubscribers() bool`

HasSubscribers returns a boolean if a field has been set.

### GetMessagesPerSec

`func (o *MqSubjectInfo) GetMessagesPerSec() float64`

GetMessagesPerSec returns the MessagesPerSec field if non-nil, zero value otherwise.

### GetMessagesPerSecOk

`func (o *MqSubjectInfo) GetMessagesPerSecOk() (*float64, bool)`

GetMessagesPerSecOk returns a tuple with the MessagesPerSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagesPerSec

`func (o *MqSubjectInfo) SetMessagesPerSec(v float64)`

SetMessagesPerSec sets MessagesPerSec field to given value.

### HasMessagesPerSec

`func (o *MqSubjectInfo) HasMessagesPerSec() bool`

HasMessagesPerSec returns a boolean if a field has been set.

### GetBytesPerSec

`func (o *MqSubjectInfo) GetBytesPerSec() int64`

GetBytesPerSec returns the BytesPerSec field if non-nil, zero value otherwise.

### GetBytesPerSecOk

`func (o *MqSubjectInfo) GetBytesPerSecOk() (*int64, bool)`

GetBytesPerSecOk returns a tuple with the BytesPerSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytesPerSec

`func (o *MqSubjectInfo) SetBytesPerSec(v int64)`

SetBytesPerSec sets BytesPerSec field to given value.

### HasBytesPerSec

`func (o *MqSubjectInfo) HasBytesPerSec() bool`

HasBytesPerSec returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


