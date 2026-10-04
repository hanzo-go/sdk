# KmsKmsSecrets

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Names** | Pointer to **[]string** | Names is the same listing reduced to bare names, which is the shape the KMS operator reads. Both are emitted so either consumer keeps working. | [optional] 
**Secrets** | Pointer to [**[]KmsSecretMeta**](KmsSecretMeta.md) | Secrets are the descriptors: name, path, environment and sealing scheme. No value and no ciphertext appears here. | [optional] 
**Total** | Pointer to **int64** | Total is how many descriptors this listing carries. | [optional] 

## Methods

### NewKmsKmsSecrets

`func NewKmsKmsSecrets() *KmsKmsSecrets`

NewKmsKmsSecrets instantiates a new KmsKmsSecrets object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKmsKmsSecretsWithDefaults

`func NewKmsKmsSecretsWithDefaults() *KmsKmsSecrets`

NewKmsKmsSecretsWithDefaults instantiates a new KmsKmsSecrets object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNames

`func (o *KmsKmsSecrets) GetNames() []string`

GetNames returns the Names field if non-nil, zero value otherwise.

### GetNamesOk

`func (o *KmsKmsSecrets) GetNamesOk() (*[]string, bool)`

GetNamesOk returns a tuple with the Names field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNames

`func (o *KmsKmsSecrets) SetNames(v []string)`

SetNames sets Names field to given value.

### HasNames

`func (o *KmsKmsSecrets) HasNames() bool`

HasNames returns a boolean if a field has been set.

### GetSecrets

`func (o *KmsKmsSecrets) GetSecrets() []KmsSecretMeta`

GetSecrets returns the Secrets field if non-nil, zero value otherwise.

### GetSecretsOk

`func (o *KmsKmsSecrets) GetSecretsOk() (*[]KmsSecretMeta, bool)`

GetSecretsOk returns a tuple with the Secrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecrets

`func (o *KmsKmsSecrets) SetSecrets(v []KmsSecretMeta)`

SetSecrets sets Secrets field to given value.

### HasSecrets

`func (o *KmsKmsSecrets) HasSecrets() bool`

HasSecrets returns a boolean if a field has been set.

### GetTotal

`func (o *KmsKmsSecrets) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *KmsKmsSecrets) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *KmsKmsSecrets) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *KmsKmsSecrets) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


