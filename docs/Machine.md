# Machine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockedReason** | Pointer to **string** |  | [optional] 
**Cluster** | Pointer to **string** |  | [optional] 
**ClusterId** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**LocalDiskGiB** | Pointer to **int32** |  | [optional] 
**MemoryMiB** | Pointer to **int32** |  | [optional] 
**MonthlyCents** | Pointer to **int32** |  | [optional] 
**Mutable** | Pointer to **bool** | Mutable reports whether this droplet may be changed DIRECTLY — deleted or resized. One predicate covers both because one fact decides both: a DOKS node belongs to a node pool, and the pool is the only thing allowed to change it. | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Pods** | Pointer to **int32** |  | [optional] 
**PrivateIp** | Pointer to **string** |  | [optional] 
**PublicIp** | Pointer to **string** |  | [optional] 
**Ready** | Pointer to **bool** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**Schedulable** | Pointer to **bool** |  | [optional] 
**SizeSlug** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Vcpus** | Pointer to **int32** |  | [optional] 
**Volumes** | Pointer to **int32** |  | [optional] 

## Methods

### NewMachine

`func NewMachine() *Machine`

NewMachine instantiates a new Machine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMachineWithDefaults

`func NewMachineWithDefaults() *Machine`

NewMachineWithDefaults instantiates a new Machine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockedReason

`func (o *Machine) GetBlockedReason() string`

GetBlockedReason returns the BlockedReason field if non-nil, zero value otherwise.

### GetBlockedReasonOk

`func (o *Machine) GetBlockedReasonOk() (*string, bool)`

GetBlockedReasonOk returns a tuple with the BlockedReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedReason

`func (o *Machine) SetBlockedReason(v string)`

SetBlockedReason sets BlockedReason field to given value.

### HasBlockedReason

`func (o *Machine) HasBlockedReason() bool`

HasBlockedReason returns a boolean if a field has been set.

### GetCluster

`func (o *Machine) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *Machine) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *Machine) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *Machine) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetClusterId

`func (o *Machine) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *Machine) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *Machine) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *Machine) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Machine) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Machine) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Machine) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Machine) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *Machine) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Machine) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Machine) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Machine) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLocalDiskGiB

`func (o *Machine) GetLocalDiskGiB() int32`

GetLocalDiskGiB returns the LocalDiskGiB field if non-nil, zero value otherwise.

### GetLocalDiskGiBOk

`func (o *Machine) GetLocalDiskGiBOk() (*int32, bool)`

GetLocalDiskGiBOk returns a tuple with the LocalDiskGiB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalDiskGiB

`func (o *Machine) SetLocalDiskGiB(v int32)`

SetLocalDiskGiB sets LocalDiskGiB field to given value.

### HasLocalDiskGiB

`func (o *Machine) HasLocalDiskGiB() bool`

HasLocalDiskGiB returns a boolean if a field has been set.

### GetMemoryMiB

`func (o *Machine) GetMemoryMiB() int32`

GetMemoryMiB returns the MemoryMiB field if non-nil, zero value otherwise.

### GetMemoryMiBOk

`func (o *Machine) GetMemoryMiBOk() (*int32, bool)`

GetMemoryMiBOk returns a tuple with the MemoryMiB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryMiB

`func (o *Machine) SetMemoryMiB(v int32)`

SetMemoryMiB sets MemoryMiB field to given value.

### HasMemoryMiB

`func (o *Machine) HasMemoryMiB() bool`

HasMemoryMiB returns a boolean if a field has been set.

### GetMonthlyCents

`func (o *Machine) GetMonthlyCents() int32`

GetMonthlyCents returns the MonthlyCents field if non-nil, zero value otherwise.

### GetMonthlyCentsOk

`func (o *Machine) GetMonthlyCentsOk() (*int32, bool)`

GetMonthlyCentsOk returns a tuple with the MonthlyCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyCents

`func (o *Machine) SetMonthlyCents(v int32)`

SetMonthlyCents sets MonthlyCents field to given value.

### HasMonthlyCents

`func (o *Machine) HasMonthlyCents() bool`

HasMonthlyCents returns a boolean if a field has been set.

### GetMutable

`func (o *Machine) GetMutable() bool`

GetMutable returns the Mutable field if non-nil, zero value otherwise.

### GetMutableOk

`func (o *Machine) GetMutableOk() (*bool, bool)`

