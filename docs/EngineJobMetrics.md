# EngineJobMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GpuUtilization** | Pointer to **float32** | Average GPU utilization percent | [optional] 
**GpuMemoryUsedMb** | Pointer to **int32** |  | [optional] 
**ThroughputSamplesPerSec** | Pointer to **float32** |  | [optional] 
**Loss** | Pointer to **float32** |  | [optional] 
**Epoch** | Pointer to **int32** |  | [optional] 
**Step** | Pointer to **int32** |  | [optional] 

## Methods

### NewEngineJobMetrics

`func NewEngineJobMetrics() *EngineJobMetrics`

NewEngineJobMetrics instantiates a new EngineJobMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEngineJobMetricsWithDefaults

`func NewEngineJobMetricsWithDefaults() *EngineJobMetrics`

NewEngineJobMetricsWithDefaults instantiates a new EngineJobMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGpuUtilization

`func (o *EngineJobMetrics) GetGpuUtilization() float32`

GetGpuUtilization returns the GpuUtilization field if non-nil, zero value otherwise.

### GetGpuUtilizationOk

`func (o *EngineJobMetrics) GetGpuUtilizationOk() (*float32, bool)`

GetGpuUtilizationOk returns a tuple with the GpuUtilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuUtilization

`func (o *EngineJobMetrics) SetGpuUtilization(v float32)`

SetGpuUtilization sets GpuUtilization field to given value.

### HasGpuUtilization

`func (o *EngineJobMetrics) HasGpuUtilization() bool`

HasGpuUtilization returns a boolean if a field has been set.

### GetGpuMemoryUsedMb

`func (o *EngineJobMetrics) GetGpuMemoryUsedMb() int32`

GetGpuMemoryUsedMb returns the GpuMemoryUsedMb field if non-nil, zero value otherwise.

### GetGpuMemoryUsedMbOk

`func (o *EngineJobMetrics) GetGpuMemoryUsedMbOk() (*int32, bool)`

GetGpuMemoryUsedMbOk returns a tuple with the GpuMemoryUsedMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuMemoryUsedMb

`func (o *EngineJobMetrics) SetGpuMemoryUsedMb(v int32)`

SetGpuMemoryUsedMb sets GpuMemoryUsedMb field to given value.

### HasGpuMemoryUsedMb

`func (o *EngineJobMetrics) HasGpuMemoryUsedMb() bool`

HasGpuMemoryUsedMb returns a boolean if a field has been set.

### GetThroughputSamplesPerSec

`func (o *EngineJobMetrics) GetThroughputSamplesPerSec() float32`

GetThroughputSamplesPerSec returns the ThroughputSamplesPerSec field if non-nil, zero value otherwise.

### GetThroughputSamplesPerSecOk

`func (o *EngineJobMetrics) GetThroughputSamplesPerSecOk() (*float32, bool)`

GetThroughputSamplesPerSecOk returns a tuple with the ThroughputSamplesPerSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThroughputSamplesPerSec

`func (o *EngineJobMetrics) SetThroughputSamplesPerSec(v float32)`

SetThroughputSamplesPerSec sets ThroughputSamplesPerSec field to given value.

### HasThroughputSamplesPerSec

`func (o *EngineJobMetrics) HasThroughputSamplesPerSec() bool`

HasThroughputSamplesPerSec returns a boolean if a field has been set.

### GetLoss

`func (o *EngineJobMetrics) GetLoss() float32`

GetLoss returns the Loss field if non-nil, zero value otherwise.

### GetLossOk

`func (o *EngineJobMetrics) GetLossOk() (*float32, bool)`

GetLossOk returns a tuple with the Loss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoss

`func (o *EngineJobMetrics) SetLoss(v float32)`

SetLoss sets Loss field to given value.

### HasLoss

`func (o *EngineJobMetrics) HasLoss() bool`

HasLoss returns a boolean if a field has been set.

### GetEpoch

`func (o *EngineJobMetrics) GetEpoch() int32`

GetEpoch returns the Epoch field if non-nil, zero value otherwise.

### GetEpochOk

`func (o *EngineJobMetrics) GetEpochOk() (*int32, bool)`

GetEpochOk returns a tuple with the Epoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpoch

`func (o *EngineJobMetrics) SetEpoch(v int32)`

SetEpoch sets Epoch field to given value.

### HasEpoch

`func (o *EngineJobMetrics) HasEpoch() bool`

HasEpoch returns a boolean if a field has been set.

### GetStep

`func (o *EngineJobMetrics) GetStep() int32`

GetStep returns the Step field if non-nil, zero value otherwise.

### GetStepOk

`func (o *EngineJobMetrics) GetStepOk() (*int32, bool)`

GetStepOk returns a tuple with the Step field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStep

`func (o *EngineJobMetrics) SetStep(v int32)`

SetStep sets Step field to given value.

### HasStep

`func (o *EngineJobMetrics) HasStep() bool`

HasStep returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


