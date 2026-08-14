# O11yMSTeamsConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotifierConfig** | Pointer to [**O11yNotifierConfig**](O11yNotifierConfig.md) |  | [optional] 
**HttpConfig** | Pointer to [**O11yHTTPClientConfig**](O11yHTTPClientConfig.md) |  | [optional] 
**Summary** | Pointer to **string** |  | [optional] 
**Text** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**WebhookUrl** | Pointer to **interface{}** |  | [optional] 
**WebhookUrlFile** | Pointer to **string** |  | [optional] 

## Methods

### NewO11yMSTeamsConfig

`func NewO11yMSTeamsConfig() *O11yMSTeamsConfig`

NewO11yMSTeamsConfig instantiates a new O11yMSTeamsConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yMSTeamsConfigWithDefaults

`func NewO11yMSTeamsConfigWithDefaults() *O11yMSTeamsConfig`

NewO11yMSTeamsConfigWithDefaults instantiates a new O11yMSTeamsConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotifierConfig

`func (o *O11yMSTeamsConfig) GetNotifierConfig() O11yNotifierConfig`

GetNotifierConfig returns the NotifierConfig field if non-nil, zero value otherwise.

### GetNotifierConfigOk

`func (o *O11yMSTeamsConfig) GetNotifierConfigOk() (*O11yNotifierConfig, bool)`

GetNotifierConfigOk returns a tuple with the NotifierConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifierConfig

`func (o *O11yMSTeamsConfig) SetNotifierConfig(v O11yNotifierConfig)`

SetNotifierConfig sets NotifierConfig field to given value.

### HasNotifierConfig

`func (o *O11yMSTeamsConfig) HasNotifierConfig() bool`

HasNotifierConfig returns a boolean if a field has been set.

### GetHttpConfig

`func (o *O11yMSTeamsConfig) GetHttpConfig() O11yHTTPClientConfig`

GetHttpConfig returns the HttpConfig field if non-nil, zero value otherwise.

### GetHttpConfigOk

`func (o *O11yMSTeamsConfig) GetHttpConfigOk() (*O11yHTTPClientConfig, bool)`

GetHttpConfigOk returns a tuple with the HttpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpConfig

`func (o *O11yMSTeamsConfig) SetHttpConfig(v O11yHTTPClientConfig)`

SetHttpConfig sets HttpConfig field to given value.

### HasHttpConfig

`func (o *O11yMSTeamsConfig) HasHttpConfig() bool`

HasHttpConfig returns a boolean if a field has been set.

### GetSummary

`func (o *O11yMSTeamsConfig) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *O11yMSTeamsConfig) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *O11yMSTeamsConfig) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *O11yMSTeamsConfig) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetText

`func (o *O11yMSTeamsConfig) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *O11yMSTeamsConfig) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *O11yMSTeamsConfig) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *O11yMSTeamsConfig) HasText() bool`

HasText returns a boolean if a field has been set.

### GetTitle

`func (o *O11yMSTeamsConfig) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *O11yMSTeamsConfig) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *O11yMSTeamsConfig) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *O11yMSTeamsConfig) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetWebhookUrl

`func (o *O11yMSTeamsConfig) GetWebhookUrl() interface{}`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *O11yMSTeamsConfig) GetWebhookUrlOk() (*interface{}, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *O11yMSTeamsConfig) SetWebhookUrl(v interface{})`

SetWebhookUrl sets WebhookUrl field to given value.

### HasWebhookUrl

`func (o *O11yMSTeamsConfig) HasWebhookUrl() bool`

HasWebhookUrl returns a boolean if a field has been set.

### SetWebhookUrlNil

`func (o *O11yMSTeamsConfig) SetWebhookUrlNil(b bool)`

 SetWebhookUrlNil sets the value for WebhookUrl to be an explicit nil

### UnsetWebhookUrl
`func (o *O11yMSTeamsConfig) UnsetWebhookUrl()`

UnsetWebhookUrl ensures that no value is present for WebhookUrl, not even an explicit nil
### GetWebhookUrlFile

`func (o *O11yMSTeamsConfig) GetWebhookUrlFile() string`

GetWebhookUrlFile returns the WebhookUrlFile field if non-nil, zero value otherwise.

### GetWebhookUrlFileOk

`func (o *O11yMSTeamsConfig) GetWebhookUrlFileOk() (*string, bool)`

GetWebhookUrlFileOk returns a tuple with the WebhookUrlFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrlFile

`func (o *O11yMSTeamsConfig) SetWebhookUrlFile(v string)`

SetWebhookUrlFile sets WebhookUrlFile field to given value.

### HasWebhookUrlFile

`func (o *O11yMSTeamsConfig) HasWebhookUrlFile() bool`

HasWebhookUrlFile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


