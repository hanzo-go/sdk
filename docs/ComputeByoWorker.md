# ComputeByoWorker

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Arch** | Pointer to **string** | Arch/CPUs/Memory are the connecting host&#39;s static CPU spec, mirrored from the registration: Arch is runtime.GOARCH (amd64 | arm64), Memory is total RAM in BYTES — the same fields a code-linked run-target carries, so the /v1/compute/fleet board renders a linked node&#39;s arch + cores + RAM like any other unit. | [optional] 
**Capabilities** | Pointer to **[]string** | Capabilities is what this worker offers the org: \&quot;studio.render\&quot; when the node can render, \&quot;engine.serve\&quot; when it serves a model endpoint. A node advertises one only once it can honour it, so an absent list means a node that has dialed in but is not ready to serve any of them yet. | [optional] 
**CpuModel** | Pointer to **string** | CPUModel is the processor as the host names it (\&quot;Apple M3 Max\&quot;), for display. | [optional] 
**Cpus** | Pointer to **int64** | CPUs is the host&#39;s logical core count. | [optional] 
**Cuda** | Pointer to **string** | Cuda is the host&#39;s CUDA toolkit version. NVIDIA hosts report it. | [optional] 
**Driver** | Pointer to **string** | Driver is the host&#39;s NVIDIA kernel driver version — distinct from Cuda, and the one that bounds which CUDA versions can run on this box. | [optional] 
**Engine** | Pointer to [**ComputeEngineAdvertisement**](ComputeEngineAdvertisement.md) | Engine is the hanzo-engine model server this node runs, when it runs one (&#x60;hanzo link --serve-engine&#x60;). Absent means the node takes jobs but serves no model endpoint. | [optional] 
**FirstSeen** | Pointer to **string** | FirstSeen is when this node first dialed in, RFC 3339 — the start of its presence record, which &#x60;hanzo unlink&#x60; ends. | [optional] 
**Gpus** | Pointer to [**[]ComputeByoGPU**](ComputeByoGPU.md) | GPUs are the accelerators the host found on itself. Empty is a real answer: a CPU-only machine can dial in and take non-GPU work. | [optional] 
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

### NewComputeByoWorker

`func NewComputeByoWorker() *ComputeByoWorker`

NewComputeByoWorker instantiates a new ComputeByoWorker object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputeByoWorkerWithDefaults

`func NewComputeByoWorkerWithDefaults() *ComputeByoWorker`

NewComputeByoWorkerWithDefaults instantiates a new ComputeByoWorker object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArch

`func (o *ComputeByoWorker) GetArch() string`

GetArch returns the Arch field if non-nil, zero value otherwise.

### GetArchOk

`func (o *ComputeByoWorker) GetArchOk() (*string, bool)`

GetArchOk returns a tuple with the Arch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArch

`func (o *ComputeByoWorker) SetArch(v string)`

SetArch sets Arch field to given value.

### HasArch

`func (o *ComputeByoWorker) HasArch() bool`

HasArch returns a boolean if a field has been set.

### GetCapabilities

`func (o *ComputeByoWorker) GetCapabilities() []string`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *ComputeByoWorker) GetCapabilitiesOk() (*[]string, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *ComputeByoWorker) SetCapabilities(v []string)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *ComputeByoWorker) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### GetCpuModel

`func (o *ComputeByoWorker) GetCpuModel() string`

GetCpuModel returns the CpuModel field if non-nil, zero value otherwise.

### GetCpuModelOk

`func (o *ComputeByoWorker) GetCpuModelOk() (*string, bool)`

GetCpuModelOk returns a tuple with the CpuModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuModel

`func (o *ComputeByoWorker) SetCpuModel(v string)`

SetCpuModel sets CpuModel field to given value.

### HasCpuModel

`func (o *ComputeByoWorker) HasCpuModel() bool`

HasCpuModel returns a boolean if a field has been set.

### GetCpus

`func (o *ComputeByoWorker) GetCpus() int64`

GetCpus returns the Cpus field if non-nil, zero value otherwise.

### GetCpusOk

`func (o *ComputeByoWorker) GetCpusOk() (*int64, bool)`

GetCpusOk returns a tuple with the Cpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpus

`func (o *ComputeByoWorker) SetCpus(v int64)`

SetCpus sets Cpus field to given value.

### HasCpus

`func (o *ComputeByoWorker) HasCpus() bool`

HasCpus returns a boolean if a field has been set.

### GetCuda

`func (o *ComputeByoWorker) GetCuda() string`

GetCuda returns the Cuda field if non-nil, zero value otherwise.

### GetCudaOk

`func (o *ComputeByoWorker) GetCudaOk() (*string, bool)`

GetCudaOk returns a tuple with the Cuda field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCuda

`func (o *ComputeByoWorker) SetCuda(v string)`

SetCuda sets Cuda field to given value.

### HasCuda

`func (o *ComputeByoWorker) HasCuda() bool`

HasCuda returns a boolean if a field has been set.

### GetDriver

`func (o *ComputeByoWorker) GetDriver() string`

GetDriver returns the Driver field if non-nil, zero value otherwise.

### GetDriverOk

`func (o *ComputeByoWorker) GetDriverOk() (*string, bool)`

GetDriverOk returns a tuple with the Driver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriver

`func (o *ComputeByoWorker) SetDriver(v string)`

SetDriver sets Driver field to given value.

### HasDriver

`func (o *ComputeByoWorker) HasDriver() bool`

HasDriver returns a boolean if a field has been set.

### GetEngine

`func (o *ComputeByoWorker) GetEngine() ComputeEngineAdvertisement`

