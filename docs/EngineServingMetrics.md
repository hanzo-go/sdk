# EngineServingMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndpointName** | Pointer to **string** |  | [optional] 
**Period** | Pointer to **string** |  | [optional] 
**RequestCount** | Pointer to **int64** |  | [optional] 
**ErrorCount** | Pointer to **int64** |  | [optional] 
**AvgLatencyMs** | Pointer to **float32** |  | [optional] 
**P50LatencyMs** | Pointer to **float32** |  | [optional] 
**P95LatencyMs** | Pointer to **float32** |  | [optional] 
**P99LatencyMs** | Pointer to **float32** |  | [optional] 
**TokensPerSecond** | Pointer to **float32** |  | [optional] 
**GpuUtilization** | Pointer to **float32** |  | [optional] 
**GpuMemoryUsedMb** | Pointer to **int32** |  | [optional] 

## Methods

### NewEngineServingMetrics

`func NewEngineServingMetrics() *EngineServingMetrics`

NewEngineServingMetrics instantiates a new EngineServingMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEngineServingMetricsWithDefaults

`func NewEngineServingMetricsWithDefaults() *EngineServingMetrics`

NewEngineServingMetricsWithDefaults instantiates a new EngineServingMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndpointName

`func (o *EngineServingMetrics) GetEndpointName() string`

GetEndpointName returns the EndpointName field if non-nil, zero value otherwise.

### GetEndpointNameOk

`func (o *EngineServingMetrics) GetEndpointNameOk() (*string, bool)`

GetEndpointNameOk returns a tuple with the EndpointName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointName

`func (o *EngineServingMetrics) SetEndpointName(v string)`

SetEndpointName sets EndpointName field to given value.

### HasEndpointName

`func (o *EngineServingMetrics) HasEndpointName() bool`

HasEndpointName returns a boolean if a field has been set.

### GetPeriod

