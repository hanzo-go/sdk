# S3UploadParts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bucket** | Pointer to **string** | Bucket is the bucket, from the path. | [optional] 
**Key** | Pointer to **string** | Key is the object key the upload was started on. | [optional] 
**Parts** | Pointer to **[]int64** | Parts are the part numbers to mint URLs for, from 1. At most 64 a call. | [optional] 
**Upload** | Pointer to **string** | Upload is the upload&#39;s id, from the path. | [optional] 

## Methods

### NewS3UploadParts

`func NewS3UploadParts() *S3UploadParts`

NewS3UploadParts instantiates a new S3UploadParts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3UploadPartsWithDefaults

`func NewS3UploadPartsWithDefaults() *S3UploadParts`

NewS3UploadPartsWithDefaults instantiates a new S3UploadParts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucket

`func (o *S3UploadParts) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *S3UploadParts) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *S3UploadParts) SetBucket(v string)`

SetBucket sets Bucket field to given value.

### HasBucket

`func (o *S3UploadParts) HasBucket() bool`

HasBucket returns a boolean if a field has been set.

### GetKey

`func (o *S3UploadParts) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *S3UploadParts) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *S3UploadParts) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *S3UploadParts) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetParts

`func (o *S3UploadParts) GetParts() []int64`

GetParts returns the Parts field if non-nil, zero value otherwise.

### GetPartsOk

`func (o *S3UploadParts) GetPartsOk() (*[]int64, bool)`

GetPartsOk returns a tuple with the Parts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParts

`func (o *S3UploadParts) SetParts(v []int64)`

SetParts sets Parts field to given value.

### HasParts

`func (o *S3UploadParts) HasParts() bool`

HasParts returns a boolean if a field has been set.

### GetUpload

`func (o *S3UploadParts) GetUpload() string`

GetUpload returns the Upload field if non-nil, zero value otherwise.

### GetUploadOk

`func (o *S3UploadParts) GetUploadOk() (*string, bool)`

GetUploadOk returns a tuple with the Upload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpload

`func (o *S3UploadParts) SetUpload(v string)`

SetUpload sets Upload field to given value.

### HasUpload

`func (o *S3UploadParts) HasUpload() bool`

HasUpload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


