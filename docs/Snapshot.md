# Snapshot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**At** | Pointer to **string** |  | [optional] 
**Clusters** | Pointer to [**[]Cluster**](Cluster.md) |  | [optional] 
**Complete** | Pointer to **bool** |  | [optional] 
**Cost** | Pointer to [**Cost**](Cost.md) |  | [optional] 
**Findings** | Pointer to [**[]Finding**](Finding.md) |  | [optional] 
**IncompleteReason** | Pointer to **string** |  | [optional] 
**LoadBalancers** | Pointer to [**[]LoadBalancer**](LoadBalancer.md) |  | [optional] 
**Nodes** | Pointer to [**[]Machine**](Machine.md) |  | [optional] 
**Sources** | Pointer to [**[]SourceStatus**](SourceStatus.md) |  | [optional] 
**Totals** | Pointer to [**Totals**](Totals.md) |  | [optional] 
**Volumes** | Pointer to [**[]Volume**](Volume.md) |  | [optional] 

## Methods

### NewSnapshot

`func NewSnapshot() *Snapshot`

NewSnapshot instantiates a new Snapshot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotWithDefaults

`func NewSnapshotWithDefaults() *Snapshot`

NewSnapshotWithDefaults instantiates a new Snapshot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAt

`func (o *Snapshot) GetAt() string`

GetAt returns the At field if non-nil, zero value otherwise.

### GetAtOk

`func (o *Snapshot) GetAtOk() (*string, bool)`

GetAtOk returns a tuple with the At field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAt

`func (o *Snapshot) SetAt(v string)`

SetAt sets At field to given value.

### HasAt

`func (o *Snapshot) HasAt() bool`

HasAt returns a boolean if a field has been set.

### GetClusters

`func (o *Snapshot) GetClusters() []Cluster`

GetClusters returns the Clusters field if non-nil, zero value otherwise.

### GetClustersOk

`func (o *Snapshot) GetClustersOk() (*[]Cluster, bool)`

GetClustersOk returns a tuple with the Clusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusters

`func (o *Snapshot) SetClusters(v []Cluster)`

SetClusters sets Clusters field to given value.

### HasClusters

`func (o *Snapshot) HasClusters() bool`

HasClusters returns a boolean if a field has been set.

### GetComplete

`func (o *Snapshot) GetComplete() bool`

GetComplete returns the Complete field if non-nil, zero value otherwise.

### GetCompleteOk

`func (o *Snapshot) GetCompleteOk() (*bool, bool)`

GetCompleteOk returns a tuple with the Complete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplete

`func (o *Snapshot) SetComplete(v bool)`

SetComplete sets Complete field to given value.

### HasComplete

`func (o *Snapshot) HasComplete() bool`

HasComplete returns a boolean if a field has been set.

### GetCost

`func (o *Snapshot) GetCost() Cost`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *Snapshot) GetCostOk() (*Cost, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *Snapshot) SetCost(v Cost)`

SetCost sets Cost field to given value.

### HasCost

`func (o *Snapshot) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetFindings

`func (o *Snapshot) GetFindings() []Finding`

GetFindings returns the Findings field if non-nil, zero value otherwise.

### GetFindingsOk

`func (o *Snapshot) GetFindingsOk() (*[]Finding, bool)`

GetFindingsOk returns a tuple with the Findings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFindings

`func (o *Snapshot) SetFindings(v []Finding)`

SetFindings sets Findings field to given value.

### HasFindings

`func (o *Snapshot) HasFindings() bool`

HasFindings returns a boolean if a field has been set.

### GetIncompleteReason

`func (o *Snapshot) GetIncompleteReason() string`

GetIncompleteReason returns the IncompleteReason field if non-nil, zero value otherwise.

### GetIncompleteReasonOk

`func (o *Snapshot) GetIncompleteReasonOk() (*string, bool)`

GetIncompleteReasonOk returns a tuple with the IncompleteReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncompleteReason

`func (o *Snapshot) SetIncompleteReason(v string)`

SetIncompleteReason sets IncompleteReason field to given value.

### HasIncompleteReason

`func (o *Snapshot) HasIncompleteReason() bool`

HasIncompleteReason returns a boolean if a field has been set.

### GetLoadBalancers

`func (o *Snapshot) GetLoadBalancers() []LoadBalancer`

GetLoadBalancers returns the LoadBalancers field if non-nil, zero value otherwise.

### GetLoadBalancersOk

`func (o *Snapshot) GetLoadBalancersOk() (*[]LoadBalancer, bool)`

GetLoadBalancersOk returns a tuple with the LoadBalancers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancers

`func (o *Snapshot) SetLoadBalancers(v []LoadBalancer)`

SetLoadBalancers sets LoadBalancers field to given value.

### HasLoadBalancers

`func (o *Snapshot) HasLoadBalancers() bool`

HasLoadBalancers returns a boolean if a field has been set.

### GetNodes

`func (o *Snapshot) GetNodes() []Machine`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *Snapshot) GetNodesOk() (*[]Machine, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *Snapshot) SetNodes(v []Machine)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *Snapshot) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetSources

`func (o *Snapshot) GetSources() []SourceStatus`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *Snapshot) GetSourcesOk() (*[]SourceStatus, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *Snapshot) SetSources(v []SourceStatus)`

SetSources sets Sources field to given value.

### HasSources

`func (o *Snapshot) HasSources() bool`

HasSources returns a boolean if a field has been set.

### GetTotals

`func (o *Snapshot) GetTotals() Totals`

GetTotals returns the Totals field if non-nil, zero value otherwise.

### GetTotalsOk

`func (o *Snapshot) GetTotalsOk() (*Totals, bool)`

GetTotalsOk returns a tuple with the Totals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotals

`func (o *Snapshot) SetTotals(v Totals)`

SetTotals sets Totals field to given value.

### HasTotals

`func (o *Snapshot) HasTotals() bool`

HasTotals returns a boolean if a field has been set.

### GetVolumes

`func (o *Snapshot) GetVolumes() []Volume`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *Snapshot) GetVolumesOk() (*[]Volume, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *Snapshot) SetVolumes(v []Volume)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *Snapshot) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


