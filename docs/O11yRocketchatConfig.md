# O11yRocketchatConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotifierConfig** | Pointer to [**O11yNotifierConfig**](O11yNotifierConfig.md) |  | [optional] 
**Actions** | Pointer to [**[]O11yRocketchatAttachmentAction**](O11yRocketchatAttachmentAction.md) |  | [optional] 
**ApiUrl** | Pointer to **interface{}** |  | [optional] 
**Channel** | Pointer to **string** | RocketChat channel override, (like #other-channel or @username). | [optional] 
**Color** | Pointer to **string** |  | [optional] 
**Emoji** | Pointer to **string** |  | [optional] 
**Fields** | Pointer to [**[]O11yRocketchatAttachmentField**](O11yRocketchatAttachmentField.md) |  | [optional] 
**HttpConfig** | Pointer to [**O11yHTTPClientConfig**](O11yHTTPClientConfig.md) |  | [optional] 
**IconUrl** | Pointer to **string** |  | [optional] 
**ImageUrl** | Pointer to **string** |  | [optional] 
**LinkNames** | Pointer to **bool** |  | [optional] 
**ShortFields** | Pointer to **bool** |  | [optional] 
**Text** | Pointer to **string** |  | [optional] 
**ThumbUrl** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**TitleLink** | Pointer to **string** |  | [optional] 
**Token** | Pointer to **interface{}** |  | [optional] 
**TokenFile** | Pointer to **string** |  | [optional] 
**TokenId** | Pointer to **interface{}** |  | [optional] 
**TokenIdFile** | Pointer to **string** |  | [optional] 

## Methods

### NewO11yRocketchatConfig

`func NewO11yRocketchatConfig() *O11yRocketchatConfig`

NewO11yRocketchatConfig instantiates a new O11yRocketchatConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yRocketchatConfigWithDefaults

`func NewO11yRocketchatConfigWithDefaults() *O11yRocketchatConfig`

NewO11yRocketchatConfigWithDefaults instantiates a new O11yRocketchatConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotifierConfig

`func (o *O11yRocketchatConfig) GetNotifierConfig() O11yNotifierConfig`

GetNotifierConfig returns the NotifierConfig field if non-nil, zero value otherwise.

### GetNotifierConfigOk

`func (o *O11yRocketchatConfig) GetNotifierConfigOk() (*O11yNotifierConfig, bool)`

GetNotifierConfigOk returns a tuple with the NotifierConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifierConfig

`func (o *O11yRocketchatConfig) SetNotifierConfig(v O11yNotifierConfig)`

SetNotifierConfig sets NotifierConfig field to given value.

### HasNotifierConfig

`func (o *O11yRocketchatConfig) HasNotifierConfig() bool`

HasNotifierConfig returns a boolean if a field has been set.

### GetActions

`func (o *O11yRocketchatConfig) GetActions() []O11yRocketchatAttachmentAction`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *O11yRocketchatConfig) GetActionsOk() (*[]O11yRocketchatAttachmentAction, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *O11yRocketchatConfig) SetActions(v []O11yRocketchatAttachmentAction)`

SetActions sets Actions field to given value.

### HasActions

`func (o *O11yRocketchatConfig) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetApiUrl

`func (o *O11yRocketchatConfig) GetApiUrl() interface{}`

GetApiUrl returns the ApiUrl field if non-nil, zero value otherwise.

### GetApiUrlOk

`func (o *O11yRocketchatConfig) GetApiUrlOk() (*interface{}, bool)`

GetApiUrlOk returns a tuple with the ApiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiUrl

`func (o *O11yRocketchatConfig) SetApiUrl(v interface{})`

SetApiUrl sets ApiUrl field to given value.

### HasApiUrl

`func (o *O11yRocketchatConfig) HasApiUrl() bool`

HasApiUrl returns a boolean if a field has been set.

### SetApiUrlNil

`func (o *O11yRocketchatConfig) SetApiUrlNil(b bool)`

 SetApiUrlNil sets the value for ApiUrl to be an explicit nil

### UnsetApiUrl
`func (o *O11yRocketchatConfig) UnsetApiUrl()`