`func (o *EngineServingMetrics) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *EngineServingMetrics) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *EngineServingMetrics) SetPeriod(v string)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *EngineServingMetrics) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetRequestCount

`func (o *EngineServingMetrics) GetRequestCount() int64`

GetRequestCount returns the RequestCount field if non-nil, zero value otherwise.

### GetRequestCountOk

`func (o *EngineServingMetrics) GetRequestCountOk() (*int64, bool)`

GetRequestCountOk returns a tuple with the RequestCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestCount

`func (o *EngineServingMetrics) SetRequestCount(v int64)`

SetRequestCount sets RequestCount field to given value.

### HasRequestCount

`func (o *EngineServingMetrics) HasRequestCount() bool`

HasRequestCount returns a boolean if a field has been set.

### GetErrorCount

`func (o *EngineServingMetrics) GetErrorCount() int64`

GetErrorCount returns the ErrorCount field if non-nil, zero value otherwise.

### GetErrorCountOk

`func (o *EngineServingMetrics) GetErrorCountOk() (*int64, bool)`

GetErrorCountOk returns a tuple with the ErrorCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCount

`func (o *EngineServingMetrics) SetErrorCount(v int64)`

SetErrorCount sets ErrorCount field to given value.

### HasErrorCount

`func (o *EngineServingMetrics) HasErrorCount() bool`

HasErrorCount returns a boolean if a field has been set.

### GetAvgLatencyMs

`func (o *EngineServingMetrics) GetAvgLatencyMs() float32`

GetAvgLatencyMs returns the AvgLatencyMs field if non-nil, zero value otherwise.

### GetAvgLatencyMsOk

`func (o *EngineServingMetrics) GetAvgLatencyMsOk() (*float32, bool)`

GetAvgLatencyMsOk returns a tuple with the AvgLatencyMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgLatencyMs

`func (o *EngineServingMetrics) SetAvgLatencyMs(v float32)`

SetAvgLatencyMs sets AvgLatencyMs field to given value.

### HasAvgLatencyMs

`func (o *EngineServingMetrics) HasAvgLatencyMs() bool`

HasAvgLatencyMs returns a boolean if a field has been set.

### GetP50LatencyMs

`func (o *EngineServingMetrics) GetP50LatencyMs() float32`

GetP50LatencyMs returns the P50LatencyMs field if non-nil, zero value otherwise.

### GetP50LatencyMsOk

`func (o *EngineServingMetrics) GetP50LatencyMsOk() (*float32, bool)`

GetP50LatencyMsOk returns a tuple with the P50LatencyMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP50LatencyMs

`func (o *EngineServingMetrics) SetP50LatencyMs(v float32)`

SetP50LatencyMs sets P50LatencyMs field to given value.

### HasP50LatencyMs

`func (o *EngineServingMetrics) HasP50LatencyMs() bool`

HasP50LatencyMs returns a boolean if a field has been set.

### GetP95LatencyMs

`func (o *EngineServingMetrics) GetP95LatencyMs() float32`

GetP95LatencyMs returns the P95LatencyMs field if non-nil, zero value otherwise.

### GetP95LatencyMsOk

`func (o *EngineServingMetrics) GetP95LatencyMsOk() (*float32, bool)`

GetP95LatencyMsOk returns a tuple with the P95LatencyMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP95LatencyMs

`func (o *EngineServingMetrics) SetP95LatencyMs(v float32)`

SetP95LatencyMs sets P95LatencyMs field to given value.

### HasP95LatencyMs

`func (o *EngineServingMetrics) HasP95LatencyMs() bool`

HasP95LatencyMs returns a boolean if a field has been set.

### GetP99LatencyMs

`func (o *EngineServingMetrics) GetP99LatencyMs() float32`

GetP99LatencyMs returns the P99LatencyMs field if non-nil, zero value otherwise.

### GetP99LatencyMsOk

`func (o *EngineServingMetrics) GetP99LatencyMsOk() (*float32, bool)`

GetP99LatencyMsOk returns a tuple with the P99LatencyMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP99LatencyMs

`func (o *EngineServingMetrics) SetP99LatencyMs(v float32)`

SetP99LatencyMs sets P99LatencyMs field to given value.

### HasP99LatencyMs

`func (o *EngineServingMetrics) HasP99LatencyMs() bool`

HasP99LatencyMs returns a boolean if a field has been set.

### GetTokensPerSecond

`func (o *EngineServingMetrics) GetTokensPerSecond() float32`

GetTokensPerSecond returns the TokensPerSecond field if non-nil, zero value otherwise.

### GetTokensPerSecondOk

`func (o *EngineServingMetrics) GetTokensPerSecondOk() (*float32, bool)`

GetTokensPerSecondOk returns a tuple with the TokensPerSecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokensPerSecond

`func (o *EngineServingMetrics) SetTokensPerSecond(v float32)`

SetTokensPerSecond sets TokensPerSecond field to given value.

### HasTokensPerSecond

`func (o *EngineServingMetrics) HasTokensPerSecond() bool`

HasTokensPerSecond returns a boolean if a field has been set.

### GetGpuUtilization

`func (o *EngineServingMetrics) GetGpuUtilization() float32`

GetGpuUtilization returns the GpuUtilization field if non-nil, zero value otherwise.

### GetGpuUtilizationOk

`func (o *EngineServingMetrics) GetGpuUtilizationOk() (*float32, bool)`

GetGpuUtilizationOk returns a tuple with the GpuUtilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuUtilization

`func (o *EngineServingMetrics) SetGpuUtilization(v float32)`

SetGpuUtilization sets GpuUtilization field to given value.

### HasGpuUtilization

`func (o *EngineServingMetrics) HasGpuUtilization() bool`

HasGpuUtilization returns a boolean if a field has been set.

### GetGpuMemoryUsedMb

`func (o *EngineServingMetrics) GetGpuMemoryUsedMb() int32`

GetGpuMemoryUsedMb returns the GpuMemoryUsedMb field if non-nil, zero value otherwise.

### GetGpuMemoryUsedMbOk

`func (o *EngineServingMetrics) GetGpuMemoryUsedMbOk() (*int32, bool)`

GetGpuMemoryUsedMbOk returns a tuple with the GpuMemoryUsedMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuMemoryUsedMb

`func (o *EngineServingMetrics) SetGpuMemoryUsedMb(v int32)`

SetGpuMemoryUsedMb sets GpuMemoryUsedMb field to given value.

### HasGpuMemoryUsedMb

`func (o *EngineServingMetrics) HasGpuMemoryUsedMb() bool`

HasGpuMemoryUsedMb returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


