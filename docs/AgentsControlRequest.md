# AgentsControlRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** | Optional steering text (max 16 KiB). Required for the message command when no payload is given. | [optional] 
**Payload** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewAgentsControlRequest

`func NewAgentsControlRequest() *AgentsControlRequest`

NewAgentsControlRequest instantiates a new AgentsControlRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentsControlRequestWithDefaults

`func NewAgentsControlRequestWithDefaults() *AgentsControlRequest`

NewAgentsControlRequestWithDefaults instantiates a new AgentsControlRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *AgentsControlRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AgentsControlRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AgentsControlRequest) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *AgentsControlRequest) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetPayload

`func (o *AgentsControlRequest) GetPayload() interface{}`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *AgentsControlRequest) GetPayloadOk() (*interface{}, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *AgentsControlRequest) SetPayload(v interface{})`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *AgentsControlRequest) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### SetPayloadNil

`func (o *AgentsControlRequest) SetPayloadNil(b bool)`

 SetPayloadNil sets the value for Payload to be an explicit nil

### UnsetPayload
`func (o *AgentsControlRequest) UnsetPayload()`

UnsetPayload ensures that no value is present for Payload, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


