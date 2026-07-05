# S3ObjectInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **int32** |  | [optional] 
**Etag** | Pointer to **string** |  | [optional] 
**ContentType** | Pointer to **string** |  | [optional] 
**LastModified** | Pointer to **time.Time** |  | [optional] 
**VersionId** | Pointer to **string** |  | [optional] 
**StorageClass** | Pointer to **string** |  | [optional] 

## Methods

### NewS3ObjectInfo

`func NewS3ObjectInfo() *S3ObjectInfo`

NewS3ObjectInfo instantiates a new S3ObjectInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3ObjectInfoWithDefaults

`func NewS3ObjectInfoWithDefaults() *S3ObjectInfo`

NewS3ObjectInfoWithDefaults instantiates a new S3ObjectInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *S3ObjectInfo) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *S3ObjectInfo) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *S3ObjectInfo) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *S3ObjectInfo) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetSize

`func (o *S3ObjectInfo) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *S3ObjectInfo) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *S3ObjectInfo) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *S3ObjectInfo) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetEtag

`func (o *S3ObjectInfo) GetEtag() string`

GetEtag returns the Etag field if non-nil, zero value otherwise.

### GetEtagOk

`func (o *S3ObjectInfo) GetEtagOk() (*string, bool)`

GetEtagOk returns a tuple with the Etag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtag

`func (o *S3ObjectInfo) SetEtag(v string)`

SetEtag sets Etag field to given value.

### HasEtag

`func (o *S3ObjectInfo) HasEtag() bool`

HasEtag returns a boolean if a field has been set.

### GetContentType

`func (o *S3ObjectInfo) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *S3ObjectInfo) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *S3ObjectInfo) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *S3ObjectInfo) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetLastModified

`func (o *S3ObjectInfo) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *S3ObjectInfo) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *S3ObjectInfo) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *S3ObjectInfo) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.

### GetVersionId

`func (o *S3ObjectInfo) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *S3ObjectInfo) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *S3ObjectInfo) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.

### HasVersionId

`func (o *S3ObjectInfo) HasVersionId() bool`

HasVersionId returns a boolean if a field has been set.

### GetStorageClass

`func (o *S3ObjectInfo) GetStorageClass() string`

GetStorageClass returns the StorageClass field if non-nil, zero value otherwise.

### GetStorageClassOk

`func (o *S3ObjectInfo) GetStorageClassOk() (*string, bool)`

GetStorageClassOk returns a tuple with the StorageClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageClass

`func (o *S3ObjectInfo) SetStorageClass(v string)`

SetStorageClass sets StorageClass field to given value.

### HasStorageClass

`func (o *S3ObjectInfo) HasStorageClass() bool`

HasStorageClass returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


