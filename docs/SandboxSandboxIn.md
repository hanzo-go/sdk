# SandboxSandboxIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Class** | Pointer to **string** | Class is what the sandbox is FOR: \&quot;exec\&quot; for a code-interpreter call, \&quot;dev\&quot; for a workspace bound to a project, \&quot;desktop\&quot; for one with a screen. It decides the image, the working directory and the isolation. | [optional] 
**Cluster** | Pointer to **string** | Cluster names one of the org&#39;s attached clusters to run the sandbox on — the fleet-local name it was registered under. Empty runs on the home cluster. The named cluster must carry the sandbox namespace and the visor runtime class; a name the org has not attached is 404. | [optional] 
**Image** | Pointer to **string** | Image overrides the image the class would pick. Honoured only for a caller the policy admits, and the sandbox that comes back names the image it GOT. | [optional] 
**Project** | Pointer to **string** | Project binds the sandbox to one of the org&#39;s projects. Required for a dev or desktop class, which are single-attach per project; an exec sandbox carries none. | [optional] 
**Runtime** | Pointer to **string** | Runtime asks for an isolation: runc, visor, kata-clh or kata-fc. It is a REQUEST, not a guarantee — the sandbox that comes back carries the runtime it was actually given, which is the field to read. | [optional] 
**TtlSec** | Pointer to **int64** | TTLSec is how long the lease runs before the reaper may take it, in seconds. Zero takes the class&#39;s own default. | [optional] 

## Methods

### NewSandboxSandboxIn

`func NewSandboxSandboxIn() *SandboxSandboxIn`

NewSandboxSandboxIn instantiates a new SandboxSandboxIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSandboxSandboxInWithDefaults

`func NewSandboxSandboxInWithDefaults() *SandboxSandboxIn`

NewSandboxSandboxInWithDefaults instantiates a new SandboxSandboxIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClass

`func (o *SandboxSandboxIn) GetClass() string`

GetClass returns the Class field if non-nil, zero value otherwise.

### GetClassOk

`func (o *SandboxSandboxIn) GetClassOk() (*string, bool)`

GetClassOk returns a tuple with the Class field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClass

`func (o *SandboxSandboxIn) SetClass(v string)`

SetClass sets Class field to given value.

### HasClass

`func (o *SandboxSandboxIn) HasClass() bool`

HasClass returns a boolean if a field has been set.

### GetCluster

`func (o *SandboxSandboxIn) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *SandboxSandboxIn) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *SandboxSandboxIn) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *SandboxSandboxIn) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetImage

`func (o *SandboxSandboxIn) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *SandboxSandboxIn) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *SandboxSandboxIn) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *SandboxSandboxIn) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetProject

`func (o *SandboxSandboxIn) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *SandboxSandboxIn) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *SandboxSandboxIn) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *SandboxSandboxIn) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetRuntime

`func (o *SandboxSandboxIn) GetRuntime() string`

GetRuntime returns the Runtime field if non-nil, zero value otherwise.

### GetRuntimeOk

`func (o *SandboxSandboxIn) GetRuntimeOk() (*string, bool)`

GetRuntimeOk returns a tuple with the Runtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntime

`func (o *SandboxSandboxIn) SetRuntime(v string)`

SetRuntime sets Runtime field to given value.

### HasRuntime

`func (o *SandboxSandboxIn) HasRuntime() bool`

HasRuntime returns a boolean if a field has been set.

### GetTtlSec

`func (o *SandboxSandboxIn) GetTtlSec() int64`

GetTtlSec returns the TtlSec field if non-nil, zero value otherwise.

### GetTtlSecOk

`func (o *SandboxSandboxIn) GetTtlSecOk() (*int64, bool)`

GetTtlSecOk returns a tuple with the TtlSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtlSec

`func (o *SandboxSandboxIn) SetTtlSec(v int64)`

SetTtlSec sets TtlSec field to given value.

### HasTtlSec

`func (o *SandboxSandboxIn) HasTtlSec() bool`

HasTtlSec returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


