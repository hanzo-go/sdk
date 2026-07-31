# CloudReadOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Messages** | Pointer to [**[]CloudMessage**](CloudMessage.md) | Messages is what was read, stream-ordered. | [optional] 

## Methods

### NewCloudReadOut

`func NewCloudReadOut() *CloudReadOut`

NewCloudReadOut instantiates a new CloudReadOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudReadOutWithDefaults

`func NewCloudReadOutWithDefaults() *CloudReadOut`

NewCloudReadOutWithDefaults instantiates a new CloudReadOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessages

`func (o *CloudReadOut) GetMessages() []CloudMessage`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *CloudReadOut) GetMessagesOk() (*[]CloudMessage, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *CloudReadOut) SetMessages(v []CloudMessage)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *CloudReadOut) HasMessages() bool`

HasMessages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


