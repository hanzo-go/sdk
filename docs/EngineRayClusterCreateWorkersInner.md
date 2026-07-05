# EngineRayClusterCreateWorkersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupName** | **string** |  | 
**Replicas** | **int32** |  | 
**MinReplicas** | Pointer to **int32** |  | [optional] 
**MaxReplicas** | Pointer to **int32** |  | [optional] 
**Cpu** | Pointer to **int32** |  | [optional] [default to 8]
**MemoryGb** | Pointer to **int32** |  | [optional] [default to 32]
**GpuCount** | Pointer to **int32** |  | [optional] [default to 1]
**GpuType** | Pointer to **string** |  | [optional] 

## Methods

### NewEngineRayClusterCreateWorkersInner

`func NewEngineRayClusterCreateWorkersInner(groupName string, replicas int32, ) *EngineRayClusterCreateWorkersInner`

NewEngineRayClusterCreateWorkersInner instantiates a new EngineRayClusterCreateWorkersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEngineRayClusterCreateWorkersInnerWithDefaults

`func NewEngineRayClusterCreateWorkersInnerWithDefaults() *EngineRayClusterCreateWorkersInner`

NewEngineRayClusterCreateWorkersInnerWithDefaults instantiates a new EngineRayClusterCreateWorkersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupName

`func (o *EngineRayClusterCreateWorkersInner) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *EngineRayClusterCreateWorkersInner) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *EngineRayClusterCreateWorkersInner) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.


### GetReplicas

`func (o *EngineRayClusterCreateWorkersInner) GetReplicas() int32`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *EngineRayClusterCreateWorkersInner) GetReplicasOk() (*int32, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *EngineRayClusterCreateWorkersInner) SetReplicas(v int32)`

SetReplicas sets Replicas field to given value.


### GetMinReplicas

`func (o *EngineRayClusterCreateWorkersInner) GetMinReplicas() int32`

GetMinReplicas returns the MinReplicas field if non-nil, zero value otherwise.

### GetMinReplicasOk

`func (o *EngineRayClusterCreateWorkersInner) GetMinReplicasOk() (*int32, bool)`

GetMinReplicasOk returns a tuple with the MinReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinReplicas

`func (o *EngineRayClusterCreateWorkersInner) SetMinReplicas(v int32)`

SetMinReplicas sets MinReplicas field to given value.

### HasMinReplicas

`func (o *EngineRayClusterCreateWorkersInner) HasMinReplicas() bool`

HasMinReplicas returns a boolean if a field has been set.

### GetMaxReplicas

`func (o *EngineRayClusterCreateWorkersInner) GetMaxReplicas() int32`

GetMaxReplicas returns the MaxReplicas field if non-nil, zero value otherwise.

### GetMaxReplicasOk

`func (o *EngineRayClusterCreateWorkersInner) GetMaxReplicasOk() (*int32, bool)`

GetMaxReplicasOk returns a tuple with the MaxReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxReplicas

`func (o *EngineRayClusterCreateWorkersInner) SetMaxReplicas(v int32)`

SetMaxReplicas sets MaxReplicas field to given value.

### HasMaxReplicas

`func (o *EngineRayClusterCreateWorkersInner) HasMaxReplicas() bool`

HasMaxReplicas returns a boolean if a field has been set.

### GetCpu

`func (o *EngineRayClusterCreateWorkersInner) GetCpu() int32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *EngineRayClusterCreateWorkersInner) GetCpuOk() (*int32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *EngineRayClusterCreateWorkersInner) SetCpu(v int32)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *EngineRayClusterCreateWorkersInner) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetMemoryGb

`func (o *EngineRayClusterCreateWorkersInner) GetMemoryGb() int32`

GetMemoryGb returns the MemoryGb field if non-nil, zero value otherwise.

### GetMemoryGbOk

`func (o *EngineRayClusterCreateWorkersInner) GetMemoryGbOk() (*int32, bool)`

GetMemoryGbOk returns a tuple with the MemoryGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryGb

`func (o *EngineRayClusterCreateWorkersInner) SetMemoryGb(v int32)`

SetMemoryGb sets MemoryGb field to given value.

### HasMemoryGb

`func (o *EngineRayClusterCreateWorkersInner) HasMemoryGb() bool`

HasMemoryGb returns a boolean if a field has been set.

### GetGpuCount

`func (o *EngineRayClusterCreateWorkersInner) GetGpuCount() int32`

GetGpuCount returns the GpuCount field if non-nil, zero value otherwise.

### GetGpuCountOk

`func (o *EngineRayClusterCreateWorkersInner) GetGpuCountOk() (*int32, bool)`

GetGpuCountOk returns a tuple with the GpuCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuCount

`func (o *EngineRayClusterCreateWorkersInner) SetGpuCount(v int32)`

SetGpuCount sets GpuCount field to given value.

### HasGpuCount

`func (o *EngineRayClusterCreateWorkersInner) HasGpuCount() bool`

HasGpuCount returns a boolean if a field has been set.

### GetGpuType

`func (o *EngineRayClusterCreateWorkersInner) GetGpuType() string`

GetGpuType returns the GpuType field if non-nil, zero value otherwise.

### GetGpuTypeOk

`func (o *EngineRayClusterCreateWorkersInner) GetGpuTypeOk() (*string, bool)`

GetGpuTypeOk returns a tuple with the GpuType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuType

`func (o *EngineRayClusterCreateWorkersInner) SetGpuType(v string)`

SetGpuType sets GpuType field to given value.

### HasGpuType

`func (o *EngineRayClusterCreateWorkersInner) HasGpuType() bool`

HasGpuType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