UnsetApiUrl ensures that no value is present for ApiUrl, not even an explicit nil
### GetChannel

`func (o *O11yRocketchatConfig) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *O11yRocketchatConfig) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *O11yRocketchatConfig) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *O11yRocketchatConfig) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetColor

`func (o *O11yRocketchatConfig) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *O11yRocketchatConfig) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *O11yRocketchatConfig) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *O11yRocketchatConfig) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetEmoji

`func (o *O11yRocketchatConfig) GetEmoji() string`

GetEmoji returns the Emoji field if non-nil, zero value otherwise.

### GetEmojiOk

`func (o *O11yRocketchatConfig) GetEmojiOk() (*string, bool)`

GetEmojiOk returns a tuple with the Emoji field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmoji

`func (o *O11yRocketchatConfig) SetEmoji(v string)`

SetEmoji sets Emoji field to given value.

### HasEmoji

`func (o *O11yRocketchatConfig) HasEmoji() bool`

HasEmoji returns a boolean if a field has been set.

### GetFields

`func (o *O11yRocketchatConfig) GetFields() []O11yRocketchatAttachmentField`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *O11yRocketchatConfig) GetFieldsOk() (*[]O11yRocketchatAttachmentField, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *O11yRocketchatConfig) SetFields(v []O11yRocketchatAttachmentField)`

SetFields sets Fields field to given value.

### HasFields

`func (o *O11yRocketchatConfig) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetHttpConfig

`func (o *O11yRocketchatConfig) GetHttpConfig() O11yHTTPClientConfig`

GetHttpConfig returns the HttpConfig field if non-nil, zero value otherwise.

### GetHttpConfigOk

`func (o *O11yRocketchatConfig) GetHttpConfigOk() (*O11yHTTPClientConfig, bool)`

GetHttpConfigOk returns a tuple with the HttpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpConfig

`func (o *O11yRocketchatConfig) SetHttpConfig(v O11yHTTPClientConfig)`

SetHttpConfig sets HttpConfig field to given value.

### HasHttpConfig

`func (o *O11yRocketchatConfig) HasHttpConfig() bool`

HasHttpConfig returns a boolean if a field has been set.

### GetIconUrl

`func (o *O11yRocketchatConfig) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *O11yRocketchatConfig) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *O11yRocketchatConfig) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.

### HasIconUrl

`func (o *O11yRocketchatConfig) HasIconUrl() bool`

HasIconUrl returns a boolean if a field has been set.

### GetImageUrl

`func (o *O11yRocketchatConfig) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *O11yRocketchatConfig) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *O11yRocketchatConfig) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *O11yRocketchatConfig) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### GetLinkNames

`func (o *O11yRocketchatConfig) GetLinkNames() bool`

GetLinkNames returns the LinkNames field if non-nil, zero value otherwise.

### GetLinkNamesOk

`func (o *O11yRocketchatConfig) GetLinkNamesOk() (*bool, bool)`

GetLinkNamesOk returns a tuple with the LinkNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkNames

`func (o *O11yRocketchatConfig) SetLinkNames(v bool)`

SetLinkNames sets LinkNames field to given value.

### HasLinkNames

`func (o *O11yRocketchatConfig) HasLinkNames() bool`

HasLinkNames returns a boolean if a field has been set.

### GetShortFields

`func (o *O11yRocketchatConfig) GetShortFields() bool`

GetShortFields returns the ShortFields field if non-nil, zero value otherwise.

### GetShortFieldsOk

`func (o *O11yRocketchatConfig) GetShortFieldsOk() (*bool, bool)`

GetShortFieldsOk returns a tuple with the ShortFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortFields

`func (o *O11yRocketchatConfig) SetShortFields(v bool)`

SetShortFields sets ShortFields field to given value.

### HasShortFields

`func (o *O11yRocketchatConfig) HasShortFields() bool`

HasShortFields returns a boolean if a field has been set.

### GetText

`func (o *O11yRocketchatConfig) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *O11yRocketchatConfig) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *O11yRocketchatConfig) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *O11yRocketchatConfig) HasText() bool`

HasText returns a boolean if a field has been set.

### GetThumbUrl

`func (o *O11yRocketchatConfig) GetThumbUrl() string`

