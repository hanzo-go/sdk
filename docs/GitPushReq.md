# GitPushReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Branch** | Pointer to **string** | Branch to advance; empty means \&quot;main\&quot;. A fresh branch that is the repo&#39;s first also becomes HEAD. | [optional] 
**Files** | Pointer to [**[]GitPushFile**](GitPushFile.md) | Files are added to or overwritten on the branch tip — files already there and not listed SURVIVE. At least one, at most 5000, 32 MiB each. | [optional] 
**Message** | Pointer to **string** | Message is the commit message; empty gets a generated one. | [optional] 
**Name** | Pointer to **string** | Name is the repo to push into, from the :name path segment. It is CREATED on first push if it does not exist. | [optional] 

## Methods

### NewGitPushReq

`func NewGitPushReq() *GitPushReq`

NewGitPushReq instantiates a new GitPushReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitPushReqWithDefaults

`func NewGitPushReqWithDefaults() *GitPushReq`

NewGitPushReqWithDefaults instantiates a new GitPushReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBranch

`func (o *GitPushReq) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *GitPushReq) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *GitPushReq) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *GitPushReq) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetFiles

`func (o *GitPushReq) GetFiles() []GitPushFile`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *GitPushReq) GetFilesOk() (*[]GitPushFile, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *GitPushReq) SetFiles(v []GitPushFile)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *GitPushReq) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetMessage

`func (o *GitPushReq) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GitPushReq) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GitPushReq) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GitPushReq) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetName

`func (o *GitPushReq) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GitPushReq) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GitPushReq) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GitPushReq) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


