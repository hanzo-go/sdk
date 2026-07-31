# CloudByoWorker

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Arch** | Pointer to **string** | Arch/CPUs/Memory are the connecting host&#39;s static CPU spec, mirrored from the registration: Arch is runtime.GOARCH (amd64 | arm64), Memory is total RAM in BYTES — the same fields a code-linked run-target carries, so the /v1/fleet board renders a linked node&#39;s arch + cores + RAM like any other unit. | [optional] 
**Capabilities** | Pointer to **[]string** | Capabilities the worker advertises (\&quot;studio.render\&quot;, \&quot;engine.serve\&quot;); Engine is present when it runs a hanzo-engine model server. Both additive + omitempty. | [optional] 
**CpuModel** | Pointer to **string** |  | [optional] 
**Cpus** | Pointer to **int32** |  | [optional] 
**Cuda** | Pointer to **string** |  | [optional] 
**Driver** | Pointer to **string** |  | [optional] 
**Engine** | Pointer to [**CloudEngineAdvertisement**](CloudEngineAdvertisement.md) |  | [optional] 
**FirstSeen** | Pointer to **string** |  | [optional] 
**Gpus** | Pointer to [**[]CloudByoGPU**](CloudByoGPU.md) |  | [optional] 
**Hip** | Pointer to **string** |  | [optional] 
**Hostname** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**JobQueue** | Pointer to **string** |  | [optional] 
**LastHeartbeat** | Pointer to **string** |  | [optional] 
**Location** | Pointer to **string** | \&quot;on-prem\&quot; (BYO has no cloud region) | [optional] 
**Memory** | Pointer to **int32** |  | [optional] 
**Os** | Pointer to **string** |  | [optional] 
**Provider** | Pointer to **string** | always \&quot;byo\&quot; | [optional] 
**Rocm** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** | online | offline | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewCloudByoWorker

`func NewCloudByoWorker() *CloudByoWorker`

NewCloudByoWorker instantiates a new CloudByoWorker object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudByoWorkerWithDefaults

`func NewCloudByoWorkerWithDefaults() *CloudByoWorker`

NewCloudByoWorkerWithDefaults instantiates a new CloudByoWorker object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArch

`func (o *CloudByoWorker) GetArch() string`

GetArch returns the Arch field if non-nil, zero value otherwise.

### GetArchOk

`func (o *CloudByoWorker) GetArchOk() (*string, bool)`

GetArchOk returns a tuple with the Arch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArch

`func (o *CloudByoWorker) SetArch(v string)`

SetArch sets Arch field to given value.

### HasArch

`func (o *CloudByoWorker) HasArch() bool`

HasArch returns a boolean if a field has been set.

### GetCapabilities

`func (o *CloudByoWorker) GetCapabilities() []string`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *CloudByoWorker) GetCapabilitiesOk() (*[]string, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *CloudByoWorker) SetCapabilities(v []string)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *CloudByoWorker) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### GetCpuModel

`func (o *CloudByoWorker) GetCpuModel() string`

GetCpuModel returns the CpuModel field if non-nil, zero value otherwise.

### GetCpuModelOk

`func (o *CloudByoWorker) GetCpuModelOk() (*string, bool)`

GetCpuModelOk returns a tuple with the CpuModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuModel

`func (o *CloudByoWorker) SetCpuModel(v string)`

SetCpuModel sets CpuModel field to given value.

### HasCpuModel

`func (o *CloudByoWorker) HasCpuModel() bool`

HasCpuModel returns a boolean if a field has been set.

### GetCpus

`func (o *CloudByoWorker) GetCpus() int32`

GetCpus returns the Cpus field if non-nil, zero value otherwise.

### GetCpusOk

`func (o *CloudByoWorker) GetCpusOk() (*int32, bool)`

GetCpusOk returns a tuple with the Cpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpus

`func (o *CloudByoWorker) SetCpus(v int32)`

SetCpus sets Cpus field to given value.

### HasCpus

`func (o *CloudByoWorker) HasCpus() bool`

HasCpus returns a boolean if a field has been set.

### GetCuda

`func (o *CloudByoWorker) GetCuda() string`

GetCuda returns the Cuda field if non-nil, zero value otherwise.

### GetCudaOk

`func (o *CloudByoWorker) GetCudaOk() (*string, bool)`

GetCudaOk returns a tuple with the Cuda field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCuda

`func (o *CloudByoWorker) SetCuda(v string)`

SetCuda sets Cuda field to given value.

### HasCuda

`func (o *CloudByoWorker) HasCuda() bool`

HasCuda returns a boolean if a field has been set.

### GetDriver

`func (o *CloudByoWorker) GetDriver() string`

GetDriver returns the Driver field if non-nil, zero value otherwise.

### GetDriverOk

`func (o *CloudByoWorker) GetDriverOk() (*string, bool)`

GetDriverOk returns a tuple with the Driver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriver

`func (o *CloudByoWorker) SetDriver(v string)`

SetDriver sets Driver field to given value.

### HasDriver

`func (o *CloudByoWorker) HasDriver() bool`

HasDriver returns a boolean if a field has been set.

### GetEngine

`func (o *CloudByoWorker) GetEngine() CloudEngineAdvertisement`

GetEngine returns the Engine field if non-nil, zero value otherwise.

### GetEngineOk

