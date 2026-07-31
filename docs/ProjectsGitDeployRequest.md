# ProjectsGitDeployRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | Pointer to **string** | Deploy source selector; \&quot;git\&quot; for the CI build path. | [optional] 
**Commit** | Pointer to **string** |  | [optional] 
**Branch** | Pointer to **string** |  | [optional] 

## Methods

### NewProjectsGitDeployRequest

`func NewProjectsGitDeployRequest() *ProjectsGitDeployRequest`

NewProjectsGitDeployRequest instantiates a new ProjectsGitDeployRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectsGitDeployRequestWithDefaults

`func NewProjectsGitDeployRequestWithDefaults() *ProjectsGitDeployRequest`

NewProjectsGitDeployRequestWithDefaults instantiates a new ProjectsGitDeployRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *ProjectsGitDeployRequest) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ProjectsGitDeployRequest) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ProjectsGitDeployRequest) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *ProjectsGitDeployRequest) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetCommit

`func (o *ProjectsGitDeployRequest) GetCommit() string`

GetCommit returns the Commit field if non-nil, zero value otherwise.

### GetCommitOk

`func (o *ProjectsGitDeployRequest) GetCommitOk() (*string, bool)`

GetCommitOk returns a tuple with the Commit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommit

`func (o *ProjectsGitDeployRequest) SetCommit(v string)`

SetCommit sets Commit field to given value.

### HasCommit

`func (o *ProjectsGitDeployRequest) HasCommit() bool`

HasCommit returns a boolean if a field has been set.

### GetBranch

`func (o *ProjectsGitDeployRequest) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *ProjectsGitDeployRequest) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *ProjectsGitDeployRequest) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *ProjectsGitDeployRequest) HasBranch() bool`

HasBranch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


