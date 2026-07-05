# KmsCreateWebhookRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkspaceId** | **string** |  | 
**Environment** | **string** |  | 
**WebhookUrl** | **string** |  | 
**SecretPath** | Pointer to **string** |  | [optional] [default to "/"]
**WebhookSecretKey** | Pointer to **string** |  | [optional] 

## Methods

### NewKmsCreateWebhookRequest

`func NewKmsCreateWebhookRequest(workspaceId string, environment string, webhookUrl string, ) *KmsCreateWebhookRequest`

NewKmsCreateWebhookRequest instantiates a new KmsCreateWebhookRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKmsCreateWebhookRequestWithDefaults

`func NewKmsCreateWebhookRequestWithDefaults() *KmsCreateWebhookRequest`

NewKmsCreateWebhookRequestWithDefaults instantiates a new KmsCreateWebhookRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkspaceId

`func (o *KmsCreateWebhookRequest) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *KmsCreateWebhookRequest) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *KmsCreateWebhookRequest) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.


### GetEnvironment

`func (o *KmsCreateWebhookRequest) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *KmsCreateWebhookRequest) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *KmsCreateWebhookRequest) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.


### GetWebhookUrl

`func (o *KmsCreateWebhookRequest) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *KmsCreateWebhookRequest) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *KmsCreateWebhookRequest) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.


### GetSecretPath

`func (o *KmsCreateWebhookRequest) GetSecretPath() string`

GetSecretPath returns the SecretPath field if non-nil, zero value otherwise.

### GetSecretPathOk

`func (o *KmsCreateWebhookRequest) GetSecretPathOk() (*string, bool)`

GetSecretPathOk returns a tuple with the SecretPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretPath

`func (o *KmsCreateWebhookRequest) SetSecretPath(v string)`

SetSecretPath sets SecretPath field to given value.

### HasSecretPath

`func (o *KmsCreateWebhookRequest) HasSecretPath() bool`

HasSecretPath returns a boolean if a field has been set.

### GetWebhookSecretKey

`func (o *KmsCreateWebhookRequest) GetWebhookSecretKey() string`

GetWebhookSecretKey returns the WebhookSecretKey field if non-nil, zero value otherwise.

### GetWebhookSecretKeyOk

`func (o *KmsCreateWebhookRequest) GetWebhookSecretKeyOk() (*string, bool)`

GetWebhookSecretKeyOk returns a tuple with the WebhookSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookSecretKey

`func (o *KmsCreateWebhookRequest) SetWebhookSecretKey(v string)`

SetWebhookSecretKey sets WebhookSecretKey field to given value.

### HasWebhookSecretKey

`func (o *KmsCreateWebhookRequest) HasWebhookSecretKey() bool`

HasWebhookSecretKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


