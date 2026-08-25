# Sandbox

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Class** | Pointer to **string** | Class is what the sandbox is FOR, and it decides the image, the working directory and the isolation: \&quot;exec\&quot; for a code-interpreter call (workdir /mnt/data, no project, bounded per org), \&quot;dev\&quot; for a workspace bound to a project (workdir /work, single-attach), \&quot;desktop\&quot; for one with a screen. | [optional] 
**ConnectedAt** | Pointer to **int32** | ConnectedAt is when somebody was last known to have this sandbox&#39;s project OPEN, Unix seconds. It is a fact with an EXPIRY rather than a flag: a watcher restamps it every beat of its stream, and it goes stale on its own when the stream dies, so nothing has to be turned off by a process that may not be there any more. The reaper reads it to choose WHICH idle allowance applies — see lifecycle.go.  Zero means nobody has said so, which puts the sandbox on the short clock. | [optional] 
**CreatedAt** | Pointer to **int32** | CreatedAt is when the lease was first taken, Unix seconds. | [optional] 
**Error** | Pointer to **string** | Error is why the sandbox could not come up, in plain words. Present only with status \&quot;error\&quot;, and it is the field to read rather than inferring a cause from the absence of a pod. | [optional] 
**ExpiresAt** | Pointer to **int32** | ExpiresAt is when the lease ends, Unix seconds. Past it the reaper may take the sandbox at any time; it is a deadline, not a guarantee of survival until then, since an idle sandbox goes sooner. | [optional] 
**Id** | Pointer to **string** | ID is the sandbox&#39;s server-minted handle and what every operation addresses it by. The caller does not choose it. | [optional] 
**Image** | Pointer to **string** | Image is the container image this sandbox is actually running — the one the class chose, or an override the policy admitted. It is what ran, not what was asked for. | [optional] 
**Kind** | Pointer to **string** | Kind is the resource family this row belongs to. Always \&quot;sandbox\&quot; here; it exists because the store this shares is keyed across kinds. | [optional] 
**LastUsedAt** | Pointer to **int32** | LastUsedAt is when the sandbox last did work, Unix seconds. The reaper reads it: a sandbox idle past the idle window is reclaimed even inside its TTL, because an idle lease is capacity nobody is using. | [optional] 
**Org** | Pointer to **string** | Org is the org that holds the lease — the validated caller&#39;s, never a value a request supplied. It is also the store&#39;s key, so a sandbox is not merely filtered out of another org&#39;s answers; it is unreachable from them. | [optional] 
**Project** | Pointer to **string** | Project is the project this sandbox is bound to. A dev or desktop sandbox has one and is SINGLE-ATTACH under it, so asking twice resumes rather than leasing a second; an exec sandbox has none. | [optional] 
**Runtime** | Pointer to **string** | Runtime is the isolation boundary this sandbox GOT, which is not always the one it asked for: a caller states a preference and runtimeFor answers with what the sandbox can actually have. Reported so a person comparing two runtimes is comparing the runtimes they got rather than the ones they typed — the difference between those two is the whole reason to record it.  Empty means the node&#39;s default runtime, which is a real answer and not a missing one.  This is not a copy that can go stale. runtimeClassName is IMMUTABLE on a pod, a sandbox&#39;s pod is created once and never recreated (restartPolicy Never, no pool), and its name is never reused — so for as long as the pod this row names exists, it is running this runtime. The alternative, asking the apiserver on every read, buys nothing and costs a round trip per row. | [optional] 
**Status** | Pointer to **string** | Status is where the sandbox is in its life: \&quot;pending\&quot; while the pod is coming up, \&quot;running\&quot; once it can take work, \&quot;error\&quot; when it cannot. Only a running sandbox takes an exec or mints an interactive ticket. | [optional] 
**Volume** | Pointer to **string** | Volume is the persistent volume attached to the sandbox, when it has one. A dev sandbox keeps its work across leases through it; an exec sandbox has none and loses everything outside /mnt/data when the lease ends. | [optional] 

