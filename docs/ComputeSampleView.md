# ComputeSampleView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**At** | Pointer to **string** | At is when the reading was MEASURED, RFC 3339 in UTC — the x-axis a chart plots against. The series is returned oldest first, so it only increases. | [optional] 
**CostCents** | Pointer to **int64** | CostCents is what this unit resold for over the hour the reading falls in, in whole US cents. 0 means UNPRICED, not free: the operator&#39;s own machines — a linked run-target, a dialed-in BYO worker — are metered for utilization and never resold, so only a priced source ever fills it. | [optional] 
**CpuTemp** | Pointer to **float64** | CPUTemp is the CPU package temperature in °C. | [optional] 
**CpuUtil** | Pointer to **float64** | CPUUtil is the busy fraction of all cores over the sample window, 0..1. | [optional] 
**Cpus** | Pointer to **int64** | CPUs is logical cores. The static capability rides every row on purpose: a chart can size load against cores without joining a registry whose row may since have been rewritten or the unit deregistered. | [optional] 
**Decode** | Pointer to **float64** | Decode is generation throughput over the sample window, in tokens per second. | [optional] 
**DiskRead** | Pointer to **float64** | DiskRead is the block-device read rate in bytes per second. | [optional] 
**DiskTotal** | Pointer to **int64** | DiskTotal is the size of those filesystems, in BYTES. | [optional] 
**DiskUsed** | Pointer to **int64** | DiskUsed is the space used across local block filesystems, in BYTES. | [optional] 
**DiskWrite** | Pointer to **float64** | DiskWrite is the block-device write rate in bytes per second. | [optional] 
**GpuMemTotal** | Pointer to **int64** | GPUMemTotal is dedicated VRAM in BYTES; absent when the GPU shares system memory (unified). | [optional] 
**GpuMemUsed** | Pointer to **int64** | GPUMemUsed is the memory held by GPU processes, in BYTES. | [optional] 
**GpuModel** | Pointer to **string** | GPUModel names the representative accelerator (\&quot;GB10\&quot;); GPUs carries how many. | [optional] 
**GpuPower** | Pointer to **float64** | GPUPower is the total GPU power draw in watts. | [optional] 
**GpuPowerLimit** | Pointer to **float64** | GPUPowerLimit is the total enforced GPU power limit in watts; absent when the driver reports none. | [optional] 
**GpuTemp** | Pointer to **float64** | GPUTemp is the hottest GPU&#39;s temperature in °C. | [optional] 
**GpuUtil** | Pointer to **float64** | GPUUtil is aggregate accelerator utilization as a FRACTION of 1 — 0.42 is 42% busy. Anything a reporter sends outside 0..1 is clamped into it on write. | [optional] 
**Gpus** | Pointer to **int64** | GPUs is how many accelerators the reading covers. | [optional] 
**Host** | Pointer to **string** | Host is the hostname the unit reported at the time of the reading. | [optional] 
**Kind** | Pointer to **string** | Kind is what the measured unit is: laptop, cloud, gpu, cluster, machine or worker. | [optional] 
**KvCache** | Pointer to **float64** | KVCache is the KV-cache occupancy, 0..1. | [optional] 
**Load1** | Pointer to **float64** | Load1 is the 1-minute load average — runnable processes, not a percentage. | [optional] 
**Load15** | Pointer to **float64** | Load15 is the 15-minute load average, the same units as Load1. | [optional] 
**Load5** | Pointer to **float64** | Load5 is the 5-minute load average, the same units as Load1. | [optional] 
**MemFree** | Pointer to **int64** | MemFree is host memory available, in BYTES, as reported rather than derived. | [optional] 
**MemUsed** | Pointer to **int64** | MemUsed is host memory in use, in BYTES. | [optional] 
**Memory** | Pointer to **int64** | Memory is total system RAM in BYTES at the time of the reading. | [optional] 
**Model** | Pointer to **string** | Model is the model id(s) served on the unit, comma-separated, at most 128 bytes. | [optional] 
**NetRx** | Pointer to **float64** | NetRx is the receive rate of physical interfaces in bytes per second. | [optional] 
**NetTx** | Pointer to **float64** | NetTx is the transmit rate of physical interfaces in bytes per second. | [optional] 
**Prefill** | Pointer to **float64** | Prefill is prompt (prefill) throughput over the sample window, in tokens per second. | [optional] 
**Running** | Pointer to **int64** | Running is the number of requests being served at the reading. | [optional] 
**Source** | Pointer to **string** | Source is the plane that reported the reading: \&quot;agent\&quot;, \&quot;byo\&quot; or \&quot;visor\&quot; — the same vocabulary the board&#39;s rows carry, and what ?source&#x3D; narrows on. | [optional] 
**Ttft** | Pointer to **float64** | TTFT is the mean time to first token, in seconds, of requests that started in the window. | [optional] 
**Unit** | Pointer to **string** | Unit is the source&#39;s own id for the measured unit. With Source it is the key the chart groups by, and the key the board joins a unit&#39;s latest reading on. | [optional] 
**Waiting** | Pointer to **int64** | Waiting is the number of requests queued at the reading. | [optional] 

