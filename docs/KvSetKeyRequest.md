# KvSetKeyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | **string** |  | 
**Ttl** | Pointer to **int32** | TTL in seconds (omit for no expiry) | [optional] 
**Nx** | Pointer to **bool** | Only set if key does not exist | [optional] 
**Xx** | Pointer to **bool** | Only set if key already exists | [optional] 

## Methods

### NewKvSetKeyRequest

`func NewKvSetKeyRequest(value string, ) *KvSetKeyRequest`

NewKvSetKeyRequest instantiates a new KvSetKeyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKvSetKeyRequestWithDefaults

`func NewKvSetKeyRequestWithDefaults() *KvSetKeyRequest`

NewKvSetKeyRequestWithDefaults instantiates a new KvSetKeyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *KvSetKeyRequest) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *KvSetKeyRequest) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *KvSetKeyRequest) SetValue(v string)`

SetValue sets Value field to given value.


### GetTtl

`func (o *KvSetKeyRequest) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *KvSetKeyRequest) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *KvSetKeyRequest) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *KvSetKeyRequest) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetNx

`func (o *KvSetKeyRequest) GetNx() bool`

GetNx returns the Nx field if non-nil, zero value otherwise.

### GetNxOk

`func (o *KvSetKeyRequest) GetNxOk() (*bool, bool)`

GetNxOk returns a tuple with the Nx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNx

`func (o *KvSetKeyRequest) SetNx(v bool)`

SetNx sets Nx field to given value.

### HasNx

`func (o *KvSetKeyRequest) HasNx() bool`

HasNx returns a boolean if a field has been set.

### GetXx

`func (o *KvSetKeyRequest) GetXx() bool`

GetXx returns the Xx field if non-nil, zero value otherwise.

### GetXxOk

`func (o *KvSetKeyRequest) GetXxOk() (*bool, bool)`

GetXxOk returns a tuple with the Xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXx

`func (o *KvSetKeyRequest) SetXx(v bool)`

SetXx sets Xx field to given value.

### HasXx

`func (o *KvSetKeyRequest) HasXx() bool`

HasXx returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


