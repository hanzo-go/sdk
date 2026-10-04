# PlatformAppOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuildType** | Pointer to **string** | BuildType is how a git app builds: &#x60;pack&#x60;, the zero-config default that detects the project, or &#x60;dockerfile&#x60;. An image app carries &#x60;image&#x60;, which means it never builds. | [optional] 
**CreatedAt** | Pointer to **int64** | CreatedAt is when the app was created, unix seconds. | [optional] 
**CurrentDeploymentId** | Pointer to **string** | CurrentDeploymentID is the deployment that is live — the pointer a deploy advances monotonically by version, so it never regresses to an older one. Empty until the first deploy reaches the cluster. | [optional] 
**Description** | Pointer to **string** | Description is free text about what the app is. Nothing derives from it. | [optional] 
**Dockerfile** | Pointer to **string** | Dockerfile is the path inside the repo to build from, for buildType &#x60;dockerfile&#x60;. The build path keys off its presence, and it is validated at create against the same allowlist the privileged build enforces. | [optional] 
**Domains** | Pointer to **[]string** | Domains are the ingress hosts rendered into the app&#39;s CR, its own &#x60;&lt;slug&gt;.&lt;org&gt;.&lt;sites host&gt;&#x60; first. That one is seeded at create and cannot be removed; a custom host joins only after add-domain and DNS verification. | [optional] 
**Env** | Pointer to [**[]PlatformEnvVarJSON**](PlatformEnvVarJSON.md) | Env is the app&#39;s environment variables, with every SECRET value masked to \&quot;\&quot; — the plaintext is in KMS and this surface never echoes it. That masking is why an empty secret value means \&quot;keep what is sealed\&quot; when posted back. | [optional] 
**Environment** | Pointer to **string** | Environment is the deploy target this app names, &#x60;production&#x60; when none was given. It is a LABEL: /v1/platform/environments derives the environment list from the apps that name one, so an environment exists as long as an app points at it and no route creates or deletes one. | [optional] 
**Health** | Pointer to **string** | Health rolls ready-vs-desired replicas up to a colour: green (all ready), yellow (some ready, or deliberately scaled to zero), red (none), or \&quot;\&quot; when the cluster reports no replica counts at all — unknown, never a guessed green. | [optional] 
**Id** | Pointer to **string** | ID is the server-minted application id (&#x60;app_…&#x60;). Routes address an app by project and slug; this is the key its deployments and builds carry. | [optional] 
**Image** | Pointer to [**PlatformImageView**](PlatformImageView.md) | Image is the image a source &#x60;image&#x60; app runs. For a git app only the tag is filled, stamped by the deploy that went live; the built ref is on the deployment. | [optional] 
**Name** | Pointer to **string** | Name is the display name. It is not an address — the slug is. | [optional] 
**Namespace** | Pointer to **string** | Namespace is where the app&#39;s cluster objects live, &#x60;tenant-&lt;org&gt;&#x60;. It is derived from the validated org and is never accepted from a request. | [optional] 
**Org** | Pointer to **string** | Org is the tenant that owns the app. It comes from the validated identity, never from the request, and it is the boundary every route is scoped to. | [optional] 
**Phase** | Pointer to **string** | Phase is the operator&#39;s own &#x60;status.phase&#x60; for the app&#39;s Service CR, read from the cluster on this request. Empty when there is no CR yet or the cluster could not be read. | [optional] 
**Port** | Pointer to **int64** | Port is the container port traffic is sent to. 8080 when the create asked for none, or for one outside 1–65535. | [optional] 
**ProjectId** | Pointer to **string** | ProjectID is the IAM project the app lives under, and it is that project&#39;s NAME — the (org,name) key IAM identifies it by, which is also what the &#x60;:project&#x60; path segment carries. There is no platform-minted project id. | [optional] 
**Replicas** | Pointer to **int64** | Replicas is how many copies the CR declares. It is CLAMPED to the deployment&#39;s ceiling rather than refused, so it can be below what was asked. | [optional] 
**Repo** | Pointer to [**PlatformGitSource**](PlatformGitSource.md) | Repo is the git origin a source &#x60;git&#x60; app builds from, and the repo+branch a landed push has to match to build it. | [optional] 
**SecretSync** | Pointer to **string** | SecretSync is how far the app&#39;s secret env has got into the cluster: \&quot;\&quot;|pending|syncing|ready|failed (secrets.go). It is best-effort and never fails a deploy, so &#x60;pending&#x60; is ordinary right after one. | [optional] 
**SecretSyncDetail** | Pointer to **string** | SecretSyncDetail is the honest reason when the sync is not ready — a missing CRD, an RBAC grant, a per-tenant credential. Empty when it is. | [optional] 
**Slug** | Pointer to **string** | Slug is the app&#39;s identity in the cluster: the operator CR&#39;s name, the first label of its default host, and the &#x60;:app&#x60; path segment. Unique per project. | [optional] 
**Source** | Pointer to **string** | Source is what the app deploys FROM: &#x60;git&#x60;, which builds Repo, or &#x60;image&#x60;, which runs Image as it is. It decides whether a deploy builds at all. | [optional] 
**Status** | Pointer to **string** | Status is the lifecycle THIS store records: draft (created, nothing in the cluster yet), building, deploying, live, stopped or error. What the cluster itself says is Phase and Health. | [optional] 
**StorageGb** | Pointer to **int64** | StorageGB is the persistent volume size in GiB. Absent means stateless — no volume at all — and it is clamped like Replicas. | [optional] 
**UpdatedAt** | Pointer to **int64** | UpdatedAt is when it last changed, unix seconds. Every lifecycle transition moves it, so it tracks deploys as well as edits. | [optional] 

## Methods

### NewPlatformAppOut

`func NewPlatformAppOut() *PlatformAppOut`

NewPlatformAppOut instantiates a new PlatformAppOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformAppOutWithDefaults

`func NewPlatformAppOutWithDefaults() *PlatformAppOut`

NewPlatformAppOutWithDefaults instantiates a new PlatformAppOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuildType

`func (o *PlatformAppOut) GetBuildType() string`

GetBuildType returns the BuildType field if non-nil, zero value otherwise.

### GetBuildTypeOk

`func (o *PlatformAppOut) GetBuildTypeOk() (*string, bool)`

GetBuildTypeOk returns a tuple with the BuildType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildType

`func (o *PlatformAppOut) SetBuildType(v string)`

SetBuildType sets BuildType field to given value.

### HasBuildType

`func (o *PlatformAppOut) HasBuildType() bool`

HasBuildType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PlatformAppOut) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PlatformAppOut) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PlatformAppOut) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PlatformAppOut) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCurrentDeploymentId

