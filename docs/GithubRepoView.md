# GithubRepoView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultBranch** | Pointer to **string** | DefaultBranch is the repo&#39;s default branch at GitHub. | [optional] 
**FullName** | Pointer to **string** | FullName is GitHub&#39;s owner/name. | [optional] 
**HtmlUrl** | Pointer to **string** | HTMLURL is the repo&#39;s page at GitHub. | [optional] 
**Imported** | Pointer to **bool** | Imported is whether this repo has been mirrored into git.hanzo.ai. | [optional] 
**LastSyncedAt** | Pointer to **string** | LastSyncedAt is the last successful mirror, RFC 3339 UTC. Absent if never. | [optional] 
**Name** | Pointer to **string** | Name is the repository&#39;s short name within the installation. | [optional] 
**Private** | Pointer to **bool** | Private is GitHub&#39;s visibility bit for the repo. | [optional] 
**SyncStatus** | Pointer to **string** | SyncStatus is \&quot;synced\&quot;, \&quot;conflict\&quot;, or \&quot;\&quot; when the repo is not imported. | [optional] 

## Methods

### NewGithubRepoView

`func NewGithubRepoView() *GithubRepoView`

NewGithubRepoView instantiates a new GithubRepoView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGithubRepoViewWithDefaults

`func NewGithubRepoViewWithDefaults() *GithubRepoView`

NewGithubRepoViewWithDefaults instantiates a new GithubRepoView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultBranch

`func (o *GithubRepoView) GetDefaultBranch() string`

GetDefaultBranch returns the DefaultBranch field if non-nil, zero value otherwise.

### GetDefaultBranchOk

`func (o *GithubRepoView) GetDefaultBranchOk() (*string, bool)`

GetDefaultBranchOk returns a tuple with the DefaultBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultBranch

`func (o *GithubRepoView) SetDefaultBranch(v string)`

SetDefaultBranch sets DefaultBranch field to given value.

### HasDefaultBranch

`func (o *GithubRepoView) HasDefaultBranch() bool`

HasDefaultBranch returns a boolean if a field has been set.

### GetFullName

`func (o *GithubRepoView) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *GithubRepoView) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *GithubRepoView) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *GithubRepoView) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### GetHtmlUrl

`func (o *GithubRepoView) GetHtmlUrl() string`

GetHtmlUrl returns the HtmlUrl field if non-nil, zero value otherwise.

### GetHtmlUrlOk

`func (o *GithubRepoView) GetHtmlUrlOk() (*string, bool)`

GetHtmlUrlOk returns a tuple with the HtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlUrl

`func (o *GithubRepoView) SetHtmlUrl(v string)`

SetHtmlUrl sets HtmlUrl field to given value.

### HasHtmlUrl

`func (o *GithubRepoView) HasHtmlUrl() bool`

HasHtmlUrl returns a boolean if a field has been set.

### GetImported

`func (o *GithubRepoView) GetImported() bool`

GetImported returns the Imported field if non-nil, zero value otherwise.

### GetImportedOk

`func (o *GithubRepoView) GetImportedOk() (*bool, bool)`

GetImportedOk returns a tuple with the Imported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImported

`func (o *GithubRepoView) SetImported(v bool)`

SetImported sets Imported field to given value.

### HasImported

`func (o *GithubRepoView) HasImported() bool`

HasImported returns a boolean if a field has been set.

### GetLastSyncedAt

`func (o *GithubRepoView) GetLastSyncedAt() string`

GetLastSyncedAt returns the LastSyncedAt field if non-nil, zero value otherwise.

### GetLastSyncedAtOk

`func (o *GithubRepoView) GetLastSyncedAtOk() (*string, bool)`

GetLastSyncedAtOk returns a tuple with the LastSyncedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSyncedAt

`func (o *GithubRepoView) SetLastSyncedAt(v string)`

SetLastSyncedAt sets LastSyncedAt field to given value.

### HasLastSyncedAt

`func (o *GithubRepoView) HasLastSyncedAt() bool`

HasLastSyncedAt returns a boolean if a field has been set.

### GetName

`func (o *GithubRepoView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GithubRepoView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GithubRepoView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GithubRepoView) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrivate

`func (o *GithubRepoView) GetPrivate() bool`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *GithubRepoView) GetPrivateOk() (*bool, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *GithubRepoView) SetPrivate(v bool)`

SetPrivate sets Private field to given value.

### HasPrivate

`func (o *GithubRepoView) HasPrivate() bool`

HasPrivate returns a boolean if a field has been set.

### GetSyncStatus

`func (o *GithubRepoView) GetSyncStatus() string`

GetSyncStatus returns the SyncStatus field if non-nil, zero value otherwise.

### GetSyncStatusOk

`func (o *GithubRepoView) GetSyncStatusOk() (*string, bool)`

GetSyncStatusOk returns a tuple with the SyncStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncStatus

`func (o *GithubRepoView) SetSyncStatus(v string)`

SetSyncStatus sets SyncStatus field to given value.

### HasSyncStatus

`func (o *GithubRepoView) HasSyncStatus() bool`

HasSyncStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


