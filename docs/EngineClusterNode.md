# EngineClusterNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Gpus** | Pointer to [**[]EngineGPUDevice**](EngineGPUDevice.md) |  | [optional] 
**CpuCores** | Pointer to **int32** |  | [optional] 
**MemoryGb** | Pointer to **int32** |  | [optional] 
**Ip** | Pointer to **string** |  | [optional] 
**Labels** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewEngineClusterNode

`func NewEngineClusterNode() *EngineClusterNode`

NewEngineClusterNode instantiates a new EngineClusterNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEngineClusterNodeWithDefaults

`func NewEngineClusterNodeWithDefaults() *EngineClusterNode`

NewEngineClusterNodeWithDefaults instantiates a new EngineClusterNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EngineClusterNode) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EngineClusterNode) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EngineClusterNode) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EngineClusterNode) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *EngineClusterNode) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EngineClusterNode) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EngineClusterNode) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EngineClusterNode) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *EngineClusterNode) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EngineClusterNode) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EngineClusterNode) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EngineClusterNode) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetGpus

`func (o *EngineClusterNode) GetGpus() []EngineGPUDevice`

GetGpus returns the Gpus field if non-nil, zero value otherwise.

### GetGpusOk

`func (o *EngineClusterNode) GetGpusOk() (*[]EngineGPUDevice, bool)`

GetGpusOk returns a tuple with the Gpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpus

`func (o *EngineClusterNode) SetGpus(v []EngineGPUDevice)`

SetGpus sets Gpus field to given value.

### HasGpus

`func (o *EngineClusterNode) HasGpus() bool`

HasGpus returns a boolean if a field has been set.

### GetCpuCores

`func (o *EngineClusterNode) GetCpuCores() int32`

GetCpuCores returns the CpuCores field if non-nil, zero value otherwise.

### GetCpuCoresOk

`func (o *EngineClusterNode) GetCpuCoresOk() (*int32, bool)`

GetCpuCoresOk returns a tuple with the CpuCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuCores

`func (o *EngineClusterNode) SetCpuCores(v int32)`

SetCpuCores sets CpuCores field to given value.

### HasCpuCores

`func (o *EngineClusterNode) HasCpuCores() bool`

HasCpuCores returns a boolean if a field has been set.

### GetMemoryGb

`func (o *EngineClusterNode) GetMemoryGb() int32`

GetMemoryGb returns the MemoryGb field if non-nil, zero value otherwise.

### GetMemoryGbOk

`func (o *EngineClusterNode) GetMemoryGbOk() (*int32, bool)`

GetMemoryGbOk returns a tuple with the MemoryGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryGb

`func (o *EngineClusterNode) SetMemoryGb(v int32)`

SetMemoryGb sets MemoryGb field to given value.

### HasMemoryGb

`func (o *EngineClusterNode) HasMemoryGb() bool`

HasMemoryGb returns a boolean if a field has been set.

### GetIp

`func (o *EngineClusterNode) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *EngineClusterNode) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *EngineClusterNode) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *EngineClusterNode) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetLabels

`func (o *EngineClusterNode) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *EngineClusterNode) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *EngineClusterNode) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *EngineClusterNode) HasLabels() bool`

HasLabels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


