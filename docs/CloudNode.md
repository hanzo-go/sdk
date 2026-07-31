# CloudNode

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

### NewCloudNode

`func NewCloudNode() *CloudNode`

NewCloudNode instantiates a new CloudNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudNodeWithDefaults

`func NewCloudNodeWithDefaults() *CloudNode`

NewCloudNodeWithDefaults instantiates a new CloudNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockedReason

`func (o *CloudNode) GetBlockedReason() string`

GetBlockedReason returns the BlockedReason field if non-nil, zero value otherwise.

### GetBlockedReasonOk

`func (o *CloudNode) GetBlockedReasonOk() (*string, bool)`

GetBlockedReasonOk returns a tuple with the BlockedReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedReason

`func (o *CloudNode) SetBlockedReason(v string)`

SetBlockedReason sets BlockedReason field to given value.

### HasBlockedReason

`func (o *CloudNode) HasBlockedReason() bool`

HasBlockedReason returns a boolean if a field has been set.

### GetCluster

`func (o *CloudNode) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *CloudNode) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *CloudNode) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *CloudNode) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetClusterId

`func (o *CloudNode) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *CloudNode) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *CloudNode) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *CloudNode) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CloudNode) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CloudNode) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CloudNode) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CloudNode) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *CloudNode) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudNode) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudNode) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CloudNode) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLocalDiskGiB

`func (o *CloudNode) GetLocalDiskGiB() int32`

GetLocalDiskGiB returns the LocalDiskGiB field if non-nil, zero value otherwise.

### GetLocalDiskGiBOk

`func (o *CloudNode) GetLocalDiskGiBOk() (*int32, bool)`

GetLocalDiskGiBOk returns a tuple with the LocalDiskGiB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalDiskGiB

`func (o *CloudNode) SetLocalDiskGiB(v int32)`

SetLocalDiskGiB sets LocalDiskGiB field to given value.

### HasLocalDiskGiB

`func (o *CloudNode) HasLocalDiskGiB() bool`

HasLocalDiskGiB returns a boolean if a field has been set.

### GetMemoryMiB

`func (o *CloudNode) GetMemoryMiB() int32`

GetMemoryMiB returns the MemoryMiB field if non-nil, zero value otherwise.

### GetMemoryMiBOk

`func (o *CloudNode) GetMemoryMiBOk() (*int32, bool)`

GetMemoryMiBOk returns a tuple with the MemoryMiB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryMiB

`func (o *CloudNode) SetMemoryMiB(v int32)`

SetMemoryMiB sets MemoryMiB field to given value.

### HasMemoryMiB

`func (o *CloudNode) HasMemoryMiB() bool`

HasMemoryMiB returns a boolean if a field has been set.

### GetMonthlyCents

`func (o *CloudNode) GetMonthlyCents() int32`

GetMonthlyCents returns the MonthlyCents field if non-nil, zero value otherwise.

### GetMonthlyCentsOk

`func (o *CloudNode) GetMonthlyCentsOk() (*int32, bool)`

GetMonthlyCentsOk returns a tuple with the MonthlyCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyCents

`func (o *CloudNode) SetMonthlyCents(v int32)`

SetMonthlyCents sets MonthlyCents field to given value.

### HasMonthlyCents

`func (o *CloudNode) HasMonthlyCents() bool`

HasMonthlyCents returns a boolean if a field has been set.

### GetMutable

`func (o *CloudNode) GetMutable() bool`

GetMutable returns the Mutable field if non-nil, zero value otherwise.

### GetMutableOk

`func (o *CloudNode) GetMutableOk() (*bool, bool)`

GetMutableOk returns a tuple with the Mutable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMutable

`func (o *CloudNode) SetMutable(v bool)`

SetMutable sets Mutable field to given value.

### HasMutable

`func (o *CloudNode) HasMutable() bool`

HasMutable returns a boolean if a field has been set.

### GetName

