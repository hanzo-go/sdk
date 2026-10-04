# AgentMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**At** | Pointer to **int64** | unix seconds, server-stamped | [optional] 
**CpuTemp** | Pointer to **float64** | CPUTemp is the CPU package temperature in °C. | [optional] 
**CpuUtil** | Pointer to **float64** | CPUUtil is the busy fraction of all cores over the sample window, 0..1. | [optional] 
**Decode** | Pointer to **float64** | Decode is generation throughput over the sample window, in tokens per second. | [optional] 
**DiskRead** | Pointer to **float64** | DiskRead is the block-device read rate in bytes per second. | [optional] 
**DiskTotal** | Pointer to **int64** | DiskTotal is the size of those filesystems, in bytes. | [optional] 
**DiskUsed** | Pointer to **int64** | DiskUsed is the space used across local block filesystems, in bytes. | [optional] 
**DiskWrite** | Pointer to **float64** | DiskWrite is the block-device write rate in bytes per second. | [optional] 
**GpuMemTotal** | Pointer to **int64** | GPUMemTotal is dedicated VRAM in bytes; absent when the GPU shares system memory (unified). | [optional] 
**GpuMemUsed** | Pointer to **int64** | GPUMemUsed is the memory held by GPU processes, in bytes. | [optional] 
**GpuPower** | Pointer to **float64** | GPUPower is the total GPU power draw in watts. | [optional] 
**GpuPowerLimit** | Pointer to **float64** | GPUPowerLimit is the total enforced GPU power limit in watts; absent when the driver reports none. | [optional] 
**GpuTemp** | Pointer to **float64** | GPUTemp is the hottest GPU&#39;s temperature in °C. | [optional] 
**GpuUtil** | Pointer to **float64** | 0..1 aggregate utilization | [optional] 
**KvCache** | Pointer to **float64** | KVCache is the KV-cache occupancy, 0..1. | [optional] 
**Load1** | Pointer to **float64** | Load1 is the machine&#39;s own one-minute load average — a count of runnable and uninterruptible tasks, NOT a percentage and NOT already divided by core count, so it is read against Spec.CPUs: 8.0 is idle on 16 cores and swamped on 4. Coerced finite and non-negative on write, so 0 means either genuinely idle or nothing reported. | [optional] 
**Load15** | Pointer to **float64** | Load15 is the same figure over fifteen. The three together are what separate a machine that is busy right now from one that has been busy all along — which is the question a dispatcher is really asking. | [optional] 
**Load5** | Pointer to **float64** | Load5 is the same figure averaged over five minutes. | [optional] 
**MemFree** | Pointer to **int64** | bytes | [optional] 
**MemUsed** | Pointer to **int64** | bytes | [optional] 
**Model** | Pointer to **string** | Model is the model id(s) served on this machine, comma-separated, at most 128 bytes. | [optional] 
**NetRx** | Pointer to **float64** | NetRx is the receive rate of physical interfaces in bytes per second. | [optional] 
**NetTx** | Pointer to **float64** | NetTx is the transmit rate of physical interfaces in bytes per second. | [optional] 
**Prefill** | Pointer to **float64** | Prefill is prompt (prefill) throughput over the sample window, in tokens per second. | [optional] 
**Running** | Pointer to **int64** | Running is the number of requests being served now. | [optional] 
**Ttft** | Pointer to **float64** | TTFT is the mean time to first token, in seconds, of requests that started in the window. | [optional] 
**Waiting** | Pointer to **int64** | Waiting is the number of requests queued. | [optional] 

## Methods

### NewAgentMetrics

`func NewAgentMetrics() *AgentMetrics`

NewAgentMetrics instantiates a new AgentMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentMetricsWithDefaults

`func NewAgentMetricsWithDefaults() *AgentMetrics`

NewAgentMetricsWithDefaults instantiates a new AgentMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAt

`func (o *AgentMetrics) GetAt() int64`

