# O11yOAuth2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TLSConfig** | Pointer to [**O11yTLSConfig**](O11yTLSConfig.md) |  | [optional] 
**Audience** | Pointer to **string** | Audience optionally specifies the intended audience of the request.  If empty, the value of TokenURL is used as the intended audience. Only used if GrantType is set to \&quot;urn:ietf:params:oauth:grant-type:jwt-bearer\&quot;. | [optional] 
**Claims** | Pointer to **map[string]map[string]interface{}** | Claims is a map of claims to be added to the JWT token. Only used if GrantType is set to \&quot;urn:ietf:params:oauth:grant-type:jwt-bearer\&quot;. | [optional] 
**ClientCertificateKey** | Pointer to **interface{}** |  | [optional] 
**ClientCertificateKeyFile** | Pointer to **string** |  | [optional] 
**ClientCertificateKeyId** | Pointer to **string** |  | [optional] 
**ClientCertificateKeyRef** | Pointer to **string** | ClientCertificateKeyRef is the name of the secret within the secret manager to use as the client secret. | [optional] 
**ClientId** | Pointer to **string** |  | [optional] 
**ClientSecret** | Pointer to **interface{}** |  | [optional] 
**ClientSecretFile** | Pointer to **string** |  | [optional] 
**ClientSecretRef** | Pointer to **string** | ClientSecretRef is the name of the secret within the secret manager to use as the client secret. | [optional] 
**EndpointParams** | Pointer to **map[string]string** |  | [optional] 
**GrantType** | Pointer to **string** | GrantType is the OAuth2 grant type to use. It can be one of \&quot;client_credentials\&quot; or \&quot;urn:ietf:params:oauth:grant-type:jwt-bearer\&quot; (RFC 7523). Default value is \&quot;client_credentials\&quot; | [optional] 
**Iss** | Pointer to **string** | Iss is the OAuth client identifier used when communicating with the configured OAuth provider. Default value is client_id. Only used if GrantType is set to \&quot;urn:ietf:params:oauth:grant-type:jwt-bearer\&quot;. | [optional] 
**NoProxy** | Pointer to **string** |  | [optional] 
**ProxyConnectHeader** | Pointer to **map[string][]interface{}** |  | [optional] 
**ProxyFromEnvironment** | Pointer to **bool** |  | [optional] 
**ProxyUrl** | Pointer to **interface{}** |  | [optional] 
**Scopes** | Pointer to **[]string** |  | [optional] 
**SignatureAlgorithm** | Pointer to **string** | SignatureAlgorithm is the RSA algorithm used to sign JWT token. Only used if GrantType is set to \&quot;urn:ietf:params:oauth:grant-type:jwt-bearer\&quot;. Default value is RS256 and valid values RS256, RS384, RS512 | [optional] 
**TokenUrl** | Pointer to **string** |  | [optional] 

## Methods

### NewO11yOAuth2

`func NewO11yOAuth2() *O11yOAuth2`

NewO11yOAuth2 instantiates a new O11yOAuth2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yOAuth2WithDefaults

`func NewO11yOAuth2WithDefaults() *O11yOAuth2`

NewO11yOAuth2WithDefaults instantiates a new O11yOAuth2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTLSConfig

`func (o *O11yOAuth2) GetTLSConfig() O11yTLSConfig`

GetTLSConfig returns the TLSConfig field if non-nil, zero value otherwise.

### GetTLSConfigOk

`func (o *O11yOAuth2) GetTLSConfigOk() (*O11yTLSConfig, bool)`

GetTLSConfigOk returns a tuple with the TLSConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTLSConfig

`func (o *O11yOAuth2) SetTLSConfig(v O11yTLSConfig)`

SetTLSConfig sets TLSConfig field to given value.

### HasTLSConfig

`func (o *O11yOAuth2) HasTLSConfig() bool`

HasTLSConfig returns a boolean if a field has been set.

### GetAudience