`func (o *CloudNode) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudNode) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudNode) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudNode) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPods

`func (o *CloudNode) GetPods() int32`

GetPods returns the Pods field if non-nil, zero value otherwise.

### GetPodsOk

`func (o *CloudNode) GetPodsOk() (*int32, bool)`

GetPodsOk returns a tuple with the Pods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPods

`func (o *CloudNode) SetPods(v int32)`

SetPods sets Pods field to given value.

### HasPods

`func (o *CloudNode) HasPods() bool`

HasPods returns a boolean if a field has been set.

### GetPrivateIp

`func (o *CloudNode) GetPrivateIp() string`

GetPrivateIp returns the PrivateIp field if non-nil, zero value otherwise.

### GetPrivateIpOk

`func (o *CloudNode) GetPrivateIpOk() (*string, bool)`

GetPrivateIpOk returns a tuple with the PrivateIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIp

`func (o *CloudNode) SetPrivateIp(v string)`

SetPrivateIp sets PrivateIp field to given value.

### HasPrivateIp

`func (o *CloudNode) HasPrivateIp() bool`

HasPrivateIp returns a boolean if a field has been set.

### GetPublicIp

`func (o *CloudNode) GetPublicIp() string`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *CloudNode) GetPublicIpOk() (*string, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIp

`func (o *CloudNode) SetPublicIp(v string)`

SetPublicIp sets PublicIp field to given value.

### HasPublicIp

`func (o *CloudNode) HasPublicIp() bool`

HasPublicIp returns a boolean if a field has been set.

### GetReady

`func (o *CloudNode) GetReady() bool`

GetReady returns the Ready field if non-nil, zero value otherwise.

### GetReadyOk

`func (o *CloudNode) GetReadyOk() (*bool, bool)`

GetReadyOk returns a tuple with the Ready field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReady

`func (o *CloudNode) SetReady(v bool)`

SetReady sets Ready field to given value.

### HasReady

`func (o *CloudNode) HasReady() bool`

HasReady returns a boolean if a field has been set.

### GetRegion

`func (o *CloudNode) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CloudNode) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CloudNode) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *CloudNode) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetSchedulable

`func (o *CloudNode) GetSchedulable() bool`

GetSchedulable returns the Schedulable field if non-nil, zero value otherwise.

### GetSchedulableOk

`func (o *CloudNode) GetSchedulableOk() (*bool, bool)`

GetSchedulableOk returns a tuple with the Schedulable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedulable

`func (o *CloudNode) SetSchedulable(v bool)`

SetSchedulable sets Schedulable field to given value.

### HasSchedulable

`func (o *CloudNode) HasSchedulable() bool`

HasSchedulable returns a boolean if a field has been set.

### GetSizeSlug

`func (o *CloudNode) GetSizeSlug() string`

GetSizeSlug returns the SizeSlug field if non-nil, zero value otherwise.

### GetSizeSlugOk

`func (o *CloudNode) GetSizeSlugOk() (*string, bool)`

GetSizeSlugOk returns a tuple with the SizeSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeSlug

`func (o *CloudNode) SetSizeSlug(v string)`

SetSizeSlug sets SizeSlug field to given value.

### HasSizeSlug

`func (o *CloudNode) HasSizeSlug() bool`

HasSizeSlug returns a boolean if a field has been set.

### GetStatus

`func (o *CloudNode) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudNode) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudNode) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudNode) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTags

`func (o *CloudNode) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CloudNode) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CloudNode) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CloudNode) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVcpus

`func (o *CloudNode) GetVcpus() int32`

GetVcpus returns the Vcpus field if non-nil, zero value otherwise.

### GetVcpusOk

`func (o *CloudNode) GetVcpusOk() (*int32, bool)`

GetVcpusOk returns a tuple with the Vcpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcpus

`func (o *CloudNode) SetVcpus(v int32)`

SetVcpus sets Vcpus field to given value.

### HasVcpus

`func (o *CloudNode) HasVcpus() bool`

HasVcpus returns a boolean if a field has been set.

### GetVolumes

`func (o *CloudNode) GetVolumes() int32`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *CloudNode) GetVolumesOk() (*int32, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *CloudNode) SetVolumes(v int32)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *CloudNode) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


