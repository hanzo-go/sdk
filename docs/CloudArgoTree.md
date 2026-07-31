# CloudArgoTree

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hosts** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Nodes** | Pointer to [**[]CloudArgoNode**](CloudArgoNode.md) |  | [optional] 
**OrphanedNodes** | Pointer to [**[]CloudArgoNode**](CloudArgoNode.md) |  | [optional] 

## Methods

### NewCloudArgoTree

`func NewCloudArgoTree() *CloudArgoTree`

NewCloudArgoTree instantiates a new CloudArgoTree object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudArgoTreeWithDefaults

`func NewCloudArgoTreeWithDefaults() *CloudArgoTree`

NewCloudArgoTreeWithDefaults instantiates a new CloudArgoTree object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHosts

`func (o *CloudArgoTree) GetHosts() []map[string]interface{}`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *CloudArgoTree) GetHostsOk() (*[]map[string]interface{}, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *CloudArgoTree) SetHosts(v []map[string]interface{})`

SetHosts sets Hosts field to given value.

### HasHosts

`func (o *CloudArgoTree) HasHosts() bool`

HasHosts returns a boolean if a field has been set.

### GetNodes

`func (o *CloudArgoTree) GetNodes() []CloudArgoNode`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *CloudArgoTree) GetNodesOk() (*[]CloudArgoNode, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *CloudArgoTree) SetNodes(v []CloudArgoNode)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *CloudArgoTree) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetOrphanedNodes

`func (o *CloudArgoTree) GetOrphanedNodes() []CloudArgoNode`

GetOrphanedNodes returns the OrphanedNodes field if non-nil, zero value otherwise.

### GetOrphanedNodesOk

`func (o *CloudArgoTree) GetOrphanedNodesOk() (*[]CloudArgoNode, bool)`

GetOrphanedNodesOk returns a tuple with the OrphanedNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrphanedNodes

`func (o *CloudArgoTree) SetOrphanedNodes(v []CloudArgoNode)`

SetOrphanedNodes sets OrphanedNodes field to given value.

### HasOrphanedNodes

`func (o *CloudArgoTree) HasOrphanedNodes() bool`

HasOrphanedNodes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


