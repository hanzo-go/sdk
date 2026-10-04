# ProviderGitlabProjectsOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** | Account is the connected GitLab username, so a client can label the list without a second call. Empty when the connection recorded none. | [optional] 
**Projects** | Pointer to [**[]ProviderGitlabProjectView**](ProviderGitlabProjectView.md) | Projects is every project the token reaches, newest activity first. Never null; [] when the account has none. | [optional] 

## Methods

### NewProviderGitlabProjectsOut

`func NewProviderGitlabProjectsOut() *ProviderGitlabProjectsOut`

NewProviderGitlabProjectsOut instantiates a new ProviderGitlabProjectsOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderGitlabProjectsOutWithDefaults

`func NewProviderGitlabProjectsOutWithDefaults() *ProviderGitlabProjectsOut`

NewProviderGitlabProjectsOutWithDefaults instantiates a new ProviderGitlabProjectsOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *ProviderGitlabProjectsOut) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ProviderGitlabProjectsOut) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ProviderGitlabProjectsOut) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ProviderGitlabProjectsOut) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetProjects

`func (o *ProviderGitlabProjectsOut) GetProjects() []ProviderGitlabProjectView`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *ProviderGitlabProjectsOut) GetProjectsOk() (*[]ProviderGitlabProjectView, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *ProviderGitlabProjectsOut) SetProjects(v []ProviderGitlabProjectView)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *ProviderGitlabProjectsOut) HasProjects() bool`

HasProjects returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


