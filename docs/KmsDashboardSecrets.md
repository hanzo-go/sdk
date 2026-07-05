# KmsDashboardSecrets

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Secrets** | Pointer to [**[]KmsSecret**](KmsSecret.md) |  | [optional] 
**Folders** | Pointer to [**[]KmsSecretFolder**](KmsSecretFolder.md) |  | [optional] 
**Imports** | Pointer to [**[]KmsSecretImport**](KmsSecretImport.md) |  | [optional] 

## Methods

### NewKmsDashboardSecrets

`func NewKmsDashboardSecrets() *KmsDashboardSecrets`

NewKmsDashboardSecrets instantiates a new KmsDashboardSecrets object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKmsDashboardSecretsWithDefaults

`func NewKmsDashboardSecretsWithDefaults() *KmsDashboardSecrets`

NewKmsDashboardSecretsWithDefaults instantiates a new KmsDashboardSecrets object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecrets

`func (o *KmsDashboardSecrets) GetSecrets() []KmsSecret`

GetSecrets returns the Secrets field if non-nil, zero value otherwise.

### GetSecretsOk

`func (o *KmsDashboardSecrets) GetSecretsOk() (*[]KmsSecret, bool)`

GetSecretsOk returns a tuple with the Secrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecrets

`func (o *KmsDashboardSecrets) SetSecrets(v []KmsSecret)`

SetSecrets sets Secrets field to given value.

### HasSecrets

`func (o *KmsDashboardSecrets) HasSecrets() bool`

HasSecrets returns a boolean if a field has been set.

### GetFolders

`func (o *KmsDashboardSecrets) GetFolders() []KmsSecretFolder`

GetFolders returns the Folders field if non-nil, zero value otherwise.

### GetFoldersOk

`func (o *KmsDashboardSecrets) GetFoldersOk() (*[]KmsSecretFolder, bool)`

GetFoldersOk returns a tuple with the Folders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolders

`func (o *KmsDashboardSecrets) SetFolders(v []KmsSecretFolder)`

SetFolders sets Folders field to given value.

### HasFolders

`func (o *KmsDashboardSecrets) HasFolders() bool`

HasFolders returns a boolean if a field has been set.

### GetImports

`func (o *KmsDashboardSecrets) GetImports() []KmsSecretImport`

GetImports returns the Imports field if non-nil, zero value otherwise.

### GetImportsOk

`func (o *KmsDashboardSecrets) GetImportsOk() (*[]KmsSecretImport, bool)`

GetImportsOk returns a tuple with the Imports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImports

`func (o *KmsDashboardSecrets) SetImports(v []KmsSecretImport)`

SetImports sets Imports field to given value.

### HasImports

`func (o *KmsDashboardSecrets) HasImports() bool`

HasImports returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


