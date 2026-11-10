# GitLinkReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name is the repository, from the :name path segment. It is made when the org has none of that name. | [optional] 
**Project** | Pointer to **string** | Project is the sub-scope to make it in; empty uses the caller&#39;s own. | [optional] 
**Source** | Pointer to **string** | Source is where its history lives: an http(s) clone address, such as https://github.com/acme/widgets.git. Its host must let the caller read it. | [optional] 

## Methods

### NewGitLinkReq

`func NewGitLinkReq() *GitLinkReq`

NewGitLinkReq instantiates a new GitLinkReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitLinkReqWithDefaults

`func NewGitLinkReqWithDefaults() *GitLinkReq`

NewGitLinkReqWithDefaults instantiates a new GitLinkReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GitLinkReq) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GitLinkReq) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GitLinkReq) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GitLinkReq) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProject

`func (o *GitLinkReq) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *GitLinkReq) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *GitLinkReq) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *GitLinkReq) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetSource

`func (o *GitLinkReq) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *GitLinkReq) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *GitLinkReq) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *GitLinkReq) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


