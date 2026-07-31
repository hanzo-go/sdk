# ProjectsProject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Org** | **string** |  | 
**Slug** | **string** |  | 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Repo** | [**ProjectsRepoRef**](ProjectsRepoRef.md) |  | 
**Framework** | **string** | Build hint. | 
**Status** | **string** | Project lifecycle status. | 
**LiveUrl** | Pointer to **string** | Canonical live URL, https://&lt;slug&gt;.&lt;apex&gt;. Set once deployed. | [optional] 
**Bucket** | Pointer to **string** | S3-origin bucket holding the site. | [optional] 
**CurrentDeploymentId** | Pointer to **string** |  | [optional] 
**CacheControl** | Pointer to **string** | Per-project HTML/document Cache-Control policy applied at the S3 origin. | [optional] 
**LastPurgeAt** | Pointer to **int64** | Unix time (seconds) of the last edge cache-tag purge. | [optional] 
**CreatedAt** | **int64** |  | 
**UpdatedAt** | **int64** |  | 

## Methods

### NewProjectsProject

`func NewProjectsProject(id string, org string, slug string, name string, repo ProjectsRepoRef, framework string, status string, createdAt int64, updatedAt int64, ) *ProjectsProject`

NewProjectsProject instantiates a new ProjectsProject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectsProjectWithDefaults

`func NewProjectsProjectWithDefaults() *ProjectsProject`

NewProjectsProjectWithDefaults instantiates a new ProjectsProject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProjectsProject) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectsProject) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectsProject) SetId(v string)`

SetId sets Id field to given value.


### GetOrg

`func (o *ProjectsProject) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *ProjectsProject) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *ProjectsProject) SetOrg(v string)`

SetOrg sets Org field to given value.


### GetSlug

`func (o *ProjectsProject) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *ProjectsProject) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *ProjectsProject) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetName

`func (o *ProjectsProject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectsProject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectsProject) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ProjectsProject) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProjectsProject) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProjectsProject) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProjectsProject) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRepo

`func (o *ProjectsProject) GetRepo() ProjectsRepoRef`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *ProjectsProject) GetRepoOk() (*ProjectsRepoRef, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *ProjectsProject) SetRepo(v ProjectsRepoRef)`

SetRepo sets Repo field to given value.


### GetFramework

`func (o *ProjectsProject) GetFramework() string`

GetFramework returns the Framework field if non-nil, zero value otherwise.

### GetFrameworkOk

`func (o *ProjectsProject) GetFrameworkOk() (*string, bool)`

GetFrameworkOk returns a tuple with the Framework field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFramework

`func (o *ProjectsProject) SetFramework(v string)`

SetFramework sets Framework field to given value.


### GetStatus

`func (o *ProjectsProject) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProjectsProject) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProjectsProject) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetLiveUrl

`func (o *ProjectsProject) GetLiveUrl() string`

GetLiveUrl returns the LiveUrl field if non-nil, zero value otherwise.

### GetLiveUrlOk

`func (o *ProjectsProject) GetLiveUrlOk() (*string, bool)`

GetLiveUrlOk returns a tuple with the LiveUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveUrl

`func (o *ProjectsProject) SetLiveUrl(v string)`

SetLiveUrl sets LiveUrl field to given value.

### HasLiveUrl

`func (o *ProjectsProject) HasLiveUrl() bool`

HasLiveUrl returns a boolean if a field has been set.

### GetBucket

`func (o *ProjectsProject) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *ProjectsProject) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *ProjectsProject) SetBucket(v string)`

SetBucket sets Bucket field to given value.

### HasBucket

`func (o *ProjectsProject) HasBucket() bool`

HasBucket returns a boolean if a field has been set.

### GetCurrentDeploymentId

`func (o *ProjectsProject) GetCurrentDeploymentId() string`

GetCurrentDeploymentId returns the CurrentDeploymentId field if non-nil, zero value otherwise.

### GetCurrentDeploymentIdOk

`func (o *ProjectsProject) GetCurrentDeploymentIdOk() (*string, bool)`

GetCurrentDeploymentIdOk returns a tuple with the CurrentDeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentDeploymentId

`func (o *ProjectsProject) SetCurrentDeploymentId(v string)`

SetCurrentDeploymentId sets CurrentDeploymentId field to given value.

### HasCurrentDeploymentId

`func (o *ProjectsProject) HasCurrentDeploymentId() bool`

HasCurrentDeploymentId returns a boolean if a field has been set.

### GetCacheControl

`func (o *ProjectsProject) GetCacheControl() string`

GetCacheControl returns the CacheControl field if non-nil, zero value otherwise.

### GetCacheControlOk

`func (o *ProjectsProject) GetCacheControlOk() (*string, bool)`

GetCacheControlOk returns a tuple with the CacheControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheControl

`func (o *ProjectsProject) SetCacheControl(v string)`

SetCacheControl sets CacheControl field to given value.

### HasCacheControl

`func (o *ProjectsProject) HasCacheControl() bool`

HasCacheControl returns a boolean if a field has been set.

### GetLastPurgeAt

`func (o *ProjectsProject) GetLastPurgeAt() int64`

GetLastPurgeAt returns the LastPurgeAt field if non-nil, zero value otherwise.

### GetLastPurgeAtOk

`func (o *ProjectsProject) GetLastPurgeAtOk() (*int64, bool)`

GetLastPurgeAtOk returns a tuple with the LastPurgeAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPurgeAt

`func (o *ProjectsProject) SetLastPurgeAt(v int64)`

SetLastPurgeAt sets LastPurgeAt field to given value.

### HasLastPurgeAt

`func (o *ProjectsProject) HasLastPurgeAt() bool`

HasLastPurgeAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ProjectsProject) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ProjectsProject) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ProjectsProject) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ProjectsProject) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ProjectsProject) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ProjectsProject) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


