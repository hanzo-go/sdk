# KmsKmsSecret

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Env** | Pointer to **string** | Env is the environment the value was resolved under. | [optional] 
**Name** | Pointer to **string** | Name is the secret&#39;s name. | [optional] 
**Value** | Pointer to **string** | Value is the opened plaintext. This body is the ONLY place it appears. | [optional] 

## Methods

### NewKmsKmsSecret

`func NewKmsKmsSecret() *KmsKmsSecret`

NewKmsKmsSecret instantiates a new KmsKmsSecret object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKmsKmsSecretWithDefaults

`func NewKmsKmsSecretWithDefaults() *KmsKmsSecret`

NewKmsKmsSecretWithDefaults instantiates a new KmsKmsSecret object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnv

`func (o *KmsKmsSecret) GetEnv() string`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *KmsKmsSecret) GetEnvOk() (*string, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *KmsKmsSecret) SetEnv(v string)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *KmsKmsSecret) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### GetName

`func (o *KmsKmsSecret) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KmsKmsSecret) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KmsKmsSecret) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KmsKmsSecret) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValue

`func (o *KmsKmsSecret) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *KmsKmsSecret) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *KmsKmsSecret) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *KmsKmsSecret) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


