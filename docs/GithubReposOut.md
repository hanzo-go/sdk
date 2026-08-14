# GithubReposOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Repos** | Pointer to [**[]GithubRepoView**](GithubRepoView.md) | Repos is every repo the installation grants. Never null; [] when none. | [optional] 
**Unread** | Pointer to **[]string** | Unread names the connected accounts this answer could NOT read, so a short list is distinguishable from a complete one. Absent when the answer is whole.  The fan-out is per installation, and one account failing used to be dropped in silence: the response stayed 200 and simply carried fewer repositories, erroring only when EVERY account failed. Measured, twice in a row, minutes apart: 1475 repositories, then 1157 — a whole installation missing with nothing in the answer to say so. Anything driven off the list then under-covers and reports success, which is the failure this field ends. | [optional] 

## Methods

### NewGithubReposOut

`func NewGithubReposOut() *GithubReposOut`

NewGithubReposOut instantiates a new GithubReposOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGithubReposOutWithDefaults

`func NewGithubReposOutWithDefaults() *GithubReposOut`

NewGithubReposOutWithDefaults instantiates a new GithubReposOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepos

`func (o *GithubReposOut) GetRepos() []GithubRepoView`

GetRepos returns the Repos field if non-nil, zero value otherwise.

### GetReposOk

`func (o *GithubReposOut) GetReposOk() (*[]GithubRepoView, bool)`

GetReposOk returns a tuple with the Repos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepos

`func (o *GithubReposOut) SetRepos(v []GithubRepoView)`

SetRepos sets Repos field to given value.

### HasRepos

`func (o *GithubReposOut) HasRepos() bool`

HasRepos returns a boolean if a field has been set.

### GetUnread

`func (o *GithubReposOut) GetUnread() []string`

GetUnread returns the Unread field if non-nil, zero value otherwise.

### GetUnreadOk

`func (o *GithubReposOut) GetUnreadOk() (*[]string, bool)`

GetUnreadOk returns a tuple with the Unread field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnread

`func (o *GithubReposOut) SetUnread(v []string)`

SetUnread sets Unread field to given value.

### HasUnread

`func (o *GithubReposOut) HasUnread() bool`

HasUnread returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


