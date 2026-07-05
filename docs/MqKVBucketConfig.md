# MqKVBucketConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Bucket name. | 
**History** | Pointer to **int32** | Maximum number of revisions per key. | [optional] [default to 1]
**Ttl** | Pointer to **string** | Default TTL for keys (e.g., \&quot;1h\&quot;, \&quot;7d\&quot;, \&quot;0\&quot; for no expiry).  | [optional] [default to "0"]
**MaxValueSize** | Pointer to **int32** | Maximum value size in bytes (-1 for default). | [optional] [default to -1]
**MaxBytes** | Pointer to **int64** | Maximum total bucket size (-1 for unlimited). | [optional] [default to -1]
**Storage** | Pointer to **string** | Storage backend. | [optional] [default to "file"]
**NumReplicas** | Pointer to **int32** | Number of replicas. | [optional] [default to 1]
**Description** | Pointer to **string** | Optional human-readable description. | [optional] 

## Methods

### NewMqKVBucketConfig

`func NewMqKVBucketConfig(name string, ) *MqKVBucketConfig`

NewMqKVBucketConfig instantiates a new MqKVBucketConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMqKVBucketConfigWithDefaults

`func NewMqKVBucketConfigWithDefaults() *MqKVBucketConfig`

NewMqKVBucketConfigWithDefaults instantiates a new MqKVBucketConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MqKVBucketConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MqKVBucketConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MqKVBucketConfig) SetName(v string)`

SetName sets Name field to given value.


### GetHistory

`func (o *MqKVBucketConfig) GetHistory() int32`

GetHistory returns the History field if non-nil, zero value otherwise.

### GetHistoryOk

`func (o *MqKVBucketConfig) GetHistoryOk() (*int32, bool)`

GetHistoryOk returns a tuple with the History field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistory

`func (o *MqKVBucketConfig) SetHistory(v int32)`

SetHistory sets History field to given value.

### HasHistory

`func (o *MqKVBucketConfig) HasHistory() bool`

HasHistory returns a boolean if a field has been set.

### GetTtl

`func (o *MqKVBucketConfig) GetTtl() string`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *MqKVBucketConfig) GetTtlOk() (*string, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *MqKVBucketConfig) SetTtl(v string)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *MqKVBucketConfig) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetMaxValueSize

`func (o *MqKVBucketConfig) GetMaxValueSize() int32`

GetMaxValueSize returns the MaxValueSize field if non-nil, zero value otherwise.

### GetMaxValueSizeOk

`func (o *MqKVBucketConfig) GetMaxValueSizeOk() (*int32, bool)`

GetMaxValueSizeOk returns a tuple with the MaxValueSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxValueSize

`func (o *MqKVBucketConfig) SetMaxValueSize(v int32)`

SetMaxValueSize sets MaxValueSize field to given value.

### HasMaxValueSize

`func (o *MqKVBucketConfig) HasMaxValueSize() bool`

HasMaxValueSize returns a boolean if a field has been set.

### GetMaxBytes

`func (o *MqKVBucketConfig) GetMaxBytes() int64`

GetMaxBytes returns the MaxBytes field if non-nil, zero value otherwise.

### GetMaxBytesOk

`func (o *MqKVBucketConfig) GetMaxBytesOk() (*int64, bool)`

GetMaxBytesOk returns a tuple with the MaxBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxBytes

`func (o *MqKVBucketConfig) SetMaxBytes(v int64)`

SetMaxBytes sets MaxBytes field to given value.

### HasMaxBytes

`func (o *MqKVBucketConfig) HasMaxBytes() bool`

HasMaxBytes returns a boolean if a field has been set.

### GetStorage

`func (o *MqKVBucketConfig) GetStorage() string`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *MqKVBucketConfig) GetStorageOk() (*string, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *MqKVBucketConfig) SetStorage(v string)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *MqKVBucketConfig) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### GetNumReplicas

`func (o *MqKVBucketConfig) GetNumReplicas() int32`

GetNumReplicas returns the NumReplicas field if non-nil, zero value otherwise.

### GetNumReplicasOk

`func (o *MqKVBucketConfig) GetNumReplicasOk() (*int32, bool)`

GetNumReplicasOk returns a tuple with the NumReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumReplicas

`func (o *MqKVBucketConfig) SetNumReplicas(v int32)`

SetNumReplicas sets NumReplicas field to given value.

### HasNumReplicas

`func (o *MqKVBucketConfig) HasNumReplicas() bool`

HasNumReplicas returns a boolean if a field has been set.

### GetDescription

`func (o *MqKVBucketConfig) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MqKVBucketConfig) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MqKVBucketConfig) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MqKVBucketConfig) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


