# KmsLogin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | ClientID is the machine identity&#39;s id, as IAM issued it. | [optional] 
**ClientSecret** | Pointer to **string** | ClientSecret is that identity&#39;s secret. It is never logged, never echoed, and never carried in an error. | [optional] 

## Methods

### NewKmsLogin

`func NewKmsLogin() *KmsLogin`

NewKmsLogin instantiates a new KmsLogin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKmsLoginWithDefaults

`func NewKmsLoginWithDefaults() *KmsLogin`

NewKmsLoginWithDefaults instantiates a new KmsLogin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *KmsLogin) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *KmsLogin) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *KmsLogin) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *KmsLogin) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *KmsLogin) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *KmsLogin) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *KmsLogin) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *KmsLogin) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


