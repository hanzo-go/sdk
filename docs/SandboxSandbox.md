# SandboxSandbox

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actor** | Pointer to **string** | Actor is the subject who LEASED this sandbox — the validated caller of the act that took the lease, never a value a request supplied. A sandbox is its lessee&#39;s: it holds their session, and a coding run&#39;s holds a write key to a repository and the codebase&#39;s secrets. So the org gates which store a row is in, and this gates who in the org may act on it (holder.holds): the lessee, or an admin of the org, or a SuperAdmin. Empty is a lease taken by the org&#39;s own background work, which no person holds. | [optional] 
**Class** | Pointer to **string** | Class is what the sandbox is FOR, and it decides the image, the working directory and the isolation: \&quot;exec\&quot; for a code-interpreter call (workdir /mnt/data, no project, bounded per org), \&quot;dev\&quot; for a workspace bound to a project (workdir /work, single-attach), \&quot;desktop\&quot; for one with a screen. | [optional] 
**Cluster** | Pointer to **string** | Cluster is the attached cluster this sandbox runs on — the fleet-local name the lease named — or empty for the home cluster. Immutable for the life of the lease, like the pod it locates: every later call into the sandbox reads it to reach the right apiserver. | [optional] 
**ConnectedAt** | Pointer to **int64** | ConnectedAt is when somebody was last known to have this sandbox&#39;s project OPEN, Unix seconds. It is a fact with an EXPIRY rather than a flag: a watcher restamps it every beat of its stream, and it goes stale on its own when the stream dies, so nothing has to be turned off by a process that may not be there any more. The reaper reads it to choose WHICH idle allowance applies — see lifecycle.go.  Zero means nobody has said so, which puts the sandbox on the short clock. | [optional] 
**CreatedAt** | Pointer to **int64** | CreatedAt is when the lease was taken, Unix seconds — by the lease that made the sandbox, and again by each resume, which gives a parked sandbox a new pod on a new lease. The ceiling on a sandbox&#39;s life runs from it. | [optional] 
**Error** | Pointer to **string** | Error is why the sandbox could not come up, in plain words. Present only with status \&quot;error\&quot;, and it is the field to read rather than inferring a cause from the absence of a pod. | [optional] 
**ExpiresAt** | Pointer to **int64** | ExpiresAt is when the lease ends, Unix seconds. Past it the reaper may take the sandbox at any time; it is a deadline, not a guarantee of survival until then, since an idle sandbox goes sooner. | [optional] 
**Id** | Pointer to **string** | ID is the sandbox&#39;s server-minted handle and what every operation addresses it by. The caller does not choose it. | [optional] 
**Image** | Pointer to **string** | Image is the container image this sandbox is actually running — the one the class chose, or an override the policy admitted. It is what ran, not what was asked for. | [optional] 
**Kind** | Pointer to **string** | Kind is the resource family this row belongs to. Always \&quot;sandbox\&quot; here; it exists because the store this shares is keyed across kinds. | [optional] 
**LastUsedAt** | Pointer to **int64** | LastUsedAt is when the sandbox last did work, Unix seconds. The reaper reads it: a sandbox idle past the idle window is reclaimed even inside its TTL, because an idle lease is capacity nobody is using. | [optional] 
**Org** | Pointer to **string** | Org is the org that holds the lease — the validated caller&#39;s, never a value a request supplied. It is also the store&#39;s key, so a sandbox is not merely filtered out of another org&#39;s answers; it is unreachable from them. | [optional] 
**Project** | Pointer to **string** | Project is the project this sandbox is bound to. A dev or desktop sandbox has one and is SINGLE-ATTACH under it, so asking twice resumes rather than leasing a second; an exec sandbox has none. | [optional] 
**Runtime** | Pointer to **string** | Runtime is the isolation boundary this sandbox GOT, which is not always the one it asked for: a caller states a preference and runtimeFor answers with what the sandbox can actually have. Reported so a person comparing two runtimes is comparing the runtimes they got rather than the ones they typed — the difference between those two is the whole reason to record it.  Empty means the node&#39;s default runtime, which is a real answer and not a missing one.  This is not a copy that can go stale. runtimeClassName is IMMUTABLE on a pod, a sandbox&#39;s pod is created once and never recreated (restartPolicy Never, no pool), and its name is never reused — so for as long as the pod this row names exists, it is running this runtime. The alternative, asking the apiserver on every read, buys nothing and costs a round trip per row. | [optional] 
**Status** | Pointer to **string** | Status is where the sandbox is in its life: \&quot;pending\&quot; while the pod is coming up, \&quot;running\&quot; once it can take work, \&quot;error\&quot; when it cannot, and \&quot;parked\&quot; when its pod is stopped and its disk kept, until a resume. Only a running sandbox takes an exec or mints an interactive ticket. | [optional] 
**Volume** | Pointer to **string** | Volume is the persistent volume attached to the sandbox, when it has one. A project sandbox keeps its work across leases through the project&#39;s disk; a sandbox with no project has one only if its lease asked for a disk of its own, which keeps its files across a park and is deleted with the lease. Without either, everything is lost with the pod. | [optional] 

