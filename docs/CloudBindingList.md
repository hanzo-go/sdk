# CloudBindingList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentBindings** | Pointer to [**[]CloudAgentBinding**](CloudAgentBinding.md) | AgentBindings is one row per bound machine, emitted verbatim as vm reports it so vm stays the single source of truth for the binding shape. | [optional] 

## Methods

### NewCloudBindingList

`func NewCloudBindingList() *CloudBindingList`

NewCloudBindingList instantiates a new CloudBindingList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudBindingListWithDefaults

`func NewCloudBindingListWithDefaults() *CloudBindingList`

NewCloudBindingListWithDefaults instantiates a new CloudBindingList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentBindings

`func (o *CloudBindingList) GetAgentBindings() []CloudAgentBinding`

GetAgentBindings returns the AgentBindings field if non-nil, zero value otherwise.

### GetAgentBindingsOk

`func (o *CloudBindingList) GetAgentBindingsOk() (*[]CloudAgentBinding, bool)`

GetAgentBindingsOk returns a tuple with the AgentBindings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentBindings

`func (o *CloudBindingList) SetAgentBindings(v []CloudAgentBinding)`

SetAgentBindings sets AgentBindings field to given value.

### HasAgentBindings

`func (o *CloudBindingList) HasAgentBindings() bool`

HasAgentBindings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


