# KvSetKeyTTLRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ttl** | **int32** | TTL in seconds (-1 to remove expiry) | 

## Methods

### NewKvSetKeyTTLRequest

`func NewKvSetKeyTTLRequest(ttl int32, ) *KvSetKeyTTLRequest`

NewKvSetKeyTTLRequest instantiates a new KvSetKeyTTLRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKvSetKeyTTLRequestWithDefaults

`func NewKvSetKeyTTLRequestWithDefaults() *KvSetKeyTTLRequest`

NewKvSetKeyTTLRequestWithDefaults instantiates a new KvSetKeyTTLRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTtl

`func (o *KvSetKeyTTLRequest) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *KvSetKeyTTLRequest) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *KvSetKeyTTLRequest) SetTtl(v int32)`

SetTtl sets Ttl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


