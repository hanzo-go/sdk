# ChatPostAgentsV1ChatCompletionsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Model** | **string** | Agent ID to use | 
**Messages** | [**[]ChatChatMessage**](ChatChatMessage.md) |  | 
**Stream** | Pointer to **bool** |  | [optional] [default to false]
**ConversationId** | Pointer to **string** |  | [optional] 
**ParentMessageId** | Pointer to **string** |  | [optional] 

## Methods

### NewChatPostAgentsV1ChatCompletionsRequest

`func NewChatPostAgentsV1ChatCompletionsRequest(model string, messages []ChatChatMessage, ) *ChatPostAgentsV1ChatCompletionsRequest`

NewChatPostAgentsV1ChatCompletionsRequest instantiates a new ChatPostAgentsV1ChatCompletionsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatPostAgentsV1ChatCompletionsRequestWithDefaults

`func NewChatPostAgentsV1ChatCompletionsRequestWithDefaults() *ChatPostAgentsV1ChatCompletionsRequest`

NewChatPostAgentsV1ChatCompletionsRequestWithDefaults instantiates a new ChatPostAgentsV1ChatCompletionsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModel

`func (o *ChatPostAgentsV1ChatCompletionsRequest) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ChatPostAgentsV1ChatCompletionsRequest) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ChatPostAgentsV1ChatCompletionsRequest) SetModel(v string)`

SetModel sets Model field to given value.


### GetMessages

`func (o *ChatPostAgentsV1ChatCompletionsRequest) GetMessages() []ChatChatMessage`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *ChatPostAgentsV1ChatCompletionsRequest) GetMessagesOk() (*[]ChatChatMessage, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *ChatPostAgentsV1ChatCompletionsRequest) SetMessages(v []ChatChatMessage)`

SetMessages sets Messages field to given value.


### GetStream

`func (o *ChatPostAgentsV1ChatCompletionsRequest) GetStream() bool`

GetStream returns the Stream field if non-nil, zero value otherwise.

### GetStreamOk

`func (o *ChatPostAgentsV1ChatCompletionsRequest) GetStreamOk() (*bool, bool)`

GetStreamOk returns a tuple with the Stream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStream

`func (o *ChatPostAgentsV1ChatCompletionsRequest) SetStream(v bool)`

SetStream sets Stream field to given value.

### HasStream

`func (o *ChatPostAgentsV1ChatCompletionsRequest) HasStream() bool`

HasStream returns a boolean if a field has been set.

### GetConversationId

`func (o *ChatPostAgentsV1ChatCompletionsRequest) GetConversationId() string`

GetConversationId returns the ConversationId field if non-nil, zero value otherwise.

### GetConversationIdOk

`func (o *ChatPostAgentsV1ChatCompletionsRequest) GetConversationIdOk() (*string, bool)`

GetConversationIdOk returns a tuple with the ConversationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversationId

`func (o *ChatPostAgentsV1ChatCompletionsRequest) SetConversationId(v string)`

SetConversationId sets ConversationId field to given value.

### HasConversationId

`func (o *ChatPostAgentsV1ChatCompletionsRequest) HasConversationId() bool`

HasConversationId returns a boolean if a field has been set.

### GetParentMessageId

`func (o *ChatPostAgentsV1ChatCompletionsRequest) GetParentMessageId() string`

GetParentMessageId returns the ParentMessageId field if non-nil, zero value otherwise.

### GetParentMessageIdOk

`func (o *ChatPostAgentsV1ChatCompletionsRequest) GetParentMessageIdOk() (*string, bool)`

GetParentMessageIdOk returns a tuple with the ParentMessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentMessageId

`func (o *ChatPostAgentsV1ChatCompletionsRequest) SetParentMessageId(v string)`

SetParentMessageId sets ParentMessageId field to given value.

### HasParentMessageId

`func (o *ChatPostAgentsV1ChatCompletionsRequest) HasParentMessageId() bool`

HasParentMessageId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


