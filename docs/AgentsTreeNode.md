# AgentsTreeNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Session** | Pointer to [**AgentsSessionView**](AgentsSessionView.md) |  | [optional] 
**Children** | Pointer to [**[]AgentsTreeNode**](AgentsTreeNode.md) |  | [optional] 

## Methods

### NewAgentsTreeNode

`func NewAgentsTreeNode() *AgentsTreeNode`

NewAgentsTreeNode instantiates a new AgentsTreeNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentsTreeNodeWithDefaults

`func NewAgentsTreeNodeWithDefaults() *AgentsTreeNode`

NewAgentsTreeNodeWithDefaults instantiates a new AgentsTreeNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSession

`func (o *AgentsTreeNode) GetSession() AgentsSessionView`

GetSession returns the Session field if non-nil, zero value otherwise.

### GetSessionOk

`func (o *AgentsTreeNode) GetSessionOk() (*AgentsSessionView, bool)`

GetSessionOk returns a tuple with the Session field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSession

`func (o *AgentsTreeNode) SetSession(v AgentsSessionView)`

SetSession sets Session field to given value.

### HasSession

`func (o *AgentsTreeNode) HasSession() bool`

HasSession returns a boolean if a field has been set.

### GetChildren

`func (o *AgentsTreeNode) GetChildren() []AgentsTreeNode`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *AgentsTreeNode) GetChildrenOk() (*[]AgentsTreeNode, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *AgentsTreeNode) SetChildren(v []AgentsTreeNode)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *AgentsTreeNode) HasChildren() bool`

HasChildren returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


