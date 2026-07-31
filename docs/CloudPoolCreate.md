# CloudPoolCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoScale** | Pointer to **bool** | AutoScale turns the provider&#39;s cluster autoscaler on for this pool. | [optional] 
**ClusterId** | Pointer to **string** | ClusterID is the cluster to add the pool to, from the URL path. | [optional] 
**Count** | Pointer to **int32** | Count is how many nodes the pool starts with. | [optional] 
**MaxNodes** | Pointer to **int32** |  | [optional] 
**MinNodes** | Pointer to **int32** | MinNodes and MaxNodes bound the autoscaler; they are ignored unless AutoScale is set. | [optional] 
**Name** | Pointer to **string** | Name is the pool&#39;s name. | [optional] 
**Provider** | Pointer to **string** | Provider is the cloud the cluster lives on (e.g. \&quot;digitalocean\&quot;). Required — Visor routes the create by it. Accepted from the body or ?provider&#x3D;. | [optional] 
**Size** | Pointer to **string** | Size is the provider size slug for each node (e.g. \&quot;s-4vcpu-8gb\&quot;). | [optional] 

## Methods

### NewCloudPoolCreate

`func NewCloudPoolCreate() *CloudPoolCreate`

NewCloudPoolCreate instantiates a new CloudPoolCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudPoolCreateWithDefaults

`func NewCloudPoolCreateWithDefaults() *CloudPoolCreate`

NewCloudPoolCreateWithDefaults instantiates a new CloudPoolCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoScale

`func (o *CloudPoolCreate) GetAutoScale() bool`

GetAutoScale returns the AutoScale field if non-nil, zero value otherwise.

### GetAutoScaleOk

`func (o *CloudPoolCreate) GetAutoScaleOk() (*bool, bool)`

GetAutoScaleOk returns a tuple with the AutoScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoScale

`func (o *CloudPoolCreate) SetAutoScale(v bool)`

SetAutoScale sets AutoScale field to given value.

### HasAutoScale

`func (o *CloudPoolCreate) HasAutoScale() bool`

HasAutoScale returns a boolean if a field has been set.

### GetClusterId

`func (o *CloudPoolCreate) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *CloudPoolCreate) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *CloudPoolCreate) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *CloudPoolCreate) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### GetCount

`func (o *CloudPoolCreate) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *CloudPoolCreate) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *CloudPoolCreate) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *CloudPoolCreate) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetMaxNodes

`func (o *CloudPoolCreate) GetMaxNodes() int32`

GetMaxNodes returns the MaxNodes field if non-nil, zero value otherwise.

### GetMaxNodesOk

`func (o *CloudPoolCreate) GetMaxNodesOk() (*int32, bool)`

GetMaxNodesOk returns a tuple with the MaxNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNodes

`func (o *CloudPoolCreate) SetMaxNodes(v int32)`

SetMaxNodes sets MaxNodes field to given value.

### HasMaxNodes

`func (o *CloudPoolCreate) HasMaxNodes() bool`

HasMaxNodes returns a boolean if a field has been set.

### GetMinNodes

`func (o *CloudPoolCreate) GetMinNodes() int32`

GetMinNodes returns the MinNodes field if non-nil, zero value otherwise.

### GetMinNodesOk

`func (o *CloudPoolCreate) GetMinNodesOk() (*int32, bool)`

GetMinNodesOk returns a tuple with the MinNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinNodes

`func (o *CloudPoolCreate) SetMinNodes(v int32)`

SetMinNodes sets MinNodes field to given value.

### HasMinNodes

`func (o *CloudPoolCreate) HasMinNodes() bool`

HasMinNodes returns a boolean if a field has been set.

### GetName

`func (o *CloudPoolCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudPoolCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudPoolCreate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudPoolCreate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProvider

`func (o *CloudPoolCreate) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *CloudPoolCreate) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *CloudPoolCreate) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *CloudPoolCreate) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetSize

`func (o *CloudPoolCreate) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *CloudPoolCreate) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *CloudPoolCreate) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *CloudPoolCreate) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


