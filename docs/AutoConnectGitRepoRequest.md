# AutoConnectGitRepoRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RemoteUrl** | **string** |  | 
**Branch** | Pointer to **string** |  | [optional] 
**SshPrivateKey** | Pointer to **string** |  | [optional] 

## Methods

### NewAutoConnectGitRepoRequest

`func NewAutoConnectGitRepoRequest(remoteUrl string, ) *AutoConnectGitRepoRequest`

NewAutoConnectGitRepoRequest instantiates a new AutoConnectGitRepoRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoConnectGitRepoRequestWithDefaults

`func NewAutoConnectGitRepoRequestWithDefaults() *AutoConnectGitRepoRequest`

NewAutoConnectGitRepoRequestWithDefaults instantiates a new AutoConnectGitRepoRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRemoteUrl

`func (o *AutoConnectGitRepoRequest) GetRemoteUrl() string`

GetRemoteUrl returns the RemoteUrl field if non-nil, zero value otherwise.

### GetRemoteUrlOk

`func (o *AutoConnectGitRepoRequest) GetRemoteUrlOk() (*string, bool)`

GetRemoteUrlOk returns a tuple with the RemoteUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteUrl

`func (o *AutoConnectGitRepoRequest) SetRemoteUrl(v string)`

SetRemoteUrl sets RemoteUrl field to given value.


### GetBranch

`func (o *AutoConnectGitRepoRequest) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *AutoConnectGitRepoRequest) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *AutoConnectGitRepoRequest) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *AutoConnectGitRepoRequest) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetSshPrivateKey

`func (o *AutoConnectGitRepoRequest) GetSshPrivateKey() string`

GetSshPrivateKey returns the SshPrivateKey field if non-nil, zero value otherwise.

### GetSshPrivateKeyOk

`func (o *AutoConnectGitRepoRequest) GetSshPrivateKeyOk() (*string, bool)`

GetSshPrivateKeyOk returns a tuple with the SshPrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPrivateKey

`func (o *AutoConnectGitRepoRequest) SetSshPrivateKey(v string)`

SetSshPrivateKey sets SshPrivateKey field to given value.

### HasSshPrivateKey

`func (o *AutoConnectGitRepoRequest) HasSshPrivateKey() bool`

HasSshPrivateKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