`func (o *O11yOAuth2) GetAudience() string`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *O11yOAuth2) GetAudienceOk() (*string, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *O11yOAuth2) SetAudience(v string)`

SetAudience sets Audience field to given value.

### HasAudience

`func (o *O11yOAuth2) HasAudience() bool`

HasAudience returns a boolean if a field has been set.

### GetClaims

`func (o *O11yOAuth2) GetClaims() map[string]map[string]interface{}`

GetClaims returns the Claims field if non-nil, zero value otherwise.

### GetClaimsOk

`func (o *O11yOAuth2) GetClaimsOk() (*map[string]map[string]interface{}, bool)`

GetClaimsOk returns a tuple with the Claims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaims

`func (o *O11yOAuth2) SetClaims(v map[string]map[string]interface{})`

SetClaims sets Claims field to given value.

### HasClaims

`func (o *O11yOAuth2) HasClaims() bool`

HasClaims returns a boolean if a field has been set.

### GetClientCertificateKey

`func (o *O11yOAuth2) GetClientCertificateKey() interface{}`

GetClientCertificateKey returns the ClientCertificateKey field if non-nil, zero value otherwise.

### GetClientCertificateKeyOk

`func (o *O11yOAuth2) GetClientCertificateKeyOk() (*interface{}, bool)`

GetClientCertificateKeyOk returns a tuple with the ClientCertificateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCertificateKey

`func (o *O11yOAuth2) SetClientCertificateKey(v interface{})`

SetClientCertificateKey sets ClientCertificateKey field to given value.

### HasClientCertificateKey

`func (o *O11yOAuth2) HasClientCertificateKey() bool`

HasClientCertificateKey returns a boolean if a field has been set.

### SetClientCertificateKeyNil

`func (o *O11yOAuth2) SetClientCertificateKeyNil(b bool)`

 SetClientCertificateKeyNil sets the value for ClientCertificateKey to be an explicit nil

### UnsetClientCertificateKey
`func (o *O11yOAuth2) UnsetClientCertificateKey()`

UnsetClientCertificateKey ensures that no value is present for ClientCertificateKey, not even an explicit nil
### GetClientCertificateKeyFile

`func (o *O11yOAuth2) GetClientCertificateKeyFile() string`

GetClientCertificateKeyFile returns the ClientCertificateKeyFile field if non-nil, zero value otherwise.

### GetClientCertificateKeyFileOk

`func (o *O11yOAuth2) GetClientCertificateKeyFileOk() (*string, bool)`

GetClientCertificateKeyFileOk returns a tuple with the ClientCertificateKeyFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCertificateKeyFile

`func (o *O11yOAuth2) SetClientCertificateKeyFile(v string)`

SetClientCertificateKeyFile sets ClientCertificateKeyFile field to given value.

### HasClientCertificateKeyFile

`func (o *O11yOAuth2) HasClientCertificateKeyFile() bool`

HasClientCertificateKeyFile returns a boolean if a field has been set.

### GetClientCertificateKeyId

`func (o *O11yOAuth2) GetClientCertificateKeyId() string`

GetClientCertificateKeyId returns the ClientCertificateKeyId field if non-nil, zero value otherwise.

### GetClientCertificateKeyIdOk

`func (o *O11yOAuth2) GetClientCertificateKeyIdOk() (*string, bool)`

GetClientCertificateKeyIdOk returns a tuple with the ClientCertificateKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCertificateKeyId

`func (o *O11yOAuth2) SetClientCertificateKeyId(v string)`

SetClientCertificateKeyId sets ClientCertificateKeyId field to given value.

### HasClientCertificateKeyId

`func (o *O11yOAuth2) HasClientCertificateKeyId() bool`

HasClientCertificateKeyId returns a boolean if a field has been set.

### GetClientCertificateKeyRef

`func (o *O11yOAuth2) GetClientCertificateKeyRef() string`

GetClientCertificateKeyRef returns the ClientCertificateKeyRef field if non-nil, zero value otherwise.

### GetClientCertificateKeyRefOk

`func (o *O11yOAuth2) GetClientCertificateKeyRefOk() (*string, bool)`

GetClientCertificateKeyRefOk returns a tuple with the ClientCertificateKeyRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCertificateKeyRef

`func (o *O11yOAuth2) SetClientCertificateKeyRef(v string)`

SetClientCertificateKeyRef sets ClientCertificateKeyRef field to given value.

### HasClientCertificateKeyRef

`func (o *O11yOAuth2) HasClientCertificateKeyRef() bool`

HasClientCertificateKeyRef returns a boolean if a field has been set.

### GetClientId

`func (o *O11yOAuth2) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *O11yOAuth2) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *O11yOAuth2) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *O11yOAuth2) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *O11yOAuth2) GetClientSecret() interface{}`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *O11yOAuth2) GetClientSecretOk() (*interface{}, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *O11yOAuth2) SetClientSecret(v interface{})`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *O11yOAuth2) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### SetClientSecretNil

