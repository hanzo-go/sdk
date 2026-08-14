# ArgoClusterList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]ArgoCluster**](ArgoCluster.md) |  | [optional] 
**Metadata** | Pointer to [**ArgoListMeta**](ArgoListMeta.md) |  | [optional] 

## Methods

### NewArgoClusterList

`func NewArgoClusterList() *ArgoClusterList`

NewArgoClusterList instantiates a new ArgoClusterList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArgoClusterListWithDefaults

`func NewArgoClusterListWithDefaults() *ArgoClusterList`

NewArgoClusterListWithDefaults instantiates a new ArgoClusterList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *ArgoClusterList) GetItems() []ArgoCluster`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ArgoClusterList) GetItemsOk() (*[]ArgoCluster, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ArgoClusterList) SetItems(v []ArgoCluster)`

SetItems sets Items field to given value.

### HasItems

`func (o *ArgoClusterList) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetMetadata

`func (o *ArgoClusterList) GetMetadata() ArgoListMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ArgoClusterList) GetMetadataOk() (*ArgoListMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ArgoClusterList) SetMetadata(v ArgoListMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ArgoClusterList) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


