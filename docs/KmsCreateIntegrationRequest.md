# KmsCreateIntegrationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IntegrationAuthId** | **string** |  | 
**Integration** | **string** |  | 
**App** | **string** |  | 
**SourceEnvironment** | **string** |  | 
**SecretPath** | **string** |  | [default to "/"]
**TargetEnvironment** | Pointer to **string** |  | [optional] 

## Methods

### NewKmsCreateIntegrationRequest

`func NewKmsCreateIntegrationRequest(integrationAuthId string, integration string, app string, sourceEnvironment string, secretPath string, ) *KmsCreateIntegrationRequest`

NewKmsCreateIntegrationRequest instantiates a new KmsCreateIntegrationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKmsCreateIntegrationRequestWithDefaults

`func NewKmsCreateIntegrationRequestWithDefaults() *KmsCreateIntegrationRequest`

NewKmsCreateIntegrationRequestWithDefaults instantiates a new KmsCreateIntegrationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIntegrationAuthId

`func (o *KmsCreateIntegrationRequest) GetIntegrationAuthId() string`

GetIntegrationAuthId returns the IntegrationAuthId field if non-nil, zero value otherwise.

### GetIntegrationAuthIdOk

`func (o *KmsCreateIntegrationRequest) GetIntegrationAuthIdOk() (*string, bool)`

GetIntegrationAuthIdOk returns a tuple with the IntegrationAuthId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationAuthId

`func (o *KmsCreateIntegrationRequest) SetIntegrationAuthId(v string)`

SetIntegrationAuthId sets IntegrationAuthId field to given value.


### GetIntegration

`func (o *KmsCreateIntegrationRequest) GetIntegration() string`

GetIntegration returns the Integration field if non-nil, zero value otherwise.

### GetIntegrationOk

`func (o *KmsCreateIntegrationRequest) GetIntegrationOk() (*string, bool)`

GetIntegrationOk returns a tuple with the Integration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegration

`func (o *KmsCreateIntegrationRequest) SetIntegration(v string)`

SetIntegration sets Integration field to given value.


### GetApp

`func (o *KmsCreateIntegrationRequest) GetApp() string`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *KmsCreateIntegrationRequest) GetAppOk() (*string, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *KmsCreateIntegrationRequest) SetApp(v string)`

SetApp sets App field to given value.


### GetSourceEnvironment

`func (o *KmsCreateIntegrationRequest) GetSourceEnvironment() string`

GetSourceEnvironment returns the SourceEnvironment field if non-nil, zero value otherwise.

### GetSourceEnvironmentOk

`func (o *KmsCreateIntegrationRequest) GetSourceEnvironmentOk() (*string, bool)`

GetSourceEnvironmentOk returns a tuple with the SourceEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceEnvironment

`func (o *KmsCreateIntegrationRequest) SetSourceEnvironment(v string)`

SetSourceEnvironment sets SourceEnvironment field to given value.


### GetSecretPath

`func (o *KmsCreateIntegrationRequest) GetSecretPath() string`

GetSecretPath returns the SecretPath field if non-nil, zero value otherwise.

### GetSecretPathOk

`func (o *KmsCreateIntegrationRequest) GetSecretPathOk() (*string, bool)`

GetSecretPathOk returns a tuple with the SecretPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretPath

`func (o *KmsCreateIntegrationRequest) SetSecretPath(v string)`

SetSecretPath sets SecretPath field to given value.


### GetTargetEnvironment

`func (o *KmsCreateIntegrationRequest) GetTargetEnvironment() string`

GetTargetEnvironment returns the TargetEnvironment field if non-nil, zero value otherwise.

### GetTargetEnvironmentOk

`func (o *KmsCreateIntegrationRequest) GetTargetEnvironmentOk() (*string, bool)`

GetTargetEnvironmentOk returns a tuple with the TargetEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetEnvironment

`func (o *KmsCreateIntegrationRequest) SetTargetEnvironment(v string)`

SetTargetEnvironment sets TargetEnvironment field to given value.

### HasTargetEnvironment

`func (o *KmsCreateIntegrationRequest) HasTargetEnvironment() bool`

HasTargetEnvironment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


