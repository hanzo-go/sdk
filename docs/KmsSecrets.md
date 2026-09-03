# KmsSecrets

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Names** | Pointer to **[]string** | Names is the same listing reduced to bare names, which is the shape the KMS operator reads. Both are emitted so either consumer keeps working. | [optional] 
**Secrets** | Pointer to [**[]SecretMeta**](SecretMeta.md) | Secrets are the descriptors: name, path, environment and sealing scheme. No value and no ciphertext appears here. | [optional] 
**Total** | Pointer to **int64** | Total is how many descriptors this listing carries. | [optional] 

## Methods

### NewKmsSecrets

`func NewKmsSecrets() *KmsSecrets`

NewKmsSecrets instantiates a new KmsSecrets object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKmsSecretsWithDefaults

`func NewKmsSecretsWithDefaults() *KmsSecrets`

NewKmsSecretsWithDefaults instantiates a new KmsSecrets object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNames

`func (o *KmsSecrets) GetNames() []string`

GetNames returns the Names field if non-nil, zero value otherwise.

### GetNamesOk

`func (o *KmsSecrets) GetNamesOk() (*[]string, bool)`

GetNamesOk returns a tuple with the Names field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNames

`func (o *KmsSecrets) SetNames(v []string)`

SetNames sets Names field to given value.

### HasNames

`func (o *KmsSecrets) HasNames() bool`

HasNames returns a boolean if a field has been set.

### GetSecrets

`func (o *KmsSecrets) GetSecrets() []SecretMeta`

GetSecrets returns the Secrets field if non-nil, zero value otherwise.

### GetSecretsOk

`func (o *KmsSecrets) GetSecretsOk() (*[]SecretMeta, bool)`

GetSecretsOk returns a tuple with the Secrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecrets

`func (o *KmsSecrets) SetSecrets(v []SecretMeta)`

SetSecrets sets Secrets field to given value.

### HasSecrets

`func (o *KmsSecrets) HasSecrets() bool`

HasSecrets returns a boolean if a field has been set.

### GetTotal

`func (o *KmsSecrets) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *KmsSecrets) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *KmsSecrets) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *KmsSecrets) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


