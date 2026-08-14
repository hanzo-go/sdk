# O11yDiscordConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotifierConfig** | Pointer to [**O11yNotifierConfig**](O11yNotifierConfig.md) |  | [optional] 
**AvatarUrl** | Pointer to **string** |  | [optional] 
**Content** | Pointer to **string** |  | [optional] 
**HttpConfig** | Pointer to [**O11yHTTPClientConfig**](O11yHTTPClientConfig.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**WebhookUrl** | Pointer to **interface{}** |  | [optional] 
**WebhookUrlFile** | Pointer to **string** |  | [optional] 

## Methods

### NewO11yDiscordConfig

`func NewO11yDiscordConfig() *O11yDiscordConfig`

NewO11yDiscordConfig instantiates a new O11yDiscordConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yDiscordConfigWithDefaults

`func NewO11yDiscordConfigWithDefaults() *O11yDiscordConfig`

NewO11yDiscordConfigWithDefaults instantiates a new O11yDiscordConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotifierConfig

`func (o *O11yDiscordConfig) GetNotifierConfig() O11yNotifierConfig`

GetNotifierConfig returns the NotifierConfig field if non-nil, zero value otherwise.

### GetNotifierConfigOk

`func (o *O11yDiscordConfig) GetNotifierConfigOk() (*O11yNotifierConfig, bool)`

GetNotifierConfigOk returns a tuple with the NotifierConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifierConfig

`func (o *O11yDiscordConfig) SetNotifierConfig(v O11yNotifierConfig)`

SetNotifierConfig sets NotifierConfig field to given value.

### HasNotifierConfig

`func (o *O11yDiscordConfig) HasNotifierConfig() bool`

HasNotifierConfig returns a boolean if a field has been set.

### GetAvatarUrl

`func (o *O11yDiscordConfig) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *O11yDiscordConfig) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *O11yDiscordConfig) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.

### HasAvatarUrl

`func (o *O11yDiscordConfig) HasAvatarUrl() bool`

HasAvatarUrl returns a boolean if a field has been set.

### GetContent

`func (o *O11yDiscordConfig) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *O11yDiscordConfig) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *O11yDiscordConfig) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *O11yDiscordConfig) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetHttpConfig

`func (o *O11yDiscordConfig) GetHttpConfig() O11yHTTPClientConfig`

GetHttpConfig returns the HttpConfig field if non-nil, zero value otherwise.

### GetHttpConfigOk

`func (o *O11yDiscordConfig) GetHttpConfigOk() (*O11yHTTPClientConfig, bool)`

GetHttpConfigOk returns a tuple with the HttpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpConfig

`func (o *O11yDiscordConfig) SetHttpConfig(v O11yHTTPClientConfig)`

SetHttpConfig sets HttpConfig field to given value.

### HasHttpConfig

`func (o *O11yDiscordConfig) HasHttpConfig() bool`

HasHttpConfig returns a boolean if a field has been set.

### GetMessage

`func (o *O11yDiscordConfig) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *O11yDiscordConfig) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *O11yDiscordConfig) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *O11yDiscordConfig) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetTitle

`func (o *O11yDiscordConfig) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *O11yDiscordConfig) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *O11yDiscordConfig) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *O11yDiscordConfig) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUsername

`func (o *O11yDiscordConfig) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *O11yDiscordConfig) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *O11yDiscordConfig) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *O11yDiscordConfig) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetWebhookUrl

`func (o *O11yDiscordConfig) GetWebhookUrl() interface{}`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *O11yDiscordConfig) GetWebhookUrlOk() (*interface{}, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *O11yDiscordConfig) SetWebhookUrl(v interface{})`

SetWebhookUrl sets WebhookUrl field to given value.

### HasWebhookUrl

`func (o *O11yDiscordConfig) HasWebhookUrl() bool`

HasWebhookUrl returns a boolean if a field has been set.

### SetWebhookUrlNil

`func (o *O11yDiscordConfig) SetWebhookUrlNil(b bool)`

 SetWebhookUrlNil sets the value for WebhookUrl to be an explicit nil

### UnsetWebhookUrl
`func (o *O11yDiscordConfig) UnsetWebhookUrl()`

UnsetWebhookUrl ensures that no value is present for WebhookUrl, not even an explicit nil
### GetWebhookUrlFile

`func (o *O11yDiscordConfig) GetWebhookUrlFile() string`

GetWebhookUrlFile returns the WebhookUrlFile field if non-nil, zero value otherwise.

### GetWebhookUrlFileOk

`func (o *O11yDiscordConfig) GetWebhookUrlFileOk() (*string, bool)`

GetWebhookUrlFileOk returns a tuple with the WebhookUrlFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrlFile

`func (o *O11yDiscordConfig) SetWebhookUrlFile(v string)`

SetWebhookUrlFile sets WebhookUrlFile field to given value.

### HasWebhookUrlFile

`func (o *O11yDiscordConfig) HasWebhookUrlFile() bool`

HasWebhookUrlFile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


