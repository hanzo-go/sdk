# CloudGraphOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Degraded** | Pointer to **bool** | Degraded is true when the store was unreachable and this graph is honestly empty rather than wrong. Absent on a normal answer. | [optional] 
**Edges** | Pointer to [**[]CloudGraphEdge**](CloudGraphEdge.md) | Edges are the parent tree, the resolved wikilinks and the connector provenance. | [optional] 
**Nodes** | Pointer to [**[]CloudGraphNode**](CloudGraphNode.md) | Nodes are the pages, memories, sources, connectors and unresolved link targets. | [optional] 

## Methods

### NewCloudGraphOut

`func NewCloudGraphOut() *CloudGraphOut`

NewCloudGraphOut instantiates a new CloudGraphOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudGraphOutWithDefaults

`func NewCloudGraphOutWithDefaults() *CloudGraphOut`

NewCloudGraphOutWithDefaults instantiates a new CloudGraphOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDegraded

`func (o *CloudGraphOut) GetDegraded() bool`

GetDegraded returns the Degraded field if non-nil, zero value otherwise.

### GetDegradedOk

`func (o *CloudGraphOut) GetDegradedOk() (*bool, bool)`

GetDegradedOk returns a tuple with the Degraded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDegraded

`func (o *CloudGraphOut) SetDegraded(v bool)`

SetDegraded sets Degraded field to given value.

### HasDegraded

`func (o *CloudGraphOut) HasDegraded() bool`

HasDegraded returns a boolean if a field has been set.

### GetEdges

`func (o *CloudGraphOut) GetEdges() []CloudGraphEdge`

GetEdges returns the Edges field if non-nil, zero value otherwise.

### GetEdgesOk

`func (o *CloudGraphOut) GetEdgesOk() (*[]CloudGraphEdge, bool)`

GetEdgesOk returns a tuple with the Edges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdges

`func (o *CloudGraphOut) SetEdges(v []CloudGraphEdge)`

SetEdges sets Edges field to given value.

### HasEdges

`func (o *CloudGraphOut) HasEdges() bool`

HasEdges returns a boolean if a field has been set.

### GetNodes

`func (o *CloudGraphOut) GetNodes() []CloudGraphNode`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *CloudGraphOut) GetNodesOk() (*[]CloudGraphNode, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *CloudGraphOut) SetNodes(v []CloudGraphNode)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *CloudGraphOut) HasNodes() bool`

HasNodes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


