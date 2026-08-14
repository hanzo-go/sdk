# BlobJSON

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Binary** | Pointer to **bool** | Binary marks content git could not treat as text; it comes back base64. | [optional] 
**Content** | Pointer to **string** | Content is the file&#39;s bytes, empty when Truncated. | [optional] 
**Encoding** | Pointer to **string** | Encoding is how Content is carried: \&quot;utf8\&quot; verbatim, or \&quot;base64\&quot;. | [optional] 
**Path** | Pointer to **string** | Path is the file&#39;s repo-relative path. | [optional] 
**Size** | Pointer to **int32** | Size is the file&#39;s byte length in the repo, whatever was returned below. | [optional] 
**Truncated** | Pointer to **bool** | Truncated marks a file past the 1 MiB view cap. No content is sent — clone the repo for it. | [optional] 

## Methods

### NewBlobJSON

`func NewBlobJSON() *BlobJSON`

NewBlobJSON instantiates a new BlobJSON object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlobJSONWithDefaults

`func NewBlobJSONWithDefaults() *BlobJSON`

NewBlobJSONWithDefaults instantiates a new BlobJSON object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBinary

`func (o *BlobJSON) GetBinary() bool`

GetBinary returns the Binary field if non-nil, zero value otherwise.

### GetBinaryOk

`func (o *BlobJSON) GetBinaryOk() (*bool, bool)`

GetBinaryOk returns a tuple with the Binary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinary

`func (o *BlobJSON) SetBinary(v bool)`

SetBinary sets Binary field to given value.

### HasBinary

`func (o *BlobJSON) HasBinary() bool`

HasBinary returns a boolean if a field has been set.

### GetContent

`func (o *BlobJSON) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *BlobJSON) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *BlobJSON) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *BlobJSON) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetEncoding

`func (o *BlobJSON) GetEncoding() string`

GetEncoding returns the Encoding field if non-nil, zero value otherwise.

### GetEncodingOk

`func (o *BlobJSON) GetEncodingOk() (*string, bool)`

GetEncodingOk returns a tuple with the Encoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncoding

`func (o *BlobJSON) SetEncoding(v string)`

SetEncoding sets Encoding field to given value.

### HasEncoding

`func (o *BlobJSON) HasEncoding() bool`

HasEncoding returns a boolean if a field has been set.

### GetPath

`func (o *BlobJSON) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *BlobJSON) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *BlobJSON) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *BlobJSON) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetSize

`func (o *BlobJSON) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *BlobJSON) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *BlobJSON) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *BlobJSON) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetTruncated

`func (o *BlobJSON) GetTruncated() bool`

GetTruncated returns the Truncated field if non-nil, zero value otherwise.

### GetTruncatedOk

`func (o *BlobJSON) GetTruncatedOk() (*bool, bool)`

GetTruncatedOk returns a tuple with the Truncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncated

`func (o *BlobJSON) SetTruncated(v bool)`

SetTruncated sets Truncated field to given value.

### HasTruncated

`func (o *BlobJSON) HasTruncated() bool`

HasTruncated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


