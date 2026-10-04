# AgentArtifact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** | Kind is file, patch, preview, pull or deploy. A file or a patch is stored, and its bytes are read at GET /v1/agent/coding/{session}/artifacts/{name}. | [optional] 
**Name** | Pointer to **string** | Name is the artifact&#39;s handle: a changed file&#39;s path in the workspace, or &#x60;changes.patch&#x60; for the whole change; for a link, its kind and what it names. | [optional] 
**Port** | Pointer to **int64** | Port is a preview&#39;s port in the run&#39;s sandbox, which POST /v1/sandbox/{id}/preview opens while the sandbox is kept. | [optional] 
**Sandbox** | Pointer to **string** | Sandbox is the sandbox a preview is served from. | [optional] 
**Sha256** | Pointer to **string** | SHA256 is a stored artifact&#39;s digest, hex. | [optional] 
**Size** | Pointer to **int64** | Size is a stored artifact&#39;s length in bytes. | [optional] 
**Url** | Pointer to **string** | URL is where a link points: a preview&#39;s host, a pull request, a deployment. | [optional] 

## Methods

### NewAgentArtifact

`func NewAgentArtifact() *AgentArtifact`

NewAgentArtifact instantiates a new AgentArtifact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentArtifactWithDefaults

`func NewAgentArtifactWithDefaults() *AgentArtifact`

NewAgentArtifactWithDefaults instantiates a new AgentArtifact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *AgentArtifact) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *AgentArtifact) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *AgentArtifact) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *AgentArtifact) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetName

`func (o *AgentArtifact) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AgentArtifact) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AgentArtifact) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AgentArtifact) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPort

`func (o *AgentArtifact) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *AgentArtifact) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *AgentArtifact) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *AgentArtifact) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetSandbox

`func (o *AgentArtifact) GetSandbox() string`

GetSandbox returns the Sandbox field if non-nil, zero value otherwise.

### GetSandboxOk

`func (o *AgentArtifact) GetSandboxOk() (*string, bool)`

GetSandboxOk returns a tuple with the Sandbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSandbox

`func (o *AgentArtifact) SetSandbox(v string)`

SetSandbox sets Sandbox field to given value.

### HasSandbox

`func (o *AgentArtifact) HasSandbox() bool`

HasSandbox returns a boolean if a field has been set.

### GetSha256

`func (o *AgentArtifact) GetSha256() string`

GetSha256 returns the Sha256 field if non-nil, zero value otherwise.

### GetSha256Ok

`func (o *AgentArtifact) GetSha256Ok() (*string, bool)`

GetSha256Ok returns a tuple with the Sha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256

`func (o *AgentArtifact) SetSha256(v string)`

SetSha256 sets Sha256 field to given value.

### HasSha256

`func (o *AgentArtifact) HasSha256() bool`

HasSha256 returns a boolean if a field has been set.

### GetSize

`func (o *AgentArtifact) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *AgentArtifact) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *AgentArtifact) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *AgentArtifact) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetUrl

`func (o *AgentArtifact) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AgentArtifact) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AgentArtifact) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AgentArtifact) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