## Methods

### NewComputeSampleView

`func NewComputeSampleView() *ComputeSampleView`

NewComputeSampleView instantiates a new ComputeSampleView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputeSampleViewWithDefaults

`func NewComputeSampleViewWithDefaults() *ComputeSampleView`

NewComputeSampleViewWithDefaults instantiates a new ComputeSampleView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAt

`func (o *ComputeSampleView) GetAt() string`

GetAt returns the At field if non-nil, zero value otherwise.

### GetAtOk

`func (o *ComputeSampleView) GetAtOk() (*string, bool)`

GetAtOk returns a tuple with the At field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAt

`func (o *ComputeSampleView) SetAt(v string)`

SetAt sets At field to given value.

### HasAt

`func (o *ComputeSampleView) HasAt() bool`

HasAt returns a boolean if a field has been set.

### GetCostCents

`func (o *ComputeSampleView) GetCostCents() int64`

GetCostCents returns the CostCents field if non-nil, zero value otherwise.

### GetCostCentsOk

`func (o *ComputeSampleView) GetCostCentsOk() (*int64, bool)`

GetCostCentsOk returns a tuple with the CostCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostCents

`func (o *ComputeSampleView) SetCostCents(v int64)`

SetCostCents sets CostCents field to given value.

### HasCostCents

`func (o *ComputeSampleView) HasCostCents() bool`

HasCostCents returns a boolean if a field has been set.

### GetCpuTemp

`func (o *ComputeSampleView) GetCpuTemp() float64`

GetCpuTemp returns the CpuTemp field if non-nil, zero value otherwise.

### GetCpuTempOk

`func (o *ComputeSampleView) GetCpuTempOk() (*float64, bool)`

GetCpuTempOk returns a tuple with the CpuTemp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuTemp

`func (o *ComputeSampleView) SetCpuTemp(v float64)`

SetCpuTemp sets CpuTemp field to given value.

### HasCpuTemp

`func (o *ComputeSampleView) HasCpuTemp() bool`

HasCpuTemp returns a boolean if a field has been set.

### GetCpuUtil

`func (o *ComputeSampleView) GetCpuUtil() float64`

GetCpuUtil returns the CpuUtil field if non-nil, zero value otherwise.

### GetCpuUtilOk

`func (o *ComputeSampleView) GetCpuUtilOk() (*float64, bool)`

GetCpuUtilOk returns a tuple with the CpuUtil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUtil

`func (o *ComputeSampleView) SetCpuUtil(v float64)`

SetCpuUtil sets CpuUtil field to given value.

### HasCpuUtil

`func (o *ComputeSampleView) HasCpuUtil() bool`

HasCpuUtil returns a boolean if a field has been set.

### GetCpus

`func (o *ComputeSampleView) GetCpus() int64`

GetCpus returns the Cpus field if non-nil, zero value otherwise.

### GetCpusOk

`func (o *ComputeSampleView) GetCpusOk() (*int64, bool)`

GetCpusOk returns a tuple with the Cpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpus

