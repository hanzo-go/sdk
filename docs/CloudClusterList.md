# CloudClusterList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Clusters** | Pointer to [**[]CloudClusterView**](CloudClusterView.md) | Clusters is the merged fleet — kind \&quot;managed\&quot; for Visor-provisioned, \&quot;byo\&quot; for an attached kubeconfig. | [optional] 

## Methods

### NewCloudClusterList

`func NewCloudClusterList() *CloudClusterList`

NewCloudClusterList instantiates a new CloudClusterList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudClusterListWithDefaults

`func NewCloudClusterListWithDefaults() *CloudClusterList`

NewCloudClusterListWithDefaults instantiates a new CloudClusterList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusters

`func (o *CloudClusterList) GetClusters() []CloudClusterView`

GetClusters returns the Clusters field if non-nil, zero value otherwise.

### GetClustersOk

`func (o *CloudClusterList) GetClustersOk() (*[]CloudClusterView, bool)`

GetClustersOk returns a tuple with the Clusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusters

`func (o *CloudClusterList) SetClusters(v []CloudClusterView)`

SetClusters sets Clusters field to given value.

### HasClusters

`func (o *CloudClusterList) HasClusters() bool`

HasClusters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


