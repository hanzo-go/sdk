# NodeList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nodes** | Pointer to [**[]MachineView**](MachineView.md) | Nodes is one row per worker node, in the SAME machineView shape the machines surface emits — a node IS a machine. | [optional] 

## Methods

### NewNodeList

`func NewNodeList() *NodeList`

NewNodeList instantiates a new NodeList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeListWithDefaults

`func NewNodeListWithDefaults() *NodeList`

NewNodeListWithDefaults instantiates a new NodeList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodes

`func (o *NodeList) GetNodes() []MachineView`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *NodeList) GetNodesOk() (*[]MachineView, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *NodeList) SetNodes(v []MachineView)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *NodeList) HasNodes() bool`

HasNodes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


