# CloudMachineView

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

### NewCloudMachineView

`func NewCloudMachineView() *CloudMachineView`

NewCloudMachineView instantiates a new CloudMachineView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudMachineViewWithDefaults

`func NewCloudMachineViewWithDefaults() *CloudMachineView`

NewCloudMachineViewWithDefaults instantiates a new CloudMachineView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedTime

`func (o *CloudMachineView) GetCreatedTime() string`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *CloudMachineView) GetCreatedTimeOk() (*string, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *CloudMachineView) SetCreatedTime(v string)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *CloudMachineView) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetGpu

`func (o *CloudMachineView) GetGpu() string`

GetGpu returns the Gpu field if non-nil, zero value otherwise.

### GetGpuOk

`func (o *CloudMachineView) GetGpuOk() (*string, bool)`

GetGpuOk returns a tuple with the Gpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpu

`func (o *CloudMachineView) SetGpu(v string)`

SetGpu sets Gpu field to given value.

### HasGpu

`func (o *CloudMachineView) HasGpu() bool`

HasGpu returns a boolean if a field has been set.

### GetId

`func (o *CloudMachineView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudMachineView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudMachineView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudMachineView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImage

`func (o *CloudMachineView) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *CloudMachineView) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *CloudMachineView) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *CloudMachineView) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetMem

`func (o *CloudMachineView) GetMem() string`

GetMem returns the Mem field if non-nil, zero value otherwise.

### GetMemOk

`func (o *CloudMachineView) GetMemOk() (*string, bool)`

GetMemOk returns a tuple with the Mem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMem

`func (o *CloudMachineView) SetMem(v string)`

SetMem sets Mem field to given value.

### HasMem

`func (o *CloudMachineView) HasMem() bool`

HasMem returns a boolean if a field has been set.

### GetName

`func (o *CloudMachineView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudMachineView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudMachineView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudMachineView) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOs

`func (o *CloudMachineView) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *CloudMachineView) GetOsOk() (*string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *CloudMachineView) SetOs(v string)`

SetOs sets Os field to given value.

### HasOs

`func (o *CloudMachineView) HasOs() bool`

HasOs returns a boolean if a field has been set.

### GetPrivateIp

`func (o *CloudMachineView) GetPrivateIp() string`

GetPrivateIp returns the PrivateIp field if non-nil, zero value otherwise.

### GetPrivateIpOk

`func (o *CloudMachineView) GetPrivateIpOk() (*string, bool)`

GetPrivateIpOk returns a tuple with the PrivateIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIp

`func (o *CloudMachineView) SetPrivateIp(v string)`

SetPrivateIp sets PrivateIp field to given value.

### HasPrivateIp

`func (o *CloudMachineView) HasPrivateIp() bool`

HasPrivateIp returns a boolean if a field has been set.

### GetProvider

`func (o *CloudMachineView) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *CloudMachineView) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *CloudMachineView) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *CloudMachineView) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetPublicIp

`func (o *CloudMachineView) GetPublicIp() string`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *CloudMachineView) GetPublicIpOk() (*string, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIp

`func (o *CloudMachineView) SetPublicIp(v string)`

SetPublicIp sets PublicIp field to given value.

### HasPublicIp

`func (o *CloudMachineView) HasPublicIp() bool`

HasPublicIp returns a boolean if a field has been set.

### GetRegion

`func (o *CloudMachineView) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CloudMachineView) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CloudMachineView) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *CloudMachineView) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetStatus

`func (o *CloudMachineView) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudMachineView) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudMachineView) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudMachineView) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *CloudMachineView) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CloudMachineView) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CloudMachineView) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CloudMachineView) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVcpu

`func (o *CloudMachineView) GetVcpu() int32`

GetVcpu returns the Vcpu field if non-nil, zero value otherwise.

### GetVcpuOk

`func (o *CloudMachineView) GetVcpuOk() (*int32, bool)`

GetVcpuOk returns a tuple with the Vcpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcpu

`func (o *CloudMachineView) SetVcpu(v int32)`

SetVcpu sets Vcpu field to given value.

### HasVcpu

`func (o *CloudMachineView) HasVcpu() bool`

HasVcpu returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


