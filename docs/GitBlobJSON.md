# GitBlobJSON

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Binary** | Pointer to **bool** | Binary marks content git could not treat as text; it comes back base64. | [optional] 
**Content** | Pointer to **string** | Content is the file&#39;s bytes, empty when Truncated. | [optional] 
**Encoding** | Pointer to **string** | Encoding is how Content is carried: \&quot;utf8\&quot; verbatim, or \&quot;base64\&quot;. | [optional] 
**Path** | Pointer to **string** | Path is the file&#39;s repo-relative path. | [optional] 
**Size** | Pointer to **int64** | Size is the file&#39;s byte length in the repo, whatever was returned below. | [optional] 
**Truncated** | Pointer to **bool** | Truncated marks a file past the 1 MiB view cap. No content is sent — clone the repo for it. | [optional] 

## Methods

### NewGitBlobJSON

`func NewGitBlobJSON() *GitBlobJSON`

NewGitBlobJSON instantiates a new GitBlobJSON object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitBlobJSONWithDefaults

`func NewGitBlobJSONWithDefaults() *GitBlobJSON`

NewGitBlobJSONWithDefaults instantiates a new GitBlobJSON object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBinary

`func (o *GitBlobJSON) GetBinary() bool`

GetBinary returns the Binary field if non-nil, zero value otherwise.

### GetBinaryOk

`func (o *GitBlobJSON) GetBinaryOk() (*bool, bool)`

GetBinaryOk returns a tuple with the Binary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinary

`func (o *GitBlobJSON) SetBinary(v bool)`

SetBinary sets Binary field to given value.

### HasBinary

`func (o *GitBlobJSON) HasBinary() bool`

HasBinary returns a boolean if a field has been set.

### GetContent

`func (o *GitBlobJSON) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *GitBlobJSON) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *GitBlobJSON) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *GitBlobJSON) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetEncoding

`func (o *GitBlobJSON) GetEncoding() string`

GetEncoding returns the Encoding field if non-nil, zero value otherwise.

### GetEncodingOk

`func (o *GitBlobJSON) GetEncodingOk() (*string, bool)`

GetEncodingOk returns a tuple with the Encoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncoding

`func (o *GitBlobJSON) SetEncoding(v string)`

SetEncoding sets Encoding field to given value.

### HasEncoding

`func (o *GitBlobJSON) HasEncoding() bool`

HasEncoding returns a boolean if a field has been set.

### GetPath

`func (o *GitBlobJSON) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *GitBlobJSON) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *GitBlobJSON) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *GitBlobJSON) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetSize

`func (o *GitBlobJSON) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *GitBlobJSON) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *GitBlobJSON) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *GitBlobJSON) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetTruncated

`func (o *GitBlobJSON) GetTruncated() bool`

GetTruncated returns the Truncated field if non-nil, zero value otherwise.

### GetTruncatedOk

`func (o *GitBlobJSON) GetTruncatedOk() (*bool, bool)`

GetTruncatedOk returns a tuple with the Truncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncated

`func (o *GitBlobJSON) SetTruncated(v bool)`

SetTruncated sets Truncated field to given value.

### HasTruncated

`func (o *GitBlobJSON) HasTruncated() bool`

HasTruncated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


