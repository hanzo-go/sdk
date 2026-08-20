# GitlabProjectsOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** | Account is the connected GitLab username, so a client can label the list without a second call. Empty when the connection recorded none. | [optional] 
**Projects** | Pointer to [**[]GitlabProjectView**](GitlabProjectView.md) | Projects is every project the token reaches, newest activity first. Never null; [] when the account has none. | [optional] 

## Methods

### NewGitlabProjectsOut

`func NewGitlabProjectsOut() *GitlabProjectsOut`

NewGitlabProjectsOut instantiates a new GitlabProjectsOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitlabProjectsOutWithDefaults

`func NewGitlabProjectsOutWithDefaults() *GitlabProjectsOut`

NewGitlabProjectsOutWithDefaults instantiates a new GitlabProjectsOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *GitlabProjectsOut) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *GitlabProjectsOut) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *GitlabProjectsOut) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *GitlabProjectsOut) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetProjects

`func (o *GitlabProjectsOut) GetProjects() []GitlabProjectView`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *GitlabProjectsOut) GetProjectsOk() (*[]GitlabProjectView, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *GitlabProjectsOut) SetProjects(v []GitlabProjectView)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *GitlabProjectsOut) HasProjects() bool`

HasProjects returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


