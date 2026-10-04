# PlatformGitOrigin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Branch** | Pointer to **string** | Branch is the branch to build; defaults to &#x60;main&#x60; for a git source. | [optional] 
**Url** | Pointer to **string** | URL is the repository clone URL. Required for source &#x60;git&#x60;, and validated against the same allowlist the privileged build enforces. | [optional] 

## Methods

### NewPlatformGitOrigin

`func NewPlatformGitOrigin() *PlatformGitOrigin`

NewPlatformGitOrigin instantiates a new PlatformGitOrigin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformGitOriginWithDefaults

`func NewPlatformGitOriginWithDefaults() *PlatformGitOrigin`

NewPlatformGitOriginWithDefaults instantiates a new PlatformGitOrigin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBranch

`func (o *PlatformGitOrigin) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *PlatformGitOrigin) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *PlatformGitOrigin) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *PlatformGitOrigin) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetUrl

`func (o *PlatformGitOrigin) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PlatformGitOrigin) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PlatformGitOrigin) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *PlatformGitOrigin) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


