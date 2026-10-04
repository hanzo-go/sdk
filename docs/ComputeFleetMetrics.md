# ComputeFleetMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**At** | Pointer to **string** | At is when this reading was MEASURED, RFC 3339 in UTC — not when the board was built. A console decides staleness by comparing it to now; the board deliberately does not decide that for it. | [optional] 
**GpuUtil** | Pointer to **float64** | GPUUtil is aggregate accelerator utilization as a FRACTION of 1 — 0.42 is 42% busy, never 42. Across all of the unit&#39;s cards, not one of them. | [optional] 
**Load1** | Pointer to **float64** | Load1 is the host&#39;s 1-minute load average — runnable processes, not a percentage, so it is read against the unit&#39;s core count and can exceed 1. | [optional] 
**MemFree** | Pointer to **int64** | MemFree is host memory still available, in BYTES. It is what the source reported, not fleetSpec.Memory minus MemUsed. | [optional] 
**MemUsed** | Pointer to **int64** | MemUsed is host memory in use, in BYTES. | [optional] 

## Methods

### NewComputeFleetMetrics

`func NewComputeFleetMetrics() *ComputeFleetMetrics`

NewComputeFleetMetrics instantiates a new ComputeFleetMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputeFleetMetricsWithDefaults

`func NewComputeFleetMetricsWithDefaults() *ComputeFleetMetrics`

NewComputeFleetMetricsWithDefaults instantiates a new ComputeFleetMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAt

`func (o *ComputeFleetMetrics) GetAt() string`

GetAt returns the At field if non-nil, zero value otherwise.

### GetAtOk

`func (o *ComputeFleetMetrics) GetAtOk() (*string, bool)`

GetAtOk returns a tuple with the At field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAt

`func (o *ComputeFleetMetrics) SetAt(v string)`

SetAt sets At field to given value.

### HasAt

`func (o *ComputeFleetMetrics) HasAt() bool`

HasAt returns a boolean if a field has been set.

### GetGpuUtil

`func (o *ComputeFleetMetrics) GetGpuUtil() float64`

GetGpuUtil returns the GpuUtil field if non-nil, zero value otherwise.

### GetGpuUtilOk

`func (o *ComputeFleetMetrics) GetGpuUtilOk() (*float64, bool)`

GetGpuUtilOk returns a tuple with the GpuUtil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuUtil

`func (o *ComputeFleetMetrics) SetGpuUtil(v float64)`

SetGpuUtil sets GpuUtil field to given value.

### HasGpuUtil

`func (o *ComputeFleetMetrics) HasGpuUtil() bool`

HasGpuUtil returns a boolean if a field has been set.

### GetLoad1

`func (o *ComputeFleetMetrics) GetLoad1() float64`

GetLoad1 returns the Load1 field if non-nil, zero value otherwise.

### GetLoad1Ok

`func (o *ComputeFleetMetrics) GetLoad1Ok() (*float64, bool)`

GetLoad1Ok returns a tuple with the Load1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoad1

`func (o *ComputeFleetMetrics) SetLoad1(v float64)`

SetLoad1 sets Load1 field to given value.

### HasLoad1

`func (o *ComputeFleetMetrics) HasLoad1() bool`

HasLoad1 returns a boolean if a field has been set.

### GetMemFree

`func (o *ComputeFleetMetrics) GetMemFree() int64`

GetMemFree returns the MemFree field if non-nil, zero value otherwise.

### GetMemFreeOk

`func (o *ComputeFleetMetrics) GetMemFreeOk() (*int64, bool)`

GetMemFreeOk returns a tuple with the MemFree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemFree

`func (o *ComputeFleetMetrics) SetMemFree(v int64)`

SetMemFree sets MemFree field to given value.

### HasMemFree

`func (o *ComputeFleetMetrics) HasMemFree() bool`

HasMemFree returns a boolean if a field has been set.

### GetMemUsed

`func (o *ComputeFleetMetrics) GetMemUsed() int64`

GetMemUsed returns the MemUsed field if non-nil, zero value otherwise.

### GetMemUsedOk

`func (o *ComputeFleetMetrics) GetMemUsedOk() (*int64, bool)`

GetMemUsedOk returns a tuple with the MemUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemUsed

`func (o *ComputeFleetMetrics) SetMemUsed(v int64)`

SetMemUsed sets MemUsed field to given value.

### HasMemUsed

`func (o *ComputeFleetMetrics) HasMemUsed() bool`

HasMemUsed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


