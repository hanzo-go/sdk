# KmsBatchSecretRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Create** | Pointer to [**[]KmsCreateSecretRequest**](KmsCreateSecretRequest.md) |  | [optional] 
**Update** | Pointer to [**[]KmsUpdateSecretRequest**](KmsUpdateSecretRequest.md) |  | [optional] 
**Delete** | Pointer to [**[]KmsBatchSecretRequestDeleteInner**](KmsBatchSecretRequestDeleteInner.md) |  | [optional] 

## Methods

### NewKmsBatchSecretRequest

`func NewKmsBatchSecretRequest() *KmsBatchSecretRequest`

NewKmsBatchSecretRequest instantiates a new KmsBatchSecretRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKmsBatchSecretRequestWithDefaults

`func NewKmsBatchSecretRequestWithDefaults() *KmsBatchSecretRequest`

NewKmsBatchSecretRequestWithDefaults instantiates a new KmsBatchSecretRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreate

`func (o *KmsBatchSecretRequest) GetCreate() []KmsCreateSecretRequest`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *KmsBatchSecretRequest) GetCreateOk() (*[]KmsCreateSecretRequest, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *KmsBatchSecretRequest) SetCreate(v []KmsCreateSecretRequest)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *KmsBatchSecretRequest) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetUpdate

`func (o *KmsBatchSecretRequest) GetUpdate() []KmsUpdateSecretRequest`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *KmsBatchSecretRequest) GetUpdateOk() (*[]KmsUpdateSecretRequest, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *KmsBatchSecretRequest) SetUpdate(v []KmsUpdateSecretRequest)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *KmsBatchSecretRequest) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetDelete

`func (o *KmsBatchSecretRequest) GetDelete() []KmsBatchSecretRequestDeleteInner`

GetDelete returns the Delete field if non-nil, zero value otherwise.

### GetDeleteOk

`func (o *KmsBatchSecretRequest) GetDeleteOk() (*[]KmsBatchSecretRequestDeleteInner, bool)`

GetDeleteOk returns a tuple with the Delete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelete

`func (o *KmsBatchSecretRequest) SetDelete(v []KmsBatchSecretRequestDeleteInner)`

SetDelete sets Delete field to given value.

### HasDelete

`func (o *KmsBatchSecretRequest) HasDelete() bool`

HasDelete returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


