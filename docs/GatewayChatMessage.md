# GatewayChatMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Role** | **string** |  | 
**Content** | [**GatewayChatMessageContent**](GatewayChatMessageContent.md) |  | 
**Name** | Pointer to **string** |  | [optional] 
**ToolCalls** | Pointer to **[]map[string]interface{}** |  | [optional] 
**ToolCallId** | Pointer to **string** |  | [optional] 

## Methods

### NewGatewayChatMessage

`func NewGatewayChatMessage(role string, content GatewayChatMessageContent, ) *GatewayChatMessage`

NewGatewayChatMessage instantiates a new GatewayChatMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayChatMessageWithDefaults

`func NewGatewayChatMessageWithDefaults() *GatewayChatMessage`

NewGatewayChatMessageWithDefaults instantiates a new GatewayChatMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRole

`func (o *GatewayChatMessage) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *GatewayChatMessage) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *GatewayChatMessage) SetRole(v string)`

SetRole sets Role field to given value.


### GetContent

`func (o *GatewayChatMessage) GetContent() GatewayChatMessageContent`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *GatewayChatMessage) GetContentOk() (*GatewayChatMessageContent, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *GatewayChatMessage) SetContent(v GatewayChatMessageContent)`

SetContent sets Content field to given value.


### GetName

`func (o *GatewayChatMessage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GatewayChatMessage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GatewayChatMessage) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GatewayChatMessage) HasName() bool`

HasName returns a boolean if a field has been set.

### GetToolCalls

`func (o *GatewayChatMessage) GetToolCalls() []map[string]interface{}`

GetToolCalls returns the ToolCalls field if non-nil, zero value otherwise.

### GetToolCallsOk

`func (o *GatewayChatMessage) GetToolCallsOk() (*[]map[string]interface{}, bool)`

GetToolCallsOk returns a tuple with the ToolCalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolCalls

`func (o *GatewayChatMessage) SetToolCalls(v []map[string]interface{})`

SetToolCalls sets ToolCalls field to given value.

### HasToolCalls

`func (o *GatewayChatMessage) HasToolCalls() bool`

HasToolCalls returns a boolean if a field has been set.

### GetToolCallId

`func (o *GatewayChatMessage) GetToolCallId() string`

GetToolCallId returns the ToolCallId field if non-nil, zero value otherwise.

### GetToolCallIdOk

`func (o *GatewayChatMessage) GetToolCallIdOk() (*string, bool)`

GetToolCallIdOk returns a tuple with the ToolCallId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolCallId

`func (o *GatewayChatMessage) SetToolCallId(v string)`

SetToolCallId sets ToolCallId field to given value.

### HasToolCallId

`func (o *GatewayChatMessage) HasToolCallId() bool`

HasToolCallId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


