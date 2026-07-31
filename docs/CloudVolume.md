# CloudVolume

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockedReason** | Pointer to **string** |  | [optional] 
**Cluster** | Pointer to **string** | Cluster/ClusterID are the PROVEN owner — resolved through a PV that names this volume, never through the tag. | [optional] 
**ClusterId** | Pointer to **string** |  | [optional] 
**Controller** | Pointer to **string** | Controller is the workload owning the pod that mounts this volume (\&quot;StatefulSet/luxd\&quot;), or \&quot;\&quot; when nothing mounts it. It names who has to act. | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**Deletable** | Pointer to **bool** |  | [optional] 
**DropletIds** | Pointer to **[]int32** |  | [optional] 
**ExpandBlockedReason** | Pointer to **string** |  | [optional] 
**Expandable** | Pointer to **bool** | Expandable/ExpandBlockedReason are the GROW verdict, kept separate from Deletable because the two ask opposite questions: a volume is deletable when nothing uses it, and expandable when something uses it in a way this board can grow completely. | [optional] 
**HasUsage** | Pointer to **bool** | HasUsage reports whether a kubelet actually MEASURED this volume&#39;s filesystem.  False means NOT MEASURED. It does NOT mean empty, and the three fields below are meaningless — not zero — when it is false. A reading exists only while a running pod has the volume mounted on a node that answered; a detached, idle or unreferenced volume has none. Rendering an unmeasured volume as \&quot;0 used / 100% wasted\&quot; would invent the single most expensive lie this board could tell, so every consumer must branch on this flag and show unknown. | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Idle** | Pointer to **bool** |  | [optional] 
**MonthlyCents** | Pointer to **int32** |  | [optional] 
**MountedBy** | Pointer to **[]string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NodeName** | Pointer to **string** |  | [optional] 
**Pv** | Pointer to **string** |  | [optional] 
**PvPhase** | Pointer to **string** |  | [optional] 
**PvcName** | Pointer to **string** |  | [optional] 
**PvcNamespace** | Pointer to **string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**SizeGiB** | Pointer to **int32** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**TagCluster** | Pointer to **string** | TagCluster is the &#x60;k8s:&lt;uuid&gt;&#x60; tag. ADVISORY ONLY: it outlives the cluster that set it. Shown so the operator can see tag-vs-truth disagree, never acted on. | [optional] 
**UsedBytes** | Pointer to **int32** | UsedBytes is the measured filesystem usage. BYTES, not GiB: the volumes this exists to catch hold a fraction of a GiB in 200, and rounding that to an integer GiB would print the very 0 the flag above exists to prevent. | [optional] 
**WastedGiB** | Pointer to **int32** | WastedGiB is provisioned minus measured, in the unit DigitalOcean BILLS: whole GiB of the volume&#39;s own size, never the filesystem&#39;s capacity — a 200 GiB volume carries a 196 GiB filesystem after format overhead, and the invoice says 200. | [optional] 
**WastedMonthlyCents** | Pointer to **int32** |  | [optional] 

## Methods

### NewCloudVolume

`func NewCloudVolume() *CloudVolume`

NewCloudVolume instantiates a new CloudVolume object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudVolumeWithDefaults

`func NewCloudVolumeWithDefaults() *CloudVolume`

NewCloudVolumeWithDefaults instantiates a new CloudVolume object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockedReason

`func (o *CloudVolume) GetBlockedReason() string`

GetBlockedReason returns the BlockedReason field if non-nil, zero value otherwise.

### GetBlockedReasonOk

`func (o *CloudVolume) GetBlockedReasonOk() (*string, bool)`

GetBlockedReasonOk returns a tuple with the BlockedReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedReason

`func (o *CloudVolume) SetBlockedReason(v string)`

SetBlockedReason sets BlockedReason field to given value.

### HasBlockedReason

`func (o *CloudVolume) HasBlockedReason() bool`

HasBlockedReason returns a boolean if a field has been set.

### GetCluster

