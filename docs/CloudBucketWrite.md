# CloudBucketWrite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bucket** | Pointer to **string** | Bucket is the bucket&#39;s name within the org, from the path: 1–64 of [A-Za-z0-9_], no dash. | [optional] 
**History** | Pointer to **int32** | History is how many revisions each key keeps, 1–64. 0 means 1. | [optional] 
**MaxValue** | Pointer to **int32** | MaxValue caps one value&#39;s size in bytes. 0 or less means the server&#39;s ceiling. | [optional] 
**Ttl** | Pointer to **int32** | TTL expires entries after this many SECONDS. 0 means no expiry. | [optional] 

## Methods

### NewCloudBucketWrite

`func NewCloudBucketWrite() *CloudBucketWrite`

NewCloudBucketWrite instantiates a new CloudBucketWrite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudBucketWriteWithDefaults

`func NewCloudBucketWriteWithDefaults() *CloudBucketWrite`

NewCloudBucketWriteWithDefaults instantiates a new CloudBucketWrite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucket

`func (o *CloudBucketWrite) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *CloudBucketWrite) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *CloudBucketWrite) SetBucket(v string)`

SetBucket sets Bucket field to given value.

### HasBucket

`func (o *CloudBucketWrite) HasBucket() bool`

HasBucket returns a boolean if a field has been set.

### GetHistory

`func (o *CloudBucketWrite) GetHistory() int32`

GetHistory returns the History field if non-nil, zero value otherwise.

### GetHistoryOk

`func (o *CloudBucketWrite) GetHistoryOk() (*int32, bool)`

GetHistoryOk returns a tuple with the History field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistory

`func (o *CloudBucketWrite) SetHistory(v int32)`

SetHistory sets History field to given value.

### HasHistory

`func (o *CloudBucketWrite) HasHistory() bool`

HasHistory returns a boolean if a field has been set.

### GetMaxValue

`func (o *CloudBucketWrite) GetMaxValue() int32`

GetMaxValue returns the MaxValue field if non-nil, zero value otherwise.

### GetMaxValueOk

`func (o *CloudBucketWrite) GetMaxValueOk() (*int32, bool)`

GetMaxValueOk returns a tuple with the MaxValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxValue

`func (o *CloudBucketWrite) SetMaxValue(v int32)`

SetMaxValue sets MaxValue field to given value.

### HasMaxValue

`func (o *CloudBucketWrite) HasMaxValue() bool`

HasMaxValue returns a boolean if a field has been set.

### GetTtl

`func (o *CloudBucketWrite) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *CloudBucketWrite) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *CloudBucketWrite) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *CloudBucketWrite) HasTtl() bool`

HasTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


