# ProjectProjectsDeployment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bucket** | Pointer to **string** | Bucket is the object-store bucket its files were written to. | [optional] 
**Bytes** | Pointer to **int64** | Bytes is their total size in bytes. | [optional] 
**Commit** | Pointer to **string** | Commit is the revision that was built, for a deployment that came from a repository. Absent for an uploaded artifact, which has no revision. | [optional] 
**CreatedAt** | Pointer to **int64** | CreatedAt is when the deployment was queued, as Unix seconds. | [optional] 
**Files** | Pointer to **int64** | Files is how many objects the deployment published. | [optional] 
**Id** | Pointer to **string** | ID identifies this one deployment attempt, and is what CI quotes back to complete it. | [optional] 
**LiveUrl** | Pointer to **string** | LiveURL is where this deployment serves, once it is live. | [optional] 
**Message** | Pointer to **string** | Message is what happened, in words — the build&#39;s own note, or on a failure why it failed. | [optional] 
**Prefix** | Pointer to **string** | Prefix is the key prefix within that bucket holding EXACTLY this deployment&#39;s objects — the unit an upload grant is scoped to, so a grant for one deployment cannot write over another. | [optional] 
**ProjectId** | Pointer to **string** | ProjectID is the project this deployment belongs to. | [optional] 
**Source** | Pointer to **string** | Source is what caused the deployment — a git push, an uploaded artifact, a generated site. | [optional] 
**Status** | Pointer to **string** | Status is where the attempt got to — queued, live, or failed. A deployment that is live is not necessarily the one SERVING: the project&#39;s own currentDeploymentId says which is. | [optional] 
**UpdatedAt** | Pointer to **int64** | UpdatedAt is when it last changed state, as Unix seconds — so the gap between the two is how long the build took. | [optional] 
**Upload** | Pointer to [**ProjectProjectsUploadGrant**](ProjectProjectsUploadGrant.md) | Upload is the prefix-scoped, short-lived S3 write grant handed to CI with a queued git deployment, so it needs no bucket credential (grant.go). Present ONLY on the 202 that creates the deployment — it is never stored and never replayed on a later read, so a grant cannot outlive the build it was minted for by being fetched again. | [optional] 
**Version** | Pointer to **int64** | Version counts deployments of this project from 1, so the history reads as an ordered sequence rather than by timestamp. It is per project, not global. | [optional] 

## Methods

### NewProjectProjectsDeployment

`func NewProjectProjectsDeployment() *ProjectProjectsDeployment`

NewProjectProjectsDeployment instantiates a new ProjectProjectsDeployment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectProjectsDeploymentWithDefaults

`func NewProjectProjectsDeploymentWithDefaults() *ProjectProjectsDeployment`

NewProjectProjectsDeploymentWithDefaults instantiates a new ProjectProjectsDeployment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucket

`func (o *ProjectProjectsDeployment) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *ProjectProjectsDeployment) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *ProjectProjectsDeployment) SetBucket(v string)`

SetBucket sets Bucket field to given value.

### HasBucket

`func (o *ProjectProjectsDeployment) HasBucket() bool`

HasBucket returns a boolean if a field has been set.

### GetBytes

`func (o *ProjectProjectsDeployment) GetBytes() int64`

GetBytes returns the Bytes field if non-nil, zero value otherwise.

### GetBytesOk

`func (o *ProjectProjectsDeployment) GetBytesOk() (*int64, bool)`

GetBytesOk returns a tuple with the Bytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytes

`func (o *ProjectProjectsDeployment) SetBytes(v int64)`

SetBytes sets Bytes field to given value.

### HasBytes

`func (o *ProjectProjectsDeployment) HasBytes() bool`

HasBytes returns a boolean if a field has been set.

### GetCommit

`func (o *ProjectProjectsDeployment) GetCommit() string`

GetCommit returns the Commit field if non-nil, zero value otherwise.

### GetCommitOk

`func (o *ProjectProjectsDeployment) GetCommitOk() (*string, bool)`

GetCommitOk returns a tuple with the Commit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommit

`func (o *ProjectProjectsDeployment) SetCommit(v string)`

SetCommit sets Commit field to given value.

### HasCommit

`func (o *ProjectProjectsDeployment) HasCommit() bool`

HasCommit returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ProjectProjectsDeployment) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ProjectProjectsDeployment) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ProjectProjectsDeployment) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ProjectProjectsDeployment) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetFiles

`func (o *ProjectProjectsDeployment) GetFiles() int64`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *ProjectProjectsDeployment) GetFilesOk() (*int64, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *ProjectProjectsDeployment) SetFiles(v int64)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *ProjectProjectsDeployment) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetId

`func (o *ProjectProjectsDeployment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectProjectsDeployment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectProjectsDeployment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProjectProjectsDeployment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLiveUrl

`func (o *ProjectProjectsDeployment) GetLiveUrl() string`

GetLiveUrl returns the LiveUrl field if non-nil, zero value otherwise.

### GetLiveUrlOk

`func (o *ProjectProjectsDeployment) GetLiveUrlOk() (*string, bool)`

GetLiveUrlOk returns a tuple with the LiveUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveUrl

`func (o *ProjectProjectsDeployment) SetLiveUrl(v string)`

SetLiveUrl sets LiveUrl field to given value.

### HasLiveUrl

`func (o *ProjectProjectsDeployment) HasLiveUrl() bool`

HasLiveUrl returns a boolean if a field has been set.

### GetMessage

`func (o *ProjectProjectsDeployment) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ProjectProjectsDeployment) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ProjectProjectsDeployment) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ProjectProjectsDeployment) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetPrefix

`func (o *ProjectProjectsDeployment) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *ProjectProjectsDeployment) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *ProjectProjectsDeployment) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *ProjectProjectsDeployment) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetProjectId

`func (o *ProjectProjectsDeployment) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ProjectProjectsDeployment) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ProjectProjectsDeployment) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *ProjectProjectsDeployment) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetSource

`func (o *ProjectProjectsDeployment) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ProjectProjectsDeployment) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ProjectProjectsDeployment) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *ProjectProjectsDeployment) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetStatus

`func (o *ProjectProjectsDeployment) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProjectProjectsDeployment) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProjectProjectsDeployment) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ProjectProjectsDeployment) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ProjectProjectsDeployment) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ProjectProjectsDeployment) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ProjectProjectsDeployment) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ProjectProjectsDeployment) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpload

`func (o *ProjectProjectsDeployment) GetUpload() ProjectProjectsUploadGrant`

GetUpload returns the Upload field if non-nil, zero value otherwise.

### GetUploadOk

`func (o *ProjectProjectsDeployment) GetUploadOk() (*ProjectProjectsUploadGrant, bool)`

GetUploadOk returns a tuple with the Upload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpload

`func (o *ProjectProjectsDeployment) SetUpload(v ProjectProjectsUploadGrant)`

SetUpload sets Upload field to given value.

### HasUpload

`func (o *ProjectProjectsDeployment) HasUpload() bool`

HasUpload returns a boolean if a field has been set.

### GetVersion

`func (o *ProjectProjectsDeployment) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ProjectProjectsDeployment) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ProjectProjectsDeployment) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ProjectProjectsDeployment) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


