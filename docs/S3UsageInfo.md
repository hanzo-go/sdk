# S3UsageInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalSize** | Pointer to **int32** | Total storage in bytes | [optional] 
**TotalObjects** | Pointer to **int32** |  | [optional] 
**Buckets** | Pointer to **int32** |  | [optional] 

## Methods

### NewS3UsageInfo

`func NewS3UsageInfo() *S3UsageInfo`

NewS3UsageInfo instantiates a new S3UsageInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3UsageInfoWithDefaults

`func NewS3UsageInfoWithDefaults() *S3UsageInfo`

NewS3UsageInfoWithDefaults instantiates a new S3UsageInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalSize

`func (o *S3UsageInfo) GetTotalSize() int32`

GetTotalSize returns the TotalSize field if non-nil, zero value otherwise.

### GetTotalSizeOk

`func (o *S3UsageInfo) GetTotalSizeOk() (*int32, bool)`

GetTotalSizeOk returns a tuple with the TotalSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSize

`func (o *S3UsageInfo) SetTotalSize(v int32)`

SetTotalSize sets TotalSize field to given value.

### HasTotalSize

`func (o *S3UsageInfo) HasTotalSize() bool`

HasTotalSize returns a boolean if a field has been set.

### GetTotalObjects

`func (o *S3UsageInfo) GetTotalObjects() int32`

GetTotalObjects returns the TotalObjects field if non-nil, zero value otherwise.

### GetTotalObjectsOk

`func (o *S3UsageInfo) GetTotalObjectsOk() (*int32, bool)`

GetTotalObjectsOk returns a tuple with the TotalObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalObjects

`func (o *S3UsageInfo) SetTotalObjects(v int32)`

SetTotalObjects sets TotalObjects field to given value.

### HasTotalObjects

`func (o *S3UsageInfo) HasTotalObjects() bool`

HasTotalObjects returns a boolean if a field has been set.

### GetBuckets

`func (o *S3UsageInfo) GetBuckets() int32`

GetBuckets returns the Buckets field if non-nil, zero value otherwise.

### GetBucketsOk

`func (o *S3UsageInfo) GetBucketsOk() (*int32, bool)`

GetBucketsOk returns a tuple with the Buckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuckets

`func (o *S3UsageInfo) SetBuckets(v int32)`

SetBuckets sets Buckets field to given value.

### HasBuckets

`func (o *S3UsageInfo) HasBuckets() bool`

HasBuckets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


