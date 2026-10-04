# ProviderGitlabProjectView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloneUrl** | Pointer to **string** | CloneURL is the https remote to clone. | [optional] 
**DefaultBranch** | Pointer to **string** | DefaultBranch is the branch a clone lands on (\&quot;main\&quot; when GitLab names none, which is what an empty project reports). | [optional] 
**Description** | Pointer to **string** | Description is the project&#39;s own, empty when it has none. | [optional] 
**FullName** | Pointer to **string** | FullName is the namespace path (\&quot;acme/widgets\&quot;, \&quot;acme/team/widgets\&quot; for a subgroup) — the string GitLab calls path_with_namespace. | [optional] 
**HtmlUrl** | Pointer to **string** | HTMLURL is the project&#39;s page. | [optional] 
**Name** | Pointer to **string** | Name is the project&#39;s path segment (\&quot;widgets\&quot;), not its display name. | [optional] 
**Private** | Pointer to **bool** | Private is true for anything not publicly visible (private or internal). | [optional] 
**PushedAt** | Pointer to **string** | PushedAt is RFC3339 last activity, so a client can sort or say \&quot;2h ago\&quot;. | [optional] 

## Methods

### NewProviderGitlabProjectView

`func NewProviderGitlabProjectView() *ProviderGitlabProjectView`

NewProviderGitlabProjectView instantiates a new ProviderGitlabProjectView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderGitlabProjectViewWithDefaults

`func NewProviderGitlabProjectViewWithDefaults() *ProviderGitlabProjectView`

NewProviderGitlabProjectViewWithDefaults instantiates a new ProviderGitlabProjectView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloneUrl

`func (o *ProviderGitlabProjectView) GetCloneUrl() string`

GetCloneUrl returns the CloneUrl field if non-nil, zero value otherwise.

### GetCloneUrlOk

`func (o *ProviderGitlabProjectView) GetCloneUrlOk() (*string, bool)`

GetCloneUrlOk returns a tuple with the CloneUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloneUrl

`func (o *ProviderGitlabProjectView) SetCloneUrl(v string)`

SetCloneUrl sets CloneUrl field to given value.

### HasCloneUrl

`func (o *ProviderGitlabProjectView) HasCloneUrl() bool`

HasCloneUrl returns a boolean if a field has been set.

### GetDefaultBranch

`func (o *ProviderGitlabProjectView) GetDefaultBranch() string`

GetDefaultBranch returns the DefaultBranch field if non-nil, zero value otherwise.

### GetDefaultBranchOk

`func (o *ProviderGitlabProjectView) GetDefaultBranchOk() (*string, bool)`

GetDefaultBranchOk returns a tuple with the DefaultBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultBranch

`func (o *ProviderGitlabProjectView) SetDefaultBranch(v string)`

SetDefaultBranch sets DefaultBranch field to given value.

### HasDefaultBranch

`func (o *ProviderGitlabProjectView) HasDefaultBranch() bool`

HasDefaultBranch returns a boolean if a field has been set.

### GetDescription

`func (o *ProviderGitlabProjectView) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProviderGitlabProjectView) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProviderGitlabProjectView) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProviderGitlabProjectView) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFullName

`func (o *ProviderGitlabProjectView) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *ProviderGitlabProjectView) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *ProviderGitlabProjectView) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *ProviderGitlabProjectView) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### GetHtmlUrl

`func (o *ProviderGitlabProjectView) GetHtmlUrl() string`

GetHtmlUrl returns the HtmlUrl field if non-nil, zero value otherwise.

### GetHtmlUrlOk

`func (o *ProviderGitlabProjectView) GetHtmlUrlOk() (*string, bool)`

GetHtmlUrlOk returns a tuple with the HtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlUrl

`func (o *ProviderGitlabProjectView) SetHtmlUrl(v string)`

SetHtmlUrl sets HtmlUrl field to given value.

### HasHtmlUrl

`func (o *ProviderGitlabProjectView) HasHtmlUrl() bool`

HasHtmlUrl returns a boolean if a field has been set.

### GetName

`func (o *ProviderGitlabProjectView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProviderGitlabProjectView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProviderGitlabProjectView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProviderGitlabProjectView) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrivate

`func (o *ProviderGitlabProjectView) GetPrivate() bool`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *ProviderGitlabProjectView) GetPrivateOk() (*bool, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *ProviderGitlabProjectView) SetPrivate(v bool)`

SetPrivate sets Private field to given value.

### HasPrivate

`func (o *ProviderGitlabProjectView) HasPrivate() bool`

HasPrivate returns a boolean if a field has been set.

### GetPushedAt

`func (o *ProviderGitlabProjectView) GetPushedAt() string`

GetPushedAt returns the PushedAt field if non-nil, zero value otherwise.

### GetPushedAtOk

`func (o *ProviderGitlabProjectView) GetPushedAtOk() (*string, bool)`

GetPushedAtOk returns a tuple with the PushedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushedAt

`func (o *ProviderGitlabProjectView) SetPushedAt(v string)`

SetPushedAt sets PushedAt field to given value.

### HasPushedAt

`func (o *ProviderGitlabProjectView) HasPushedAt() bool`

HasPushedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


