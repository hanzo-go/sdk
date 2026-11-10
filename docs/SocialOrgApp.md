# SocialOrgApp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | ClientID is the public app id, in the clear. The screen masks it. | [optional] 
**Provider** | Pointer to **string** | Provider is the network: x, facebook, instagram, linkedin, tiktok, youtube or threads. | [optional] 
**SecretSet** | Pointer to **bool** | SecretSet is whether a secret is sealed for this app. The secret itself is not in this response. | [optional] 

## Methods

### NewSocialOrgApp

`func NewSocialOrgApp() *SocialOrgApp`

NewSocialOrgApp instantiates a new SocialOrgApp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSocialOrgAppWithDefaults

`func NewSocialOrgAppWithDefaults() *SocialOrgApp`

NewSocialOrgAppWithDefaults instantiates a new SocialOrgApp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *SocialOrgApp) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *SocialOrgApp) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *SocialOrgApp) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *SocialOrgApp) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetProvider

`func (o *SocialOrgApp) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *SocialOrgApp) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *SocialOrgApp) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *SocialOrgApp) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetSecretSet

`func (o *SocialOrgApp) GetSecretSet() bool`

GetSecretSet returns the SecretSet field if non-nil, zero value otherwise.

### GetSecretSetOk

`func (o *SocialOrgApp) GetSecretSetOk() (*bool, bool)`

GetSecretSetOk returns a tuple with the SecretSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretSet

`func (o *SocialOrgApp) SetSecretSet(v bool)`

SetSecretSet sets SecretSet field to given value.

### HasSecretSet

`func (o *SocialOrgApp) HasSecretSet() bool`

HasSecretSet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


