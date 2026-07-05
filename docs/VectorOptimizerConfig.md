# VectorOptimizerConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeletedThreshold** | Pointer to **float32** |  | [optional] 
**VacuumMinVectorNumber** | Pointer to **int32** |  | [optional] 
**DefaultSegmentNumber** | Pointer to **int32** |  | [optional] 
**MaxSegmentSize** | Pointer to **int32** |  | [optional] 
**MemmapThreshold** | Pointer to **int32** |  | [optional] 
**IndexingThreshold** | Pointer to **int32** |  | [optional] 
**FlushIntervalSec** | Pointer to **int32** |  | [optional] 

## Methods

### NewVectorOptimizerConfig

`func NewVectorOptimizerConfig() *VectorOptimizerConfig`

NewVectorOptimizerConfig instantiates a new VectorOptimizerConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVectorOptimizerConfigWithDefaults

`func NewVectorOptimizerConfigWithDefaults() *VectorOptimizerConfig`

NewVectorOptimizerConfigWithDefaults instantiates a new VectorOptimizerConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeletedThreshold

`func (o *VectorOptimizerConfig) GetDeletedThreshold() float32`

GetDeletedThreshold returns the DeletedThreshold field if non-nil, zero value otherwise.

### GetDeletedThresholdOk

`func (o *VectorOptimizerConfig) GetDeletedThresholdOk() (*float32, bool)`

GetDeletedThresholdOk returns a tuple with the DeletedThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedThreshold

`func (o *VectorOptimizerConfig) SetDeletedThreshold(v float32)`

SetDeletedThreshold sets DeletedThreshold field to given value.

### HasDeletedThreshold

`func (o *VectorOptimizerConfig) HasDeletedThreshold() bool`

HasDeletedThreshold returns a boolean if a field has been set.

### GetVacuumMinVectorNumber

`func (o *VectorOptimizerConfig) GetVacuumMinVectorNumber() int32`

GetVacuumMinVectorNumber returns the VacuumMinVectorNumber field if non-nil, zero value otherwise.

### GetVacuumMinVectorNumberOk

`func (o *VectorOptimizerConfig) GetVacuumMinVectorNumberOk() (*int32, bool)`

GetVacuumMinVectorNumberOk returns a tuple with the VacuumMinVectorNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVacuumMinVectorNumber

`func (o *VectorOptimizerConfig) SetVacuumMinVectorNumber(v int32)`

SetVacuumMinVectorNumber sets VacuumMinVectorNumber field to given value.

### HasVacuumMinVectorNumber

`func (o *VectorOptimizerConfig) HasVacuumMinVectorNumber() bool`

HasVacuumMinVectorNumber returns a boolean if a field has been set.

### GetDefaultSegmentNumber

`func (o *VectorOptimizerConfig) GetDefaultSegmentNumber() int32`

GetDefaultSegmentNumber returns the DefaultSegmentNumber field if non-nil, zero value otherwise.

### GetDefaultSegmentNumberOk

`func (o *VectorOptimizerConfig) GetDefaultSegmentNumberOk() (*int32, bool)`

GetDefaultSegmentNumberOk returns a tuple with the DefaultSegmentNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSegmentNumber

`func (o *VectorOptimizerConfig) SetDefaultSegmentNumber(v int32)`

SetDefaultSegmentNumber sets DefaultSegmentNumber field to given value.

### HasDefaultSegmentNumber

`func (o *VectorOptimizerConfig) HasDefaultSegmentNumber() bool`

HasDefaultSegmentNumber returns a boolean if a field has been set.

### GetMaxSegmentSize

`func (o *VectorOptimizerConfig) GetMaxSegmentSize() int32`

GetMaxSegmentSize returns the MaxSegmentSize field if non-nil, zero value otherwise.

### GetMaxSegmentSizeOk

`func (o *VectorOptimizerConfig) GetMaxSegmentSizeOk() (*int32, bool)`

GetMaxSegmentSizeOk returns a tuple with the MaxSegmentSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSegmentSize

`func (o *VectorOptimizerConfig) SetMaxSegmentSize(v int32)`

SetMaxSegmentSize sets MaxSegmentSize field to given value.

### HasMaxSegmentSize

`func (o *VectorOptimizerConfig) HasMaxSegmentSize() bool`

HasMaxSegmentSize returns a boolean if a field has been set.

### GetMemmapThreshold

`func (o *VectorOptimizerConfig) GetMemmapThreshold() int32`

GetMemmapThreshold returns the MemmapThreshold field if non-nil, zero value otherwise.

### GetMemmapThresholdOk

`func (o *VectorOptimizerConfig) GetMemmapThresholdOk() (*int32, bool)`

GetMemmapThresholdOk returns a tuple with the MemmapThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemmapThreshold

`func (o *VectorOptimizerConfig) SetMemmapThreshold(v int32)`

SetMemmapThreshold sets MemmapThreshold field to given value.

### HasMemmapThreshold

`func (o *VectorOptimizerConfig) HasMemmapThreshold() bool`

HasMemmapThreshold returns a boolean if a field has been set.

### GetIndexingThreshold

`func (o *VectorOptimizerConfig) GetIndexingThreshold() int32`

GetIndexingThreshold returns the IndexingThreshold field if non-nil, zero value otherwise.

### GetIndexingThresholdOk

`func (o *VectorOptimizerConfig) GetIndexingThresholdOk() (*int32, bool)`

GetIndexingThresholdOk returns a tuple with the IndexingThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexingThreshold

`func (o *VectorOptimizerConfig) SetIndexingThreshold(v int32)`

SetIndexingThreshold sets IndexingThreshold field to given value.

### HasIndexingThreshold

`func (o *VectorOptimizerConfig) HasIndexingThreshold() bool`

HasIndexingThreshold returns a boolean if a field has been set.

### GetFlushIntervalSec

`func (o *VectorOptimizerConfig) GetFlushIntervalSec() int32`

GetFlushIntervalSec returns the FlushIntervalSec field if non-nil, zero value otherwise.

### GetFlushIntervalSecOk

`func (o *VectorOptimizerConfig) GetFlushIntervalSecOk() (*int32, bool)`

GetFlushIntervalSecOk returns a tuple with the FlushIntervalSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlushIntervalSec

`func (o *VectorOptimizerConfig) SetFlushIntervalSec(v int32)`

SetFlushIntervalSec sets FlushIntervalSec field to given value.

### HasFlushIntervalSec

`func (o *VectorOptimizerConfig) HasFlushIntervalSec() bool`

HasFlushIntervalSec returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


