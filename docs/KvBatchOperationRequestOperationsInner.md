# KvBatchOperationRequestOperationsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Op** | **string** |  | 
**Key** | **string** |  | 
**Value** | Pointer to **string** |  | [optional] 
**Ttl** | Pointer to **int32** |  | [optional] 

## Methods

### NewKvBatchOperationRequestOperationsInner

`func NewKvBatchOperationRequestOperationsInner(op string, key string, ) *KvBatchOperationRequestOperationsInner`

NewKvBatchOperationRequestOperationsInner instantiates a new KvBatchOperationRequestOperationsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKvBatchOperationRequestOperationsInnerWithDefaults

`func NewKvBatchOperationRequestOperationsInnerWithDefaults() *KvBatchOperationRequestOperationsInner`

NewKvBatchOperationRequestOperationsInnerWithDefaults instantiates a new KvBatchOperationRequestOperationsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOp

`func (o *KvBatchOperationRequestOperationsInner) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *KvBatchOperationRequestOperationsInner) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *KvBatchOperationRequestOperationsInner) SetOp(v string)`

SetOp sets Op field to given value.


### GetKey

`func (o *KvBatchOperationRequestOperationsInner) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *KvBatchOperationRequestOperationsInner) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *KvBatchOperationRequestOperationsInner) SetKey(v string)`

SetKey sets Key field to given value.


### GetValue

`func (o *KvBatchOperationRequestOperationsInner) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *KvBatchOperationRequestOperationsInner) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *KvBatchOperationRequestOperationsInner) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *KvBatchOperationRequestOperationsInner) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetTtl

`func (o *KvBatchOperationRequestOperationsInner) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *KvBatchOperationRequestOperationsInner) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *KvBatchOperationRequestOperationsInner) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *KvBatchOperationRequestOperationsInner) HasTtl() bool`

HasTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


