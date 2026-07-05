# EngineServingEndpointCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Model** | **string** |  | 
**Framework** | Pointer to **string** |  | [optional] [default to "vllm"]
**GpuType** | Pointer to **string** |  | [optional] [default to "A100"]
**GpuPerReplica** | Pointer to **int32** |  | [optional] [default to 1]
**MinReplicas** | Pointer to **int32** |  | [optional] [default to 1]
**MaxReplicas** | Pointer to **int32** |  | [optional] [default to 4]
**Env** | Pointer to **map[string]string** |  | [optional] 
**ScaleToZero** | Pointer to **bool** |  | [optional] [default to false]
**MaxBatchSize** | Pointer to **int32** |  | [optional] 
**MaxConcurrentRequests** | Pointer to **int32** |  | [optional] 

## Methods

### NewEngineServingEndpointCreate

`func NewEngineServingEndpointCreate(name string, model string, ) *EngineServingEndpointCreate`

NewEngineServingEndpointCreate instantiates a new EngineServingEndpointCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEngineServingEndpointCreateWithDefaults

`func NewEngineServingEndpointCreateWithDefaults() *EngineServingEndpointCreate`

NewEngineServingEndpointCreateWithDefaults instantiates a new EngineServingEndpointCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EngineServingEndpointCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EngineServingEndpointCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EngineServingEndpointCreate) SetName(v string)`

SetName sets Name field to given value.


### GetModel

`func (o *EngineServingEndpointCreate) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *EngineServingEndpointCreate) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *EngineServingEndpointCreate) SetModel(v string)`

SetModel sets Model field to given value.


### GetFramework

`func (o *EngineServingEndpointCreate) GetFramework() string`

GetFramework returns the Framework field if non-nil, zero value otherwise.

### GetFrameworkOk

`func (o *EngineServingEndpointCreate) GetFrameworkOk() (*string, bool)`

GetFrameworkOk returns a tuple with the Framework field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFramework

`func (o *EngineServingEndpointCreate) SetFramework(v string)`

SetFramework sets Framework field to given value.

### HasFramework

`func (o *EngineServingEndpointCreate) HasFramework() bool`

HasFramework returns a boolean if a field has been set.

### GetGpuType

`func (o *EngineServingEndpointCreate) GetGpuType() string`

GetGpuType returns the GpuType field if non-nil, zero value otherwise.

### GetGpuTypeOk

`func (o *EngineServingEndpointCreate) GetGpuTypeOk() (*string, bool)`

GetGpuTypeOk returns a tuple with the GpuType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuType

`func (o *EngineServingEndpointCreate) SetGpuType(v string)`

SetGpuType sets GpuType field to given value.

### HasGpuType

`func (o *EngineServingEndpointCreate) HasGpuType() bool`

HasGpuType returns a boolean if a field has been set.

### GetGpuPerReplica

`func (o *EngineServingEndpointCreate) GetGpuPerReplica() int32`

GetGpuPerReplica returns the GpuPerReplica field if non-nil, zero value otherwise.

### GetGpuPerReplicaOk

`func (o *EngineServingEndpointCreate) GetGpuPerReplicaOk() (*int32, bool)`

GetGpuPerReplicaOk returns a tuple with the GpuPerReplica field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuPerReplica

`func (o *EngineServingEndpointCreate) SetGpuPerReplica(v int32)`

SetGpuPerReplica sets GpuPerReplica field to given value.

### HasGpuPerReplica

`func (o *EngineServingEndpointCreate) HasGpuPerReplica() bool`

HasGpuPerReplica returns a boolean if a field has been set.

### GetMinReplicas

`func (o *EngineServingEndpointCreate) GetMinReplicas() int32`

GetMinReplicas returns the MinReplicas field if non-nil, zero value otherwise.

### GetMinReplicasOk

`func (o *EngineServingEndpointCreate) GetMinReplicasOk() (*int32, bool)`

GetMinReplicasOk returns a tuple with the MinReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinReplicas

`func (o *EngineServingEndpointCreate) SetMinReplicas(v int32)`

SetMinReplicas sets MinReplicas field to given value.

### HasMinReplicas

`func (o *EngineServingEndpointCreate) HasMinReplicas() bool`

HasMinReplicas returns a boolean if a field has been set.

### GetMaxReplicas

`func (o *EngineServingEndpointCreate) GetMaxReplicas() int32`

GetMaxReplicas returns the MaxReplicas field if non-nil, zero value otherwise.

### GetMaxReplicasOk

`func (o *EngineServingEndpointCreate) GetMaxReplicasOk() (*int32, bool)`

GetMaxReplicasOk returns a tuple with the MaxReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxReplicas

`func (o *EngineServingEndpointCreate) SetMaxReplicas(v int32)`

SetMaxReplicas sets MaxReplicas field to given value.

### HasMaxReplicas

`func (o *EngineServingEndpointCreate) HasMaxReplicas() bool`

HasMaxReplicas returns a boolean if a field has been set.

### GetEnv

`func (o *EngineServingEndpointCreate) GetEnv() map[string]string`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *EngineServingEndpointCreate) GetEnvOk() (*map[string]string, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *EngineServingEndpointCreate) SetEnv(v map[string]string)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *EngineServingEndpointCreate) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### GetScaleToZero

`func (o *EngineServingEndpointCreate) GetScaleToZero() bool`

GetScaleToZero returns the ScaleToZero field if non-nil, zero value otherwise.

### GetScaleToZeroOk

`func (o *EngineServingEndpointCreate) GetScaleToZeroOk() (*bool, bool)`

GetScaleToZeroOk returns a tuple with the ScaleToZero field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaleToZero

`func (o *EngineServingEndpointCreate) SetScaleToZero(v bool)`

SetScaleToZero sets ScaleToZero field to given value.

### HasScaleToZero

`func (o *EngineServingEndpointCreate) HasScaleToZero() bool`

HasScaleToZero returns a boolean if a field has been set.

### GetMaxBatchSize

`func (o *EngineServingEndpointCreate) GetMaxBatchSize() int32`

GetMaxBatchSize returns the MaxBatchSize field if non-nil, zero value otherwise.

### GetMaxBatchSizeOk

`func (o *EngineServingEndpointCreate) GetMaxBatchSizeOk() (*int32, bool)`

GetMaxBatchSizeOk returns a tuple with the MaxBatchSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxBatchSize

`func (o *EngineServingEndpointCreate) SetMaxBatchSize(v int32)`

SetMaxBatchSize sets MaxBatchSize field to given value.

### HasMaxBatchSize

`func (o *EngineServingEndpointCreate) HasMaxBatchSize() bool`

HasMaxBatchSize returns a boolean if a field has been set.

### GetMaxConcurrentRequests

`func (o *EngineServingEndpointCreate) GetMaxConcurrentRequests() int32`

GetMaxConcurrentRequests returns the MaxConcurrentRequests field if non-nil, zero value otherwise.

### GetMaxConcurrentRequestsOk

`func (o *EngineServingEndpointCreate) GetMaxConcurrentRequestsOk() (*int32, bool)`

GetMaxConcurrentRequestsOk returns a tuple with the MaxConcurrentRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConcurrentRequests

`func (o *EngineServingEndpointCreate) SetMaxConcurrentRequests(v int32)`

SetMaxConcurrentRequests sets MaxConcurrentRequests field to given value.

### HasMaxConcurrentRequests

`func (o *EngineServingEndpointCreate) HasMaxConcurrentRequests() bool`

HasMaxConcurrentRequests returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


