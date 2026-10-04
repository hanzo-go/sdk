# DeployArgoTree

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hosts** | Pointer to **[]interface{}** | Hosts is ArgoCD&#39;s per-node machine inventory. Always empty: this plane projects applications and serves no cluster-node view. | [optional] 
**Nodes** | Pointer to [**[]DeployArgoNode**](DeployArgoNode.md) | Nodes is the FLAT node list, root first: the App CR, then the objects the operator owns, then their ReplicaSets and Pods. The hierarchy is in ParentRefs, not in the ordering. | [optional] 
**OrphanedNodes** | Pointer to [**[]DeployArgoNode**](DeployArgoNode.md) | OrphanedNodes are objects in the namespace belonging to no application. Always empty: this walk reaches an object only THROUGH ownership from the App CR, so it can never hold one that is orphaned. | [optional] 

## Methods

### NewDeployArgoTree

`func NewDeployArgoTree() *DeployArgoTree`

NewDeployArgoTree instantiates a new DeployArgoTree object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeployArgoTreeWithDefaults

`func NewDeployArgoTreeWithDefaults() *DeployArgoTree`

NewDeployArgoTreeWithDefaults instantiates a new DeployArgoTree object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHosts

`func (o *DeployArgoTree) GetHosts() []interface{}`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *DeployArgoTree) GetHostsOk() (*[]interface{}, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *DeployArgoTree) SetHosts(v []interface{})`

SetHosts sets Hosts field to given value.

### HasHosts

`func (o *DeployArgoTree) HasHosts() bool`

HasHosts returns a boolean if a field has been set.

### GetNodes

`func (o *DeployArgoTree) GetNodes() []DeployArgoNode`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *DeployArgoTree) GetNodesOk() (*[]DeployArgoNode, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *DeployArgoTree) SetNodes(v []DeployArgoNode)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *DeployArgoTree) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetOrphanedNodes

`func (o *DeployArgoTree) GetOrphanedNodes() []DeployArgoNode`

GetOrphanedNodes returns the OrphanedNodes field if non-nil, zero value otherwise.

### GetOrphanedNodesOk

`func (o *DeployArgoTree) GetOrphanedNodesOk() (*[]DeployArgoNode, bool)`

GetOrphanedNodesOk returns a tuple with the OrphanedNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrphanedNodes

`func (o *DeployArgoTree) SetOrphanedNodes(v []DeployArgoNode)`

SetOrphanedNodes sets OrphanedNodes field to given value.

### HasOrphanedNodes

`func (o *DeployArgoTree) HasOrphanedNodes() bool`

HasOrphanedNodes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


