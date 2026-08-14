# O11yEmailConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotifierConfig** | Pointer to [**O11yNotifierConfig**](O11yNotifierConfig.md) |  | [optional] 
**AuthIdentity** | Pointer to **string** |  | [optional] 
**AuthPassword** | Pointer to **interface{}** |  | [optional] 
**AuthPasswordFile** | Pointer to **string** |  | [optional] 
**AuthSecret** | Pointer to **interface{}** |  | [optional] 
**AuthSecretFile** | Pointer to **string** |  | [optional] 
**AuthUsername** | Pointer to **string** |  | [optional] 
**ForceImplicitTls** | Pointer to **bool** | ForceImplicitTLS controls whether to use implicit TLS (direct TLS connection). true: force use of implicit TLS (direct TLS connection) false: force disable implicit TLS (use explicit TLS/STARTTLS if required) nil (default): auto-detect based on port (465&#x3D;implicit, other&#x3D;explicit) for backward compatibility | [optional] 
**From** | Pointer to **string** |  | [optional] 
**Headers** | Pointer to **map[string]string** |  | [optional] 
**Hello** | Pointer to **string** |  | [optional] 
**Html** | Pointer to **string** |  | [optional] 
**RequireTls** | Pointer to **bool** |  | [optional] 
**Smarthost** | Pointer to **interface{}** |  | [optional] 
**Text** | Pointer to **string** |  | [optional] 
**Threading** | Pointer to [**O11yThreadingConfig**](O11yThreadingConfig.md) |  | [optional] 
**TlsConfig** | Pointer to [**O11yTLSConfig**](O11yTLSConfig.md) |  | [optional] 
**To** | Pointer to **string** | Email address to notify. | [optional] 

## Methods

### NewO11yEmailConfig

`func NewO11yEmailConfig() *O11yEmailConfig`

NewO11yEmailConfig instantiates a new O11yEmailConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yEmailConfigWithDefaults

`func NewO11yEmailConfigWithDefaults() *O11yEmailConfig`

NewO11yEmailConfigWithDefaults instantiates a new O11yEmailConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotifierConfig

`func (o *O11yEmailConfig) GetNotifierConfig() O11yNotifierConfig`

GetNotifierConfig returns the NotifierConfig field if non-nil, zero value otherwise.

### GetNotifierConfigOk

`func (o *O11yEmailConfig) GetNotifierConfigOk() (*O11yNotifierConfig, bool)`

GetNotifierConfigOk returns a tuple with the NotifierConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifierConfig

`func (o *O11yEmailConfig) SetNotifierConfig(v O11yNotifierConfig)`

SetNotifierConfig sets NotifierConfig field to given value.

### HasNotifierConfig

`func (o *O11yEmailConfig) HasNotifierConfig() bool`

HasNotifierConfig returns a boolean if a field has been set.

### GetAuthIdentity

`func (o *O11yEmailConfig) GetAuthIdentity() string`

GetAuthIdentity returns the AuthIdentity field if non-nil, zero value otherwise.

### GetAuthIdentityOk

`func (o *O11yEmailConfig) GetAuthIdentityOk() (*string, bool)`

GetAuthIdentityOk returns a tuple with the AuthIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthIdentity

`func (o *O11yEmailConfig) SetAuthIdentity(v string)`

SetAuthIdentity sets AuthIdentity field to given value.

### HasAuthIdentity

`func (o *O11yEmailConfig) HasAuthIdentity() bool`

HasAuthIdentity returns a boolean if a field has been set.

### GetAuthPassword

`func (o *O11yEmailConfig) GetAuthPassword() interface{}`

GetAuthPassword returns the AuthPassword field if non-nil, zero value otherwise.

### GetAuthPasswordOk

`func (o *O11yEmailConfig) GetAuthPasswordOk() (*interface{}, bool)`

GetAuthPasswordOk returns a tuple with the AuthPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthPassword

`func (o *O11yEmailConfig) SetAuthPassword(v interface{})`

SetAuthPassword sets AuthPassword field to given value.

### HasAuthPassword

`func (o *O11yEmailConfig) HasAuthPassword() bool`

HasAuthPassword returns a boolean if a field has been set.

### SetAuthPasswordNil

`func (o *O11yEmailConfig) SetAuthPasswordNil(b bool)`

 SetAuthPasswordNil sets the value for AuthPassword to be an explicit nil

### UnsetAuthPassword
`func (o *O11yEmailConfig) UnsetAuthPassword()`

UnsetAuthPassword ensures that no value is present for AuthPassword, not even an explicit nil
### GetAuthPasswordFile

`func (o *O11yEmailConfig) GetAuthPasswordFile() string`

GetAuthPasswordFile returns the AuthPasswordFile field if non-nil, zero value otherwise.

### GetAuthPasswordFileOk

`func (o *O11yEmailConfig) GetAuthPasswordFileOk() (*string, bool)`

GetAuthPasswordFileOk returns a tuple with the AuthPasswordFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthPasswordFile

`func (o *O11yEmailConfig) SetAuthPasswordFile(v string)`

SetAuthPasswordFile sets AuthPasswordFile field to given value.

### HasAuthPasswordFile

`func (o *O11yEmailConfig) HasAuthPasswordFile() bool`

HasAuthPasswordFile returns a boolean if a field has been set.

### GetAuthSecret

`func (o *O11yEmailConfig) GetAuthSecret() interface{}`

GetAuthSecret returns the AuthSecret field if non-nil, zero value otherwise.

### GetAuthSecretOk

`func (o *O11yEmailConfig) GetAuthSecretOk() (*interface{}, bool)`

GetAuthSecretOk returns a tuple with the AuthSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthSecret

`func (o *O11yEmailConfig) SetAuthSecret(v interface{})`

SetAuthSecret sets AuthSecret field to given value.

### HasAuthSecret

`func (o *O11yEmailConfig) HasAuthSecret() bool`

HasAuthSecret returns a boolean if a field has been set.

### SetAuthSecretNil

`func (o *O11yEmailConfig) SetAuthSecretNil(b bool)`

 SetAuthSecretNil sets the value for AuthSecret to be an explicit nil

### UnsetAuthSecret
`func (o *O11yEmailConfig) UnsetAuthSecret()`

UnsetAuthSecret ensures that no value is present for AuthSecret, not even an explicit nil
### GetAuthSecretFile

`func (o *O11yEmailConfig) GetAuthSecretFile() string`

GetAuthSecretFile returns the AuthSecretFile field if non-nil, zero value otherwise.

### GetAuthSecretFileOk

`func (o *O11yEmailConfig) GetAuthSecretFileOk() (*string, bool)`

GetAuthSecretFileOk returns a tuple with the AuthSecretFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthSecretFile

`func (o *O11yEmailConfig) SetAuthSecretFile(v string)`

SetAuthSecretFile sets AuthSecretFile field to given value.

### HasAuthSecretFile

`func (o *O11yEmailConfig) HasAuthSecretFile() bool`

HasAuthSecretFile returns a boolean if a field has been set.

### GetAuthUsername

`func (o *O11yEmailConfig) GetAuthUsername() string`

GetAuthUsername returns the AuthUsername field if non-nil, zero value otherwise.

### GetAuthUsernameOk

`func (o *O11yEmailConfig) GetAuthUsernameOk() (*string, bool)`

GetAuthUsernameOk returns a tuple with the AuthUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthUsername

`func (o *O11yEmailConfig) SetAuthUsername(v string)`

SetAuthUsername sets AuthUsername field to given value.

### HasAuthUsername

`func (o *O11yEmailConfig) HasAuthUsername() bool`

HasAuthUsername returns a boolean if a field has been set.

### GetForceImplicitTls

`func (o *O11yEmailConfig) GetForceImplicitTls() bool`

GetForceImplicitTls returns the ForceImplicitTls field if non-nil, zero value otherwise.

### GetForceImplicitTlsOk

`func (o *O11yEmailConfig) GetForceImplicitTlsOk() (*bool, bool)`

GetForceImplicitTlsOk returns a tuple with the ForceImplicitTls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceImplicitTls

`func (o *O11yEmailConfig) SetForceImplicitTls(v bool)`

SetForceImplicitTls sets ForceImplicitTls field to given value.

### HasForceImplicitTls

`func (o *O11yEmailConfig) HasForceImplicitTls() bool`

HasForceImplicitTls returns a boolean if a field has been set.

### GetFrom

`func (o *O11yEmailConfig) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *O11yEmailConfig) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *O11yEmailConfig) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *O11yEmailConfig) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetHeaders

`func (o *O11yEmailConfig) GetHeaders() map[string]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *O11yEmailConfig) GetHeadersOk() (*map[string]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *O11yEmailConfig) SetHeaders(v map[string]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *O11yEmailConfig) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetHello

