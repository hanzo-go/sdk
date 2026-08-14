# GithubForkReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Org** | Pointer to **string** | Org is the GitHub account to fork INTO; empty forks to the installation&#39;s own account, which is the common case. | [optional] 
**Repo** | Pointer to **string** | Repo is a repository the org&#39;s installation was GRANTED, by name. | [optional] 

## Methods

### NewGithubForkReq

`func NewGithubForkReq() *GithubForkReq`

NewGithubForkReq instantiates a new GithubForkReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGithubForkReqWithDefaults

`func NewGithubForkReqWithDefaults() *GithubForkReq`

NewGithubForkReqWithDefaults instantiates a new GithubForkReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrg

`func (o *GithubForkReq) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *GithubForkReq) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *GithubForkReq) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *GithubForkReq) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetRepo

`func (o *GithubForkReq) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *GithubForkReq) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *GithubForkReq) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *GithubForkReq) HasRepo() bool`

HasRepo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