`func (o *O11yOAuth2) SetClientSecretNil(b bool)`

 SetClientSecretNil sets the value for ClientSecret to be an explicit nil

### UnsetClientSecret
`func (o *O11yOAuth2) UnsetClientSecret()`

UnsetClientSecret ensures that no value is present for ClientSecret, not even an explicit nil
### GetClientSecretFile

`func (o *O11yOAuth2) GetClientSecretFile() string`

GetClientSecretFile returns the ClientSecretFile field if non-nil, zero value otherwise.

### GetClientSecretFileOk

`func (o *O11yOAuth2) GetClientSecretFileOk() (*string, bool)`

GetClientSecretFileOk returns a tuple with the ClientSecretFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecretFile

`func (o *O11yOAuth2) SetClientSecretFile(v string)`

SetClientSecretFile sets ClientSecretFile field to given value.

### HasClientSecretFile

`func (o *O11yOAuth2) HasClientSecretFile() bool`

HasClientSecretFile returns a boolean if a field has been set.

### GetClientSecretRef

`func (o *O11yOAuth2) GetClientSecretRef() string`

GetClientSecretRef returns the ClientSecretRef field if non-nil, zero value otherwise.

### GetClientSecretRefOk

`func (o *O11yOAuth2) GetClientSecretRefOk() (*string, bool)`

GetClientSecretRefOk returns a tuple with the ClientSecretRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecretRef

`func (o *O11yOAuth2) SetClientSecretRef(v string)`

SetClientSecretRef sets ClientSecretRef field to given value.

### HasClientSecretRef

`func (o *O11yOAuth2) HasClientSecretRef() bool`

HasClientSecretRef returns a boolean if a field has been set.

### GetEndpointParams

`func (o *O11yOAuth2) GetEndpointParams() map[string]string`

GetEndpointParams returns the EndpointParams field if non-nil, zero value otherwise.

### GetEndpointParamsOk

`func (o *O11yOAuth2) GetEndpointParamsOk() (*map[string]string, bool)`

GetEndpointParamsOk returns a tuple with the EndpointParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointParams

`func (o *O11yOAuth2) SetEndpointParams(v map[string]string)`

SetEndpointParams sets EndpointParams field to given value.

### HasEndpointParams

`func (o *O11yOAuth2) HasEndpointParams() bool`

HasEndpointParams returns a boolean if a field has been set.

### GetGrantType

`func (o *O11yOAuth2) GetGrantType() string`

GetGrantType returns the GrantType field if non-nil, zero value otherwise.

### GetGrantTypeOk

`func (o *O11yOAuth2) GetGrantTypeOk() (*string, bool)`

GetGrantTypeOk returns a tuple with the GrantType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantType

`func (o *O11yOAuth2) SetGrantType(v string)`

SetGrantType sets GrantType field to given value.

### HasGrantType

`func (o *O11yOAuth2) HasGrantType() bool`

HasGrantType returns a boolean if a field has been set.

### GetIss

`func (o *O11yOAuth2) GetIss() string`

GetIss returns the Iss field if non-nil, zero value otherwise.

### GetIssOk

`func (o *O11yOAuth2) GetIssOk() (*string, bool)`

GetIssOk returns a tuple with the Iss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIss

`func (o *O11yOAuth2) SetIss(v string)`

SetIss sets Iss field to given value.

### HasIss

`func (o *O11yOAuth2) HasIss() bool`

HasIss returns a boolean if a field has been set.

### GetNoProxy

`func (o *O11yOAuth2) GetNoProxy() string`

GetNoProxy returns the NoProxy field if non-nil, zero value otherwise.

### GetNoProxyOk

`func (o *O11yOAuth2) GetNoProxyOk() (*string, bool)`

GetNoProxyOk returns a tuple with the NoProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoProxy

`func (o *O11yOAuth2) SetNoProxy(v string)`

SetNoProxy sets NoProxy field to given value.

### HasNoProxy

