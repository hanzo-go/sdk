# GithubPagesUpdatedOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Repo** | Pointer to **string** | Repo is the repository that was updated. | [optional] 
**Updated** | Pointer to **bool** | Updated is always true — a failure is an HTTP error, never this shape. | [optional] 

## Methods

### NewGithubPagesUpdatedOut

`func NewGithubPagesUpdatedOut() *GithubPagesUpdatedOut`

NewGithubPagesUpdatedOut instantiates a new GithubPagesUpdatedOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGithubPagesUpdatedOutWithDefaults

`func NewGithubPagesUpdatedOutWithDefaults() *GithubPagesUpdatedOut`

NewGithubPagesUpdatedOutWithDefaults instantiates a new GithubPagesUpdatedOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepo

`func (o *GithubPagesUpdatedOut) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *GithubPagesUpdatedOut) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *GithubPagesUpdatedOut) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *GithubPagesUpdatedOut) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### GetUpdated

`func (o *GithubPagesUpdatedOut) GetUpdated() bool`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *GithubPagesUpdatedOut) GetUpdatedOk() (*bool, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *GithubPagesUpdatedOut) SetUpdated(v bool)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *GithubPagesUpdatedOut) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


