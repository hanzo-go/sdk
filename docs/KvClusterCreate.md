# KvClusterCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Mode** | Pointer to **string** |  | [optional] [default to "standalone"]
**Version** | Pointer to **string** |  | [optional] [default to "8.0"]
**MaxMemoryMb** | Pointer to **int32** |  | [optional] [default to 256]
**Replicas** | Pointer to **int32** |  | [optional] [default to 0]
**Tls** | Pointer to **bool** |  | [optional] [default to true]
**EvictionPolicy** | Pointer to **string** |  | [optional] [default to "allkeys-lru"]

## Methods

### NewKvClusterCreate

`func NewKvClusterCreate(name string, ) *KvClusterCreate`

NewKvClusterCreate instantiates a new KvClusterCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKvClusterCreateWithDefaults

`func NewKvClusterCreateWithDefaults() *KvClusterCreate`

NewKvClusterCreateWithDefaults instantiates a new KvClusterCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *KvClusterCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KvClusterCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KvClusterCreate) SetName(v string)`

SetName sets Name field to given value.


### GetMode

`func (o *KvClusterCreate) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *KvClusterCreate) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *KvClusterCreate) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *KvClusterCreate) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetVersion

`func (o *KvClusterCreate) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *KvClusterCreate) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *KvClusterCreate) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *KvClusterCreate) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetMaxMemoryMb

`func (o *KvClusterCreate) GetMaxMemoryMb() int32`

GetMaxMemoryMb returns the MaxMemoryMb field if non-nil, zero value otherwise.

### GetMaxMemoryMbOk

`func (o *KvClusterCreate) GetMaxMemoryMbOk() (*int32, bool)`

GetMaxMemoryMbOk returns a tuple with the MaxMemoryMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemoryMb

`func (o *KvClusterCreate) SetMaxMemoryMb(v int32)`

SetMaxMemoryMb sets MaxMemoryMb field to given value.

### HasMaxMemoryMb

`func (o *KvClusterCreate) HasMaxMemoryMb() bool`

HasMaxMemoryMb returns a boolean if a field has been set.

### GetReplicas

`func (o *KvClusterCreate) GetReplicas() int32`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *KvClusterCreate) GetReplicasOk() (*int32, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *KvClusterCreate) SetReplicas(v int32)`

SetReplicas sets Replicas field to given value.

### HasReplicas

`func (o *KvClusterCreate) HasReplicas() bool`

HasReplicas returns a boolean if a field has been set.

### GetTls

`func (o *KvClusterCreate) GetTls() bool`

GetTls returns the Tls field if non-nil, zero value otherwise.

### GetTlsOk

`func (o *KvClusterCreate) GetTlsOk() (*bool, bool)`

GetTlsOk returns a tuple with the Tls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTls

`func (o *KvClusterCreate) SetTls(v bool)`

SetTls sets Tls field to given value.

### HasTls

`func (o *KvClusterCreate) HasTls() bool`

HasTls returns a boolean if a field has been set.

### GetEvictionPolicy

`func (o *KvClusterCreate) GetEvictionPolicy() string`

GetEvictionPolicy returns the EvictionPolicy field if non-nil, zero value otherwise.

### GetEvictionPolicyOk

`func (o *KvClusterCreate) GetEvictionPolicyOk() (*string, bool)`

GetEvictionPolicyOk returns a tuple with the EvictionPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvictionPolicy

`func (o *KvClusterCreate) SetEvictionPolicy(v string)`

SetEvictionPolicy sets EvictionPolicy field to given value.

### HasEvictionPolicy

`func (o *KvClusterCreate) HasEvictionPolicy() bool`

HasEvictionPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


