# PlatformAppView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuildType** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **int32** |  | [optional] 
**CurrentDeploymentId** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Dockerfile** | Pointer to **string** |  | [optional] 
**Domains** | Pointer to **[]string** |  | [optional] 
**Env** | Pointer to [**[]PlatformEnvVarJSON**](PlatformEnvVarJSON.md) |  | [optional] 
**Environment** | Pointer to **string** |  | [optional] 
**Health** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Image** | Pointer to [**PlatformImageView**](PlatformImageView.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**Org** | Pointer to **string** |  | [optional] 
**Phase** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**ProjectId** | Pointer to **string** |  | [optional] 
**Replicas** | Pointer to **int32** |  | [optional] 
**Repo** | Pointer to [**PlatformGitSource**](PlatformGitSource.md) |  | [optional] 
**SecretSync** | Pointer to **string** |  | [optional] 
**SecretSyncDetail** | Pointer to **string** |  | [optional] 
**Slug** | Pointer to **string** |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StorageGb** | Pointer to **int32** |  | [optional] 
**UpdatedAt** | Pointer to **int32** |  | [optional] 

## Methods

### NewPlatformAppView

`func NewPlatformAppView() *PlatformAppView`

NewPlatformAppView instantiates a new PlatformAppView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformAppViewWithDefaults

`func NewPlatformAppViewWithDefaults() *PlatformAppView`

NewPlatformAppViewWithDefaults instantiates a new PlatformAppView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuildType

`func (o *PlatformAppView) GetBuildType() string`

GetBuildType returns the BuildType field if non-nil, zero value otherwise.

### GetBuildTypeOk

`func (o *PlatformAppView) GetBuildTypeOk() (*string, bool)`

GetBuildTypeOk returns a tuple with the BuildType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildType

`func (o *PlatformAppView) SetBuildType(v string)`

SetBuildType sets BuildType field to given value.

### HasBuildType

`func (o *PlatformAppView) HasBuildType() bool`

HasBuildType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PlatformAppView) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PlatformAppView) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PlatformAppView) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PlatformAppView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCurrentDeploymentId

`func (o *PlatformAppView) GetCurrentDeploymentId() string`

GetCurrentDeploymentId returns the CurrentDeploymentId field if non-nil, zero value otherwise.

### GetCurrentDeploymentIdOk

`func (o *PlatformAppView) GetCurrentDeploymentIdOk() (*string, bool)`

GetCurrentDeploymentIdOk returns a tuple with the CurrentDeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentDeploymentId

`func (o *PlatformAppView) SetCurrentDeploymentId(v string)`

SetCurrentDeploymentId sets CurrentDeploymentId field to given value.

### HasCurrentDeploymentId

`func (o *PlatformAppView) HasCurrentDeploymentId() bool`

HasCurrentDeploymentId returns a boolean if a field has been set.

### GetDescription

`func (o *PlatformAppView) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PlatformAppView) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PlatformAppView) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PlatformAppView) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDockerfile

`func (o *PlatformAppView) GetDockerfile() string`

GetDockerfile returns the Dockerfile field if non-nil, zero value otherwise.

### GetDockerfileOk

`func (o *PlatformAppView) GetDockerfileOk() (*string, bool)`

GetDockerfileOk returns a tuple with the Dockerfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerfile

`func (o *PlatformAppView) SetDockerfile(v string)`

SetDockerfile sets Dockerfile field to given value.

### HasDockerfile

`func (o *PlatformAppView) HasDockerfile() bool`

HasDockerfile returns a boolean if a field has been set.

### GetDomains

`func (o *PlatformAppView) GetDomains() []string`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *PlatformAppView) GetDomainsOk() (*[]string, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *PlatformAppView) SetDomains(v []string)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *PlatformAppView) HasDomains() bool`

HasDomains returns a boolean if a field has been set.

### GetEnv

`func (o *PlatformAppView) GetEnv() []PlatformEnvVarJSON`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *PlatformAppView) GetEnvOk() (*[]PlatformEnvVarJSON, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *PlatformAppView) SetEnv(v []PlatformEnvVarJSON)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *PlatformAppView) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### GetEnvironment

