# SocialSocialProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CredentialsConfigured** | Pointer to **bool** | CredentialsConfigured is whether this deployment holds every OAuth application credential the network needs. It is a statement about the DEPLOYMENT, not about the caller: a connected account also needs its own access token before a post can go out. | [optional] 
**MissingCredentials** | Pointer to **[]string** | MissingCredentials names the environment variables still unset for this network, so the answer is an installation instruction rather than a refusal. Absent when the credentials are complete. Only the NAMES appear here; a credential value is never reported. | [optional] 
**Provider** | Pointer to **string** | Provider is the network: x, facebook, instagram, linkedin, tiktok, youtube or threads.  Example: \&quot;x\&quot; | [optional] 

## Methods

### NewSocialSocialProvider

`func NewSocialSocialProvider() *SocialSocialProvider`

NewSocialSocialProvider instantiates a new SocialSocialProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSocialSocialProviderWithDefaults

`func NewSocialSocialProviderWithDefaults() *SocialSocialProvider`

NewSocialSocialProviderWithDefaults instantiates a new SocialSocialProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentialsConfigured

`func (o *SocialSocialProvider) GetCredentialsConfigured() bool`

GetCredentialsConfigured returns the CredentialsConfigured field if non-nil, zero value otherwise.

### GetCredentialsConfiguredOk

`func (o *SocialSocialProvider) GetCredentialsConfiguredOk() (*bool, bool)`

GetCredentialsConfiguredOk returns a tuple with the CredentialsConfigured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialsConfigured

`func (o *SocialSocialProvider) SetCredentialsConfigured(v bool)`

SetCredentialsConfigured sets CredentialsConfigured field to given value.

### HasCredentialsConfigured

`func (o *SocialSocialProvider) HasCredentialsConfigured() bool`

HasCredentialsConfigured returns a boolean if a field has been set.

### GetMissingCredentials

`func (o *SocialSocialProvider) GetMissingCredentials() []string`

GetMissingCredentials returns the MissingCredentials field if non-nil, zero value otherwise.

### GetMissingCredentialsOk

`func (o *SocialSocialProvider) GetMissingCredentialsOk() (*[]string, bool)`

GetMissingCredentialsOk returns a tuple with the MissingCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissingCredentials

`func (o *SocialSocialProvider) SetMissingCredentials(v []string)`

SetMissingCredentials sets MissingCredentials field to given value.

### HasMissingCredentials

`func (o *SocialSocialProvider) HasMissingCredentials() bool`

HasMissingCredentials returns a boolean if a field has been set.

### GetProvider

`func (o *SocialSocialProvider) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *SocialSocialProvider) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *SocialSocialProvider) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *SocialSocialProvider) HasProvider() bool`

HasProvider returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


