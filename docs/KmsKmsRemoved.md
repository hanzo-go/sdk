# KmsKmsRemoved

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Deleted** | Pointer to **bool** | Deleted is true; a delete confirms by not failing. | [optional] 
**Env** | Pointer to **string** | Env is the environment the secret was removed from. | [optional] 
**Name** | Pointer to **string** | Name is the secret&#39;s name. | [optional] 

## Methods

### NewKmsKmsRemoved

`func NewKmsKmsRemoved() *KmsKmsRemoved`

NewKmsKmsRemoved instantiates a new KmsKmsRemoved object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKmsKmsRemovedWithDefaults

`func NewKmsKmsRemovedWithDefaults() *KmsKmsRemoved`

NewKmsKmsRemovedWithDefaults instantiates a new KmsKmsRemoved object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeleted

`func (o *KmsKmsRemoved) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *KmsKmsRemoved) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *KmsKmsRemoved) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *KmsKmsRemoved) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetEnv

`func (o *KmsKmsRemoved) GetEnv() string`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *KmsKmsRemoved) GetEnvOk() (*string, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *KmsKmsRemoved) SetEnv(v string)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *KmsKmsRemoved) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### GetName

`func (o *KmsKmsRemoved) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KmsKmsRemoved) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KmsKmsRemoved) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KmsKmsRemoved) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


