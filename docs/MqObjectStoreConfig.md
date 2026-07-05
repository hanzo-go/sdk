# MqObjectStoreConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Store name. | 
**MaxChunkSize** | Pointer to **int32** | Maximum chunk size in bytes (default 128KB). | [optional] [default to 131072]
**MaxBytes** | Pointer to **int64** | Maximum total store size (-1 for unlimited). | [optional] [default to -1]
**Storage** | Pointer to **string** | Storage backend. | [optional] [default to "file"]
**NumReplicas** | Pointer to **int32** | Number of replicas. | [optional] [default to 1]
**Description** | Pointer to **string** | Optional human-readable description. | [optional] 

## Methods

### NewMqObjectStoreConfig

`func NewMqObjectStoreConfig(name string, ) *MqObjectStoreConfig`

NewMqObjectStoreConfig instantiates a new MqObjectStoreConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMqObjectStoreConfigWithDefaults

`func NewMqObjectStoreConfigWithDefaults() *MqObjectStoreConfig`

NewMqObjectStoreConfigWithDefaults instantiates a new MqObjectStoreConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MqObjectStoreConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MqObjectStoreConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MqObjectStoreConfig) SetName(v string)`

SetName sets Name field to given value.


### GetMaxChunkSize

`func (o *MqObjectStoreConfig) GetMaxChunkSize() int32`

GetMaxChunkSize returns the MaxChunkSize field if non-nil, zero value otherwise.

### GetMaxChunkSizeOk

`func (o *MqObjectStoreConfig) GetMaxChunkSizeOk() (*int32, bool)`

GetMaxChunkSizeOk returns a tuple with the MaxChunkSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxChunkSize

`func (o *MqObjectStoreConfig) SetMaxChunkSize(v int32)`

SetMaxChunkSize sets MaxChunkSize field to given value.

### HasMaxChunkSize

`func (o *MqObjectStoreConfig) HasMaxChunkSize() bool`

HasMaxChunkSize returns a boolean if a field has been set.

### GetMaxBytes

`func (o *MqObjectStoreConfig) GetMaxBytes() int64`

GetMaxBytes returns the MaxBytes field if non-nil, zero value otherwise.

### GetMaxBytesOk

`func (o *MqObjectStoreConfig) GetMaxBytesOk() (*int64, bool)`

GetMaxBytesOk returns a tuple with the MaxBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxBytes

`func (o *MqObjectStoreConfig) SetMaxBytes(v int64)`

SetMaxBytes sets MaxBytes field to given value.

### HasMaxBytes

`func (o *MqObjectStoreConfig) HasMaxBytes() bool`

HasMaxBytes returns a boolean if a field has been set.

### GetStorage

`func (o *MqObjectStoreConfig) GetStorage() string`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *MqObjectStoreConfig) GetStorageOk() (*string, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *MqObjectStoreConfig) SetStorage(v string)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *MqObjectStoreConfig) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### GetNumReplicas

`func (o *MqObjectStoreConfig) GetNumReplicas() int32`

GetNumReplicas returns the NumReplicas field if non-nil, zero value otherwise.

### GetNumReplicasOk

`func (o *MqObjectStoreConfig) GetNumReplicasOk() (*int32, bool)`

GetNumReplicasOk returns a tuple with the NumReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumReplicas

`func (o *MqObjectStoreConfig) SetNumReplicas(v int32)`

SetNumReplicas sets NumReplicas field to given value.

### HasNumReplicas

`func (o *MqObjectStoreConfig) HasNumReplicas() bool`

HasNumReplicas returns a boolean if a field has been set.

### GetDescription

`func (o *MqObjectStoreConfig) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MqObjectStoreConfig) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MqObjectStoreConfig) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MqObjectStoreConfig) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


