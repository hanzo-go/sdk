# AgentBuildList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Builds** | Pointer to [**[]AgentBuildSummary**](AgentBuildSummary.md) | Builds is every published build, most recently updated first. | [optional] 

## Methods

### NewAgentBuildList

`func NewAgentBuildList() *AgentBuildList`

NewAgentBuildList instantiates a new AgentBuildList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentBuildListWithDefaults

`func NewAgentBuildListWithDefaults() *AgentBuildList`

NewAgentBuildListWithDefaults instantiates a new AgentBuildList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuilds

`func (o *AgentBuildList) GetBuilds() []AgentBuildSummary`

GetBuilds returns the Builds field if non-nil, zero value otherwise.

### GetBuildsOk

`func (o *AgentBuildList) GetBuildsOk() (*[]AgentBuildSummary, bool)`

GetBuildsOk returns a tuple with the Builds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuilds

`func (o *AgentBuildList) SetBuilds(v []AgentBuildSummary)`

SetBuilds sets Builds field to given value.

### HasBuilds

`func (o *AgentBuildList) HasBuilds() bool`

HasBuilds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


