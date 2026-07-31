# CloudGPU

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Memory** | Pointer to **int32** | VRAM bytes, 0 &#x3D; unknown | [optional] 
**Model** | Pointer to **string** | \&quot;GB10\&quot;, \&quot;8060S\&quot;, \&quot;RTX 4090\&quot; | [optional] 
**Vendor** | Pointer to **string** | nvidia | amd | apple | intel | ... | [optional] 

## Methods

### NewCloudGPU

`func NewCloudGPU() *CloudGPU`

NewCloudGPU instantiates a new CloudGPU object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudGPUWithDefaults

`func NewCloudGPUWithDefaults() *CloudGPU`

NewCloudGPUWithDefaults instantiates a new CloudGPU object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMemory

`func (o *CloudGPU) GetMemory() int32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *CloudGPU) GetMemoryOk() (*int32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *CloudGPU) SetMemory(v int32)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *CloudGPU) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetModel

`func (o *CloudGPU) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *CloudGPU) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *CloudGPU) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *CloudGPU) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetVendor

`func (o *CloudGPU) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *CloudGPU) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *CloudGPU) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *CloudGPU) HasVendor() bool`

HasVendor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


