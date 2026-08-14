# Volume

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

### NewVolume

`func NewVolume() *Volume`

NewVolume instantiates a new Volume object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeWithDefaults

`func NewVolumeWithDefaults() *Volume`

NewVolumeWithDefaults instantiates a new Volume object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockedReason

`func (o *Volume) GetBlockedReason() string`

GetBlockedReason returns the BlockedReason field if non-nil, zero value otherwise.

### GetBlockedReasonOk

`func (o *Volume) GetBlockedReasonOk() (*string, bool)`

GetBlockedReasonOk returns a tuple with the BlockedReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedReason

`func (o *Volume) SetBlockedReason(v string)`

SetBlockedReason sets BlockedReason field to given value.

### HasBlockedReason

`func (o *Volume) HasBlockedReason() bool`

HasBlockedReason returns a boolean if a field has been set.

### GetCluster

`func (o *Volume) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *Volume) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *Volume) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *Volume) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetClusterId

`func (o *Volume) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *Volume) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *Volume) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *Volume) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### GetController

`func (o *Volume) GetController() string`

GetController returns the Controller field if non-nil, zero value otherwise.

### GetControllerOk

`func (o *Volume) GetControllerOk() (*string, bool)`

GetControllerOk returns a tuple with the Controller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetController

`func (o *Volume) SetController(v string)`

SetController sets Controller field to given value.

### HasController

`func (o *Volume) HasController() bool`

HasController returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Volume) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Volume) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Volume) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Volume) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeletable

`func (o *Volume) GetDeletable() bool`

GetDeletable returns the Deletable field if non-nil, zero value otherwise.

### GetDeletableOk

`func (o *Volume) GetDeletableOk() (*bool, bool)`

GetDeletableOk returns a tuple with the Deletable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletable

`func (o *Volume) SetDeletable(v bool)`

SetDeletable sets Deletable field to given value.

### HasDeletable

`func (o *Volume) HasDeletable() bool`

HasDeletable returns a boolean if a field has been set.

### GetDropletIds

`func (o *Volume) GetDropletIds() []int32`

GetDropletIds returns the DropletIds field if non-nil, zero value otherwise.

### GetDropletIdsOk

`func (o *Volume) GetDropletIdsOk() (*[]int32, bool)`

GetDropletIdsOk returns a tuple with the DropletIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDropletIds

`func (o *Volume) SetDropletIds(v []int32)`

SetDropletIds sets DropletIds field to given value.

### HasDropletIds

`func (o *Volume) HasDropletIds() bool`

HasDropletIds returns a boolean if a field has been set.

### GetExpandBlockedReason

`func (o *Volume) GetExpandBlockedReason() string`

GetExpandBlockedReason returns the ExpandBlockedReason field if non-nil, zero value otherwise.

### GetExpandBlockedReasonOk

`func (o *Volume) GetExpandBlockedReasonOk() (*string, bool)`

GetExpandBlockedReasonOk returns a tuple with the ExpandBlockedReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpandBlockedReason

`func (o *Volume) SetExpandBlockedReason(v string)`

SetExpandBlockedReason sets ExpandBlockedReason field to given value.

### HasExpandBlockedReason

`func (o *Volume) HasExpandBlockedReason() bool`

HasExpandBlockedReason returns a boolean if a field has been set.

### GetExpandable

`func (o *Volume) GetExpandable() bool`

GetExpandable returns the Expandable field if non-nil, zero value otherwise.

### GetExpandableOk

`func (o *Volume) GetExpandableOk() (*bool, bool)`

GetExpandableOk returns a tuple with the Expandable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpandable

`func (o *Volume) SetExpandable(v bool)`

SetExpandable sets Expandable field to given value.

### HasExpandable

`func (o *Volume) HasExpandable() bool`

HasExpandable returns a boolean if a field has been set.

### GetHasUsage

`func (o *Volume) GetHasUsage() bool`

GetHasUsage returns the HasUsage field if non-nil, zero value otherwise.

### GetHasUsageOk

