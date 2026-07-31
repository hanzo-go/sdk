# CloudGithubPagesEnableReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Branch** | Pointer to **string** | Branch is the legacy source branch; empty defaults to the repo&#39;s own default branch. Ignored when buildType is \&quot;workflow\&quot;. | [optional] 
**BuildType** | Pointer to **string** | BuildType selects the builder: \&quot;workflow\&quot; builds via GitHub Actions, anything else builds from the branch source above. | [optional] 
**Path** | Pointer to **string** | Path is the source directory within the branch: \&quot;/\&quot; (the default) or \&quot;/docs\&quot;. GitHub allows no others. | [optional] 
**Repo** | Pointer to **string** | Repo is the repository, from the :repo path segment. | [optional] 

## Methods

### NewCloudGithubPagesEnableReq

`func NewCloudGithubPagesEnableReq() *CloudGithubPagesEnableReq`

NewCloudGithubPagesEnableReq instantiates a new CloudGithubPagesEnableReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudGithubPagesEnableReqWithDefaults

`func NewCloudGithubPagesEnableReqWithDefaults() *CloudGithubPagesEnableReq`

NewCloudGithubPagesEnableReqWithDefaults instantiates a new CloudGithubPagesEnableReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBranch

`func (o *CloudGithubPagesEnableReq) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *CloudGithubPagesEnableReq) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *CloudGithubPagesEnableReq) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *CloudGithubPagesEnableReq) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetBuildType

`func (o *CloudGithubPagesEnableReq) GetBuildType() string`

GetBuildType returns the BuildType field if non-nil, zero value otherwise.

### GetBuildTypeOk

`func (o *CloudGithubPagesEnableReq) GetBuildTypeOk() (*string, bool)`

GetBuildTypeOk returns a tuple with the BuildType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildType

`func (o *CloudGithubPagesEnableReq) SetBuildType(v string)`

SetBuildType sets BuildType field to given value.

### HasBuildType

`func (o *CloudGithubPagesEnableReq) HasBuildType() bool`

HasBuildType returns a boolean if a field has been set.

### GetPath

`func (o *CloudGithubPagesEnableReq) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *CloudGithubPagesEnableReq) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *CloudGithubPagesEnableReq) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *CloudGithubPagesEnableReq) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetRepo

`func (o *CloudGithubPagesEnableReq) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *CloudGithubPagesEnableReq) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *CloudGithubPagesEnableReq) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *CloudGithubPagesEnableReq) HasRepo() bool`

HasRepo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