`func (o *PlatformAppOut) GetCurrentDeploymentId() string`

GetCurrentDeploymentId returns the CurrentDeploymentId field if non-nil, zero value otherwise.

### GetCurrentDeploymentIdOk

`func (o *PlatformAppOut) GetCurrentDeploymentIdOk() (*string, bool)`

GetCurrentDeploymentIdOk returns a tuple with the CurrentDeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentDeploymentId

`func (o *PlatformAppOut) SetCurrentDeploymentId(v string)`

SetCurrentDeploymentId sets CurrentDeploymentId field to given value.

### HasCurrentDeploymentId

`func (o *PlatformAppOut) HasCurrentDeploymentId() bool`

HasCurrentDeploymentId returns a boolean if a field has been set.

### GetDescription

`func (o *PlatformAppOut) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PlatformAppOut) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PlatformAppOut) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PlatformAppOut) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDockerfile

`func (o *PlatformAppOut) GetDockerfile() string`

GetDockerfile returns the Dockerfile field if non-nil, zero value otherwise.

### GetDockerfileOk

`func (o *PlatformAppOut) GetDockerfileOk() (*string, bool)`

GetDockerfileOk returns a tuple with the Dockerfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerfile

`func (o *PlatformAppOut) SetDockerfile(v string)`

SetDockerfile sets Dockerfile field to given value.

### HasDockerfile

`func (o *PlatformAppOut) HasDockerfile() bool`

HasDockerfile returns a boolean if a field has been set.

### GetDomains

`func (o *PlatformAppOut) GetDomains() []string`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *PlatformAppOut) GetDomainsOk() (*[]string, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *PlatformAppOut) SetDomains(v []string)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *PlatformAppOut) HasDomains() bool`

HasDomains returns a boolean if a field has been set.

### GetEnv

`func (o *PlatformAppOut) GetEnv() []PlatformEnvVarJSON`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *PlatformAppOut) GetEnvOk() (*[]PlatformEnvVarJSON, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *PlatformAppOut) SetEnv(v []PlatformEnvVarJSON)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *PlatformAppOut) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### GetEnvironment

