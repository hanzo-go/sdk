# DeployArgoClusterList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]DeployArgoCluster**](DeployArgoCluster.md) | Items is one entry per distinct destination server, in first-seen order with the in-cluster destination first. Never empty: an empty fleet still has the one cluster it would deploy into. | [optional] 
**Metadata** | Pointer to [**DeployArgoListMeta**](DeployArgoListMeta.md) | Metadata is the list envelope the SPA expects; it carries no resume point. | [optional] 

## Methods

### NewDeployArgoClusterList

`func NewDeployArgoClusterList() *DeployArgoClusterList`

NewDeployArgoClusterList instantiates a new DeployArgoClusterList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeployArgoClusterListWithDefaults

`func NewDeployArgoClusterListWithDefaults() *DeployArgoClusterList`

NewDeployArgoClusterListWithDefaults instantiates a new DeployArgoClusterList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *DeployArgoClusterList) GetItems() []DeployArgoCluster`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *DeployArgoClusterList) GetItemsOk() (*[]DeployArgoCluster, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *DeployArgoClusterList) SetItems(v []DeployArgoCluster)`

SetItems sets Items field to given value.

### HasItems

`func (o *DeployArgoClusterList) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetMetadata

`func (o *DeployArgoClusterList) GetMetadata() DeployArgoListMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DeployArgoClusterList) GetMetadataOk() (*DeployArgoListMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DeployArgoClusterList) SetMetadata(v DeployArgoListMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DeployArgoClusterList) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


