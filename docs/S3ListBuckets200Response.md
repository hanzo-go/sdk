# S3ListBuckets200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Buckets** | Pointer to [**[]S3Bucket**](S3Bucket.md) |  | [optional] 

## Methods

### NewS3ListBuckets200Response

`func NewS3ListBuckets200Response() *S3ListBuckets200Response`

NewS3ListBuckets200Response instantiates a new S3ListBuckets200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3ListBuckets200ResponseWithDefaults

`func NewS3ListBuckets200ResponseWithDefaults() *S3ListBuckets200Response`

NewS3ListBuckets200ResponseWithDefaults instantiates a new S3ListBuckets200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuckets

`func (o *S3ListBuckets200Response) GetBuckets() []S3Bucket`

GetBuckets returns the Buckets field if non-nil, zero value otherwise.

### GetBucketsOk

`func (o *S3ListBuckets200Response) GetBucketsOk() (*[]S3Bucket, bool)`

GetBucketsOk returns a tuple with the Buckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuckets

`func (o *S3ListBuckets200Response) SetBuckets(v []S3Bucket)`

SetBuckets sets Buckets field to given value.

### HasBuckets

`func (o *S3ListBuckets200Response) HasBuckets() bool`

HasBuckets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


