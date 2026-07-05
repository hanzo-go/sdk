# FlowConnectGitRepoRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RemoteUrl** | **string** |  | 
**Branch** | Pointer to **string** |  | [optional] 
**Slug** | Pointer to **string** |  | [optional] 
**SshPrivateKey** | Pointer to **string** |  | [optional] 

## Methods

### NewFlowConnectGitRepoRequest

`func NewFlowConnectGitRepoRequest(remoteUrl string, ) *FlowConnectGitRepoRequest`

NewFlowConnectGitRepoRequest instantiates a new FlowConnectGitRepoRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlowConnectGitRepoRequestWithDefaults

`func NewFlowConnectGitRepoRequestWithDefaults() *FlowConnectGitRepoRequest`

NewFlowConnectGitRepoRequestWithDefaults instantiates a new FlowConnectGitRepoRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRemoteUrl

`func (o *FlowConnectGitRepoRequest) GetRemoteUrl() string`

GetRemoteUrl returns the RemoteUrl field if non-nil, zero value otherwise.

### GetRemoteUrlOk

`func (o *FlowConnectGitRepoRequest) GetRemoteUrlOk() (*string, bool)`

GetRemoteUrlOk returns a tuple with the RemoteUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteUrl

`func (o *FlowConnectGitRepoRequest) SetRemoteUrl(v string)`

SetRemoteUrl sets RemoteUrl field to given value.


### GetBranch

`func (o *FlowConnectGitRepoRequest) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *FlowConnectGitRepoRequest) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *FlowConnectGitRepoRequest) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *FlowConnectGitRepoRequest) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetSlug

`func (o *FlowConnectGitRepoRequest) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *FlowConnectGitRepoRequest) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *FlowConnectGitRepoRequest) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *FlowConnectGitRepoRequest) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetSshPrivateKey

`func (o *FlowConnectGitRepoRequest) GetSshPrivateKey() string`

GetSshPrivateKey returns the SshPrivateKey field if non-nil, zero value otherwise.

### GetSshPrivateKeyOk

`func (o *FlowConnectGitRepoRequest) GetSshPrivateKeyOk() (*string, bool)`

GetSshPrivateKeyOk returns a tuple with the SshPrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPrivateKey

`func (o *FlowConnectGitRepoRequest) SetSshPrivateKey(v string)`

SetSshPrivateKey sets SshPrivateKey field to given value.

### HasSshPrivateKey

`func (o *FlowConnectGitRepoRequest) HasSshPrivateKey() bool`

HasSshPrivateKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


