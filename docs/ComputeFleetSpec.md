# ComputeFleetSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Arch** | Pointer to **string** | Arch is the CPU architecture, amd64 or arm64, and it is what decides whether a binary built for the fleet will run here. Only the sources that report one carry it (a linked run-target, a BYO worker). | [optional] 
**Cpus** | Pointer to **int64** | CPUs is logical cores on the unit. | [optional] 
**GpuModel** | Pointer to **string** | GPUModel names the FIRST accelerator (\&quot;NVIDIA GB10\&quot;) as the representative of the set; GPUs carries how many. Empty for a cluster, whose cards are counted rather than modelled, and for a unit with none. | [optional] 
**Gpus** | Pointer to **int64** | GPUs is how many accelerators the unit has. For a cluster it is the vendor totals summed across every node, so it counts cards, not machines. | [optional] 
**Memory** | Pointer to **int64** | Memory is total system RAM in BYTES — not GB, and not what is free right now (fleetMetrics carries that). Absent when the source reports no RAM figure. | [optional] 
**Os** | Pointer to **string** | OS is the operating system the unit runs: linux, darwin or windows. Empty when the source does not report one — a cluster row does not. | [optional] 

## Methods

### NewComputeFleetSpec

`func NewComputeFleetSpec() *ComputeFleetSpec`

NewComputeFleetSpec instantiates a new ComputeFleetSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputeFleetSpecWithDefaults

`func NewComputeFleetSpecWithDefaults() *ComputeFleetSpec`

NewComputeFleetSpecWithDefaults instantiates a new ComputeFleetSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArch

`func (o *ComputeFleetSpec) GetArch() string`

GetArch returns the Arch field if non-nil, zero value otherwise.

### GetArchOk

`func (o *ComputeFleetSpec) GetArchOk() (*string, bool)`

GetArchOk returns a tuple with the Arch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArch

`func (o *ComputeFleetSpec) SetArch(v string)`

SetArch sets Arch field to given value.

### HasArch

`func (o *ComputeFleetSpec) HasArch() bool`

HasArch returns a boolean if a field has been set.

### GetCpus

`func (o *ComputeFleetSpec) GetCpus() int64`

GetCpus returns the Cpus field if non-nil, zero value otherwise.

### GetCpusOk

`func (o *ComputeFleetSpec) GetCpusOk() (*int64, bool)`

GetCpusOk returns a tuple with the Cpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpus

`func (o *ComputeFleetSpec) SetCpus(v int64)`

SetCpus sets Cpus field to given value.

### HasCpus

`func (o *ComputeFleetSpec) HasCpus() bool`

HasCpus returns a boolean if a field has been set.

### GetGpuModel

`func (o *ComputeFleetSpec) GetGpuModel() string`

GetGpuModel returns the GpuModel field if non-nil, zero value otherwise.

### GetGpuModelOk

`func (o *ComputeFleetSpec) GetGpuModelOk() (*string, bool)`

GetGpuModelOk returns a tuple with the GpuModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuModel

`func (o *ComputeFleetSpec) SetGpuModel(v string)`

SetGpuModel sets GpuModel field to given value.

### HasGpuModel

`func (o *ComputeFleetSpec) HasGpuModel() bool`

HasGpuModel returns a boolean if a field has been set.

### GetGpus

`func (o *ComputeFleetSpec) GetGpus() int64`

GetGpus returns the Gpus field if non-nil, zero value otherwise.

### GetGpusOk

`func (o *ComputeFleetSpec) GetGpusOk() (*int64, bool)`

GetGpusOk returns a tuple with the Gpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpus

`func (o *ComputeFleetSpec) SetGpus(v int64)`

SetGpus sets Gpus field to given value.

### HasGpus

`func (o *ComputeFleetSpec) HasGpus() bool`

HasGpus returns a boolean if a field has been set.

### GetMemory

`func (o *ComputeFleetSpec) GetMemory() int64`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *ComputeFleetSpec) GetMemoryOk() (*int64, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *ComputeFleetSpec) SetMemory(v int64)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *ComputeFleetSpec) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetOs

`func (o *ComputeFleetSpec) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *ComputeFleetSpec) GetOsOk() (*string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *ComputeFleetSpec) SetOs(v string)`

SetOs sets Os field to given value.

### HasOs

`func (o *ComputeFleetSpec) HasOs() bool`

HasOs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


