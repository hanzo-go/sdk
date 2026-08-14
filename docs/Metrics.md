# Metrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**At** | Pointer to **int32** | unix seconds, server-stamped | [optional] 
**GpuUtil** | Pointer to **float32** | 0..1 aggregate utilization | [optional] 
**Load1** | Pointer to **float32** |  | [optional] 
**Load5** | Pointer to **float32** |  | [optional] 
**Load15** | Pointer to **float32** |  | [optional] 
**MemFree** | Pointer to **int32** | bytes | [optional] 
**MemUsed** | Pointer to **int32** | bytes | [optional] 

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

`func (o *Metrics) GetAt() int32`

GetAt returns the At field if non-nil, zero value otherwise.

### GetAtOk

`func (o *Metrics) GetAtOk() (*int32, bool)`

GetAtOk returns a tuple with the At field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAt

`func (o *Metrics) SetAt(v int32)`

SetAt sets At field to given value.

### HasAt

`func (o *Metrics) HasAt() bool`

HasAt returns a boolean if a field has been set.

### GetGpuUtil

`func (o *Metrics) GetGpuUtil() float32`

GetGpuUtil returns the GpuUtil field if non-nil, zero value otherwise.

### GetGpuUtilOk

`func (o *Metrics) GetGpuUtilOk() (*float32, bool)`

GetGpuUtilOk returns a tuple with the GpuUtil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuUtil

`func (o *Metrics) SetGpuUtil(v float32)`

SetGpuUtil sets GpuUtil field to given value.

### HasGpuUtil

`func (o *Metrics) HasGpuUtil() bool`

HasGpuUtil returns a boolean if a field has been set.

### GetLoad1

`func (o *Metrics) GetLoad1() float32`

GetLoad1 returns the Load1 field if non-nil, zero value otherwise.

### GetLoad1Ok

`func (o *Metrics) GetLoad1Ok() (*float32, bool)`

GetLoad1Ok returns a tuple with the Load1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoad1

`func (o *Metrics) SetLoad1(v float32)`

SetLoad1 sets Load1 field to given value.

### HasLoad1

`func (o *Metrics) HasLoad1() bool`

HasLoad1 returns a boolean if a field has been set.

### GetLoad5

`func (o *Metrics) GetLoad5() float32`

GetLoad5 returns the Load5 field if non-nil, zero value otherwise.

### GetLoad5Ok

`func (o *Metrics) GetLoad5Ok() (*float32, bool)`

GetLoad5Ok returns a tuple with the Load5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoad5

`func (o *Metrics) SetLoad5(v float32)`

SetLoad5 sets Load5 field to given value.

### HasLoad5

`func (o *Metrics) HasLoad5() bool`

HasLoad5 returns a boolean if a field has been set.

### GetLoad15

`func (o *Metrics) GetLoad15() float32`

GetLoad15 returns the Load15 field if non-nil, zero value otherwise.

### GetLoad15Ok

`func (o *Metrics) GetLoad15Ok() (*float32, bool)`

GetLoad15Ok returns a tuple with the Load15 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoad15

`func (o *Metrics) SetLoad15(v float32)`

SetLoad15 sets Load15 field to given value.

### HasLoad15

`func (o *Metrics) HasLoad15() bool`

HasLoad15 returns a boolean if a field has been set.

### GetMemFree

`func (o *Metrics) GetMemFree() int32`

GetMemFree returns the MemFree field if non-nil, zero value otherwise.

### GetMemFreeOk

`func (o *Metrics) GetMemFreeOk() (*int32, bool)`

GetMemFreeOk returns a tuple with the MemFree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemFree

`func (o *Metrics) SetMemFree(v int32)`

SetMemFree sets MemFree field to given value.

### HasMemFree

`func (o *Metrics) HasMemFree() bool`

HasMemFree returns a boolean if a field has been set.

### GetMemUsed

`func (o *Metrics) GetMemUsed() int32`

GetMemUsed returns the MemUsed field if non-nil, zero value otherwise.

### GetMemUsedOk

`func (o *Metrics) GetMemUsedOk() (*int32, bool)`

GetMemUsedOk returns a tuple with the MemUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemUsed

`func (o *Metrics) SetMemUsed(v int32)`

SetMemUsed sets MemUsed field to given value.

### HasMemUsed

`func (o *Metrics) HasMemUsed() bool`

HasMemUsed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


