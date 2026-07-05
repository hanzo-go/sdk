# EngineJobResources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GpuCount** | **int32** |  | 
**GpuType** | Pointer to **string** | Required GPU model (e.g. A100, H100) | [optional] 
**CpuCores** | Pointer to **int32** |  | [optional] 
**MemoryGb** | Pointer to **int32** |  | [optional] 
**SharedMemoryGb** | Pointer to **int32** | Shared memory (/dev/shm) size | [optional] 
**StorageGb** | Pointer to **int32** |  | [optional] 

## Methods

### NewEngineJobResources

`func NewEngineJobResources(gpuCount int32, ) *EngineJobResources`

NewEngineJobResources instantiates a new EngineJobResources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEngineJobResourcesWithDefaults

`func NewEngineJobResourcesWithDefaults() *EngineJobResources`

NewEngineJobResourcesWithDefaults instantiates a new EngineJobResources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGpuCount

`func (o *EngineJobResources) GetGpuCount() int32`

GetGpuCount returns the GpuCount field if non-nil, zero value otherwise.

### GetGpuCountOk

`func (o *EngineJobResources) GetGpuCountOk() (*int32, bool)`

GetGpuCountOk returns a tuple with the GpuCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuCount

`func (o *EngineJobResources) SetGpuCount(v int32)`

SetGpuCount sets GpuCount field to given value.


### GetGpuType

`func (o *EngineJobResources) GetGpuType() string`

GetGpuType returns the GpuType field if non-nil, zero value otherwise.

### GetGpuTypeOk

`func (o *EngineJobResources) GetGpuTypeOk() (*string, bool)`

GetGpuTypeOk returns a tuple with the GpuType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuType

`func (o *EngineJobResources) SetGpuType(v string)`

SetGpuType sets GpuType field to given value.

### HasGpuType

`func (o *EngineJobResources) HasGpuType() bool`

HasGpuType returns a boolean if a field has been set.

### GetCpuCores

`func (o *EngineJobResources) GetCpuCores() int32`

GetCpuCores returns the CpuCores field if non-nil, zero value otherwise.

### GetCpuCoresOk

`func (o *EngineJobResources) GetCpuCoresOk() (*int32, bool)`

GetCpuCoresOk returns a tuple with the CpuCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuCores

`func (o *EngineJobResources) SetCpuCores(v int32)`

SetCpuCores sets CpuCores field to given value.

### HasCpuCores

`func (o *EngineJobResources) HasCpuCores() bool`

HasCpuCores returns a boolean if a field has been set.

### GetMemoryGb

`func (o *EngineJobResources) GetMemoryGb() int32`

GetMemoryGb returns the MemoryGb field if non-nil, zero value otherwise.

### GetMemoryGbOk

`func (o *EngineJobResources) GetMemoryGbOk() (*int32, bool)`

GetMemoryGbOk returns a tuple with the MemoryGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryGb

`func (o *EngineJobResources) SetMemoryGb(v int32)`

SetMemoryGb sets MemoryGb field to given value.

### HasMemoryGb

`func (o *EngineJobResources) HasMemoryGb() bool`

HasMemoryGb returns a boolean if a field has been set.

### GetSharedMemoryGb

`func (o *EngineJobResources) GetSharedMemoryGb() int32`

GetSharedMemoryGb returns the SharedMemoryGb field if non-nil, zero value otherwise.

### GetSharedMemoryGbOk

`func (o *EngineJobResources) GetSharedMemoryGbOk() (*int32, bool)`

GetSharedMemoryGbOk returns a tuple with the SharedMemoryGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedMemoryGb

`func (o *EngineJobResources) SetSharedMemoryGb(v int32)`

SetSharedMemoryGb sets SharedMemoryGb field to given value.

### HasSharedMemoryGb

`func (o *EngineJobResources) HasSharedMemoryGb() bool`

HasSharedMemoryGb returns a boolean if a field has been set.

### GetStorageGb

`func (o *EngineJobResources) GetStorageGb() int32`

GetStorageGb returns the StorageGb field if non-nil, zero value otherwise.

### GetStorageGbOk

`func (o *EngineJobResources) GetStorageGbOk() (*int32, bool)`

GetStorageGbOk returns a tuple with the StorageGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageGb

`func (o *EngineJobResources) SetStorageGb(v int32)`

SetStorageGb sets StorageGb field to given value.

### HasStorageGb

`func (o *EngineJobResources) HasStorageGb() bool`

HasStorageGb returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


