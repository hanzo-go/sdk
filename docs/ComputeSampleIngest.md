# ComputeSampleIngest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GpuModel** | Pointer to **string** | GPUModel names the representative accelerator (\&quot;GB10\&quot;); GPUs carries how many. A heterogeneous host names its first card rather than inventing a summary. | [optional] 
**GpuUtil** | Pointer to **float64** | GPUUtil is accelerator utilization as a fraction 0..1; the warehouse clamps anything outside that. | [optional] 
**Gpus** | Pointer to **int64** | GPUs is how many accelerators this reading covers. | [optional] 
**Host** | Pointer to **string** | Host is the node&#39;s hostname, for display. | [optional] 
**MemFree** | Pointer to **int64** | MemFree is host memory still available, in BYTES. | [optional] 
**MemUsed** | Pointer to **int64** | MemUsed is host memory in use, in BYTES. | [optional] 
**Unit** | Pointer to **string** | Unit is the reporting node&#39;s own id — the same id it registered under, and the key the board joins this series onto. Required. | [optional] 

## Methods

### NewComputeSampleIngest

`func NewComputeSampleIngest() *ComputeSampleIngest`

NewComputeSampleIngest instantiates a new ComputeSampleIngest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputeSampleIngestWithDefaults

`func NewComputeSampleIngestWithDefaults() *ComputeSampleIngest`

NewComputeSampleIngestWithDefaults instantiates a new ComputeSampleIngest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGpuModel

`func (o *ComputeSampleIngest) GetGpuModel() string`

GetGpuModel returns the GpuModel field if non-nil, zero value otherwise.

### GetGpuModelOk

`func (o *ComputeSampleIngest) GetGpuModelOk() (*string, bool)`

GetGpuModelOk returns a tuple with the GpuModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuModel

`func (o *ComputeSampleIngest) SetGpuModel(v string)`

SetGpuModel sets GpuModel field to given value.

### HasGpuModel

`func (o *ComputeSampleIngest) HasGpuModel() bool`

HasGpuModel returns a boolean if a field has been set.

### GetGpuUtil

`func (o *ComputeSampleIngest) GetGpuUtil() float64`

GetGpuUtil returns the GpuUtil field if non-nil, zero value otherwise.

### GetGpuUtilOk

`func (o *ComputeSampleIngest) GetGpuUtilOk() (*float64, bool)`

GetGpuUtilOk returns a tuple with the GpuUtil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuUtil

`func (o *ComputeSampleIngest) SetGpuUtil(v float64)`

SetGpuUtil sets GpuUtil field to given value.

### HasGpuUtil

`func (o *ComputeSampleIngest) HasGpuUtil() bool`

HasGpuUtil returns a boolean if a field has been set.

### GetGpus

`func (o *ComputeSampleIngest) GetGpus() int64`

GetGpus returns the Gpus field if non-nil, zero value otherwise.

### GetGpusOk

`func (o *ComputeSampleIngest) GetGpusOk() (*int64, bool)`

GetGpusOk returns a tuple with the Gpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpus

`func (o *ComputeSampleIngest) SetGpus(v int64)`

SetGpus sets Gpus field to given value.

### HasGpus

`func (o *ComputeSampleIngest) HasGpus() bool`

HasGpus returns a boolean if a field has been set.

### GetHost

`func (o *ComputeSampleIngest) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *ComputeSampleIngest) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *ComputeSampleIngest) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *ComputeSampleIngest) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetMemFree

`func (o *ComputeSampleIngest) GetMemFree() int64`

GetMemFree returns the MemFree field if non-nil, zero value otherwise.

### GetMemFreeOk

`func (o *ComputeSampleIngest) GetMemFreeOk() (*int64, bool)`

GetMemFreeOk returns a tuple with the MemFree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemFree

`func (o *ComputeSampleIngest) SetMemFree(v int64)`

SetMemFree sets MemFree field to given value.

### HasMemFree

`func (o *ComputeSampleIngest) HasMemFree() bool`

HasMemFree returns a boolean if a field has been set.

### GetMemUsed

`func (o *ComputeSampleIngest) GetMemUsed() int64`

GetMemUsed returns the MemUsed field if non-nil, zero value otherwise.

### GetMemUsedOk

`func (o *ComputeSampleIngest) GetMemUsedOk() (*int64, bool)`

GetMemUsedOk returns a tuple with the MemUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemUsed

`func (o *ComputeSampleIngest) SetMemUsed(v int64)`

SetMemUsed sets MemUsed field to given value.

### HasMemUsed

`func (o *ComputeSampleIngest) HasMemUsed() bool`

HasMemUsed returns a boolean if a field has been set.

### GetUnit

`func (o *ComputeSampleIngest) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *ComputeSampleIngest) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *ComputeSampleIngest) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *ComputeSampleIngest) HasUnit() bool`

HasUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


