# GitCreateRepo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Repo name; must match ^[A-Za-z0-9][A-Za-z0-9._-]{0,63}$ | 
**Project** | Pointer to **string** | Optional project sub-scope (else the X-Project-Id header) | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewGitCreateRepo

`func NewGitCreateRepo(name string, ) *GitCreateRepo`

NewGitCreateRepo instantiates a new GitCreateRepo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitCreateRepoWithDefaults

`func NewGitCreateRepoWithDefaults() *GitCreateRepo`

NewGitCreateRepoWithDefaults instantiates a new GitCreateRepo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GitCreateRepo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GitCreateRepo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GitCreateRepo) SetName(v string)`

SetName sets Name field to given value.


### GetProject

`func (o *GitCreateRepo) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *GitCreateRepo) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *GitCreateRepo) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *GitCreateRepo) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetDescription

`func (o *GitCreateRepo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GitCreateRepo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GitCreateRepo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GitCreateRepo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


