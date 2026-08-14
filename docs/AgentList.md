# AgentList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Agents** | Pointer to [**[]AgentView**](AgentView.md) | Agents is the org&#39;s agents, each carrying its recorded run count. | [optional] 

## Methods

### NewAgentList

`func NewAgentList() *AgentList`

NewAgentList instantiates a new AgentList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentListWithDefaults

`func NewAgentListWithDefaults() *AgentList`

NewAgentListWithDefaults instantiates a new AgentList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgents

`func (o *AgentList) GetAgents() []AgentView`

GetAgents returns the Agents field if non-nil, zero value otherwise.

### GetAgentsOk

`func (o *AgentList) GetAgentsOk() (*[]AgentView, bool)`

GetAgentsOk returns a tuple with the Agents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgents

`func (o *AgentList) SetAgents(v []AgentView)`

SetAgents sets Agents field to given value.

### HasAgents

`func (o *AgentList) HasAgents() bool`

HasAgents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


