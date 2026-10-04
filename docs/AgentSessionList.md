# AgentSessionList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Next** | Pointer to **string** | Next is the cursor for the page after this one; empty on the last page. | [optional] 
**Sessions** | Pointer to [**[]AgentSessionView**](AgentSessionView.md) | Sessions is the matching sessions, each with its event and child counts and a one-line preview of its latest event. | [optional] 

## Methods

### NewAgentSessionList

`func NewAgentSessionList() *AgentSessionList`

NewAgentSessionList instantiates a new AgentSessionList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentSessionListWithDefaults

`func NewAgentSessionListWithDefaults() *AgentSessionList`

NewAgentSessionListWithDefaults instantiates a new AgentSessionList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNext

`func (o *AgentSessionList) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *AgentSessionList) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *AgentSessionList) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *AgentSessionList) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetSessions

`func (o *AgentSessionList) GetSessions() []AgentSessionView`

GetSessions returns the Sessions field if non-nil, zero value otherwise.

### GetSessionsOk

`func (o *AgentSessionList) GetSessionsOk() (*[]AgentSessionView, bool)`

GetSessionsOk returns a tuple with the Sessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessions

`func (o *AgentSessionList) SetSessions(v []AgentSessionView)`

SetSessions sets Sessions field to given value.

### HasSessions

`func (o *AgentSessionList) HasSessions() bool`

HasSessions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


