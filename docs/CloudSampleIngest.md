# CloudSampleIngest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GpuModel** | Pointer to **string** |  | [optional] 
**GpuUtil** | Pointer to **float32** | GPUUtil is accelerator utilization as a fraction 0..1; the warehouse clamps anything outside that. | [optional] 
**Gpus** | Pointer to **int32** | GPUs is how many accelerators the reading covers, GPUModel the representative model name. | [optional] 
**Host** | Pointer to **string** | Host is the node&#39;s hostname, for display. | [optional] 
**MemFree** | Pointer to **int32** |  | [optional] 
**MemUsed** | Pointer to **int32** | MemUsed and MemFree are host memory in bytes. | [optional] 
**Unit** | Pointer to **string** | Unit is the reporting node&#39;s own id — the same id it registered under, and the key the board joins this series onto. Required. | [optional] 

## Methods

### NewCloudSampleIngest

`func NewCloudSampleIngest() *CloudSampleIngest`

NewCloudSampleIngest instantiates a new CloudSampleIngest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudSampleIngestWithDefaults

`func NewCloudSampleIngestWithDefaults() *CloudSampleIngest`

NewCloudSampleIngestWithDefaults instantiates a new CloudSampleIngest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGpuModel

`func (o *CloudSampleIngest) GetGpuModel() string`

GetGpuModel returns the GpuModel field if non-nil, zero value otherwise.

### GetGpuModelOk

`func (o *CloudSampleIngest) GetGpuModelOk() (*string, bool)`

GetGpuModelOk returns a tuple with the GpuModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuModel

`func (o *CloudSampleIngest) SetGpuModel(v string)`

SetGpuModel sets GpuModel field to given value.

### HasGpuModel

`func (o *CloudSampleIngest) HasGpuModel() bool`

HasGpuModel returns a boolean if a field has been set.

### GetGpuUtil

`func (o *CloudSampleIngest) GetGpuUtil() float32`

GetGpuUtil returns the GpuUtil field if non-nil, zero value otherwise.

### GetGpuUtilOk

`func (o *CloudSampleIngest) GetGpuUtilOk() (*float32, bool)`

GetGpuUtilOk returns a tuple with the GpuUtil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuUtil

`func (o *CloudSampleIngest) SetGpuUtil(v float32)`

SetGpuUtil sets GpuUtil field to given value.

### HasGpuUtil

`func (o *CloudSampleIngest) HasGpuUtil() bool`

HasGpuUtil returns a boolean if a field has been set.

### GetGpus

`func (o *CloudSampleIngest) GetGpus() int32`

GetGpus returns the Gpus field if non-nil, zero value otherwise.

### GetGpusOk

`func (o *CloudSampleIngest) GetGpusOk() (*int32, bool)`

GetGpusOk returns a tuple with the Gpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpus

`func (o *CloudSampleIngest) SetGpus(v int32)`

SetGpus sets Gpus field to given value.

### HasGpus

`func (o *CloudSampleIngest) HasGpus() bool`

HasGpus returns a boolean if a field has been set.

### GetHost

`func (o *CloudSampleIngest) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *CloudSampleIngest) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *CloudSampleIngest) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *CloudSampleIngest) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetMemFree

`func (o *CloudSampleIngest) GetMemFree() int32`

GetMemFree returns the MemFree field if non-nil, zero value otherwise.

### GetMemFreeOk

`func (o *CloudSampleIngest) GetMemFreeOk() (*int32, bool)`

GetMemFreeOk returns a tuple with the MemFree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemFree

`func (o *CloudSampleIngest) SetMemFree(v int32)`

SetMemFree sets MemFree field to given value.

### HasMemFree

`func (o *CloudSampleIngest) HasMemFree() bool`

HasMemFree returns a boolean if a field has been set.

### GetMemUsed

`func (o *CloudSampleIngest) GetMemUsed() int32`

GetMemUsed returns the MemUsed field if non-nil, zero value otherwise.

### GetMemUsedOk

`func (o *CloudSampleIngest) GetMemUsedOk() (*int32, bool)`

GetMemUsedOk returns a tuple with the MemUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemUsed

`func (o *CloudSampleIngest) SetMemUsed(v int32)`

SetMemUsed sets MemUsed field to given value.

### HasMemUsed

`func (o *CloudSampleIngest) HasMemUsed() bool`

HasMemUsed returns a boolean if a field has been set.

### GetUnit

`func (o *CloudSampleIngest) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *CloudSampleIngest) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *CloudSampleIngest) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *CloudSampleIngest) HasUnit() bool`

HasUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


