# KvUpdateClusterRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxMemoryMb** | Pointer to **int32** |  | [optional] 
**Replicas** | Pointer to **int32** |  | [optional] 
**EvictionPolicy** | Pointer to **string** |  | [optional] 

## Methods

### NewKvUpdateClusterRequest

`func NewKvUpdateClusterRequest() *KvUpdateClusterRequest`

NewKvUpdateClusterRequest instantiates a new KvUpdateClusterRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKvUpdateClusterRequestWithDefaults

`func NewKvUpdateClusterRequestWithDefaults() *KvUpdateClusterRequest`

NewKvUpdateClusterRequestWithDefaults instantiates a new KvUpdateClusterRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxMemoryMb

`func (o *KvUpdateClusterRequest) GetMaxMemoryMb() int32`

GetMaxMemoryMb returns the MaxMemoryMb field if non-nil, zero value otherwise.

### GetMaxMemoryMbOk

`func (o *KvUpdateClusterRequest) GetMaxMemoryMbOk() (*int32, bool)`

GetMaxMemoryMbOk returns a tuple with the MaxMemoryMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemoryMb

`func (o *KvUpdateClusterRequest) SetMaxMemoryMb(v int32)`

SetMaxMemoryMb sets MaxMemoryMb field to given value.

### HasMaxMemoryMb

`func (o *KvUpdateClusterRequest) HasMaxMemoryMb() bool`

HasMaxMemoryMb returns a boolean if a field has been set.

### GetReplicas

`func (o *KvUpdateClusterRequest) GetReplicas() int32`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *KvUpdateClusterRequest) GetReplicasOk() (*int32, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *KvUpdateClusterRequest) SetReplicas(v int32)`

SetReplicas sets Replicas field to given value.

### HasReplicas

`func (o *KvUpdateClusterRequest) HasReplicas() bool`

HasReplicas returns a boolean if a field has been set.

### GetEvictionPolicy

`func (o *KvUpdateClusterRequest) GetEvictionPolicy() string`

GetEvictionPolicy returns the EvictionPolicy field if non-nil, zero value otherwise.

### GetEvictionPolicyOk

`func (o *KvUpdateClusterRequest) GetEvictionPolicyOk() (*string, bool)`

GetEvictionPolicyOk returns a tuple with the EvictionPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvictionPolicy

`func (o *KvUpdateClusterRequest) SetEvictionPolicy(v string)`

SetEvictionPolicy sets EvictionPolicy field to given value.

### HasEvictionPolicy

`func (o *KvUpdateClusterRequest) HasEvictionPolicy() bool`

HasEvictionPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


