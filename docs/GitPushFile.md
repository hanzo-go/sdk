# GitPushFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | **string** | Repo-relative path (no leading slash, no traversal) | 
**Content** | **string** | File content; UTF-8 by default, base64 when encoding&#x3D;base64 | 
**Encoding** | Pointer to **string** | Content encoding (default utf-8) | [optional] 

## Methods

### NewGitPushFile

`func NewGitPushFile(path string, content string, ) *GitPushFile`

NewGitPushFile instantiates a new GitPushFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitPushFileWithDefaults

`func NewGitPushFileWithDefaults() *GitPushFile`

NewGitPushFileWithDefaults instantiates a new GitPushFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *GitPushFile) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *GitPushFile) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *GitPushFile) SetPath(v string)`

SetPath sets Path field to given value.


### GetContent

`func (o *GitPushFile) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *GitPushFile) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *GitPushFile) SetContent(v string)`

SetContent sets Content field to given value.


### GetEncoding

`func (o *GitPushFile) GetEncoding() string`

GetEncoding returns the Encoding field if non-nil, zero value otherwise.

### GetEncodingOk

`func (o *GitPushFile) GetEncodingOk() (*string, bool)`

GetEncodingOk returns a tuple with the Encoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncoding

`func (o *GitPushFile) SetEncoding(v string)`

SetEncoding sets Encoding field to given value.

### HasEncoding

`func (o *GitPushFile) HasEncoding() bool`

HasEncoding returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


