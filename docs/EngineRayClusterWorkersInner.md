# EngineRayClusterWorkersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupName** | Pointer to **string** |  | [optional] 
**Replicas** | Pointer to **int32** |  | [optional] 
**MinReplicas** | Pointer to **int32** |  | [optional] 
**MaxReplicas** | Pointer to **int32** |  | [optional] 
**Cpu** | Pointer to **int32** |  | [optional] 
**MemoryGb** | Pointer to **int32** |  | [optional] 
**GpuCount** | Pointer to **int32** |  | [optional] 
**GpuType** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewEngineRayClusterWorkersInner

`func NewEngineRayClusterWorkersInner() *EngineRayClusterWorkersInner`

NewEngineRayClusterWorkersInner instantiates a new EngineRayClusterWorkersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEngineRayClusterWorkersInnerWithDefaults

`func NewEngineRayClusterWorkersInnerWithDefaults() *EngineRayClusterWorkersInner`

NewEngineRayClusterWorkersInnerWithDefaults instantiates a new EngineRayClusterWorkersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupName

`func (o *EngineRayClusterWorkersInner) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *EngineRayClusterWorkersInner) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *EngineRayClusterWorkersInner) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *EngineRayClusterWorkersInner) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetReplicas

`func (o *EngineRayClusterWorkersInner) GetReplicas() int32`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *EngineRayClusterWorkersInner) GetReplicasOk() (*int32, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *EngineRayClusterWorkersInner) SetReplicas(v int32)`

SetReplicas sets Replicas field to given value.

### HasReplicas

`func (o *EngineRayClusterWorkersInner) HasReplicas() bool`

HasReplicas returns a boolean if a field has been set.

### GetMinReplicas

`func (o *EngineRayClusterWorkersInner) GetMinReplicas() int32`

GetMinReplicas returns the MinReplicas field if non-nil, zero value otherwise.

### GetMinReplicasOk

`func (o *EngineRayClusterWorkersInner) GetMinReplicasOk() (*int32, bool)`

GetMinReplicasOk returns a tuple with the MinReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinReplicas

`func (o *EngineRayClusterWorkersInner) SetMinReplicas(v int32)`

SetMinReplicas sets MinReplicas field to given value.

### HasMinReplicas

`func (o *EngineRayClusterWorkersInner) HasMinReplicas() bool`

HasMinReplicas returns a boolean if a field has been set.

### GetMaxReplicas

`func (o *EngineRayClusterWorkersInner) GetMaxReplicas() int32`

GetMaxReplicas returns the MaxReplicas field if non-nil, zero value otherwise.

### GetMaxReplicasOk

`func (o *EngineRayClusterWorkersInner) GetMaxReplicasOk() (*int32, bool)`

GetMaxReplicasOk returns a tuple with the MaxReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxReplicas

`func (o *EngineRayClusterWorkersInner) SetMaxReplicas(v int32)`

SetMaxReplicas sets MaxReplicas field to given value.

### HasMaxReplicas

`func (o *EngineRayClusterWorkersInner) HasMaxReplicas() bool`

HasMaxReplicas returns a boolean if a field has been set.

### GetCpu

`func (o *EngineRayClusterWorkersInner) GetCpu() int32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *EngineRayClusterWorkersInner) GetCpuOk() (*int32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *EngineRayClusterWorkersInner) SetCpu(v int32)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *EngineRayClusterWorkersInner) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetMemoryGb

`func (o *EngineRayClusterWorkersInner) GetMemoryGb() int32`

GetMemoryGb returns the MemoryGb field if non-nil, zero value otherwise.

### GetMemoryGbOk

`func (o *EngineRayClusterWorkersInner) GetMemoryGbOk() (*int32, bool)`

GetMemoryGbOk returns a tuple with the MemoryGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryGb

`func (o *EngineRayClusterWorkersInner) SetMemoryGb(v int32)`

SetMemoryGb sets MemoryGb field to given value.

### HasMemoryGb

`func (o *EngineRayClusterWorkersInner) HasMemoryGb() bool`

HasMemoryGb returns a boolean if a field has been set.

### GetGpuCount

`func (o *EngineRayClusterWorkersInner) GetGpuCount() int32`

GetGpuCount returns the GpuCount field if non-nil, zero value otherwise.

### GetGpuCountOk

`func (o *EngineRayClusterWorkersInner) GetGpuCountOk() (*int32, bool)`

GetGpuCountOk returns a tuple with the GpuCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuCount

`func (o *EngineRayClusterWorkersInner) SetGpuCount(v int32)`

SetGpuCount sets GpuCount field to given value.

### HasGpuCount

`func (o *EngineRayClusterWorkersInner) HasGpuCount() bool`

HasGpuCount returns a boolean if a field has been set.

### GetGpuType

`func (o *EngineRayClusterWorkersInner) GetGpuType() string`

GetGpuType returns the GpuType field if non-nil, zero value otherwise.

### GetGpuTypeOk

`func (o *EngineRayClusterWorkersInner) GetGpuTypeOk() (*string, bool)`

GetGpuTypeOk returns a tuple with the GpuType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuType

`func (o *EngineRayClusterWorkersInner) SetGpuType(v string)`

SetGpuType sets GpuType field to given value.

### HasGpuType

`func (o *EngineRayClusterWorkersInner) HasGpuType() bool`

HasGpuType returns a boolean if a field has been set.

### GetStatus

`func (o *EngineRayClusterWorkersInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EngineRayClusterWorkersInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EngineRayClusterWorkersInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EngineRayClusterWorkersInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


