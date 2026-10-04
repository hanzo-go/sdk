# S3UploadDone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Etag** | Pointer to **string** | ETag is the assembled object&#39;s tag, quotes stripped. | [optional] 
**Key** | Pointer to **string** | Key is the object&#39;s key. | [optional] 
**Parts** | Pointer to **int64** | Parts is how many parts it was assembled from. | [optional] 
**Size** | Pointer to **int64** | Size is the object&#39;s length in bytes: the sum of its parts. | [optional] 

## Methods

### NewS3UploadDone

`func NewS3UploadDone() *S3UploadDone`

NewS3UploadDone instantiates a new S3UploadDone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3UploadDoneWithDefaults

`func NewS3UploadDoneWithDefaults() *S3UploadDone`

NewS3UploadDoneWithDefaults instantiates a new S3UploadDone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEtag

`func (o *S3UploadDone) GetEtag() string`

GetEtag returns the Etag field if non-nil, zero value otherwise.

### GetEtagOk

`func (o *S3UploadDone) GetEtagOk() (*string, bool)`

GetEtagOk returns a tuple with the Etag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtag

`func (o *S3UploadDone) SetEtag(v string)`

SetEtag sets Etag field to given value.

### HasEtag

`func (o *S3UploadDone) HasEtag() bool`

HasEtag returns a boolean if a field has been set.

### GetKey

`func (o *S3UploadDone) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *S3UploadDone) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *S3UploadDone) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *S3UploadDone) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetParts

`func (o *S3UploadDone) GetParts() int64`

GetParts returns the Parts field if non-nil, zero value otherwise.

### GetPartsOk

`func (o *S3UploadDone) GetPartsOk() (*int64, bool)`

GetPartsOk returns a tuple with the Parts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParts

`func (o *S3UploadDone) SetParts(v int64)`

SetParts sets Parts field to given value.

### HasParts

`func (o *S3UploadDone) HasParts() bool`

HasParts returns a boolean if a field has been set.

### GetSize

`func (o *S3UploadDone) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *S3UploadDone) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *S3UploadDone) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *S3UploadDone) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


