# KmsListSecrets200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Secrets** | Pointer to [**[]KmsSecret**](KmsSecret.md) |  | [optional] 
**Imports** | Pointer to [**[]KmsListSecrets200ResponseImportsInner**](KmsListSecrets200ResponseImportsInner.md) |  | [optional] 

## Methods

### NewKmsListSecrets200Response

`func NewKmsListSecrets200Response() *KmsListSecrets200Response`

NewKmsListSecrets200Response instantiates a new KmsListSecrets200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKmsListSecrets200ResponseWithDefaults

`func NewKmsListSecrets200ResponseWithDefaults() *KmsListSecrets200Response`

NewKmsListSecrets200ResponseWithDefaults instantiates a new KmsListSecrets200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecrets

`func (o *KmsListSecrets200Response) GetSecrets() []KmsSecret`

GetSecrets returns the Secrets field if non-nil, zero value otherwise.

### GetSecretsOk

`func (o *KmsListSecrets200Response) GetSecretsOk() (*[]KmsSecret, bool)`

GetSecretsOk returns a tuple with the Secrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecrets

`func (o *KmsListSecrets200Response) SetSecrets(v []KmsSecret)`

SetSecrets sets Secrets field to given value.

### HasSecrets

`func (o *KmsListSecrets200Response) HasSecrets() bool`

HasSecrets returns a boolean if a field has been set.

### GetImports

`func (o *KmsListSecrets200Response) GetImports() []KmsListSecrets200ResponseImportsInner`

GetImports returns the Imports field if non-nil, zero value otherwise.

### GetImportsOk

`func (o *KmsListSecrets200Response) GetImportsOk() (*[]KmsListSecrets200ResponseImportsInner, bool)`

GetImportsOk returns a tuple with the Imports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImports

`func (o *KmsListSecrets200Response) SetImports(v []KmsListSecrets200ResponseImportsInner)`

SetImports sets Imports field to given value.

### HasImports

`func (o *KmsListSecrets200Response) HasImports() bool`

HasImports returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