GetAt returns the At field if non-nil, zero value otherwise.

### GetAtOk

`func (o *AgentMetrics) GetAtOk() (*int64, bool)`

GetAtOk returns a tuple with the At field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAt

`func (o *AgentMetrics) SetAt(v int64)`

SetAt sets At field to given value.

### HasAt

`func (o *AgentMetrics) HasAt() bool`

HasAt returns a boolean if a field has been set.

### GetCpuTemp

`func (o *AgentMetrics) GetCpuTemp() float64`

GetCpuTemp returns the CpuTemp field if non-nil, zero value otherwise.

### GetCpuTempOk

`func (o *AgentMetrics) GetCpuTempOk() (*float64, bool)`

GetCpuTempOk returns a tuple with the CpuTemp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuTemp

`func (o *AgentMetrics) SetCpuTemp(v float64)`

SetCpuTemp sets CpuTemp field to given value.

### HasCpuTemp

`func (o *AgentMetrics) HasCpuTemp() bool`

HasCpuTemp returns a boolean if a field has been set.

### GetCpuUtil

`func (o *AgentMetrics) GetCpuUtil() float64`

GetCpuUtil returns the CpuUtil field if non-nil, zero value otherwise.

### GetCpuUtilOk

`func (o *AgentMetrics) GetCpuUtilOk() (*float64, bool)`

GetCpuUtilOk returns a tuple with the CpuUtil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUtil

`func (o *AgentMetrics) SetCpuUtil(v float64)`

SetCpuUtil sets CpuUtil field to given value.

### HasCpuUtil

`func (o *AgentMetrics) HasCpuUtil() bool`

HasCpuUtil returns a boolean if a field has been set.

### GetDecode

`func (o *AgentMetrics) GetDecode() float64`

GetDecode returns the Decode field if non-nil, zero value otherwise.

### GetDecodeOk

`func (o *AgentMetrics) GetDecodeOk() (*float64, bool)`

GetDecodeOk returns a tuple with the Decode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecode

`func (o *AgentMetrics) SetDecode(v float64)`

SetDecode sets Decode field to given value.

### HasDecode

`func (o *AgentMetrics) HasDecode() bool`

HasDecode returns a boolean if a field has been set.

### GetDiskRead

`func (o *AgentMetrics) GetDiskRead() float64`

GetDiskRead returns the DiskRead field if non-nil, zero value otherwise.

### GetDiskReadOk

`func (o *AgentMetrics) GetDiskReadOk() (*float64, bool)`

GetDiskReadOk returns a tuple with the DiskRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskRead

`func (o *AgentMetrics) SetDiskRead(v float64)`

SetDiskRead sets DiskRead field to given value.

### HasDiskRead

`func (o *AgentMetrics) HasDiskRead() bool`

HasDiskRead returns a boolean if a field has been set.

### GetDiskTotal

`func (o *AgentMetrics) GetDiskTotal() int64`

GetDiskTotal returns the DiskTotal field if non-nil, zero value otherwise.

### GetDiskTotalOk

`func (o *AgentMetrics) GetDiskTotalOk() (*int64, bool)`

GetDiskTotalOk returns a tuple with the DiskTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskTotal

`func (o *AgentMetrics) SetDiskTotal(v int64)`

SetDiskTotal sets DiskTotal field to given value.

### HasDiskTotal

`func (o *AgentMetrics) HasDiskTotal() bool`

HasDiskTotal returns a boolean if a field has been set.

### GetDiskUsed

`func (o *AgentMetrics) GetDiskUsed() int64`

GetDiskUsed returns the DiskUsed field if non-nil, zero value otherwise.

### GetDiskUsedOk

`func (o *AgentMetrics) GetDiskUsedOk() (*int64, bool)`

GetDiskUsedOk returns a tuple with the DiskUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskUsed

`func (o *AgentMetrics) SetDiskUsed(v int64)`

SetDiskUsed sets DiskUsed field to given value.