`func (o *ComputeSampleView) SetCpus(v int64)`

SetCpus sets Cpus field to given value.

### HasCpus

`func (o *ComputeSampleView) HasCpus() bool`

HasCpus returns a boolean if a field has been set.

### GetDecode

`func (o *ComputeSampleView) GetDecode() float64`

GetDecode returns the Decode field if non-nil, zero value otherwise.

### GetDecodeOk

`func (o *ComputeSampleView) GetDecodeOk() (*float64, bool)`

GetDecodeOk returns a tuple with the Decode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecode

`func (o *ComputeSampleView) SetDecode(v float64)`

SetDecode sets Decode field to given value.

### HasDecode

`func (o *ComputeSampleView) HasDecode() bool`

HasDecode returns a boolean if a field has been set.

### GetDiskRead

`func (o *ComputeSampleView) GetDiskRead() float64`

GetDiskRead returns the DiskRead field if non-nil, zero value otherwise.

### GetDiskReadOk

`func (o *ComputeSampleView) GetDiskReadOk() (*float64, bool)`

GetDiskReadOk returns a tuple with the DiskRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskRead

`func (o *ComputeSampleView) SetDiskRead(v float64)`

SetDiskRead sets DiskRead field to given value.

### HasDiskRead

`func (o *ComputeSampleView) HasDiskRead() bool`

HasDiskRead returns a boolean if a field has been set.

### GetDiskTotal

`func (o *ComputeSampleView) GetDiskTotal() int64`

GetDiskTotal returns the DiskTotal field if non-nil, zero value otherwise.

### GetDiskTotalOk

`func (o *ComputeSampleView) GetDiskTotalOk() (*int64, bool)`

GetDiskTotalOk returns a tuple with the DiskTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskTotal

`func (o *ComputeSampleView) SetDiskTotal(v int64)`

SetDiskTotal sets DiskTotal field to given value.

### HasDiskTotal

`func (o *ComputeSampleView) HasDiskTotal() bool`

HasDiskTotal returns a boolean if a field has been set.

### GetDiskUsed

`func (o *ComputeSampleView) GetDiskUsed() int64`

GetDiskUsed returns the DiskUsed field if non-nil, zero value otherwise.

### GetDiskUsedOk

`func (o *ComputeSampleView) GetDiskUsedOk() (*int64, bool)`

GetDiskUsedOk returns a tuple with the DiskUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskUsed

`func (o *ComputeSampleView) SetDiskUsed(v int64)`

SetDiskUsed sets DiskUsed field to given value.

### HasDiskUsed

`func (o *ComputeSampleView) HasDiskUsed() bool`

HasDiskUsed returns a boolean if a field has been set.

### GetDiskWrite

`func (o *ComputeSampleView) GetDiskWrite() float64`

GetDiskWrite returns the DiskWrite field if non-nil, zero value otherwise.

### GetDiskWriteOk

`func (o *ComputeSampleView) GetDiskWriteOk() (*float64, bool)`

GetDiskWriteOk returns a tuple with the DiskWrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskWrite

`func (o *ComputeSampleView) SetDiskWrite(v float64)`

SetDiskWrite sets DiskWrite field to given value.

### HasDiskWrite

`func (o *ComputeSampleView) HasDiskWrite() bool`

HasDiskWrite returns a boolean if a field has been set.

### GetGpuMemTotal

`func (o *ComputeSampleView) GetGpuMemTotal() int64`

GetGpuMemTotal returns the GpuMemTotal field if non-nil, zero value otherwise.

### GetGpuMemTotalOk

`func (o *ComputeSampleView) GetGpuMemTotalOk() (*int64, bool)`

GetGpuMemTotalOk returns a tuple with the GpuMemTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuMemTotal

`func (o *ComputeSampleView) SetGpuMemTotal(v int64)`

SetGpuMemTotal sets GpuMemTotal field to given value.

### HasGpuMemTotal

`func (o *ComputeSampleView) HasGpuMemTotal() bool`

HasGpuMemTotal returns a boolean if a field has been set.

### GetGpuMemUsed

