# CloudByoGPU

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Arch** | Pointer to **string** | native target, e.g. \&quot;gfx1151\&quot; | [optional] 
**MemoryTotal** | Pointer to **string** | VRAM (or unified pool), e.g. \&quot;122880 MiB\&quot; | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Unified** | Pointer to **bool** | unified CPU/GPU memory pool (APU / SoC) | [optional] 

## Methods

### NewCloudByoGPU

`func NewCloudByoGPU() *CloudByoGPU`

NewCloudByoGPU instantiates a new CloudByoGPU object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudByoGPUWithDefaults

`func NewCloudByoGPUWithDefaults() *CloudByoGPU`

NewCloudByoGPUWithDefaults instantiates a new CloudByoGPU object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArch

`func (o *CloudByoGPU) GetArch() string`

GetArch returns the Arch field if non-nil, zero value otherwise.

### GetArchOk

`func (o *CloudByoGPU) GetArchOk() (*string, bool)`

GetArchOk returns a tuple with the Arch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArch

`func (o *CloudByoGPU) SetArch(v string)`

SetArch sets Arch field to given value.

### HasArch

`func (o *CloudByoGPU) HasArch() bool`

HasArch returns a boolean if a field has been set.

### GetMemoryTotal

`func (o *CloudByoGPU) GetMemoryTotal() string`

GetMemoryTotal returns the MemoryTotal field if non-nil, zero value otherwise.

### GetMemoryTotalOk

`func (o *CloudByoGPU) GetMemoryTotalOk() (*string, bool)`

GetMemoryTotalOk returns a tuple with the MemoryTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryTotal

`func (o *CloudByoGPU) SetMemoryTotal(v string)`

SetMemoryTotal sets MemoryTotal field to given value.

### HasMemoryTotal

`func (o *CloudByoGPU) HasMemoryTotal() bool`

HasMemoryTotal returns a boolean if a field has been set.

### GetName

`func (o *CloudByoGPU) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudByoGPU) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudByoGPU) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudByoGPU) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUnified

`func (o *CloudByoGPU) GetUnified() bool`

GetUnified returns the Unified field if non-nil, zero value otherwise.

### GetUnifiedOk

`func (o *CloudByoGPU) GetUnifiedOk() (*bool, bool)`

GetUnifiedOk returns a tuple with the Unified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnified

`func (o *CloudByoGPU) SetUnified(v bool)`

SetUnified sets Unified field to given value.

### HasUnified

`func (o *CloudByoGPU) HasUnified() bool`

HasUnified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


