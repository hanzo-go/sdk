# GPU

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Memory** | Pointer to **int32** | VRAM bytes, 0 &#x3D; unknown | [optional] 
**Model** | Pointer to **string** | \&quot;GB10\&quot;, \&quot;8060S\&quot;, \&quot;RTX 4090\&quot; | [optional] 
**Vendor** | Pointer to **string** | nvidia | amd | apple | intel | ... | [optional] 

## Methods

### NewGPU

`func NewGPU() *GPU`

NewGPU instantiates a new GPU object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGPUWithDefaults

`func NewGPUWithDefaults() *GPU`

NewGPUWithDefaults instantiates a new GPU object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMemory

`func (o *GPU) GetMemory() int32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *GPU) GetMemoryOk() (*int32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *GPU) SetMemory(v int32)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *GPU) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetModel

`func (o *GPU) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *GPU) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *GPU) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *GPU) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetVendor

`func (o *GPU) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *GPU) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *GPU) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *GPU) HasVendor() bool`

HasVendor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


