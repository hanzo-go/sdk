# AgentRunList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Runs** | Pointer to [**[]AgentAgentRunView**](AgentAgentRunView.md) | Runs is the agent&#39;s executions, newest first. | [optional] 

## Methods

### NewAgentRunList

`func NewAgentRunList() *AgentRunList`

NewAgentRunList instantiates a new AgentRunList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentRunListWithDefaults

`func NewAgentRunListWithDefaults() *AgentRunList`

NewAgentRunListWithDefaults instantiates a new AgentRunList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRuns

`func (o *AgentRunList) GetRuns() []AgentAgentRunView`

GetRuns returns the Runs field if non-nil, zero value otherwise.

### GetRunsOk

`func (o *AgentRunList) GetRunsOk() (*[]AgentAgentRunView, bool)`

GetRunsOk returns a tuple with the Runs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuns

`func (o *AgentRunList) SetRuns(v []AgentAgentRunView)`

SetRuns sets Runs field to given value.

### HasRuns

`func (o *AgentRunList) HasRuns() bool`

HasRuns returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