`func (o *CloudVolume) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *CloudVolume) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *CloudVolume) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *CloudVolume) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetClusterId

`func (o *CloudVolume) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *CloudVolume) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *CloudVolume) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *CloudVolume) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### GetController

`func (o *CloudVolume) GetController() string`

GetController returns the Controller field if non-nil, zero value otherwise.

### GetControllerOk

`func (o *CloudVolume) GetControllerOk() (*string, bool)`

GetControllerOk returns a tuple with the Controller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetController

`func (o *CloudVolume) SetController(v string)`

SetController sets Controller field to given value.

### HasController

`func (o *CloudVolume) HasController() bool`

HasController returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CloudVolume) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CloudVolume) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CloudVolume) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CloudVolume) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeletable

`func (o *CloudVolume) GetDeletable() bool`

GetDeletable returns the Deletable field if non-nil, zero value otherwise.

### GetDeletableOk

`func (o *CloudVolume) GetDeletableOk() (*bool, bool)`

GetDeletableOk returns a tuple with the Deletable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletable

`func (o *CloudVolume) SetDeletable(v bool)`

SetDeletable sets Deletable field to given value.

### HasDeletable

`func (o *CloudVolume) HasDeletable() bool`

HasDeletable returns a boolean if a field has been set.

### GetDropletIds

`func (o *CloudVolume) GetDropletIds() []int32`

GetDropletIds returns the DropletIds field if non-nil, zero value otherwise.

### GetDropletIdsOk

`func (o *CloudVolume) GetDropletIdsOk() (*[]int32, bool)`

GetDropletIdsOk returns a tuple with the DropletIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDropletIds

`func (o *CloudVolume) SetDropletIds(v []int32)`

SetDropletIds sets DropletIds field to given value.

### HasDropletIds

`func (o *CloudVolume) HasDropletIds() bool`

HasDropletIds returns a boolean if a field has been set.

### GetExpandBlockedReason

`func (o *CloudVolume) GetExpandBlockedReason() string`

GetExpandBlockedReason returns the ExpandBlockedReason field if non-nil, zero value otherwise.

### GetExpandBlockedReasonOk

`func (o *CloudVolume) GetExpandBlockedReasonOk() (*string, bool)`

GetExpandBlockedReasonOk returns a tuple with the ExpandBlockedReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpandBlockedReason

`func (o *CloudVolume) SetExpandBlockedReason(v string)`

SetExpandBlockedReason sets ExpandBlockedReason field to given value.

### HasExpandBlockedReason

`func (o *CloudVolume) HasExpandBlockedReason() bool`

HasExpandBlockedReason returns a boolean if a field has been set.

### GetExpandable

`func (o *CloudVolume) GetExpandable() bool`

GetExpandable returns the Expandable field if non-nil, zero value otherwise.

### GetExpandableOk

`func (o *CloudVolume) GetExpandableOk() (*bool, bool)`

GetExpandableOk returns a tuple with the Expandable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpandable

`func (o *CloudVolume) SetExpandable(v bool)`

SetExpandable sets Expandable field to given value.

### HasExpandable

`func (o *CloudVolume) HasExpandable() bool`

HasExpandable returns a boolean if a field has been set.

### GetHasUsage

`func (o *CloudVolume) GetHasUsage() bool`

GetHasUsage returns the HasUsage field if non-nil, zero value otherwise.

### GetHasUsageOk

`func (o *CloudVolume) GetHasUsageOk() (*bool, bool)`

GetHasUsageOk returns a tuple with the HasUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasUsage

`func (o *CloudVolume) SetHasUsage(v bool)`

SetHasUsage sets HasUsage field to given value.

### HasHasUsage

`func (o *CloudVolume) HasHasUsage() bool`

HasHasUsage returns a boolean if a field has been set.

### GetId