`func (o *Volume) GetHasUsageOk() (*bool, bool)`

GetHasUsageOk returns a tuple with the HasUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasUsage

`func (o *Volume) SetHasUsage(v bool)`

SetHasUsage sets HasUsage field to given value.

### HasHasUsage

`func (o *Volume) HasHasUsage() bool`

HasHasUsage returns a boolean if a field has been set.

### GetId

`func (o *Volume) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Volume) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Volume) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Volume) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdle

`func (o *Volume) GetIdle() bool`

GetIdle returns the Idle field if non-nil, zero value otherwise.

### GetIdleOk

`func (o *Volume) GetIdleOk() (*bool, bool)`

GetIdleOk returns a tuple with the Idle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdle

`func (o *Volume) SetIdle(v bool)`

SetIdle sets Idle field to given value.

### HasIdle

`func (o *Volume) HasIdle() bool`

HasIdle returns a boolean if a field has been set.

### GetMonthlyCents

`func (o *Volume) GetMonthlyCents() int32`

GetMonthlyCents returns the MonthlyCents field if non-nil, zero value otherwise.

### GetMonthlyCentsOk

`func (o *Volume) GetMonthlyCentsOk() (*int32, bool)`

GetMonthlyCentsOk returns a tuple with the MonthlyCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyCents

`func (o *Volume) SetMonthlyCents(v int32)`

SetMonthlyCents sets MonthlyCents field to given value.

### HasMonthlyCents

`func (o *Volume) HasMonthlyCents() bool`

HasMonthlyCents returns a boolean if a field has been set.

### GetMountedBy

`func (o *Volume) GetMountedBy() []string`

GetMountedBy returns the MountedBy field if non-nil, zero value otherwise.

### GetMountedByOk

`func (o *Volume) GetMountedByOk() (*[]string, bool)`

GetMountedByOk returns a tuple with the MountedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountedBy

`func (o *Volume) SetMountedBy(v []string)`

SetMountedBy sets MountedBy field to given value.

### HasMountedBy

`func (o *Volume) HasMountedBy() bool`

HasMountedBy returns a boolean if a field has been set.

### GetName

`func (o *Volume) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Volume) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Volume) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Volume) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNodeName

`func (o *Volume) GetNodeName() string`

GetNodeName returns the NodeName field if non-nil, zero value otherwise.

### GetNodeNameOk

`func (o *Volume) GetNodeNameOk() (*string, bool)`

GetNodeNameOk returns a tuple with the NodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeName

`func (o *Volume) SetNodeName(v string)`

SetNodeName sets NodeName field to given value.

### HasNodeName

`func (o *Volume) HasNodeName() bool`

HasNodeName returns a boolean if a field has been set.

### GetPv

`func (o *Volume) GetPv() string`

GetPv returns the Pv field if non-nil, zero value otherwise.

### GetPvOk

`func (o *Volume) GetPvOk() (*string, bool)`

GetPvOk returns a tuple with the Pv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPv

`func (o *Volume) SetPv(v string)`

SetPv sets Pv field to given value.

### HasPv

`func (o *Volume) HasPv() bool`

HasPv returns a boolean if a field has been set.

### GetPvPhase

`func (o *Volume) GetPvPhase() string`

GetPvPhase returns the PvPhase field if non-nil, zero value otherwise.

### GetPvPhaseOk

`func (o *Volume) GetPvPhaseOk() (*string, bool)`

GetPvPhaseOk returns a tuple with the PvPhase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPvPhase

`func (o *Volume) SetPvPhase(v string)`

SetPvPhase sets PvPhase field to given value.

### HasPvPhase

`func (o *Volume) HasPvPhase() bool`

HasPvPhase returns a boolean if a field has been set.

### GetPvcName

`func (o *Volume) GetPvcName() string`

GetPvcName returns the PvcName field if non-nil, zero value otherwise.

### GetPvcNameOk

`func (o *Volume) GetPvcNameOk() (*string, bool)`

GetPvcNameOk returns a tuple with the PvcName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPvcName

`func (o *Volume) SetPvcName(v string)`

