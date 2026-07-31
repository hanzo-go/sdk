# ProjectsRepoRef

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** | Git remote URL. | [optional] 
**Branch** | Pointer to **string** | Branch to build (defaults to \&quot;main\&quot; when a url is set). | [optional] 

## Methods

### NewProjectsRepoRef

`func NewProjectsRepoRef() *ProjectsRepoRef`

NewProjectsRepoRef instantiates a new ProjectsRepoRef object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectsRepoRefWithDefaults

`func NewProjectsRepoRefWithDefaults() *ProjectsRepoRef`

NewProjectsRepoRefWithDefaults instantiates a new ProjectsRepoRef object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *ProjectsRepoRef) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ProjectsRepoRef) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ProjectsRepoRef) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ProjectsRepoRef) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetBranch

`func (o *ProjectsRepoRef) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *ProjectsRepoRef) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *ProjectsRepoRef) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *ProjectsRepoRef) HasBranch() bool`

HasBranch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


