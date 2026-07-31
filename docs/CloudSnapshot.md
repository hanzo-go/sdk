# CloudSnapshot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**At** | Pointer to **string** |  | [optional] 
**Clusters** | Pointer to [**[]CloudCluster**](CloudCluster.md) |  | [optional] 
**Complete** | Pointer to **bool** |  | [optional] 
**Cost** | Pointer to [**CloudCost**](CloudCost.md) |  | [optional] 
**Findings** | Pointer to [**[]CloudFinding**](CloudFinding.md) |  | [optional] 
**IncompleteReason** | Pointer to **string** |  | [optional] 
**LoadBalancers** | Pointer to [**[]CloudLoadBalancer**](CloudLoadBalancer.md) |  | [optional] 
**Nodes** | Pointer to [**[]CloudNode**](CloudNode.md) |  | [optional] 
**Sources** | Pointer to [**[]CloudSourceStatus**](CloudSourceStatus.md) |  | [optional] 
**Totals** | Pointer to [**CloudTotals**](CloudTotals.md) |  | [optional] 
**Volumes** | Pointer to [**[]CloudVolume**](CloudVolume.md) |  | [optional] 

## Methods

### NewCloudSnapshot

`func NewCloudSnapshot() *CloudSnapshot`

NewCloudSnapshot instantiates a new CloudSnapshot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudSnapshotWithDefaults

`func NewCloudSnapshotWithDefaults() *CloudSnapshot`

NewCloudSnapshotWithDefaults instantiates a new CloudSnapshot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAt

`func (o *CloudSnapshot) GetAt() string`

GetAt returns the At field if non-nil, zero value otherwise.

### GetAtOk

`func (o *CloudSnapshot) GetAtOk() (*string, bool)`

GetAtOk returns a tuple with the At field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAt

`func (o *CloudSnapshot) SetAt(v string)`

SetAt sets At field to given value.

### HasAt

`func (o *CloudSnapshot) HasAt() bool`

HasAt returns a boolean if a field has been set.

### GetClusters

`func (o *CloudSnapshot) GetClusters() []CloudCluster`

GetClusters returns the Clusters field if non-nil, zero value otherwise.

### GetClustersOk

`func (o *CloudSnapshot) GetClustersOk() (*[]CloudCluster, bool)`

GetClustersOk returns a tuple with the Clusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusters

`func (o *CloudSnapshot) SetClusters(v []CloudCluster)`

SetClusters sets Clusters field to given value.

### HasClusters

`func (o *CloudSnapshot) HasClusters() bool`

HasClusters returns a boolean if a field has been set.

### GetComplete

`func (o *CloudSnapshot) GetComplete() bool`

GetComplete returns the Complete field if non-nil, zero value otherwise.

### GetCompleteOk

`func (o *CloudSnapshot) GetCompleteOk() (*bool, bool)`

GetCompleteOk returns a tuple with the Complete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplete

`func (o *CloudSnapshot) SetComplete(v bool)`

SetComplete sets Complete field to given value.

### HasComplete

`func (o *CloudSnapshot) HasComplete() bool`

HasComplete returns a boolean if a field has been set.

### GetCost

`func (o *CloudSnapshot) GetCost() CloudCost`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *CloudSnapshot) GetCostOk() (*CloudCost, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *CloudSnapshot) SetCost(v CloudCost)`

SetCost sets Cost field to given value.

### HasCost

`func (o *CloudSnapshot) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetFindings

`func (o *CloudSnapshot) GetFindings() []CloudFinding`

GetFindings returns the Findings field if non-nil, zero value otherwise.

### GetFindingsOk

`func (o *CloudSnapshot) GetFindingsOk() (*[]CloudFinding, bool)`

GetFindingsOk returns a tuple with the Findings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFindings

`func (o *CloudSnapshot) SetFindings(v []CloudFinding)`

SetFindings sets Findings field to given value.

### HasFindings

`func (o *CloudSnapshot) HasFindings() bool`

HasFindings returns a boolean if a field has been set.

### GetIncompleteReason

`func (o *CloudSnapshot) GetIncompleteReason() string`

GetIncompleteReason returns the IncompleteReason field if non-nil, zero value otherwise.

### GetIncompleteReasonOk

`func (o *CloudSnapshot) GetIncompleteReasonOk() (*string, bool)`

GetIncompleteReasonOk returns a tuple with the IncompleteReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncompleteReason

`func (o *CloudSnapshot) SetIncompleteReason(v string)`

SetIncompleteReason sets IncompleteReason field to given value.

### HasIncompleteReason

`func (o *CloudSnapshot) HasIncompleteReason() bool`

HasIncompleteReason returns a boolean if a field has been set.

### GetLoadBalancers

`func (o *CloudSnapshot) GetLoadBalancers() []CloudLoadBalancer`

GetLoadBalancers returns the LoadBalancers field if non-nil, zero value otherwise.

### GetLoadBalancersOk

`func (o *CloudSnapshot) GetLoadBalancersOk() (*[]CloudLoadBalancer, bool)`

GetLoadBalancersOk returns a tuple with the LoadBalancers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancers

`func (o *CloudSnapshot) SetLoadBalancers(v []CloudLoadBalancer)`

SetLoadBalancers sets LoadBalancers field to given value.

### HasLoadBalancers

`func (o *CloudSnapshot) HasLoadBalancers() bool`

HasLoadBalancers returns a boolean if a field has been set.

### GetNodes

`func (o *CloudSnapshot) GetNodes() []CloudNode`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *CloudSnapshot) GetNodesOk() (*[]CloudNode, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *CloudSnapshot) SetNodes(v []CloudNode)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *CloudSnapshot) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetSources

`func (o *CloudSnapshot) GetSources() []CloudSourceStatus`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *CloudSnapshot) GetSourcesOk() (*[]CloudSourceStatus, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *CloudSnapshot) SetSources(v []CloudSourceStatus)`

SetSources sets Sources field to given value.

### HasSources

`func (o *CloudSnapshot) HasSources() bool`

HasSources returns a boolean if a field has been set.

### GetTotals

`func (o *CloudSnapshot) GetTotals() CloudTotals`

GetTotals returns the Totals field if non-nil, zero value otherwise.

### GetTotalsOk

`func (o *CloudSnapshot) GetTotalsOk() (*CloudTotals, bool)`

GetTotalsOk returns a tuple with the Totals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotals

`func (o *CloudSnapshot) SetTotals(v CloudTotals)`

SetTotals sets Totals field to given value.

### HasTotals

`func (o *CloudSnapshot) HasTotals() bool`

HasTotals returns a boolean if a field has been set.

### GetVolumes

`func (o *CloudSnapshot) GetVolumes() []CloudVolume`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *CloudSnapshot) GetVolumesOk() (*[]CloudVolume, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *CloudSnapshot) SetVolumes(v []CloudVolume)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *CloudSnapshot) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


