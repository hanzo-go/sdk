# BindingList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentBindings** | Pointer to [**[]AgentBinding**](AgentBinding.md) | AgentBindings is one row per bound machine, emitted verbatim as vm reports it. | [optional] 

## Methods

### NewBindingList

`func NewBindingList() *BindingList`

NewBindingList instantiates a new BindingList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBindingListWithDefaults

`func NewBindingListWithDefaults() *BindingList`

NewBindingListWithDefaults instantiates a new BindingList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentBindings

`func (o *BindingList) GetAgentBindings() []AgentBinding`

GetAgentBindings returns the AgentBindings field if non-nil, zero value otherwise.

### GetAgentBindingsOk

`func (o *BindingList) GetAgentBindingsOk() (*[]AgentBinding, bool)`

GetAgentBindingsOk returns a tuple with the AgentBindings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentBindings

`func (o *BindingList) SetAgentBindings(v []AgentBinding)`

SetAgentBindings sets AgentBindings field to given value.

### HasAgentBindings

`func (o *BindingList) HasAgentBindings() bool`

HasAgentBindings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