`func (o *ComputeSampleView) GetGpuMemUsed() int64`

GetGpuMemUsed returns the GpuMemUsed field if non-nil, zero value otherwise.

### GetGpuMemUsedOk

`func (o *ComputeSampleView) GetGpuMemUsedOk() (*int64, bool)`

GetGpuMemUsedOk returns a tuple with the GpuMemUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuMemUsed

`func (o *ComputeSampleView) SetGpuMemUsed(v int64)`

SetGpuMemUsed sets GpuMemUsed field to given value.

### HasGpuMemUsed

`func (o *ComputeSampleView) HasGpuMemUsed() bool`

HasGpuMemUsed returns a boolean if a field has been set.

### GetGpuModel

`func (o *ComputeSampleView) GetGpuModel() string`

GetGpuModel returns the GpuModel field if non-nil, zero value otherwise.

### GetGpuModelOk

`func (o *ComputeSampleView) GetGpuModelOk() (*string, bool)`

GetGpuModelOk returns a tuple with the GpuModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuModel

`func (o *ComputeSampleView) SetGpuModel(v string)`

SetGpuModel sets GpuModel field to given value.

### HasGpuModel

`func (o *ComputeSampleView) HasGpuModel() bool`

HasGpuModel returns a boolean if a field has been set.

### GetGpuPower

`func (o *ComputeSampleView) GetGpuPower() float64`

GetGpuPower returns the GpuPower field if non-nil, zero value otherwise.

### GetGpuPowerOk

`func (o *ComputeSampleView) GetGpuPowerOk() (*float64, bool)`

GetGpuPowerOk returns a tuple with the GpuPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuPower

`func (o *ComputeSampleView) SetGpuPower(v float64)`

SetGpuPower sets GpuPower field to given value.

### HasGpuPower

`func (o *ComputeSampleView) HasGpuPower() bool`

HasGpuPower returns a boolean if a field has been set.

### GetGpuPowerLimit

`func (o *ComputeSampleView) GetGpuPowerLimit() float64`

GetGpuPowerLimit returns the GpuPowerLimit field if non-nil, zero value otherwise.

### GetGpuPowerLimitOk

`func (o *ComputeSampleView) GetGpuPowerLimitOk() (*float64, bool)`

GetGpuPowerLimitOk returns a tuple with the GpuPowerLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuPowerLimit

`func (o *ComputeSampleView) SetGpuPowerLimit(v float64)`

SetGpuPowerLimit sets GpuPowerLimit field to given value.

### HasGpuPowerLimit

`func (o *ComputeSampleView) HasGpuPowerLimit() bool`

HasGpuPowerLimit returns a boolean if a field has been set.

### GetGpuTemp

`func (o *ComputeSampleView) GetGpuTemp() float64`

GetGpuTemp returns the GpuTemp field if non-nil, zero value otherwise.

### GetGpuTempOk

`func (o *ComputeSampleView) GetGpuTempOk() (*float64, bool)`

GetGpuTempOk returns a tuple with the GpuTemp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuTemp

`func (o *ComputeSampleView) SetGpuTemp(v float64)`

SetGpuTemp sets GpuTemp field to given value.

### HasGpuTemp

`func (o *ComputeSampleView) HasGpuTemp() bool`

HasGpuTemp returns a boolean if a field has been set.

### GetGpuUtil

`func (o *ComputeSampleView) GetGpuUtil() float64`

GetGpuUtil returns the GpuUtil field if non-nil, zero value otherwise.

### GetGpuUtilOk

`func (o *ComputeSampleView) GetGpuUtilOk() (*float64, bool)`

GetGpuUtilOk returns a tuple with the GpuUtil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuUtil

`func (o *ComputeSampleView) SetGpuUtil(v float64)`

SetGpuUtil sets GpuUtil field to given value.

### HasGpuUtil

`func (o *ComputeSampleView) HasGpuUtil() bool`

HasGpuUtil returns a boolean if a field has been set.

### GetGpus

`func (o *ComputeSampleView) GetGpus() int64`

GetGpus returns the Gpus field if non-nil, zero value otherwise.

