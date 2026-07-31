# CloudAppView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuildType** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **int32** |  | [optional] 
**CurrentDeploymentId** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Dockerfile** | Pointer to **string** |  | [optional] 
**Domains** | Pointer to **[]string** |  | [optional] 
**Env** | Pointer to [**[]CloudEnvVarJSON**](CloudEnvVarJSON.md) |  | [optional] 
**Environment** | Pointer to **string** |  | [optional] 
**Health** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Image** | Pointer to [**CloudImageView**](CloudImageView.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**Org** | Pointer to **string** |  | [optional] 
**Phase** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**ProjectId** | Pointer to **string** |  | [optional] 
**Replicas** | Pointer to **int32** |  | [optional] 
**Repo** | Pointer to [**CloudGitSource**](CloudGitSource.md) |  | [optional] 
**SecretSync** | Pointer to **string** |  | [optional] 
**SecretSyncDetail** | Pointer to **string** |  | [optional] 
**Slug** | Pointer to **string** |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StorageGb** | Pointer to **int32** |  | [optional] 
**UpdatedAt** | Pointer to **int32** |  | [optional] 

## Methods

### NewCloudAppView

`func NewCloudAppView() *CloudAppView`

NewCloudAppView instantiates a new CloudAppView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudAppViewWithDefaults

`func NewCloudAppViewWithDefaults() *CloudAppView`

NewCloudAppViewWithDefaults instantiates a new CloudAppView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuildType

`func (o *CloudAppView) GetBuildType() string`

GetBuildType returns the BuildType field if non-nil, zero value otherwise.

### GetBuildTypeOk

`func (o *CloudAppView) GetBuildTypeOk() (*string, bool)`

GetBuildTypeOk returns a tuple with the BuildType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildType

`func (o *CloudAppView) SetBuildType(v string)`

SetBuildType sets BuildType field to given value.

### HasBuildType

`func (o *CloudAppView) HasBuildType() bool`

HasBuildType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CloudAppView) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CloudAppView) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CloudAppView) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CloudAppView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCurrentDeploymentId

`func (o *CloudAppView) GetCurrentDeploymentId() string`

GetCurrentDeploymentId returns the CurrentDeploymentId field if non-nil, zero value otherwise.

### GetCurrentDeploymentIdOk

`func (o *CloudAppView) GetCurrentDeploymentIdOk() (*string, bool)`

GetCurrentDeploymentIdOk returns a tuple with the CurrentDeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentDeploymentId

`func (o *CloudAppView) SetCurrentDeploymentId(v string)`

SetCurrentDeploymentId sets CurrentDeploymentId field to given value.

### HasCurrentDeploymentId

`func (o *CloudAppView) HasCurrentDeploymentId() bool`

HasCurrentDeploymentId returns a boolean if a field has been set.

### GetDescription

`func (o *CloudAppView) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CloudAppView) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CloudAppView) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CloudAppView) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDockerfile

`func (o *CloudAppView) GetDockerfile() string`

GetDockerfile returns the Dockerfile field if non-nil, zero value otherwise.

### GetDockerfileOk

`func (o *CloudAppView) GetDockerfileOk() (*string, bool)`

GetDockerfileOk returns a tuple with the Dockerfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerfile

`func (o *CloudAppView) SetDockerfile(v string)`

SetDockerfile sets Dockerfile field to given value.

### HasDockerfile

`func (o *CloudAppView) HasDockerfile() bool`

HasDockerfile returns a boolean if a field has been set.

### GetDomains

`func (o *CloudAppView) GetDomains() []string`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *CloudAppView) GetDomainsOk() (*[]string, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *CloudAppView) SetDomains(v []string)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *CloudAppView) HasDomains() bool`

HasDomains returns a boolean if a field has been set.

### GetEnv

`func (o *CloudAppView) GetEnv() []CloudEnvVarJSON`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *CloudAppView) GetEnvOk() (*[]CloudEnvVarJSON, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *CloudAppView) SetEnv(v []CloudEnvVarJSON)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *CloudAppView) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### GetEnvironment

