# CloudCreateClusterReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name is the cluster&#39;s name. Required. | [optional] 
**NodePool** | Pointer to [**CloudCreateClusterReqNodePool**](CloudCreateClusterReqNodePool.md) |  | [optional] 
**Region** | Pointer to **string** | Region is the provider region slug (e.g. \&quot;nyc3\&quot;). Required. | [optional] 
**Version** | Pointer to **string** | Version is the Kubernetes version slug; empty takes the provider default. | [optional] 

## Methods

### NewCloudCreateClusterReq

`func NewCloudCreateClusterReq() *CloudCreateClusterReq`

NewCloudCreateClusterReq instantiates a new CloudCreateClusterReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudCreateClusterReqWithDefaults

`func NewCloudCreateClusterReqWithDefaults() *CloudCreateClusterReq`

NewCloudCreateClusterReqWithDefaults instantiates a new CloudCreateClusterReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CloudCreateClusterReq) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudCreateClusterReq) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudCreateClusterReq) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudCreateClusterReq) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNodePool

`func (o *CloudCreateClusterReq) GetNodePool() CloudCreateClusterReqNodePool`

GetNodePool returns the NodePool field if non-nil, zero value otherwise.

### GetNodePoolOk

`func (o *CloudCreateClusterReq) GetNodePoolOk() (*CloudCreateClusterReqNodePool, bool)`

GetNodePoolOk returns a tuple with the NodePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodePool

`func (o *CloudCreateClusterReq) SetNodePool(v CloudCreateClusterReqNodePool)`

SetNodePool sets NodePool field to given value.

### HasNodePool

`func (o *CloudCreateClusterReq) HasNodePool() bool`

HasNodePool returns a boolean if a field has been set.

### GetRegion

`func (o *CloudCreateClusterReq) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CloudCreateClusterReq) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CloudCreateClusterReq) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *CloudCreateClusterReq) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetVersion

`func (o *CloudCreateClusterReq) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CloudCreateClusterReq) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CloudCreateClusterReq) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CloudCreateClusterReq) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


