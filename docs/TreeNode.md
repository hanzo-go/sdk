# TreeNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Children** | Pointer to [**[]TreeNode**](TreeNode.md) | Children is this node&#39;s direct children, each a whole node, so the array nests to the depth of the flow. A leaf carries null rather than an empty array. The subtree is materialised in full, up to 10000 nodes, out of one indexed read of the root; nothing is walked node by node. | [optional] 
**Session** | Pointer to [**SessionView**](SessionView.md) | Session is this node&#39;s own session, carrying its event count and its direct fan-out. It is the same shape the list and detail reads answer with, minus the last-event preview, which the tree does not fetch. | [optional] 

## Methods

### NewTreeNode

`func NewTreeNode() *TreeNode`

NewTreeNode instantiates a new TreeNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTreeNodeWithDefaults

`func NewTreeNodeWithDefaults() *TreeNode`

NewTreeNodeWithDefaults instantiates a new TreeNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChildren

`func (o *TreeNode) GetChildren() []TreeNode`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *TreeNode) GetChildrenOk() (*[]TreeNode, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *TreeNode) SetChildren(v []TreeNode)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *TreeNode) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### GetSession

`func (o *TreeNode) GetSession() SessionView`

GetSession returns the Session field if non-nil, zero value otherwise.

### GetSessionOk

`func (o *TreeNode) GetSessionOk() (*SessionView, bool)`

GetSessionOk returns a tuple with the Session field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSession

`func (o *TreeNode) SetSession(v SessionView)`

SetSession sets Session field to given value.

### HasSession

`func (o *TreeNode) HasSession() bool`

HasSession returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