`func (o *CloudAppView) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *CloudAppView) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *CloudAppView) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *CloudAppView) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetHealth

`func (o *CloudAppView) GetHealth() string`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *CloudAppView) GetHealthOk() (*string, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *CloudAppView) SetHealth(v string)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *CloudAppView) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetId

`func (o *CloudAppView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudAppView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudAppView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudAppView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImage

`func (o *CloudAppView) GetImage() CloudImageView`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *CloudAppView) GetImageOk() (*CloudImageView, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *CloudAppView) SetImage(v CloudImageView)`

SetImage sets Image field to given value.

### HasImage

`func (o *CloudAppView) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetName

`func (o *CloudAppView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudAppView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudAppView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudAppView) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *CloudAppView) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *CloudAppView) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *CloudAppView) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *CloudAppView) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetOrg

`func (o *CloudAppView) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *CloudAppView) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *CloudAppView) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *CloudAppView) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetPhase

`func (o *CloudAppView) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *CloudAppView) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *CloudAppView) SetPhase(v string)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *CloudAppView) HasPhase() bool`

HasPhase returns a boolean if a field has been set.

### GetPort

`func (o *CloudAppView) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *CloudAppView) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *CloudAppView) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *CloudAppView) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetProjectId

`func (o *CloudAppView) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *CloudAppView) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *CloudAppView) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *CloudAppView) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetReplicas

`func (o *CloudAppView) GetReplicas() int32`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *CloudAppView) GetReplicasOk() (*int32, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *CloudAppView) SetReplicas(v int32)`

SetReplicas sets Replicas field to given value.

### HasReplicas

`func (o *CloudAppView) HasReplicas() bool`

HasReplicas returns a boolean if a field has been set.

### GetRepo

`func (o *CloudAppView) GetRepo() CloudGitSource`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *CloudAppView) GetRepoOk() (*CloudGitSource, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *CloudAppView) SetRepo(v CloudGitSource)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *CloudAppView) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### GetSecretSync

`func (o *CloudAppView) GetSecretSync() string`

GetSecretSync returns the SecretSync field if non-nil, zero value otherwise.

### GetSecretSyncOk

`func (o *CloudAppView) GetSecretSyncOk() (*string, bool)`

GetSecretSyncOk returns a tuple with the SecretSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretSync

`func (o *CloudAppView) SetSecretSync(v string)`

SetSecretSync sets SecretSync field to given value.

### HasSecretSync

`func (o *CloudAppView) HasSecretSync() bool`

HasSecretSync returns a boolean if a field has been set.

### GetSecretSyncDetail

`func (o *CloudAppView) GetSecretSyncDetail() string`

GetSecretSyncDetail returns the SecretSyncDetail field if non-nil, zero value otherwise.

### GetSecretSyncDetailOk

`func (o *CloudAppView) GetSecretSyncDetailOk() (*string, bool)`

GetSecretSyncDetailOk returns a tuple with the SecretSyncDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretSyncDetail

`func (o *CloudAppView) SetSecretSyncDetail(v string)`

SetSecretSyncDetail sets SecretSyncDetail field to given value.

### HasSecretSyncDetail

`func (o *CloudAppView) HasSecretSyncDetail() bool`

HasSecretSyncDetail returns a boolean if a field has been set.

### GetSlug

`func (o *CloudAppView) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *CloudAppView) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *CloudAppView) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *CloudAppView) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetSource

`func (o *CloudAppView) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *CloudAppView) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *CloudAppView) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *CloudAppView) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetStatus

`func (o *CloudAppView) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudAppView) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudAppView) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudAppView) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStorageGb

`func (o *CloudAppView) GetStorageGb() int32`

GetStorageGb returns the StorageGb field if non-nil, zero value otherwise.

### GetStorageGbOk

`func (o *CloudAppView) GetStorageGbOk() (*int32, bool)`

GetStorageGbOk returns a tuple with the StorageGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageGb

`func (o *CloudAppView) SetStorageGb(v int32)`

SetStorageGb sets StorageGb field to given value.

### HasStorageGb

`func (o *CloudAppView) HasStorageGb() bool`

HasStorageGb returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CloudAppView) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CloudAppView) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CloudAppView) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CloudAppView) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


