# GraphEdge

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | Pointer to **string** | From is the id of the node the edge starts at: the child page for a parent edge, the page holding the wikilink for a link edge, the kb.source for a provenance edge. Always one of Nodes. | [optional] 
**Kind** | Pointer to **string** | parent | link | provenance | [optional] 
**To** | Pointer to **string** | To is the id of the node the edge points at: the parent page, the linked page, the kb.connector. Always one of Nodes — a wikilink matching no page points at a synthetic \&quot;unresolved:&lt;lowercased title&gt;\&quot; node rather than dangling. | [optional] 

## Methods

### NewGraphEdge

`func NewGraphEdge() *GraphEdge`

NewGraphEdge instantiates a new GraphEdge object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphEdgeWithDefaults

`func NewGraphEdgeWithDefaults() *GraphEdge`

NewGraphEdgeWithDefaults instantiates a new GraphEdge object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *GraphEdge) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *GraphEdge) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *GraphEdge) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *GraphEdge) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetKind

`func (o *GraphEdge) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *GraphEdge) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *GraphEdge) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *GraphEdge) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetTo

`func (o *GraphEdge) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *GraphEdge) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *GraphEdge) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *GraphEdge) HasTo() bool`

HasTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