`func (o *O11yEmailConfig) GetHello() string`

GetHello returns the Hello field if non-nil, zero value otherwise.

### GetHelloOk

`func (o *O11yEmailConfig) GetHelloOk() (*string, bool)`

GetHelloOk returns a tuple with the Hello field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHello

`func (o *O11yEmailConfig) SetHello(v string)`

SetHello sets Hello field to given value.

### HasHello

`func (o *O11yEmailConfig) HasHello() bool`

HasHello returns a boolean if a field has been set.

### GetHtml

`func (o *O11yEmailConfig) GetHtml() string`

GetHtml returns the Html field if non-nil, zero value otherwise.

### GetHtmlOk

`func (o *O11yEmailConfig) GetHtmlOk() (*string, bool)`

GetHtmlOk returns a tuple with the Html field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtml

`func (o *O11yEmailConfig) SetHtml(v string)`

SetHtml sets Html field to given value.

### HasHtml

`func (o *O11yEmailConfig) HasHtml() bool`

HasHtml returns a boolean if a field has been set.

### GetRequireTls

`func (o *O11yEmailConfig) GetRequireTls() bool`

GetRequireTls returns the RequireTls field if non-nil, zero value otherwise.

### GetRequireTlsOk

`func (o *O11yEmailConfig) GetRequireTlsOk() (*bool, bool)`

