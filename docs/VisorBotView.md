# VisorBotView

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
**Agent** | Pointer to **string** | The bound agent name | [optional] 
**Binding** | Pointer to [**VisorAgentBinding**](VisorAgentBinding.md) |  | [optional] 

## Methods

### NewVisorBotView

`func NewVisorBotView() *VisorBotView`

NewVisorBotView instantiates a new VisorBotView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVisorBotViewWithDefaults

`func NewVisorBotViewWithDefaults() *VisorBotView`

NewVisorBotViewWithDefaults instantiates a new VisorBotView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VisorBotView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VisorBotView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VisorBotView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VisorBotView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *VisorBotView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VisorBotView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VisorBotView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VisorBotView) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRegion

`func (o *VisorBotView) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *VisorBotView) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *VisorBotView) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *VisorBotView) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetType

`func (o *VisorBotView) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VisorBotView) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VisorBotView) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *VisorBotView) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStatus

`func (o *VisorBotView) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VisorBotView) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VisorBotView) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VisorBotView) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetProvider

`func (o *VisorBotView) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *VisorBotView) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *VisorBotView) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *VisorBotView) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetPublicIp

`func (o *VisorBotView) GetPublicIp() string`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *VisorBotView) GetPublicIpOk() (*string, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIp

`func (o *VisorBotView) SetPublicIp(v string)`

SetPublicIp sets PublicIp field to given value.

### HasPublicIp

`func (o *VisorBotView) HasPublicIp() bool`

HasPublicIp returns a boolean if a field has been set.

### GetPrivateIp

`func (o *VisorBotView) GetPrivateIp() string`

GetPrivateIp returns the PrivateIp field if non-nil, zero value otherwise.

### GetPrivateIpOk

`func (o *VisorBotView) GetPrivateIpOk() (*string, bool)`

GetPrivateIpOk returns a tuple with the PrivateIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIp

`func (o *VisorBotView) SetPrivateIp(v string)`

SetPrivateIp sets PrivateIp field to given value.

### HasPrivateIp

`func (o *VisorBotView) HasPrivateIp() bool`

HasPrivateIp returns a boolean if a field has been set.

### GetCreatedTime

`func (o *VisorBotView) GetCreatedTime() string`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *VisorBotView) GetCreatedTimeOk() (*string, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *VisorBotView) SetCreatedTime(v string)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *VisorBotView) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetVcpu

`func (o *VisorBotView) GetVcpu() int32`

GetVcpu returns the Vcpu field if non-nil, zero value otherwise.

### GetVcpuOk

`func (o *VisorBotView) GetVcpuOk() (*int32, bool)`

GetVcpuOk returns a tuple with the Vcpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcpu

`func (o *VisorBotView) SetVcpu(v int32)`

SetVcpu sets Vcpu field to given value.

### HasVcpu

`func (o *VisorBotView) HasVcpu() bool`

HasVcpu returns a boolean if a field has been set.

### GetGpu

`func (o *VisorBotView) GetGpu() string`

GetGpu returns the Gpu field if non-nil, zero value otherwise.

### GetGpuOk

`func (o *VisorBotView) GetGpuOk() (*string, bool)`

GetGpuOk returns a tuple with the Gpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpu

`func (o *VisorBotView) SetGpu(v string)`

SetGpu sets Gpu field to given value.

### HasGpu

`func (o *VisorBotView) HasGpu() bool`

HasGpu returns a boolean if a field has been set.

### GetImage

`func (o *VisorBotView) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *VisorBotView) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *VisorBotView) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *VisorBotView) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetOs

`func (o *VisorBotView) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *VisorBotView) GetOsOk() (*string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *VisorBotView) SetOs(v string)`

SetOs sets Os field to given value.

### HasOs

`func (o *VisorBotView) HasOs() bool`

HasOs returns a boolean if a field has been set.

### GetAgent

`func (o *VisorBotView) GetAgent() string`

GetAgent returns the Agent field if non-nil, zero value otherwise.

### GetAgentOk

`func (o *VisorBotView) GetAgentOk() (*string, bool)`

GetAgentOk returns a tuple with the Agent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgent

`func (o *VisorBotView) SetAgent(v string)`

SetAgent sets Agent field to given value.

### HasAgent

`func (o *VisorBotView) HasAgent() bool`

HasAgent returns a boolean if a field has been set.

### GetBinding

`func (o *VisorBotView) GetBinding() VisorAgentBinding`

GetBinding returns the Binding field if non-nil, zero value otherwise.

### GetBindingOk

`func (o *VisorBotView) GetBindingOk() (*VisorAgentBinding, bool)`

GetBindingOk returns a tuple with the Binding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinding

`func (o *VisorBotView) SetBinding(v VisorAgentBinding)`

SetBinding sets Binding field to given value.

### HasBinding

`func (o *VisorBotView) HasBinding() bool`

HasBinding returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


