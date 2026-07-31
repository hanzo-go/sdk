# VectorCreateCollectionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Vectors** | Pointer to [**VectorCollectionConfigParamsVectors**](VectorCollectionConfigParamsVectors.md) |  | [optional] 
**ShardNumber** | Pointer to **int32** |  | [optional] 
**ReplicationFactor** | Pointer to **int32** |  | [optional] 
**WriteConsistencyFactor** | Pointer to **int32** |  | [optional] 
**OnDiskPayload** | Pointer to **bool** |  | [optional] 
**HnswConfig** | Pointer to [**VectorHnswConfig**](VectorHnswConfig.md) |  | [optional] 
**OptimizersConfig** | Pointer to [**VectorOptimizerConfig**](VectorOptimizerConfig.md) |  | [optional] 
**WalConfig** | Pointer to [**VectorWalConfig**](VectorWalConfig.md) |  | [optional] 

## Methods

### NewVectorCreateCollectionRequest

`func NewVectorCreateCollectionRequest() *VectorCreateCollectionRequest`

NewVectorCreateCollectionRequest instantiates a new VectorCreateCollectionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVectorCreateCollectionRequestWithDefaults

`func NewVectorCreateCollectionRequestWithDefaults() *VectorCreateCollectionRequest`

NewVectorCreateCollectionRequestWithDefaults instantiates a new VectorCreateCollectionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVectors

`func (o *VectorCreateCollectionRequest) GetVectors() VectorCollectionConfigParamsVectors`

GetVectors returns the Vectors field if non-nil, zero value otherwise.

### GetVectorsOk

`func (o *VectorCreateCollectionRequest) GetVectorsOk() (*VectorCollectionConfigParamsVectors, bool)`

GetVectorsOk returns a tuple with the Vectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVectors

`func (o *VectorCreateCollectionRequest) SetVectors(v VectorCollectionConfigParamsVectors)`

SetVectors sets Vectors field to given value.

### HasVectors

`func (o *VectorCreateCollectionRequest) HasVectors() bool`

HasVectors returns a boolean if a field has been set.

### GetShardNumber

`func (o *VectorCreateCollectionRequest) GetShardNumber() int32`

GetShardNumber returns the ShardNumber field if non-nil, zero value otherwise.

### GetShardNumberOk

`func (o *VectorCreateCollectionRequest) GetShardNumberOk() (*int32, bool)`

GetShardNumberOk returns a tuple with the ShardNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShardNumber

`func (o *VectorCreateCollectionRequest) SetShardNumber(v int32)`

SetShardNumber sets ShardNumber field to given value.

### HasShardNumber

`func (o *VectorCreateCollectionRequest) HasShardNumber() bool`

HasShardNumber returns a boolean if a field has been set.

### GetReplicationFactor

`func (o *VectorCreateCollectionRequest) GetReplicationFactor() int32`

GetReplicationFactor returns the ReplicationFactor field if non-nil, zero value otherwise.

### GetReplicationFactorOk

`func (o *VectorCreateCollectionRequest) GetReplicationFactorOk() (*int32, bool)`

GetReplicationFactorOk returns a tuple with the ReplicationFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationFactor

`func (o *VectorCreateCollectionRequest) SetReplicationFactor(v int32)`

SetReplicationFactor sets ReplicationFactor field to given value.

### HasReplicationFactor

`func (o *VectorCreateCollectionRequest) HasReplicationFactor() bool`

HasReplicationFactor returns a boolean if a field has been set.

### GetWriteConsistencyFactor

`func (o *VectorCreateCollectionRequest) GetWriteConsistencyFactor() int32`

GetWriteConsistencyFactor returns the WriteConsistencyFactor field if non-nil, zero value otherwise.

### GetWriteConsistencyFactorOk

`func (o *VectorCreateCollectionRequest) GetWriteConsistencyFactorOk() (*int32, bool)`

GetWriteConsistencyFactorOk returns a tuple with the WriteConsistencyFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteConsistencyFactor

`func (o *VectorCreateCollectionRequest) SetWriteConsistencyFactor(v int32)`

SetWriteConsistencyFactor sets WriteConsistencyFactor field to given value.

### HasWriteConsistencyFactor

`func (o *VectorCreateCollectionRequest) HasWriteConsistencyFactor() bool`

HasWriteConsistencyFactor returns a boolean if a field has been set.

### GetOnDiskPayload

`func (o *VectorCreateCollectionRequest) GetOnDiskPayload() bool`

GetOnDiskPayload returns the OnDiskPayload field if non-nil, zero value otherwise.

### GetOnDiskPayloadOk

`func (o *VectorCreateCollectionRequest) GetOnDiskPayloadOk() (*bool, bool)`

GetOnDiskPayloadOk returns a tuple with the OnDiskPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnDiskPayload

`func (o *VectorCreateCollectionRequest) SetOnDiskPayload(v bool)`

SetOnDiskPayload sets OnDiskPayload field to given value.

### HasOnDiskPayload

`func (o *VectorCreateCollectionRequest) HasOnDiskPayload() bool`

HasOnDiskPayload returns a boolean if a field has been set.

### GetHnswConfig

`func (o *VectorCreateCollectionRequest) GetHnswConfig() VectorHnswConfig`

GetHnswConfig returns the HnswConfig field if non-nil, zero value otherwise.

### GetHnswConfigOk

`func (o *VectorCreateCollectionRequest) GetHnswConfigOk() (*VectorHnswConfig, bool)`

GetHnswConfigOk returns a tuple with the HnswConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHnswConfig

`func (o *VectorCreateCollectionRequest) SetHnswConfig(v VectorHnswConfig)`

SetHnswConfig sets HnswConfig field to given value.

### HasHnswConfig

`func (o *VectorCreateCollectionRequest) HasHnswConfig() bool`

HasHnswConfig returns a boolean if a field has been set.

### GetOptimizersConfig

`func (o *VectorCreateCollectionRequest) GetOptimizersConfig() VectorOptimizerConfig`

GetOptimizersConfig returns the OptimizersConfig field if non-nil, zero value otherwise.

### GetOptimizersConfigOk

`func (o *VectorCreateCollectionRequest) GetOptimizersConfigOk() (*VectorOptimizerConfig, bool)`

GetOptimizersConfigOk returns a tuple with the OptimizersConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptimizersConfig

`func (o *VectorCreateCollectionRequest) SetOptimizersConfig(v VectorOptimizerConfig)`

SetOptimizersConfig sets OptimizersConfig field to given value.

### HasOptimizersConfig

`func (o *VectorCreateCollectionRequest) HasOptimizersConfig() bool`

HasOptimizersConfig returns a boolean if a field has been set.

### GetWalConfig

`func (o *VectorCreateCollectionRequest) GetWalConfig() VectorWalConfig`

GetWalConfig returns the WalConfig field if non-nil, zero value otherwise.

### GetWalConfigOk

`func (o *VectorCreateCollectionRequest) GetWalConfigOk() (*VectorWalConfig, bool)`

GetWalConfigOk returns a tuple with the WalConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalConfig

`func (o *VectorCreateCollectionRequest) SetWalConfig(v VectorWalConfig)`

SetWalConfig sets WalConfig field to given value.

### HasWalConfig

`func (o *VectorCreateCollectionRequest) HasWalConfig() bool`

HasWalConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


