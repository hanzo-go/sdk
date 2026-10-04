# ComputeClusterList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Clusters** | Pointer to [**[]ComputeClusterView**](ComputeClusterView.md) | Clusters is the merged fleet — kind \&quot;managed\&quot; for Visor-provisioned, \&quot;byo\&quot; for an attached kubeconfig. | [optional] 
**Degraded** | Pointer to [**[]ComputeSourceFailure**](ComputeSourceFailure.md) | Degraded names any source that did not answer, so an empty Clusters means \&quot;you have none\&quot; only when this is absent. Omitted when everything answered, so a healthy response is unchanged. See degraded.go. | [optional] 

## Methods

### NewComputeClusterList

`func NewComputeClusterList() *ComputeClusterList`

NewComputeClusterList instantiates a new ComputeClusterList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputeClusterListWithDefaults

`func NewComputeClusterListWithDefaults() *ComputeClusterList`

NewComputeClusterListWithDefaults instantiates a new ComputeClusterList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusters

`func (o *ComputeClusterList) GetClusters() []ComputeClusterView`

GetClusters returns the Clusters field if non-nil, zero value otherwise.

### GetClustersOk

`func (o *ComputeClusterList) GetClustersOk() (*[]ComputeClusterView, bool)`

GetClustersOk returns a tuple with the Clusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusters

`func (o *ComputeClusterList) SetClusters(v []ComputeClusterView)`

SetClusters sets Clusters field to given value.

### HasClusters

`func (o *ComputeClusterList) HasClusters() bool`

HasClusters returns a boolean if a field has been set.

### GetDegraded

`func (o *ComputeClusterList) GetDegraded() []ComputeSourceFailure`

GetDegraded returns the Degraded field if non-nil, zero value otherwise.

### GetDegradedOk

`func (o *ComputeClusterList) GetDegradedOk() (*[]ComputeSourceFailure, bool)`

GetDegradedOk returns a tuple with the Degraded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDegraded

`func (o *ComputeClusterList) SetDegraded(v []ComputeSourceFailure)`

SetDegraded sets Degraded field to given value.

### HasDegraded

`func (o *ComputeClusterList) HasDegraded() bool`

HasDegraded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


