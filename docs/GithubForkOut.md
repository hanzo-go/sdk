# GithubForkOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloneUrl** | Pointer to **string** |  | [optional] 
**DefaultBranch** | Pointer to **string** |  | [optional] 
**Existing** | Pointer to **bool** | Existing reports that the fork was already there. GitHub answers 202 either way, so without this a caller cannot tell \&quot;made you one\&quot; from \&quot;you had one\&quot;. | [optional] 
**FullName** | Pointer to **string** |  | [optional] 
**HtmlUrl** | Pointer to **string** |  | [optional] 

## Methods

### NewGithubForkOut

`func NewGithubForkOut() *GithubForkOut`

NewGithubForkOut instantiates a new GithubForkOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGithubForkOutWithDefaults

`func NewGithubForkOutWithDefaults() *GithubForkOut`

NewGithubForkOutWithDefaults instantiates a new GithubForkOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloneUrl

`func (o *GithubForkOut) GetCloneUrl() string`

GetCloneUrl returns the CloneUrl field if non-nil, zero value otherwise.

### GetCloneUrlOk

`func (o *GithubForkOut) GetCloneUrlOk() (*string, bool)`

GetCloneUrlOk returns a tuple with the CloneUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloneUrl

`func (o *GithubForkOut) SetCloneUrl(v string)`

SetCloneUrl sets CloneUrl field to given value.

### HasCloneUrl

`func (o *GithubForkOut) HasCloneUrl() bool`

HasCloneUrl returns a boolean if a field has been set.

### GetDefaultBranch

`func (o *GithubForkOut) GetDefaultBranch() string`

GetDefaultBranch returns the DefaultBranch field if non-nil, zero value otherwise.

### GetDefaultBranchOk

`func (o *GithubForkOut) GetDefaultBranchOk() (*string, bool)`

GetDefaultBranchOk returns a tuple with the DefaultBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultBranch

`func (o *GithubForkOut) SetDefaultBranch(v string)`

SetDefaultBranch sets DefaultBranch field to given value.

### HasDefaultBranch

`func (o *GithubForkOut) HasDefaultBranch() bool`

HasDefaultBranch returns a boolean if a field has been set.

### GetExisting

`func (o *GithubForkOut) GetExisting() bool`

GetExisting returns the Existing field if non-nil, zero value otherwise.

### GetExistingOk

`func (o *GithubForkOut) GetExistingOk() (*bool, bool)`

GetExistingOk returns a tuple with the Existing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExisting

`func (o *GithubForkOut) SetExisting(v bool)`

SetExisting sets Existing field to given value.

### HasExisting

`func (o *GithubForkOut) HasExisting() bool`

HasExisting returns a boolean if a field has been set.

### GetFullName

`func (o *GithubForkOut) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *GithubForkOut) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *GithubForkOut) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *GithubForkOut) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### GetHtmlUrl

`func (o *GithubForkOut) GetHtmlUrl() string`

GetHtmlUrl returns the HtmlUrl field if non-nil, zero value otherwise.

### GetHtmlUrlOk

`func (o *GithubForkOut) GetHtmlUrlOk() (*string, bool)`

GetHtmlUrlOk returns a tuple with the HtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlUrl

`func (o *GithubForkOut) SetHtmlUrl(v string)`

SetHtmlUrl sets HtmlUrl field to given value.

### HasHtmlUrl

`func (o *GithubForkOut) HasHtmlUrl() bool`

HasHtmlUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


