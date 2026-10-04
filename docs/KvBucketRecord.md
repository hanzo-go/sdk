# KvBucketRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bucket** | Pointer to **string** | Bucket is the bucket&#39;s name within the org. | [optional] 
**History** | Pointer to **int64** | History is how many revisions each key keeps. | [optional] 
**Ttl** | Pointer to **int64** | TTL is the entry expiry in seconds; 0 means none. | [optional] 
**Values** | Pointer to **int32** | Values is how many values the bucket holds right now. | [optional] 

## Methods

### NewKvBucketRecord

`func NewKvBucketRecord() *KvBucketRecord`

NewKvBucketRecord instantiates a new KvBucketRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKvBucketRecordWithDefaults

`func NewKvBucketRecordWithDefaults() *KvBucketRecord`

NewKvBucketRecordWithDefaults instantiates a new KvBucketRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucket

`func (o *KvBucketRecord) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *KvBucketRecord) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *KvBucketRecord) SetBucket(v string)`

SetBucket sets Bucket field to given value.

### HasBucket

`func (o *KvBucketRecord) HasBucket() bool`

HasBucket returns a boolean if a field has been set.

### GetHistory

`func (o *KvBucketRecord) GetHistory() int64`

GetHistory returns the History field if non-nil, zero value otherwise.

### GetHistoryOk

`func (o *KvBucketRecord) GetHistoryOk() (*int64, bool)`

GetHistoryOk returns a tuple with the History field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistory

`func (o *KvBucketRecord) SetHistory(v int64)`

SetHistory sets History field to given value.

### HasHistory

`func (o *KvBucketRecord) HasHistory() bool`

HasHistory returns a boolean if a field has been set.

### GetTtl

`func (o *KvBucketRecord) GetTtl() int64`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *KvBucketRecord) GetTtlOk() (*int64, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *KvBucketRecord) SetTtl(v int64)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *KvBucketRecord) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetValues

`func (o *KvBucketRecord) GetValues() int32`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *KvBucketRecord) GetValuesOk() (*int32, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *KvBucketRecord) SetValues(v int32)`

SetValues sets Values field to given value.

### HasValues

`func (o *KvBucketRecord) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


