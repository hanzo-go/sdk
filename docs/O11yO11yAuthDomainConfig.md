# O11yO11yAuthDomainConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GoogleAuthConfig** | Pointer to [**O11yO11yGoogleConfig**](O11yO11yGoogleConfig.md) | Google is the Google provider&#39;s settings, when SSOType is google_auth. | [optional] 
**OidcConfig** | Pointer to [**O11yO11yOIDCConfig**](O11yO11yOIDCConfig.md) | OIDC is the OIDC provider&#39;s settings, when SSOType is oidc. | [optional] 
**RoleMapping** | Pointer to [**O11yO11yRoleMapping**](O11yO11yRoleMapping.md) | RoleMapping maps the provider&#39;s groups onto roles for new users. | [optional] 
**SamlConfig** | Pointer to [**O11yO11ySAMLConfig**](O11yO11ySAMLConfig.md) | SAML is the SAML provider&#39;s settings, when SSOType is saml. | [optional] 
**SsoEnabled** | Pointer to **bool** | SSOEnabled turns enforced SSO on for the domain. | [optional] 
**SsoType** | Pointer to **string** | SSOType picks the provider — saml, google_auth or oidc. | [optional] 

## Methods

### NewO11yO11yAuthDomainConfig

`func NewO11yO11yAuthDomainConfig() *O11yO11yAuthDomainConfig`

NewO11yO11yAuthDomainConfig instantiates a new O11yO11yAuthDomainConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yAuthDomainConfigWithDefaults

`func NewO11yO11yAuthDomainConfigWithDefaults() *O11yO11yAuthDomainConfig`

NewO11yO11yAuthDomainConfigWithDefaults instantiates a new O11yO11yAuthDomainConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGoogleAuthConfig

`func (o *O11yO11yAuthDomainConfig) GetGoogleAuthConfig() O11yO11yGoogleConfig`

GetGoogleAuthConfig returns the GoogleAuthConfig field if non-nil, zero value otherwise.

### GetGoogleAuthConfigOk

`func (o *O11yO11yAuthDomainConfig) GetGoogleAuthConfigOk() (*O11yO11yGoogleConfig, bool)`

GetGoogleAuthConfigOk returns a tuple with the GoogleAuthConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleAuthConfig

`func (o *O11yO11yAuthDomainConfig) SetGoogleAuthConfig(v O11yO11yGoogleConfig)`

SetGoogleAuthConfig sets GoogleAuthConfig field to given value.

### HasGoogleAuthConfig

`func (o *O11yO11yAuthDomainConfig) HasGoogleAuthConfig() bool`

HasGoogleAuthConfig returns a boolean if a field has been set.

### GetOidcConfig

`func (o *O11yO11yAuthDomainConfig) GetOidcConfig() O11yO11yOIDCConfig`

GetOidcConfig returns the OidcConfig field if non-nil, zero value otherwise.

### GetOidcConfigOk

`func (o *O11yO11yAuthDomainConfig) GetOidcConfigOk() (*O11yO11yOIDCConfig, bool)`

GetOidcConfigOk returns a tuple with the OidcConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcConfig

`func (o *O11yO11yAuthDomainConfig) SetOidcConfig(v O11yO11yOIDCConfig)`

SetOidcConfig sets OidcConfig field to given value.

### HasOidcConfig

`func (o *O11yO11yAuthDomainConfig) HasOidcConfig() bool`

HasOidcConfig returns a boolean if a field has been set.

### GetRoleMapping

`func (o *O11yO11yAuthDomainConfig) GetRoleMapping() O11yO11yRoleMapping`

GetRoleMapping returns the RoleMapping field if non-nil, zero value otherwise.

### GetRoleMappingOk

`func (o *O11yO11yAuthDomainConfig) GetRoleMappingOk() (*O11yO11yRoleMapping, bool)`

GetRoleMappingOk returns a tuple with the RoleMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleMapping

`func (o *O11yO11yAuthDomainConfig) SetRoleMapping(v O11yO11yRoleMapping)`

SetRoleMapping sets RoleMapping field to given value.

### HasRoleMapping

`func (o *O11yO11yAuthDomainConfig) HasRoleMapping() bool`

HasRoleMapping returns a boolean if a field has been set.

### GetSamlConfig

`func (o *O11yO11yAuthDomainConfig) GetSamlConfig() O11yO11ySAMLConfig`

GetSamlConfig returns the SamlConfig field if non-nil, zero value otherwise.

### GetSamlConfigOk

`func (o *O11yO11yAuthDomainConfig) GetSamlConfigOk() (*O11yO11ySAMLConfig, bool)`

GetSamlConfigOk returns a tuple with the SamlConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlConfig

`func (o *O11yO11yAuthDomainConfig) SetSamlConfig(v O11yO11ySAMLConfig)`

SetSamlConfig sets SamlConfig field to given value.

### HasSamlConfig

`func (o *O11yO11yAuthDomainConfig) HasSamlConfig() bool`

HasSamlConfig returns a boolean if a field has been set.

### GetSsoEnabled

`func (o *O11yO11yAuthDomainConfig) GetSsoEnabled() bool`

GetSsoEnabled returns the SsoEnabled field if non-nil, zero value otherwise.

### GetSsoEnabledOk

`func (o *O11yO11yAuthDomainConfig) GetSsoEnabledOk() (*bool, bool)`

GetSsoEnabledOk returns a tuple with the SsoEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoEnabled

`func (o *O11yO11yAuthDomainConfig) SetSsoEnabled(v bool)`

SetSsoEnabled sets SsoEnabled field to given value.

### HasSsoEnabled

`func (o *O11yO11yAuthDomainConfig) HasSsoEnabled() bool`

HasSsoEnabled returns a boolean if a field has been set.

### GetSsoType

`func (o *O11yO11yAuthDomainConfig) GetSsoType() string`

GetSsoType returns the SsoType field if non-nil, zero value otherwise.

### GetSsoTypeOk

`func (o *O11yO11yAuthDomainConfig) GetSsoTypeOk() (*string, bool)`

GetSsoTypeOk returns a tuple with the SsoType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoType

`func (o *O11yO11yAuthDomainConfig) SetSsoType(v string)`

SetSsoType sets SsoType field to given value.

### HasSsoType

`func (o *O11yO11yAuthDomainConfig) HasSsoType() bool`

HasSsoType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


