# AgentsControlResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Command** | Pointer to **string** |  | [optional] 
**Event** | Pointer to [**AgentsEventView**](AgentsEventView.md) |  | [optional] 
**Forwarded** | Pointer to **bool** | True when the command was forwarded to the hanzoai/tasks engine. | [optional] 

## Methods

### NewAgentsControlResult

`func NewAgentsControlResult() *AgentsControlResult`

NewAgentsControlResult instantiates a new AgentsControlResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentsControlResultWithDefaults

`func NewAgentsControlResultWithDefaults() *AgentsControlResult`

NewAgentsControlResultWithDefaults instantiates a new AgentsControlResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommand

`func (o *AgentsControlResult) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *AgentsControlResult) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *AgentsControlResult) SetCommand(v string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *AgentsControlResult) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### GetEvent

`func (o *AgentsControlResult) GetEvent() AgentsEventView`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *AgentsControlResult) GetEventOk() (*AgentsEventView, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *AgentsControlResult) SetEvent(v AgentsEventView)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *AgentsControlResult) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetForwarded

`func (o *AgentsControlResult) GetForwarded() bool`

GetForwarded returns the Forwarded field if non-nil, zero value otherwise.

### GetForwardedOk

`func (o *AgentsControlResult) GetForwardedOk() (*bool, bool)`

GetForwardedOk returns a tuple with the Forwarded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwarded

`func (o *AgentsControlResult) SetForwarded(v bool)`

SetForwarded sets Forwarded field to given value.

### HasForwarded

`func (o *AgentsControlResult) HasForwarded() bool`

HasForwarded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


