# GithubSearchHit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloneUrl** | Pointer to **string** | CloneURL is the repository&#39;s https git remote. | [optional] 
**DefaultBranch** | Pointer to **string** | DefaultBranch is the branch a clone checks out. | [optional] 
**Description** | Pointer to **string** | Description is the blurb the repository&#39;s owner wrote. Empty when it has none. | [optional] 
**FullName** | Pointer to **string** | FullName is the repository&#39;s \&quot;owner/repo\&quot; on GitHub. Finding it here does NOT make it forkable: githubFork takes a repo the org&#39;s installation was granted, and a hit from the public index usually is not one. | [optional] 
**HtmlUrl** | Pointer to **string** | HTMLURL is the repository&#39;s page on github.com. | [optional] 
**Language** | Pointer to **string** | Language is the primary language GitHub detected from the file mix (\&quot;Go\&quot;, \&quot;TypeScript\&quot;). Empty when GitHub attributes none. | [optional] 
**Private** | Pointer to **bool** | Private is GitHub&#39;s visibility flag, passed through. This op reads the public index — the org&#39;s token only charges the rate limit to the installation — so it is false for everything a search can reach. | [optional] 
**Stars** | Pointer to **int64** | Stars is GitHub&#39;s stargazers_count as the SEARCH INDEX held it when the query ran — a snapshot, not a live count off the repository. | [optional] 

## Methods

### NewGithubSearchHit

`func NewGithubSearchHit() *GithubSearchHit`

NewGithubSearchHit instantiates a new GithubSearchHit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGithubSearchHitWithDefaults

`func NewGithubSearchHitWithDefaults() *GithubSearchHit`

NewGithubSearchHitWithDefaults instantiates a new GithubSearchHit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloneUrl

`func (o *GithubSearchHit) GetCloneUrl() string`

GetCloneUrl returns the CloneUrl field if non-nil, zero value otherwise.

### GetCloneUrlOk

`func (o *GithubSearchHit) GetCloneUrlOk() (*string, bool)`

GetCloneUrlOk returns a tuple with the CloneUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloneUrl

`func (o *GithubSearchHit) SetCloneUrl(v string)`

SetCloneUrl sets CloneUrl field to given value.

### HasCloneUrl

`func (o *GithubSearchHit) HasCloneUrl() bool`

HasCloneUrl returns a boolean if a field has been set.

### GetDefaultBranch

`func (o *GithubSearchHit) GetDefaultBranch() string`

GetDefaultBranch returns the DefaultBranch field if non-nil, zero value otherwise.

### GetDefaultBranchOk

`func (o *GithubSearchHit) GetDefaultBranchOk() (*string, bool)`

GetDefaultBranchOk returns a tuple with the DefaultBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultBranch

`func (o *GithubSearchHit) SetDefaultBranch(v string)`

SetDefaultBranch sets DefaultBranch field to given value.

### HasDefaultBranch

`func (o *GithubSearchHit) HasDefaultBranch() bool`

HasDefaultBranch returns a boolean if a field has been set.

### GetDescription

`func (o *GithubSearchHit) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GithubSearchHit) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GithubSearchHit) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GithubSearchHit) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFullName

`func (o *GithubSearchHit) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *GithubSearchHit) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *GithubSearchHit) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *GithubSearchHit) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### GetHtmlUrl

`func (o *GithubSearchHit) GetHtmlUrl() string`

GetHtmlUrl returns the HtmlUrl field if non-nil, zero value otherwise.

### GetHtmlUrlOk

`func (o *GithubSearchHit) GetHtmlUrlOk() (*string, bool)`

GetHtmlUrlOk returns a tuple with the HtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlUrl

`func (o *GithubSearchHit) SetHtmlUrl(v string)`

SetHtmlUrl sets HtmlUrl field to given value.

### HasHtmlUrl

`func (o *GithubSearchHit) HasHtmlUrl() bool`

HasHtmlUrl returns a boolean if a field has been set.

### GetLanguage

`func (o *GithubSearchHit) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *GithubSearchHit) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *GithubSearchHit) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *GithubSearchHit) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetPrivate

`func (o *GithubSearchHit) GetPrivate() bool`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *GithubSearchHit) GetPrivateOk() (*bool, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *GithubSearchHit) SetPrivate(v bool)`

SetPrivate sets Private field to given value.

### HasPrivate

`func (o *GithubSearchHit) HasPrivate() bool`

HasPrivate returns a boolean if a field has been set.

### GetStars

`func (o *GithubSearchHit) GetStars() int64`

GetStars returns the Stars field if non-nil, zero value otherwise.

### GetStarsOk

`func (o *GithubSearchHit) GetStarsOk() (*int64, bool)`

GetStarsOk returns a tuple with the Stars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStars

`func (o *GithubSearchHit) SetStars(v int64)`

SetStars sets Stars field to given value.

### HasStars

`func (o *GithubSearchHit) HasStars() bool`

HasStars returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