`func (o *CloudVolume) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudVolume) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudVolume) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudVolume) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdle

`func (o *CloudVolume) GetIdle() bool`

GetIdle returns the Idle field if non-nil, zero value otherwise.

### GetIdleOk

`func (o *CloudVolume) GetIdleOk() (*bool, bool)`

GetIdleOk returns a tuple with the Idle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdle

`func (o *CloudVolume) SetIdle(v bool)`

SetIdle sets Idle field to given value.

### HasIdle

`func (o *CloudVolume) HasIdle() bool`

HasIdle returns a boolean if a field has been set.

### GetMonthlyCents

`func (o *CloudVolume) GetMonthlyCents() int32`

GetMonthlyCents returns the MonthlyCents field if non-nil, zero value otherwise.

### GetMonthlyCentsOk

`func (o *CloudVolume) GetMonthlyCentsOk() (*int32, bool)`

GetMonthlyCentsOk returns a tuple with the MonthlyCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyCents

`func (o *CloudVolume) SetMonthlyCents(v int32)`

SetMonthlyCents sets MonthlyCents field to given value.

### HasMonthlyCents

`func (o *CloudVolume) HasMonthlyCents() bool`

HasMonthlyCents returns a boolean if a field has been set.

### GetMountedBy

`func (o *CloudVolume) GetMountedBy() []string`

GetMountedBy returns the MountedBy field if non-nil, zero value otherwise.

### GetMountedByOk

`func (o *CloudVolume) GetMountedByOk() (*[]string, bool)`

GetMountedByOk returns a tuple with the MountedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountedBy

`func (o *CloudVolume) SetMountedBy(v []string)`

SetMountedBy sets MountedBy field to given value.

### HasMountedBy

`func (o *CloudVolume) HasMountedBy() bool`

HasMountedBy returns a boolean if a field has been set.

### GetName

`func (o *CloudVolume) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudVolume) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudVolume) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudVolume) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNodeName

`func (o *CloudVolume) GetNodeName() string`

GetNodeName returns the NodeName field if non-nil, zero value otherwise.

### GetNodeNameOk

`func (o *CloudVolume) GetNodeNameOk() (*string, bool)`

GetNodeNameOk returns a tuple with the NodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeName

`func (o *CloudVolume) SetNodeName(v string)`

SetNodeName sets NodeName field to given value.

### HasNodeName

`func (o *CloudVolume) HasNodeName() bool`

HasNodeName returns a boolean if a field has been set.

### GetPv

`func (o *CloudVolume) GetPv() string`

GetPv returns the Pv field if non-nil, zero value otherwise.

### GetPvOk

`func (o *CloudVolume) GetPvOk() (*string, bool)`

GetPvOk returns a tuple with the Pv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPv

`func (o *CloudVolume) SetPv(v string)`

SetPv sets Pv field to given value.

### HasPv

`func (o *CloudVolume) HasPv() bool`

HasPv returns a boolean if a field has been set.

### GetPvPhase

`func (o *CloudVolume) GetPvPhase() string`

GetPvPhase returns the PvPhase field if non-nil, zero value otherwise.

### GetPvPhaseOk

`func (o *CloudVolume) GetPvPhaseOk() (*string, bool)`

GetPvPhaseOk returns a tuple with the PvPhase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPvPhase

`func (o *CloudVolume) SetPvPhase(v string)`

SetPvPhase sets PvPhase field to given value.

### HasPvPhase

`func (o *CloudVolume) HasPvPhase() bool`

HasPvPhase returns a boolean if a field has been set.

### GetPvcName

`func (o *CloudVolume) GetPvcName() string`

GetPvcName returns the PvcName field if non-nil, zero value otherwise.

### GetPvcNameOk

`func (o *CloudVolume) GetPvcNameOk() (*string, bool)`

GetPvcNameOk returns a tuple with the PvcName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPvcName

`func (o *CloudVolume) SetPvcName(v string)`

SetPvcName sets PvcName field to given value.

### HasPvcName

`func (o *CloudVolume) HasPvcName() bool`

HasPvcName returns a boolean if a field has been set.

### GetPvcNamespace

`func (o *CloudVolume) GetPvcNamespace() string`

