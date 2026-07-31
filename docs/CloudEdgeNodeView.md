# CloudEdgeNodeView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID is the ZT edge-router&#39;s id. | [optional] 
**Name** | Pointer to **string** | Name is the edge-router&#39;s name, falling back to its id when it has none. | [optional] 
**Region** | Pointer to **string** | Region comes from a \&quot;region-&lt;slug&gt;\&quot; role attribute and is omitted when the router carries none, so the column renders \&quot;—\&quot; rather than a guess. | [optional] 
**Status** | Pointer to **string** | Status is the controller&#39;s own health signal: \&quot;online\&quot; when connected, \&quot;disabled\&quot; when administratively disabled, \&quot;offline\&quot; otherwise. | [optional] 

## Methods

### NewCloudEdgeNodeView

`func NewCloudEdgeNodeView() *CloudEdgeNodeView`

NewCloudEdgeNodeView instantiates a new CloudEdgeNodeView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudEdgeNodeViewWithDefaults

`func NewCloudEdgeNodeViewWithDefaults() *CloudEdgeNodeView`

NewCloudEdgeNodeViewWithDefaults instantiates a new CloudEdgeNodeView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CloudEdgeNodeView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudEdgeNodeView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudEdgeNodeView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudEdgeNodeView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CloudEdgeNodeView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudEdgeNodeView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudEdgeNodeView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudEdgeNodeView) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRegion

`func (o *CloudEdgeNodeView) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CloudEdgeNodeView) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CloudEdgeNodeView) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *CloudEdgeNodeView) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetStatus

`func (o *CloudEdgeNodeView) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudEdgeNodeView) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudEdgeNodeView) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudEdgeNodeView) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


