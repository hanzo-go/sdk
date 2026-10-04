# KvKvWrite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bucket** | Pointer to **string** | Bucket is the bucket, from the path. | [optional] 
**Key** | Pointer to **string** | Key is the key, from the path. | [optional] 
**Value** | Pointer to **string** | Value is the value, carried verbatim as UTF-8 text (typically JSON). | [optional] 

## Methods

### NewKvKvWrite

`func NewKvKvWrite() *KvKvWrite`

NewKvKvWrite instantiates a new KvKvWrite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKvKvWriteWithDefaults

`func NewKvKvWriteWithDefaults() *KvKvWrite`

NewKvKvWriteWithDefaults instantiates a new KvKvWrite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucket

`func (o *KvKvWrite) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *KvKvWrite) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *KvKvWrite) SetBucket(v string)`

SetBucket sets Bucket field to given value.

### HasBucket

`func (o *KvKvWrite) HasBucket() bool`

HasBucket returns a boolean if a field has been set.

### GetKey

`func (o *KvKvWrite) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *KvKvWrite) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *KvKvWrite) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *KvKvWrite) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetValue

`func (o *KvKvWrite) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *KvKvWrite) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *KvKvWrite) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *KvKvWrite) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