GetMutableOk returns a tuple with the Mutable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMutable

`func (o *Machine) SetMutable(v bool)`

SetMutable sets Mutable field to given value.

### HasMutable

`func (o *Machine) HasMutable() bool`

HasMutable returns a boolean if a field has been set.

### GetName

`func (o *Machine) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Machine) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Machine) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Machine) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPods

`func (o *Machine) GetPods() int32`

GetPods returns the Pods field if non-nil, zero value otherwise.

### GetPodsOk

`func (o *Machine) GetPodsOk() (*int32, bool)`

GetPodsOk returns a tuple with the Pods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPods

`func (o *Machine) SetPods(v int32)`

SetPods sets Pods field to given value.

### HasPods

`func (o *Machine) HasPods() bool`

HasPods returns a boolean if a field has been set.

### GetPrivateIp

`func (o *Machine) GetPrivateIp() string`

GetPrivateIp returns the PrivateIp field if non-nil, zero value otherwise.

### GetPrivateIpOk

`func (o *Machine) GetPrivateIpOk() (*string, bool)`

GetPrivateIpOk returns a tuple with the PrivateIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIp

`func (o *Machine) SetPrivateIp(v string)`

SetPrivateIp sets PrivateIp field to given value.

### HasPrivateIp

`func (o *Machine) HasPrivateIp() bool`

HasPrivateIp returns a boolean if a field has been set.

### GetPublicIp

`func (o *Machine) GetPublicIp() string`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *Machine) GetPublicIpOk() (*string, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIp

`func (o *Machine) SetPublicIp(v string)`

SetPublicIp sets PublicIp field to given value.

### HasPublicIp

`func (o *Machine) HasPublicIp() bool`

HasPublicIp returns a boolean if a field has been set.

### GetReady

`func (o *Machine) GetReady() bool`

GetReady returns the Ready field if non-nil, zero value otherwise.

### GetReadyOk

`func (o *Machine) GetReadyOk() (*bool, bool)`

GetReadyOk returns a tuple with the Ready field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReady

`func (o *Machine) SetReady(v bool)`

SetReady sets Ready field to given value.

### HasReady

`func (o *Machine) HasReady() bool`

HasReady returns a boolean if a field has been set.

### GetRegion

`func (o *Machine) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *Machine) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *Machine) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *Machine) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetSchedulable

`func (o *Machine) GetSchedulable() bool`

GetSchedulable returns the Schedulable field if non-nil, zero value otherwise.

### GetSchedulableOk

`func (o *Machine) GetSchedulableOk() (*bool, bool)`

GetSchedulableOk returns a tuple with the Schedulable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedulable

`func (o *Machine) SetSchedulable(v bool)`

SetSchedulable sets Schedulable field to given value.

### HasSchedulable

`func (o *Machine) HasSchedulable() bool`

HasSchedulable returns a boolean if a field has been set.

### GetSizeSlug

`func (o *Machine) GetSizeSlug() string`

GetSizeSlug returns the SizeSlug field if non-nil, zero value otherwise.

### GetSizeSlugOk

`func (o *Machine) GetSizeSlugOk() (*string, bool)`

GetSizeSlugOk returns a tuple with the SizeSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeSlug

`func (o *Machine) SetSizeSlug(v string)`

SetSizeSlug sets SizeSlug field to given value.

### HasSizeSlug

`func (o *Machine) HasSizeSlug() bool`

HasSizeSlug returns a boolean if a field has been set.

### GetStatus

`func (o *Machine) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Machine) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Machine) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Machine) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTags

`func (o *Machine) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Machine) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Machine) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Machine) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVcpus

`func (o *Machine) GetVcpus() int32`

GetVcpus returns the Vcpus field if non-nil, zero value otherwise.

### GetVcpusOk

`func (o *Machine) GetVcpusOk() (*int32, bool)`

GetVcpusOk returns a tuple with the Vcpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcpus

`func (o *Machine) SetVcpus(v int32)`

SetVcpus sets Vcpus field to given value.

### HasVcpus

`func (o *Machine) HasVcpus() bool`

HasVcpus returns a boolean if a field has been set.

### GetVolumes

`func (o *Machine) GetVolumes() int32`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *Machine) GetVolumesOk() (*int32, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *Machine) SetVolumes(v int32)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *Machine) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


