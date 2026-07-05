# VectorCollectionConfigParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Vectors** | Pointer to [**VectorCreateCollectionRequestVectors**](VectorCreateCollectionRequestVectors.md) |  | [optional] 
**ShardNumber** | Pointer to **int32** |  | [optional] 
**ReplicationFactor** | Pointer to **int32** |  | [optional] 

## Methods

### NewVectorCollectionConfigParams

`func NewVectorCollectionConfigParams() *VectorCollectionConfigParams`

NewVectorCollectionConfigParams instantiates a new VectorCollectionConfigParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVectorCollectionConfigParamsWithDefaults

`func NewVectorCollectionConfigParamsWithDefaults() *VectorCollectionConfigParams`

NewVectorCollectionConfigParamsWithDefaults instantiates a new VectorCollectionConfigParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVectors

`func (o *VectorCollectionConfigParams) GetVectors() VectorCreateCollectionRequestVectors`

GetVectors returns the Vectors field if non-nil, zero value otherwise.

### GetVectorsOk

`func (o *VectorCollectionConfigParams) GetVectorsOk() (*VectorCreateCollectionRequestVectors, bool)`

GetVectorsOk returns a tuple with the Vectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVectors

`func (o *VectorCollectionConfigParams) SetVectors(v VectorCreateCollectionRequestVectors)`

SetVectors sets Vectors field to given value.

### HasVectors

`func (o *VectorCollectionConfigParams) HasVectors() bool`

HasVectors returns a boolean if a field has been set.

### GetShardNumber

`func (o *VectorCollectionConfigParams) GetShardNumber() int32`

GetShardNumber returns the ShardNumber field if non-nil, zero value otherwise.

### GetShardNumberOk

`func (o *VectorCollectionConfigParams) GetShardNumberOk() (*int32, bool)`

GetShardNumberOk returns a tuple with the ShardNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShardNumber

`func (o *VectorCollectionConfigParams) SetShardNumber(v int32)`

SetShardNumber sets ShardNumber field to given value.

### HasShardNumber

`func (o *VectorCollectionConfigParams) HasShardNumber() bool`

HasShardNumber returns a boolean if a field has been set.

### GetReplicationFactor

`func (o *VectorCollectionConfigParams) GetReplicationFactor() int32`

GetReplicationFactor returns the ReplicationFactor field if non-nil, zero value otherwise.

### GetReplicationFactorOk

`func (o *VectorCollectionConfigParams) GetReplicationFactorOk() (*int32, bool)`

GetReplicationFactorOk returns a tuple with the ReplicationFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationFactor

`func (o *VectorCollectionConfigParams) SetReplicationFactor(v int32)`

SetReplicationFactor sets ReplicationFactor field to given value.

### HasReplicationFactor

`func (o *VectorCollectionConfigParams) HasReplicationFactor() bool`

HasReplicationFactor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


