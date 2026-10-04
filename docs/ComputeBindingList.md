# ComputeBindingList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentBindings** | Pointer to [**[]ComputeAgentBinding**](ComputeAgentBinding.md) | AgentBindings is one row per bound machine, emitted verbatim as vm reports it. | [optional] 

## Methods

### NewComputeBindingList

`func NewComputeBindingList() *ComputeBindingList`

NewComputeBindingList instantiates a new ComputeBindingList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputeBindingListWithDefaults

`func NewComputeBindingListWithDefaults() *ComputeBindingList`

NewComputeBindingListWithDefaults instantiates a new ComputeBindingList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentBindings

`func (o *ComputeBindingList) GetAgentBindings() []ComputeAgentBinding`

GetAgentBindings returns the AgentBindings field if non-nil, zero value otherwise.

### GetAgentBindingsOk

`func (o *ComputeBindingList) GetAgentBindingsOk() (*[]ComputeAgentBinding, bool)`

GetAgentBindingsOk returns a tuple with the AgentBindings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentBindings

`func (o *ComputeBindingList) SetAgentBindings(v []ComputeAgentBinding)`

SetAgentBindings sets AgentBindings field to given value.

### HasAgentBindings

`func (o *ComputeBindingList) HasAgentBindings() bool`

HasAgentBindings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


