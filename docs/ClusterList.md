# ClusterList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Clusters** | Pointer to [**[]ClusterView**](ClusterView.md) | Clusters is the merged fleet — kind \&quot;managed\&quot; for Visor-provisioned, \&quot;byo\&quot; for an attached kubeconfig. | [optional] 
**Degraded** | Pointer to [**[]SourceFailure**](SourceFailure.md) | Degraded names any source that did not answer, so an empty Clusters means \&quot;you have none\&quot; only when this is absent. Omitted when everything answered, so a healthy response is unchanged. See degraded.go. | [optional] 

## Methods

### NewClusterList

`func NewClusterList() *ClusterList`

NewClusterList instantiates a new ClusterList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterListWithDefaults

`func NewClusterListWithDefaults() *ClusterList`

NewClusterListWithDefaults instantiates a new ClusterList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusters

`func (o *ClusterList) GetClusters() []ClusterView`

GetClusters returns the Clusters field if non-nil, zero value otherwise.

### GetClustersOk

`func (o *ClusterList) GetClustersOk() (*[]ClusterView, bool)`

GetClustersOk returns a tuple with the Clusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusters

`func (o *ClusterList) SetClusters(v []ClusterView)`

SetClusters sets Clusters field to given value.

### HasClusters

`func (o *ClusterList) HasClusters() bool`

HasClusters returns a boolean if a field has been set.

### GetDegraded

`func (o *ClusterList) GetDegraded() []SourceFailure`

GetDegraded returns the Degraded field if non-nil, zero value otherwise.

### GetDegradedOk

`func (o *ClusterList) GetDegradedOk() (*[]SourceFailure, bool)`

GetDegradedOk returns a tuple with the Degraded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDegraded

`func (o *ClusterList) SetDegraded(v []SourceFailure)`

SetDegraded sets Degraded field to given value.

### HasDegraded

`func (o *ClusterList) HasDegraded() bool`

HasDegraded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


