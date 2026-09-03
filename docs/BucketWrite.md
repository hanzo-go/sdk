# BucketWrite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bucket** | Pointer to **string** | Bucket is the bucket&#39;s name within the org, from the path: 1–64 of [A-Za-z0-9_], no dash. | [optional] 
**History** | Pointer to **int64** | History is how many revisions each key keeps, 1–64. 0 means 1. | [optional] 
**MaxValue** | Pointer to **int64** | MaxValue caps one value&#39;s size in bytes. 0 or less means the server&#39;s ceiling. | [optional] 
**Ttl** | Pointer to **int64** | TTL expires entries after this many SECONDS. 0 means no expiry. | [optional] 

## Methods

### NewBucketWrite

`func NewBucketWrite() *BucketWrite`

NewBucketWrite instantiates a new BucketWrite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBucketWriteWithDefaults

`func NewBucketWriteWithDefaults() *BucketWrite`

NewBucketWriteWithDefaults instantiates a new BucketWrite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucket

`func (o *BucketWrite) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *BucketWrite) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *BucketWrite) SetBucket(v string)`

SetBucket sets Bucket field to given value.

### HasBucket

`func (o *BucketWrite) HasBucket() bool`

HasBucket returns a boolean if a field has been set.

### GetHistory

`func (o *BucketWrite) GetHistory() int64`

GetHistory returns the History field if non-nil, zero value otherwise.

### GetHistoryOk

`func (o *BucketWrite) GetHistoryOk() (*int64, bool)`

GetHistoryOk returns a tuple with the History field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistory

`func (o *BucketWrite) SetHistory(v int64)`

SetHistory sets History field to given value.

### HasHistory

`func (o *BucketWrite) HasHistory() bool`

HasHistory returns a boolean if a field has been set.

### GetMaxValue

`func (o *BucketWrite) GetMaxValue() int64`

GetMaxValue returns the MaxValue field if non-nil, zero value otherwise.

### GetMaxValueOk

`func (o *BucketWrite) GetMaxValueOk() (*int64, bool)`

GetMaxValueOk returns a tuple with the MaxValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxValue

`func (o *BucketWrite) SetMaxValue(v int64)`

SetMaxValue sets MaxValue field to given value.

### HasMaxValue

`func (o *BucketWrite) HasMaxValue() bool`

HasMaxValue returns a boolean if a field has been set.

### GetTtl

`func (o *BucketWrite) GetTtl() int64`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *BucketWrite) GetTtlOk() (*int64, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *BucketWrite) SetTtl(v int64)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *BucketWrite) HasTtl() bool`

HasTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


