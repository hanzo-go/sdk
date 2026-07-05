# EngineServingEndpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Model** | Pointer to **string** | Model identifier or path | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Replicas** | Pointer to **int32** |  | [optional] 
**MinReplicas** | Pointer to **int32** |  | [optional] 
**MaxReplicas** | Pointer to **int32** |  | [optional] 
**GpuType** | Pointer to **string** |  | [optional] 
**GpuPerReplica** | Pointer to **int32** |  | [optional] 
**Framework** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** | Inference endpoint URL | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewEngineServingEndpoint

`func NewEngineServingEndpoint() *EngineServingEndpoint`

NewEngineServingEndpoint instantiates a new EngineServingEndpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEngineServingEndpointWithDefaults

`func NewEngineServingEndpointWithDefaults() *EngineServingEndpoint`

NewEngineServingEndpointWithDefaults instantiates a new EngineServingEndpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EngineServingEndpoint) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EngineServingEndpoint) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EngineServingEndpoint) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EngineServingEndpoint) HasName() bool`

HasName returns a boolean if a field has been set.

### GetModel

`func (o *EngineServingEndpoint) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *EngineServingEndpoint) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *EngineServingEndpoint) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *EngineServingEndpoint) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetStatus

`func (o *EngineServingEndpoint) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EngineServingEndpoint) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EngineServingEndpoint) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EngineServingEndpoint) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetReplicas

`func (o *EngineServingEndpoint) GetReplicas() int32`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *EngineServingEndpoint) GetReplicasOk() (*int32, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *EngineServingEndpoint) SetReplicas(v int32)`

SetReplicas sets Replicas field to given value.

### HasReplicas

`func (o *EngineServingEndpoint) HasReplicas() bool`

HasReplicas returns a boolean if a field has been set.

### GetMinReplicas

`func (o *EngineServingEndpoint) GetMinReplicas() int32`

GetMinReplicas returns the MinReplicas field if non-nil, zero value otherwise.

### GetMinReplicasOk

`func (o *EngineServingEndpoint) GetMinReplicasOk() (*int32, bool)`

GetMinReplicasOk returns a tuple with the MinReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinReplicas

`func (o *EngineServingEndpoint) SetMinReplicas(v int32)`

SetMinReplicas sets MinReplicas field to given value.

### HasMinReplicas

`func (o *EngineServingEndpoint) HasMinReplicas() bool`

HasMinReplicas returns a boolean if a field has been set.

### GetMaxReplicas

`func (o *EngineServingEndpoint) GetMaxReplicas() int32`

GetMaxReplicas returns the MaxReplicas field if non-nil, zero value otherwise.

### GetMaxReplicasOk

`func (o *EngineServingEndpoint) GetMaxReplicasOk() (*int32, bool)`

GetMaxReplicasOk returns a tuple with the MaxReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxReplicas

`func (o *EngineServingEndpoint) SetMaxReplicas(v int32)`

SetMaxReplicas sets MaxReplicas field to given value.

### HasMaxReplicas

`func (o *EngineServingEndpoint) HasMaxReplicas() bool`

HasMaxReplicas returns a boolean if a field has been set.

### GetGpuType

`func (o *EngineServingEndpoint) GetGpuType() string`

GetGpuType returns the GpuType field if non-nil, zero value otherwise.

### GetGpuTypeOk

`func (o *EngineServingEndpoint) GetGpuTypeOk() (*string, bool)`

GetGpuTypeOk returns a tuple with the GpuType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuType

`func (o *EngineServingEndpoint) SetGpuType(v string)`

SetGpuType sets GpuType field to given value.

### HasGpuType

`func (o *EngineServingEndpoint) HasGpuType() bool`

HasGpuType returns a boolean if a field has been set.

### GetGpuPerReplica

`func (o *EngineServingEndpoint) GetGpuPerReplica() int32`

GetGpuPerReplica returns the GpuPerReplica field if non-nil, zero value otherwise.

### GetGpuPerReplicaOk

`func (o *EngineServingEndpoint) GetGpuPerReplicaOk() (*int32, bool)`

GetGpuPerReplicaOk returns a tuple with the GpuPerReplica field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuPerReplica

`func (o *EngineServingEndpoint) SetGpuPerReplica(v int32)`

SetGpuPerReplica sets GpuPerReplica field to given value.

### HasGpuPerReplica

`func (o *EngineServingEndpoint) HasGpuPerReplica() bool`

HasGpuPerReplica returns a boolean if a field has been set.

### GetFramework

`func (o *EngineServingEndpoint) GetFramework() string`

GetFramework returns the Framework field if non-nil, zero value otherwise.

### GetFrameworkOk

`func (o *EngineServingEndpoint) GetFrameworkOk() (*string, bool)`

GetFrameworkOk returns a tuple with the Framework field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFramework

`func (o *EngineServingEndpoint) SetFramework(v string)`

SetFramework sets Framework field to given value.

### HasFramework

`func (o *EngineServingEndpoint) HasFramework() bool`

HasFramework returns a boolean if a field has been set.

### GetUrl

`func (o *EngineServingEndpoint) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *EngineServingEndpoint) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *EngineServingEndpoint) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *EngineServingEndpoint) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetCreatedAt

`func (o *EngineServingEndpoint) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EngineServingEndpoint) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EngineServingEndpoint) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *EngineServingEndpoint) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *EngineServingEndpoint) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EngineServingEndpoint) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EngineServingEndpoint) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *EngineServingEndpoint) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