### HasDiskUsed

`func (o *AgentMetrics) HasDiskUsed() bool`

HasDiskUsed returns a boolean if a field has been set.

### GetDiskWrite

`func (o *AgentMetrics) GetDiskWrite() float64`

GetDiskWrite returns the DiskWrite field if non-nil, zero value otherwise.

### GetDiskWriteOk

`func (o *AgentMetrics) GetDiskWriteOk() (*float64, bool)`

GetDiskWriteOk returns a tuple with the DiskWrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskWrite

`func (o *AgentMetrics) SetDiskWrite(v float64)`

SetDiskWrite sets DiskWrite field to given value.

### HasDiskWrite

`func (o *AgentMetrics) HasDiskWrite() bool`

HasDiskWrite returns a boolean if a field has been set.

### GetGpuMemTotal

`func (o *AgentMetrics) GetGpuMemTotal() int64`

GetGpuMemTotal returns the GpuMemTotal field if non-nil, zero value otherwise.

### GetGpuMemTotalOk

`func (o *AgentMetrics) GetGpuMemTotalOk() (*int64, bool)`

GetGpuMemTotalOk returns a tuple with the GpuMemTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuMemTotal

`func (o *AgentMetrics) SetGpuMemTotal(v int64)`

SetGpuMemTotal sets GpuMemTotal field to given value.

### HasGpuMemTotal

`func (o *AgentMetrics) HasGpuMemTotal() bool`

HasGpuMemTotal returns a boolean if a field has been set.

### GetGpuMemUsed

`func (o *AgentMetrics) GetGpuMemUsed() int64`

GetGpuMemUsed returns the GpuMemUsed field if non-nil, zero value otherwise.

### GetGpuMemUsedOk

`func (o *AgentMetrics) GetGpuMemUsedOk() (*int64, bool)`

GetGpuMemUsedOk returns a tuple with the GpuMemUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuMemUsed

`func (o *AgentMetrics) SetGpuMemUsed(v int64)`

SetGpuMemUsed sets GpuMemUsed field to given value.

### HasGpuMemUsed

`func (o *AgentMetrics) HasGpuMemUsed() bool`

HasGpuMemUsed returns a boolean if a field has been set.

### GetGpuPower

`func (o *AgentMetrics) GetGpuPower() float64`

GetGpuPower returns the GpuPower field if non-nil, zero value otherwise.

### GetGpuPowerOk

`func (o *AgentMetrics) GetGpuPowerOk() (*float64, bool)`

GetGpuPowerOk returns a tuple with the GpuPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuPower

`func (o *AgentMetrics) SetGpuPower(v float64)`

SetGpuPower sets GpuPower field to given value.

### HasGpuPower

`func (o *AgentMetrics) HasGpuPower() bool`

HasGpuPower returns a boolean if a field has been set.

### GetGpuPowerLimit

`func (o *AgentMetrics) GetGpuPowerLimit() float64`

GetGpuPowerLimit returns the GpuPowerLimit field if non-nil, zero value otherwise.

### GetGpuPowerLimitOk

`func (o *AgentMetrics) GetGpuPowerLimitOk() (*float64, bool)`

GetGpuPowerLimitOk returns a tuple with the GpuPowerLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuPowerLimit

`func (o *AgentMetrics) SetGpuPowerLimit(v float64)`

SetGpuPowerLimit sets GpuPowerLimit field to given value.

### HasGpuPowerLimit

`func (o *AgentMetrics) HasGpuPowerLimit() bool`

HasGpuPowerLimit returns a boolean if a field has been set.

### GetGpuTemp

`func (o *AgentMetrics) GetGpuTemp() float64`

GetGpuTemp returns the GpuTemp field if non-nil, zero value otherwise.

### GetGpuTempOk

`func (o *AgentMetrics) GetGpuTempOk() (*float64, bool)`

GetGpuTempOk returns a tuple with the GpuTemp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuTemp

