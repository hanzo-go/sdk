# KmsPostV1KmsSecretsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Env** | **string** |  | 
**Value** | **string** | The secret value. &#x60;format: password&#x60; is the ONE marker that makes a client read it from stdin instead of offering a flag, so it can never land in argv, &#x60;ps&#x60; or shell history. | 

## Methods

### NewKmsPostV1KmsSecretsRequest

`func NewKmsPostV1KmsSecretsRequest(name string, env string, value string, ) *KmsPostV1KmsSecretsRequest`

NewKmsPostV1KmsSecretsRequest instantiates a new KmsPostV1KmsSecretsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKmsPostV1KmsSecretsRequestWithDefaults

`func NewKmsPostV1KmsSecretsRequestWithDefaults() *KmsPostV1KmsSecretsRequest`

NewKmsPostV1KmsSecretsRequestWithDefaults instantiates a new KmsPostV1KmsSecretsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *KmsPostV1KmsSecretsRequest) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *KmsPostV1KmsSecretsRequest) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *KmsPostV1KmsSecretsRequest) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *KmsPostV1KmsSecretsRequest) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetName

`func (o *KmsPostV1KmsSecretsRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KmsPostV1KmsSecretsRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KmsPostV1KmsSecretsRequest) SetName(v string)`

SetName sets Name field to given value.


### GetEnv

`func (o *KmsPostV1KmsSecretsRequest) GetEnv() string`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *KmsPostV1KmsSecretsRequest) GetEnvOk() (*string, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *KmsPostV1KmsSecretsRequest) SetEnv(v string)`

SetEnv sets Env field to given value.


### GetValue

`func (o *KmsPostV1KmsSecretsRequest) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *KmsPostV1KmsSecretsRequest) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *KmsPostV1KmsSecretsRequest) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


