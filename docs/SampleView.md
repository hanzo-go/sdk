# SampleView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**At** | Pointer to **string** | At is when the reading was MEASURED, RFC 3339 in UTC — the x-axis a chart plots against. The series is returned oldest first, so it only increases. | [optional] 
**CostCents** | Pointer to **int64** | CostCents is what this unit resold for over the hour the reading falls in, in whole US cents. 0 means UNPRICED, not free: the operator&#39;s own machines — a linked run-target, a dialed-in BYO worker — are metered for utilization and never resold, so only a priced source ever fills it. | [optional] 
**Cpus** | Pointer to **int64** | CPUs is logical cores. The static capability rides every row on purpose: a chart can size load against cores without joining a registry whose row may since have been rewritten or the unit deregistered. | [optional] 
**GpuModel** | Pointer to **string** | GPUModel names the representative accelerator (\&quot;GB10\&quot;); GPUs carries how many. | [optional] 
**GpuUtil** | Pointer to **float64** | GPUUtil is aggregate accelerator utilization as a FRACTION of 1 — 0.42 is 42% busy. Anything a reporter sends outside 0..1 is clamped into it on write. | [optional] 
**Gpus** | Pointer to **int64** | GPUs is how many accelerators the reading covers. | [optional] 
**Host** | Pointer to **string** | Host is the hostname the unit reported at the time of the reading. | [optional] 
**Kind** | Pointer to **string** | Kind is what the measured unit is: laptop, cloud, gpu, cluster, machine or worker. | [optional] 
**Load1** | Pointer to **float64** | Load1 is the 1-minute load average — runnable processes, not a percentage. | [optional] 
**Load15** | Pointer to **float64** | Load15 is the 15-minute load average, the same units as Load1. | [optional] 
**Load5** | Pointer to **float64** | Load5 is the 5-minute load average, the same units as Load1. | [optional] 
**MemFree** | Pointer to **int64** | MemFree is host memory available, in BYTES, as reported rather than derived. | [optional] 
**MemUsed** | Pointer to **int64** | MemUsed is host memory in use, in BYTES. | [optional] 
**Memory** | Pointer to **int64** | Memory is total system RAM in BYTES at the time of the reading. | [optional] 
**Source** | Pointer to **string** | Source is the plane that reported the reading: \&quot;agent\&quot;, \&quot;byo\&quot; or \&quot;visor\&quot; — the same vocabulary the board&#39;s rows carry, and what ?source&#x3D; narrows on. | [optional] 
**Unit** | Pointer to **string** | Unit is the source&#39;s own id for the measured unit. With Source it is the key the chart groups by, and the key the board joins a unit&#39;s latest reading on. | [optional] 

## Methods

### NewSampleView

`func NewSampleView() *SampleView`

NewSampleView instantiates a new SampleView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSampleViewWithDefaults

`func NewSampleViewWithDefaults() *SampleView`

NewSampleViewWithDefaults instantiates a new SampleView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAt

`func (o *SampleView) GetAt() string`

GetAt returns the At field if non-nil, zero value otherwise.

### GetAtOk

`func (o *SampleView) GetAtOk() (*string, bool)`

GetAtOk returns a tuple with the At field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAt

`func (o *SampleView) SetAt(v string)`

SetAt sets At field to given value.

### HasAt

`func (o *SampleView) HasAt() bool`

HasAt returns a boolean if a field has been set.

### GetCostCents

`func (o *SampleView) GetCostCents() int64`

GetCostCents returns the CostCents field if non-nil, zero value otherwise.

### GetCostCentsOk

`func (o *SampleView) GetCostCentsOk() (*int64, bool)`

GetCostCentsOk returns a tuple with the CostCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostCents

`func (o *SampleView) SetCostCents(v int64)`

SetCostCents sets CostCents field to given value.

### HasCostCents

`func (o *SampleView) HasCostCents() bool`

HasCostCents returns a boolean if a field has been set.

### GetCpus

`func (o *SampleView) GetCpus() int64`

GetCpus returns the Cpus field if non-nil, zero value otherwise.

### GetCpusOk

`func (o *SampleView) GetCpusOk() (*int64, bool)`

GetCpusOk returns a tuple with the Cpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpus

`func (o *SampleView) SetCpus(v int64)`

SetCpus sets Cpus field to given value.

### HasCpus

`func (o *SampleView) HasCpus() bool`

HasCpus returns a boolean if a field has been set.

### GetGpuModel

`func (o *SampleView) GetGpuModel() string`

GetGpuModel returns the GpuModel field if non-nil, zero value otherwise.

### GetGpuModelOk

`func (o *SampleView) GetGpuModelOk() (*string, bool)`

