# ProjectProjectsRepo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Branch** | Pointer to **string** | Branch is the ref a push has to touch for this project to rebuild. Pushes to any other branch are ignored. | [optional] 
**Provider** | Pointer to **string** | Provider is the forge the URL was recognised as — it decides which webhook and which credential reach the repository, and is DERIVED from the URL rather than chosen by the caller. | [optional] 
**Url** | Pointer to **string** | URL is the clone address of the repository this project builds from. | [optional] 

## Methods

### NewProjectProjectsRepo

`func NewProjectProjectsRepo() *ProjectProjectsRepo`

NewProjectProjectsRepo instantiates a new ProjectProjectsRepo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectProjectsRepoWithDefaults

`func NewProjectProjectsRepoWithDefaults() *ProjectProjectsRepo`

NewProjectProjectsRepoWithDefaults instantiates a new ProjectProjectsRepo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBranch

`func (o *ProjectProjectsRepo) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *ProjectProjectsRepo) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *ProjectProjectsRepo) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *ProjectProjectsRepo) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetProvider

`func (o *ProjectProjectsRepo) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ProjectProjectsRepo) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ProjectProjectsRepo) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *ProjectProjectsRepo) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetUrl

`func (o *ProjectProjectsRepo) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ProjectProjectsRepo) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ProjectProjectsRepo) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ProjectProjectsRepo) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


