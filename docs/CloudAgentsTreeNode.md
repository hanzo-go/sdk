# CloudAgentsTreeNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Session** | Pointer to [**CloudAgentsSession**](CloudAgentsSession.md) |  | [optional] 
**Children** | Pointer to [**[]CloudAgentsTreeNode**](CloudAgentsTreeNode.md) |  | [optional] 

## Methods

### NewCloudAgentsTreeNode

`func NewCloudAgentsTreeNode() *CloudAgentsTreeNode`

NewCloudAgentsTreeNode instantiates a new CloudAgentsTreeNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudAgentsTreeNodeWithDefaults

`func NewCloudAgentsTreeNodeWithDefaults() *CloudAgentsTreeNode`

NewCloudAgentsTreeNodeWithDefaults instantiates a new CloudAgentsTreeNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSession

`func (o *CloudAgentsTreeNode) GetSession() CloudAgentsSession`

GetSession returns the Session field if non-nil, zero value otherwise.

### GetSessionOk

`func (o *CloudAgentsTreeNode) GetSessionOk() (*CloudAgentsSession, bool)`

GetSessionOk returns a tuple with the Session field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSession

`func (o *CloudAgentsTreeNode) SetSession(v CloudAgentsSession)`

SetSession sets Session field to given value.

### HasSession

`func (o *CloudAgentsTreeNode) HasSession() bool`

HasSession returns a boolean if a field has been set.

### GetChildren

`func (o *CloudAgentsTreeNode) GetChildren() []CloudAgentsTreeNode`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *CloudAgentsTreeNode) GetChildrenOk() (*[]CloudAgentsTreeNode, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *CloudAgentsTreeNode) SetChildren(v []CloudAgentsTreeNode)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *CloudAgentsTreeNode) HasChildren() bool`

HasChildren returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


