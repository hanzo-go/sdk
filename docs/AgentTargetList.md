# AgentTargetList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Targets** | Pointer to [**[]AgentTargetView**](AgentTargetView.md) | Targets is every target registered to the caller&#39;s org. | [optional] 

## Methods

### NewAgentTargetList

`func NewAgentTargetList() *AgentTargetList`

NewAgentTargetList instantiates a new AgentTargetList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentTargetListWithDefaults

`func NewAgentTargetListWithDefaults() *AgentTargetList`

NewAgentTargetListWithDefaults instantiates a new AgentTargetList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargets

`func (o *AgentTargetList) GetTargets() []AgentTargetView`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *AgentTargetList) GetTargetsOk() (*[]AgentTargetView, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *AgentTargetList) SetTargets(v []AgentTargetView)`

SetTargets sets Targets field to given value.

### HasTargets

`func (o *AgentTargetList) HasTargets() bool`

HasTargets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


