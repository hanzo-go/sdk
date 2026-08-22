# SampleIngest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GpuModel** | Pointer to **string** | GPUModel names the representative accelerator (\&quot;GB10\&quot;); GPUs carries how many. A heterogeneous host names its first card rather than inventing a summary. | [optional] 
**GpuUtil** | Pointer to **float32** | GPUUtil is accelerator utilization as a fraction 0..1; the warehouse clamps anything outside that. | [optional] 
**Gpus** | Pointer to **int32** | GPUs is how many accelerators this reading covers. | [optional] 
**Host** | Pointer to **string** | Host is the node&#39;s hostname, for display. | [optional] 
**MemFree** | Pointer to **int32** | MemFree is host memory still available, in BYTES. | [optional] 
**MemUsed** | Pointer to **int32** | MemUsed is host memory in use, in BYTES. | [optional] 
**Unit** | Pointer to **string** | Unit is the reporting node&#39;s own id — the same id it registered under, and the key the board joins this series onto. Required. | [optional] 

## Methods

### NewSampleIngest

`func NewSampleIngest() *SampleIngest`

NewSampleIngest instantiates a new SampleIngest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSampleIngestWithDefaults

`func NewSampleIngestWithDefaults() *SampleIngest`

NewSampleIngestWithDefaults instantiates a new SampleIngest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGpuModel

`func (o *SampleIngest) GetGpuModel() string`

GetGpuModel returns the GpuModel field if non-nil, zero value otherwise.

### GetGpuModelOk

`func (o *SampleIngest) GetGpuModelOk() (*string, bool)`

GetGpuModelOk returns a tuple with the GpuModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuModel

`func (o *SampleIngest) SetGpuModel(v string)`

SetGpuModel sets GpuModel field to given value.

### HasGpuModel

`func (o *SampleIngest) HasGpuModel() bool`

HasGpuModel returns a boolean if a field has been set.

### GetGpuUtil

`func (o *SampleIngest) GetGpuUtil() float32`

GetGpuUtil returns the GpuUtil field if non-nil, zero value otherwise.

### GetGpuUtilOk

`func (o *SampleIngest) GetGpuUtilOk() (*float32, bool)`

GetGpuUtilOk returns a tuple with the GpuUtil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuUtil

`func (o *SampleIngest) SetGpuUtil(v float32)`

SetGpuUtil sets GpuUtil field to given value.

### HasGpuUtil

`func (o *SampleIngest) HasGpuUtil() bool`

HasGpuUtil returns a boolean if a field has been set.

### GetGpus

`func (o *SampleIngest) GetGpus() int32`

GetGpus returns the Gpus field if non-nil, zero value otherwise.

### GetGpusOk

`func (o *SampleIngest) GetGpusOk() (*int32, bool)`

GetGpusOk returns a tuple with the Gpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpus

`func (o *SampleIngest) SetGpus(v int32)`

SetGpus sets Gpus field to given value.

### HasGpus

`func (o *SampleIngest) HasGpus() bool`

HasGpus returns a boolean if a field has been set.

### GetHost

`func (o *SampleIngest) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *SampleIngest) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *SampleIngest) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *SampleIngest) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetMemFree

`func (o *SampleIngest) GetMemFree() int32`

GetMemFree returns the MemFree field if non-nil, zero value otherwise.

### GetMemFreeOk

`func (o *SampleIngest) GetMemFreeOk() (*int32, bool)`

GetMemFreeOk returns a tuple with the MemFree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemFree

`func (o *SampleIngest) SetMemFree(v int32)`

SetMemFree sets MemFree field to given value.

### HasMemFree

`func (o *SampleIngest) HasMemFree() bool`

HasMemFree returns a boolean if a field has been set.

### GetMemUsed

`func (o *SampleIngest) GetMemUsed() int32`

GetMemUsed returns the MemUsed field if non-nil, zero value otherwise.

### GetMemUsedOk

`func (o *SampleIngest) GetMemUsedOk() (*int32, bool)`

GetMemUsedOk returns a tuple with the MemUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemUsed

`func (o *SampleIngest) SetMemUsed(v int32)`

SetMemUsed sets MemUsed field to given value.

### HasMemUsed

`func (o *SampleIngest) HasMemUsed() bool`

HasMemUsed returns a boolean if a field has been set.

### GetUnit

`func (o *SampleIngest) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *SampleIngest) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *SampleIngest) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *SampleIngest) HasUnit() bool`

HasUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


