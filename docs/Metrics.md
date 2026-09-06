# Metrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**At** | Pointer to **int64** | unix seconds, server-stamped | [optional] 
**GpuUtil** | Pointer to **float64** | 0..1 aggregate utilization | [optional] 
**Load1** | Pointer to **float64** | Load1 is the machine&#39;s own one-minute load average — a count of runnable and uninterruptible tasks, NOT a percentage and NOT already divided by core count, so it is read against Spec.CPUs: 8.0 is idle on 16 cores and swamped on 4. Coerced finite and non-negative on write, so 0 means either genuinely idle or nothing reported. | [optional] 
**Load15** | Pointer to **float64** | Load15 is the same figure over fifteen. The three together are what separate a machine that is busy right now from one that has been busy all along — which is the question a dispatcher is really asking. | [optional] 
**Load5** | Pointer to **float64** | Load5 is the same figure averaged over five minutes. | [optional] 
**MemFree** | Pointer to **int64** | bytes | [optional] 
**MemUsed** | Pointer to **int64** | bytes | [optional] 

## Methods

### NewMetrics

`func NewMetrics() *Metrics`

NewMetrics instantiates a new Metrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricsWithDefaults

`func NewMetricsWithDefaults() *Metrics`

NewMetricsWithDefaults instantiates a new Metrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAt

`func (o *Metrics) GetAt() int64`

GetAt returns the At field if non-nil, zero value otherwise.

### GetAtOk

`func (o *Metrics) GetAtOk() (*int64, bool)`

GetAtOk returns a tuple with the At field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAt

`func (o *Metrics) SetAt(v int64)`

SetAt sets At field to given value.

### HasAt

`func (o *Metrics) HasAt() bool`

HasAt returns a boolean if a field has been set.

### GetGpuUtil

`func (o *Metrics) GetGpuUtil() float64`

GetGpuUtil returns the GpuUtil field if non-nil, zero value otherwise.

### GetGpuUtilOk

`func (o *Metrics) GetGpuUtilOk() (*float64, bool)`

GetGpuUtilOk returns a tuple with the GpuUtil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuUtil

`func (o *Metrics) SetGpuUtil(v float64)`

SetGpuUtil sets GpuUtil field to given value.

### HasGpuUtil

`func (o *Metrics) HasGpuUtil() bool`

HasGpuUtil returns a boolean if a field has been set.

### GetLoad1

`func (o *Metrics) GetLoad1() float64`

GetLoad1 returns the Load1 field if non-nil, zero value otherwise.

### GetLoad1Ok

`func (o *Metrics) GetLoad1Ok() (*float64, bool)`

GetLoad1Ok returns a tuple with the Load1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoad1

`func (o *Metrics) SetLoad1(v float64)`

SetLoad1 sets Load1 field to given value.

### HasLoad1

`func (o *Metrics) HasLoad1() bool`

HasLoad1 returns a boolean if a field has been set.

### GetLoad15

`func (o *Metrics) GetLoad15() float64`

GetLoad15 returns the Load15 field if non-nil, zero value otherwise.

### GetLoad15Ok

`func (o *Metrics) GetLoad15Ok() (*float64, bool)`

GetLoad15Ok returns a tuple with the Load15 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoad15

`func (o *Metrics) SetLoad15(v float64)`

SetLoad15 sets Load15 field to given value.

### HasLoad15

`func (o *Metrics) HasLoad15() bool`

HasLoad15 returns a boolean if a field has been set.

### GetLoad5

`func (o *Metrics) GetLoad5() float64`

GetLoad5 returns the Load5 field if non-nil, zero value otherwise.

### GetLoad5Ok

`func (o *Metrics) GetLoad5Ok() (*float64, bool)`

GetLoad5Ok returns a tuple with the Load5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoad5

`func (o *Metrics) SetLoad5(v float64)`

SetLoad5 sets Load5 field to given value.

### HasLoad5

`func (o *Metrics) HasLoad5() bool`

HasLoad5 returns a boolean if a field has been set.

### GetMemFree

`func (o *Metrics) GetMemFree() int64`

GetMemFree returns the MemFree field if non-nil, zero value otherwise.

### GetMemFreeOk

`func (o *Metrics) GetMemFreeOk() (*int64, bool)`

GetMemFreeOk returns a tuple with the MemFree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemFree

`func (o *Metrics) SetMemFree(v int64)`

SetMemFree sets MemFree field to given value.

### HasMemFree

`func (o *Metrics) HasMemFree() bool`

HasMemFree returns a boolean if a field has been set.

### GetMemUsed

`func (o *Metrics) GetMemUsed() int64`

GetMemUsed returns the MemUsed field if non-nil, zero value otherwise.

### GetMemUsedOk

`func (o *Metrics) GetMemUsedOk() (*int64, bool)`

GetMemUsedOk returns a tuple with the MemUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemUsed

`func (o *Metrics) SetMemUsed(v int64)`

SetMemUsed sets MemUsed field to given value.

### HasMemUsed

`func (o *Metrics) HasMemUsed() bool`

HasMemUsed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


