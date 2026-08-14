# CreateClusterReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name is the cluster&#39;s name. Required. | [optional] 
**NodePool** | Pointer to [**CreateClusterReqNodePool**](CreateClusterReqNodePool.md) |  | [optional] 
**Region** | Pointer to **string** | Region is the provider region slug (e.g. \&quot;nyc3\&quot;). Required. | [optional] 
**Version** | Pointer to **string** | Version is the Kubernetes version slug; empty takes the provider default. | [optional] 

## Methods

### NewCreateClusterReq

`func NewCreateClusterReq() *CreateClusterReq`

NewCreateClusterReq instantiates a new CreateClusterReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateClusterReqWithDefaults

`func NewCreateClusterReqWithDefaults() *CreateClusterReq`

NewCreateClusterReqWithDefaults instantiates a new CreateClusterReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateClusterReq) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateClusterReq) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateClusterReq) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateClusterReq) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNodePool

`func (o *CreateClusterReq) GetNodePool() CreateClusterReqNodePool`

GetNodePool returns the NodePool field if non-nil, zero value otherwise.

### GetNodePoolOk

`func (o *CreateClusterReq) GetNodePoolOk() (*CreateClusterReqNodePool, bool)`

GetNodePoolOk returns a tuple with the NodePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodePool

`func (o *CreateClusterReq) SetNodePool(v CreateClusterReqNodePool)`

SetNodePool sets NodePool field to given value.

### HasNodePool

`func (o *CreateClusterReq) HasNodePool() bool`

HasNodePool returns a boolean if a field has been set.

### GetRegion

`func (o *CreateClusterReq) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CreateClusterReq) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CreateClusterReq) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *CreateClusterReq) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetVersion

`func (o *CreateClusterReq) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CreateClusterReq) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CreateClusterReq) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CreateClusterReq) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


