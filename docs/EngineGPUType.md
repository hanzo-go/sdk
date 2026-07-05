# EngineGPUType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Model** | Pointer to **string** | GPU model name | [optional] 
**MemoryGb** | Pointer to **int32** |  | [optional] 
**Architecture** | Pointer to **string** |  | [optional] 
**Fp16Tflops** | Pointer to **float32** |  | [optional] 
**Bf16Tflops** | Pointer to **float32** |  | [optional] 
**Fp8Tflops** | Pointer to **float32** |  | [optional] 
**Interconnect** | Pointer to **string** |  | [optional] 
**TotalCount** | Pointer to **int32** |  | [optional] 
**AvailableCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewEngineGPUType

`func NewEngineGPUType() *EngineGPUType`

NewEngineGPUType instantiates a new EngineGPUType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEngineGPUTypeWithDefaults

`func NewEngineGPUTypeWithDefaults() *EngineGPUType`

NewEngineGPUTypeWithDefaults instantiates a new EngineGPUType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModel

`func (o *EngineGPUType) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *EngineGPUType) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *EngineGPUType) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *EngineGPUType) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetMemoryGb

`func (o *EngineGPUType) GetMemoryGb() int32`

GetMemoryGb returns the MemoryGb field if non-nil, zero value otherwise.

### GetMemoryGbOk

`func (o *EngineGPUType) GetMemoryGbOk() (*int32, bool)`

GetMemoryGbOk returns a tuple with the MemoryGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryGb

`func (o *EngineGPUType) SetMemoryGb(v int32)`

SetMemoryGb sets MemoryGb field to given value.

### HasMemoryGb

`func (o *EngineGPUType) HasMemoryGb() bool`

HasMemoryGb returns a boolean if a field has been set.

### GetArchitecture

`func (o *EngineGPUType) GetArchitecture() string`

GetArchitecture returns the Architecture field if non-nil, zero value otherwise.

### GetArchitectureOk

`func (o *EngineGPUType) GetArchitectureOk() (*string, bool)`

GetArchitectureOk returns a tuple with the Architecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchitecture

`func (o *EngineGPUType) SetArchitecture(v string)`

SetArchitecture sets Architecture field to given value.

### HasArchitecture

`func (o *EngineGPUType) HasArchitecture() bool`

HasArchitecture returns a boolean if a field has been set.

### GetFp16Tflops

`func (o *EngineGPUType) GetFp16Tflops() float32`

GetFp16Tflops returns the Fp16Tflops field if non-nil, zero value otherwise.

### GetFp16TflopsOk

`func (o *EngineGPUType) GetFp16TflopsOk() (*float32, bool)`

GetFp16TflopsOk returns a tuple with the Fp16Tflops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFp16Tflops

`func (o *EngineGPUType) SetFp16Tflops(v float32)`

SetFp16Tflops sets Fp16Tflops field to given value.

### HasFp16Tflops

`func (o *EngineGPUType) HasFp16Tflops() bool`

HasFp16Tflops returns a boolean if a field has been set.

### GetBf16Tflops

`func (o *EngineGPUType) GetBf16Tflops() float32`

GetBf16Tflops returns the Bf16Tflops field if non-nil, zero value otherwise.

### GetBf16TflopsOk

`func (o *EngineGPUType) GetBf16TflopsOk() (*float32, bool)`

GetBf16TflopsOk returns a tuple with the Bf16Tflops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBf16Tflops

`func (o *EngineGPUType) SetBf16Tflops(v float32)`

SetBf16Tflops sets Bf16Tflops field to given value.

### HasBf16Tflops

`func (o *EngineGPUType) HasBf16Tflops() bool`

HasBf16Tflops returns a boolean if a field has been set.

### GetFp8Tflops

`func (o *EngineGPUType) GetFp8Tflops() float32`

GetFp8Tflops returns the Fp8Tflops field if non-nil, zero value otherwise.

### GetFp8TflopsOk

`func (o *EngineGPUType) GetFp8TflopsOk() (*float32, bool)`

GetFp8TflopsOk returns a tuple with the Fp8Tflops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFp8Tflops

`func (o *EngineGPUType) SetFp8Tflops(v float32)`

SetFp8Tflops sets Fp8Tflops field to given value.

### HasFp8Tflops

`func (o *EngineGPUType) HasFp8Tflops() bool`

HasFp8Tflops returns a boolean if a field has been set.

### GetInterconnect

`func (o *EngineGPUType) GetInterconnect() string`

GetInterconnect returns the Interconnect field if non-nil, zero value otherwise.

### GetInterconnectOk

`func (o *EngineGPUType) GetInterconnectOk() (*string, bool)`

GetInterconnectOk returns a tuple with the Interconnect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterconnect

`func (o *EngineGPUType) SetInterconnect(v string)`

SetInterconnect sets Interconnect field to given value.

### HasInterconnect

`func (o *EngineGPUType) HasInterconnect() bool`

HasInterconnect returns a boolean if a field has been set.

### GetTotalCount

`func (o *EngineGPUType) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *EngineGPUType) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *EngineGPUType) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *EngineGPUType) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetAvailableCount

`func (o *EngineGPUType) GetAvailableCount() int32`

GetAvailableCount returns the AvailableCount field if non-nil, zero value otherwise.

### GetAvailableCountOk

`func (o *EngineGPUType) GetAvailableCountOk() (*int32, bool)`

GetAvailableCountOk returns a tuple with the AvailableCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableCount

`func (o *EngineGPUType) SetAvailableCount(v int32)`

SetAvailableCount sets AvailableCount field to given value.

### HasAvailableCount

`func (o *EngineGPUType) HasAvailableCount() bool`

HasAvailableCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


