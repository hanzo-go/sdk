# O11yMSTeamsV2Config

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotifierConfig** | Pointer to [**O11yNotifierConfig**](O11yNotifierConfig.md) |  | [optional] 
**HttpConfig** | Pointer to [**O11yHTTPClientConfig**](O11yHTTPClientConfig.md) |  | [optional] 
**Text** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**WebhookUrl** | Pointer to **interface{}** |  | [optional] 
**WebhookUrlFile** | Pointer to **string** |  | [optional] 

## Methods

### NewO11yMSTeamsV2Config

`func NewO11yMSTeamsV2Config() *O11yMSTeamsV2Config`

NewO11yMSTeamsV2Config instantiates a new O11yMSTeamsV2Config object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yMSTeamsV2ConfigWithDefaults

`func NewO11yMSTeamsV2ConfigWithDefaults() *O11yMSTeamsV2Config`

NewO11yMSTeamsV2ConfigWithDefaults instantiates a new O11yMSTeamsV2Config object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotifierConfig

`func (o *O11yMSTeamsV2Config) GetNotifierConfig() O11yNotifierConfig`

GetNotifierConfig returns the NotifierConfig field if non-nil, zero value otherwise.

### GetNotifierConfigOk

`func (o *O11yMSTeamsV2Config) GetNotifierConfigOk() (*O11yNotifierConfig, bool)`

GetNotifierConfigOk returns a tuple with the NotifierConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifierConfig

`func (o *O11yMSTeamsV2Config) SetNotifierConfig(v O11yNotifierConfig)`

SetNotifierConfig sets NotifierConfig field to given value.

### HasNotifierConfig

`func (o *O11yMSTeamsV2Config) HasNotifierConfig() bool`

HasNotifierConfig returns a boolean if a field has been set.

### GetHttpConfig

`func (o *O11yMSTeamsV2Config) GetHttpConfig() O11yHTTPClientConfig`

GetHttpConfig returns the HttpConfig field if non-nil, zero value otherwise.

### GetHttpConfigOk

`func (o *O11yMSTeamsV2Config) GetHttpConfigOk() (*O11yHTTPClientConfig, bool)`

GetHttpConfigOk returns a tuple with the HttpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpConfig

`func (o *O11yMSTeamsV2Config) SetHttpConfig(v O11yHTTPClientConfig)`

SetHttpConfig sets HttpConfig field to given value.

### HasHttpConfig

`func (o *O11yMSTeamsV2Config) HasHttpConfig() bool`

HasHttpConfig returns a boolean if a field has been set.

### GetText

`func (o *O11yMSTeamsV2Config) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *O11yMSTeamsV2Config) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *O11yMSTeamsV2Config) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *O11yMSTeamsV2Config) HasText() bool`

HasText returns a boolean if a field has been set.

### GetTitle

`func (o *O11yMSTeamsV2Config) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *O11yMSTeamsV2Config) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *O11yMSTeamsV2Config) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *O11yMSTeamsV2Config) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetWebhookUrl

`func (o *O11yMSTeamsV2Config) GetWebhookUrl() interface{}`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *O11yMSTeamsV2Config) GetWebhookUrlOk() (*interface{}, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *O11yMSTeamsV2Config) SetWebhookUrl(v interface{})`

SetWebhookUrl sets WebhookUrl field to given value.

### HasWebhookUrl

`func (o *O11yMSTeamsV2Config) HasWebhookUrl() bool`

HasWebhookUrl returns a boolean if a field has been set.

### SetWebhookUrlNil

`func (o *O11yMSTeamsV2Config) SetWebhookUrlNil(b bool)`

 SetWebhookUrlNil sets the value for WebhookUrl to be an explicit nil

### UnsetWebhookUrl
`func (o *O11yMSTeamsV2Config) UnsetWebhookUrl()`

UnsetWebhookUrl ensures that no value is present for WebhookUrl, not even an explicit nil
### GetWebhookUrlFile

`func (o *O11yMSTeamsV2Config) GetWebhookUrlFile() string`

GetWebhookUrlFile returns the WebhookUrlFile field if non-nil, zero value otherwise.

### GetWebhookUrlFileOk

`func (o *O11yMSTeamsV2Config) GetWebhookUrlFileOk() (*string, bool)`

GetWebhookUrlFileOk returns a tuple with the WebhookUrlFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrlFile

`func (o *O11yMSTeamsV2Config) SetWebhookUrlFile(v string)`

SetWebhookUrlFile sets WebhookUrlFile field to given value.

### HasWebhookUrlFile

`func (o *O11yMSTeamsV2Config) HasWebhookUrlFile() bool`

HasWebhookUrlFile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


