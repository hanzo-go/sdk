# ArgoTree

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hosts** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Nodes** | Pointer to [**[]ArgoNode**](ArgoNode.md) |  | [optional] 
**OrphanedNodes** | Pointer to [**[]ArgoNode**](ArgoNode.md) |  | [optional] 

## Methods

### NewArgoTree

`func NewArgoTree() *ArgoTree`

NewArgoTree instantiates a new ArgoTree object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArgoTreeWithDefaults

`func NewArgoTreeWithDefaults() *ArgoTree`

NewArgoTreeWithDefaults instantiates a new ArgoTree object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHosts

`func (o *ArgoTree) GetHosts() []map[string]interface{}`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *ArgoTree) GetHostsOk() (*[]map[string]interface{}, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *ArgoTree) SetHosts(v []map[string]interface{})`

SetHosts sets Hosts field to given value.

### HasHosts

`func (o *ArgoTree) HasHosts() bool`

HasHosts returns a boolean if a field has been set.

### GetNodes

`func (o *ArgoTree) GetNodes() []ArgoNode`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *ArgoTree) GetNodesOk() (*[]ArgoNode, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *ArgoTree) SetNodes(v []ArgoNode)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *ArgoTree) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetOrphanedNodes

`func (o *ArgoTree) GetOrphanedNodes() []ArgoNode`

GetOrphanedNodes returns the OrphanedNodes field if non-nil, zero value otherwise.

### GetOrphanedNodesOk

`func (o *ArgoTree) GetOrphanedNodesOk() (*[]ArgoNode, bool)`

GetOrphanedNodesOk returns a tuple with the OrphanedNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrphanedNodes

`func (o *ArgoTree) SetOrphanedNodes(v []ArgoNode)`

SetOrphanedNodes sets OrphanedNodes field to given value.

### HasOrphanedNodes

`func (o *ArgoTree) HasOrphanedNodes() bool`

HasOrphanedNodes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


