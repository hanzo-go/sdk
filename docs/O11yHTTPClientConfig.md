# O11yHTTPClientConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Authorization** | Pointer to [**O11yAuthorization**](O11yAuthorization.md) | The HTTP authorization credentials for the targets. | [optional] 
**BasicAuth** | Pointer to [**O11yBasicAuth**](O11yBasicAuth.md) | The HTTP basic authentication credentials for the targets. | [optional] 
**BearerToken** | Pointer to **interface{}** |  | [optional] 
**BearerTokenFile** | Pointer to **string** | The bearer token file for the targets. Deprecated in favour of Authorization.CredentialsFile. | [optional] 
**EnableHttp2** | Pointer to **bool** | EnableHTTP2 specifies whether the client should configure HTTP2. The omitempty flag is not set, because it would be hidden from the marshalled configuration when set to false. | [optional] 
**FollowRedirects** | Pointer to **bool** | FollowRedirects specifies whether the client should follow HTTP 3xx redirects. The omitempty flag is not set, because it would be hidden from the marshalled configuration when set to false. | [optional] 
**HttpHeaders** | Pointer to **interface{}** |  | [optional] 
**NoProxy** | Pointer to **string** |  | [optional] 
**Oauth2** | Pointer to [**O11yOAuth2**](O11yOAuth2.md) | The OAuth2 client credentials used to fetch a token for the targets. | [optional] 
**ProxyConnectHeader** | Pointer to **map[string][]interface{}** |  | [optional] 
**ProxyFromEnvironment** | Pointer to **bool** |  | [optional] 
**ProxyUrl** | Pointer to **interface{}** |  | [optional] 
**TlsConfig** | Pointer to [**O11yTLSConfig**](O11yTLSConfig.md) | TLSConfig to use to connect to the targets. | [optional] 

## Methods

### NewO11yHTTPClientConfig

`func NewO11yHTTPClientConfig() *O11yHTTPClientConfig`

NewO11yHTTPClientConfig instantiates a new O11yHTTPClientConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yHTTPClientConfigWithDefaults

`func NewO11yHTTPClientConfigWithDefaults() *O11yHTTPClientConfig`

NewO11yHTTPClientConfigWithDefaults instantiates a new O11yHTTPClientConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorization

`func (o *O11yHTTPClientConfig) GetAuthorization() O11yAuthorization`

GetAuthorization returns the Authorization field if non-nil, zero value otherwise.

### GetAuthorizationOk

`func (o *O11yHTTPClientConfig) GetAuthorizationOk() (*O11yAuthorization, bool)`

GetAuthorizationOk returns a tuple with the Authorization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorization

`func (o *O11yHTTPClientConfig) SetAuthorization(v O11yAuthorization)`

SetAuthorization sets Authorization field to given value.

### HasAuthorization

`func (o *O11yHTTPClientConfig) HasAuthorization() bool`

HasAuthorization returns a boolean if a field has been set.

### GetBasicAuth

`func (o *O11yHTTPClientConfig) GetBasicAuth() O11yBasicAuth`

GetBasicAuth returns the BasicAuth field if non-nil, zero value otherwise.

### GetBasicAuthOk

`func (o *O11yHTTPClientConfig) GetBasicAuthOk() (*O11yBasicAuth, bool)`

GetBasicAuthOk returns a tuple with the BasicAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicAuth

`func (o *O11yHTTPClientConfig) SetBasicAuth(v O11yBasicAuth)`

SetBasicAuth sets BasicAuth field to given value.

### HasBasicAuth

`func (o *O11yHTTPClientConfig) HasBasicAuth() bool`

HasBasicAuth returns a boolean if a field has been set.

### GetBearerToken

`func (o *O11yHTTPClientConfig) GetBearerToken() interface{}`

GetBearerToken returns the BearerToken field if non-nil, zero value otherwise.

### GetBearerTokenOk

`func (o *O11yHTTPClientConfig) GetBearerTokenOk() (*interface{}, bool)`

GetBearerTokenOk returns a tuple with the BearerToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBearerToken

`func (o *O11yHTTPClientConfig) SetBearerToken(v interface{})`

SetBearerToken sets BearerToken field to given value.

### HasBearerToken

`func (o *O11yHTTPClientConfig) HasBearerToken() bool`

HasBearerToken returns a boolean if a field has been set.

