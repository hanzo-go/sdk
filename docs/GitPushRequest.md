# GitPushRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Branch** | Pointer to **string** | Target branch (default main) | [optional] 
**Message** | Pointer to **string** | Commit message | [optional] 
**Files** | [**[]GitPushFile**](GitPushFile.md) |  | 

## Methods

### NewGitPushRequest

`func NewGitPushRequest(files []GitPushFile, ) *GitPushRequest`

NewGitPushRequest instantiates a new GitPushRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitPushRequestWithDefaults

`func NewGitPushRequestWithDefaults() *GitPushRequest`

NewGitPushRequestWithDefaults instantiates a new GitPushRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBranch

`func (o *GitPushRequest) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *GitPushRequest) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *GitPushRequest) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *GitPushRequest) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetMessage

`func (o *GitPushRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GitPushRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GitPushRequest) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GitPushRequest) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetFiles

`func (o *GitPushRequest) GetFiles() []GitPushFile`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *GitPushRequest) GetFilesOk() (*[]GitPushFile, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *GitPushRequest) SetFiles(v []GitPushFile)`

SetFiles sets Files field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


