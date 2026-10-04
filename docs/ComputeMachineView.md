# ComputeMachineView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Agent** | Pointer to **string** | Agent is the cloud Agent this machine runs, lifted out of the binding so a list reads without following one. Empty means nothing is bound — for a kind&#x3D;bot machine that means it costs money and answers nothing. | [optional] 
**Binding** | Pointer to [**ComputeAgentBinding**](ComputeAgentBinding.md) | Binding is the record joining this machine to that agent, carrying vm&#39;s own reconciled status and its reason. Absent means no runtime is bound, which is also what a stopped bot looks like: stopping unbinds and leaves the machine running. | [optional] 
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

### NewComputeMachineView

`func NewComputeMachineView() *ComputeMachineView`

NewComputeMachineView instantiates a new ComputeMachineView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputeMachineViewWithDefaults

`func NewComputeMachineViewWithDefaults() *ComputeMachineView`

NewComputeMachineViewWithDefaults instantiates a new ComputeMachineView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgent

`func (o *ComputeMachineView) GetAgent() string`

GetAgent returns the Agent field if non-nil, zero value otherwise.

### GetAgentOk

`func (o *ComputeMachineView) GetAgentOk() (*string, bool)`

GetAgentOk returns a tuple with the Agent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgent

`func (o *ComputeMachineView) SetAgent(v string)`

SetAgent sets Agent field to given value.

### HasAgent

`func (o *ComputeMachineView) HasAgent() bool`

HasAgent returns a boolean if a field has been set.

### GetBinding

`func (o *ComputeMachineView) GetBinding() ComputeAgentBinding`

GetBinding returns the Binding field if non-nil, zero value otherwise.

### GetBindingOk

`func (o *ComputeMachineView) GetBindingOk() (*ComputeAgentBinding, bool)`

GetBindingOk returns a tuple with the Binding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinding

`func (o *ComputeMachineView) SetBinding(v ComputeAgentBinding)`

SetBinding sets Binding field to given value.

### HasBinding

`func (o *ComputeMachineView) HasBinding() bool`

HasBinding returns a boolean if a field has been set.

### GetCreatedTime

`func (o *ComputeMachineView) GetCreatedTime() string`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *ComputeMachineView) GetCreatedTimeOk() (*string, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *ComputeMachineView) SetCreatedTime(v string)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *ComputeMachineView) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetGpu

`func (o *ComputeMachineView) GetGpu() string`

GetGpu returns the Gpu field if non-nil, zero value otherwise.

### GetGpuOk

`func (o *ComputeMachineView) GetGpuOk() (*string, bool)`

GetGpuOk returns a tuple with the Gpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpu

`func (o *ComputeMachineView) SetGpu(v string)`

SetGpu sets Gpu field to given value.

### HasGpu

`func (o *ComputeMachineView) HasGpu() bool`

HasGpu returns a boolean if a field has been set.

### GetId

`func (o *ComputeMachineView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ComputeMachineView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ComputeMachineView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ComputeMachineView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImage

`func (o *ComputeMachineView) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *ComputeMachineView) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *ComputeMachineView) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *ComputeMachineView) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetMem

`func (o *ComputeMachineView) GetMem() string`

GetMem returns the Mem field if non-nil, zero value otherwise.

### GetMemOk

`func (o *ComputeMachineView) GetMemOk() (*string, bool)`

GetMemOk returns a tuple with the Mem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMem

`func (o *ComputeMachineView) SetMem(v string)`

SetMem sets Mem field to given value.

### HasMem

`func (o *ComputeMachineView) HasMem() bool`

HasMem returns a boolean if a field has been set.

### GetName

`func (o *ComputeMachineView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ComputeMachineView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ComputeMachineView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ComputeMachineView) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOs

`func (o *ComputeMachineView) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *ComputeMachineView) GetOsOk() (*string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *ComputeMachineView) SetOs(v string)`

SetOs sets Os field to given value.

### HasOs

`func (o *ComputeMachineView) HasOs() bool`

HasOs returns a boolean if a field has been set.

### GetPrivateIp

`func (o *ComputeMachineView) GetPrivateIp() string`

GetPrivateIp returns the PrivateIp field if non-nil, zero value otherwise.

### GetPrivateIpOk

`func (o *ComputeMachineView) GetPrivateIpOk() (*string, bool)`

GetPrivateIpOk returns a tuple with the PrivateIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIp

`func (o *ComputeMachineView) SetPrivateIp(v string)`

SetPrivateIp sets PrivateIp field to given value.

### HasPrivateIp

`func (o *ComputeMachineView) HasPrivateIp() bool`

HasPrivateIp returns a boolean if a field has been set.

### GetProvider

`func (o *ComputeMachineView) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ComputeMachineView) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ComputeMachineView) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *ComputeMachineView) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetPublicIp

`func (o *ComputeMachineView) GetPublicIp() string`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *ComputeMachineView) GetPublicIpOk() (*string, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIp

`func (o *ComputeMachineView) SetPublicIp(v string)`

SetPublicIp sets PublicIp field to given value.

### HasPublicIp

`func (o *ComputeMachineView) HasPublicIp() bool`

HasPublicIp returns a boolean if a field has been set.

### GetRegion

`func (o *ComputeMachineView) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ComputeMachineView) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ComputeMachineView) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *ComputeMachineView) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetStatus

`func (o *ComputeMachineView) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ComputeMachineView) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ComputeMachineView) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ComputeMachineView) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *ComputeMachineView) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ComputeMachineView) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ComputeMachineView) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ComputeMachineView) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVcpu

`func (o *ComputeMachineView) GetVcpu() int64`

GetVcpu returns the Vcpu field if non-nil, zero value otherwise.

### GetVcpuOk

`func (o *ComputeMachineView) GetVcpuOk() (*int64, bool)`

GetVcpuOk returns a tuple with the Vcpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcpu

`func (o *ComputeMachineView) SetVcpu(v int64)`

SetVcpu sets Vcpu field to given value.

### HasVcpu

`func (o *ComputeMachineView) HasVcpu() bool`

HasVcpu returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