### SetBearerTokenNil

`func (o *O11yHTTPClientConfig) SetBearerTokenNil(b bool)`

 SetBearerTokenNil sets the value for BearerToken to be an explicit nil

### UnsetBearerToken
`func (o *O11yHTTPClientConfig) UnsetBearerToken()`

UnsetBearerToken ensures that no value is present for BearerToken, not even an explicit nil
### GetBearerTokenFile

`func (o *O11yHTTPClientConfig) GetBearerTokenFile() string`

GetBearerTokenFile returns the BearerTokenFile field if non-nil, zero value otherwise.

### GetBearerTokenFileOk

`func (o *O11yHTTPClientConfig) GetBearerTokenFileOk() (*string, bool)`

GetBearerTokenFileOk returns a tuple with the BearerTokenFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBearerTokenFile

`func (o *O11yHTTPClientConfig) SetBearerTokenFile(v string)`

SetBearerTokenFile sets BearerTokenFile field to given value.

### HasBearerTokenFile

`func (o *O11yHTTPClientConfig) HasBearerTokenFile() bool`

HasBearerTokenFile returns a boolean if a field has been set.

### GetEnableHttp2

`func (o *O11yHTTPClientConfig) GetEnableHttp2() bool`

GetEnableHttp2 returns the EnableHttp2 field if non-nil, zero value otherwise.

### GetEnableHttp2Ok

`func (o *O11yHTTPClientConfig) GetEnableHttp2Ok() (*bool, bool)`

GetEnableHttp2Ok returns a tuple with the EnableHttp2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableHttp2

`func (o *O11yHTTPClientConfig) SetEnableHttp2(v bool)`

SetEnableHttp2 sets EnableHttp2 field to given value.

### HasEnableHttp2

`func (o *O11yHTTPClientConfig) HasEnableHttp2() bool`

HasEnableHttp2 returns a boolean if a field has been set.

### GetFollowRedirects

`func (o *O11yHTTPClientConfig) GetFollowRedirects() bool`

GetFollowRedirects returns the FollowRedirects field if non-nil, zero value otherwise.

### GetFollowRedirectsOk

`func (o *O11yHTTPClientConfig) GetFollowRedirectsOk() (*bool, bool)`

GetFollowRedirectsOk returns a tuple with the FollowRedirects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowRedirects

`func (o *O11yHTTPClientConfig) SetFollowRedirects(v bool)`

SetFollowRedirects sets FollowRedirects field to given value.

### HasFollowRedirects

`func (o *O11yHTTPClientConfig) HasFollowRedirects() bool`

HasFollowRedirects returns a boolean if a field has been set.

### GetHttpHeaders

`func (o *O11yHTTPClientConfig) GetHttpHeaders() interface{}`

GetHttpHeaders returns the HttpHeaders field if non-nil, zero value otherwise.

### GetHttpHeadersOk

`func (o *O11yHTTPClientConfig) GetHttpHeadersOk() (*interface{}, bool)`

GetHttpHeadersOk returns a tuple with the HttpHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpHeaders

`func (o *O11yHTTPClientConfig) SetHttpHeaders(v interface{})`

SetHttpHeaders sets HttpHeaders field to given value.

### HasHttpHeaders

`func (o *O11yHTTPClientConfig) HasHttpHeaders() bool`

HasHttpHeaders returns a boolean if a field has been set.

### SetHttpHeadersNil

`func (o *O11yHTTPClientConfig) SetHttpHeadersNil(b bool)`

 SetHttpHeadersNil sets the value for HttpHeaders to be an explicit nil

### UnsetHttpHeaders
`func (o *O11yHTTPClientConfig) UnsetHttpHeaders()`

UnsetHttpHeaders ensures that no value is present for HttpHeaders, not even an explicit nil
### GetNoProxy

`func (o *O11yHTTPClientConfig) GetNoProxy() string`

GetNoProxy returns the NoProxy field if non-nil, zero value otherwise.

### GetNoProxyOk

`func (o *O11yHTTPClientConfig) GetNoProxyOk() (*string, bool)`

GetNoProxyOk returns a tuple with the NoProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoProxy

`func (o *O11yHTTPClientConfig) SetNoProxy(v string)`

SetNoProxy sets NoProxy field to given value.

### HasNoProxy

