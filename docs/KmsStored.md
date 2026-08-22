# KmsStored

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Env** | Pointer to **string** | Env is the environment the secret was written under. | [optional] 
**Name** | Pointer to **string** | Name is the secret&#39;s name. | [optional] 
**Stored** | Pointer to **bool** | Stored is true; a write confirms by not failing. | [optional] 

## Methods

### NewKmsStored

`func NewKmsStored() *KmsStored`

NewKmsStored instantiates a new KmsStored object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKmsStoredWithDefaults

`func NewKmsStoredWithDefaults() *KmsStored`

NewKmsStoredWithDefaults instantiates a new KmsStored object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnv

`func (o *KmsStored) GetEnv() string`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *KmsStored) GetEnvOk() (*string, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *KmsStored) SetEnv(v string)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *KmsStored) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### GetName

`func (o *KmsStored) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KmsStored) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KmsStored) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KmsStored) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStored

`func (o *KmsStored) GetStored() bool`

GetStored returns the Stored field if non-nil, zero value otherwise.

### GetStoredOk

`func (o *KmsStored) GetStoredOk() (*bool, bool)`

GetStoredOk returns a tuple with the Stored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStored

`func (o *KmsStored) SetStored(v bool)`

SetStored sets Stored field to given value.

### HasStored

`func (o *KmsStored) HasStored() bool`

HasStored returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


