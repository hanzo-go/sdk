# S3StoredPart

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Etag** | Pointer to **string** | ETag is the store&#39;s tag for its bytes, quotes stripped. | [optional] 
**Part** | Pointer to **int64** | Part is the part number. | [optional] 
**Size** | Pointer to **int64** | Size is its length in bytes. | [optional] 

## Methods

### NewS3StoredPart

`func NewS3StoredPart() *S3StoredPart`

NewS3StoredPart instantiates a new S3StoredPart object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3StoredPartWithDefaults

`func NewS3StoredPartWithDefaults() *S3StoredPart`

NewS3StoredPartWithDefaults instantiates a new S3StoredPart object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEtag

`func (o *S3StoredPart) GetEtag() string`

GetEtag returns the Etag field if non-nil, zero value otherwise.

### GetEtagOk

`func (o *S3StoredPart) GetEtagOk() (*string, bool)`

GetEtagOk returns a tuple with the Etag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtag

`func (o *S3StoredPart) SetEtag(v string)`

SetEtag sets Etag field to given value.

### HasEtag

`func (o *S3StoredPart) HasEtag() bool`

HasEtag returns a boolean if a field has been set.

### GetPart

`func (o *S3StoredPart) GetPart() int64`

GetPart returns the Part field if non-nil, zero value otherwise.

### GetPartOk

`func (o *S3StoredPart) GetPartOk() (*int64, bool)`

GetPartOk returns a tuple with the Part field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPart

`func (o *S3StoredPart) SetPart(v int64)`

SetPart sets Part field to given value.

### HasPart

`func (o *S3StoredPart) HasPart() bool`

HasPart returns a boolean if a field has been set.

### GetSize

`func (o *S3StoredPart) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *S3StoredPart) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *S3StoredPart) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *S3StoredPart) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