`func (o *O11yHTTPClientConfig) HasNoProxy() bool`

HasNoProxy returns a boolean if a field has been set.

### GetOauth2

`func (o *O11yHTTPClientConfig) GetOauth2() O11yOAuth2`

GetOauth2 returns the Oauth2 field if non-nil, zero value otherwise.

### GetOauth2Ok

`func (o *O11yHTTPClientConfig) GetOauth2Ok() (*O11yOAuth2, bool)`

GetOauth2Ok returns a tuple with the Oauth2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth2

`func (o *O11yHTTPClientConfig) SetOauth2(v O11yOAuth2)`

SetOauth2 sets Oauth2 field to given value.

### HasOauth2

`func (o *O11yHTTPClientConfig) HasOauth2() bool`

HasOauth2 returns a boolean if a field has been set.

### GetProxyConnectHeader

`func (o *O11yHTTPClientConfig) GetProxyConnectHeader() map[string][]interface{}`

GetProxyConnectHeader returns the ProxyConnectHeader field if non-nil, zero value otherwise.

### GetProxyConnectHeaderOk

`func (o *O11yHTTPClientConfig) GetProxyConnectHeaderOk() (*map[string][]interface{}, bool)`

GetProxyConnectHeaderOk returns a tuple with the ProxyConnectHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyConnectHeader

`func (o *O11yHTTPClientConfig) SetProxyConnectHeader(v map[string][]interface{})`

SetProxyConnectHeader sets ProxyConnectHeader field to given value.

### HasProxyConnectHeader

`func (o *O11yHTTPClientConfig) HasProxyConnectHeader() bool`

HasProxyConnectHeader returns a boolean if a field has been set.

### GetProxyFromEnvironment

`func (o *O11yHTTPClientConfig) GetProxyFromEnvironment() bool`

GetProxyFromEnvironment returns the ProxyFromEnvironment field if non-nil, zero value otherwise.

### GetProxyFromEnvironmentOk

`func (o *O11yHTTPClientConfig) GetProxyFromEnvironmentOk() (*bool, bool)`

GetProxyFromEnvironmentOk returns a tuple with the ProxyFromEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyFromEnvironment

`func (o *O11yHTTPClientConfig) SetProxyFromEnvironment(v bool)`

SetProxyFromEnvironment sets ProxyFromEnvironment field to given value.

### HasProxyFromEnvironment

`func (o *O11yHTTPClientConfig) HasProxyFromEnvironment() bool`

HasProxyFromEnvironment returns a boolean if a field has been set.

### GetProxyUrl

`func (o *O11yHTTPClientConfig) GetProxyUrl() interface{}`

GetProxyUrl returns the ProxyUrl field if non-nil, zero value otherwise.

### GetProxyUrlOk

`func (o *O11yHTTPClientConfig) GetProxyUrlOk() (*interface{}, bool)`

GetProxyUrlOk returns a tuple with the ProxyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyUrl

`func (o *O11yHTTPClientConfig) SetProxyUrl(v interface{})`

SetProxyUrl sets ProxyUrl field to given value.

### HasProxyUrl

`func (o *O11yHTTPClientConfig) HasProxyUrl() bool`

HasProxyUrl returns a boolean if a field has been set.

### SetProxyUrlNil

`func (o *O11yHTTPClientConfig) SetProxyUrlNil(b bool)`

 SetProxyUrlNil sets the value for ProxyUrl to be an explicit nil

### UnsetProxyUrl
`func (o *O11yHTTPClientConfig) UnsetProxyUrl()`

UnsetProxyUrl ensures that no value is present for ProxyUrl, not even an explicit nil
### GetTlsConfig

`func (o *O11yHTTPClientConfig) GetTlsConfig() O11yTLSConfig`

GetTlsConfig returns the TlsConfig field if non-nil, zero value otherwise.

### GetTlsConfigOk

`func (o *O11yHTTPClientConfig) GetTlsConfigOk() (*O11yTLSConfig, bool)`

GetTlsConfigOk returns a tuple with the TlsConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsConfig

`func (o *O11yHTTPClientConfig) SetTlsConfig(v O11yTLSConfig)`

SetTlsConfig sets TlsConfig field to given value.

### HasTlsConfig

`func (o *O11yHTTPClientConfig) HasTlsConfig() bool`

HasTlsConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


