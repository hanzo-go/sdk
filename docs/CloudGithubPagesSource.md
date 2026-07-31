# CloudGithubPagesSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Branch** | Pointer to **string** | Branch is the branch the site builds from. | [optional] 
**Path** | Pointer to **string** | Path is the directory within that branch: \&quot;/\&quot; or \&quot;/docs\&quot;. | [optional] 

## Methods

### NewCloudGithubPagesSource

`func NewCloudGithubPagesSource() *CloudGithubPagesSource`

NewCloudGithubPagesSource instantiates a new CloudGithubPagesSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudGithubPagesSourceWithDefaults

`func NewCloudGithubPagesSourceWithDefaults() *CloudGithubPagesSource`

NewCloudGithubPagesSourceWithDefaults instantiates a new CloudGithubPagesSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBranch

`func (o *CloudGithubPagesSource) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *CloudGithubPagesSource) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *CloudGithubPagesSource) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *CloudGithubPagesSource) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetPath

`func (o *CloudGithubPagesSource) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *CloudGithubPagesSource) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *CloudGithubPagesSource) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *CloudGithubPagesSource) HasPath() bool`

HasPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


