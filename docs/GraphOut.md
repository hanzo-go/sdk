# GraphOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Degraded** | Pointer to **bool** | Degraded is true when the store was unreachable and this graph is honestly empty rather than wrong. Absent on a normal answer. | [optional] 
**Edges** | Pointer to [**[]GraphEdge**](GraphEdge.md) | Edges are the parent tree, the resolved wikilinks and the connector provenance. | [optional] 
**Nodes** | Pointer to [**[]GraphNode**](GraphNode.md) | Nodes are the pages, memories, sources, connectors and unresolved link targets. | [optional] 

## Methods

### NewGraphOut

`func NewGraphOut() *GraphOut`

NewGraphOut instantiates a new GraphOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphOutWithDefaults

`func NewGraphOutWithDefaults() *GraphOut`

NewGraphOutWithDefaults instantiates a new GraphOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDegraded

`func (o *GraphOut) GetDegraded() bool`

GetDegraded returns the Degraded field if non-nil, zero value otherwise.

### GetDegradedOk

`func (o *GraphOut) GetDegradedOk() (*bool, bool)`

GetDegradedOk returns a tuple with the Degraded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDegraded

`func (o *GraphOut) SetDegraded(v bool)`

SetDegraded sets Degraded field to given value.

### HasDegraded

`func (o *GraphOut) HasDegraded() bool`

HasDegraded returns a boolean if a field has been set.

### GetEdges

`func (o *GraphOut) GetEdges() []GraphEdge`

GetEdges returns the Edges field if non-nil, zero value otherwise.

### GetEdgesOk

`func (o *GraphOut) GetEdgesOk() (*[]GraphEdge, bool)`

GetEdgesOk returns a tuple with the Edges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdges

`func (o *GraphOut) SetEdges(v []GraphEdge)`

SetEdges sets Edges field to given value.

### HasEdges

`func (o *GraphOut) HasEdges() bool`

HasEdges returns a boolean if a field has been set.

### GetNodes

`func (o *GraphOut) GetNodes() []GraphNode`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *GraphOut) GetNodesOk() (*[]GraphNode, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *GraphOut) SetNodes(v []GraphNode)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *GraphOut) HasNodes() bool`

HasNodes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