GetGpuModelOk returns a tuple with the GpuModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuModel

`func (o *SampleView) SetGpuModel(v string)`

SetGpuModel sets GpuModel field to given value.

### HasGpuModel

`func (o *SampleView) HasGpuModel() bool`

HasGpuModel returns a boolean if a field has been set.

### GetGpuUtil

`func (o *SampleView) GetGpuUtil() float64`

GetGpuUtil returns the GpuUtil field if non-nil, zero value otherwise.

### GetGpuUtilOk

`func (o *SampleView) GetGpuUtilOk() (*float64, bool)`

GetGpuUtilOk returns a tuple with the GpuUtil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuUtil

`func (o *SampleView) SetGpuUtil(v float64)`

SetGpuUtil sets GpuUtil field to given value.

### HasGpuUtil

`func (o *SampleView) HasGpuUtil() bool`

HasGpuUtil returns a boolean if a field has been set.

### GetGpus

`func (o *SampleView) GetGpus() int64`

GetGpus returns the Gpus field if non-nil, zero value otherwise.

### GetGpusOk

`func (o *SampleView) GetGpusOk() (*int64, bool)`

GetGpusOk returns a tuple with the Gpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpus

`func (o *SampleView) SetGpus(v int64)`

SetGpus sets Gpus field to given value.

### HasGpus

`func (o *SampleView) HasGpus() bool`

HasGpus returns a boolean if a field has been set.

### GetHost

`func (o *SampleView) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *SampleView) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *SampleView) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *SampleView) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetKind

`func (o *SampleView) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *SampleView) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *SampleView) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *SampleView) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLoad1

`func (o *SampleView) GetLoad1() float64`

GetLoad1 returns the Load1 field if non-nil, zero value otherwise.

### GetLoad1Ok

`func (o *SampleView) GetLoad1Ok() (*float64, bool)`

GetLoad1Ok returns a tuple with the Load1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoad1

`func (o *SampleView) SetLoad1(v float64)`

SetLoad1 sets Load1 field to given value.

### HasLoad1

`func (o *SampleView) HasLoad1() bool`

HasLoad1 returns a boolean if a field has been set.

### GetLoad15

`func (o *SampleView) GetLoad15() float64`

GetLoad15 returns the Load15 field if non-nil, zero value otherwise.

### GetLoad15Ok

`func (o *SampleView) GetLoad15Ok() (*float64, bool)`

GetLoad15Ok returns a tuple with the Load15 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoad15

`func (o *SampleView) SetLoad15(v float64)`

SetLoad15 sets Load15 field to given value.

### HasLoad15

`func (o *SampleView) HasLoad15() bool`

HasLoad15 returns a boolean if a field has been set.

### GetLoad5

`func (o *SampleView) GetLoad5() float64`

GetLoad5 returns the Load5 field if non-nil, zero value otherwise.

### GetLoad5Ok

`func (o *SampleView) GetLoad5Ok() (*float64, bool)`

GetLoad5Ok returns a tuple with the Load5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoad5

`func (o *SampleView) SetLoad5(v float64)`

SetLoad5 sets Load5 field to given value.

### HasLoad5

`func (o *SampleView) HasLoad5() bool`

HasLoad5 returns a boolean if a field has been set.

### GetMemFree

`func (o *SampleView) GetMemFree() int64`

GetMemFree returns the MemFree field if non-nil, zero value otherwise.

### GetMemFreeOk

`func (o *SampleView) GetMemFreeOk() (*int64, bool)`

GetMemFreeOk returns a tuple with the MemFree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemFree

`func (o *SampleView) SetMemFree(v int64)`

SetMemFree sets MemFree field to given value.

### HasMemFree

`func (o *SampleView) HasMemFree() bool`

HasMemFree returns a boolean if a field has been set.

### GetMemUsed

`func (o *SampleView) GetMemUsed() int64`

GetMemUsed returns the MemUsed field if non-nil, zero value otherwise.

### GetMemUsedOk

`func (o *SampleView) GetMemUsedOk() (*int64, bool)`

GetMemUsedOk returns a tuple with the MemUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemUsed

`func (o *SampleView) SetMemUsed(v int64)`

SetMemUsed sets MemUsed field to given value.

### HasMemUsed

`func (o *SampleView) HasMemUsed() bool`

HasMemUsed returns a boolean if a field has been set.

### GetMemory

`func (o *SampleView) GetMemory() int64`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *SampleView) GetMemoryOk() (*int64, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *SampleView) SetMemory(v int64)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *SampleView) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetSource

`func (o *SampleView) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *SampleView) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *SampleView) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *SampleView) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetUnit

`func (o *SampleView) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *SampleView) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *SampleView) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *SampleView) HasUnit() bool`

HasUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