GetPvcNamespace returns the PvcNamespace field if non-nil, zero value otherwise.

### GetPvcNamespaceOk

`func (o *CloudVolume) GetPvcNamespaceOk() (*string, bool)`

GetPvcNamespaceOk returns a tuple with the PvcNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPvcNamespace

`func (o *CloudVolume) SetPvcNamespace(v string)`

SetPvcNamespace sets PvcNamespace field to given value.

### HasPvcNamespace

`func (o *CloudVolume) HasPvcNamespace() bool`

HasPvcNamespace returns a boolean if a field has been set.

### GetRegion

`func (o *CloudVolume) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CloudVolume) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CloudVolume) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *CloudVolume) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetSizeGiB

`func (o *CloudVolume) GetSizeGiB() int32`

GetSizeGiB returns the SizeGiB field if non-nil, zero value otherwise.

### GetSizeGiBOk

`func (o *CloudVolume) GetSizeGiBOk() (*int32, bool)`

GetSizeGiBOk returns a tuple with the SizeGiB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeGiB

`func (o *CloudVolume) SetSizeGiB(v int32)`

SetSizeGiB sets SizeGiB field to given value.

### HasSizeGiB

`func (o *CloudVolume) HasSizeGiB() bool`

HasSizeGiB returns a boolean if a field has been set.

### GetState

`func (o *CloudVolume) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CloudVolume) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CloudVolume) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *CloudVolume) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTagCluster

`func (o *CloudVolume) GetTagCluster() string`

GetTagCluster returns the TagCluster field if non-nil, zero value otherwise.

### GetTagClusterOk

`func (o *CloudVolume) GetTagClusterOk() (*string, bool)`

GetTagClusterOk returns a tuple with the TagCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagCluster

`func (o *CloudVolume) SetTagCluster(v string)`

SetTagCluster sets TagCluster field to given value.

### HasTagCluster

`func (o *CloudVolume) HasTagCluster() bool`

HasTagCluster returns a boolean if a field has been set.

### GetUsedBytes

`func (o *CloudVolume) GetUsedBytes() int32`

GetUsedBytes returns the UsedBytes field if non-nil, zero value otherwise.

### GetUsedBytesOk

`func (o *CloudVolume) GetUsedBytesOk() (*int32, bool)`

GetUsedBytesOk returns a tuple with the UsedBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedBytes

`func (o *CloudVolume) SetUsedBytes(v int32)`

SetUsedBytes sets UsedBytes field to given value.

### HasUsedBytes

`func (o *CloudVolume) HasUsedBytes() bool`

HasUsedBytes returns a boolean if a field has been set.

### GetWastedGiB

`func (o *CloudVolume) GetWastedGiB() int32`

GetWastedGiB returns the WastedGiB field if non-nil, zero value otherwise.

### GetWastedGiBOk

`func (o *CloudVolume) GetWastedGiBOk() (*int32, bool)`

GetWastedGiBOk returns a tuple with the WastedGiB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWastedGiB

`func (o *CloudVolume) SetWastedGiB(v int32)`

SetWastedGiB sets WastedGiB field to given value.

### HasWastedGiB

`func (o *CloudVolume) HasWastedGiB() bool`

HasWastedGiB returns a boolean if a field has been set.

### GetWastedMonthlyCents

`func (o *CloudVolume) GetWastedMonthlyCents() int32`

GetWastedMonthlyCents returns the WastedMonthlyCents field if non-nil, zero value otherwise.

### GetWastedMonthlyCentsOk

`func (o *CloudVolume) GetWastedMonthlyCentsOk() (*int32, bool)`

GetWastedMonthlyCentsOk returns a tuple with the WastedMonthlyCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWastedMonthlyCents

`func (o *CloudVolume) SetWastedMonthlyCents(v int32)`

SetWastedMonthlyCents sets WastedMonthlyCents field to given value.

### HasWastedMonthlyCents

`func (o *CloudVolume) HasWastedMonthlyCents() bool`

HasWastedMonthlyCents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


