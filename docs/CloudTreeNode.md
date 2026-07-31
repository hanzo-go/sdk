# CloudTreeNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Children** | Pointer to [**[]CloudTreeNode**](CloudTreeNode.md) |  | [optional] 
**Session** | Pointer to [**CloudSessionView**](CloudSessionView.md) |  | [optional] 

## Methods

### NewCloudTreeNode

`func NewCloudTreeNode() *CloudTreeNode`

NewCloudTreeNode instantiates a new CloudTreeNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudTreeNodeWithDefaults

`func NewCloudTreeNodeWithDefaults() *CloudTreeNode`

NewCloudTreeNodeWithDefaults instantiates a new CloudTreeNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChildren

`func (o *CloudTreeNode) GetChildren() []CloudTreeNode`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *CloudTreeNode) GetChildrenOk() (*[]CloudTreeNode, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *CloudTreeNode) SetChildren(v []CloudTreeNode)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *CloudTreeNode) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### GetSession

`func (o *CloudTreeNode) GetSession() CloudSessionView`

GetSession returns the Session field if non-nil, zero value otherwise.

### GetSessionOk

`func (o *CloudTreeNode) GetSessionOk() (*CloudSessionView, bool)`

GetSessionOk returns a tuple with the Session field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSession

`func (o *CloudTreeNode) SetSession(v CloudSessionView)`

SetSession sets Session field to given value.

### HasSession

`func (o *CloudTreeNode) HasSession() bool`

HasSession returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


