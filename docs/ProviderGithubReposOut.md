# ProviderGithubReposOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Connected** | Pointer to **bool** | Connected is whether this answer is the caller&#39;s OWN GitHub. False with repositories is an org admin seeing the org&#39;s installations; false with none is the cue to connect. | [optional] 
**Next** | Pointer to **string** | Next is the cursor for the page after this one; empty on the last page. | [optional] 
**Repos** | Pointer to [**[]ProviderGithubRepoItem**](ProviderGithubRepoItem.md) | Repos is this page, most recently pushed first. Never null. | [optional] 
**Total** | Pointer to **int64** | Total is how many repositories match q and owner, across every page. | [optional] 
**Unread** | Pointer to **[]string** | Unread names the installations that could not be read, so a short list is distinguishable from a complete one. Absent when the answer is whole. | [optional] 

## Methods

### NewProviderGithubReposOut

`func NewProviderGithubReposOut() *ProviderGithubReposOut`

NewProviderGithubReposOut instantiates a new ProviderGithubReposOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderGithubReposOutWithDefaults

`func NewProviderGithubReposOutWithDefaults() *ProviderGithubReposOut`

NewProviderGithubReposOutWithDefaults instantiates a new ProviderGithubReposOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnected

`func (o *ProviderGithubReposOut) GetConnected() bool`

GetConnected returns the Connected field if non-nil, zero value otherwise.

### GetConnectedOk

`func (o *ProviderGithubReposOut) GetConnectedOk() (*bool, bool)`

GetConnectedOk returns a tuple with the Connected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnected

`func (o *ProviderGithubReposOut) SetConnected(v bool)`

SetConnected sets Connected field to given value.

### HasConnected

`func (o *ProviderGithubReposOut) HasConnected() bool`

HasConnected returns a boolean if a field has been set.

### GetNext

`func (o *ProviderGithubReposOut) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *ProviderGithubReposOut) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *ProviderGithubReposOut) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *ProviderGithubReposOut) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetRepos

`func (o *ProviderGithubReposOut) GetRepos() []ProviderGithubRepoItem`

GetRepos returns the Repos field if non-nil, zero value otherwise.

### GetReposOk

`func (o *ProviderGithubReposOut) GetReposOk() (*[]ProviderGithubRepoItem, bool)`

GetReposOk returns a tuple with the Repos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepos

`func (o *ProviderGithubReposOut) SetRepos(v []ProviderGithubRepoItem)`

SetRepos sets Repos field to given value.

### HasRepos

`func (o *ProviderGithubReposOut) HasRepos() bool`

HasRepos returns a boolean if a field has been set.

### GetTotal

`func (o *ProviderGithubReposOut) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ProviderGithubReposOut) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ProviderGithubReposOut) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ProviderGithubReposOut) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetUnread

`func (o *ProviderGithubReposOut) GetUnread() []string`

GetUnread returns the Unread field if non-nil, zero value otherwise.

### GetUnreadOk

`func (o *ProviderGithubReposOut) GetUnreadOk() (*[]string, bool)`

GetUnreadOk returns a tuple with the Unread field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnread

`func (o *ProviderGithubReposOut) SetUnread(v []string)`

SetUnread sets Unread field to given value.

### HasUnread

`func (o *ProviderGithubReposOut) HasUnread() bool`

HasUnread returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


