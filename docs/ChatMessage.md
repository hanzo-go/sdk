# ChatMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessageId** | Pointer to **string** |  | [optional] 
**ConversationId** | Pointer to **string** |  | [optional] 
**ParentMessageId** | Pointer to **string** |  | [optional] 
**Text** | Pointer to **string** |  | [optional] 
**Sender** | Pointer to **string** |  | [optional] 
**IsCreatedByUser** | Pointer to **bool** |  | [optional] 
**Model** | Pointer to **string** |  | [optional] 
**Endpoint** | Pointer to **string** |  | [optional] 
**Content** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Unfinished** | Pointer to **bool** |  | [optional] 
**Error** | Pointer to **bool** |  | [optional] 
**IconURL** | Pointer to **string** |  | [optional] 
**Feedback** | Pointer to **map[string]interface{}** |  | [optional] 
**TokenCount** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewChatMessage

`func NewChatMessage() *ChatMessage`

NewChatMessage instantiates a new ChatMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatMessageWithDefaults

`func NewChatMessageWithDefaults() *ChatMessage`

NewChatMessageWithDefaults instantiates a new ChatMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessageId

`func (o *ChatMessage) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *ChatMessage) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *ChatMessage) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *ChatMessage) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.

### GetConversationId

`func (o *ChatMessage) GetConversationId() string`

GetConversationId returns the ConversationId field if non-nil, zero value otherwise.

### GetConversationIdOk

`func (o *ChatMessage) GetConversationIdOk() (*string, bool)`

GetConversationIdOk returns a tuple with the ConversationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversationId

`func (o *ChatMessage) SetConversationId(v string)`

SetConversationId sets ConversationId field to given value.

### HasConversationId

`func (o *ChatMessage) HasConversationId() bool`

HasConversationId returns a boolean if a field has been set.

### GetParentMessageId

`func (o *ChatMessage) GetParentMessageId() string`

GetParentMessageId returns the ParentMessageId field if non-nil, zero value otherwise.

### GetParentMessageIdOk

`func (o *ChatMessage) GetParentMessageIdOk() (*string, bool)`

GetParentMessageIdOk returns a tuple with the ParentMessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentMessageId

`func (o *ChatMessage) SetParentMessageId(v string)`

SetParentMessageId sets ParentMessageId field to given value.

### HasParentMessageId

`func (o *ChatMessage) HasParentMessageId() bool`

HasParentMessageId returns a boolean if a field has been set.

### GetText

`func (o *ChatMessage) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *ChatMessage) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *ChatMessage) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *ChatMessage) HasText() bool`

HasText returns a boolean if a field has been set.

### GetSender

`func (o *ChatMessage) GetSender() string`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *ChatMessage) GetSenderOk() (*string, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *ChatMessage) SetSender(v string)`

SetSender sets Sender field to given value.

### HasSender

`func (o *ChatMessage) HasSender() bool`

HasSender returns a boolean if a field has been set.

### GetIsCreatedByUser

`func (o *ChatMessage) GetIsCreatedByUser() bool`

GetIsCreatedByUser returns the IsCreatedByUser field if non-nil, zero value otherwise.

### GetIsCreatedByUserOk

`func (o *ChatMessage) GetIsCreatedByUserOk() (*bool, bool)`

GetIsCreatedByUserOk returns a tuple with the IsCreatedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCreatedByUser

`func (o *ChatMessage) SetIsCreatedByUser(v bool)`

SetIsCreatedByUser sets IsCreatedByUser field to given value.

### HasIsCreatedByUser

`func (o *ChatMessage) HasIsCreatedByUser() bool`

HasIsCreatedByUser returns a boolean if a field has been set.

### GetModel

`func (o *ChatMessage) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ChatMessage) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ChatMessage) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *ChatMessage) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetEndpoint

`func (o *ChatMessage) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *ChatMessage) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *ChatMessage) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *ChatMessage) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetContent

`func (o *ChatMessage) GetContent() []map[string]interface{}`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ChatMessage) GetContentOk() (*[]map[string]interface{}, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ChatMessage) SetContent(v []map[string]interface{})`

SetContent sets Content field to given value.

### HasContent

`func (o *ChatMessage) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetUnfinished

`func (o *ChatMessage) GetUnfinished() bool`

GetUnfinished returns the Unfinished field if non-nil, zero value otherwise.

### GetUnfinishedOk

`func (o *ChatMessage) GetUnfinishedOk() (*bool, bool)`

GetUnfinishedOk returns a tuple with the Unfinished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnfinished

`func (o *ChatMessage) SetUnfinished(v bool)`

SetUnfinished sets Unfinished field to given value.

### HasUnfinished

`func (o *ChatMessage) HasUnfinished() bool`

HasUnfinished returns a boolean if a field has been set.

### GetError

`func (o *ChatMessage) GetError() bool`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ChatMessage) GetErrorOk() (*bool, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ChatMessage) SetError(v bool)`

SetError sets Error field to given value.

### HasError

`func (o *ChatMessage) HasError() bool`

HasError returns a boolean if a field has been set.

### GetIconURL

`func (o *ChatMessage) GetIconURL() string`

GetIconURL returns the IconURL field if non-nil, zero value otherwise.

### GetIconURLOk

`func (o *ChatMessage) GetIconURLOk() (*string, bool)`

GetIconURLOk returns a tuple with the IconURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconURL

`func (o *ChatMessage) SetIconURL(v string)`

SetIconURL sets IconURL field to given value.

### HasIconURL

`func (o *ChatMessage) HasIconURL() bool`

HasIconURL returns a boolean if a field has been set.

### GetFeedback

`func (o *ChatMessage) GetFeedback() map[string]interface{}`

GetFeedback returns the Feedback field if non-nil, zero value otherwise.

### GetFeedbackOk

`func (o *ChatMessage) GetFeedbackOk() (*map[string]interface{}, bool)`

GetFeedbackOk returns a tuple with the Feedback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedback

`func (o *ChatMessage) SetFeedback(v map[string]interface{})`

SetFeedback sets Feedback field to given value.

### HasFeedback

`func (o *ChatMessage) HasFeedback() bool`

HasFeedback returns a boolean if a field has been set.

### GetTokenCount

`func (o *ChatMessage) GetTokenCount() int32`

GetTokenCount returns the TokenCount field if non-nil, zero value otherwise.

### GetTokenCountOk

`func (o *ChatMessage) GetTokenCountOk() (*int32, bool)`

GetTokenCountOk returns a tuple with the TokenCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenCount

`func (o *ChatMessage) SetTokenCount(v int32)`

SetTokenCount sets TokenCount field to given value.

### HasTokenCount

`func (o *ChatMessage) HasTokenCount() bool`

HasTokenCount returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ChatMessage) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ChatMessage) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ChatMessage) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ChatMessage) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


