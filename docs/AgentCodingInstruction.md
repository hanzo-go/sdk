# AgentCodingInstruction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Session** | Pointer to **string** | Session identifies the running or paused coding session from the path. | [optional] 
**Text** | Pointer to **string** | Text is the new instruction to carry into the follow-up run. | [optional] 

## Methods

### NewAgentCodingInstruction

`func NewAgentCodingInstruction() *AgentCodingInstruction`

NewAgentCodingInstruction instantiates a new AgentCodingInstruction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentCodingInstructionWithDefaults

`func NewAgentCodingInstructionWithDefaults() *AgentCodingInstruction`

NewAgentCodingInstructionWithDefaults instantiates a new AgentCodingInstruction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSession

`func (o *AgentCodingInstruction) GetSession() string`

GetSession returns the Session field if non-nil, zero value otherwise.

### GetSessionOk

`func (o *AgentCodingInstruction) GetSessionOk() (*string, bool)`

GetSessionOk returns a tuple with the Session field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSession

`func (o *AgentCodingInstruction) SetSession(v string)`

SetSession sets Session field to given value.

### HasSession

`func (o *AgentCodingInstruction) HasSession() bool`

HasSession returns a boolean if a field has been set.

### GetText

`func (o *AgentCodingInstruction) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *AgentCodingInstruction) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *AgentCodingInstruction) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *AgentCodingInstruction) HasText() bool`

HasText returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