GetThumbUrl returns the ThumbUrl field if non-nil, zero value otherwise.

### GetThumbUrlOk

`func (o *O11yRocketchatConfig) GetThumbUrlOk() (*string, bool)`

GetThumbUrlOk returns a tuple with the ThumbUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbUrl

`func (o *O11yRocketchatConfig) SetThumbUrl(v string)`

SetThumbUrl sets ThumbUrl field to given value.

### HasThumbUrl

`func (o *O11yRocketchatConfig) HasThumbUrl() bool`

HasThumbUrl returns a boolean if a field has been set.

### GetTitle

`func (o *O11yRocketchatConfig) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *O11yRocketchatConfig) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *O11yRocketchatConfig) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *O11yRocketchatConfig) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetTitleLink

`func (o *O11yRocketchatConfig) GetTitleLink() string`

GetTitleLink returns the TitleLink field if non-nil, zero value otherwise.

### GetTitleLinkOk

`func (o *O11yRocketchatConfig) GetTitleLinkOk() (*string, bool)`

GetTitleLinkOk returns a tuple with the TitleLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitleLink

`func (o *O11yRocketchatConfig) SetTitleLink(v string)`

SetTitleLink sets TitleLink field to given value.

### HasTitleLink

`func (o *O11yRocketchatConfig) HasTitleLink() bool`

HasTitleLink returns a boolean if a field has been set.

### GetToken

`func (o *O11yRocketchatConfig) GetToken() interface{}`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *O11yRocketchatConfig) GetTokenOk() (*interface{}, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *O11yRocketchatConfig) SetToken(v interface{})`

SetToken sets Token field to given value.

### HasToken

`func (o *O11yRocketchatConfig) HasToken() bool`

HasToken returns a boolean if a field has been set.

### SetTokenNil

`func (o *O11yRocketchatConfig) SetTokenNil(b bool)`

 SetTokenNil sets the value for Token to be an explicit nil

### UnsetToken
`func (o *O11yRocketchatConfig) UnsetToken()`

UnsetToken ensures that no value is present for Token, not even an explicit nil
### GetTokenFile

`func (o *O11yRocketchatConfig) GetTokenFile() string`

GetTokenFile returns the TokenFile field if non-nil, zero value otherwise.

### GetTokenFileOk

`func (o *O11yRocketchatConfig) GetTokenFileOk() (*string, bool)`

GetTokenFileOk returns a tuple with the TokenFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenFile

`func (o *O11yRocketchatConfig) SetTokenFile(v string)`

SetTokenFile sets TokenFile field to given value.

### HasTokenFile

`func (o *O11yRocketchatConfig) HasTokenFile() bool`

HasTokenFile returns a boolean if a field has been set.

### GetTokenId

`func (o *O11yRocketchatConfig) GetTokenId() interface{}`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *O11yRocketchatConfig) GetTokenIdOk() (*interface{}, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *O11yRocketchatConfig) SetTokenId(v interface{})`

SetTokenId sets TokenId field to given value.

### HasTokenId

`func (o *O11yRocketchatConfig) HasTokenId() bool`

HasTokenId returns a boolean if a field has been set.

### SetTokenIdNil

`func (o *O11yRocketchatConfig) SetTokenIdNil(b bool)`

 SetTokenIdNil sets the value for TokenId to be an explicit nil

### UnsetTokenId
`func (o *O11yRocketchatConfig) UnsetTokenId()`

UnsetTokenId ensures that no value is present for TokenId, not even an explicit nil
### GetTokenIdFile

`func (o *O11yRocketchatConfig) GetTokenIdFile() string`

GetTokenIdFile returns the TokenIdFile field if non-nil, zero value otherwise.

### GetTokenIdFileOk

`func (o *O11yRocketchatConfig) GetTokenIdFileOk() (*string, bool)`

GetTokenIdFileOk returns a tuple with the TokenIdFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenIdFile

`func (o *O11yRocketchatConfig) SetTokenIdFile(v string)`

SetTokenIdFile sets TokenIdFile field to given value.

### HasTokenIdFile

`func (o *O11yRocketchatConfig) HasTokenIdFile() bool`

HasTokenIdFile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


