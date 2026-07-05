# EngineGPUDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | Pointer to **int32** |  | [optional] 
**Model** | Pointer to **string** | GPU model (e.g. A100-SXM4-80GB, H100-SXM5-80GB) | [optional] 
**MemoryMb** | Pointer to **int32** |  | [optional] 
**UtilizationPercent** | Pointer to **float32** |  | [optional] 
**MemoryUsedMb** | Pointer to **int32** |  | [optional] 
**TemperatureC** | Pointer to **int32** |  | [optional] 
**PowerDrawW** | Pointer to **float32** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewEngineGPUDevice

`func NewEngineGPUDevice() *EngineGPUDevice`

NewEngineGPUDevice instantiates a new EngineGPUDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEngineGPUDeviceWithDefaults

`func NewEngineGPUDeviceWithDefaults() *EngineGPUDevice`

NewEngineGPUDeviceWithDefaults instantiates a new EngineGPUDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *EngineGPUDevice) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *EngineGPUDevice) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *EngineGPUDevice) SetIndex(v int32)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *EngineGPUDevice) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetModel

`func (o *EngineGPUDevice) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *EngineGPUDevice) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *EngineGPUDevice) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *EngineGPUDevice) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetMemoryMb

`func (o *EngineGPUDevice) GetMemoryMb() int32`

GetMemoryMb returns the MemoryMb field if non-nil, zero value otherwise.

### GetMemoryMbOk

`func (o *EngineGPUDevice) GetMemoryMbOk() (*int32, bool)`

GetMemoryMbOk returns a tuple with the MemoryMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryMb

`func (o *EngineGPUDevice) SetMemoryMb(v int32)`

SetMemoryMb sets MemoryMb field to given value.

### HasMemoryMb

`func (o *EngineGPUDevice) HasMemoryMb() bool`

HasMemoryMb returns a boolean if a field has been set.

### GetUtilizationPercent

`func (o *EngineGPUDevice) GetUtilizationPercent() float32`

GetUtilizationPercent returns the UtilizationPercent field if non-nil, zero value otherwise.

### GetUtilizationPercentOk

`func (o *EngineGPUDevice) GetUtilizationPercentOk() (*float32, bool)`

GetUtilizationPercentOk returns a tuple with the UtilizationPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtilizationPercent

`func (o *EngineGPUDevice) SetUtilizationPercent(v float32)`

SetUtilizationPercent sets UtilizationPercent field to given value.

### HasUtilizationPercent

`func (o *EngineGPUDevice) HasUtilizationPercent() bool`

HasUtilizationPercent returns a boolean if a field has been set.

### GetMemoryUsedMb

`func (o *EngineGPUDevice) GetMemoryUsedMb() int32`

GetMemoryUsedMb returns the MemoryUsedMb field if non-nil, zero value otherwise.

### GetMemoryUsedMbOk

`func (o *EngineGPUDevice) GetMemoryUsedMbOk() (*int32, bool)`

GetMemoryUsedMbOk returns a tuple with the MemoryUsedMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryUsedMb

`func (o *EngineGPUDevice) SetMemoryUsedMb(v int32)`

SetMemoryUsedMb sets MemoryUsedMb field to given value.

### HasMemoryUsedMb

`func (o *EngineGPUDevice) HasMemoryUsedMb() bool`

HasMemoryUsedMb returns a boolean if a field has been set.

### GetTemperatureC

`func (o *EngineGPUDevice) GetTemperatureC() int32`

GetTemperatureC returns the TemperatureC field if non-nil, zero value otherwise.

### GetTemperatureCOk

`func (o *EngineGPUDevice) GetTemperatureCOk() (*int32, bool)`

GetTemperatureCOk returns a tuple with the TemperatureC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperatureC

`func (o *EngineGPUDevice) SetTemperatureC(v int32)`

SetTemperatureC sets TemperatureC field to given value.

### HasTemperatureC

`func (o *EngineGPUDevice) HasTemperatureC() bool`

HasTemperatureC returns a boolean if a field has been set.

### GetPowerDrawW

`func (o *EngineGPUDevice) GetPowerDrawW() float32`

GetPowerDrawW returns the PowerDrawW field if non-nil, zero value otherwise.

### GetPowerDrawWOk

`func (o *EngineGPUDevice) GetPowerDrawWOk() (*float32, bool)`

GetPowerDrawWOk returns a tuple with the PowerDrawW field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerDrawW

`func (o *EngineGPUDevice) SetPowerDrawW(v float32)`

SetPowerDrawW sets PowerDrawW field to given value.

### HasPowerDrawW

`func (o *EngineGPUDevice) HasPowerDrawW() bool`

HasPowerDrawW returns a boolean if a field has been set.

### GetStatus

`func (o *EngineGPUDevice) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EngineGPUDevice) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EngineGPUDevice) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EngineGPUDevice) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


