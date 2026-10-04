# SandboxLeaseIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Class** | Pointer to **string** | Class is what KIND of computer to lease, and the set is closed:   exec     a throwaway one that keeps nothing. Seconds to minutes.  dev      a coding one, with the project&#39;s own disk attached. Hours.  desktop  a dev one that also has a screen.  android  a desktop with a phone running on that screen.  Empty leases an &#x60;exec&#x60;, which is the right answer for running a program and the wrong one for working on a repository, because it keeps nothing.  An &#x60;android&#x60; needs a node that can virtualise a CPU, so it is the one class a deployment may not be able to place. Where the fleet has none, the lease succeeds and the pod stays Pending naming the device it is waiting for — which is the honest answer, because the alternative is an emulator running on an interpreted CPU and never finishing its boot. | [optional] 
**Cluster** | Pointer to **string** | Cluster names one of the org&#39;s ATTACHED clusters to run the lease on — the fleet-local name a BYO cluster was registered under. Empty runs where leases have always run, the home cluster. The named cluster must carry the sandbox namespace and the gvisor runtime class; a name the org has not attached is refused 404. | [optional] 
**Disk** | Pointer to **bool** | Disk gives a sandbox with no Project a disk of its OWN, so what it writes survives a park: small — the size its working directory always had — billed as storage for as long as it exists, and deleted with the lease. Without it a project-less sandbox keeps its files in the pod and loses them with it. A project sandbox has the project&#39;s disk and ignores this. | [optional] 
**Id** | Pointer to **string** | ID names a sandbox to RESUME, and is the id an earlier lease answered with. Empty asks for a new one. A caller that holds an id and omits it does not get a second view of the same computer, it gets a second computer. | [optional] 
**Project** | Pointer to **string** | Project names the disk to attach, and is REQUIRED for every class but &#x60;exec&#x60;.  One live sandbox per project: the disk attaches to one computer at a time, so a second lease over a project that already has one is refused by name rather than handed a silently empty disk. | [optional] 
**Runtime** | Pointer to **string** | Runtime is the isolation boundary asked for: &#x60;gvisor&#x60; shares a filesystem and holds a project volume, &#x60;kata-fc&#x60; is a microVM that boots slower and reads files faster but has no shared filesystem at all. Empty asks for the fleet&#39;s default, which is the right answer unless you are measuring.  It is a REQUEST. The owner decides, and refuses a combination it cannot honour — a volume under a runtime with no shared filesystem would write into a tmpfs and lose the bytes at exit. Read Leased.Runtime for what the sandbox actually got. | [optional] 
**TtlSec** | Pointer to **int64** | TTLSec bounds the lease in seconds. Unset takes the class default. Nothing runs forever, because a sandbox is somebody else&#39;s code on our nodes. | [optional] 

## Methods

### NewSandboxLeaseIn

`func NewSandboxLeaseIn() *SandboxLeaseIn`

NewSandboxLeaseIn instantiates a new SandboxLeaseIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSandboxLeaseInWithDefaults

`func NewSandboxLeaseInWithDefaults() *SandboxLeaseIn`

NewSandboxLeaseInWithDefaults instantiates a new SandboxLeaseIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClass

`func (o *SandboxLeaseIn) GetClass() string`

GetClass returns the Class field if non-nil, zero value otherwise.

### GetClassOk

`func (o *SandboxLeaseIn) GetClassOk() (*string, bool)`

GetClassOk returns a tuple with the Class field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClass

`func (o *SandboxLeaseIn) SetClass(v string)`

SetClass sets Class field to given value.

### HasClass

`func (o *SandboxLeaseIn) HasClass() bool`

HasClass returns a boolean if a field has been set.

### GetCluster

`func (o *SandboxLeaseIn) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *SandboxLeaseIn) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *SandboxLeaseIn) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *SandboxLeaseIn) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetDisk

`func (o *SandboxLeaseIn) GetDisk() bool`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *SandboxLeaseIn) GetDiskOk() (*bool, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *SandboxLeaseIn) SetDisk(v bool)`

SetDisk sets Disk field to given value.

### HasDisk

`func (o *SandboxLeaseIn) HasDisk() bool`

HasDisk returns a boolean if a field has been set.

### GetId

`func (o *SandboxLeaseIn) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SandboxLeaseIn) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SandboxLeaseIn) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SandboxLeaseIn) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProject

`func (o *SandboxLeaseIn) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *SandboxLeaseIn) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *SandboxLeaseIn) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *SandboxLeaseIn) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetRuntime

`func (o *SandboxLeaseIn) GetRuntime() string`

GetRuntime returns the Runtime field if non-nil, zero value otherwise.

### GetRuntimeOk

`func (o *SandboxLeaseIn) GetRuntimeOk() (*string, bool)`

GetRuntimeOk returns a tuple with the Runtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntime

`func (o *SandboxLeaseIn) SetRuntime(v string)`

SetRuntime sets Runtime field to given value.

### HasRuntime

`func (o *SandboxLeaseIn) HasRuntime() bool`

HasRuntime returns a boolean if a field has been set.

### GetTtlSec

`func (o *SandboxLeaseIn) GetTtlSec() int64`

GetTtlSec returns the TtlSec field if non-nil, zero value otherwise.

### GetTtlSecOk

`func (o *SandboxLeaseIn) GetTtlSecOk() (*int64, bool)`

GetTtlSecOk returns a tuple with the TtlSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtlSec

`func (o *SandboxLeaseIn) SetTtlSec(v int64)`

SetTtlSec sets TtlSec field to given value.

### HasTtlSec

`func (o *SandboxLeaseIn) HasTtlSec() bool`

HasTtlSec returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


