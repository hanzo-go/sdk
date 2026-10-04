# GitMirrorTargetReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | Pointer to **string** | Host is an optional assertion of the target&#39;s hostname. The authoritative host is the one in URL; a value that disagrees with it is refused. | [optional] 
**Name** | Pointer to **string** | Name is the repo whose advanced refs are pushed downstream, from the :name path segment. | [optional] 
**Url** | Pointer to **string** | URL is the downstream https git remote. Must be https to an allowlisted host (github.com / gitlab.com); any embedded credentials are stripped. Required. | [optional] 

## Methods

### NewGitMirrorTargetReq

`func NewGitMirrorTargetReq() *GitMirrorTargetReq`

NewGitMirrorTargetReq instantiates a new GitMirrorTargetReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitMirrorTargetReqWithDefaults

`func NewGitMirrorTargetReqWithDefaults() *GitMirrorTargetReq`

NewGitMirrorTargetReqWithDefaults instantiates a new GitMirrorTargetReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *GitMirrorTargetReq) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *GitMirrorTargetReq) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *GitMirrorTargetReq) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *GitMirrorTargetReq) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetName

`func (o *GitMirrorTargetReq) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GitMirrorTargetReq) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GitMirrorTargetReq) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GitMirrorTargetReq) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUrl

`func (o *GitMirrorTargetReq) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *GitMirrorTargetReq) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *GitMirrorTargetReq) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *GitMirrorTargetReq) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


