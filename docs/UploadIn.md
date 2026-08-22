# UploadIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bucket** | Pointer to **string** | Bucket is the bucket to upload into, from the path. | [optional] 
**Key** | Pointer to **string** | Key is the object key relative to the bucket root. It is path-cleaned, so a \&quot;../\&quot; cannot escape the bucket, and an empty or unclean key is 400. | [optional] 

## Methods

### NewUploadIn

`func NewUploadIn() *UploadIn`

NewUploadIn instantiates a new UploadIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUploadInWithDefaults

`func NewUploadInWithDefaults() *UploadIn`

NewUploadInWithDefaults instantiates a new UploadIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucket

`func (o *UploadIn) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *UploadIn) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *UploadIn) SetBucket(v string)`

SetBucket sets Bucket field to given value.

### HasBucket

`func (o *UploadIn) HasBucket() bool`

HasBucket returns a boolean if a field has been set.

### GetKey

`func (o *UploadIn) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *UploadIn) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *UploadIn) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *UploadIn) HasKey() bool`

HasKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


