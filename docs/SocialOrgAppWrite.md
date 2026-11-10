# SocialOrgAppWrite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | ClientID is the full app id, never a mask. Empty clears the app. | [optional] 
**ClientSecret** | Pointer to **string** | ClientSecret is the secret. Absent leaves the sealed value. An empty string clears it. | [optional] 
**Provider** | Pointer to **string** | Provider is the network to write. | [optional] 

## Methods

### NewSocialOrgAppWrite

`func NewSocialOrgAppWrite() *SocialOrgAppWrite`

NewSocialOrgAppWrite instantiates a new SocialOrgAppWrite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSocialOrgAppWriteWithDefaults

`func NewSocialOrgAppWriteWithDefaults() *SocialOrgAppWrite`

NewSocialOrgAppWriteWithDefaults instantiates a new SocialOrgAppWrite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *SocialOrgAppWrite) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *SocialOrgAppWrite) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *SocialOrgAppWrite) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *SocialOrgAppWrite) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *SocialOrgAppWrite) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *SocialOrgAppWrite) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *SocialOrgAppWrite) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *SocialOrgAppWrite) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetProvider

`func (o *SocialOrgAppWrite) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *SocialOrgAppWrite) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *SocialOrgAppWrite) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *SocialOrgAppWrite) HasProvider() bool`

HasProvider returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


