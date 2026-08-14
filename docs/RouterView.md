# RouterView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID is the ZT edge-router&#39;s id. | [optional] 
**Name** | Pointer to **string** | Name is the edge-router&#39;s name, falling back to its id when it has none. | [optional] 
**Region** | Pointer to **string** | Region comes from a \&quot;region-&lt;slug&gt;\&quot; role attribute and is omitted when the router carries none, so the column renders \&quot;—\&quot; rather than a guess. | [optional] 
**Status** | Pointer to **string** | Status is the controller&#39;s own health signal: \&quot;online\&quot; when connected, \&quot;disabled\&quot; when administratively disabled, \&quot;offline\&quot; otherwise. | [optional] 

## Methods

### NewRouterView

`func NewRouterView() *RouterView`

NewRouterView instantiates a new RouterView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouterViewWithDefaults

`func NewRouterViewWithDefaults() *RouterView`

NewRouterViewWithDefaults instantiates a new RouterView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RouterView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RouterView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RouterView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RouterView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *RouterView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RouterView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RouterView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RouterView) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRegion

`func (o *RouterView) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *RouterView) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *RouterView) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *RouterView) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetStatus

`func (o *RouterView) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RouterView) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RouterView) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RouterView) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