GetRequireTlsOk returns a tuple with the RequireTls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireTls

`func (o *O11yEmailConfig) SetRequireTls(v bool)`

SetRequireTls sets RequireTls field to given value.

### HasRequireTls

`func (o *O11yEmailConfig) HasRequireTls() bool`

HasRequireTls returns a boolean if a field has been set.

### GetSmarthost

`func (o *O11yEmailConfig) GetSmarthost() interface{}`

GetSmarthost returns the Smarthost field if non-nil, zero value otherwise.

### GetSmarthostOk

`func (o *O11yEmailConfig) GetSmarthostOk() (*interface{}, bool)`

GetSmarthostOk returns a tuple with the Smarthost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmarthost

`func (o *O11yEmailConfig) SetSmarthost(v interface{})`

SetSmarthost sets Smarthost field to given value.

### HasSmarthost

`func (o *O11yEmailConfig) HasSmarthost() bool`

HasSmarthost returns a boolean if a field has been set.

### SetSmarthostNil

`func (o *O11yEmailConfig) SetSmarthostNil(b bool)`

 SetSmarthostNil sets the value for Smarthost to be an explicit nil

### UnsetSmarthost
`func (o *O11yEmailConfig) UnsetSmarthost()`

UnsetSmarthost ensures that no value is present for Smarthost, not even an explicit nil
### GetText

`func (o *O11yEmailConfig) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *O11yEmailConfig) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *O11yEmailConfig) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *O11yEmailConfig) HasText() bool`

HasText returns a boolean if a field has been set.

### GetThreading

`func (o *O11yEmailConfig) GetThreading() O11yThreadingConfig`

GetThreading returns the Threading field if non-nil, zero value otherwise.

### GetThreadingOk

`func (o *O11yEmailConfig) GetThreadingOk() (*O11yThreadingConfig, bool)`

GetThreadingOk returns a tuple with the Threading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreading

`func (o *O11yEmailConfig) SetThreading(v O11yThreadingConfig)`

SetThreading sets Threading field to given value.

### HasThreading

`func (o *O11yEmailConfig) HasThreading() bool`

HasThreading returns a boolean if a field has been set.

### GetTlsConfig

`func (o *O11yEmailConfig) GetTlsConfig() O11yTLSConfig`

GetTlsConfig returns the TlsConfig field if non-nil, zero value otherwise.

### GetTlsConfigOk

`func (o *O11yEmailConfig) GetTlsConfigOk() (*O11yTLSConfig, bool)`

GetTlsConfigOk returns a tuple with the TlsConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsConfig

`func (o *O11yEmailConfig) SetTlsConfig(v O11yTLSConfig)`

SetTlsConfig sets TlsConfig field to given value.

### HasTlsConfig

`func (o *O11yEmailConfig) HasTlsConfig() bool`

HasTlsConfig returns a boolean if a field has been set.

### GetTo

`func (o *O11yEmailConfig) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *O11yEmailConfig) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *O11yEmailConfig) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *O11yEmailConfig) HasTo() bool`

HasTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


