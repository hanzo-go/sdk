# S3UploadRef

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bucket** | Pointer to **string** | Bucket is the bucket, from the path. | [optional] 
**Key** | Pointer to **string** | Key is the object key the upload was started on. In the query for a read or an abandon, in the body to complete. | [optional] 
**Upload** | Pointer to **string** | Upload is the upload&#39;s id, from the path. | [optional] 

## Methods

### NewS3UploadRef

`func NewS3UploadRef() *S3UploadRef`

NewS3UploadRef instantiates a new S3UploadRef object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3UploadRefWithDefaults

`func NewS3UploadRefWithDefaults() *S3UploadRef`

NewS3UploadRefWithDefaults instantiates a new S3UploadRef object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucket

`func (o *S3UploadRef) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *S3UploadRef) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *S3UploadRef) SetBucket(v string)`

SetBucket sets Bucket field to given value.

### HasBucket

`func (o *S3UploadRef) HasBucket() bool`

HasBucket returns a boolean if a field has been set.

### GetKey

`func (o *S3UploadRef) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *S3UploadRef) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *S3UploadRef) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *S3UploadRef) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetUpload

`func (o *S3UploadRef) GetUpload() string`

GetUpload returns the Upload field if non-nil, zero value otherwise.

### GetUploadOk

`func (o *S3UploadRef) GetUploadOk() (*string, bool)`

GetUploadOk returns a tuple with the Upload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpload

`func (o *S3UploadRef) SetUpload(v string)`

SetUpload sets Upload field to given value.

### HasUpload

`func (o *S3UploadRef) HasUpload() bool`

HasUpload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


