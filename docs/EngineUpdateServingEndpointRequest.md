# EngineUpdateServingEndpointRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Model** | Pointer to **string** |  | [optional] 
**MinReplicas** | Pointer to **int32** |  | [optional] 
**MaxReplicas** | Pointer to **int32** |  | [optional] 
**GpuType** | Pointer to **string** |  | [optional] 
**GpuPerReplica** | Pointer to **int32** |  | [optional] 
**Env** | Pointer to **map[string]string** |  | [optional] 
**ScaleToZero** | Pointer to **bool** |  | [optional] 

## Methods

### NewEngineUpdateServingEndpointRequest

`func NewEngineUpdateServingEndpointRequest() *EngineUpdateServingEndpointRequest`

NewEngineUpdateServingEndpointRequest instantiates a new EngineUpdateServingEndpointRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEngineUpdateServingEndpointRequestWithDefaults

`func NewEngineUpdateServingEndpointRequestWithDefaults() *EngineUpdateServingEndpointRequest`

NewEngineUpdateServingEndpointRequestWithDefaults instantiates a new EngineUpdateServingEndpointRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModel

`func (o *EngineUpdateServingEndpointRequest) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *EngineUpdateServingEndpointRequest) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *EngineUpdateServingEndpointRequest) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *EngineUpdateServingEndpointRequest) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetMinReplicas

`func (o *EngineUpdateServingEndpointRequest) GetMinReplicas() int32`

GetMinReplicas returns the MinReplicas field if non-nil, zero value otherwise.

### GetMinReplicasOk

`func (o *EngineUpdateServingEndpointRequest) GetMinReplicasOk() (*int32, bool)`

GetMinReplicasOk returns a tuple with the MinReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinReplicas

`func (o *EngineUpdateServingEndpointRequest) SetMinReplicas(v int32)`

SetMinReplicas sets MinReplicas field to given value.

### HasMinReplicas

`func (o *EngineUpdateServingEndpointRequest) HasMinReplicas() bool`

HasMinReplicas returns a boolean if a field has been set.

### GetMaxReplicas

`func (o *EngineUpdateServingEndpointRequest) GetMaxReplicas() int32`

GetMaxReplicas returns the MaxReplicas field if non-nil, zero value otherwise.

### GetMaxReplicasOk

`func (o *EngineUpdateServingEndpointRequest) GetMaxReplicasOk() (*int32, bool)`

GetMaxReplicasOk returns a tuple with the MaxReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxReplicas

`func (o *EngineUpdateServingEndpointRequest) SetMaxReplicas(v int32)`

SetMaxReplicas sets MaxReplicas field to given value.

### HasMaxReplicas

`func (o *EngineUpdateServingEndpointRequest) HasMaxReplicas() bool`

HasMaxReplicas returns a boolean if a field has been set.

### GetGpuType

`func (o *EngineUpdateServingEndpointRequest) GetGpuType() string`

GetGpuType returns the GpuType field if non-nil, zero value otherwise.

### GetGpuTypeOk

`func (o *EngineUpdateServingEndpointRequest) GetGpuTypeOk() (*string, bool)`

GetGpuTypeOk returns a tuple with the GpuType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuType

`func (o *EngineUpdateServingEndpointRequest) SetGpuType(v string)`

SetGpuType sets GpuType field to given value.

### HasGpuType

`func (o *EngineUpdateServingEndpointRequest) HasGpuType() bool`

HasGpuType returns a boolean if a field has been set.

### GetGpuPerReplica

`func (o *EngineUpdateServingEndpointRequest) GetGpuPerReplica() int32`

GetGpuPerReplica returns the GpuPerReplica field if non-nil, zero value otherwise.

### GetGpuPerReplicaOk

`func (o *EngineUpdateServingEndpointRequest) GetGpuPerReplicaOk() (*int32, bool)`

GetGpuPerReplicaOk returns a tuple with the GpuPerReplica field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuPerReplica

`func (o *EngineUpdateServingEndpointRequest) SetGpuPerReplica(v int32)`

SetGpuPerReplica sets GpuPerReplica field to given value.

### HasGpuPerReplica

`func (o *EngineUpdateServingEndpointRequest) HasGpuPerReplica() bool`

HasGpuPerReplica returns a boolean if a field has been set.

### GetEnv

`func (o *EngineUpdateServingEndpointRequest) GetEnv() map[string]string`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *EngineUpdateServingEndpointRequest) GetEnvOk() (*map[string]string, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *EngineUpdateServingEndpointRequest) SetEnv(v map[string]string)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *EngineUpdateServingEndpointRequest) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### GetScaleToZero

`func (o *EngineUpdateServingEndpointRequest) GetScaleToZero() bool`

GetScaleToZero returns the ScaleToZero field if non-nil, zero value otherwise.

### GetScaleToZeroOk

`func (o *EngineUpdateServingEndpointRequest) GetScaleToZeroOk() (*bool, bool)`

GetScaleToZeroOk returns a tuple with the ScaleToZero field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaleToZero

`func (o *EngineUpdateServingEndpointRequest) SetScaleToZero(v bool)`

SetScaleToZero sets ScaleToZero field to given value.

### HasScaleToZero

`func (o *EngineUpdateServingEndpointRequest) HasScaleToZero() bool`

HasScaleToZero returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


