# S3UploadIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bucket** | Pointer to **string** | Bucket is the bucket to upload into, from the path. | [optional] 
**Key** | Pointer to **string** | Key is the object key relative to the bucket root. It is path-cleaned, so a \&quot;../\&quot; cannot escape the bucket, and an empty or unclean key is 400. | [optional] 

## Methods

### NewS3UploadIn

`func NewS3UploadIn() *S3UploadIn`

NewS3UploadIn instantiates a new S3UploadIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3UploadInWithDefaults

`func NewS3UploadInWithDefaults() *S3UploadIn`

NewS3UploadInWithDefaults instantiates a new S3UploadIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucket

`func (o *S3UploadIn) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *S3UploadIn) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *S3UploadIn) SetBucket(v string)`

SetBucket sets Bucket field to given value.

### HasBucket

`func (o *S3UploadIn) HasBucket() bool`

HasBucket returns a boolean if a field has been set.

### GetKey

`func (o *S3UploadIn) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *S3UploadIn) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *S3UploadIn) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *S3UploadIn) HasKey() bool`

HasKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


