# SandboxLeased

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Class** | Pointer to **string** | Class is what was actually leased, from the closed set LeaseIn.Class names: exec | dev | desktop | android. A request that named none leased an &#x60;exec&#x60;, so this is where a caller learns which kind of computer it is holding, and it is what Workdir below follows from. | [optional] 
**Cluster** | Pointer to **string** | Cluster is the attached cluster this sandbox runs on, when one was named. Empty is the home cluster. | [optional] 
**Id** | Pointer to **string** | ID names this computer for every later call — run, read, write, stop and end all take it, and a LeaseIn carrying it resumes THIS sandbox instead of leasing a second one. Minted here; a caller cannot choose it, and a resumed lease that had expired comes back under a new one. | [optional] 
**Runtime** | Pointer to **string** | Runtime is the boundary this sandbox GOT, which need not be the one asked for — carried for the same reason Workdir is, that it is a fact only the owner knows and a caller assuming it would be holding a second copy. Empty is the node&#39;s default runtime, and a real answer. | [optional] 
**Status** | Pointer to **string** | Status is where the pod stands, from the store&#39;s three: pending | running | error. A lease that ANSWERS has already waited for the pod, so this reads &#x60;running&#x60; — a start that failed is a 503 and no sandbox at all. Read it anyway: exec refuses a sandbox that is not running, so anything else here is the reason the next call will not work. | [optional] 
**Workdir** | Pointer to **string** | Workdir is the absolute directory this sandbox keeps files in, and what a relative path in a later read, write or run resolves against — /work for dev, desktop and android (the project volume&#39;s mount point), /mnt/data for exec (the artifact directory the code tool tells the model to write to). A path that climbs above it is refused rather than rewritten. | [optional] 

## Methods

### NewSandboxLeased

`func NewSandboxLeased() *SandboxLeased`

NewSandboxLeased instantiates a new SandboxLeased object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSandboxLeasedWithDefaults

`func NewSandboxLeasedWithDefaults() *SandboxLeased`

NewSandboxLeasedWithDefaults instantiates a new SandboxLeased object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClass

`func (o *SandboxLeased) GetClass() string`

GetClass returns the Class field if non-nil, zero value otherwise.

### GetClassOk

`func (o *SandboxLeased) GetClassOk() (*string, bool)`

GetClassOk returns a tuple with the Class field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClass

`func (o *SandboxLeased) SetClass(v string)`

SetClass sets Class field to given value.

### HasClass

`func (o *SandboxLeased) HasClass() bool`

HasClass returns a boolean if a field has been set.

### GetCluster

`func (o *SandboxLeased) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *SandboxLeased) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *SandboxLeased) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *SandboxLeased) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetId

`func (o *SandboxLeased) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SandboxLeased) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SandboxLeased) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SandboxLeased) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRuntime

`func (o *SandboxLeased) GetRuntime() string`

GetRuntime returns the Runtime field if non-nil, zero value otherwise.

### GetRuntimeOk

`func (o *SandboxLeased) GetRuntimeOk() (*string, bool)`

GetRuntimeOk returns a tuple with the Runtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntime

`func (o *SandboxLeased) SetRuntime(v string)`

SetRuntime sets Runtime field to given value.

### HasRuntime

`func (o *SandboxLeased) HasRuntime() bool`

HasRuntime returns a boolean if a field has been set.

### GetStatus

`func (o *SandboxLeased) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SandboxLeased) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SandboxLeased) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SandboxLeased) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetWorkdir

`func (o *SandboxLeased) GetWorkdir() string`

GetWorkdir returns the Workdir field if non-nil, zero value otherwise.

### GetWorkdirOk

`func (o *SandboxLeased) GetWorkdirOk() (*string, bool)`

GetWorkdirOk returns a tuple with the Workdir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkdir

`func (o *SandboxLeased) SetWorkdir(v string)`

SetWorkdir sets Workdir field to given value.

### HasWorkdir

`func (o *SandboxLeased) HasWorkdir() bool`

HasWorkdir returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


