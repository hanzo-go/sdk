# O11yTelegramConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotifierConfig** | Pointer to [**O11yNotifierConfig**](O11yNotifierConfig.md) |  | [optional] 
**ApiUrl** | Pointer to **interface{}** |  | [optional] 
**Chat** | Pointer to **int32** |  | [optional] 
**ChatFile** | Pointer to **string** |  | [optional] 
**DisableNotifications** | Pointer to **bool** |  | [optional] 
**HttpConfig** | Pointer to [**O11yHTTPClientConfig**](O11yHTTPClientConfig.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**MessageThreadId** | Pointer to **int32** |  | [optional] 
**ParseMode** | Pointer to **string** |  | [optional] 
**Token** | Pointer to **interface{}** |  | [optional] 
**TokenFile** | Pointer to **string** |  | [optional] 

## Methods

### NewO11yTelegramConfig

`func NewO11yTelegramConfig() *O11yTelegramConfig`

NewO11yTelegramConfig instantiates a new O11yTelegramConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yTelegramConfigWithDefaults

`func NewO11yTelegramConfigWithDefaults() *O11yTelegramConfig`

NewO11yTelegramConfigWithDefaults instantiates a new O11yTelegramConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotifierConfig

`func (o *O11yTelegramConfig) GetNotifierConfig() O11yNotifierConfig`

GetNotifierConfig returns the NotifierConfig field if non-nil, zero value otherwise.

### GetNotifierConfigOk

`func (o *O11yTelegramConfig) GetNotifierConfigOk() (*O11yNotifierConfig, bool)`

GetNotifierConfigOk returns a tuple with the NotifierConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifierConfig

`func (o *O11yTelegramConfig) SetNotifierConfig(v O11yNotifierConfig)`

SetNotifierConfig sets NotifierConfig field to given value.

### HasNotifierConfig

`func (o *O11yTelegramConfig) HasNotifierConfig() bool`

HasNotifierConfig returns a boolean if a field has been set.

### GetApiUrl

`func (o *O11yTelegramConfig) GetApiUrl() interface{}`

GetApiUrl returns the ApiUrl field if non-nil, zero value otherwise.

### GetApiUrlOk

`func (o *O11yTelegramConfig) GetApiUrlOk() (*interface{}, bool)`

GetApiUrlOk returns a tuple with the ApiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiUrl

`func (o *O11yTelegramConfig) SetApiUrl(v interface{})`

SetApiUrl sets ApiUrl field to given value.

### HasApiUrl

`func (o *O11yTelegramConfig) HasApiUrl() bool`

HasApiUrl returns a boolean if a field has been set.

### SetApiUrlNil

`func (o *O11yTelegramConfig) SetApiUrlNil(b bool)`

 SetApiUrlNil sets the value for ApiUrl to be an explicit nil

### UnsetApiUrl
`func (o *O11yTelegramConfig) UnsetApiUrl()`

UnsetApiUrl ensures that no value is present for ApiUrl, not even an explicit nil
### GetChat

`func (o *O11yTelegramConfig) GetChat() int32`

GetChat returns the Chat field if non-nil, zero value otherwise.

### GetChatOk

`func (o *O11yTelegramConfig) GetChatOk() (*int32, bool)`

GetChatOk returns a tuple with the Chat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChat

`func (o *O11yTelegramConfig) SetChat(v int32)`

SetChat sets Chat field to given value.

### HasChat

`func (o *O11yTelegramConfig) HasChat() bool`

HasChat returns a boolean if a field has been set.

### GetChatFile

`func (o *O11yTelegramConfig) GetChatFile() string`

GetChatFile returns the ChatFile field if non-nil, zero value otherwise.

### GetChatFileOk

`func (o *O11yTelegramConfig) GetChatFileOk() (*string, bool)`

GetChatFileOk returns a tuple with the ChatFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatFile

`func (o *O11yTelegramConfig) SetChatFile(v string)`

SetChatFile sets ChatFile field to given value.

### HasChatFile

`func (o *O11yTelegramConfig) HasChatFile() bool`

HasChatFile returns a boolean if a field has been set.

### GetDisableNotifications

`func (o *O11yTelegramConfig) GetDisableNotifications() bool`

GetDisableNotifications returns the DisableNotifications field if non-nil, zero value otherwise.

### GetDisableNotificationsOk

`func (o *O11yTelegramConfig) GetDisableNotificationsOk() (*bool, bool)`

GetDisableNotificationsOk returns a tuple with the DisableNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableNotifications

`func (o *O11yTelegramConfig) SetDisableNotifications(v bool)`

SetDisableNotifications sets DisableNotifications field to given value.

### HasDisableNotifications

`func (o *O11yTelegramConfig) HasDisableNotifications() bool`

HasDisableNotifications returns a boolean if a field has been set.

### GetHttpConfig

`func (o *O11yTelegramConfig) GetHttpConfig() O11yHTTPClientConfig`

GetHttpConfig returns the HttpConfig field if non-nil, zero value otherwise.

### GetHttpConfigOk

`func (o *O11yTelegramConfig) GetHttpConfigOk() (*O11yHTTPClientConfig, bool)`

GetHttpConfigOk returns a tuple with the HttpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpConfig

`func (o *O11yTelegramConfig) SetHttpConfig(v O11yHTTPClientConfig)`

SetHttpConfig sets HttpConfig field to given value.

### HasHttpConfig

`func (o *O11yTelegramConfig) HasHttpConfig() bool`

HasHttpConfig returns a boolean if a field has been set.

### GetMessage

`func (o *O11yTelegramConfig) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *O11yTelegramConfig) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *O11yTelegramConfig) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *O11yTelegramConfig) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetMessageThreadId

`func (o *O11yTelegramConfig) GetMessageThreadId() int32`

GetMessageThreadId returns the MessageThreadId field if non-nil, zero value otherwise.

### GetMessageThreadIdOk

`func (o *O11yTelegramConfig) GetMessageThreadIdOk() (*int32, bool)`

GetMessageThreadIdOk returns a tuple with the MessageThreadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageThreadId

`func (o *O11yTelegramConfig) SetMessageThreadId(v int32)`

SetMessageThreadId sets MessageThreadId field to given value.

### HasMessageThreadId

`func (o *O11yTelegramConfig) HasMessageThreadId() bool`

HasMessageThreadId returns a boolean if a field has been set.

### GetParseMode

`func (o *O11yTelegramConfig) GetParseMode() string`

GetParseMode returns the ParseMode field if non-nil, zero value otherwise.

### GetParseModeOk

`func (o *O11yTelegramConfig) GetParseModeOk() (*string, bool)`

GetParseModeOk returns a tuple with the ParseMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParseMode

`func (o *O11yTelegramConfig) SetParseMode(v string)`

SetParseMode sets ParseMode field to given value.

### HasParseMode

`func (o *O11yTelegramConfig) HasParseMode() bool`

HasParseMode returns a boolean if a field has been set.

### GetToken

`func (o *O11yTelegramConfig) GetToken() interface{}`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *O11yTelegramConfig) GetTokenOk() (*interface{}, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *O11yTelegramConfig) SetToken(v interface{})`

SetToken sets Token field to given value.

### HasToken

`func (o *O11yTelegramConfig) HasToken() bool`

HasToken returns a boolean if a field has been set.

### SetTokenNil

`func (o *O11yTelegramConfig) SetTokenNil(b bool)`

 SetTokenNil sets the value for Token to be an explicit nil

### UnsetToken
`func (o *O11yTelegramConfig) UnsetToken()`

UnsetToken ensures that no value is present for Token, not even an explicit nil
### GetTokenFile

`func (o *O11yTelegramConfig) GetTokenFile() string`

GetTokenFile returns the TokenFile field if non-nil, zero value otherwise.

### GetTokenFileOk

`func (o *O11yTelegramConfig) GetTokenFileOk() (*string, bool)`

GetTokenFileOk returns a tuple with the TokenFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenFile

`func (o *O11yTelegramConfig) SetTokenFile(v string)`

SetTokenFile sets TokenFile field to given value.

### HasTokenFile

`func (o *O11yTelegramConfig) HasTokenFile() bool`

HasTokenFile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


