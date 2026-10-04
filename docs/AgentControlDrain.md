# AgentControlDrain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Commands** | Pointer to [**[]AgentControlCommandView**](AgentControlCommandView.md) | Commands is the session&#39;s control commands newer than the cursor, oldest first. | [optional] 
**Cursor** | Pointer to **int64** | Cursor is the seq to send as &#x60;after&#x60; on the next poll — the highest seq in this page, or the cursor sent in when the page is empty. | [optional] 

## Methods

### NewAgentControlDrain

`func NewAgentControlDrain() *AgentControlDrain`

NewAgentControlDrain instantiates a new AgentControlDrain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentControlDrainWithDefaults

`func NewAgentControlDrainWithDefaults() *AgentControlDrain`

NewAgentControlDrainWithDefaults instantiates a new AgentControlDrain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommands

`func (o *AgentControlDrain) GetCommands() []AgentControlCommandView`

GetCommands returns the Commands field if non-nil, zero value otherwise.

### GetCommandsOk

`func (o *AgentControlDrain) GetCommandsOk() (*[]AgentControlCommandView, bool)`

GetCommandsOk returns a tuple with the Commands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommands

`func (o *AgentControlDrain) SetCommands(v []AgentControlCommandView)`

SetCommands sets Commands field to given value.

### HasCommands

`func (o *AgentControlDrain) HasCommands() bool`

HasCommands returns a boolean if a field has been set.

### GetCursor

`func (o *AgentControlDrain) GetCursor() int64`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *AgentControlDrain) GetCursorOk() (*int64, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *AgentControlDrain) SetCursor(v int64)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *AgentControlDrain) HasCursor() bool`

HasCursor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


