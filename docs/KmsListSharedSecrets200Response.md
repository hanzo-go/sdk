# KmsListSharedSecrets200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Secrets** | Pointer to [**[]KmsSharedSecret**](KmsSharedSecret.md) |  | [optional] 

## Methods

### NewKmsListSharedSecrets200Response

`func NewKmsListSharedSecrets200Response() *KmsListSharedSecrets200Response`

NewKmsListSharedSecrets200Response instantiates a new KmsListSharedSecrets200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKmsListSharedSecrets200ResponseWithDefaults

`func NewKmsListSharedSecrets200ResponseWithDefaults() *KmsListSharedSecrets200Response`

NewKmsListSharedSecrets200ResponseWithDefaults instantiates a new KmsListSharedSecrets200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecrets

`func (o *KmsListSharedSecrets200Response) GetSecrets() []KmsSharedSecret`

GetSecrets returns the Secrets field if non-nil, zero value otherwise.

### GetSecretsOk

`func (o *KmsListSharedSecrets200Response) GetSecretsOk() (*[]KmsSharedSecret, bool)`

GetSecretsOk returns a tuple with the Secrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecrets

`func (o *KmsListSharedSecrets200Response) SetSecrets(v []KmsSharedSecret)`

SetSecrets sets Secrets field to given value.

### HasSecrets

`func (o *KmsListSharedSecrets200Response) HasSecrets() bool`

HasSecrets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


