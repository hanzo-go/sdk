# GitFileJSON

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to **string** | Content is the file&#39;s bytes, empty when Truncated. | [optional] 
**Encoding** | Pointer to **string** | Encoding is how Content is carried: \&quot;utf8\&quot; verbatim, or \&quot;base64\&quot;. | [optional] 
**Path** | Pointer to **string** | Path is the file&#39;s repo-relative path. | [optional] 
**Size** | Pointer to **int64** | Size is the file&#39;s byte length in the repo. | [optional] 
**Truncated** | Pointer to **bool** | Truncated marks a file past the read cap; no content is sent. A caller assembling a desired set must treat this as INCOMPLETE, never as empty. | [optional] 

## Methods

### NewGitFileJSON

`func NewGitFileJSON() *GitFileJSON`

NewGitFileJSON instantiates a new GitFileJSON object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitFileJSONWithDefaults

`func NewGitFileJSONWithDefaults() *GitFileJSON`

NewGitFileJSONWithDefaults instantiates a new GitFileJSON object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *GitFileJSON) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *GitFileJSON) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *GitFileJSON) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *GitFileJSON) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetEncoding

`func (o *GitFileJSON) GetEncoding() string`

GetEncoding returns the Encoding field if non-nil, zero value otherwise.

### GetEncodingOk

`func (o *GitFileJSON) GetEncodingOk() (*string, bool)`

GetEncodingOk returns a tuple with the Encoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncoding

`func (o *GitFileJSON) SetEncoding(v string)`

SetEncoding sets Encoding field to given value.

### HasEncoding

`func (o *GitFileJSON) HasEncoding() bool`

HasEncoding returns a boolean if a field has been set.

### GetPath

`func (o *GitFileJSON) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *GitFileJSON) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *GitFileJSON) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *GitFileJSON) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetSize

`func (o *GitFileJSON) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *GitFileJSON) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *GitFileJSON) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *GitFileJSON) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetTruncated

`func (o *GitFileJSON) GetTruncated() bool`

GetTruncated returns the Truncated field if non-nil, zero value otherwise.

### GetTruncatedOk

`func (o *GitFileJSON) GetTruncatedOk() (*bool, bool)`

GetTruncatedOk returns a tuple with the Truncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncated

`func (o *GitFileJSON) SetTruncated(v bool)`

SetTruncated sets Truncated field to given value.

### HasTruncated

`func (o *GitFileJSON) HasTruncated() bool`

HasTruncated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


