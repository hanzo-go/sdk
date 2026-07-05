# MqGetStreamMessages200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Messages** | Pointer to [**[]MqStreamMessage**](MqStreamMessage.md) |  | [optional] 

## Methods

### NewMqGetStreamMessages200Response

`func NewMqGetStreamMessages200Response() *MqGetStreamMessages200Response`

NewMqGetStreamMessages200Response instantiates a new MqGetStreamMessages200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMqGetStreamMessages200ResponseWithDefaults

`func NewMqGetStreamMessages200ResponseWithDefaults() *MqGetStreamMessages200Response`

NewMqGetStreamMessages200ResponseWithDefaults instantiates a new MqGetStreamMessages200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessages

`func (o *MqGetStreamMessages200Response) GetMessages() []MqStreamMessage`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *MqGetStreamMessages200Response) GetMessagesOk() (*[]MqStreamMessage, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *MqGetStreamMessages200Response) SetMessages(v []MqStreamMessage)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *MqGetStreamMessages200Response) HasMessages() bool`

HasMessages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


