# MachineView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Agent** | Pointer to **string** | Agent is the cloud Agent this machine runs, lifted out of the binding so a list reads without following one. Empty means nothing is bound — for a kind&#x3D;bot machine that means it costs money and answers nothing. | [optional] 
**Binding** | Pointer to [**AgentBinding**](AgentBinding.md) | Binding is the record joining this machine to that agent, carrying vm&#39;s own reconciled status and its reason. Absent means no runtime is bound, which is also what a stopped bot looks like: stopping unbinds and leaves the machine running. | [optional] 
**CreatedTime** | Pointer to **string** | CreatedTime is when the machine came into being: the provider&#39;s own creation timestamp for a Visor machine, passed through in whatever form it states it, and for a BYO machine the RFC 3339 moment it first dialed in. | [optional] 
**Gpu** | Pointer to **string** | GPU names the accelerators this machine holds (\&quot;H100\&quot;, or \&quot;2× NVIDIA GB10\&quot; for a BYO machine reporting a matched pair). Empty means the machine is not a GPU machine — the size slug does not parse as one, or nvidia-smi found nothing. | [optional] 
**Id** | Pointer to **string** | ID addresses this machine on the /v1/compute/machines/:id routes: the org-scoped NAME Visor keys a machine by, falling back to the provider id for a machine that has no name. A BYO machine&#39;s is the id it dialed in under. | [optional] 
**Image** | Pointer to **string** | Image is the OS image the machine booted from, as the provider names it. | [optional] 
**Mem** | Pointer to **string** | Mem is system RAM rendered for a human (\&quot;8 GB\&quot;), not a number to compute with. Empty when the provider&#39;s figure is ambiguous, or when the only figure available is a GPU slug&#39;s gb — that is VRAM, and reporting it as system RAM would be a fabrication. A BYO machine&#39;s RAM is on /v1/compute/fleet/workers. | [optional] 
**Name** | Pointer to **string** | Name is the label to show a human — Visor&#39;s displayName, or the machine name when it carries none. A BYO machine&#39;s is its hostname. It is not an address: ID is what the routes take. | [optional] 
**Os** | Pointer to **string** | Os is the operating system on the machine — Visor&#39;s record for a provisioned one, the host&#39;s own report (linux, darwin, windows) for a BYO one. | [optional] 
**PrivateIp** | Pointer to **string** | PrivateIp is the address on the provider&#39;s own network, reachable from the org&#39;s other machines in the same region. Empty on the same terms as PublicIp. | [optional] 
**Provider** | Pointer to **string** | Provider is the cloud that runs the machine (\&quot;digitalocean\&quot;), or \&quot;byo\&quot; for one the operator dialed in with &#x60;hanzo link&#x60;. | [optional] 
**PublicIp** | Pointer to **string** | PublicIp is the internet-facing address the provider assigned. Empty while a machine is still provisioning, and empty for a BYO machine — it dials out from behind NAT, so no address is ever learned for it. | [optional] 
**Region** | Pointer to **string** | Region is the provider region slug (\&quot;sfo3\&quot;), or the zone when the provider reports only that. \&quot;on-prem\&quot; for a BYO machine, which has no cloud region. | [optional] 
**Status** | Pointer to **string** | Status is the lifecycle state in the PROVIDER&#39;s own words (\&quot;active\&quot;, \&quot;running\&quot;, \&quot;off\&quot;), passed through rather than mapped onto a vocabulary of ours. A BYO machine&#39;s is \&quot;online\&quot; or \&quot;offline\&quot;, decided by whether its last heartbeat is within 90s. | [optional] 
**Type** | Pointer to **string** | Type is the provider SIZE SLUG the machine runs at (\&quot;s-2vcpu-4gb\&quot;, \&quot;gpu-h100x8-640gb\&quot;) — the value a launch asks for, and what Vcpu/Mem/GPU are read out of when the provider states them no other way. \&quot;byo-gpu\&quot; for a dialed-in machine, which was never bought from a size catalog. | [optional] 
**Vcpu** | Pointer to **int64** | Vcpu is logical cores — the provider&#39;s own cpuSize when that is a clean integer, else the count read out of the size slug (4 from \&quot;s-4vcpu-8gb\&quot;). ABSENT, never 0, when neither says. A BYO machine leaves it absent here; its real core count is on GET /v1/compute/fleet/workers. | [optional] 

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

### GetAgent

`func (o *MachineView) GetAgent() string`

GetAgent returns the Agent field if non-nil, zero value otherwise.

### GetAgentOk

`func (o *MachineView) GetAgentOk() (*string, bool)`

GetAgentOk returns a tuple with the Agent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgent

`func (o *MachineView) SetAgent(v string)`

SetAgent sets Agent field to given value.

### HasAgent

`func (o *MachineView) HasAgent() bool`

HasAgent returns a boolean if a field has been set.

### GetBinding

`func (o *MachineView) GetBinding() AgentBinding`

GetBinding returns the Binding field if non-nil, zero value otherwise.

### GetBindingOk

`func (o *MachineView) GetBindingOk() (*AgentBinding, bool)`

GetBindingOk returns a tuple with the Binding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinding

`func (o *MachineView) SetBinding(v AgentBinding)`

SetBinding sets Binding field to given value.

### HasBinding

`func (o *MachineView) HasBinding() bool`

HasBinding returns a boolean if a field has been set.

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

`func (o *MachineView) GetVcpu() int64`

GetVcpu returns the Vcpu field if non-nil, zero value otherwise.

### GetVcpuOk

`func (o *MachineView) GetVcpuOk() (*int64, bool)`

GetVcpuOk returns a tuple with the Vcpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcpu

`func (o *MachineView) SetVcpu(v int64)`

SetVcpu sets Vcpu field to given value.

### HasVcpu

`func (o *MachineView) HasVcpu() bool`

HasVcpu returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


