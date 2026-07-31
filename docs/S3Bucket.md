# S3Bucket

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**CreationDate** | Pointer to **time.Time** |  | [optional] 
**Size** | Pointer to **int32** | Total size in bytes | [optional] 
**Objects** | Pointer to **int32** | Number of objects | [optional] 

## Methods

### NewS3Bucket

`func NewS3Bucket() *S3Bucket`

NewS3Bucket instantiates a new S3Bucket object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3BucketWithDefaults

`func NewS3BucketWithDefaults() *S3Bucket`

NewS3BucketWithDefaults instantiates a new S3Bucket object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *S3Bucket) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *S3Bucket) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *S3Bucket) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *S3Bucket) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCreationDate

`func (o *S3Bucket) GetCreationDate() time.Time`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *S3Bucket) GetCreationDateOk() (*time.Time, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *S3Bucket) SetCreationDate(v time.Time)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *S3Bucket) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetSize

`func (o *S3Bucket) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *S3Bucket) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *S3Bucket) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *S3Bucket) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetObjects

`func (o *S3Bucket) GetObjects() int32`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *S3Bucket) GetObjectsOk() (*int32, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *S3Bucket) SetObjects(v int32)`

SetObjects sets Objects field to given value.

### HasObjects

`func (o *S3Bucket) HasObjects() bool`

HasObjects returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