`func (o *AgentMetrics) SetGpuTemp(v float64)`

SetGpuTemp sets GpuTemp field to given value.

### HasGpuTemp

`func (o *AgentMetrics) HasGpuTemp() bool`

HasGpuTemp returns a boolean if a field has been set.

### GetGpuUtil

`func (o *AgentMetrics) GetGpuUtil() float64`

GetGpuUtil returns the GpuUtil field if non-nil, zero value otherwise.

### GetGpuUtilOk

`func (o *AgentMetrics) GetGpuUtilOk() (*float64, bool)`

GetGpuUtilOk returns a tuple with the GpuUtil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuUtil

`func (o *AgentMetrics) SetGpuUtil(v float64)`

SetGpuUtil sets GpuUtil field to given value.

### HasGpuUtil

`func (o *AgentMetrics) HasGpuUtil() bool`

HasGpuUtil returns a boolean if a field has been set.

### GetKvCache

`func (o *AgentMetrics) GetKvCache() float64`

GetKvCache returns the KvCache field if non-nil, zero value otherwise.

### GetKvCacheOk

`func (o *AgentMetrics) GetKvCacheOk() (*float64, bool)`

GetKvCacheOk returns a tuple with the KvCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKvCache

`func (o *AgentMetrics) SetKvCache(v float64)`

SetKvCache sets KvCache field to given value.

### HasKvCache

`func (o *AgentMetrics) HasKvCache() bool`

HasKvCache returns a boolean if a field has been set.

### GetLoad1

`func (o *AgentMetrics) GetLoad1() float64`

GetLoad1 returns the Load1 field if non-nil, zero value otherwise.

### GetLoad1Ok

`func (o *AgentMetrics) GetLoad1Ok() (*float64, bool)`

GetLoad1Ok returns a tuple with the Load1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoad1

`func (o *AgentMetrics) SetLoad1(v float64)`

SetLoad1 sets Load1 field to given value.

### HasLoad1

`func (o *AgentMetrics) HasLoad1() bool`

HasLoad1 returns a boolean if a field has been set.

### GetLoad15

`func (o *AgentMetrics) GetLoad15() float64`

GetLoad15 returns the Load15 field if non-nil, zero value otherwise.

### GetLoad15Ok

`func (o *AgentMetrics) GetLoad15Ok() (*float64, bool)`

GetLoad15Ok returns a tuple with the Load15 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoad15

`func (o *AgentMetrics) SetLoad15(v float64)`

SetLoad15 sets Load15 field to given value.

### HasLoad15

`func (o *AgentMetrics) HasLoad15() bool`

HasLoad15 returns a boolean if a field has been set.

### GetLoad5

`func (o *AgentMetrics) GetLoad5() float64`

GetLoad5 returns the Load5 field if non-nil, zero value otherwise.

### GetLoad5Ok

`func (o *AgentMetrics) GetLoad5Ok() (*float64, bool)`

GetLoad5Ok returns a tuple with the Load5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoad5

`func (o *AgentMetrics) SetLoad5(v float64)`

SetLoad5 sets Load5 field to given value.

### HasLoad5

`func (o *AgentMetrics) HasLoad5() bool`

HasLoad5 returns a boolean if a field has been set.

### GetMemFree

`func (o *AgentMetrics) GetMemFree() int64`

GetMemFree returns the MemFree field if non-nil, zero value otherwise.

### GetMemFreeOk

`func (o *AgentMetrics) GetMemFreeOk() (*int64, bool)`

GetMemFreeOk returns a tuple with the MemFree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemFree

`func (o *AgentMetrics) SetMemFree(v int64)`

SetMemFree sets MemFree field to given value.

### HasMemFree

`func (o *AgentMetrics) HasMemFree() bool`

HasMemFree returns a boolean if a field has been set.

### GetMemUsed

`func (o *AgentMetrics) GetMemUsed() int64`