`func (o *CloudByoWorker) GetEngineOk() (*CloudEngineAdvertisement, bool)`

GetEngineOk returns a tuple with the Engine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngine

`func (o *CloudByoWorker) SetEngine(v CloudEngineAdvertisement)`

SetEngine sets Engine field to given value.

### HasEngine

`func (o *CloudByoWorker) HasEngine() bool`

HasEngine returns a boolean if a field has been set.

### GetFirstSeen

`func (o *CloudByoWorker) GetFirstSeen() string`

GetFirstSeen returns the FirstSeen field if non-nil, zero value otherwise.

### GetFirstSeenOk

`func (o *CloudByoWorker) GetFirstSeenOk() (*string, bool)`

GetFirstSeenOk returns a tuple with the FirstSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstSeen

`func (o *CloudByoWorker) SetFirstSeen(v string)`

SetFirstSeen sets FirstSeen field to given value.

### HasFirstSeen

`func (o *CloudByoWorker) HasFirstSeen() bool`

HasFirstSeen returns a boolean if a field has been set.

### GetGpus

`func (o *CloudByoWorker) GetGpus() []CloudByoGPU`

GetGpus returns the Gpus field if non-nil, zero value otherwise.

### GetGpusOk

`func (o *CloudByoWorker) GetGpusOk() (*[]CloudByoGPU, bool)`

GetGpusOk returns a tuple with the Gpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpus

`func (o *CloudByoWorker) SetGpus(v []CloudByoGPU)`

SetGpus sets Gpus field to given value.

### HasGpus

`func (o *CloudByoWorker) HasGpus() bool`

HasGpus returns a boolean if a field has been set.

### GetHip

`func (o *CloudByoWorker) GetHip() string`

GetHip returns the Hip field if non-nil, zero value otherwise.

### GetHipOk

`func (o *CloudByoWorker) GetHipOk() (*string, bool)`

GetHipOk returns a tuple with the Hip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHip

`func (o *CloudByoWorker) SetHip(v string)`

SetHip sets Hip field to given value.

### HasHip

`func (o *CloudByoWorker) HasHip() bool`

HasHip returns a boolean if a field has been set.

### GetHostname

`func (o *CloudByoWorker) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *CloudByoWorker) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *CloudByoWorker) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *CloudByoWorker) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetId

`func (o *CloudByoWorker) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudByoWorker) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudByoWorker) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudByoWorker) HasId() bool`

HasId returns a boolean if a field has been set.

### GetJobQueue

`func (o *CloudByoWorker) GetJobQueue() string`

GetJobQueue returns the JobQueue field if non-nil, zero value otherwise.

### GetJobQueueOk

`func (o *CloudByoWorker) GetJobQueueOk() (*string, bool)`

GetJobQueueOk returns a tuple with the JobQueue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobQueue

`func (o *CloudByoWorker) SetJobQueue(v string)`

SetJobQueue sets JobQueue field to given value.

### HasJobQueue

`func (o *CloudByoWorker) HasJobQueue() bool`

HasJobQueue returns a boolean if a field has been set.

### GetLastHeartbeat

`func (o *CloudByoWorker) GetLastHeartbeat() string`

GetLastHeartbeat returns the LastHeartbeat field if non-nil, zero value otherwise.

### GetLastHeartbeatOk

`func (o *CloudByoWorker) GetLastHeartbeatOk() (*string, bool)`

GetLastHeartbeatOk returns a tuple with the LastHeartbeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastHeartbeat

`func (o *CloudByoWorker) SetLastHeartbeat(v string)`

SetLastHeartbeat sets LastHeartbeat field to given value.

### HasLastHeartbeat

`func (o *CloudByoWorker) HasLastHeartbeat() bool`

HasLastHeartbeat returns a boolean if a field has been set.

### GetLocation

`func (o *CloudByoWorker) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *CloudByoWorker) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *CloudByoWorker) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *CloudByoWorker) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetMemory

`func (o *CloudByoWorker) GetMemory() int32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *CloudByoWorker) GetMemoryOk() (*int32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *CloudByoWorker) SetMemory(v int32)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *CloudByoWorker) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetOs

`func (o *CloudByoWorker) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *CloudByoWorker) GetOsOk() (*string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *CloudByoWorker) SetOs(v string)`

SetOs sets Os field to given value.

### HasOs

`func (o *CloudByoWorker) HasOs() bool`

HasOs returns a boolean if a field has been set.

### GetProvider

`func (o *CloudByoWorker) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *CloudByoWorker) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *CloudByoWorker) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *CloudByoWorker) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetRocm

`func (o *CloudByoWorker) GetRocm() string`

GetRocm returns the Rocm field if non-nil, zero value otherwise.

### GetRocmOk

`func (o *CloudByoWorker) GetRocmOk() (*string, bool)`

GetRocmOk returns a tuple with the Rocm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRocm

`func (o *CloudByoWorker) SetRocm(v string)`

SetRocm sets Rocm field to given value.

### HasRocm

`func (o *CloudByoWorker) HasRocm() bool`

HasRocm returns a boolean if a field has been set.

### GetStatus

`func (o *CloudByoWorker) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudByoWorker) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudByoWorker) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudByoWorker) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVersion

`func (o *CloudByoWorker) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CloudByoWorker) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CloudByoWorker) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CloudByoWorker) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


