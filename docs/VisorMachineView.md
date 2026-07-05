# VisorMachineView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Org-scoped machine name (stable key the :id routes address) | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Provider** | Pointer to **string** |  | [optional] 
**PublicIp** | Pointer to **string** |  | [optional] 
**PrivateIp** | Pointer to **string** |  | [optional] 
**CreatedTime** | Pointer to **string** |  | [optional] 
**Vcpu** | Pointer to **int32** |  | [optional] 
**Gpu** | Pointer to **string** | GPU model when the size slug is a GPU accelerator | [optional] 
**Image** | Pointer to **string** |  | [optional] 
**Os** | Pointer to **string** |  | [optional] 

## Methods

### NewVisorMachineView

`func NewVisorMachineView() *VisorMachineView`

NewVisorMachineView instantiates a new VisorMachineView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVisorMachineViewWithDefaults

`func NewVisorMachineViewWithDefaults() *VisorMachineView`

NewVisorMachineViewWithDefaults instantiates a new VisorMachineView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VisorMachineView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VisorMachineView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VisorMachineView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VisorMachineView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *VisorMachineView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VisorMachineView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VisorMachineView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VisorMachineView) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRegion

`func (o *VisorMachineView) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *VisorMachineView) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *VisorMachineView) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *VisorMachineView) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetType

`func (o *VisorMachineView) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VisorMachineView) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VisorMachineView) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *VisorMachineView) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStatus

`func (o *VisorMachineView) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VisorMachineView) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VisorMachineView) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VisorMachineView) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetProvider

`func (o *VisorMachineView) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *VisorMachineView) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *VisorMachineView) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *VisorMachineView) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetPublicIp

`func (o *VisorMachineView) GetPublicIp() string`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *VisorMachineView) GetPublicIpOk() (*string, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIp

`func (o *VisorMachineView) SetPublicIp(v string)`

SetPublicIp sets PublicIp field to given value.

### HasPublicIp

`func (o *VisorMachineView) HasPublicIp() bool`

HasPublicIp returns a boolean if a field has been set.

### GetPrivateIp

`func (o *VisorMachineView) GetPrivateIp() string`

GetPrivateIp returns the PrivateIp field if non-nil, zero value otherwise.

### GetPrivateIpOk

`func (o *VisorMachineView) GetPrivateIpOk() (*string, bool)`

GetPrivateIpOk returns a tuple with the PrivateIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIp

`func (o *VisorMachineView) SetPrivateIp(v string)`

SetPrivateIp sets PrivateIp field to given value.

### HasPrivateIp

`func (o *VisorMachineView) HasPrivateIp() bool`

HasPrivateIp returns a boolean if a field has been set.

### GetCreatedTime

`func (o *VisorMachineView) GetCreatedTime() string`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *VisorMachineView) GetCreatedTimeOk() (*string, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *VisorMachineView) SetCreatedTime(v string)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *VisorMachineView) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetVcpu

`func (o *VisorMachineView) GetVcpu() int32`

GetVcpu returns the Vcpu field if non-nil, zero value otherwise.

### GetVcpuOk

`func (o *VisorMachineView) GetVcpuOk() (*int32, bool)`

GetVcpuOk returns a tuple with the Vcpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcpu

`func (o *VisorMachineView) SetVcpu(v int32)`

SetVcpu sets Vcpu field to given value.

### HasVcpu

`func (o *VisorMachineView) HasVcpu() bool`

HasVcpu returns a boolean if a field has been set.

### GetGpu

`func (o *VisorMachineView) GetGpu() string`

GetGpu returns the Gpu field if non-nil, zero value otherwise.

### GetGpuOk

`func (o *VisorMachineView) GetGpuOk() (*string, bool)`

GetGpuOk returns a tuple with the Gpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpu

`func (o *VisorMachineView) SetGpu(v string)`

SetGpu sets Gpu field to given value.

### HasGpu

`func (o *VisorMachineView) HasGpu() bool`

HasGpu returns a boolean if a field has been set.

### GetImage

`func (o *VisorMachineView) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *VisorMachineView) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *VisorMachineView) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *VisorMachineView) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetOs

`func (o *VisorMachineView) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *VisorMachineView) GetOsOk() (*string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *VisorMachineView) SetOs(v string)`

SetOs sets Os field to given value.

### HasOs

`func (o *VisorMachineView) HasOs() bool`

HasOs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


