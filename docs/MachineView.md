# MachineView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedTime** | Pointer to **string** |  | [optional] 
**Gpu** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Image** | Pointer to **string** |  | [optional] 
**Mem** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Os** | Pointer to **string** |  | [optional] 
**PrivateIp** | Pointer to **string** |  | [optional] 
**Provider** | Pointer to **string** |  | [optional] 
**PublicIp** | Pointer to **string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Vcpu** | Pointer to **int32** |  | [optional] 

## Methods

### NewMachineView

`func NewMachineView() *MachineView`

NewMachineView instantiates a new MachineView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMachineViewWithDefaults

`func NewMachineViewWithDefaults() *MachineView`

NewMachineViewWithDefaults instantiates a new MachineView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedTime

`func (o *MachineView) GetCreatedTime() string`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *MachineView) GetCreatedTimeOk() (*string, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *MachineView) SetCreatedTime(v string)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *MachineView) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetGpu

`func (o *MachineView) GetGpu() string`

GetGpu returns the Gpu field if non-nil, zero value otherwise.

### GetGpuOk

`func (o *MachineView) GetGpuOk() (*string, bool)`

GetGpuOk returns a tuple with the Gpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpu

`func (o *MachineView) SetGpu(v string)`

SetGpu sets Gpu field to given value.

### HasGpu

`func (o *MachineView) HasGpu() bool`

HasGpu returns a boolean if a field has been set.

### GetId

`func (o *MachineView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MachineView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MachineView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MachineView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImage

`func (o *MachineView) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *MachineView) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *MachineView) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *MachineView) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetMem

`func (o *MachineView) GetMem() string`

GetMem returns the Mem field if non-nil, zero value otherwise.

### GetMemOk

`func (o *MachineView) GetMemOk() (*string, bool)`

GetMemOk returns a tuple with the Mem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMem

`func (o *MachineView) SetMem(v string)`

SetMem sets Mem field to given value.

### HasMem

`func (o *MachineView) HasMem() bool`

HasMem returns a boolean if a field has been set.

### GetName

`func (o *MachineView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MachineView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MachineView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MachineView) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOs

`func (o *MachineView) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *MachineView) GetOsOk() (*string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *MachineView) SetOs(v string)`

SetOs sets Os field to given value.

### HasOs

`func (o *MachineView) HasOs() bool`

HasOs returns a boolean if a field has been set.

### GetPrivateIp

`func (o *MachineView) GetPrivateIp() string`

GetPrivateIp returns the PrivateIp field if non-nil, zero value otherwise.

### GetPrivateIpOk

`func (o *MachineView) GetPrivateIpOk() (*string, bool)`

GetPrivateIpOk returns a tuple with the PrivateIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIp

`func (o *MachineView) SetPrivateIp(v string)`

SetPrivateIp sets PrivateIp field to given value.

### HasPrivateIp

`func (o *MachineView) HasPrivateIp() bool`

HasPrivateIp returns a boolean if a field has been set.

### GetProvider

`func (o *MachineView) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *MachineView) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *MachineView) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *MachineView) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetPublicIp

`func (o *MachineView) GetPublicIp() string`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *MachineView) GetPublicIpOk() (*string, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIp

`func (o *MachineView) SetPublicIp(v string)`

SetPublicIp sets PublicIp field to given value.

### HasPublicIp

`func (o *MachineView) HasPublicIp() bool`

HasPublicIp returns a boolean if a field has been set.

### GetRegion

`func (o *MachineView) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *MachineView) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *MachineView) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *MachineView) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetStatus

`func (o *MachineView) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MachineView) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MachineView) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MachineView) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *MachineView) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MachineView) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MachineView) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MachineView) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVcpu

`func (o *MachineView) GetVcpu() int32`

GetVcpu returns the Vcpu field if non-nil, zero value otherwise.

### GetVcpuOk

`func (o *MachineView) GetVcpuOk() (*int32, bool)`

GetVcpuOk returns a tuple with the Vcpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcpu

`func (o *MachineView) SetVcpu(v int32)`

SetVcpu sets Vcpu field to given value.

### HasVcpu

`func (o *MachineView) HasVcpu() bool`

HasVcpu returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