## Methods

### NewSandbox

`func NewSandbox() *Sandbox`

NewSandbox instantiates a new Sandbox object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSandboxWithDefaults

`func NewSandboxWithDefaults() *Sandbox`

NewSandboxWithDefaults instantiates a new Sandbox object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClass

`func (o *Sandbox) GetClass() string`

GetClass returns the Class field if non-nil, zero value otherwise.

### GetClassOk

`func (o *Sandbox) GetClassOk() (*string, bool)`

GetClassOk returns a tuple with the Class field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClass

`func (o *Sandbox) SetClass(v string)`

SetClass sets Class field to given value.

### HasClass

`func (o *Sandbox) HasClass() bool`

HasClass returns a boolean if a field has been set.

### GetConnectedAt

`func (o *Sandbox) GetConnectedAt() int32`

GetConnectedAt returns the ConnectedAt field if non-nil, zero value otherwise.

### GetConnectedAtOk

`func (o *Sandbox) GetConnectedAtOk() (*int32, bool)`

GetConnectedAtOk returns a tuple with the ConnectedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedAt

`func (o *Sandbox) SetConnectedAt(v int32)`

SetConnectedAt sets ConnectedAt field to given value.

### HasConnectedAt

`func (o *Sandbox) HasConnectedAt() bool`

HasConnectedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Sandbox) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Sandbox) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Sandbox) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Sandbox) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetError

`func (o *Sandbox) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *Sandbox) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *Sandbox) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *Sandbox) HasError() bool`

HasError returns a boolean if a field has been set.

### GetExpiresAt

`func (o *Sandbox) GetExpiresAt() int32`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *Sandbox) GetExpiresAtOk() (*int32, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *Sandbox) SetExpiresAt(v int32)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *Sandbox) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetId

`func (o *Sandbox) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Sandbox) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Sandbox) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Sandbox) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImage

`func (o *Sandbox) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *Sandbox) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *Sandbox) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *Sandbox) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetKind

`func (o *Sandbox) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *Sandbox) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *Sandbox) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *Sandbox) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLastUsedAt

`func (o *Sandbox) GetLastUsedAt() int32`

GetLastUsedAt returns the LastUsedAt field if non-nil, zero value otherwise.

### GetLastUsedAtOk

`func (o *Sandbox) GetLastUsedAtOk() (*int32, bool)`

GetLastUsedAtOk returns a tuple with the LastUsedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsedAt

`func (o *Sandbox) SetLastUsedAt(v int32)`

SetLastUsedAt sets LastUsedAt field to given value.

### HasLastUsedAt

`func (o *Sandbox) HasLastUsedAt() bool`

HasLastUsedAt returns a boolean if a field has been set.

### GetOrg

`func (o *Sandbox) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *Sandbox) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *Sandbox) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *Sandbox) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetProject

`func (o *Sandbox) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *Sandbox) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *Sandbox) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *Sandbox) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetRuntime

`func (o *Sandbox) GetRuntime() string`

GetRuntime returns the Runtime field if non-nil, zero value otherwise.

### GetRuntimeOk

`func (o *Sandbox) GetRuntimeOk() (*string, bool)`

GetRuntimeOk returns a tuple with the Runtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntime

`func (o *Sandbox) SetRuntime(v string)`

SetRuntime sets Runtime field to given value.

### HasRuntime

`func (o *Sandbox) HasRuntime() bool`

HasRuntime returns a boolean if a field has been set.

### GetStatus

`func (o *Sandbox) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Sandbox) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Sandbox) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Sandbox) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVolume

`func (o *Sandbox) GetVolume() string`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *Sandbox) GetVolumeOk() (*string, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *Sandbox) SetVolume(v string)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *Sandbox) HasVolume() bool`

HasVolume returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


