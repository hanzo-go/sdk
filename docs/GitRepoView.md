# GitRepoView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Branches** | Pointer to **[]string** | Branches are the repo&#39;s branch names. Read live, so the detail view carries them and a list row does not. | [optional] 
**CloneUrl** | Pointer to **string** | CloneURL is the HTTPS smart-HTTP remote &#x60;git clone&#x60; takes. | [optional] 
**CreatedAt** | Pointer to **string** | CreatedAt is RFC 3339 UTC. | [optional] 
**DefaultBranch** | Pointer to **string** | DefaultBranch is where HEAD points on a fresh repo (\&quot;main\&quot;). | [optional] 
**Description** | Pointer to **string** | Description is the caller-supplied blurb (max 4KiB). | [optional] 
**Head** | Pointer to **string** | Head is the resolved HEAD commit, empty on an empty repo. | [optional] 
**Id** | Pointer to **string** | ID is the repo&#39;s stable, prefixed identifier (\&quot;repo_\&quot; + 128 random bits). | [optional] 
**Name** | Pointer to **string** | Name is the org-unique handle, and the last path segment of both URLs below. | [optional] 
**Org** | Pointer to **string** | Org owns the repo — the gateway-minted X-Org-Id, and the isolation key. | [optional] 
**Project** | Pointer to **string** | Project is the optional sub-scope the repo lives in; absent for the org&#39;s default scope. | [optional] 
**Public** | Pointer to **bool** | Public grants ANONYMOUS read (fetch) only; push and the whole control plane stay org-authed. | [optional] 
**SizeBytes** | Pointer to **int64** | SizeBytes is the repo&#39;s measured on-disk size, re-measured on create, after each push, and after a gc. This is the number billing meters. | [optional] 
**SshUrl** | Pointer to **string** | SSHURL is the scp-style SSH remote (git@host:org/repo.git). | [optional] 
**UpdatedAt** | Pointer to **string** | UpdatedAt is RFC 3339 UTC, empty until the first write. | [optional] 

## Methods

### NewGitRepoView

`func NewGitRepoView() *GitRepoView`

NewGitRepoView instantiates a new GitRepoView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitRepoViewWithDefaults

`func NewGitRepoViewWithDefaults() *GitRepoView`

NewGitRepoViewWithDefaults instantiates a new GitRepoView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBranches

`func (o *GitRepoView) GetBranches() []string`

GetBranches returns the Branches field if non-nil, zero value otherwise.

### GetBranchesOk

`func (o *GitRepoView) GetBranchesOk() (*[]string, bool)`

GetBranchesOk returns a tuple with the Branches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranches

`func (o *GitRepoView) SetBranches(v []string)`

SetBranches sets Branches field to given value.

### HasBranches

`func (o *GitRepoView) HasBranches() bool`

HasBranches returns a boolean if a field has been set.

### GetCloneUrl

`func (o *GitRepoView) GetCloneUrl() string`

GetCloneUrl returns the CloneUrl field if non-nil, zero value otherwise.

### GetCloneUrlOk

`func (o *GitRepoView) GetCloneUrlOk() (*string, bool)`

GetCloneUrlOk returns a tuple with the CloneUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloneUrl

`func (o *GitRepoView) SetCloneUrl(v string)`

SetCloneUrl sets CloneUrl field to given value.

### HasCloneUrl

`func (o *GitRepoView) HasCloneUrl() bool`

HasCloneUrl returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GitRepoView) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GitRepoView) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GitRepoView) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GitRepoView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDefaultBranch

`func (o *GitRepoView) GetDefaultBranch() string`

GetDefaultBranch returns the DefaultBranch field if non-nil, zero value otherwise.

### GetDefaultBranchOk

`func (o *GitRepoView) GetDefaultBranchOk() (*string, bool)`

GetDefaultBranchOk returns a tuple with the DefaultBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultBranch

`func (o *GitRepoView) SetDefaultBranch(v string)`

SetDefaultBranch sets DefaultBranch field to given value.

### HasDefaultBranch

`func (o *GitRepoView) HasDefaultBranch() bool`

HasDefaultBranch returns a boolean if a field has been set.

### GetDescription

`func (o *GitRepoView) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GitRepoView) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GitRepoView) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GitRepoView) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHead

`func (o *GitRepoView) GetHead() string`

GetHead returns the Head field if non-nil, zero value otherwise.

### GetHeadOk

`func (o *GitRepoView) GetHeadOk() (*string, bool)`

GetHeadOk returns a tuple with the Head field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHead

`func (o *GitRepoView) SetHead(v string)`

SetHead sets Head field to given value.

### HasHead

`func (o *GitRepoView) HasHead() bool`

HasHead returns a boolean if a field has been set.

### GetId

`func (o *GitRepoView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GitRepoView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GitRepoView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GitRepoView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GitRepoView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GitRepoView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GitRepoView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GitRepoView) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrg

`func (o *GitRepoView) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *GitRepoView) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *GitRepoView) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *GitRepoView) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetProject

`func (o *GitRepoView) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *GitRepoView) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *GitRepoView) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *GitRepoView) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetPublic

`func (o *GitRepoView) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *GitRepoView) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *GitRepoView) SetPublic(v bool)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *GitRepoView) HasPublic() bool`

HasPublic returns a boolean if a field has been set.

### GetSizeBytes

`func (o *GitRepoView) GetSizeBytes() int64`

GetSizeBytes returns the SizeBytes field if non-nil, zero value otherwise.

### GetSizeBytesOk

`func (o *GitRepoView) GetSizeBytesOk() (*int64, bool)`

GetSizeBytesOk returns a tuple with the SizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeBytes

`func (o *GitRepoView) SetSizeBytes(v int64)`

SetSizeBytes sets SizeBytes field to given value.

### HasSizeBytes

`func (o *GitRepoView) HasSizeBytes() bool`

HasSizeBytes returns a boolean if a field has been set.

### GetSshUrl

`func (o *GitRepoView) GetSshUrl() string`

GetSshUrl returns the SshUrl field if non-nil, zero value otherwise.

### GetSshUrlOk

`func (o *GitRepoView) GetSshUrlOk() (*string, bool)`

GetSshUrlOk returns a tuple with the SshUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshUrl

`func (o *GitRepoView) SetSshUrl(v string)`

SetSshUrl sets SshUrl field to given value.

### HasSshUrl

`func (o *GitRepoView) HasSshUrl() bool`

HasSshUrl returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GitRepoView) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GitRepoView) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GitRepoView) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GitRepoView) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