`func (o *O11yOAuth2) HasNoProxy() bool`

HasNoProxy returns a boolean if a field has been set.

### GetProxyConnectHeader

`func (o *O11yOAuth2) GetProxyConnectHeader() map[string][]interface{}`

GetProxyConnectHeader returns the ProxyConnectHeader field if non-nil, zero value otherwise.

### GetProxyConnectHeaderOk

`func (o *O11yOAuth2) GetProxyConnectHeaderOk() (*map[string][]interface{}, bool)`

GetProxyConnectHeaderOk returns a tuple with the ProxyConnectHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyConnectHeader

`func (o *O11yOAuth2) SetProxyConnectHeader(v map[string][]interface{})`

SetProxyConnectHeader sets ProxyConnectHeader field to given value.

### HasProxyConnectHeader

`func (o *O11yOAuth2) HasProxyConnectHeader() bool`

HasProxyConnectHeader returns a boolean if a field has been set.

### GetProxyFromEnvironment

`func (o *O11yOAuth2) GetProxyFromEnvironment() bool`

GetProxyFromEnvironment returns the ProxyFromEnvironment field if non-nil, zero value otherwise.

### GetProxyFromEnvironmentOk

`func (o *O11yOAuth2) GetProxyFromEnvironmentOk() (*bool, bool)`

GetProxyFromEnvironmentOk returns a tuple with the ProxyFromEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyFromEnvironment

`func (o *O11yOAuth2) SetProxyFromEnvironment(v bool)`

SetProxyFromEnvironment sets ProxyFromEnvironment field to given value.

### HasProxyFromEnvironment

`func (o *O11yOAuth2) HasProxyFromEnvironment() bool`

HasProxyFromEnvironment returns a boolean if a field has been set.

### GetProxyUrl

`func (o *O11yOAuth2) GetProxyUrl() interface{}`

GetProxyUrl returns the ProxyUrl field if non-nil, zero value otherwise.

### GetProxyUrlOk

`func (o *O11yOAuth2) GetProxyUrlOk() (*interface{}, bool)`

GetProxyUrlOk returns a tuple with the ProxyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyUrl

`func (o *O11yOAuth2) SetProxyUrl(v interface{})`

SetProxyUrl sets ProxyUrl field to given value.

### HasProxyUrl

`func (o *O11yOAuth2) HasProxyUrl() bool`

HasProxyUrl returns a boolean if a field has been set.

### SetProxyUrlNil

`func (o *O11yOAuth2) SetProxyUrlNil(b bool)`

 SetProxyUrlNil sets the value for ProxyUrl to be an explicit nil

### UnsetProxyUrl
`func (o *O11yOAuth2) UnsetProxyUrl()`

UnsetProxyUrl ensures that no value is present for ProxyUrl, not even an explicit nil
### GetScopes

`func (o *O11yOAuth2) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *O11yOAuth2) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *O11yOAuth2) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *O11yOAuth2) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetSignatureAlgorithm

`func (o *O11yOAuth2) GetSignatureAlgorithm() string`

GetSignatureAlgorithm returns the SignatureAlgorithm field if non-nil, zero value otherwise.

### GetSignatureAlgorithmOk

`func (o *O11yOAuth2) GetSignatureAlgorithmOk() (*string, bool)`

GetSignatureAlgorithmOk returns a tuple with the SignatureAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureAlgorithm

`func (o *O11yOAuth2) SetSignatureAlgorithm(v string)`

SetSignatureAlgorithm sets SignatureAlgorithm field to given value.

### HasSignatureAlgorithm

`func (o *O11yOAuth2) HasSignatureAlgorithm() bool`

HasSignatureAlgorithm returns a boolean if a field has been set.

### GetTokenUrl

`func (o *O11yOAuth2) GetTokenUrl() string`

GetTokenUrl returns the TokenUrl field if non-nil, zero value otherwise.

### GetTokenUrlOk

`func (o *O11yOAuth2) GetTokenUrlOk() (*string, bool)`

GetTokenUrlOk returns a tuple with the TokenUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenUrl

`func (o *O11yOAuth2) SetTokenUrl(v string)`

SetTokenUrl sets TokenUrl field to given value.

### HasTokenUrl

`func (o *O11yOAuth2) HasTokenUrl() bool`

HasTokenUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


