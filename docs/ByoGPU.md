# ByoGPU

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Arch** | Pointer to **string** | native target, e.g. \&quot;gfx1151\&quot; | [optional] 
**MemoryTotal** | Pointer to **string** | VRAM (or unified pool), e.g. \&quot;122880 MiB\&quot; | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Unified** | Pointer to **bool** | unified CPU/GPU memory pool (APU / SoC) | [optional] 

## Methods

### NewByoGPU

`func NewByoGPU() *ByoGPU`

NewByoGPU instantiates a new ByoGPU object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewByoGPUWithDefaults

`func NewByoGPUWithDefaults() *ByoGPU`

NewByoGPUWithDefaults instantiates a new ByoGPU object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArch

`func (o *ByoGPU) GetArch() string`

GetArch returns the Arch field if non-nil, zero value otherwise.

### GetArchOk

`func (o *ByoGPU) GetArchOk() (*string, bool)`

GetArchOk returns a tuple with the Arch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArch

`func (o *ByoGPU) SetArch(v string)`

SetArch sets Arch field to given value.

### HasArch

`func (o *ByoGPU) HasArch() bool`

HasArch returns a boolean if a field has been set.

### GetMemoryTotal

`func (o *ByoGPU) GetMemoryTotal() string`

GetMemoryTotal returns the MemoryTotal field if non-nil, zero value otherwise.

### GetMemoryTotalOk

`func (o *ByoGPU) GetMemoryTotalOk() (*string, bool)`

GetMemoryTotalOk returns a tuple with the MemoryTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryTotal

`func (o *ByoGPU) SetMemoryTotal(v string)`

SetMemoryTotal sets MemoryTotal field to given value.

### HasMemoryTotal

`func (o *ByoGPU) HasMemoryTotal() bool`

HasMemoryTotal returns a boolean if a field has been set.

### GetName

`func (o *ByoGPU) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ByoGPU) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ByoGPU) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ByoGPU) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUnified

`func (o *ByoGPU) GetUnified() bool`

GetUnified returns the Unified field if non-nil, zero value otherwise.

### GetUnifiedOk

`func (o *ByoGPU) GetUnifiedOk() (*bool, bool)`

GetUnifiedOk returns a tuple with the Unified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnified

`func (o *ByoGPU) SetUnified(v bool)`

SetUnified sets Unified field to given value.

### HasUnified

`func (o *ByoGPU) HasUnified() bool`

HasUnified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


