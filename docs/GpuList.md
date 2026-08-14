# GpuList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gpus** | Pointer to [**[]GpuView**](GpuView.md) | GPUs is every accelerator the org has, from Visor GPU droplets and from BYO workers alike. | [optional] 

## Methods

### NewGpuList

`func NewGpuList() *GpuList`

NewGpuList instantiates a new GpuList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGpuListWithDefaults

`func NewGpuListWithDefaults() *GpuList`

NewGpuListWithDefaults instantiates a new GpuList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGpus

`func (o *GpuList) GetGpus() []GpuView`

GetGpus returns the Gpus field if non-nil, zero value otherwise.

### GetGpusOk

`func (o *GpuList) GetGpusOk() (*[]GpuView, bool)`

GetGpusOk returns a tuple with the Gpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpus

`func (o *GpuList) SetGpus(v []GpuView)`

SetGpus sets Gpus field to given value.

### HasGpus

`func (o *GpuList) HasGpus() bool`

HasGpus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


