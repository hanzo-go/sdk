# FileJSON

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to **string** | Content is the file&#39;s bytes, empty when Truncated. | [optional] 
**Encoding** | Pointer to **string** | Encoding is how Content is carried: \&quot;utf8\&quot; verbatim, or \&quot;base64\&quot;. | [optional] 
**Path** | Pointer to **string** | Path is the file&#39;s repo-relative path. | [optional] 
**Size** | Pointer to **int32** | Size is the file&#39;s byte length in the repo. | [optional] 
**Truncated** | Pointer to **bool** | Truncated marks a file past the read cap; no content is sent. A caller assembling a desired set must treat this as INCOMPLETE, never as empty. | [optional] 

## Methods

### NewFileJSON

`func NewFileJSON() *FileJSON`

NewFileJSON instantiates a new FileJSON object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileJSONWithDefaults

`func NewFileJSONWithDefaults() *FileJSON`

NewFileJSONWithDefaults instantiates a new FileJSON object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *FileJSON) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *FileJSON) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *FileJSON) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *FileJSON) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetEncoding

`func (o *FileJSON) GetEncoding() string`

GetEncoding returns the Encoding field if non-nil, zero value otherwise.

### GetEncodingOk

`func (o *FileJSON) GetEncodingOk() (*string, bool)`

GetEncodingOk returns a tuple with the Encoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncoding

`func (o *FileJSON) SetEncoding(v string)`

SetEncoding sets Encoding field to given value.

### HasEncoding

`func (o *FileJSON) HasEncoding() bool`

HasEncoding returns a boolean if a field has been set.

### GetPath

`func (o *FileJSON) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *FileJSON) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *FileJSON) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *FileJSON) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetSize

`func (o *FileJSON) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *FileJSON) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *FileJSON) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *FileJSON) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetTruncated

`func (o *FileJSON) GetTruncated() bool`

GetTruncated returns the Truncated field if non-nil, zero value otherwise.

### GetTruncatedOk

`func (o *FileJSON) GetTruncatedOk() (*bool, bool)`

GetTruncatedOk returns a tuple with the Truncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncated

`func (o *FileJSON) SetTruncated(v bool)`

SetTruncated sets Truncated field to given value.

### HasTruncated

`func (o *FileJSON) HasTruncated() bool`

HasTruncated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


