# AgentControlResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Command** | Pointer to **string** | Command is the verb that was recorded: pause, resume, stop or message. | [optional] 
**Event** | Pointer to [**AgentEventView**](AgentEventView.md) | Event is the durable control event the command became. The intent is recorded whether or not it reached an engine, which is what makes a stream-consuming surface able to act on it. | [optional] 
**Forwarded** | Pointer to **bool** | Forwarded is whether the command also reached the durable-execution engine. FALSE IS NOT A FAILURE: a session with no workflow link, or a deployment with no tasks backend, is record-only by design. A forward that was attempted and failed is a 502, never a false here. | [optional] 
**Run** | Pointer to **string** | Run is the run a resume started, under the same session. Empty for every other command. | [optional] 

## Methods

### NewAgentControlResult

`func NewAgentControlResult() *AgentControlResult`

NewAgentControlResult instantiates a new AgentControlResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentControlResultWithDefaults

`func NewAgentControlResultWithDefaults() *AgentControlResult`

NewAgentControlResultWithDefaults instantiates a new AgentControlResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommand

`func (o *AgentControlResult) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *AgentControlResult) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *AgentControlResult) SetCommand(v string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *AgentControlResult) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### GetEvent

`func (o *AgentControlResult) GetEvent() AgentEventView`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *AgentControlResult) GetEventOk() (*AgentEventView, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *AgentControlResult) SetEvent(v AgentEventView)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *AgentControlResult) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetForwarded

`func (o *AgentControlResult) GetForwarded() bool`

GetForwarded returns the Forwarded field if non-nil, zero value otherwise.

### GetForwardedOk

`func (o *AgentControlResult) GetForwardedOk() (*bool, bool)`

GetForwardedOk returns a tuple with the Forwarded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwarded

`func (o *AgentControlResult) SetForwarded(v bool)`

SetForwarded sets Forwarded field to given value.

### HasForwarded

`func (o *AgentControlResult) HasForwarded() bool`

HasForwarded returns a boolean if a field has been set.

### GetRun

`func (o *AgentControlResult) GetRun() string`

GetRun returns the Run field if non-nil, zero value otherwise.

### GetRunOk

`func (o *AgentControlResult) GetRunOk() (*string, bool)`

GetRunOk returns a tuple with the Run field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRun

`func (o *AgentControlResult) SetRun(v string)`

SetRun sets Run field to given value.

### HasRun

`func (o *AgentControlResult) HasRun() bool`

HasRun returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


