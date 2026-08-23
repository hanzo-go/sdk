# ControlResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Command** | Pointer to **string** | Command is the verb that was recorded: pause, resume, stop or message. | [optional] 
**Event** | Pointer to [**EventView**](EventView.md) | Event is the durable control event the command became. The intent is recorded whether or not it reached an engine, which is what makes a stream-consuming surface able to act on it. | [optional] 
**Forwarded** | Pointer to **bool** | Forwarded is whether the command also reached the durable-execution engine. FALSE IS NOT A FAILURE: a session with no workflow link, or a deployment with no tasks backend, is record-only by design. A forward that was attempted and failed is a 502, never a false here. | [optional] 

## Methods

### NewControlResult

`func NewControlResult() *ControlResult`

NewControlResult instantiates a new ControlResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewControlResultWithDefaults

`func NewControlResultWithDefaults() *ControlResult`

NewControlResultWithDefaults instantiates a new ControlResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommand

`func (o *ControlResult) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *ControlResult) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *ControlResult) SetCommand(v string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *ControlResult) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### GetEvent

`func (o *ControlResult) GetEvent() EventView`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *ControlResult) GetEventOk() (*EventView, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *ControlResult) SetEvent(v EventView)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *ControlResult) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetForwarded

`func (o *ControlResult) GetForwarded() bool`

GetForwarded returns the Forwarded field if non-nil, zero value otherwise.

### GetForwardedOk

`func (o *ControlResult) GetForwardedOk() (*bool, bool)`

GetForwardedOk returns a tuple with the Forwarded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwarded

`func (o *ControlResult) SetForwarded(v bool)`

SetForwarded sets Forwarded field to given value.

### HasForwarded

`func (o *ControlResult) HasForwarded() bool`

HasForwarded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


