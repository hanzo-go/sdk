# GithubPagesDisabledOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Disabled** | Pointer to **bool** | Disabled is always true — a failure is an HTTP error, never this shape. | [optional] 
**Repo** | Pointer to **string** | Repo is the repository whose site was deleted. | [optional] 

## Methods

### NewGithubPagesDisabledOut

`func NewGithubPagesDisabledOut() *GithubPagesDisabledOut`

NewGithubPagesDisabledOut instantiates a new GithubPagesDisabledOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGithubPagesDisabledOutWithDefaults

`func NewGithubPagesDisabledOutWithDefaults() *GithubPagesDisabledOut`

NewGithubPagesDisabledOutWithDefaults instantiates a new GithubPagesDisabledOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisabled

`func (o *GithubPagesDisabledOut) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *GithubPagesDisabledOut) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *GithubPagesDisabledOut) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *GithubPagesDisabledOut) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetRepo

`func (o *GithubPagesDisabledOut) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *GithubPagesDisabledOut) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *GithubPagesDisabledOut) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *GithubPagesDisabledOut) HasRepo() bool`

HasRepo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


