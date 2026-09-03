# O11ySlackConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotifierConfig** | Pointer to [**O11yNotifierConfig**](O11yNotifierConfig.md) |  | [optional] 
**Actions** | Pointer to [**[]O11ySlackAction**](O11ySlackAction.md) |  | [optional] 
**ApiUrl** | Pointer to **interface{}** |  | [optional] 
**ApiUrlFile** | Pointer to **string** |  | [optional] 
**AppToken** | Pointer to **interface{}** |  | [optional] 
**AppTokenFile** | Pointer to **string** |  | [optional] 
**AppUrl** | Pointer to **interface{}** |  | [optional] 
**CallbackId** | Pointer to **string** |  | [optional] 
**Channel** | Pointer to **string** | Slack channel override, (like #other-channel or @username). | [optional] 
**Color** | Pointer to **string** |  | [optional] 
**Fallback** | Pointer to **string** |  | [optional] 
**Fields** | Pointer to [**[]O11ySlackField**](O11ySlackField.md) |  | [optional] 
**Footer** | Pointer to **string** |  | [optional] 
**HttpConfig** | Pointer to [**O11yHTTPClientConfig**](O11yHTTPClientConfig.md) |  | [optional] 
**IconEmoji** | Pointer to **string** |  | [optional] 
**IconUrl** | Pointer to **string** |  | [optional] 
**ImageUrl** | Pointer to **string** |  | [optional] 
**LinkNames** | Pointer to **bool** |  | [optional] 
**MessageText** | Pointer to **string** |  | [optional] 
**MrkdwnIn** | Pointer to **[]string** |  | [optional] 
**Pretext** | Pointer to **string** |  | [optional] 
**ShortFields** | Pointer to **bool** |  | [optional] 
**Text** | Pointer to **string** |  | [optional] 
**ThumbUrl** | Pointer to **string** |  | [optional] 
**Timeout** | Pointer to **int64** | Timeout is the maximum time allowed to invoke the slack. Setting this to 0 does not impose a timeout. | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**TitleLink** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 

## Methods

### NewO11ySlackConfig

`func NewO11ySlackConfig() *O11ySlackConfig`

NewO11ySlackConfig instantiates a new O11ySlackConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11ySlackConfigWithDefaults

`func NewO11ySlackConfigWithDefaults() *O11ySlackConfig`

NewO11ySlackConfigWithDefaults instantiates a new O11ySlackConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotifierConfig

`func (o *O11ySlackConfig) GetNotifierConfig() O11yNotifierConfig`

GetNotifierConfig returns the NotifierConfig field if non-nil, zero value otherwise.

### GetNotifierConfigOk

`func (o *O11ySlackConfig) GetNotifierConfigOk() (*O11yNotifierConfig, bool)`

GetNotifierConfigOk returns a tuple with the NotifierConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifierConfig

`func (o *O11ySlackConfig) SetNotifierConfig(v O11yNotifierConfig)`

SetNotifierConfig sets NotifierConfig field to given value.

### HasNotifierConfig

`func (o *O11ySlackConfig) HasNotifierConfig() bool`

HasNotifierConfig returns a boolean if a field has been set.

### GetActions

`func (o *O11ySlackConfig) GetActions() []O11ySlackAction`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *O11ySlackConfig) GetActionsOk() (*[]O11ySlackAction, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *O11ySlackConfig) SetActions(v []O11ySlackAction)`

SetActions sets Actions field to given value.

### HasActions

`func (o *O11ySlackConfig) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetApiUrl

`func (o *O11ySlackConfig) GetApiUrl() interface{}`

GetApiUrl returns the ApiUrl field if non-nil, zero value otherwise.

### GetApiUrlOk

`func (o *O11ySlackConfig) GetApiUrlOk() (*interface{}, bool)`

GetApiUrlOk returns a tuple with the ApiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiUrl

`func (o *O11ySlackConfig) SetApiUrl(v interface{})`

SetApiUrl sets ApiUrl field to given value.

### HasApiUrl

`func (o *O11ySlackConfig) HasApiUrl() bool`

HasApiUrl returns a boolean if a field has been set.

### SetApiUrlNil

`func (o *O11ySlackConfig) SetApiUrlNil(b bool)`

 SetApiUrlNil sets the value for ApiUrl to be an explicit nil

