# KvCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** | Server version | [optional] 
**Mode** | Pointer to **string** |  | [optional] 
**Nodes** | Pointer to [**[]KvClusterNodesInner**](KvClusterNodesInner.md) |  | [optional] 
**Replicas** | Pointer to **int32** |  | [optional] 
**MaxMemoryMb** | Pointer to **int32** |  | [optional] 
**ConnectionUri** | Pointer to **string** | Connection string (valkey://...) | [optional] 
**Tls** | Pointer to **bool** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewKvCluster

`func NewKvCluster() *KvCluster`

NewKvCluster instantiates a new KvCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKvClusterWithDefaults

`func NewKvClusterWithDefaults() *KvCluster`

NewKvClusterWithDefaults instantiates a new KvCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KvCluster) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KvCluster) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KvCluster) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *KvCluster) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *KvCluster) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KvCluster) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KvCluster) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KvCluster) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *KvCluster) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KvCluster) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KvCluster) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *KvCluster) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVersion

`func (o *KvCluster) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *KvCluster) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *KvCluster) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *KvCluster) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetMode

`func (o *KvCluster) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *KvCluster) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *KvCluster) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *KvCluster) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetNodes

`func (o *KvCluster) GetNodes() []KvClusterNodesInner`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *KvCluster) GetNodesOk() (*[]KvClusterNodesInner, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *KvCluster) SetNodes(v []KvClusterNodesInner)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *KvCluster) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetReplicas

`func (o *KvCluster) GetReplicas() int32`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *KvCluster) GetReplicasOk() (*int32, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *KvCluster) SetReplicas(v int32)`

SetReplicas sets Replicas field to given value.

### HasReplicas

`func (o *KvCluster) HasReplicas() bool`

HasReplicas returns a boolean if a field has been set.

### GetMaxMemoryMb

`func (o *KvCluster) GetMaxMemoryMb() int32`

GetMaxMemoryMb returns the MaxMemoryMb field if non-nil, zero value otherwise.

### GetMaxMemoryMbOk

`func (o *KvCluster) GetMaxMemoryMbOk() (*int32, bool)`

GetMaxMemoryMbOk returns a tuple with the MaxMemoryMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemoryMb

`func (o *KvCluster) SetMaxMemoryMb(v int32)`

SetMaxMemoryMb sets MaxMemoryMb field to given value.

### HasMaxMemoryMb

`func (o *KvCluster) HasMaxMemoryMb() bool`

HasMaxMemoryMb returns a boolean if a field has been set.

### GetConnectionUri

`func (o *KvCluster) GetConnectionUri() string`

GetConnectionUri returns the ConnectionUri field if non-nil, zero value otherwise.

### GetConnectionUriOk

`func (o *KvCluster) GetConnectionUriOk() (*string, bool)`

GetConnectionUriOk returns a tuple with the ConnectionUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionUri

`func (o *KvCluster) SetConnectionUri(v string)`

SetConnectionUri sets ConnectionUri field to given value.

### HasConnectionUri

`func (o *KvCluster) HasConnectionUri() bool`

HasConnectionUri returns a boolean if a field has been set.

### GetTls

`func (o *KvCluster) GetTls() bool`

GetTls returns the Tls field if non-nil, zero value otherwise.

### GetTlsOk

`func (o *KvCluster) GetTlsOk() (*bool, bool)`

GetTlsOk returns a tuple with the Tls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTls

`func (o *KvCluster) SetTls(v bool)`

SetTls sets Tls field to given value.

### HasTls

`func (o *KvCluster) HasTls() bool`

HasTls returns a boolean if a field has been set.

### GetCreatedAt

`func (o *KvCluster) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KvCluster) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KvCluster) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *KvCluster) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


