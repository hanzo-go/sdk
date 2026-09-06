# ByoWorker

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Arch** | Pointer to **string** | Arch/CPUs/Memory are the connecting host&#39;s static CPU spec, mirrored from the registration: Arch is runtime.GOARCH (amd64 | arm64), Memory is total RAM in BYTES — the same fields a code-linked run-target carries, so the /v1/compute/fleet board renders a linked node&#39;s arch + cores + RAM like any other unit. | [optional] 
**Capabilities** | Pointer to **[]string** | Capabilities is what this worker offers the org: \&quot;studio.render\&quot; when the node can render, \&quot;engine.serve\&quot; when it serves a model endpoint. A node advertises one only once it can honour it, so an absent list means a node that has dialed in but is not ready to serve any of them yet. | [optional] 
**CpuModel** | Pointer to **string** | CPUModel is the processor as the host names it (\&quot;Apple M3 Max\&quot;), for display. | [optional] 
**Cpus** | Pointer to **int64** | CPUs is the host&#39;s logical core count. | [optional] 
**Cuda** | Pointer to **string** | Cuda is the host&#39;s CUDA toolkit version. NVIDIA hosts report it. | [optional] 
**Driver** | Pointer to **string** | Driver is the host&#39;s NVIDIA kernel driver version — distinct from Cuda, and the one that bounds which CUDA versions can run on this box. | [optional] 
**Engine** | Pointer to [**EngineAdvertisement**](EngineAdvertisement.md) | Engine is the hanzo-engine model server this node runs, when it runs one (&#x60;hanzo link --serve-engine&#x60;). Absent means the node takes jobs but serves no model endpoint. | [optional] 
**FirstSeen** | Pointer to **string** | FirstSeen is when this node first dialed in, RFC 3339 — the start of its presence record, which &#x60;hanzo unlink&#x60; ends. | [optional] 
**Gpus** | Pointer to [**[]ByoGPU**](ByoGPU.md) | GPUs are the accelerators the host found on itself. Empty is a real answer: a CPU-only machine can dial in and take non-GPU work. | [optional] 
**Hip** | Pointer to **string** | Hip is the host&#39;s HIP runtime version, the AMD counterpart to Cuda. | [optional] 
**Hostname** | Pointer to **string** | Hostname is what the host calls itself. It equals ID for any hostname already in the [a-z0-9-] alphabet, and differs when sanitizing had to change it. | [optional] 
**Id** | Pointer to **string** | ID is the node&#39;s id in the fleet — the sanitized hostname it registered under, which is also the &#x60;unit&#x60; its samples and its gpu-jobs lane key on. This is the id to use everywhere else on the compute surface. | [optional] 
**JobQueue** | Pointer to **string** | JobQueue is the tasks NAMESPACE this worker claims render jobs out of — \&quot;gpu-jobs\&quot; unless &#x60;hanzo link&#x60; was pointed at another. Within it, a job aimed at this node alone rides the task-queue value \&quot;gpu:&lt;id&gt;\&quot;. | [optional] 
**LastHeartbeat** | Pointer to **string** | LastHeartbeat is the most recent beat this node sent, RFC 3339. It is what Status is computed from, so a reader can check the judgement. | [optional] 
**Location** | Pointer to **string** | Location is always \&quot;on-prem\&quot; — a machine that dialed in has no cloud region, and inventing one would put it somewhere it is not. | [optional] 
**Memory** | Pointer to **int64** | Memory is the host&#39;s total RAM in BYTES. | [optional] 
**Os** | Pointer to **string** | Os is the host&#39;s operating system: linux, darwin or windows. | [optional] 
**Provider** | Pointer to **string** | Provider is always \&quot;byo\&quot;: this machine is the operator&#39;s, not one Hanzo provisioned. It exists so a fold into the machines/GPUs pages says which rows are rented and which are the customer&#39;s own. | [optional] 
**Rocm** | Pointer to **string** | Rocm is the host&#39;s ROCm version. AMD hosts report it; empty otherwise. | [optional] 
**Status** | Pointer to **string** | Status is \&quot;online\&quot; when the last heartbeat landed within 90s, else \&quot;offline\&quot; — so it is a fact about heartbeat freshness, not about the box being powered on. A worker that has never beaten reads offline. | [optional] 
**Version** | Pointer to **string** | Version is the &#x60;hanzo&#x60; CLI version running on the node. It is what to check when a worker is missing a field a newer registration reports. | [optional] 

## Methods

### NewByoWorker

`func NewByoWorker() *ByoWorker`

NewByoWorker instantiates a new ByoWorker object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewByoWorkerWithDefaults

`func NewByoWorkerWithDefaults() *ByoWorker`

NewByoWorkerWithDefaults instantiates a new ByoWorker object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArch

`func (o *ByoWorker) GetArch() string`

GetArch returns the Arch field if non-nil, zero value otherwise.

### GetArchOk

`func (o *ByoWorker) GetArchOk() (*string, bool)`

GetArchOk returns a tuple with the Arch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArch

`func (o *ByoWorker) SetArch(v string)`

SetArch sets Arch field to given value.

### HasArch

`func (o *ByoWorker) HasArch() bool`

HasArch returns a boolean if a field has been set.

### GetCapabilities

`func (o *ByoWorker) GetCapabilities() []string`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *ByoWorker) GetCapabilitiesOk() (*[]string, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *ByoWorker) SetCapabilities(v []string)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *ByoWorker) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### GetCpuModel