### UnsetApiUrl
`func (o *O11ySlackConfig) UnsetApiUrl()`

UnsetApiUrl ensures that no value is present for ApiUrl, not even an explicit nil
### GetApiUrlFile

`func (o *O11ySlackConfig) GetApiUrlFile() string`

GetApiUrlFile returns the ApiUrlFile field if non-nil, zero value otherwise.

### GetApiUrlFileOk

`func (o *O11ySlackConfig) GetApiUrlFileOk() (*string, bool)`

GetApiUrlFileOk returns a tuple with the ApiUrlFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiUrlFile

`func (o *O11ySlackConfig) SetApiUrlFile(v string)`

SetApiUrlFile sets ApiUrlFile field to given value.

### HasApiUrlFile

`func (o *O11ySlackConfig) HasApiUrlFile() bool`

HasApiUrlFile returns a boolean if a field has been set.

### GetAppToken

`func (o *O11ySlackConfig) GetAppToken() interface{}`

GetAppToken returns the AppToken field if non-nil, zero value otherwise.

### GetAppTokenOk

`func (o *O11ySlackConfig) GetAppTokenOk() (*interface{}, bool)`

GetAppTokenOk returns a tuple with the AppToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppToken

`func (o *O11ySlackConfig) SetAppToken(v interface{})`

SetAppToken sets AppToken field to given value.

### HasAppToken

`func (o *O11ySlackConfig) HasAppToken() bool`

HasAppToken returns a boolean if a field has been set.

### SetAppTokenNil

`func (o *O11ySlackConfig) SetAppTokenNil(b bool)`

 SetAppTokenNil sets the value for AppToken to be an explicit nil

### UnsetAppToken
`func (o *O11ySlackConfig) UnsetAppToken()`

UnsetAppToken ensures that no value is present for AppToken, not even an explicit nil
### GetAppTokenFile

`func (o *O11ySlackConfig) GetAppTokenFile() string`

GetAppTokenFile returns the AppTokenFile field if non-nil, zero value otherwise.

### GetAppTokenFileOk

`func (o *O11ySlackConfig) GetAppTokenFileOk() (*string, bool)`

GetAppTokenFileOk returns a tuple with the AppTokenFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppTokenFile

`func (o *O11ySlackConfig) SetAppTokenFile(v string)`

SetAppTokenFile sets AppTokenFile field to given value.

### HasAppTokenFile

`func (o *O11ySlackConfig) HasAppTokenFile() bool`

HasAppTokenFile returns a boolean if a field has been set.

### GetAppUrl

`func (o *O11ySlackConfig) GetAppUrl() interface{}`

GetAppUrl returns the AppUrl field if non-nil, zero value otherwise.

### GetAppUrlOk

`func (o *O11ySlackConfig) GetAppUrlOk() (*interface{}, bool)`

GetAppUrlOk returns a tuple with the AppUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppUrl

`func (o *O11ySlackConfig) SetAppUrl(v interface{})`

SetAppUrl sets AppUrl field to given value.

### HasAppUrl

`func (o *O11ySlackConfig) HasAppUrl() bool`

HasAppUrl returns a boolean if a field has been set.

### SetAppUrlNil

`func (o *O11ySlackConfig) SetAppUrlNil(b bool)`

 SetAppUrlNil sets the value for AppUrl to be an explicit nil

### UnsetAppUrl
`func (o *O11ySlackConfig) UnsetAppUrl()`

UnsetAppUrl ensures that no value is present for AppUrl, not even an explicit nil
### GetCallbackId

`func (o *O11ySlackConfig) GetCallbackId() string`

GetCallbackId returns the CallbackId field if non-nil, zero value otherwise.

### GetCallbackIdOk

`func (o *O11ySlackConfig) GetCallbackIdOk() (*string, bool)`

GetCallbackIdOk returns a tuple with the CallbackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackId

`func (o *O11ySlackConfig) SetCallbackId(v string)`

SetCallbackId sets CallbackId field to given value.

### HasCallbackId

`func (o *O11ySlackConfig) HasCallbackId() bool`

HasCallbackId returns a boolean if a field has been set.

### GetChannel

