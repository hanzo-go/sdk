# CloudTotals

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttachedGiB** | Pointer to **int32** |  | [optional] 
**AttachedVolumes** | Pointer to **int32** |  | [optional] 
**Clusters** | Pointer to **int32** |  | [optional] 
**DetachedGiB** | Pointer to **int32** |  | [optional] 
**DetachedVolumes** | Pointer to **int32** |  | [optional] 
**IdlePVCs** | Pointer to **int32** |  | [optional] 
**LoadBalancers** | Pointer to **int32** |  | [optional] 
**LocalDiskGiB** | Pointer to **int32** |  | [optional] 
**MeasuredGiB** | Pointer to **int32** |  | [optional] 
**MeasuredVolumes** | Pointer to **int32** | Fill. MeasuredVolumes/UnmeasuredVolumes are the honesty denominator: UsedGiB and WastedGiB describe the measured set ONLY, so a board showing waste must show how much of the fleet the figure was computed from. Unmeasured capacity contributes nothing to either — it is not assumed empty, and it is not assumed full. | [optional] 
**Nodes** | Pointer to **int32** |  | [optional] 
**UnmeasuredGiB** | Pointer to **int32** |  | [optional] 
**UnmeasuredVolumes** | Pointer to **int32** |  | [optional] 
**UnreferencedGiB** | Pointer to **int32** |  | [optional] 
**UnreferencedVolumes** | Pointer to **int32** |  | [optional] 
**UsedGiB** | Pointer to **int32** |  | [optional] 
**VolumeGiB** | Pointer to **int32** |  | [optional] 
**Volumes** | Pointer to **int32** |  | [optional] 
**WastedGiB** | Pointer to **int32** |  | [optional] 

## Methods

### NewCloudTotals

`func NewCloudTotals() *CloudTotals`

NewCloudTotals instantiates a new CloudTotals object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudTotalsWithDefaults

`func NewCloudTotalsWithDefaults() *CloudTotals`

NewCloudTotalsWithDefaults instantiates a new CloudTotals object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttachedGiB

`func (o *CloudTotals) GetAttachedGiB() int32`

GetAttachedGiB returns the AttachedGiB field if non-nil, zero value otherwise.

### GetAttachedGiBOk

`func (o *CloudTotals) GetAttachedGiBOk() (*int32, bool)`

GetAttachedGiBOk returns a tuple with the AttachedGiB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachedGiB

`func (o *CloudTotals) SetAttachedGiB(v int32)`

SetAttachedGiB sets AttachedGiB field to given value.

### HasAttachedGiB

`func (o *CloudTotals) HasAttachedGiB() bool`

HasAttachedGiB returns a boolean if a field has been set.

### GetAttachedVolumes

`func (o *CloudTotals) GetAttachedVolumes() int32`

GetAttachedVolumes returns the AttachedVolumes field if non-nil, zero value otherwise.

### GetAttachedVolumesOk

`func (o *CloudTotals) GetAttachedVolumesOk() (*int32, bool)`

GetAttachedVolumesOk returns a tuple with the AttachedVolumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachedVolumes

`func (o *CloudTotals) SetAttachedVolumes(v int32)`

SetAttachedVolumes sets AttachedVolumes field to given value.

### HasAttachedVolumes

`func (o *CloudTotals) HasAttachedVolumes() bool`

HasAttachedVolumes returns a boolean if a field has been set.

### GetClusters

`func (o *CloudTotals) GetClusters() int32`

GetClusters returns the Clusters field if non-nil, zero value otherwise.

### GetClustersOk

`func (o *CloudTotals) GetClustersOk() (*int32, bool)`

GetClustersOk returns a tuple with the Clusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusters

`func (o *CloudTotals) SetClusters(v int32)`

SetClusters sets Clusters field to given value.

### HasClusters

`func (o *CloudTotals) HasClusters() bool`

HasClusters returns a boolean if a field has been set.

### GetDetachedGiB

`func (o *CloudTotals) GetDetachedGiB() int32`

GetDetachedGiB returns the DetachedGiB field if non-nil, zero value otherwise.

### GetDetachedGiBOk

`func (o *CloudTotals) GetDetachedGiBOk() (*int32, bool)`

GetDetachedGiBOk returns a tuple with the DetachedGiB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetachedGiB

`func (o *CloudTotals) SetDetachedGiB(v int32)`

SetDetachedGiB sets DetachedGiB field to given value.

### HasDetachedGiB

`func (o *CloudTotals) HasDetachedGiB() bool`

HasDetachedGiB returns a boolean if a field has been set.

### GetDetachedVolumes

`func (o *CloudTotals) GetDetachedVolumes() int32`

GetDetachedVolumes returns the DetachedVolumes field if non-nil, zero value otherwise.

