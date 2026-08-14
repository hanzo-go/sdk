# O11yPushoverConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotifierConfig** | Pointer to [**O11yNotifierConfig**](O11yNotifierConfig.md) |  | [optional] 
**Device** | Pointer to **string** |  | [optional] 
**Expire** | Pointer to **int32** |  | [optional] 
**Html** | Pointer to **bool** |  | [optional] 
**HttpConfig** | Pointer to [**O11yHTTPClientConfig**](O11yHTTPClientConfig.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Monospace** | Pointer to **bool** |  | [optional] 
**Priority** | Pointer to **string** |  | [optional] 
**Retry** | Pointer to **int32** |  | [optional] 
**Sound** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Token** | Pointer to **interface{}** |  | [optional] 
**TokenFile** | Pointer to **string** |  | [optional] 
**Ttl** | Pointer to **int32** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**UrlTitle** | Pointer to **string** |  | [optional] 
**UserKey** | Pointer to **interface{}** |  | [optional] 
**UserKeyFile** | Pointer to **string** |  | [optional] 

## Methods

### NewO11yPushoverConfig

`func NewO11yPushoverConfig() *O11yPushoverConfig`

NewO11yPushoverConfig instantiates a new O11yPushoverConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yPushoverConfigWithDefaults

`func NewO11yPushoverConfigWithDefaults() *O11yPushoverConfig`

NewO11yPushoverConfigWithDefaults instantiates a new O11yPushoverConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotifierConfig

`func (o *O11yPushoverConfig) GetNotifierConfig() O11yNotifierConfig`

GetNotifierConfig returns the NotifierConfig field if non-nil, zero value otherwise.

### GetNotifierConfigOk

`func (o *O11yPushoverConfig) GetNotifierConfigOk() (*O11yNotifierConfig, bool)`

GetNotifierConfigOk returns a tuple with the NotifierConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifierConfig

`func (o *O11yPushoverConfig) SetNotifierConfig(v O11yNotifierConfig)`

SetNotifierConfig sets NotifierConfig field to given value.

### HasNotifierConfig

`func (o *O11yPushoverConfig) HasNotifierConfig() bool`

HasNotifierConfig returns a boolean if a field has been set.

### GetDevice

`func (o *O11yPushoverConfig) GetDevice() string`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *O11yPushoverConfig) GetDeviceOk() (*string, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *O11yPushoverConfig) SetDevice(v string)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *O11yPushoverConfig) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetExpire

`func (o *O11yPushoverConfig) GetExpire() int32`

GetExpire returns the Expire field if non-nil, zero value otherwise.

### GetExpireOk

`func (o *O11yPushoverConfig) GetExpireOk() (*int32, bool)`

GetExpireOk returns a tuple with the Expire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpire

`func (o *O11yPushoverConfig) SetExpire(v int32)`

SetExpire sets Expire field to given value.

### HasExpire

`func (o *O11yPushoverConfig) HasExpire() bool`

HasExpire returns a boolean if a field has been set.

### GetHtml

`func (o *O11yPushoverConfig) GetHtml() bool`

GetHtml returns the Html field if non-nil, zero value otherwise.

### GetHtmlOk

`func (o *O11yPushoverConfig) GetHtmlOk() (*bool, bool)`

GetHtmlOk returns a tuple with the Html field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtml

`func (o *O11yPushoverConfig) SetHtml(v bool)`

SetHtml sets Html field to given value.

### HasHtml

`func (o *O11yPushoverConfig) HasHtml() bool`

HasHtml returns a boolean if a field has been set.

### GetHttpConfig

`func (o *O11yPushoverConfig) GetHttpConfig() O11yHTTPClientConfig`

GetHttpConfig returns the HttpConfig field if non-nil, zero value otherwise.

### GetHttpConfigOk

`func (o *O11yPushoverConfig) GetHttpConfigOk() (*O11yHTTPClientConfig, bool)`

GetHttpConfigOk returns a tuple with the HttpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpConfig

`func (o *O11yPushoverConfig) SetHttpConfig(v O11yHTTPClientConfig)`

SetHttpConfig sets HttpConfig field to given value.

### HasHttpConfig

`func (o *O11yPushoverConfig) HasHttpConfig() bool`

HasHttpConfig returns a boolean if a field has been set.

### GetMessage

`func (o *O11yPushoverConfig) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *O11yPushoverConfig) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *O11yPushoverConfig) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *O11yPushoverConfig) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetMonospace

`func (o *O11yPushoverConfig) GetMonospace() bool`

GetMonospace returns the Monospace field if non-nil, zero value otherwise.

### GetMonospaceOk

`func (o *O11yPushoverConfig) GetMonospaceOk() (*bool, bool)`

GetMonospaceOk returns a tuple with the Monospace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonospace

`func (o *O11yPushoverConfig) SetMonospace(v bool)`

SetMonospace sets Monospace field to given value.

### HasMonospace

`func (o *O11yPushoverConfig) HasMonospace() bool`

HasMonospace returns a boolean if a field has been set.

### GetPriority

`func (o *O11yPushoverConfig) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *O11yPushoverConfig) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *O11yPushoverConfig) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *O11yPushoverConfig) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetRetry

