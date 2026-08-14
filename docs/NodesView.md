# NodesView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nodes** | Pointer to [**[]NodeView**](NodeView.md) | Nodes is every node of the caller&#39;s org with a live socket to THIS replica, ordered by id. A node connected to a different replica is not in it. | [optional] 

## Methods

### NewNodesView

`func NewNodesView() *NodesView`

NewNodesView instantiates a new NodesView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodesViewWithDefaults

`func NewNodesViewWithDefaults() *NodesView`

NewNodesViewWithDefaults instantiates a new NodesView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodes

`func (o *NodesView) GetNodes() []NodeView`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *NodesView) GetNodesOk() (*[]NodeView, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *NodesView) SetNodes(v []NodeView)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *NodesView) HasNodes() bool`

HasNodes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


