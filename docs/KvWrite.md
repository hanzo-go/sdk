# KvWrite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bucket** | Pointer to **string** | Bucket is the bucket, from the path. | [optional] 
**Key** | Pointer to **string** | Key is the key, from the path. | [optional] 
**Value** | Pointer to **string** | Value is the value, carried verbatim as UTF-8 text (typically JSON). | [optional] 

## Methods

### NewKvWrite

`func NewKvWrite() *KvWrite`

NewKvWrite instantiates a new KvWrite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKvWriteWithDefaults

`func NewKvWriteWithDefaults() *KvWrite`

NewKvWriteWithDefaults instantiates a new KvWrite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucket

`func (o *KvWrite) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *KvWrite) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *KvWrite) SetBucket(v string)`

SetBucket sets Bucket field to given value.

### HasBucket

`func (o *KvWrite) HasBucket() bool`

HasBucket returns a boolean if a field has been set.

### GetKey

`func (o *KvWrite) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *KvWrite) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *KvWrite) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *KvWrite) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetValue

`func (o *KvWrite) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *KvWrite) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *KvWrite) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *KvWrite) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


