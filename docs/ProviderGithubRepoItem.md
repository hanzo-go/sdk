# ProviderGithubRepoItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultBranch** | Pointer to **string** | DefaultBranch is the branch GitHub opens it on. | [optional] 
**FullName** | Pointer to **string** | FullName is owner/name, the address a coding run, a branch listing and a deploy take. | [optional] 
**InstallationId** | Pointer to **int64** | InstallationID is the installation of the Hanzo Platform App that reaches it — the grant a run on it is minted from. | [optional] 
**Name** | Pointer to **string** | Name is its name within that account. | [optional] 
**Owner** | Pointer to **string** | Owner is the GitHub account that holds it. | [optional] 
**Private** | Pointer to **bool** | Private is GitHub&#39;s visibility bit. | [optional] 
**PushedAt** | Pointer to **string** | PushedAt is when anything was last pushed to it, RFC 3339 UTC. Absent for a repository nothing was ever pushed to. | [optional] 

## Methods

### NewProviderGithubRepoItem

`func NewProviderGithubRepoItem() *ProviderGithubRepoItem`

NewProviderGithubRepoItem instantiates a new ProviderGithubRepoItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderGithubRepoItemWithDefaults

`func NewProviderGithubRepoItemWithDefaults() *ProviderGithubRepoItem`

NewProviderGithubRepoItemWithDefaults instantiates a new ProviderGithubRepoItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultBranch

`func (o *ProviderGithubRepoItem) GetDefaultBranch() string`

GetDefaultBranch returns the DefaultBranch field if non-nil, zero value otherwise.

### GetDefaultBranchOk

`func (o *ProviderGithubRepoItem) GetDefaultBranchOk() (*string, bool)`

GetDefaultBranchOk returns a tuple with the DefaultBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultBranch

`func (o *ProviderGithubRepoItem) SetDefaultBranch(v string)`

SetDefaultBranch sets DefaultBranch field to given value.

### HasDefaultBranch

`func (o *ProviderGithubRepoItem) HasDefaultBranch() bool`

HasDefaultBranch returns a boolean if a field has been set.

### GetFullName

`func (o *ProviderGithubRepoItem) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *ProviderGithubRepoItem) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *ProviderGithubRepoItem) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *ProviderGithubRepoItem) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### GetInstallationId

`func (o *ProviderGithubRepoItem) GetInstallationId() int64`

GetInstallationId returns the InstallationId field if non-nil, zero value otherwise.

### GetInstallationIdOk

`func (o *ProviderGithubRepoItem) GetInstallationIdOk() (*int64, bool)`

GetInstallationIdOk returns a tuple with the InstallationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallationId

`func (o *ProviderGithubRepoItem) SetInstallationId(v int64)`

SetInstallationId sets InstallationId field to given value.

### HasInstallationId

`func (o *ProviderGithubRepoItem) HasInstallationId() bool`

HasInstallationId returns a boolean if a field has been set.

### GetName

`func (o *ProviderGithubRepoItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProviderGithubRepoItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProviderGithubRepoItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProviderGithubRepoItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwner

`func (o *ProviderGithubRepoItem) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ProviderGithubRepoItem) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ProviderGithubRepoItem) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ProviderGithubRepoItem) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetPrivate

`func (o *ProviderGithubRepoItem) GetPrivate() bool`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *ProviderGithubRepoItem) GetPrivateOk() (*bool, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *ProviderGithubRepoItem) SetPrivate(v bool)`

SetPrivate sets Private field to given value.

### HasPrivate

`func (o *ProviderGithubRepoItem) HasPrivate() bool`

HasPrivate returns a boolean if a field has been set.

### GetPushedAt

`func (o *ProviderGithubRepoItem) GetPushedAt() string`

GetPushedAt returns the PushedAt field if non-nil, zero value otherwise.

### GetPushedAtOk

`func (o *ProviderGithubRepoItem) GetPushedAtOk() (*string, bool)`

GetPushedAtOk returns a tuple with the PushedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushedAt

`func (o *ProviderGithubRepoItem) SetPushedAt(v string)`

SetPushedAt sets PushedAt field to given value.

### HasPushedAt

`func (o *ProviderGithubRepoItem) HasPushedAt() bool`

HasPushedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


