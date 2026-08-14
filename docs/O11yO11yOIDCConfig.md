# O11yO11yOIDCConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClaimMapping** | Pointer to [**O11yO11yAttributeMapping**](O11yO11yAttributeMapping.md) | ClaimMapping names the token claims to read identity from. | [optional] 
**ClientId** | Pointer to **string** | ClientID is the OAuth application&#39;s id. | [optional] 
**ClientSecret** | Pointer to **string** | ClientSecret is the OAuth application&#39;s secret. | [optional] 
**GetUserInfo** | Pointer to **bool** | GetUserInfo also queries the userinfo endpoint, for providers whose id tokens are thin. | [optional] 
**InsecureSkipEmailVerified** | Pointer to **bool** | InsecureSkipEmailVerified admits addresses the provider has not verified. | [optional] 
**Issuer** | Pointer to **string** | Issuer is the provider&#39;s issuer URL. | [optional] 
**IssuerAlias** | Pointer to **string** | IssuerAlias overrides the issuer for providers whose discovery document disagrees with their issuer URL. | [optional] 

## Methods

### NewO11yO11yOIDCConfig

`func NewO11yO11yOIDCConfig() *O11yO11yOIDCConfig`

NewO11yO11yOIDCConfig instantiates a new O11yO11yOIDCConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yOIDCConfigWithDefaults

`func NewO11yO11yOIDCConfigWithDefaults() *O11yO11yOIDCConfig`

NewO11yO11yOIDCConfigWithDefaults instantiates a new O11yO11yOIDCConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClaimMapping

`func (o *O11yO11yOIDCConfig) GetClaimMapping() O11yO11yAttributeMapping`

GetClaimMapping returns the ClaimMapping field if non-nil, zero value otherwise.

### GetClaimMappingOk

`func (o *O11yO11yOIDCConfig) GetClaimMappingOk() (*O11yO11yAttributeMapping, bool)`

GetClaimMappingOk returns a tuple with the ClaimMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimMapping

`func (o *O11yO11yOIDCConfig) SetClaimMapping(v O11yO11yAttributeMapping)`

SetClaimMapping sets ClaimMapping field to given value.

### HasClaimMapping

`func (o *O11yO11yOIDCConfig) HasClaimMapping() bool`

HasClaimMapping returns a boolean if a field has been set.

### GetClientId

`func (o *O11yO11yOIDCConfig) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *O11yO11yOIDCConfig) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *O11yO11yOIDCConfig) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *O11yO11yOIDCConfig) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *O11yO11yOIDCConfig) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *O11yO11yOIDCConfig) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *O11yO11yOIDCConfig) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *O11yO11yOIDCConfig) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetGetUserInfo

`func (o *O11yO11yOIDCConfig) GetGetUserInfo() bool`

GetGetUserInfo returns the GetUserInfo field if non-nil, zero value otherwise.

### GetGetUserInfoOk

`func (o *O11yO11yOIDCConfig) GetGetUserInfoOk() (*bool, bool)`

GetGetUserInfoOk returns a tuple with the GetUserInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGetUserInfo

`func (o *O11yO11yOIDCConfig) SetGetUserInfo(v bool)`

SetGetUserInfo sets GetUserInfo field to given value.

### HasGetUserInfo

`func (o *O11yO11yOIDCConfig) HasGetUserInfo() bool`

HasGetUserInfo returns a boolean if a field has been set.

### GetInsecureSkipEmailVerified

`func (o *O11yO11yOIDCConfig) GetInsecureSkipEmailVerified() bool`

GetInsecureSkipEmailVerified returns the InsecureSkipEmailVerified field if non-nil, zero value otherwise.

### GetInsecureSkipEmailVerifiedOk

`func (o *O11yO11yOIDCConfig) GetInsecureSkipEmailVerifiedOk() (*bool, bool)`

GetInsecureSkipEmailVerifiedOk returns a tuple with the InsecureSkipEmailVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsecureSkipEmailVerified

`func (o *O11yO11yOIDCConfig) SetInsecureSkipEmailVerified(v bool)`

SetInsecureSkipEmailVerified sets InsecureSkipEmailVerified field to given value.

### HasInsecureSkipEmailVerified

`func (o *O11yO11yOIDCConfig) HasInsecureSkipEmailVerified() bool`

HasInsecureSkipEmailVerified returns a boolean if a field has been set.

### GetIssuer

`func (o *O11yO11yOIDCConfig) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *O11yO11yOIDCConfig) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *O11yO11yOIDCConfig) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *O11yO11yOIDCConfig) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetIssuerAlias

`func (o *O11yO11yOIDCConfig) GetIssuerAlias() string`

GetIssuerAlias returns the IssuerAlias field if non-nil, zero value otherwise.

### GetIssuerAliasOk

`func (o *O11yO11yOIDCConfig) GetIssuerAliasOk() (*string, bool)`

GetIssuerAliasOk returns a tuple with the IssuerAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerAlias

`func (o *O11yO11yOIDCConfig) SetIssuerAlias(v string)`

SetIssuerAlias sets IssuerAlias field to given value.

### HasIssuerAlias

`func (o *O11yO11yOIDCConfig) HasIssuerAlias() bool`

HasIssuerAlias returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


