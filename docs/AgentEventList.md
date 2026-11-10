# AgentEventList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | Pointer to [**[]AgentEventView**](AgentEventView.md) | Events is the page, in seq order. | [optional] 
**Next** | Pointer to **int64** | Next is the seq to pass as after for the page beyond this one; absent when this page reached the end of the log as it stands now. | [optional] 

## Methods

### NewAgentEventList

`func NewAgentEventList() *AgentEventList`

NewAgentEventList instantiates a new AgentEventList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentEventListWithDefaults

`func NewAgentEventListWithDefaults() *AgentEventList`

NewAgentEventListWithDefaults instantiates a new AgentEventList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *AgentEventList) GetEvents() []AgentEventView`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *AgentEventList) GetEventsOk() (*[]AgentEventView, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *AgentEventList) SetEvents(v []AgentEventView)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *AgentEventList) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetNext

`func (o *AgentEventList) GetNext() int64`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *AgentEventList) GetNextOk() (*int64, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *AgentEventList) SetNext(v int64)`

SetNext sets Next field to given value.

### HasNext

`func (o *AgentEventList) HasNext() bool`

HasNext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


