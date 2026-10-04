# ComputePoolCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoScale** | Pointer to **bool** | AutoScale turns the provider&#39;s cluster autoscaler on for this pool. | [optional] 
**ClusterId** | Pointer to **string** | ClusterID is the cluster to add the pool to, from the URL path. | [optional] 
**Count** | Pointer to **int64** | Count is how many nodes the pool starts with. | [optional] 
**MaxNodes** | Pointer to **int64** | MaxNodes is the ceiling the autoscaler may not grow the pool past, and so the bound on what this pool can spend. Ignored unless AutoScale is set. | [optional] 
**MinNodes** | Pointer to **int64** | MinNodes is the floor the autoscaler may not shrink the pool below. Ignored unless AutoScale is set. | [optional] 
**Name** | Pointer to **string** | Name is the pool&#39;s name. | [optional] 
**Provider** | Pointer to **string** | Provider is the cloud the cluster lives on (e.g. \&quot;digitalocean\&quot;). Required — Visor routes the create by it. Accepted from the body or ?provider&#x3D;. | [optional] 
**Size** | Pointer to **string** | Size is the provider size slug for each node (e.g. \&quot;s-4vcpu-8gb\&quot;). | [optional] 

## Methods

### NewComputePoolCreate

`func NewComputePoolCreate() *ComputePoolCreate`

NewComputePoolCreate instantiates a new ComputePoolCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputePoolCreateWithDefaults

`func NewComputePoolCreateWithDefaults() *ComputePoolCreate`

NewComputePoolCreateWithDefaults instantiates a new ComputePoolCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoScale

`func (o *ComputePoolCreate) GetAutoScale() bool`

GetAutoScale returns the AutoScale field if non-nil, zero value otherwise.

### GetAutoScaleOk

`func (o *ComputePoolCreate) GetAutoScaleOk() (*bool, bool)`

GetAutoScaleOk returns a tuple with the AutoScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoScale

`func (o *ComputePoolCreate) SetAutoScale(v bool)`

SetAutoScale sets AutoScale field to given value.

### HasAutoScale

`func (o *ComputePoolCreate) HasAutoScale() bool`

HasAutoScale returns a boolean if a field has been set.

### GetClusterId

`func (o *ComputePoolCreate) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *ComputePoolCreate) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *ComputePoolCreate) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *ComputePoolCreate) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### GetCount

`func (o *ComputePoolCreate) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ComputePoolCreate) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ComputePoolCreate) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *ComputePoolCreate) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetMaxNodes

`func (o *ComputePoolCreate) GetMaxNodes() int64`

GetMaxNodes returns the MaxNodes field if non-nil, zero value otherwise.

### GetMaxNodesOk

`func (o *ComputePoolCreate) GetMaxNodesOk() (*int64, bool)`

GetMaxNodesOk returns a tuple with the MaxNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNodes

`func (o *ComputePoolCreate) SetMaxNodes(v int64)`

SetMaxNodes sets MaxNodes field to given value.

### HasMaxNodes

`func (o *ComputePoolCreate) HasMaxNodes() bool`

HasMaxNodes returns a boolean if a field has been set.

### GetMinNodes

`func (o *ComputePoolCreate) GetMinNodes() int64`

GetMinNodes returns the MinNodes field if non-nil, zero value otherwise.

### GetMinNodesOk

`func (o *ComputePoolCreate) GetMinNodesOk() (*int64, bool)`

GetMinNodesOk returns a tuple with the MinNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinNodes

`func (o *ComputePoolCreate) SetMinNodes(v int64)`

SetMinNodes sets MinNodes field to given value.

### HasMinNodes

`func (o *ComputePoolCreate) HasMinNodes() bool`

HasMinNodes returns a boolean if a field has been set.

### GetName

`func (o *ComputePoolCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ComputePoolCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ComputePoolCreate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ComputePoolCreate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProvider

`func (o *ComputePoolCreate) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ComputePoolCreate) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ComputePoolCreate) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *ComputePoolCreate) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetSize

`func (o *ComputePoolCreate) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ComputePoolCreate) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ComputePoolCreate) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *ComputePoolCreate) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