`func (o *O11yPushoverConfig) GetRetry() int32`

GetRetry returns the Retry field if non-nil, zero value otherwise.

### GetRetryOk

`func (o *O11yPushoverConfig) GetRetryOk() (*int32, bool)`

GetRetryOk returns a tuple with the Retry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetry

`func (o *O11yPushoverConfig) SetRetry(v int32)`

SetRetry sets Retry field to given value.

### HasRetry

`func (o *O11yPushoverConfig) HasRetry() bool`

HasRetry returns a boolean if a field has been set.

### GetSound

`func (o *O11yPushoverConfig) GetSound() string`

GetSound returns the Sound field if non-nil, zero value otherwise.

### GetSoundOk

`func (o *O11yPushoverConfig) GetSoundOk() (*string, bool)`

GetSoundOk returns a tuple with the Sound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSound

`func (o *O11yPushoverConfig) SetSound(v string)`

SetSound sets Sound field to given value.

### HasSound

`func (o *O11yPushoverConfig) HasSound() bool`

HasSound returns a boolean if a field has been set.

### GetTitle

`func (o *O11yPushoverConfig) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *O11yPushoverConfig) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *O11yPushoverConfig) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *O11yPushoverConfig) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetToken

`func (o *O11yPushoverConfig) GetToken() interface{}`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *O11yPushoverConfig) GetTokenOk() (*interface{}, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *O11yPushoverConfig) SetToken(v interface{})`

SetToken sets Token field to given value.

### HasToken

`func (o *O11yPushoverConfig) HasToken() bool`

HasToken returns a boolean if a field has been set.

### SetTokenNil

`func (o *O11yPushoverConfig) SetTokenNil(b bool)`

 SetTokenNil sets the value for Token to be an explicit nil

### UnsetToken
`func (o *O11yPushoverConfig) UnsetToken()`

UnsetToken ensures that no value is present for Token, not even an explicit nil
### GetTokenFile

`func (o *O11yPushoverConfig) GetTokenFile() string`

GetTokenFile returns the TokenFile field if non-nil, zero value otherwise.

### GetTokenFileOk

`func (o *O11yPushoverConfig) GetTokenFileOk() (*string, bool)`

GetTokenFileOk returns a tuple with the TokenFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenFile

`func (o *O11yPushoverConfig) SetTokenFile(v string)`

SetTokenFile sets TokenFile field to given value.

### HasTokenFile

`func (o *O11yPushoverConfig) HasTokenFile() bool`

HasTokenFile returns a boolean if a field has been set.

### GetTtl

`func (o *O11yPushoverConfig) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *O11yPushoverConfig) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *O11yPushoverConfig) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *O11yPushoverConfig) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetUrl

`func (o *O11yPushoverConfig) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *O11yPushoverConfig) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *O11yPushoverConfig) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *O11yPushoverConfig) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetUrlTitle

`func (o *O11yPushoverConfig) GetUrlTitle() string`

GetUrlTitle returns the UrlTitle field if non-nil, zero value otherwise.

### GetUrlTitleOk

`func (o *O11yPushoverConfig) GetUrlTitleOk() (*string, bool)`

GetUrlTitleOk returns a tuple with the UrlTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlTitle

`func (o *O11yPushoverConfig) SetUrlTitle(v string)`

SetUrlTitle sets UrlTitle field to given value.

### HasUrlTitle

`func (o *O11yPushoverConfig) HasUrlTitle() bool`

HasUrlTitle returns a boolean if a field has been set.

### GetUserKey

`func (o *O11yPushoverConfig) GetUserKey() interface{}`

GetUserKey returns the UserKey field if non-nil, zero value otherwise.

### GetUserKeyOk

`func (o *O11yPushoverConfig) GetUserKeyOk() (*interface{}, bool)`

GetUserKeyOk returns a tuple with the UserKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserKey

`func (o *O11yPushoverConfig) SetUserKey(v interface{})`

SetUserKey sets UserKey field to given value.

### HasUserKey

`func (o *O11yPushoverConfig) HasUserKey() bool`

HasUserKey returns a boolean if a field has been set.

### SetUserKeyNil

`func (o *O11yPushoverConfig) SetUserKeyNil(b bool)`

 SetUserKeyNil sets the value for UserKey to be an explicit nil

### UnsetUserKey
`func (o *O11yPushoverConfig) UnsetUserKey()`

UnsetUserKey ensures that no value is present for UserKey, not even an explicit nil
### GetUserKeyFile

`func (o *O11yPushoverConfig) GetUserKeyFile() string`

GetUserKeyFile returns the UserKeyFile field if non-nil, zero value otherwise.

### GetUserKeyFileOk

`func (o *O11yPushoverConfig) GetUserKeyFileOk() (*string, bool)`

GetUserKeyFileOk returns a tuple with the UserKeyFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserKeyFile

`func (o *O11yPushoverConfig) SetUserKeyFile(v string)`

SetUserKeyFile sets UserKeyFile field to given value.

### HasUserKeyFile

`func (o *O11yPushoverConfig) HasUserKeyFile() bool`

HasUserKeyFile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