`func (o *PlatformAppView) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *PlatformAppView) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *PlatformAppView) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *PlatformAppView) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetHealth

`func (o *PlatformAppView) GetHealth() string`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *PlatformAppView) GetHealthOk() (*string, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *PlatformAppView) SetHealth(v string)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *PlatformAppView) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetId

`func (o *PlatformAppView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PlatformAppView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PlatformAppView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PlatformAppView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImage

`func (o *PlatformAppView) GetImage() PlatformImageView`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *PlatformAppView) GetImageOk() (*PlatformImageView, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *PlatformAppView) SetImage(v PlatformImageView)`

SetImage sets Image field to given value.

### HasImage

`func (o *PlatformAppView) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetName

`func (o *PlatformAppView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PlatformAppView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PlatformAppView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PlatformAppView) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *PlatformAppView) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *PlatformAppView) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *PlatformAppView) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *PlatformAppView) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetOrg

`func (o *PlatformAppView) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *PlatformAppView) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *PlatformAppView) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *PlatformAppView) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetPhase

`func (o *PlatformAppView) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *PlatformAppView) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *PlatformAppView) SetPhase(v string)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *PlatformAppView) HasPhase() bool`

HasPhase returns a boolean if a field has been set.

### GetPort

`func (o *PlatformAppView) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *PlatformAppView) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *PlatformAppView) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *PlatformAppView) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetProjectId

`func (o *PlatformAppView) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *PlatformAppView) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *PlatformAppView) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *PlatformAppView) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetReplicas

`func (o *PlatformAppView) GetReplicas() int32`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *PlatformAppView) GetReplicasOk() (*int32, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *PlatformAppView) SetReplicas(v int32)`

SetReplicas sets Replicas field to given value.

### HasReplicas

`func (o *PlatformAppView) HasReplicas() bool`

HasReplicas returns a boolean if a field has been set.

### GetRepo

`func (o *PlatformAppView) GetRepo() PlatformGitSource`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *PlatformAppView) GetRepoOk() (*PlatformGitSource, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *PlatformAppView) SetRepo(v PlatformGitSource)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *PlatformAppView) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### GetSecretSync

`func (o *PlatformAppView) GetSecretSync() string`

GetSecretSync returns the SecretSync field if non-nil, zero value otherwise.

### GetSecretSyncOk

`func (o *PlatformAppView) GetSecretSyncOk() (*string, bool)`

GetSecretSyncOk returns a tuple with the SecretSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretSync

`func (o *PlatformAppView) SetSecretSync(v string)`

SetSecretSync sets SecretSync field to given value.

### HasSecretSync

`func (o *PlatformAppView) HasSecretSync() bool`

HasSecretSync returns a boolean if a field has been set.

### GetSecretSyncDetail

`func (o *PlatformAppView) GetSecretSyncDetail() string`

GetSecretSyncDetail returns the SecretSyncDetail field if non-nil, zero value otherwise.

### GetSecretSyncDetailOk

`func (o *PlatformAppView) GetSecretSyncDetailOk() (*string, bool)`

GetSecretSyncDetailOk returns a tuple with the SecretSyncDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretSyncDetail

`func (o *PlatformAppView) SetSecretSyncDetail(v string)`

SetSecretSyncDetail sets SecretSyncDetail field to given value.

### HasSecretSyncDetail

`func (o *PlatformAppView) HasSecretSyncDetail() bool`

HasSecretSyncDetail returns a boolean if a field has been set.

### GetSlug

`func (o *PlatformAppView) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *PlatformAppView) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *PlatformAppView) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *PlatformAppView) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetSource

`func (o *PlatformAppView) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *PlatformAppView) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *PlatformAppView) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *PlatformAppView) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetStatus

`func (o *PlatformAppView) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PlatformAppView) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PlatformAppView) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PlatformAppView) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStorageGb

`func (o *PlatformAppView) GetStorageGb() int32`

GetStorageGb returns the StorageGb field if non-nil, zero value otherwise.

### GetStorageGbOk

`func (o *PlatformAppView) GetStorageGbOk() (*int32, bool)`

GetStorageGbOk returns a tuple with the StorageGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageGb

`func (o *PlatformAppView) SetStorageGb(v int32)`

SetStorageGb sets StorageGb field to given value.

### HasStorageGb

`func (o *PlatformAppView) HasStorageGb() bool`

HasStorageGb returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PlatformAppView) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PlatformAppView) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PlatformAppView) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PlatformAppView) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


