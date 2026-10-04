# ProviderGithubPagesUpdateReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Branch** | Pointer to **string** | Branch switches the legacy source branch. Empty leaves the source alone. | [optional] 
**BuildType** | Pointer to **string** | BuildType switches the builder: \&quot;legacy\&quot; or \&quot;workflow\&quot;. Empty leaves it. | [optional] 
**Cname** | Pointer to **string** | CNAME is the custom domain. Omit to leave it alone, \&quot;\&quot; to clear it, or a valid FQDN to set it. | [optional] 
**HttpsEnforced** | Pointer to **bool** | HTTPSEnforced toggles GitHub&#39;s enforce-HTTPS bit. Omit to leave it alone. | [optional] 
**Path** | Pointer to **string** | Path is the source directory to pair with Branch: \&quot;/\&quot; (the default) or \&quot;/docs\&quot;. Read only when Branch is given. | [optional] 
**Repo** | Pointer to **string** | Repo is the repository, from the :repo path segment. | [optional] 

## Methods

### NewProviderGithubPagesUpdateReq

`func NewProviderGithubPagesUpdateReq() *ProviderGithubPagesUpdateReq`

NewProviderGithubPagesUpdateReq instantiates a new ProviderGithubPagesUpdateReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderGithubPagesUpdateReqWithDefaults

`func NewProviderGithubPagesUpdateReqWithDefaults() *ProviderGithubPagesUpdateReq`

NewProviderGithubPagesUpdateReqWithDefaults instantiates a new ProviderGithubPagesUpdateReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBranch

`func (o *ProviderGithubPagesUpdateReq) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *ProviderGithubPagesUpdateReq) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *ProviderGithubPagesUpdateReq) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *ProviderGithubPagesUpdateReq) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetBuildType

`func (o *ProviderGithubPagesUpdateReq) GetBuildType() string`

GetBuildType returns the BuildType field if non-nil, zero value otherwise.

### GetBuildTypeOk

`func (o *ProviderGithubPagesUpdateReq) GetBuildTypeOk() (*string, bool)`

GetBuildTypeOk returns a tuple with the BuildType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildType

`func (o *ProviderGithubPagesUpdateReq) SetBuildType(v string)`

SetBuildType sets BuildType field to given value.

### HasBuildType

`func (o *ProviderGithubPagesUpdateReq) HasBuildType() bool`

HasBuildType returns a boolean if a field has been set.

### GetCname

`func (o *ProviderGithubPagesUpdateReq) GetCname() string`

GetCname returns the Cname field if non-nil, zero value otherwise.

### GetCnameOk

`func (o *ProviderGithubPagesUpdateReq) GetCnameOk() (*string, bool)`

GetCnameOk returns a tuple with the Cname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCname

`func (o *ProviderGithubPagesUpdateReq) SetCname(v string)`

SetCname sets Cname field to given value.

### HasCname

`func (o *ProviderGithubPagesUpdateReq) HasCname() bool`

HasCname returns a boolean if a field has been set.

### GetHttpsEnforced

`func (o *ProviderGithubPagesUpdateReq) GetHttpsEnforced() bool`

GetHttpsEnforced returns the HttpsEnforced field if non-nil, zero value otherwise.

### GetHttpsEnforcedOk

`func (o *ProviderGithubPagesUpdateReq) GetHttpsEnforcedOk() (*bool, bool)`

GetHttpsEnforcedOk returns a tuple with the HttpsEnforced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpsEnforced

`func (o *ProviderGithubPagesUpdateReq) SetHttpsEnforced(v bool)`

SetHttpsEnforced sets HttpsEnforced field to given value.

### HasHttpsEnforced

`func (o *ProviderGithubPagesUpdateReq) HasHttpsEnforced() bool`

HasHttpsEnforced returns a boolean if a field has been set.

### GetPath

`func (o *ProviderGithubPagesUpdateReq) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ProviderGithubPagesUpdateReq) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ProviderGithubPagesUpdateReq) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *ProviderGithubPagesUpdateReq) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetRepo

`func (o *ProviderGithubPagesUpdateReq) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *ProviderGithubPagesUpdateReq) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *ProviderGithubPagesUpdateReq) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *ProviderGithubPagesUpdateReq) HasRepo() bool`

HasRepo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


