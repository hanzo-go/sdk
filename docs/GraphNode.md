# GraphNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | \&quot;&lt;doctype&gt;:&lt;name&gt;\&quot; — globally unique, click-to-open key | [optional] 
**Name** | Pointer to **string** | the document name (empty for synthetic nodes) | [optional] 
**Project** | Pointer to **string** | Project is the project scope the underlying document was saved under. Absent for a document saved with none, and for the synthetic nodes — unresolved link targets and connectors belong to no project. When ?project&#x3D; narrows the graph, every page, memory and source node carries that value. | [optional] 
**Title** | Pointer to **string** | display label | [optional] 
**Type** | Pointer to **string** | kb.page | kb.memory | kb.source | kb.connector | unresolved | [optional] 

## Methods

### NewGraphNode

`func NewGraphNode() *GraphNode`

NewGraphNode instantiates a new GraphNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphNodeWithDefaults

`func NewGraphNodeWithDefaults() *GraphNode`

NewGraphNodeWithDefaults instantiates a new GraphNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GraphNode) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GraphNode) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GraphNode) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GraphNode) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GraphNode) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GraphNode) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GraphNode) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GraphNode) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProject

`func (o *GraphNode) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *GraphNode) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *GraphNode) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *GraphNode) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetTitle

`func (o *GraphNode) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *GraphNode) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *GraphNode) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *GraphNode) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetType

`func (o *GraphNode) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GraphNode) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GraphNode) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GraphNode) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