`func (o *ByoWorker) GetCpuModel() string`

GetCpuModel returns the CpuModel field if non-nil, zero value otherwise.

### GetCpuModelOk

`func (o *ByoWorker) GetCpuModelOk() (*string, bool)`

GetCpuModelOk returns a tuple with the CpuModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuModel

`func (o *ByoWorker) SetCpuModel(v string)`

SetCpuModel sets CpuModel field to given value.

### HasCpuModel

`func (o *ByoWorker) HasCpuModel() bool`

HasCpuModel returns a boolean if a field has been set.

### GetCpus

`func (o *ByoWorker) GetCpus() int64`

GetCpus returns the Cpus field if non-nil, zero value otherwise.

### GetCpusOk

`func (o *ByoWorker) GetCpusOk() (*int64, bool)`

GetCpusOk returns a tuple with the Cpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpus

`func (o *ByoWorker) SetCpus(v int64)`

SetCpus sets Cpus field to given value.

### HasCpus

`func (o *ByoWorker) HasCpus() bool`

HasCpus returns a boolean if a field has been set.

### GetCuda

`func (o *ByoWorker) GetCuda() string`

GetCuda returns the Cuda field if non-nil, zero value otherwise.

### GetCudaOk

`func (o *ByoWorker) GetCudaOk() (*string, bool)`

GetCudaOk returns a tuple with the Cuda field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCuda

`func (o *ByoWorker) SetCuda(v string)`

SetCuda sets Cuda field to given value.

### HasCuda

`func (o *ByoWorker) HasCuda() bool`

HasCuda returns a boolean if a field has been set.

### GetDriver

`func (o *ByoWorker) GetDriver() string`

GetDriver returns the Driver field if non-nil, zero value otherwise.

### GetDriverOk

`func (o *ByoWorker) GetDriverOk() (*string, bool)`

GetDriverOk returns a tuple with the Driver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriver

`func (o *ByoWorker) SetDriver(v string)`

SetDriver sets Driver field to given value.

### HasDriver

`func (o *ByoWorker) HasDriver() bool`

HasDriver returns a boolean if a field has been set.

### GetEngine

`func (o *ByoWorker) GetEngine() EngineAdvertisement`

GetEngine returns the Engine field if non-nil, zero value otherwise.

### GetEngineOk

`func (o *ByoWorker) GetEngineOk() (*EngineAdvertisement, bool)`

GetEngineOk returns a tuple with the Engine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngine

`func (o *ByoWorker) SetEngine(v EngineAdvertisement)`

SetEngine sets Engine field to given value.

### HasEngine

`func (o *ByoWorker) HasEngine() bool`

HasEngine returns a boolean if a field has been set.

### GetFirstSeen

`func (o *ByoWorker) GetFirstSeen() string`

GetFirstSeen returns the FirstSeen field if non-nil, zero value otherwise.

### GetFirstSeenOk

`func (o *ByoWorker) GetFirstSeenOk() (*string, bool)`

GetFirstSeenOk returns a tuple with the FirstSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstSeen

`func (o *ByoWorker) SetFirstSeen(v string)`

SetFirstSeen sets FirstSeen field to given value.

### HasFirstSeen

`func (o *ByoWorker) HasFirstSeen() bool`

HasFirstSeen returns a boolean if a field has been set.

### GetGpus

`func (o *ByoWorker) GetGpus() []ByoGPU`

GetGpus returns the Gpus field if non-nil, zero value otherwise.

### GetGpusOk

`func (o *ByoWorker) GetGpusOk() (*[]ByoGPU, bool)`

GetGpusOk returns a tuple with the Gpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpus

`func (o *ByoWorker) SetGpus(v []ByoGPU)`

SetGpus sets Gpus field to given value.

### HasGpus

`func (o *ByoWorker) HasGpus() bool`

HasGpus returns a boolean if a field has been set.

### GetHip

`func (o *ByoWorker) GetHip() string`

GetHip returns the Hip field if non-nil, zero value otherwise.

### GetHipOk

`func (o *ByoWorker) GetHipOk() (*string, bool)`

GetHipOk returns a tuple with the Hip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHip

`func (o *ByoWorker) SetHip(v string)`

SetHip sets Hip field to given value.

### HasHip

`func (o *ByoWorker) HasHip() bool`

HasHip returns a boolean if a field has been set.

### GetHostname

`func (o *ByoWorker) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *ByoWorker) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *ByoWorker) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *ByoWorker) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetId

`func (o *ByoWorker) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ByoWorker) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ByoWorker) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ByoWorker) HasId() bool`

HasId returns a boolean if a field has been set.

### GetJobQueue

`func (o *ByoWorker) GetJobQueue() string`

GetJobQueue returns the JobQueue field if non-nil, zero value otherwise.

### GetJobQueueOk

`func (o *ByoWorker) GetJobQueueOk() (*string, bool)`

GetJobQueueOk returns a tuple with the JobQueue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobQueue

`func (o *ByoWorker) SetJobQueue(v string)`

SetJobQueue sets JobQueue field to given value.

### HasJobQueue

`func (o *ByoWorker) HasJobQueue() bool`

HasJobQueue returns a boolean if a field has been set.

### GetLastHeartbeat

`func (o *ByoWorker) GetLastHeartbeat() string`

GetLastHeartbeat returns the LastHeartbeat field if non-nil, zero value otherwise.

### GetLastHeartbeatOk

`func (o *ByoWorker) GetLastHeartbeatOk() (*string, bool)`

GetLastHeartbeatOk returns a tuple with the LastHeartbeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastHeartbeat

`func (o *ByoWorker) SetLastHeartbeat(v string)`

SetLastHeartbeat sets LastHeartbeat field to given value.

### HasLastHeartbeat

`func (o *ByoWorker) HasLastHeartbeat() bool`

HasLastHeartbeat returns a boolean if a field has been set.

### GetLocation

`func (o *ByoWorker) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ByoWorker) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ByoWorker) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *ByoWorker) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetMemory

`func (o *ByoWorker) GetMemory() int64`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *ByoWorker) GetMemoryOk() (*int64, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *ByoWorker) SetMemory(v int64)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *ByoWorker) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetOs

`func (o *ByoWorker) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *ByoWorker) GetOsOk() (*string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *ByoWorker) SetOs(v string)`

SetOs sets Os field to given value.

### HasOs

`func (o *ByoWorker) HasOs() bool`

HasOs returns a boolean if a field has been set.

### GetProvider

`func (o *ByoWorker) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ByoWorker) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ByoWorker) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *ByoWorker) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetRocm

`func (o *ByoWorker) GetRocm() string`

GetRocm returns the Rocm field if non-nil, zero value otherwise.

### GetRocmOk

`func (o *ByoWorker) GetRocmOk() (*string, bool)`

GetRocmOk returns a tuple with the Rocm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRocm

`func (o *ByoWorker) SetRocm(v string)`

SetRocm sets Rocm field to given value.

### HasRocm

`func (o *ByoWorker) HasRocm() bool`

HasRocm returns a boolean if a field has been set.

### GetStatus

`func (o *ByoWorker) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ByoWorker) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ByoWorker) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ByoWorker) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVersion

`func (o *ByoWorker) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ByoWorker) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ByoWorker) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ByoWorker) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