`func (o *O11ySlackConfig) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *O11ySlackConfig) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *O11ySlackConfig) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *O11ySlackConfig) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetColor

`func (o *O11ySlackConfig) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *O11ySlackConfig) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *O11ySlackConfig) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *O11ySlackConfig) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetFallback

`func (o *O11ySlackConfig) GetFallback() string`

GetFallback returns the Fallback field if non-nil, zero value otherwise.

### GetFallbackOk

`func (o *O11ySlackConfig) GetFallbackOk() (*string, bool)`

GetFallbackOk returns a tuple with the Fallback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallback

`func (o *O11ySlackConfig) SetFallback(v string)`

SetFallback sets Fallback field to given value.

### HasFallback

`func (o *O11ySlackConfig) HasFallback() bool`

HasFallback returns a boolean if a field has been set.

### GetFields

`func (o *O11ySlackConfig) GetFields() []O11ySlackField`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *O11ySlackConfig) GetFieldsOk() (*[]O11ySlackField, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *O11ySlackConfig) SetFields(v []O11ySlackField)`

SetFields sets Fields field to given value.

### HasFields

`func (o *O11ySlackConfig) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetFooter

`func (o *O11ySlackConfig) GetFooter() string`

GetFooter returns the Footer field if non-nil, zero value otherwise.

### GetFooterOk

`func (o *O11ySlackConfig) GetFooterOk() (*string, bool)`

GetFooterOk returns a tuple with the Footer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFooter

`func (o *O11ySlackConfig) SetFooter(v string)`

SetFooter sets Footer field to given value.

### HasFooter

`func (o *O11ySlackConfig) HasFooter() bool`

HasFooter returns a boolean if a field has been set.

### GetHttpConfig

`func (o *O11ySlackConfig) GetHttpConfig() O11yHTTPClientConfig`

GetHttpConfig returns the HttpConfig field if non-nil, zero value otherwise.

### GetHttpConfigOk

`func (o *O11ySlackConfig) GetHttpConfigOk() (*O11yHTTPClientConfig, bool)`

GetHttpConfigOk returns a tuple with the HttpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpConfig

`func (o *O11ySlackConfig) SetHttpConfig(v O11yHTTPClientConfig)`

SetHttpConfig sets HttpConfig field to given value.

### HasHttpConfig

`func (o *O11ySlackConfig) HasHttpConfig() bool`

HasHttpConfig returns a boolean if a field has been set.

### GetIconEmoji

`func (o *O11ySlackConfig) GetIconEmoji() string`

GetIconEmoji returns the IconEmoji field if non-nil, zero value otherwise.

### GetIconEmojiOk

`func (o *O11ySlackConfig) GetIconEmojiOk() (*string, bool)`

GetIconEmojiOk returns a tuple with the IconEmoji field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconEmoji

`func (o *O11ySlackConfig) SetIconEmoji(v string)`

SetIconEmoji sets IconEmoji field to given value.

### HasIconEmoji

`func (o *O11ySlackConfig) HasIconEmoji() bool`

HasIconEmoji returns a boolean if a field has been set.

### GetIconUrl

`func (o *O11ySlackConfig) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *O11ySlackConfig) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *O11ySlackConfig) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.

### HasIconUrl

`func (o *O11ySlackConfig) HasIconUrl() bool`

HasIconUrl returns a boolean if a field has been set.

### GetImageUrl

`func (o *O11ySlackConfig) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *O11ySlackConfig) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *O11ySlackConfig) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *O11ySlackConfig) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### GetLinkNames

`func (o *O11ySlackConfig) GetLinkNames() bool`

GetLinkNames returns the LinkNames field if non-nil, zero value otherwise.

### GetLinkNamesOk

`func (o *O11ySlackConfig) GetLinkNamesOk() (*bool, bool)`

GetLinkNamesOk returns a tuple with the LinkNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkNames

`func (o *O11ySlackConfig) SetLinkNames(v bool)`

SetLinkNames sets LinkNames field to given value.

### HasLinkNames

`func (o *O11ySlackConfig) HasLinkNames() bool`

HasLinkNames returns a boolean if a field has been set.

### GetMessageText

`func (o *O11ySlackConfig) GetMessageText() string`

GetMessageText returns the MessageText field if non-nil, zero value otherwise.

### GetMessageTextOk

