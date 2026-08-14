# GitSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Branch** | Pointer to **string** |  | [optional] 
**Provider** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewGitSource

`func NewGitSource() *GitSource`

NewGitSource instantiates a new GitSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitSourceWithDefaults

`func NewGitSourceWithDefaults() *GitSource`

NewGitSourceWithDefaults instantiates a new GitSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBranch

`func (o *GitSource) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *GitSource) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *GitSource) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *GitSource) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetProvider

`func (o *GitSource) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *GitSource) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *GitSource) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *GitSource) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetUrl

`func (o *GitSource) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *GitSource) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *GitSource) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *GitSource) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


