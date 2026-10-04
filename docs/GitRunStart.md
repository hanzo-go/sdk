# GitRunStart

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | Ref is the branch to run; empty means the repository&#39;s default. | [optional] 
**Repo** | Pointer to **string** | Repo is the repository to run. | [optional] 

## Methods

### NewGitRunStart

`func NewGitRunStart() *GitRunStart`

NewGitRunStart instantiates a new GitRunStart object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitRunStartWithDefaults

`func NewGitRunStartWithDefaults() *GitRunStart`

NewGitRunStartWithDefaults instantiates a new GitRunStart object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GitRunStart) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GitRunStart) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GitRunStart) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GitRunStart) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetRepo

`func (o *GitRunStart) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *GitRunStart) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *GitRunStart) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *GitRunStart) HasRepo() bool`

HasRepo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


