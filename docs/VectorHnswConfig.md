# VectorHnswConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**M** | Pointer to **int32** |  | [optional] [default to 16]
**EfConstruct** | Pointer to **int32** |  | [optional] [default to 100]
**FullScanThreshold** | Pointer to **int32** |  | [optional] 
**MaxIndexingThreads** | Pointer to **int32** |  | [optional] 
**OnDisk** | Pointer to **bool** |  | [optional] 

## Methods

### NewVectorHnswConfig

`func NewVectorHnswConfig() *VectorHnswConfig`

NewVectorHnswConfig instantiates a new VectorHnswConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVectorHnswConfigWithDefaults

`func NewVectorHnswConfigWithDefaults() *VectorHnswConfig`

NewVectorHnswConfigWithDefaults instantiates a new VectorHnswConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetM

`func (o *VectorHnswConfig) GetM() int32`

GetM returns the M field if non-nil, zero value otherwise.

### GetMOk

`func (o *VectorHnswConfig) GetMOk() (*int32, bool)`

GetMOk returns a tuple with the M field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetM

`func (o *VectorHnswConfig) SetM(v int32)`

SetM sets M field to given value.

### HasM

`func (o *VectorHnswConfig) HasM() bool`

HasM returns a boolean if a field has been set.

### GetEfConstruct

`func (o *VectorHnswConfig) GetEfConstruct() int32`

GetEfConstruct returns the EfConstruct field if non-nil, zero value otherwise.

### GetEfConstructOk

`func (o *VectorHnswConfig) GetEfConstructOk() (*int32, bool)`

GetEfConstructOk returns a tuple with the EfConstruct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEfConstruct

`func (o *VectorHnswConfig) SetEfConstruct(v int32)`

SetEfConstruct sets EfConstruct field to given value.

### HasEfConstruct

`func (o *VectorHnswConfig) HasEfConstruct() bool`

HasEfConstruct returns a boolean if a field has been set.

### GetFullScanThreshold

`func (o *VectorHnswConfig) GetFullScanThreshold() int32`

GetFullScanThreshold returns the FullScanThreshold field if non-nil, zero value otherwise.

### GetFullScanThresholdOk

`func (o *VectorHnswConfig) GetFullScanThresholdOk() (*int32, bool)`

GetFullScanThresholdOk returns a tuple with the FullScanThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullScanThreshold

`func (o *VectorHnswConfig) SetFullScanThreshold(v int32)`

SetFullScanThreshold sets FullScanThreshold field to given value.

### HasFullScanThreshold

`func (o *VectorHnswConfig) HasFullScanThreshold() bool`

HasFullScanThreshold returns a boolean if a field has been set.

### GetMaxIndexingThreads

`func (o *VectorHnswConfig) GetMaxIndexingThreads() int32`

GetMaxIndexingThreads returns the MaxIndexingThreads field if non-nil, zero value otherwise.

### GetMaxIndexingThreadsOk

`func (o *VectorHnswConfig) GetMaxIndexingThreadsOk() (*int32, bool)`

GetMaxIndexingThreadsOk returns a tuple with the MaxIndexingThreads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxIndexingThreads

`func (o *VectorHnswConfig) SetMaxIndexingThreads(v int32)`

SetMaxIndexingThreads sets MaxIndexingThreads field to given value.

### HasMaxIndexingThreads

`func (o *VectorHnswConfig) HasMaxIndexingThreads() bool`

HasMaxIndexingThreads returns a boolean if a field has been set.

### GetOnDisk

`func (o *VectorHnswConfig) GetOnDisk() bool`

GetOnDisk returns the OnDisk field if non-nil, zero value otherwise.

### GetOnDiskOk

`func (o *VectorHnswConfig) GetOnDiskOk() (*bool, bool)`

GetOnDiskOk returns a tuple with the OnDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnDisk

`func (o *VectorHnswConfig) SetOnDisk(v bool)`

SetOnDisk sets OnDisk field to given value.

### HasOnDisk

`func (o *VectorHnswConfig) HasOnDisk() bool`

HasOnDisk returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


