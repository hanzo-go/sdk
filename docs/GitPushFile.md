# GitPushFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to **string** | Content is the file&#39;s bytes, carried per Encoding. | [optional] 
**Encoding** | Pointer to **string** | Encoding is \&quot;base64\&quot;, or \&quot;utf-8\&quot; (the default, also \&quot;utf8\&quot; / \&quot;text\&quot;). | [optional] 
**Path** | Pointer to **string** | Path is repo-relative. Absolute or traversing paths, and paths inside .git, are refused. A path that changes CI (.hanzo/workflows, .github/workflows, .forgejo, .gitea) needs an org admin. | [optional] 

## Methods

### NewGitPushFile

`func NewGitPushFile() *GitPushFile`

NewGitPushFile instantiates a new GitPushFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitPushFileWithDefaults

`func NewGitPushFileWithDefaults() *GitPushFile`

NewGitPushFileWithDefaults instantiates a new GitPushFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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

### HasContent

`func (o *GitPushFile) HasContent() bool`

HasContent returns a boolean if a field has been set.

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

### HasPath

`func (o *GitPushFile) HasPath() bool`

HasPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