## Methods

### NewSandboxSandbox

`func NewSandboxSandbox() *SandboxSandbox`

NewSandboxSandbox instantiates a new SandboxSandbox object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSandboxSandboxWithDefaults

`func NewSandboxSandboxWithDefaults() *SandboxSandbox`

NewSandboxSandboxWithDefaults instantiates a new SandboxSandbox object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActor

`func (o *SandboxSandbox) GetActor() string`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *SandboxSandbox) GetActorOk() (*string, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *SandboxSandbox) SetActor(v string)`

SetActor sets Actor field to given value.

### HasActor

`func (o *SandboxSandbox) HasActor() bool`

HasActor returns a boolean if a field has been set.

### GetClass

`func (o *SandboxSandbox) GetClass() string`

GetClass returns the Class field if non-nil, zero value otherwise.

### GetClassOk

`func (o *SandboxSandbox) GetClassOk() (*string, bool)`

GetClassOk returns a tuple with the Class field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClass

`func (o *SandboxSandbox) SetClass(v string)`

SetClass sets Class field to given value.

### HasClass

`func (o *SandboxSandbox) HasClass() bool`

HasClass returns a boolean if a field has been set.

### GetCluster

`func (o *SandboxSandbox) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *SandboxSandbox) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *SandboxSandbox) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *SandboxSandbox) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetConnectedAt

`func (o *SandboxSandbox) GetConnectedAt() int64`

GetConnectedAt returns the ConnectedAt field if non-nil, zero value otherwise.

### GetConnectedAtOk

`func (o *SandboxSandbox) GetConnectedAtOk() (*int64, bool)`

GetConnectedAtOk returns a tuple with the ConnectedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedAt

`func (o *SandboxSandbox) SetConnectedAt(v int64)`

SetConnectedAt sets ConnectedAt field to given value.

### HasConnectedAt

`func (o *SandboxSandbox) HasConnectedAt() bool`

HasConnectedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SandboxSandbox) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SandboxSandbox) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SandboxSandbox) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SandboxSandbox) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetError

`func (o *SandboxSandbox) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *SandboxSandbox) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *SandboxSandbox) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *SandboxSandbox) HasError() bool`

HasError returns a boolean if a field has been set.

### GetExpiresAt

`func (o *SandboxSandbox) GetExpiresAt() int64`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *SandboxSandbox) GetExpiresAtOk() (*int64, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *SandboxSandbox) SetExpiresAt(v int64)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *SandboxSandbox) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetId

`func (o *SandboxSandbox) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SandboxSandbox) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SandboxSandbox) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SandboxSandbox) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImage

`func (o *SandboxSandbox) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *SandboxSandbox) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *SandboxSandbox) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *SandboxSandbox) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetKind

`func (o *SandboxSandbox) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *SandboxSandbox) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *SandboxSandbox) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *SandboxSandbox) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLastUsedAt

`func (o *SandboxSandbox) GetLastUsedAt() int64`

GetLastUsedAt returns the LastUsedAt field if non-nil, zero value otherwise.

### GetLastUsedAtOk

`func (o *SandboxSandbox) GetLastUsedAtOk() (*int64, bool)`

GetLastUsedAtOk returns a tuple with the LastUsedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsedAt

`func (o *SandboxSandbox) SetLastUsedAt(v int64)`

SetLastUsedAt sets LastUsedAt field to given value.

### HasLastUsedAt

`func (o *SandboxSandbox) HasLastUsedAt() bool`

HasLastUsedAt returns a boolean if a field has been set.

### GetOrg

`func (o *SandboxSandbox) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *SandboxSandbox) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *SandboxSandbox) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *SandboxSandbox) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetProject

`func (o *SandboxSandbox) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *SandboxSandbox) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *SandboxSandbox) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *SandboxSandbox) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetRuntime

`func (o *SandboxSandbox) GetRuntime() string`

GetRuntime returns the Runtime field if non-nil, zero value otherwise.

### GetRuntimeOk

`func (o *SandboxSandbox) GetRuntimeOk() (*string, bool)`

GetRuntimeOk returns a tuple with the Runtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntime

`func (o *SandboxSandbox) SetRuntime(v string)`

SetRuntime sets Runtime field to given value.

### HasRuntime

`func (o *SandboxSandbox) HasRuntime() bool`

HasRuntime returns a boolean if a field has been set.

### GetStatus

`func (o *SandboxSandbox) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SandboxSandbox) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SandboxSandbox) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SandboxSandbox) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVolume

`func (o *SandboxSandbox) GetVolume() string`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *SandboxSandbox) GetVolumeOk() (*string, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *SandboxSandbox) SetVolume(v string)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *SandboxSandbox) HasVolume() bool`

HasVolume returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


