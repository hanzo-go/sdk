# PoolCreate

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

### NewPoolCreate

`func NewPoolCreate() *PoolCreate`

NewPoolCreate instantiates a new PoolCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoolCreateWithDefaults

`func NewPoolCreateWithDefaults() *PoolCreate`

NewPoolCreateWithDefaults instantiates a new PoolCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoScale

`func (o *PoolCreate) GetAutoScale() bool`

GetAutoScale returns the AutoScale field if non-nil, zero value otherwise.

### GetAutoScaleOk

`func (o *PoolCreate) GetAutoScaleOk() (*bool, bool)`

GetAutoScaleOk returns a tuple with the AutoScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoScale

`func (o *PoolCreate) SetAutoScale(v bool)`

SetAutoScale sets AutoScale field to given value.

### HasAutoScale

`func (o *PoolCreate) HasAutoScale() bool`

HasAutoScale returns a boolean if a field has been set.

### GetClusterId

`func (o *PoolCreate) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *PoolCreate) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *PoolCreate) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *PoolCreate) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### GetCount

`func (o *PoolCreate) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PoolCreate) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PoolCreate) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *PoolCreate) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetMaxNodes

`func (o *PoolCreate) GetMaxNodes() int32`

GetMaxNodes returns the MaxNodes field if non-nil, zero value otherwise.

### GetMaxNodesOk

`func (o *PoolCreate) GetMaxNodesOk() (*int32, bool)`

GetMaxNodesOk returns a tuple with the MaxNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNodes

`func (o *PoolCreate) SetMaxNodes(v int32)`

SetMaxNodes sets MaxNodes field to given value.

### HasMaxNodes

`func (o *PoolCreate) HasMaxNodes() bool`

HasMaxNodes returns a boolean if a field has been set.

### GetMinNodes

`func (o *PoolCreate) GetMinNodes() int32`

GetMinNodes returns the MinNodes field if non-nil, zero value otherwise.

### GetMinNodesOk

`func (o *PoolCreate) GetMinNodesOk() (*int32, bool)`

GetMinNodesOk returns a tuple with the MinNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinNodes

`func (o *PoolCreate) SetMinNodes(v int32)`

SetMinNodes sets MinNodes field to given value.

### HasMinNodes

`func (o *PoolCreate) HasMinNodes() bool`

HasMinNodes returns a boolean if a field has been set.

### GetName

`func (o *PoolCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PoolCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PoolCreate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PoolCreate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProvider

`func (o *PoolCreate) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *PoolCreate) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *PoolCreate) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *PoolCreate) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetSize

`func (o *PoolCreate) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *PoolCreate) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *PoolCreate) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *PoolCreate) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


