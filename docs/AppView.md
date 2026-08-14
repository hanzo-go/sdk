# AppView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuildType** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **int32** |  | [optional] 
**CurrentDeploymentId** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Dockerfile** | Pointer to **string** |  | [optional] 
**Domains** | Pointer to **[]string** |  | [optional] 
**Env** | Pointer to [**[]EnvVarJSON**](EnvVarJSON.md) |  | [optional] 
**Environment** | Pointer to **string** |  | [optional] 
**Health** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Image** | Pointer to [**ImageView**](ImageView.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**Org** | Pointer to **string** |  | [optional] 
**Phase** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**ProjectId** | Pointer to **string** |  | [optional] 
**Replicas** | Pointer to **int32** |  | [optional] 
**Repo** | Pointer to [**GitSource**](GitSource.md) |  | [optional] 
**SecretSync** | Pointer to **string** | \&quot;\&quot;|pending|syncing|ready|failed (secrets.go) | [optional] 
**SecretSyncDetail** | Pointer to **string** | honest reason when not ready | [optional] 
**Slug** | Pointer to **string** |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StorageGb** | Pointer to **int32** | GiB; absent means stateless | [optional] 
**UpdatedAt** | Pointer to **int32** |  | [optional] 

## Methods

### NewAppView

`func NewAppView() *AppView`

NewAppView instantiates a new AppView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppViewWithDefaults

`func NewAppViewWithDefaults() *AppView`

NewAppViewWithDefaults instantiates a new AppView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuildType

`func (o *AppView) GetBuildType() string`

GetBuildType returns the BuildType field if non-nil, zero value otherwise.

### GetBuildTypeOk

`func (o *AppView) GetBuildTypeOk() (*string, bool)`

GetBuildTypeOk returns a tuple with the BuildType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildType

`func (o *AppView) SetBuildType(v string)`

SetBuildType sets BuildType field to given value.

### HasBuildType

`func (o *AppView) HasBuildType() bool`

HasBuildType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AppView) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AppView) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AppView) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AppView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCurrentDeploymentId

`func (o *AppView) GetCurrentDeploymentId() string`

GetCurrentDeploymentId returns the CurrentDeploymentId field if non-nil, zero value otherwise.

### GetCurrentDeploymentIdOk

`func (o *AppView) GetCurrentDeploymentIdOk() (*string, bool)`

GetCurrentDeploymentIdOk returns a tuple with the CurrentDeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentDeploymentId

`func (o *AppView) SetCurrentDeploymentId(v string)`

SetCurrentDeploymentId sets CurrentDeploymentId field to given value.

### HasCurrentDeploymentId

`func (o *AppView) HasCurrentDeploymentId() bool`

HasCurrentDeploymentId returns a boolean if a field has been set.

### GetDescription

`func (o *AppView) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AppView) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AppView) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AppView) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDockerfile

`func (o *AppView) GetDockerfile() string`

GetDockerfile returns the Dockerfile field if non-nil, zero value otherwise.

### GetDockerfileOk

`func (o *AppView) GetDockerfileOk() (*string, bool)`

GetDockerfileOk returns a tuple with the Dockerfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerfile

`func (o *AppView) SetDockerfile(v string)`

SetDockerfile sets Dockerfile field to given value.

### HasDockerfile

`func (o *AppView) HasDockerfile() bool`

HasDockerfile returns a boolean if a field has been set.

### GetDomains

`func (o *AppView) GetDomains() []string`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *AppView) GetDomainsOk() (*[]string, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *AppView) SetDomains(v []string)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *AppView) HasDomains() bool`

HasDomains returns a boolean if a field has been set.

### GetEnv

`func (o *AppView) GetEnv() []EnvVarJSON`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *AppView) GetEnvOk() (*[]EnvVarJSON, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *AppView) SetEnv(v []EnvVarJSON)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *AppView) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### GetEnvironment

`func (o *AppView) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *AppView) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *AppView) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *AppView) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetHealth

`func (o *AppView) GetHealth() string`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *AppView) GetHealthOk() (*string, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *AppView) SetHealth(v string)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *AppView) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetId

`func (o *AppView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AppView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImage

`func (o *AppView) GetImage() ImageView`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *AppView) GetImageOk() (*ImageView, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *AppView) SetImage(v ImageView)`

SetImage sets Image field to given value.

### HasImage

`func (o *AppView) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetName

`func (o *AppView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AppView) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *AppView) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *AppView) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *AppView) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *AppView) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetOrg

`func (o *AppView) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *AppView) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *AppView) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *AppView) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetPhase

`func (o *AppView) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *AppView) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *AppView) SetPhase(v string)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *AppView) HasPhase() bool`

HasPhase returns a boolean if a field has been set.

### GetPort

`func (o *AppView) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *AppView) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *AppView) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *AppView) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetProjectId

`func (o *AppView) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *AppView) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *AppView) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *AppView) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetReplicas

`func (o *AppView) GetReplicas() int32`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *AppView) GetReplicasOk() (*int32, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *AppView) SetReplicas(v int32)`

SetReplicas sets Replicas field to given value.

### HasReplicas

`func (o *AppView) HasReplicas() bool`

HasReplicas returns a boolean if a field has been set.

### GetRepo

`func (o *AppView) GetRepo() GitSource`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *AppView) GetRepoOk() (*GitSource, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *AppView) SetRepo(v GitSource)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *AppView) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### GetSecretSync

`func (o *AppView) GetSecretSync() string`

GetSecretSync returns the SecretSync field if non-nil, zero value otherwise.

### GetSecretSyncOk

`func (o *AppView) GetSecretSyncOk() (*string, bool)`

GetSecretSyncOk returns a tuple with the SecretSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretSync

`func (o *AppView) SetSecretSync(v string)`

SetSecretSync sets SecretSync field to given value.

### HasSecretSync

`func (o *AppView) HasSecretSync() bool`

HasSecretSync returns a boolean if a field has been set.

### GetSecretSyncDetail

`func (o *AppView) GetSecretSyncDetail() string`

GetSecretSyncDetail returns the SecretSyncDetail field if non-nil, zero value otherwise.

### GetSecretSyncDetailOk

`func (o *AppView) GetSecretSyncDetailOk() (*string, bool)`

GetSecretSyncDetailOk returns a tuple with the SecretSyncDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretSyncDetail

`func (o *AppView) SetSecretSyncDetail(v string)`

SetSecretSyncDetail sets SecretSyncDetail field to given value.

### HasSecretSyncDetail

`func (o *AppView) HasSecretSyncDetail() bool`

HasSecretSyncDetail returns a boolean if a field has been set.

### GetSlug

`func (o *AppView) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *AppView) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *AppView) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *AppView) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetSource

`func (o *AppView) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *AppView) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *AppView) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *AppView) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetStatus

`func (o *AppView) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AppView) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AppView) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AppView) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStorageGb

`func (o *AppView) GetStorageGb() int32`

GetStorageGb returns the StorageGb field if non-nil, zero value otherwise.

### GetStorageGbOk

`func (o *AppView) GetStorageGbOk() (*int32, bool)`

GetStorageGbOk returns a tuple with the StorageGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageGb

`func (o *AppView) SetStorageGb(v int32)`

SetStorageGb sets StorageGb field to given value.

### HasStorageGb

`func (o *AppView) HasStorageGb() bool`

HasStorageGb returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AppView) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AppView) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AppView) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AppView) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


