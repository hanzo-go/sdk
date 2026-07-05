# KvCreateNamespaceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**MaxMemoryMb** | Pointer to **int32** |  | [optional] [default to 64]
**EvictionPolicy** | Pointer to **string** |  | [optional] [default to "allkeys-lru"]

## Methods

### NewKvCreateNamespaceRequest

`func NewKvCreateNamespaceRequest(name string, ) *KvCreateNamespaceRequest`

NewKvCreateNamespaceRequest instantiates a new KvCreateNamespaceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKvCreateNamespaceRequestWithDefaults

`func NewKvCreateNamespaceRequestWithDefaults() *KvCreateNamespaceRequest`

NewKvCreateNamespaceRequestWithDefaults instantiates a new KvCreateNamespaceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *KvCreateNamespaceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KvCreateNamespaceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KvCreateNamespaceRequest) SetName(v string)`

SetName sets Name field to given value.


### GetMaxMemoryMb

`func (o *KvCreateNamespaceRequest) GetMaxMemoryMb() int32`

GetMaxMemoryMb returns the MaxMemoryMb field if non-nil, zero value otherwise.

### GetMaxMemoryMbOk

`func (o *KvCreateNamespaceRequest) GetMaxMemoryMbOk() (*int32, bool)`

GetMaxMemoryMbOk returns a tuple with the MaxMemoryMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemoryMb

`func (o *KvCreateNamespaceRequest) SetMaxMemoryMb(v int32)`

SetMaxMemoryMb sets MaxMemoryMb field to given value.

### HasMaxMemoryMb

`func (o *KvCreateNamespaceRequest) HasMaxMemoryMb() bool`

HasMaxMemoryMb returns a boolean if a field has been set.

### GetEvictionPolicy

`func (o *KvCreateNamespaceRequest) GetEvictionPolicy() string`

GetEvictionPolicy returns the EvictionPolicy field if non-nil, zero value otherwise.

### GetEvictionPolicyOk

`func (o *KvCreateNamespaceRequest) GetEvictionPolicyOk() (*string, bool)`

GetEvictionPolicyOk returns a tuple with the EvictionPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvictionPolicy

`func (o *KvCreateNamespaceRequest) SetEvictionPolicy(v string)`

SetEvictionPolicy sets EvictionPolicy field to given value.

### HasEvictionPolicy

`func (o *KvCreateNamespaceRequest) HasEvictionPolicy() bool`

HasEvictionPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


