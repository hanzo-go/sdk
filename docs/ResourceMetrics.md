# ResourceMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CpuPercentage** | Pointer to **float32** |  | [optional] 
**CpuUsage** | Pointer to **string** |  | [optional] 
**MemoryPercentage** | Pointer to **float32** |  | [optional] 
**MemoryUsage** | Pointer to **string** |  | [optional] 
**PodCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewResourceMetrics

`func NewResourceMetrics() *ResourceMetrics`

NewResourceMetrics instantiates a new ResourceMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceMetricsWithDefaults

`func NewResourceMetricsWithDefaults() *ResourceMetrics`

NewResourceMetricsWithDefaults instantiates a new ResourceMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpuPercentage

`func (o *ResourceMetrics) GetCpuPercentage() float32`

GetCpuPercentage returns the CpuPercentage field if non-nil, zero value otherwise.

### GetCpuPercentageOk

`func (o *ResourceMetrics) GetCpuPercentageOk() (*float32, bool)`

GetCpuPercentageOk returns a tuple with the CpuPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuPercentage

`func (o *ResourceMetrics) SetCpuPercentage(v float32)`

SetCpuPercentage sets CpuPercentage field to given value.

### HasCpuPercentage

`func (o *ResourceMetrics) HasCpuPercentage() bool`

HasCpuPercentage returns a boolean if a field has been set.

### GetCpuUsage

`func (o *ResourceMetrics) GetCpuUsage() string`

GetCpuUsage returns the CpuUsage field if non-nil, zero value otherwise.

### GetCpuUsageOk

`func (o *ResourceMetrics) GetCpuUsageOk() (*string, bool)`

GetCpuUsageOk returns a tuple with the CpuUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUsage

`func (o *ResourceMetrics) SetCpuUsage(v string)`

SetCpuUsage sets CpuUsage field to given value.

### HasCpuUsage

`func (o *ResourceMetrics) HasCpuUsage() bool`

HasCpuUsage returns a boolean if a field has been set.

### GetMemoryPercentage

`func (o *ResourceMetrics) GetMemoryPercentage() float32`

GetMemoryPercentage returns the MemoryPercentage field if non-nil, zero value otherwise.

### GetMemoryPercentageOk

`func (o *ResourceMetrics) GetMemoryPercentageOk() (*float32, bool)`

GetMemoryPercentageOk returns a tuple with the MemoryPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryPercentage

`func (o *ResourceMetrics) SetMemoryPercentage(v float32)`

SetMemoryPercentage sets MemoryPercentage field to given value.

### HasMemoryPercentage

`func (o *ResourceMetrics) HasMemoryPercentage() bool`

HasMemoryPercentage returns a boolean if a field has been set.

### GetMemoryUsage

`func (o *ResourceMetrics) GetMemoryUsage() string`

GetMemoryUsage returns the MemoryUsage field if non-nil, zero value otherwise.

### GetMemoryUsageOk

`func (o *ResourceMetrics) GetMemoryUsageOk() (*string, bool)`

GetMemoryUsageOk returns a tuple with the MemoryUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryUsage

`func (o *ResourceMetrics) SetMemoryUsage(v string)`

SetMemoryUsage sets MemoryUsage field to given value.

### HasMemoryUsage

`func (o *ResourceMetrics) HasMemoryUsage() bool`

HasMemoryUsage returns a boolean if a field has been set.

### GetPodCount

`func (o *ResourceMetrics) GetPodCount() int32`

GetPodCount returns the PodCount field if non-nil, zero value otherwise.

### GetPodCountOk

`func (o *ResourceMetrics) GetPodCountOk() (*int32, bool)`

GetPodCountOk returns a tuple with the PodCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodCount

`func (o *ResourceMetrics) SetPodCount(v int32)`

SetPodCount sets PodCount field to given value.

### HasPodCount

`func (o *ResourceMetrics) HasPodCount() bool`

HasPodCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


