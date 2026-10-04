# MqReadOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Messages** | Pointer to [**[]MqDelivery**](MqDelivery.md) | Messages is what was read, stream-ordered. | [optional] 

## Methods

### NewMqReadOut

`func NewMqReadOut() *MqReadOut`

NewMqReadOut instantiates a new MqReadOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMqReadOutWithDefaults

`func NewMqReadOutWithDefaults() *MqReadOut`

NewMqReadOutWithDefaults instantiates a new MqReadOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessages

`func (o *MqReadOut) GetMessages() []MqDelivery`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *MqReadOut) GetMessagesOk() (*[]MqDelivery, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *MqReadOut) SetMessages(v []MqDelivery)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *MqReadOut) HasMessages() bool`

HasMessages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