### GetGpusOk

`func (o *ComputeSampleView) GetGpusOk() (*int64, bool)`

GetGpusOk returns a tuple with the Gpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpus

`func (o *ComputeSampleView) SetGpus(v int64)`

SetGpus sets Gpus field to given value.

### HasGpus

`func (o *ComputeSampleView) HasGpus() bool`

HasGpus returns a boolean if a field has been set.

### GetHost

`func (o *ComputeSampleView) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *ComputeSampleView) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *ComputeSampleView) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *ComputeSampleView) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetKind

`func (o *ComputeSampleView) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ComputeSampleView) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ComputeSampleView) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ComputeSampleView) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetKvCache

`func (o *ComputeSampleView) GetKvCache() float64`

GetKvCache returns the KvCache field if non-nil, zero value otherwise.

### GetKvCacheOk

`func (o *ComputeSampleView) GetKvCacheOk() (*float64, bool)`

GetKvCacheOk returns a tuple with the KvCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKvCache

`func (o *ComputeSampleView) SetKvCache(v float64)`

SetKvCache sets KvCache field to given value.

### HasKvCache

`func (o *ComputeSampleView) HasKvCache() bool`

HasKvCache returns a boolean if a field has been set.

### GetLoad1

`func (o *ComputeSampleView) GetLoad1() float64`

GetLoad1 returns the Load1 field if non-nil, zero value otherwise.

### GetLoad1Ok

`func (o *ComputeSampleView) GetLoad1Ok() (*float64, bool)`

GetLoad1Ok returns a tuple with the Load1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoad1

`func (o *ComputeSampleView) SetLoad1(v float64)`

SetLoad1 sets Load1 field to given value.

### HasLoad1

`func (o *ComputeSampleView) HasLoad1() bool`

HasLoad1 returns a boolean if a field has been set.

### GetLoad15

`func (o *ComputeSampleView) GetLoad15() float64`

GetLoad15 returns the Load15 field if non-nil, zero value otherwise.

### GetLoad15Ok

`func (o *ComputeSampleView) GetLoad15Ok() (*float64, bool)`

GetLoad15Ok returns a tuple with the Load15 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoad15

`func (o *ComputeSampleView) SetLoad15(v float64)`

SetLoad15 sets Load15 field to given value.

### HasLoad15

`func (o *ComputeSampleView) HasLoad15() bool`

HasLoad15 returns a boolean if a field has been set.

### GetLoad5

`func (o *ComputeSampleView) GetLoad5() float64`

GetLoad5 returns the Load5 field if non-nil, zero value otherwise.

### GetLoad5Ok

`func (o *ComputeSampleView) GetLoad5Ok() (*float64, bool)`

GetLoad5Ok returns a tuple with the Load5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoad5

`func (o *ComputeSampleView) SetLoad5(v float64)`

SetLoad5 sets Load5 field to given value.

### HasLoad5

`func (o *ComputeSampleView) HasLoad5() bool`

HasLoad5 returns a boolean if a field has been set.

### GetMemFree

`func (o *ComputeSampleView) GetMemFree() int64`

GetMemFree returns the MemFree field if non-nil, zero value otherwise.

### GetMemFreeOk

`func (o *ComputeSampleView) GetMemFreeOk() (*int64, bool)`

GetMemFreeOk returns a tuple with the MemFree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemFree

`func (o *ComputeSampleView) SetMemFree(v int64)`

SetMemFree sets MemFree field to given value.

### HasMemFree

`func (o *ComputeSampleView) HasMemFree() bool`

HasMemFree returns a boolean if a field has been set.

### GetMemUsed

`func (o *ComputeSampleView) GetMemUsed() int64`

GetMemUsed returns the MemUsed field if non-nil, zero value otherwise.

### GetMemUsedOk

`func (o *ComputeSampleView) GetMemUsedOk() (*int64, bool)`

GetMemUsedOk returns a tuple with the MemUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemUsed

`func (o *ComputeSampleView) SetMemUsed(v int64)`

SetMemUsed sets MemUsed field to given value.

### HasMemUsed

