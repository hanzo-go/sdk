# KvNamespace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**MaxMemoryMb** | Pointer to **int32** | Memory limit in MB | [optional] 
**UsedMemoryMb** | Pointer to **float32** |  | [optional] 
**KeyCount** | Pointer to **int32** |  | [optional] 
**EvictionPolicy** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewKvNamespace

`func NewKvNamespace() *KvNamespace`

NewKvNamespace instantiates a new KvNamespace object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKvNamespaceWithDefaults

`func NewKvNamespaceWithDefaults() *KvNamespace`

NewKvNamespaceWithDefaults instantiates a new KvNamespace object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KvNamespace) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KvNamespace) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KvNamespace) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *KvNamespace) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *KvNamespace) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KvNamespace) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KvNamespace) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KvNamespace) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMaxMemoryMb

`func (o *KvNamespace) GetMaxMemoryMb() int32`

GetMaxMemoryMb returns the MaxMemoryMb field if non-nil, zero value otherwise.

### GetMaxMemoryMbOk

`func (o *KvNamespace) GetMaxMemoryMbOk() (*int32, bool)`

GetMaxMemoryMbOk returns a tuple with the MaxMemoryMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemoryMb

`func (o *KvNamespace) SetMaxMemoryMb(v int32)`

SetMaxMemoryMb sets MaxMemoryMb field to given value.

### HasMaxMemoryMb

`func (o *KvNamespace) HasMaxMemoryMb() bool`

HasMaxMemoryMb returns a boolean if a field has been set.

### GetUsedMemoryMb

`func (o *KvNamespace) GetUsedMemoryMb() float32`

GetUsedMemoryMb returns the UsedMemoryMb field if non-nil, zero value otherwise.

### GetUsedMemoryMbOk

`func (o *KvNamespace) GetUsedMemoryMbOk() (*float32, bool)`

GetUsedMemoryMbOk returns a tuple with the UsedMemoryMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedMemoryMb

`func (o *KvNamespace) SetUsedMemoryMb(v float32)`

SetUsedMemoryMb sets UsedMemoryMb field to given value.

### HasUsedMemoryMb

`func (o *KvNamespace) HasUsedMemoryMb() bool`

HasUsedMemoryMb returns a boolean if a field has been set.

### GetKeyCount

`func (o *KvNamespace) GetKeyCount() int32`

GetKeyCount returns the KeyCount field if non-nil, zero value otherwise.

### GetKeyCountOk

`func (o *KvNamespace) GetKeyCountOk() (*int32, bool)`

GetKeyCountOk returns a tuple with the KeyCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyCount

`func (o *KvNamespace) SetKeyCount(v int32)`

SetKeyCount sets KeyCount field to given value.

### HasKeyCount

`func (o *KvNamespace) HasKeyCount() bool`

HasKeyCount returns a boolean if a field has been set.

### GetEvictionPolicy

`func (o *KvNamespace) GetEvictionPolicy() string`

GetEvictionPolicy returns the EvictionPolicy field if non-nil, zero value otherwise.

### GetEvictionPolicyOk

`func (o *KvNamespace) GetEvictionPolicyOk() (*string, bool)`

GetEvictionPolicyOk returns a tuple with the EvictionPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvictionPolicy

`func (o *KvNamespace) SetEvictionPolicy(v string)`

SetEvictionPolicy sets EvictionPolicy field to given value.

### HasEvictionPolicy

`func (o *KvNamespace) HasEvictionPolicy() bool`

HasEvictionPolicy returns a boolean if a field has been set.

### GetCreatedAt

`func (o *KvNamespace) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KvNamespace) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KvNamespace) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *KvNamespace) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