`func (o *PlatformAppOut) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *PlatformAppOut) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *PlatformAppOut) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *PlatformAppOut) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetHealth

`func (o *PlatformAppOut) GetHealth() string`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *PlatformAppOut) GetHealthOk() (*string, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *PlatformAppOut) SetHealth(v string)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *PlatformAppOut) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetId

`func (o *PlatformAppOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PlatformAppOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PlatformAppOut) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PlatformAppOut) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImage

`func (o *PlatformAppOut) GetImage() PlatformImageView`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *PlatformAppOut) GetImageOk() (*PlatformImageView, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *PlatformAppOut) SetImage(v PlatformImageView)`

SetImage sets Image field to given value.

### HasImage

`func (o *PlatformAppOut) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetName

`func (o *PlatformAppOut) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PlatformAppOut) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PlatformAppOut) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PlatformAppOut) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *PlatformAppOut) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *PlatformAppOut) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *PlatformAppOut) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *PlatformAppOut) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetOrg

`func (o *PlatformAppOut) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *PlatformAppOut) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *PlatformAppOut) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *PlatformAppOut) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetPhase

`func (o *PlatformAppOut) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *PlatformAppOut) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *PlatformAppOut) SetPhase(v string)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *PlatformAppOut) HasPhase() bool`

HasPhase returns a boolean if a field has been set.

### GetPort

`func (o *PlatformAppOut) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *PlatformAppOut) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *PlatformAppOut) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *PlatformAppOut) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetProjectId

`func (o *PlatformAppOut) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *PlatformAppOut) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *PlatformAppOut) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *PlatformAppOut) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetReplicas

`func (o *PlatformAppOut) GetReplicas() int64`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *PlatformAppOut) GetReplicasOk() (*int64, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *PlatformAppOut) SetReplicas(v int64)`

SetReplicas sets Replicas field to given value.

### HasReplicas

`func (o *PlatformAppOut) HasReplicas() bool`

HasReplicas returns a boolean if a field has been set.

### GetRepo

`func (o *PlatformAppOut) GetRepo() PlatformGitSource`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *PlatformAppOut) GetRepoOk() (*PlatformGitSource, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *PlatformAppOut) SetRepo(v PlatformGitSource)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *PlatformAppOut) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### GetSecretSync

`func (o *PlatformAppOut) GetSecretSync() string`

GetSecretSync returns the SecretSync field if non-nil, zero value otherwise.

### GetSecretSyncOk

`func (o *PlatformAppOut) GetSecretSyncOk() (*string, bool)`

GetSecretSyncOk returns a tuple with the SecretSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretSync

`func (o *PlatformAppOut) SetSecretSync(v string)`

SetSecretSync sets SecretSync field to given value.

### HasSecretSync

`func (o *PlatformAppOut) HasSecretSync() bool`

HasSecretSync returns a boolean if a field has been set.

### GetSecretSyncDetail

`func (o *PlatformAppOut) GetSecretSyncDetail() string`

GetSecretSyncDetail returns the SecretSyncDetail field if non-nil, zero value otherwise.

### GetSecretSyncDetailOk

`func (o *PlatformAppOut) GetSecretSyncDetailOk() (*string, bool)`

GetSecretSyncDetailOk returns a tuple with the SecretSyncDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretSyncDetail

`func (o *PlatformAppOut) SetSecretSyncDetail(v string)`

SetSecretSyncDetail sets SecretSyncDetail field to given value.

### HasSecretSyncDetail

`func (o *PlatformAppOut) HasSecretSyncDetail() bool`

HasSecretSyncDetail returns a boolean if a field has been set.

### GetSlug

`func (o *PlatformAppOut) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *PlatformAppOut) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *PlatformAppOut) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *PlatformAppOut) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetSource

`func (o *PlatformAppOut) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *PlatformAppOut) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *PlatformAppOut) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *PlatformAppOut) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetStatus

`func (o *PlatformAppOut) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PlatformAppOut) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PlatformAppOut) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PlatformAppOut) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStorageGb

`func (o *PlatformAppOut) GetStorageGb() int64`

GetStorageGb returns the StorageGb field if non-nil, zero value otherwise.

### GetStorageGbOk

`func (o *PlatformAppOut) GetStorageGbOk() (*int64, bool)`

GetStorageGbOk returns a tuple with the StorageGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageGb

`func (o *PlatformAppOut) SetStorageGb(v int64)`

SetStorageGb sets StorageGb field to given value.

### HasStorageGb

`func (o *PlatformAppOut) HasStorageGb() bool`

HasStorageGb returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PlatformAppOut) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PlatformAppOut) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PlatformAppOut) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PlatformAppOut) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