GetEngine returns the Engine field if non-nil, zero value otherwise.

### GetEngineOk

`func (o *ComputeByoWorker) GetEngineOk() (*ComputeEngineAdvertisement, bool)`

GetEngineOk returns a tuple with the Engine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngine

`func (o *ComputeByoWorker) SetEngine(v ComputeEngineAdvertisement)`

SetEngine sets Engine field to given value.

### HasEngine

`func (o *ComputeByoWorker) HasEngine() bool`

HasEngine returns a boolean if a field has been set.

### GetFirstSeen

`func (o *ComputeByoWorker) GetFirstSeen() string`

GetFirstSeen returns the FirstSeen field if non-nil, zero value otherwise.

### GetFirstSeenOk

`func (o *ComputeByoWorker) GetFirstSeenOk() (*string, bool)`

GetFirstSeenOk returns a tuple with the FirstSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstSeen

`func (o *ComputeByoWorker) SetFirstSeen(v string)`

SetFirstSeen sets FirstSeen field to given value.

### HasFirstSeen

`func (o *ComputeByoWorker) HasFirstSeen() bool`

HasFirstSeen returns a boolean if a field has been set.

### GetGpus

`func (o *ComputeByoWorker) GetGpus() []ComputeByoGPU`

GetGpus returns the Gpus field if non-nil, zero value otherwise.

### GetGpusOk

`func (o *ComputeByoWorker) GetGpusOk() (*[]ComputeByoGPU, bool)`

GetGpusOk returns a tuple with the Gpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpus

`func (o *ComputeByoWorker) SetGpus(v []ComputeByoGPU)`

SetGpus sets Gpus field to given value.

### HasGpus

`func (o *ComputeByoWorker) HasGpus() bool`

HasGpus returns a boolean if a field has been set.

### GetHip

`func (o *ComputeByoWorker) GetHip() string`

GetHip returns the Hip field if non-nil, zero value otherwise.

### GetHipOk

`func (o *ComputeByoWorker) GetHipOk() (*string, bool)`

GetHipOk returns a tuple with the Hip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHip

`func (o *ComputeByoWorker) SetHip(v string)`

SetHip sets Hip field to given value.

### HasHip

`func (o *ComputeByoWorker) HasHip() bool`

HasHip returns a boolean if a field has been set.

### GetHostname

`func (o *ComputeByoWorker) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *ComputeByoWorker) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *ComputeByoWorker) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *ComputeByoWorker) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetId

`func (o *ComputeByoWorker) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ComputeByoWorker) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ComputeByoWorker) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ComputeByoWorker) HasId() bool`

HasId returns a boolean if a field has been set.

### GetJobQueue

`func (o *ComputeByoWorker) GetJobQueue() string`

GetJobQueue returns the JobQueue field if non-nil, zero value otherwise.

### GetJobQueueOk

`func (o *ComputeByoWorker) GetJobQueueOk() (*string, bool)`

GetJobQueueOk returns a tuple with the JobQueue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobQueue

`func (o *ComputeByoWorker) SetJobQueue(v string)`

SetJobQueue sets JobQueue field to given value.

### HasJobQueue

`func (o *ComputeByoWorker) HasJobQueue() bool`

HasJobQueue returns a boolean if a field has been set.

### GetLastHeartbeat

`func (o *ComputeByoWorker) GetLastHeartbeat() string`

GetLastHeartbeat returns the LastHeartbeat field if non-nil, zero value otherwise.

### GetLastHeartbeatOk

`func (o *ComputeByoWorker) GetLastHeartbeatOk() (*string, bool)`

GetLastHeartbeatOk returns a tuple with the LastHeartbeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastHeartbeat

`func (o *ComputeByoWorker) SetLastHeartbeat(v string)`

SetLastHeartbeat sets LastHeartbeat field to given value.

### HasLastHeartbeat

`func (o *ComputeByoWorker) HasLastHeartbeat() bool`

HasLastHeartbeat returns a boolean if a field has been set.

### GetLocation

`func (o *ComputeByoWorker) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ComputeByoWorker) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ComputeByoWorker) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *ComputeByoWorker) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetMemory

`func (o *ComputeByoWorker) GetMemory() int64`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *ComputeByoWorker) GetMemoryOk() (*int64, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *ComputeByoWorker) SetMemory(v int64)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *ComputeByoWorker) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetOs

`func (o *ComputeByoWorker) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *ComputeByoWorker) GetOsOk() (*string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *ComputeByoWorker) SetOs(v string)`

SetOs sets Os field to given value.

### HasOs

`func (o *ComputeByoWorker) HasOs() bool`

HasOs returns a boolean if a field has been set.

### GetProvider

`func (o *ComputeByoWorker) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ComputeByoWorker) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ComputeByoWorker) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *ComputeByoWorker) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetRocm

`func (o *ComputeByoWorker) GetRocm() string`

GetRocm returns the Rocm field if non-nil, zero value otherwise.

### GetRocmOk

`func (o *ComputeByoWorker) GetRocmOk() (*string, bool)`

GetRocmOk returns a tuple with the Rocm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRocm

`func (o *ComputeByoWorker) SetRocm(v string)`

SetRocm sets Rocm field to given value.

### HasRocm

`func (o *ComputeByoWorker) HasRocm() bool`

HasRocm returns a boolean if a field has been set.

### GetStatus

`func (o *ComputeByoWorker) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ComputeByoWorker) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ComputeByoWorker) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ComputeByoWorker) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVersion

`func (o *ComputeByoWorker) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ComputeByoWorker) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ComputeByoWorker) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ComputeByoWorker) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


