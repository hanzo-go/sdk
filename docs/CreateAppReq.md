# CreateAppReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuildType** | Pointer to **string** | BuildType is &#x60;pack&#x60; — the zero-config default that detects any project — or &#x60;dockerfile&#x60;, the explicit escape hatch. An image app never builds. | [optional] 
**Description** | Pointer to **string** | Description is free text about what the application is. | [optional] 
**Dockerfile** | Pointer to **string** | Dockerfile is the path to build from, for buildType &#x60;dockerfile&#x60;. | [optional] 
**Domains** | Pointer to **[]string** | Domains are extra ingress hosts. The canonical default host is always attached; a bare custom host is refused here and must go through add-domain → verify first. | [optional] 
**Env** | Pointer to [**[]EnvVarJSON**](EnvVarJSON.md) | Env is the application&#39;s environment. Keys must match &#x60;^[A-Za-z_][A-Za-z0-9_]*$&#x60;; a variable marked &#x60;secret: true&#x60; is sealed into KMS and its plaintext is never written to the database. | [optional] 
**Environment** | Pointer to **string** | Environment is the deploy target this app names (\&quot;production\&quot; by default). | [optional] 
**Image** | Pointer to [**ImageOrigin**](ImageOrigin.md) | Image is the container image to run, for source &#x60;image&#x60;. | [optional] 
**Name** | Pointer to **string** | Name is the application&#39;s display name. Required; the slug is derived from it when none is given. | [optional] 
**Port** | Pointer to **int64** | Port is the container port the app listens on. | [optional] 
**Project** | Pointer to **string** | Project is the project to create the application under, from the path. | [optional] 
**Replicas** | Pointer to **int64** | Replicas is how many copies to run; clamped to the deployment&#39;s limit rather than refused. | [optional] 
**Repo** | Pointer to [**GitOrigin**](GitOrigin.md) | Repo is the git source to build from, for source &#x60;git&#x60;. | [optional] 
**Slug** | Pointer to **string** | Slug is the app&#39;s identity in the cluster — its CR name and part of its host. Given or derived from Name, it must match &#x60;^[a-z0-9]([a-z0-9-]{0,38}[a-z0-9])?$&#x60;, and one already used in this project is 409. | [optional] 
**Source** | Pointer to **string** | Source is &#x60;git&#x60;, which requires repo.url, or &#x60;image&#x60;, which requires image.repository. Anything else is 400. | [optional] 
**StorageGb** | Pointer to **int64** | StorageGB is the persistent volume size in GiB; absent means stateless. Clamped to the deployment&#39;s limit rather than refused. | [optional] 

## Methods

### NewCreateAppReq

`func NewCreateAppReq() *CreateAppReq`

NewCreateAppReq instantiates a new CreateAppReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAppReqWithDefaults

`func NewCreateAppReqWithDefaults() *CreateAppReq`

NewCreateAppReqWithDefaults instantiates a new CreateAppReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuildType

`func (o *CreateAppReq) GetBuildType() string`

GetBuildType returns the BuildType field if non-nil, zero value otherwise.

### GetBuildTypeOk

`func (o *CreateAppReq) GetBuildTypeOk() (*string, bool)`

GetBuildTypeOk returns a tuple with the BuildType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildType

`func (o *CreateAppReq) SetBuildType(v string)`

SetBuildType sets BuildType field to given value.

### HasBuildType

`func (o *CreateAppReq) HasBuildType() bool`

HasBuildType returns a boolean if a field has been set.

### GetDescription

`func (o *CreateAppReq) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateAppReq) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateAppReq) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateAppReq) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDockerfile

`func (o *CreateAppReq) GetDockerfile() string`

GetDockerfile returns the Dockerfile field if non-nil, zero value otherwise.

### GetDockerfileOk

`func (o *CreateAppReq) GetDockerfileOk() (*string, bool)`

GetDockerfileOk returns a tuple with the Dockerfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerfile

`func (o *CreateAppReq) SetDockerfile(v string)`

SetDockerfile sets Dockerfile field to given value.

### HasDockerfile

`func (o *CreateAppReq) HasDockerfile() bool`

HasDockerfile returns a boolean if a field has been set.

### GetDomains

`func (o *CreateAppReq) GetDomains() []string`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *CreateAppReq) GetDomainsOk() (*[]string, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *CreateAppReq) SetDomains(v []string)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *CreateAppReq) HasDomains() bool`

HasDomains returns a boolean if a field has been set.

### GetEnv

`func (o *CreateAppReq) GetEnv() []EnvVarJSON`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *CreateAppReq) GetEnvOk() (*[]EnvVarJSON, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *CreateAppReq) SetEnv(v []EnvVarJSON)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *CreateAppReq) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### GetEnvironment

`func (o *CreateAppReq) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *CreateAppReq) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *CreateAppReq) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *CreateAppReq) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetImage

`func (o *CreateAppReq) GetImage() ImageOrigin`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *CreateAppReq) GetImageOk() (*ImageOrigin, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *CreateAppReq) SetImage(v ImageOrigin)`

SetImage sets Image field to given value.

### HasImage

`func (o *CreateAppReq) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetName

`func (o *CreateAppReq) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateAppReq) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateAppReq) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateAppReq) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPort

`func (o *CreateAppReq) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *CreateAppReq) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *CreateAppReq) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *CreateAppReq) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetProject

`func (o *CreateAppReq) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *CreateAppReq) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *CreateAppReq) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *CreateAppReq) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetReplicas

`func (o *CreateAppReq) GetReplicas() int64`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *CreateAppReq) GetReplicasOk() (*int64, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *CreateAppReq) SetReplicas(v int64)`

SetReplicas sets Replicas field to given value.

### HasReplicas

`func (o *CreateAppReq) HasReplicas() bool`

HasReplicas returns a boolean if a field has been set.

### GetRepo

`func (o *CreateAppReq) GetRepo() GitOrigin`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *CreateAppReq) GetRepoOk() (*GitOrigin, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *CreateAppReq) SetRepo(v GitOrigin)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *CreateAppReq) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### GetSlug

`func (o *CreateAppReq) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *CreateAppReq) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *CreateAppReq) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *CreateAppReq) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetSource

`func (o *CreateAppReq) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *CreateAppReq) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *CreateAppReq) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *CreateAppReq) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetStorageGb

`func (o *CreateAppReq) GetStorageGb() int64`

GetStorageGb returns the StorageGb field if non-nil, zero value otherwise.

### GetStorageGbOk

`func (o *CreateAppReq) GetStorageGbOk() (*int64, bool)`

GetStorageGbOk returns a tuple with the StorageGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageGb

`func (o *CreateAppReq) SetStorageGb(v int64)`

SetStorageGb sets StorageGb field to given value.

### HasStorageGb

`func (o *CreateAppReq) HasStorageGb() bool`

HasStorageGb returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


