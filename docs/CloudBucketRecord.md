# CloudBucketRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bucket** | Pointer to **string** | Bucket is the bucket&#39;s name within the org. | [optional] 
**History** | Pointer to **int32** | History is how many revisions each key keeps. | [optional] 
**Ttl** | Pointer to **int32** | TTL is the entry expiry in seconds; 0 means none. | [optional] 
**Values** | Pointer to **int32** | Values is how many values the bucket holds right now. | [optional] 

## Methods

### NewCloudBucketRecord

`func NewCloudBucketRecord() *CloudBucketRecord`

NewCloudBucketRecord instantiates a new CloudBucketRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudBucketRecordWithDefaults

`func NewCloudBucketRecordWithDefaults() *CloudBucketRecord`

NewCloudBucketRecordWithDefaults instantiates a new CloudBucketRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucket

`func (o *CloudBucketRecord) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *CloudBucketRecord) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *CloudBucketRecord) SetBucket(v string)`

SetBucket sets Bucket field to given value.

### HasBucket

`func (o *CloudBucketRecord) HasBucket() bool`

HasBucket returns a boolean if a field has been set.

### GetHistory

`func (o *CloudBucketRecord) GetHistory() int32`

GetHistory returns the History field if non-nil, zero value otherwise.

### GetHistoryOk

`func (o *CloudBucketRecord) GetHistoryOk() (*int32, bool)`

GetHistoryOk returns a tuple with the History field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistory

`func (o *CloudBucketRecord) SetHistory(v int32)`

SetHistory sets History field to given value.

### HasHistory

`func (o *CloudBucketRecord) HasHistory() bool`

HasHistory returns a boolean if a field has been set.

### GetTtl

`func (o *CloudBucketRecord) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *CloudBucketRecord) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *CloudBucketRecord) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *CloudBucketRecord) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetValues

`func (o *CloudBucketRecord) GetValues() int32`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *CloudBucketRecord) GetValuesOk() (*int32, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *CloudBucketRecord) SetValues(v int32)`

SetValues sets Values field to given value.

### HasValues

`func (o *CloudBucketRecord) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


