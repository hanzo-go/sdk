# ProviderSlackMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BotId** | Pointer to **string** | BotID identifies a bot author when present. | [optional] 
**Files** | Pointer to [**[]ProviderSlackFile**](ProviderSlackFile.md) | Files are the files attached to this message. Read one&#39;s bytes with GET /v1/provider/slack/file?id&#x3D;&lt;id&gt;; nothing here is the content. | [optional] 
**ReplyCount** | Pointer to **int64** | ReplyCount indicates whether a parent has replies to fetch separately. | [optional] 
**Subtype** | Pointer to **string** | Subtype distinguishes messages from channel events. | [optional] 
**Text** | Pointer to **string** | Text is the message in Slack&#39;s formatting. | [optional] 
**ThreadTs** | Pointer to **string** | ThreadTS is the parent timestamp for a threaded message. | [optional] 
**Ts** | Pointer to **string** | TS is Slack&#39;s exact message timestamp and ID; retain it as a string. | [optional] 
**User** | Pointer to **string** | User is the author&#39;s Slack user ID. | [optional] 

## Methods

### NewProviderSlackMessage

`func NewProviderSlackMessage() *ProviderSlackMessage`

NewProviderSlackMessage instantiates a new ProviderSlackMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderSlackMessageWithDefaults

`func NewProviderSlackMessageWithDefaults() *ProviderSlackMessage`

NewProviderSlackMessageWithDefaults instantiates a new ProviderSlackMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBotId

`func (o *ProviderSlackMessage) GetBotId() string`

GetBotId returns the BotId field if non-nil, zero value otherwise.

### GetBotIdOk

`func (o *ProviderSlackMessage) GetBotIdOk() (*string, bool)`

GetBotIdOk returns a tuple with the BotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBotId

`func (o *ProviderSlackMessage) SetBotId(v string)`

SetBotId sets BotId field to given value.

### HasBotId

`func (o *ProviderSlackMessage) HasBotId() bool`

HasBotId returns a boolean if a field has been set.

### GetFiles

`func (o *ProviderSlackMessage) GetFiles() []ProviderSlackFile`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *ProviderSlackMessage) GetFilesOk() (*[]ProviderSlackFile, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *ProviderSlackMessage) SetFiles(v []ProviderSlackFile)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *ProviderSlackMessage) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetReplyCount

`func (o *ProviderSlackMessage) GetReplyCount() int64`

GetReplyCount returns the ReplyCount field if non-nil, zero value otherwise.

### GetReplyCountOk

`func (o *ProviderSlackMessage) GetReplyCountOk() (*int64, bool)`

GetReplyCountOk returns a tuple with the ReplyCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyCount

`func (o *ProviderSlackMessage) SetReplyCount(v int64)`

SetReplyCount sets ReplyCount field to given value.

### HasReplyCount

`func (o *ProviderSlackMessage) HasReplyCount() bool`

HasReplyCount returns a boolean if a field has been set.

### GetSubtype

`func (o *ProviderSlackMessage) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *ProviderSlackMessage) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *ProviderSlackMessage) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.

### HasSubtype

`func (o *ProviderSlackMessage) HasSubtype() bool`

HasSubtype returns a boolean if a field has been set.

### GetText

`func (o *ProviderSlackMessage) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *ProviderSlackMessage) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *ProviderSlackMessage) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *ProviderSlackMessage) HasText() bool`

HasText returns a boolean if a field has been set.

### GetThreadTs

`func (o *ProviderSlackMessage) GetThreadTs() string`

GetThreadTs returns the ThreadTs field if non-nil, zero value otherwise.

### GetThreadTsOk

`func (o *ProviderSlackMessage) GetThreadTsOk() (*string, bool)`

GetThreadTsOk returns a tuple with the ThreadTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadTs

`func (o *ProviderSlackMessage) SetThreadTs(v string)`

SetThreadTs sets ThreadTs field to given value.

### HasThreadTs

`func (o *ProviderSlackMessage) HasThreadTs() bool`

HasThreadTs returns a boolean if a field has been set.

### GetTs

`func (o *ProviderSlackMessage) GetTs() string`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *ProviderSlackMessage) GetTsOk() (*string, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *ProviderSlackMessage) SetTs(v string)`

SetTs sets Ts field to given value.

### HasTs

`func (o *ProviderSlackMessage) HasTs() bool`

HasTs returns a boolean if a field has been set.

### GetUser

`func (o *ProviderSlackMessage) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ProviderSlackMessage) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ProviderSlackMessage) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *ProviderSlackMessage) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


