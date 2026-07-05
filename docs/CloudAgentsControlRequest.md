# CloudAgentsControlRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**Payload** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewCloudAgentsControlRequest

`func NewCloudAgentsControlRequest() *CloudAgentsControlRequest`

NewCloudAgentsControlRequest instantiates a new CloudAgentsControlRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudAgentsControlRequestWithDefaults

`func NewCloudAgentsControlRequestWithDefaults() *CloudAgentsControlRequest`

NewCloudAgentsControlRequestWithDefaults instantiates a new CloudAgentsControlRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *CloudAgentsControlRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CloudAgentsControlRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CloudAgentsControlRequest) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *CloudAgentsControlRequest) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetPayload

`func (o *CloudAgentsControlRequest) GetPayload() map[string]interface{}`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *CloudAgentsControlRequest) GetPayloadOk() (*map[string]interface{}, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *CloudAgentsControlRequest) SetPayload(v map[string]interface{})`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *CloudAgentsControlRequest) HasPayload() bool`

HasPayload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


