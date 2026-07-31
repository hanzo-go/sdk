# CloudGraphNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | \&quot;&lt;doctype&gt;:&lt;name&gt;\&quot; — globally unique, click-to-open key | [optional] 
**Name** | Pointer to **string** | the document name (empty for synthetic nodes) | [optional] 
**Project** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** | display label | [optional] 
**Type** | Pointer to **string** | kb-page | kb-memory | kb-source | kb-connector | unresolved | [optional] 

## Methods

### NewCloudGraphNode

`func NewCloudGraphNode() *CloudGraphNode`

NewCloudGraphNode instantiates a new CloudGraphNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudGraphNodeWithDefaults

`func NewCloudGraphNodeWithDefaults() *CloudGraphNode`

NewCloudGraphNodeWithDefaults instantiates a new CloudGraphNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CloudGraphNode) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudGraphNode) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudGraphNode) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudGraphNode) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CloudGraphNode) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudGraphNode) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudGraphNode) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudGraphNode) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProject

`func (o *CloudGraphNode) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *CloudGraphNode) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *CloudGraphNode) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *CloudGraphNode) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetTitle

`func (o *CloudGraphNode) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CloudGraphNode) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CloudGraphNode) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CloudGraphNode) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetType

`func (o *CloudGraphNode) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CloudGraphNode) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CloudGraphNode) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CloudGraphNode) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


