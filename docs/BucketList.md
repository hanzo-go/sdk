# BucketList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Buckets** | Pointer to [**[]BucketItem**](BucketItem.md) | Buckets are the caller org&#39;s buckets, friendly names, oldest first as the store returns them. | [optional] 
**Total** | Pointer to **int32** | Total is how many buckets this org has. It equals len(buckets): the listing is not paged, because an org&#39;s bucket count is small by construction. | [optional] 

## Methods

### NewBucketList

`func NewBucketList() *BucketList`

NewBucketList instantiates a new BucketList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBucketListWithDefaults

`func NewBucketListWithDefaults() *BucketList`

NewBucketListWithDefaults instantiates a new BucketList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuckets

`func (o *BucketList) GetBuckets() []BucketItem`

GetBuckets returns the Buckets field if non-nil, zero value otherwise.

### GetBucketsOk

`func (o *BucketList) GetBucketsOk() (*[]BucketItem, bool)`

GetBucketsOk returns a tuple with the Buckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuckets

`func (o *BucketList) SetBuckets(v []BucketItem)`

SetBuckets sets Buckets field to given value.

### HasBuckets

`func (o *BucketList) HasBuckets() bool`

HasBuckets returns a boolean if a field has been set.

### GetTotal

`func (o *BucketList) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *BucketList) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *BucketList) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *BucketList) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


