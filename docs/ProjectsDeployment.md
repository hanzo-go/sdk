# ProjectsDeployment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**ProjectId** | **string** |  | 
**Version** | **int32** | Monotonic per project, 1-based. | 
**Status** | **string** | Deployment status. | 
**Source** | **string** | How the artifact was produced. | 
**Commit** | Pointer to **string** |  | [optional] 
**LiveUrl** | Pointer to **string** | Canonical live URL, https://&lt;slug&gt;.&lt;apex&gt;. | [optional] 
**Bucket** | Pointer to **string** | S3-origin bucket. | [optional] 
**Prefix** | Pointer to **string** | S3-origin key prefix the site is served from (&lt;org&gt;/&lt;slug&gt;). | [optional] 
**Files** | **int32** |  | 
**Bytes** | **int64** |  | 
**Message** | Pointer to **string** |  | [optional] 
**CreatedAt** | **int64** |  | 
**UpdatedAt** | **int64** |  | 

## Methods

### NewProjectsDeployment

`func NewProjectsDeployment(id string, projectId string, version int32, status string, source string, files int32, bytes int64, createdAt int64, updatedAt int64, ) *ProjectsDeployment`

NewProjectsDeployment instantiates a new ProjectsDeployment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectsDeploymentWithDefaults

`func NewProjectsDeploymentWithDefaults() *ProjectsDeployment`

NewProjectsDeploymentWithDefaults instantiates a new ProjectsDeployment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProjectsDeployment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectsDeployment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectsDeployment) SetId(v string)`

SetId sets Id field to given value.


### GetProjectId

`func (o *ProjectsDeployment) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ProjectsDeployment) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ProjectsDeployment) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetVersion

`func (o *ProjectsDeployment) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ProjectsDeployment) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ProjectsDeployment) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetStatus

`func (o *ProjectsDeployment) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProjectsDeployment) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProjectsDeployment) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetSource

`func (o *ProjectsDeployment) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ProjectsDeployment) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ProjectsDeployment) SetSource(v string)`

SetSource sets Source field to given value.


### GetCommit

`func (o *ProjectsDeployment) GetCommit() string`

GetCommit returns the Commit field if non-nil, zero value otherwise.

### GetCommitOk

`func (o *ProjectsDeployment) GetCommitOk() (*string, bool)`

GetCommitOk returns a tuple with the Commit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommit

`func (o *ProjectsDeployment) SetCommit(v string)`

SetCommit sets Commit field to given value.

### HasCommit

`func (o *ProjectsDeployment) HasCommit() bool`

HasCommit returns a boolean if a field has been set.

### GetLiveUrl

`func (o *ProjectsDeployment) GetLiveUrl() string`

GetLiveUrl returns the LiveUrl field if non-nil, zero value otherwise.

### GetLiveUrlOk

`func (o *ProjectsDeployment) GetLiveUrlOk() (*string, bool)`

GetLiveUrlOk returns a tuple with the LiveUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveUrl

`func (o *ProjectsDeployment) SetLiveUrl(v string)`

SetLiveUrl sets LiveUrl field to given value.

### HasLiveUrl

`func (o *ProjectsDeployment) HasLiveUrl() bool`

HasLiveUrl returns a boolean if a field has been set.

### GetBucket

`func (o *ProjectsDeployment) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *ProjectsDeployment) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *ProjectsDeployment) SetBucket(v string)`

SetBucket sets Bucket field to given value.

### HasBucket

`func (o *ProjectsDeployment) HasBucket() bool`

HasBucket returns a boolean if a field has been set.

### GetPrefix

`func (o *ProjectsDeployment) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *ProjectsDeployment) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *ProjectsDeployment) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *ProjectsDeployment) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetFiles

`func (o *ProjectsDeployment) GetFiles() int32`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *ProjectsDeployment) GetFilesOk() (*int32, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *ProjectsDeployment) SetFiles(v int32)`

SetFiles sets Files field to given value.


### GetBytes

`func (o *ProjectsDeployment) GetBytes() int64`

GetBytes returns the Bytes field if non-nil, zero value otherwise.

### GetBytesOk

`func (o *ProjectsDeployment) GetBytesOk() (*int64, bool)`

GetBytesOk returns a tuple with the Bytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytes

`func (o *ProjectsDeployment) SetBytes(v int64)`

SetBytes sets Bytes field to given value.


### GetMessage

`func (o *ProjectsDeployment) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ProjectsDeployment) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ProjectsDeployment) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ProjectsDeployment) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ProjectsDeployment) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ProjectsDeployment) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ProjectsDeployment) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ProjectsDeployment) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ProjectsDeployment) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ProjectsDeployment) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