### GetDetachedVolumesOk

`func (o *CloudTotals) GetDetachedVolumesOk() (*int32, bool)`

GetDetachedVolumesOk returns a tuple with the DetachedVolumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetachedVolumes

`func (o *CloudTotals) SetDetachedVolumes(v int32)`

SetDetachedVolumes sets DetachedVolumes field to given value.

### HasDetachedVolumes

`func (o *CloudTotals) HasDetachedVolumes() bool`

HasDetachedVolumes returns a boolean if a field has been set.

### GetIdlePVCs

`func (o *CloudTotals) GetIdlePVCs() int32`

GetIdlePVCs returns the IdlePVCs field if non-nil, zero value otherwise.

### GetIdlePVCsOk

`func (o *CloudTotals) GetIdlePVCsOk() (*int32, bool)`

GetIdlePVCsOk returns a tuple with the IdlePVCs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdlePVCs

`func (o *CloudTotals) SetIdlePVCs(v int32)`

SetIdlePVCs sets IdlePVCs field to given value.

### HasIdlePVCs

`func (o *CloudTotals) HasIdlePVCs() bool`

HasIdlePVCs returns a boolean if a field has been set.

### GetLoadBalancers

`func (o *CloudTotals) GetLoadBalancers() int32`

GetLoadBalancers returns the LoadBalancers field if non-nil, zero value otherwise.

### GetLoadBalancersOk

`func (o *CloudTotals) GetLoadBalancersOk() (*int32, bool)`

GetLoadBalancersOk returns a tuple with the LoadBalancers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancers

`func (o *CloudTotals) SetLoadBalancers(v int32)`

SetLoadBalancers sets LoadBalancers field to given value.

### HasLoadBalancers

`func (o *CloudTotals) HasLoadBalancers() bool`

HasLoadBalancers returns a boolean if a field has been set.

### GetLocalDiskGiB

`func (o *CloudTotals) GetLocalDiskGiB() int32`

GetLocalDiskGiB returns the LocalDiskGiB field if non-nil, zero value otherwise.

### GetLocalDiskGiBOk

`func (o *CloudTotals) GetLocalDiskGiBOk() (*int32, bool)`

GetLocalDiskGiBOk returns a tuple with the LocalDiskGiB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalDiskGiB

`func (o *CloudTotals) SetLocalDiskGiB(v int32)`

SetLocalDiskGiB sets LocalDiskGiB field to given value.

### HasLocalDiskGiB

`func (o *CloudTotals) HasLocalDiskGiB() bool`

HasLocalDiskGiB returns a boolean if a field has been set.

### GetMeasuredGiB

`func (o *CloudTotals) GetMeasuredGiB() int32`

GetMeasuredGiB returns the MeasuredGiB field if non-nil, zero value otherwise.

### GetMeasuredGiBOk

`func (o *CloudTotals) GetMeasuredGiBOk() (*int32, bool)`

GetMeasuredGiBOk returns a tuple with the MeasuredGiB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasuredGiB

`func (o *CloudTotals) SetMeasuredGiB(v int32)`

SetMeasuredGiB sets MeasuredGiB field to given value.

### HasMeasuredGiB

`func (o *CloudTotals) HasMeasuredGiB() bool`

HasMeasuredGiB returns a boolean if a field has been set.

### GetMeasuredVolumes

`func (o *CloudTotals) GetMeasuredVolumes() int32`

GetMeasuredVolumes returns the MeasuredVolumes field if non-nil, zero value otherwise.

### GetMeasuredVolumesOk

`func (o *CloudTotals) GetMeasuredVolumesOk() (*int32, bool)`

GetMeasuredVolumesOk returns a tuple with the MeasuredVolumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasuredVolumes

`func (o *CloudTotals) SetMeasuredVolumes(v int32)`

SetMeasuredVolumes sets MeasuredVolumes field to given value.

### HasMeasuredVolumes

`func (o *CloudTotals) HasMeasuredVolumes() bool`

HasMeasuredVolumes returns a boolean if a field has been set.

### GetNodes

`func (o *CloudTotals) GetNodes() int32`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *CloudTotals) GetNodesOk() (*int32, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *CloudTotals) SetNodes(v int32)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *CloudTotals) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetUnmeasuredGiB

`func (o *CloudTotals) GetUnmeasuredGiB() int32`

GetUnmeasuredGiB returns the UnmeasuredGiB field if non-nil, zero value otherwise.

### GetUnmeasuredGiBOk

`func (o *CloudTotals) GetUnmeasuredGiBOk() (*int32, bool)`

GetUnmeasuredGiBOk returns a tuple with the UnmeasuredGiB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnmeasuredGiB

`func (o *CloudTotals) SetUnmeasuredGiB(v int32)`

