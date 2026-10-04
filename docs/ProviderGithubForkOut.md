# ProviderGithubForkOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloneUrl** | Pointer to **string** | CloneURL is the fork&#39;s https git remote. GitHub populates a new fork in the background, so a clone issued the moment this answers can still find it empty. | [optional] 
**DefaultBranch** | Pointer to **string** | DefaultBranch is the branch the fork checks out, inherited from upstream. | [optional] 
**Existing** | Pointer to **bool** | Existing reports that the fork was already there. GitHub answers 202 either way, so without this a caller cannot tell \&quot;made you one\&quot; from \&quot;you had one\&quot;. | [optional] 
**FullName** | Pointer to **string** | FullName is the fork&#39;s \&quot;owner/repo\&quot;. The owner is the account it landed in — the request&#39;s org, or the installation&#39;s own account when none was named. | [optional] 
**HtmlUrl** | Pointer to **string** | HTMLURL is the fork&#39;s page on github.com. | [optional] 

## Methods

### NewProviderGithubForkOut

`func NewProviderGithubForkOut() *ProviderGithubForkOut`

NewProviderGithubForkOut instantiates a new ProviderGithubForkOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderGithubForkOutWithDefaults

`func NewProviderGithubForkOutWithDefaults() *ProviderGithubForkOut`

NewProviderGithubForkOutWithDefaults instantiates a new ProviderGithubForkOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloneUrl

`func (o *ProviderGithubForkOut) GetCloneUrl() string`

GetCloneUrl returns the CloneUrl field if non-nil, zero value otherwise.

### GetCloneUrlOk

`func (o *ProviderGithubForkOut) GetCloneUrlOk() (*string, bool)`

GetCloneUrlOk returns a tuple with the CloneUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloneUrl

`func (o *ProviderGithubForkOut) SetCloneUrl(v string)`

SetCloneUrl sets CloneUrl field to given value.

### HasCloneUrl

`func (o *ProviderGithubForkOut) HasCloneUrl() bool`

HasCloneUrl returns a boolean if a field has been set.

### GetDefaultBranch

`func (o *ProviderGithubForkOut) GetDefaultBranch() string`

GetDefaultBranch returns the DefaultBranch field if non-nil, zero value otherwise.

### GetDefaultBranchOk

`func (o *ProviderGithubForkOut) GetDefaultBranchOk() (*string, bool)`

GetDefaultBranchOk returns a tuple with the DefaultBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultBranch

`func (o *ProviderGithubForkOut) SetDefaultBranch(v string)`

SetDefaultBranch sets DefaultBranch field to given value.

### HasDefaultBranch

`func (o *ProviderGithubForkOut) HasDefaultBranch() bool`

HasDefaultBranch returns a boolean if a field has been set.

### GetExisting

`func (o *ProviderGithubForkOut) GetExisting() bool`

GetExisting returns the Existing field if non-nil, zero value otherwise.

### GetExistingOk

`func (o *ProviderGithubForkOut) GetExistingOk() (*bool, bool)`

GetExistingOk returns a tuple with the Existing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExisting

`func (o *ProviderGithubForkOut) SetExisting(v bool)`

SetExisting sets Existing field to given value.

### HasExisting

`func (o *ProviderGithubForkOut) HasExisting() bool`

HasExisting returns a boolean if a field has been set.

### GetFullName

`func (o *ProviderGithubForkOut) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *ProviderGithubForkOut) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *ProviderGithubForkOut) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *ProviderGithubForkOut) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### GetHtmlUrl

`func (o *ProviderGithubForkOut) GetHtmlUrl() string`

GetHtmlUrl returns the HtmlUrl field if non-nil, zero value otherwise.

### GetHtmlUrlOk

`func (o *ProviderGithubForkOut) GetHtmlUrlOk() (*string, bool)`

GetHtmlUrlOk returns a tuple with the HtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlUrl

`func (o *ProviderGithubForkOut) SetHtmlUrl(v string)`

SetHtmlUrl sets HtmlUrl field to given value.

### HasHtmlUrl

`func (o *ProviderGithubForkOut) HasHtmlUrl() bool`

HasHtmlUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


