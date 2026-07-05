# KmsKmsPutOrgSecretRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Env** | Pointer to **string** |  | [optional] [default to "default"]
**Value** | **string** |  | 

## Methods

### NewKmsKmsPutOrgSecretRequest

`func NewKmsKmsPutOrgSecretRequest(name string, value string, ) *KmsKmsPutOrgSecretRequest`

NewKmsKmsPutOrgSecretRequest instantiates a new KmsKmsPutOrgSecretRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKmsKmsPutOrgSecretRequestWithDefaults

`func NewKmsKmsPutOrgSecretRequestWithDefaults() *KmsKmsPutOrgSecretRequest`

NewKmsKmsPutOrgSecretRequestWithDefaults instantiates a new KmsKmsPutOrgSecretRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *KmsKmsPutOrgSecretRequest) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *KmsKmsPutOrgSecretRequest) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *KmsKmsPutOrgSecretRequest) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *KmsKmsPutOrgSecretRequest) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetName

`func (o *KmsKmsPutOrgSecretRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KmsKmsPutOrgSecretRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KmsKmsPutOrgSecretRequest) SetName(v string)`

SetName sets Name field to given value.


### GetEnv

`func (o *KmsKmsPutOrgSecretRequest) GetEnv() string`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *KmsKmsPutOrgSecretRequest) GetEnvOk() (*string, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *KmsKmsPutOrgSecretRequest) SetEnv(v string)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *KmsKmsPutOrgSecretRequest) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### GetValue

`func (o *KmsKmsPutOrgSecretRequest) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *KmsKmsPutOrgSecretRequest) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *KmsKmsPutOrgSecretRequest) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


