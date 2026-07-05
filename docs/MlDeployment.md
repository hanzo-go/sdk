# MlDeployment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ModelId** | Pointer to **string** |  | [optional] 
**ModelVersion** | Pointer to **string** |  | [optional] 
**Runtime** | Pointer to **string** |  | [optional] 
**Gpu** | Pointer to **string** |  | [optional] 
**Replicas** | Pointer to **int32** |  | [optional] 
**Environment** | Pointer to **string** |  | [optional] 
**Endpoint** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewMlDeployment

`func NewMlDeployment() *MlDeployment`

NewMlDeployment instantiates a new MlDeployment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMlDeploymentWithDefaults

`func NewMlDeploymentWithDefaults() *MlDeployment`

NewMlDeploymentWithDefaults instantiates a new MlDeployment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MlDeployment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MlDeployment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MlDeployment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MlDeployment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModelId

`func (o *MlDeployment) GetModelId() string`

GetModelId returns the ModelId field if non-nil, zero value otherwise.

### GetModelIdOk

`func (o *MlDeployment) GetModelIdOk() (*string, bool)`

GetModelIdOk returns a tuple with the ModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelId

`func (o *MlDeployment) SetModelId(v string)`

SetModelId sets ModelId field to given value.

### HasModelId

`func (o *MlDeployment) HasModelId() bool`

HasModelId returns a boolean if a field has been set.

### GetModelVersion

`func (o *MlDeployment) GetModelVersion() string`

GetModelVersion returns the ModelVersion field if non-nil, zero value otherwise.

### GetModelVersionOk

`func (o *MlDeployment) GetModelVersionOk() (*string, bool)`

GetModelVersionOk returns a tuple with the ModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVersion

`func (o *MlDeployment) SetModelVersion(v string)`

SetModelVersion sets ModelVersion field to given value.

### HasModelVersion

`func (o *MlDeployment) HasModelVersion() bool`

HasModelVersion returns a boolean if a field has been set.

### GetRuntime

`func (o *MlDeployment) GetRuntime() string`

GetRuntime returns the Runtime field if non-nil, zero value otherwise.

### GetRuntimeOk

`func (o *MlDeployment) GetRuntimeOk() (*string, bool)`

GetRuntimeOk returns a tuple with the Runtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntime

`func (o *MlDeployment) SetRuntime(v string)`

SetRuntime sets Runtime field to given value.

### HasRuntime

`func (o *MlDeployment) HasRuntime() bool`

HasRuntime returns a boolean if a field has been set.

### GetGpu

`func (o *MlDeployment) GetGpu() string`

GetGpu returns the Gpu field if non-nil, zero value otherwise.

### GetGpuOk

`func (o *MlDeployment) GetGpuOk() (*string, bool)`

GetGpuOk returns a tuple with the Gpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpu

`func (o *MlDeployment) SetGpu(v string)`

SetGpu sets Gpu field to given value.

### HasGpu

`func (o *MlDeployment) HasGpu() bool`

HasGpu returns a boolean if a field has been set.

### GetReplicas

`func (o *MlDeployment) GetReplicas() int32`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *MlDeployment) GetReplicasOk() (*int32, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *MlDeployment) SetReplicas(v int32)`

SetReplicas sets Replicas field to given value.

### HasReplicas

`func (o *MlDeployment) HasReplicas() bool`

HasReplicas returns a boolean if a field has been set.

### GetEnvironment

`func (o *MlDeployment) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *MlDeployment) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *MlDeployment) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *MlDeployment) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetEndpoint

`func (o *MlDeployment) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *MlDeployment) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *MlDeployment) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *MlDeployment) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetStatus

`func (o *MlDeployment) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MlDeployment) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MlDeployment) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MlDeployment) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreatedAt

`func (o *MlDeployment) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MlDeployment) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MlDeployment) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *MlDeployment) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


