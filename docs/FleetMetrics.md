# FleetMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**At** | Pointer to **string** | At is when this reading was MEASURED, RFC 3339 in UTC — not when the board was built. A console decides staleness by comparing it to now; the board deliberately does not decide that for it. | [optional] 
**GpuUtil** | Pointer to **float32** | GPUUtil is aggregate accelerator utilization as a FRACTION of 1 — 0.42 is 42% busy, never 42. Across all of the unit&#39;s cards, not one of them. | [optional] 
**Load1** | Pointer to **float32** | Load1 is the host&#39;s 1-minute load average — runnable processes, not a percentage, so it is read against the unit&#39;s core count and can exceed 1. | [optional] 
**MemFree** | Pointer to **int32** | MemFree is host memory still available, in BYTES. It is what the source reported, not fleetSpec.Memory minus MemUsed. | [optional] 
**MemUsed** | Pointer to **int32** | MemUsed is host memory in use, in BYTES. | [optional] 

## Methods

### NewFleetMetrics

`func NewFleetMetrics() *FleetMetrics`

NewFleetMetrics instantiates a new FleetMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFleetMetricsWithDefaults

`func NewFleetMetricsWithDefaults() *FleetMetrics`

NewFleetMetricsWithDefaults instantiates a new FleetMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAt

`func (o *FleetMetrics) GetAt() string`

GetAt returns the At field if non-nil, zero value otherwise.

### GetAtOk

`func (o *FleetMetrics) GetAtOk() (*string, bool)`

GetAtOk returns a tuple with the At field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAt

`func (o *FleetMetrics) SetAt(v string)`

SetAt sets At field to given value.

### HasAt

`func (o *FleetMetrics) HasAt() bool`

HasAt returns a boolean if a field has been set.

### GetGpuUtil

`func (o *FleetMetrics) GetGpuUtil() float32`

GetGpuUtil returns the GpuUtil field if non-nil, zero value otherwise.

### GetGpuUtilOk

`func (o *FleetMetrics) GetGpuUtilOk() (*float32, bool)`

GetGpuUtilOk returns a tuple with the GpuUtil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuUtil

`func (o *FleetMetrics) SetGpuUtil(v float32)`

SetGpuUtil sets GpuUtil field to given value.

### HasGpuUtil

`func (o *FleetMetrics) HasGpuUtil() bool`

HasGpuUtil returns a boolean if a field has been set.

### GetLoad1

`func (o *FleetMetrics) GetLoad1() float32`

GetLoad1 returns the Load1 field if non-nil, zero value otherwise.

### GetLoad1Ok

`func (o *FleetMetrics) GetLoad1Ok() (*float32, bool)`

GetLoad1Ok returns a tuple with the Load1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoad1

`func (o *FleetMetrics) SetLoad1(v float32)`

SetLoad1 sets Load1 field to given value.

### HasLoad1

`func (o *FleetMetrics) HasLoad1() bool`

HasLoad1 returns a boolean if a field has been set.

### GetMemFree

`func (o *FleetMetrics) GetMemFree() int32`

GetMemFree returns the MemFree field if non-nil, zero value otherwise.

### GetMemFreeOk

`func (o *FleetMetrics) GetMemFreeOk() (*int32, bool)`

GetMemFreeOk returns a tuple with the MemFree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemFree

`func (o *FleetMetrics) SetMemFree(v int32)`

SetMemFree sets MemFree field to given value.

### HasMemFree

`func (o *FleetMetrics) HasMemFree() bool`

HasMemFree returns a boolean if a field has been set.

### GetMemUsed

`func (o *FleetMetrics) GetMemUsed() int32`

GetMemUsed returns the MemUsed field if non-nil, zero value otherwise.

### GetMemUsedOk

`func (o *FleetMetrics) GetMemUsedOk() (*int32, bool)`

GetMemUsedOk returns a tuple with the MemUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemUsed

`func (o *FleetMetrics) SetMemUsed(v int32)`

SetMemUsed sets MemUsed field to given value.

### HasMemUsed

`func (o *FleetMetrics) HasMemUsed() bool`

HasMemUsed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