GetMemUsed returns the MemUsed field if non-nil, zero value otherwise.

### GetMemUsedOk

`func (o *AgentMetrics) GetMemUsedOk() (*int64, bool)`

GetMemUsedOk returns a tuple with the MemUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemUsed

`func (o *AgentMetrics) SetMemUsed(v int64)`

SetMemUsed sets MemUsed field to given value.

### HasMemUsed

`func (o *AgentMetrics) HasMemUsed() bool`

HasMemUsed returns a boolean if a field has been set.

### GetModel

`func (o *AgentMetrics) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *AgentMetrics) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *AgentMetrics) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *AgentMetrics) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetNetRx

`func (o *AgentMetrics) GetNetRx() float64`

GetNetRx returns the NetRx field if non-nil, zero value otherwise.

### GetNetRxOk

`func (o *AgentMetrics) GetNetRxOk() (*float64, bool)`

GetNetRxOk returns a tuple with the NetRx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetRx

`func (o *AgentMetrics) SetNetRx(v float64)`

SetNetRx sets NetRx field to given value.

### HasNetRx

`func (o *AgentMetrics) HasNetRx() bool`

HasNetRx returns a boolean if a field has been set.

### GetNetTx

`func (o *AgentMetrics) GetNetTx() float64`

GetNetTx returns the NetTx field if non-nil, zero value otherwise.

### GetNetTxOk

`func (o *AgentMetrics) GetNetTxOk() (*float64, bool)`

GetNetTxOk returns a tuple with the NetTx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetTx

`func (o *AgentMetrics) SetNetTx(v float64)`

SetNetTx sets NetTx field to given value.

### HasNetTx

`func (o *AgentMetrics) HasNetTx() bool`

HasNetTx returns a boolean if a field has been set.

### GetPrefill

`func (o *AgentMetrics) GetPrefill() float64`

GetPrefill returns the Prefill field if non-nil, zero value otherwise.

### GetPrefillOk

`func (o *AgentMetrics) GetPrefillOk() (*float64, bool)`

GetPrefillOk returns a tuple with the Prefill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefill

`func (o *AgentMetrics) SetPrefill(v float64)`

SetPrefill sets Prefill field to given value.

### HasPrefill

`func (o *AgentMetrics) HasPrefill() bool`

HasPrefill returns a boolean if a field has been set.

### GetRunning

`func (o *AgentMetrics) GetRunning() int64`

GetRunning returns the Running field if non-nil, zero value otherwise.

### GetRunningOk

`func (o *AgentMetrics) GetRunningOk() (*int64, bool)`

GetRunningOk returns a tuple with the Running field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunning

`func (o *AgentMetrics) SetRunning(v int64)`

SetRunning sets Running field to given value.

### HasRunning

`func (o *AgentMetrics) HasRunning() bool`

HasRunning returns a boolean if a field has been set.

### GetTtft

`func (o *AgentMetrics) GetTtft() float64`

GetTtft returns the Ttft field if non-nil, zero value otherwise.

### GetTtftOk

`func (o *AgentMetrics) GetTtftOk() (*float64, bool)`

GetTtftOk returns a tuple with the Ttft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtft

`func (o *AgentMetrics) SetTtft(v float64)`

SetTtft sets Ttft field to given value.

### HasTtft

`func (o *AgentMetrics) HasTtft() bool`

HasTtft returns a boolean if a field has been set.

### GetWaiting

`func (o *AgentMetrics) GetWaiting() int64`

GetWaiting returns the Waiting field if non-nil, zero value otherwise.

### GetWaitingOk

`func (o *AgentMetrics) GetWaitingOk() (*int64, bool)`

GetWaitingOk returns a tuple with the Waiting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaiting

`func (o *AgentMetrics) SetWaiting(v int64)`

SetWaiting sets Waiting field to given value.

### HasWaiting

`func (o *AgentMetrics) HasWaiting() bool`

HasWaiting returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


