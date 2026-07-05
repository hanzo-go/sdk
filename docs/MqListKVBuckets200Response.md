# MqListKVBuckets200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Buckets** | Pointer to [**[]MqKVBucket**](MqKVBucket.md) |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 

## Methods

### NewMqListKVBuckets200Response

`func NewMqListKVBuckets200Response() *MqListKVBuckets200Response`

NewMqListKVBuckets200Response instantiates a new MqListKVBuckets200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMqListKVBuckets200ResponseWithDefaults

`func NewMqListKVBuckets200ResponseWithDefaults() *MqListKVBuckets200Response`

NewMqListKVBuckets200ResponseWithDefaults instantiates a new MqListKVBuckets200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuckets

`func (o *MqListKVBuckets200Response) GetBuckets() []MqKVBucket`

GetBuckets returns the Buckets field if non-nil, zero value otherwise.

### GetBucketsOk

`func (o *MqListKVBuckets200Response) GetBucketsOk() (*[]MqKVBucket, bool)`

GetBucketsOk returns a tuple with the Buckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuckets

`func (o *MqListKVBuckets200Response) SetBuckets(v []MqKVBucket)`

SetBuckets sets Buckets field to given value.

### HasBuckets

`func (o *MqListKVBuckets200Response) HasBuckets() bool`

HasBuckets returns a boolean if a field has been set.

### GetTotal

`func (o *MqListKVBuckets200Response) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *MqListKVBuckets200Response) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *MqListKVBuckets200Response) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *MqListKVBuckets200Response) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


