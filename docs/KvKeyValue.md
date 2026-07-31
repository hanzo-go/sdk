# KvKeyValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Ttl** | Pointer to **int32** | Time-to-live in seconds (-1 for no expiry) | [optional] 
**Size** | Pointer to **int32** | Size in bytes | [optional] 

## Methods

### NewKvKeyValue

`func NewKvKeyValue() *KvKeyValue`

NewKvKeyValue instantiates a new KvKeyValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKvKeyValueWithDefaults

`func NewKvKeyValueWithDefaults() *KvKeyValue`

NewKvKeyValueWithDefaults instantiates a new KvKeyValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *KvKeyValue) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *KvKeyValue) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *KvKeyValue) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *KvKeyValue) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetValue

`func (o *KvKeyValue) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *KvKeyValue) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *KvKeyValue) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *KvKeyValue) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetType

`func (o *KvKeyValue) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *KvKeyValue) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *KvKeyValue) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *KvKeyValue) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTtl

`func (o *KvKeyValue) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *KvKeyValue) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *KvKeyValue) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *KvKeyValue) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetSize

`func (o *KvKeyValue) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *KvKeyValue) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *KvKeyValue) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *KvKeyValue) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


