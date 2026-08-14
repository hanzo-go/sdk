# O11yGoogleChatReceiverConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotifierConfig** | Pointer to [**O11yNotifierConfig**](O11yNotifierConfig.md) |  | [optional] 
**HttpConfig** | Pointer to [**O11yHTTPClientConfig**](O11yHTTPClientConfig.md) |  | [optional] 
**Text** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**WebhookUrl** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewO11yGoogleChatReceiverConfig

`func NewO11yGoogleChatReceiverConfig() *O11yGoogleChatReceiverConfig`

NewO11yGoogleChatReceiverConfig instantiates a new O11yGoogleChatReceiverConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yGoogleChatReceiverConfigWithDefaults

`func NewO11yGoogleChatReceiverConfigWithDefaults() *O11yGoogleChatReceiverConfig`

NewO11yGoogleChatReceiverConfigWithDefaults instantiates a new O11yGoogleChatReceiverConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotifierConfig

`func (o *O11yGoogleChatReceiverConfig) GetNotifierConfig() O11yNotifierConfig`

GetNotifierConfig returns the NotifierConfig field if non-nil, zero value otherwise.

### GetNotifierConfigOk

`func (o *O11yGoogleChatReceiverConfig) GetNotifierConfigOk() (*O11yNotifierConfig, bool)`

GetNotifierConfigOk returns a tuple with the NotifierConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifierConfig

`func (o *O11yGoogleChatReceiverConfig) SetNotifierConfig(v O11yNotifierConfig)`

SetNotifierConfig sets NotifierConfig field to given value.

### HasNotifierConfig

`func (o *O11yGoogleChatReceiverConfig) HasNotifierConfig() bool`

HasNotifierConfig returns a boolean if a field has been set.

### GetHttpConfig

`func (o *O11yGoogleChatReceiverConfig) GetHttpConfig() O11yHTTPClientConfig`

GetHttpConfig returns the HttpConfig field if non-nil, zero value otherwise.

### GetHttpConfigOk

`func (o *O11yGoogleChatReceiverConfig) GetHttpConfigOk() (*O11yHTTPClientConfig, bool)`

GetHttpConfigOk returns a tuple with the HttpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpConfig

`func (o *O11yGoogleChatReceiverConfig) SetHttpConfig(v O11yHTTPClientConfig)`

SetHttpConfig sets HttpConfig field to given value.

### HasHttpConfig

`func (o *O11yGoogleChatReceiverConfig) HasHttpConfig() bool`

HasHttpConfig returns a boolean if a field has been set.

### GetText

`func (o *O11yGoogleChatReceiverConfig) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *O11yGoogleChatReceiverConfig) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *O11yGoogleChatReceiverConfig) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *O11yGoogleChatReceiverConfig) HasText() bool`

HasText returns a boolean if a field has been set.

### GetTitle

`func (o *O11yGoogleChatReceiverConfig) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *O11yGoogleChatReceiverConfig) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *O11yGoogleChatReceiverConfig) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *O11yGoogleChatReceiverConfig) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetWebhookUrl

`func (o *O11yGoogleChatReceiverConfig) GetWebhookUrl() interface{}`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *O11yGoogleChatReceiverConfig) GetWebhookUrlOk() (*interface{}, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *O11yGoogleChatReceiverConfig) SetWebhookUrl(v interface{})`

SetWebhookUrl sets WebhookUrl field to given value.

### HasWebhookUrl

`func (o *O11yGoogleChatReceiverConfig) HasWebhookUrl() bool`

HasWebhookUrl returns a boolean if a field has been set.

### SetWebhookUrlNil

`func (o *O11yGoogleChatReceiverConfig) SetWebhookUrlNil(b bool)`

 SetWebhookUrlNil sets the value for WebhookUrl to be an explicit nil

### UnsetWebhookUrl
`func (o *O11yGoogleChatReceiverConfig) UnsetWebhookUrl()`

UnsetWebhookUrl ensures that no value is present for WebhookUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


