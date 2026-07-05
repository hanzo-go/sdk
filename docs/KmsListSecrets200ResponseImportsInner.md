# KmsListSecrets200ResponseImportsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Environment** | Pointer to **string** |  | [optional] 
**SecretPath** | Pointer to **string** |  | [optional] 
**FolderId** | Pointer to **string** |  | [optional] 
**Secrets** | Pointer to [**[]KmsSecret**](KmsSecret.md) |  | [optional] 

## Methods

### NewKmsListSecrets200ResponseImportsInner

`func NewKmsListSecrets200ResponseImportsInner() *KmsListSecrets200ResponseImportsInner`

NewKmsListSecrets200ResponseImportsInner instantiates a new KmsListSecrets200ResponseImportsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKmsListSecrets200ResponseImportsInnerWithDefaults

`func NewKmsListSecrets200ResponseImportsInnerWithDefaults() *KmsListSecrets200ResponseImportsInner`

NewKmsListSecrets200ResponseImportsInnerWithDefaults instantiates a new KmsListSecrets200ResponseImportsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironment

`func (o *KmsListSecrets200ResponseImportsInner) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *KmsListSecrets200ResponseImportsInner) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *KmsListSecrets200ResponseImportsInner) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *KmsListSecrets200ResponseImportsInner) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetSecretPath

`func (o *KmsListSecrets200ResponseImportsInner) GetSecretPath() string`

GetSecretPath returns the SecretPath field if non-nil, zero value otherwise.

### GetSecretPathOk

`func (o *KmsListSecrets200ResponseImportsInner) GetSecretPathOk() (*string, bool)`

GetSecretPathOk returns a tuple with the SecretPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretPath

`func (o *KmsListSecrets200ResponseImportsInner) SetSecretPath(v string)`

SetSecretPath sets SecretPath field to given value.

### HasSecretPath

`func (o *KmsListSecrets200ResponseImportsInner) HasSecretPath() bool`

HasSecretPath returns a boolean if a field has been set.

### GetFolderId

`func (o *KmsListSecrets200ResponseImportsInner) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *KmsListSecrets200ResponseImportsInner) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *KmsListSecrets200ResponseImportsInner) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *KmsListSecrets200ResponseImportsInner) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### GetSecrets

`func (o *KmsListSecrets200ResponseImportsInner) GetSecrets() []KmsSecret`

GetSecrets returns the Secrets field if non-nil, zero value otherwise.

### GetSecretsOk

`func (o *KmsListSecrets200ResponseImportsInner) GetSecretsOk() (*[]KmsSecret, bool)`

GetSecretsOk returns a tuple with the Secrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecrets

`func (o *KmsListSecrets200ResponseImportsInner) SetSecrets(v []KmsSecret)`

SetSecrets sets Secrets field to given value.

### HasSecrets

`func (o *KmsListSecrets200ResponseImportsInner) HasSecrets() bool`

HasSecrets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


