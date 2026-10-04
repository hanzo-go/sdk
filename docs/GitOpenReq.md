# GitOpenReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Base** | Pointer to **string** | Base is the branch the work is proposed INTO, by short name. Defaults to the repo&#39;s default branch, which is where a proposal goes when nobody says otherwise. | [optional] 
**Body** | Pointer to **string** | Body is the longer description. Optional. | [optional] 
**Head** | Pointer to **string** | Head is the branch holding the work, by short name (agent/fix-503). Required, and must already exist. | [optional] 
**Name** | Pointer to **string** | Name is the repo the proposal belongs to, from the :name path segment. | [optional] 
**Title** | Pointer to **string** | Title is the one-line summary of what is being proposed. Required. | [optional] 

## Methods

### NewGitOpenReq

`func NewGitOpenReq() *GitOpenReq`

NewGitOpenReq instantiates a new GitOpenReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitOpenReqWithDefaults

`func NewGitOpenReqWithDefaults() *GitOpenReq`

NewGitOpenReqWithDefaults instantiates a new GitOpenReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBase

`func (o *GitOpenReq) GetBase() string`

GetBase returns the Base field if non-nil, zero value otherwise.

### GetBaseOk

`func (o *GitOpenReq) GetBaseOk() (*string, bool)`

GetBaseOk returns a tuple with the Base field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBase

`func (o *GitOpenReq) SetBase(v string)`

SetBase sets Base field to given value.

### HasBase

`func (o *GitOpenReq) HasBase() bool`

HasBase returns a boolean if a field has been set.

### GetBody

`func (o *GitOpenReq) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *GitOpenReq) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *GitOpenReq) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *GitOpenReq) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetHead

`func (o *GitOpenReq) GetHead() string`

GetHead returns the Head field if non-nil, zero value otherwise.

### GetHeadOk

`func (o *GitOpenReq) GetHeadOk() (*string, bool)`

GetHeadOk returns a tuple with the Head field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHead

`func (o *GitOpenReq) SetHead(v string)`

SetHead sets Head field to given value.

### HasHead

`func (o *GitOpenReq) HasHead() bool`

HasHead returns a boolean if a field has been set.

### GetName

`func (o *GitOpenReq) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GitOpenReq) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GitOpenReq) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GitOpenReq) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTitle

`func (o *GitOpenReq) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *GitOpenReq) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *GitOpenReq) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *GitOpenReq) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