`func (o *ComputeSampleView) HasMemUsed() bool`

HasMemUsed returns a boolean if a field has been set.

### GetMemory

`func (o *ComputeSampleView) GetMemory() int64`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *ComputeSampleView) GetMemoryOk() (*int64, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *ComputeSampleView) SetMemory(v int64)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *ComputeSampleView) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetModel

`func (o *ComputeSampleView) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ComputeSampleView) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ComputeSampleView) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *ComputeSampleView) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetNetRx

`func (o *ComputeSampleView) GetNetRx() float64`

GetNetRx returns the NetRx field if non-nil, zero value otherwise.

### GetNetRxOk

`func (o *ComputeSampleView) GetNetRxOk() (*float64, bool)`

GetNetRxOk returns a tuple with the NetRx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetRx

`func (o *ComputeSampleView) SetNetRx(v float64)`

SetNetRx sets NetRx field to given value.

### HasNetRx

`func (o *ComputeSampleView) HasNetRx() bool`

HasNetRx returns a boolean if a field has been set.

### GetNetTx

`func (o *ComputeSampleView) GetNetTx() float64`

GetNetTx returns the NetTx field if non-nil, zero value otherwise.

### GetNetTxOk

`func (o *ComputeSampleView) GetNetTxOk() (*float64, bool)`

GetNetTxOk returns a tuple with the NetTx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetTx

`func (o *ComputeSampleView) SetNetTx(v float64)`

SetNetTx sets NetTx field to given value.

### HasNetTx

`func (o *ComputeSampleView) HasNetTx() bool`

HasNetTx returns a boolean if a field has been set.

### GetPrefill

`func (o *ComputeSampleView) GetPrefill() float64`

GetPrefill returns the Prefill field if non-nil, zero value otherwise.

### GetPrefillOk

`func (o *ComputeSampleView) GetPrefillOk() (*float64, bool)`

GetPrefillOk returns a tuple with the Prefill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefill

`func (o *ComputeSampleView) SetPrefill(v float64)`

SetPrefill sets Prefill field to given value.

### HasPrefill

`func (o *ComputeSampleView) HasPrefill() bool`

HasPrefill returns a boolean if a field has been set.

### GetRunning

`func (o *ComputeSampleView) GetRunning() int64`

GetRunning returns the Running field if non-nil, zero value otherwise.

### GetRunningOk

`func (o *ComputeSampleView) GetRunningOk() (*int64, bool)`

GetRunningOk returns a tuple with the Running field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunning

`func (o *ComputeSampleView) SetRunning(v int64)`

SetRunning sets Running field to given value.

### HasRunning

`func (o *ComputeSampleView) HasRunning() bool`

HasRunning returns a boolean if a field has been set.

### GetSource

`func (o *ComputeSampleView) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ComputeSampleView) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ComputeSampleView) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *ComputeSampleView) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetTtft

`func (o *ComputeSampleView) GetTtft() float64`

GetTtft returns the Ttft field if non-nil, zero value otherwise.

### GetTtftOk

`func (o *ComputeSampleView) GetTtftOk() (*float64, bool)`

GetTtftOk returns a tuple with the Ttft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtft

`func (o *ComputeSampleView) SetTtft(v float64)`

SetTtft sets Ttft field to given value.

### HasTtft

`func (o *ComputeSampleView) HasTtft() bool`

HasTtft returns a boolean if a field has been set.

### GetUnit

`func (o *ComputeSampleView) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *ComputeSampleView) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *ComputeSampleView) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *ComputeSampleView) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetWaiting

`func (o *ComputeSampleView) GetWaiting() int64`

GetWaiting returns the Waiting field if non-nil, zero value otherwise.

### GetWaitingOk

`func (o *ComputeSampleView) GetWaitingOk() (*int64, bool)`

GetWaitingOk returns a tuple with the Waiting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaiting

`func (o *ComputeSampleView) SetWaiting(v int64)`

SetWaiting sets Waiting field to given value.

### HasWaiting

`func (o *ComputeSampleView) HasWaiting() bool`

HasWaiting returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


