# O11yMattermostConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotifierConfig** | Pointer to [**O11yNotifierConfig**](O11yNotifierConfig.md) |  | [optional] 
**Attachments** | Pointer to [**[]O11yMattermostAttachment**](O11yMattermostAttachment.md) |  | [optional] 
**Channel** | Pointer to **string** |  | [optional] 
**HttpConfig** | Pointer to [**O11yHTTPClientConfig**](O11yHTTPClientConfig.md) |  | [optional] 
**IconEmoji** | Pointer to **string** |  | [optional] 
**IconUrl** | Pointer to **string** |  | [optional] 
**Priority** | Pointer to [**O11yMattermostPriority**](O11yMattermostPriority.md) |  | [optional] 
**Props** | Pointer to [**O11yMattermostProps**](O11yMattermostProps.md) |  | [optional] 
**Text** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**WebhookUrl** | Pointer to **interface{}** |  | [optional] 
**WebhookUrlFile** | Pointer to **string** |  | [optional] 

## Methods

### NewO11yMattermostConfig

`func NewO11yMattermostConfig() *O11yMattermostConfig`

NewO11yMattermostConfig instantiates a new O11yMattermostConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yMattermostConfigWithDefaults

`func NewO11yMattermostConfigWithDefaults() *O11yMattermostConfig`

NewO11yMattermostConfigWithDefaults instantiates a new O11yMattermostConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotifierConfig

`func (o *O11yMattermostConfig) GetNotifierConfig() O11yNotifierConfig`

GetNotifierConfig returns the NotifierConfig field if non-nil, zero value otherwise.

### GetNotifierConfigOk

`func (o *O11yMattermostConfig) GetNotifierConfigOk() (*O11yNotifierConfig, bool)`

GetNotifierConfigOk returns a tuple with the NotifierConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifierConfig

`func (o *O11yMattermostConfig) SetNotifierConfig(v O11yNotifierConfig)`

SetNotifierConfig sets NotifierConfig field to given value.

### HasNotifierConfig

`func (o *O11yMattermostConfig) HasNotifierConfig() bool`

HasNotifierConfig returns a boolean if a field has been set.

### GetAttachments

`func (o *O11yMattermostConfig) GetAttachments() []O11yMattermostAttachment`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *O11yMattermostConfig) GetAttachmentsOk() (*[]O11yMattermostAttachment, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *O11yMattermostConfig) SetAttachments(v []O11yMattermostAttachment)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *O11yMattermostConfig) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetChannel

`func (o *O11yMattermostConfig) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *O11yMattermostConfig) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *O11yMattermostConfig) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *O11yMattermostConfig) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetHttpConfig

`func (o *O11yMattermostConfig) GetHttpConfig() O11yHTTPClientConfig`

GetHttpConfig returns the HttpConfig field if non-nil, zero value otherwise.

### GetHttpConfigOk

`func (o *O11yMattermostConfig) GetHttpConfigOk() (*O11yHTTPClientConfig, bool)`

GetHttpConfigOk returns a tuple with the HttpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpConfig

`func (o *O11yMattermostConfig) SetHttpConfig(v O11yHTTPClientConfig)`

SetHttpConfig sets HttpConfig field to given value.

### HasHttpConfig

`func (o *O11yMattermostConfig) HasHttpConfig() bool`

HasHttpConfig returns a boolean if a field has been set.

### GetIconEmoji

`func (o *O11yMattermostConfig) GetIconEmoji() string`

GetIconEmoji returns the IconEmoji field if non-nil, zero value otherwise.

### GetIconEmojiOk

`func (o *O11yMattermostConfig) GetIconEmojiOk() (*string, bool)`

GetIconEmojiOk returns a tuple with the IconEmoji field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconEmoji

`func (o *O11yMattermostConfig) SetIconEmoji(v string)`

SetIconEmoji sets IconEmoji field to given value.

### HasIconEmoji

`func (o *O11yMattermostConfig) HasIconEmoji() bool`

HasIconEmoji returns a boolean if a field has been set.

### GetIconUrl

`func (o *O11yMattermostConfig) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *O11yMattermostConfig) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *O11yMattermostConfig) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.

### HasIconUrl

`func (o *O11yMattermostConfig) HasIconUrl() bool`

HasIconUrl returns a boolean if a field has been set.

### GetPriority

`func (o *O11yMattermostConfig) GetPriority() O11yMattermostPriority`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *O11yMattermostConfig) GetPriorityOk() (*O11yMattermostPriority, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *O11yMattermostConfig) SetPriority(v O11yMattermostPriority)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *O11yMattermostConfig) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetProps

`func (o *O11yMattermostConfig) GetProps() O11yMattermostProps`

GetProps returns the Props field if non-nil, zero value otherwise.

### GetPropsOk

`func (o *O11yMattermostConfig) GetPropsOk() (*O11yMattermostProps, bool)`

GetPropsOk returns a tuple with the Props field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProps

`func (o *O11yMattermostConfig) SetProps(v O11yMattermostProps)`

SetProps sets Props field to given value.

### HasProps

`func (o *O11yMattermostConfig) HasProps() bool`

HasProps returns a boolean if a field has been set.

### GetText

`func (o *O11yMattermostConfig) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *O11yMattermostConfig) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *O11yMattermostConfig) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *O11yMattermostConfig) HasText() bool`

HasText returns a boolean if a field has been set.

### GetType

`func (o *O11yMattermostConfig) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *O11yMattermostConfig) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *O11yMattermostConfig) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *O11yMattermostConfig) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUsername

`func (o *O11yMattermostConfig) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *O11yMattermostConfig) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *O11yMattermostConfig) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *O11yMattermostConfig) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetWebhookUrl

`func (o *O11yMattermostConfig) GetWebhookUrl() interface{}`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *O11yMattermostConfig) GetWebhookUrlOk() (*interface{}, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *O11yMattermostConfig) SetWebhookUrl(v interface{})`

SetWebhookUrl sets WebhookUrl field to given value.

### HasWebhookUrl

`func (o *O11yMattermostConfig) HasWebhookUrl() bool`

HasWebhookUrl returns a boolean if a field has been set.

### SetWebhookUrlNil

`func (o *O11yMattermostConfig) SetWebhookUrlNil(b bool)`

 SetWebhookUrlNil sets the value for WebhookUrl to be an explicit nil

### UnsetWebhookUrl
`func (o *O11yMattermostConfig) UnsetWebhookUrl()`

UnsetWebhookUrl ensures that no value is present for WebhookUrl, not even an explicit nil
### GetWebhookUrlFile

`func (o *O11yMattermostConfig) GetWebhookUrlFile() string`

GetWebhookUrlFile returns the WebhookUrlFile field if non-nil, zero value otherwise.

### GetWebhookUrlFileOk

`func (o *O11yMattermostConfig) GetWebhookUrlFileOk() (*string, bool)`

GetWebhookUrlFileOk returns a tuple with the WebhookUrlFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrlFile

`func (o *O11yMattermostConfig) SetWebhookUrlFile(v string)`

SetWebhookUrlFile sets WebhookUrlFile field to given value.

### HasWebhookUrlFile

`func (o *O11yMattermostConfig) HasWebhookUrlFile() bool`

HasWebhookUrlFile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


