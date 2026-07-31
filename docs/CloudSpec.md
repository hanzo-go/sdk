# CloudSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Arch** | Pointer to **string** | amd64 | arm64 | ... | [optional] 
**Cpus** | Pointer to **int32** | logical cores | [optional] 
**Gpus** | Pointer to [**[]CloudGPU**](CloudGPU.md) |  | [optional] 
**Memory** | Pointer to **int32** | total RAM, bytes | [optional] 
**Os** | Pointer to **string** | linux | darwin | windows | [optional] 

## Methods

### NewCloudSpec

`func NewCloudSpec() *CloudSpec`

NewCloudSpec instantiates a new CloudSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudSpecWithDefaults

`func NewCloudSpecWithDefaults() *CloudSpec`

NewCloudSpecWithDefaults instantiates a new CloudSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArch

`func (o *CloudSpec) GetArch() string`

GetArch returns the Arch field if non-nil, zero value otherwise.

### GetArchOk

`func (o *CloudSpec) GetArchOk() (*string, bool)`

GetArchOk returns a tuple with the Arch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArch

`func (o *CloudSpec) SetArch(v string)`

SetArch sets Arch field to given value.

### HasArch

`func (o *CloudSpec) HasArch() bool`

HasArch returns a boolean if a field has been set.

### GetCpus

`func (o *CloudSpec) GetCpus() int32`

GetCpus returns the Cpus field if non-nil, zero value otherwise.

### GetCpusOk

`func (o *CloudSpec) GetCpusOk() (*int32, bool)`

GetCpusOk returns a tuple with the Cpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpus

`func (o *CloudSpec) SetCpus(v int32)`

SetCpus sets Cpus field to given value.

### HasCpus

`func (o *CloudSpec) HasCpus() bool`

HasCpus returns a boolean if a field has been set.

### GetGpus

`func (o *CloudSpec) GetGpus() []CloudGPU`

GetGpus returns the Gpus field if non-nil, zero value otherwise.

### GetGpusOk

`func (o *CloudSpec) GetGpusOk() (*[]CloudGPU, bool)`

GetGpusOk returns a tuple with the Gpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpus

`func (o *CloudSpec) SetGpus(v []CloudGPU)`

SetGpus sets Gpus field to given value.

### HasGpus

`func (o *CloudSpec) HasGpus() bool`

HasGpus returns a boolean if a field has been set.

### GetMemory

`func (o *CloudSpec) GetMemory() int32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *CloudSpec) GetMemoryOk() (*int32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *CloudSpec) SetMemory(v int32)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *CloudSpec) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetOs

`func (o *CloudSpec) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *CloudSpec) GetOsOk() (*string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *CloudSpec) SetOs(v string)`

SetOs sets Os field to given value.

### HasOs

`func (o *CloudSpec) HasOs() bool`

HasOs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


