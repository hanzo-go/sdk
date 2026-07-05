# PubsubKVBucketConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bucket** | **string** | Bucket name | 
**History** | Pointer to **int32** | Number of historical values per key | [optional] [default to 1]
**Ttl** | Pointer to **int32** | Default TTL in nanoseconds (0 &#x3D; no expiry) | [optional] [default to 0]
**Replicas** | Pointer to **int32** |  | [optional] [default to 1]
**MaxValueSize** | Pointer to **int32** | Maximum value size in bytes | [optional] [default to -1]

## Methods

### NewPubsubKVBucketConfig

`func NewPubsubKVBucketConfig(bucket string, ) *PubsubKVBucketConfig`

NewPubsubKVBucketConfig instantiates a new PubsubKVBucketConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPubsubKVBucketConfigWithDefaults

`func NewPubsubKVBucketConfigWithDefaults() *PubsubKVBucketConfig`

NewPubsubKVBucketConfigWithDefaults instantiates a new PubsubKVBucketConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucket

`func (o *PubsubKVBucketConfig) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *PubsubKVBucketConfig) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *PubsubKVBucketConfig) SetBucket(v string)`

SetBucket sets Bucket field to given value.


### GetHistory

`func (o *PubsubKVBucketConfig) GetHistory() int32`

GetHistory returns the History field if non-nil, zero value otherwise.

### GetHistoryOk

`func (o *PubsubKVBucketConfig) GetHistoryOk() (*int32, bool)`

GetHistoryOk returns a tuple with the History field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistory

`func (o *PubsubKVBucketConfig) SetHistory(v int32)`

SetHistory sets History field to given value.

### HasHistory

`func (o *PubsubKVBucketConfig) HasHistory() bool`

HasHistory returns a boolean if a field has been set.

### GetTtl

`func (o *PubsubKVBucketConfig) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *PubsubKVBucketConfig) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *PubsubKVBucketConfig) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *PubsubKVBucketConfig) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetReplicas

`func (o *PubsubKVBucketConfig) GetReplicas() int32`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *PubsubKVBucketConfig) GetReplicasOk() (*int32, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *PubsubKVBucketConfig) SetReplicas(v int32)`

SetReplicas sets Replicas field to given value.

### HasReplicas

`func (o *PubsubKVBucketConfig) HasReplicas() bool`

HasReplicas returns a boolean if a field has been set.

### GetMaxValueSize

`func (o *PubsubKVBucketConfig) GetMaxValueSize() int32`

GetMaxValueSize returns the MaxValueSize field if non-nil, zero value otherwise.

### GetMaxValueSizeOk

`func (o *PubsubKVBucketConfig) GetMaxValueSizeOk() (*int32, bool)`

GetMaxValueSizeOk returns a tuple with the MaxValueSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxValueSize

`func (o *PubsubKVBucketConfig) SetMaxValueSize(v int32)`

SetMaxValueSize sets MaxValueSize field to given value.

### HasMaxValueSize

`func (o *PubsubKVBucketConfig) HasMaxValueSize() bool`

HasMaxValueSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


