# KmsIntegration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**IsActive** | Pointer to **bool** |  | [optional] 
**Integration** | Pointer to **string** | Target platform (e.g., aws-parameter-store, github, vercel) | [optional] 
**IntegrationAuthId** | Pointer to **string** |  | [optional] 
**EnvId** | Pointer to **string** |  | [optional] 
**SecretPath** | Pointer to **string** |  | [optional] 
**TargetEnvironment** | Pointer to **string** |  | [optional] 
**App** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewKmsIntegration

`func NewKmsIntegration() *KmsIntegration`

NewKmsIntegration instantiates a new KmsIntegration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKmsIntegrationWithDefaults

`func NewKmsIntegrationWithDefaults() *KmsIntegration`

NewKmsIntegrationWithDefaults instantiates a new KmsIntegration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KmsIntegration) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KmsIntegration) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KmsIntegration) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *KmsIntegration) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsActive

`func (o *KmsIntegration) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *KmsIntegration) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *KmsIntegration) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *KmsIntegration) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetIntegration

`func (o *KmsIntegration) GetIntegration() string`

GetIntegration returns the Integration field if non-nil, zero value otherwise.

### GetIntegrationOk

`func (o *KmsIntegration) GetIntegrationOk() (*string, bool)`

GetIntegrationOk returns a tuple with the Integration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegration

`func (o *KmsIntegration) SetIntegration(v string)`

SetIntegration sets Integration field to given value.

### HasIntegration

`func (o *KmsIntegration) HasIntegration() bool`

HasIntegration returns a boolean if a field has been set.

### GetIntegrationAuthId

`func (o *KmsIntegration) GetIntegrationAuthId() string`

GetIntegrationAuthId returns the IntegrationAuthId field if non-nil, zero value otherwise.

### GetIntegrationAuthIdOk

`func (o *KmsIntegration) GetIntegrationAuthIdOk() (*string, bool)`

GetIntegrationAuthIdOk returns a tuple with the IntegrationAuthId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationAuthId

`func (o *KmsIntegration) SetIntegrationAuthId(v string)`

SetIntegrationAuthId sets IntegrationAuthId field to given value.

### HasIntegrationAuthId

`func (o *KmsIntegration) HasIntegrationAuthId() bool`

HasIntegrationAuthId returns a boolean if a field has been set.

### GetEnvId

`func (o *KmsIntegration) GetEnvId() string`

GetEnvId returns the EnvId field if non-nil, zero value otherwise.

### GetEnvIdOk

`func (o *KmsIntegration) GetEnvIdOk() (*string, bool)`

GetEnvIdOk returns a tuple with the EnvId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvId

`func (o *KmsIntegration) SetEnvId(v string)`

SetEnvId sets EnvId field to given value.

### HasEnvId

`func (o *KmsIntegration) HasEnvId() bool`

HasEnvId returns a boolean if a field has been set.

### GetSecretPath

`func (o *KmsIntegration) GetSecretPath() string`

GetSecretPath returns the SecretPath field if non-nil, zero value otherwise.

### GetSecretPathOk

`func (o *KmsIntegration) GetSecretPathOk() (*string, bool)`

GetSecretPathOk returns a tuple with the SecretPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretPath

`func (o *KmsIntegration) SetSecretPath(v string)`

SetSecretPath sets SecretPath field to given value.

### HasSecretPath

`func (o *KmsIntegration) HasSecretPath() bool`

HasSecretPath returns a boolean if a field has been set.

### GetTargetEnvironment

`func (o *KmsIntegration) GetTargetEnvironment() string`

GetTargetEnvironment returns the TargetEnvironment field if non-nil, zero value otherwise.

### GetTargetEnvironmentOk

`func (o *KmsIntegration) GetTargetEnvironmentOk() (*string, bool)`

GetTargetEnvironmentOk returns a tuple with the TargetEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetEnvironment

`func (o *KmsIntegration) SetTargetEnvironment(v string)`

SetTargetEnvironment sets TargetEnvironment field to given value.

### HasTargetEnvironment

`func (o *KmsIntegration) HasTargetEnvironment() bool`

HasTargetEnvironment returns a boolean if a field has been set.

### GetApp

`func (o *KmsIntegration) GetApp() string`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *KmsIntegration) GetAppOk() (*string, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *KmsIntegration) SetApp(v string)`

SetApp sets App field to given value.

### HasApp

`func (o *KmsIntegration) HasApp() bool`

HasApp returns a boolean if a field has been set.

### GetCreatedAt

`func (o *KmsIntegration) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KmsIntegration) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KmsIntegration) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *KmsIntegration) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *KmsIntegration) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *KmsIntegration) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *KmsIntegration) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *KmsIntegration) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