`func (o *O11ySlackConfig) GetMessageTextOk() (*string, bool)`

GetMessageTextOk returns a tuple with the MessageText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageText

`func (o *O11ySlackConfig) SetMessageText(v string)`

SetMessageText sets MessageText field to given value.

### HasMessageText

`func (o *O11ySlackConfig) HasMessageText() bool`

HasMessageText returns a boolean if a field has been set.

### GetMrkdwnIn

`func (o *O11ySlackConfig) GetMrkdwnIn() []string`

GetMrkdwnIn returns the MrkdwnIn field if non-nil, zero value otherwise.

### GetMrkdwnInOk

`func (o *O11ySlackConfig) GetMrkdwnInOk() (*[]string, bool)`

GetMrkdwnInOk returns a tuple with the MrkdwnIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMrkdwnIn

`func (o *O11ySlackConfig) SetMrkdwnIn(v []string)`

SetMrkdwnIn sets MrkdwnIn field to given value.

### HasMrkdwnIn

`func (o *O11ySlackConfig) HasMrkdwnIn() bool`

HasMrkdwnIn returns a boolean if a field has been set.

### GetPretext

`func (o *O11ySlackConfig) GetPretext() string`

GetPretext returns the Pretext field if non-nil, zero value otherwise.

### GetPretextOk

`func (o *O11ySlackConfig) GetPretextOk() (*string, bool)`

GetPretextOk returns a tuple with the Pretext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPretext

`func (o *O11ySlackConfig) SetPretext(v string)`

SetPretext sets Pretext field to given value.

### HasPretext

`func (o *O11ySlackConfig) HasPretext() bool`

HasPretext returns a boolean if a field has been set.

### GetShortFields

`func (o *O11ySlackConfig) GetShortFields() bool`

GetShortFields returns the ShortFields field if non-nil, zero value otherwise.

### GetShortFieldsOk

`func (o *O11ySlackConfig) GetShortFieldsOk() (*bool, bool)`

GetShortFieldsOk returns a tuple with the ShortFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortFields

`func (o *O11ySlackConfig) SetShortFields(v bool)`

SetShortFields sets ShortFields field to given value.

### HasShortFields

`func (o *O11ySlackConfig) HasShortFields() bool`

HasShortFields returns a boolean if a field has been set.

### GetText

`func (o *O11ySlackConfig) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *O11ySlackConfig) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *O11ySlackConfig) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *O11ySlackConfig) HasText() bool`

HasText returns a boolean if a field has been set.

### GetThumbUrl

`func (o *O11ySlackConfig) GetThumbUrl() string`

GetThumbUrl returns the ThumbUrl field if non-nil, zero value otherwise.

### GetThumbUrlOk

`func (o *O11ySlackConfig) GetThumbUrlOk() (*string, bool)`

GetThumbUrlOk returns a tuple with the ThumbUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbUrl

`func (o *O11ySlackConfig) SetThumbUrl(v string)`

SetThumbUrl sets ThumbUrl field to given value.

### HasThumbUrl

`func (o *O11ySlackConfig) HasThumbUrl() bool`

HasThumbUrl returns a boolean if a field has been set.

### GetTimeout

`func (o *O11ySlackConfig) GetTimeout() int64`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *O11ySlackConfig) GetTimeoutOk() (*int64, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *O11ySlackConfig) SetTimeout(v int64)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *O11ySlackConfig) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetTitle

`func (o *O11ySlackConfig) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *O11ySlackConfig) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *O11ySlackConfig) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *O11ySlackConfig) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetTitleLink

`func (o *O11ySlackConfig) GetTitleLink() string`

GetTitleLink returns the TitleLink field if non-nil, zero value otherwise.

### GetTitleLinkOk

`func (o *O11ySlackConfig) GetTitleLinkOk() (*string, bool)`

GetTitleLinkOk returns a tuple with the TitleLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitleLink

`func (o *O11ySlackConfig) SetTitleLink(v string)`

SetTitleLink sets TitleLink field to given value.

### HasTitleLink

`func (o *O11ySlackConfig) HasTitleLink() bool`

HasTitleLink returns a boolean if a field has been set.

### GetUsername

`func (o *O11ySlackConfig) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *O11ySlackConfig) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *O11ySlackConfig) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *O11ySlackConfig) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


