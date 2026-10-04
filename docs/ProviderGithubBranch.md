# ProviderGithubBranch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Commit** | Pointer to **string** | Commit is the sha the branch points at. | [optional] 
**Default** | Pointer to **bool** | Default is whether it is the repository&#39;s default branch, which always leads the first page when it matches. | [optional] 
**Name** | Pointer to **string** | Name is the branch&#39;s short name (\&quot;main\&quot;, \&quot;feat/login\&quot;). | [optional] 

## Methods

### NewProviderGithubBranch

`func NewProviderGithubBranch() *ProviderGithubBranch`

NewProviderGithubBranch instantiates a new ProviderGithubBranch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderGithubBranchWithDefaults

`func NewProviderGithubBranchWithDefaults() *ProviderGithubBranch`

NewProviderGithubBranchWithDefaults instantiates a new ProviderGithubBranch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommit

`func (o *ProviderGithubBranch) GetCommit() string`

GetCommit returns the Commit field if non-nil, zero value otherwise.

### GetCommitOk

`func (o *ProviderGithubBranch) GetCommitOk() (*string, bool)`

GetCommitOk returns a tuple with the Commit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommit

`func (o *ProviderGithubBranch) SetCommit(v string)`

SetCommit sets Commit field to given value.

### HasCommit

`func (o *ProviderGithubBranch) HasCommit() bool`

HasCommit returns a boolean if a field has been set.

### GetDefault

`func (o *ProviderGithubBranch) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *ProviderGithubBranch) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *ProviderGithubBranch) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *ProviderGithubBranch) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetName

`func (o *ProviderGithubBranch) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProviderGithubBranch) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProviderGithubBranch) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProviderGithubBranch) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


