# AgentControlIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID is the session to steer, from the path. | [optional] 
**Message** | Pointer to **string** | Message is free text for the running agent, up to 16 KiB. On a stop it is recorded as the cancellation reason. | [optional] 
**Payload** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewAgentControlIn

`func NewAgentControlIn() *AgentControlIn`

NewAgentControlIn instantiates a new AgentControlIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentControlInWithDefaults

`func NewAgentControlInWithDefaults() *AgentControlIn`

NewAgentControlInWithDefaults instantiates a new AgentControlIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AgentControlIn) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AgentControlIn) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AgentControlIn) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AgentControlIn) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMessage

`func (o *AgentControlIn) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AgentControlIn) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AgentControlIn) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *AgentControlIn) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetPayload

`func (o *AgentControlIn) GetPayload() interface{}`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *AgentControlIn) GetPayloadOk() (*interface{}, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *AgentControlIn) SetPayload(v interface{})`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *AgentControlIn) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### SetPayloadNil

`func (o *AgentControlIn) SetPayloadNil(b bool)`

 SetPayloadNil sets the value for Payload to be an explicit nil

### UnsetPayload
`func (o *AgentControlIn) UnsetPayload()`

UnsetPayload ensures that no value is present for Payload, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


