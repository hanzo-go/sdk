# S3BucketItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **int64** | unix seconds | [optional] 
**Name** | Pointer to **string** | friendly name | [optional] 

## Methods

### NewS3BucketItem

`func NewS3BucketItem() *S3BucketItem`

NewS3BucketItem instantiates a new S3BucketItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3BucketItemWithDefaults

`func NewS3BucketItemWithDefaults() *S3BucketItem`

NewS3BucketItemWithDefaults instantiates a new S3BucketItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *S3BucketItem) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *S3BucketItem) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *S3BucketItem) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *S3BucketItem) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetName

`func (o *S3BucketItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *S3BucketItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *S3BucketItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *S3BucketItem) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