SetUnmeasuredGiB sets UnmeasuredGiB field to given value.

### HasUnmeasuredGiB

`func (o *CloudTotals) HasUnmeasuredGiB() bool`

HasUnmeasuredGiB returns a boolean if a field has been set.

### GetUnmeasuredVolumes

`func (o *CloudTotals) GetUnmeasuredVolumes() int32`

GetUnmeasuredVolumes returns the UnmeasuredVolumes field if non-nil, zero value otherwise.

### GetUnmeasuredVolumesOk

`func (o *CloudTotals) GetUnmeasuredVolumesOk() (*int32, bool)`

GetUnmeasuredVolumesOk returns a tuple with the UnmeasuredVolumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnmeasuredVolumes

`func (o *CloudTotals) SetUnmeasuredVolumes(v int32)`

SetUnmeasuredVolumes sets UnmeasuredVolumes field to given value.

### HasUnmeasuredVolumes

`func (o *CloudTotals) HasUnmeasuredVolumes() bool`

HasUnmeasuredVolumes returns a boolean if a field has been set.

### GetUnreferencedGiB

`func (o *CloudTotals) GetUnreferencedGiB() int32`

GetUnreferencedGiB returns the UnreferencedGiB field if non-nil, zero value otherwise.

### GetUnreferencedGiBOk

`func (o *CloudTotals) GetUnreferencedGiBOk() (*int32, bool)`

GetUnreferencedGiBOk returns a tuple with the UnreferencedGiB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnreferencedGiB

`func (o *CloudTotals) SetUnreferencedGiB(v int32)`

SetUnreferencedGiB sets UnreferencedGiB field to given value.

### HasUnreferencedGiB

`func (o *CloudTotals) HasUnreferencedGiB() bool`

HasUnreferencedGiB returns a boolean if a field has been set.

### GetUnreferencedVolumes

`func (o *CloudTotals) GetUnreferencedVolumes() int32`

GetUnreferencedVolumes returns the UnreferencedVolumes field if non-nil, zero value otherwise.

### GetUnreferencedVolumesOk

`func (o *CloudTotals) GetUnreferencedVolumesOk() (*int32, bool)`

GetUnreferencedVolumesOk returns a tuple with the UnreferencedVolumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnreferencedVolumes

`func (o *CloudTotals) SetUnreferencedVolumes(v int32)`

SetUnreferencedVolumes sets UnreferencedVolumes field to given value.

### HasUnreferencedVolumes

`func (o *CloudTotals) HasUnreferencedVolumes() bool`

HasUnreferencedVolumes returns a boolean if a field has been set.

### GetUsedGiB

`func (o *CloudTotals) GetUsedGiB() int32`

GetUsedGiB returns the UsedGiB field if non-nil, zero value otherwise.

### GetUsedGiBOk

`func (o *CloudTotals) GetUsedGiBOk() (*int32, bool)`

GetUsedGiBOk returns a tuple with the UsedGiB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedGiB

`func (o *CloudTotals) SetUsedGiB(v int32)`

SetUsedGiB sets UsedGiB field to given value.

### HasUsedGiB

`func (o *CloudTotals) HasUsedGiB() bool`

HasUsedGiB returns a boolean if a field has been set.

### GetVolumeGiB

`func (o *CloudTotals) GetVolumeGiB() int32`

GetVolumeGiB returns the VolumeGiB field if non-nil, zero value otherwise.

### GetVolumeGiBOk

`func (o *CloudTotals) GetVolumeGiBOk() (*int32, bool)`

GetVolumeGiBOk returns a tuple with the VolumeGiB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeGiB

`func (o *CloudTotals) SetVolumeGiB(v int32)`

SetVolumeGiB sets VolumeGiB field to given value.

### HasVolumeGiB

`func (o *CloudTotals) HasVolumeGiB() bool`

HasVolumeGiB returns a boolean if a field has been set.

### GetVolumes

`func (o *CloudTotals) GetVolumes() int32`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *CloudTotals) GetVolumesOk() (*int32, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *CloudTotals) SetVolumes(v int32)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *CloudTotals) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.

### GetWastedGiB

`func (o *CloudTotals) GetWastedGiB() int32`

GetWastedGiB returns the WastedGiB field if non-nil, zero value otherwise.

### GetWastedGiBOk

`func (o *CloudTotals) GetWastedGiBOk() (*int32, bool)`

GetWastedGiBOk returns a tuple with the WastedGiB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWastedGiB

`func (o *CloudTotals) SetWastedGiB(v int32)`

SetWastedGiB sets WastedGiB field to given value.

### HasWastedGiB

`func (o *CloudTotals) HasWastedGiB() bool`

HasWastedGiB returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


