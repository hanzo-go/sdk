# CloudGithubRepoView

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

### NewCloudGithubRepoView

`func NewCloudGithubRepoView() *CloudGithubRepoView`

NewCloudGithubRepoView instantiates a new CloudGithubRepoView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudGithubRepoViewWithDefaults

`func NewCloudGithubRepoViewWithDefaults() *CloudGithubRepoView`

NewCloudGithubRepoViewWithDefaults instantiates a new CloudGithubRepoView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultBranch

`func (o *CloudGithubRepoView) GetDefaultBranch() string`

GetDefaultBranch returns the DefaultBranch field if non-nil, zero value otherwise.

### GetDefaultBranchOk

`func (o *CloudGithubRepoView) GetDefaultBranchOk() (*string, bool)`

GetDefaultBranchOk returns a tuple with the DefaultBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultBranch

`func (o *CloudGithubRepoView) SetDefaultBranch(v string)`

SetDefaultBranch sets DefaultBranch field to given value.

### HasDefaultBranch

`func (o *CloudGithubRepoView) HasDefaultBranch() bool`

HasDefaultBranch returns a boolean if a field has been set.

### GetFullName

`func (o *CloudGithubRepoView) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *CloudGithubRepoView) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *CloudGithubRepoView) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *CloudGithubRepoView) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### GetHtmlUrl

`func (o *CloudGithubRepoView) GetHtmlUrl() string`

GetHtmlUrl returns the HtmlUrl field if non-nil, zero value otherwise.

### GetHtmlUrlOk

`func (o *CloudGithubRepoView) GetHtmlUrlOk() (*string, bool)`

GetHtmlUrlOk returns a tuple with the HtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlUrl

`func (o *CloudGithubRepoView) SetHtmlUrl(v string)`

SetHtmlUrl sets HtmlUrl field to given value.

### HasHtmlUrl

`func (o *CloudGithubRepoView) HasHtmlUrl() bool`

HasHtmlUrl returns a boolean if a field has been set.

### GetImported

`func (o *CloudGithubRepoView) GetImported() bool`

GetImported returns the Imported field if non-nil, zero value otherwise.

### GetImportedOk

`func (o *CloudGithubRepoView) GetImportedOk() (*bool, bool)`

GetImportedOk returns a tuple with the Imported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImported

`func (o *CloudGithubRepoView) SetImported(v bool)`

SetImported sets Imported field to given value.

### HasImported

`func (o *CloudGithubRepoView) HasImported() bool`

HasImported returns a boolean if a field has been set.

### GetLastSyncedAt

`func (o *CloudGithubRepoView) GetLastSyncedAt() string`

GetLastSyncedAt returns the LastSyncedAt field if non-nil, zero value otherwise.

### GetLastSyncedAtOk

`func (o *CloudGithubRepoView) GetLastSyncedAtOk() (*string, bool)`

GetLastSyncedAtOk returns a tuple with the LastSyncedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSyncedAt

`func (o *CloudGithubRepoView) SetLastSyncedAt(v string)`

SetLastSyncedAt sets LastSyncedAt field to given value.

### HasLastSyncedAt

`func (o *CloudGithubRepoView) HasLastSyncedAt() bool`

HasLastSyncedAt returns a boolean if a field has been set.

### GetName

`func (o *CloudGithubRepoView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudGithubRepoView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudGithubRepoView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudGithubRepoView) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrivate

`func (o *CloudGithubRepoView) GetPrivate() bool`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *CloudGithubRepoView) GetPrivateOk() (*bool, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *CloudGithubRepoView) SetPrivate(v bool)`

SetPrivate sets Private field to given value.

### HasPrivate

`func (o *CloudGithubRepoView) HasPrivate() bool`

HasPrivate returns a boolean if a field has been set.

### GetSyncStatus

`func (o *CloudGithubRepoView) GetSyncStatus() string`

GetSyncStatus returns the SyncStatus field if non-nil, zero value otherwise.

### GetSyncStatusOk

`func (o *CloudGithubRepoView) GetSyncStatusOk() (*string, bool)`

GetSyncStatusOk returns a tuple with the SyncStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncStatus

`func (o *CloudGithubRepoView) SetSyncStatus(v string)`

SetSyncStatus sets SyncStatus field to given value.

### HasSyncStatus

`func (o *CloudGithubRepoView) HasSyncStatus() bool`

HasSyncStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


