# CloudBotView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Agent** | Pointer to **string** |  | [optional] 
**Binding** | Pointer to [**CloudAgentBinding**](CloudAgentBinding.md) |  | [optional] 
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

### NewCloudBotView

`func NewCloudBotView() *CloudBotView`

NewCloudBotView instantiates a new CloudBotView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudBotViewWithDefaults

`func NewCloudBotViewWithDefaults() *CloudBotView`

NewCloudBotViewWithDefaults instantiates a new CloudBotView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgent

`func (o *CloudBotView) GetAgent() string`

GetAgent returns the Agent field if non-nil, zero value otherwise.

### GetAgentOk

`func (o *CloudBotView) GetAgentOk() (*string, bool)`

GetAgentOk returns a tuple with the Agent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgent

`func (o *CloudBotView) SetAgent(v string)`

SetAgent sets Agent field to given value.

### HasAgent

`func (o *CloudBotView) HasAgent() bool`

HasAgent returns a boolean if a field has been set.

### GetBinding

`func (o *CloudBotView) GetBinding() CloudAgentBinding`

GetBinding returns the Binding field if non-nil, zero value otherwise.

### GetBindingOk

`func (o *CloudBotView) GetBindingOk() (*CloudAgentBinding, bool)`

GetBindingOk returns a tuple with the Binding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinding

`func (o *CloudBotView) SetBinding(v CloudAgentBinding)`

SetBinding sets Binding field to given value.

### HasBinding

`func (o *CloudBotView) HasBinding() bool`

HasBinding returns a boolean if a field has been set.

### GetCreatedTime

`func (o *CloudBotView) GetCreatedTime() string`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *CloudBotView) GetCreatedTimeOk() (*string, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *CloudBotView) SetCreatedTime(v string)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *CloudBotView) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetGpu

`func (o *CloudBotView) GetGpu() string`

GetGpu returns the Gpu field if non-nil, zero value otherwise.

### GetGpuOk

`func (o *CloudBotView) GetGpuOk() (*string, bool)`

GetGpuOk returns a tuple with the Gpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpu

`func (o *CloudBotView) SetGpu(v string)`

SetGpu sets Gpu field to given value.

### HasGpu

`func (o *CloudBotView) HasGpu() bool`

HasGpu returns a boolean if a field has been set.

### GetId

`func (o *CloudBotView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudBotView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudBotView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudBotView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImage

`func (o *CloudBotView) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *CloudBotView) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *CloudBotView) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *CloudBotView) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetMem

`func (o *CloudBotView) GetMem() string`

GetMem returns the Mem field if non-nil, zero value otherwise.

### GetMemOk

`func (o *CloudBotView) GetMemOk() (*string, bool)`

GetMemOk returns a tuple with the Mem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMem

`func (o *CloudBotView) SetMem(v string)`

SetMem sets Mem field to given value.

### HasMem

`func (o *CloudBotView) HasMem() bool`

HasMem returns a boolean if a field has been set.

### GetName

`func (o *CloudBotView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudBotView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudBotView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudBotView) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOs

`func (o *CloudBotView) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *CloudBotView) GetOsOk() (*string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *CloudBotView) SetOs(v string)`

SetOs sets Os field to given value.

### HasOs

`func (o *CloudBotView) HasOs() bool`

HasOs returns a boolean if a field has been set.

### GetPrivateIp

`func (o *CloudBotView) GetPrivateIp() string`

GetPrivateIp returns the PrivateIp field if non-nil, zero value otherwise.

### GetPrivateIpOk

`func (o *CloudBotView) GetPrivateIpOk() (*string, bool)`

GetPrivateIpOk returns a tuple with the PrivateIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIp

`func (o *CloudBotView) SetPrivateIp(v string)`

SetPrivateIp sets PrivateIp field to given value.

### HasPrivateIp

`func (o *CloudBotView) HasPrivateIp() bool`

HasPrivateIp returns a boolean if a field has been set.

### GetProvider

`func (o *CloudBotView) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *CloudBotView) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *CloudBotView) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *CloudBotView) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetPublicIp

`func (o *CloudBotView) GetPublicIp() string`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *CloudBotView) GetPublicIpOk() (*string, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIp

`func (o *CloudBotView) SetPublicIp(v string)`

SetPublicIp sets PublicIp field to given value.

### HasPublicIp

`func (o *CloudBotView) HasPublicIp() bool`

HasPublicIp returns a boolean if a field has been set.

### GetRegion

`func (o *CloudBotView) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CloudBotView) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CloudBotView) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *CloudBotView) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetStatus

`func (o *CloudBotView) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudBotView) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudBotView) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudBotView) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *CloudBotView) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CloudBotView) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CloudBotView) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CloudBotView) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVcpu

`func (o *CloudBotView) GetVcpu() int32`

GetVcpu returns the Vcpu field if non-nil, zero value otherwise.

### GetVcpuOk

`func (o *CloudBotView) GetVcpuOk() (*int32, bool)`

GetVcpuOk returns a tuple with the Vcpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcpu

`func (o *CloudBotView) SetVcpu(v int32)`

SetVcpu sets Vcpu field to given value.

### HasVcpu

`func (o *CloudBotView) HasVcpu() bool`

HasVcpu returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


