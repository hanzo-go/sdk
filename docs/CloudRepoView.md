# CloudRepoView

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
**SizeBytes** | Pointer to **int32** | SizeBytes is the repo&#39;s measured on-disk size, re-measured on create, after each push, and after a gc. This is the number billing meters. | [optional] 
**SshUrl** | Pointer to **string** | SSHURL is the scp-style SSH remote (git@host:org/repo.git). | [optional] 
**UpdatedAt** | Pointer to **string** | UpdatedAt is RFC 3339 UTC, empty until the first write. | [optional] 

## Methods

### NewCloudRepoView

`func NewCloudRepoView() *CloudRepoView`

NewCloudRepoView instantiates a new CloudRepoView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudRepoViewWithDefaults

`func NewCloudRepoViewWithDefaults() *CloudRepoView`

NewCloudRepoViewWithDefaults instantiates a new CloudRepoView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBranches

`func (o *CloudRepoView) GetBranches() []string`

GetBranches returns the Branches field if non-nil, zero value otherwise.

### GetBranchesOk

`func (o *CloudRepoView) GetBranchesOk() (*[]string, bool)`

GetBranchesOk returns a tuple with the Branches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranches

`func (o *CloudRepoView) SetBranches(v []string)`

SetBranches sets Branches field to given value.

### HasBranches

`func (o *CloudRepoView) HasBranches() bool`

HasBranches returns a boolean if a field has been set.

### GetCloneUrl

`func (o *CloudRepoView) GetCloneUrl() string`

GetCloneUrl returns the CloneUrl field if non-nil, zero value otherwise.

### GetCloneUrlOk

`func (o *CloudRepoView) GetCloneUrlOk() (*string, bool)`

GetCloneUrlOk returns a tuple with the CloneUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloneUrl

`func (o *CloudRepoView) SetCloneUrl(v string)`

SetCloneUrl sets CloneUrl field to given value.

### HasCloneUrl

`func (o *CloudRepoView) HasCloneUrl() bool`

HasCloneUrl returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CloudRepoView) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CloudRepoView) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CloudRepoView) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CloudRepoView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDefaultBranch

`func (o *CloudRepoView) GetDefaultBranch() string`

GetDefaultBranch returns the DefaultBranch field if non-nil, zero value otherwise.

### GetDefaultBranchOk

`func (o *CloudRepoView) GetDefaultBranchOk() (*string, bool)`

GetDefaultBranchOk returns a tuple with the DefaultBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultBranch

`func (o *CloudRepoView) SetDefaultBranch(v string)`

SetDefaultBranch sets DefaultBranch field to given value.

### HasDefaultBranch

`func (o *CloudRepoView) HasDefaultBranch() bool`

HasDefaultBranch returns a boolean if a field has been set.

### GetDescription

`func (o *CloudRepoView) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CloudRepoView) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CloudRepoView) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CloudRepoView) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHead

`func (o *CloudRepoView) GetHead() string`

GetHead returns the Head field if non-nil, zero value otherwise.

### GetHeadOk

`func (o *CloudRepoView) GetHeadOk() (*string, bool)`

GetHeadOk returns a tuple with the Head field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHead

`func (o *CloudRepoView) SetHead(v string)`

SetHead sets Head field to given value.

### HasHead

`func (o *CloudRepoView) HasHead() bool`

HasHead returns a boolean if a field has been set.

### GetId

`func (o *CloudRepoView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudRepoView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudRepoView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudRepoView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CloudRepoView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudRepoView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudRepoView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudRepoView) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrg

`func (o *CloudRepoView) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *CloudRepoView) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *CloudRepoView) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *CloudRepoView) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetProject

`func (o *CloudRepoView) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *CloudRepoView) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *CloudRepoView) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *CloudRepoView) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetPublic

`func (o *CloudRepoView) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *CloudRepoView) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *CloudRepoView) SetPublic(v bool)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *CloudRepoView) HasPublic() bool`

HasPublic returns a boolean if a field has been set.

### GetSizeBytes

`func (o *CloudRepoView) GetSizeBytes() int32`

GetSizeBytes returns the SizeBytes field if non-nil, zero value otherwise.

### GetSizeBytesOk

`func (o *CloudRepoView) GetSizeBytesOk() (*int32, bool)`

GetSizeBytesOk returns a tuple with the SizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeBytes

`func (o *CloudRepoView) SetSizeBytes(v int32)`

SetSizeBytes sets SizeBytes field to given value.

### HasSizeBytes

`func (o *CloudRepoView) HasSizeBytes() bool`

HasSizeBytes returns a boolean if a field has been set.

### GetSshUrl

`func (o *CloudRepoView) GetSshUrl() string`

GetSshUrl returns the SshUrl field if non-nil, zero value otherwise.

### GetSshUrlOk

`func (o *CloudRepoView) GetSshUrlOk() (*string, bool)`

GetSshUrlOk returns a tuple with the SshUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshUrl

`func (o *CloudRepoView) SetSshUrl(v string)`

SetSshUrl sets SshUrl field to given value.

### HasSshUrl

`func (o *CloudRepoView) HasSshUrl() bool`

HasSshUrl returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CloudRepoView) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CloudRepoView) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CloudRepoView) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CloudRepoView) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


