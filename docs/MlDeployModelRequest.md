# MlDeployModelRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ModelId** | **string** |  | 
**ModelVersion** | Pointer to **string** |  | [optional] 
**Runtime** | Pointer to **string** |  | [optional] [default to "vllm"]
**Gpu** | Pointer to **string** |  | [optional] [default to "a100"]
**Replicas** | Pointer to **int32** |  | [optional] [default to 1]
**Environment** | Pointer to **string** |  | [optional] [default to "dev"]

## Methods

### NewMlDeployModelRequest

`func NewMlDeployModelRequest(modelId string, ) *MlDeployModelRequest`

NewMlDeployModelRequest instantiates a new MlDeployModelRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMlDeployModelRequestWithDefaults

`func NewMlDeployModelRequestWithDefaults() *MlDeployModelRequest`

NewMlDeployModelRequestWithDefaults instantiates a new MlDeployModelRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModelId

`func (o *MlDeployModelRequest) GetModelId() string`

GetModelId returns the ModelId field if non-nil, zero value otherwise.

### GetModelIdOk

`func (o *MlDeployModelRequest) GetModelIdOk() (*string, bool)`

GetModelIdOk returns a tuple with the ModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelId

`func (o *MlDeployModelRequest) SetModelId(v string)`

SetModelId sets ModelId field to given value.


### GetModelVersion

`func (o *MlDeployModelRequest) GetModelVersion() string`

GetModelVersion returns the ModelVersion field if non-nil, zero value otherwise.

### GetModelVersionOk

`func (o *MlDeployModelRequest) GetModelVersionOk() (*string, bool)`

GetModelVersionOk returns a tuple with the ModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVersion

`func (o *MlDeployModelRequest) SetModelVersion(v string)`

SetModelVersion sets ModelVersion field to given value.

### HasModelVersion

`func (o *MlDeployModelRequest) HasModelVersion() bool`

HasModelVersion returns a boolean if a field has been set.

### GetRuntime

`func (o *MlDeployModelRequest) GetRuntime() string`

GetRuntime returns the Runtime field if non-nil, zero value otherwise.

### GetRuntimeOk

`func (o *MlDeployModelRequest) GetRuntimeOk() (*string, bool)`

GetRuntimeOk returns a tuple with the Runtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntime

`func (o *MlDeployModelRequest) SetRuntime(v string)`

SetRuntime sets Runtime field to given value.

### HasRuntime

`func (o *MlDeployModelRequest) HasRuntime() bool`

HasRuntime returns a boolean if a field has been set.

### GetGpu

`func (o *MlDeployModelRequest) GetGpu() string`

GetGpu returns the Gpu field if non-nil, zero value otherwise.

### GetGpuOk

`func (o *MlDeployModelRequest) GetGpuOk() (*string, bool)`

GetGpuOk returns a tuple with the Gpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpu

`func (o *MlDeployModelRequest) SetGpu(v string)`

SetGpu sets Gpu field to given value.

### HasGpu

`func (o *MlDeployModelRequest) HasGpu() bool`

HasGpu returns a boolean if a field has been set.

### GetReplicas

`func (o *MlDeployModelRequest) GetReplicas() int32`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *MlDeployModelRequest) GetReplicasOk() (*int32, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *MlDeployModelRequest) SetReplicas(v int32)`

SetReplicas sets Replicas field to given value.

### HasReplicas

`func (o *MlDeployModelRequest) HasReplicas() bool`

HasReplicas returns a boolean if a field has been set.

### GetEnvironment

`func (o *MlDeployModelRequest) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *MlDeployModelRequest) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *MlDeployModelRequest) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *MlDeployModelRequest) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


