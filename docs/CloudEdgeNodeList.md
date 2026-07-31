# CloudEdgeNodeList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nodes** | Pointer to [**[]CloudEdgeNodeView**](CloudEdgeNodeView.md) | Nodes is one row per ZT edge-router tagged with the caller&#39;s org role. | [optional] 

## Methods

### NewCloudEdgeNodeList

`func NewCloudEdgeNodeList() *CloudEdgeNodeList`

NewCloudEdgeNodeList instantiates a new CloudEdgeNodeList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudEdgeNodeListWithDefaults

`func NewCloudEdgeNodeListWithDefaults() *CloudEdgeNodeList`

NewCloudEdgeNodeListWithDefaults instantiates a new CloudEdgeNodeList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodes

`func (o *CloudEdgeNodeList) GetNodes() []CloudEdgeNodeView`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *CloudEdgeNodeList) GetNodesOk() (*[]CloudEdgeNodeView, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *CloudEdgeNodeList) SetNodes(v []CloudEdgeNodeView)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *CloudEdgeNodeList) HasNodes() bool`

HasNodes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


