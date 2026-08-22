# BotView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Agent** | Pointer to **string** | Agent is the cloud Agent this machine runs, lifted out of the binding so a list of bots reads without following one. Empty means the machine is a bot machine with nothing bound — it costs money and answers nothing. | [optional] 
**Binding** | Pointer to [**AgentBinding**](AgentBinding.md) | Binding is the record joining this machine to that agent, carrying vm&#39;s own reconciled status and its reason. Absent means no runtime is bound, which is also what a stopped bot looks like: stopping unbinds and leaves the machine running. | [optional] 
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

### NewBotView

`func NewBotView() *BotView`

NewBotView instantiates a new BotView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBotViewWithDefaults

`func NewBotViewWithDefaults() *BotView`

NewBotViewWithDefaults instantiates a new BotView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgent

`func (o *BotView) GetAgent() string`

GetAgent returns the Agent field if non-nil, zero value otherwise.

### GetAgentOk

`func (o *BotView) GetAgentOk() (*string, bool)`

GetAgentOk returns a tuple with the Agent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgent

`func (o *BotView) SetAgent(v string)`

SetAgent sets Agent field to given value.

### HasAgent

`func (o *BotView) HasAgent() bool`

HasAgent returns a boolean if a field has been set.

### GetBinding

`func (o *BotView) GetBinding() AgentBinding`

GetBinding returns the Binding field if non-nil, zero value otherwise.

### GetBindingOk

`func (o *BotView) GetBindingOk() (*AgentBinding, bool)`

GetBindingOk returns a tuple with the Binding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinding

`func (o *BotView) SetBinding(v AgentBinding)`

SetBinding sets Binding field to given value.

### HasBinding

`func (o *BotView) HasBinding() bool`

HasBinding returns a boolean if a field has been set.

### GetCreatedTime

`func (o *BotView) GetCreatedTime() string`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *BotView) GetCreatedTimeOk() (*string, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *BotView) SetCreatedTime(v string)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *BotView) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetGpu

`func (o *BotView) GetGpu() string`

GetGpu returns the Gpu field if non-nil, zero value otherwise.

### GetGpuOk

`func (o *BotView) GetGpuOk() (*string, bool)`

GetGpuOk returns a tuple with the Gpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpu

`func (o *BotView) SetGpu(v string)`

SetGpu sets Gpu field to given value.

### HasGpu

`func (o *BotView) HasGpu() bool`

HasGpu returns a boolean if a field has been set.

### GetId

`func (o *BotView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BotView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BotView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BotView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImage

`func (o *BotView) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *BotView) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *BotView) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *BotView) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetMem

`func (o *BotView) GetMem() string`

GetMem returns the Mem field if non-nil, zero value otherwise.

### GetMemOk

`func (o *BotView) GetMemOk() (*string, bool)`

GetMemOk returns a tuple with the Mem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMem

`func (o *BotView) SetMem(v string)`

SetMem sets Mem field to given value.

### HasMem

`func (o *BotView) HasMem() bool`

HasMem returns a boolean if a field has been set.

### GetName

`func (o *BotView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BotView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BotView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BotView) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOs

`func (o *BotView) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *BotView) GetOsOk() (*string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *BotView) SetOs(v string)`

SetOs sets Os field to given value.

### HasOs

`func (o *BotView) HasOs() bool`

HasOs returns a boolean if a field has been set.

### GetPrivateIp

`func (o *BotView) GetPrivateIp() string`

GetPrivateIp returns the PrivateIp field if non-nil, zero value otherwise.

### GetPrivateIpOk

`func (o *BotView) GetPrivateIpOk() (*string, bool)`

GetPrivateIpOk returns a tuple with the PrivateIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIp

`func (o *BotView) SetPrivateIp(v string)`

SetPrivateIp sets PrivateIp field to given value.

### HasPrivateIp

`func (o *BotView) HasPrivateIp() bool`

HasPrivateIp returns a boolean if a field has been set.

### GetProvider

`func (o *BotView) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *BotView) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *BotView) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *BotView) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetPublicIp

`func (o *BotView) GetPublicIp() string`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *BotView) GetPublicIpOk() (*string, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIp

`func (o *BotView) SetPublicIp(v string)`

SetPublicIp sets PublicIp field to given value.

### HasPublicIp

`func (o *BotView) HasPublicIp() bool`

HasPublicIp returns a boolean if a field has been set.

### GetRegion

`func (o *BotView) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *BotView) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *BotView) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *BotView) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetStatus

`func (o *BotView) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BotView) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BotView) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BotView) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *BotView) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BotView) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BotView) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *BotView) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVcpu

`func (o *BotView) GetVcpu() int32`

GetVcpu returns the Vcpu field if non-nil, zero value otherwise.

### GetVcpuOk

`func (o *BotView) GetVcpuOk() (*int32, bool)`

GetVcpuOk returns a tuple with the Vcpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcpu

`func (o *BotView) SetVcpu(v int32)`

SetVcpu sets Vcpu field to given value.

### HasVcpu

`func (o *BotView) HasVcpu() bool`

HasVcpu returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