SetPvcName sets PvcName field to given value.

### HasPvcName

`func (o *Volume) HasPvcName() bool`

HasPvcName returns a boolean if a field has been set.

### GetPvcNamespace

`func (o *Volume) GetPvcNamespace() string`

GetPvcNamespace returns the PvcNamespace field if non-nil, zero value otherwise.

### GetPvcNamespaceOk

`func (o *Volume) GetPvcNamespaceOk() (*string, bool)`

GetPvcNamespaceOk returns a tuple with the PvcNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPvcNamespace

`func (o *Volume) SetPvcNamespace(v string)`

SetPvcNamespace sets PvcNamespace field to given value.

### HasPvcNamespace

`func (o *Volume) HasPvcNamespace() bool`

HasPvcNamespace returns a boolean if a field has been set.

### GetRegion

`func (o *Volume) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *Volume) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *Volume) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *Volume) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetSizeGiB

`func (o *Volume) GetSizeGiB() int32`

GetSizeGiB returns the SizeGiB field if non-nil, zero value otherwise.

### GetSizeGiBOk

`func (o *Volume) GetSizeGiBOk() (*int32, bool)`

GetSizeGiBOk returns a tuple with the SizeGiB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeGiB

`func (o *Volume) SetSizeGiB(v int32)`

SetSizeGiB sets SizeGiB field to given value.

### HasSizeGiB

`func (o *Volume) HasSizeGiB() bool`

HasSizeGiB returns a boolean if a field has been set.

### GetState

`func (o *Volume) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Volume) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Volume) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Volume) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTagCluster

`func (o *Volume) GetTagCluster() string`

GetTagCluster returns the TagCluster field if non-nil, zero value otherwise.

### GetTagClusterOk

`func (o *Volume) GetTagClusterOk() (*string, bool)`

GetTagClusterOk returns a tuple with the TagCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagCluster

`func (o *Volume) SetTagCluster(v string)`

SetTagCluster sets TagCluster field to given value.

### HasTagCluster

`func (o *Volume) HasTagCluster() bool`

HasTagCluster returns a boolean if a field has been set.

### GetUsedBytes

`func (o *Volume) GetUsedBytes() int32`

GetUsedBytes returns the UsedBytes field if non-nil, zero value otherwise.

### GetUsedBytesOk

`func (o *Volume) GetUsedBytesOk() (*int32, bool)`

GetUsedBytesOk returns a tuple with the UsedBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedBytes

`func (o *Volume) SetUsedBytes(v int32)`

SetUsedBytes sets UsedBytes field to given value.

### HasUsedBytes

`func (o *Volume) HasUsedBytes() bool`

HasUsedBytes returns a boolean if a field has been set.

### GetWastedGiB

`func (o *Volume) GetWastedGiB() int32`

GetWastedGiB returns the WastedGiB field if non-nil, zero value otherwise.

### GetWastedGiBOk

`func (o *Volume) GetWastedGiBOk() (*int32, bool)`

GetWastedGiBOk returns a tuple with the WastedGiB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWastedGiB

`func (o *Volume) SetWastedGiB(v int32)`

SetWastedGiB sets WastedGiB field to given value.

### HasWastedGiB

`func (o *Volume) HasWastedGiB() bool`

HasWastedGiB returns a boolean if a field has been set.

### GetWastedMonthlyCents

`func (o *Volume) GetWastedMonthlyCents() int32`

GetWastedMonthlyCents returns the WastedMonthlyCents field if non-nil, zero value otherwise.

### GetWastedMonthlyCentsOk

`func (o *Volume) GetWastedMonthlyCentsOk() (*int32, bool)`

GetWastedMonthlyCentsOk returns a tuple with the WastedMonthlyCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWastedMonthlyCents

`func (o *Volume) SetWastedMonthlyCents(v int32)`

SetWastedMonthlyCents sets WastedMonthlyCents field to given value.

### HasWastedMonthlyCents

`func (o *Volume) HasWastedMonthlyCents() bool`

HasWastedMonthlyCents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


