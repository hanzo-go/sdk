# KvBatchOperationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Namespace** | Pointer to **string** |  | [optional] 
**Operations** | Pointer to [**[]KvBatchOperationRequestOperationsInner**](KvBatchOperationRequestOperationsInner.md) |  | [optional] 

## Methods

### NewKvBatchOperationRequest

`func NewKvBatchOperationRequest() *KvBatchOperationRequest`

NewKvBatchOperationRequest instantiates a new KvBatchOperationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKvBatchOperationRequestWithDefaults

`func NewKvBatchOperationRequestWithDefaults() *KvBatchOperationRequest`

NewKvBatchOperationRequestWithDefaults instantiates a new KvBatchOperationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNamespace

`func (o *KvBatchOperationRequest) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *KvBatchOperationRequest) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *KvBatchOperationRequest) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *KvBatchOperationRequest) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetOperations

`func (o *KvBatchOperationRequest) GetOperations() []KvBatchOperationRequestOperationsInner`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *KvBatchOperationRequest) GetOperationsOk() (*[]KvBatchOperationRequestOperationsInner, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *KvBatchOperationRequest) SetOperations(v []KvBatchOperationRequestOperationsInner)`

SetOperations sets Operations field to given value.

### HasOperations

`func (o *KvBatchOperationRequest) HasOperations() bool`

HasOperations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


